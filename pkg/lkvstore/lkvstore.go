// Local Key-Value Store based on sync.Map and file I/O
package lkvstore

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
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
	Value string
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

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(dbFilePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(dbFilePath)
	if err != nil {
		return fmt.Errorf("failed to create db file: %w", err)
	}
	defer file.Close()

	tempMap := make(map[string]interface{})
	lkvstore.Range(func(key, value interface{}) bool {
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

	var tempMap map[string]interface{}
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
func Get(key string) (string, bool) {
	if value, ok := lkvstore.Load(key); ok {
		return value.(string), true
	}
	return "", false
}

// GetWithPrefix returns the values for a given key prefix.
func GetWithPrefix(keyPrefix string) ([]string, bool) {
	var results []string
	var exists bool
	lkvstore.Range(func(key, value interface{}) bool {
		if strings.HasPrefix(key.(string), keyPrefix) {
			results = append(results, value.(string))
			exists = true
		}
		return true
	})
	return results, exists
}

// GetKv returns the key-value pair for a given key.
func GetKv(key string) (KeyValue, bool) {
	var kv KeyValue
	var exists bool
	if value, ok := lkvstore.Load(key); ok {
		kv.Key = key
		kv.Value = value.(string)
		exists = true
	}
	return kv, exists
}

// GetKvWIthPrefix returns the key-value pairs for a given key prefix.
func GetKvWIthPrefix(keyPrefix string) ([]KeyValue, bool) {
	var results []KeyValue
	var exists bool
	lkvstore.Range(func(key, value interface{}) bool {
		if strings.HasPrefix(key.(string), keyPrefix) {
			kv := KeyValue{Key: key.(string), Value: value.(string)}
			results = append(results, kv)
			exists = true
		}
		return true
	})
	return results, exists
}

// Put the key-value pair.
func Put(key string, value string) {
	lkvstore.Store(key, value)
}

// Delete the key-value pair for a given key.
func Delete(key string) {
	lkvstore.Delete(key)
}
