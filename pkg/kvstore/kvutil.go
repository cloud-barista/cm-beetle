package kvstore

import (
	"fmt"
	"strings"
)

// FilterKVsBy filters a KeyValue map based on the given prefix key.
// It returns a new KeyValue map containing only the key-value pairs that match the prefix criteria.
func FilterKVsBy(kvs KeyValueMap, prefixKey string) KeyValueMap {
	result := make(KeyValueMap)
	prefix := strings.TrimSuffix(prefixKey, "/")

	for key, value := range kvs {
		if strings.HasPrefix(key, prefix) && hasOnlyNextSegment(prefix, key) {
			result[key] = value
		}
	}

	return result
}

// hasOnlyNextSegment checks if a key has only one additional key segment after the prefix key.
func hasOnlyNextSegment(prefix, key string) bool {
	// Trim the prefix from the key and split the remaining part
	trimmed := strings.TrimPrefix(key, prefix)
	trimmed = strings.Trim(trimmed, "/")
	parts := strings.Split(trimmed, "/")

	// Check if the key has only one additional key segment
	return len(parts) == 1
}

// ExtractIDsFromKey extracts specific IDs from a given key based on the provided structure.
// It returns a slice of extracted ID values in the order of provided ID types.
// If any ID type is not found in the key, it returns an error.
func ExtractIDsFromKey(key string, idTypes ...string) ([]string, error) {
	parts := strings.Split(strings.Trim(key, "/"), "/")
	if len(parts) < len(idTypes)*2 {
		return nil, fmt.Errorf("key does not contain all requested ID types")
	}

	ids := make([]string, len(idTypes))
	for i, idType := range idTypes {
		index := indexOf(parts, idType)
		if index == -1 || index+1 >= len(parts) {
			return nil, fmt.Errorf("could not find ID for type: %s", idType)
		}
		ids[i] = parts[index+1]
	}
	return ids, nil
}

// ContainsIDs checks if a key contains specific ID values.
// It returns true if the key contains all ID types and values specified in the ids map.
func ContainsIDs(key string, ids map[string]string) bool {
	parts := strings.Split(strings.Trim(key, "/"), "/")
	for idType, idValue := range ids {
		index := indexOf(parts, idType)
		if index == -1 || index+1 >= len(parts) || parts[index+1] != idValue {
			return false
		}
	}
	return true
}

// BuildKey constructs a key from given ID types and values.
// It returns a string representing the constructed key.
func BuildKey(ids map[string]string) string {
	var parts []string
	for idType, idValue := range ids {
		parts = append(parts, idType, idValue)
	}
	return "/" + strings.Join(parts, "/")
}

// indexOf finds the index of a string in a slice.
// It returns the index if found, or -1 if not found.
func indexOf(slice []string, item string) int {
	for i, s := range slice {
		if s == item {
			return i
		}
	}
	return -1
}
