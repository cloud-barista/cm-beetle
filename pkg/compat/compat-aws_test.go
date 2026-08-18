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

func TestGetAwsCpuVendor(t *testing.T) {
	tests := []struct {
		name        string
		cspSpecName string
		wantVendor  string
	}{
		{"m6a is AMD", "m6a.large", "amd"},
		{"c6a is AMD", "c6a.xlarge", "amd"},
		{"r6a is AMD", "r6a.large", "amd"},
		{"t3a is AMD", "t3a.medium", "amd"},
		{"m5ad is AMD with local disk", "m5ad.large", "amd"},
		{"m6i is Intel", "m6i.large", "intel"},
		{"c6i is Intel", "c6i.xlarge", "intel"},
		{"m5 (no letter) is Intel", "m5.large", "intel"},
		{"c5 (no letter) is Intel", "c5.large", "intel"},
		{"t3 (no letter) is Intel", "t3.medium", "intel"},
		{"m5dn (disk+network, no vendor letter) is Intel", "m5dn.large", "intel"},
		{"m6g is Graviton/ARM, unclassified for amd/intel", "m6g.large", ""},
		{"t4g is Graviton/ARM, unclassified", "t4g.medium", ""},
		{"c6gn is Graviton/ARM, unclassified", "c6gn.large", ""},
		{"g5 (GPU family, excluded)", "g5.xlarge", ""},
		{"p3 (GPU family, excluded)", "p3.2xlarge", ""},
		{"uppercase input", "M6A.LARGE", "amd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAwsCpuVendor(tt.cspSpecName); got != tt.wantVendor {
				t.Errorf("getAwsCpuVendor(%q) = %q, want %q", tt.cspSpecName, got, tt.wantVendor)
			}
		})
	}
}
