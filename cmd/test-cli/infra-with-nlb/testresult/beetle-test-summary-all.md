# CM-Beetle integration test summary (with NLB)

> [!IMPORTANT]
> This document provides an overall summary of automated integration test results for all provider-region pairs.

## Execution details

- **Test Date**: November 3, 2026
- **Start Time**: 16:41:08 KST
- **End Time**: 18:01:11 KST
- **Total Execution Duration**: 19m34s
- **CM-Beetle Version**: v0.6.3+ (d67a8e3)
- **imdl Version**: v0.1.19+ (d67a8e3)
- **CB-Tumblebug Version**: v0.13.8
- **CB-Spider Version**: v0.13.20
- **CB-MapUI Version**: v0.13.14

## High-level test status

| Metric | Count | Description |
|--------|-------|-------------|
| **Total CSP Pairs** | **10** | Number of unique CSP-Region configurations evaluated |
| Passed CSP Pairs | 7 | Pairs where all test steps succeeded |
| Failed CSP Pairs | 0 | Pairs where at least one test step failed |
| Skipped CSP Pairs | 3 | Pairs that were disabled in config |
| **Total Test Steps** | **150** | Total individual endpoint tests triggered |
| Passed Steps | 105 | Individual tests that succeeded |
| Failed Steps | 0 | Individual tests that failed |

## Provider-specific summary

| Provider-Region | Status | Duration | Steps Passed | Details |
|-----------------|--------|----------|--------------|---------|
| **AWS-Seoul** | ✅ **PASS** | 7m24s | 15 / 15 | [View Report](beetle-test-results-aws.md) |
| **Alibaba-Seoul** | ✅ **PASS** | 4m18s | 15 / 15 | [View Report](beetle-test-results-alibaba.md) |
| **Azure-Busan** | ✅ **PASS** | 13m24s | 15 / 15 | [View Report](beetle-test-results-azure.md) |
| **GCP-Seoul** | ✅ **PASS** | 11m12s | 15 / 15 | [View Report](beetle-test-results-gcp.md) |
| **IBMCloud-Sydney** | ✅ **PASS** | 19m34s | 15 / 15 | [View Report](beetle-test-results-ibm.md) |
| **NHNCloud-Pangyo** | ✅ **PASS** | 8m50s | 15 / 15 | [View Report](beetle-test-results-nhn.md) |
| **Tencent-Seoul** | ✅ **PASS** | 6m17s | 15 / 15 | [View Report](beetle-test-results-tencent.md) |
