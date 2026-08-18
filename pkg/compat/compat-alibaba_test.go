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

import (
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
)

func TestGetAlibabaCpuVendor(t *testing.T) {
	tests := []struct {
		name       string
		spec       cloudmodel.SpecInfo
		wantVendor string
	}{
		{
			name: "Intel PhysicalProcessorModel",
			spec: cloudmodel.SpecInfo{Details: []cloudmodel.KeyValue{
				{Key: "PhysicalProcessorModel", Value: "Intel Xeon(Ice Lake) Platinum 8369B"},
			}},
			wantVendor: "intel",
		},
		{
			name: "AMD PhysicalProcessorModel",
			spec: cloudmodel.SpecInfo{Details: []cloudmodel.KeyValue{
				{Key: "PhysicalProcessorModel", Value: "AMD EPYC(Milan) 7T83"},
			}},
			wantVendor: "amd",
		},
		{
			name:       "missing PhysicalProcessorModel key",
			spec:       cloudmodel.SpecInfo{Details: []cloudmodel.KeyValue{{Key: "InstanceTypeFamily", Value: "ecs.g6"}}},
			wantVendor: "",
		},
		{
			name:       "no Details at all",
			spec:       cloudmodel.SpecInfo{},
			wantVendor: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAlibabaCpuVendor(tt.spec); got != tt.wantVendor {
				t.Errorf("getAlibabaCpuVendor(%+v) = %q, want %q", tt.spec, got, tt.wantVendor)
			}
		})
	}
}
