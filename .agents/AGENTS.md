# Antigravity Project Rules for CM-Beetle / Beetle UX Lab

## UI Typography, Styling & Terminology Rules (Beetle UX Lab)

1. **Key-Value Styling & Font Hierarchy**:
   - **Key (Labels / Field Names)**: Always use standard font weight (**`font-normal text-text-muted`**). Do NOT use bold weights for key labels.
   - **Value (Data / Values / Metrics)**: Always use bold font weight (**`font-extrabold text-text-main`**) to emphasize actual data values.
   - **Minimum Font Size**: Main labels and readable data values must be at least **`text-sm`** (16px). Avoid microscopic font sizes (`text-[9px]`, `text-[10px]`, `text-[11px]`).

2. **Text Casing & Project Terminology**:
   - **Text Casing**: Section headers & labels must use **Title Case** or **Sentence Case** only (e.g., `Node Spec`, `Node Image`, `Root Disk`, `Security Group(s)`, `Node(s)`).
   - **No Forced Uppercase**: Do NOT apply Tailwind `uppercase` or `tracking-wider` to static labels or headers. Reserve `uppercase` only for dynamic data badges (e.g., `{status}`, `{protocol}`).
   - **Project Terminology**:
     - Always use **`Nodes`** or **`Node Groups`** across all UI labels, tabs, and headers (never `VM Instances` or `Active Instances`).
     - Main step tab name: **`2. Target Cloud Optimizer`** (never `Cloud Target Optimizer`).
     - Spec header: **`onpremiseInfraModel`** (omit `Spec` postfix).

3. **Tab Description Header Boxes**:
   - Single-line flex container (`flex flex-wrap items-center gap-x-3 gap-y-1.5 px-6 py-4.5 rounded-2xl`).
   - Title: `text-base font-extrabold text-text-main` accompanied by its tab Lucide icon (`w-5 h-5 text-emerald-500`).
   - Subtitle description: `text-sm text-text-muted`.

4. **Brand Palette & Light Mode Contrast**:
   - **Brand Colors**: Emerald (#10b981) / Teal (#14b8a6) **ONLY**. Never use `cyan-*`, `purple-*`, or `sky-*`.
   - **Light Mode Contrast**: Ensure clear inputs and readable previews in light mode. Use `--border-input` (`#cbd5e1`) for form borders and high-contrast text (`text-slate-800` or `text-text-main`) for code/previews.
