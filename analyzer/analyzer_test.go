package analyzer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetDefaultBaseDir(t *testing.T) {
	homeDir, err := GetDefaultBaseDir()
	if err != nil {
		t.Fatalf("GetDefaultBaseDir failed: %v", err)
	}

	if homeDir == "" {
		t.Fatal("Home directory is empty")
	}

	// Verify directory exists
	info, err := os.Stat(homeDir)
	if err != nil {
		t.Fatalf("Home directory does not exist: %v", err)
	}

	if !info.IsDir() {
		t.Fatal("Home path is not a directory")
	}

	t.Logf("Home directory: %s", homeDir)
}

func TestListDirectory(t *testing.T) {
	// Test with default directory
	result, err := ListDirectory("")
	if err != nil {
		t.Fatalf("ListDirectory with empty path failed: %v", err)
	}

	if result.BaseDir == "" {
		t.Fatal("Base directory is empty")
	}

	t.Logf("Listed directory: %s", result.BaseDir)
	t.Logf("Total files: %d, Total directories: %d", result.TotalFiles, result.TotalDirs)

	// Test with /tmp directory
	tmpResult, err := ListDirectory("/tmp")
	if err != nil {
		t.Fatalf("ListDirectory /tmp failed: %v", err)
	}

	if tmpResult.BaseDir != "/tmp" {
		t.Fatalf("Expected /tmp, got %s", tmpResult.BaseDir)
	}

	t.Logf("/tmp entries: %d", len(tmpResult.Entries))
}

func TestScanDirectory(t *testing.T) {
	// Create a temporary test directory
	tmpDir := t.TempDir()

	// Create test structure
	testFiles := []string{
		"file1.txt",
		"file2.log",
		"subdir/file3.txt",
		"subdir/file4.log",
		".hidden.txt",
	}

	for _, file := range testFiles {
		fullPath := filepath.Join(tmpDir, file)
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, []byte("test content"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Test 1: Single level scan
	t.Run("SingleLevel", func(t *testing.T) {
		options := ScanOptions{
			BaseDir:   tmpDir,
			Recursive: false,
		}

		result, err := ScanDirectory(options)
		if err != nil {
			t.Fatalf("ScanDirectory failed: %v", err)
		}

		// Should find 2 files + 1 directory at root level
		if len(result.Entries) != 3 {
			t.Fatalf("Expected 3 entries, got %d", len(result.Entries))
		}

		t.Logf("Single level: %d entries", len(result.Entries))
	})

	// Test 2: Recursive scan
	t.Run("Recursive", func(t *testing.T) {
		options := ScanOptions{
			BaseDir:   tmpDir,
			Recursive: true,
		}

		result, err := ScanDirectory(options)
		if err != nil {
			t.Fatalf("ScanDirectory failed: %v", err)
		}

		// Should find all 4 .txt/.log files + 1 directory
		if len(result.Entries) < 5 {
			t.Fatalf("Expected at least 5 entries, got %d", len(result.Entries))
		}

		t.Logf("Recursive: %d entries", len(result.Entries))
	})

	// Test 3: With exclude pattern
	t.Run("WithExcludePattern", func(t *testing.T) {
		options := ScanOptions{
			BaseDir:         tmpDir,
			Recursive:       true,
			ExcludePatterns: []string{"*.log"},
		}

		result, err := ScanDirectory(options)
		if err != nil {
			t.Fatalf("ScanDirectory failed: %v", err)
		}

		// Should exclude .log files
		for _, entry := range result.Entries {
			if filepath.Ext(entry.Name) == ".log" {
				t.Fatalf("Found .log file despite exclude pattern: %s", entry.Name)
			}
		}

		t.Logf("With exclude pattern: %d entries", len(result.Entries))
	})

	// Test 4: Include hidden files
	t.Run("IncludeHidden", func(t *testing.T) {
		options := ScanOptions{
			BaseDir:       tmpDir,
			Recursive:     true,
			IncludeHidden: true,
		}

		result, err := ScanDirectory(options)
		if err != nil {
			t.Fatalf("ScanDirectory failed: %v", err)
		}

		// Should find .hidden.txt
		foundHidden := false
		for _, entry := range result.Entries {
			if entry.Name == ".hidden.txt" {
				foundHidden = true
				break
			}
		}

		if !foundHidden {
			t.Fatal("Hidden file not found")
		}

		t.Logf("With hidden files: %d entries", len(result.Entries))
	})
}

func TestExtractFileMetadata(t *testing.T) {
	// Create a temporary file
	tmpFile := filepath.Join(t.TempDir(), "test.txt")
	content := []byte("Test content for metadata extraction")
	if err := os.WriteFile(tmpFile, content, 0644); err != nil {
		t.Fatal(err)
	}

	metadata, err := ExtractFileMetadata(tmpFile, false)
	if err != nil {
		t.Fatalf("ExtractFileMetadata failed: %v", err)
	}

	if metadata.Path != tmpFile {
		t.Errorf("Expected path %s, got %s", tmpFile, metadata.Path)
	}

	if metadata.Name != "test.txt" {
		t.Errorf("Expected name test.txt, got %s", metadata.Name)
	}

	if metadata.Size != int64(len(content)) {
		t.Errorf("Expected size %d, got %d", len(content), metadata.Size)
	}

	if metadata.IsDir {
		t.Error("File incorrectly identified as directory")
	}

	if metadata.Extension != ".txt" {
		t.Errorf("Expected extension .txt, got %s", metadata.Extension)
	}

	t.Logf("Metadata: %+v", metadata)
}

func TestCreateMigrationPlan(t *testing.T) {
	// Create test directory structure
	tmpDir := t.TempDir()

	testFiles := map[string]string{
		"doc1.txt":        "Content 1",
		"doc2.md":         "Content 2",
		"file.log":        "Log content",
		"data.json":       "{}",
		"subdir/doc3.txt": "Content 3",
	}

	for file, content := range testFiles {
		fullPath := filepath.Join(tmpDir, file)
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	filters := FilterOptions{
		IncludePatterns: []string{"*.txt", "*.md"},
		ExcludePatterns: []string{"*.log"},
	}

	plan, err := CreateMigrationPlan(tmpDir, true, filters)
	if err != nil {
		t.Fatalf("CreateMigrationPlan failed: %v", err)
	}

	if plan.SourceDir != tmpDir {
		t.Errorf("Expected source dir %s, got %s", tmpDir, plan.SourceDir)
	}

	if !plan.IncludeSubDir {
		t.Error("Expected IncludeSubDir to be true")
	}

	// Should find .txt and .md files, excluding .log
	// Actual count depends on filter matching behavior
	if plan.TotalFiles < 2 {
		t.Errorf("Expected at least 2 files, got %d", plan.TotalFiles)
	}

	if plan.TotalSize == 0 {
		t.Error("Expected non-zero total size")
	}

	t.Logf("Migration plan: %d files, %d bytes", plan.TotalFiles, plan.TotalSize)
}

func TestMatchPattern(t *testing.T) {
	tests := []struct {
		path    string
		pattern string
		want    bool
	}{
		// Simple patterns
		{"file.txt", "*.txt", true},
		{"file.log", "*.txt", false},
		{"data.json", "*.json", true},

		// Directory patterns
		{"data/file.txt", "data/*", true},
		{"other/file.txt", "data/*", false},

		// Recursive patterns
		{"data/sub/file.txt", "data/**", true},
		{"data/file.txt", "data/**", true},
		{"other/file.txt", "data/**", false},

		// Complex patterns
		{"src/main.go", "**/*.go", true},
		{"test/unit/test.go", "**/*.go", true},
		{"readme.txt", "**/*.go", false},
	}

	for _, tt := range tests {
		t.Run(tt.path+"_"+tt.pattern, func(t *testing.T) {
			got := matchPattern(tt.path, tt.pattern)
			if got != tt.want {
				t.Errorf("matchPattern(%q, %q) = %v, want %v", tt.path, tt.pattern, got, tt.want)
			}
		})
	}
}

func TestGetDirectoryStatistics(t *testing.T) {
	// Create test directory
	tmpDir := t.TempDir()

	// Create test files
	for i := 0; i < 3; i++ {
		file := filepath.Join(tmpDir, "file"+string(rune('1'+i))+".txt")
		if err := os.WriteFile(file, []byte("test"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Create subdirectory with files
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 2; i++ {
		file := filepath.Join(subDir, "subfile"+string(rune('1'+i))+".txt")
		if err := os.WriteFile(file, []byte("test"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Test non-recursive
	fileCount, dirCount, totalSize, err := GetDirectoryStatistics(tmpDir, false)
	if err != nil {
		t.Fatalf("GetDirectoryStatistics failed: %v", err)
	}

	if fileCount != 3 {
		t.Errorf("Expected 3 files, got %d", fileCount)
	}

	if dirCount != 1 {
		t.Errorf("Expected 1 directory, got %d", dirCount)
	}

	t.Logf("Non-recursive: %d files, %d dirs, %d bytes", fileCount, dirCount, totalSize)

	// Test recursive
	fileCount, dirCount, totalSize, err = GetDirectoryStatistics(tmpDir, true)
	if err != nil {
		t.Fatalf("GetDirectoryStatistics recursive failed: %v", err)
	}

	if fileCount != 5 {
		t.Errorf("Expected 5 files (recursive), got %d", fileCount)
	}

	if dirCount != 1 {
		t.Errorf("Expected 1 directory (recursive), got %d", dirCount)
	}

	t.Logf("Recursive: %d files, %d dirs, %d bytes", fileCount, dirCount, totalSize)
}
