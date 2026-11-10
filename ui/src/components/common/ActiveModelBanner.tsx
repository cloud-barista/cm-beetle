'use client';

import React from 'react';
import {
  cleanModelDisplayName,
  getStageBadgeMeta,
  ModelStageKind,
  normalizeModelVersion,
} from '../../utils/modelFormat';

export interface ActiveModelBannerMetric {
  label: string;
  value: string;
  tone?: 'default' | 'emerald' | 'teal';
}

interface ActiveModelBannerProps {
  mode: 'loaded' | 'draft';
  name: string;
  version?: string;
  stageKind: ModelStageKind;
  hasUnsavedChanges?: boolean;
  metrics?: ActiveModelBannerMetric[];
  updatedTime?: string;
}

// Shared status strip showing the currently loaded or draft model across workflow steps.
export const ActiveModelBanner: React.FC<ActiveModelBannerProps> = ({
  mode,
  name,
  version,
  stageKind,
  hasUnsavedChanges = false,
  metrics = [],
  updatedTime,
}) => {
  const cleanName = cleanModelDisplayName(name, mode === 'draft' ? 'Unsaved Model' : 'Unnamed Model');
  const badge = getStageBadgeMeta(stageKind);
  const normalizedVersion = version ? normalizeModelVersion(version, '1.0') : '';

  return (
    <div className="bg-bg-panel/40 border border-border-main/40 rounded-xl p-3.5 flex flex-col md:flex-row md:items-center justify-between text-sm gap-2.5">
      <div className="flex flex-wrap items-center gap-2">
        <span className="font-normal text-text-muted">
          {mode === 'draft' ? 'Current Draft:' : 'Loaded Model:'}
        </span>
        <span className="font-extrabold text-emerald-600 dark:text-emerald-400 text-sm">
          {cleanName}
        </span>
        {normalizedVersion && (
          <span className="text-xs font-extrabold text-text-main font-mono bg-bg-panel px-2 py-0.5 rounded border border-border-main/50">
            v{normalizedVersion}
          </span>
        )}
        <span className={`px-2.5 py-0.5 rounded-md text-xs font-extrabold border ${badge.className}`}>
          {badge.label}
        </span>
        {hasUnsavedChanges && (
          <span className="px-2.5 py-0.5 rounded-md text-xs font-extrabold bg-amber-500/15 text-amber-600 dark:text-amber-400 border border-amber-500/30">
            Unsaved
          </span>
        )}
        {metrics.map((m, idx) => {
          const boxTone =
            m.tone === 'emerald'
              ? 'bg-emerald-500/15 border-emerald-500/30'
              : m.tone === 'teal'
                ? 'bg-teal-500/15 border-teal-500/30'
                : 'bg-bg-panel border-border-main/50';
          const valTone =
            m.tone === 'emerald'
              ? 'font-extrabold text-emerald-600 dark:text-emerald-400'
              : m.tone === 'teal'
                ? 'font-extrabold text-teal-600 dark:text-teal-400'
                : 'font-extrabold text-text-main';
          return (
            <span key={idx} className={`px-2.5 py-0.5 rounded-md text-xs border ${boxTone}`}>
              <span className="font-normal text-text-muted">{m.label}: </span>
              <span className={valTone}>{m.value}</span>
            </span>
          );
        })}
      </div>
      {updatedTime && (
        <div className="text-xs shrink-0">
          <span className="font-normal text-text-muted">Updated: </span>
          <span className="font-extrabold text-text-main">
            {Number.isNaN(new Date(updatedTime).getTime())
              ? updatedTime
              : new Date(updatedTime).toLocaleString()}
          </span>
        </div>
      )}
    </div>
  );
};
