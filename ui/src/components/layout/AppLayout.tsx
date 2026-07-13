'use client';

import React from 'react';
import { useMigrationStore } from '../../store/migrationStore';
import { Database, Cpu, Compass, HardDrive, Sun, Moon, Maximize2 } from 'lucide-react';

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

  const navItems = [
    { id: 'source', label: '1. Source Infrastructures', icon: Database, desc: 'Register host connections & extract system models' },
    { id: 'design', label: '2. Target Cloud Optimizer', icon: Compass, desc: 'Compose & compare multiple cloud architecture combinations' },
    { id: 'migrate', label: '3. Migration Execution', icon: Cpu, desc: 'Trigger deployment & monitor build logs' },
  ] as const;

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
                CM-Beetle v0.5.4
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
                min={1500}
                max={2500}
                step={1}
                value={layoutWidth}
                onChange={(e) => setLayoutWidth(parseInt(e.target.value, 10))}
                className="w-24 h-1 bg-bg-input rounded-lg appearance-none cursor-pointer accent-emerald-500 focus:outline-none"
                title="Adjust Page layout width dynamically"
              />
              <button
                onClick={() => setLayoutWidth(1920)}
                className="p-1 px-2 text-[11px] font-bold bg-bg-input hover:bg-emerald-500/10 border border-border-main hover:text-emerald-600 dark:hover:text-emerald-400 rounded-md text-text-muted cursor-pointer transition whitespace-nowrap"
                title="Reset to default 1920px"
              >
                Reset
              </button>
            </div>

            {/* 1. Light/Dark Theme Mode Toggle Button */}
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

      {/* Main Tab Navigation bar */}
      <div className="bg-bg-panel border-b border-border-main">
        <div className="mx-auto px-6 py-4" style={{ width: `${layoutWidth}px`, maxWidth: '100%' }}>
          <nav className="grid grid-cols-1 md:grid-cols-3 gap-3">
            {navItems.map((item) => {
              const Icon = item.icon;
              const isActive = activeTab === item.id;
              return (
                <button
                  key={item.id}
                  onClick={() => setActiveTab(item.id)}
                  className={`flex items-start p-4 rounded-xl text-left border transition-all duration-300 relative overflow-hidden group cursor-pointer ${isActive
                    ? 'border-emerald-500 dark:border-emerald-500/40 bg-bg-panel shadow-lg shadow-emerald-500/10'
                    : 'border-border-main bg-bg-main/20 hover:border-emerald-500/30 hover:bg-bg-panel'
                    }`}
                >
                  {/* Glowing border glow-effect for active tab */}
                  {isActive && (
                    <div className="absolute top-0 left-0 w-full h-[2px] bg-gradient-to-r from-emerald-500 via-teal-400 to-teal-500" />
                  )}
                  <div className={`p-2.5 rounded-lg mr-4 border transition-colors ${isActive
                    ? 'bg-emerald-100 dark:bg-emerald-950/40 border-emerald-300 dark:border-emerald-800/40 text-emerald-600 dark:text-emerald-400'
                    : 'bg-bg-input border-transparent text-text-muted group-hover:text-emerald-600 dark:group-hover:text-emerald-400'
                    }`}>
                    <Icon className="w-5 h-5" />
                  </div>
                  <div>
                    <h3 className={`font-bold text-base transition-colors duration-200 ${isActive
                      ? 'text-emerald-600 dark:text-emerald-400 font-extrabold'
                      : 'text-text-muted group-hover:text-emerald-600 dark:group-hover:text-emerald-400'
                      }`}>
                      {item.label}
                    </h3>
                    <p className="text-xs text-text-muted mt-1 leading-relaxed">
                      {item.desc}
                    </p>
                  </div>
                </button>
              );
            })}
          </nav>
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
