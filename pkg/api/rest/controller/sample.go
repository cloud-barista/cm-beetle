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

package controller

import (
	"net/http"
	"strconv"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/labstack/echo/v4"
)

// [Note]
// No RequestBody required for "GET /users"

type GetUsersResponse struct {
	Users []model.MyUser `json:"users"`
}

// GetUsers godoc
// @ID GetUsers
// @Summary Get a list of users
// @Description Get information of all users.
// @Tags [Sample API] Users
// @Accept  json
// @Produce  json
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} GetUsersResponse "(sample) This is a sample description for success response in Swagger UI"
// @Failure 404 {object} object "User Not Found"
// @Router /sample/users [get]
func GetUsers(c echo.Context) error {

	// In this example, hardcoded data is returned
	users := []model.MyUser{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Anne Jacqueline Hathaway", Email: "Anne@example.com"},
		{ID: 3, Name: "Robert John Downey Jr.", Email: "Robert@example.com"},
	}
	return c.JSON(http.StatusOK, users)
}

// [Note]
// No RequestBody required for "GET /users"

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type GetUserResponse struct {
	model.MyUser
}

// GetUser godoc
// @ID GetUser
// @Summary Get specific user information
// @Description Get information of a user with a specific ID.
// @Tags [Sample API] Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} GetUserResponse "(Sample) This is a sample description for success response in Swagger UI"
// @Failure 404 {object} object "User Not Found"
// @Router /sample/users/{id} [get]
func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	// Implement user retrieval logic (this is a simple example)
	if id != 1 {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	// In this example, hardcoded data is returned
	user := model.MyUser{ID: 1, Name: "John Doe", Email: "john@example.com"}
	return c.JSON(http.StatusOK, user)
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type CreateUserRequest struct {
	model.MyUser
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type CreateUserResponse struct {
	model.MyUser
}

// CreateUser godoc
// @ID CreateUser
// @Summary Create a new user
// @Description Create a new user with the given information.
// @Tags [Sample API] Users
// @Accept  json
// @Produce  json
// @Param User body CreateUserRequest true "User information"
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 201 {object} GetUserResponse "(Sample) This is a sample description for success response in Swagger UI"
// @Failure 400 {object} object "Invalid Request"
// @Router /sample/users [post]
func CreateUser(c echo.Context) error {
	u := new(model.MyUser)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	// Implement user creation logic (this is a simple example)
	u.ID = 100 // Unique ID generation logic needed in actual implementation

	return c.JSON(http.StatusCreated, u)
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type UpdateUserRequest struct {
	model.MyUser
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type UpdateUserResponse struct {
	model.MyUser
}

// UpdateUser godoc
// @ID UpdateUser
// @Summary Update a user
// @Description Update a user with the given information.
// @Tags [Sample API] Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param User body UpdateUserRequest true "User information to update"
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 201 {object} UpdateUserResponse "(Sample) This is a sample description for success response in Swagger UI"
// @Failure 400 {object} object "Invalid Request"
// @Router /sample/users/{id} [put]
func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	u := new(model.MyUser)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Implement user update logic (this is a simple example)
	u.ID = id // Update the information of the user with the corresponding ID in the actual implementation

	return c.JSON(http.StatusOK, u)
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type PatchUserRequest struct {
	model.MyUser
}

// [Note]
// Struct Embedding is used to inherit the fields of MyUser
type PatchUserResponse struct {
	model.MyUser
}

// PatchUser godoc
// @ID PatchUser
// @Summary Patch a user
// @Description Patch a user with the given information.
// @Tags [Sample API] Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param User body PatchUserRequest true "User information to update"
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} PatchUserResponse "(Sample) This is a sample description for success response in Swagger UI"
// @Failure 400 {object} object "Invalid Request"
// @Failure 404 {object} object "User Not Found"
// @Router /sample/users/{id} [patch]
func PatchUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	u := new(model.MyUser)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Implement user update logic (this is a simple example)
	u.ID = id // Update the information of the user with the corresponding ID in the actual implementation

	return c.JSON(http.StatusOK, u)
}

// [Note]
// No RequestBody required for "DELETE /users/{id}"

// [Note]
// No ResponseBody required for "DELETE /users/{id}"

// DeleteUser godoc
// @ID DeleteUser
// @Summary Delete a user
// @Description Delete a user with the given information.
// @Tags [Sample API] Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {string} string "User deletion successful"
// @Failure 400 {object} object "Invalid Request"
// @Failure 404 {object} object "User Not Found"
// @Router /sample/users/{id} [delete]
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	// Implement user update logic (this is a simple example)
	// In this example, hardcoded data is returned
	user := model.MyUser{ID: 1, Name: "John Doe", Email: "john@example.com"}
	if id != user.ID {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, "User deletion successful")
}
