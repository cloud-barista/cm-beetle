'use client';

import React, { useEffect } from 'react';
import { CheckCircle2, AlertCircle, X } from 'lucide-react';

interface ModelToastProps {
  message: string | null;
  tone?: 'success' | 'error';
  durationMs?: number;
  onClose: () => void;
}

// Unified bottom-right toast notification for model save, update, and delete events.
export const ModelToast: React.FC<ModelToastProps> = ({
  message,
  tone = 'success',
  durationMs = 4000,
  onClose,
}) => {
  useEffect(() => {
    if (!message) return;
    const timer = setTimeout(() => {
      onClose();
    }, durationMs);
    return () => clearTimeout(timer);
  }, [message, durationMs, onClose]);

  if (!message) return null;

  const isError = tone === 'error';

  return (
    <div
      className={`fixed bottom-6 right-6 z-50 flex items-center gap-3 px-5 py-4 rounded-2xl shadow-2xl animate-fade-in text-sm font-extrabold backdrop-blur-md border max-w-lg ${
        isError
          ? 'bg-slate-950/95 border-red-500/40 text-red-400 shadow-red-500/10'
          : 'bg-slate-950/95 border-emerald-500/40 text-emerald-400 shadow-emerald-500/10'
      }`}
    >
      {isError ? (
        <AlertCircle className="w-5 h-5 text-red-500 shrink-0" />
      ) : (
        <CheckCircle2 className="w-5 h-5 text-emerald-500 shrink-0" />
      )}
      <span className="flex-1">{message}</span>
      <button
        type="button"
        onClick={onClose}
        className="p-1 rounded-lg text-slate-400 hover:text-white transition cursor-pointer shrink-0"
        title="Dismiss notification"
      >
        <X className="w-4 h-4" />
      </button>
    </div>
  );
};
