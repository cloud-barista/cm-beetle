package fieldsec

import (
	"strings"
)

// getFieldByPath retrieves a value from a nested map using a dot-separated path.
// Example: "source.filesystem.ssh.privateKey"
func getFieldByPath(m map[string]any, path string) (string, bool) {
	parts := strings.Split(path, ".")
	current := any(m)

	for _, part := range parts {
		cm, ok := current.(map[string]any)
		if !ok {
			return "", false
		}

		val, exists := cm[part]
		if !exists {
			return "", false
		}

		current = val
	}

	// Final value should be a string
	str, ok := current.(string)
	if !ok {
		return "", false
	}

	return str, true
}

// setFieldByPath sets a value in a nested map using a dot-separated path.
// Creates intermediate maps if they don't exist.
// Example: "source.filesystem.ssh.privateKey"
func setFieldByPath(m map[string]any, path string, value string) bool {
	parts := strings.Split(path, ".")
	current := m

	// Navigate to the parent of the target field
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]

		val, exists := current[part]
		if !exists {
			return false
		}

		cm, ok := val.(map[string]any)
		if !ok {
			return false
		}

		current = cm
	}

	// Set the final field
	lastPart := parts[len(parts)-1]
	if _, exists := current[lastPart]; !exists {
		return false
	}

	current[lastPart] = value
	return true
}
