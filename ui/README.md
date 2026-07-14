# CM-Beetle UI

Web-based dashboard for **CM-Beetle** — the Computing Infrastructure Migration sub-system of the [Cloud-Barista](https://github.com/cloud-barista) platform.

## Tech Stack

- **[Next.js 15](https://nextjs.org/docs)** (App Router) — framework and server-side API proxy
- **[TypeScript](https://www.typescriptlang.org/docs)** — type safety
- **[Tailwind CSS](https://tailwindcss.com/docs)** — utility-first CSS
- **[shadcn/ui](https://ui.shadcn.com/docs)** — UI component library
- **[Zustand](https://zustand-demo.pmnd.rs/)** — client-side state management
- **[React Flow](https://reactflow.dev/)** — topology visualization

## Architecture

The UI is a **standalone Next.js application** served separately from the Beetle Go backend.
Next.js API route handlers act as a server-side reverse proxy, forwarding browser requests to each backend service using environment variables for endpoint configuration.

```
Browser → Next.js (port 3000)
             ├── /beetle/*    → BEETLE_ENDPOINT (cm-beetle:8056)
             ├── /tumblebug/* → TUMBLEBUG_ENDPOINT (cb-tumblebug:1323)
             ├── /honeybee/*  → HONEYBEE_ENDPOINT (cm-honeybee:8081)
             └── /damselfly/* → DAMSELFLY_ENDPOINT (cm-damselfly:8088)
```

## Getting Started

### Prerequisites

- Node.js 20+

### Local Development

```bash
# Install dependencies
npm install

# Create local environment file
cp .env.example .env.local
# Edit .env.local with the actual backend service URLs

# Start dev server (http://localhost:3000)
npm run dev
```

### Production Build

```bash
npm run build
npm start
```

### Docker

```bash
docker build -t cm-beetle-ui .
docker run -p 3000:3000 \
  -e BEETLE_ENDPOINT=http://cm-beetle:8056 \
  -e TUMBLEBUG_ENDPOINT=http://cb-tumblebug:1323 \
  -e HONEYBEE_ENDPOINT=http://cm-honeybee:8081 \
  -e DAMSELFLY_ENDPOINT=http://cm-damselfly:8088 \
  cm-beetle-ui
```

## Environment Variables

| Variable             | Default                 | Description             |
| -------------------- | ----------------------- | ----------------------- |
| `BEETLE_ENDPOINT`    | `http://localhost:8056` | CM-Beetle API server    |
| `TUMBLEBUG_ENDPOINT` | `http://localhost:1323` | CB-Tumblebug API server |
| `HONEYBEE_ENDPOINT`  | `http://localhost:8081` | CM-Honeybee API server  |
| `DAMSELFLY_ENDPOINT` | `http://localhost:8088` | CM-Damselfly API server |

## License

Apache License 2.0 — see the [LICENSE](../LICENSE) file for details.

The React Compiler is not enabled on this template because of its impact on dev & build performances. To add it, see [this documentation](https://react.dev/learn/react-compiler/installation).

## Expanding the Oxlint configuration

If you are developing a production application, we recommend enabling type-aware lint rules by installing `oxlint-tsgolint` and editing `.oxlintrc.json`:

```json
{
  "$schema": "./node_modules/oxlint/configuration_schema.json",
  "plugins": ["react", "typescript", "oxc"],
  "options": {
    "typeAware": true
  },
  "rules": {
    "react/rules-of-hooks": "error",
    "react/only-export-components": ["warn", { "allowConstantExport": true }]
  }
}
```

See the [Oxlint rules documentation](https://oxc.rs/docs/guide/usage/linter/rules) for the full list of rules and categories.
