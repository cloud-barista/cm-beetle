package main

import (
	"encoding/json"
	"fmt"

	"github.com/cloud-barista/cm-beetle/pkg/lkvstore"
)

type ExampleStruct struct {
	Name string
	Age  int
}

func main() {
	// Initialize the key-value store with the specified file path
	config := lkvstore.Config{
		DbFilePath: "./lkvstore.db",
	}
	lkvstore.Init(config)

	// Example usage with various keys
	lkvstore.Put("prefix1_key1", string(123))
	lkvstore.Put("prefix1_key2", "Hello, world!")
	example := ExampleStruct{Name: "John Doe", Age: 30}
	bytes, err := json.Marshal(example)
	if err != nil {
		fmt.Printf("Error marshalling: %v\n", err)
	}

	lkvstore.Put("prefix2_key1", string(bytes))
	lkvstore.Put("prefix1_key3", "Another string")

	// Save the current state of the key-value store to file
	if err := lkvstore.SaveLkvStore(); err != nil {
		fmt.Printf("Error saving: %v\n", err)
	} else {
		fmt.Println("Successfully saved the lkvstore to file.")
	}

	// Clear the in-memory store
	lkvstore.Delete("prefix1_key1")
	lkvstore.Delete("prefix1_key2")
	lkvstore.Delete("prefix1_key3")
	lkvstore.Delete("prefix2_key1")

	// Load the state from the file back into the key-value store
	if err := lkvstore.LoadLkvStore(); err != nil {
		fmt.Printf("Error loading: %v\n", err)
	} else {
		fmt.Println("Successfully loaded the lkvstore from file.")
	}

	// Verify loaded data with prefix
	values, exists := lkvstore.GetWithPrefix("prefix1_")
	if exists {
		fmt.Println("Values with prefix 'prefix1_':")
		for _, value := range values {
			fmt.Printf("%v\n", value)
		}
	} else {
		fmt.Println("No values found with prefix 'prefix1_'")
	}

	// Verify loaded data without prefix
	value, exists := lkvstore.Get("prefix2_key1")
	if exists {
		fmt.Printf("Loaded value for 'prefix2_key1': %v\n", value)
	} else {
		fmt.Println("No value found for 'prefix2_key1'")
	}
}
