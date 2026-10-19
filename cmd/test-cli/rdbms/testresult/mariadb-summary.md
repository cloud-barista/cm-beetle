# CM-Beetle Managed RDBMS (MariaDB) Test Run Summary

- **Run Timestamp:** 2026-10-19 11:12:27
- **Run Duration:** 19m55s
- **Active Test Cases in this Run:** 3

## 1. Current Run Results (Tested in this Run)

| Test Items / Phase                       |     **AWS**      |   **ALIBABA**    |    **NHN**    |
| :--------------------------------------- | :--------------: | :--------------: | :-----------: |
| **Region**                               | `ap-northeast-2` | `ap-northeast-2` |     `kr1`     |
| **Duration**                             |      19m45s      |      16m27s      |    18m20s     |
| **Pre-flight Spec & Image Review**       |        🟢        |        🟢        |      🟢       |
| **Create Pre-requisite Infra (VNet/SG)** |        🟢        |        🟢        |      🟢       |
| **Get RDBMS Support**                    |        🟢        |        🟢        |      🟢       |
| **Get RDBMS Capability**                 |        🟢        |        🟢        |      🟢       |
| **Recommend RDBMS**                      |        🟢        |        🟢        |      🟢       |
| **Validate Recommendation**              |        🟢        |        🟢        |      🟢       |
| **Migrate RDBMS (Provisioning)**         |        🟢        |        🟢        |      🟢       |
| **Get RDBMS Info & List**                |        🟢        |        🟢        |      🟢       |
| **Create Database**                      |        🟢        |        🟢        |      🟢       |
| **Get Secure Transport**                 |        🟢        |        🟢        |      🟢       |
| **External Data I/O**                    |        🟢        |        🟢        |      🟢       |
| **Internal Data I/O**                    |        🟢        |        🟢        |      🟢       |
| **Delete Database**                      |        🟢        |        🟢        |      🟢       |
| **Delete RDBMS**                         |        🟢        |        🟢        |      🟢       |
| **Delete SecurityGroup**                 |        🟢        |        🟢        |      🟢       |
| **Delete VNet**                          |        🟢        |        🟢        |      🟢       |
| **Overall Result**                       |  **✅ PASSED**   |  **✅ PASSED**   | **✅ PASSED** |

## 2. Cumulative Multi-Cloud Matrix (Overall Compatibility Status)

- **Matrix Scope:** All Supported MariaDB Providers (4 Providers)
- **Last Matrix Update:** 2026-10-19 11:12:27

| Test Items / Phase                       |        **AWS**        |      **ALIBABA**      |        **NHN**        |  **OPENSTACK**  |
| :--------------------------------------- | :-------------------: | :-------------------: | :-------------------: | :-------------: |
| **Region**                               |   `ap-northeast-2`    |   `ap-northeast-2`    |         `kr1`         |   `RegionOne`   |
| **Last Verified**                        | `2026-10-19 10:54:05` | `2026-10-19 10:54:05` | `2026-10-19 10:54:05` |        —        |
| **Current Status**                       |      🟢 Verified      |      🟢 Verified      |      🟢 Verified      |   ⏸️ Excluded   |
| **Pre-flight Spec & Image Review**       |          🟢           |          🟢           |          🟢           |        —        |
| **Create Pre-requisite Infra (VNet/SG)** |          🟢           |          🟢           |          🟢           |        —        |
| **Get RDBMS Support**                    |          🟢           |          🟢           |          🟢           |        —        |
| **Get RDBMS Capability**                 |          🟢           |          🟢           |          🟢           |        —        |
| **Recommend RDBMS**                      |          🟢           |          🟢           |          🟢           |        —        |
| **Validate Recommendation**              |          🟢           |          🟢           |          🟢           |        —        |
| **Migrate RDBMS (Provisioning)**         |          🟢           |          🟢           |          🟢           |        —        |
| **Get RDBMS Info & List**                |          🟢           |          🟢           |          🟢           |        —        |
| **Create Database**                      |          🟢           |          🟢           |          🟢           |        —        |
| **Get Secure Transport**                 |          🟢           |          🟢           |          🟢           |        —        |
| **External Data I/O**                    |          🟢           |          🟢           |          🟢           |        —        |
| **Internal Data I/O**                    |          🟢           |          🟢           |          🟢           |        —        |
| **Delete Database**                      |          🟢           |          🟢           |          🟢           |        —        |
| **Delete RDBMS**                         |          🟢           |          🟢           |          🟢           |        —        |
| **Delete SecurityGroup**                 |          🟢           |          🟢           |          🟢           |        —        |
| **Delete VNet**                          |          🟢           |          🟢           |          🟢           |        —        |
| **Overall Result**                       |        **✅**         |        **✅**         |        **✅**         | **⏸️ EXCLUDED** |

> [!NOTE]
> **Notes on Excluded / Pending Test Cases**:
>
> - **OPENSTACK**: Pending Spider OpenStack MariaDB < 10.5 Secure Transport (cb-spider#1869) and privilege grant fix (cb-spider#1866) (`execute: false`)

---

_Generated by CM-Beetle Managed RDBMS Test CLI_
