# CM-Beetle API 기반 GPU 인프라 추천 및 마이그레이션 가이드

본 문서는 **자체적으로 추출한 소스 GPU 서버(NVIDIA Tesla T4 기준)의 사양 정보**를 바탕으로, **CM-Beetle의 REST API를 호출하여 타겟 클라우드 GPU 인프라를 추천받고 마이그레이션(프로비저닝)하는 전체 절차와 API 명세**를 안내합니다.

---

## 1. 전체 마이그레이션 시나리오 및 단계별 범위

본 가이드가 다루는 전체 워크플로우와 단계별 작업 범위는 다음과 같습니다.

```mermaid
flowchart LR
    subgraph PreStep["[사전 단계] 소스 구축 및 사양 추출"]
        direction TB
        S0["0. 소스 GPU 인프라 구축\n(CB-Tumblebug 활용)\n• GPU: NVIDIA Tesla T4 (16GB)"]
        S1["1. 소스 서버 사양 자체 추출\n(CM-Honeybee 미사용)\n• CPU, RAM, Disk, OS, 방화벽\n• GPU (Tesla T4, 16GB VRAM)"]
        S0 --> S1
    end

    subgraph BeetleAPI["[핵심 연동] CM-Beetle REST API (:8056)"]
        direction TB
        S2["2. 타겟 GPU 인프라 추천 API\nPOST /beetle/recommendation/infra\n• VRAM 하한선(>=16GB) 보장 스펙 추천\n• GPU 전용 OS 이미지(DLAMI 등) 매칭"]
        S3["3. 타겟 GPU 인프라 마이그레이션 API\nPOST /beetle/migration/ns/{nsId}/infra\n• vNet, SG, SSH Key, GPU VM 일괄 생성"]
        S4["4. 타겟 VM 상태 및 SSH 준비 확인 API\nGET .../infra/{infraId}/ssh-ready\n• Public IP, 접속 계정, SSH 가용성 확인"]
        S2 -->|"추천 결과(data[0])\n그대로 전달"| S3 --> S4
    end

    subgraph PostStep["[사후 단계] AI 모델 / SW / 데이터 배포"]
        direction TB
        S5["5. AI 모델 / SW / 데이터 배포\n• 확보된 타겟 VM 접속 정보(IP, SSH Key)로\n  별도 배포 도구/Agent 실행"]
    end

    PreStep -->|"OnpremInfra JSON 구성"| BeetleAPI -->|"타겟 VM 접속 정보"| PostStep
```

|    단계    |       구분        | 주요 작업 내용                                                                                              | 사용 시스템 / API                                                                                 |
| :--------: | :---------------: | :---------------------------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------ |
| **Step 0** |     사전 준비     | 소스 GPU 인프라(`Tesla T4`) 구축                                                                            | `CB-Tumblebug` (`:1323`)                                                                          |
| **Step 1** |     사전 준비     | `CM-Honeybee`를 사용하지 않고 소스 서버에서 직접 하드웨어 및 GPU(`Tesla T4`) 정보를 추출하여 요청 JSON 구성 | 자체 추출 스크립트 / 에이전트                                                                     |
| **Step 2** | **CM-Beetle API** | 소스 사양 기반 타겟 CSP/리전의 최적 GPU VM 스펙 및 GPU OS 이미지 추천                                       | **`POST /beetle/recommendation/infra`**                                                           |
| **Step 3** | **CM-Beetle API** | 추천받은 후보 명세를 기반으로 타겟 클라우드 GPU 인프라 일괄 생성                                            | **`POST /beetle/migration/ns/{nsId}/infra`**                                                      |
| **Step 4** | **CM-Beetle API** | 생성된 타겟 GPU VM의 상세 정보 조회 및 SSH 접속 가능 여부 확인                                              | **`GET /beetle/migration/ns/{nsId}/infra/{infraId}`**<br/>**`GET .../infra/{infraId}/ssh-ready`** |
| **Step 5** |     사후 배포     | 생성된 타겟 GPU VM 위에 별도 방법(Agent 등)으로 AI 구동 환경, 모델, 응용 데이터 배포                        | 별도 배포 도구 / 에이전트                                                                         |
| **Step 6** | **CM-Beetle API** | (테스트 완료 시) 마이그레이션된 타겟 GPU 인프라 및 종속 자원 일괄 삭제                                      | **`DELETE /beetle/migration/ns/{nsId}/infra/{infraId}`**                                          |

---

## 2. CM-Beetle API 호출 공통 규격

1. **기본 엔드포인트 (Base URL)**:
   - `http://<BEETLE_HOST>:8056/beetle`
   - Swagger API 문서 확인: `http://<BEETLE_HOST>:8056/beetle/api`
2. **인증 방식 (Authentication)**:
   - HTTP Basic Authentication 사용 (기본 계정: `default` / `default`, 운영 환경의 경우 설정된 `API_USERNAME` / `API_PASSWORD` 적용)
   - 요청 헤더 예: `Authorization: Basic ZGVmYXVsdDpkZWZhdWx0` (`curl` 사용 시 `-u default:default`)
3. **공통 응답 포맷 (`ApiResponse[T]`)**:
   - 모든 CM-Beetle API는 아래와 같은 표준 JSON 래퍼 구조로 응답을 반환하며, 호출한 API의 반환 타입(`T`)에 따라 **`data`** 필드에는 **단일 객체(`{ ... }`)** 또는 **배열(`[ ... ]`)**이 담깁니다 (단순 삭제/메시지 응답의 경우 `data` 생략).
   ```json
   {
     "success": true,
     "status": 200,
     "message": "성공 또는 상태 요약 메시지",
     "data": { ... } // 단일 객체(예: 마이그레이션 결과 VmInfraInfo) 또는 배열 [ ... ] (예: 추천 후보 목록)
   }
   ```
4. **동기(Sync) 및 비동기(Async) 호출 지원**:
   - **기본 동작(동기)**: 별도 헤더 없이 호출하면 작업이 완료될 때까지 대기한 후 `200 OK` 또는 `201 Created`를 반환합니다. (추천 API는 약 2~5초, 인프라 마이그레이션 API는 VM 생성에 약 2~5분 소요)
   - **비동기 호출(선택)**: HTTP 클라이언트의 타임아웃이 짧은 경우 요청 헤더에 **`Prefer: respond-async`**를 추가하면 즉시 `202 Accepted`와 `reqId`를 반환하며, **`GET /beetle/request/{reqId}`**로 진행 상태(`Handling` $\rightarrow$ `Success` / `Error`)를 폴링할 수 있습니다.

---

## 3. [Step 1] 자체 추출 기반 소스 서버(`Tesla T4`) 모델(`onpremiseInfraModel`) 구성 방법

`CM-Honeybee`를 사용하지 않고 자체적으로 소스 서버 정보를 수집할 때, CM-Beetle 추천 API가 요구하는 **`onpremiseInfraModel` (`OnpremInfra` 구조체)**의 핵심 필드와 리눅스 명령어 기준 추출 방법은 다음과 같습니다.

### 3.1 추천 알고리즘이 참조하는 핵심 필드 및 추출 명령어 매핑

| JSON 경로 (`onpremiseInfraModel` 내부) | 데이터 타입 |  Tesla T4 소스 예시값  | 추천 엔진에서의 역할 및 자체 추출 방법                                                                                                                                                                        |
| :------------------------------------- | :---------: | :--------------------: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `nodes[].hostname`                     |  `string`   |     `"nvidiat4-1"`     | 소스 호스트명 (`hostname` 명령어)                                                                                                                                                                             |
| `nodes[].machineId`                    |  `string`   |     `"nvidiat4-1"`     | **필수**: 소스 서버 고유 식별자. 타겟 노드그룹 이름(`vm-{machineId}`) 및 추적 라벨(`cm-source-machine-ids`) 생성의 기준이 됩니다. (`network.ipv4Networks.defaultGateways[].machineId`와 동일한 값 사용)       |
| `nodes[].cpu.architecture`             |  `string`   |       `"x86_64"`       | **필수**: CPU 아키텍처 (`uname -m` 결과: `x86_64` 또는 `amd64`). 타겟 VM 스펙 및 OS 이미지 아키텍처 필터링에 사용됩니다.                                                                                      |
| `nodes[].cpu.cpus`                     |  `uint32`   |          `1`           | 물리 소켓 수 (일반 VM인 경우 `1`로 설정)                                                                                                                                                                      |
| `nodes[].cpu.cores`                    |  `uint32`   |          `2`           | 물리 코어 수 (`nproc`의 절반 또는 코어 수)                                                                                                                                                                    |
| `nodes[].cpu.threads`                  |  `uint32`   |          `4`           | **필수**: 논리 vCPU 스레드 수 (`nproc` 출력값). CM-Beetle은 `cpus * threads` (예: `1 * 4 = 4 vCPU`)를 기준으로 타겟 vCPU 크기를 매칭합니다.                                                                   |
| `nodes[].memory.totalSize`             |  `uint64`   |          `16`          | **필수**: 총 시스템 메모리 크기 (**단위: GiB, 정수**). `free -g` 실행 시 커널 예약분 때문에 `14` 또는 `15`로 출력될 수 있으므로 표준 클라우드 메모리 크기인 **`16`**으로 정규화하여 입력하는 것을 권장합니다. |
| `nodes[].rootDisk.totalSize`           |  `uint64`   |         `100`          | **필수**: 루트 디스크 크기 (**단위: GB, 정수**, `df -BG /` 기준). 타겟 VM의 `rootDiskSize` 산정에 사용됩니다.                                                                                                 |
| `nodes[].os.id`                        |  `string`   |       `"ubuntu"`       | **필수**: `/etc/os-release`의 `ID` 값 (소문자, 예: `"ubuntu"`). 타겟 OS 이미지 매칭의 핵심 기준입니다.                                                                                                        |
| `nodes[].os.versionId`                 |  `string`   |       `"22.04"`        | **필수**: `/etc/os-release`의 `VERSION_ID` 값 (예: `"22.04"`).                                                                                                                                                |
| `nodes[].os.prettyName`                |  `string`   | `"Ubuntu 22.04.5 LTS"` | `/etc/os-release`의 `PRETTY_NAME` 값.                                                                                                                                                                         |
| **`nodes[].gpuCards[]`**               |   `array`   |  _(아래 3.2절 참조)_   | **필수 (GPU 추천 활성화 키)**: 이 배열에 최소 1개 이상의 GPU 객체가 들어있어야 CM-Beetle이 **GPU 전용 스펙 추천 알고리즘**과 **GPU 전용 OS 이미지(`isGPUImage=true`) 추천 모드**로 동작합니다.                |
| `nodes[].firewallTable[]`              |   `array`   | TCP `22`, 서비스 포트  | 소스 서버에 열려 있는 방화벽 규칙 목록. 여기에 정의한 인바운드 포트들이 타겟 클라우드의 **Security Group(`targetSecurityGroupList`) 방화벽 규칙으로 자동 변환**됩니다.                                        |

### 3.2 `gpuCards` 배열 작성 시 주의사항 (NVIDIA Tesla T4 기준)

소스 서버에서 아래 `nvidia-smi` 명령어로 GPU 정보를 추출할 수 있습니다.

```bash
nvidia-smi --query-gpu=index,uuid,name,driver_version,pci.bus_id,memory.total --format=csv,noheader,nounits
# 출력 예시: 0, GPU-8fae0166-355e-0223-3b69-cd8d9ef7eb92, Tesla T4, 595.91.07, 00000000:00:1E.0, 15360
```

- **`vendor`**: `"NVIDIA"`로 입력합니다.
- **`model`**: `"Tesla T4"` (또는 `"T4"`)로 입력합니다.
- **`memoryTotalGB`**: **`16`**으로 입력합니다.
  - _(참고)_ `nvidia-smi`가 보고하는 `15360 MiB`를 `1024`로 나누면 `15.0 GiB`이지만, 실제 하드웨어 사양은 **`16 GB`**입니다. CM-Beetle 내부에도 `0.5 GB` 허용 오차 보정 로직이 있으나, 요청 JSON 작성 시 표준 VRAM 용량인 **`16`**으로 명시하는 것이 가장 정확합니다.

---

## 4. [Step 2] 타겟 GPU 인프라 추천 API 호출 (`POST /beetle/recommendation/infra`)

### 4.1 API 명세

- **Method & Path**: `POST /beetle/recommendation/infra`
- **Query Parameters**:
  | 파라미터 | 필수 여부 | 기본값 | 설명 |
  | :--- | :---: | :---: | :--- |
  | `desiredCsp` | 조건부 필수 | `aws` | 타겟 클라우드 프로바이더 (`aws`, `azure`, `gcp`, `alibaba`, `ncp`, `nhn` 등). Request Body의 `desiredCspAndRegionPair`에 값을 넣으면 Body 값이 우선 적용됩니다. |
  | `desiredRegion` | 조건부 필수 | `ap-northeast-2` | 타겟 클라우드 리전 (예: `ap-northeast-2`, `koreacentral`, `asia-northeast3` 등). |
  | `limit` | 선택 | `3` | 반환받을 추천 인프라 후보(Candidate) 최대 개수 |
  | `minMatchRate` | 선택 | `90.0` | `highly-matched` 분류 기준 점수 (0~100) |

### 4.2 요청 JSON 예시 (`recommend-request.json`)

아래 JSON은 `1x NVIDIA Tesla T4 (16GB VRAM)`, `4 vCPU`, `16 GiB RAM`, `100 GB Disk`, `Ubuntu 22.04` 소스 서버를 기준으로 작성된 완전한 요청 페이로드입니다.

```json
{
  "desiredCspAndRegionPair": {
    "csp": "aws",
    "region": "ap-northeast-2"
  },
  "onpremiseInfraModel": {
    "network": {
      "ipv4Networks": {
        "defaultGateways": [
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "nvidiat4-1"
          }
        ]
      },
      "ipv6Networks": {}
    },
    "nodes": [
      {
        "hostname": "nvidiat4-1",
        "machineId": "nvidiat4-1",
        "role": "standalone",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 1,
          "cores": 2,
          "threads": 4
        },
        "memory": {
          "type": "DDR4",
          "totalSize": 16,
          "available": 12,
          "used": 4
        },
        "rootDisk": {
          "label": "/",
          "type": "SSD",
          "totalSize": 100,
          "available": 70,
          "used": 30
        },
        "interfaces": [
          {
            "name": "ens5",
            "macAddress": "06:91:bc:3c:2a:01",
            "ipv4CidrBlocks": ["10.0.1.10/24"],
            "mtu": 9001,
            "state": "up"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0/0",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 100,
            "protocol": "dhcp",
            "scope": "global",
            "linkState": "up"
          }
        ],
        "firewallTable": [
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "22",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "3000",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "11434",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "allow"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.5 LTS",
          "version": "22.04.5 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        },
        "gpuCards": [
          {
            "driverIndex": "0",
            "uuid": "GPU-8fae0166-355e-0223-3b69-cd8d9ef7eb92",
            "vendor": "NVIDIA",
            "model": "Tesla T4",
            "type": "GPU",
            "driverVersion": "595.91.07",
            "cudaVersion": "13.2",
            "pciBusId": "00000000:00:1E.0",
            "memoryTotalGB": 16
          }
        ]
      }
    ]
  }
}
```

### 4.3 `curl` 호출 예시

```bash
curl -s -X POST "http://localhost:8056/beetle/recommendation/infra?desiredCsp=aws&desiredRegion=ap-northeast-2&limit=3" \
  -u default:default \
  -H "Content-Type: application/json" \
  -d @recommend-request.json > recommend-response.json
```

### 4.4 응답 JSON 구조 및 핵심 필드 해설 (`recommend-response.json`)

응답의 `data` 필드는 우선순위별로 정렬된 추천 후보 배열(`[]RecommendedInfra`)입니다. 첫 번째 요소인 **`data[0]`**이 최우선 추천 후보(1순위)입니다.

```json
{
  "success": true,
  "status": 200,
  "message": "The result of recommended infrastructure",
  "data": [
    {
      "status": "partially-matched",
      "description": "GPU Candidate #1 | partially-matched | Overall Match Rate: Min=63.0% Max=100.0% Avg=87.7% | VMs: 1 total, 0 matched, 1 acceptable",
      "targetCloud": {
        "csp": "aws",
        "region": "ap-northeast-2"
      },
      "targetInfra": {
        "name": "gpu-infra-01",
        "nodeGroups": [
          {
            "name": "vm-nvidiat4-1",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "nvidiat4-1"
            },
            "description": "Recommended VM 01 for nvidiat4-1",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+g4dn.xlarge",
            "imageId": "ami-00ffffe72a332c865",
            "cspImageName": "ami-00ffffe72a332c865",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": ["sg-01"],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100
          }
        ]
      },
      "targetVNet": {
        "name": "vnet-01",
        "connectionName": "aws-ap-northeast-2",
        "cidrBlock": "10.0.0.0/21",
        "subnetInfoList": [
          {
            "name": "subnet-01",
            "ipv4_CIDR": "10.0.1.0/24",
            "description": "a recommended subnet for GPU migration"
          }
        ]
      },
      "targetSshKey": {
        "name": "sshkey-01",
        "connectionName": "aws-ap-northeast-2"
      },
      "targetSpecList": [
        {
          "id": "aws+ap-northeast-2+g4dn.xlarge",
          "cspSpecName": "g4dn.xlarge",
          "vCPU": 4,
          "memoryGiB": 16,
          "acceleratorModel": "NVIDIA T4",
          "acceleratorCount": 1,
          "acceleratorMemoryGB": 17,
          "acceleratorType": "gpu",
          "costPerHour": 0.647
        }
      ],
      "targetOsImageList": [
        {
          "id": "ami-00ffffe72a332c865",
          "cspImageName": "ami-00ffffe72a332c865",
          "isGPUImage": true,
          "osType": "Ubuntu 22.04",
          "osDistribution": "Deep Learning Base OSS Nvidia Driver GPU AMI (Ubuntu 22.04) 20260818"
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "sg-01",
          "connectionName": "aws-ap-northeast-2",
          "vNetId": "vnet-01",
          "firewallRules": [
            {
              "Ports": "22",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "3000",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "11434",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            }
          ]
        }
      ]
    }
  ]
}
```

> [!IMPORTANT]
> **추천 응답(`data[0]`)을 다음 단계인 마이그레이션 API에 전달할 때의 핵심 원칙**:
>
> - 추천 결과인 `data[0]` 객체 안에는 타겟 VM 생성(`targetInfra`)뿐만 아니라 CM-Beetle이 함께 생성해야 할 네트워크(`targetVNet`), 보안그룹(`targetSecurityGroupList`), SSH 키(`targetSshKey`), 스펙/이미지 검증 메타데이터(`cspImageName`, `targetSpecList`, `targetOsImageList`)가 모두 상호 참조(ID 연결) 형태로 묶여 있습니다.
> - 따라서 **추천 응답에서 `data[0]` 객체 전체를 수정 없이 그대로 추출하여 Step 3 마이그레이션 API의 Request Body로 전달**하는 것이 가장 안전하고 표준적인 연동 방법입니다.

---

## 5. [Step 3] 타겟 GPU 인프라 마이그레이션(생성) API 호출 (`POST /beetle/migration/ns/{nsId}/infra`)

### 5.1 API 명세

- **Method & Path**: `POST /beetle/migration/ns/{nsId}/infra`
- **Path Parameters**:
  | 파라미터 | 필수 여부 | 기본값 / 예시값 | 설명 |
  | :--- | :---: | :---: | :--- |
  | `nsId` | **필수** | `mig01` | 타겟 인프라를 생성할 네임스페이스 ID.<br/>• **`mig01`**은 CM-Beetle 서버 기동 시 CB-Tumblebug에 **자동 생성되는 기본 네임스페이스**이므로 별도 사전 작업 없이 바로 사용할 수 있습니다 (`default` 네임스페이스도 동일하게 사용 가능).<br/>• 만약 `mig01`, `default` 이외의 **별도 네임스페이스(예: `custom-ns`)**를 지정하려는 경우에도 정상 동작하며, 단 호출 전에 CB-Tumblebug API(`POST /tumblebug/ns`, `{"name": "custom-ns", "description": "..."}`)로 해당 네임스페이스를 먼저 생성해 두어야 합니다. |
- **Query Parameters**:
  | 파라미터 | 필수 여부 | 기본값 | 설명 |
  | :--- | :---: | :---: | :--- |
  | `nameSeed` | 선택 (권장) | `""` | 생성되는 모든 클라우드 자원 이름 앞에 붙일 접두사(Prefix). 예를 들어 `nameSeed=aimx`로 지정하면 `gpu-infra-01` $\rightarrow$ `aimx-gpu-infra-01`, `vnet-01` $\rightarrow$ `aimx-vnet-01`로 자동 변환되어 **동일 네임스페이스(`mig01`) 내 자원 이름 충돌을 방지**합니다. |
  | `useExisting` | 선택 | `true` | 동일한 이름의 하위 자원(`vNet`, `sshKey`, `securityGroup`)이 이미 존재할 경우 새로 만들지 않고 재사용할지 여부 |

### 5.2 요청 바디 준비 및 `curl` 호출 예시

Step 2의 추천 응답 파일(`recommend-response.json`)에서 1순위 후보(`data[0]`)만 추출하여 바로 마이그레이션 API로 전달합니다.

```bash
# 1. 추천 응답에서 1순위 후보(data[0])만 추출하여 migrate-request.json으로 저장
jq '.data[0]' recommend-response.json > migrate-request.json

# 2. CM-Beetle 인프라 마이그레이션 API 호출 (기본 네임스페이스: nsId=mig01, nameSeed=aimx)
curl -s -X POST "http://localhost:8056/beetle/migration/ns/mig01/infra?nameSeed=aimx" \
  -u default:default \
  -H "Content-Type: application/json" \
  -d @migrate-request.json > migrate-response.json
```

> 💡 **팁 (비동기 호출이 필요한 경우)**:
> GPU VM 프로비저닝은 클라우드 벤더에 따라 약 2~5분이 소요됩니다. 만약 호출 측 HTTP 클라이언트의 타임아웃이 짧다면 아래와 같이 `Prefer: respond-async` 헤더를 붙여 비동기로 실행하고 상태를 폴링하십시오.
>
> ```bash
> # 비동기 실행 요청 (즉시 202 Accepted 및 reqId 반환)
> curl -s -X POST "http://localhost:8056/beetle/migration/ns/mig01/infra?nameSeed=aimx" \
>   -u default:default \
>   -H "Content-Type: application/json" \
>   -H "Prefer: respond-async" \
>   -d @migrate-request.json
>
> # 반환된 reqId로 진행 상태 폴링 (status가 "Success"가 될 때까지 조회)
> curl -s -X GET "http://localhost:8056/beetle/request/{reqId}" -u default:default
> ```

### 5.3 마이그레이션 완료 응답 확인 (`migrate-response.json`)

생성이 완료되면 `201 Created`와 함께 프로비저닝된 타겟 GPU 인프라 정보(`VmInfraInfo`)가 `data` 필드에 반환됩니다.

```json
{
  "success": true,
  "status": 201,
  "message": "The result of migrating the multi-cloud infrastructure",
  "data": {
    "id": "aimx-gpu-infra-01",
    "name": "aimx-gpu-infra-01",
    "status": "Running (1/1)",
    "node": [
      {
        "id": "aimx-vm-nvidiat4-1-1",
        "name": "aimx-vm-nvidiat4-1-1",
        "nodeGroupId": "aimx-vm-nvidiat4-1",
        "connectionName": "aws-ap-northeast-2",
        "specId": "aws+ap-northeast-2+g4dn.xlarge",
        "imageId": "ami-00ffffe72a332c865",
        "vNetId": "aimx-vnet-01",
        "subnetId": "aimx-subnet-01",
        "securityGroupIds": ["aimx-sg-01"],
        "sshKeyId": "aimx-sshkey-01",
        "nodeUserName": "cb-user",
        "publicIP": "43.200.254.107",
        "privateIP": "10.0.1.60",
        "sshPort": "22",
        "status": "Running"
      }
    ]
  }
}
```

---

## 6. [Step 4 & Step 5] 타겟 VM SSH 준비 확인 및 SW/데이터 배포 연계

타겟 GPU VM의 클라우드 상태가 `Running`으로 바뀌더라도, OS 내부의 `cloud-init`(SSH 계정 생성)과 NVIDIA 드라이버 초기화가 끝나기까지 약 30초~1분가량 추가 시간이 걸릴 수 있습니다.

### 6.1 타겟 인프라 SSH 접속 준비 상태 확인 API (`GET .../ssh-ready`)

배포 도구(Agent, SCP, Ansible 등)를 실행하기 전에 CM-Beetle의 **`ssh-ready`** API를 호출하면 모든 노드가 실제로 SSH 명령을 받을 수 있는 상태인지 안전하게 확인할 수 있습니다.

```bash
# 주의: 본 API는 내부적으로 최대 30초 동안(10초 간격) SSH 응답을 확인한 후 결과를 반환하며,
# 동일 인프라(nsId:infraId)에 대해 30초당 1회 호출 제한(Rate Limit)이 적용됩니다.
curl -s -X GET "http://localhost:8056/beetle/migration/ns/mig01/infra/aimx-gpu-infra-01/ssh-ready?option=detail" \
  -u default:default | jq .
```

**응답 예시 (`ready: true` 확인 후 SW 배포 시작)**:

```json
{
  "success": true,
  "status": 200,
  "message": "SSH ready - All nodes in the infrastructure are SSH-accessible.",
  "data": {
    "ready": true,
    "totalNodes": 1,
    "readyNodes": 1,
    "nodeStatus": [
      {
        "id": "aimx-vm-nvidiat4-1-1",
        "name": "aimx-vm-nvidiat4-1-1",
        "publicIP": "43.200.254.107",
        "privateIP": "10.0.1.60",
        "username": "cb-user",
        "status": "Running",
        "sshReady": true,
        "sshPort": 22
      }
    ],
    "message": "All nodes are SSH-accessible",
    "checkedAt": "2026-10-16T18:00:00+09:00"
  }
}
```

### 6.2 SW/Data 배포를 위한 타겟 VM 접속 정보 획득 방법

타겟 GPU VM(`43.200.254.107`)에 접속하여 별도 방법으로 SW나 데이터를 배포하려면 다음 3가지 정보를 사용하면 됩니다.

1. **타겟 VM 공인 IP (`publicIP`) 및 SSH 포트 (`sshPort`)**:
   - Step 3 마이그레이션 응답 또는 `GET /beetle/migration/ns/{nsId}/infra/{infraId}` 응답의 `data.node[0].publicIP` (`43.200.254.107`), `data.node[0].sshPort` (`22`)
2. **SSH 접속 계정명 (`nodeUserName`)**:
   - `data.node[0].nodeUserName` (기본값: `"cb-user"` 또는 `"ubuntu"`)
3. **SSH 개인키 (`privateKey`) 조회**:
   - `data.node[0].sshKeyId` (예: `"aimx-sshkey-01"`)를 확인한 후, CB-Tumblebug의 SSH Key 조회 API를 호출하면 PEM 형식의 개인키(`privateKey`) 원문을 가져올 수 있습니다.

   ```bash
   # CB-Tumblebug에서 생성된 타겟 VM의 SSH 개인키(PEM) 추출 (nsId=mig01 기준)
   curl -s -X GET "http://localhost:1323/tumblebug/ns/mig01/resources/sshKey/aimx-sshkey-01" \
     -u default:default | jq -r '.privateKey' > target-vm.pem
   chmod 600 target-vm.pem

   # 추출한 키로 타겟 GPU VM에 직접 SSH 접속 또는 배포 스크립트 실행
   ssh -i target-vm.pem cb-user@43.200.254.107 "nvidia-smi"
   ```

   _(또는 SSH 직접 접속 대신 CB-Tumblebug의 원격 명령 실행 API인 `POST /tumblebug/ns/{nsId}/cmd/infra/{infraId}`를 활용하여 배포 스크립트를 전달할 수도 있습니다.)_

---

## 7. [Step 6] 마이그레이션 인프라 조회 및 정리(삭제) API

테스트 완료 후 생성된 타겟 GPU VM과 하위 네트워크/보안 자원을 정리할 때 사용하는 API입니다.

| 기능                                   | HTTP Method & Path                                                    | 설명 및 예시                                                                                                  |
| :------------------------------------- | :-------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------ |
| **마이그레이션된 인프라 목록 조회**    | `GET /beetle/migration/ns/{nsId}/infra`                               | 네임스페이스 내 전체 인프라 상세 목록 조회 (`?option=id` 추가 시 ID 목록만 조회)                              |
| **단일 인프라 상세 조회**              | `GET /beetle/migration/ns/{nsId}/infra/{infraId}`                     | 특정 타겟 인프라(`aimx-gpu-infra-01`)의 최신 노드 상태 및 IP 정보 조회                                        |
| **타겟 인프라 및 종속 자원 일괄 삭제** | `DELETE /beetle/migration/ns/{nsId}/infra/{infraId}?option=terminate` | 타겟 GPU VM 종료 후 종속된 `SecurityGroup` $\rightarrow$ `SSHKey` $\rightarrow$ `vNet`까지 순서대로 일괄 삭제 |

```bash
# 타겟 GPU 인프라 및 하위 자원(SG, SSHKey, vNet) 일괄 삭제 (nsId=mig01 기준)
curl -s -X DELETE "http://localhost:8056/beetle/migration/ns/mig01/infra/aimx-gpu-infra-01?option=terminate" \
  -u default:default | jq .
```
