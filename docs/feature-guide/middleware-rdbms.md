# Managed RDBMS (RDS) Feature Guide

## Overview

CM-Beetle provides end-to-end **Managed RDBMS (Relational Database Service)** recommendation, validation, and migration capabilities for heterogeneous multi-cloud environments. This feature empowers users to transition on-premises or existing cloud databases to optimal managed database services (e.g., AWS RDS, Google Cloud SQL, Azure Database, Alibaba ApsaraDB, Tencent CDB, IBM Cloud Databases, NCP Cloud DB, NHN RDS) with zero guesswork.

Key capabilities include:

- **Capability-Driven Recommendation**: Evaluates source database parameters (engine, version, vCPU, memory, storage) and maps them to the optimal target CSP DB instance spec and configuration using real-time CSP capability metrics.
- **Pre-flight & Referential Validation**: Performs multi-tier dry-run validations against CSP API rules and referential integrity (e.g., matching VNet/Subnet/SecurityGroup connection profiles) before touching cloud resources.
- **Automated Migration & Provisioning**: Seamlessly deploys managed database clusters and instances via CB-Tumblebug with robust state tracking.
- **Logical Database Lifecycle Management**: Supports creation, listing, and deletion of logical tenant databases inside the provisioned RDBMS instance.
- **Data I/O Verification**: Supports end-to-end connectivity and SQL I/O verification for both public-accessible and private-isolated database topologies.

---

## Architecture & Workflow

The Managed RDBMS workflow spans from initial capability discovery to post-migration database operations:

```mermaid
flowchart TD
    subgraph Phase1["1. Discovery & Recommendation"]
        A["Source DB Metadata"] --> B["GET /recommendation/middleware/rdbms/support"]
        B --> C["GET /recommendation/middleware/rdbms/capability"]
        C --> D["POST /recommendation/middleware/rdbms"]
        D --> E["Recommended RDBMS Spec"]
    end

    subgraph Phase2["2. Validation & Autofill"]
        E --> F["POST /recommendation/middleware/rdbms/validate"]
        F --> G{"Referential & CSP Dry-run Check"}
        G -->|Pass| H["Migration-Ready Payload"]
        G -->|Fail| I["Actionable Diagnostics"]
    end

    subgraph Phase3["3. Migration & Provisioning"]
        H --> J["POST /migration/middleware/ns/{nsId}/rdbms"]
        J --> K["CB-Tumblebug & Cloud Provider Provisioning"]
        K --> L["RDBMS Instance Ready"]
    end

    subgraph Phase4["4. Database Operations & Verification"]
        L --> M["POST /migration/.../database - Create DB"]
        M --> N["Internal / External SQL Data I/O Verification"]
        N --> O["DELETE /migration/.../rdbms - Clean-up"]
    end
```

---

## CSP Support Matrix

CM-Beetle Managed RDBMS supports multi-cloud database provisioning across 9 major Cloud Service Providers (CSPs), with dedicated verification for **MySQL** and **MariaDB**.

### 1. MySQL Support Matrix

MySQL is verified and supported across all 9 CSPs:

| CSP           | Service Brand                    | Supported Versions     |   Access Mode    | Subnet Requirement                              | Provisioning Characteristics                                              |
| :------------ | :------------------------------- | :--------------------- | :--------------: | :---------------------------------------------- | :------------------------------------------------------------------------ |
| **AWS**       | Amazon RDS for MySQL             | 8.4, 8.0, 5.7          | Public / Private | Multi-AZ Subnet Group (>= 2 subnets across AZs) | High availability and automated failover                                  |
| **Azure**     | Azure Database for MySQL         | 8.4, 8.0 (8.0.21), 5.7 | Public / Private | Single or Multi-Subnet                          | Flexible server architecture; supports 8.4 & 9.5                          |
| **GCP**       | Google Cloud SQL for MySQL       | 8.4, 8.0, 5.7          | Public / Private | Authorized Networks / VPC Peering               | Highly performant storage scaling                                         |
| **Alibaba**   | Alibaba Cloud ApsaraDB for MySQL | 8.4, 8.0, 5.7          | Public / Private | VNet Subnet Binding                             | Broad engine version choices                                              |
| **Tencent**   | Tencent Cloud CDB for MySQL      | 8.4, 8.0, 5.7, 5.6     | Public / Private | Multi-AZ VPC Subnet Group                       | Fast regional provisioning                                                |
| **IBM**       | IBM Cloud Databases for MySQL    | 8.4 (replaces 8.0)     | Public / Private | Resource Group / VPC bound                      | `multitenant` (~9-10m) vs dedicated `b3c.*` (35-45+m) hosting models      |
| **NCP**       | NAVER Cloud DB for MySQL         | 8.0 (8.0.36)           |   Private Only   | Single Subnet + Dedicated DB Port               | Private access enforced; verified via Internal Runner VM                  |
| **NHN**       | NHN Cloud RDS for MySQL          | 8.4, 8.0, 5.7          | Public / Private | Standard VPC Subnet                             | Dedicated DB Security Group required (nhnDBSGToAllowAllInbound supported) |
| **OpenStack** | OpenStack Trove (MySQL)          | 5.7.29                 | Public / Private | Standard VPC Subnet / Floating IP               | Flexible private cloud deployment; verified external & internal SQL I/O   |

### 2. MariaDB Support Matrix

MariaDB is natively supported across 4 CSPs:

| CSP           | Service Brand                      | Supported Versions            |   Access Mode    | Subnet Requirement                              | Provisioning Characteristics                      |
| :------------ | :--------------------------------- | :---------------------------- | :--------------: | :---------------------------------------------- | :------------------------------------------------ |
| **AWS**       | Amazon RDS for MariaDB             | 10.5, 10.6, 10.11, 11.4, 11.8 | Public / Private | Multi-AZ Subnet Group (>= 2 subnets across AZs) | Automated failover, parameter group customization |
| **Alibaba**   | Alibaba Cloud ApsaraDB for MariaDB | 10.6, 10.3                    | Public / Private | VNet Subnet Binding                             | High-availability primary/standby architecture    |
| **NHN**       | NHN Cloud RDS for MariaDB          | 10.6, 10.11, 11.4 (10.3~11.8) | Public / Private | Standard VPC Subnet                             | AppKey credential integration required            |
| **OpenStack** | OpenStack Trove (MariaDB)          | 10.4                          | Public / Private | Standard VPC Subnet / Floating IP               | Reference version 10.4 in RegionOne datastore     |

> [!NOTE]
> **Automatic Compatibility Fallback**:
> For CSPs that do not provide native MariaDB managed services (**Azure**, **GCP**, **Tencent**, **IBM**, **NCP**, **KT**), CM-Beetle's recommendation engine detects the absence of MariaDB support via CB-Tumblebug capability metadata and automatically recommends **MySQL** as an API-compatible fallback with actionable warnings.

---

## API Reference

### 1. Discovery & Recommendation APIs

#### 1.1 Get CSP Support Matrix

Retrieve supported database engines, features, and operation methods across all or specific CSPs.

- **Endpoint**: `GET /recommendation/middleware/rdbms/support`
- **Query Parameter**:
  - `providerName` (optional): Filter by CSP provider (e.g., `aws`, `azure`, `gcp`, `ibm`, `ncp`, `nhn`).

**Sample Request**:

```http
GET /beetle/recommendation/middleware/rdbms/support?providerName=aws HTTP/1.1
Host: localhost:8056
Authorization: Basic <credentials>
```

**Sample Response**:

```json
{
  "resourceType": "rdbms",
  "supports": {
    "aws": {
      "supported": true,
      "supportedDBEngines": ["mysql", "mariadb"],
      "dbOperationMethod": "sqlFallback",
      "supportsTag": true,
      "storageTypeSelectable": true
    }
  }
}
```

---

#### 1.2 Get Real-time Capability Options

Fetch live DB instance spec catalogs, supported engine versions, and storage configurations for a given connection.

- **Endpoint**: `GET /recommendation/middleware/rdbms/capability`
- **Query Parameter**:
  - `connectionName` (required): Target cloud connection name (e.g., `aws-ap-northeast-2`).

**Sample Response**:

```json
{
  "resourceType": "rdbms",
  "connectionName": "aws-ap-northeast-2",
  "supportedDBEngines": ["mysql", "mariadb"],
  "dbmsRequirements": {
    "mysql": {
      "referenceDBInstanceSpec": "db.t3.medium",
      "rootDiskType": "gp2",
      "rootDiskSize": "100"
    }
  },
  "dbInstanceSpecs": [
    {
      "cspSpecName": "db.t3.medium",
      "vCPU": 2,
      "memoryGiB": 4.0
    }
  ]
}
```

---

#### 1.3 Recommend Managed RDBMS

Recommends optimal target RDBMS specifications based on source database workloads and desired cloud properties.

- **Endpoint**: `POST /recommendation/middleware/rdbms`
- **Request Body**: `RecommendRDBMSRequest`

**Sample Request**:

```json
{
  "desiredCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
  },
  "autoFillSourceDefaults": true,
  "targetPreferences": {
    "adminUserName": "dbadmin",
    "highAvailability": true,
    "publicAccess": true,
    "backupRetentionDays": 7
  },
  "sourceRDBMSInstances": [
    {
      "displayName": "source-mysql-01",
      "description": "Production user database instance",
      "dbNode": {
        "hostname": "db-server-01",
        "cpu": {
          "cpus": 1,
          "cores": 4,
          "threads": 4
        },
        "memory": {
          "totalSize": 16
        },
        "rootDisk": {
          "label": "/",
          "totalSize": 50,
          "type": "SSD"
        },
        "dataDisks": [
          {
            "label": "/data",
            "totalSize": 50,
            "type": "SSD"
          }
        ]
      },
      "dbEngine": {
        "engine": "mysql",
        "engineVersion": "8.0",
        "port": 3306,
        "role": "primary"
      },
      "innerDatabases": [
        {
          "databaseName": "userdb",
          "characterSet": "utf8mb4",
          "collation": "utf8mb4_unicode_ci"
        }
      ]
    }
  ]
}
```

**Sample Response**:

```json
{
  "desiredCloud": {
    "provider": "aws",
    "region": "ap-northeast-2"
  },
  "recommendedRDBMS": [
    {
      "sourceInstanceId": "source-mysql-01",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db.m5.large",
      "storageType": "gp3",
      "storageSize": 100,
      "highAvailability": true,
      "publicAccess": true
    }
  ]
}
```

---

### 2. Validation & Autofill API

#### Validate RDBMS Recommendation

Executes pre-flight referential integrity checks and dry-run validation without provisioning cloud resources.

- **Endpoint**: `POST /recommendation/middleware/rdbms/validate`
- **Query Parameter**:
  - `nsId` (optional): Namespace identifier for referential resource checks.
- **Request Body**: `RDBMSCreateRequest`

**Validation Checks Performed**:

1. **Connection Consistency**: Verifies that the VNet, Subnet(s), and Security Group(s) belong to the same CSP connection profile as the target RDBMS.
2. **Engine & Version Compatibility**: Checks if the specified DB engine version is supported by the target CSP capability catalog.
3. **Multi-AZ Subnet Requirements**: Validates whether the CSP requires multiple subnets across distinct Availability Zones (e.g., AWS, Tencent).
4. **Access Constraint Alignment**: Adjusts and validates `publicAccess` flags for CSPs with private-only policies (e.g., NCP).

---

### 3. Migration & Management APIs

#### 3.1 Migrate (Provision) Managed RDBMS

Provisions the recommended Managed RDBMS instance in the specified namespace.

- **Endpoint**: `POST /migration/middleware/ns/{nsId}/rdbms`
- **Query Parameter**:
  - `nameSeed` (optional): Prefix seed for deterministic instance naming.

**Sample Request**:

```json
{
  "name": "prod-rdbms-mysql",
  "connectionName": "aws-ap-northeast-2",
  "vNetId": "vpc-01",
  "subnetIds": ["subnet-1a", "subnet-1c"],
  "securityGroupIds": ["sg-db-01"],
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "db.t3.medium",
  "storageType": "gp3",
  "storageSize": 100,
  "adminUserName": "dbadmin",
  "adminUserPassword": "SecurePassword123!",
  "publicAccess": true,
  "highAvailability": false
}
```

---

#### 3.2 Logical Database Management (CRUD)

Once the managed RDBMS is provisioned, logical tenant databases can be created, inspected, and deleted.

> [!NOTE]
> **Zero-Credential Security Policy**:
> Neither CB-Tumblebug nor CM-Beetle persists or stores database administrator credentials.
> All logical database management requests require callers to explicitly provide credentials via HTTP headers:
>
> - `X-Admin-User-Name`: Database master/administrator username
> - `X-Admin-User-Password`: Database master/administrator password

- **Create Database**: `POST /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database`
  - **Headers**: `X-Admin-User-Name: <username>`, `X-Admin-User-Password: <password>`
  - **Body**:
    ```json
    {
      "databaseName": "customerdb"
    }
    ```
- **List Databases**: `GET /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database`
  - **Headers**: `X-Admin-User-Name: <username>`, `X-Admin-User-Password: <password>`
- **Delete Database**: `DELETE /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database/{databaseName}`
  - **Headers**: `X-Admin-User-Name: <username>`, `X-Admin-User-Password: <password>`

#### Private Mode (`publicAccess=false`) Fast-Fail Policy

> [!NOTE]
> **Version Requirement**: Private RDBMS Fast-Fail (`HTTP 422 Unprocessable Entity`) and Private NHN `nhnDBSGToAllowAllInbound` support are available starting from **CB-Tumblebug v0.13.9** and later.

CM-Beetle aligns with CB-Tumblebug (`cb-tumblebug#2829`) by separating Control Plane API operations from Data Plane SQL/TLS connections. When an RDBMS instance is provisioned with `publicAccess=false`, CB-Spider cannot reach private DB ports (3306, 5432) from outside the target VPC/VNet. CB-Tumblebug returns `HTTP 422 Unprocessable Entity` immediately, and CM-Beetle preserves and propagates this `422` status to the caller:

| CM-Beetle API / Step                                                     | Private RDBMS (`publicAccess=false`) Behavior                                                                                                                                                                                  | Recommended Action (Data Migration / Ops)                                                                                                                                                                                                                                                                     |
| :----------------------------------------------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `POST /recommendation/middleware/rdbms`                                  | Preserves original `charset`/`collation` for Tencent and Azure (no silent mutation); emits `warnings` for AWS/IBM, Tencent (`collation`), and Azure (`charset`-only). NCP/NHN omit `charset`/`collation` for native REST API.  | Review `warnings` and choose between in-VPC DB creation (exact spec) or adjusting options for native API.                                                                                                                                                                                                     |
| `POST /migration/middleware/ns/{nsId}/rdbms`                             | **AWS, IBM**: Provisions the instance and **skips** remote `/database` creation with a warning log.<br/>**Other CSPs**: Invokes `/database` via Tumblebug (if `422` occurs, logs warning and completes instance provisioning). | Create logical databases inside the target VPC/VNet via Bastion, `POST /migration/data` (`transx` hooks), or CM-Centipede.                                                                                                                                                                                    |
| `GET .../rdbms/{rdbmsId}/secure-transport`                               | **Fast-Fail (`HTTP 422`)** on all CSPs                                                                                                                                                                                         | Inspect TLS enforcement (`require_secure_transport`, `Ssl_cipher`) and CA certificates from inside the VPC/VNet, or connect with `--ssl-mode=REQUIRED` / CSP official Root CA bundle.                                                                                                                         |
| `POST/GET/DELETE .../rdbms/{rdbmsId}/database` (AWS, IBM)                | **Fast-Fail (`HTTP 422`)** (no native logical DB API)                                                                                                                                                                          | Manage logical databases from inside the VPC/VNet (Bastion, `transx` hooks, or CM-Centipede).                                                                                                                                                                                                                 |
| `POST .../rdbms/{rdbmsId}/database` (Azure, Tencent, NCP, NHN)           | **Allowed** on the native API path; **Fast-Fail (`HTTP 422`)** when parameters trigger direct SQL fallback                                                                                                                     | - **Azure**: specify both `charset` and `collation` (or omit both), or create inside VNet.<br/>- **Tencent**: omit `collation` for native API, or create inside VPC to preserve exact collation.<br/>- **NCP / NHN**: omit both `charset` and `collation` for native API, or run `ALTER DATABASE` inside VPC. |
| `POST/GET/DELETE .../rdbms/{rdbmsId}/database` (GCP, Alibaba, OpenStack) | **Allowed** (100% CSP Native API; Alibaba MariaDB accepts `charset` only)                                                                                                                                                      | No extra action needed for DB CRUD.                                                                                                                                                                                                                                                                           |

---

#### 3.3 Secure Transport & Server CA Certificate Inspection

Inspects live TLS/SSL enforcement status, active cipher suite, and server CA certificate from the cloud provider.

> [!IMPORTANT]
> **Private RDBMS (`publicAccess=false`)**: Across all CSPs, `GET /secure-transport` requires direct TLS/SQL probing to the database port and immediately returns `HTTP 422 Unprocessable Entity`. Inspect TLS enforcement and server CA certificates from inside the target VPC/VNet (e.g., via a Bastion host).

- **Endpoint**: `GET /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/secure-transport`
- **Headers**:
  - `X-Admin-User-Name`: Database master/administrator username
  - `X-Admin-User-Password`: Database master/administrator password
- **Sample Response**:
  ```json
  {
    "status": "SUCCESS",
    "message": "Successfully retrieved secure transport info",
    "data": {
      "engine": "mysql",
      "requireSecureTransport": "ON",
      "enforced": true,
      "rules": "REQUIRE SSL",
      "tlsInUse": true,
      "tlsCipher": "ECDHE-RSA-AES128-GCM-SHA256",
      "caCertificate": {
        "pem": "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----",
        "subject": "CN=rds.amazonaws.com",
        "issuer": "CN=Amazon Root CA 1",
        "notAfter": "2027-10-08T00:00:00Z",
        "isSelfSigned": false
      },
      "recommendedSSLMode": "VERIFY_IDENTITY"
    }
  }
  ```

---

#### 3.4 Delete Managed RDBMS

Terminates and removes the RDBMS instance.

- **Endpoint**: `DELETE /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}`
- **Query Parameter**:
  - `option=force`: Forces deletion if cloud-side instances are in an error or intermediate state.

---

## Post-Migration Data Transfer Guide (Public vs Private RDBMS)

After provisioning the target managed RDBMS instance via CM-Beetle, schema and data migration can be performed either with **CM-Beetle's built-in Data Migration API (`transx`)** or with **CM-Centipede**. The workflow differs depending on `publicAccess`:

### Option 1: CM-Beetle Data Migration API (`POST /beetle/migration/data` / `transx` Hooks)

For **Private RDBMS (`publicAccess=false`)**, external hosts cannot connect directly to the target database port. Use a VM inside the target VPC/VNet (such as a Bastion host or migrated application node) as the SSH destination for `POST /beetle/migration/data` with symmetric `preCmd`/`postCmd` hooks:

1. **Stage 1 (`source.preCmd`)**: Export databases on the source server with `mysqldump --databases` so `CREATE DATABASE /*!32312 IF NOT EXISTS*/ ... CHARACTER SET ... COLLATE ...` statements are embedded in the dump file.
2. **Stage 2 (`Transfer`)**: `transx` transfers the dump file to the target in-VPC VM over SSH (`rsync`) or via Object Storage staging.
3. **Stage 3 (`destination.postCmd`)**: Restore the dump from the in-VPC VM into the private RDBMS endpoint over the private network, then remove temporary dump files.

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/tmp/rdbms-transfer",
    "preCmd": "mkdir -p /tmp/rdbms-transfer && mysqldump -h 127.0.0.1 -u root -p\"$SRC_DB_PASS\" --databases appdb --single-transaction --routines --triggers > /tmp/rdbms-transfer/dump.sql",
    "postCmd": "rm -f /tmp/rdbms-transfer/dump.sql",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "source-db.example.com",
        "port": 22,
        "username": "ubuntu",
        "privateKey": "<encrypted>"
      }
    }
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/tmp/rdbms-transfer",
    "preCmd": "mkdir -p /tmp/rdbms-transfer && command -v mysql >/dev/null || command -v mariadb >/dev/null",
    "postCmd": "mysql -h <private-rdbms-endpoint> -P 3306 -u dbadmin -p\"$DST_DB_PASS\" --ssl-mode=REQUIRED < /tmp/rdbms-transfer/dump.sql && rm -f /tmp/rdbms-transfer/dump.sql",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "<target-vpc-bastion-or-app-vm-ip>",
        "port": 22,
        "username": "cb-user",
        "privateKey": "<encrypted>"
      }
    }
  },
  "strategy": "auto"
}
```

> [!TIP]
> **Security & Pre-Created Default Databases**:
>
> - Encrypt SSH keys and credentials before calling `POST /beetle/migration/data` using `GET /beetle/migration/data/encryptionKey` (see [Data Transfer Feature Guide](data-transfer.md)). Avoid putting raw DB passwords in shell command history; prefer a temporary `0600` `.my.cnf` file or environment setup on the node.
> - On **NCP** and **NHN**, CM-Beetle's recommendation omits `charset`/`collation` so the native REST API creates an empty database with server defaults. If exact source `charset`/`collation` must be preserved, execute `ALTER DATABASE <db> CHARACTER SET <cs> COLLATE <col>;` inside the VPC before restoring tables.

### Option 2: CM-Centipede Integration Contract

When **CM-Centipede** orchestrates RDBMS data migration on top of CM-Beetle:

| Integration Point          | `publicAccess=true` (Public)                                                                            | `publicAccess=false` (Private)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| :------------------------- | :------------------------------------------------------------------------------------------------------ | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Endpoint Discovery**     | Call `GET /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}` to read `endpoint` and `publicAccess`.       | Same. Detect `publicAccess == false` and route SQL operations via an in-VPC Bastion / SSH tunnel / in-VPC agent.                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| **TLS & CA Certificate**   | Call `GET .../secure-transport` with admin headers; apply `recommendedSSLMode` and `caCertificate.pem`. | `GET .../secure-transport` returns **`HTTP 422`**. Treat `422` as expected (do not fail with 500); connect inside the VPC with `--ssl-mode=REQUIRED` or CSP official Root CA bundle, or inspect `SHOW VARIABLES LIKE 'require_secure_transport'` inside the VPC.                                                                                                                                                                                                                                                                                                                     |
| **Logical DB Preparation** | `POST .../database` works across all CSPs.                                                              | - **AWS, IBM**: `POST .../database` returns **`HTTP 422`**; create logical DBs inside the VPC (`CREATE DATABASE IF NOT EXISTS ... CHARACTER SET ... COLLATE ...`).<br/>- **Tencent (`collation`) / Azure (`charset`-only)**: `POST .../database` returns **`HTTP 422`**; create inside the VPC to preserve exact charset/collation (or omit `collation` on Tencent / supply explicit `collation` on Azure for native API).<br/>- **NCP, NHN**: Native API creates DB with server default charset/collation; run `ALTER DATABASE` inside the VPC if source charset/collation differs. |

---

## Best Practices & Operational Guidelines

1. **Pre-requisite Infrastructure Alignment**:
   - Always ensure pre-requisite Virtual Networks (`VNet`), `Subnets`, and `Security Groups` are provisioned prior to issuing the RDBMS migration request.
   - For AWS and Tencent, ensure at least **two subnets in different Availability Zones** are attached to satisfy CSP subnet group requirements.
   - For **Azure Private (`publicAccess=false`)**, provision at least **two subnets** in the VNet: one subnet for application/Bastion VMs and a separate empty subnet delegated exclusively to Azure Database for MySQL Flexible Server.
2. **Private Network Isolation & Access Control (All CSPs)**:
   - Provision a Bastion or target application VM in the same VNet so both `transx` (`destination.postCmd`) and CM-Centipede can reach the private DB endpoint on port 3306/5432.
   - **NCP ACG Default Deny Policy**: NCP Cloud DB ACGs default to **0 inbound rules (Default Deny)**. Even internal traffic from a VM within the same VPC is blocked on port 3306 unless an inbound ACG rule is configured in NCP Console (`Database > Cloud DB for MySQL > ACG`) or `ncpDBACGToAllowAllInbound: true` is set (test/dev only).
   - **NHN Cloud DB Security Group**: `nhnDBSGToAllowAllInbound: true` is supported for both `publicAccess=true` and `publicAccess=false` (test/dev only; restrict inbound rules to the VPC CIDR in production).
   - **GCP / IBM Private Constraints**: Automated GCP Private IP Peering is not yet automated in CB-Spider (`cb-tumblebug#2829`), and IBM Cloud Databases typically requires `publicAccess=true` for standard remote operations.
3. **Provisioning Timeouts & Asynchronous Execution**:
   - Cloud database provisioning typically takes 5 to 15 minutes across most CSPs.
   - For **IBM Cloud Databases**:
     - The **`multitenant`** hosting model provisions in approximately **9 to 10 minutes** and is recommended by default.
     - Dedicated/encrypted host flavors (e.g., **`b3c.4x16.encrypted`**) take **35 to 45+ minutes** due to dedicated bare-metal/VPC hardware slice allocation and storage encryption.
   - Ensure client and reverse-proxy timeouts are configured with at least **50 minutes** if calling synchronously with dedicated host flavors, or execute via the asynchronous API (`Prefer: respond-async`).
4. **Storage Capacity Metric Standardization**:
   - Storage size ranges and capabilities are standardized to 10-decimal Gigabytes (`storageSizeRangeGB`, $10^9$ bytes) per CB-Spider issue #1820. Validation and UI sliders should align with `storageSizeRangeGB` returned from capability APIs.
