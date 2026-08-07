package cloudmodel

// The following structs are used temporarily.

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
	RootDiskType string `json:"rootDiskType" example:"cloud_essd" enum:"default, TYPE1, ..."` // "", "default", "TYPE1", AWS: ["standard", "gp2", "gp3"], Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"], GCP: ["pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"], ALIBABA: ["cloud_efficiency", "cloud", "cloud_ssd"], TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
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
