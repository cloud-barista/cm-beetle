# CM-Beetle integration test summary (with NLB)

> [!IMPORTANT]
> This document provides an overall summary of automated integration test results for all provider-region pairs.

## Execution details

- **Test Date**: October 30, 2026
- **Start Time**: 21:45:19 KST
- **End Time**: 22:06:39 KST
- **Total Execution Duration**: 21m19s
- **CM-Beetle Version**: v0.6.2+ (83efe14)
- **imdl Version**: v0.1.19+ (83efe14)
- **CB-Tumblebug Version**: v0.13.8
- **CB-Spider Version**: v0.13.20
- **CB-MapUI Version**: v0.13.14

## High-level test status

| Metric | Count | Description |
|--------|-------|-------------|
| **Total CSP Pairs** | **10** | Number of unique CSP-Region configurations evaluated |
| Passed CSP Pairs | 7 | Pairs where all test steps succeeded |
| Failed CSP Pairs | 1 | Pairs where at least one test step failed |
| Skipped CSP Pairs | 2 | Pairs that were disabled in config |
| **Total Test Steps** | **150** | Total individual endpoint tests triggered |
| Passed Steps | 108 | Individual tests that succeeded |
| Failed Steps | 1 | Individual tests that failed |
| Skipped Steps | 11 | Tests skipped due to pre-requisite step failure |

## Provider-specific summary

| Provider-Region | Status | Duration | Steps Passed | Details |
|-----------------|--------|----------|--------------|---------|
| **AWS-Seoul** | ✅ **PASS** | 6m59s | 15 / 15 | [View Report](beetle-test-results-aws.md) |
| **Alibaba-Seoul** | ✅ **PASS** | 4m53s | 15 / 15 | [View Report](beetle-test-results-alibaba.md) |
| **Azure-Busan** | ✅ **PASS** | 13m57s | 15 / 15 | [View Report](beetle-test-results-azure.md) |
| **GCP-Seoul** | ✅ **PASS** | 10m53s | 15 / 15 | [View Report](beetle-test-results-gcp.md) |
| **IBMCloud-Sydney** | ✅ **PASS** | 18m44s | 15 / 15 | [View Report](beetle-test-results-ibm.md) |
| **NCP-Seoul** | ✅ **PASS** | 10m39s | 15 / 15 | [View Report](beetle-test-results-ncp.md) |
| **NHNCloud-Pangyo** | ❌ **FAIL** | 48s | 3 / 15 | [View Report](beetle-test-results-nhn.md) |
| **Tencent-Seoul** | ✅ **PASS** | 6m21s | 15 / 15 | [View Report](beetle-test-results-tencent.md) |

