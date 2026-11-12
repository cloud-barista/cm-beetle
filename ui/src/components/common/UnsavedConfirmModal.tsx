'use client';

import React from 'react';
import { AlertTriangle, X } from 'lucide-react';

interface UnsavedConfirmModalProps {
  isOpen: boolean;
  onStay: () => void;
  onProceed: () => void;
  description?: string;
  stayLabel?: string;
  proceedLabel?: string;
}

// Lightweight 2-button confirmation popup shown when advancing or loading with unsaved model changes.
export const UnsavedConfirmModal: React.FC<UnsavedConfirmModalProps> = ({
  isOpen,
  onStay,
  onProceed,
  description = 'Your current model has not been saved yet. If you continue, it will be carried forward as an Unsaved Draft until saved. Would you like to proceed without saving?',
  stayLabel = 'Stay on This Step',
  proceedLabel = 'Proceed without Saving',
}) => {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
      <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up space-y-4">
        <div className="flex justify-between items-center">
          <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
            <AlertTriangle className="w-5 h-5 text-amber-500" />
            <span>Unsaved Model Changes</span>
          </h3>
          <button
            type="button"
            onClick={onStay}
            className="text-text-muted hover:text-text-main transition p-1 hover:bg-bg-input rounded-lg cursor-pointer"
          >
            <X className="w-5 h-5" />
          </button>
        </div>

        <p className="text-sm text-text-muted leading-relaxed">
          {description}
        </p>

        <div className="flex justify-end gap-3 pt-2 text-sm font-extrabold">
          <button
            type="button"
            onClick={onStay}
            className="px-4 py-2.5 bg-bg-panel border border-border-main text-text-main hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-xl transition cursor-pointer"
          >
            {stayLabel}
          </button>
          <button
            type="button"
            onClick={() => {
              onStay();
              onProceed();
            }}
            className="px-4 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white rounded-xl transition cursor-pointer shadow-md shadow-emerald-500/20"
          >
            {proceedLabel}
          </button>
        </div>
      </div>
    </div>
  );
};
