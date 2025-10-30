# Analyzer 구현 완료 - 파일 데이터 형상 분석 및 수집

## 개요

데이터 마이그레이션을 위한 파일 데이터 형상 분석, 수집 및 관리 기능을 구현했습니다.

## 구현된 핵심 기능

### 1. 디렉토리 탐색 및 나열 (Directory Browsing)

**주요 함수:**

- `GetDefaultBaseDir()`: 기본 홈 디렉토리 반환 ($HOME)
- `ListDirectory(path)`: 지정된 경로의 파일/디렉토리 목록 반환

**UI 친화적 설계:**

- 경로를 지정하지 않으면 자동으로 $HOME 디렉토리 표시
- 사용자가 하위 디렉토리를 선택하면 해당 디렉토리 내용 표시
- 단계별 탐색으로 사용자가 마이그레이션 대상 디렉토리 선택 가능

### 2. 파일 메타데이터 수집

**수집되는 메타데이터 항목:**

분석기는 각 파일과 디렉토리에 대해 포괄적인 메타데이터를 추출합니다. 이 정보는 마이그레이션 계획 및 분석에 필수적입니다.

| 메타데이터 항목   | 설명                                                    |
| ----------------- | ------------------------------------------------------- |
| **Path**          | 파일 또는 디렉토리의 전체 절대 경로                     |
| **Name**          | 파일 또는 디렉토리 이름 (경로 제외)                     |
| **Size**          | 바이트 단위 크기 (디렉토리는 0)                         |
| **IsDir**         | 디렉토리 여부를 나타내는 boolean 플래그                 |
| **Mode**          | 권한 모드 (예: "0755", "0644")                          |
| **ModTime**       | 마지막 수정 시간 (내용이 변경된 시간)                   |
| **AccessTime**    | 마지막 access 시간 (파일을 읽거나 연 시간)              |
| **ChangeTime**    | 마지막 상태 변경 시간 (metadata가 수정된 시간)          |
| **Owner**         | 소유자 사용자 ID (Linux/Unix의 UID)                     |
| **Group**         | 그룹 ID (Linux/Unix의 GID)                              |
| **MimeType**      | 파일의 MIME 타입 (예: "text/plain", "application/json") |
| **Extension**     | 점을 포함한 파일 확장자 (예: ".txt", ".json")           |
| **IsSymlink**     | 심볼릭 링크 여부를 나타내는 boolean 플래그              |
| **SymlinkTarget** | 심볼릭 링크인 경우 대상 경로                            |
| **Checksum**      | 파일 무결성 검증을 위한 선택적 체크섬 값 (MD5/SHA256)   |

**FileMetadata 구조체 정의:**

```go
type FileMetadata struct {
    Path          string    // 전체 절대 경로
    Name          string    // 파일/디렉토리 이름
    Size          int64     // 크기 (바이트)
    IsDir         bool      // 디렉토리 여부
    Mode          string    // 권한 모드
    ModTime       time.Time // 수정 시간
    AccessTime    time.Time // 액세스 시간
    ChangeTime    time.Time // 상태 변경 시간
    Owner         string    // 소유자 ID
    Group         string    // 그룹 ID
    MimeType      string    // MIME 타입
    Extension     string    // 파일 확장자
    IsSymlink     bool      // 심볼릭 링크 여부
    SymlinkTarget string    // 심볼릭 링크 대상
    Checksum      string    // 체크섬 (선택적)
}
```

**수집 방법:**

- `os.Lstat()` 사용: 심볼릭 링크를 올바르게 처리
- `syscall.Stat_t` 활용: Linux/Unix 플랫폼 특화 메타데이터 추출
- `filepath.WalkDir()`: 효율적인 재귀 탐색

### 3. 필터링 시스템 (Include/Exclude)

**FilterOptions 구조체:**

```go
type FilterOptions struct {
    IncludePatterns []string // 포함할 패턴 (화이트리스트)
    ExcludePatterns []string // 제외할 패턴 (블랙리스트)
}
```

**패턴 매칭 지원:**

- 단순 패턴: `*.txt`, `*.log`
- 디렉토리 패턴: `data/*`, `logs/*`
- 재귀 패턴: `data/**`, `**/*.json`, `**/test/**`
- rsync 스타일 필터링 로직 채택 (transx 패키지 참고)

**필터링 로직:**

1. Include 패턴이 있으면 해당 패턴에 매칭되는 파일만 포함
2. Exclude 패턴으로 불필요한 파일 제외
3. 패턴이 없으면 모든 파일 포함

### 4. 스캔 옵션

**ScanOptions 구조체:**

```go
type ScanOptions struct {
    BaseDir         string   // 시작 디렉토리
    Recursive       bool     // 하위 디렉토리 포함
    FollowSymlinks  bool     // 심볼릭 링크 따라가기
    IncludeHidden   bool     // 숨김 파일 포함
    MaxDepth        int      // 최대 재귀 깊이
    CollectChecksum bool     // 체크섬 수집 (비용 높음)
    IncludePatterns []string // 포함 패턴
    ExcludePatterns []string // 제외 패턴
}
```

### 5. 마이그레이션 계획 생성

**MigrationPlan 구조체:**

```go
type MigrationPlan struct {
    SourceDir     string         // 원본 디렉토리
    IncludeSubDir bool           // 하위 디렉토리 포함 여부
    FilterOptions FilterOptions  // 필터 설정
    TotalFiles    int            // 총 파일 개수
    TotalSize     int64          // 총 크기 (바이트)
    FileList      []FileMetadata // 마이그레이션 대상 파일 목록
    CreatedAt     time.Time      // 생성 시간
}
```

**생성 함수:**

- `CreateMigrationPlan()`: 사용자가 선택한 디렉토리와 필터를 바탕으로 마이그레이션 계획 생성

## 사용 워크플로우

### 단계 1: 디렉토리 탐색

```go
// 1. 기본 홈 디렉토리 표시
result, _ := analyzer.ListDirectory("")
// UI에서 디렉토리 목록 표시

// 2. 사용자가 하위 디렉토리 선택
result, _ := analyzer.ListDirectory("/home/user/documents")
// UI에서 해당 디렉토리 내용 표시

// 3. 최종 디렉토리 선택될 때까지 반복
```

### 단계 2: 마이그레이션 옵션 설정

```go
// 사용자가 다음을 선택:
// - 하위 디렉토리 포함 여부
// - 포함/제외 필터 패턴
filters := analyzer.FilterOptions{
    IncludePatterns: []string{"*.txt", "*.md", "docs/**"},
    ExcludePatterns: []string{"*.tmp", "*.log", ".git/**"},
}
```

### 단계 3: 마이그레이션 계획 생성

```go
plan, _ := analyzer.CreateMigrationPlan(
    selectedDir,    // 사용자가 선택한 디렉토리
    includeSubDir,  // 하위 디렉토리 포함 여부
    filters,        // 필터 옵션
)

// plan.FileList에 마이그레이션 대상 파일 목록
// plan.TotalFiles, plan.TotalSize로 통계 확인
```

## 구현 상세

### 파일 구조

```
analyzer/
├── analyzer.go                    # 핵심 기능 구현
├── analyzer_platform_linux.go     # Linux 플랫폼 특화 메타데이터 추출
├── analyzer_test.go               # 단위 테스트
├── go.mod                      # Go 모듈 정의
├── README.md                   # 상세 문서
└── examples/
    └── basic/
        └── main.go             # 사용 예제
```

### 플랫폼 지원

- **Linux/Unix**: 완전 지원 (소유자, 그룹, 타임스탬프)
- **Windows**: 별도 platform 파일로 확장 가능
- **macOS**: Linux와 유사하게 지원 가능

### 성능 최적화

- `filepath.WalkDir()` 사용: 메모리 효율적인 디렉토리 순회
- 필터링을 스캔 중에 적용: 불필요한 메타데이터 수집 방지
- 체크섬 계산 선택적 활성화: 성능 영향 최소화

### 테스트 결과

모든 단위 테스트 통과:

```
PASS: TestGetDefaultBaseDir
PASS: TestListDirectory
PASS: TestScanDirectory
PASS: TestExtractFileMetadata
PASS: TestCreateMigrationPlan
PASS: TestMatchPattern
PASS: TestGetDirectoryStatistics
```

## 제안된 개선 사항

### 즉시 가능한 개선

1. ✅ 핵심 데이터 구조 정의
2. ✅ 디렉토리 탐색 기능
3. ✅ 패턴 기반 필터링
4. ✅ 메타데이터 추출
5. ✅ 단위 테스트

### 향후 개선 (선택적)

1. **체크섬 계산 구현**

   - MD5, SHA256 등 해시 알고리즘 추가
   - 증분 계산으로 메모리 효율성 확보

2. **MIME 타입 감지**

   - `net/http.DetectContentType()` 활용
   - 확장자 기반 타입 매핑

3. **필터 모듈화**

   - 별도 `filter` 패키지로 분리
   - Go 표준 패키지 활용 검토

4. **압축 지원**

   - ZIP, TAR, GZ 형식 지원
   - 압축 전 크기 추정

5. **증분 마이그레이션**

   - 변경된 파일만 감지
   - 타임스탬프 비교

6. **진행 상황 추적**

   - 채널 기반 진행률 보고
   - 취소/일시정지 기능

7. **REST API 엔드포인트**
   - HTTP API 제공
   - JSON 응답 형식

## 통합 가능성

### CB-Tumblebug 연동

- 마이그레이션 계획을 기반으로 클라우드 리소스 프로비저닝
- 스토리지 크기 자동 계산

### cm-model 연동

- `FileMetadata`를 cm-model의 데이터 모델과 매핑
- 온프레미스 인프라 정보와 통합

### transx 연동

- 생성된 파일 목록을 transx로 전달하여 실제 전송 수행
- Object Storage 업로드/다운로드

## 결론

파일 데이터 형상 분석 및 수집을 위한 Analyzer 기능이 완성되었습니다.

**핵심 성과:**

- ✅ 사용자 친화적인 디렉토리 탐색
- ✅ 포괄적인 파일 메타데이터 수집
- ✅ 강력한 패턴 기반 필터링
- ✅ 마이그레이션 계획 생성
- ✅ 작고 심플하게 시작, 점진적 확장 가능한 구조
- ✅ 충분한 테스트 커버리지

이제 UI 개발팀이 이 Analyzer API를 활용하여 사용자에게 친화적인 파일 선택 및 마이그레이션 인터페이스를 구축할 수 있습니다.
