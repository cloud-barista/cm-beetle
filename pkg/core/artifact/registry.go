package artifact

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremisemodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	"github.com/cloud-barista/cm-beetle/transx"
)

// SchemaDescriptor holds reflection metadata and default classification for a registered schema type.
type SchemaDescriptor struct {
	SchemaType      string
	ReflectType     reflect.Type
	DefaultCategory Category
	DefaultStage    Stage
}

var registeredSchemas = map[string]SchemaDescriptor{
	"OnpremInfra": {
		SchemaType:      "OnpremInfra",
		ReflectType:     reflect.TypeOf(onpremisemodel.OnpremInfra{}),
		DefaultCategory: CategoryInfra,
		DefaultStage:    StageSource,
	},
	"RecommendedInfra": {
		SchemaType:      "RecommendedInfra",
		ReflectType:     reflect.TypeOf(cloudmodel.RecommendedInfra{}),
		DefaultCategory: CategoryInfra,
		DefaultStage:    StageTarget,
	},
	"SourceObjectStorage": {
		SchemaType:      "SourceObjectStorage",
		ReflectType:     reflect.TypeOf(storagemodel.SourceObjectStorage{}),
		DefaultCategory: CategoryInfra,
		DefaultStage:    StageSource,
	},
	"RecommendedObjectStorage": {
		SchemaType:      "RecommendedObjectStorage",
		ReflectType:     reflect.TypeOf(storagemodel.RecommendedObjectStorage{}),
		DefaultCategory: CategoryInfra,
		DefaultStage:    StageTarget,
	},
	"SourceRDBMS": {
		SchemaType:      "SourceRDBMS",
		ReflectType:     reflect.TypeOf(rdbmsmodel.SourceRDBMS{}),
		DefaultCategory: CategoryInfra,
		DefaultStage:    StageSource,
	},
	"RecommendedRDBMS": {
		SchemaType:      "RecommendedRDBMS",
		ReflectType:     reflect.TypeOf(rdbmsmodel.RecommendedRDBMS{}),
		DefaultCategory: CategoryInfra,
		DefaultStage:    StageTarget,
	},
	"DataMigrationModel": {
		SchemaType:      "DataMigrationModel",
		ReflectType:     reflect.TypeOf(transx.DataMigrationModel{}),
		DefaultCategory: CategoryData,
		DefaultStage:    StageTarget,
	},
}

// ResolveSchema resolves the target reflect.Type, schemaType name, Category, and Stage for a payload.
func ResolveSchema(schemaType string, category Category, stage Stage, rawPayload []byte) (reflect.Type, string, Category, Stage) {
	resolvedType := strings.TrimSpace(schemaType)
	if resolvedType != "" {
		if desc, ok := lookupDescriptorCaseInsensitive(resolvedType); ok {
			if category == "" {
				category = desc.DefaultCategory
			}
			if stage == "" {
				stage = desc.DefaultStage
			}
			return desc.ReflectType, desc.SchemaType, category, stage
		}
	}

	inferredName := inferSchemaTypeFromJSON(rawPayload, category, stage)
	if inferredName != "" {
		if desc, ok := registeredSchemas[inferredName]; ok {
			if category == "" {
				category = desc.DefaultCategory
			}
			if stage == "" {
				stage = desc.DefaultStage
			}
			if resolvedType == "" {
				resolvedType = desc.SchemaType
			}
			return desc.ReflectType, resolvedType, category, stage
		}
	}

	return nil, resolvedType, category, stage
}

func lookupDescriptorCaseInsensitive(name string) (SchemaDescriptor, bool) {
	if desc, ok := registeredSchemas[name]; ok {
		return desc, true
	}
	for k, desc := range registeredSchemas {
		if strings.EqualFold(k, name) {
			return desc, true
		}
	}
	return SchemaDescriptor{}, false
}

func inferSchemaTypeFromJSON(rawPayload []byte, category Category, stage Stage) string {
	var topMap map[string]any
	if err := json.Unmarshal(rawPayload, &topMap); err != nil {
		return ""
	}

	switch {
	case hasAnyKey(topMap, "sourceObjectStorageModel", "sourceObjectStorages", "buckets", "objectStorages"):
		return "SourceObjectStorage"
	case hasAnyKey(topMap, "recommendedObjectStorageModel", "targetObjectStorages", "targetBuckets"):
		return "RecommendedObjectStorage"
	case hasAnyKey(topMap, "sourceRDBMSModel", "sourceRDBMSInstances", "rdbmsList", "sourceRdbmsList"):
		return "SourceRDBMS"
	case hasAnyKey(topMap, "recommendedRDBMSModel", "targetRDBMSInstances", "targetRdbmsList"):
		return "RecommendedRDBMS"
	case hasAnyKey(topMap, "recommendedInfraModel", "cloudInfraModel", "targetInfra", "targetVmInfra", "targetSpecList", "targetVmSpecList"):
		return "RecommendedInfra"
	case hasAnyKey(topMap, "onpremiseInfraModel", "nodes", "servers", "vms", "k8sCluster", "nlbs"):
		return "OnpremInfra"
	case hasAnyKey(topMap, "dataMigrationModel") || (hasAnyKey(topMap, "source") && hasAnyKey(topMap, "destination")):
		return "DataMigrationModel"
	}

	if category == CategoryInfra {
		switch stage {
		case StageSource, StagePlanning:
			return "OnpremInfra"
		case StageTarget:
			return "RecommendedInfra"
		}
	}
	if category == CategoryData {
		return "DataMigrationModel"
	}

	return ""
}

func hasAnyKey(m map[string]any, keys ...string) bool {
	for _, k := range keys {
		if _, exists := m[k]; exists {
			return true
		}
	}
	return false
}

// ValidateTypeGuard verifies that the requested Go type matches the stored artifact schemaType.
func ValidateTypeGuard(storedSchemaType string, requestedType reflect.Type) error {
	for requestedType.Kind() == reflect.Ptr {
		requestedType = requestedType.Elem()
	}
	reqName := requestedType.Name()
	if reqName == "" || reqName == "RawMessage" || storedSchemaType == "" {
		return nil
	}
	if !strings.EqualFold(storedSchemaType, reqName) {
		return fmt.Errorf("schemaType mismatch: artifact has schemaType '%s', cannot cast to '%s'", storedSchemaType, reqName)
	}
	return nil
}

