'use client';

import React from 'react';
import { Folder, FileText, Database, CheckCircle2, XCircle } from 'lucide-react';

export interface DirectoryTreeItem {
  id: string;
  displayPath: string;
  pattern: string;
  sizeLabel: string;
  description: string;
  kind: 'dir' | 'file' | 'db';
  defaultRole: 'include' | 'exclude';
}

export const AIMX_DIRECTORY_TREE_ITEMS: DirectoryTreeItem[] = [
  {
    id: 'ollama-blobs',
    displayPath: '/root/.ollama/models/blobs',
    pattern: '~/.ollama/models/blobs/**',
    sizeLabel: '4.8 GB',
    description: 'Gemma & LLM GGUF weight blobs',
    kind: 'dir',
    defaultRole: 'include',
  },
  {
    id: 'ollama-manifests',
    displayPath: '/root/.ollama/models/manifests',
    pattern: '~/.ollama/models/manifests/**',
    sizeLabel: '128 KB',
    description: 'Ollama model registry manifests',
    kind: 'dir',
    defaultRole: 'include',
  },
  {
    id: 'openwebui-db',
    displayPath: '/app/backend/data/webui.db',
    pattern: '/app/backend/data/webui.db',
    sizeLabel: '42 MB',
    description: 'OpenWebUI users, prompts & chat history SQLite DB',
    kind: 'db',
    defaultRole: 'include',
  },
  {
    id: 'openwebui-uploads',
    displayPath: '/app/backend/data/uploads',
    pattern: '/app/backend/data/uploads/**',
    sizeLabel: '310 MB',
    description: 'Knowledge base & RAG uploaded documents',
    kind: 'dir',
    defaultRole: 'include',
  },
  {
    id: 'var-log-syslog',
    displayPath: '/var/log/syslog',
    pattern: '*.log',
    sizeLabel: '1.2 GB',
    description: 'System & daemon log files (recommended exclude)',
    kind: 'file',
    defaultRole: 'exclude',
  },
  {
    id: 'tmp-cache',
    displayPath: '/tmp',
    pattern: '*.tmp',
    sizeLabel: '240 MB',
    description: 'Temporary runtime files (recommended exclude)',
    kind: 'dir',
    defaultRole: 'exclude',
  },
  {
    id: 'python-pycache',
    displayPath: '/app/backend/__pycache__',
    pattern: '__pycache__/**',
    sizeLabel: '18 MB',
    description: 'Compiled Python bytecode cache (recommended exclude)',
    kind: 'dir',
    defaultRole: 'exclude',
  },
];

interface DirectoryTreeExplorerProps {
  includeFilter: string;
  excludeFilter: string;
  onToggleInclude: (pattern: string) => void;
  onToggleExclude: (pattern: string) => void;
}

export const DirectoryTreeExplorer: React.FC<DirectoryTreeExplorerProps> = ({
  includeFilter,
  excludeFilter,
  onToggleInclude,
  onToggleExclude,
}) => {
  const includeSet = new Set(
    includeFilter
      .split(',')
      .map((s) => s.trim())
      .filter(Boolean)
  );
  const excludeSet = new Set(
    excludeFilter
      .split(',')
      .map((s) => s.trim())
      .filter(Boolean)
  );

  return (
    <div className="bg-bg-panel/80 border border-border-main rounded-xl p-4 space-y-3 animate-fade-in">
      <div className="flex flex-wrap items-center justify-between gap-2 border-b border-border-main/30 pb-2.5">
        <div className="flex items-center gap-2">
          <Folder className="w-4 h-4 text-emerald-500" />
          <span className="text-sm font-extrabold text-text-main">
            Source Filesystem Directory Tree
          </span>
        </div>
        <span className="text-xs text-text-muted">
          Click Include or Exclude to sync patterns with Data Filter Settings above
        </span>
      </div>

      <div className="divide-y divide-border-main/30">
        {AIMX_DIRECTORY_TREE_ITEMS.map((item) => {
          const isIncluded = includeSet.has(item.pattern);
          const isExcluded = excludeSet.has(item.pattern);
          const Icon =
            item.kind === 'db'
              ? Database
              : item.kind === 'dir'
              ? Folder
              : FileText;

          return (
            <div
              key={item.id}
              className="py-2.5 px-2 flex flex-col sm:flex-row sm:items-center justify-between gap-2 hover:bg-bg-input/40 rounded-lg transition"
            >
              <div className="flex items-start sm:items-center gap-2.5 min-w-0">
                <Icon
                  className={`w-4 h-4 shrink-0 mt-0.5 sm:mt-0 ${
                    isIncluded
                      ? 'text-emerald-500'
                      : isExcluded
                      ? 'text-rose-500'
                      : 'text-text-muted'
                  }`}
                />
                <div className="min-w-0">
                  <div className="flex flex-wrap items-center gap-2">
                    <span className="text-sm font-extrabold text-text-main font-mono">
                      {item.displayPath}
                    </span>
                    <span className="text-xs font-mono px-2 py-0.5 rounded bg-bg-input border border-border-main text-text-muted">
                      Pattern: {item.pattern}
                    </span>
                    <span className="text-xs font-extrabold font-mono text-teal-600 dark:text-teal-400">
                      {item.sizeLabel}
                    </span>
                  </div>
                  <p className="text-xs text-text-muted mt-0.5">{item.description}</p>
                </div>
              </div>

              <div className="flex items-center gap-2 shrink-0">
                <button
                  type="button"
                  onClick={() => onToggleInclude(item.pattern)}
                  className={`px-2.5 py-1 rounded-lg text-xs font-extrabold border transition cursor-pointer flex items-center gap-1 ${
                    isIncluded
                      ? 'bg-emerald-500 text-slate-950 border-emerald-500 shadow-xs'
                      : 'bg-bg-input border-border-main text-text-muted hover:text-emerald-500 hover:border-emerald-500/40'
                  }`}
                >
                  <CheckCircle2 className="w-3.5 h-3.5" />
                  <span>{isIncluded ? 'Included' : 'Include'}</span>
                </button>
                <button
                  type="button"
                  onClick={() => onToggleExclude(item.pattern)}
                  className={`px-2.5 py-1 rounded-lg text-xs font-extrabold border transition cursor-pointer flex items-center gap-1 ${
                    isExcluded
                      ? 'bg-rose-500 text-white border-rose-500 shadow-xs'
                      : 'bg-bg-input border-border-main text-text-muted hover:text-rose-500 hover:border-rose-500/40'
                  }`}
                >
                  <XCircle className="w-3.5 h-3.5" />
                  <span>{isExcluded ? 'Excluded' : 'Exclude'}</span>
                </button>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};
