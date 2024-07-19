package etcd

import (
	"context"
	"fmt"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

// type KeyValue struct {
// 	Key   string
// 	Value string
// }
//
// type Store interface {
//     InitDB() error
//     InitData() error
//     Put(key string, value string) error
//     Get(key string) (*KeyValue, error)
//     GetList(key string, sortAscend bool) ([]*KeyValue, error)
//     Delete(key string) error
//     Close() error

//     // --- management, nutsdb only
//     Merge() error
// }

// KeyValueStore defines the interface for EtcdStore operations.
// This interface ensures that the EtcdStore struct implements all necessary methods for interacting with etcd.
type KeyValueStore interface {
	NewSession(ctx context.Context) (*concurrency.Session, error)
	NewLock(ctx context.Context, session *concurrency.Session, lockKey string) (*concurrency.Mutex, error)
	Put(key, value string) error
	PutWith(ctx context.Context, key, value string) error
	Get(key string) (string, error)
	GetWith(ctx context.Context, key string) (string, error)
	GetList(keyPrefix string) ([]string, error)
	GetListWith(ctx context.Context, keyPrefix string) ([]string, error)
	Delete(key string) error
	DeleteWith(ctx context.Context, key string) error
	Close() error
	// CloseSession(session *concurrency.Session) error
	// Unlock(ctx context.Context, mutex *concurrency.Mutex) error
	// GetSortedList(keyPrefix string, sortBy clientv3.SortTarget, order clientv3.SortOrder) ([]string, error)
	// GetSortedListWith(ctx context.Context, keyPrefix string, sortBy clientv3.SortTarget, order clientv3.SortOrder) ([]string, error)
}

// EtcdStore represents an etcd client.
type EtcdStore struct {
	cli *clientv3.Client
	ctx context.Context
}

// Config holds the configuration for EtcdStore.
type Config struct {
	Endpoints   []string
	DialTimeout time.Duration
}

var (
	instance *EtcdStore
	once     sync.Once
)

// NewEtcdStore creates a new instance of EtcdStore (singleton).
// It initializes the etcd client with the provided configuration and ensures only one instance is created.
func NewEtcdStore(ctx context.Context, config Config) (KeyValueStore, error) {
	var err error
	once.Do(func() {
		cli, cliErr := clientv3.New(clientv3.Config{
			Endpoints:   config.Endpoints,
			DialTimeout: config.DialTimeout,
		})
		if cliErr != nil {
			err = fmt.Errorf("failed to create etcd client: %w", cliErr)
			return
		}

		if ctx == nil {
			ctx = context.Background()
		}

		instance = &EtcdStore{
			cli: cli,
			ctx: ctx,
		}
	})

	if err != nil {
		return nil, err
	}
	return instance, nil
}

// GetEtcdStore returns the singleton instance of EtcdStore.
func GetEtcdStore() *EtcdStore {
	return instance
}

// OpenSession creates a new etcd session.
// A session is needed for acquiring locks.
func (s *EtcdStore) NewSession(ctx context.Context) (*concurrency.Session, error) {
	return concurrency.NewSession(s.cli)
}

// NewLock acquires a lock on the given key and returns the mutex.
// It uses the provided session to ensure the lock's lifecycle is tied to the session.
func (s *EtcdStore) NewLock(ctx context.Context, session *concurrency.Session, lockKey string) (*concurrency.Mutex, error) {
	mutex := concurrency.NewMutex(session, lockKey)
	err := mutex.Lock(ctx)
	if err != nil {
		return nil, err
	}
	return mutex, nil
}

// Put stores a key-value pair in etcd.
func (s *EtcdStore) Put(key, value string) error {
	return s.PutWith(s.ctx, key, value)
}

// PutWith stores a key-value pair in etcd using the provided context.
func (s *EtcdStore) PutWith(ctx context.Context, key, value string) error {
	_, err := s.cli.Put(ctx, key, value)
	if err != nil {
		return fmt.Errorf("failed to put key-value: %w", err)
	}
	return nil
}

// Get retrieves the value for a given key from etcd without using a context.
func (s *EtcdStore) Get(key string) (string, error) {
	return s.GetWith(s.ctx, key)
}

// GetWith retrieves the value for a given key from etcd using the provided context.
func (s *EtcdStore) GetWith(ctx context.Context, key string) (string, error) {
	resp, err := s.cli.Get(ctx, key)
	if err != nil {
		return "", fmt.Errorf("failed to get key: %w", err)
	}
	if len(resp.Kvs) == 0 {
		return "", fmt.Errorf("key not found: %s", key)
	}
	return string(resp.Kvs[0].Value), nil
}

// GetListWith retrieves multiple values for keys with the given keyPrefix from etcd.
func (s *EtcdStore) GetList(keyPrefix string) ([]string, error) {
	return s.GetListWith(s.ctx, keyPrefix)
}

// GetListWith retrieves multiple values for keys with the given keyPrefix from etcd using the provided context.
func (s *EtcdStore) GetListWith(ctx context.Context, keyPrefix string) ([]string, error) {
	// ascending by key as a default sort order
	optAscendByKey := clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend)

	// Get all values with the given keyPrefix
	resp, err := s.cli.Get(ctx, keyPrefix, clientv3.WithPrefix(), optAscendByKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get list with keyPrefix: %w", err)
	}

	values := make([]string, len(resp.Kvs))
	for i, kv := range resp.Kvs {
		values[i] = string(kv.Value)
	}
	return values, nil
}

// GetSortedList retrieves multiple values for keys with the given keyPrefix, sortBy, and order from etcd.
func (s *EtcdStore) GetSortedList(keyPrefix string, sortBy clientv3.SortTarget, order clientv3.SortOrder) ([]string, error) {
	return s.GetSortedListWith(s.ctx, keyPrefix, sortBy, order)
}

// GetSortedListWith retrieves multiple values for keys with  the given keyPrefix, sortBy, and order from etcd using the provided context.
func (s *EtcdStore) GetSortedListWith(ctx context.Context, keyPrefix string, sortBy clientv3.SortTarget, order clientv3.SortOrder) ([]string, error) {
	sortOp := clientv3.WithSort(sortBy, order)
	resp, err := s.cli.Get(ctx, keyPrefix, clientv3.WithPrefix(), sortOp)
	if err != nil {
		return nil, fmt.Errorf("failed to get list with keyPrefix: %w", err)
	}

	values := make([]string, len(resp.Kvs))
	for i, kv := range resp.Kvs {
		values[i] = string(kv.Value)
	}
	return values, nil
}

// Delete removes a key-value pair from etcd without using a context.
func (s *EtcdStore) Delete(key string) error {
	return s.DeleteWith(s.ctx, key)
}

// DeleteWith removes a key-value pair from etcd using the provided context.
func (s *EtcdStore) DeleteWith(ctx context.Context, key string) error {
	_, err := s.cli.Delete(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to delete key: %w", err)
	}
	return nil
}

// Close closes the etcd client.
// This is necessary to release resources associated with the client.
func (s *EtcdStore) Close() error {
	return s.cli.Close()
}

// // CloseSession closes the given etcd session.
// // It's important to close sessions to release resources.
// func (s *EtcdStore) CloseSession(session *concurrency.Session) error {
// 	return session.Close()
// }

// // Unlock releases the given lock.
// // It is important to release the lock to allow other clients to acquire it.
// func (s *EtcdStore) Unlock(ctx context.Context, mutex *concurrency.Mutex) error {
// 	return mutex.Unlock(ctx)
// }
