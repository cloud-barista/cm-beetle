'use client';

import React from 'react';
import { useMigrationStore } from '@/store/migrationStore';
import { AppLayout } from '@/components/layout/AppLayout';
import { SourceMetadataExtraction } from '@/components/source/SourceMetadataExtraction';
import { SourceInfraRefinement } from '@/components/refine/SourceInfraRefinement';
import { CloudInfraOptimizer } from '@/components/design/CloudInfraOptimizer';
import { MigrationExecution } from '@/components/center/MigrationExecution';
import { MigratedInfraManagement } from '@/components/operations/MigratedInfraManagement';

export default function Home() {
  const { activeTab } = useMigrationStore();

  return (
    <AppLayout>
      {activeTab === 'source' && <SourceMetadataExtraction />}
      {activeTab === 'refine' && <SourceInfraRefinement />}
      {activeTab === 'design' && <CloudInfraOptimizer />}
      {activeTab === 'migrate' && <MigrationExecution />}
      {activeTab === 'operations' && <MigratedInfraManagement />}
    </AppLayout>
  );
}
