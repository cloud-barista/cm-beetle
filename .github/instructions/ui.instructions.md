---
applyTo:
  - ui/**
---

# Beetle Lab UI Instructions

## Quick Reference

**📘 Full Design System:** Read [`ui/DESIGN_SYSTEM.md`](../../ui/DESIGN_SYSTEM.md) for comprehensive guidelines on branding, colors, components, and tech stack.

## Essential Rules

### Branding & Colors

- **Brand Name:** Beetle Lab (UI), CM-Beetle (documentation/API)
- **Color Palette:** Emerald (#10b981) / Teal (#14b8a6) **only**
- **Forbidden Colors:** ❌ `cyan-*`, `purple-*`, `sky-*` (removed from codebase)
- **Gradients:** `from-emerald-400 via-teal-400 to-teal-400`

### Tech Stack

- **Framework:** Next.js 15 App Router with standalone output
- **React:** v19 with `'use client'` for interactive components
- **Styling:** Tailwind CSS 4 (@theme inline configuration)
- **State:** Zustand v5 for global state
- **Icons:** Lucide React (HardDrive, Database, Cpu, Compass, etc.)

### Architecture

**Server-Side API Proxy Pattern:**

```typescript
// app/beetle/[[...path]]/route.ts
import { proxy } from '@/lib/proxy';

const TARGET = process.env.BEETLE_ENDPOINT || 'http://localhost:8056';

export async function GET(req: Request, { params }: { params: { path?: string[] } }) {
  return proxy(req, TARGET, '/beetle', params.path);
}
```

**Client Components:**

```tsx
'use client'; // Required for useState, Zustand, event handlers

import { useMigrationStore } from '@/store/migrationStore';
```

### Component Patterns

**Glass Panel:**

```tsx
<div className="bg-bg-panel backdrop-blur-sm border border-border-main rounded-xl p-6">
  {/* Content */}
</div>
```

**Primary Button:**

```tsx
<button className="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white rounded-lg">
  Action
</button>
```

**Icon Usage:**

```tsx
import { Database } from 'lucide-react';

<Database className="w-5 h-5 text-emerald-500" />
```

## File Organization

```
ui/src/
├── app/
│   ├── layout.tsx              # Root layout
│   ├── page.tsx                # Main entry with tab routing
│   ├── globals.css             # Tailwind theme (emerald/teal)
│   ├── beetle/[[...path]]/     # API proxy routes
│   ├── tumblebug/[[...path]]/
│   ├── honeybee/[[...path]]/
│   └── damselfly/[[...path]]/
├── components/
│   ├── layout/AppLayout.tsx    # Header, navigation, theme toggle
│   ├── source/SourceCenter.tsx
│   ├── design/MigrationDesigner.tsx
│   └── center/MigrationCenter.tsx
├── store/migrationStore.ts     # Zustand global state
└── lib/proxy.ts                # Generic reverse-proxy helper
```

## Code Quality Standards

### TypeScript

- Enable strict mode (`strict: true`)
- Define interfaces for all props and state
- Use `React.FC<PropsType>` for components

### Styling

- Tailwind utilities first, custom CSS only when necessary
- Group classes: `bg-* text-* border-* rounded-* p-* m-*`
- Use `clsx()` for conditional classes
- Dark mode by default (optional `dark:` prefix)

### Text Casing

- **Labels and section headers**: Title Case only — **never ALL CAPS**
  - ✅ `Spec`, `Root Disk`, `Security Group`, `Node Count`
  - ❌ `SPEC`, `ROOT DISK`, `SECURITY GROUP`, `NODE COUNT`
- **Tailwind**: Do **not** use `uppercase` + `tracking-wide/tracking-wider` on labels; reserve `uppercase` for dynamic data-value badges (e.g., `{status}`, `{direction}`)
- **Data value badges** (e.g., status pills, direction indicators): `uppercase` is acceptable to visually distinguish computed values from labels

### Badge vs Plain Text

Use **colored background badges** for:
- Resource identifiers compared across candidates: vCPU, memory, instance type (emerald), OS image (teal), security group name (orange)
- Status / categorical values: match status, direction, type

Use **plain text** for:
- Scalar measurements: disk size (GB), node count
- Form field values inside inputs

### State Management

```typescript
// Zustand store pattern
import { create } from 'zustand';

interface Store {
  value: string;
  setValue: (value: string) => void;
}

export const useStore = create<Store>((set) => ({
  value: '',
  setValue: (value) => set({ value }),
}));
```

## Docker Deployment

**Standalone Build:**

```dockerfile
# Multi-stage build with node:20-bookworm-slim
# Copies .next/standalone for production
```

**Environment Variables:**

```bash
BEETLE_ENDPOINT=http://cm-beetle:8056
TUMBLEBUG_ENDPOINT=http://cb-tumblebug:1323
HONEYBEE_ENDPOINT=http://cm-honeybee:8081
DAMSELFLY_ENDPOINT=http://cm-damselfly:8088
```

## Common Tasks

### Add New API Route

1. Create `app/[service]/[[...path]]/route.ts`
2. Import `proxy` from `@/lib/proxy`
3. Define `TARGET` from environment variable
4. Export GET, POST, PUT, DELETE methods

### Add New Component

1. Create in appropriate `components/` subdirectory
2. Add `'use client'` if using hooks/interactivity
3. Use Lucide React icons with `text-emerald-*` or `text-teal-*`
4. Follow glass panel pattern for containers

### Update Colors

1. Edit `src/app/globals.css` @theme variables
2. Use `emerald-*` / `teal-*` Tailwind classes
3. Update CSS custom properties in `:root` for dark/light modes

## Design Philosophy

1. **Emerald Elegance:** Beetle-inspired emerald/teal palette
2. **Glass & Glow:** Premium glassmorphic effects
3. **Dark-First:** Default dark mode (#071a10 emerald-tinted background)
4. **Server-Side Security:** Environment variables never exposed to browser
5. **Clarity Over Complexity:** Prioritize user comprehension

---

**For Detailed Guidelines:** Always consult [`ui/DESIGN_SYSTEM.md`](../../ui/DESIGN_SYSTEM.md) before making significant changes to branding, colors, or component patterns.
