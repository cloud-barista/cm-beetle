# CM-Beetle Managed RDBMS (MySQL) Test Run Summary

- **Access Mode:** Public / Default (per test-config.yaml)
- **Run Timestamp:** 2026-11-25 19:17:57
- **Run Duration:** 24m49s
- **Active Test Cases in this Run:** 9

## 1. Current Run Results (Tested in this Run)

| Test Items / Phase                       |     **AWS**      |   **AZURE**    |    **GCP**    |   **ALIBABA**    |  **TENCENT**  |    **IBM**    |    **NCP**    |    **NHN**    | **OPENSTACK** |
| :--------------------------------------- | :--------------: | :------------: | :-----------: | :--------------: | :-----------: | :-----------: | :-----------: | :-----------: | :-----------: |
| **Region**                               | `ap-northeast-2` | `koreacentral` | `us-central1` | `ap-northeast-2` |  `ap-seoul`   |  `us-south`   |     `kr`      |     `kr1`     |  `RegionOne`  |
| **Duration**                             |      18m32s      |     12m48s     |    13m38s     |      19m25s      |    12m30s     |    17m31s     |    23m52s     |    24m39s     |     5m53s     |
| **Pre-flight Spec & Image Review**       |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      🟢       |
| **Create Pre-requisite Infra (VNet/SG)** |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      🟢       |
| **Get RDBMS Support**                    |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      🟢       |
| **Get RDBMS Capability**                 |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      🟢       |
| **Recommend RDBMS**                      |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      🟢       |
| **Validate Recommendation**              |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      🟢       |
| **Migrate RDBMS (Provisioning)**         |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      🔴       |
| **Get RDBMS Info & List**                |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      ⚪       |
| **Create Database**                      |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      ⚪       |
| **Get Secure Transport**                 |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      ⚪       |      🟢       |      ⚪       |
| **External Data I/O**                    |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |       —       |      🟢       |      ⚪       |
| **Internal Data I/O**                    |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      ⚪       |
| **Delete Database**                      |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      ⚪       |
| **Delete RDBMS**                         |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      🟢       |
| **Delete SecurityGroup**                 |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      🟢       |
| **Delete VNet**                          |        🟢        |       🟢       |      🟢       |        🟢        |      🟢       |      🟢       |      🟢       |      🟢       |      🟢       |
| **Overall Result**                       |  **✅ PASSED**   | **✅ PASSED**  | **✅ PASSED** |  **✅ PASSED**   | **✅ PASSED** | **✅ PASSED** | **✅ PASSED** | **✅ PASSED** | **❌ FAILED** |

## 2. Cumulative Multi-Cloud Matrix (Overall Compatibility Status)

- **Matrix Scope:** All Supported MySQL Providers (9 Providers)
- **Last Matrix Update:** 2026-11-25 19:17:57

| Test Items / Phase                       |        **AWS**        |       **AZURE**       |        **GCP**        |      **ALIBABA**      |      **TENCENT**      |        **IBM**        |        **NCP**        |        **NHN**        |     **OPENSTACK**     |
| :--------------------------------------- | :-------------------: | :-------------------: | :-------------------: | :-------------------: | :-------------------: | :-------------------: | :-------------------: | :-------------------: | :-------------------: |
| **Region**                               |   `ap-northeast-2`    |    `koreacentral`     |     `us-central1`     |   `ap-northeast-2`    |      `ap-seoul`       |      `us-south`       |         `kr`          |         `kr1`         |      `RegionOne`      |
| **Last Verified**                        | `2026-11-25 18:54:32` | `2026-11-25 18:54:32` | `2026-11-25 18:54:32` | `2026-11-25 18:54:32` | `2026-11-25 18:54:32` | `2026-11-25 18:54:32` | `2026-11-25 18:54:32` | `2026-11-25 18:54:32` | `2026-11-25 18:54:32` |
| **Current Status**                       |      🟢 Verified      |      🟢 Verified      |      🟢 Verified      |      🟢 Verified      |      🟢 Verified      |      🟢 Verified      |      🟢 Verified      |      🟢 Verified      |      🟢 Verified      |
| **Pre-flight Spec & Image Review**       |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |
| **Create Pre-requisite Infra (VNet/SG)** |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |
| **Get RDBMS Support**                    |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |
| **Get RDBMS Capability**                 |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |
| **Recommend RDBMS**                      |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |
| **Validate Recommendation**              |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |
| **Migrate RDBMS (Provisioning)**         |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🔴           |
| **Get RDBMS Info & List**                |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          ⚪           |
| **Create Database**                      |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          ⚪           |
| **Get Secure Transport**                 |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          ⚪           |          🟢           |          ⚪           |
| **External Data I/O**                    |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |           —           |          🟢           |          ⚪           |
| **Internal Data I/O**                    |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          ⚪           |
| **Delete Database**                      |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          ⚪           |
| **Delete RDBMS**                         |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |
| **Delete SecurityGroup**                 |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |
| **Delete VNet**                          |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |          🟢           |
| **Overall Result**                       |        **✅**         |        **✅**         |        **✅**         |        **✅**         |        **✅**         |        **✅**         |        **✅**         |        **✅**         |        **❌**         |

---

_Generated by CM-Beetle Managed RDBMS Test CLI_
