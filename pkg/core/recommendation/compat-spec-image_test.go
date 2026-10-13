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

func TestFilterTencentVmSpecs(t *testing.T) {
	specs := []cloudmodel.SpecInfo{
		{CspSpecName: "BF1.LARGE16", VCPU: 4, MemoryGiB: 16},
		{
			CspSpecName: "SA2.LARGE16", VCPU: 4, MemoryGiB: 16,
			Details: []cloudmodel.KeyValue{
				{Key: "Status", Value: "SOLD_OUT"},
				{Key: "StatusCategory", Value: "WithoutStock"},
			},
		},
		{
			CspSpecName: "S2.LARGE16", VCPU: 4, MemoryGiB: 16,
			Details: []cloudmodel.KeyValue{
				{Key: "Status", Value: "SELL"},
				{Key: "StatusCategory", Value: "EnoughStock"},
			},
		},
		{
			CspSpecName: "C2.MEDIUM4", VCPU: 2, MemoryGiB: 4,
			Details: []cloudmodel.KeyValue{
				{Key: "Status", Value: "SELL"},
				{Key: "StatusCategory", Value: "EnoughStock"},
			},
		},
		{
			CspSpecName: "S6.LARGE16", VCPU: 4, MemoryGiB: 16,
			Details: []cloudmodel.KeyValue{
				{Key: "Status", Value: "SELL"},
				{Key: "StatusCategory", Value: "EnoughStock"},
			},
		},
		{CspSpecName: "BF1.MEDIUM2", VCPU: 2, MemoryGiB: 2},
		{
			CspSpecName: "SA5.MEDIUM2", VCPU: 2, MemoryGiB: 2,
			Details: []cloudmodel.KeyValue{
				{Key: "Status", Value: "SELL"},
				{Key: "StatusCategory", Value: "NormalStock"},
			},
		},
	}

	filtered := filterTencentVmSpecs(specs)

	if len(filtered) != 2 {
		t.Fatalf("expected 2 modern in-stock specs (S6.LARGE16, SA5.MEDIUM2), got %d: %+v", len(filtered), filtered)
	}

	expected := map[string]bool{"S6.LARGE16": true, "SA5.MEDIUM2": true}
	for _, s := range filtered {
		if !expected[s.CspSpecName] {
			t.Errorf("unexpected spec found in filtered results: %s", s.CspSpecName)
		}
	}
}

func TestPreFilterByCsp_Tencent(t *testing.T) {
	specs := []cloudmodel.SpecInfo{
		{CspSpecName: "BF1.LARGE16", VCPU: 4, MemoryGiB: 16},
		{
			CspSpecName: "S2.LARGE16", VCPU: 4, MemoryGiB: 16,
			Details: []cloudmodel.KeyValue{
				{Key: "Status", Value: "SELL"},
				{Key: "StatusCategory", Value: "EnoughStock"},
			},
		},
		{
			CspSpecName: "S6.LARGE16", VCPU: 4, MemoryGiB: 16,
			Details: []cloudmodel.KeyValue{
				{Key: "Status", Value: "SELL"},
				{Key: "StatusCategory", Value: "EnoughStock"},
			},
		},
	}
	images := []cloudmodel.ImageInfo{
		{CspImageName: "img-7rotv4ux"},
	}

	filteredSpecs, filteredImages := preFilterByCsp("tencent", specs, images)

	if len(filteredSpecs) != 1 || filteredSpecs[0].CspSpecName != "S6.LARGE16" {
		t.Errorf("preFilterByCsp for tencent failed: expected [S6.LARGE16], got %+v", filteredSpecs)
	}

	if len(filteredImages) != 1 {
		t.Errorf("preFilterByCsp for tencent should not modify images: got %d", len(filteredImages))
	}
}
