# AIMX (AI Migration X): AI Service Migration Prototype CLI

**AIMX** (**AI Migration X**, _AI Migration eXecutor_) is an end-to-end Proof-of-Concept (PoC) CLI tool that automates **AI Semiconductor (GPU) Infrastructure & AI Model/Application Migration** by orchestrating **CB-Tumblebug** and **CM-Beetle (`transx`)** REST APIs.

It supports two stateful AI web applications out of the box via `aiService.appType`:
- **`kotaemon` (Recommended for Conference RAG Demo)**: **`Kotaemon` (`ghcr.io/cinnamon/kotaemon:main-lite`, `:7860`) + `Ollama` (`gemma4:e2b` + `bge-m3`)** featuring an integrated PDF slide viewer with yellow citation highlighting.
- **`open-webui`**: **`Open WebUI` (`ghcr.io/open-webui/open-webui:main`, `:3000`) + `Ollama` (`gemma4:e2b`)**.

---

## 1. Overview & 4-Layer Migration Scope (Source -> Target 1:1 Mapping)

`AIMX` migrates a live GPU-accelerated AI service stack across cloud environments in four layers:

1. **Layer 1: AI Semiconductor (GPU) VM Infrastructure & Cloud Block Storage**:
   - **Source**: AWS GPU VM (`mc-daisy`, `aws+ap-northeast-2+g4dn.xlarge`, `1x NVIDIA Tesla T4 16GB VRAM`, 4 vCPU, 16 GiB RAM, Ubuntu 22.04 Deep Learning AMI `ami-00ffffe72a332c865`, 100GB EBS Root Disk) provisioned via CB-Tumblebug's `POST /tumblebug/ns/{nsId}/infraDynamic` (or via `mc-terrarium/examples/aws/ai-stack-testbed`).
   - **Target**: Recommended via CM-Beetle's GPU-aware recommender (`POST /beetle/recommendation/infra`) and provisioned via CM-Beetle (`POST /beetle/migration/ns/{nsId}/infra`).
2. **Layer 2: AI Semiconductor Runtime**:
   - NVIDIA GPU Driver, CUDA Runtime, Docker Engine + Compose v2, and `nvidia-container-toolkit` (`driver: nvidia`, `count: all`).
3. **Layer 3: Self-Hosted LLM & Multilingual Embedding Volume (`Ollama` + `Gemma 4` + `bge-m3`)**:
   - **Ollama (`ollama/ollama:latest`, `:11434`)** serving **`gemma4:e2b`** (generation) and **`bge-m3`** (Korean/English conference slide embeddings) co-resident in 16GB GPU VRAM (`OLLAMA_MAX_LOADED_MODELS=2`), persisted under `/opt/aimx/ai-service/ollama`.
4. **Layer 4: Stateful RAG Web Application & Database Volume (`Kotaemon` + `SQLite` + `ChromaDB` / `LanceDB`)**:
   - **Kotaemon (`:7860`)** connected to Ollama via internal Docker DNS (`KH_OLLAMA_URL=http://ollama:11434/v1/`), persisting accounts, chat threads (`ktem_app_data/user_data/sql.db`), uploaded conference PDFs (`user_data/files/`), embeddings (`user_data/vectorstore/` & `docstore/`), and slide highlight coordinates (`chunks_cache_dir/`).

---

## 2. Conference Live Demo Walkthrough (`Kotaemon` + `gemma4:e2b` + `bge-m3`)

### Step 1. Provision Source GPU VM & Deploy `Kotaemon` + `Ollama`

From the repository root (`cm-beetle`):

```bash
# Option A: Using the dedicated Makefile alias
make test-aimx-kotaemon EXTRA_ARGS="-mode setup-source"

# Option B: Passing -config explicitly to test-aimx
make test-aimx EXTRA_ARGS="-config testconf/kotaemon-config.yaml -mode setup-source"
```

### Step 2. Upload Conference Slides & Build Citation History (Browser)

1. Open `http://<SOURCE_PUBLIC_IP>:7860` and sign in (`admin` / `admin`).
2. In the **File Index** tab, upload Cloud-Barista conference presentation PDF files and wait for indexing (`bge-m3` GPU embeddings).
3. Switch to the **Chat** tab, select the indexed files, and ask 2-3 questions in Korean or English.
4. Confirm that **chat history** is saved on the left and **original PDF slides with yellow citation highlights** appear on the right viewer.

### Step 3. Recommend Target Cloud GPU & Migrate Entire Stack (`CLI` or `Beetle UX Lab` UI)

#### Path A: Automated CLI Migration (`recommend` -> `migrate` -> `verify`)

```bash
# 1) Inspect live Source GPU specs (generates cmd/aimx/testresult/source-infra-discovered.json) & get Target GPU candidates
make test-aimx-kotaemon EXTRA_ARGS="-mode recommend"

# 2) Quiesce Source containers, record sql.db SHA-256, transfer /opt/aimx/ai-service via transx, & warm up Target GPU
make test-aimx-kotaemon EXTRA_ARGS="-mode migrate"

# 3) Verify Target GPU tok/s, Kotaemon HTTP (:7860), and 100% identical sql.db SHA-256, then write Markdown report
make test-aimx-kotaemon EXTRA_ARGS="-mode verify"
```

#### Path B: Interactive `Beetle UX Lab` UI Migration

1. Upload `cmd/aimx/testresult/source-infra-discovered.json` (or extract from `ai-stack-testbed` via Honeybee) in **1. Source Extraction**.
2. Recommend and provision the Target Cloud GPU infrastructure (`7860` & `11434` security group rules auto-mapped) in **2. Target Cloud Optimizer**.
3. In **Data Transfer Center**, click **Launch New Data Migration** -> **`+ Preset: AIMX (Ollama + Kotaemon)`** to transfer `/opt/aimx/ai-service` (`ktem_app_data/**`, `ollama/**`, `docker-compose.yml`) and auto-start/warm up the Target GPU stack.

### Step 4. Live Target Verification & Follow-Up Question (Browser)

1. Open `http://<TARGET_PUBLIC_IP>:7860` and sign in (`admin` / `admin`).
2. Verify that **uploaded conference PDFs, previous chat threads, and yellow slide citation highlights** are immediately intact without re-downloading models or re-indexing PDFs.
3. Ask a **follow-up question** inside the migrated thread to demonstrate live Target GPU (`Ollama`) inference and new slide citation highlights.

### Step 5. Cleanup

```bash
# Delete Target GPU infrastructure only
make test-aimx-kotaemon EXTRA_ARGS="-mode cleanup"

# Delete both Target and Source GPU infrastructure
make test-aimx-kotaemon EXTRA_ARGS="-mode cleanup -cleanup-source=true"
```

---

## 3. Key Configuration Options (`testconf/kotaemon-config.yaml`)

| Section     | Key                    | Kotaemon Preset                | Description                                                                                      |
| :---------- | :--------------------- | :----------------------------- | :----------------------------------------------------------------------------------------------- |
| `test`      | `cases[].preferredZone`| `""` (optional)                | Optional AZ override (e.g., `ap-northeast-2a`); defaults to `AvailableZones[0]` on zonal CSPs.   |
| `aiService` | `appType`              | `"kotaemon"`                   | Explicit app selector (`"kotaemon"` or `"open-webui"`, required without fallback).               |
| `aiService` | `rootDir`              | `"/opt/aimx/ai-service"`       | Host root directory migrated end-to-end via `transx`.                                            |
| `aiService` | `modelName`            | `"gemma4:e2b"`                 | Generation LLM tag (`gemma4:e2b` recommended on 16GB T4; `gemma4:e4b` optional on >=24GB GPUs).  |
| `aiService` | `embeddingModel`       | `"bge-m3"`                     | Multilingual embedding model tag pulled into Ollama for Korean/English PDF indexing.             |
| `aiService` | `appPort`              | `7860`                         | Host TCP port mapped to Kotaemon (`7860`) or Open WebUI (`3000`).                                |
| `aiService` | `transferModelWeights` | `true`                         | Transfers `/opt/aimx/ai-service/ollama` GGUF blobs directly via `transx` without re-pulling.     |
| `aiService` | `useDataEncryption`    | `true`                         | Encrypts ephemeral SSH private keys using CM-Beetle's 1-time RSA-OAEP/AES-GCM KeyStore API.      |
