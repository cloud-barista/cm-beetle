# Go Development Philosophy & Coding Rules

## 1. Core Engineering Philosophy

- **"Simple is better than complex. Clear is better than clever."**
  - Prioritize direct, transparent, and readable Go code over multi-layered abstractions, premature frameworks, or complex indirection.
- **"Design for what exists today, not an imaginary future."**
  - Solve concrete requirements grounded in existing data models and real system constraints. Avoid speculative engineering for hypothetical use cases.

## 2. Architectural & Design Standards

- **Go Single Responsibility Principle (SRP):**
  - Every file, struct, and function must have a single, well-defined responsibility and a single reason to change.
  - Independent feature domains (e.g., standard VM infra, NLB clustering, GPU hardware acceleration) should remain focused and self-contained.
- **No Fallback (Deterministic Logic & Explicit Errors):**
  - Never implement silent fallback degradation or hidden catch-all defaults that mask underlying issues.
  - Validate inputs and preconditions explicitly; fail fast with clear, actionable error messages.
- **No Circular Dependencies (순환 참조 불가):**
  - Maintain strict acyclic dependency graphs across packages, interfaces, and types.
  - Decouple shared utilities and domain models to ensure clean unidirectional imports.
- **Maintainability & Continuous Improvement (유지보수 및 지속 개선):**
  - Code must be linearly readable and straightforward to debug, test, and profile.
  - As new resources are introduced, consider the concrete inter-resource relationships (연관 관계) explicitly rather than blindly wrapping them in generic layers.

## 3. Code Block Comment Rules

- **Maximum 1 Line per Comment Block:**
  - Comments inside code blocks must be limited to **1 single line**.
  - Restate only non-obvious intent or critical constraints, not plain syntax.
- **Reference / Example Exception (+1 Line):**
  - An additional **1 line** is permitted exclusively when citing an external reference URL, issue number, or brief code example.
- **Language & Tone:**
  - All comments must be written in **English** with technical accuracy and conciseness.
  - Never write historical diff explanations or "do not reintroduce X" narratives in source code comments.
