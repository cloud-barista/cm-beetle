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

// Package controller handles REST API requests for CM-Beetle.
package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/artifact"
	"github.com/labstack/echo/v4"
)

// CreateArtifact godoc
// @ID CreateArtifact
// @Summary Create a migration artifact with automatic schema upcasting and lineage tracking
// @Description Creates and persists a migration artifact (`source`, `planning`, or `target`) across
// @Description `infra`, `software`, or `data` categories. When `parentId` is provided, lineage fields
// @Description (`sourceGroupId`, `sourceModelId`, `planningModelId`, `targetModelId`) are inherited automatically.
// @Description Legacy JSON payloads are automatically normalized to the latest schema using reflect upcasting.
// @Tags [Artifact] Migration Artifact Management
// @Accept json
// @Produce json
// @Param Artifact body artifact.CreateReq true "Artifact creation request"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 201 {object} model.ApiResponse[artifact.RawArtifact] "Created migration artifact"
// @Failure 400 {object} model.ApiResponse[any] "Invalid artifact payload or metadata"
// @Router /artifacts [post]
func CreateArtifact(c echo.Context) error {
	var req artifact.CreateReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request body: "+err.Error()))
	}

	created, err := artifact.Create(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, model.SuccessResponse(*created))
}

// ListArtifacts godoc
// @ID ListArtifacts
// @Summary List migration artifacts with optional filtering
// @Description Lists stored migration artifacts filtered by `category` (`infra`, `software`, `data`),
// @Description `stage` (`source`, `planning`, `target`), `schemaType` (e.g., `OnpremInfra`, `RecommendedInfra`),
// @Description or `sourceGroupId`.
// @Tags [Artifact] Migration Artifact Management
// @Produce json
// @Param category query string false "Artifact domain category" Enums(infra,software,data)
// @Param stage query string false "Artifact lifecycle stage" Enums(source,planning,target)
// @Param schemaType query string false "Target Go struct schema type (e.g., OnpremInfra, RecommendedInfra, SourceObjectStorage)"
// @Param sourceGroupId query string false "Source group ID for lineage filtering"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[[]artifact.RawArtifact] "Filtered migration artifacts"
// @Failure 400 {object} model.ApiResponse[any] "Invalid filter query parameters"
// @Router /artifacts [get]
func ListArtifacts(c echo.Context) error {
	filter := artifact.Filter{
		Category:      artifact.Category(strings.TrimSpace(c.QueryParam("category"))),
		Stage:         artifact.Stage(strings.TrimSpace(c.QueryParam("stage"))),
		SchemaType:    strings.TrimSpace(c.QueryParam("schemaType")),
		SourceGroupId: strings.TrimSpace(c.QueryParam("sourceGroupId")),
	}

	items, err := artifact.List(filter)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, model.SuccessListResponse(items))
}

// GetArtifact godoc
// @ID GetArtifact
// @Summary Get a migration artifact by ID
// @Description Retrieves a single migration artifact by its unique ID and ensures its payload is normalized to the latest schema.
// @Tags [Artifact] Migration Artifact Management
// @Produce json
// @Param id path string true "Artifact ID"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[artifact.RawArtifact] "Migration artifact"
// @Failure 404 {object} model.ApiResponse[any] "Artifact not found"
// @Router /artifacts/{id} [get]
func GetArtifact(c echo.Context) error {
	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Path parameter 'id' is required"))
	}

	item, err := artifact.Get(id)
	if err != nil {
		if errors.Is(err, artifact.ErrArtifactNotFound) {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, model.SuccessResponse(*item))
}

// UpdateArtifact godoc
// @ID UpdateArtifact
// @Summary Update an existing migration artifact in-place
// @Description Overwrites an existing migration artifact's metadata or payload by ID and normalizes the payload to the latest schema.
// @Tags [Artifact] Migration Artifact Management
// @Accept json
// @Produce json
// @Param id path string true "Artifact ID"
// @Param Artifact body artifact.UpdateReq true "Artifact update request"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[artifact.RawArtifact] "Updated migration artifact"
// @Failure 400 {object} model.ApiResponse[any] "Invalid update payload"
// @Failure 404 {object} model.ApiResponse[any] "Artifact not found"
// @Router /artifacts/{id} [put]
func UpdateArtifact(c echo.Context) error {
	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Path parameter 'id' is required"))
	}

	var req artifact.UpdateReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request body: "+err.Error()))
	}

	updated, err := artifact.Update(id, req)
	if err != nil {
		if errors.Is(err, artifact.ErrArtifactNotFound) {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, model.SuccessResponse(*updated))
}

// DeleteArtifact godoc
// @ID DeleteArtifact
// @Summary Delete a migration artifact by ID
// @Description Removes a migration artifact file and its index entry by ID.
// @Tags [Artifact] Migration Artifact Management
// @Produce json
// @Param id path string true "Artifact ID"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[any] "Artifact deleted successfully"
// @Failure 404 {object} model.ApiResponse[any] "Artifact not found"
// @Router /artifacts/{id} [delete]
func DeleteArtifact(c echo.Context) error {
	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Path parameter 'id' is required"))
	}

	if err := artifact.Delete(id); err != nil {
		if errors.Is(err, artifact.ErrArtifactNotFound) {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, model.SimpleSuccessResponse("Artifact deleted successfully: "+id))
}

// GetArtifactLineage godoc
// @ID GetArtifactLineage
// @Summary Get the derivation lineage chain and group artifacts for an artifact ID
// @Description Returns the connected `source`, `planning`, and `target` artifacts in the derivation chain
// @Description along with all artifacts sharing the same `sourceGroupId`.
// @Tags [Artifact] Migration Artifact Management
// @Produce json
// @Param id path string true "Artifact ID"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[artifact.LineageGraph] "Artifact lineage graph"
// @Failure 404 {object} model.ApiResponse[any] "Artifact not found"
// @Router /artifacts/{id}/lineage [get]
func GetArtifactLineage(c echo.Context) error {
	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Path parameter 'id' is required"))
	}

	graph, err := artifact.GetLineageGraph(id)
	if err != nil {
		if errors.Is(err, artifact.ErrArtifactNotFound) {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, model.SuccessResponse(*graph))
}

