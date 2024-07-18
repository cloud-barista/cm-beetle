package etcd

import (
	"context"
	"fmt"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

// KeyValueStore defines the interface for EtcdStore operations.
// This interface ensures that the EtcdStore struct implements all necessary methods for interacting with etcd.
type KeyValueStore interface {
	CreateSession(ctx context.Context) (*concurrency.Session, error)
	CloseSession(session *concurrency.Session) error
	Lock(ctx context.Context, session *concurrency.Session, lockKey string) (*concurrency.Mutex, error)
	Unlock(ctx context.Context, mutex *concurrency.Mutex) error
	Put(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	GetList(ctx context.Context, prefix string) ([]string, error)
	Delete(ctx context.Context, key string) error
	Close() error
}

// EtcdStore represents an etcd client.
type EtcdStore struct {
	cli *clientv3.Client
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
func NewEtcdStore(config Config) (KeyValueStore, error) {
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

		instance = &EtcdStore{
			cli: cli,
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

// CreateSession creates a new etcd session.
// A session is needed for acquiring locks.
func (s *EtcdStore) CreateSession(ctx context.Context) (*concurrency.Session, error) {
	return concurrency.NewSession(s.cli)
}

// CloseSession closes the given etcd session.
// It's important to close sessions to release resources.
func (s *EtcdStore) CloseSession(session *concurrency.Session) error {
	return session.Close()
}

// Lock acquires a lock on the given key and returns the mutex.
// It uses the provided session to ensure the lock's lifecycle is tied to the session.
func (s *EtcdStore) Lock(ctx context.Context, session *concurrency.Session, lockKey string) (*concurrency.Mutex, error) {
	mutex := concurrency.NewMutex(session, lockKey)
	err := mutex.Lock(ctx)
	if err != nil {
		return nil, err
	}
	return mutex, nil
}

// Unlock releases the given lock.
// It is important to release the lock to allow other clients to acquire it.
func (s *EtcdStore) Unlock(ctx context.Context, mutex *concurrency.Mutex) error {
	return mutex.Unlock(ctx)
}

// Put stores a key-value pair in etcd.
// It uses the etcd client to perform a Put operation.
func (s *EtcdStore) Put(ctx context.Context, key, value string) error {
	_, err := s.cli.Put(ctx, key, value)
	if err != nil {
		return fmt.Errorf("failed to put key-value: %w", err)
	}
	return nil
}

// Get retrieves the value for a given key from etcd.
// If the key is not found, it returns an error.
func (s *EtcdStore) Get(ctx context.Context, key string) (string, error) {
	resp, err := s.cli.Get(ctx, key)
	if err != nil {
		return "", fmt.Errorf("failed to get key: %w", err)
	}
	if len(resp.Kvs) == 0 {
		return "", fmt.Errorf("key not found: %s", key)
	}
	return string(resp.Kvs[0].Value), nil
}

// GetList retrieves multiple values for keys with the given prefix from etcd.
// It returns a slice of values corresponding to the keys with the given prefix.
func (s *EtcdStore) GetList(ctx context.Context, keyPrefix string) ([]string, error) {
	resp, err := s.cli.Get(ctx, keyPrefix, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("failed to get list with prefix: %w", err)
	}

	values := make([]string, len(resp.Kvs))
	for i, kv := range resp.Kvs {
		values[i] = string(kv.Value)
	}
	return values, nil
}

// Delete removes a key-value pair from etcd.
// It uses the etcd client to perform a Delete operation.
func (s *EtcdStore) Delete(ctx context.Context, key string) error {
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
