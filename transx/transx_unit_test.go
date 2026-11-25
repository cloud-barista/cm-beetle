package transx

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// mockPresignedS3Provider serves fixed presigned URL results with RequiredHeaders for unit tests.
type mockPresignedS3Provider struct {
	downloadURL     string
	requiredHeaders map[string]string
}

// GeneratePresignedURL returns the configured test server URL and headers.
func (m *mockPresignedS3Provider) GeneratePresignedURL(action, key string) (PresignedURLResult, error) {
	return PresignedURLResult{
		URL:             m.downloadURL,
		RequiredHeaders: m.requiredHeaders,
	}, nil
}

// ListObjects is unused in downloadObject header tests.
func (m *mockPresignedS3Provider) ListObjects(prefix string) ([]ObjectInfo, error) {
	return nil, nil
}

// GetBucket returns the mock bucket name.
func (m *mockPresignedS3Provider) GetBucket() string {
	return "mock-bucket"
}

func TestRsyncQuotingAndEndOfOptions(t *testing.T) {
	if got := shellQuote("it's $(id) & spaces"); got != "'it'\\''s $(id) & spaces'" {
		t.Errorf("Unexpected shellQuote output: %s", got)
	}

	srcLoc := DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        "-weird/dir with 'quotes'/$(id)",
		Filter: &FilterOption{
			Include: []string{"*.txt", "sub'dir/*"},
			Exclude: []string{"$(rm -rf)"},
		},
		Filesystem: &FilesystemAccess{
			AccessType: AccessTypeSSH,
			SSH: &SSHConfig{
				Host:     "10.0.0.1",
				Username: "user1",
			},
		},
	}
	dstLoc := DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        "/dst/path with space",
		Filesystem: &FilesystemAccess{
			AccessType: AccessTypeSSH,
			SSH: &SSHConfig{
				Host:     "10.0.0.2",
				Username: "user2",
			},
		},
	}

	execPush, err := NewRsyncExecutorWithMode(srcLoc, dstLoc, TransferModeSSHAgentFwdPush)
	if err != nil {
		t.Fatalf("NewRsyncExecutorWithMode(Push) failed: %v", err)
	}
	pushCmd := execPush.buildRemoteRsyncCommand(srcLoc, dstLoc, dstLoc.Filesystem.SSH)
	if !strings.Contains(pushCmd, "--include='sub'\\''dir/*'") {
		t.Errorf("Expected quoted include pattern in pushCmd: %s", pushCmd)
	}
	if !strings.Contains(pushCmd, "--exclude='$(rm -rf)'") {
		t.Errorf("Expected quoted exclude pattern in pushCmd: %s", pushCmd)
	}
	if !strings.Contains(pushCmd, "-- '-weird/dir with '\\''quotes'\\''/$(id)/' 'user2@10.0.0.2:/dst/path with space'") {
		t.Errorf("Expected -- and single-quoted src/dst in pushCmd: %s", pushCmd)
	}

	execPull, err := NewRsyncExecutorWithMode(srcLoc, dstLoc, TransferModeSSHAgentFwdPull)
	if err != nil {
		t.Fatalf("NewRsyncExecutorWithMode(Pull) failed: %v", err)
	}
	pullCmd := execPull.buildRemoteRsyncPullCommand(srcLoc, dstLoc, srcLoc.Filesystem.SSH)
	if !strings.Contains(pullCmd, "-- 'user1@10.0.0.1:-weird/dir with '\\''quotes'\\''/$(id)/' '/dst/path with space'") {
		t.Errorf("Expected -- and single-quoted src/dst in pullCmd: %s", pullCmd)
	}

	localSrc := DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        "-dash-src",
		Filesystem:  &FilesystemAccess{AccessType: AccessTypeLocal},
	}
	localDst := DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        "-dash-dst",
		Filesystem:  &FilesystemAccess{AccessType: AccessTypeLocal},
	}
	execLocal, err := NewRsyncExecutor(localSrc, localDst)
	if err != nil {
		t.Fatalf("NewRsyncExecutor(local) failed: %v", err)
	}
	localArgs := execLocal.buildLocalRsyncArgs(localSrc, localDst)
	if len(localArgs) < 3 || localArgs[len(localArgs)-3] != "--" || localArgs[len(localArgs)-2] != "-dash-src/" || localArgs[len(localArgs)-1] != "-dash-dst/" {
		t.Errorf("Expected '--' right before source/dest in localArgs, got %v", localArgs)
	}
}

func TestS3ExecutorDownloadObjectRequiredHeaders(t *testing.T) {
	transport, ok := s3TransferHTTPClient.Transport.(*http.Transport)
	if !ok || transport.ResponseHeaderTimeout != 60*time.Second {
		t.Fatalf("Expected s3TransferHTTPClient ResponseHeaderTimeout=60s, got %#v", s3TransferHTTPClient.Transport)
	}

	var receivedHeader string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		receivedHeader = r.Header.Get("X-Ms-Version")
		_, _ = w.Write([]byte("downloaded-object-bytes"))
	}))
	defer srv.Close()

	tmpDir := t.TempDir()
	dstFile := filepath.Join(tmpDir, "object.bin")
	exec := NewS3Executor(&mockPresignedS3Provider{
		downloadURL:     srv.URL + "/object.bin",
		requiredHeaders: map[string]string{"x-ms-version": "2023-11-03"},
	}, TransferModePull)

	if err := exec.downloadObject("object.bin", dstFile); err != nil {
		t.Fatalf("downloadObject failed: %v", err)
	}
	if receivedHeader != "2023-11-03" {
		t.Errorf("Expected RequiredHeaders to be sent on GET, got %q", receivedHeader)
	}
	gotBytes, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %v", err)
	}
	if string(gotBytes) != "downloaded-object-bytes" {
		t.Errorf("Downloaded file content mismatch: got %q", string(gotBytes))
	}
}

func TestSpiderProviderListObjectsPagination(t *testing.T) {
	t.Run("paginates with NextMarker", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			marker := r.URL.Query().Get("marker")
			w.Header().Set("Content-Type", "application/xml")
			if marker == "" {
				_, _ = w.Write([]byte(`<ListBucketResult><Name>b</Name><IsTruncated>true</IsTruncated><NextMarker>page-2</NextMarker><Contents><Key>a.txt</Key><Size>1</Size></Contents></ListBucketResult>`))
				return
			}
			if marker == "page-2" {
				_, _ = w.Write([]byte(`<ListBucketResult><Name>b</Name><IsTruncated>false</IsTruncated><Contents><Key>b.txt</Key><Size>2</Size></Contents></ListBucketResult>`))
				return
			}
			t.Errorf("Unexpected marker %q", marker)
		}))
		defer srv.Close()

		p, err := NewSpiderProvider(&SpiderConfig{Endpoint: srv.URL, ConnectionName: "aws-conn"}, "b")
		if err != nil {
			t.Fatalf("NewSpiderProvider failed: %v", err)
		}
		objs, err := p.ListObjects("")
		if err != nil {
			t.Fatalf("ListObjects failed: %v", err)
		}
		if len(objs) != 2 || objs[0].Key != "a.txt" || objs[1].Key != "b.txt" {
			t.Errorf("Unexpected paginated objects: %+v", objs)
		}
	})

	t.Run("paginates using last object key when NextMarker omitted", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			marker := r.URL.Query().Get("marker")
			w.Header().Set("Content-Type", "application/xml")
			if marker == "" {
				_, _ = w.Write([]byte(`<ListBucketResult><Name>b</Name><IsTruncated>true</IsTruncated><Contents><Key>k1</Key><Size>10</Size></Contents></ListBucketResult>`))
				return
			}
			if marker == "k1" {
				_, _ = w.Write([]byte(`<ListBucketResult><Name>b</Name><IsTruncated>false</IsTruncated><Contents><Key>k2</Key><Size>20</Size></Contents></ListBucketResult>`))
				return
			}
			t.Errorf("Unexpected fallback marker %q", marker)
		}))
		defer srv.Close()

		p, err := NewSpiderProvider(&SpiderConfig{Endpoint: srv.URL, ConnectionName: "aws-conn"}, "b")
		if err != nil {
			t.Fatalf("NewSpiderProvider failed: %v", err)
		}
		objs, err := p.ListObjects("k")
		if err != nil {
			t.Fatalf("ListObjects failed: %v", err)
		}
		if len(objs) != 2 || objs[0].Key != "k1" || objs[1].Key != "k2" {
			t.Errorf("Unexpected fallback-paginated objects: %+v", objs)
		}
	})

	t.Run("fails fast on truncated empty marker stall", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`<ListBucketResult><Name>b</Name><IsTruncated>true</IsTruncated></ListBucketResult>`))
		}))
		defer srv.Close()

		p, err := NewSpiderProvider(&SpiderConfig{Endpoint: srv.URL, ConnectionName: "aws-conn"}, "b")
		if err != nil {
			t.Fatalf("NewSpiderProvider failed: %v", err)
		}
		if _, err := p.ListObjects(""); err == nil || !strings.Contains(err.Error(), "marker did not advance") {
			t.Errorf("Expected marker stall error, got: %v", err)
		}
	})
}

func TestTumblebugProviderListObjectsTruncated(t *testing.T) {
	truncated := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"isTruncated":%t,"contents":[{"key":"dir/a.txt","size":10},{"key":"other/b.txt","size":20}]}`, truncated)
	}))
	defer srv.Close()

	p, err := NewTumblebugProvider(&TumblebugConfig{Endpoint: srv.URL, NsId: "ns1", OsId: "os1"})
	if err != nil {
		t.Fatalf("NewTumblebugProvider failed: %v", err)
	}

	objs, err := p.ListObjects("dir/")
	if err != nil {
		t.Fatalf("ListObjects failed: %v", err)
	}
	if len(objs) != 1 || objs[0].Key != "dir/a.txt" {
		t.Errorf("Expected filtered key dir/a.txt, got %+v", objs)
	}

	truncated = true
	if _, err := p.ListObjects(""); err == nil || !strings.Contains(err.Error(), "truncated") {
		t.Errorf("Expected truncated error, got: %v", err)
	}
}
