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

func TestGetNcpCpuVendor(t *testing.T) {
	tests := []struct {
		name        string
		cspSpecName string
		wantVendor  string
	}{
		{"c2-g3 is Intel", "c2-g3", "intel"},
		{"s2-g3 is Intel", "s2-g3", "intel"},
		{"m2-g2 is Intel", "m2-g2", "intel"},
		{"ci8-g3 is Intel", "ci8-g3", "intel"},
		{"mi1-g2 is Intel", "mi1-g2", "intel"},
		{"c2-g3a is AMD", "c2-g3a", "amd"},
		{"s4-g3a is AMD", "s4-g3a", "amd"},
		{"ci8-g3a is AMD", "ci8-g3a", "amd"},
		{"unrecognized family is unclassified", "g2-g3a", ""},
		{"unrecognized code format is unclassified", "SVR.VSVR.STAND.C002.M008", ""},
		{"uppercase input", "C2-G3A", "amd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNcpCpuVendor(tt.cspSpecName); got != tt.wantVendor {
				t.Errorf("getNcpCpuVendor(%q) = %q, want %q", tt.cspSpecName, got, tt.wantVendor)
			}
		})
	}
}
