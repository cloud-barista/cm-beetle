package transx

import (
	"fmt"
)

// ============================================================================
// Pipeline Definition
// ============================================================================

// Pipeline represents a planned transfer with multiple steps.
type Pipeline struct {
	Name     string
	Strategy string
	Steps    []Step
}

// Step represents a single transfer step in the pipeline.
type Step struct {
	Name        string
	Source      DataLocation
	Destination DataLocation
	Executor    Executor
}

// Execute runs all steps in the pipeline sequentially.
func (p *Pipeline) Execute() error {
	for i, step := range p.Steps {
		if err := step.Executor.Execute(step.Source, step.Destination); err != nil {
			return fmt.Errorf("step %d (%s) failed: %w", i+1, step.Name, err)
		}
	}
	return nil
}

// ============================================================================
// Plan Function: Single Entry Point
// ============================================================================

// Plan analyzes the DataMigrationModel and returns the optimal transfer Pipeline.
// The routing is based on StorageType combinations (3 cases only).
func Plan(model DataMigrationModel) (*Pipeline, error) {
	if err := Validate(model); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	srcStorage := model.Source.StorageType
	dstStorage := model.Destination.StorageType

	switch {
	// Case 1: filesystem ↔ filesystem
	case srcStorage == StorageTypeFilesystem && dstStorage == StorageTypeFilesystem:
		return planFilesystemTransfer(model)

	// Case 2: objectstorage ↔ objectstorage
	case srcStorage == StorageTypeObjectStorage && dstStorage == StorageTypeObjectStorage:
		return planObjectStorageTransfer(model)

	// Case 3: Cross-storage (filesystem ↔ objectstorage)
	default:
		return planCrossStorageTransfer(model)
	}
}

// ============================================================================
// Filesystem Transfer (rsync-based)
// ============================================================================

// planFilesystemTransfer handles all filesystem ↔ filesystem transfers.
// Supported scenarios:
//   - local→ssh (Push): rsync from local to remote
//   - ssh→local (Pull): rsync from remote to local
//   - ssh→ssh with strategy "relay": Pull to local staging, then Push to destination
//   - ssh→ssh with strategy "agent-forward" or "auto": AgentForward mode
//
// Note: local-to-local is not supported (use standard file copy utilities).
func planFilesystemTransfer(model DataMigrationModel) (*Pipeline, error) {
	srcIsRemote := model.Source.Filesystem != nil && model.Source.Filesystem.SSH != nil
	dstIsRemote := model.Destination.Filesystem != nil && model.Destination.Filesystem.SSH != nil

	// SSH → SSH with relay strategy: use local staging
	if srcIsRemote && dstIsRemote && model.Strategy == StrategyRelay {
		return planRelayFilesystemTransfer(model)
	}

	// Default: direct transfer (Pull, Push, or AgentForward)
	rsyncExec, err := NewRsyncExecutor(model.Source, model.Destination)
	if err != nil {
		return nil, fmt.Errorf("failed to create rsync executor: %w", err)
	}

	return &Pipeline{
		Name:     PipelineFilesystemTransfer,
		Strategy: model.Strategy,
		Steps: []Step{
			{
				Name:        StepRsyncTransfer,
				Source:      model.Source,
				Destination: model.Destination,
				Executor:    rsyncExec,
			},
		},
	}, nil
}

// planRelayFilesystemTransfer creates a two-step transfer via local staging.
// Step 1: Pull from remote source to local staging
// Step 2: Push from local staging to remote destination
func planRelayFilesystemTransfer(model DataMigrationModel) (*Pipeline, error) {
	// Create local staging location
	stagingLoc := createLocalStagingLocation()

	// Step 1: Pull (SSH source → local staging)
	pullExec, err := NewRsyncExecutor(model.Source, stagingLoc)
	if err != nil {
		return nil, fmt.Errorf("failed to create pull executor: %w", err)
	}

	// Step 2: Push (local staging → SSH destination)
	pushExec, err := NewRsyncExecutor(stagingLoc, model.Destination)
	if err != nil {
		return nil, fmt.Errorf("failed to create push executor: %w", err)
	}

	return &Pipeline{
		Name:     PipelineFilesystemTransfer,
		Strategy: model.Strategy,
		Steps: []Step{
			{
				Name:        "pull-to-staging",
				Source:      model.Source,
				Destination: stagingLoc,
				Executor:    pullExec,
			},
			{
				Name:        "push-from-staging",
				Source:      stagingLoc,
				Destination: model.Destination,
				Executor:    pushExec,
			},
		},
	}, nil
}

// ============================================================================
// Object Storage Transfer (relay-based)
// ============================================================================

// planObjectStorageTransfer handles objectstorage ↔ objectstorage transfers.
// Always uses relay via local staging: S3 → local → S3
func planObjectStorageTransfer(model DataMigrationModel) (*Pipeline, error) {
	// Create staging location
	stagingLoc := createLocalStagingLocation()

	// Create providers for source and destination
	srcProvider, err := NewS3Provider(model.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to create source S3 provider: %w", err)
	}
	dstProvider, err := NewS3Provider(model.Destination)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination S3 provider: %w", err)
	}

	srcS3Exec := NewS3Executor(srcProvider)
	dstS3Exec := NewS3Executor(dstProvider)

	return &Pipeline{
		Name:     PipelineObjectStorageTransfer,
		Strategy: model.Strategy,
		Steps: []Step{
			{
				Name:        StepDownloadFromS3,
				Source:      model.Source,
				Destination: stagingLoc,
				Executor:    srcS3Exec,
			},
			{
				Name:        StepUploadToS3,
				Source:      stagingLoc,
				Destination: model.Destination,
				Executor:    dstS3Exec,
			},
		},
	}, nil
}

// ============================================================================
// Cross-Storage Transfer
// ============================================================================

// planCrossStorageTransfer handles filesystem ↔ objectstorage transfers.
// Determines direction and whether relay is needed.
func planCrossStorageTransfer(model DataMigrationModel) (*Pipeline, error) {
	srcIsFS := model.Source.IsFilesystem()

	if srcIsFS {
		// filesystem → objectstorage
		return planFilesystemToObjectStorage(model)
	}
	// objectstorage → filesystem
	return planObjectStorageToFilesystem(model)
}

// planFilesystemToObjectStorage handles filesystem → objectstorage.
// If source is SSH, uses relay: ssh → local → s3
// If source is local, direct: local → s3
func planFilesystemToObjectStorage(model DataMigrationModel) (*Pipeline, error) {
	s3Provider, err := NewS3Provider(model.Destination)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 provider: %w", err)
	}
	s3Exec := NewS3Executor(s3Provider)

	// Direct transfer if source is local filesystem
	if model.Source.IsLocal() {
		return &Pipeline{
			Name:     PipelineCrossStorageTransfer,
			Strategy: model.Strategy,
			Steps: []Step{
				{
					Name:        StepUploadToS3,
					Source:      model.Source,
					Destination: model.Destination,
					Executor:    s3Exec,
				},
			},
		}, nil
	}

	// Relay transfer: ssh → local → s3
	stagingLoc := createLocalStagingLocation()
	rsyncExec, err := NewRsyncExecutor(model.Source, stagingLoc)
	if err != nil {
		return nil, fmt.Errorf("failed to create rsync executor: %w", err)
	}

	return &Pipeline{
		Name:     PipelineCrossStorageTransfer,
		Strategy: model.Strategy,
		Steps: []Step{
			{
				Name:        StepRsyncFromServer,
				Source:      model.Source,
				Destination: stagingLoc,
				Executor:    rsyncExec,
			},
			{
				Name:        StepUploadToS3,
				Source:      stagingLoc,
				Destination: model.Destination,
				Executor:    s3Exec,
			},
		},
	}, nil
}

// planObjectStorageToFilesystem handles objectstorage → filesystem.
// If destination is SSH, uses relay: s3 → local → ssh
// If destination is local, direct: s3 → local
func planObjectStorageToFilesystem(model DataMigrationModel) (*Pipeline, error) {
	s3Provider, err := NewS3Provider(model.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 provider: %w", err)
	}
	s3Exec := NewS3Executor(s3Provider)

	// Direct transfer if destination is local filesystem
	if model.Destination.IsLocal() {
		return &Pipeline{
			Name:     PipelineCrossStorageTransfer,
			Strategy: model.Strategy,
			Steps: []Step{
				{
					Name:        StepDownloadFromS3,
					Source:      model.Source,
					Destination: model.Destination,
					Executor:    s3Exec,
				},
			},
		}, nil
	}

	// Relay transfer: s3 → local → ssh
	stagingLoc := createLocalStagingLocation()
	rsyncExec, err := NewRsyncExecutor(stagingLoc, model.Destination)
	if err != nil {
		return nil, fmt.Errorf("failed to create rsync executor: %w", err)
	}

	return &Pipeline{
		Name:     PipelineCrossStorageTransfer,
		Strategy: model.Strategy,
		Steps: []Step{
			{
				Name:        StepDownloadFromS3,
				Source:      model.Source,
				Destination: stagingLoc,
				Executor:    s3Exec,
			},
			{
				Name:        StepRsyncToServer,
				Source:      stagingLoc,
				Destination: model.Destination,
				Executor:    rsyncExec,
			},
		},
	}, nil
}

// ============================================================================
// Helper Functions
// ============================================================================

// createLocalStagingLocation creates a local filesystem location for staging.
func createLocalStagingLocation() DataLocation {
	return DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        DefaultStagingPath,
		Filesystem: &FilesystemAccess{
			AccessType: AccessTypeLocal,
		},
	}
}

// NewS3Provider creates an S3 provider from DataLocation.
func NewS3Provider(loc DataLocation) (S3Provider, error) {
	if !loc.IsObjectStorage() || loc.ObjectStorage == nil {
		return nil, fmt.Errorf("location is not object storage")
	}

	bucket, _ := ParseBucketAndKey(loc.Path)
	if bucket == "" {
		return nil, fmt.Errorf("bucket name is required in path")
	}

	osAccess := loc.ObjectStorage

	switch osAccess.AccessType {
	case AccessTypeMinio:
		if osAccess.Minio == nil {
			return nil, fmt.Errorf("minio S3 config is required")
		}
		cfg := &MinioConfig{
			Endpoint:        osAccess.Minio.Endpoint,
			AccessKeyId:     osAccess.Minio.AccessKeyId,
			SecretAccessKey: osAccess.Minio.SecretAccessKey,
			Region:          osAccess.Minio.Region,
			UseSSL:          osAccess.Minio.UseSSL,
		}
		return NewMinioProvider(cfg, bucket)

	case AccessTypeSpider:
		if osAccess.Spider == nil {
			return nil, fmt.Errorf("Spider config is required")
		}
		return NewSpiderProvider(osAccess.Spider, bucket)

	case AccessTypeTumblebug:
		if osAccess.Tumblebug == nil {
			return nil, fmt.Errorf("Tumblebug config is required")
		}
		return NewTumblebugProvider(osAccess.Tumblebug)

	default:
		return nil, fmt.Errorf("unsupported object storage access type: %s", osAccess.AccessType)
	}
}
