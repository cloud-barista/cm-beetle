# AI-Assisted Development: Context for Us

Welcome! This document explains how we use AI context files to work together more effectively on CM-Beetle.
Think of these files as our **shared understanding** with AI—when we maintain them well,
**everyone benefits** from consistent, high-quality assistance.

## Table of Contents

- [Context Files](#context-files)
- [GitHub Copilot Integration](#github-copilot-integration)
- [Supporting Tools](#supporting-tools)
- [Troubleshooting](#troubleshooting)

## Context Files

Context files are the **foundation** of our AI-assisted development workflow.
They **teach AI tools** about CM-Beetle's architecture, patterns, and standards.

### Copilot Instructions

✅ **"This file helps Copilot understand CM-Beetle like a maintainer or contributor would."**

**`.github/copilot-instructions.md`** is our main guide for GitHub Copilot. It covers:

- Project architecture and design patterns
- How we integrate with CB-Tumblebug and cm-model
- Echo handler patterns we follow
- Our coding standards and conventions

### Component Instructions

✅ **"These ensure Copilot gives us the right guidance depending on the file we're editing."**

**`.github/instructions/`** contains specialized rules that apply based on what we're working on:

| File Pattern  | Instruction File           | What It Covers                        |
| ------------- | -------------------------- | ------------------------------------- |
| `**/*.go`     | `go.instructions.md`       | Go standards, imports, error handling |
| `**/*.md`     | `markdown.instructions.md` | Documentation style and grammar       |
| `analyzer/**` | `analyzer.instructions.md` | Analyzer module-specific constraints  |
| `transx/**`   | `transx.instructions.md`   | TransX library guidelines             |

### Prompt Templates

✅ **"These prompts save time and ensure consistency across this project."**

**`.github/prompts/`** contains reusable workflows that automate common tasks:

**`git-commit.prompt.md`** helps write conventional commit messages:

- Analyzes our **staged changes**
- Determines the appropriate commit type and scope
- Formats everything according to our conventions

> **Note:** Proposed messages are reviewed and finalized by contributors.

**`api-guide.prompt.md`** generates comprehensive API documentation:

- Reads our Swagger specifications
- Creates realistic JSON request/response examples
- Includes prerequisites and best practices

## GitHub Copilot in VS Code

When we open CM-Beetle in VS Code, GitHub Copilot **automatically loads** our context files:

1. `.github/copilot-instructions.md` - Core project knowledge
2. `.github/instructions/*.instructions.md` - Component-specific expertise

**Try it out in VS Code:**

Open Copilot Chat (`Ctrl+Alt+I`) and ask:

```
"What are the key coding standards for this project?"
"How do I integrate with CB-Tumblebug?"
"Show me the Echo handler pattern we use"
```

Copilot will reference our instruction files and give us **project-specific answers**—like having an **experienced maintainer or contributor** available at any time.

## How to Use Prompt Templates

Our prompt templates work with different AI tools depending on the task.

### Generate Commit Messages (VS Code Copilot)

Let VS Code Copilot write our commit messages using our conventions:

1. Stage our changes: `git add .`
2. Open Copilot Chat in VS Code (`Ctrl+Alt+I`)
3. Type: `/git-commit`

Copilot analyzes our changes and generates a properly formatted conventional commit message:

```
feat(migration): add proximity-based VM sorting

- Sort by vCPU/memory distance for all machine types
- Add Azure hypervisor generation compatibility
```

The proposed commit title/message is then reviewed and finalized by the contributor.

### Generate API Documentation (GitHub Copilot CLI)

Need documentation for an API endpoint? Use the Copilot CLI tool:

```bash
make api-guide API_PATH=/migration/ns/{nsId}/mci
```

This uses GitHub Copilot CLI to read our Swagger specs and generate comprehensive guides with realistic examples.

**First-time setup (Copilot CLI):**

```bash
npx @github/copilot
# Type: /login and follow the prompts
```

## Troubleshooting

**VS Code Copilot not following our guidelines?**

- Try reloading VS Code: `Ctrl+Shift+P` > "Reload Window"
- Verify it's working: Ask in Copilot Chat `"What instructions do you have for this project?"`

**GitHub Copilot CLI authentication issues?**

```bash
npx @github/copilot
# Type: /login
```

**Getting incomplete API examples?**

- Run `make swag` to regenerate Swagger docs
- Check that `api/swagger.yaml` has complete schema definitions

**Component-specific rules not being applied in VS Code?**

- Verify our file location matches the `applyTo` pattern in the instruction file
- Try reloading VS Code: `Ctrl+Shift+P` > "Reload Window"

---

**Need more help?** Check out: [GitHub Copilot Docs](https://docs.github.com/en/copilot) | [Conventional Commits](https://www.conventionalcommits.org/) | [Our API Specs](../api/swagger.yaml)

**Want to improve our AI context?** Great! These files live alongside our code—update them as the project evolves and submit your improvements via pull request. **Better context helps everyone.**
