/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

// Response represents a legacy API response structure.
// @Description **(To be deprecated)** This structure is currently in use but will be replaced by the generic `ApiResponse[T]` in the future.
type Response struct {
	Success bool                   `json:"success" example:"true"`
	Text    string                 `json:"text" example:"Any text"`
	Detail  string                 `json:"details,omitempty" example:"Any details"`
	Object  map[string]interface{} `json:"object,omitempty"`
	List    []interface{}          `json:"list,omitempty"`
}

// ApiResponse represents a standardized API response structure.
// It uses generics to support different data types for successful responses.
//
// For successful responses:
//   - Data field contains the result object, array, or paginated result
//
// For error responses:
//   - Error field contains the error message
type ApiResponse[T any] struct {
	Success bool `json:"success" example:"true"` // Indicates whether the API call was successful

	// Code is reserved for internal error/status codes.
	// Uncomment this field when internal status codes are needed.
	// Code int `json:"code,omitempty"`

	Data    *T     `json:"data,omitempty"`                                     // Contains the actual response data (single object, list, or page)
	Message string `json:"message,omitempty" example:"Operation successful"`   // Optional message for additional context
	Error   string `json:"error,omitempty" example:"Error message if failure"` // Error message for failed responses
}

// Page represents a standardized structure for paginated results.
// This struct is intended to be used within the ApiResponse.Data field.
type Page[T any] struct {
	Items      []T   `json:"items"`                   // List of items in the current page
	TotalCount int64 `json:"total_count" example:"1"` // Total number of items across all pages
	Page       int   `json:"page" example:"1"`        // Current page number (1-based index)
	Size       int   `json:"size" example:"10"`       // Number of items per page
	HasNext    bool  `json:"has_next" example:"true"` // Indicates if there is a next page
}

// -------------------------------------------------------------------
// Helper functions to create API responses
// -------------------------------------------------------------------

// SuccessResponse creates a successful API response for a single object or any generic data.
// Usage:
//   - Single Object: SuccessResponse(user) -> Data: User
//   - Raw List:      SuccessResponse(users) -> Data: []User
func SuccessResponse[T any](data T) ApiResponse[T] {
	return ApiResponse[T]{
		Success: true,
		Data:    &data,
	}
}

// SuccessResponseWithMessage creates a successful API response with a custom message.
func SuccessResponseWithMessage[T any](data T, message string) ApiResponse[T] {
	return ApiResponse[T]{
		Success: true,
		Data:    &data,
		Message: message,
	}
}

// SuccessListResponse is a convenience wrapper for list responses.
// It explicitly takes a slice and returns ApiResponse[[]T].
// This ensures that the Data field is always a JSON array.
func SuccessListResponse[T any](items []T) ApiResponse[[]T] {
	return ApiResponse[[]T]{
		Success: true,
		Data:    &items,
	}
}

// SuccessListResponseWithMessage creates a successful list API response with a custom message.
func SuccessListResponseWithMessage[T any](items []T, message string) ApiResponse[[]T] {
	return ApiResponse[[]T]{
		Success: true,
		Data:    &items,
		Message: message,
	}
}

// SuccessPagedResponse creates a successful API response with pagination details.
// It automatically wraps the items and metadata into a Page struct.
func SuccessPagedResponse[T any](items []T, totalCount int64, page, size int) ApiResponse[Page[T]] {
	hasNext := totalCount > int64(page*size)

	pageData := Page[T]{
		Items:      items,
		TotalCount: totalCount,
		Page:       page,
		Size:       size,
		HasNext:    hasNext,
	}

	return ApiResponse[Page[T]]{
		Success: true,
		Data:    &pageData,
	}
}

// ErrorResponse creates an error API response with a structured error.
// T is set to 'any' as there is no data to return.
func ErrorResponse(errorMessage, message string) ApiResponse[any] {
	return ApiResponse[any]{
		Success: false,
		Error:   errorMessage,
		Message: message,
	}
}

// SimpleErrorResponse creates an error API response from a simple error message.
func SimpleErrorResponse(errorMessage string) ApiResponse[any] {
	return ApiResponse[any]{
		Success: false,
		Error:   errorMessage,
	}
}
