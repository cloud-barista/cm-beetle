# transx Plan for Strategy and Generalized TransferMode Architecture

> **Status**: Implemented (Finalized Architecture & Specification)  
> **Author**: Cloud-Barista CM-Beetle Development Team  
> **Created**: 2026-11-03  
> **Updated**: 2026-11-04  
> **Target Components**: `transx`, `pkg/api/rest/controller`, `cmd/aimx`, `ui`, `docs`

---

## 1. 개요 및 배경 (Overview & Background)

`transx`는 CM-Beetle 서버 내부 모듈로 동작할 뿐만 아니라, 소스/목표 서버 상의 에이전트(Agent)나 독립 실행형 CLI 환경에서도 범용적으로 활용될 수 있도록 모듈화된 데이터 마이그레이션 패키지입니다. 또한 Filesystem(`rsync`), Object Storage(`S3` Presigned URL), 나아가 향후 RDBMS/NoSQL 등 다양한 스토리지 및 데이터베이스 엔드포인트 간의 전송을 단일 파이프라인(`Pipeline` -> `[]Step` -> `Executor`)으로 추상화하는 것을 목표로 합니다.

본 계획서는 1, 2차 아키텍처 리뷰(`/grill-me`)를 통해 최종 확정된 설계 철학과 구현 명세를 수록한 문서입니다.

### 핵심 목표 요약
1. **단일 `Strategy` 필드: 기본 3종(`auto`, `direct`, `relay`) + 선택적 SSH 전용 2종(`ssh-agent-fwd-push`, `ssh-agent-fwd-pull`)**:
   - **기본 3종**:
     - `"auto"` (기본값): 엔드포인트 조합 및 방화벽 상태에 따라 최적 경로를 자동 선택합니다. `SSH -> SSH`에서는 **1차 `ssh-agent-fwd-push` -> 2차 `ssh-agent-fwd-pull` -> 3차 `relay`** 순서로 자동 탐색 및 폴백합니다.
     - `"direct"`: 중간 스테이징 없이 직접 전송합니다. `SSH -> SSH`에서는 **1차 `ssh-agent-fwd-push` -> 2차 `ssh-agent-fwd-pull`**까지 자동 탐색하며, 양방향 모두 막혀 있으면 `relay`로 가지 않고 즉시 에러를 반환합니다.
     - `"relay"`: `transx` 호스트의 로컬 스테이징(`/tmp/transx-staging/<subdir>`)을 경유하여 2-Step(`pull` -> `push`)으로 전송합니다.
   - **선택적 SSH 전용 2종**:
     - `"ssh-agent-fwd-push"`: `Source -> Destination` 방향 에이전트 포워딩 Push만 시도합니다.
     - `"ssh-agent-fwd-pull"`: `Destination <- Source` 방향 에이전트 포워딩 Pull만 시도합니다.
   - `transx` 오케스트레이터는 양쪽 서버의 SSH Private Key를 메모리에만 로드(`agent.NewKeyring()`)하여 SSH Agent Forwarding으로 인증을 중계하므로, 소스/목표 서버 디스크에 상대방 비밀키 파일을 남기지 않습니다.
2. **일반화된 `TransferMode` 및 미사용 레거시 코드 정리 (`transx/executor-interface.go`)**:
   - 미사용 구버전 심볼(`MethodLocal`, `MethodSSH`, `MethodS3`, `CategoryRsync`, `CategoryObjectStorage`, `IsRsyncMethod`, `IsObjectStorageMethod`, `GetCategory`)을 제거합니다.
   - 내부 스텝 실행 양식(`TransferMode`)을 `transx/executor-interface.go`에 정의하여 모든 실행기(`RsyncExecutor`, `S3Executor`)가 공유하도록 일반화합니다:
     - `"pull"`, `"push"`, `"ssh-agent-fwd-auto"`, `"ssh-agent-fwd-push"`, `"ssh-agent-fwd-pull"`.
3. **사전 연결 탐색(`Pre-Flight Auto-Probing`), 실행 경로 추적(`EffectiveStrategy`), 스테이징 정리 정책**:
   - `transx` <-> 원격 서버 자체의 SSH 접속 실패나 `rsync` 미설치는 치명적 오류로서 즉시 중단하며, 오직 **원격 서버 간(`Primary -> Peer`) SSH 연결/도달 실패**만 `*PreFlightError`로 분류하여 재시도 대기 없이 즉시 다음 방향(`Pull` 또는 `Relay`)으로 전환합니다.
   - `Pipeline.EffectiveStrategy`와 `RsyncExecutor.ResolvedMode`에 실제로 성사된 전략/모드를 기록하고 CM-Beetle API 로그 및 응답(`effectiveStrategy`)에 노출합니다.
   - `Relay` 스테이징 디렉토리는 항상 `/tmp/transx-staging/<subdir>` 하위 폴더로 격리 생성하며, 파이프라인 실행이 **성공했을 때만** 자동 삭제하고 실패 시에는 디버깅을 위해 보존합니다.
   - `local -> local` 전송은 `transx` 라이브러리 차원에서 충실히 지원하며(`auto`, `direct`), 외부 API 호출자의 로컬 파일시스템 접근 차단은 상위 계층인 CM-Beetle API 컨트롤러(`pkg/api/rest/controller/migration-data.go`)에서 수행합니다.
4. **`RsyncExecutor` 전송 복원력 강화 및 `SSHConfig.UnmarshalJSON` 기본값 처리**:
   - 기본값: `Archive: true`, `Compress: true`, `Partial: true`, `MaxRetries: 3`.
   - `SSHConfig.UnmarshalJSON`을 통해 JSON에서 생략된 필드(기본값 `true` 유지)와 명시적으로 해제한 필드(`false` 적용)를 정확히 구분합니다.
5. **Beetle UX Lab (`ui/`) 동적 옵션 제어**:
   - `4. Data Transfer` 화면에서 Source/Target 엔드포인트 종류(`remote-ssh` vs `object-storage`)에 따라 기본 3종 및 선택적 SSH 전용 2종 `Strategy`, `Rsync Options`, 그리고 엔드포인트별 `preCmd`/`postCmd` 입력창을 동적으로 비활성화합니다.

---

## 2. 아키텍처 설계: `Strategy` vs 일반화된 `TransferMode`

### 2.1 2계층 역할 분리 다이어그램

| 계층 | 타입 및 정의 위치 | 정규 상수 값 | 핵심 역할 및 철학 |
| :--- | :--- | :--- | :--- |
| **1. 파이프라인 계층 (User API)** | `Strategy` (`transx/model.go`) | 기본 3종: `"auto"`, `"direct"`, `"relay"`<br>선택적 SSH 전용 2종: `"ssh-agent-fwd-push"`, `"ssh-agent-fwd-pull"` | **"데이터가 어떤 파이프라인 경로로 이동하는가?"**<br>기본값 `"auto"`만으로도 `Push -> Pull -> Relay`를 자동 탐색하며, 필요 시 직접 전송(`direct`), 경유(`relay`), 또는 특정 방향(`ssh-agent-fwd-*`)을 강제 |
| **2. 스텝/실행기 계층 (Internal)** | `TransferMode` (`transx/executor-interface.go`) | `"pull"`, `"push"`, `"ssh-agent-fwd-auto"`, `"ssh-agent-fwd-push"`, `"ssh-agent-fwd-pull"` | **"오케스트레이터(`transx`)가 단일 스텝을 어떤 방향으로 실행하는가?"**<br>`Step`, `RsyncExecutor`, `S3Executor`가 공유하는 일반화된 실행 방향 |

### 2.2 스토리지 조합별 `Strategy` 및 `TransferMode` 검증/매핑 매트릭스

| 출발지 (`Source`) | 목적지 (`Destination`) | 허용되는 `Strategy` | 거부되는 `Strategy` (`Validate` 에러) | 생성되는 `Pipeline` 스텝 및 `Step.Mode` |
| :--- | :--- | :--- | :--- | :--- |
| **SSH** | **SSH** | `auto` (기본)<br>`direct`<br>`relay`<br>`ssh-agent-fwd-push`<br>`ssh-agent-fwd-pull` | 없음 | - `auto`: 1-Step (`Mode: ssh-agent-fwd-auto`) + `Fallback: 2-Step Relay`<br>- `direct`: 1-Step (`Mode: ssh-agent-fwd-auto`, Fallback 없음)<br>- `ssh-agent-fwd-push`: 1-Step (`Mode: ssh-agent-fwd-push`)<br>- `ssh-agent-fwd-pull`: 1-Step (`Mode: ssh-agent-fwd-pull`)<br>- `relay`: 2-Step (`Step 1: pull`, `Step 2: push`) |
| **ObjectStorage (S3)** | **ObjectStorage (S3)** | `auto` (기본: `relay`)<br>`relay` | `direct`<br>`ssh-agent-fwd-push`<br>`ssh-agent-fwd-pull` | - 2-Step (`Step 1: pull` [S3 Download], `Step 2: push` [S3 Upload]) |
| **SSH** | **ObjectStorage (S3)** | `auto` (기본: `relay`)<br>`relay` | `direct`<br>`ssh-agent-fwd-push`<br>`ssh-agent-fwd-pull` | - 2-Step (`Step 1: pull` [Rsync Pull], `Step 2: push` [S3 Upload]) |
| **ObjectStorage (S3)** | **SSH** | `auto` (기본: `relay`)<br>`relay` | `direct`<br>`ssh-agent-fwd-push`<br>`ssh-agent-fwd-pull` | - 2-Step (`Step 1: pull` [S3 Download], `Step 2: push` [Rsync Push]) |
| **Local** | **SSH / S3 / Local** | `auto`, `direct` | `relay`, `ssh-agent-fwd-*` | - 1-Step (`Mode: push`) |
| **SSH / S3** | **Local** | `auto`, `direct` | `relay`, `ssh-agent-fwd-*` | - 1-Step (`Mode: pull`) |

---

## 3. 백엔드(`transx/`, `pkg/`, `cmd/aimx/`) 파일별 상세 구현 명세

### 3.1 `transx/executor-interface.go`

미사용 구버전 심볼(`Method*`, `Category*`, `IsRsyncMethod`, `IsObjectStorageMethod`, `GetCategory`)을 제거하고, `TransferMode` 타입과 5개 정규 상수를 `transx/executor-interface.go`에 정의하여 `Step`, `RsyncExecutor`, `S3Executor`가 공통으로 사용하도록 합니다.

```go
package transx

// Executor defines the interface for transfer operations.
type Executor interface {
	// Execute performs the transfer from source to destination.
	Execute(source, destination DataLocation) error
}

// TransferMode defines the execution direction of a transfer step across Rsync and S3 executors.
type TransferMode string

const (
	// TransferModePull pulls data from a remote endpoint to local filesystem (remote-ssh -> local, or object-storage -> local).
	TransferModePull TransferMode = "pull"

	// TransferModePush pushes data from local filesystem to a destination (local -> remote-ssh, local -> object-storage, or local -> local).
	TransferModePush TransferMode = "push"

	// TransferModeSSHAgentFwdAuto tries ssh-agent-fwd-push first and falls back to ssh-agent-fwd-pull on pre-flight failure.
	TransferModeSSHAgentFwdAuto TransferMode = "ssh-agent-fwd-auto"

	// TransferModeSSHAgentFwdPush connects to source with SSH agent forwarding and pushes directly to destination (src -> dst).
	TransferModeSSHAgentFwdPush TransferMode = "ssh-agent-fwd-push"

	// TransferModeSSHAgentFwdPull connects to destination with SSH agent forwarding and pulls directly from source (dst <- src).
	TransferModeSSHAgentFwdPull TransferMode = "ssh-agent-fwd-pull"
)
```

---

### 3.2 `transx/model.go`

1. **전송 전략 상수(`Strategy*`) 및 스텝 상수 정의**:
   기본 3종(`auto`, `direct`, `relay`)과 선택적 SSH 전용 2종(`ssh-agent-fwd-push`, `ssh-agent-fwd-pull`)으로 정의합니다.
   ```go
   const (
   	// StrategyAuto automatically selects the optimal route (SSH -> SSH: Push -> Pull -> Relay).
   	StrategyAuto = "auto"

   	// StrategyDirect enforces direct transfer without local staging (SSH -> SSH: Push -> Pull).
   	StrategyDirect = "direct"

   	// StrategyRelay enforces a 2-step relay via local staging (/tmp/transx-staging/<subdir>).
   	StrategyRelay = "relay"

   	// StrategySSHAgentFwdPush enforces SSH agent forwarding Push only (Source -> Destination).
   	StrategySSHAgentFwdPush = "ssh-agent-fwd-push"

   	// StrategySSHAgentFwdPull enforces SSH agent forwarding Pull only (Destination <- Source).
   	StrategySSHAgentFwdPull = "ssh-agent-fwd-pull"
   )
   ```

2. **`DataMigrationModel.Strategy` 태그 수정**:
   ```go
   // Strategy determines how the transfer is orchestrated (auto, direct, relay, ssh-agent-fwd-push, ssh-agent-fwd-pull).
   Strategy string `json:"strategy,omitempty" default:"auto" validate:"omitempty,oneof=auto direct relay ssh-agent-fwd-push ssh-agent-fwd-pull"`
   ```

3. **`SSHConfig` 복원력 옵션 필드 및 `UnmarshalJSON` 기본값 처리**:
   ```go
   	// Rsync options
   	Archive    bool `json:"archive,omitempty" default:"true"`
   	Compress   bool `json:"compress,omitempty" default:"true"`
   	Partial    bool `json:"partial,omitempty" default:"true"`
   	MaxRetries int  `json:"maxRetries,omitempty" default:"3"`
   	Delete     bool `json:"delete,omitempty"`
   	Verbose    bool `json:"verbose,omitempty"`
   	DryRun     bool `json:"dryRun,omitempty"`

   	explicitOptions bool
   }

   // UnmarshalJSON pre-populates default SSH and rsync options before decoding JSON.
   func (c *SSHConfig) UnmarshalJSON(data []byte) error {
   	type sshConfigAlias SSHConfig
   	alias := sshConfigAlias{
   		Port:            22,
   		ConnectTimeout:  30,
   		Archive:         true,
   		Compress:        true,
   		Partial:         true,
   		MaxRetries:      3,
   		explicitOptions: true,
   	}
   	if err := json.Unmarshal(data, &alias); err != nil {
   		return err
   	}
   	*c = SSHConfig(alias)
   	return nil
   }
   ```

4. **`Validate(dmm DataMigrationModel)` 및 `validateStrategy(dmm)`**:
   - `local -> local` 전송은 `transx` 라이브러리에서 자체 지원(`auto`, `direct`)하며, 외부 API 호출자의 로컬 파일시스템 접근 차단은 상위 계층(`pkg/api/rest/controller/migration-data.go`)에서 수행합니다.
   - 비-SSH 구간에서 `ssh-agent-fwd-*` 지정, 로컬 엔드포인트가 포함된 구간에서 `relay` 지정, 원격 Object Storage 구간에서 `direct` 지정 등 비호환 조합을 즉시 거부합니다.

---

### 3.3 `transx/executor-s3.go`

`S3Executor`에 `Mode TransferMode`를 부여하여 `TransferModePull`(다운로드)과 `TransferModePush`(업로드)를 일관되게 기록합니다.

```go
// S3Executor implements Executor for S3 object storage transfers using presigned URLs.
type S3Executor struct {
	Provider S3Provider
	Mode     TransferMode
}

// NewS3Executor creates a new S3Executor with the given provider and optional TransferMode.
func NewS3Executor(provider S3Provider, mode ...TransferMode) *S3Executor {
	var m TransferMode
	if len(mode) > 0 {
		m = mode[0]
	}
	return &S3Executor{
		Provider: provider,
		Mode:     m,
	}
}
```

---

### 3.4 `transx/executor-rsync.go`

`RsyncExecutor`에 다음 기능을 구현합니다:
1. `PreFlightError`: 오직 원격 서버 간(`Primary -> Peer`) SSH 도달 실패만 래핑하며, `transx -> Primary` 접속 실패 및 `rsync` 미설치(exit code `10`, `11`)는 일반 치명적 에러로 반환하여 불필요한 폴백을 방지합니다.
2. `ResolvedMode TransferMode`: `ssh-agent-fwd-auto` 실행 시 실제로 성사된 방향(`ssh-agent-fwd-push` 또는 `ssh-agent-fwd-pull`)을 기록합니다.
3. `runWithRetry`: `*PreFlightError` 발생 시에는 재시도 없이 즉시 다음 경로로 전환하고, 사전 점검 통과 후 본 `rsync` 전송 중 발생한 일시적 오류만 최대 `MaxRetries`회 재시도합니다.
4. `executeSSHAgentFwdAuto`, `executeSSHAgentFwdPush`, `executeSSHAgentFwdPull`, `verifyAgentFwdPreFlight`, `buildBaseRsyncFlags`, `buildRemoteRsyncPullCommand` 구현.
5. `determineTransferMode`에서 `local -> local`도 `TransferModePush`(로컬 `rsync` 실행)로 지원합니다.

---

### 3.5 `transx/planner.go` & `transx/transx.go`

1. `Pipeline` 구조체에 `EffectiveStrategy string` 및 `Fallback *Pipeline`을 추가하고, `Step` 구조체에 `Mode TransferMode`를 추가합니다.
2. `Pipeline.Execute()`:
   - `*PreFlightError` 발생 시 `p.Fallback != nil`이면 즉시 2-Step `Relay` 폴백 파이프라인을 실행하고 `p.EffectiveStrategy = StrategyRelay`로 기록합니다.
   - `RsyncExecutor` 실행 성공 시 `ResolvedMode`를 바탕으로 `p.EffectiveStrategy`를 갱신합니다.
   - 모든 스텝이 **성공했을 때만** 이번 파이프라인에서 생성한 특정 하위 스테이징 디렉토리(`/tmp/transx-staging/<subdir>`)를 삭제하고, 실패 시에는 디버깅을 위해 보존합니다.
3. `createLocalStagingLocation(sourcePath)`: `sourcePath`가 `"/"`이거나 비어 있어도 항상 `/tmp/transx-staging/root` 등 하위 폴더를 반환하도록 보장합니다.
4. `MigrateDataWithPipeline(dmm)` 및 `TransferWithPipeline(dmm)`을 제공하여 상위 호출자가 `EffectiveStrategy`를 조회할 수 있도록 합니다.

---

### 3.6 `transx/transx-sec_test.go` 단위 테스트 명세

1. `TestSSHAgentFwdAndStrategyValidation`:
   - 5종 `Strategy`(`auto`, `direct`, `relay`, `ssh-agent-fwd-push`, `ssh-agent-fwd-pull`) 플래닝 검증
   - `local -> local` 지원(`auto`, `direct`) 및 비호환 전략(`relay`, `ssh-agent-fwd-*`) 거부 검증
   - `buildRemoteRsyncCommand` 및 `buildRemoteRsyncPullCommand` 명령어/플래그 생성 검증
2. `TestSSHConfigUnmarshalJSONDefaultsAndOverrides`:
   - JSON에서 생략된 필드 기본값(`Archive: true`, `Compress: true`, `Partial: true`, `MaxRetries: 3`) 유지 및 명시적 `false` 오버라이드 검증
3. `TestPreFlightFallbackAndStagingCleanup`:
   - `*PreFlightError` 발생 시 `Fallback`(`Relay`) 전환 및 `EffectiveStrategy` 기록 검증
   - 성공 시 스테이징 하위 디렉토리 자동 삭제 및 실패 시 보존 검증

---

### 3.7 `pkg/api/rest/controller/migration-data.go` 및 `cmd/aimx/` 업데이트

1. **`pkg/api/rest/controller/migration-data.go`**:
   - Swagger 주석에 5종 `Strategy` 명시
   - `runDataMigration`에서 `transx.MigrateDataWithPipeline(req)`을 호출하여 `effectiveStrategy`를 로그와 응답에 포함
2. **`cmd/aimx/`**:
   - `cmd/aimx/service_migration.go` (`buildTransxMigrationModel`): `Archive: true`, `Compress: true`, `Partial: true`, `MaxRetries: 3` 설정
   - `cmd/aimx/testconf/template-test-config.yaml`, `cmd/aimx/main_test.go`, `cmd/aimx/README.md`: 기본 `transxStrategy`를 `"auto"`로 통일하고 5종 전략 설명 반영

---

## 4. 프론트엔드(`ui/`) 상세 구현 명세

### 4.1 `ui/src/components/center/MigrationExecution.tsx`

`Transfer Strategy` 드롭다운 옵션을 기본 3종(`auto`, `relay`, `direct`)과 선택적 SSH 전용 2종(`ssh-agent-fwd-push`, `ssh-agent-fwd-pull`)으로 업데이트합니다:

```tsx
<option value="auto">Auto (Optimal Path Selection)</option>
<option value="relay">Relay (2-Step Staging via Beetle)</option>
<option value="direct" disabled>Direct (Push -&gt; Pull, Remote SSH Only)</option>
<option value="ssh-agent-fwd-push" disabled>SSH Agent Fwd Push (Src -&gt; Dst, Remote SSH Only)</option>
<option value="ssh-agent-fwd-pull" disabled>SSH Agent Fwd Pull (Dst &lt;- Src, Remote SSH Only)</option>
```

### 4.2 `ui/src/components/data/DataTransferCenter.tsx`

1. **State 및 엔드포인트 플래그**:
   ```tsx
   const [dataStrategy, setDataStrategy] = useState('auto');
   const [rsyncCompress, setRsyncCompress] = useState(true);
   const [rsyncPartial, setRsyncPartial] = useState(true);
   const [rsyncMaxRetries, setRsyncMaxRetries] = useState(3);
   const [rsyncArchive, setRsyncArchive] = useState(true);

   const isBothRemoteSsh = sourceAccessType === 'remote-ssh' && targetAccessType === 'remote-ssh';
   const hasAnyRemoteSsh = sourceAccessType === 'remote-ssh' || targetAccessType === 'remote-ssh';
   const isSourceRemoteSsh = sourceAccessType === 'remote-ssh';
   const isTargetRemoteSsh = targetAccessType === 'remote-ssh';
   ```
2. **엔드포인트 타입 변경 시 자동 보정 (`useEffect`)**:
   - `!isBothRemoteSsh`일 때 `'direct'`, `'ssh-agent-fwd-push'`, `'ssh-agent-fwd-pull'`이 선택되어 있으면 자동으로 `'auto'`로 전환합니다.
   - `!hasAnyRemoteSsh`일 때 `showCommandsSection`을 자동으로 닫습니다.
3. **`Transfer Strategy & Rsync Resilience Options` 섹션**:
   - 기본 3종(`auto`, `direct`, `relay`) 버튼과 선택적 SSH 전용 2종(`ssh-agent-fwd-push`, `ssh-agent-fwd-pull`) 버튼을 구분하여 배치하고, `!isBothRemoteSsh`일 때 SSH 전용 전략들을 비활성화합니다.
   - `Compress (-z)`, `Partial Resume`, `Archive (-a)`, `Max Retries` 컨트롤을 제공하며 `!hasAnyRemoteSsh`일 때 비활성화합니다.
4. **`Source & Destination Commands (preCmd / postCmd)` 개별 제어**:
   - 양쪽 모두 `object-storage`일 때는 섹션 전체를 비활성화합니다.
   - 한쪽만 `object-storage`일 때는 해당 방향의 명령어 입력창(`source.preCmd` 또는 `destination.postCmd`)만 개별 비활성화하고 API 요청 페이로드에서 제외합니다.

---

## 5. 검증 체크리스트 (Verification Checklist)

1. **Go 단위 테스트 및 전체 빌드 검증**:
   ```bash
   go test -v ./transx/... ./cmd/aimx/...
   go build ./...
   ```
2. **문서 및 코드 기호/스타일 검증**:
   - 모든 수정 파일에서 금지된 문장 부호(em-dash, middle-dot), 임의 픽셀 폰트(`text-[10px]` 등), 금지 색상(`cyan-*`, `purple-*`, `sky-*`) 미사용 확인

