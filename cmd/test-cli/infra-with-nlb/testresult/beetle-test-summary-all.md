# CM-Beetle integration test summary (with NLB)

> [!IMPORTANT]
> This document provides an overall summary of automated integration test results for all provider-region pairs.

## Execution details

- **Test Date**: October 12, 2026
- **Start Time**: 21:19:03 KST
- **End Time**: 21:25:13 KST
- **Total Execution Duration**: 6m9s
- **CM-Beetle Version**: v0.6.1+ (3649ed7)
- **imdl Version**: v0.1.16+ (3649ed7)
- **CB-Tumblebug Version**: unknown
- **CB-Spider Version**: unknown
- **CB-MapUI Version**: v0.13.13

## High-level test status

| Metric | Count | Description |
|--------|-------|-------------|
| **Total CSP Pairs** | **10** | Number of unique CSP-Region configurations evaluated |
| Passed CSP Pairs | 1 | Pairs where all test steps succeeded |
| Failed CSP Pairs | 0 | Pairs where at least one test step failed |
| Skipped CSP Pairs | 9 | Pairs that were disabled in config |
| **Total Test Steps** | **150** | Total individual endpoint tests triggered |
| Passed Steps | 15 | Individual tests that succeeded |
| Failed Steps | 0 | Individual tests that failed |

## Provider-specific summary

| Provider-Region | Status | Duration | Steps Passed | Details |
|-----------------|--------|----------|--------------|---------|
| **Tencent-Seoul** | ✅ **PASS** | 6m4s | 15 / 15 | [View Report](beetle-test-results-tencent.md) |

