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

func TestGetGcpCpuVendor(t *testing.T) {
	tests := []struct {
		name        string
		cspSpecName string
		wantVendor  string
	}{
		{"n2d is AMD", "n2d-standard-4", "amd"},
		{"c2d is AMD", "c2d-standard-4", "amd"},
		{"c3d is AMD", "c3d-standard-4", "amd"},
		{"t2d is AMD", "t2d-standard-4", "amd"},
		{"n2 is Intel", "n2-standard-4", "intel"},
		{"c2 is Intel", "c2-standard-4", "intel"},
		{"n1 is Intel", "n1-standard-1", "intel"},
		{"t2a is Ampere/ARM, unclassified", "t2a-standard-4", ""},
		{"e2 is explicitly unclassified (GCP auto-selects vendor)", "e2-medium", ""},
		{"a2 GPU family is unclassified", "a2-highgpu-1g", ""},
		{"uppercase input", "N2D-STANDARD-4", "amd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGcpCpuVendor(tt.cspSpecName); got != tt.wantVendor {
				t.Errorf("getGcpCpuVendor(%q) = %q, want %q", tt.cspSpecName, got, tt.wantVendor)
			}
		})
	}
}
