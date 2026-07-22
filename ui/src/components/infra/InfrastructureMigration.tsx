'use client';

import React, { useState } from 'react';
import { useMigrationStore } from '@/store/migrationStore';
import { SourceMetadataExtraction } from '@/components/source/SourceMetadataExtraction';
import { SourceInfraRefinement } from '@/components/refine/SourceInfraRefinement';
import { CloudInfraOptimizer } from '@/components/design/CloudInfraOptimizer';
import { MigrationExecution } from '@/components/center/MigrationExecution';
import { Database, Sliders, Compass, Cpu, ArrowRight, ArrowLeft } from 'lucide-react';

export const InfrastructureMigration: React.FC = () => {
  const { activeTab } = useMigrationStore();
  const [subTab, setSubTab] = useState<'source' | 'refine' | 'design' | 'migrate'>(() => {
    if (['source', 'refine', 'design', 'migrate'].includes(activeTab)) {
      return activeTab as any;
    }
    return 'source';
  });

  const subSteps = [
    { id: 'source', label: '1. Source Analysis', icon: Database, desc: 'Register hosts & extract metadata' },
    { id: 'refine', label: '2. Refinement', icon: Sliders, desc: 'Review & refine source specs' },
    { id: 'design', label: '3. Target Infra Optimization', icon: Compass, desc: 'Customize target cloud model' },
    { id: 'migrate', label: '4. Migration Execution', icon: Cpu, desc: 'Execute VM migration & check status' },
  ] as const;

  return (
    <div className="space-y-6">
      {/* Unified Workflow Container Box */}
      <div className="bg-bg-panel border border-border-main rounded-2xl p-4 shadow-sm space-y-3">
        {/* Row 1: Workflow Title Line */}
        <div className="flex items-center space-x-2.5 border-b border-border-main pb-3 px-1">
          <Cpu className="w-5 h-5 text-emerald-500" />
          <h2 className="text-base font-extrabold text-text-main">Infrastructure Migration Workflow</h2>
        </div>

        {/* Row 2: Workflow Tab Cards */}
        <nav className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-2.5">
          {subSteps.map((step) => {
            const Icon = step.icon;
            const isActive = subTab === step.id;
            return (
              <button
                key={step.id}
                onClick={() => setSubTab(step.id)}
                className={`flex items-center p-3.5 rounded-xl border text-left transition-all duration-200 cursor-pointer ${
                  isActive
                    ? 'border-emerald-500 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 font-extrabold shadow-sm'
                    : 'border-border-main bg-bg-main/20 text-text-muted hover:border-emerald-500/30 hover:bg-bg-panel'
                }`}
              >
                <Icon className={`w-5 h-5 mr-3 shrink-0 ${isActive ? 'text-emerald-500' : 'text-text-muted'}`} />
                <div className="min-w-0">
                  <div className="text-sm font-bold truncate">{step.label}</div>
                  <div className="text-xs text-text-muted truncate mt-0.5">{step.desc}</div>
                </div>
              </button>
            );
          })}
        </nav>
      </div>

      {/* Sub-step View Render */}
      <div className="animate-fade-in">
        {subTab === 'source' && <SourceMetadataExtraction onNext={() => setSubTab('refine')} />}
        {subTab === 'refine' && <SourceInfraRefinement onNext={() => setSubTab('design')} onBack={() => setSubTab('source')} />}
        {subTab === 'design' && <CloudInfraOptimizer onNext={() => setSubTab('migrate')} onBack={() => setSubTab('refine')} />}
        {subTab === 'migrate' && <MigrationExecution onBack={() => setSubTab('design')} />}
      </div>
    </div>
  );
};
