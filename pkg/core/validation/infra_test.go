/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validation

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog"
)

// testNsId is the namespace used by every test in this file.
const testNsId = "ns-test"

// fixture is a canned Tumblebug response served for one exact request path.
type fixture struct {
	status int
	body   any
}

// fixtures holds every canned response the fake Tumblebug server can return.
// Any path not present here 404s, which is exactly the "resource does not
// exist" case ValidateTargetInfra needs to exercise - so only "found" cases
// need an explicit entry.
var fixtures = map[string]fixture{
	"/ns/" + testNsId + "/resources/vNet/vnet-exists": {http.StatusOK, tbmodel.VNetInfo{
		Id: "vnet-exists", Name: "vnet-exists", ConnectionName: "testcsp-region1",
		SubnetInfoList: []tbmodel.SubnetInfo{{Name: "subnet-a"}},
	}},
	"/ns/" + testNsId + "/resources/vNet/vnet-mismatch": {http.StatusOK, tbmodel.VNetInfo{
		Id: "vnet-mismatch", Name: "vnet-mismatch", ConnectionName: "othercsp-region2",
		SubnetInfoList: []tbmodel.SubnetInfo{{Name: "subnet-a"}},
	}},
	"/ns/" + testNsId + "/resources/sshKey/sshkey-exists": {http.StatusOK, tbmodel.SshKeyInfo{
		Id: "sshkey-exists", Name: "sshkey-exists", ConnectionName: "testcsp-region1",
	}},
	"/ns/" + testNsId + "/resources/sshKey/sshkey-mismatch": {http.StatusOK, tbmodel.SshKeyInfo{
		Id: "sshkey-mismatch", Name: "sshkey-mismatch", ConnectionName: "othercsp-region2",
	}},
	"/ns/" + testNsId + "/resources/securityGroup/sg-exists": {http.StatusOK, tbmodel.SecurityGroupInfo{
		Id: "sg-exists", Name: "sg-exists", ConnectionName: "testcsp-region1",
	}},
	"/ns/" + testNsId + "/resources/securityGroup/sg-mismatch": {http.StatusOK, tbmodel.SecurityGroupInfo{
		Id: "sg-mismatch", Name: "sg-mismatch", ConnectionName: "othercsp-region2",
	}},
	"/ns/" + testNsId + "/infra/infra-exists": {http.StatusOK, tbmodel.InfraInfo{
		Id: "infra-exists", Name: "infra-exists",
	}},
	// Spec/Image lookups are always issued against the "system" namespace.
	"/ns/system/resources/spec/spec-x86": {http.StatusOK, tbmodel.SpecInfo{
		Id: "spec-x86", Architecture: "x86_64",
	}},
	"/ns/system/resources/image/testcsp+image-x86": {http.StatusOK, tbmodel.ImageInfo{
		Id: "testcsp+image-x86", OSArchitecture: tbmodel.X86_64,
	}},
	"/ns/system/resources/image/testcsp+image-arm": {http.StatusOK, tbmodel.ImageInfo{
		Id: "testcsp+image-arm", OSArchitecture: tbmodel.ARM64,
	}},
}

// TestMain wires tbclient to a fake Tumblebug server before any test runs.
// tbclient.NewSession() (used throughout this package) always goes through a
// process-global client protected by sync.Once, so it must be initialized
// exactly once, here, rather than per test.
//
// Silences zerolog's global level: tbclient logs every request (including
// the expected 404s "does not exist" checks rely on) at Info/Debug/Warn/Error,
// which otherwise buries the -v test output. Test failures are unaffected -
// t.Fatalf/t.Errorf print through Go's own test runner, not zerolog.
//
// Uses an httptest TLS server (rather than a plain HTTP one) so resty's own
// "Using Basic Auth in HTTP mode is not secure" warning - printed via its own
// private *log.Logger, so it can't be silenced through zerolog or the stdlib
// log package - never triggers in the first place. InsecureSkipVerify trusts
// the server's self-signed test certificate.
func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.Disabled)

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f, ok := fixtures[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
			return
		}
		w.WriteHeader(f.status)
		_ = json.NewEncoder(w).Encode(f.body)
	}))
	defer server.Close()

	tbclient.Init(tbclient.ApiConfig{
		RestUrl:  server.URL,
		Username: "test",
		Password: "test",
		HttpTransport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // trusts only the local httptest TLS server's self-signed cert
		},
	})

	os.Exit(m.Run())
}

// baseTarget returns a fresh, fully valid single-NodeGroup RecommendedInfra:
// one VNet/Subnet, one SSH key, one security group, all named "*-new" (i.e.
// not present in `fixtures`, so they read as "does not exist yet"), and a
// spec/image pair that is architecture-compatible under CSP "testcsp" (not
// one of compat.CheckCompatibility's special-cased CSPs, so only the common
// architecture check applies).
func baseTarget() *cloudmodel.RecommendedInfra {
	return &cloudmodel.RecommendedInfra{
		TargetInfra: cloudmodel.InfraReq{
			Name: "infra-new",
			NodeGroups: []cloudmodel.CreateNodeGroupReq{
				{
					Name:             "ng1",
					ConnectionName:   "testcsp-region1",
					SpecId:           "spec-x86",
					ImageId:          "image-x86",
					VNetId:           "vnet-new",
					SubnetId:         "subnet-a",
					SecurityGroupIds: []string{"sg-new"},
					SshKeyId:         "sshkey-new",
				},
			},
		},
		TargetVNet: cloudmodel.VNetReq{
			Name:           "vnet-new",
			ConnectionName: "testcsp-region1",
			CidrBlock:      "10.0.0.0/16",
			SubnetInfoList: []cloudmodel.SubnetReq{
				{Name: "subnet-a", IPv4_CIDR: "10.0.1.0/24"},
			},
		},
		TargetSshKey: cloudmodel.SshKeyReq{
			Name:           "sshkey-new",
			ConnectionName: "testcsp-region1",
		},
		TargetSecurityGroupList: []cloudmodel.SecurityGroupReq{
			{Name: "sg-new", ConnectionName: "testcsp-region1", VNetId: "vnet-new"},
		},
	}
}

// targetOption mutates a baseTarget() to set up one test scenario.
type targetOption func(*cloudmodel.RecommendedInfra)

// withVNetName points the sole NodeGroup, the sole SecurityGroup, and
// TargetVNet at the same name, keeping referential integrity intact while
// swapping which VNet fixture (or non-existent name) the test exercises.
func withVNetName(name string) targetOption {
	return func(t *cloudmodel.RecommendedInfra) {
		t.TargetVNet.Name = name
		t.TargetInfra.NodeGroups[0].VNetId = name
		t.TargetSecurityGroupList[0].VNetId = name
	}
}

func withSshKeyName(name string) targetOption {
	return func(t *cloudmodel.RecommendedInfra) {
		t.TargetSshKey.Name = name
		t.TargetInfra.NodeGroups[0].SshKeyId = name
	}
}

func withSecurityGroupName(name string) targetOption {
	return func(t *cloudmodel.RecommendedInfra) {
		t.TargetSecurityGroupList[0].Name = name
		t.TargetInfra.NodeGroups[0].SecurityGroupIds = []string{name}
	}
}

func TestValidateTargetInfra(t *testing.T) {
	tests := []struct {
		name        string
		useExisting bool
		opts        []targetOption
		wantValid   bool
		wantCode    string // checked only when wantValid is false
	}{
		{
			name:        "fresh creation: all new resources with compatible spec/image is valid",
			useExisting: false,
			wantValid:   true,
		},
		{
			name:        "fresh creation: vNet already exists",
			useExisting: false,
			opts:        []targetOption{withVNetName("vnet-exists")},
			wantValid:   false,
			wantCode:    CodeResourceAlreadyExists,
		},
		{
			name:        "reuse existing: all resources already exist under the requested connection",
			useExisting: true,
			opts: []targetOption{
				withVNetName("vnet-exists"),
				withSshKeyName("sshkey-exists"),
				withSecurityGroupName("sg-exists"),
			},
			wantValid: true,
		},
		{
			name:        "reuse existing: vNet exists under a different CSP/region connection",
			useExisting: true,
			opts: []targetOption{
				withVNetName("vnet-mismatch"),
				withSshKeyName("sshkey-exists"),
				withSecurityGroupName("sg-exists"),
			},
			wantValid: false,
			wantCode:  CodeConnectionMismatch,
		},
		{
			name:        "reuse existing: SSH key exists under a different CSP/region connection",
			useExisting: true,
			opts: []targetOption{
				withVNetName("vnet-exists"),
				withSshKeyName("sshkey-mismatch"),
				withSecurityGroupName("sg-exists"),
			},
			wantValid: false,
			wantCode:  CodeConnectionMismatch,
		},
		{
			name:        "reuse existing: security group exists under a different CSP/region connection",
			useExisting: true,
			opts: []targetOption{
				withVNetName("vnet-exists"),
				withSshKeyName("sshkey-exists"),
				withSecurityGroupName("sg-mismatch"),
			},
			wantValid: false,
			wantCode:  CodeConnectionMismatch,
		},
		{
			name:        "reuse existing: missing vNet with no fallback creation data (cidrBlock)",
			useExisting: true,
			opts: []targetOption{
				withVNetName("vnet-does-not-exist"),
				withSshKeyName("sshkey-exists"),
				withSecurityGroupName("sg-exists"),
				func(t *cloudmodel.RecommendedInfra) { t.TargetVNet.CidrBlock = "" },
			},
			wantValid: false,
			wantCode:  CodeResourceNotAvailable,
		},
		{
			name:        "spec and image architecture mismatch",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.NodeGroups[0].ImageId = "image-arm" },
			},
			wantValid: false,
			wantCode:  CodeSpecImageIncompatible,
		},
		{
			name:        "nodegroup missing connectionName",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.NodeGroups[0].ConnectionName = "" },
			},
			wantValid: false,
			wantCode:  CodeRequiredFieldMissing,
		},
		{
			name:        "Infra (MCI) name already exists",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.Name = "infra-exists" },
			},
			wantValid: false,
			wantCode:  CodeResourceAlreadyExists,
		},
		{
			name:        "fresh creation: targetVNet.name is required",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetVNet.Name = "" },
			},
			wantValid: false,
			wantCode:  CodeRequiredFieldMissing,
		},
		{
			name:        "fresh creation: targetSshKey.name is required",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetSshKey.Name = "" },
			},
			wantValid: false,
			wantCode:  CodeRequiredFieldMissing,
		},
		{
			name:        "fresh creation: security group name is required",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetSecurityGroupList[0].Name = "" },
			},
			wantValid: false,
			wantCode:  CodeRequiredFieldMissing,
		},
		{
			name:        "fresh creation: SSH key already exists",
			useExisting: false,
			opts:        []targetOption{withSshKeyName("sshkey-exists")},
			wantValid:   false,
			wantCode:    CodeResourceAlreadyExists,
		},
		{
			name:        "fresh creation: security group already exists",
			useExisting: false,
			opts:        []targetOption{withSecurityGroupName("sg-exists")},
			wantValid:   false,
			wantCode:    CodeResourceAlreadyExists,
		},
		{
			name:        "reuse existing: nodegroup vNetId is required",
			useExisting: true,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.NodeGroups[0].VNetId = "" },
			},
			wantValid: false,
			wantCode:  CodeRequiredFieldMissing,
		},
		{
			name:        "reuse existing: nodegroup sshKeyId is required",
			useExisting: true,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.NodeGroups[0].SshKeyId = "" },
			},
			wantValid: false,
			wantCode:  CodeRequiredFieldMissing,
		},
		{
			name:        "reuse existing: nodegroup securityGroupIds is required",
			useExisting: true,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.NodeGroups[0].SecurityGroupIds = nil },
			},
			wantValid: false,
			wantCode:  CodeRequiredFieldMissing,
		},
		{
			name:        "nodegroup missing specId",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.NodeGroups[0].SpecId = "" },
			},
			wantValid: false,
			wantCode:  CodeRequiredFieldMissing,
		},
		{
			name:        "nodegroup missing imageId",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.NodeGroups[0].ImageId = "" },
			},
			wantValid: false,
			wantCode:  CodeRequiredFieldMissing,
		},
		{
			name:        "nodegroup connectionName is not in csp-region format",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.NodeGroups[0].ConnectionName = "nodash" },
			},
			wantValid: false,
			wantCode:  CodeInvalidConnectionName,
		},
		{
			name:        "nodegroup spec lookup fails (spec not found)",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.NodeGroups[0].SpecId = "spec-missing" },
			},
			wantValid: false,
			wantCode:  CodeSpecOrImageLookupFailed,
		},
		{
			name:        "nodegroup image lookup fails (image not found)",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.NodeGroups[0].ImageId = "image-missing" },
			},
			wantValid: false,
			wantCode:  CodeSpecOrImageLookupFailed,
		},
		{
			name:        "targetInfra.name is required",
			useExisting: false,
			opts: []targetOption{
				func(t *cloudmodel.RecommendedInfra) { t.TargetInfra.Name = "" },
			},
			wantValid: false,
			wantCode:  CodeRequiredFieldMissing,
		},
		{
			name:        "nil target model",
			useExisting: false,
			wantValid:   false,
			wantCode:    CodeRequiredFieldMissing,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var target *cloudmodel.RecommendedInfra
			if tc.name != "nil target model" {
				target = baseTarget()
				for _, opt := range tc.opts {
					opt(target)
				}
			}

			result := ValidateTargetInfra(testNsId, target, tc.useExisting)

			if result.Valid != tc.wantValid {
				t.Fatalf("Valid = %v, want %v; issues: %+v", result.Valid, tc.wantValid, result.Issues)
			}
			if !tc.wantValid {
				found := false
				for _, issue := range result.Issues {
					if issue.Code == tc.wantCode {
						found = true
						break
					}
				}
				if !found {
					t.Fatalf("expected an issue with code %q, got: %+v", tc.wantCode, result.Issues)
				}
			}
		})
	}
}
