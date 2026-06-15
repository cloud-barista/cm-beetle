# CM-Beetle Test Suite Summary

> [!NOTE]
> This document summarizes all CM-Beetle integration tests from a single test run.

## Test Run Information

- **Date**: June 15, 2026
- **Start Time**: 20:17:43 KST
- **Total Duration**: 20m4.948s
- **Execution Mode**: parallel
- **CM-Beetle**: v0.5.1+ (412492f)
- **CB-Tumblebug**: v0.12.15
- **CB-Spider**: v0.12.30

## Overall Results

**Status**: ✅ All Passed

| Metric | Count |
|--------|-------|
| Total CSP-Region Pairs | 10 |
| Passed Pairs | 6 |
| Failed Pairs | 0 |
| Skipped Pairs | 4 |
| Total API Tests | 90 |
| Passed API Tests | 54 |
| Failed API Tests | 0 |

## Per CSP-Region Pair Results

| # | Display Name | CSP | Region | Status | Duration | Passed | Failed | Skipped |
|---|--------------|-----|--------|--------|----------|--------|--------|--------|
| 1 | AWS-Seoul | AWS | ap-northeast-2 | ✅ PASS | 2m51.693s | 9/9 | 0/9 | 0/9 |
| 2 | Azure-Busan | AZURE | koreasouth | ✅ PASS | 5m20.981s | 9/9 | 0/9 | 0/9 |
| 3 | GCP-Seoul | GCP | asia-northeast3 | ✅ PASS | 19m54.526s | 9/9 | 0/9 | 0/9 |
| 4 | Alibaba-Seoul | ALIBABA | ap-northeast-2 | ✅ PASS | 2m52.083s | 9/9 | 0/9 | 0/9 |
| 5 | IBMCloud-Sydney | IBM | au-syd | ✅ PASS | 6m24.302s | 9/9 | 0/9 | 0/9 |
| 6 | NCP-Seoul | NCP | kr | ✅ PASS | 14m13.68s | 9/9 | 0/9 | 0/9 |

## Detailed Results Per CSP-Region Pair

### 1. AWS-Seoul (AWS, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-15 20:17:48 KST
- **Duration**: 2m51.693s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 2.658s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 47.865s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 44ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 13ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 23ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.39s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.28s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 47.88s |

### 2. Azure-Busan (AZURE, koreasouth) — ✅ PASS

- **Start Time**: 2026-06-15 20:17:51 KST
- **Duration**: 5m20.981s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 7.63s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 1m54.598s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 254ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 7ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 38ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.314s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.307s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 1m56.115s |

### 3. GCP-Seoul (GCP, asia-northeast3) — ✅ PASS

- **Start Time**: 2026-06-15 20:17:53 KST
- **Duration**: 19m54.526s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 15.431s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 9m10.659s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 33ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 4ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 21ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.352s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.326s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 9m2.729s |

### 4. Alibaba-Seoul (ALIBABA, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-15 20:17:55 KST
- **Duration**: 2m52.083s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 3.289s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 1m9.903s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 51ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 4ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 19ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.46s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.326s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 26.319s |

### 5. IBMCloud-Sydney (IBM, au-syd) — ✅ PASS

- **Start Time**: 2026-06-15 20:17:57 KST
- **Duration**: 6m24.302s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 17.574s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 2m57.526s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 612ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 3ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 27ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.361s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.372s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 1m28s |

### 6. NCP-Seoul (NCP, kr) — ✅ PASS

- **Start Time**: 2026-06-15 20:17:58 KST
- **Duration**: 14m13.68s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 2m27.677s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 8m6.362s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 1.218s |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 5ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 21ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.339s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.363s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 2m13.196s |

---

*Generated by CM-Beetle Test CLI on 2026-06-15 20:17:43 KST*
