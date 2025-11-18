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

import "fmt"

// ApiResponse represents a standardized API response structure.
// It uses generics to support different data types for successful responses.
// For successful responses, either Item (single object) or Items (array) will be populated.
// For error responses, the Error field will contain the error details.
type ApiResponse[T any] struct {
	Success bool `json:"success"`         // Indicates whether the API call was successful
	Item    *T   `json:"item,omitempty"`  // Contains a single item for successful responses
	Items   []T  `json:"items,omitempty"` // Contains multiple items for successful list responses

	Error *ApiError `json:"error,omitempty"` // Contains error details for failed responses
}

// ApiError represents an error response structure for API operations.
// It contains a machine-readable error code and a human-readable message.
type ApiError struct {
	Code    string `json:"code"`    // Error code that can be parsed by the client
	Message string `json:"message"` // Message that can be displayed to the user
}

// Error implements the error interface for ApiError.
// It returns a formatted string representation of the error with code and message.
func (e *ApiError) Error() string {
	return fmt.Sprintf("API Error [%s]: %s", e.Code, e.Message)
}

// Helper functions to create API responses

// SuccessResponse creates a successful API response for a single item.
// It returns an ApiResponse with the success flag set to true and the provided item.
func SuccessResponse[T any](item T) ApiResponse[T] {
	return ApiResponse[T]{
		Success: true,
		Item:    &item,
	}
}

// SuccessListResponse creates a successful API response for a list of items.
// It returns an ApiResponse with the success flag set to true and the provided items array.
func SuccessListResponse[T any](items []T) ApiResponse[T] {
	return ApiResponse[T]{
		Success: true,
		Items:   items,
	}
}

// ErrorResponse creates an error API response.
// When an error occurs, T cannot be a specific type, so we use 'any'.
// It returns an ApiResponse with the success flag set to false and the provided error.
func ErrorResponse(err *ApiError) ApiResponse[any] {
	return ApiResponse[any]{
		Success: false,
		Error:   err,
	}
}
