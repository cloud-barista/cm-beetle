package artifact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloud-barista/cm-beetle/imdl/lineage"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
)

// CurrentSchemaVersion defines the current imdl schema version stamped on normalized artifacts.
const CurrentSchemaVersion = "v0.1.18"

// Category classifies the domain of a migration artifact.
type Category string

const (
	CategoryInfra    Category = "infra"
	CategorySoftware Category = "software"
	CategoryData     Category = "data"
)

// IsValid checks whether Category is one of the supported domain categories.
func (c Category) IsValid() bool {
	switch c {
	case CategoryInfra, CategorySoftware, CategoryData:
		return true
	default:
		return false
	}
}

// Stage classifies the migration lifecycle stage of an artifact.
type Stage string

const (
	StageSource   Stage = "source"
	StagePlanning Stage = "planning"
	StageTarget   Stage = "target"
)

// IsValid checks whether Stage is one of the supported lifecycle stages.
func (s Stage) IsValid() bool {
	switch s {
	case StageSource, StagePlanning, StageTarget:
		return true
	default:
		return false
	}
}

// Artifact is a generic envelope for any migration payload type T.
type Artifact[T any] struct {
	Id               string                      `json:"id" example:"art-d0g535l40272926020g0"`
	Name             string                      `json:"name" validate:"required" example:"ecommerce-prod-plan"`
	Description      string                      `json:"description,omitempty" example:"Refined migration artifact"`
	Version          string                      `json:"version,omitempty" example:"1.0.0"`
	Category         Category                    `json:"category" validate:"required" enums:"infra,software,data" example:"infra"`
	Stage            Stage                       `json:"stage" validate:"required" enums:"source,planning,target" example:"planning"`
	SchemaType       string                      `json:"schemaType,omitempty" example:"OnpremInfra"`
	SchemaVersion    string                      `json:"schemaVersion,omitempty" example:"v0.1.18"`
	Lineage          lineage.MigrationLineage    `json:"lineage"`
	Payload          T                           `json:"payload" validate:"required"`
	ConversionReport *modelconv.ConversionReport `json:"conversionReport,omitempty"`
	CreateTime       string                      `json:"createTime" example:"2026-10-23T12:00:00Z"`
	UpdateTime       string                      `json:"updateTime,omitempty" example:"2026-10-23T12:05:00Z"`
}

// RawArtifact is the untyped storage and wire representation of an Artifact.
type RawArtifact = Artifact[json.RawMessage]

// CreateMetadata holds metadata fields used when creating a typed artifact via CreateAs[T].
type CreateMetadata struct {
	Id            string                   `json:"id,omitempty"`
	Name          string                   `json:"name"`
	Description   string                   `json:"description,omitempty"`
	Version       string                   `json:"version,omitempty"`
	Category      Category                 `json:"category,omitempty"`
	Stage         Stage                    `json:"stage,omitempty"`
	SchemaType    string                   `json:"schemaType,omitempty"`
	ParentId      string                   `json:"parentId,omitempty"`
	SourceGroupId string                   `json:"sourceGroupId,omitempty"`
	Lineage       lineage.MigrationLineage `json:"lineage,omitempty"`
}

// CreateReq represents the JSON request body for POST /beetle/artifacts.
type CreateReq struct {
	Id            string                   `json:"id,omitempty" example:"art-custom-01"`
	Name          string                   `json:"name" validate:"required" example:"ecommerce-source-infra"`
	Description   string                   `json:"description,omitempty" example:"Extracted source infrastructure"`
	Version       string                   `json:"version,omitempty" example:"1.0.0"`
	Category      Category                 `json:"category" validate:"required" enums:"infra,software,data" example:"infra"`
	Stage         Stage                    `json:"stage" validate:"required" enums:"source,planning,target" example:"source"`
	SchemaType    string                   `json:"schemaType,omitempty" example:"OnpremInfra"`
	ParentId      string                   `json:"parentId,omitempty" example:"art-src-01"`
	SourceGroupId string                   `json:"sourceGroupId,omitempty" example:"sg-ecommerce-prod"`
	Lineage       lineage.MigrationLineage `json:"lineage,omitempty"`
	Payload       json.RawMessage          `json:"payload" validate:"required"`

	// Legacy envelope compatibility fields accepted on import and normalized automatically.
	LegacyModelType     string          `json:"modelType,omitempty" swaggerignore:"true"`
	LegacyOnpremModel   json.RawMessage `json:"onpremiseInfraModel,omitempty" swaggerignore:"true"`
	LegacyCloudModel    json.RawMessage `json:"cloudInfraModel,omitempty" swaggerignore:"true"`
	LegacyRecInfraModel json.RawMessage `json:"recommendedInfraModel,omitempty" swaggerignore:"true"`
}

// Normalize fills missing Category, Stage, SchemaType, and Payload from legacy fields when provided.
func (r *CreateReq) Normalize() error {
	if len(r.Payload) == 0 {
		switch {
		case len(r.LegacyOnpremModel) > 0:
			r.Payload = r.LegacyOnpremModel
			if r.Category == "" {
				r.Category = CategoryInfra
			}
			if r.Stage == "" {
				r.Stage = StageSource
			}
		case len(r.LegacyCloudModel) > 0:
			r.Payload = r.LegacyCloudModel
			if r.Category == "" {
				r.Category = CategoryInfra
			}
			if r.Stage == "" {
				r.Stage = StageTarget
			}
		case len(r.LegacyRecInfraModel) > 0:
			r.Payload = r.LegacyRecInfraModel
			if r.Category == "" {
				r.Category = CategoryInfra
			}
			if r.Stage == "" {
				r.Stage = StageTarget
			}
		}
	}

	if r.Stage == "" && r.LegacyModelType != "" {
		switch strings.ToLower(strings.TrimSpace(r.LegacyModelType)) {
		case "onprem", "source":
			r.Stage = StageSource
		case "planning", "plan":
			r.Stage = StagePlanning
		case "cloud", "target", "recommended":
			r.Stage = StageTarget
		}
	}

	if strings.TrimSpace(r.Name) == "" {
		return fmt.Errorf("artifact 'name' is required")
	}
	if len(bytes.TrimSpace(r.Payload)) == 0 {
		return fmt.Errorf("artifact 'payload' is required")
	}
	return nil
}

// UpdateReq represents the JSON request body for PUT /beetle/artifacts/:id.
type UpdateReq struct {
	Name          string                   `json:"name,omitempty" example:"ecommerce-plan-v2"`
	Description   string                   `json:"description,omitempty" example:"Updated node spec"`
	Version       string                   `json:"version,omitempty" example:"1.1.0"`
	Category      Category                 `json:"category,omitempty" enums:"infra,software,data" example:"infra"`
	Stage         Stage                    `json:"stage,omitempty" enums:"source,planning,target" example:"planning"`
	SchemaType    string                   `json:"schemaType,omitempty" example:"OnpremInfra"`
	ParentId      string                   `json:"parentId,omitempty"`
	SourceGroupId string                   `json:"sourceGroupId,omitempty"`
	Lineage       lineage.MigrationLineage `json:"lineage,omitempty"`
	Payload       json.RawMessage          `json:"payload,omitempty"`

	// Legacy envelope compatibility fields accepted on update and normalized automatically.
	LegacyOnpremModel   json.RawMessage `json:"onpremiseInfraModel,omitempty" swaggerignore:"true"`
	LegacyCloudModel    json.RawMessage `json:"cloudInfraModel,omitempty" swaggerignore:"true"`
	LegacyRecInfraModel json.RawMessage `json:"recommendedInfraModel,omitempty" swaggerignore:"true"`
}

// Normalize populates Payload from legacy wrapper fields if Payload was omitted.
func (r *UpdateReq) Normalize() {
	if len(r.Payload) == 0 {
		switch {
		case len(r.LegacyOnpremModel) > 0:
			r.Payload = r.LegacyOnpremModel
		case len(r.LegacyCloudModel) > 0:
			r.Payload = r.LegacyCloudModel
		case len(r.LegacyRecInfraModel) > 0:
			r.Payload = r.LegacyRecInfraModel
		}
	}
}

// Filter defines query parameters for listing artifacts via GET /beetle/artifacts.
type Filter struct {
	Category      Category `json:"category,omitempty"`
	Stage         Stage    `json:"stage,omitempty"`
	SchemaType    string   `json:"schemaType,omitempty"`
	SourceGroupId string   `json:"sourceGroupId,omitempty"`
}

// LineageGraph represents the derivation chain and group-related artifacts for GET /beetle/artifacts/:id/lineage.
type LineageGraph struct {
	ArtifactId       string                   `json:"artifactId"`
	SourceGroupId    string                   `json:"sourceGroupId,omitempty"`
	Lineage          lineage.MigrationLineage `json:"lineage"`
	SourceArtifact   *RawArtifact             `json:"sourceArtifact,omitempty"`
	PlanningArtifact *RawArtifact             `json:"planningArtifact,omitempty"`
	TargetArtifact   *RawArtifact             `json:"targetArtifact,omitempty"`
	GroupArtifacts   []RawArtifact            `json:"groupArtifacts,omitempty"`
}

// ListResponse wraps a slice of artifacts for consistent REST API responses.
type ListResponse struct {
	Count     int           `json:"count" example:"2"`
	Artifacts []RawArtifact `json:"artifacts"`
}
