# CM-Beetle integration test summary (with NLB)

> [!IMPORTANT]
> This document provides an overall summary of automated integration test results for all provider-region pairs.

## Execution details

- **Test Date**: November 25, 2026
- **Start Time**: 18:14:08 KST
- **End Time**: 18:34:11 KST
- **Total Execution Duration**: 21m13s
- **CM-Beetle Version**: v0.6.5+ (83361e8)
- **imdl Version**: v0.1.22+ (83361e8)
- **CB-Tumblebug Version**: v0.14.0
- **CB-Spider Version**: v0.14.0
- **CB-MapUI Version**: v0.14.0

## High-level test status

| Metric | Count | Description |
|--------|-------|-------------|
| **Total CSP Pairs** | **10** | Number of unique CSP-Region configurations evaluated |
| Passed CSP Pairs | 8 | Pairs where all test steps succeeded |
| Failed CSP Pairs | 0 | Pairs where at least one test step failed |
| Skipped CSP Pairs | 2 | Pairs that were disabled in config |
| **Total Test Steps** | **150** | Total individual endpoint tests triggered |
| Passed Steps | 120 | Individual tests that succeeded |
| Failed Steps | 0 | Individual tests that failed |

## Provider-specific summary

| Provider-Region | Status | Duration | Steps Passed | Details |
|-----------------|--------|----------|--------------|---------|
| **AWS-Seoul** | ✅ **PASS** | 7m8s | 15 / 15 | [View Report](beetle-test-results-aws.md) |
| **Alibaba-Seoul** | ✅ **PASS** | 5m3s | 15 / 15 | [View Report](beetle-test-results-alibaba.md) |
| **Azure-Busan** | ✅ **PASS** | 10m53s | 15 / 15 | [View Report](beetle-test-results-azure.md) |
| **GCP-Seoul** | ✅ **PASS** | 12m46s | 15 / 15 | [View Report](beetle-test-results-gcp.md) |
| **IBMCloud-Sydney** | ✅ **PASS** | 18m38s | 15 / 15 | [View Report](beetle-test-results-ibm.md) |
| **NCP-Seoul** | ✅ **PASS** | 9m37s | 15 / 15 | [View Report](beetle-test-results-ncp.md) |
| **NHNCloud-Pangyo** | ✅ **PASS** | 14m27s | 15 / 15 | [View Report](beetle-test-results-nhn.md) |
| **Tencent-Seoul** | ✅ **PASS** | 6m25s | 15 / 15 | [View Report](beetle-test-results-tencent.md) |

