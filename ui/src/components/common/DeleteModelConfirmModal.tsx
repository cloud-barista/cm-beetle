'use client';

import React, { useEffect, useState } from 'react';
import { Trash2, X, Loader2 } from 'lucide-react';
import { cleanModelDisplayName } from '../../utils/modelFormat';

interface DeleteModelConfirmModalProps {
  isOpen: boolean;
  onClose: () => void;
  title?: string;
  modelName: string;
  descriptionText?: string;
  onConfirm: () => Promise<void>;
}

// Type-to-confirm deletion modal shared across Infrastructure and Object Storage model selectors.
export const DeleteModelConfirmModal: React.FC<DeleteModelConfirmModalProps> = ({
  isOpen,
  onClose,
  title = 'Delete Model',
  modelName,
  descriptionText,
  onConfirm,
}) => {
  const [confirmText, setConfirmText] = useState('');
  const [isDeleting, setIsDeleting] = useState(false);
  const [error, setError] = useState('');

  const expectedText = cleanModelDisplayName(modelName, modelName);

  useEffect(() => {
    if (isOpen) {
      setConfirmText('');
      setIsDeleting(false);
      setError('');
    }
  }, [isOpen, expectedText]);

  if (!isOpen) return null;

  const isMatch = confirmText === expectedText;

  const handleConfirm = async () => {
    if (!isMatch) return;
    setIsDeleting(true);
    setError('');
    try {
      await onConfirm();
      onClose();
    } catch (err: any) {
      setError(err?.response?.data?.message || err?.response?.data?.error || err?.message || 'Failed to delete model');
    } finally {
      setIsDeleting(false);
    }
  };

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
      <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
            <Trash2 className="w-4 h-4 text-red-500" /> {title}
          </h3>
          <button
            type="button"
            onClick={onClose}
            disabled={isDeleting}
            className="text-text-muted hover:text-text-main transition p-1 hover:bg-bg-input rounded-lg cursor-pointer disabled:opacity-50"
          >
            <X className="w-5 h-5" />
          </button>
        </div>

        <div className="space-y-4">
          <p className="text-sm text-text-muted leading-relaxed">
            {descriptionText || (
              <>
                Are you sure you want to delete the model{' '}
                <strong className="font-extrabold text-text-main">&quot;{expectedText}&quot;</strong>? This action cannot be undone.
              </>
            )}
          </p>

          <div className="space-y-1.5 pt-1">
            <label className="block text-sm font-normal text-text-muted">
              To confirm, type{' '}
              <span className="font-mono font-extrabold bg-bg-panel px-1.5 py-0.5 rounded border border-border-main/60 text-text-main select-all">
                {expectedText}
              </span>{' '}
              below:
            </label>
            <input
              type="text"
              value={confirmText}
              onChange={(e) => setConfirmText(e.target.value)}
              placeholder="Type to confirm deletion"
              disabled={isDeleting}
              className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm font-extrabold font-mono focus:outline-none focus:ring-1 focus:ring-red-500"
            />
          </div>

          {error && (
            <div className="flex items-center gap-2 bg-red-500/10 text-red-500 px-4 py-3 rounded-xl text-sm border border-red-500/20">
              <span>{error}</span>
            </div>
          )}

          <div className="flex justify-end gap-3 pt-2 text-sm font-extrabold">
            <button
              type="button"
              onClick={onClose}
              disabled={isDeleting}
              className="px-4 py-2 bg-bg-panel border border-border-main text-text-main rounded-xl hover:bg-bg-input transition cursor-pointer disabled:opacity-50"
            >
              Cancel
            </button>
            <button
              type="button"
              onClick={handleConfirm}
              disabled={isDeleting || !isMatch}
              className={`px-4 py-2 rounded-xl transition flex items-center gap-1.5 ${
                isDeleting || !isMatch
                  ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                  : 'bg-red-500 hover:bg-red-600 text-white cursor-pointer shadow-md shadow-red-500/20'
              }`}
            >
              {isDeleting && <Loader2 className="w-4 h-4 animate-spin" />}
              <span>Confirm Delete</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};
