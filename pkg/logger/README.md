# pkg/logger

Structured logging with real-time sensitive data masking and OpenTelemetry tracing support, built on [`github.com/rs/zerolog`](https://github.com/rs/zerolog) and [`gopkg.in/natefinch/lumberjack.v2`](https://gopkg.in/natefinch/lumberjack.v2).

This package provides logging primitives for CM-Beetle, featuring automatic redacting of cloud credentials, private keys, database passwords, and Bearer tokens before log events reach terminal or file outputs.

## Architecture

`MaskingWriter` sits directly between zerolog and destination writers (standard output, log files, or both), acting as an intercepting stream filter:

```text
[zerolog Event]
       |
       v
[MaskingWriter]
       |
       +---> Fast-Path Check (sensitiveKeywords: bytes.Contains)
       |        |
       |        +-- [No Match] ----> Skip regex, forward original bytes
       |        |
       |        +-- [Match Found] -> Apply Regex Masking Rules
       |                                |
       |                                +-- PEM blocks (certificates/keys)
       |                                +-- Bearer tokens
       |                                +-- JSON & Escaped JSON credentials
       |                                +-- Config tags (<TB_POSTGRES_PASSWORD>)
       |                                +-- Key-Value pairs (password=..., auth: ...)
       v
[Destination Writer]
       |
       +---> ConsoleWriter (os.Stdout)
       +---> Lumberjack Logger (sharedLogFile with rotation)
       +---> MultiLevelWriter (both)
```

## Masking Pipeline

Log masking operates in two distinct stages to maintain high logging throughput:

1. **Stage 1: Fast-Path Filter (Zero Allocations)**
   Before running regular expression replacements, the log buffer is converted to lowercase and scanned against `sensitiveKeywords` using `bytes.Contains`. If none of the 13 sensitive keywords (`password`, `secret`, `token`, `privatekey`, `auth`, `-----begin`, `pem`, `bearer `, etc.) are present, the original byte slice is returned immediately without regular expression evaluation.

2. **Stage 2: Pattern Redaction**
   When a keyword is detected, regular expressions scan and replace sensitive patterns with fixed-length asterisks (`********`) or structured tags (`[MASKED ...]`).

### Supported Masking Patterns

| Pattern Category | Target Rule             | Sample Raw Log                       | Masked Output                            |
| :--------------- | :---------------------- | :----------------------------------- | :--------------------------------------- |
| **PEM Block**    | `-----BEGIN ...-----`   | `-----BEGIN CERTIFICATE----- ...`    | `[MASKED CERTIFICATE]`                   |
| **Bearer Token** | `Bearer <token>`        | `Authorization: Bearer eyJhbGci...`  | `Authorization: Bearer ********`         |
| **JSON Fields**  | `"password": "val"`     | `{"user":"admin","password":"p123"}` | `{"user":"admin","password":"********"}` |
| **Escaped JSON** | `\"password\": \"val\"` | `{\"api_token\":\"secret123\"}`      | `{\"api_token\":\"********\"}`           |
| **Config Tags**  | `<CONFIG_TAG> val`      | `<TB_POSTGRES_PASSWORD> p123`        | `<TB_POSTGRES_PASSWORD> ********`        |
| **Key-Value**    | `password=val`          | `pass=secret123&user=admin`          | `pass=********&user=admin`               |

## Components

### `MaskingWriter`

`MaskingWriter` wraps any `io.Writer` and satisfies both `io.Writer` and `zerolog.LevelWriter`:

```go
type MaskingWriter struct {
    w io.Writer
}

func NewMaskingWriter(w io.Writer) *MaskingWriter
func (mw *MaskingWriter) Write(p []byte) (n int, err error)
func (mw *MaskingWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error)
```

- Satisfies `io.Writer`: Always returns `len(p)` on success so callers never encounter short write errors when masking changes byte length.
- Satisfies `zerolog.LevelWriter`: Preserves zerolog log-level metadata if the underlying writer implements `zerolog.LevelWriter`.

### Configuration

`Config` controls the logger and writer initialization:

| Field               | Type     | Default                     | Description                                                                        |
| ------------------- | -------- | --------------------------- | ---------------------------------------------------------------------------------- |
| `LogLevel`          | `string` | `"debug"`                   | Minimum log level (`trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic`).   |
| `LogWriter`         | `string` | `"console"`                 | Output target (`"stdout"`, `"file"`, `"both"`).                                    |
| `LogFilePath`       | `string` | `"./log/app.log"`           | Destination path for file logging.                                                 |
| `MaxSize`           | `int`    | `1000`                      | Maximum log file size in megabytes before rotation.                                |
| `MaxBackups`        | `int`    | `3`                         | Maximum number of retained old log files.                                          |
| `MaxAge`            | `int`    | `30`                        | Maximum number of days to retain old log files.                                    |
| `Compress`          | `bool`   | `false`                     | Whether to compress rotated log files with gzip.                                   |
| `LogMaskingEnabled` | `*bool`  | `nil` (evaluates to `true`) | Toggles real-time masking on destination writers. Set `BoolPtr(false)` to disable. |

## Usage Examples

### 1. Initializing CM-Beetle Application Logger

```go
package main

import (
    "github.com/cloud-barista/cm-beetle/pkg/logger"
)

func main() {
    log := logger.NewLogger(logger.Config{
        LogLevel:          "info",
        LogWriter:         "both",
        LogFilePath:       "./log/beetle.log",
        LogMaskingEnabled: logger.BoolPtr(true),
    })

    // MaskingWriter automatically intercepts and masks sensitive data
    log.Info().
        Str("user", "admin").
        Str("MasterUserPassword", "PlainPassword123!").
        Msg("Connecting to database")
    // Output: {"level":"info","user":"admin","MasterUserPassword":"********","message":"Connecting to database"}
}
```

### 2. Standalone MaskingWriter Wrapping Custom Writer

```go
package main

import (
    "os"

    "github.com/cloud-barista/cm-beetle/pkg/logger"
    "github.com/rs/zerolog"
)

func main() {
    // Wrap standard output with sensitive data masking
    maskingWriter := logger.NewMaskingWriter(os.Stdout)
    log := zerolog.New(maskingWriter).With().Timestamp().Logger()

    log.Info().Msg("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.token123")
    // Output: {"level":"info","message":"Bearer ********"}
}
```

### 3. TracingHook Integration

Loggers initialized via `NewLogger` include `TracingHook` automatically. When request contexts hold `trace_id` and `span_id`, they are embedded into structured logs:

```go
ctx := context.WithValue(context.Background(), logger.TraceIdKey, "trace-abc-123")
log.WithContext(ctx).Info().Msg("Processing migration step")
// Output includes: "trace_id":"trace-abc-123"
```

## Testing & Benchmarks

Run the unit tests to verify masking rules across supported formats:

```bash
go test -v ./pkg/logger -run "TestMaskSensitiveData|TestMaskingWriter"
```

Run benchmarks to evaluate fast-path versus regex masking overhead:

```bash
go test -bench=. ./pkg/logger -benchmem
```
