'use client';

import React from 'react';
import { useMigrationStore } from '@/store/migrationStore';
import { AppLayout } from '@/components/layout/AppLayout';
import { InfrastructureMigration } from '@/components/infra/InfrastructureMigration';
import { ObjectStorageMigration } from '@/components/storage/ObjectStorageMigration';
import { DataTransferCenter } from '@/components/data/DataTransferCenter';
import { CredentialManagement } from '@/components/credential/CredentialManagement';
import { MigratedInfraManagement } from '@/components/operations/MigratedInfraManagement';

export default function Home() {
  const { activeTab } = useMigrationStore();

  const isInfraTab = ['infra', 'source', 'refine', 'design', 'migrate'].includes(activeTab);
  const isOverviewTab = ['overview', 'operations'].includes(activeTab);

  return (
    <AppLayout>
      {isInfraTab && <InfrastructureMigration />}
      {activeTab === 'storage' && <ObjectStorageMigration />}
      {activeTab === 'data' && <DataTransferCenter />}
      {activeTab === 'credential' && <CredentialManagement />}
      {isOverviewTab && <MigratedInfraManagement />}
    </AppLayout>
  );
}
