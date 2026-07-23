---
name: ui-development
description: Guidelines and instructions for developing the Beetle UX Lab (Next.js 15, React 19, Tailwind CSS, Zustand) in the ui/ directory.
---

# Beetle UX Lab Instructions

## Quick Reference

**📘 Full Design System:** Read [DESIGN_SYSTEM.md](file:///home/ubuntu/dev/cloud-barista/cm-beetle/ui/DESIGN_SYSTEM.md) for comprehensive guidelines on branding, colors, components, and tech stack.

## Essential Rules

### Branding & Colors

- **Brand Name:** Beetle UX Lab, CM-Beetle (documentation/API)
- **Color Palette:** Emerald (#10b981) / Teal (#14b8a6) **only**
- **Forbidden Colors:** ❌ `cyan-*`, `purple-*`, `sky-*` (removed from codebase)
- **Gradients:** `from-emerald-400 via-teal-400 to-teal-400`

### Tech Stack

- **Framework:** Next.js 15 App Router with standalone output
- **React:** v19 with `'use client'` for interactive components
- **Styling:** Tailwind CSS 4 (@theme inline configuration)
- **State:** Zustand v5 for global state
- **Icons:** Lucide React (HardDrive, Database, Cpu, Compass, etc.)

### Typography & Key-Value Rules

- **Key (Labels / Field Names):** Always standard weight (`font-normal text-text-muted`). Never bold.
- **Value (Data / Metrics):** Always bold (`font-extrabold text-text-main`).
- **Minimum Font Size:** Main labels and data values must be at least `text-sm` (16px). Avoid `text-[9px]`, `text-[10px]`, `text-[11px]`.
- **Text Casing:** Title Case only — never `ALL CAPS` on static labels.
  - ✅ `Node Spec`, `Node Image`, `Root Disk`, `Security Group(s)`, `Node(s)`
  - ❌ `NODE SPEC`, `NODE IMAGE`, `ROOT DISK`, `SECURITY GROUP`
- **Tailwind:** Do **not** use `uppercase` + `tracking-wider` on static labels. Reserve `uppercase` for dynamic value badges (e.g., `{status}`, `{protocol}`).

### Project Terminology Standards

- **Nodes / Node Groups:** Use `Node` or `Node Group` across all UI labels and cards (never `VM Instance` or `Active Instances`).
- **Step 2 Tab Name:** `2. Target Cloud Optimizer` (never `Cloud Target Optimizer`).
- **Spec Header:** `onpremiseInfraModel` (omit `Spec` postfix).

### Component Patterns

**Glass Panel:**

```tsx
<div className="bg-bg-panel backdrop-blur-sm border border-border-main rounded-xl p-6">
  {/* Content */}
</div>
```

**Primary Button:**

```tsx
<button className="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white rounded-lg font-bold">
  Action
</button>
```

**Icon Usage:**

```tsx
import { Database } from "lucide-react";

<Database className="w-5 h-5 text-emerald-500" />;
```

## Flexible Topology Visual Patterns (UX Scenario Reference)

> **Note:** The following visual structures are recommended reference patterns for topology visualization. They should be adapted flexibly if UX scenarios, user flows, or design layouts evolve.

- **NLB Traffic Flow Pattern:**
  - Ingress traffic: `Traffic Ingress ➔ Listener Port: {port}`
  - Target routing: `Target NodeGroup: {nodeGroupName}`
  - Destination node tree: Display specific target nodes (`nodeGroup-01`, `nodeGroup-02`) and ports (`Port: 8086`) with tree connectors (`├─`, `└─`).
- **Node Group Details Pattern:**
  - 3-column responsive spec grid (`Node Spec`, `Node Image`, `Root Disk`) using `grid-cols-1 md:grid-cols-3`.
  - Associated Security Group(s) list filtered to display matching security groups per Node Group.

## File Organization

- [ui/src/app/layout.tsx](file:///home/ubuntu/dev/cloud-barista/cm-beetle/ui/src/app/layout.tsx) # Root layout
- [ui/src/app/page.tsx](file:///home/ubuntu/dev/cloud-barista/cm-beetle/ui/src/app/page.tsx) # Main entry with tab routing
- [ui/src/app/globals.css](file:///home/ubuntu/dev/cloud-barista/cm-beetle/ui/src/app/globals.css) # Tailwind theme (emerald/teal)
- [ui/src/components/layout/AppLayout.tsx](file:///home/ubuntu/dev/cloud-barista/cm-beetle/ui/src/components/layout/AppLayout.tsx)
- [ui/src/components/source/SourceCenter.tsx](file:///home/ubuntu/dev/cloud-barista/cm-beetle/ui/src/components/source/SourceCenter.tsx)
- [ui/src/components/design/MigrationDesigner.tsx](file:///home/ubuntu/dev/cloud-barista/cm-beetle/ui/src/components/design/MigrationDesigner.tsx)
- [ui/src/components/center/MigrationCenter.tsx](file:///home/ubuntu/dev/cloud-barista/cm-beetle/ui/src/components/center/MigrationCenter.tsx)
- [ui/src/store/migrationStore.ts](file:///home/ubuntu/dev/cloud-barista/cm-beetle/ui/src/store/migrationStore.ts)

## Code Quality Standards

### TypeScript

- Enable strict mode (`strict: true`)
- Define interfaces for all props and state
- Use `React.FC<PropsType>` for components

### Styling

- Tailwind utilities first, custom CSS only when necessary
- Group classes: `bg-* text-* border-* rounded-* p-* m-*`
- Dark mode by default (`dark:` class support)
- Light mode high contrast: Use `--border-input` (`#cbd5e1`) for form borders and high-contrast text (`text-slate-800` or `text-text-main`) for previews.
