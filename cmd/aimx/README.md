# AIMX (AI Migration X) — AI Service Migration Prototype CLI

**AIMX** (**AI Migration X** — _AI Migration eXecutor_) is an end-to-end Proof-of-Concept (PoC) CLI tool that automates **AI Semiconductor (GPU) Infrastructure & AI Model/Application Migration** by orchestrating **CB-Tumblebug** and **CM-Beetle (`transx`)** REST APIs.

> 📄 **For Top-Down Architecture, Source ↔ Target 1:1 Mapping, Storage/Volume/DB Migration Spec & Chronological Sequence Diagrams**:
> See [`../../docs/poc/ai-service-migration-design.md`](../../docs/poc/ai-service-migration-design.md).

---

## 1. Overview & 4-Layer Migration Scope (Source ↔ Target 1:1 Mapping)

`AIMX` migrates a live GPU-accelerated AI service stack across cloud environments in four layers:

1. **Layer 1 — AI Semiconductor (GPU) VM Infrastructure & Cloud Block Storage**:
   - **Source**: AWS GPU VM (`mc-daisy`, `aws+ap-northeast-2+g4dn.xlarge`, `1x NVIDIA Tesla T4 16GB VRAM`, 4 vCPU, 16 GiB RAM, Ubuntu 22.04 Deep Learning AMI `ami-00ffffe72a332c865`, 100GB EBS Root Disk) provisioned directly via CB-Tumblebug's `POST /tumblebug/ns/{nsId}/infraDynamic`.
   - **Target**: Recommended via CM-Beetle's GPU-aware recommender (`POST /beetle/recommendation/infra`) and provisioned via CM-Beetle (`POST /beetle/migration/ns/{nsId}/infra`).
2. **Layer 2 — AI Semiconductor Runtime**:
   - NVIDIA GPU Driver (`595.91.07`), CUDA Runtime (`13.2`), Docker Engine + Compose v2, and `nvidia-container-toolkit` (`driver: nvidia`, `count: all`).
3. **Layer 3 — Self-Hosted LLM Engine & Model Weights Volume (`Ollama` + `Gemma 4`)**:
   - **Ollama (`ollama/ollama:latest`, `:11434`)** serving **Google Gemma 4 (`gemma4:e2b`)** with host bind-mounted model weights and manifests in `/opt/aimx/ai-service/ollama` $\rightarrow$ `/root/.ollama` (`~4.6 GB` GGUF blobs).
4. **Layer 4 — Stateful AI Web Application & Database Volume (`Open WebUI` + `SQLite`)**:
   - **Open WebUI (`ghcr.io/open-webui/open-webui:main`, `:3000`)** connected to Ollama, storing user accounts, prompt templates, chat history, and RAG vector indices in `/opt/aimx/ai-service/open-webui` $\rightarrow$ `/app/backend/data` (`webui.db` SQLite RDBMS + `vector_db/` + `uploads/`).

---

## 2. Directory Structure

```text
cmd/aimx/
├── main.go                           # CLI entrypoint (-config, -mode, -cleanup-source router & seq pipeline)
├── config.go                         # YAML/JSON configuration, Resty clients, & state persistence
├── wait.go                           # Async SSE remote command streamer, countdowns & readiness polling
├── source.go                         # Stage 1: Tumblebug GPU infraDynamic provisioning & hybrid GPU discovery
├── infra_migration.go                # Stage 2 & 3-1: Beetle GPU recommendation & target GPU VM provisioning
├── service_migration.go              # Stage 3-2..6: Ephemeral SSH keys, container quiesce, & transx transfer
├── verification.go                   # Stage 4 & 5: GPU/Gemma tok/s benchmark, SHA-256 DB check, & cleanup
├── report.go                         # Console comparison table & Markdown report generator
├── main_test.go                      # Unit tests for config, GPU parsing, transx payload, and reporting
├── README.md                         # Quickstart & CLI execution guide
├── testconf/
│   ├── template-auth-config.json     # Template credentials for CM-Beetle & CB-Tumblebug
│   ├── template-test-config.yaml     # Template execution modes, GPU stack settings & timing controls
│   └── source-infra-base.json        # Base OnpremInfra JSON enriched by live nvidia-smi inspection
└── testresult/                       # Generated state (.aimx-state.json) & Markdown verification reports
```

---

## 3. Quickstart (Makefile & CLI)

### Run End-to-End Sequential Migration (`seq`)

From the repository root (`cm-beetle`):

```bash
# Run all 4 stages sequentially (setup-source -> recommend -> migrate -> verify)
make test-aimx
```

### Run Individual Stages in Chronological Order

Each stage persists its state to `cmd/aimx/testresult/.aimx-state.json`, allowing step-by-step execution and verification:

```bash
# Stage 1: Provision Source AWS Tesla T4 GPU VM & deploy Ollama (gemma4:e2b) + Open WebUI
make test-aimx EXTRA_ARGS="-mode setup-source"

# Stage 2: Inspect Source GPU VM via SSH & get Target GPU VM recommendations from CM-Beetle
make test-aimx EXTRA_ARGS="-mode recommend"

# Stage 3: Provision Target GPU VM & migrate /opt/aimx/ai-service (Gemma blobs + webui.db) via transx
make test-aimx EXTRA_ARGS="-mode migrate"

# Stage 4: Benchmark Gemma GPU tok/s & verify SHA-256 SQLite DB integrity on Source vs. Target
make test-aimx EXTRA_ARGS="-mode verify"

# Stage 5 (Cleanup Target Only): Delete migrated Target GPU infrastructure
make test-aimx EXTRA_ARGS="-mode cleanup"

# Stage 5 (Cleanup All): Delete both Target and Source GPU infrastructure + default resources
make test-aimx EXTRA_ARGS="-mode cleanup -cleanup-source=true"
```

---

## 4. Key Configuration Options (`testconf/test-config.yaml`)

| Section     | Key                      | Default                            | Description                                                                                                                              |
| :---------- | :----------------------- | :--------------------------------- | :--------------------------------------------------------------------------------------------------------------------------------------- |
| `test`      | `executionMode`          | `"seq"`                            | Default execution mode (`seq`, `setup-source`, `recommend`, `migrate`, `verify`, `cleanup`).                                             |
| `source`    | `infraId`                | `"mc-daisy"`                       | Source GPU Infra ID in CB-Tumblebug. Reused if already running and healthy.                                                              |
| `source`    | `specId`                 | `"aws+ap-northeast-2+g4dn.xlarge"` | GPU VM spec (`4 vCPU, 16 GiB RAM, 1x NVIDIA Tesla T4 16GB VRAM`) provisioned via `POST /tumblebug/ns/{nsId}/infraDynamic`.               |
| `source`    | `imageId`                | `"ami-00ffffe72a332c865"`          | Ubuntu 22.04 Deep Learning OSS Nvidia Driver GPU AMI in `ap-northeast-2`.                                                                |
| `aiService` | `rootDir`                | `"/opt/aimx/ai-service"`           | Host root directory containing `docker-compose.yml`, `ollama/` (model volume), and `open-webui/` (SQLite/vector DB volume).              |
| `aiService` | `modelName`              | `"gemma4:e2b"`                     | Google Gemma 4 model tag pulled and benchmarked in Ollama (`~4.6 GB`).                                                                   |
| `aiService` | `transferModelWeights`   | `true`                             | `true`: transfers `/opt/aimx/ai-service/ollama` blobs via `transx`. `false`: excludes `ollama/**` and pulls model on Target.             |
| `aiService` | `transxStrategy`         | `"relay"`                          | `transx` execution strategy (`"relay"` stages via CM-Beetle; `"direct"` uses in-memory SSH agent forwarding; `"auto"`).                  |
| `aiService` | `filter.include/exclude` | `exclude: cache/**, *.log, *.tmp`  | Maps directly to `transx.FilterOption` (`rsync --include` / `--exclude`) for selective AI model/DB transfer.                             |
| `aiService` | `useDataEncryption`      | `true`                             | `true`: encrypts ephemeral Tumblebug SSH keys in `DataMigrationModel` via `GET /beetle/migration/data/encryptionKey` (1-time KeyStore).  |
