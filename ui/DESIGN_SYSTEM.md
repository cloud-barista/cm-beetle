# Beetle Lab Design System

> Unified design guidelines for CM-Beetle UI across all AI coding assistants (GitHub Copilot, Claude, Cursor, Windsurf, etc.)

## Overview

**Beetle Lab** is the demonstration and testing interface for CM-Beetle, a Cloud-Barista sub-system enabling infrastructure migration from on-premise to multi-cloud environments.

## Branding

### Identity

- **Name:** Beetle Lab
- **Tagline:** Cloud Infrastructure Migration Laboratory
- **Version:** v0.5.6
- **Concept:** Inspired by the iridescent emerald-teal colors of scarab beetles in nature

### Logo

- **Icon:** HardDrive (Lucide React)
- **Background:** Gradient from `emerald-500` to `teal-500`
- **Shadow:** `shadow-emerald-500/20`

### Color Palette

Primary brand colors inspired by beetle exoskeletons:

```css
/* Brand Colors */
--color-brand-emerald: #10b981  /* Primary brand color */
--color-brand-teal: #14b8a6      /* Secondary brand color */

/* Dark Mode (Default) */
--bg-main: #071a10               /* Emerald-tinted dark background */
--bg-panel: rgba(16, 185, 129, 0.05)  /* Subtle emerald glass panels */
--bg-input: #0a1f13              /* Input backgrounds */
--text-main: #f8fafc             /* Primary text */
--text-muted: #94a3b8            /* Secondary text */
--border-main: rgba(16, 185, 129, 0.08)    /* Subtle emerald borders */
--border-input: rgba(16, 185, 129, 0.2)    /* Input borders */

/* Light Mode */
--bg-main: #f1f5f9
--bg-panel: #ffffff
--bg-input: #f8fafc
--text-main: #0f172a
--text-muted: #64748b
--border-main: rgba(0, 0, 0, 0.08)
--border-input: rgba(0, 0, 0, 0.15)
```

### Tailwind Color Usage

**Always use emerald/teal variants:**

- ✅ `bg-emerald-500`, `text-emerald-400`, `border-teal-300`
- ✅ `from-emerald-400 via-teal-400 to-teal-400` (gradients)
- ✅ `hover:bg-emerald-500/10` (hover states)
- ❌ `bg-cyan-500`, `text-purple-400` (deprecated, removed from codebase)

### Typography

Enlarged font scales for improved readability:

```css
--text-xs: 0.95rem;    /* 15.2px */
--text-sm: 1.1rem;     /* 17.6px */
--text-base: 1.25rem;  /* 20px */
--text-lg: 1.4rem;     /* 22.4px */
--text-xl: 1.6rem;     /* 25.6px */
--text-2xl: 1.85rem;   /* 29.6px */
--text-3xl: 2.2rem;    /* 35.2px */
```

## Tech Stack

### Core Framework

- **Next.js:** v15.1.0 (App Router with standalone output mode)
- **React:** v19.0.0 (Client components with `'use client'` directive)
- **TypeScript:** v5.7.0
- **Tailwind CSS:** v4.0.0 (@theme inline configuration)

### Key Libraries

- **State Management:** Zustand v5.0.3
- **HTTP Client:** Axios v1.7.9
- **Icons:** Lucide React v0.475.0
- **Charts:** Recharts v2.15.1
- **Diagrams:** XYFlow v12.4.2
- **Styling Utilities:** clsx, tailwind-merge, class-variance-authority

### Backend Integration

Server-side API proxy routes forward requests to backend services:

- `/beetle/[[...path]]` → CM-Beetle API (`BEETLE_ENDPOINT`)
- `/tumblebug/[[...path]]` → CB-Tumblebug API (`TUMBLEBUG_ENDPOINT`)
- `/honeybee/[[...path]]` → CM-Honeybee API (`HONEYBEE_ENDPOINT`)
- `/damselfly/[[...path]]` → CM-Damselfly API (`DAMSELFLY_ENDPOINT`)

Environment variables are resolved server-side in Next.js API routes, not exposed to browser.

## Component Patterns

### Glass Panel Effect

Premium glassmorphic panels with emerald tints:

```tsx
<div className="bg-bg-panel backdrop-blur-sm border border-border-main rounded-xl p-6">
  {/* Content */}
</div>
```

**Interactive variant:**

```tsx
<div className="glass-panel-interactive hover:border-emerald-400/40 hover:shadow-emerald-500/8">
  {/* Content */}
</div>
```

### Gradients

Use emerald-to-teal gradients for visual hierarchy:

```tsx
{/* Text gradient */}
<span className="bg-gradient-to-r from-emerald-400 via-teal-400 to-teal-400 bg-clip-text text-transparent">
  Beetle Lab
</span>

{/* Background gradient */}
<div className="bg-gradient-to-br from-emerald-500 to-teal-500">
  {/* Content */}
</div>
```

### Buttons

**Primary button:**

```tsx
<button className="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white rounded-lg transition-colors">
  Action
</button>
```

**Secondary button:**

```tsx
<button className="px-4 py-2 bg-bg-panel border border-border-main hover:border-emerald-400/40 rounded-lg transition-colors">
  Action
</button>
```

### Icons

Use Lucide React icons with emerald/teal colors:

```tsx
import { Database, Cpu, Compass } from 'lucide-react';

<Database className="w-5 h-5 text-emerald-500" />
```

### Glow Effects

For emphasis on interactive nodes or text:

```css
/* globals.css */
.glow-emerald {
  text-shadow: 0 0 10px rgba(16, 185, 129, 0.4);
}

.glow-teal {
  text-shadow: 0 0 10px rgba(20, 184, 166, 0.4);
}
```

## File Structure

```
ui/
├── src/
│   ├── app/
│   │   ├── layout.tsx              # Root layout with metadata
│   │   ├── page.tsx                # Main entry with tab routing
│   │   ├── globals.css             # Tailwind theme & custom styles
│   │   ├── beetle/[[...path]]/     # API proxy routes
│   │   ├── tumblebug/[[...path]]/
│   │   ├── honeybee/[[...path]]/
│   │   └── damselfly/[[...path]]/
│   ├── components/
│   │   ├── layout/
│   │   │   └── AppLayout.tsx       # Header, navigation, theme toggle
│   │   ├── source/
│   │   │   └── SourceCenter.tsx    # Source infrastructure management
│   │   ├── design/
│   │   │   └── MigrationDesigner.tsx  # Target cloud optimizer
│   │   └── center/
│   │       └── MigrationCenter.tsx    # Migration execution
│   ├── store/
│   │   └── migrationStore.ts       # Zustand global state
│   ├── lib/
│   │   └── proxy.ts                # Generic reverse-proxy helper
│   └── global.d.ts                 # TypeScript declarations for CSS
├── public/
│   └── favicon.svg                 # Emerald→teal gradient beetle icon
├── Dockerfile                      # Multi-stage Next.js build
├── next.config.ts                  # output: 'standalone'
├── package.json
└── DESIGN_SYSTEM.md               # This file
```

## Coding Conventions

### Next.js App Router

- **Server Components:** Default (no `'use client'`)
- **Client Components:** Add `'use client'` for interactivity (useState, Zustand, event handlers)
- **API Routes:** Place in `app/[service]/[[...path]]/route.ts` with all HTTP methods

### Component Structure

```tsx
'use client'; // Only if using hooks or interactivity

import React from 'react';
import { SomeIcon } from 'lucide-react';

interface ComponentNameProps {
  // Props
}

export const ComponentName: React.FC<ComponentNameProps> = ({ prop }) => {
  // Hooks
  // Handlers
  // Render
  return (
    <div className="...">
      {/* Content */}
    </div>
  );
};
```

### Tailwind Best Practices

- Use semantic class grouping: `bg-* text-* border-* rounded-* p-* m-*`
- Prefer Tailwind utilities over custom CSS
- Use `clsx()` or `cn()` for conditional classes
- Leverage dark mode with `dark:` prefix (optional, defaults to dark)

### State Management

Use Zustand for global state:

```tsx
// store/migrationStore.ts
import { create } from 'zustand';

interface MigrationStore {
  activeTab: string;
  setActiveTab: (tab: string) => void;
}

export const useMigrationStore = create<MigrationStore>((set) => ({
  activeTab: 'source',
  setActiveTab: (tab) => set({ activeTab: tab }),
}));
```

### API Integration

Use server-side proxy routes:

```tsx
// Client component
const response = await fetch('/beetle/recommendation/mci', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify(payload),
});
```

Backend services are accessed server-side via environment variables.

## Docker Deployment

### Standalone Build

```dockerfile
# ui/Dockerfile
FROM node:20-bookworm-slim AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM node:20-bookworm-slim AS runner
WORKDIR /app
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public
ENV NODE_ENV=production
EXPOSE 3000
CMD ["node", "server.js"]
```

### Environment Variables

```bash
BEETLE_ENDPOINT=http://cm-beetle:8056
TUMBLEBUG_ENDPOINT=http://cb-tumblebug:1323
HONEYBEE_ENDPOINT=http://cm-honeybee:8081
DAMSELFLY_ENDPOINT=http://cm-damselfly:8088
```

### Docker Compose

Use overlay pattern for optional UI services:

```bash
# Base services only
docker compose up

# With UI + Honeybee + Damselfly
COMPOSE_FILE=docker-compose.yaml:docker-compose.ui.yaml docker compose up
```

## Development Workflow

### Local Development

```bash
cd ui
npm install
npm run dev  # http://localhost:3000
```

### Build & Test

```bash
npm run build   # Generates .next/standalone
npm run start   # Production server
```

### Linting

```bash
npm run lint    # ESLint with Next.js rules
```

## Design Philosophy

1. **Clarity Over Complexity:** Prioritize user comprehension of cloud migration workflows
2. **Emerald Elegance:** Consistent use of emerald/teal palette inspired by nature
3. **Glass & Glow:** Premium glassmorphic effects with subtle emerald tints
4. **Dark-First:** Default to dark mode with emerald-tinted backgrounds (#071a10)
5. **Responsive Typography:** Enlarged font scales for readability
6. **Server-Side Security:** Environment variables never exposed to browser

## AI Assistant Guidelines

When working with Beetle Lab UI:

- **Colors:** Always use `emerald-*` and `teal-*`, never `cyan-*` or `purple-*`
- **Branding:** Refer to the project as "Beetle Lab" in UI, "CM-Beetle" in documentation
- **Components:** Use Lucide React icons, avoid importing other icon libraries
- **Patterns:** Follow glass panel and gradient patterns from this document
- **Architecture:** Server-side API proxy routes for backend communication
- **State:** Use Zustand for global state, React hooks for local state
- **Styling:** Tailwind utilities first, custom CSS only when necessary

---

**Last Updated:** July 2026  
**Maintained By:** CM-Beetle Team  
**For Questions:** Refer to `ui/README.md` or `.github/copilot-instructions.md`
