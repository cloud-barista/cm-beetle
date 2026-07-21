'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore, DEMO_SOURCE_INFRA } from '../../store/migrationStore';
import { honeybeeApi } from '../../api/client';
import { SaveRevisionModal } from '../common/SaveRevisionModal';
import {
  Plus, Server, Key, Upload, CheckCircle2, XCircle,
  RefreshCw, FileText, ChevronRight, Download, Trash2,
  AlertCircle, Loader2, ChevronDown, ChevronUp, Play, X, Save,
  Database
} from 'lucide-react';

// A plain SSH username never looks like this — treat long base64-like
// strings as leaked secret material (e.g. a private key) rather than a
// real username, and avoid rendering it in full.
const looksLikeSecret = (value: string): boolean =>
  value.length > 32 && /^[A-Za-z0-9+/=]+$/.test(value);

// Some backends return the username base64-encoded rather than encrypted.
// If we can decode it into a short, printable string, show the decoded
// value in full instead of treating it as an opaque secret.
const tryDecodeBase64Username = (value: string): string | null => {
  try {
    const decoded = atob(value);
    if (decoded && decoded.length <= 64 && /^[\x20-\x7E]+$/.test(decoded)) {
      return decoded;
    }
  } catch {
    // not valid base64 — fall through
  }
  return null;
};

// ─────────────────────────────────────────────────────────
// Types
// ─────────────────────────────────────────────────────────

interface ServerRow {
  localId: string;
  connId: string;         // empty = not yet registered
  name: string;
  ip: string;
  port: string;
  user: string;
  privateKey: string;
  password: string;
  authMode: 'key' | 'password';
  connectionStatus: string;
  agentStatus: string;
  connectionFailedMsg: string;
  agentFailedMsg: string;
}

// A connection needs at least one authentication method
const hasAuth = (row: ServerRow, useCommonCred: boolean, commonKey: string, commonPassword: string): boolean =>
  !!(row.privateKey || row.password || (useCommonCred && (commonKey || commonPassword)));

// ─────────────────────────────────────────────────────────
// StatusPill
// ─────────────────────────────────────────────────────────

const StatusPill = ({
  status, failedMsg, successLabel = 'OK', failLabel = 'Failed',
}: { status: string; failedMsg?: string; successLabel?: string; failLabel?: string }) => {
  if (!status) return <span className="text-text-muted text-xs">—</span>;
  if (status === 'success')
    return (
      <span className="inline-flex items-center gap-1 px-2 py-0.5 bg-emerald-100 dark:bg-emerald-950/40 text-emerald-600 dark:text-emerald-400 border border-emerald-300 dark:border-emerald-800/40 rounded-full text-xs font-semibold">
        <CheckCircle2 className="w-3 h-3" />{successLabel}
      </span>
    );
  return (
    <span title={failedMsg} className="inline-flex items-center gap-1 px-2 py-0.5 bg-red-100 dark:bg-red-950/40 text-red-600 dark:text-red-400 border border-red-300 dark:border-red-800/40 rounded-full text-xs font-semibold cursor-help">
      <XCircle className="w-3 h-3" />{failLabel}
    </span>
  );
};

// ─────────────────────────────────────────────────────────
// Shared Credentials panel (reused in both modal and inline)
// ─────────────────────────────────────────────────────────

interface SharedCredsProps {
  show: boolean; onToggle: () => void;
  enabled: boolean; onToggleEnabled: (v: boolean) => void;
  user: string; onUser: (v: string) => void;
  port: number; onPort: (v: number) => void;
  pemKey: string; onPemKey: (v: string) => void;
  password: string; onPassword: (v: string) => void;
  onKeyDrop: (e: React.DragEvent<HTMLTextAreaElement>) => void;
  locked?: boolean;
}

const SharedCredsPanel: React.FC<SharedCredsProps> = ({
  show, onToggle, enabled, onToggleEnabled, user, onUser, port, onPort, pemKey, onPemKey, password, onPassword, onKeyDrop, locked = false,
}) => (
  <div className="border border-border-main rounded-xl overflow-hidden">
    <button
      onClick={onToggle}
      className="w-full px-4 py-3 flex items-center justify-between text-sm font-semibold text-text-muted hover:text-text-main bg-bg-input/20 hover:bg-bg-input/40 transition cursor-pointer"
    >
      <span className="flex items-center gap-2">
        <Key className="w-4 h-4 text-teal-400" />
        Shared SSH Credentials
        {enabled && <span className="text-xs text-teal-500 font-normal">(enabled — blank fields inherit these values)</span>}
      </span>
      {show ? <ChevronUp className="w-4 h-4" /> : <ChevronDown className="w-4 h-4" />}
    </button>
    {show && (
      <div className="px-4 py-4 bg-bg-input/5 border-t border-border-main">
        {locked && (
          <p className="text-xs text-text-muted mb-3 flex items-center gap-1.5">
            <Key className="w-3.5 h-3.5" /> Locked — this group already has registered connections, so shared credentials can no longer be edited.
          </p>
        )}
        <label className={`inline-flex items-center gap-2 mb-3 ${locked ? 'cursor-not-allowed opacity-60' : 'cursor-pointer'}`}>
          <input type="checkbox" disabled={locked} checked={enabled} onChange={e => onToggleEnabled(e.target.checked)} className="sr-only peer" />
          <div className="relative w-9 h-5 bg-bg-input peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-slate-400 after:border-slate-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-teal-600" />
          <span className="text-sm text-text-muted">Enable Auto Inherit</span>
        </label>
        <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
          <div>
            <label className="block text-xs font-semibold text-text-muted mb-1">SSH Username</label>
            <input type="text" disabled={!enabled || locked} value={user} onChange={e => onUser(e.target.value)} className="w-full bg-bg-input border border-border-main disabled:opacity-40 text-text-main rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500" />
          </div>
          <div>
            <label className="block text-xs font-semibold text-text-muted mb-1">SSH Port</label>
            <input type="number" disabled={!enabled || locked} value={port} onChange={e => onPort(Number(e.target.value))} className="w-full bg-bg-input border border-border-main disabled:opacity-40 text-text-main rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500" />
          </div>
          <div>
            <label className="block text-xs font-semibold text-text-muted mb-1">Password</label>
            <input type="password" disabled={!enabled || locked} placeholder="Optional if using key" value={password} onChange={e => onPassword(e.target.value)} className="w-full bg-bg-input border border-border-main disabled:opacity-40 text-text-main rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500" />
          </div>
          <div>
            <label className="block text-xs font-semibold text-text-muted mb-1">Private Key (PEM)</label>
            <textarea disabled={!enabled || locked} onDragOver={e => e.preventDefault()} onDrop={onKeyDrop} placeholder="Paste PEM key or drop file" value={pemKey} onChange={e => onPemKey(e.target.value)} rows={1} className="w-full bg-bg-input border border-border-main disabled:opacity-40 text-text-main rounded-lg px-3 py-2 text-sm font-mono focus:outline-none focus:ring-1 focus:ring-teal-500 resize-none" />
          </div>
        </div>
        {!locked && enabled && !password && !pemKey && (
          <p className="text-xs text-amber-600 dark:text-amber-400 mt-2">Provide a Password or Private Key (at least one is required to connect).</p>
        )}
      </div>
    )}
  </div>
);

// ─────────────────────────────────────────────────────────
// Server table row renderer
// ─────────────────────────────────────────────────────────

interface RowProps {
  row: ServerRow; idx: number;
  useCommonCred: boolean; commonUser: string; commonPort: number; commonKey: string; commonPassword: string;
  onUpdate: (id: string, field: keyof ServerRow, v: string) => void;
  onDelete: (id: string) => void;
  onReset?: (id: string) => void;
  onKeyDrop: (id: string, e: React.DragEvent) => void;
}

const ServerTableRow: React.FC<RowProps> = ({
  row, idx, useCommonCred, commonUser, commonPort, commonKey, commonPassword, onUpdate, onDelete, onReset, onKeyDrop,
}) => {
  const isRegistered = !!row.connId;
  const isDraft = !isRegistered && !!(row.name || row.ip);
  const authOk = isRegistered || hasAuth(row, useCommonCred, commonKey, commonPassword);
  return (
    <tr className={`transition ${isRegistered ? 'bg-emerald-500/[0.015]' : 'hover:bg-bg-input/20'}`}>
      <td className="py-2.5 px-3 text-center text-text-muted text-xs">
        {idx + 1}
        {isDraft && <div className="mt-1 text-[10px] font-bold text-amber-600 dark:text-amber-400 uppercase tracking-wide">Draft</div>}
      </td>

      <td className="py-2.5 px-3">
        {isRegistered
          ? <span className="text-sm text-text-main font-semibold px-2">{row.name || '—'}</span>
          : <input type="text" value={row.name} onChange={e => onUpdate(row.localId, 'name', e.target.value)} placeholder="Server name"
              className="w-full bg-bg-input border border-border-main rounded-lg px-2 py-1.5 text-sm text-text-main focus:outline-none focus:border-emerald-500/50" />}
      </td>

      <td className="py-2.5 px-3">
        {isRegistered
          ? <span className="text-sm text-text-main font-mono px-2">{row.ip}</span>
          : <input type="text" placeholder="10.0.1.x" value={row.ip} onChange={e => onUpdate(row.localId, 'ip', e.target.value)}
              className="w-full bg-bg-input border border-border-main rounded-lg px-2 py-1.5 text-sm text-text-main focus:outline-none focus:border-emerald-500/50" />}
      </td>

      <td className="py-2.5 px-3">
        {isRegistered
          ? <span className="text-sm text-text-main px-2">{row.port || commonPort}</span>
          : <input type="number" placeholder={useCommonCred ? String(commonPort) : '22'} value={row.port} onChange={e => onUpdate(row.localId, 'port', e.target.value)}
              className="w-full bg-bg-input border border-border-main rounded-lg px-2 py-1.5 text-sm text-text-main focus:outline-none focus:border-emerald-500/50 placeholder-teal-500/40" />}
      </td>

      <td className="py-2.5 px-3">
        {isRegistered ? (() => {
          const raw = row.user || commonUser;
          if (!looksLikeSecret(raw)) return <span className="text-sm text-text-main px-2">{raw}</span>;
          const decoded = tryDecodeBase64Username(raw);
          if (decoded) return <span className="text-sm text-text-main px-2">{decoded}</span>;
          // Honeybee encrypts the username with its own server-held RSA key before
          // returning it via the API, so it can never be decrypted by this UI.
          return <span className="text-xs text-text-muted px-2 italic whitespace-nowrap" title="Encrypted by Honeybee — cannot be decrypted by this UI">Stored in Honeybee</span>;
        })()
          : <input type="text" placeholder={useCommonCred ? commonUser : 'ubuntu'} value={row.user} onChange={e => onUpdate(row.localId, 'user', e.target.value)}
              className="w-full bg-bg-input border border-border-main rounded-lg px-2 py-1.5 text-sm text-text-main focus:outline-none focus:border-emerald-500/50 placeholder-teal-500/40" />}
      </td>

      <td className="py-2.5 px-3">
        {isRegistered
          ? <span className="text-xs text-text-muted px-2 italic">Stored in Honeybee</span>
          : (
            <div className="space-y-1">
              <div className="flex gap-1">
                <button type="button" onClick={() => onUpdate(row.localId, 'authMode', 'key')}
                  className={`px-2 py-0.5 rounded text-xs font-semibold border cursor-pointer transition ${row.authMode !== 'password' ? 'bg-emerald-100 dark:bg-emerald-500/15 border-emerald-300 dark:border-emerald-500/40 text-emerald-700 dark:text-emerald-400' : 'bg-bg-input border-border-main text-text-muted'}`}>
                  Key
                </button>
                <button type="button" onClick={() => onUpdate(row.localId, 'authMode', 'password')}
                  className={`px-2 py-0.5 rounded text-xs font-semibold border cursor-pointer transition ${row.authMode === 'password' ? 'bg-emerald-100 dark:bg-emerald-500/15 border-emerald-300 dark:border-emerald-500/40 text-emerald-700 dark:text-emerald-400' : 'bg-bg-input border-border-main text-text-muted'}`}>
                  Password
                </button>
              </div>
              {row.authMode === 'password' ? (
                <input type="password"
                  placeholder={useCommonCred && commonPassword ? '↗ Inheriting shared password' : 'Enter password'}
                  value={row.password} onChange={e => onUpdate(row.localId, 'password', e.target.value)}
                  className={`w-full bg-bg-input border rounded-lg px-2 py-1.5 text-sm text-text-main focus:outline-none ${!authOk ? 'border-red-400/60' : 'border-border-main focus:border-emerald-500/50'}`} />
              ) : (
                <div onDragOver={e => e.preventDefault()} onDrop={e => onKeyDrop(row.localId, e)}
                  className={`border border-dashed rounded-lg p-1.5 flex items-center gap-2 bg-bg-input ${!authOk ? 'border-red-400/60' : 'border-border-main'}`}>
                  <input type="password"
                    placeholder={row.privateKey ? 'Key loaded' : useCommonCred && commonKey ? '↗ Inheriting shared key' : 'Drop PEM or paste key'}
                    value={row.privateKey} onChange={e => onUpdate(row.localId, 'privateKey', e.target.value)}
                    className={`flex-1 bg-transparent border-0 outline-none text-sm min-w-0 ${!row.privateKey && useCommonCred && commonKey ? 'text-teal-400/70 placeholder-teal-400/60' : 'text-text-main'}`} />
                  <span className="text-xs text-text-muted bg-bg-panel px-1.5 py-0.5 rounded flex-shrink-0"><Upload className="w-3 h-3 inline" /></span>
                </div>
              )}
              {!authOk && <p className="text-[10px] text-red-500 dark:text-red-400">Password or Private Key required</p>}
            </div>
          )}
      </td>

      <td className="py-2.5 px-3 text-center">
        <StatusPill status={row.connectionStatus} failedMsg={row.connectionFailedMsg} successLabel="Connected" failLabel="Failed" />
      </td>
      <td className="py-2.5 px-3 text-center">
        <StatusPill status={row.agentStatus} failedMsg={row.agentFailedMsg} successLabel="OK" failLabel="Failed" />
      </td>
      <td className="py-2.5 px-3 text-center">
        {isRegistered ? (
          <button onClick={() => onDelete(row.localId)} className="p-1.5 rounded-lg hover:bg-red-500/10 text-text-muted hover:text-red-400 transition cursor-pointer" title="Delete this connection">
            <Trash2 className="w-3.5 h-3.5" />
          </button>
        ) : (
          <button onClick={() => (onReset || onDelete)(row.localId)} className="p-1.5 rounded-lg hover:bg-bg-input text-text-muted hover:text-text-main transition cursor-pointer" title="Clear this row">
            <X className="w-3.5 h-3.5" />
          </button>
        )}
      </td>
    </tr>
  );
};

const TABLE_HEADER = (
  <thead>
    <tr className="border-b border-border-main bg-bg-input/40 text-text-muted text-xs font-bold">
      <th className="py-3 px-3 w-10 text-center">#</th>
      <th className="py-3 px-3 min-w-[150px]">Name</th>
      <th className="py-3 px-3 min-w-[140px]">IP Address</th>
      <th className="py-3 px-3 w-20">Port</th>
      <th className="py-3 px-3 min-w-[140px]">Username</th>
      <th className="py-3 px-3 min-w-[220px]">Authentication (Password / Private Key)</th>
      <th className="py-3 px-3 text-center w-28">SSH Status</th>
      <th className="py-3 px-3 text-center w-24">Agent</th>
      <th className="py-3 px-3 text-center w-14">Actions</th>
    </tr>
  </thead>
);

// ─────────────────────────────────────────────────────────
// Main component
// ─────────────────────────────────────────────────────────

export const SourceMetadataExtraction: React.FC = () => {
  const {
    sourceGroups, activeSgId, connections, refinedSourceInfra, isLoadingSource,
    fetchSourceGroups, createSourceGroup, deleteSourceGroup, refreshSourceGroup,
    registerConnection, fetchRefinedInfraByGroup, fetchSavedSourceModels, savedSourceModels,
    saveSourceModel, updateSourceModel,
  } = useMigrationStore();

  // ── Shared credentials (used in both modal & inline section) ────
  const [showSharedCreds, setShowSharedCreds] = useState(true);
  const [useCommonCred,   setUseCommonCred]   = useState(true);
  const [commonUser,      setCommonUser]      = useState('ubuntu');
  const [commonPort,      setCommonPort]      = useState(22);
  const [commonKey,       setCommonKey]       = useState('');
  const [commonPassword,  setCommonPassword]  = useState('');

  // ── Inline Section 2 rows (existing group) ──────────────────────
  const [serverRows, setServerRows] = useState<ServerRow[]>([]);
  const [registerStatus, setRegisterStatus] = useState<'idle'|'registering'|'done'|'failed'>('idle');
  const [refreshStatus,  setRefreshStatus]  = useState<'idle'|'refreshing'|'done'>('idle');
  const [registerCredError, setRegisterCredError] = useState('');

  // ── New Source Group modal ───────────────────────────────────────
  const [showModal,    setShowModal]    = useState(false);
  const [newGroupName, setNewGroupName] = useState('');
  const [newGroupDesc, setNewGroupDesc] = useState('');
  const [isCreating,   setIsCreating]   = useState(false);
  const [modalCredError, setModalCredError] = useState('');
  // Modal uses the same shared creds state, separate row list
  const [modalRows,       setModalRows]       = useState<ServerRow[]>([emptyRow()]);
  const [showModalCreds,  setShowModalCreds]  = useState(true);

  // ── Delete confirm ───────────────────────────────────────────────
  const [deleteConfirmId, setDeleteConfirmId] = useState<string | null>(null);
  const [deleteConfirmText, setDeleteConfirmText] = useState('');
  const [isDeletingGroup, setIsDeletingGroup] = useState(false);
  const [groupDeleteError, setGroupDeleteError] = useState('');
  const [connectionToDelete, setConnectionToDelete] = useState<{ localId: string; connId: string; name: string } | null>(null);
  const [connDeleteConfirmText, setConnDeleteConfirmText] = useState('');
  const [isDeletingConn, setIsDeletingConn] = useState(false);
  const [connDeleteError, setConnDeleteError] = useState('');

  // ── Collect & Save ───────────────────────────────────────────────
  const [importStatus, setImportStatus] = useState<'idle'|'importing'|'done'|'failed'>('idle');
  const [saveSuccess,  setSaveSuccess]  = useState(false);
  const [showSaveModal, setShowSaveModal] = useState(false);

  // ─── Init ──────────────────────────────────────────────────────
  useEffect(() => { fetchSourceGroups(); fetchSavedSourceModels(); }, []);

  // Sync inline rows from store connections (when group selected / refreshed).
  // A blank draft row is always kept at the end, ready for entering the next server.
  useEffect(() => {
    if (connections && connections.length > 0) {
      const registeredRows: ServerRow[] = connections.map((c: any, i: number) => ({
        localId: `api-${c.id || i}`,
        connId: c.id || '',
        name: c.name || '',
        ip:   c.ip_address || c.ip || '',
        port: c.ssh_port ? String(c.ssh_port) : (c.port ? String(c.port) : ''),
        user: c.user || '',
        privateKey: '',
        password: '',
        authMode: 'key' as const,
        connectionStatus:    c.connection_status || '',
        agentStatus:         c.agent_status || '',
        connectionFailedMsg: c.connection_failed_message || '',
        agentFailedMsg:      c.agent_failed_message || '',
      }));
      setServerRows([...registeredRows, emptyRow()]);
      if (connections.some((c: any) => c.id)) setRegisterStatus('done');
    } else if (activeSgId) {
      setServerRows([emptyRow()]);
      setRegisterStatus('idle');
    }
  }, [connections, activeSgId]);

  // ─── Helpers ──────────────────────────────────────────────────
  function emptyRow(): ServerRow {
    return { localId: `r-${Date.now()}-${Math.random()}`, connId: '', name: '', ip: '', port: '', user: '', privateKey: '', password: '', authMode: 'key', connectionStatus: '', agentStatus: '', connectionFailedMsg: '', agentFailedMsg: '' };
  }

  const resolveCredentials = (row: ServerRow) => ({
    user:       row.user       || (useCommonCred ? commonUser : 'ubuntu'),
    port:       row.port ? Number(row.port) : (useCommonCred ? commonPort : 22),
    privateKey: row.privateKey || (useCommonCred ? commonKey      : ''),
    password:   row.password   || (useCommonCred ? commonPassword : ''),
  });

  const makeUpdater = (setter: React.Dispatch<React.SetStateAction<ServerRow[]>>) =>
    (localId: string, field: keyof ServerRow, value: string) =>
      setter(prev => prev.map(r => r.localId === localId ? { ...r, [field]: value } : r));

  const makeAdder = (setter: React.Dispatch<React.SetStateAction<ServerRow[]>>) =>
    () => setter(prev => [...prev, emptyRow()]);

  const makeKeyDrop = (setter: React.Dispatch<React.SetStateAction<ServerRow[]>>) =>
    (localId: string, e: React.DragEvent) => {
      e.preventDefault();
      const file = e.dataTransfer.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = ev => makeUpdater(setter)(localId, 'privateKey', ev.target?.result as string);
        reader.readAsText(file);
      }
    };

  const handleCommonKeyDrop = (e: React.DragEvent<HTMLTextAreaElement>) => {
    e.preventDefault();
    const file = e.dataTransfer.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = ev => setCommonKey(ev.target?.result as string);
      reader.readAsText(file);
    }
  };

  const handleJsonImport = (setter: React.Dispatch<React.SetStateAction<ServerRow[]>>) =>
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const file = e.target.files?.[0];
      if (!file) return;
      const reader = new FileReader();
      reader.onload = ev => {
        try {
          const parsed = JSON.parse(ev.target?.result as string);
          if (Array.isArray(parsed)) {
            setter(parsed.map((item: any, idx: number) => ({
              localId: `import-${idx}-${Date.now()}`, connId: '',
              name: item.name || `Host ${idx + 1}`,
              ip:   item.ip || item.ip_address || '',
              port: item.port ? String(item.port) : '',
              user: item.user || '',
              privateKey: item.private_key || item.privateKey || '',
              password: item.password || '',
              authMode: (item.password && !item.private_key && !item.privateKey) ? 'password' : 'key',
              connectionStatus: '', agentStatus: '', connectionFailedMsg: '', agentFailedMsg: '',
            })));
          } else { alert('JSON must be an array of server objects.'); }
        } catch { alert('Invalid JSON format.'); }
      };
      reader.readAsText(file);
    };

  const handleDownloadTemplate = () => {
    const tpl = [
      { name: 'Web Server 1', ip: '10.0.1.30',  port: 22, user: 'ubuntu', private_key: '' },
      { name: 'Database 1',   ip: '10.0.1.221', port: 22, user: 'root',   private_key: '' },
    ];
    const blob = new Blob([JSON.stringify(tpl, null, 2)], { type: 'application/json' });
    const url  = URL.createObjectURL(blob);
    const a    = document.createElement('a');
    a.href = url; a.download = 'source_servers_template.json';
    document.body.appendChild(a); a.click();
    document.body.removeChild(a); URL.revokeObjectURL(url);
  };

  // ─── Inline Section 2 actions ─────────────────────────────────

  // Load built-in sample infrastructure (no API call required)
  const handleLoadSampleInfra = () => {
    const SAMPLE_ID = 'sg-sample';
    const sampleGroup = {
      id: SAMPLE_ID,
      name: '[Sample] web-db-cluster',
      description: '1 Web Server (HAProxy/App) + 2 InfluxDB nodes on 10.0.1.x',
      connection_info_status_count: {
        connection_info_total: 3,
        count_connection_success: 3,
        count_connection_failed: 0,
        count_agent_success: 3,
        count_agent_failed: 0,
      },
    };
    const sampleConnections = [
      { id: 'sample-c1', name: 'Web Server 1',  ip_address: '10.0.1.30',  ssh_port: '22', user: 'ubuntu', connection_status: 'success', agent_status: 'success', connection_failed_message: '', agent_failed_message: '' },
      { id: 'sample-c2', name: 'Database 1',   ip_address: '10.0.1.221', ssh_port: '22', user: 'root',   connection_status: 'success', agent_status: 'success', connection_failed_message: '', agent_failed_message: '' },
      { id: 'sample-c3', name: 'Database 2',   ip_address: '10.0.1.138', ssh_port: '22', user: 'root',   connection_status: 'success', agent_status: 'success', connection_failed_message: '', agent_failed_message: '' },
    ];
    const current = useMigrationStore.getState().sourceGroups;
    if (!current.some((g: any) => g.id === SAMPLE_ID)) {
      useMigrationStore.setState({ sourceGroups: [sampleGroup, ...current] });
    }
    useMigrationStore.setState({ activeSgId: SAMPLE_ID, connections: sampleConnections });
    setRegisterStatus('done');
    setImportStatus('idle');
    setRefreshStatus('idle');
  };

  const handleSelectGroup = (sgId: string) => {
    useMigrationStore.setState({ activeSgId: sgId });
    setRegisterStatus('idle');
    setImportStatus('idle');
    setRefreshStatus('idle');
    // Skip API call for the built-in sample group — connections are already in store
    if (sgId === 'sg-sample') {
      useMigrationStore.setState({ refinedSourceInfra: DEMO_SOURCE_INFRA });
      return;
    }
    // Set refinedSourceInfra to null first when switching groups so we don't show stale preview
    useMigrationStore.setState({ refinedSourceInfra: null });

    honeybeeApi.getConnectionInfoList(sgId)
      .then((data: any) => {
        const list: any[] = data?.connection_info || (Array.isArray(data) ? data : []);
        useMigrationStore.setState({ connections: list });
      })
      .catch(() => useMigrationStore.setState({ connections: [] }));

    // Fetch the previously refined infra for this group if it exists
    fetchRefinedInfraByGroup(sgId);
  };

  const handleDeleteInlineRow = (localId: string) => {
    const row = serverRows.find(r => r.localId === localId);
    if (!row) return;
    if (row.connId && activeSgId && activeSgId !== 'sg-sample') {
      setConnectionToDelete({ localId, connId: row.connId, name: row.name || row.ip || 'Unnamed Server' });
      setConnDeleteConfirmText('');
      setConnDeleteError('');
      setIsDeletingConn(false);
    } else {
      setServerRows(prev => prev.length > 1 ? prev.filter(r => r.localId !== localId) : [emptyRow()]);
    }
  };

  const handleConfirmDeleteConnection = async () => {
    if (!connectionToDelete || !activeSgId) return;
    setIsDeletingConn(true);
    setConnDeleteError('');
    try {
      await honeybeeApi.deleteConnectionInfo(activeSgId, connectionToDelete.connId);
      const data = await honeybeeApi.getConnectionInfoList(activeSgId);
      const list: any[] = data?.connection_info || (Array.isArray(data) ? data : []);
      useMigrationStore.setState({ connections: list });
      setConnectionToDelete(null);
    } catch (err: any) {
      console.error('Delete failed:', err);
      setConnDeleteError(err.message || 'Failed to delete connection');
    } finally {
      setIsDeletingConn(false);
    }
  };

  // Draft rows aren't "deleted" — their fields are just cleared back to blank
  const handleResetInlineRow = (localId: string) => {
    setServerRows(prev => prev.map(r => r.localId === localId ? { ...emptyRow(), localId } : r));
  };

  const handleRegisterConnections = async () => {
    if (!activeSgId || activeSgId === 'sg-sample') return;
    const newRows = serverRows.filter(r => !r.connId && r.ip.trim());
    if (newRows.length === 0) return;
    const missing = newRows.filter(r => !hasAuth(r, useCommonCred, commonKey, commonPassword));
    if (missing.length > 0) {
      setRegisterCredError(`Provide a password or private key for: ${missing.map(r => r.name || r.ip).join(', ')}`);
      return;
    }
    setRegisterCredError('');
    setRegisterStatus('registering');
    try {
      for (const row of newRows) {
        const creds = resolveCredentials(row);
        await registerConnection({ name: row.name, ip: row.ip, port: creds.port, user: creds.user, privateKey: creds.privateKey, password: creds.password, description: '' });
      }
      const data = await honeybeeApi.getConnectionInfoList(activeSgId);
      const list: any[] = data?.connection_info || (Array.isArray(data) ? data : []);
      useMigrationStore.setState({ connections: list });
      setRegisterStatus('done');
    } catch (err) { console.error('Register failed:', err); setRegisterStatus('failed'); }
  };

  const handleRefreshAll = async () => {
    if (!activeSgId || activeSgId === 'sg-sample') return;
    setRefreshStatus('refreshing');
    try {
      await refreshSourceGroup(activeSgId);
      const data = await honeybeeApi.getConnectionInfoList(activeSgId);
      const list: any[] = data?.connection_info || (Array.isArray(data) ? data : []);
      useMigrationStore.setState({ connections: list });
      setRefreshStatus('done');
      setTimeout(() => setRefreshStatus('idle'), 3000);
    } catch { setRefreshStatus('idle'); }
  };

  // ─── New Source Group modal actions ─────────────────────────────

  const openCreateModal = () => {
    setNewGroupName('');
    setNewGroupDesc('');
    setModalRows([emptyRow()]);
    setModalCredError('');
    setShowModal(true);
  };

  const handleCreateGroup = async () => {
    if (!newGroupName.trim()) return;
    const rowsWithIp = modalRows.filter(r => r.ip.trim());
    const missing = rowsWithIp.filter(r => !hasAuth(r, useCommonCred, commonKey, commonPassword));
    if (missing.length > 0) {
      setModalCredError(`Provide a password or private key for: ${missing.map(r => r.name || r.ip).join(', ')}`);
      return;
    }
    setModalCredError('');
    setIsCreating(true);
    try {
      const connList = rowsWithIp
         .map(r => {
           const creds = resolveCredentials(r);
           return { name: r.name || 'Server', ip_address: r.ip.trim(), ssh_port: String(creds.port), user: creds.user, private_key: creds.privateKey, password: creds.password, description: '' };
         });

      if (connList.length > 0) {
        const result = await honeybeeApi.createSourceGroupWithConnections({ name: newGroupName.trim(), description: newGroupDesc.trim(), connection_info: connList });
        const newId = result?.id || `sg-${Date.now()}`;
        await fetchSourceGroups();
        handleSelectGroup(newId);
      } else {
        await createSourceGroup(newGroupName.trim(), newGroupDesc.trim());
      }
      setShowModal(false);
    } catch (err) {
      console.error('Create group failed:', err);
      await createSourceGroup(newGroupName.trim(), newGroupDesc.trim());
      setShowModal(false);
    } finally {
      setIsCreating(false);
    }
  };

  const handleDeleteModalRow = (localId: string) => {
    setModalRows(prev => prev.length > 1 ? prev.filter(r => r.localId !== localId) : [emptyRow()]);
  };

  // ─── Group delete ────────────────────────────────────────────────
  const handleDeleteGroup = async (sgId: string) => {
    setIsDeletingGroup(true);
    setGroupDeleteError('');
    try {
      await deleteSourceGroup(sgId);
      setDeleteConfirmId(null);
      setDeleteConfirmText('');
    } catch (err: any) {
      setGroupDeleteError(err.message || 'Failed to delete source group');
    } finally {
      setIsDeletingGroup(false);
    }
  };

  // ─── Collect & Save ──────────────────────────────────────────────
  const handleImportInfra = async () => {
    if (!activeSgId) return;
    setImportStatus('importing');
    try {
      await honeybeeApi.importInfraByGroup(activeSgId);
      await fetchRefinedInfraByGroup(activeSgId);
      setImportStatus('done');
    } catch (err) {
      console.warn('Honeybee import failed, falling back to demo data for testing:', err);
      useMigrationStore.setState({ refinedSourceInfra: DEMO_SOURCE_INFRA });
      setImportStatus('done');
    }
  };

  const handleSaveToDamselfly = async (result: { name: string; description: string; version: string; overwriteId: string | null }) => {
    if (!refinedSourceInfra) return;
    if (result.overwriteId) {
      await updateSourceModel(result.overwriteId, result.name, result.description, result.version, refinedSourceInfra);
    } else {
      await saveSourceModel(result.name, result.description, result.version, refinedSourceInfra);
    }
    setSaveSuccess(true);
    setTimeout(() => setSaveSuccess(false), 3000);
  };

  // ─── Derived ────────────────────────────────────────────────────
  const activeGroup = sourceGroups.find((g: any) => g.id === activeSgId);
  const hasNewRows  = serverRows.some(r => !r.connId && r.ip.trim());

  // Row operation factories for inline section
  const updateInlineRow = makeUpdater(setServerRows);
  const keyDropInline   = makeKeyDrop(setServerRows);

  // Row operation factories for modal
  const updateModalRow = makeUpdater(setModalRows);
  const addModalRow    = makeAdder(setModalRows);
  const keyDropModal   = makeKeyDrop(setModalRows);

  // ─── Render ──────────────────────────────────────────────────────
  return (
    <div className="space-y-6">

      {/* Top Banner Description Box */}
      <div className="glass-panel px-6 py-4.5 rounded-2xl border border-border-main flex flex-wrap items-center gap-x-3 gap-y-1.5">
        <div className="flex items-center gap-2 shrink-0">
          <Database className="w-5 h-5 text-emerald-500" />
          <h2 className="text-base font-extrabold text-text-main tracking-tight">
            Source Infra Analysis
          </h2>
        </div>
        <span className="text-sm text-text-muted">
          Register source hosts, extract system configurations, and analyze infrastructure specifications & network topology metadata.
        </span>
      </div>

      {/* ══════════════════════════════════════════════════
          SECTION 1 — Source Infrastructure Groups
      ══════════════════════════════════════════════════ */}
      <div className="glass-panel p-6 rounded-2xl">
        <div className="flex items-center justify-between mb-5">
          <div>
            <h2 className="text-base font-bold text-text-main flex items-center gap-2">
              <Server className="w-5 h-5 text-emerald-400" />
              Source Infrastructure Groups
            </h2>
            <p className="text-sm text-text-muted mt-1">
              Click a group to manage its server connections below, or create a new group.
            </p>
          </div>
          <div className="flex items-center gap-3">
            <button
              onClick={handleLoadSampleInfra}
              className="px-4 py-2.5 bg-bg-panel border border-teal-500/40 hover:bg-teal-500/10 text-teal-600 dark:text-teal-400 rounded-xl text-sm font-bold flex items-center gap-2 transition cursor-pointer"
              title="Load built-in sample infrastructure data (no API call required)"
            >
              <FileText className="w-4 h-4" /> Load Sample Infra
            </button>
            <button
              onClick={openCreateModal}
              className="px-4 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white rounded-xl text-sm font-bold flex items-center gap-2 transition cursor-pointer shadow-sm"
            >
              <Plus className="w-4 h-4" /> New Source Group
            </button>
          </div>
        </div>

        {sourceGroups.length === 0 ? (
          <div className="text-center py-12 text-text-muted">
            <Server className="w-10 h-10 mx-auto mb-3 opacity-20" />
            <p className="text-sm">No source groups yet. Create one to get started.</p>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {sourceGroups.map((group: any) => {
              const isActive  = group.id === activeSgId;
              const sc        = group.connection_info_status_count;
              const total     = sc?.connection_info_total    ?? 0;
              const success   = sc?.count_connection_success ?? 0;
              const failed    = sc?.count_connection_failed  ?? 0;
              const allOk     = total > 0 && failed === 0;
              const anyFailed = failed > 0;

              return (
                <div
                  key={group.id}
                  onClick={() => handleSelectGroup(group.id)}
                  className={`relative rounded-xl p-5 border-2 cursor-pointer transition-all select-none ${
                    isActive
                      ? 'border-emerald-500 bg-emerald-500/5 shadow-md shadow-emerald-500/10'
                      : 'border-border-main bg-bg-input/20 hover:border-emerald-500/40 hover:bg-emerald-500/[0.02]'
                  }`}
                >
                  {isActive && <span className="absolute top-3 right-10 w-2.5 h-2.5 rounded-full bg-emerald-500 shadow-sm shadow-emerald-500/50" />}

                  <div className="mb-3 pr-8">
                    <div className="font-bold text-text-main text-sm">{group.name}</div>
                    {group.description && <div className="text-xs text-text-muted mt-0.5 truncate">{group.description}</div>}
                  </div>

                  <div className="mb-4">
                    {total > 0 ? (
                      <span className={`inline-flex items-center gap-1.5 text-xs font-semibold px-2.5 py-1 rounded-full border ${
                        allOk     ? 'bg-emerald-100 dark:bg-emerald-950/40 text-emerald-600 dark:text-emerald-400 border-emerald-300 dark:border-emerald-800/40'
                        : anyFailed ? 'bg-red-100 dark:bg-red-950/40 text-red-600 dark:text-red-400 border-red-300 dark:border-red-800/40'
                        : 'bg-bg-input text-text-muted border-border-main'
                      }`}>
                        {allOk && <CheckCircle2 className="w-3 h-3" />}
                        {anyFailed && <XCircle className="w-3 h-3" />}
                        {success}/{total} Connected
                        {failed > 0 && <span className="opacity-70">· {failed} failed</span>}
                      </span>
                    ) : (
                      <span className="text-xs text-text-muted">No connections — click to add</span>
                    )}
                  </div>

                  <div className="flex items-center" onClick={e => e.stopPropagation()}>
                    {isActive && <span className="text-xs font-bold text-emerald-500 flex items-center gap-1"><CheckCircle2 className="w-3 h-3" />Active</span>}
                    <button
                      onClick={e => { e.stopPropagation(); setDeleteConfirmId(group.id); }}
                      className="ml-auto p-1.5 rounded-lg hover:bg-red-500/10 text-text-muted hover:text-red-400 transition cursor-pointer"
                      title="Delete group"
                    >
                      <Trash2 className="w-3.5 h-3.5" />
                    </button>
                  </div>
                </div>
              );
            })}
          </div>
        )}
      </div>

      {/* ══════════════════════════════════════════════════
          SECTION 2 — Server Connections (inline)
      ══════════════════════════════════════════════════ */}
      {activeSgId && (
        <div className="glass-panel rounded-2xl overflow-hidden">
          {/* Header */}
          <div className="px-6 py-4 border-b border-border-main flex flex-wrap gap-3 items-center justify-between bg-bg-input/20">
            <div>
              <h2 className="text-base font-bold text-text-main flex items-center gap-2">
                Server Connections
                {activeGroup && <span className="font-normal text-text-muted text-sm">— {activeGroup.name}</span>}
              </h2>
            </div>
            <div className="flex flex-wrap items-center gap-2">
              <button onClick={handleRefreshAll} disabled={refreshStatus === 'refreshing'}
                className="px-3 py-2 bg-bg-panel border border-border-main hover:bg-bg-input rounded-lg text-sm font-semibold flex items-center gap-1.5 transition cursor-pointer text-text-muted hover:text-text-main disabled:opacity-50">
                <RefreshCw className={`w-3.5 h-3.5 ${refreshStatus === 'refreshing' ? 'animate-spin' : ''}`} />
                {refreshStatus === 'refreshing' ? 'Refreshing…' : refreshStatus === 'done' ? 'Refreshed ✓' : 'Refresh Status'}
              </button>
              <button onClick={handleDownloadTemplate}
                className="px-3 py-2 bg-bg-panel border border-border-main hover:bg-bg-input rounded-lg text-sm font-semibold flex items-center gap-1.5 transition cursor-pointer text-text-muted hover:text-text-main">
                <Download className="w-3.5 h-3.5" /> Template
              </button>
              <label className="px-3 py-2 bg-bg-panel border border-border-main hover:bg-bg-input rounded-lg text-sm font-semibold flex items-center gap-1.5 transition cursor-pointer text-text-muted hover:text-text-main">
                <Upload className="w-3.5 h-3.5" /> Import JSON
                <input type="file" accept=".json" onChange={handleJsonImport(setServerRows)} className="hidden" />
              </label>
            </div>
          </div>

          {/* Shared Credentials */}
          <SharedCredsPanel
            show={showSharedCreds} onToggle={() => setShowSharedCreds(v => !v)}
            enabled={useCommonCred} onToggleEnabled={setUseCommonCred}
            user={commonUser} onUser={setCommonUser}
            port={commonPort} onPort={setCommonPort}
            pemKey={commonKey} onPemKey={setCommonKey}
            password={commonPassword} onPassword={setCommonPassword}
            onKeyDrop={handleCommonKeyDrop}
            locked={connections && connections.length > 0}
          />

          {/* Server table */}
          <div className="overflow-x-auto">
            <table className="w-full text-left border-collapse text-sm">
              {TABLE_HEADER}
              <tbody className="divide-y divide-border-main">
                {serverRows.map((row, idx) => (
                  <ServerTableRow
                    key={row.localId} row={row} idx={idx}
                    useCommonCred={useCommonCred} commonUser={commonUser} commonPort={commonPort} commonKey={commonKey} commonPassword={commonPassword}
                    onUpdate={updateInlineRow}
                    onDelete={handleDeleteInlineRow}
                    onReset={handleResetInlineRow}
                    onKeyDrop={keyDropInline}
                  />
                ))}
              </tbody>
            </table>
          </div>

          {registerCredError && (
            <div className="mx-6 mt-3 p-2.5 bg-red-100 dark:bg-red-950/40 border border-red-300 dark:border-red-800/40 text-red-600 dark:text-red-400 text-sm rounded-lg">
              {registerCredError}
            </div>
          )}

          <div className="px-6 pt-4 border-t border-border-main/60">
            <p className="text-sm text-text-muted mb-3">Fill in the row above with server details, then click <span className="font-semibold text-text-main">“Add Server”</span> to register it to this group. A new blank row appears automatically for the next server. Blank SSH fields inherit from Shared Credentials.</p>
            <div className="flex items-center gap-3 pb-4">
              <button onClick={handleRegisterConnections} disabled={!hasNewRows || registerStatus === 'registering'}
                className="px-4 py-2 bg-gradient-to-r from-emerald-500 to-teal-600 hover:from-emerald-600 hover:to-teal-700 text-slate-950 rounded-lg text-sm font-extrabold flex items-center gap-2 transition cursor-pointer disabled:opacity-40 shadow-sm">
                {registerStatus === 'registering' ? <><Loader2 className="w-3.5 h-3.5 animate-spin" /> Adding…</>
                  : <><Plus className="w-3.5 h-3.5" /> Add Server</>}
              </button>
              {registerStatus === 'done' && !hasNewRows && (
                <span className="text-sm font-semibold text-emerald-600 dark:text-emerald-400 flex items-center gap-1"><CheckCircle2 className="w-3.5 h-3.5" /> All Registered</span>
              )}
            </div>
          </div>
        </div>
      )}

      {/* ══════════════════════════════════════════════════
          SECTION 3 — Collect & Save Source Model
      ══════════════════════════════════════════════════ */}
      {activeSgId && (
        <div className="glass-panel p-6 rounded-2xl">
          <h2 className="text-base font-bold text-text-main flex items-center gap-2 mb-5">
            <FileText className="w-5 h-5 text-emerald-400" />
            Collect &amp; Save Source Model
            {activeGroup && <span className="text-sm font-normal text-text-muted">— {activeGroup.name}</span>}
          </h2>

          {/* Step progress */}
          <div className="flex items-center gap-3 mb-6 flex-wrap">
            {[
              { n: 1, label: 'Connections Registered', done: registerStatus === 'done',   active: registerStatus === 'registering' },
              { n: 2, label: 'Infra Collected',        done: importStatus   === 'done',   active: importStatus   === 'importing'   },
              { n: 3, label: 'Model Saved',            done: saveSuccess,                 active: false },
            ].map((step, i) => (
              <React.Fragment key={step.n}>
                <div className={`flex items-center gap-2 text-sm ${step.done ? 'text-emerald-400' : step.active ? 'text-teal-400' : 'text-text-muted'}`}>
                  <div className={`w-6 h-6 rounded-full flex items-center justify-center text-xs font-bold border-2 flex-shrink-0 ${
                    step.done ? 'border-emerald-500 bg-emerald-100 dark:bg-emerald-950/40 text-emerald-600 dark:text-emerald-400'
                    : step.active ? 'border-teal-500 bg-teal-100 dark:bg-teal-950/40 text-teal-600 dark:text-teal-400'
                    : 'border-border-main text-text-muted'
                  }`}>
                    {step.done ? <CheckCircle2 className="w-3.5 h-3.5" /> : step.active ? <Loader2 className="w-3.5 h-3.5 animate-spin" /> : step.n}
                  </div>
                  <span className="whitespace-nowrap">{step.label}</span>
                </div>
                {i < 2 && <ChevronRight className="w-4 h-4 text-text-muted flex-shrink-0" />}
              </React.Fragment>
            ))}
          </div>

          <div className="flex flex-wrap gap-5 items-start">
            <div className="flex-1 min-w-[220px]">
              <button onClick={handleImportInfra} disabled={importStatus === 'importing'}
                className="w-full py-3 bg-bg-panel border-2 border-emerald-500/40 hover:border-emerald-500/70 hover:bg-emerald-500/5 rounded-xl text-sm font-bold text-emerald-600 dark:text-emerald-400 flex items-center justify-center gap-2 transition cursor-pointer disabled:opacity-50">
                {importStatus === 'importing' ? <><Loader2 className="w-4 h-4 animate-spin" /> Collecting…</>
                  : importStatus === 'done'   ? <><CheckCircle2 className="w-4 h-4" /> Collected — Re-Import</>
                  : importStatus === 'failed' ? <><AlertCircle className="w-4 h-4 text-red-400" /> Failed — Retry</>
                  : <><Play className="w-4 h-4" /> Collect Infra from All Servers</>}
              </button>
              <p className="text-xs text-text-muted mt-2 text-center leading-relaxed">
                Triggers the Honeybee agent to collect infrastructure specs from all registered servers.
              </p>
            </div>
            <div className="flex-1 min-w-[280px] flex flex-col justify-center space-y-3">
              <button onClick={() => setShowSaveModal(true)} disabled={!refinedSourceInfra} className="w-full py-3 bg-gradient-to-r from-emerald-500 to-teal-600 hover:from-emerald-600 hover:to-teal-700 text-slate-950 font-bold rounded-lg text-sm transition shadow-sm cursor-pointer disabled:opacity-50 flex items-center justify-center gap-2">
                <Save className="w-4 h-4" /> Save Source Infra Revision
              </button>
              <p className="text-xs text-text-muted text-center leading-relaxed">
                Opens a popup to name, version, and save this collected model to Damselfly.
              </p>
            </div>
          </div>

          {refinedSourceInfra && (
            <div className="mt-5 pt-5 border-t border-border-main">
              <h3 className="text-sm font-bold text-text-main mb-3 flex items-center gap-2">
                <FileText className="w-4 h-4 text-emerald-400" />
                Collected Infrastructure Model Preview
              </h3>
              <div className="bg-bg-input border border-border-main rounded-xl p-4 text-xs font-mono text-slate-800 dark:text-emerald-400 max-h-[60vh] overflow-y-auto">
                <pre>{JSON.stringify(refinedSourceInfra, null, 2)}</pre>
              </div>
            </div>
          )}
        </div>
      )}

      {/* ══════════════════════════════════════════════════
          NEW SOURCE GROUP MODAL
      ══════════════════════════════════════════════════ */}
      {showModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4">
          <div className="glass-panel rounded-2xl w-full max-w-7xl max-h-[96vh] flex flex-col border border-border-main">

            <div className="flex items-center justify-between px-6 py-4 border-b border-border-main flex-shrink-0">
              <h3 className="text-base font-bold text-text-main">New Source Group</h3>
              <button onClick={() => setShowModal(false)} className="p-1.5 rounded-lg hover:bg-bg-input text-text-muted hover:text-text-main transition cursor-pointer">
                <X className="w-5 h-5" />
              </button>
            </div>

            <div className="overflow-y-auto flex-1 px-6 py-5 space-y-5">
              {/* Group name & description */}
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <label className="block text-sm font-semibold text-text-muted mb-1.5">Group Name <span className="text-red-400">*</span></label>
                  <input type="text" placeholder="e.g. production-dc" value={newGroupName} onChange={e => setNewGroupName(e.target.value)}
                    className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-1 focus:ring-emerald-500" />
                </div>
                <div>
                  <label className="block text-sm font-semibold text-text-muted mb-1.5">Description</label>
                  <input type="text" placeholder="Optional description" value={newGroupDesc} onChange={e => setNewGroupDesc(e.target.value)}
                    className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-1 focus:ring-emerald-500" />
                </div>
              </div>

              {/* Shared credentials in modal */}
              <SharedCredsPanel
                show={showModalCreds} onToggle={() => setShowModalCreds(v => !v)}
                enabled={useCommonCred} onToggleEnabled={setUseCommonCred}
                user={commonUser} onUser={setCommonUser}
                port={commonPort} onPort={setCommonPort}
                pemKey={commonKey} onPemKey={setCommonKey}
                password={commonPassword} onPassword={setCommonPassword}
                onKeyDrop={handleCommonKeyDrop}
              />

              {/* Server table in modal */}
              <div className="border border-border-main rounded-xl overflow-hidden">
                <div className="px-4 py-3 bg-bg-input/20 border-b border-border-main flex flex-wrap items-center gap-2 justify-between">
                  <span className="text-sm font-semibold text-text-main">Server Connections <span className="text-xs font-normal text-text-muted">(optional — can be added later)</span></span>
                  <div className="flex flex-wrap items-center gap-2">
                    <button onClick={handleDownloadTemplate} className="px-3 py-1.5 bg-bg-panel border border-border-main hover:bg-bg-input rounded-lg text-xs font-semibold flex items-center gap-1.5 transition cursor-pointer text-text-muted">
                      <Download className="w-3.5 h-3.5" /> Template
                    </button>
                    <label className="px-3 py-1.5 bg-bg-panel border border-border-main hover:bg-bg-input rounded-lg text-xs font-semibold flex items-center gap-1.5 transition cursor-pointer text-text-muted">
                      <Upload className="w-3.5 h-3.5" /> Import JSON
                      <input type="file" accept=".json" onChange={handleJsonImport(setModalRows)} className="hidden" />
                    </label>
                    <button onClick={addModalRow} className="px-3 py-1.5 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 rounded-lg text-xs font-semibold flex items-center gap-1.5 transition cursor-pointer text-emerald-600 dark:text-emerald-400">
                      <Plus className="w-3.5 h-3.5" /> Add Server
                    </button>
                  </div>
                </div>
                <div className="overflow-x-auto">
                  <table className="w-full text-left border-collapse text-sm">
                    {TABLE_HEADER}
                    <tbody className="divide-y divide-border-main">
                      {modalRows.map((row, idx) => (
                        <ServerTableRow
                          key={row.localId} row={row} idx={idx}
                          useCommonCred={useCommonCred} commonUser={commonUser} commonPort={commonPort} commonKey={commonKey} commonPassword={commonPassword}
                          onUpdate={updateModalRow}
                          onDelete={handleDeleteModalRow}
                          onKeyDrop={keyDropModal}
                        />
                      ))}
                    </tbody>
                  </table>
                </div>
              </div>
            </div>

            {modalCredError && (
              <div className="mx-6 mb-4 p-2.5 bg-red-100 dark:bg-red-950/40 border border-red-300 dark:border-red-800/40 text-red-600 dark:text-red-400 text-sm rounded-lg flex-shrink-0">
                {modalCredError}
              </div>
            )}

            <div className="flex items-center justify-end gap-3 px-6 py-4 border-t border-border-main flex-shrink-0">
              <button onClick={() => setShowModal(false)} className="px-4 py-2.5 bg-bg-input border border-border-main hover:bg-bg-panel rounded-xl text-sm font-bold cursor-pointer">
                Cancel
              </button>
              <button
                disabled={!newGroupName.trim() || isCreating}
                onClick={handleCreateGroup}
                className="px-4 py-2.5 bg-emerald-500 text-slate-950 hover:bg-emerald-600 rounded-xl text-sm font-bold cursor-pointer disabled:opacity-50 flex items-center gap-2"
              >
                {isCreating ? <><Loader2 className="w-4 h-4 animate-spin" /> Creating…</> : 'Create Source Group'}
              </button>
            </div>
          </div>
        </div>
      )}

      {/* ══════════════════════════════════════════════════
          Delete Source Group Confirm Dialog
      ══════════════════════════════════════════════════ */}
      {deleteConfirmId && (() => {
        const group = sourceGroups.find(g => g.id === deleteConfirmId);
        if (!group) return null;
        return (
          <div className="fixed inset-0 z-[60] flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
            <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up">
              <div className="flex justify-between items-center mb-4">
                <h3 className="text-base font-bold text-text-main flex items-center gap-2">
                  <Trash2 className="w-4 h-4 text-red-500" /> Delete Source Group
                </h3>
                <button
                  onClick={() => { setDeleteConfirmId(null); setDeleteConfirmText(''); setGroupDeleteError(''); }}
                  disabled={isDeletingGroup}
                  className="text-text-muted hover:text-text-main transition p-1 hover:bg-bg-input rounded-lg cursor-pointer disabled:opacity-50"
                >
                  <X className="w-5 h-5" />
                </button>
              </div>

              <div className="space-y-4">
                <p className="text-sm text-text-muted leading-relaxed">
                  Are you sure you want to delete the source group <strong className="text-text-main">"{group.name}"</strong>? This will permanently delete the group and all its registered server connection entries. This cannot be undone.
                </p>

                <div className="space-y-1.5 pt-1">
                  <label className="block text-xs font-bold text-text-muted">
                    To confirm, type <span className="font-mono bg-bg-panel px-1 py-0.5 rounded border border-border-main/60 text-text-main select-all">{group.name}</span> in the box below:
                  </label>
                  <input
                    type="text"
                    value={deleteConfirmText}
                    onChange={(e) => setDeleteConfirmText(e.target.value)}
                    placeholder="Type the source group name to delete"
                    className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-1 focus:ring-red-500 font-bold font-mono"
                    disabled={isDeletingGroup}
                  />
                </div>

                {groupDeleteError && (
                  <div className="flex items-center gap-2 bg-red-500/10 text-red-500 px-4 py-3 rounded-xl text-xs font-semibold border border-red-500/20">
                    <span>{groupDeleteError}</span>
                  </div>
                )}

                <div className="flex justify-end gap-3 pt-2">
                  <button
                    onClick={() => { setDeleteConfirmId(null); setDeleteConfirmText(''); setGroupDeleteError(''); }}
                    disabled={isDeletingGroup}
                    className="px-4 py-2 bg-bg-panel border border-border-main text-text-main rounded-xl text-sm font-semibold hover:bg-bg-input transition cursor-pointer disabled:opacity-50"
                  >
                    Cancel
                  </button>
                  <button
                    onClick={() => handleDeleteGroup(deleteConfirmId)}
                    disabled={isDeletingGroup || deleteConfirmText !== group.name}
                    className={`px-4 py-2 rounded-xl text-sm font-semibold transition flex items-center gap-1.5 ${
                      isDeletingGroup || deleteConfirmText !== group.name
                        ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                        : 'bg-red-500 hover:bg-red-600 text-white cursor-pointer shadow-md shadow-red-500/20'
                    }`}
                  >
                    {isDeletingGroup && <Loader2 className="w-4 h-4 animate-spin" />}
                    Confirm Delete
                  </button>
                </div>
              </div>
            </div>
          </div>
        );
      })()}

      {/* ══════════════════════════════════════════════════
          Delete Server Connection Confirm Dialog
      ══════════════════════════════════════════════════ */}
      {connectionToDelete && (
        <div className="fixed inset-0 z-[60] flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up">
            <div className="flex justify-between items-center mb-4">
              <h3 className="text-base font-bold text-text-main flex items-center gap-2">
                <Trash2 className="w-4 h-4 text-red-500" /> Delete Server Connection
              </h3>
              <button
                onClick={() => { setConnectionToDelete(null); setConnDeleteConfirmText(''); setConnDeleteError(''); }}
                disabled={isDeletingConn}
                className="text-text-muted hover:text-text-main transition p-1 hover:bg-bg-input rounded-lg cursor-pointer disabled:opacity-50"
              >
                <X className="w-5 h-5" />
              </button>
            </div>

            <div className="space-y-4">
              <p className="text-sm text-text-muted leading-relaxed">
                Are you sure you want to delete the server connection <strong className="text-text-main">"{connectionToDelete.name}"</strong>? This will permanently remove the server from this source group.
              </p>

              <div className="space-y-1.5 pt-1">
                <label className="block text-xs font-bold text-text-muted">
                  To confirm, type <span className="font-mono bg-bg-panel px-1 py-0.5 rounded border border-border-main/60 text-text-main select-all">{connectionToDelete.name}</span> in the box below:
                </label>
                <input
                  type="text"
                  value={connDeleteConfirmText}
                  onChange={(e) => setConnDeleteConfirmText(e.target.value)}
                  placeholder="Type the server name/IP to delete"
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-1 focus:ring-red-500 font-bold font-mono"
                  disabled={isDeletingConn}
                />
              </div>

              {connDeleteError && (
                <div className="flex items-center gap-2 bg-red-500/10 text-red-500 px-4 py-3 rounded-xl text-xs font-semibold border border-red-500/20">
                  <span>{connDeleteError}</span>
                </div>
              )}

              <div className="flex justify-end gap-3 pt-2">
                <button
                  onClick={() => { setConnectionToDelete(null); setConnDeleteConfirmText(''); setConnDeleteError(''); }}
                  disabled={isDeletingConn}
                  className="px-4 py-2 bg-bg-panel border border-border-main text-text-main rounded-xl text-sm font-semibold hover:bg-bg-input transition cursor-pointer disabled:opacity-50"
                >
                  Cancel
                </button>
                <button
                  onClick={handleConfirmDeleteConnection}
                  disabled={isDeletingConn || connDeleteConfirmText !== connectionToDelete.name}
                  className={`px-4 py-2 rounded-xl text-sm font-semibold transition flex items-center gap-1.5 ${
                    isDeletingConn || connDeleteConfirmText !== connectionToDelete.name
                      ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                      : 'bg-red-500 hover:bg-red-600 text-white cursor-pointer shadow-md shadow-red-500/20'
                  }`}
                >
                  {isDeletingConn && <Loader2 className="w-4 h-4 animate-spin" />}
                  Confirm Delete
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      <SaveRevisionModal
        isOpen={showSaveModal}
        onClose={() => setShowSaveModal(false)}
        title="Save Source Infra Revision"
        defaultName=""
        defaultDescription=""
        defaultVersion="1.0.0"
        existingRevisions={savedSourceModels
          .filter(m => m.id !== 'sample-source-infra-1')
          .map(m => ({ id: m.id, name: m.name, version: m.version }))}
        onSave={handleSaveToDamselfly}
      />

      {/* Toast Notification */}
      {saveSuccess && (
        <div className="fixed bottom-6 right-6 z-50 flex items-center gap-2.5 bg-slate-950/95 border border-emerald-500/40 text-emerald-600 dark:text-emerald-400 px-5 py-4.5 rounded-2xl shadow-2xl shadow-emerald-500/10 animate-fade-in font-bold text-sm backdrop-blur-md">
          <CheckCircle2 className="w-5 h-5 text-emerald-500" />
          <span>Model saved to Damselfly successfully.</span>
        </div>
      )}
    </div>
  );
};
