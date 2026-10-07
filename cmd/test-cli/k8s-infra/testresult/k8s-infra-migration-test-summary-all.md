# CM-Beetle K8s Infra Migration Test Summary

- CM-Beetle Version: v0.6.1+ (cd1f3a2)
- Test Date: 2026-10-07 16:41:05 KST

| Target | CSP / Region | Steps Passed | Cluster ID | Result |
|--------|--------------|--------------|------------|--------|
| AWS-Seoul | aws / ap-northeast-2 | 8/8 | `mig01-on-prem-k8s-cluster` | ✅ PASS |
| Azure-Busan | azure / koreasouth | 8/8 | `mig02-on-prem-k8s-cluster` | ✅ PASS |
| GCP-Seoul | gcp / asia-northeast3 | 8/8 | `mig03-on-prem-k8s-cluster` | ✅ PASS |
| Alibaba-Seoul | alibaba / ap-northeast-2 | 8/8 | `mig04-on-prem-k8s-cluster` | ✅ PASS |
| Tencent-Seoul | tencent / ap-seoul | 7/8 | `mig05-on-prem-k8s-cluster` | ❌ FAIL |
| IBMCloud-Sydney | ibm / au-syd | 8/8 | `mig06-on-prem-k8s-cluster` | ✅ PASS |
| NCP-Seoul | ncp / kr | 8/8 | `mig07-on-prem-k8s-cluster` | ✅ PASS |
| NHNCloud-Pangyo | nhn / kr1 | 8/8 | `mig08-on-prem-k8s-cluster` | ✅ PASS |

## Per-CSP Reports

- [AWS-Seoul](k8s-infra-migration-test-results-aws.md)
- [Azure-Busan](k8s-infra-migration-test-results-azure.md)
- [GCP-Seoul](k8s-infra-migration-test-results-gcp.md)
- [Alibaba-Seoul](k8s-infra-migration-test-results-alibaba.md)
- [Tencent-Seoul](k8s-infra-migration-test-results-tencent.md)
- [IBMCloud-Sydney](k8s-infra-migration-test-results-ibm.md)
- [NCP-Seoul](k8s-infra-migration-test-results-ncp.md)
- [NHNCloud-Pangyo](k8s-infra-migration-test-results-nhn.md)
