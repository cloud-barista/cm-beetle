package artifact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/cloud-barista/cm-beetle/imdl/lineage"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
	"github.com/rs/xid"
)

// Create validates, upcasts, and persists a new artifact in the default store.
func Create(req CreateReq) (*RawArtifact, error) {
	st, err := GetDefaultStore()
	if err != nil {
		return nil, err
	}
	return CreateInStore(st, req)
}

// CreateInStore validates, upcasts, and persists a new artifact in the given store.
func CreateInStore(st *Store, req CreateReq) (*RawArtifact, error) {
	if err := req.Normalize(); err != nil {
		return nil, err
	}

	_, resolvedSchemaType, resolvedCategory, resolvedStage := ResolveSchema(
		req.SchemaType,
		req.Category,
		req.Stage,
		req.Payload,
	)

	if !resolvedCategory.IsValid() {
		return nil, fmt.Errorf("invalid or missing category '%s': must be one of 'infra', 'software', 'data'", resolvedCategory)
	}
	if !resolvedStage.IsValid() {
		return nil, fmt.Errorf("invalid or missing stage '%s': must be one of 'source', 'planning', 'target'", resolvedStage)
	}

	id := strings.TrimSpace(req.Id)
	if id == "" {
		id = "art-" + xid.New().String()
	}

	lin, err := resolveLineage(st, id, resolvedStage, req.ParentId, req.SourceGroupId, req.Lineage)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC().Format(time.RFC3339)
	version := strings.TrimSpace(req.Version)
	if version == "" {
		version = "1.0.0"
	}

	art := RawArtifact{
		Id:            id,
		Name:          strings.TrimSpace(req.Name),
		Description:   strings.TrimSpace(req.Description),
		Version:       version,
		Category:      resolvedCategory,
		Stage:         resolvedStage,
		SchemaType:    resolvedSchemaType,
		SchemaVersion: CurrentSchemaVersion,
		Lineage:       lin,
		Payload:       req.Payload,
		CreateTime:    now,
		UpdateTime:    now,
	}

	normalized, err := normalizeAndUpcastArtifact(art)
	if err != nil {
		return nil, err
	}

	if err := st.Save(normalized); err != nil {
		return nil, err
	}
	return &normalized, nil
}

// Get retrieves an artifact by ID and ensures its payload is upcasted to the latest registered schema.
func Get(id string) (*RawArtifact, error) {
	st, err := GetDefaultStore()
	if err != nil {
		return nil, err
	}
	return GetFromStore(st, id)
}

// GetFromStore retrieves an artifact by ID from the given store and ensures latest schema normalization.
func GetFromStore(st *Store, id string) (*RawArtifact, error) {
	art, err := st.Get(id)
	if err != nil {
		return nil, err
	}
	normalized, err := normalizeAndUpcastArtifact(art)
	if err != nil {
		return nil, err
	}
	return &normalized, nil
}

// List returns all artifacts matching Filter from the default store.
func List(filter Filter) ([]RawArtifact, error) {
	st, err := GetDefaultStore()
	if err != nil {
		return nil, err
	}
	return ListFromStore(st, filter)
}

// ListFromStore returns all artifacts matching Filter from the given store.
func ListFromStore(st *Store, filter Filter) ([]RawArtifact, error) {
	if filter.Category != "" && !filter.Category.IsValid() {
		return nil, fmt.Errorf("invalid category filter '%s': must be one of 'infra', 'software', 'data'", filter.Category)
	}
	if filter.Stage != "" && !filter.Stage.IsValid() {
		return nil, fmt.Errorf("invalid stage filter '%s': must be one of 'source', 'planning', 'target'", filter.Stage)
	}
	items := st.List(filter)
	out := make([]RawArtifact, 0, len(items))
	for _, item := range items {
		norm, err := normalizeAndUpcastArtifact(item)
		if err != nil {
			out = append(out, item)
			continue
		}
		out = append(out, norm)
	}
	return out, nil
}

// Update modifies an existing artifact in-place, upcasts the updated payload, and persists it atomically.
func Update(id string, req UpdateReq) (*RawArtifact, error) {
	st, err := GetDefaultStore()
	if err != nil {
		return nil, err
	}
	return UpdateInStore(st, id, req)
}

// UpdateInStore modifies an existing artifact in the given store.
func UpdateInStore(st *Store, id string, req UpdateReq) (*RawArtifact, error) {
	req.Normalize()

	existing, err := st.Get(id)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(req.Name) != "" {
		existing.Name = strings.TrimSpace(req.Name)
	}
	if req.Description != "" {
		existing.Description = strings.TrimSpace(req.Description)
	}
	if strings.TrimSpace(req.Version) != "" {
		existing.Version = strings.TrimSpace(req.Version)
	}
	if req.Category != "" {
		if !req.Category.IsValid() {
			return nil, fmt.Errorf("invalid category '%s': must be one of 'infra', 'software', 'data'", req.Category)
		}
		existing.Category = req.Category
	}
	if req.Stage != "" {
		if !req.Stage.IsValid() {
			return nil, fmt.Errorf("invalid stage '%s': must be one of 'source', 'planning', 'target'", req.Stage)
		}
		existing.Stage = req.Stage
	}
	if strings.TrimSpace(req.SchemaType) != "" {
		existing.SchemaType = strings.TrimSpace(req.SchemaType)
	}
	if len(bytes.TrimSpace(req.Payload)) > 0 {
		existing.Payload = req.Payload
	}

	mergedLineage := existing.Lineage
	if req.Lineage.SourceGroupId != "" {
		mergedLineage.SourceGroupId = req.Lineage.SourceGroupId
	}
	if req.Lineage.SourceModelId != "" {
		mergedLineage.SourceModelId = req.Lineage.SourceModelId
	}
	if req.Lineage.PlanningModelId != "" {
		mergedLineage.PlanningModelId = req.Lineage.PlanningModelId
	}
	if req.Lineage.TargetModelId != "" {
		mergedLineage.TargetModelId = req.Lineage.TargetModelId
	}

	lin, err := resolveLineage(st, existing.Id, existing.Stage, req.ParentId, req.SourceGroupId, mergedLineage)
	if err != nil {
		return nil, err
	}
	existing.Lineage = lin
	existing.UpdateTime = time.Now().UTC().Format(time.RFC3339)

	normalized, err := normalizeAndUpcastArtifact(existing)
	if err != nil {
		return nil, err
	}

	if err := st.Save(normalized); err != nil {
		return nil, err
	}
	return &normalized, nil
}

// Delete removes an artifact by ID from the default store.
func Delete(id string) error {
	st, err := GetDefaultStore()
	if err != nil {
		return err
	}
	return st.Delete(id)
}

// GetLineageGraph builds the full derivation chain and group-related artifacts for an artifact ID.
func GetLineageGraph(id string) (*LineageGraph, error) {
	st, err := GetDefaultStore()
	if err != nil {
		return nil, err
	}
	return GetLineageGraphFromStore(st, id)
}

// GetLineageGraphFromStore builds the full derivation chain from the given store.
func GetLineageGraphFromStore(st *Store, id string) (*LineageGraph, error) {
	art, err := GetFromStore(st, id)
	if err != nil {
		return nil, err
	}

	graph := &LineageGraph{
		ArtifactId:    art.Id,
		SourceGroupId: art.Lineage.SourceGroupId,
		Lineage:       art.Lineage,
	}

	if art.Lineage.SourceModelId != "" {
		if src, err := GetFromStore(st, art.Lineage.SourceModelId); err == nil {
			graph.SourceArtifact = src
		}
	}
	if art.Lineage.PlanningModelId != "" {
		if plan, err := GetFromStore(st, art.Lineage.PlanningModelId); err == nil {
			graph.PlanningArtifact = plan
		}
	}
	if art.Lineage.TargetModelId != "" {
		if tgt, err := GetFromStore(st, art.Lineage.TargetModelId); err == nil {
			graph.TargetArtifact = tgt
		}
	}

	if art.Lineage.SourceGroupId != "" {
		groupItems, err := ListFromStore(st, Filter{SourceGroupId: art.Lineage.SourceGroupId})
		if err == nil {
			graph.GroupArtifacts = groupItems
		}
	}
	return graph, nil
}

// CreateAs stores a typed payload T as a normalized Artifact and returns the typed Artifact[T].
func CreateAs[T any](meta CreateMetadata, payload T) (*Artifact[T], error) {
	st, err := GetDefaultStore()
	if err != nil {
		return nil, err
	}
	return CreateAsInStore[T](st, meta, payload)
}

// CreateAsInStore stores a typed payload T in the specified store and returns the typed Artifact[T].
func CreateAsInStore[T any](st *Store, meta CreateMetadata, payload T) (*Artifact[T], error) {
	tType := reflect.TypeOf((*T)(nil)).Elem()
	for tType.Kind() == reflect.Ptr {
		tType = tType.Elem()
	}

	schemaType := strings.TrimSpace(meta.SchemaType)
	if schemaType == "" && tType.Name() != "" && tType.Name() != "RawMessage" {
		schemaType = tType.Name()
	}

	rawBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal typed payload: %w", err)
	}

	rawArt, err := CreateInStore(st, CreateReq{
		Id:            meta.Id,
		Name:          meta.Name,
		Description:   meta.Description,
		Version:       meta.Version,
		Category:      meta.Category,
		Stage:         meta.Stage,
		SchemaType:    schemaType,
		ParentId:      meta.ParentId,
		SourceGroupId: meta.SourceGroupId,
		Lineage:       meta.Lineage,
		Payload:       rawBytes,
	})
	if err != nil {
		return nil, err
	}
	return CastPayload[T](rawArt)
}

// GetAs retrieves an artifact by ID and casts/upcasts its payload into target struct T using reflect.
func GetAs[T any](id string) (*Artifact[T], error) {
	st, err := GetDefaultStore()
	if err != nil {
		return nil, err
	}
	return GetAsFromStore[T](st, id)
}

// GetAsFromStore retrieves an artifact by ID from the given store and casts/upcasts its payload into T.
func GetAsFromStore[T any](st *Store, id string) (*Artifact[T], error) {
	rawArt, err := GetFromStore(st, id)
	if err != nil {
		return nil, err
	}
	return CastPayload[T](rawArt)
}

// CastPayload validates schemaType compatibility and upcasts a RawArtifact into a typed Artifact[T].
func CastPayload[T any](raw *RawArtifact) (*Artifact[T], error) {
	if raw == nil {
		return nil, fmt.Errorf("cannot cast nil RawArtifact")
	}

	targetType := reflect.TypeOf((*T)(nil)).Elem()
	if err := ValidateTypeGuard(raw.SchemaType, targetType); err != nil {
		return nil, err
	}

	typedPayload, report, err := modelconv.UpcastTo[T](raw.Payload)
	if err != nil {
		return nil, err
	}

	finalReport := raw.ConversionReport
	if report != nil && report.Upcasted {
		finalReport = report
	}

	return &Artifact[T]{
		Id:               raw.Id,
		Name:             raw.Name,
		Description:      raw.Description,
		Version:          raw.Version,
		Category:         raw.Category,
		Stage:            raw.Stage,
		SchemaType:       raw.SchemaType,
		SchemaVersion:    raw.SchemaVersion,
		Lineage:          raw.Lineage,
		Payload:          typedPayload,
		ConversionReport: finalReport,
		CreateTime:       raw.CreateTime,
		UpdateTime:       raw.UpdateTime,
	}, nil
}

func resolveLineage(
	st *Store,
	selfId string,
	stage Stage,
	parentId string,
	sourceGroupId string,
	explicit lineage.MigrationLineage,
) (lineage.MigrationLineage, error) {
	result := explicit

	if strings.TrimSpace(parentId) != "" {
		cleanParentId := strings.TrimSpace(parentId)
		if parent, err := st.Get(cleanParentId); err == nil {
			if result.SourceGroupId == "" {
				result.SourceGroupId = parent.Lineage.SourceGroupId
			}
			if result.SourceModelId == "" {
				result.SourceModelId = parent.Lineage.SourceModelId
			}
			if result.PlanningModelId == "" {
				result.PlanningModelId = parent.Lineage.PlanningModelId
			}
			if result.TargetModelId == "" {
				result.TargetModelId = parent.Lineage.TargetModelId
			}

			switch parent.Stage {
			case StageSource:
				if result.SourceModelId == "" {
					result.SourceModelId = parent.Id
				}
			case StagePlanning:
				if result.PlanningModelId == "" {
					result.PlanningModelId = parent.Id
				}
			case StageTarget:
				if result.TargetModelId == "" {
					result.TargetModelId = parent.Id
				}
			}
		} else if result.SourceModelId == "" {
			result.SourceModelId = cleanParentId
		}
	}

	if strings.TrimSpace(sourceGroupId) != "" {
		result.SourceGroupId = strings.TrimSpace(sourceGroupId)
	}

	switch stage {
	case StageSource:
		result.SourceModelId = selfId
	case StagePlanning:
		result.PlanningModelId = selfId
	case StageTarget:
		result.TargetModelId = selfId
	}

	return result, nil
}

func normalizeAndUpcastArtifact(art RawArtifact) (RawArtifact, error) {
	targetType, schemaType, category, stage := ResolveSchema(art.SchemaType, art.Category, art.Stage, art.Payload)
	if schemaType != "" {
		art.SchemaType = schemaType
	}
	if category != "" {
		art.Category = category
	}
	if stage != "" {
		art.Stage = stage
	}
	if art.SchemaVersion == "" {
		art.SchemaVersion = CurrentSchemaVersion
	}

	if targetType != nil {
		canonicalBytes, report, err := modelconv.UpcastJSON(art.Payload, targetType)
		if err != nil {
			return RawArtifact{}, err
		}
		art.Payload = canonicalBytes
		if report != nil && report.Upcasted {
			art.ConversionReport = report
		}
		return art, nil
	}

	// For custom or software payloads without a registered struct, ensure valid JSON formatting.
	var check any
	if err := json.Unmarshal(art.Payload, &check); err != nil {
		return RawArtifact{}, fmt.Errorf("invalid JSON payload: %w", err)
	}
	return art, nil
}
