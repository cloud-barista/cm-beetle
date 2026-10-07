package lineage_test

import (
	"testing"

	"github.com/cloud-barista/cm-beetle/imdl/lineage"
)

func TestMigrationLineage_ToLabelMap(t *testing.T) {
	// Verify label mapping with all migration lineage fields populated
	lin := lineage.MigrationLineage{
		SourceGroupId:   "sg-01",
		SourceModelId:   "src-01",
		PlanningModelId: "plan-01",
		TargetModelId:   "tgt-01",
	}

	labels := lin.ToLabelMap()
	if labels[lineage.LabelSourceGroupId] != "sg-01" {
		t.Fatalf("expected sg-01, got %s", labels[lineage.LabelSourceGroupId])
	}
	if labels[lineage.LabelSourceModelId] != "src-01" {
		t.Fatalf("expected src-01, got %s", labels[lineage.LabelSourceModelId])
	}
	if labels[lineage.LabelPlanningModelId] != "plan-01" {
		t.Fatalf("expected plan-01, got %s", labels[lineage.LabelPlanningModelId])
	}
	if labels[lineage.LabelTargetModelId] != "tgt-01" {
		t.Fatalf("expected tgt-01, got %s", labels[lineage.LabelTargetModelId])
	}
}

func TestMigrationLineage_FromLabelMap(t *testing.T) {
	// Test roundtrip conversion from label map to struct
	raw := map[string]string{
		lineage.LabelSourceGroupId:   "sg-02",
		lineage.LabelSourceModelId:   "src-02",
		lineage.LabelPlanningModelId: "plan-02",
		lineage.LabelTargetModelId:   "tgt-02",
	}

	lin := lineage.FromLabelMap(raw)
	if lin.SourceGroupId != "sg-02" || lin.SourceModelId != "src-02" || lin.PlanningModelId != "plan-02" || lin.TargetModelId != "tgt-02" {
		t.Fatalf("unexpected unmarshaled struct: %+v", lin)
	}

	// Verify nil map returns empty struct without panicking
	emptyLin := lineage.FromLabelMap(nil)
	if emptyLin.TargetModelId != "" {
		t.Fatalf("expected empty struct on nil map, got %+v", emptyLin)
	}
}
