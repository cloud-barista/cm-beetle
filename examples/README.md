# CM-Beetle Migration Examples

This directory provides practical, ready-to-run examples and curl scripts for **Object Storage** and **Data Migration** using CM-Beetle REST APIs.

These examples are designed for teams seeking to automate bucket provisioning and data transfers while official portal features are in active development (e.g., prior to the full rollout of the CM-Centipede data migration framework).

---

## End-to-End Migration Flow

```text
┌─────────────────────────────────────────────────────────────────────────────┐
│ Step 1: Bucket Infrastructure Provisioning                                   │
│ Directory: examples/object-storage-migration/                               │
│                                                                             │
│  [1. Input Source KT Info] ──> [2. Recommend for NCP] ──> [3. Create Bucket]│
│  (kt-test-source-bucket)        (warnings: CORS disabled)     (test-os-01)  │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
                                       ▼ Target bucket 'test-os-01' ready
┌─────────────────────────────────────────────────────────────────────────────┐
│ Step 2: Data Transfer & Migration                                           │
│ Directory: examples/data-migration/                                         │
│                                                                             │
│  [1. Get One-time Key]   ──> [2. Encrypt Credentials] ──> [3. Relay Transfer]│
│  (RSA-OAEP Key Bundle)       (MinIO S3 Access/Secret)      (KT kr1 -> NCP kr)│
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Prerequisites & Pre-flight Health Check

Before running any examples, ensure:
1. CM-Beetle and CB-Tumblebug are running (`docker compose up -d` or `make up`).
2. The necessary CLI tools are installed: `curl`, `jq`, and `bash`.

Verify CM-Beetle is running and ready with a single curl command:

```bash
curl -s http://localhost:8056/beetle/readyz
```
- **Expected Output:** `{"message":"CM-Beetle is ready"}`

---

## Migration Workflow Roadmap

Follow the tutorials in order using the step-by-step curl commands documented in each directory:

1. **Step 1 — Provision Target Cloud Bucket ([`object-storage-migration/`](./object-storage-migration/)):**
   - **Recommendation**: Submit source KT Cloud bucket properties (`POST /beetle/recommendation/middleware/objectStorage`) to generate an NCP-compatible bucket recommendation.
   - **Provisioning**: Provision the recommended bucket in CB-Tumblebug with late-binding naming (`POST /beetle/migration/middleware/ns/mig01/objectStorage?nameSeed=test`), creating `test-os-01`.
   - *(Keep `test-os-01` active to proceed to Step 2.)*

2. **Step 2 — Transfer Data into Target Bucket ([`data-migration/`](./data-migration/)):**
   - **One-time RSA Key**: Fetch a fresh public key bundle (`GET /beetle/migration/data/encryptionKey`).
   - **In-Transit Encryption**: Encrypt sensitive credentials (`POST /beetle/migration/data/test/encrypt`).
   - **Data Transfer**: Launch the asynchronous transfer job (`POST /beetle/migration/data` with `Prefer: respond-async`) to migrate objects from KT Cloud into `test-os-01`.
   - **Job Tracking**: Poll the migration status (`GET /beetle/request/{reqId}`).

> [!NOTE]
> Each directory also includes an illustrative reference script (`run.sh`) showing how these steps can be automated in shell scripts. Users should follow the individual curl commands in the respective READMEs to observe and verify intermediate states.

---

## Example Directories

| Directory | Description | Key APIs Demonstrated | API Documentation |
| :--- | :--- | :--- | :--- |
| **[`object-storage-migration/`](./object-storage-migration/)** | Target bucket recommendation & infrastructure provisioning *(Requires source bucket info)* | `POST /beetle/recommendation/middleware/objectStorage`<br>`POST /beetle/migration/middleware/ns/:nsId/objectStorage` | [Recommendation API](https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BRecommendation%5D%20Managed%20Object%20Storage)<br>[Migration API](https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BMigration%5D%20Managed%20Object%20Storage) |
| **[`data-migration/`](./data-migration/)** | Data transfer from Source KT Cloud kr1 (via MinIO SDK) to Destination NCP kr bucket (sync/async) | `POST /beetle/migration/data`<br>`GET /beetle/request/:reqId` | [Data Migration API](https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BMigration%5D%20Data%20(incubating)) |

> [!NOTE]
> **Prerequisites and Preparation**:
> - **For Bucket Recommendation & Provisioning**: Source bucket metadata (name, versioning, CORS, etc.) must be provided in the JSON payload (or discovered via CM-Beetle's temporary Scan/Inspect test APIs).
> - **For Data Migration**: The target bucket **must be provisioned beforehand** (using `object-storage-migration`), and valid S3 read credentials for the source storage are required.

---

## Official API Documentation Links

- **Object Storage Recommendation API:**  
  [https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BRecommendation%5D%20Managed%20Object%20Storage](https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BRecommendation%5D%20Managed%20Object%20Storage)
- **Object Storage Migration API:**  
  [https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BMigration%5D%20Managed%20Object%20Storage](https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BMigration%5D%20Managed%20Object%20Storage)
- **Data Migration API (incubating):**  
  [https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BMigration%5D%20Data%20(incubating)](https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BMigration%5D%20Data%20(incubating))

---

## Alternative: Using Beetle UX Lab (GUI Test Portal)

In addition to using the REST APIs and CLI scripts provided here, you can also use **Beetle UX Lab**, the built-in testing dashboard:

1. Open your browser and navigate to:
   ```
   http://localhost:8056/beetle/ui/
   ```
2. Key features available in the portal:
   - **Object Storage Migration Lifecycle**: 22-step workflow for scanning source buckets, inspecting objects, generating recommendation models, and provisioning target buckets.
   - **Data Transfer Center**: Visual monitoring and execution of data migration tasks across storage systems.
