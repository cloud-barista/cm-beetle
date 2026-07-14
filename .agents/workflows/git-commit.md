---
description: Generate conventional commit messages for CM-Beetle based on staged changes
---

This workflow helps you generate a high-quality, conventional commit message tailored for the CM-Beetle project by analyzing your staged Git changes.

1. **Identify Staged Changes**  
   Check the current staging area to see which files are ready to be committed.
   // turbo
   Run `git diff --cached --name-status` to list staged files and their status (Added, Modified, Deleted).

2. **Handle Empty Staging Area**  
   If no files are staged, notify the user that they need to stage changes first (e.g., using `git add`) before generating a commit message.

3. **Analyze Diff Content**  
   For the staged files, examine the actual code changes to understand the intent and impact.
   // turbo
   Run `git diff --cached` to view the detailed modifications. Focus on `pkg/` and `cmd/` for core logic, while treating `go.sum`, `swagger.yaml`, and `*.mock.go` as lower-priority context.

4. **Determine Commit Type and Scope**  
   Select the most appropriate type and scope based on the CM-Beetle project structure:

   **Commit Types:**
   - `feat`: New features (APIs, migration capabilities, recommendation algorithms)
   - `fix`: Bug fixes
   - `refactor`: Structural code changes without changing behavior
   - `enhance`: Minor improvements to existing features
   - `update`: Update existing functionality or dependencies
   - `docs`: Documentation changes
   - `test`: Test code updates
   - `context`: AI context files (instructions, prompts)
   - `chore`: Maintenance (dependencies, build config)
   - `release`: Release staging (version bumps, prepping releases)

   **Common Scopes:**
   - `migration`: Infrastructure migration (`pkg/core/migration/`)
   - `recommendation`: VM/Infrastructure recommendation (`pkg/core/recommendation/`)
   - `api`: REST API handlers/models (`pkg/api/rest/`)
   - `client`: CB-Tumblebug client (`pkg/client/tumblebug/`)
   - `test-cli`: Test CLI tool (`cmd/test-cli/`)
   - `core`: Core logic/utilities (`pkg/core/`)
   - `config`: Configuration management (`pkg/config/`)
   - `transx`: Data transfer utilities (`transx/`)
   - `analyzer`: Infrastructure analysis (`analyzer/`)

5. **Generate the Commit Message**  
   Compose the message following these strict requirements:
   - **Title**: `type(scope): description` (Max 50 characters, imperative mood).
   - **Body**: **Max 3 bullet points**, each **≤ 40 characters**. Focus on functional impact only. Omit obvious or low-value lines (e.g., "update README", "regenerate swagger").
   - **Breaking Changes**: Add `BREAKING CHANGE: <description>` in the footer if public APIs or configs changed.
   - **Release Staging**: If bumping versions or refreshing multiple test results, use `release: staging vX.Y.Z with [highlights]`.

6. **Output the Result**  
   Present the final commit message to the user in a code block, ready for use.
