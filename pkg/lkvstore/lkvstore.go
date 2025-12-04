// Local Key-Value Store based on sync.Map and file I/O
package lkvstore

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

var (
	lkvstore   sync.Map
	dbFilePath string
)

type Config struct {
	DbFilePath string
}

type KeyValue struct {
	Key   string
	Value any
}

func Init(config Config) {
	if config.DbFilePath != "" {
		dbFilePath = config.DbFilePath
	} else {
		dbFilePath = ".lkvstore/lkvstore.db"
	}
}

// Save lkvstore to file
func SaveLkvStore() error {
	if dbFilePath == "" {
		return fmt.Errorf("db file path is not set")
	}

	// Ensure the DB file directory exists before creating the log file
	dir := filepath.Dir(dbFilePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Create the directory if it does not exist
		err = os.MkdirAll(dir, 0755) // Set permissions as needed
		if err != nil {
			log.Error().Msgf("Failed to create the DB directory: [%v]", err)
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	file, err := os.Create(dbFilePath)
	if err != nil {
		return fmt.Errorf("failed to create db file: %w", err)
	}
	defer file.Close()

	tempMap := make(map[string]any)
	lkvstore.Range(func(key, value any) bool {
		tempMap[key.(string)] = value
		return true
	})

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tempMap); err != nil {
		return fmt.Errorf("failed to encode map: %w", err)
	}

	return nil
}

// Load the info from file
func LoadLkvStore() error {
	if dbFilePath == "" {
		return fmt.Errorf("db file path is not set")
	}

	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		return fmt.Errorf("db file does not exist: %w", err)
	}

	file, err := os.Open(dbFilePath)
	if err != nil {
		return fmt.Errorf("failed to open db file: %w", err)
	}
	defer file.Close()

	var tempMap map[string]any
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tempMap); err != nil {
		return fmt.Errorf("failed to decode map: %w", err)
	}

	for key, value := range tempMap {
		lkvstore.Store(key, value)
	}

	return nil
}

// Get returns the value for a given key.
func Get(key string) (any, bool) {
	value, ok := lkvstore.Load(key)
	if !ok {
		return nil, false
	}

	var result any

	switch v := value.(type) {
	case string:
		// Unmarshal JSON if value is string
		if err := json.Unmarshal([]byte(v), &result); err != nil {
			return nil, false // Stop when failing to unmarshal
		}
	case map[string]any:
		// already map type
		result = v
	default:
		// unknown type
		result = v
	}

	return result, true
}

// GetWithPrefix returns the values for a given key prefix.
func GetWithPrefix(keyPrefix string) ([]any, bool) {
	var results []any
	var exists bool

	lkvstore.Range(func(key, value any) bool {
		if strings.HasPrefix(key.(string), keyPrefix) {
			var result any

			switch v := value.(type) {
			case string:
				// Unmarshal JSON if value is string
				if err := json.Unmarshal([]byte(v), &result); err != nil {
					return false // Stop when failing to unmarshal
				}
			case map[string]any:
				// already map type
				result = v
			default:
				// unknown type
				result = v
			}

			results = append(results, result)
			exists = true
		}
		return true
	})

	if !exists {
		return nil, false
	}

	return results, true
}

// GetKv returns the key-value pair for a given key.
func GetKv(key string) (KeyValue, bool) {
	value, ok := Get(key)
	if !ok {
		return KeyValue{}, false
	}
	return KeyValue{Key: key, Value: value}, true
}

// GetKvWithPrefix returns the key-value pairs for a given key prefix.
func GetKvWithPrefix(keyPrefix string) ([]KeyValue, bool) {
	var results []KeyValue
	var exists bool

	lkvstore.Range(func(key, value any) bool {
		if strings.HasPrefix(key.(string), keyPrefix) {
			var result any

			switch v := value.(type) {
			case string:
				// Unmarshal JSON if value is string
				if err := json.Unmarshal([]byte(v), &result); err != nil {
					return false // Stop when failing to unmarshal
				}
			case map[string]any:
				// already map type
				result = v
			default:
				// unknown type
				result = v
			}

			results = append(results, KeyValue{Key: key.(string), Value: result})
			exists = true
		}
		return true
	})

	if !exists {
		return nil, false
	}

	return results, true
}

// Put stores the key-value pair.
func Put(key string, value any) error {
	// Marshal the value to JSON
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	// Store the JSON string
	lkvstore.Store(key, string(jsonValue))
	return nil
}

// Delete removes the key-value pair for a given key.
func Delete(key string) {
	lkvstore.Delete(key)
}
