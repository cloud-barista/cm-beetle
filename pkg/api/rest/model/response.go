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

// ApiResponse represents a standardized API response structure.
// It uses generics to support different data types for successful responses.
// For successful responses:
//   - Item field contains a single object
//   - Items field contains an array of objects
//
// For error responses:
//   - Error field contains the error message
//
// The Message field is optional and can provide additional context when needed.
type ApiResponse[T any] struct {
	Success bool   `json:"success"`           // Indicates whether the API call was successful
	Item    *T     `json:"item,omitempty"`    // Contains a single object for successful responses
	Items   []T    `json:"items,omitempty"`   // Contains an array of objects for successful list responses
	Message string `json:"message,omitempty"` // Optional message for additional context
	Error   string `json:"error,omitempty"`   // Error message for failed responses
}

// Helper functions to create API responses

// SuccessResponse creates a successful API response for a single object.
// It returns an ApiResponse with success=true and the provided item.
func SuccessResponse[T any](item T) ApiResponse[T] {
	return ApiResponse[T]{
		Success: true,
		Item:    &item,
	}
}

// SuccessResponseWithMessage creates a successful API response with a custom message.
// The message field can provide additional context about the successful operation.
func SuccessResponseWithMessage[T any](item T, message string) ApiResponse[T] {
	return ApiResponse[T]{
		Success: true,
		Item:    &item,
		Message: message,
	}
}

// SuccessListResponse creates a successful API response for a list of objects.
// It returns an ApiResponse with success=true and the provided items array.
func SuccessListResponse[T any](items []T) ApiResponse[T] {
	return ApiResponse[T]{
		Success: true,
		Items:   items,
	}
}

// SuccessListResponseWithMessage creates a successful list API response with a custom message.
// The message field can provide additional context about the successful operation.
func SuccessListResponseWithMessage[T any](items []T, message string) ApiResponse[T] {
	return ApiResponse[T]{
		Success: true,
		Items:   items,
		Message: message,
	}
}

// ErrorResponse creates an error API response with a structured error.
// When an error occurs, T cannot be a specific type, so we use 'any'.
// It returns an ApiResponse with success=false and the error message.
func ErrorResponse(errorMessage, Message string) ApiResponse[any] {
	return ApiResponse[any]{
		Success: false,
		Error:   errorMessage,
		Message: Message,
	}
}

// SimpleErrorResponse creates an error API response from a simple error message.
// This is a convenience function for quick error responses.
// It returns an ApiResponse with success=false and the error message.
func SimpleErrorResponse(errorMessage string) ApiResponse[any] {
	return ApiResponse[any]{
		Success: false,
		Error:   errorMessage,
	}
}
