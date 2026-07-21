# API Response Policy

This document outlines the policies for API response codes, success messages, and error handling in CM-Beetle.

## HTTP Status Codes

| Status Code | Description | Usage |
| :--- | :--- | :--- |
| **200 OK** | Success | Standard response for successful GET, PUT, DELETE operations. |
| **201 Created** | Created | Standard response for successful POST operations resulting in resource creation. |
| **202 Accepted** | Async job accepted | Reserved repo-wide for async responses (see [Async Responses](#async-responses)). Never return 202 for any other reason. |
| **400 Bad Request** | Bad Request | Invalid input parameters or request body. |
| **404 Not Found** | Not Found | Resource does not exist. Used for GET, DELETE, etc., when the target ID is invalid. |
| **500 Internal Server Error** | Server Error | Unexpected server-side errors. |
| **503 Service Unavailable** | At capacity | Returned by async-capable endpoints when too many async jobs are already running (see [Async Responses](#async-responses)). |

## Response Body

### Success
- **Resource Return**: Return the requested resource object directly (e.g., `VNetInfo`, `MCI`).
- **Message Return**: For operations without a specific resource return (e.g., Delete), return a standard success message (e.g., `SimpleMsg`).

### Error
- **Structure**: Return a standard error message structure (e.g., `SimpleMsg` with a `message` field).
- **Handling Upstream Errors**: 
  - If an upstream service (e.g., Tumblebug) returns an error indicating a resource is missing (e.g., "not found"), the controller must catch this and return `404 Not Found` instead of `500 Internal Server Error`.
  - This may require parsing error strings if typed errors are not available from the client library.

## Proxy Behavior
- **Status Code Preservation**: When proxying requests to Tumblebug or other services, the proxy **must** preserve the upstream HTTP status code.
- **Transparency**: Do not override specific upstream codes (like `404`) with generic codes (like `500`) unless it is a true internal proxy error.

## Async Responses

Some long-running APIs (data/infra migration, recommendation) support an optional asynchronous mode.

- **Opt-in**: send header `Prefer: respond-async`. Without it, the API responds synchronously as usual.
  Only the `respond-async` token is recognized; other `Prefer` tokens (e.g. `wait=N`) are ignored.
- **Accepted**: `202` with `{reqId, status: "Handling", statusUrl}` and header `Preference-Applied: respond-async`.
- **Poll**: `GET /request/{reqId}` until `status` is `Success` or `Error`.
- **At capacity**: `503` if too many async jobs are already running (process-wide limit, currently 20),
  with a `Retry-After` header. The sync path is never affected by this limit.
- **Implementation**: handlers use `common.RunAsync` (`pkg/core/common/request-manager.go`), which finalizes
  the request record itself. Middleware skips status finalization for any `202` response — this is why
  `202` must never be returned for a non-async reason.
