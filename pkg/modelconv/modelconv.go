package modelconv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// ConvertWithUpcast converts source value T into target struct U with reflect-based legacy field migration.
func ConvertWithUpcast[T any, U any](source T) (U, *ConversionReport, error) {
	var target U
	if isNil(source) {
		return target, nil, nil
	}
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return target, nil, fmt.Errorf("failed to marshal source: %w", err)
	}
	return UpcastTo[U](sourceJSON)
}

// ConvertWithValidation combines type validation and conversion
func ConvertWithValidation[T any, U any](source T) (U, error) {
	var target U

	// Validate compatibility first
	err := ValidateTypeCompatibility[T, U]()
	if err != nil {
		return target, fmt.Errorf("type compatibility validation failed: %w", err)
	}

	// Then convert
	return ConvertAny[T, U](source)
}

// ConvertAny converts any type T to any type U using JSON marshaling/unmarshaling
func ConvertAny[T any, U any](source T) (U, error) {
	var target U

	// Handle nil pointers
	if isNil(source) {
		return target, nil
	}

	// Marshal source to JSON
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return target, fmt.Errorf("failed to marshal source: %w", err)
	}

	// Unmarshal JSON to target
	err = json.Unmarshal(sourceJSON, &target)
	if err != nil {
		return target, fmt.Errorf("failed to unmarshal to target: %w", err)
	}

	return target, nil
}

// ValidateTypeCompatibility checks if two types are compatible for conversion
func ValidateTypeCompatibility[T any, U any]() error {
	var sourceType T
	var targetType U

	sourceReflect := reflect.TypeOf(sourceType)
	targetReflect := reflect.TypeOf(targetType)

	return validateReflectTypes(sourceReflect, targetReflect)
}

// validateReflectTypes performs deep type compatibility validation
func validateReflectTypes(sourceType, targetType reflect.Type) error {
	// Handle nil types
	if sourceType == nil || targetType == nil {
		return nil
	}

	// Handle pointers
	if sourceType.Kind() == reflect.Ptr {
		sourceType = sourceType.Elem()
	}
	if targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}

	// Check basic type compatibility
	if sourceType.Kind() != targetType.Kind() {
		return fmt.Errorf("incompatible kinds: source=%v, target=%v", sourceType.Kind(), targetType.Kind())
	}

	switch sourceType.Kind() {
	case reflect.Struct:
		return validateStructTypes(sourceType, targetType)
	case reflect.Slice:
		return validateSliceTypes(sourceType, targetType)
	case reflect.Map:
		return validateMapTypes(sourceType, targetType)
	case reflect.String:
		// For string-based custom types, check if they have the same underlying type
		if sourceType.Kind() == reflect.String && targetType.Kind() == reflect.String {
			return nil // Allow conversion between string-based custom types
		}
		fallthrough
	default:
		// For basic types, check underlying type compatibility
		if sourceType.Kind() == targetType.Kind() {
			// Same underlying kind, allow conversion
			return nil
		}
		if sourceType != targetType {
			return fmt.Errorf("incompatible basic types: source=%v, target=%v", sourceType, targetType)
		}
	}

	return nil
}

// validateStructTypes validates struct field compatibility
func validateStructTypes(sourceType, targetType reflect.Type) error {
	sourceFields := getFieldMap(sourceType)
	targetFields := getFieldMap(targetType)

	var missingFields []string

	for fieldName, targetField := range targetFields {
		if sourceField, exists := sourceFields[fieldName]; !exists {
			missingFields = append(missingFields, fieldName)
		} else {
			// Recursively check field type compatibility
			if err := validateReflectTypes(sourceField.Type, targetField.Type); err != nil {
				return fmt.Errorf("field '%s': %w", fieldName, err)
			}
		}
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("missing fields in source struct: %v", missingFields)
	}

	return nil
}

// validateSliceTypes validates slice element type compatibility
func validateSliceTypes(sourceType, targetType reflect.Type) error {
	return validateReflectTypes(sourceType.Elem(), targetType.Elem())
}

// validateMapTypes validates map key and value type compatibility
func validateMapTypes(sourceType, targetType reflect.Type) error {
	// Check key types
	if err := validateReflectTypes(sourceType.Key(), targetType.Key()); err != nil {
		return fmt.Errorf("incompatible map key types: %w", err)
	}

	// Check value types
	return validateReflectTypes(sourceType.Elem(), targetType.Elem())
}

// getFieldMap creates a map of JSON field names to exported struct fields, flattening anonymous embedded structs.
func getFieldMap(structType reflect.Type) map[string]reflect.StructField {
	fieldMap := make(map[string]reflect.StructField)
	collectFieldMap(structType, fieldMap)
	return fieldMap
}

// collectFieldMap populates fieldMap matching encoding/json rules (outer fields shadow embedded ones).
func collectFieldMap(structType reflect.Type, fieldMap map[string]reflect.StructField) {
	var embedded []reflect.Type

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if !field.IsExported() && !field.Anonymous {
			continue
		}

		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" {
			continue
		}

		tagName, _, _ := strings.Cut(jsonTag, ",")
		if field.Anonymous && tagName == "" {
			embeddedType := field.Type
			if embeddedType.Kind() == reflect.Ptr {
				embeddedType = embeddedType.Elem()
			}
			if embeddedType.Kind() == reflect.Struct {
				embedded = append(embedded, embeddedType)
			}
			continue
		}

		key := field.Name
		if tagName != "" {
			key = tagName
		}
		if _, exists := fieldMap[key]; !exists {
			fieldMap[key] = field
		}
	}

	for _, embeddedType := range embedded {
		collectFieldMap(embeddedType, fieldMap)
	}
}

// isNil checks if a value is nil
func isNil(v interface{}) bool {
	if v == nil {
		return true
	}

	valueOf := reflect.ValueOf(v)
	switch valueOf.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return valueOf.IsNil()
	default:
		return false
	}
}
