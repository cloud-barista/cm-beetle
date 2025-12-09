# API Development Guide

This guide provides instructions and best practices for developing REST APIs in **CM-Beetle**.
It covers the project structure, standard response formats, and step-by-step instructions for adding new APIs.

## 1. Project Structure

CM-Beetle follows a standard Go project layout with a focus on modularity.
The key directories for API development are:

```
.
├── cmd/
│   └── cm-beetle/          # Main application entry point
├── pkg/
│   └── api/
│       └── rest/
│           ├── controller/ # API Handlers (Business Logic Interface)
│           ├── model/      # Request/Response Models
│           ├── route/      # (Optional) Route definitions
│           └── server.go   # Server setup and Route registration
└── docs/                   # Documentation
```

## 2. Standard API Response Structure

To ensure consistency across all APIs, CM-Beetle uses a standardized generic response structure defined in `pkg/api/rest/model/response.go`.

### 2.1. ApiResponse[T]

The `ApiResponse[T]` struct is designed to handle both success and error responses uniformly.

```go
type ApiResponse[T any] struct {
    Success bool   `json:"success"`           // Indicates success/failure

    // Code is reserved for internal error/status codes.
    // Uncomment this field when internal status codes are needed.
    // Code int `json:"code,omitempty"`

    Data    T      `json:"data,omitempty"`    // Payload (Object, List, or Page)
    Message string `json:"message,omitempty"` // Contextual message
    Error   string `json:"error,omitempty"`   // Error details (if failed)
}
```

### 2.2. Helper Functions

Use the provided helper functions in your controllers to generate responses:

| Function                                      | Usage                      | Example                                                                          |
| :-------------------------------------------- | :------------------------- | :------------------------------------------------------------------------------- |
| `model.SuccessResponse(data)`                 | Return a single object     | `return c.JSON(http.StatusOK, model.SuccessResponse(user))`                      |
| `model.SuccessListResponse(items)`            | Return a list of items     | `return c.JSON(http.StatusOK, model.SuccessListResponse(users))`                 |
| `model.SuccessResponseWithMessage(data, msg)` | Return data with a message | `return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(nil, "Deleted"))` |
| `model.SimpleErrorResponse(err)`              | Return a simple error      | `return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))`   |

## 3. How to Add a New API

Follow these steps to add a new API endpoint (e.g., `GET /beetle/sample/users`).

### Step 1: Define the Model

Create or update a file in `pkg/api/rest/model/` (e.g., `user.go`) to define your data structures.

```go
package model

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

### Step 2: Implement the Controller

Create a new controller file in `pkg/api/rest/controller/` (e.g., `user.go`).
Implement the handler function using the **Echo** context and **ApiResponse** helpers.
Don't forget to add **Swagger annotations**.

```go
package controller

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
)

// GetUsers godoc
// @Summary Get a list of users
// @Description Retrieve all registered users.
// @Tags [Sample] User Management
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ApiResponse[[]model.User] "List of users"
// @Failure 500 {object} model.ApiResponse[any] "Internal Server Error"
// @Router /sample/users [get]
func GetUsers(c echo.Context) error {
    // Mock data (Replace with actual logic)
    users := []model.User{
        {ID: "u1", Name: "Alice", Email: "alice@example.com"},
        {ID: "u2", Name: "Bob", Email: "bob@example.com"},
    }

    // Return standardized list response
    return c.JSON(http.StatusOK, model.SuccessListResponse(users))
}
```

### Step 3: Register the Route

Open `pkg/api/rest/server.go` and register your new handler in the `RunServer` function.

```go
// ... inside RunServer function ...

// Create a new group if needed
gSample := gBeetle.Group("/sample")

// Register the endpoint
gSample.GET("/users", controller.GetUsers)
```

### Step 4: Generate API Documentation

Run the following command to regenerate Swagger documentation:

```bash
make swag
```

This will update `pkg/api/rest/docs/docs.go`, `swagger.json`, and `swagger.yaml`.

## 4. Testing

### 4.1. Run the Server

```bash
make run
```

### 4.2. Verify via Swagger UI

Open your browser and navigate to:
`http://localhost:8056/beetle/swagger/index.html`

You should see your new API listed under the **[Sample] User Management** tag.

### 4.3. Verify via cURL

```bash
curl -X GET http://localhost:8056/beetle/sample/users
```

**Expected Output:**

```json
{
  "success": true,
  "data": [
    { "id": "u1", "name": "Alice", "email": "alice@example.com" },
    { "id": "u2", "name": "Bob", "email": "bob@example.com" }
  ]
}
```
