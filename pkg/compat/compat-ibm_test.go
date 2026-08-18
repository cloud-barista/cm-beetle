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

package compat

import "testing"

func TestGetIbmCpuVendor(t *testing.T) {
	tests := []struct {
		name        string
		cspSpecName string
		wantVendor  string
	}{
		{"hx4a is AMD", "hx4a-32x64", "amd"},
		{"hx4da is AMD with local disk", "hx4da-176x1424", "amd"},
		{"bx2 is Intel", "bx2-2x8", "intel"},
		{"bx3d is Intel", "bx3d-2x8", "intel"},
		{"cx2 is Intel", "cx2-2x4", "intel"},
		{"mx2 is Intel", "mx2-2x16", "intel"},
		{"vx2d is Intel", "vx2d-2x56", "intel"},
		{"ux2d is Intel", "ux2d-2x112", "intel"},
		{"ox2 is Intel", "ox2-2x16", "intel"},
		{"bxf Flex is unclassified (mixed Intel/AMD host pool)", "bxf-2x8", ""},
		{"cxf Flex is unclassified", "cxf-2x4", ""},
		{"mxf Flex is unclassified", "mxf-2x16", ""},
		{"gx2 GPU family is unclassified", "gx2-8x64x1v100", ""},
		{"uppercase input", "HX4A-32X64", "amd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIbmCpuVendor(tt.cspSpecName); got != tt.wantVendor {
				t.Errorf("getIbmCpuVendor(%q) = %q, want %q", tt.cspSpecName, got, tt.wantVendor)
			}
		})
	}
}
