package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/etcd"
)

func main() {
	// EtcdStore configuration
	config := etcd.Config{
		Endpoints:   []string{"localhost:2379"}, // Replace with your etcd server endpoints
		DialTimeout: 5 * time.Second,
	}

	// Create EtcdStore instance (singleton)
	store, err := etcd.NewEtcdStore(config)
	if err != nil {
		log.Fatalf("Failed to create EtcdStore: %v", err)
	}
	defer store.Close()

	ctx := context.Background() // Create context for etcd operations

	// Basic CRUD operations test
	basicCRUDTest(ctx, store)

	// Race condition test
	raceConditionTest(ctx, store)

	fmt.Println("All operations completed successfully!")
}

func basicCRUDTest(ctx context.Context, store etcd.KeyValueStore) {
	key := "test_key"
	value := "Hello, Etcd!"

	// Put (Store) a key-value pair
	err := store.Put(ctx, key, value)
	if err != nil {
		log.Fatalf("Failed to put key-value: %v", err)
	}
	fmt.Printf("Successfully put key '%s' with value '%s'\n", key, value)

	// Get (Retrieve) the value
	retrievedValue, err := store.Get(ctx, key)
	if err != nil {
		log.Fatalf("Failed to get value: %v", err)
	}
	fmt.Printf("Retrieved value for key '%s': %s\n", key, retrievedValue)

	// Update the value
	newValue := "Updated Etcd Value"
	err = store.Put(ctx, key, newValue)
	if err != nil {
		log.Fatalf("Failed to update value: %v", err)
	}
	fmt.Printf("Successfully updated key '%s' with new value '%s'\n", key, newValue)

	// Get (Retrieve) the updated value
	retrievedValue, err = store.Get(ctx, key)
	if err != nil {
		log.Fatalf("Failed to get updated value: %v", err)
	}
	fmt.Printf("Retrieved updated value for key '%s': %s\n", key, retrievedValue)

	// Delete the key-value pair
	err = store.Delete(ctx, key)
	if err != nil {
		log.Fatalf("Failed to delete key: %v", err)
	}
	fmt.Printf("Successfully deleted key '%s'\n", key)

	// Verify deletion
	_, err = store.Get(ctx, key)
	if err != nil {
		fmt.Printf("As expected, failed to get deleted key '%s': %v\n", key, err)
	} else {
		log.Fatalf("Unexpectedly succeeded in getting deleted key '%s'", key)
	}
}

func raceConditionTest(ctx context.Context, store etcd.KeyValueStore) {
	fmt.Println("Starting race condition test...")

	key := "race_test_key"
	iterations := 100
	goroutines := 5

	// Initialize the key with 0
	err := store.Put(ctx, key, "0")
	if err != nil {
		log.Fatalf("Failed to initialize key: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(goroutines)

	// Start goroutines
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()

			// Create a persistent session
			session, err := store.CreateSession(ctx)
			if err != nil {
				log.Fatalf("Failed to create etcd session: %v", err)
			}
			defer store.CloseSession(session)

			for j := 0; j < iterations; j++ {

				// Lock
				lockKey := key
				mutex, err := store.Lock(ctx, session, lockKey)
				if err != nil {
					log.Printf("Failed to acquire lock: %v", err)
					continue
				}

				// Get current value, increment, and put new value within the lock
				value, err := store.Get(ctx, key)
				if err != nil {
					log.Printf("Failed to get value: %v", err)
					store.Unlock(ctx, mutex)
					continue
				}

				intValue, _ := strconv.Atoi(value)
				newValue := fmt.Sprintf("%d", intValue+1)

				err = store.Put(ctx, key, newValue)
				if err != nil {
					log.Printf("Failed to put value: %v", err)
					store.Unlock(ctx, mutex)
					continue
				}
				log.Printf("Put value: %s", newValue)

				// Unlock
				err = store.Unlock(ctx, mutex)
				if err != nil {
					log.Printf("Failed to release lock: %v", err)
					continue
				}
			}
		}()
	}

	wg.Wait()

	// Verify the final value
	finalValue, err := store.Get(ctx, key)
	if err != nil {
		log.Fatalf("Failed to get final value: %v", err)
	}

	expectedValue := goroutines * iterations
	actualValue, _ := strconv.Atoi(finalValue)
	if actualValue != expectedValue {
		log.Fatalf("Race condition detected. Expected %d, but got %d", expectedValue, actualValue)
	}

	fmt.Printf("Race condition test finished. Final value: %s\n", finalValue)

	// Clean up
	store.Delete(ctx, key)
}
