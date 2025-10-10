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

// Package common is to include common methods for managing multi-cloud infra
package common

type KeyValue struct {
	Key   string
	Value string
}

type IdList struct {
	IdList []string `json:"output"`
}

// SystemReady is global variable for checking SystemReady status
var SystemReady bool

var TumblebugRestUrl string
var AutocontrolDurationMs string
var err error

const (
	StrTumblebugRestUrl           string = "TUMBLEBUG_REST_URL"
	StrAutocontrolDurationMs      string = "AUTOCONTROL_DURATION_MS"
	CbStoreKeyNotFoundErrorString string = "key not found"
	StrAdd                        string = "add"
	StrDelete                     string = "delete"
	StrSSHKey                     string = "sshKey"
	StrImage                      string = "image"
	StrCustomImage                string = "customImage"
	StrSecurityGroup              string = "securityGroup"
	StrSpec                       string = "spec"
	StrVNet                       string = "vNet"
	StrSubnet                     string = "subnet"
	StrDataDisk                   string = "dataDisk"
	StrNLB                        string = "nlb"
	StrVM                         string = "vm"
	StrMCI                        string = "mci"
	StrDefaultResourceName        string = "-systemdefault-"

	// SystemCommonNs is const for SystemCommon NameSpace ID
	SystemCommonNs string = "system-purpose-common-ns"
)

var StartTime string

// Spider 2020-03-30 https://github.com/cloud-barista/cb-spider/blob/master/cloud-control-manager/cloud-driver/interfaces/resources/IId.go
type IID struct {
	NameId   string // NameID by user
	SystemId string // SystemID by CloudOS
}

type SpiderConnectionName struct {
	ConnectionName string `json:"ConnectionName"`
}

var DefaulNamespaceId = "mig01"

// NsReq is struct for namespace creation
type NsReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// NsInfo is struct for namespace information
type NsInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Response structure for RestGetAllNs
type RestGetAllNsResponse struct {
	//Name string     `json:"name"`
	Ns []NsInfo `json:"ns"`
}
