'use client';

import React from 'react';
import { useMigrationStore } from '../../store/migrationStore';
import { Database, Cpu, Compass, HardDrive, Sun, Moon, Maximize2, Sliders, Server, Key } from 'lucide-react';

interface AppLayoutProps {
  children: React.ReactNode;
}

export const AppLayout: React.FC<AppLayoutProps> = ({ children }) => {
  const { activeTab, setActiveTab, themeMode, toggleTheme } = useMigrationStore();

  // Synchronize documentElement theme class with themeMode on mount & updates
  React.useEffect(() => {
    const root = window.document.documentElement;
    if (themeMode === 'light') {
      root.classList.add('light');
      root.classList.remove('dark');
    } else {
      root.classList.add('dark');
      root.classList.remove('light');
    }
  }, [themeMode]);

  // Dynamic Page Width Controller State (Default 1920px - User preferred environment width)
  const [layoutWidth, setLayoutWidth] = React.useState<number>(1920);

  const dashboardNavItems = [
    { id: 'overview', label: 'Overview', icon: Server, desc: 'Multi-cloud topology & resource overview' },
    { id: 'credential', label: 'Credential', icon: Key, desc: 'CSP Access credentials & Vault secrets', badge: 'WIP' },
  ] as const;

  const migrationNavItems = [
    { id: 'infra', label: 'Infrastructure', icon: Cpu, desc: 'Host nodes & VM migration pipeline' },
    { id: 'storage', label: 'Object Storage', icon: HardDrive, desc: 'Cloud bucket recommendation & creation', badge: 'WIP' },
    { id: 'data', label: 'Data', icon: Database, desc: 'Encrypted file & object data transfer', badge: 'WIP' },
  ] as const;

  const isTabActive = (itemId: string) => {
    if (itemId === 'infra') return ['infra', 'source', 'refine', 'design', 'migrate'].includes(activeTab);
    if (itemId === 'overview') return ['overview', 'operations'].includes(activeTab);
    return activeTab === itemId;
  };

  return (
    <div className="min-h-screen bg-bg-main text-text-main flex flex-col selection:bg-emerald-500/30 selection:text-emerald-200 overflow-y-scroll">
      {/* Premium Header */}
      <header className="sticky top-0 z-40 w-full border-b border-border-main bg-bg-main/80 backdrop-blur-md">
        <div className="w-full max-w-[1920px] mx-auto px-6 h-16 flex items-center justify-between">
          <div className="flex items-center space-x-3">
            <div className="p-2 bg-gradient-to-tr from-emerald-500 to-teal-500 rounded-lg shadow-lg shadow-emerald-500/20">
              <HardDrive className="w-6 h-6 text-slate-950" />
            </div>
            <div>
              <span className="font-extrabold text-2xl tracking-tight bg-gradient-to-r from-emerald-400 via-teal-400 to-teal-400 bg-clip-text text-transparent">
                Beetle UX Lab
              </span>
              <span className="ml-2.5 text-xs font-semibold px-2.5 py-1 bg-bg-input border border-border-main rounded-full text-text-muted">
                CM-Beetle v0.5.6
              </span>
            </div>
          </div>

          <div className="flex items-center space-x-4">
            {/* Dynamic Width Controller Slider */}
            <div className="flex items-center space-x-2 bg-bg-panel border border-border-main px-2.5 py-1 rounded-xl">
              <Maximize2 className="w-3.5 h-3.5 text-emerald-600 dark:text-emerald-400" />
              <span className="text-xs font-bold text-text-muted font-mono whitespace-nowrap">{layoutWidth}px</span>
              <input
                type="range"
                min="1200"
                max="2560"
                step="40"
                value={layoutWidth}
                onChange={(e) => setLayoutWidth(Number(e.target.value))}
                className="w-24 accent-emerald-500 cursor-pointer"
                title="Adjust Dashboard Layout Container Width"
              />
              <button
                onClick={() => setLayoutWidth(1920)}
                className="text-[10px] font-bold px-1.5 py-0.5 bg-bg-input hover:bg-bg-main text-text-muted hover:text-emerald-500 border border-border-main rounded cursor-pointer transition"
                title="Reset layout width to standard 1920px"
              >
                Reset
              </button>
            </div>

            {/* Light/Dark Theme Mode Toggle Button */}
            <button
              onClick={toggleTheme}
              aria-label="Toggle Theme Mode"
              className="p-2 bg-bg-panel border border-border-main hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-xl transition-all duration-200 text-emerald-600 dark:text-emerald-400 cursor-pointer"
            >
              {themeMode === 'dark' ? <Sun className="w-5 h-5" /> : <Moon className="w-5 h-5" />}
            </button>

            <div className="flex items-center space-x-2 text-xs text-text-muted bg-bg-panel border border-border-main px-3 py-1.5 rounded-full">
              <span className="pulse-dot" />
              <span>Multi-Cloud Engine Connected</span>
            </div>
          </div>
        </div>
      </header>

      {/* Main Tab Navigation bar - Grouped Structure (Dashboard Left, Migration Right) */}
      <div className="bg-bg-panel border-b border-border-main">
        <div className="mx-auto px-6 py-3 flex flex-wrap items-stretch gap-4" style={{ width: `${layoutWidth}px`, maxWidth: '100%' }}>
          
          {/* Group 1: Dashboard (Left Side) */}
          <div className="flex-[2] min-w-[500px] border border-border-main rounded-2xl p-3 bg-bg-main/30 flex flex-col justify-between">
            <div className="flex items-center mb-2 px-1">
              <span className="px-3 py-1 text-xs font-extrabold text-teal-600 dark:text-teal-400 bg-teal-500/10 border border-teal-500/20 rounded-md font-mono">
                Dashboard
              </span>
            </div>
            <nav className="grid grid-cols-1 md:grid-cols-2 gap-3">
              {dashboardNavItems.map((item) => {
                const Icon = item.icon;
                const active = isTabActive(item.id);
                return (
                  <button
                    key={item.id}
                    onClick={() => setActiveTab(item.id as any)}
                    className={`flex items-start p-3.5 rounded-xl text-left border transition-all duration-300 relative overflow-hidden group cursor-pointer ${active
                      ? 'border-teal-500 dark:border-teal-500/40 bg-bg-panel shadow-md shadow-teal-500/10'
                      : 'border-border-main bg-bg-main/20 hover:border-teal-500/30 hover:bg-bg-panel'
                      }`}
                  >
                    {active && (
                      <div className="absolute top-0 left-0 w-full h-[2px] bg-gradient-to-r from-teal-500 via-emerald-400 to-emerald-500" />
                    )}
                    <div className={`p-2.5 rounded-lg mr-3.5 border transition-colors shrink-0 ${active
                      ? 'bg-teal-100 dark:bg-teal-950/40 border-teal-300 dark:border-teal-800/40 text-teal-600 dark:text-teal-400'
                      : 'bg-bg-input border-transparent text-text-muted group-hover:text-teal-600 dark:group-hover:text-teal-400'
                      }`}>
                      <Icon className="w-5 h-5" />
                    </div>
                    <div className="min-w-0">
                      <h3 className={`font-extrabold text-sm transition-colors duration-200 flex items-center space-x-1.5 ${active
                        ? 'text-teal-600 dark:text-teal-400'
                        : 'text-text-main group-hover:text-teal-600 dark:group-hover:text-teal-400'
                        }`}>
                        <span>{item.label}</span>
                        {(item as any).badge && (
                          <span className="px-1.5 py-0.5 text-[10px] font-mono font-extrabold bg-amber-500/10 text-amber-600 dark:text-amber-400 border border-amber-500/20 rounded-md">
                            {(item as any).badge}
                          </span>
                        )}
                      </h3>
                      <p className="text-xs text-text-muted mt-1 leading-snug truncate">
                        {item.desc}
                      </p>
                    </div>
                  </button>
                );
              })}
            </nav>
          </div>

          {/* Group 2: Migration (Right Side) */}
          <div className="flex-[3] min-w-[700px] border border-border-main rounded-2xl p-3 bg-bg-main/30 flex flex-col justify-between">
            <div className="flex items-center mb-2 px-1">
              <span className="px-3 py-1 text-xs font-extrabold text-emerald-600 dark:text-emerald-400 bg-emerald-500/10 border border-emerald-500/20 rounded-md font-mono">
                Migration
              </span>
            </div>
            <nav className="grid grid-cols-1 md:grid-cols-3 gap-3">
              {migrationNavItems.map((item) => {
                const Icon = item.icon;
                const active = isTabActive(item.id);
                return (
                  <button
                    key={item.id}
                    onClick={() => setActiveTab(item.id as any)}
                    className={`flex items-start p-3.5 rounded-xl text-left border transition-all duration-300 relative overflow-hidden group cursor-pointer ${active
                      ? 'border-emerald-500 dark:border-emerald-500/40 bg-bg-panel shadow-md shadow-emerald-500/10'
                      : 'border-border-main bg-bg-main/20 hover:border-emerald-500/30 hover:bg-bg-panel'
                      }`}
                  >
                    {active && (
                      <div className="absolute top-0 left-0 w-full h-[2px] bg-gradient-to-r from-emerald-500 via-teal-400 to-teal-500" />
                    )}
                    <div className={`p-2.5 rounded-lg mr-3.5 border transition-colors shrink-0 ${active
                      ? 'bg-emerald-100 dark:bg-emerald-950/40 border-emerald-300 dark:border-emerald-800/40 text-emerald-600 dark:text-emerald-400'
                      : 'bg-bg-input border-transparent text-text-muted group-hover:text-emerald-600 dark:group-hover:text-emerald-400'
                      }`}>
                      <Icon className="w-5 h-5" />
                    </div>
                    <div className="min-w-0">
                      <h3 className={`font-extrabold text-sm transition-colors duration-200 flex items-center space-x-1.5 ${active
                        ? 'text-emerald-600 dark:text-emerald-400'
                        : 'text-text-main group-hover:text-emerald-600 dark:group-hover:text-emerald-400'
                        }`}>
                        <span>{item.label}</span>
                        {(item as any).badge && (
                          <span className="px-1.5 py-0.5 text-[10px] font-mono font-extrabold bg-amber-500/10 text-amber-600 dark:text-amber-400 border border-amber-500/20 rounded-md">
                            {(item as any).badge}
                          </span>
                        )}
                      </h3>
                      <p className="text-xs text-text-muted mt-1 leading-snug truncate">
                        {item.desc}
                      </p>
                    </div>
                  </button>
                );
              })}
            </nav>
          </div>

        </div>
      </div>

      {/* Content wrapper - Dynamic layout width */}
      <main className="flex-1 mx-auto px-6 py-8" style={{ width: `${layoutWidth}px`, maxWidth: '100%' }}>
        <div className="animate-fade-in">
          {children}
        </div>
      </main>

      {/* Footer */}
      <footer className="border-t border-border-main bg-bg-panel py-6 text-center text-xs text-text-muted">
        <p>© 2026 Cloud-Barista Authors. Licensed under the Apache License, Version 2.0.</p>
      </footer>
    </div>
  );
};
