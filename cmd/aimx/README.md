# AIMX (AI Migration X): AI Service Migration Prototype CLI

**AIMX** (**AI Migration X**, _AI Migration eXecutor_) is an end-to-end Proof-of-Concept (PoC) CLI tool that automates **AI Semiconductor (GPU) Infrastructure & AI Model/Application Migration** by orchestrating **CB-Tumblebug** and **CM-Beetle (`transx`)** REST APIs.

It supports two stateful AI web applications out of the box via `aiService.appType`:
- **`kotaemon` (Recommended for Conference RAG Demo)**: **`Kotaemon Full` (`ghcr.io/cinnamon/kotaemon:main-full`, `v0.12.0`, `:7860`, GPU pass-through) + `Ollama` (`gemma4:12b` or `gemma4:e2b` + `bge-m3`, `:11434`)**, pre-seeded with `Docling` figure/table extraction, Korean response language, `inline` footnotes (`[1]`, `[2]`) with right-panel `[Preview]` evidence highlighting, and `Mindmap` visualization.
- **`open-webui`**: **`Open WebUI` (`ghcr.io/open-webui/open-webui:main`, `:3000`) + `Ollama` (`gemma4:e2b`)**.

---

## 1. Overview & 4-Layer Migration Scope (Source -> Target 1:1 Mapping)

`AIMX` migrates a live GPU-accelerated AI service stack across cloud environments in four layers:

1. **Layer 1: AI Semiconductor (GPU) VM Infrastructure & Cloud Block Storage**:
   - **Source Mode A (External OpenTofu Testbed - `mc-terrarium/examples/aws/ai-stack-testbed`)**: Primary AI inference node `ai-gpu-01` (`aws+ap-northeast-2+g6.2xlarge`, `1x NVIDIA L4 24GB VRAM`, 8 vCPU, 32 GiB RAM, 150GB Root Disk) from the 10-node testbed, connected directly via `source_connections_public.json` (or `source_connections_private.json` when running inside `cloud-migrator`).
   - **Source Mode B (Managed CB-Tumblebug Source - `mc-daisy`)**: `aws+ap-northeast-2+g4dn.xlarge` (`1x NVIDIA Tesla T4 16GB VRAM`, 4 vCPU, 16 GiB RAM, 100GB Root Disk) provisioned via CB-Tumblebug's `POST /tumblebug/ns/{nsId}/infraDynamic`.
   - **Target**: Recommended via CM-Beetle's zone- and availability-aware GPU recommender (`POST /beetle/recommendation/infra`), validated via `POST /beetle/validation/ns/{nsId}/infra`, and provisioned via `POST /beetle/migration/ns/{nsId}/infra` (e.g., `gcp+asia-northeast3+g2-standard-8` with `1x NVIDIA L4 24GB VRAM`).
2. **Layer 2: AI Semiconductor Runtime**:
   - NVIDIA GPU Driver (`550+`), CUDA Runtime, Docker Engine + Compose v2, and `nvidia-container-toolkit` (`driver: nvidia`, `count: all`). When the target cloud uses a bare OS image without pre-installed drivers (e.g., standard GCP Ubuntu 22.04), `scripts/ensure-gpu-runtime.sh` automatically installs and configures the complete GPU container runtime before starting containers.
3. **Layer 3: Self-Hosted LLM & Multilingual Embedding Volume (`Ollama` + `Gemma 4` + `bge-m3`)**:
   - **Ollama (`ollama/ollama:latest`, `:11434`)** serving **`gemma4:12b`** on 24GB L4/A10G GPUs (or **`gemma4:e2b`** on 16GB T4 GPUs) alongside **`bge-m3`** multilingual embeddings co-resident in GPU VRAM (`OLLAMA_MAX_LOADED_MODELS=2`), persisted under `/opt/aimx/ai-service/ollama`.
4. **Layer 4: Stateful RAG Web Application & Database Volume (`Kotaemon` + `SQLite` + `ChromaDB` / `LanceDB`)**:
   - **Kotaemon (`:7860`)** connected to Ollama via internal Docker DNS (`KH_OLLAMA_URL=http://ollama:11434/v1/`), persisting accounts, default RAG settings, and chat threads (`ktem_app_data/user_data/sql.db`), uploaded PDFs (`user_data/files/`), embeddings (`user_data/vectorstore/` & `docstore/`), evidence chunk coordinates (`chunks_cache_dir/`), and `.nv` CUDA cache under `/opt/aimx/ai-service/`.

---

## 2. Conference Live Demo Walkthroughs

### Workflow A: External `ai-stack-testbed` Source (`ai-gpu-01`) -> Multi-Cloud Target GPU

Use this workflow when the 10-node source environment has been provisioned via `../mc-terrarium/examples/aws/ai-stack-testbed` (`tofu apply` + `./setup-ai-stack.sh`).

#### Step 1. Verify Source `Kotaemon` (`:7860`) & Build RAG History

1. In `mc-terrarium/examples/aws/ai-stack-testbed`, open the URL from `tofu output -json demo_endpoints | jq -r .source_kotaemon_url` and sign in (`admin` / `admin`).
2. Upload PDF documents in **`Files` -> `File Collection`** (`Docling` parser, Korean response language, `inline` citation, and `Mindmap` are pre-seeded in `sql.db`). Confirm indexing completion in the `note` column (`{"tokens": ..., "loader": "DoclingReader"}`).
3. In **`Chat`**, select **`Search In File(s)`** for document-specific questions (or **`Search All`** across all files), ask questions in Korean, and verify inline footnotes (`[1]`, `[2]`), right-panel `[Preview]` highlighted evidence chunks, and mindmap diagrams.

#### Step 2. Prepare External Testbed Config & Recommend Target Cloud GPU

From the `cm-beetle` repository root:

```bash
cp -n cmd/aimx/testconf/template-auth-config.json cmd/aimx/testconf/auth-config.json
cp -n cmd/aimx/testconf/template-ai-stack-testbed-config.yaml cmd/aimx/testconf/ai-stack-testbed-config.yaml

# Inspect live ai-gpu-01 hardware/GPU/LLM tags and get zone-aware Target GPU candidates
make test-aimx EXTRA_ARGS="-config testconf/ai-stack-testbed-config.yaml -mode recommend"
```

> **Note**: When running CM-Beetle inside the `cloud-migrator` VM (`10.0.1.0/24`), set `source.connectionsFile` in `ai-stack-testbed-config.yaml` to `source_connections_private.json` (copied to `~/cm-beetle/` by `./setup-cloud-migrator.sh`).

#### Step 3. Provision Target GPU, Migrate Stack, & Verify Integrity

```bash
# Provision Target GPU VM, quiesce ai-gpu-01 (prepare-for-migration.sh), transfer /opt/aimx/ai-service via transx, & warm up Target GPU
make test-aimx EXTRA_ARGS="-config testconf/ai-stack-testbed-config.yaml -mode migrate"

# Verify Target GPU tok/s, Kotaemon HTTP (:7860), and 100% identical sql.db SHA-256, then write Markdown report
make test-aimx EXTRA_ARGS="-config testconf/ai-stack-testbed-config.yaml -mode verify"
```

- **Incremental Delta Re-Sync**: If you add or clean up chat threads or settings on Source `ai-gpu-01` after initial migration, re-running `-mode migrate` reuses the running Target GPU VM, stops Target containers cleanly before `transx` delta transfer (~2s for `sql.db`), restarts Kotaemon + Ollama, and re-verifies SHA-256 parity.

---

### Workflow B: Standalone CB-Tumblebug Source (`mc-daisy`) -> Target GPU

```bash
# 1) Provision Source GPU VM (g4dn.xlarge T4 16GB) & deploy Kotaemon + Ollama (gemma4:e2b + bge-m3)
make test-aimx-kotaemon EXTRA_ARGS="-mode setup-source"

# 2) Upload PDFs & ask questions at http://<SOURCE_PUBLIC_IP>:7860 (admin / admin)

# 3) Recommend, migrate, and verify Target Cloud GPU stack
make test-aimx-kotaemon EXTRA_ARGS="-mode recommend"
make test-aimx-kotaemon EXTRA_ARGS="-mode migrate"
make test-aimx-kotaemon EXTRA_ARGS="-mode verify"

# 4) Cleanup Target (add -cleanup-source=true to also delete mc-daisy)
make test-aimx-kotaemon EXTRA_ARGS="-mode cleanup"
```

---

### Workflow C: Interactive `Beetle UX Lab` UI Migration

1. Import `source_connections_public.{csv,json,yaml}` (or `source_connections_private.{csv,json,yaml}` on `cloud-migrator`) or upload `cmd/aimx/testresult/source-infra-discovered.json` in **1. Source Extraction**.
2. Recommend and provision the Target Cloud GPU infrastructure (`7860` & `11434` security group rules auto-mapped) in **2. Target Cloud Optimizer**.
3. In **Data Transfer Center**, click **Launch New Data Migration** -> **`+ Preset: AIMX (Ollama + Kotaemon)`** to quiesce Source, stop any running Target containers, transfer `/opt/aimx/ai-service` (`ktem_app_data/**`, `ollama/**`, `docker-compose.yml`), ensure Target GPU drivers/Docker/`nvidia-ctk`, and auto-warm up the models registered in `docker-compose.yml`.

---

## 3. Key Configuration Options (`kotaemon-config.yaml` / `ai-stack-testbed-config.yaml`)

| Section     | Key                     | Example / Preset                | Description                                                                                                             |
| :---------- | :---------------------- | :------------------------------ | :---------------------------------------------------------------------------------------------------------------------- |
| `test`      | `cases[].preferredZone` | `""` (optional)                 | Optional AZ override (e.g., `ap-northeast-2c`); defaults to `AvailableZones[0]` on zonal CSPs.                          |
| `source`    | `connectionsFile`       | `".../source_connections_public.json"` | Optional path to `mc-terrarium` JSON connection file; enables External SSH Source mode instead of CB-Tumblebug lookup. |
| `source`    | `nodeName`              | `"ai-gpu-01"`                   | Entry `name` to select inside `source.connectionsFile` (supports direct SSH inspection and `transx` transfer).          |
| `aiService` | `appType`               | `"kotaemon"`                    | Explicit app selector (`"kotaemon"` or `"open-webui"`, required without fallback).                                      |
| `aiService` | `rootDir`               | `"/opt/aimx/ai-service"`        | Host root directory migrated end-to-end via `transx`.                                                                   |
| `aiService` | `modelName`             | `"auto"` or `"gemma4:12b"`      | Generation LLM tag (`"auto"` reads `LOCAL_MODEL` from Source `docker-compose.yml`; `gemma4:12b` for 24GB, `gemma4:e2b` for 16GB). |
| `aiService` | `embeddingModel`        | `"auto"` or `"bge-m3"`          | Multilingual embedding model tag (`"auto"` reads `LOCAL_MODEL_EMBEDDINGS` from Source `docker-compose.yml`).            |
| `aiService` | `appPort`               | `7860`                          | Host TCP port mapped to Kotaemon (`7860`) or Open WebUI (`3000`).                                                       |
| `aiService` | `transferModelWeights`  | `true`                          | Transfers `/opt/aimx/ai-service/ollama` GGUF blobs directly via `transx` without re-pulling from the internet.          |
| `aiService` | `useDataEncryption`     | `true`                          | Encrypts ephemeral SSH private keys using CM-Beetle's 1-time RSA-OAEP/AES-GCM KeyStore API.                             |
