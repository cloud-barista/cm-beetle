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
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
)

func TestSortByProximityWithCost_VendorTieBreak(t *testing.T) {
	tests := []struct {
		name            string
		vcpus           uint32
		memory          uint32
		csp             string
		sourceCpuVendor string
		specs           []cloudmodel.SpecInfo
		wantFirst       string // CspSpecName expected to rank first
	}{
		{
			name:            "azure: intel wins over cheaper AMD at equal fit (the reported bug, fixed)",
			vcpus:           2,
			memory:          8,
			csp:             "azure",
			sourceCpuVendor: "intel",
			specs: []cloudmodel.SpecInfo{
				{CspSpecName: "Standard_D2as_v5", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.10},
				{CspSpecName: "Standard_D2s_v5", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.12},
			},
			wantFirst: "Standard_D2s_v5",
		},
		{
			name:            "aws: intel wins over cheaper AMD at equal fit",
			vcpus:           2,
			memory:          8,
			csp:             "aws",
			sourceCpuVendor: "intel",
			specs: []cloudmodel.SpecInfo{
				{CspSpecName: "m6a.large", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.086},
				{CspSpecName: "m6i.large", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.096},
			},
			wantFirst: "m6i.large",
		},
		{
			name:            "gcp: intel wins over cheaper AMD at equal fit",
			vcpus:           2,
			memory:          8,
			csp:             "gcp",
			sourceCpuVendor: "intel",
			specs: []cloudmodel.SpecInfo{
				{CspSpecName: "n2d-standard-2", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.09},
				{CspSpecName: "n2-standard-2", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.10},
			},
			wantFirst: "n2-standard-2",
		},
		{
			name:            "alibaba: intel wins over cheaper AMD at equal fit, via Details PhysicalProcessorModel",
			vcpus:           2,
			memory:          8,
			csp:             "alibaba",
			sourceCpuVendor: "intel",
			specs: []cloudmodel.SpecInfo{
				{CspSpecName: "ecs.g6a.large", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.09, Details: []cloudmodel.KeyValue{
					{Key: "PhysicalProcessorModel", Value: "AMD EPYC(Milan) 7T83"},
				}},
				{CspSpecName: "ecs.g6.large", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.10, Details: []cloudmodel.KeyValue{
					{Key: "PhysicalProcessorModel", Value: "Intel Xeon(Ice Lake) Platinum 8369B"},
				}},
			},
			wantFirst: "ecs.g6.large",
		},
		{
			name:            "regression: unknown source vendor keeps today's cost-only order",
			vcpus:           2,
			memory:          8,
			csp:             "azure",
			sourceCpuVendor: "",
			specs: []cloudmodel.SpecInfo{
				{CspSpecName: "Standard_D2as_v5", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.10},
				{CspSpecName: "Standard_D2s_v5", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.12},
			},
			wantFirst: "Standard_D2as_v5",
		},
		{
			name:            "regression: CSP without a vendor detector keeps today's cost-only order",
			vcpus:           2,
			memory:          8,
			csp:             "ibm",
			sourceCpuVendor: "intel",
			specs: []cloudmodel.SpecInfo{
				{CspSpecName: "Standard_D2as_v5", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.10},
				{CspSpecName: "Standard_D2s_v5", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.12},
			},
			wantFirst: "Standard_D2as_v5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			specs := append([]cloudmodel.SpecInfo(nil), tt.specs...)
			sortByProximityWithCost(specs, tt.vcpus, tt.memory, tt.csp, tt.sourceCpuVendor)
			if specs[0].CspSpecName != tt.wantFirst {
				t.Errorf("sortByProximityWithCost() first = %q, want %q", specs[0].CspSpecName, tt.wantFirst)
			}
		})
	}
}

// When the source vendor is known, vendor match dominates resource proximity by design (per
// explicit user direction): a worse-fitting spec matching the source vendor ranks ahead of an
// exact-fitting spec of a different vendor.
func TestSortByProximityWithCost_VendorDominatesResourceProximityWhenKnown(t *testing.T) {
	exactFitWrongVendor := cloudmodel.SpecInfo{CspSpecName: "Standard_D2as_v5", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.10} // AMD, exact fit
	worseFitRightVendor := cloudmodel.SpecInfo{CspSpecName: "Standard_D4s_v5", VCPU: 4, MemoryGiB: 16, CostPerHour: 0.02} // Intel, worse fit, far cheaper

	specs := []cloudmodel.SpecInfo{exactFitWrongVendor, worseFitRightVendor}
	sortByProximityWithCost(specs, 2, 8, "azure", "intel")

	if specs[0].CspSpecName != worseFitRightVendor.CspSpecName {
		t.Errorf("sortByProximityWithCost() first = %q, want %q (vendor match must dominate proximity when vendor is known)",
			specs[0].CspSpecName, worseFitRightVendor.CspSpecName)
	}
}

// When the source vendor is unknown, ranking must fall back to proximity-then-cost, unaffected by
// vendor - vendorMatch always ties in this case regardless of its position in the criteria list.
func TestSortByProximityWithCost_ProximityDecidesWhenVendorUnknown(t *testing.T) {
	exactFit := cloudmodel.SpecInfo{CspSpecName: "Standard_D2as_v5", VCPU: 2, MemoryGiB: 8, CostPerHour: 0.10}
	worseFit := cloudmodel.SpecInfo{CspSpecName: "Standard_D4s_v5", VCPU: 4, MemoryGiB: 16, CostPerHour: 0.02}

	specs := []cloudmodel.SpecInfo{worseFit, exactFit}
	sortByProximityWithCost(specs, 2, 8, "azure", "")

	if specs[0].CspSpecName != exactFit.CspSpecName {
		t.Errorf("sortByProximityWithCost() first = %q, want %q (proximity must decide when source vendor is unknown)",
			specs[0].CspSpecName, exactFit.CspSpecName)
	}
}

// The vendor tie-break must apply identically in the compute-intensive and memory-intensive
// branches, not just the general-purpose default.
func TestSortByProximityWithCost_AllMachineTypeBranches(t *testing.T) {
	tests := []struct {
		name      string
		vcpus     uint32
		memory    uint32
		amdSpec   cloudmodel.SpecInfo
		intelSpec cloudmodel.SpecInfo
	}{
		{
			name:      "compute-intensive branch (memory:vCPU ratio <= 3)",
			vcpus:     8,
			memory:    16,
			amdSpec:   cloudmodel.SpecInfo{CspSpecName: "Standard_F8as_v6", VCPU: 8, MemoryGiB: 16, CostPerHour: 0.20},
			intelSpec: cloudmodel.SpecInfo{CspSpecName: "Standard_F8s_v2", VCPU: 8, MemoryGiB: 16, CostPerHour: 0.24},
		},
		{
			name:      "memory-intensive branch (memory:vCPU ratio >= 7)",
			vcpus:     2,
			memory:    16,
			amdSpec:   cloudmodel.SpecInfo{CspSpecName: "Standard_E2as_v5", VCPU: 2, MemoryGiB: 16, CostPerHour: 0.15},
			intelSpec: cloudmodel.SpecInfo{CspSpecName: "Standard_E2s_v5", VCPU: 2, MemoryGiB: 16, CostPerHour: 0.18},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			specs := []cloudmodel.SpecInfo{tt.amdSpec, tt.intelSpec}
			sortByProximityWithCost(specs, tt.vcpus, tt.memory, "azure", "intel")
			if specs[0].CspSpecName != tt.intelSpec.CspSpecName {
				t.Errorf("sortByProximityWithCost() first = %q, want %q (intel should win the vendor tie-break)",
					specs[0].CspSpecName, tt.intelSpec.CspSpecName)
			}
		})
	}
}

func TestExtractCpuVendor(t *testing.T) {
	tests := []struct {
		name   string
		vendor string
		want   string
	}{
		{"GenuineIntel", "GenuineIntel", "intel"},
		{"AuthenticAMD", "AuthenticAMD", "amd"},
		{"empty", "", ""},
		{"unrecognized vendor string", "some-unknown-vendor", ""},
		{"full model string containing Intel", "Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz", "intel"},
		{"full model string containing AMD", "AMD EPYC 7302 16-Core Processor", "amd"},
		{"ARM vendor string", "ARM Limited", "arm"},
		{"Ampere model string", "Ampere(R) Altra(R) Processor", "arm"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractCpuVendor(tt.vendor); got != tt.want {
				t.Errorf("extractCpuVendor(%q) = %q, want %q", tt.vendor, got, tt.want)
			}
		})
	}
}

func TestVendorRank(t *testing.T) {
	intelSpec := cloudmodel.SpecInfo{CspSpecName: "intel-spec"}
	amdSpec := cloudmodel.SpecInfo{CspSpecName: "amd-spec"}
	unknownSpec := cloudmodel.SpecInfo{CspSpecName: "unknown-spec"}
	vendorByName := map[string]string{"intel-spec": "intel", "amd-spec": "amd", "unknown-spec": ""}

	tests := []struct {
		name         string
		sourceVendor string
		vendorByName map[string]string
		spec         cloudmodel.SpecInfo
		want         int
	}{
		{"matching vendor ranks 0", "intel", vendorByName, intelSpec, 0},
		{"mismatched vendor ranks 1", "intel", vendorByName, amdSpec, 1},
		{"unknown spec vendor ranks 1 (no special trust)", "intel", vendorByName, unknownSpec, 1},
		{"empty source vendor is a no-op (always 0)", "", vendorByName, amdSpec, 0},
		{"nil vendorByName is a no-op (always 0)", "intel", nil, amdSpec, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vendorRank(tt.sourceVendor, tt.vendorByName, tt.spec); got != tt.want {
				t.Errorf("vendorRank(%q, ..., %+v) = %d, want %d", tt.sourceVendor, tt.spec, got, tt.want)
			}
		})
	}
}
