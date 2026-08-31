package cloudmodel

import "time"

// * To avoid circular dependencies, the following structs are copied from the cb-tumblebug framework.
// TODO: When the cb-tumblebug framework is updated, we should synchronize these structs.
// * Version: CB-Tumblebug latest main (commit: a9a104734c622ff44d5ad0698d6df9a4b12a74d0)
// * Synchronized: 2026-08-31

// IID is a struct to handle Identifier Information of CSP resources from CB-Spider.
type IID struct {
	NameId   string `json:"NameId"`
	SystemId string `json:"SystemId"`
}

// K8sClusterStatus represents status of K8s cluster
type K8sClusterStatus string

const (
	K8sClusterCreating K8sClusterStatus = "Creating"
	K8sClusterActive   K8sClusterStatus = "Active"
	K8sClusterInactive K8sClusterStatus = "Inactive"
	K8sClusterUpdating K8sClusterStatus = "Updating"
	K8sClusterDeleting K8sClusterStatus = "Deleting"
)

// K8sNodeGroupStatus represents status of K8s node group
type K8sNodeGroupStatus string

const (
	K8sNodeGroupCreating K8sNodeGroupStatus = "Creating"
	K8sNodeGroupActive   K8sNodeGroupStatus = "Active"
	K8sNodeGroupInactive K8sNodeGroupStatus = "Inactive"
	K8sNodeGroupUpdating K8sNodeGroupStatus = "Updating"
	K8sNodeGroupDeleting K8sNodeGroupStatus = "Deleting"
)

type SpiderClusterStatus string

const (
	SpiderClusterCreating SpiderClusterStatus = "Creating"
	SpiderClusterActive   SpiderClusterStatus = "Active"
	SpiderClusterInactive SpiderClusterStatus = "Inactive"
	SpiderClusterUpdating SpiderClusterStatus = "Updating"
	SpiderClusterDeleting SpiderClusterStatus = "Deleting"
)

type SpiderNodeGroupStatus string

const (
	SpiderNodeGroupCreating SpiderNodeGroupStatus = "Creating"
	SpiderNodeGroupActive   SpiderNodeGroupStatus = "Active"
	SpiderNodeGroupInactive SpiderNodeGroupStatus = "Inactive"
	SpiderNodeGroupUpdating SpiderNodeGroupStatus = "Updating"
	SpiderNodeGroupDeleting SpiderNodeGroupStatus = "Deleting"
)

// K8sClusterReq is a struct to handle 'Create K8sCluster' request toward CB-Tumblebug.
type K8sClusterReq struct {
	//Namespace      string `json:"namespace" validate:"required" example:"default"`
	ConnectionName string `json:"connectionName" validate:"required" example:"alibaba-ap-northeast-2"`
	Description    string `json:"description" example:"My K8sCluster"`

	// (1) K8sCluster Info
	Name    string `json:"name" validate:"required" example:"k8scluster01"`
	Version string `json:"version" example:"1.30.1-aliyun.1"`

	// (2) Network Info
	VNetId           string   `json:"vNetId" validate:"required" example:"vpc-01"`
	SubnetIds        []string `json:"subnetIds" validate:"required" example:"subnet-01"`
	SecurityGroupIds []string `json:"securityGroupIds" validate:"required" example:"sg-01"`

	// (3) NodeGroupInfo List
	K8sNodeGroupList []K8sNodeGroupReq `json:"k8sNodeGroupList"`

	// Fields for "Register existing K8sCluster" feature
	// @description CspResourceId is required to register a k8s cluster from CSP (option=register)
	CspResourceId string `json:"cspResourceId" example:"required when option is register"`

	// Label is for describing the object by keywords
	Label map[string]string `json:"label"`

	// SystemLabel is for describing the k8scluster in a keyword (any string can be used) for special System purpose
	SystemLabel string `json:"systemLabel" example:"" default:""`
}

// K8sNodeGroupReq is a struct to handle requests related to K8sNodeGroup toward CB-Tumblebug.
type K8sNodeGroupReq struct {
	Name         string `json:"name" example:"k8sng01"`
	ImageId      string `json:"imageId" example:"image-01"`
	SpecId       string `json:"specId" example:"spec-01"`
	RootDiskType string `json:"rootDiskType" example:"cloud_essd" enum:"default, TYPE1, ..."` // "", "default", "TYPE1", AWS: ["standard", "gp2", "gp3"], Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"], GCP: ["pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"], ALIBABA: ["cloud_efficiency", "cloud", "cloud_essd"], TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
	RootDiskSize int    `json:"rootDiskSize" example:"40"`                                    // Root disk size in GB. 0 = use CSP default.
	SshKeyId     string `json:"sshKeyId" example:"sshkey-01"`

	// autoscale config.
	OnAutoScaling   string `json:"onAutoScaling" example:"true"`
	DesiredNodeSize int    `json:"desiredNodeSize" example:"1"`
	MinNodeSize     int    `json:"minNodeSize" example:"1"`
	MaxNodeSize     int    `json:"maxNodeSize" example:"3"`

	// Label is for describing the object by keywords
	Label map[string]string `json:"label"`

	Description string `json:"description" example:"Description"`
}

// K8sClusterDynamicReq is struct for requirements to create K8sCluster dynamically (with default resource option)
type K8sClusterDynamicReq struct {
	// K8sCluster name if it is not empty. Optional when used with namePrefix in multi-cluster creation.
	Name string `json:"name" example:"k8scluster01"`

	// K8s Clsuter version
	Version string `json:"version,omitempty" example:"1.29"`

	// Label is for describing the object by keywords
	Label map[string]string `json:"label,omitempty"`

	Description string `json:"description,omitempty" example:"Description"`

	// NodeGroup name if it is not empty
	NodeGroupName string `json:"nodeGroupName,omitempty" example:"k8sng01"`

	// SpecId is field for id of a spec in common namespace
	SpecId string `json:"specId" validate:"required" example:"tencent+ap-seoul+S2.MEDIUM4"`

	// ImageId is field for id of a image in common namespace
	ImageId string `json:"imageId" validate:"required" example:"default, tencent+ap-seoul+ubuntu20.04"`

	RootDiskType string `json:"rootDiskType,omitempty" example:"default, TYPE1, ..." default:"default"` // "", "default", "TYPE1", AWS: ["standard", "gp2", "gp3"], Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"], GCP: ["pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"], ALIBABA: ["cloud_efficiency", "cloud", "cloud_essd"], TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
	RootDiskSize int    `json:"rootDiskSize,omitempty" example:"30"`                                    // Root disk size in GB. 0 = use CSP default.

	OnAutoScaling   string `json:"onAutoScaling,omitempty" default:"true" example:"true"`
	DesiredNodeSize int    `json:"desiredNodeSize,omitempty" example:"1"`
	MinNodeSize     int    `json:"minNodeSize,omitempty" example:"1"`
	MaxNodeSize     int    `json:"maxNodeSize,omitempty" example:"3"`

	// if ConnectionName is given, the VM tries to use associated credential.
	// if not, it will use predefined ConnectionName in Spec objects
	ConnectionName string `json:"connectionName,omitempty" default:"tencent-ap-seoul"`
}

// K8sNodeGroupDynamicReq is struct for requirements to create K8sNodeGroup dynamically (with default resource option)
type K8sNodeGroupDynamicReq struct {
	// K8sNodeGroup name if it is not empty.
	Name string `json:"name" validate:"required" example:"k8sng01"`

	// Label is for describing the object by keywords
	Label map[string]string `json:"label,omitempty"`

	Description string `json:"description,omitempty" example:"Description"`

	// SpecId is field for id of a spec in common namespace
	SpecId string `json:"specId" validate:"required" example:"tencent+ap-seoul+S2.MEDIUM4"`

	// ImageId is field for id of a image in common namespace
	ImageId string `json:"imageId" validate:"required" example:"default, tencent+ap-seoul+ubuntu20.04"`

	RootDiskType string `json:"rootDiskType,omitempty" example:"default, TYPE1, ..." default:"default"` // "", "default", "TYPE1", AWS: ["standard", "gp2", "gp3"], Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"], GCP: ["pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"], ALIBABA: ["cloud_efficiency", "cloud", "cloud_essd"], TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
	RootDiskSize int    `json:"rootDiskSize,omitempty" example:"30"`                                    // Root disk size in GB. 0 = use CSP default.

	OnAutoScaling   string `json:"onAutoScaling,omitempty" default:"true" example:"true"`
	DesiredNodeSize int    `json:"desiredNodeSize,omitempty" example:"1"`
	MinNodeSize     int    `json:"minNodeSize,omitempty" example:"1"`
	MaxNodeSize     int    `json:"maxNodeSize,omitempty" example:"3"`
}

// K8sMultiClusterDynamicReq is a wrapper struct for creating multiple K8sClusters in parallel
type K8sMultiClusterDynamicReq struct {
	NamePrefix string                 `json:"namePrefix" example:"across"`
	Clusters   []K8sClusterDynamicReq `json:"clusters" validate:"required,dive"`
}

// K8sMultiClusterInfo is a wrapper struct for multiple K8sCluster creation results
type K8sMultiClusterInfo struct {
	Clusters       []K8sClusterInfo       `json:"clusters"`
	FailedClusters []K8sClusterFailedInfo `json:"failedClusters,omitempty"`
}

// K8sClusterFailedInfo contains information about a failed cluster creation attempt
type K8sClusterFailedInfo struct {
	Name           string `json:"name" example:"k8s-cluster-01"`
	ConnectionName string `json:"connectionName,omitempty" example:"aws-ap-northeast-2"`
	SpecId         string `json:"specId,omitempty" example:"aws+ap-northeast-2+t3.medium"`
	Error          string `json:"error" example:"failed to create cluster: resource quota exceeded"`
}

// SetK8sNodeGroupAutoscalingReq is a struct to handle 'Set K8sNodeGroup's Autoscaling' request toward CB-Tumblebug.
type SetK8sNodeGroupAutoscalingReq struct {
	OnAutoScaling string `json:"onAutoScaling" example:"true"`
}

// SetK8sNodeGroupAutoscalingRes is a struct to handle 'Set K8sNodeGroup's Autoscaling' response from CB-Tumblebug.
type SetK8sNodeGroupAutoscalingRes struct {
	Result string `json:"result" example:"true"`
}

// ChangeK8sNodeGroupAutoscaleSizeReq is a struct to handle 'Change K8sNodeGroup's Autoscale Size' request toward CB-Tumblebug.
type ChangeK8sNodeGroupAutoscaleSizeReq struct {
	DesiredNodeSize int `json:"desiredNodeSize" example:"1"`
	MinNodeSize     int `json:"minNodeSize" example:"1"`
	MaxNodeSize     int `json:"maxNodeSize" example:"3"`
}

// ChangeK8sNodeGroupAutoscaleSizeRes is a struct to handle 'Change K8sNodeGroup's Autoscale Size' response from CB-Tumblebug.
type ChangeK8sNodeGroupAutoscaleSizeRes struct {
	K8sNodeGroupInfo
}

// UpgradeK8sClusterReq is a struct to handle 'Upgrade K8sCluster' request toward CB-Tumblebug.
type UpgradeK8sClusterReq struct {
	Version string `json:"version" example:"1.30.1-alyun.1"`
}

// K8sClusterInfo is a struct that represents TB K8sCluster object.
type K8sClusterInfo struct {
	// ResourceType is the type of the resource
	ResourceType string `json:"resourceType"`

	// Id is unique identifier for the object, same as Name
	Id string `json:"id" example:"k8scluster01"`

	// Uid is universally unique identifier for the object, used for labelSelector
	Uid string `json:"uid,omitempty" example:"wef12awefadf1221edcf"`

	// Name is human-readable string to represent the object
	Name           string `json:"name" example:"k8scluster01"`
	ConnectionName string `json:"connectionName" example:"alibaba-ap-northeast-2"`

	// ConnectionConfig shows connection info to cloud service provider
	ConnectionConfig ConnConfig `json:"connectionConfig"`

	Description string `json:"description" example:"My K8sCluster"`

	// Latest system message such as error message
	SystemMessage string `json:"systemMessage" example:"Failed because ..." default:""` // systeam-given string message

	// Label is for describing the object by keywords
	Label map[string]string `json:"label"`

	// SystemLabel is for describing the Resource in a keyword (any string can be used) for special System purpose
	SystemLabel string `json:"systemLabel" example:"Managed by CB-Tumblebug" default:""`

	// Version is for kubernetes version
	Version string `json:"version" example:"1.30.1"` // Kubernetes Version, ex) 1.30.1

	// Network is for describing network information about the cluster
	Network K8sClusterNetworkInfo `json:"network"`

	// K8sNodeGroupList is for describing network information about the cluster
	K8sNodeGroupList []K8sNodeGroupInfo `json:"k8sNodeGroupList"`
	AccessInfo       K8sAccessInfo      `json:"accessInfo"`
	Addons           K8sAddonsInfo      `json:"addons"`

	Status K8sClusterStatus `json:"status" example:"Active"` // Creating, Active, Inactive, Updating, Deleting

	CreatedTime  time.Time  `json:"createdTime" example:"1970-01-01T00:00:00.00Z"`
	KeyValueList []KeyValue `json:"keyValueList"`

	// CspResourceName is name assigned to the CSP resource. This name is internally used to handle the resource.
	CspResourceName string `json:"cspResourceName,omitempty" example:"we12fawefadf1221edcf"`

	// CspResourceId is resource identifier managed by CSP
	CspResourceId string `json:"cspResourceId,omitempty" example:"csp-06eb41e14121c550a"`

	SpiderViewK8sClusterDetail SpiderClusterInfo `json:"spiderViewK8sClusterDetail"`
}

// K8sClusterNetworkInfo is a struct to handle K8sCluster Network information from the CB-Tumblebug's REST API response
type K8sClusterNetworkInfo struct {
	VNetId           string   `json:"vNetId" example:"vpc-01"`
	SubnetIds        []string `json:"subnetIds" example:"subnet-01"`
	SecurityGroupIds []string `json:"securityGroupIds" example:"sg-01"`

	// ---

	KeyValueList []KeyValue `json:"keyValueList"`
}

// K8sNodeGroupInfo is a struct to handle K8sCluster's Node Group information from the CB-Tumblebug's REST API response
type K8sNodeGroupInfo struct {
	// Id is unique identifier for the object
	Id string `json:"id" example:"aws-ap-southeast-1"`

	// Name is human-readable string to represent the object
	Name string `json:"name" example:"aws-ap-southeast-1"`

	ImageId         string             `json:"imageId"`
	SpecId          string             `json:"specId"`
	RootDiskType    string             `json:"rootDiskType"`
	RootDiskSize    int                `json:"rootDiskSize"`
	SshKeyId        string             `json:"sshKeyId"`
	OnAutoScaling   bool               `json:"onAutoScaling"`
	DesiredNodeSize int                `json:"desiredNodeSize"`
	MinNodeSize     int                `json:"minNodeSize"`
	MaxNodeSize     int                `json:"maxNodeSize"`
	Status          K8sNodeGroupStatus `json:"status" example:"Active"` // Creating, Active, Inactive, Updating, Deleting
	K8sNodes        []K8sNodeInfo      `json:"k8sNodes"`
	KeyValueList    []KeyValue         `json:"keyValueList"`

	// IsInitialNodeGroup indicates whether this node group was created during cluster creation.
	// For some CSPs (Alibaba ACK, Tencent TKE), this node group cannot be deleted independently;
	// it is automatically removed when the cluster is deleted.
	IsInitialNodeGroup bool `json:"isInitialNodeGroup,omitempty"`

	// CspResourceName is name assigned to the CSP resource. This name is internally used to handle the resource.
	CspResourceName string `json:"cspResourceName,omitempty" example:"we12fawefadf1221edcf"`

	// CspResourceId is resource identifier managed by CSP
	CspResourceId string `json:"cspResourceId,omitempty" example:"csp-06eb41e14121c550a"`

	SpiderViewK8sNodeGroupDetail SpiderNodeGroupInfo `json:"spiderViewK8sNodeGroupDetail"`
}

// K8sNodeInfo is a struct to handle K8sCluster's Node information
type K8sNodeInfo struct {
	// CspResourceName is name assigned to the CSP resource. This name is internally used to handle the resource.
	CspResourceName string `json:"cspResourceName,omitempty" example:"we12fawefadf1221edcf"`

	// CspResourceId is resource identifier managed by CSP
	CspResourceId string `json:"cspResourceId,omitempty" example:"csp-06eb41e14121c550a"`
}

// K8sAccessInfo is a struct to handle K8sCluster Access information from the CB-Tumblebug's REST API response
type K8sAccessInfo struct {
	Endpoint   string `json:"endpoint" example:"http://1.2.3.4:6443"`
	Kubeconfig string `json:"kubeconfig" example:"apiVersion: v1\nclusters:\n- cluster:\n certificate-authority-data: LS0..."`
}

// K8sAddonsInfo is a struct to handle K8sCluster Addons information from the CB-Tumblebug's REST API response
type K8sAddonsInfo struct {
	KeyValueList []KeyValue `json:"keyValueList"`
}

// SpiderClusterInfo is a struct to handle Cluster information from the CB-Spider's REST API response
type SpiderClusterInfo struct {
	IId IID // {NameId, SystemId}

	Version string // Kubernetes Version, ex) 1.23.3
	Network SpiderNetworkInfo

	// ---

	NodeGroupList []SpiderNodeGroupInfo
	AccessInfo    SpiderAccessInfo
	Addons        SpiderAddonsInfo

	Status SpiderClusterStatus

	CreatedTime  time.Time
	KeyValueList []KeyValue
}

// SpiderNetworkInfo is a struct to handle Cluster Network information from the CB-Spider's REST API response
type SpiderNetworkInfo struct {
	VpcIID            IID // {NameId, SystemId}
	SubnetIIDs        []IID
	SecurityGroupIIDs []IID

	// ---

	KeyValueList []KeyValue
}

// SpiderNodeGroupInfo is a struct to handle Cluster Node Group information from the CB-Spider's REST API response
type SpiderNodeGroupInfo struct {
	IId IID `json:"IId" validate:"required"` // {NameId, SystemId}

	// VM config.
	ImageIID     IID    `json:"ImageIID" validate:"required"`
	VMSpecName   string `json:"VMSpecName" validate:"required" example:"t3.medium"`
	RootDiskType string `json:"RootDiskType,omitempty" validate:"omitempty"`              // "SSD(gp2)", "Premium SSD", ...
	RootDiskSize string `json:"RootDiskSize,omitempty" validate:"omitempty" example:"50"` // "", "default", "50", "1000" (GB)
	KeyPairIID   IID    `json:"KeyPairIID" validate:"required"`

	// Scaling config.
	OnAutoScaling   bool `json:"OnAutoScaling" validate:"required" example:"true"`
	DesiredNodeSize int  `json:"DesiredNodeSize" validate:"required" example:"2"`
	MinNodeSize     int  `json:"MinNodeSize" validate:"required" example:"1"`
	MaxNodeSize     int  `json:"MaxNodeSize" validate:"required" example:"3"`

	// ---

	Status SpiderNodeGroupStatus `json:"Status" validate:"required" example:"Active"`

	Nodes []IID `json:"Nodes,omitempty" validate:"omitempty"`

	KeyValueList []KeyValue `json:"KeyValueList,omitempty" validate:"omitempty"`
}

// SpiderAccessInfo is a struct to handle Cluster Access information from the CB-Spider's REST API response
type SpiderAccessInfo struct {
	Endpoint   string // ex) https://1.2.3.4:6443
	Kubeconfig string
}

// SpiderAddonsInfo is a struct to handle Cluster Addons information from the CB-Spider's REST API response
type SpiderAddonsInfo struct {
	KeyValueList []KeyValue
}

// ExecCredential mirrors the Kubernetes ExecCredential format (client.authentication.k8s.io/v1).
type ExecCredential struct {
	ApiVersion string               `json:"apiVersion"`
	Kind       string               `json:"kind"`
	Status     ExecCredentialStatus `json:"status"`
}

// ExecCredentialStatus holds credentials for the transport to use.
type ExecCredentialStatus struct {
	Token string `json:"token"`
}

// K8sClusterTokenResponse is the response struct for the K8sCluster token API.
type K8sClusterTokenResponse struct {
	ExecCredential ExecCredential `json:"execCredential"`
}

// K8sClusterKubeconfigResponse is the response struct for the K8sCluster kubeconfig API.
type K8sClusterKubeconfigResponse struct {
	Kubeconfig string `json:"kubeconfig" example:"apiVersion: v1\nkind: Config\n..."`
}
