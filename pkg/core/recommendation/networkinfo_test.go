/*
Copyright 2024 The Cloud-Barista Authors.
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

package recommendation

import (
	"net"
	"strings"
	"testing"
)

func TestGetCspNetworkConfig(t *testing.T) {
	expectedCSPs := []string{"ncp", "aws", "azure", "gcp", "alibaba", "ibm", "tencent", "nhn"}

	for _, csp := range expectedCSPs {
		cfg, ok := GetCspNetworkConfig(csp)
		if !ok {
			t.Errorf("GetCspNetworkConfig(%q) not found", csp)
			continue
		}
		if cfg.Description == "" {
			t.Errorf("GetCspNetworkConfig(%q) description is empty", csp)
		}
	}
}

func TestAdjustVNetCidrForCsp_NCP(t *testing.T) {
	// NCP VPC requires prefix-length min 16, max 28

	// Case 1: /24 subnet expanded with 4 networks should become /22, which is >= 16 (valid)
	got, _, err := AdjustVNetCidrForCsp("ncp", "10.1.0.0/24", 4)
	if err != nil {
		t.Fatalf("unexpected error for 10.1.0.0/24: %v", err)
	}
	_, ipNet, _ := net.ParseCIDR(got)
	prefix, _ := ipNet.Mask.Size()
	if prefix < 16 || prefix > 28 {
		t.Errorf("AdjustVNetCidrForCsp(ncp) got prefix /%d, want within [16, 28]", prefix)
	}

	// Case 2: /16 supernet expanded should clamp to /16 instead of overflowing to /14
	got16, note16, err := AdjustVNetCidrForCsp("ncp", "10.0.0.0/16", 4)
	if err != nil {
		t.Fatalf("unexpected error for 10.0.0.0/16: %v", err)
	}
	_, ipNet16, _ := net.ParseCIDR(got16)
	prefix16, _ := ipNet16.Mask.Size()
	if prefix16 != 16 {
		t.Errorf("AdjustVNetCidrForCsp(ncp, 10.0.0.0/16) got /%d, want /16 clamp", prefix16)
	}
	if !strings.Contains(note16, "Clamped") {
		t.Errorf("expected note to contain 'Clamped', got %q", note16)
	}

	// Case 3: Raw supernet already wider than /16 (e.g., /14) should clamp to /16 and provide Note
	gotClamped, noteClamped, err := AdjustVNetCidrForCsp("ncp", "10.0.0.0/14", 1)
	if err != nil {
		t.Fatalf("AdjustVNetCidrForCsp(ncp, /14) unexpected error: %v", err)
	}
	_, ipNetClamped, _ := net.ParseCIDR(gotClamped)
	prefixClamped, _ := ipNetClamped.Mask.Size()
	if prefixClamped != 16 {
		t.Errorf("AdjustVNetCidrForCsp(ncp, /14) got prefix /%d, want /16", prefixClamped)
	}
	if !strings.Contains(noteClamped, "Clamped from /14 to /16") {
		t.Errorf("expected note with clamp explanation, got %q", noteClamped)
	}
}

func TestAdjustVNetCidrForCsp_AllCSPs(t *testing.T) {
	testCases := []struct {
		csp         string
		rawCidr     string
		networks    int
		wantMinPref int
		wantMaxPref int
	}{
		{"aws", "10.0.0.0/24", 4, 16, 28},
		{"azure", "10.0.0.0/24", 4, 8, 29},
		{"alibaba", "10.0.0.0/24", 4, 8, 28},
		{"ibm", "10.0.0.0/24", 4, 9, 28},
		{"tencent", "10.0.0.0/24", 4, 12, 28},
		{"nhn", "10.0.0.0/24", 4, 8, 24},
		{"gcp", "10.0.0.0/24", 4, 1, 32}, // GCP has no VPC CIDR constraint
	}

	for _, tc := range testCases {
		t.Run(tc.csp, func(t *testing.T) {
			got, _, err := AdjustVNetCidrForCsp(tc.csp, tc.rawCidr, tc.networks)
			if err != nil {
				t.Fatalf("AdjustVNetCidrForCsp(%s) unexpected error: %v", tc.csp, err)
			}
			_, ipNet, err := net.ParseCIDR(got)
			if err != nil {
				t.Fatalf("invalid CIDR returned: %s", got)
			}
			prefix, _ := ipNet.Mask.Size()
			if prefix < tc.wantMinPref || prefix > tc.wantMaxPref {
				t.Errorf("AdjustVNetCidrForCsp(%s) got prefix /%d, want within [%d, %d]",
					tc.csp, prefix, tc.wantMinPref, tc.wantMaxPref)
			}
		})
	}
}

func TestAdjustSubnetCidrForCsp(t *testing.T) {
	// NCP: subnet prefix range [16, 28]

	// Case 1: Subnet /29 exceeds max prefix 28 -> should be adjusted to /28
	adjusted, err := AdjustSubnetCidrForCsp("ncp", "10.0.0.0/16", "10.0.1.0/29")
	if err != nil {
		t.Fatalf("unexpected error adjusting /29 subnet for NCP: %v", err)
	}
	_, subNet, _ := net.ParseCIDR(adjusted)
	subPrefix, _ := subNet.Mask.Size()
	if subPrefix != 28 {
		t.Errorf("AdjustSubnetCidrForCsp(ncp, 10.0.1.0/29) got /%d, want /28", subPrefix)
	}

	// Case 2: Subnet /15 is larger than NCP min prefix 16 -> should return error
	_, err = AdjustSubnetCidrForCsp("ncp", "10.0.0.0/16", "10.0.0.0/15")
	if err == nil {
		t.Error("AdjustSubnetCidrForCsp(ncp, /15) expected error for exceeding max subnet size, got nil")
	}

	// Case 3: Subnet not in VNet range -> should return error
	_, err = AdjustSubnetCidrForCsp("ncp", "10.0.0.0/16", "192.168.1.0/24")
	if err == nil {
		t.Error("AdjustSubnetCidrForCsp outside vNet expected error, got nil")
	}
}

func TestBuildRecommendedVNet_NcpConstraints(t *testing.T) {
	cidrs := []string{
		"10.0.1.0/24",
		"10.0.2.0/29", // Needs adjustment to /28 in NCP
	}

	vnet, err := buildRecommendedVNet("ncp", "KR-1", "10.0.0.0/8", cidrs, 4)
	if err != nil {
		t.Fatalf("buildRecommendedVNet(ncp) unexpected error: %v", err)
	}

	if vnet == nil {
		t.Fatal("buildRecommendedVNet(ncp) returned nil vnet")
	}

	_, vNetNet, err := net.ParseCIDR(vnet.CidrBlock)
	if err != nil {
		t.Fatalf("invalid recommended VNet CIDR: %s", vnet.CidrBlock)
	}
	vNetPrefix, _ := vNetNet.Mask.Size()

	// Verify NCP VNet prefix constraint: 16 <= prefix <= 28
	if vNetPrefix < 16 || vNetPrefix > 28 {
		t.Errorf("VNet CIDR %s prefix /%d outside NCP range [16, 28]", vnet.CidrBlock, vNetPrefix)
	}

	// Verify Subnet prefix constraints and containment
	for _, sub := range vnet.SubnetInfoList {
		_, subNet, err := net.ParseCIDR(sub.IPv4_CIDR)
		if err != nil {
			t.Errorf("invalid subnet CIDR: %s", sub.IPv4_CIDR)
			continue
		}
		subPrefix, _ := subNet.Mask.Size()
		if subPrefix < 16 || subPrefix > 28 {
			t.Errorf("Subnet CIDR %s prefix /%d outside NCP range [16, 28]", sub.IPv4_CIDR, subPrefix)
		}
		if !vNetNet.Contains(subNet.IP) {
			t.Errorf("Subnet %s is not contained in VNet %s", sub.IPv4_CIDR, vnet.CidrBlock)
		}
	}
}

func TestBuildRecommendedVNet_ClampedSupernetNote(t *testing.T) {
	// Source with /16 subnet that would expand beyond /16 without clamp
	cidrs := []string{
		"10.1.0.0/16",
	}

	vnet, err := buildRecommendedVNet("ncp", "KR-1", "10.0.0.0/8", cidrs, 4)
	if err != nil {
		t.Fatalf("buildRecommendedVNet(ncp) unexpected error: %v", err)
	}

	if vnet == nil {
		t.Fatal("buildRecommendedVNet(ncp) returned nil vnet")
	}

	// Description must contain Clamped Note
	if !strings.Contains(vnet.Description, "Note: Clamped") {
		t.Errorf("expected vnet Description to contain 'Note: Clamped', got %q", vnet.Description)
	}
}
