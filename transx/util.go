package transx

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// performDirectTransfer handles direct transfer where at least one endpoint is local
func performDirectTransfer(dmm DataMigrationModel) error {

	// Validate
	if dmm.Source.IsRemote() && (dmm.SourceTransferOptions == nil || dmm.SourceTransferOptions.Method == "") {
		return fmt.Errorf("source transfer method must be explicitly specified in SourceTransferOptions.Method")
	}
	if dmm.Destination.IsRemote() && (dmm.DestinationTransferOptions == nil || dmm.DestinationTransferOptions.Method == "") {
		return fmt.Errorf("destination transfer method must be explicitly specified in DestinationTransferOptions.Method")
	}

	// Determine transfer method based on which endpoint is local/remote
	transferMethod := ""
	if dmm.Source.IsLocal() && dmm.Destination.IsRemote() {
		transferMethod = dmm.DestinationTransferOptions.Method
	} else if dmm.Source.IsRemote() && dmm.Destination.IsLocal() {
		transferMethod = dmm.SourceTransferOptions.Method
	} else {
		return fmt.Errorf("direct transfer requires at least one endpoint to be local")
	}

	// Perform transfer based on the determined method
	switch transferMethod {
	case TransferMethodRsync:
		// Use rsync for transfer
		if err := performRsyncTransfer(dmm); err != nil {
			// If it's already an OperationError, return as is
			if _, ok := err.(*OperationError); ok {
				return err
			}
			// Otherwise, wrap in OperationError
			return &OperationError{
				Operation:   "transfer",
				Method:      transferMethod,
				Source:      dmm.Source.DataPath,
				Destination: dmm.Destination.DataPath,
				IsRelayMode: false,
				Err:         err,
			}
		}
		return nil

	case TransferMethodObjectStorage:
		// Use Object Storage API for transfer
		if err := performObjectStorageTransfer(dmm); err != nil {
			// If it's already an OperationError, return as is
			if _, ok := err.(*OperationError); ok {
				return err
			}
			// Otherwise, wrap in OperationError
			return &OperationError{
				Operation:   "transfer",
				Method:      transferMethod,
				Source:      dmm.Source.DataPath,
				Destination: dmm.Destination.DataPath,
				IsRelayMode: false,
				Err:         err,
			}
		}
		return nil

	default:
		return &OperationError{
			Operation:   "transfer",
			Method:      transferMethod,
			Source:      dmm.Source.DataPath,
			Destination: dmm.Destination.DataPath,
			IsRelayMode: false,
			Err:         fmt.Errorf("unsupported direct transfer method: %s", transferMethod),
		}
	}
}

// performRelayTransfer handles relay transfer where both endpoints are remote
func performRelayTransfer(dmm DataMigrationModel) error {

	// Create temporary directory for relay
	tempDir, err := os.MkdirTemp("", "transx-relay-*")
	if err != nil {
		return &OperationError{
			Operation:   "transfer",
			Method:      "relay",
			Source:      dmm.Source.DataPath,
			Destination: dmm.Destination.DataPath,
			IsRelayMode: true,
			Err:         fmt.Errorf("failed to create temporary directory: %w", err),
		}
	}
	defer os.RemoveAll(tempDir)

	// Step 1: Download from source to local temp
	tempDmm := dmm
	tempDmm.Destination = EndpointDetails{DataPath: tempDir}

	if err := performDirectTransfer(tempDmm); err != nil {
		return &OperationError{
			Operation:   "transfer",
			Method:      "relay",
			Source:      dmm.Source.DataPath,
			Destination: dmm.Destination.DataPath,
			IsRelayMode: true,
			Err:         fmt.Errorf("relay step 1 (source → relay node) failed: %w", err),
		}
	}

	// Step 2: Upload from local temp to destination
	tempDmm = dmm
	tempDmm.Source = EndpointDetails{DataPath: tempDir}

	if err := performDirectTransfer(tempDmm); err != nil {
		return &OperationError{
			Operation:   "transfer",
			Method:      "relay",
			Source:      dmm.Source.DataPath,
			Destination: dmm.Destination.DataPath,
			IsRelayMode: true,
			Err:         fmt.Errorf("relay step 2 (relay node → destination) failed: %w", err),
		}
	}

	return nil
}

// IsRelayMode determines if both source and destination endpoints are remote.
// This is used to identify relay migration scenarios where data needs to flow through the local machine
// as an intermediary between two remote endpoints.
func (dmm *DataMigrationModel) IsRelayMode() bool {
	return dmm.Source.IsRemote() && dmm.Destination.IsRemote()
}

// GetEndpoint returns the endpoint (SSH host/IP or Object Storage endpoint).
func (e *EndpointDetails) GetEndpoint() string {
	return e.Endpoint
}

// GetPort returns the SSH port.
func (e *EndpointDetails) GetPort() int {
	return e.Port
}

// GetBucketAndObjectKey extracts bucket name and object key from DataPath.
// DataPath formats:
//   - "bucket-name/object/key/path" (e.g., "spider-test-bucket/a/b/c/test.txt")
//   - "bucket-name/" (download all objects in bucket)
//   - "bucket-name/prefix/" (download objects with prefix)
func (e *EndpointDetails) GetBucketAndObjectKey() (string, string, error) {
	if e.DataPath == "" {
		return "", "", fmt.Errorf("DataPath is empty")
	}

	parts := strings.SplitN(e.DataPath, "/", 2)
	if len(parts) < 1 {
		return "", "", fmt.Errorf("DataPath must contain at least bucket name (e.g., 'bucket' or 'bucket/object-key')")
	}

	// If only bucket name is provided (no "/" or ends with "/")
	if len(parts) == 1 || (len(parts) == 2 && parts[1] == "") {
		return parts[0], "", nil // Empty object key means list all objects in bucket
	}

	return parts[0], parts[1], nil
}

// IsObjectStorageEndpoint checks if this endpoint is for Object Storage
func (e *EndpointDetails) IsObjectStorageEndpoint() bool {
	return strings.Contains(e.Endpoint, "/spider/s3")
}

// IsRemote determines if the EndpointDetails represent a remote endpoint.
// Returns true if the endpoint has a non-empty host/endpoint configured.
func (e *EndpointDetails) IsRemote() bool {
	return strings.TrimSpace(e.GetEndpoint()) != ""
}

// IsLocal determines if the EndpointDetails represent a local endpoint.
func (e *EndpointDetails) IsLocal() bool {
	return strings.TrimSpace(e.GetEndpoint()) == ""
}

// GetRsyncPath constructs the path string suitable for rsync (e.g., "user@host:/path" or "/local/path").
func (e *EndpointDetails) GetRsyncPath(options *TransferOptions) string {
	if e.IsRemote() {
		var username string
		if options != nil && options.RsyncOptions != nil {
			username = options.RsyncOptions.Username
		}
		host := e.GetEndpoint()
		if strings.TrimSpace(username) != "" {
			return fmt.Sprintf("%s@%s:%s", username, host, e.DataPath)
		}
		return fmt.Sprintf("%s:%s", host, e.DataPath) // Username might be optional if SSH config handles it
	}
	return e.DataPath
}

/*
 * Rsync transfer functions
 */

// performRsyncTransfer performs rsync-based transfer
func performRsyncTransfer(dmm DataMigrationModel) error {
	rsyncCmdPath := "rsync"
	sourceOptions := dmm.SourceTransferOptions
	if sourceOptions != nil && sourceOptions.RsyncOptions != nil && sourceOptions.RsyncOptions.RsyncPath != "" {
		rsyncCmdPath = sourceOptions.RsyncOptions.RsyncPath
	}

	var args []string
	// Configure basic rsync options
	if sourceOptions != nil && sourceOptions.RsyncOptions != nil && sourceOptions.RsyncOptions.Archive {
		args = append(args, "-a")
	}
	if sourceOptions != nil && sourceOptions.RsyncOptions != nil && sourceOptions.RsyncOptions.Compress {
		args = append(args, "-z")
	}
	if sourceOptions != nil && sourceOptions.RsyncOptions != nil && sourceOptions.RsyncOptions.Verbose {
		args = append(args, "-v")
	}
	if sourceOptions != nil && sourceOptions.RsyncOptions != nil && sourceOptions.RsyncOptions.Delete {
		args = append(args, "--delete")
	}
	if sourceOptions != nil && sourceOptions.RsyncOptions != nil && sourceOptions.RsyncOptions.Progress {
		args = append(args, "--progress")
	}
	if sourceOptions != nil && sourceOptions.RsyncOptions != nil && sourceOptions.RsyncOptions.DryRun {
		args = append(args, "-n") // or "--dry-run"
	}

	// Configure Exclude and Include options
	if sourceOptions != nil && sourceOptions.RsyncOptions != nil && sourceOptions.RsyncOptions.Filter != nil {
		filter := sourceOptions.RsyncOptions.Filter
		for _, ex := range filter.Exclude {
			if strings.TrimSpace(ex) != "" {
				args = append(args, "--exclude="+ex)
			}
		}
		for _, inc := range filter.Include {
			if strings.TrimSpace(inc) != "" {
				args = append(args, "--include="+inc)
			}
		}
	}

	// Configure SSH options (-e)
	var sshOptString string
	var activeRemoteEndpointForRsync EndpointDetails
	operationInvolvesRemoteRsync := false

	if dmm.Source.IsRemote() {
		activeRemoteEndpointForRsync = dmm.Source
		operationInvolvesRemoteRsync = true
	} else if dmm.Destination.IsRemote() {
		activeRemoteEndpointForRsync = dmm.Destination
		operationInvolvesRemoteRsync = true
	}

	if operationInvolvesRemoteRsync && sourceOptions != nil && sourceOptions.RsyncOptions != nil && sourceOptions.RsyncOptions.SSHPrivateKeyPath != "" {
		var sshCmdParts []string
		sshCmdParts = append(sshCmdParts, "ssh")
		if strings.TrimSpace(sourceOptions.RsyncOptions.SSHPrivateKeyPath) != "" {
			sshCmdParts = append(sshCmdParts, "-i", sourceOptions.RsyncOptions.SSHPrivateKeyPath)
		}
		if activeRemoteEndpointForRsync.GetPort() != 0 { // If 0, use default port (22)
			sshCmdParts = append(sshCmdParts, "-p", strconv.Itoa(activeRemoteEndpointForRsync.GetPort()))
		}
		if sourceOptions.RsyncOptions.InsecureSkipHostKeyVerification {
			sshCmdParts = append(sshCmdParts, "-o", "StrictHostKeyChecking=accept-new")
			sshCmdParts = append(sshCmdParts, "-o", "UserKnownHostsFile=/dev/null")
		}

		// Set connection timeout if specified
		if sourceOptions.RsyncOptions.ConnectTimeout > 0 {
			connectTimeout := sourceOptions.RsyncOptions.ConnectTimeout
			sshCmdParts = append(sshCmdParts, "-o", fmt.Sprintf("ConnectTimeout=%d", connectTimeout))
		}

		sshOptString = strings.Join(sshCmdParts, " ")
	}

	if sshOptString != "" {
		args = append(args, "-e", sshOptString)
	}

	// Add source and destination paths
	sourceRsyncPath := dmm.Source.GetRsyncPath(sourceOptions)
	destinationRsyncPath := dmm.Destination.GetRsyncPath(sourceOptions) // Use sourceOptions for destination path too since it contains the SSH config

	// Handle TransferDirContentsOnly option
	if sourceOptions != nil && sourceOptions.RsyncOptions != nil && sourceOptions.RsyncOptions.TransferDirContentsOnly && !strings.HasSuffix(sourceRsyncPath, "/") {
		sourceRsyncPath += "/"
	}

	args = append(args, sourceRsyncPath, destinationRsyncPath)

	// Execute rsync command
	cmd := exec.Command(rsyncCmdPath, args...)

	// Capture output for error reporting
	output, err := cmd.CombinedOutput()
	if err != nil {
		return &OperationError{
			Operation:   "transfer",
			Method:      TransferMethodRsync,
			Source:      sourceRsyncPath,
			Destination: destinationRsyncPath,
			IsRelayMode: false, // Direct transfer
			Err:         fmt.Errorf("%w (command: %s %s)\nOutput: %s", err, rsyncCmdPath, strings.Join(args, " "), string(output)),
		}
	}

	return nil
}

/*
 * Object Storage transfer functions
 */

// performObjectStorageTransfer performs Object Storage API-based transfer for directories and files
func performObjectStorageTransfer(dmm DataMigrationModel) error {
	// Determine transfer direction and validate options
	var transferOptions *TransferOptions
	var isUpload bool

	if dmm.Source.IsLocal() && dmm.Destination.IsRemote() {
		// Upload: Local to Object Storage
		isUpload = true
		transferOptions = dmm.DestinationTransferOptions
		if transferOptions == nil || transferOptions.ObjectStorageOptions == nil {
			return fmt.Errorf("destination ObjectStorageOptions is required for upload")
		}
	} else if dmm.Source.IsRemote() && dmm.Destination.IsLocal() {
		// Download: Object Storage to Local
		isUpload = false
		transferOptions = dmm.SourceTransferOptions
		if transferOptions == nil || transferOptions.ObjectStorageOptions == nil {
			return fmt.Errorf("source ObjectStorageOptions is required for download")
		}
	} else {
		return fmt.Errorf("object Storage transfer requires one local and one remote endpoint")
	}

	// Check if bucket exists before proceeding with transfer
	var bucketEndpoint EndpointDetails
	if isUpload {
		bucketEndpoint = dmm.Destination
	} else {
		bucketEndpoint = dmm.Source
	}

	if err := checkBucketExists(bucketEndpoint, transferOptions.ObjectStorageOptions); err != nil {
		return fmt.Errorf("bucket validation failed: %w", err)
	}

	if isUpload {
		// Upload: Local to Object Storage
		localPath := dmm.Source.DataPath

		// Check if source is a file or directory
		fileInfo, err := os.Stat(localPath)
		if err != nil {
			return fmt.Errorf("failed to stat source path %s: %w", localPath, err)
		}

		if fileInfo.IsDir() {
			// Handle directory upload recursively
			return filepath.Walk(localPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				// Sleep a bit for stability
				time.Sleep(1 * time.Millisecond)

				// Skip directories themselves, only process files
				if info.IsDir() {
					return nil
				}

				// Calculate relative path from base directory
				relPath, err := filepath.Rel(localPath, path)
				if err != nil {
					return fmt.Errorf("failed to calculate relative path for %s: %w", path, err)
				}

				// Apply filtering if filter patterns are specified
				if transferOptions.ObjectStorageOptions != nil && transferOptions.ObjectStorageOptions.Filter != nil {
					filter := transferOptions.ObjectStorageOptions.Filter
					if len(filter.Exclude) > 0 || len(filter.Include) > 0 {
						// Use relative path for pattern matching
						if !shouldTransferObject(relPath, filter.Exclude, filter.Include) {
							return nil // Skip this file
						}
					}
				}

				// Construct object key by combining destination path with relative path
				objectPath := filepath.ToSlash(filepath.Join(dmm.Destination.DataPath, relPath))

				// Upload the individual file
				return uploadFileToObjectStorage(path, objectPath, dmm.Destination, transferOptions)
			})
		} else {
			// Handle single file upload
			return uploadFileToObjectStorage(localPath, dmm.Destination.DataPath, dmm.Destination, transferOptions)
		}
	} else {
		// Download: Object Storage to Local
		// Create destination directory if it doesn't exist
		if err := os.MkdirAll(dmm.Destination.DataPath, 0755); err != nil {
			return fmt.Errorf("failed to create destination directory: %w", err)
		}

		// Check if source DataPath represents a single file or directory prefix
		bucket, objectKey, err := dmm.Source.GetBucketAndObjectKey()
		if err != nil {
			return fmt.Errorf("failed to parse source path: %w", err)
		}

		// List objects with the specified prefix
		objects, err := listBucketObjects(dmm.Source, objectKey, transferOptions.ObjectStorageOptions)
		if err != nil {
			return fmt.Errorf("failed to list objects: %w", err)
		}

		// Apply filtering if filter patterns are specified
		if transferOptions.ObjectStorageOptions != nil && transferOptions.ObjectStorageOptions.Filter != nil {
			filter := transferOptions.ObjectStorageOptions.Filter
			if len(filter.Exclude) > 0 || len(filter.Include) > 0 {
				objects = filterObjectList(objects, filter.Exclude, filter.Include)
			}
		}

		if len(objects) == 0 {
			return fmt.Errorf("no objects found with prefix '%s' in bucket '%s' (after filtering)", objectKey, bucket)
		}

		// Debug: Print first few objects
		// for i, obj := range objects {
		// 	fmt.Printf("Object List[%d]: %v\n", i, obj)
		// 	if i >= 4 {
		// 		break
		// 	}
		// }

		// Download each object
		for _, obj := range objects {
			// Calculate local file path
			// Remove the prefix from object key to get relative path
			// Note: objectKey is the prefix specified in DataPath (e.g., "resources")
			// obj.Key is the full object key from bucket (e.g., "resources/story/original/file.png")
			relativePath := strings.TrimPrefix(obj.Key, objectKey)
			if relativePath == "" {
				// If the object key exactly matches the prefix, use the filename only
				relativePath = filepath.Base(obj.Key)
			} else if strings.HasPrefix(relativePath, "/") {
				// Remove leading slash
				relativePath = strings.TrimPrefix(relativePath, "/")
			}

			localFilePath := filepath.Join(dmm.Destination.DataPath, relativePath)

			// Construct object path for download
			// Both Spider API and MinIO SDK expect format: "bucket/object-key"
			// - Spider API: presigned URL format is /presigned/{operation}/{bucket}/{object-key}
			// - MinIO SDK: Internally splits "bucket/object-key" into separate parameters
			// obj.Key already contains the full object key (e.g., "resources/story/original/file.png")
			objectPath := filepath.ToSlash(filepath.Join(bucket, obj.Key))

			// Debug: Print mapping info
			//fmt.Printf("ObjectPath: %s -> LocalFilePath: %s\n", objectPath, localFilePath)

			// Download the individual file
			if err := downloadFileFromObjectStorage(localFilePath, objectPath, dmm.Source, transferOptions); err != nil {
				return fmt.Errorf("failed to download object '%s' (from bucket '%s', key '%s'): %w", objectPath, bucket, obj.Key, err)
			}
		}

		return nil
	}
}
