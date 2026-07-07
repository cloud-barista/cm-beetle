import React from 'react';
import { useMigrationStore } from './store/migrationStore';
import { AppLayout } from './components/layout/AppLayout';
import { SourceCenter } from './components/source/SourceCenter';
import { MigrationDesigner } from './components/design/MigrationDesigner';
import { MigrationCenter } from './components/center/MigrationCenter';

const App: React.FC = () => {
  const { activeTab } = useMigrationStore();

  return (
    <AppLayout>
      {activeTab === 'source' && <SourceCenter />}
      {activeTab === 'design' && <MigrationDesigner />}
      {activeTab === 'migrate' && <MigrationCenter />}
    </AppLayout>
  );
};

export default App;
