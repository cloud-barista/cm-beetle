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
	"reflect"
	"strings"
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
)

func TestDeriveNetworkRequirements(t *testing.T) {
	nodeGroups := []cloudmodel.CreateNodeGroupReq{
		{Name: "ng1", VNetId: "v1", ConnectionName: "c1", SubnetId: "s1"},
		// Same VNet as ng1: dedup, subnet dedup, and connectionName must NOT
		// be overwritten by a later nodegroup once already set.
		{Name: "ng2", VNetId: "v1", ConnectionName: "c-ignored", SubnetId: "s1"},
		// Same VNet again: a genuinely new subnet gets appended.
		{Name: "ng3", VNetId: "v1", SubnetId: "s2"},
		// No VNetId: skipped entirely.
		{Name: "ng4", VNetId: ""},
		// A second, distinct VNet with no connection/subnet info.
		{Name: "ng5", VNetId: "v2"},
	}

	got := DeriveNetworkRequirements(nodeGroups)
	want := []NetworkRequirement{
		{VNetId: "v1", SubnetIds: []string{"s1", "s2"}, ConnectionName: "c1"},
		{VNetId: "v2"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("DeriveNetworkRequirements = %+v, want %+v", got, want)
	}
}

func TestCheckNetworkAvailability(t *testing.T) {
	tests := []struct {
		name            string
		req             NetworkRequirement
		creationReq     cloudmodel.VNetReq
		wantNeedsCreate bool
		wantIssueCode   string // empty means no issue expected
	}{
		{
			name:            "exists and all required subnets present",
			req:             NetworkRequirement{VNetId: "vnet-exists", SubnetIds: []string{"subnet-a"}, ConnectionName: "testcsp-region1"},
			wantNeedsCreate: false,
		},
		{
			name:          "exists under a different CSP/region connection",
			req:           NetworkRequirement{VNetId: "vnet-mismatch", ConnectionName: "testcsp-region1"},
			wantIssueCode: CodeConnectionMismatch,
		},
		{
			name:            "exists but missing a required subnet, with fallback creation data",
			req:             NetworkRequirement{VNetId: "vnet-exists", SubnetIds: []string{"subnet-a", "subnet-missing"}, ConnectionName: "testcsp-region1"},
			creationReq:     cloudmodel.VNetReq{CidrBlock: "10.0.0.0/16"},
			wantNeedsCreate: true,
		},
		{
			name:            "does not exist, has fallback creation data",
			req:             NetworkRequirement{VNetId: "vnet-does-not-exist"},
			creationReq:     cloudmodel.VNetReq{CidrBlock: "10.0.0.0/16"},
			wantNeedsCreate: true,
		},
		{
			name:          "does not exist, no fallback creation data",
			req:           NetworkRequirement{VNetId: "vnet-does-not-exist"},
			wantIssueCode: CodeResourceNotAvailable,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			needsCreate, issue := CheckNetworkAvailability(testNsId, tc.req, tc.creationReq)
			assertAvailability(t, needsCreate, issue, tc.wantNeedsCreate, tc.wantIssueCode)
		})
	}
}

func TestCheckSshKeyAvailability(t *testing.T) {
	tests := []struct {
		name            string
		req             SshKeyRequirement
		creationReq     cloudmodel.SshKeyReq
		wantNeedsCreate bool
		wantIssueCode   string
	}{
		{
			name:            "exists under the requested connection",
			req:             SshKeyRequirement{SshKeyId: "sshkey-exists", ConnectionName: "testcsp-region1"},
			wantNeedsCreate: false,
		},
		{
			name:          "exists under a different CSP/region connection",
			req:           SshKeyRequirement{SshKeyId: "sshkey-mismatch", ConnectionName: "testcsp-region1"},
			wantIssueCode: CodeConnectionMismatch,
		},
		{
			name:            "does not exist, has fallback creation data",
			req:             SshKeyRequirement{SshKeyId: "sshkey-does-not-exist"},
			creationReq:     cloudmodel.SshKeyReq{Name: "sshkey-does-not-exist"},
			wantNeedsCreate: true,
		},
		{
			name:          "does not exist, no fallback creation data",
			req:           SshKeyRequirement{SshKeyId: "sshkey-does-not-exist"},
			wantIssueCode: CodeResourceNotAvailable,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			needsCreate, issue := CheckSshKeyAvailability(testNsId, tc.req, tc.creationReq)
			assertAvailability(t, needsCreate, issue, tc.wantNeedsCreate, tc.wantIssueCode)
		})
	}
}

func TestDeriveSecurityGroupRequirements(t *testing.T) {
	nodeGroups := []cloudmodel.CreateNodeGroupReq{
		{Name: "ng1", VNetId: "v1", ConnectionName: "c1", SecurityGroupIds: []string{"sg1", "sg2"}},
		// "sg1" repeated (dedup) and an empty id (skipped); the ignored
		// connection/VNet here prove the *first* sighting of each id wins.
		{Name: "ng2", VNetId: "v1-ignored", ConnectionName: "c-ignored", SecurityGroupIds: []string{"sg1", ""}},
		{Name: "ng3", VNetId: "v3", ConnectionName: "c3", SecurityGroupIds: []string{"sg3"}},
	}

	got := DeriveSecurityGroupRequirements(nodeGroups)
	want := []SecurityGroupRequirement{
		{SecurityGroupId: "sg1", VNetId: "v1", ConnectionName: "c1"},
		{SecurityGroupId: "sg2", VNetId: "v1", ConnectionName: "c1"},
		{SecurityGroupId: "sg3", VNetId: "v3", ConnectionName: "c3"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("DeriveSecurityGroupRequirements = %+v, want %+v", got, want)
	}
}

func TestCheckSecurityGroupAvailability(t *testing.T) {
	tests := []struct {
		name            string
		req             SecurityGroupRequirement
		creationReqList []cloudmodel.SecurityGroupReq
		wantNeedsCreate bool
		wantIssueCode   string
	}{
		{
			name:            "exists under the requested connection",
			req:             SecurityGroupRequirement{SecurityGroupId: "sg-exists", ConnectionName: "testcsp-region1"},
			wantNeedsCreate: false,
		},
		{
			name:          "exists under a different CSP/region connection",
			req:           SecurityGroupRequirement{SecurityGroupId: "sg-mismatch", ConnectionName: "testcsp-region1"},
			wantIssueCode: CodeConnectionMismatch,
		},
		{
			name: "does not exist, matching creation entry carries its own connection/vNet",
			req:  SecurityGroupRequirement{SecurityGroupId: "sg-new"},
			creationReqList: []cloudmodel.SecurityGroupReq{
				{Name: "sg-new", ConnectionName: "testcsp-region1", VNetId: "vnet-new"},
			},
			wantNeedsCreate: true,
		},
		{
			name: "does not exist, matching creation entry is missing connection/vNet, falls back to the requirement",
			req:  SecurityGroupRequirement{SecurityGroupId: "sg-new", ConnectionName: "testcsp-region1", VNetId: "vnet-new"},
			creationReqList: []cloudmodel.SecurityGroupReq{
				{Name: "sg-new"},
			},
			wantNeedsCreate: true,
		},
		{
			name:            "does not exist, no matching creation entry, falls back to the requirement",
			req:             SecurityGroupRequirement{SecurityGroupId: "sg-new", ConnectionName: "testcsp-region1", VNetId: "vnet-new"},
			wantNeedsCreate: true,
		},
		{
			name:          "does not exist, no matching creation entry, and the requirement itself lacks connection/vNet",
			req:           SecurityGroupRequirement{SecurityGroupId: "sg-new"},
			wantIssueCode: CodeResourceNotAvailable,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			needsCreate, issue := CheckSecurityGroupAvailability(testNsId, tc.req, tc.creationReqList)
			assertAvailability(t, needsCreate, issue, tc.wantNeedsCreate, tc.wantIssueCode)
		})
	}
}

// assertAvailability is the common assertion shared by the Check*Availability
// table tests: either no issue is expected (and needsCreate must match), or a
// specific issue code is expected.
func assertAvailability(t *testing.T, needsCreate bool, issue *ValidationIssue, wantNeedsCreate bool, wantIssueCode string) {
	t.Helper()
	if wantIssueCode == "" {
		if issue != nil {
			t.Fatalf("unexpected issue: %+v", issue)
		}
		if needsCreate != wantNeedsCreate {
			t.Fatalf("needsCreate = %v, want %v", needsCreate, wantNeedsCreate)
		}
		return
	}
	if issue == nil || issue.Code != wantIssueCode {
		t.Fatalf("issue = %+v, want code %q", issue, wantIssueCode)
	}
}

func TestValidationResultErr(t *testing.T) {
	valid := ValidationResult{Valid: true}
	if err := valid.Err(); err != nil {
		t.Fatalf("Err() = %v, want nil for a valid result", err)
	}

	invalid := newResult([]ValidationIssue{
		{Severity: SeverityError, Message: "boom"},
		{Severity: SeverityWarning, Message: "just a warning"},
	})
	err := invalid.Err()
	if err == nil {
		t.Fatal("Err() = nil, want a non-nil error for an invalid result")
	}
	msg := err.Error()
	if !strings.Contains(msg, "boom") {
		t.Fatalf("Err() message %q does not mention the error-severity issue", msg)
	}
	if strings.Contains(msg, "just a warning") {
		t.Fatalf("Err() message %q should not mention warning-severity issues", msg)
	}
}
