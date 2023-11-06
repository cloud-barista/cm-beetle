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

// Package migration is to handle REST API for migration
package migration

import (
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

// [NOTE] a struct to test initially
// type Infrastructure struct {
// 	Network        string
// 	Disk           string
// 	Compute        string
// 	SecurityGroup  string
// 	VirtualMachine string
// }

// TbMcisDynamicReq is sturct for requirements to create MCIS dynamically (with default resource option)
type TbMcisDynamicReq struct {
	Name string `json:"name" validate:"required" example:"mcis01"`

	// InstallMonAgent Option for CB-Dragonfly agent installation ([yes/no] default:yes)
	InstallMonAgent string `json:"installMonAgent" example:"no" default:"yes" enums:"yes,no"` // yes or no

	// Label is for describing the mcis in a keyword (any string can be used)
	Label string `json:"label" example:"DynamicVM" default:""`

	// SystemLabel is for describing the mcis in a keyword (any string can be used) for special System purpose
	SystemLabel string `json:"systemLabel" example:"" default:""`

	Description string `json:"description" example:"Made in CB-TB"`

	Vm []TbVmDynamicReq `json:"vm" validate:"required"`
}

// TbVmDynamicReq is struct to get requirements to create a new server instance dynamically (with default resource option)
type TbVmDynamicReq struct {
	// VM name or subGroup name if is (not empty) && (> 0). If it is a group, actual VM name will be generated with -N postfix.
	Name string `json:"name" example:"g1-1"`

	// if subGroupSize is (not empty) && (> 0), subGroup will be gernetad. VMs will be created accordingly.
	SubGroupSize string `json:"subGroupSize" example:"3" default:""`

	Label string `json:"label" example:"DynamicVM"`

	Description string `json:"description" example:"Description"`

	// CommonSpec is field for id of a spec in common namespace
	CommonSpec string `json:"commonSpec" validate:"required" example:"aws-ap-northeast-2-t2-small"`
	// CommonImage is field for id of a image in common namespace
	CommonImage string `json:"commonImage" validate:"required" example:"ubuntu18.04"`

	RootDiskType string `json:"rootDiskType,omitempty" example:"default, TYPE1, ..."`  // "", "default", "TYPE1", AWS: ["standard", "gp2", "gp3"], Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"], GCP: ["pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"], ALIBABA: ["cloud_efficiency", "cloud", "cloud_essd"], TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
	RootDiskSize string `json:"rootDiskSize,omitempty" example:"default, 30, 42, ..."` // "default", Integer (GB): ["50", ..., "1000"]

	VmUserPassword string `json:"vmUserPassword default:""`
	// if ConnectionName is given, the VM tries to use associtated credential.
	// if not, it will use predefined ConnectionName in Spec objects
	ConnectionName string `json:"connectionName,omitempty" default:""`
}

// StatusCountInfo is struct to count the number of VMs in each status. ex: Running=4, Suspended=8.
type StatusCountInfo struct {

	// CountTotal is for Total VMs
	CountTotal int `json:"countTotal"`

	// CountCreating is for counting Creating
	CountCreating int `json:"countCreating"`

	// CountRunning is for counting Running
	CountRunning int `json:"countRunning"`

	// CountFailed is for counting Failed
	CountFailed int `json:"countFailed"`

	// CountSuspended is for counting Suspended
	CountSuspended int `json:"countSuspended"`

	// CountRebooting is for counting Rebooting
	CountRebooting int `json:"countRebooting"`

	// CountTerminated is for counting Terminated
	CountTerminated int `json:"countTerminated"`

	// CountSuspending is for counting Suspending
	CountSuspending int `json:"countSuspending"`

	// CountResuming is for counting Resuming
	CountResuming int `json:"countResuming"`

	// CountTerminating is for counting Terminating
	CountTerminating int `json:"countTerminating"`

	// CountUndefined is for counting Undefined
	CountUndefined int `json:"countUndefined"`
}

// RegionInfo is struct for region information
type RegionInfo struct {
	Region string
	Zone   string
}

type SpiderImageType string

// Ref: cb-spider/cloud-control-manager/cloud-driver/interfaces/resources/VMHandler.go
// SpiderVMInfo is struct from CB-Spider for VM information
type SpiderVMInfo struct {
	// Fields for request
	Name               string
	ImageName          string
	VPCName            string
	SubnetName         string
	SecurityGroupNames []string
	KeyPairName        string
	CSPid              string // VM ID given by CSP (required for registering VM)
	DataDiskNames      []string

	// Fields for both request and response
	VMSpecName   string // instance type or flavour, etc... ex) t2.micro or f1.micro
	VMUserId     string // ex) user1
	VMUserPasswd string
	RootDiskType string // "SSD(gp2)", "Premium SSD", ...
	RootDiskSize string // "default", "50", "1000" (GB)
	ImageType    SpiderImageType

	// Fields for response
	IId               common.IID // {NameId, SystemId}
	ImageIId          common.IID
	VpcIID            common.IID
	SubnetIID         common.IID   // AWS, ex) subnet-8c4a53e4
	SecurityGroupIIds []common.IID // AWS, ex) sg-0b7452563e1121bb6
	KeyPairIId        common.IID
	DataDiskIIDs      []common.IID
	StartTime         time.Time
	Region            RegionInfo //  ex) {us-east1, us-east1-c} or {ap-northeast-2}
	NetworkInterface  string     // ex) eth0
	PublicIP          string
	PublicDNS         string
	PrivateIP         string
	PrivateDNS        string
	RootDeviceName    string // "/dev/sda1", ...
	SSHAccessPoint    string
	KeyValueList      []common.KeyValue
}

// TbVmInfo is struct to define a server instance object
type TbVmInfo struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	IdByCSP string `json:"idByCSP"` // CSP managed ID or Name

	// defined if the VM is in a group
	SubGroupId string `json:"subGroupId"`

	Location common.GeoLocation `json:"location"`

	// Required by CB-Tumblebug
	Status       string `json:"status"`
	TargetStatus string `json:"targetStatus"`
	TargetAction string `json:"targetAction"`

	// Montoring agent status
	MonAgentStatus string `json:"monAgentStatus" example:"[installed, notInstalled, failed]"` // yes or no// installed, notInstalled, failed

	// NetworkAgent status
	NetworkAgentStatus string `json:"networkAgentStatus" example:"[notInstalled, installing, installed, failed]"` // notInstalled, installing, installed, failed

	// Latest system message such as error message
	SystemMessage string `json:"systemMessage" example:"Failed because ..." default:""` // systeam-given string message

	// Created time
	CreatedTime string `json:"createdTime" example:"2022-11-10 23:00:00" default:""`

	Label       string `json:"label"`
	Description string `json:"description"`

	Region         RegionInfo `json:"region"` // AWS, ex) {us-east1, us-east1-c} or {ap-northeast-2}
	PublicIP       string     `json:"publicIP"`
	SSHPort        string     `json:"sshPort"`
	PublicDNS      string     `json:"publicDNS"`
	PrivateIP      string     `json:"privateIP"`
	PrivateDNS     string     `json:"privateDNS"`
	RootDiskType   string     `json:"rootDiskType"`
	RootDiskSize   string     `json:"rootDiskSize"`
	RootDeviceName string     `json:"rootDeviceName"`

	ConnectionName   string            `json:"connectionName"`
	ConnectionConfig common.ConnConfig `json:"connectionConfig"`
	SpecId           string            `json:"specId"`
	ImageId          string            `json:"imageId"`
	VNetId           string            `json:"vNetId"`
	SubnetId         string            `json:"subnetId"`
	SecurityGroupIds []string          `json:"securityGroupIds"`
	DataDiskIds      []string          `json:"dataDiskIds"`
	SshKeyId         string            `json:"sshKeyId"`
	VmUserAccount    string            `json:"vmUserAccount,omitempty"`
	VmUserPassword   string            `json:"vmUserPassword,omitempty"`

	CspViewVmDetail SpiderVMInfo `json:"cspViewVmDetail,omitempty"`
}

// TbMcisInfo is struct for MCIS info
type TbMcisInfo struct {
	Id           string          `json:"id"`
	Name         string          `json:"name"`
	Status       string          `json:"status"`
	StatusCount  StatusCountInfo `json:"statusCount"`
	TargetStatus string          `json:"targetStatus"`
	TargetAction string          `json:"targetAction"`

	// InstallMonAgent Option for CB-Dragonfly agent installation ([yes/no] default:yes)
	InstallMonAgent string `json:"installMonAgent" example:"yes" default:"yes" enums:"yes,no"` // yes or no

	// ConfigureCloudAdaptiveNetwork is an option to configure Cloud Adaptive Network (CLADNet) ([yes/no] default:yes)
	ConfigureCloudAdaptiveNetwork string `json:"configureCloudAdaptiveNetwork" example:"yes" default:"no" enums:"yes,no"` // yes or no

	// Label is for describing the mcis in a keyword (any string can be used)
	Label string `json:"label" example:"User custom label"`

	// SystemLabel is for describing the mcis in a keyword (any string can be used) for special System purpose
	SystemLabel string `json:"systemLabel" example:"Managed by CB-Tumblebug" default:""`

	// Latest system message such as error message
	SystemMessage string `json:"systemMessage" example:"Failed because ..." default:""` // systeam-given string message

	PlacementAlgo string     `json:"placementAlgo,omitempty"`
	Description   string     `json:"description"`
	Vm            []TbVmInfo `json:"vm"`

	// List of IDs for new VMs. Return IDs if the VMs are newly added. This field should be used for return body only.
	NewVmList []string `json:"newVmList"`
}
