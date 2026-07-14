'use client';

import React, { useEffect, useState } from 'react';
import { Save, X, CheckCircle2, Loader2, AlertCircle } from 'lucide-react';

export interface SaveRevisionTarget {
  id: string;
  name: string;
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
  defaultName: string;
  defaultDescription?: string;
  defaultVersion?: string;
  existingRevisions: SaveRevisionTarget[];
  onSave: (result: SaveRevisionResult) => Promise<void>;
  successMessage?: string;
}

// Shared Damselfly "save revision" popup used by Source Infra Metadata Extraction,
// Source Infra Metadata Refinement, and Target Cloud Infra Optimization pages.
export const SaveRevisionModal: React.FC<SaveRevisionModalProps> = ({
  isOpen,
  onClose,
  title,
  defaultName,
  defaultDescription = '',
  defaultVersion = '1.0.0',
  existingRevisions,
  onSave,
  successMessage = 'Revision saved to Damselfly successfully.',
}) => {
  const [name, setName] = useState(defaultName);
  const [description, setDescription] = useState(defaultDescription);
  const [version, setVersion] = useState(defaultVersion);
  const [overwrite, setOverwrite] = useState(false);
  const [overwriteId, setOverwriteId] = useState(existingRevisions[0]?.id || '');
  const [isSaving, setIsSaving] = useState(false);
  const [saveSuccess, setSaveSuccess] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    if (isOpen) {
      setName(defaultName);
      setDescription(defaultDescription);
      setVersion(defaultVersion);
      setOverwrite(false);
      setOverwriteId(existingRevisions[0]?.id || '');
      setIsSaving(false);
      setSaveSuccess(false);
      setError('');
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isOpen]);

  if (!isOpen) return null;

  const handleSave = async () => {
    if (!name.trim()) { setError('Name required'); return; }
    if (overwrite && !overwriteId) { setError('Select a revision to overwrite'); return; }
    setError('');
    setIsSaving(true);
    try {
      await onSave({
        name: name.trim(),
        description: description.trim(),
        version: version.trim() || '1.0.0',
        overwriteId: overwrite ? overwriteId : null,
      });
      setSaveSuccess(true);
      setTimeout(() => { onClose(); }, 1200);
    } catch (err: any) {
      setError(err?.response?.data?.message || err?.message || 'Save failed');
    } finally {
      setIsSaving(false);
    }
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4">
      <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-base font-bold text-text-main flex items-center gap-2">
            <Save className="w-4 h-4 text-emerald-400" /> {title}
          </h3>
          <button
            onClick={onClose}
            disabled={isSaving}
            className="text-text-muted hover:text-text-main transition p-1 hover:bg-bg-input rounded-lg cursor-pointer disabled:opacity-50"
          >
            <X className="w-5 h-5" />
          </button>
        </div>

        <div className="space-y-4">
          <div>
            <label className="block text-sm font-semibold text-text-muted mb-1.5">Name</label>
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-emerald-500"
            />
          </div>
          <div>
            <label className="block text-sm font-semibold text-text-muted mb-1.5">Description</label>
            <textarea
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              rows={2}
              className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-emerald-500 resize-none"
            />
          </div>
          <div>
            <label className="block text-sm font-semibold text-text-muted mb-1.5">Version (SemVer)</label>
            <input
              type="text"
              value={version}
              onChange={(e) => setVersion(e.target.value)}
              placeholder="1.0.0"
              className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2 text-sm font-mono focus:outline-none focus:ring-1 focus:ring-emerald-500"
            />
          </div>

          {existingRevisions.length > 0 && (
            <div className="pt-3 border-t border-border-main/30">
              <label className="flex items-center gap-2 text-sm font-semibold text-text-muted cursor-pointer">
                <input
                  type="checkbox"
                  checked={overwrite}
                  onChange={(e) => setOverwrite(e.target.checked)}
                  className="w-4 h-4 rounded border-border-main accent-emerald-500 cursor-pointer"
                />
                Overwrite an existing revision (PUT)
              </label>
              {overwrite && (
                <select
                  value={overwriteId}
                  onChange={(e) => setOverwriteId(e.target.value)}
                  className="w-full mt-2 bg-bg-input border border-border-main text-text-main rounded-xl px-3 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
                >
                  {existingRevisions.map((r) => (
                    <option key={r.id} value={r.id}>{r.name}{r.version ? ` (v${r.version})` : ''}</option>
                  ))}
                </select>
              )}
            </div>
          )}

          {error && (
            <div className="p-2 bg-red-500/10 border border-red-500/20 text-red-400 rounded-xl text-sm text-center flex items-center justify-center gap-1.5">
              <AlertCircle className="w-4 h-4" /> {error}
            </div>
          )}
          {saveSuccess && (
            <div className="p-2 bg-emerald-500/10 border border-emerald-500/20 text-emerald-400 rounded-xl text-sm text-center font-medium flex items-center justify-center gap-1.5">
              <CheckCircle2 className="w-4 h-4" /> {successMessage}
            </div>
          )}
        </div>

        <div className="mt-6 flex justify-end gap-3 text-sm font-bold">
          <button
            onClick={onClose}
            disabled={isSaving}
            className="px-4 py-2.5 bg-bg-panel border border-border-main text-text-main hover:bg-emerald-500/10 hover:border-emerald-500/20 rounded-xl cursor-pointer disabled:opacity-50"
          >
            Cancel
          </button>
          <button
            onClick={handleSave}
            disabled={isSaving}
            className="px-4 py-2.5 bg-gradient-to-r from-emerald-500 to-sky-500 text-slate-950 rounded-xl cursor-pointer disabled:opacity-50 flex items-center gap-1.5"
          >
            {isSaving ? <><Loader2 className="w-4 h-4 animate-spin" /> Saving…</> : (overwrite ? 'Overwrite' : 'Save New')}
          </button>
        </div>
      </div>
    </div>
  );
};
