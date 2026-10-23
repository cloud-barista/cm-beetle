package artifact

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

// ErrArtifactNotFound indicates that the requested artifact ID does not exist in the store.
var ErrArtifactNotFound = errors.New("artifact not found")

var validArtifactIdPattern = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9._-]{0,127}$`)

// Store manages thread-safe in-memory indexing and atomic JSON file persistence for artifacts.
type Store struct {
	mu      sync.RWMutex
	baseDir string
	items   map[string]RawArtifact
}

var (
	defaultStoreMu sync.RWMutex
	defaultStore   *Store
)

// InitStore initializes the package-level default artifact store at the given directory path.
func InitStore(dirPath string) (*Store, error) {
	if strings.TrimSpace(dirPath) == "" {
		dirPath = "./db/artifacts"
	}
	st, err := NewStore(dirPath)
	if err != nil {
		return nil, err
	}
	defaultStoreMu.Lock()
	defaultStore = st
	defaultStoreMu.Unlock()
	return st, nil
}

// GetDefaultStore returns the initialized default Store, or initializes one at ./db/artifacts if nil.
func GetDefaultStore() (*Store, error) {
	defaultStoreMu.RLock()
	st := defaultStore
	defaultStoreMu.RUnlock()
	if st != nil {
		return st, nil
	}
	return InitStore("./db/artifacts")
}

// NewStore creates a new Store backed by dirPath and loads existing JSON artifact files into memory.
func NewStore(dirPath string) (*Store, error) {
	absDir, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve artifact directory path: %w", err)
	}
	if err := os.MkdirAll(absDir, 0750); err != nil {
		return nil, fmt.Errorf("failed to create artifact directory %s: %w", absDir, err)
	}

	st := &Store{
		baseDir: absDir,
		items:   make(map[string]RawArtifact),
	}
	if err := st.loadFromDisk(); err != nil {
		return nil, err
	}
	return st, nil
}

func (s *Store) resolveSafeFilePath(id string) (string, error) {
	cleanId := strings.TrimSpace(id)
	if !validArtifactIdPattern.MatchString(cleanId) || strings.Contains(cleanId, "..") {
		return "", fmt.Errorf("invalid artifact id '%s': must contain only alphanumeric, '.', '_', or '-' characters", id)
	}
	fileName := filepath.Base(cleanId) + ".json"
	fullPath := filepath.Join(s.baseDir, fileName)
	absFull, err := filepath.Abs(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve artifact file path: %w", err)
	}
	if !strings.HasPrefix(absFull, s.baseDir+string(filepath.Separator)) {
		return "", fmt.Errorf("artifact file path escapes base directory")
	}
	return absFull, nil
}

func (s *Store) loadFromDisk() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	entries, err := os.ReadDir(s.baseDir)
	if err != nil {
		return fmt.Errorf("failed to read artifact directory %s: %w", s.baseDir, err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		filePath := filepath.Join(s.baseDir, entry.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Warn().Err(err).Str("file", filePath).Msg("Skipping unreadable artifact file")
			continue
		}
		art, err := parseStoredArtifactBytes(data, strings.TrimSuffix(entry.Name(), ".json"))
		if err != nil {
			log.Warn().Err(err).Str("file", filePath).Msg("Skipping malformed artifact file")
			continue
		}
		s.items[art.Id] = art
	}
	return nil
}

func parseStoredArtifactBytes(data []byte, fallbackId string) (RawArtifact, error) {
	var art RawArtifact
	if err := json.Unmarshal(data, &art); err != nil {
		return RawArtifact{}, err
	}

	// Support legacy envelope files dropped directly into the artifact directory.
	if len(art.Payload) == 0 {
		var legacy CreateReq
		if err := json.Unmarshal(data, &legacy); err == nil {
			_ = legacy.Normalize()
			if len(legacy.Payload) > 0 {
				art.Payload = legacy.Payload
				if art.Category == "" {
					art.Category = legacy.Category
				}
				if art.Stage == "" {
					art.Stage = legacy.Stage
				}
			}
		}
	}

	if art.Id == "" {
		art.Id = fallbackId
	}
	normalized, err := normalizeAndUpcastArtifact(art)
	if err != nil {
		return art, nil
	}
	return normalized, nil
}

// Save atomically persists a RawArtifact to disk using Write-Sync-Rename and updates the in-memory index.
func (s *Store) Save(art RawArtifact) error {
	targetPath, err := s.resolveSafeFilePath(art.Id)
	if err != nil {
		return err
	}

	encoded, err := json.MarshalIndent(art, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal artifact '%s': %w", art.Id, err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	tmpFile, err := os.CreateTemp(s.baseDir, art.Id+".*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp file for artifact '%s': %w", art.Id, err)
	}
	tmpPath := tmpFile.Name()
	cleanupTmp := func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpPath)
	}

	if _, err := tmpFile.Write(encoded); err != nil {
		cleanupTmp()
		return fmt.Errorf("failed to write temp file for artifact '%s': %w", art.Id, err)
	}
	if err := tmpFile.Sync(); err != nil {
		cleanupTmp()
		return fmt.Errorf("failed to fsync temp file for artifact '%s': %w", art.Id, err)
	}
	if err := tmpFile.Close(); err != nil {
		_ = os.Remove(tmpPath)
		return fmt.Errorf("failed to close temp file for artifact '%s': %w", art.Id, err)
	}
	if err := os.Rename(tmpPath, targetPath); err != nil {
		_ = os.Remove(tmpPath)
		return fmt.Errorf("failed to atomically rename artifact file '%s': %w", art.Id, err)
	}

	s.items[art.Id] = art
	return nil
}

// Get retrieves a RawArtifact by ID from the in-memory index.
func (s *Store) Get(id string) (RawArtifact, error) {
	s.mu.RLock()
	art, exists := s.items[id]
	s.mu.RUnlock()
	if !exists {
		return RawArtifact{}, fmt.Errorf("%w: %s", ErrArtifactNotFound, id)
	}
	return art, nil
}

// List returns all artifacts matching Filter, sorted newest first by CreateTime/UpdateTime.
func (s *Store) List(filter Filter) []RawArtifact {
	s.mu.RLock()
	result := make([]RawArtifact, 0, len(s.items))
	for _, item := range s.items {
		if filter.Category != "" && item.Category != filter.Category {
			continue
		}
		if filter.Stage != "" && item.Stage != filter.Stage {
			continue
		}
		if filter.SchemaType != "" && !strings.EqualFold(item.SchemaType, filter.SchemaType) {
			continue
		}
		if filter.SourceGroupId != "" && item.Lineage.SourceGroupId != filter.SourceGroupId {
			continue
		}
		result = append(result, item)
	}
	s.mu.RUnlock()

	sort.Slice(result, func(i, j int) bool {
		ti := result[i].UpdateTime
		if ti == "" {
			ti = result[i].CreateTime
		}
		tj := result[j].UpdateTime
		if tj == "" {
			tj = result[j].CreateTime
		}
		if ti == tj {
			return result[i].Id < result[j].Id
		}
		return ti > tj
	})
	return result
}

// Delete removes an artifact file from disk and deletes it from the in-memory index.
func (s *Store) Delete(id string) error {
	targetPath, err := s.resolveSafeFilePath(id)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[id]; !exists {
		return fmt.Errorf("%w: %s", ErrArtifactNotFound, id)
	}
	if err := os.Remove(targetPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to delete artifact file '%s': %w", id, err)
	}
	delete(s.items, id)
	return nil
}

