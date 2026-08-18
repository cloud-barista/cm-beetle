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

func TestGetAzureCpuVendor(t *testing.T) {
	tests := []struct {
		name        string
		cspSpecName string
		wantVendor  string
	}{
		{"Dasv5 is AMD", "Standard_D2as_v5", "amd"},
		{"Dasv4 is AMD, case-insensitive", "standard_d2as_v4", "amd"},
		{"Dadsv5 is AMD", "Standard_D2ads_v5", "amd"},
		{"Dsv5 is Intel", "Standard_D2s_v5", "intel"},
		{"Ddsv5 is Intel (local disk, no AMD marker)", "Standard_D2ds_v5", "intel"},
		{"Easv5 is AMD", "Standard_E2as_v5", "amd"},
		{"Esv5 is Intel", "Standard_E2s_v5", "intel"},
		{"M-series is unclassified", "Standard_M128s", ""},
		{"classic A-series is unclassified", "Standard_A2_v2", ""},
		{"HBv3 (actually AMD HPC) is intentionally unclassified", "Standard_HB176rs_v3", ""},
		{"DCasv5 is AMD", "Standard_DC8as_v5", "amd"},
		{"DCadsv5 is AMD", "Standard_DC8ads_v5", "amd"},
		{"ECasv5 is AMD", "Standard_EC8as_v5", "amd"},
		{"DCsv3 is Intel", "Standard_DC8s_v3", "intel"},
		{"uppercase input", "STANDARD_D2AS_V5", "amd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAzureCpuVendor(tt.cspSpecName); got != tt.wantVendor {
				t.Errorf("getAzureCpuVendor(%q) = %q, want %q", tt.cspSpecName, got, tt.wantVendor)
			}
		})
	}
}
