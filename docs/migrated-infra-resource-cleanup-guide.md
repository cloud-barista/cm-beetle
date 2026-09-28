# 마이그레이션된 인프라 자원 조회 및 삭제 가이드

본 문서는 CM-Beetle을 통해 배포된 마이그레이션된 인프라 및 개별 인프라 자원(VM, Security Group, SSH Key, vNet, Subnet)의 **기본 조회 및 순차 삭제 절차**와, CSP(NCP, AWS 등) 콘솔에서 자원이 먼저 삭제되었거나 의존성 오류로 정상 삭제가 불가능할 때 사용하는 **Force 삭제 절차**를 안내합니다.

> [!IMPORTANT]
> **🚨 핵심 주의사항**
>
> 1. **Platform을 통한 자원 삭제 원칙**:  
>    Cloud-Migrator Platform을 통해 마이그레이션된 인프라 자원은 반드시 **Platform(API 또는 UI)에서 삭제**해야 합니다. CSP(NCP, AWS, Azure, GCP 등) 콘솔에서 자원을 임의로 직접 삭제할 경우, 내부 메타데이터와 실제 CSP 자원 간의 상태 불일치가 발생하여 플랫폼에서의 정상 삭제가 실패(`404 Not Found`, 의존성 오류 등)하게 됩니다.
> 2. **Force 삭제 시 CSP Console 크로스 체크 필수**:  
>    **Force 삭제는 플랫폼 내부 메타데이터의 강제 정리를 의미하며, CSP 상의 실제 자원 삭제 완료 여부를 보장하지 않습니다.** CSP API 오류나 제약으로 인해 CSP 콘솔에 고아 자원이 잔존하여 예상치 못한 과금이 발생할 위험이 있으므로, **Force 삭제를 수행한 후에는 반드시 대상 CSP 콘솔(예: NCP 콘솔)에 접속하여 실제 자원이 완전히 삭제되었는지 직접 크로스 체크(교차 확인)**해야 합니다.

---

## 1. 사전 준비 및 기본 정보

### 1.1 접속 정보 및 파라미터 기본값

- **Beetle API Base URL**: `http://localhost:8056/beetle` (원격 환경의 경우 IP 또는 도메인 변경 필요)
- **Basic Auth**: `-u "default:default"`
- **Namespace**: `mig01`
- **Infra ID**: `infra101`

> **⚠️ API 경로 안내: MCI ➔ Infra**  
> CM-Beetle 최신 API 명세에서는 기존 `/mci` 경로가 **`/infra`**로 표준화되었습니다.  
> 마이그레이션된 전체 인프라(VM 세트) 조회 및 삭제 시 `/migration/ns/{nsId}/infra` 엔드포인트를 사용합니다.

### 1.2 자원 삭제 권장 순서

클라우드 자원 간의 종속성으로 인해 반드시 아래 순서대로 삭제를 진행해야 합니다.

```text
[1단계] Infra
   └── VM이 서브넷, 보안 그룹, SSH 키를 점유하고 있으므로 가장 먼저 종료 및 삭제
[2단계] Security Group
   └── VM 네트워크 인터페이스가 해제된 후 보안 그룹 삭제
[3단계] SSH Key
   └── VM 접근용 키페어 삭제
[4단계] vNet 및 Subnet
   └── 상위 VPC 및 내부 서브넷 최종 삭제
```

---

## 2. 마이그레이션된 인프라 자원 조회 및 정상 삭제

일반적인 환경에서는 아래의 표준 명령어를 통해 자원 상태를 조회하고 순차적으로 정상 삭제합니다.

### 2.1 자원 목록 조회

```bash
# 1. Infra 목록 조회
curl -s -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/infra' \
  -H 'accept: application/json'

# 2. Security Group 목록 조회
curl -s -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/securityGroup' \
  -H 'accept: application/json'

# 3. SSH Key 목록 조회
curl -s -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/sshKey' \
  -H 'accept: application/json'

# 4. vNet 목록 조회
curl -s -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/vNet' \
  -H 'accept: application/json'
```

---

### 2.2 순차적 정상 삭제

#### Step 1: Infra 정상 삭제 (option=terminate)

VM 인스턴스를 종료하고 인프라를 삭제합니다.

```bash
curl -s -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/infra/infra101?option=terminate' \
  -H 'accept: application/json'
```

_인프라 삭제 요청 후 CSP 상에서 VM 인스턴스가 완전히 정리될 때까지 잠시 대기합니다(약 10~30초)._

#### Step 2: Security Group 삭제

```bash
# 네임스페이스 내 보안 그룹 전체 일괄 삭제
curl -s -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/securityGroup' \
  -H 'accept: application/json'

# 또는 특정 보안 그룹 단건 삭제 (예: mig-sg-01)
curl -s -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/securityGroup/mig-sg-01' \
  -H 'accept: application/json'
```

#### Step 3: SSH Key 삭제

```bash
# 특정 SSH Key 단건 삭제 (예: mig-sshkey-01)
curl -s -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/sshKey/mig-sshkey-01' \
  -H 'accept: application/json'
```

#### Step 4: vNet 및 Subnet 삭제

`action=withsubnets` 옵션을 사용하여 vNet과 연관된 하위 서브넷을 함께 삭제합니다.

```bash
curl -s -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/vNet/mig-vnet-01?action=withsubnets' \
  -H 'accept: application/json'
```

---

## 3. 마이그레이션된 잔여 인프라 자원 조회 및 Force 삭제

> [!WARNING]
> **CSP Console 교차 확인 필수**  
> Force 삭제는 CSP 실제 자원의 강제 삭제를 시도하고, 성공/실패 여부에 관계 없이 플랫폼 내부 메타데이터를 강제로 삭제합니다. 삭제되지 않은 VM 인스턴스, 공인 IP, 디스크 등이 남아 지속적인 요금이 청구될 수 있으므로, **Force 삭제 완료 후 반드시 해당 CSP 콘솔(NCP, AWS 등)에서 자원이 실제로 삭제되었는지 직접 크로스 체크**해야 합니다.

### 3.1 Force 삭제가 필요한 경우

- **CSP 콘솔에서 직접 수동 삭제한 경우**: NCP, AWS 등 콘솔에서 VM이나 리소스를 먼저 삭제하여 Beetle/Tumblebug 메타데이터와 불일치가 발생하고 정상 삭제 시 `404 Not Found` 또는 통신 오류가 날 때
- **Dependency Lock 발생 시**: 자원 간 연결 정보가 남아있어 삭제가 차단될 때
- **Spider 또는 CSP 오류 지속 시**: CSP API 타임아웃이나 응답 지연으로 정상 삭제 사이클이 실패할 때

---

### 3.2 잔여 자원 ID 확인

Force 삭제를 진행하기 전, 시스템에 남아있는 정확한 자원 ID를 조회합니다.

```bash
# 특정 인프라 상세 조회 (존재 여부 및 상태 확인)
curl -s -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/infra/infra101' \
  -H 'accept: application/json'

# 잔여 Security Group ID 목록 조회 (option=id 사용 시 ID 목록만 반환)
curl -s -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/securityGroup?option=id' \
  -H 'accept: application/json'

# 잔여 SSH Key ID 목록 조회
curl -s -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/sshKey?option=id' \
  -H 'accept: application/json'

# 잔여 vNet ID 목록 조회
curl -s -u "default:default" -X 'GET' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/vNet?option=id' \
  -H 'accept: application/json'
```

---

### 3.3 Force 삭제 실행

#### Step 1: Infra 강제 삭제 (`option=force`)

CSP 상에 VM이 이미 없거나 오류가 발생해도 시스템 내부 메타데이터를 강제로 삭제합니다.

```bash
curl -s -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/infra/infra101?option=force' \
  -H 'accept: application/json'
```

#### Step 2: Security Group 강제 삭제 (`force=true`)

리소스 의존성 검사를 우회하고 Spider 및 CSP 실패 여부와 상관없이 내부 레코드를 즉시 삭제합니다.

```bash
# 전체 Security Group 일괄 강제 삭제
curl -s -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/securityGroup?force=true' \
  -H 'accept: application/json'

# 또는 특정 Security Group 단건 강제 삭제
curl -s -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/securityGroup/mig-sg-01?force=true' \
  -H 'accept: application/json'
```

#### Step 3: SSH Key 강제 삭제 (`force=true`)

CSP에 키페어가 없더라도 시스템 내부 키 메타데이터를 즉시 정리합니다.

```bash
# 특정 SSH Key 단건 강제 삭제
curl -s -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/sshKey/mig-sshkey-01?force=true' \
  -H 'accept: application/json'

# 또는 네임스페이스 내 SSH Key 일괄 강제 삭제
curl -s -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/sshKey?force=true' \
  -H 'accept: application/json'
```

#### Step 4: vNet 및 Subnet 강제 삭제 (`action=force`)

하위 서브넷을 포함하여 CSP로 강제 삭제(`force=true`)를 전달하고 잔존 레코드를 제거합니다.

```bash
curl -s -u "default:default" -X 'DELETE' \
  'http://localhost:8056/beetle/migration/ns/mig01/resources/vNet/mig-vnet-01?action=force' \
  -H 'accept: application/json'
```

---

## 4. 파라미터 요약 비교표

| 대상 자원          | 정상 삭제 옵션        | Force 삭제 옵션     | Force 삭제 동작 설명                                          |
| :----------------- | :-------------------- | :------------------ | :------------------------------------------------------------ |
| **Infra**          | `?option=terminate`   | **`?option=force`** | CSP VM 상태 무시 및 내부 메타데이터 즉시 제거                 |
| **Security Group** | _(파라미터 없음)_     | **`?force=true`**   | 자원 간 의존성 검사 우회, Spider 실패 시에도 레코드 강제 삭제 |
| **SSH Key**        | _(파라미터 없음)_     | **`?force=true`**   | CSP 키 부재 시에도 내부 레코드 즉시 삭제                      |
| **vNet & Subnet**  | `?action=withsubnets` | **`?action=force`** | CSP 강제 삭제 호출 및 CSP 존재 검증 우회 후 레코드 삭제       |

---

## 5. 자동화 스크립트 예시: force-cleanup.sh

복구 및 일괄 정리를 위해 아래 쉘 스크립트를 활용할 수 있습니다.

```bash
#!/usr/bin/env bash
# force-cleanup.sh: NCP 등 외부 콘솔 삭제로 인한 마이그레이션된 인프라 자원 일괄 강제 정리 스크립트

BASE_URL="http://localhost:8056/beetle"
AUTH="default:default"
NS="mig01"
INFRA_ID="infra101"
SSHKEY_ID="mig-sshkey-01"
VNET_ID="mig-vnet-01"

echo "=========================================================="
echo "⚠️  CM-Beetle Force Cleanup을 시작합니다."
echo "=========================================================="

# 1. Infra 강제 삭제
echo -e "\n[1/4] Infra 강제 삭제 중: ${INFRA_ID}..."
curl -s -u "${AUTH}" -X 'DELETE' \
  "${BASE_URL}/migration/ns/${NS}/infra/${INFRA_ID}?option=force" \
  -H 'accept: application/json'
echo ""

sleep 3

# 2. Security Group 일괄 강제 삭제
echo -e "\n[2/4] Security Group 일괄 강제 삭제 중..."
curl -s -u "${AUTH}" -X 'DELETE' \
  "${BASE_URL}/migration/ns/${NS}/resources/securityGroup?force=true" \
  -H 'accept: application/json'
echo ""

# 3. SSH Key 강제 삭제
echo -e "\n[3/4] SSH Key 강제 삭제 중: ${SSHKEY_ID}..."
curl -s -u "${AUTH}" -X 'DELETE' \
  "${BASE_URL}/migration/ns/${NS}/resources/sshKey/${SSHKEY_ID}?force=true" \
  -H 'accept: application/json'
echo ""

# 4. vNet 및 Subnet 강제 삭제
echo -e "\n[4/4] vNet 및 Subnet 강제 삭제 중: ${VNET_ID}..."
curl -s -u "${AUTH}" -X 'DELETE' \
  "${BASE_URL}/migration/ns/${NS}/resources/vNet/${VNET_ID}?action=force" \
  -H 'accept: application/json'
echo ""

echo -e "\n=========================================================="
echo "✅ Force Cleanup 작업이 완료되었습니다. 최종 상태를 확인합니다."
echo "=========================================================="

echo -e "\n[최종 상태 확인 - Infra 목록]"
curl -s -u "${AUTH}" -X 'GET' "${BASE_URL}/migration/ns/${NS}/infra" -H 'accept: application/json'
echo ""
```
