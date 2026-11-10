'use client';

import React, { useEffect, useState } from 'react';
import { Save, X, Loader2, AlertCircle } from 'lucide-react';
import {
  cleanModelDisplayName,
  formatSourceOptionText,
  getStageBadgeMeta,
  ModelStageKind,
  normalizeModelVersion,
} from '../../utils/modelFormat';

export interface SaveRevisionTarget {
  id: string;
  name: string;
  description?: string;
  version?: string;
}

export interface SaveRevisionResult {
  name: string;
  description: string;
  version: string;
  overwriteId: string | null;
}

interface SaveRevisionModalProps {
  isOpen: boolean;
  onClose: () => void;
  title: string;
  stageKind?: Exclude<ModelStageKind, 'sample' | 'draft'>;
  defaultName: string;
  defaultDescription?: string;
  defaultVersion?: string;
  defaultOverwriteId?: string | null;
  existingRevisions: SaveRevisionTarget[];
  onSave: (result: SaveRevisionResult) => Promise<void>;
  successMessage?: string;
}

// Shared save model modal used across Infrastructure and Object Storage workflows.
export const SaveRevisionModal: React.FC<SaveRevisionModalProps> = ({
  isOpen,
  onClose,
  title,
  stageKind,
  defaultName,
  defaultDescription = '',
  defaultVersion = '1.0.0',
  defaultOverwriteId = null,
  existingRevisions,
  onSave,
}) => {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [version, setVersion] = useState('1.0.0');
  const [overwrite, setOverwrite] = useState(false);
  const [overwriteId, setOverwriteId] = useState('');
  const [isSaving, setIsSaving] = useState(false);
  const [error, setError] = useState('');

  const resolvePreferredOverwriteTarget = (): SaveRevisionTarget | undefined => {
    if (defaultOverwriteId) {
      const matched = existingRevisions.find((r) => r.id === defaultOverwriteId);
      if (matched) return matched;
    }
    return existingRevisions[0];
  };

  useEffect(() => {
    if (isOpen) {
      const cleanDefault = defaultName ? cleanModelDisplayName(defaultName, '') : '';
      const preferredTarget = resolvePreferredOverwriteTarget();
      setName(cleanDefault);
      setDescription(defaultDescription);
      setVersion(normalizeModelVersion(defaultVersion, '1.0.0'));
      setOverwrite(false);
      setOverwriteId(preferredTarget?.id || '');
      setIsSaving(false);
      setError('');
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isOpen]);

  if (!isOpen) return null;

  const handleToggleOverwrite = (checked: boolean) => {
    setOverwrite(checked);
    setError('');
    if (checked) {
      const target =
        existingRevisions.find((r) => r.id === overwriteId) || resolvePreferredOverwriteTarget();
      if (target) {
        setOverwriteId(target.id);
        setName(cleanModelDisplayName(target.name, ''));
        setDescription(target.description ?? defaultDescription);
        setVersion(normalizeModelVersion(target.version, defaultVersion || '1.0.0'));
      }
    } else {
      const cleanDefault = defaultName ? cleanModelDisplayName(defaultName, '') : '';
      setName(cleanDefault);
      setDescription(defaultDescription);
      setVersion(normalizeModelVersion(defaultVersion, '1.0.0'));
    }
  };

  const handleSelectOverwriteTarget = (targetId: string) => {
    setOverwriteId(targetId);
    const target = existingRevisions.find((r) => r.id === targetId);
    if (target) {
      setName(cleanModelDisplayName(target.name, ''));
      setDescription(target.description ?? '');
      setVersion(normalizeModelVersion(target.version, '1.0.0'));
    }
  };

  const handleSave = async () => {
    const cleanName = cleanModelDisplayName(name, '');
    if (!cleanName) {
      setError('Model name is required.');
      return;
    }
    if (overwrite && !overwriteId) {
      setError('Select an existing model to overwrite.');
      return;
    }
    setError('');
    setIsSaving(true);
    try {
      await onSave({
        name: cleanName,
        description: description.trim(),
        version: normalizeModelVersion(version, '1.0.0'),
        overwriteId: overwrite ? overwriteId : null,
      });
      onClose();
    } catch (err: any) {
      setError(err?.response?.data?.message || err?.response?.data?.error || err?.message || 'Save failed');
    } finally {
      setIsSaving(false);
    }
  };

  const stageBadge = stageKind ? getStageBadgeMeta(stageKind) : null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4">
      <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up">
        <div className="flex justify-between items-center mb-4">
          <div className="flex items-center gap-2">
            <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
              <Save className="w-4 h-4 text-emerald-500" /> {title}
            </h3>
            {stageBadge && (
              <span className={`px-2 py-0.5 rounded-md text-xs font-extrabold border ${stageBadge.className}`}>
                {stageBadge.label}
              </span>
            )}
          </div>
          <button
            type="button"
            onClick={onClose}
            disabled={isSaving}
            className="text-text-muted hover:text-text-main transition p-1 hover:bg-bg-input rounded-lg cursor-pointer disabled:opacity-50"
          >
            <X className="w-5 h-5" />
          </button>
        </div>

        <div className="space-y-4">
          <div>
            <label className="block text-sm font-normal text-text-muted mb-1.5">Model Name</label>
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="e.g. prod-web-cluster"
              className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2 text-sm font-extrabold focus:outline-none focus:ring-1 focus:ring-emerald-500"
            />
          </div>
          <div>
            <label className="block text-sm font-normal text-text-muted mb-1.5">Description</label>
            <textarea
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              rows={2}
              placeholder="Optional description"
              className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2 text-sm font-extrabold focus:outline-none focus:ring-1 focus:ring-emerald-500 resize-none"
            />
          </div>
          <div>
            <label className="block text-sm font-normal text-text-muted mb-1.5">Version (SemVer)</label>
            <input
              type="text"
              value={version}
              onChange={(e) => setVersion(e.target.value)}
              placeholder="1.0.0"
              className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2 text-sm font-extrabold font-mono focus:outline-none focus:ring-1 focus:ring-emerald-500"
            />
          </div>

          {existingRevisions.length > 0 && (
            <div className="pt-3 border-t border-border-main/30">
              <label className="flex items-center gap-2 text-sm font-normal text-text-muted cursor-pointer">
                <input
                  type="checkbox"
                  checked={overwrite}
                  onChange={(e) => handleToggleOverwrite(e.target.checked)}
                  className="w-4 h-4 rounded border-border-main accent-emerald-500 cursor-pointer"
                />
                <span>Overwrite an existing model</span>
              </label>
              {overwrite && (
                <select
                  value={overwriteId}
                  onChange={(e) => handleSelectOverwriteTarget(e.target.value)}
                  className="w-full mt-2 bg-bg-input border border-border-main text-text-main rounded-xl px-3 py-2 text-sm font-extrabold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
                >
                  {existingRevisions.map((r) => (
                    <option key={r.id} value={r.id}>
                      {formatSourceOptionText(r, '1.0.0')}
                    </option>
                  ))}
                </select>
              )}
            </div>
          )}

          {error && (
            <div className="p-2.5 bg-red-500/10 border border-red-500/20 text-red-500 dark:text-red-400 rounded-xl text-sm flex items-center justify-center gap-1.5">
              <AlertCircle className="w-4 h-4 shrink-0" />
              <span>{error}</span>
            </div>
          )}
        </div>

        <div className="mt-6 flex justify-end gap-3 text-sm font-extrabold">
          <button
            type="button"
            onClick={onClose}
            disabled={isSaving}
            className="px-4 py-2.5 bg-bg-panel border border-border-main text-text-main hover:bg-emerald-500/10 hover:border-emerald-500/20 rounded-xl cursor-pointer disabled:opacity-50"
          >
            Cancel
          </button>
          <button
            type="button"
            onClick={handleSave}
            disabled={isSaving}
            className="px-4 py-2.5 bg-gradient-to-r from-emerald-400 via-teal-400 to-teal-500 hover:from-emerald-500 hover:to-teal-600 text-slate-950 rounded-xl cursor-pointer disabled:opacity-50 flex items-center gap-1.5 shadow-sm"
          >
            {isSaving ? (
              <>
                <Loader2 className="w-4 h-4 animate-spin" />
                <span>Saving...</span>
              </>
            ) : overwrite ? (
              'Overwrite Model'
            ) : (
              'Save Model'
            )}
          </button>
        </div>
      </div>
    </div>
  );
};
