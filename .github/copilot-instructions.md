# CM-Beetle Copilot Instructions

## Project Rules

- **Language & Framework:** Go 1.21+, Echo Framework.
- **Documentation Style:** Technical accuracy with a friendly and professional tone.
- **Language Setting:** All responses and explanations must be in **English**.
- **Code Standards:**
  - Echo handlers must follow the standard pattern using `context.Bind` and `context.JSON`.
  - Refer to `pkg/core/common/error.go` for standard error types.

This repository contains the source code for **CM-Beetle** (Computing Infrastructure Migration), a sub-system of the Cloud-Barista platform.

## Project Overview

- **Objectives:** Recommend optimal target computing infrastructure (cloud infrastructure) based on source infrastructure information, and execute migration according to the recommendation.
- **Languages:** Go (1.25.0+).
- **Frameworks:** Echo (Web Framework), Viper (Config), Zerolog (Logging), Swag (API Docs).
- **Dependencies:** `cb-tumblebug`, `cm-model`.
- **Supported Clouds:** AWS, Azure, GCP, NCP, Alibaba Cloud.
- **Architecture:** REST API server (Echo) with a modular architecture separating API handlers from core logic (migration, recommendation). It functions as an orchestration layer, tightly integrated with CB-Tumblebug for managing multi-cloud resources.

## Key Dependencies & Integration Patterns

### Key Reference Files

- **API Specification:** `api/swagger.yaml` (Source of truth for API contracts).
- **Data Models:** `pkg/api/rest/model/` (Request/Response structures).
- **Configuration:** `conf/config.yaml` (Available configuration options).
- **Error Definitions:** `pkg/core/common/error.go` (Standard error types).

### CB-Tumblebug Integration

- **Core Model Import:** `tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"`
- **Resource Interface Import:** `tbresource "github.com/cloud-barista/cb-tumblebug/src/interface/rest/server/resource"`
- **Network Utilities Import:** `"github.com/cloud-barista/cb-tumblebug/src/core/common/netutil"`
- **Client Package:** `pkg/client/tumblebug/` (Contains all Tumblebug API clients)
- **Key Components:**
  - `TumblebugClient`: Main client for interactions (MCI, Namespace, VM, VNet, SecurityGroup, SSH key).

### cm-model Integration

- **Cloud Model Import:** `cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"`
- **On-Premise Model Import:** `onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"`
- **Usage:** Shared data models for cloud and on-premise infrastructure.

## Architecture & Code Organization

### Core Packages Structure

- `cmd/cm-beetle/`: Main application entry point (`main.go`).
- `pkg/api/rest/`: REST API handlers (`controller/`), middlewares (`middlewares/`), and models (`model/`).
- `pkg/client/tumblebug/`: CB-Tumblebug API client implementations.
- `pkg/core/`: Core logic and algorithms.
  - `migration/`: Infrastructure migration logic (CreateVMInfra, etc.).
  - `recommendation/`: Infrastructure recommendation logic (VM Specs, OS Images).
  - `common/`: Shared utilities.
- `pkg/config/`: Configuration management (Viper).
- `pkg/modelconv/`: Model conversion utilities between internal, Tumblebug, and cm-model formats.

### Key Functional Areas

1.  **Infrastructure Migration (`pkg/core/migration/`)**:
    - Migrate on-premise infrastructure to multi-cloud.
    - `CreateVMInfra()`: Create target cloud infrastructure via CB-Tumblebug.
2.  **Infrastructure Recommendation (`pkg/core/recommendation/`)**:
    - Recommend optimal cloud configurations (VM Specs, OS Images, VNet).
    - Uses similarity-based matching algorithms.

### Key Design Patterns

- **Middleware Pattern:**
  - `TumblebugInitChecker`: Ensures Tumblebug is initialized before processing requests.
  - Authentication and CORS handling.
- **Model Conversion Pattern:**
  - Use `modelconv` package for converting between different model formats.
  - Heavy use of CB-Tumblebug and cm-model structures.

## Coding Standards & Conventions

### General

- Write all code comments and documentation in English.
- Use English for struct field comments, function comments, inline comments, and TODO/FIXME notes.
- Ensure consistency and accessibility for international contributors.

### API Response Messages

**User-Centric Message Guidelines:**

- Write messages from the API user's perspective, not the developer's.
- Prioritize **clarity** and **conciseness** while maintaining technical accuracy.
- Remove unnecessary words and technical jargon, but preserve essential information.

**Best Practices:**

- ✅ **Do:** Use short, action-oriented messages
  - `"Provider required"` instead of `"invalid request: 'desiredProvider' is required"`
  - `"Invalid request format"` instead of `"Invalid request body: " + err.Error()`
  - `"Data migrated successfully (2.5s)"` instead of `"Data migration completed successfully (elapsed: 2.5s)"`
- ❌ **Don't:** Include redundant prefixes or verbose explanations
  - Avoid: `"Failed to..."`, `"Error:"`, `"Invalid request:"`
  - Avoid: Technical error details in user-facing messages
  - Avoid: Repetitive parameter names in quotes

**Validation Error Patterns:**

```go
// Required parameter
"Provider required"          // ✅ Clear and concise
"desiredProvider is required"  // ❌ Too technical

// Invalid format
"Invalid request format"     // ✅ Simple
"Invalid request body: ..."  // ❌ Too verbose

// Missing resources
"At least one source server required"  // ✅ Actionable
"Source infrastructure must contain at least one server"  // ❌ Too formal
```

**Success Message Patterns:**

```go
// Operation success
"Recommended 3 object storage(s) for aws ap-northeast-2"  // ✅ Informative
"Successfully generated 3 recommendations..."  // ❌ Redundant "Successfully"

// Completion with timing
"Data migrated successfully (2.5s)"  // ✅ Includes timing
"Data migration completed successfully (elapsed: 2.5s)"  // ❌ Too verbose
```

**When Using Helper Functions:**

- `model.SimpleErrorResponse(msg)`: Keep `msg` concise yet clear enough to understand the issue
- `model.SuccessResponseWithMessage(data, msg)`: Focus on what was achieved, not how
- Always review messages: "Is this clear to someone unfamiliar with the implementation?"
- **Priority order:** Clarity first, then conciseness

### Go

- Detailed Go coding standards are managed in `.github/instructions/go.instructions.md`.

## Common Patterns

### API Handler Pattern

```go
func HandlerName(c echo.Context) error {
    nsId := c.Param("nsId")
    var req RequestModel
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, model.Response{Message: err.Error()})
    }

    // Core logic
    result, err := coreLogic(nsId, req)
    if err != nil {
        log.Error().Err(err).Msg("Operation failed")
        return c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
    }
    return c.JSON(http.StatusOK, result)
}
```

### Client Integration Pattern

```go
// Initialize client
client := tbclient.NewDefaultClient()

// Execute request with proper error handling
result, err := client.SomeOperation(params)
if err != nil {
    log.Error().Err(err).Msg("Tumblebug operation failed")
    return nil, err
}
return result, nil
```

### Model Conversion Pattern

```go
// Convert between different model formats
targetModel := modelconv.ConvertToTargetFormat(sourceModel)

// Use appropriate imports for model types
tbModel := tbmodel.SomeModel{}
cloudModel := cloudmodel.SomeModel{}
```

## Development Workflow

### Working with External Dependencies

#### CB-Tumblebug Integration

- Always use the client packages in `pkg/client/tumblebug/`.
- Handle CB-Tumblebug initialization requirements (e.g., using `TumblebugInitChecker`).
- Respect CB-Tumblebug's REST API patterns and authentication.
- Use proper namespace management.

#### cm-model Integration

- Import shared models from `cm-model` package.
- Ensure model compatibility between versions.
- Use conversion utilities (`pkg/modelconv/`) when transforming data.

### Environment Setup

- **Local Development:** Use `replace` directives in `go.mod` for local dependency testing (e.g., `cb-tumblebug`, `cm-model`).
- **Production:** Ensure `go.mod` specifies production versions.

## Build and Run Instructions

Always use the `Makefile` for build and run tasks.

- **Build:** `make build` (builds `cm-beetle` binary in `cmd/cm-beetle/`).
- **Run:** `make run` (sources `conf/setup.env` and runs the binary).
- **Clean:** `make clean` (removes build artifacts).
- **Lint:** `make lint` (runs `golangci-lint`).
- **Test CLI:** `make test-cli` (builds and runs the test CLI).

## API Documentation

- The API is documented using **Swagger**.
- If you modify API handlers, update the Swagger comments.
- Regenerate documentation with: `make swag`.
- Swagger files are located in `api/`.
- **Example:**
  ```go
  // CreateMCI godoc
  // @ID CreateMCI
  // @Summary Create Multi-Cloud Infrastructure
  // @Description Create a multi-cloud infrastructure
  // @Tags [Migration] Infrastructure
  // @Accept json
  // @Produce json
  // @Param nsId path string true "Namespace ID"
  // @Success 200 {object} ResponseModel
  // @Router /migration/ns/{nsId}/mci [post]
  ```

## Common Tasks

- **Adding a new API endpoint:**

  1. Define the handler in `pkg/api/rest/`.
  2. Add the route in `cmd/cm-beetle/main.go` (or relevant router setup).
  3. Add Swagger comments to the handler.
  4. Run `make swag`.

- **Running locally:**

  1. Ensure `conf/setup.env` and `conf/config.yaml` are configured.
  2. Run `make run`.

- **Testing:**
  - Use `go test ./...` for unit tests.
  - Use `make test-cli` for integration-like testing with the CLI.
