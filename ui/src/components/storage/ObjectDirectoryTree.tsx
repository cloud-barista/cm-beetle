'use client';

import React, { useState, useMemo } from 'react';
import {
  Folder,
  FolderOpen,
  FileText,
  ChevronRight,
  ChevronDown,
  Search,
  Maximize2,
  Minimize2,
  HardDrive
} from 'lucide-react';
import { buildObjectTree, ObjectNode, formatBytes } from '@/utils/objectTree';

export interface ObjectDirectoryTreeProps {
  objects: Array<{ key?: string; bucketName?: string; name?: string; totalSizeBytes?: number; sizeBytes?: number; lastModified?: string; creationDate?: string }>;
  bucketName?: string;
  className?: string;
}

export const ObjectDirectoryTree: React.FC<ObjectDirectoryTreeProps> = ({
  objects,
  bucketName = 'bucket',
  className = ''
}) => {
  const [searchTerm, setSearchTerm] = useState('');
  const [expandedNodes, setExpandedNodes] = useState<Record<string, boolean>>({});

  const tree = useMemo(() => buildObjectTree(objects), [objects]);

  const toggleNode = (nodeId: string) => {
    setExpandedNodes((prev) => ({ ...prev, [nodeId]: !prev[nodeId] }));
  };

  const expandAll = () => {
    const newExpanded: Record<string, boolean> = {};
    const traverse = (nodes: ObjectNode[]) => {
      nodes.forEach((n) => {
        if (n.isFolder) {
          newExpanded[n.id] = true;
          traverse(n.children);
        }
      });
    };
    traverse(tree);
    setExpandedNodes(newExpanded);
  };

  const collapseAll = () => {
    setExpandedNodes({});
  };

  // Filter nodes recursively if search term is provided
  const filterNodes = (nodes: ObjectNode[], term: string): ObjectNode[] => {
    if (!term.trim()) return nodes;
    const lower = term.toLowerCase();

    return nodes.reduce((acc: ObjectNode[], node) => {
      const nameMatches = node.name.toLowerCase().includes(lower) || node.fullPath.toLowerCase().includes(lower);
      if (node.isFolder) {
        const filteredChildren = filterNodes(node.children, term);
        if (nameMatches || filteredChildren.length > 0) {
          acc.push({ ...node, children: filteredChildren });
        }
      } else if (nameMatches) {
        acc.push(node);
      }
      return acc;
    }, []);
  };

  const displayTree = useMemo(() => filterNodes(tree, searchTerm), [tree, searchTerm]);

  const renderNode = (node: ObjectNode, depth: number = 0) => {
    const isExpanded = expandedNodes[node.id] || Boolean(searchTerm.trim());

    if (node.isFolder) {
      return (
        <div key={node.id} className="select-none font-mono text-xs">
          <div
            onClick={() => toggleNode(node.id)}
            className="flex items-center justify-between py-1.5 px-2.5 rounded-lg hover:bg-emerald-500/10 cursor-pointer transition text-text-main group"
            style={{ paddingLeft: `${depth * 1.25 + 0.5}rem` }}
          >
            <div className="flex items-center space-x-2 min-w-0">
              <button type="button" className="text-text-muted hover:text-emerald-500 transition shrink-0">
                {isExpanded ? <ChevronDown className="w-3.5 h-3.5 text-emerald-500" /> : <ChevronRight className="w-3.5 h-3.5 text-text-muted" />}
              </button>
              {isExpanded ? (
                <FolderOpen className="w-4 h-4 text-emerald-500 shrink-0" />
              ) : (
                <Folder className="w-4 h-4 text-emerald-600 dark:text-emerald-400 shrink-0" />
              )}
              <span className="font-extrabold text-text-main truncate text-xs group-hover:text-emerald-600 dark:group-hover:text-emerald-400">
                {node.name}
              </span>
              <span className="text-[11px] text-text-muted font-normal shrink-0">
                ({node.objectCount.toLocaleString()} {node.objectCount === 1 ? 'item' : 'items'})
              </span>
            </div>

            <div className="flex items-center space-x-3 text-[11px] text-text-muted shrink-0 pl-2">
              <span className="font-semibold text-text-main">{formatBytes(node.sizeBytes)}</span>
            </div>
          </div>

          {isExpanded && node.children.length > 0 && (
            <div className="border-l border-border-main/30 ml-4">
              {node.children.map((child) => renderNode(child, depth + 1))}
            </div>
          )}
        </div>
      );
    }

    // File Node
    return (
      <div
        key={node.id}
        className="flex items-center justify-between py-1.5 px-2.5 rounded-lg hover:bg-bg-input/60 transition text-text-main text-xs font-mono select-text"
        style={{ paddingLeft: `${depth * 1.25 + 1.75}rem` }}
      >
        <div className="flex items-center space-x-2 min-w-0">
          <FileText className="w-3.5 h-3.5 text-teal-500 shrink-0" />
          <span className="font-bold text-text-main truncate text-xs hover:text-emerald-500 cursor-pointer" title={node.fullPath}>
            {node.name}
          </span>
        </div>

        <div className="flex items-center space-x-3 text-[11px] text-text-muted shrink-0 pl-2">
          {node.lastModified && <span className="text-text-muted/70 hidden sm:inline">{node.lastModified}</span>}
          <span className="font-mono font-extrabold text-emerald-600 dark:text-emerald-400">{formatBytes(node.sizeBytes)}</span>
        </div>
      </div>
    );
  };

  return (
    <div className={`bg-bg-panel border border-border-main rounded-2xl p-4 space-y-3 font-sans shadow-sm ${className}`}>
      {/* Top Header Bar & Search Controls */}
      <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-2 pb-3 border-b border-border-main/40">
        <div className="flex items-center space-x-2">
          <HardDrive className="w-4.5 h-4.5 text-emerald-500 shrink-0" />
          <span className="font-extrabold text-sm text-text-main font-mono">
            {bucketName}
          </span>
          <span className="text-xs text-text-muted font-normal font-mono bg-bg-input px-2 py-0.5 rounded border border-border-main/50">
            {objects.length.toLocaleString()} total objects
          </span>
        </div>

        <div className="flex items-center space-x-2">
          <div className="relative flex-1 sm:w-64">
            <Search className="w-3.5 h-3.5 absolute left-3 top-1/2 -translate-y-1/2 text-text-muted" />
            <input
              type="text"
              placeholder="Search object path..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="w-full pl-8 pr-3 py-1.5 bg-bg-input border border-border-main rounded-xl text-text-main font-mono text-xs focus:outline-none focus:border-emerald-500"
            />
          </div>

          <button
            type="button"
            onClick={expandAll}
            className="p-1.5 bg-bg-input hover:bg-bg-main border border-border-main text-text-muted hover:text-emerald-500 rounded-lg transition cursor-pointer"
            title="Expand All Directories"
          >
            <Maximize2 className="w-3.5 h-3.5" />
          </button>
          <button
            type="button"
            onClick={collapseAll}
            className="p-1.5 bg-bg-input hover:bg-bg-main border border-border-main text-text-muted hover:text-emerald-500 rounded-lg transition cursor-pointer"
            title="Collapse All Directories"
          >
            <Minimize2 className="w-3.5 h-3.5" />
          </button>
        </div>
      </div>

      {/* Directory Tree View Container */}
      <div className="max-h-[26rem] overflow-y-auto pr-1 space-y-0.5">
        {displayTree.length === 0 ? (
          <div className="py-8 text-center text-xs text-text-muted font-mono italic">
            {searchTerm ? `No objects match search "${searchTerm}"` : 'No objects in bucket'}
          </div>
        ) : (
          displayTree.map((node) => renderNode(node, 0))
        )}
      </div>
    </div>
  );
};
