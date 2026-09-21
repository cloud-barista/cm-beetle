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

// Package csp defines standard Cloud Service Provider constants and identifiers for CM-Beetle.
package csp

import "strings"

// Cloud Service Providers
const (
	Alibaba   = "alibaba"
	AWS       = "aws"
	Azure     = "azure"
	GCP       = "gcp"
	IBM       = "ibm"
	Tencent   = "tencent"
	NCP       = "ncp"
	NHN       = "nhn"
	KT        = "kt"
	OpenStack = "openstack"
)

// AllCSPs is the list of all Cloud Service Providers recognized in Cloud-Barista
var AllCSPs = []string{
	AWS, Azure, GCP, Alibaba, Tencent, IBM, OpenStack, NCP, NHN, KT,
}

var knownCSPs = func() map[string]bool {
	m := make(map[string]bool, len(AllCSPs))
	for _, c := range AllCSPs {
		m[c] = true
	}
	return m
}()

// IsKnown returns whether the given CSP name is a recognized Cloud Service Provider
func IsKnown(cspName string) bool {
	return knownCSPs[strings.ToLower(strings.TrimSpace(cspName))]
}
