package lineage

// Standard MigrationLineage label keys.
const (
	LabelSourceGroupId    = "cm-source-group-id"
	LabelSourceModelId    = "cm-source-model-id"
	LabelPlanningModelId  = "cm-planning-model-id"
	LabelTargetModelId    = "cm-target-model-id"
	LabelSourceMachineIds = "cm-source-machine-ids"
	LabelMigratedAt       = "cm-migrated-at"
)

// MigrationLineage encapsulates migration traceability metadata across Cloud-Barista.
type MigrationLineage struct {
	SourceGroupId   string `json:"sourceGroupId,omitempty" example:"sg-ecommerce-prod"`
	SourceModelId   string `json:"sourceModelId,omitempty" example:"src-mdl-101"`
	PlanningModelId string `json:"planningModelId,omitempty" example:"plan-mdl-101"`
	TargetModelId   string `json:"targetModelId,omitempty" example:"tgt-mdl-201"`
}

// ToLabelMap converts MigrationLineage into a standard label map for cloud resources.
func (l MigrationLineage) ToLabelMap() map[string]string {
	labels := make(map[string]string)
	if l.SourceGroupId != "" {
		labels[LabelSourceGroupId] = l.SourceGroupId
	}
	if l.SourceModelId != "" {
		labels[LabelSourceModelId] = l.SourceModelId
	}
	if l.PlanningModelId != "" {
		labels[LabelPlanningModelId] = l.PlanningModelId
	}
	if l.TargetModelId != "" {
		labels[LabelTargetModelId] = l.TargetModelId
	}
	return labels
}

// FromLabelMap reconstructs MigrationLineage from a label map.
func FromLabelMap(labels map[string]string) MigrationLineage {
	if labels == nil {
		return MigrationLineage{}
	}
	return MigrationLineage{
		SourceGroupId:   labels[LabelSourceGroupId],
		SourceModelId:   labels[LabelSourceModelId],
		PlanningModelId: labels[LabelPlanningModelId],
		TargetModelId:   labels[LabelTargetModelId],
	}
}
