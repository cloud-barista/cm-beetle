package transx

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ============================================================================
// Pipeline Definition
// ============================================================================

// Pipeline represents a planned transfer with one or more steps.
type Pipeline struct {
	Name              string
	Strategy          string
	EffectiveStrategy string
	Steps             []Step
	Fallback          *Pipeline
}

// Step represents a single transfer step in the pipeline.
type Step struct {
	Name        string
	Mode        TransferMode
	Source      DataLocation
	Destination DataLocation
	Executor    Executor
}

// Execute runs all steps in the pipeline sequentially and cleans up staging directories on success.
func (p *Pipeline) Execute() error {
	for _, step := range p.Steps {
		if step.Destination.IsLocal() && strings.HasPrefix(step.Destination.Path, DefaultStagingPath+"/") {
			_ = os.RemoveAll(step.Destination.Path)
			if err := os.MkdirAll(step.Destination.Path, 0755); err != nil {
				return fmt.Errorf("failed to prepare staging directory %s: %w", step.Destination.Path, err)
			}
		}
	}

	for i, step := range p.Steps {
		if err := step.Executor.Execute(step.Source, step.Destination); err != nil {
			var preFlightErr *PreFlightError
			if errors.As(err, &preFlightErr) && p.Fallback != nil {
				if fallbackErr := p.Fallback.Execute(); fallbackErr != nil {
					return fmt.Errorf("auto-probing pre-flight failed (%v) and relay fallback also failed: %w", err, fallbackErr)
				}
				p.EffectiveStrategy = p.Fallback.EffectiveStrategy
				if p.EffectiveStrategy == "" {
					p.EffectiveStrategy = StrategyRelay
				}
				return nil
			}
			return fmt.Errorf("step %d (%s) failed: %w", i+1, step.Name, err)
		}

		if rsyncExec, ok := step.Executor.(*RsyncExecutor); ok && rsyncExec.ResolvedMode != "" {
			p.Steps[i].Mode = rsyncExec.ResolvedMode
			if len(p.Steps) == 1 && (rsyncExec.ResolvedMode == TransferModeSSHAgentFwdPush || rsyncExec.ResolvedMode == TransferModeSSHAgentFwdPull) {
				p.EffectiveStrategy = string(rsyncExec.ResolvedMode)
			}
		}
	}

	for _, step := range p.Steps {
		if step.Destination.IsLocal() && strings.HasPrefix(step.Destination.Path, DefaultStagingPath+"/") {
			_ = os.RemoveAll(step.Destination.Path)
		}
	}

	return nil
}

// ============================================================================
// Plan Function: Single Entry Point
// ============================================================================

// Plan analyzes the DataMigrationModel and returns the optimal transfer Pipeline.
func Plan(model DataMigrationModel) (*Pipeline, error) {
	if err := Validate(model); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	if strings.TrimSpace(model.Strategy) == "" {
		model.Strategy = StrategyAuto
	}

	srcStorage := model.Source.StorageType
	dstStorage := model.Destination.StorageType

	switch {
	case srcStorage == StorageTypeFilesystem && dstStorage == StorageTypeFilesystem:
		return planFilesystemTransfer(model)

	case srcStorage == StorageTypeObjectStorage && dstStorage == StorageTypeObjectStorage:
		return planObjectStorageTransfer(model)

	default:
		return planCrossStorageTransfer(model)
	}
}

// ============================================================================
// Filesystem Transfer (rsync-based)
// ============================================================================

// planFilesystemTransfer handles all filesystem-to-filesystem transfers.
func planFilesystemTransfer(model DataMigrationModel) (*Pipeline, error) {
	srcIsRemote := model.Source.Filesystem != nil && model.Source.Filesystem.SSH != nil
	dstIsRemote := model.Destination.Filesystem != nil && model.Destination.Filesystem.SSH != nil
	strategy := model.Strategy

	if srcIsRemote && dstIsRemote {
		switch strategy {
		case StrategyRelay:
			return planRelayFilesystemTransfer(model)

		case StrategySSHAgentFwdPush:
			rsyncExec, err := NewRsyncExecutorWithMode(model.Source, model.Destination, TransferModeSSHAgentFwdPush)
			if err != nil {
				return nil, fmt.Errorf("failed to create rsync executor: %w", err)
			}
			return &Pipeline{
				Name:              PipelineFilesystemTransfer,
				Strategy:          strategy,
				EffectiveStrategy: StrategySSHAgentFwdPush,
				Steps: []Step{
					{
						Name:        StepRsyncTransfer,
						Mode:        TransferModeSSHAgentFwdPush,
						Source:      model.Source,
						Destination: model.Destination,
						Executor:    rsyncExec,
					},
				},
			}, nil

		case StrategySSHAgentFwdPull:
			rsyncExec, err := NewRsyncExecutorWithMode(model.Source, model.Destination, TransferModeSSHAgentFwdPull)
			if err != nil {
				return nil, fmt.Errorf("failed to create rsync executor: %w", err)
			}
			return &Pipeline{
				Name:              PipelineFilesystemTransfer,
				Strategy:          strategy,
				EffectiveStrategy: StrategySSHAgentFwdPull,
				Steps: []Step{
					{
						Name:        StepRsyncTransfer,
						Mode:        TransferModeSSHAgentFwdPull,
						Source:      model.Source,
						Destination: model.Destination,
						Executor:    rsyncExec,
					},
				},
			}, nil

		case StrategyDirect:
			rsyncExec, err := NewRsyncExecutorWithMode(model.Source, model.Destination, TransferModeSSHAgentFwdAuto)
			if err != nil {
				return nil, fmt.Errorf("failed to create rsync executor: %w", err)
			}
			return &Pipeline{
				Name:              PipelineFilesystemTransfer,
				Strategy:          strategy,
				EffectiveStrategy: StrategyDirect,
				Steps: []Step{
					{
						Name:        StepRsyncTransfer,
						Mode:        TransferModeSSHAgentFwdAuto,
						Source:      model.Source,
						Destination: model.Destination,
						Executor:    rsyncExec,
					},
				},
			}, nil

		default:
			rsyncExec, err := NewRsyncExecutorWithMode(model.Source, model.Destination, TransferModeSSHAgentFwdAuto)
			if err != nil {
				return nil, fmt.Errorf("failed to create rsync executor: %w", err)
			}
			relayFallback, err := planRelayFilesystemTransfer(model)
			if err != nil {
				return nil, fmt.Errorf("failed to create relay fallback pipeline: %w", err)
			}
			return &Pipeline{
				Name:              PipelineFilesystemTransfer,
				Strategy:          strategy,
				EffectiveStrategy: StrategyAuto,
				Steps: []Step{
					{
						Name:        StepRsyncTransfer,
						Mode:        TransferModeSSHAgentFwdAuto,
						Source:      model.Source,
						Destination: model.Destination,
						Executor:    rsyncExec,
					},
				},
				Fallback: relayFallback,
			}, nil
		}
	}

	rsyncExec, err := NewRsyncExecutor(model.Source, model.Destination)
	if err != nil {
		return nil, fmt.Errorf("failed to create rsync executor: %w", err)
	}

	return &Pipeline{
		Name:              PipelineFilesystemTransfer,
		Strategy:          strategy,
		EffectiveStrategy: StrategyDirect,
		Steps: []Step{
			{
				Name:        StepRsyncTransfer,
				Mode:        rsyncExec.Mode,
				Source:      model.Source,
				Destination: model.Destination,
				Executor:    rsyncExec,
			},
		},
	}, nil
}

// planRelayFilesystemTransfer creates a two-step transfer via local staging (pull -> push).
func planRelayFilesystemTransfer(model DataMigrationModel) (*Pipeline, error) {
	stagingLoc := createLocalStagingLocation(model.Source.Path)

	pullExec, err := NewRsyncExecutorWithMode(model.Source, stagingLoc, TransferModePull)
	if err != nil {
		return nil, fmt.Errorf("failed to create pull executor: %w", err)
	}

	pushExec, err := NewRsyncExecutorWithMode(stagingLoc, model.Destination, TransferModePush)
	if err != nil {
		return nil, fmt.Errorf("failed to create push executor: %w", err)
	}

	return &Pipeline{
		Name:              PipelineFilesystemTransfer,
		Strategy:          StrategyRelay,
		EffectiveStrategy: StrategyRelay,
		Steps: []Step{
			{
				Name:        StepPullToStaging,
				Mode:        TransferModePull,
				Source:      model.Source,
				Destination: stagingLoc,
				Executor:    pullExec,
			},
			{
				Name:        StepPushFromStaging,
				Mode:        TransferModePush,
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

// planObjectStorageTransfer handles objectstorage-to-objectstorage transfers via local staging.
func planObjectStorageTransfer(model DataMigrationModel) (*Pipeline, error) {
	stagingLoc := createLocalStagingLocation(model.Source.Path)

	srcProvider, err := NewS3Provider(model.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to create source S3 provider: %w", err)
	}
	dstProvider, err := NewS3Provider(model.Destination)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination S3 provider: %w", err)
	}

	srcS3Exec := NewS3Executor(srcProvider, TransferModePull)
	dstS3Exec := NewS3Executor(dstProvider, TransferModePush)

	return &Pipeline{
		Name:              PipelineObjectStorageTransfer,
		Strategy:          model.Strategy,
		EffectiveStrategy: StrategyRelay,
		Steps: []Step{
			{
				Name:        StepDownloadFromS3,
				Mode:        TransferModePull,
				Source:      model.Source,
				Destination: stagingLoc,
				Executor:    srcS3Exec,
			},
			{
				Name:        StepUploadToS3,
				Mode:        TransferModePush,
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

// planCrossStorageTransfer handles filesystem-to-objectstorage and objectstorage-to-filesystem transfers.
func planCrossStorageTransfer(model DataMigrationModel) (*Pipeline, error) {
	if model.Source.IsFilesystem() {
		return planFilesystemToObjectStorage(model)
	}
	return planObjectStorageToFilesystem(model)
}

// planFilesystemToObjectStorage handles filesystem-to-objectstorage transfers.
func planFilesystemToObjectStorage(model DataMigrationModel) (*Pipeline, error) {
	s3Provider, err := NewS3Provider(model.Destination)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 provider: %w", err)
	}
	s3Exec := NewS3Executor(s3Provider, TransferModePush)

	if model.Source.IsLocal() {
		return &Pipeline{
			Name:              PipelineCrossStorageTransfer,
			Strategy:          model.Strategy,
			EffectiveStrategy: StrategyDirect,
			Steps: []Step{
				{
					Name:        StepUploadToS3,
					Mode:        TransferModePush,
					Source:      model.Source,
					Destination: model.Destination,
					Executor:    s3Exec,
				},
			},
		}, nil
	}

	stagingLoc := createLocalStagingLocation(model.Source.Path)
	rsyncExec, err := NewRsyncExecutorWithMode(model.Source, stagingLoc, TransferModePull)
	if err != nil {
		return nil, fmt.Errorf("failed to create rsync executor: %w", err)
	}

	return &Pipeline{
		Name:              PipelineCrossStorageTransfer,
		Strategy:          model.Strategy,
		EffectiveStrategy: StrategyRelay,
		Steps: []Step{
			{
				Name:        StepRsyncFromServer,
				Mode:        TransferModePull,
				Source:      model.Source,
				Destination: stagingLoc,
				Executor:    rsyncExec,
			},
			{
				Name:        StepUploadToS3,
				Mode:        TransferModePush,
				Source:      stagingLoc,
				Destination: model.Destination,
				Executor:    s3Exec,
			},
		},
	}, nil
}

// planObjectStorageToFilesystem handles objectstorage-to-filesystem transfers.
func planObjectStorageToFilesystem(model DataMigrationModel) (*Pipeline, error) {
	s3Provider, err := NewS3Provider(model.Source)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 provider: %w", err)
	}
	s3Exec := NewS3Executor(s3Provider, TransferModePull)

	if model.Destination.IsLocal() {
		return &Pipeline{
			Name:              PipelineCrossStorageTransfer,
			Strategy:          model.Strategy,
			EffectiveStrategy: StrategyDirect,
			Steps: []Step{
				{
					Name:        StepDownloadFromS3,
					Mode:        TransferModePull,
					Source:      model.Source,
					Destination: model.Destination,
					Executor:    s3Exec,
				},
			},
		}, nil
	}

	stagingLoc := createLocalStagingLocation(model.Source.Path)
	rsyncExec, err := NewRsyncExecutorWithMode(stagingLoc, model.Destination, TransferModePush)
	if err != nil {
		return nil, fmt.Errorf("failed to create rsync executor: %w", err)
	}

	return &Pipeline{
		Name:              PipelineCrossStorageTransfer,
		Strategy:          model.Strategy,
		EffectiveStrategy: StrategyRelay,
		Steps: []Step{
			{
				Name:        StepDownloadFromS3,
				Mode:        TransferModePull,
				Source:      model.Source,
				Destination: stagingLoc,
				Executor:    s3Exec,
			},
			{
				Name:        StepRsyncToServer,
				Mode:        TransferModePush,
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

// createLocalStagingLocation creates an isolated local filesystem staging subdirectory under DefaultStagingPath.
func createLocalStagingLocation(sourcePath string) DataLocation {
	subDir := "root"
	if bucket, _ := ParseBucketAndKey(sourcePath); bucket != "" {
		cleanBucket := filepath.Base(bucket)
		if cleanBucket != "" && cleanBucket != "." && cleanBucket != "/" {
			subDir = cleanBucket
		}
	} else if sourcePath != "" {
		cleanPath := filepath.Base(sourcePath)
		if cleanPath != "" && cleanPath != "." && cleanPath != "/" {
			subDir = cleanPath
		}
	}
	stagingPath := filepath.Join(DefaultStagingPath, subDir)
	return DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        stagingPath,
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
