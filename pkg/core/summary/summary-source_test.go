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

package summary

import "testing"

func TestIsMainInterface(t *testing.T) {
	cases := []struct {
		name string
		want bool
	}{
		{"lo", true},
		{"eth0", true},
		{"eno1", true},
		{"ens5", true},   // systemd PCI-slot predictable name (e.g. AWS/KVM guests)
		{"enp0s3", true}, // systemd PCI-ID predictable name
		{"enx001122334455", true},
		{"br-ex", true},
		{"tap0", false},
		{"veth1234", false},
		{"docker0", false},
		{"", false},
	}

	for _, c := range cases {
		if got := isMainInterface(c.name); got != c.want {
			t.Errorf("isMainInterface(%q) = %v, want %v", c.name, got, c.want)
		}
	}
}
