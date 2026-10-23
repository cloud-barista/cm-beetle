package modelconv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// FieldMigration records a legacy field or key that was mapped to a canonical target field.
type FieldMigration struct {
	Path string `json:"path" example:"targetVmInfra"`
	From string `json:"from" example:"targetVmInfra"`
	To   string `json:"to" example:"targetInfra"`
}

// ConversionReport summarizes all schema normalizations performed during reflect upcasting.
type ConversionReport struct {
	TargetSchemaType string           `json:"targetSchemaType" example:"RecommendedInfra"`
	Upcasted         bool             `json:"upcasted" example:"true"`
	MigratedFields   []FieldMigration `json:"migratedFields,omitempty"`
	DroppedFields    []string         `json:"droppedFields,omitempty"`
}

// DefaultLegacyFieldAliases maps current canonical JSON field names to known legacy field names.
var DefaultLegacyFieldAliases = map[string][]string{
	"targetInfra":          {"targetVmInfra", "vmInfra"},
	"nodeGroups":           {"subGroups", "vm", "vms", "vmGroups"},
	"nodeGroupSize":        {"subGroupSize", "vmGroupSize"},
	"nodeUserName":         {"vmUserName"},
	"nodeUserPassword":     {"vmUserPassword"},
	"targetSpecList":       {"targetVmSpecList"},
	"targetOsImageList":    {"targetVmOsImageList"},
	"targetSpec":           {"targetVmSpec"},
	"targetOsImage":        {"targetVmOsImage"},
	"nodes":                {"servers", "vms", "machines"},
	"sourceRDBMSInstances": {"rdbmsList", "sourceRdbmsList"},
	"targetRDBMSInstances": {"targetRdbmsList"},
	"sourceObjectStorages": {"buckets", "objectStorages"},
	"targetObjectStorages": {"targetBuckets"},
}

// DefaultLegacyWrapperKeys lists known single-field wrapper keys used in older envelopes.
var DefaultLegacyWrapperKeys = []string{
	"onpremiseInfraModel",
	"cloudInfraModel",
	"recommendedInfraModel",
	"recommendedVmInfraModel",
	"sourceObjectStorageModel",
	"recommendedObjectStorageModel",
	"sourceRDBMSModel",
	"recommendedRDBMSModel",
	"sourceNlbModel",
	"dataMigrationModel",
}

// DefaultLegacyLabelKeyAliases maps legacy label keys to canonical Cloud-Barista lineage label keys.
var DefaultLegacyLabelKeyAliases = map[string]string{
	"sourceMachineId":  "cm-source-machine-ids",
	"sourceMachineIds": "cm-source-machine-ids",
	"sourceGroupId":    "cm-source-group-id",
	"sourceModelId":    "cm-source-model-id",
	"planningModelId":  "cm-planning-model-id",
	"targetModelId":    "cm-target-model-id",
}

// UpcastTo converts raw JSON bytes into the target Go type T with reflect-based legacy field migration.
func UpcastTo[T any](rawJSON []byte) (T, *ConversionReport, error) {
	var zero T
	targetType := reflect.TypeOf((*T)(nil)).Elem()
	normalizedBytes, report, err := UpcastJSON(rawJSON, targetType)
	if err != nil {
		return zero, nil, err
	}

	var result T
	dec := json.NewDecoder(bytes.NewReader(normalizedBytes))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&result); err != nil {
		return zero, nil, fmt.Errorf("failed to decode normalized JSON into %s: %w", report.TargetSchemaType, err)
	}

	return result, report, nil
}

// UpcastJSON normalizes raw JSON bytes against targetType using reflect and returns canonical JSON bytes.
func UpcastJSON(rawJSON []byte, targetType reflect.Type) ([]byte, *ConversionReport, error) {
	if targetType == nil {
		return nil, nil, fmt.Errorf("target reflect.Type must not be nil")
	}
	for targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}

	trimmed := bytes.TrimSpace(rawJSON)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		return nil, nil, fmt.Errorf("input JSON payload is empty or null")
	}

	var rawVal any
	dec := json.NewDecoder(bytes.NewReader(trimmed))
	dec.UseNumber()
	if err := dec.Decode(&rawVal); err != nil {
		return nil, nil, fmt.Errorf("invalid JSON payload: %w", err)
	}

	report := &ConversionReport{
		TargetSchemaType: targetType.Name(),
	}

	normalizedVal, err := normalizeValue(rawVal, targetType, "", report)
	if err != nil {
		return nil, nil, err
	}

	outBytes, err := json.Marshal(normalizedVal)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal normalized value: %w", err)
	}

	var checkVal reflect.Value
	if targetType.Kind() == reflect.Struct {
		checkVal = reflect.New(targetType)
	} else {
		checkVal = reflect.New(targetType)
	}
	strictDec := json.NewDecoder(bytes.NewReader(outBytes))
	strictDec.DisallowUnknownFields()
	if err := strictDec.Decode(checkVal.Interface()); err != nil {
		return nil, nil, fmt.Errorf("schema validation failed for %s: %w", targetType.Name(), err)
	}

	canonicalBytes, err := json.Marshal(checkVal.Elem().Interface())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal canonical struct %s: %w", targetType.Name(), err)
	}

	sort.Strings(report.DroppedFields)
	report.Upcasted = len(report.MigratedFields) > 0 || len(report.DroppedFields) > 0
	return canonicalBytes, report, nil
}

func normalizeValue(rawVal any, targetType reflect.Type, path string, report *ConversionReport) (any, error) {
	if rawVal == nil {
		return nil, nil
	}
	for targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}

	switch targetType.Kind() {
	case reflect.Struct:
		objMap, ok := rawVal.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("type mismatch at '%s': expected JSON object for struct %s, got %T", pathOrRoot(path), targetType.Name(), rawVal)
		}
		return normalizeStructMap(objMap, targetType, path, report)

	case reflect.Slice:
		arr, ok := rawVal.([]any)
		if !ok {
			return nil, fmt.Errorf("type mismatch at '%s': expected JSON array for slice %s, got %T", pathOrRoot(path), targetType.String(), rawVal)
		}
		outSlice := make([]any, len(arr))
		elemType := targetType.Elem()
		for i, item := range arr {
			elemPath := fmt.Sprintf("%s[%d]", path, i)
			normItem, err := normalizeValue(item, elemType, elemPath, report)
			if err != nil {
				return nil, err
			}
			outSlice[i] = normItem
		}
		return outSlice, nil

	case reflect.Map:
		inMap, ok := rawVal.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("type mismatch at '%s': expected JSON object for map %s, got %T", pathOrRoot(path), targetType.String(), rawVal)
		}
		return normalizeMap(inMap, targetType, path, report)

	case reflect.String:
		return normalizeScalarString(rawVal, path)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return normalizeScalarInt(rawVal, path, report)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return normalizeScalarUint(rawVal, path, report)

	case reflect.Float32, reflect.Float64:
		return normalizeScalarFloat(rawVal, path, report)

	case reflect.Bool:
		return normalizeScalarBool(rawVal, path, report)

	default:
		return rawVal, nil
	}
}

type structFieldMeta struct {
	jsonName  string
	fieldName string
	fieldType reflect.Type
}

func collectStructFields(structType reflect.Type) map[string]structFieldMeta {
	fields := make(map[string]structFieldMeta)
	for i := 0; i < structType.NumField(); i++ {
		f := structType.Field(i)
		if f.PkgPath != "" && !f.Anonymous {
			continue
		}
		tag := f.Tag.Get("json")
		if tag == "-" {
			continue
		}
		tagName := tag
		if idx := strings.IndexByte(tag, ','); idx != -1 {
			tagName = tag[:idx]
		}
		if f.Anonymous && tagName == "" {
			ft := f.Type
			for ft.Kind() == reflect.Ptr {
				ft = ft.Elem()
			}
			if ft.Kind() == reflect.Struct {
				for k, v := range collectStructFields(ft) {
					fields[k] = v
				}
				continue
			}
		}
		if tagName == "" {
			tagName = f.Name
		}
		fields[tagName] = structFieldMeta{
			jsonName:  tagName,
			fieldName: f.Name,
			fieldType: f.Type,
		}
	}
	return fields
}

func normalizeStructMap(objMap map[string]any, structType reflect.Type, path string, report *ConversionReport) (map[string]any, error) {
	fields := collectStructFields(structType)

	// Unwrap single-key legacy model envelopes when target struct is the inner payload.
	if path == "" {
		if unwrapped, wrapperKey, ok := tryUnwrapLegacyEnvelope(objMap, fields); ok {
			report.MigratedFields = append(report.MigratedFields, FieldMigration{
				Path: wrapperKey,
				From: wrapperKey,
				To:   "<root>",
			})
			objMap = unwrapped
		}
	}

	consumedKeys := make(map[string]bool, len(objMap))
	ciLookup := make(map[string]string, len(objMap))
	for k := range objMap {
		ciLookup[strings.ToLower(k)] = k
	}

	result := make(map[string]any, len(fields))

	for canonicalKey, meta := range fields {
		fieldPath := joinPath(path, canonicalKey)
		var rawFieldVal any
		found := false

		// 1. Check exact canonical key match.
		if val, exists := objMap[canonicalKey]; exists {
			rawFieldVal = val
			consumedKeys[canonicalKey] = true
			found = true
		}

		// 2. Check registered legacy field aliases.
		if !found {
			if aliases, hasAliases := DefaultLegacyFieldAliases[canonicalKey]; hasAliases {
				for _, alias := range aliases {
					if val, exists := objMap[alias]; exists {
						rawFieldVal = val
						consumedKeys[alias] = true
						found = true
						report.MigratedFields = append(report.MigratedFields, FieldMigration{
							Path: joinPath(path, alias),
							From: alias,
							To:   canonicalKey,
						})
						break
					}
				}
			}
		}

		// 3. Check case-insensitive match on canonical key or struct field name.
		if !found {
			if actualKey, exists := ciLookup[strings.ToLower(canonicalKey)]; exists && !consumedKeys[actualKey] {
				rawFieldVal = objMap[actualKey]
				consumedKeys[actualKey] = true
				found = true
				if actualKey != canonicalKey {
					report.MigratedFields = append(report.MigratedFields, FieldMigration{
						Path: joinPath(path, actualKey),
						From: actualKey,
						To:   canonicalKey,
					})
				}
			}
		}

		if !found {
			continue
		}

		normVal, err := normalizeValue(rawFieldVal, meta.fieldType, fieldPath, report)
		if err != nil {
			return nil, err
		}
		result[canonicalKey] = normVal
	}

	// Record any remaining unconsumed keys as DroppedFields instead of failing.
	for k := range objMap {
		if !consumedKeys[k] {
			report.DroppedFields = append(report.DroppedFields, joinPath(path, k))
		}
	}

	return result, nil
}

func tryUnwrapLegacyEnvelope(objMap map[string]any, targetFields map[string]structFieldMeta) (map[string]any, string, bool) {
	for _, wrapperKey := range DefaultLegacyWrapperKeys {
		// If the target struct itself expects this wrapper key, do not unwrap.
		if _, expectsWrapper := targetFields[wrapperKey]; expectsWrapper {
			continue
		}
		if inner, exists := objMap[wrapperKey]; exists {
			if innerMap, ok := inner.(map[string]any); ok {
				return innerMap, wrapperKey, true
			}
		}
	}
	return nil, "", false
}

func normalizeMap(inMap map[string]any, mapType reflect.Type, path string, report *ConversionReport) (map[string]any, error) {
	elemType := mapType.Elem()
	outMap := make(map[string]any, len(inMap))
	isLabelMap := strings.HasSuffix(strings.ToLower(path), "label") || strings.HasSuffix(strings.ToLower(path), "labels")

	for k, v := range inMap {
		targetKey := k
		if isLabelMap {
			if mappedKey, hasAlias := DefaultLegacyLabelKeyAliases[k]; hasAlias {
				// Migrate legacy label key if canonical key is not already present.
				if _, exists := inMap[mappedKey]; !exists {
					targetKey = mappedKey
				}
				report.MigratedFields = append(report.MigratedFields, FieldMigration{
					Path: joinPath(path, k),
					From: k,
					To:   mappedKey,
				})
			}
		}
		normVal, err := normalizeValue(v, elemType, joinPath(path, targetKey), report)
		if err != nil {
			return nil, err
		}
		outMap[targetKey] = normVal
	}
	return outMap, nil
}

func normalizeScalarString(rawVal any, path string) (string, error) {
	switch v := rawVal.(type) {
	case string:
		return v, nil
	case json.Number:
		return v.String(), nil
	default:
		return "", fmt.Errorf("type mismatch at '%s': expected string, got %T", pathOrRoot(path), rawVal)
	}
}

func normalizeScalarInt(rawVal any, path string, report *ConversionReport) (int64, error) {
	switch v := rawVal.(type) {
	case json.Number:
		if i, err := v.Int64(); err == nil {
			return i, nil
		}
		if f, err := v.Float64(); err == nil {
			return int64(f), nil
		}
		return 0, fmt.Errorf("invalid integer number at '%s': %s", pathOrRoot(path), v.String())
	case float64:
		return int64(v), nil
	case string:
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			return 0, nil
		}
		i, err := strconv.ParseInt(trimmed, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("type mismatch at '%s': expected integer, got string %q", pathOrRoot(path), v)
		}
		report.MigratedFields = append(report.MigratedFields, FieldMigration{
			Path: pathOrRoot(path),
			From: "string",
			To:   "int",
		})
		return i, nil
	default:
		return 0, fmt.Errorf("type mismatch at '%s': expected integer, got %T", pathOrRoot(path), rawVal)
	}
}

func normalizeScalarUint(rawVal any, path string, report *ConversionReport) (uint64, error) {
	switch v := rawVal.(type) {
	case json.Number:
		if u, err := strconv.ParseUint(v.String(), 10, 64); err == nil {
			return u, nil
		}
		if f, err := v.Float64(); err == nil && f >= 0 {
			return uint64(f), nil
		}
		return 0, fmt.Errorf("invalid unsigned integer at '%s': %s", pathOrRoot(path), v.String())
	case float64:
		if v < 0 {
			return 0, fmt.Errorf("negative number for unsigned field at '%s': %v", pathOrRoot(path), v)
		}
		return uint64(v), nil
	case string:
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			return 0, nil
		}
		u, err := strconv.ParseUint(trimmed, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("type mismatch at '%s': expected uint, got string %q", pathOrRoot(path), v)
		}
		report.MigratedFields = append(report.MigratedFields, FieldMigration{
			Path: pathOrRoot(path),
			From: "string",
			To:   "uint",
		})
		return u, nil
	default:
		return 0, fmt.Errorf("type mismatch at '%s': expected uint, got %T", pathOrRoot(path), rawVal)
	}
}

func normalizeScalarFloat(rawVal any, path string, report *ConversionReport) (float64, error) {
	switch v := rawVal.(type) {
	case json.Number:
		f, err := v.Float64()
		if err != nil {
			return 0, fmt.Errorf("invalid float number at '%s': %s", pathOrRoot(path), v.String())
		}
		return f, nil
	case float64:
		return v, nil
	case string:
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			return 0, nil
		}
		f, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			return 0, fmt.Errorf("type mismatch at '%s': expected float, got string %q", pathOrRoot(path), v)
		}
		report.MigratedFields = append(report.MigratedFields, FieldMigration{
			Path: pathOrRoot(path),
			From: "string",
			To:   "float",
		})
		return f, nil
	default:
		return 0, fmt.Errorf("type mismatch at '%s': expected float, got %T", pathOrRoot(path), rawVal)
	}
}

func normalizeScalarBool(rawVal any, path string, report *ConversionReport) (bool, error) {
	switch v := rawVal.(type) {
	case bool:
		return v, nil
	case string:
		b, err := strconv.ParseBool(strings.TrimSpace(v))
		if err != nil {
			return false, fmt.Errorf("type mismatch at '%s': expected bool, got string %q", pathOrRoot(path), v)
		}
		report.MigratedFields = append(report.MigratedFields, FieldMigration{
			Path: pathOrRoot(path),
			From: "string",
			To:   "bool",
		})
		return b, nil
	default:
		return false, fmt.Errorf("type mismatch at '%s': expected bool, got %T", pathOrRoot(path), rawVal)
	}
}

func joinPath(parent, child string) string {
	if parent == "" {
		return child
	}
	return parent + "." + child
}

func pathOrRoot(path string) string {
	if path == "" {
		return "<root>"
	}
	return path
}

