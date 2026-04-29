package cloudmodel

// RecommendedVmInfraModel represents the recommended virtual machine infrastructure model.
type RecommendedVmInfraModel struct {
	RecommendedVmInfraModel RecommendedVmInfra `json:"recommendedVmInfraModel" validate:"required"`
}

// RecommendedVmInfra represents the recommended virtual machine infrastructure information.
type RecommendedVmInfra struct {
	NameSeed                string             `json:"nameSeed"`
	Status                  string             `json:"status"`
	Description             string             `json:"description"`
	TargetCloud             CloudProperty      `json:"targetCloud"`
	TargetVmInfra           InfraReq           `json:"targetVmInfra"`
	TargetVNet              VNetReq            `json:"targetVNet"`
	TargetSshKey            SshKeyReq          `json:"targetSshKey"`
	TargetVmSpecList        []SpecInfo         `json:"targetVmSpecList"`
	TargetVmOsImageList     []ImageInfo        `json:"targetVmOsImageList"`
	TargetSecurityGroupList []SecurityGroupReq `json:"targetSecurityGroupList"`
}

// CloudProperty represents the cloud service provider (CSP) information.
type CloudProperty struct {
	Csp    string `json:"csp" example:"aws"`
	Region string `json:"region" example:"ap-northeast-2"`
}

// RecommendedVmInfraDynamic represents the recommended virtual machine infrastructure information.
type RecommendedVmInfraDynamic struct {
	Status        string        `json:"status"`
	Description   string        `json:"description"`
	TargetVmInfra InfraDynamicReq `json:"targetVmInfra"`
}

// RecommendedVmInfraDynamicList represents a list of recommended virtual machine infrastructure information.
type RecommendedVmInfraDynamicList struct {
	Description       string                      `json:"description"`
	Count             int                         `json:"count"`
	TargetVmInfraList []RecommendedVmInfraDynamic `json:"targetVmInfraList"`
}

// RecommendedVNet represents the recommended virtual network information.
// * May be mainly used this object
type RecommendedVNet struct {
	Status      string  `json:"status"`
	Description string  `json:"description"`
	TargetVNet  VNetReq `json:"targetVNet"`
}

// RecommendedVNetList represents a list of recommended virtual network information.
type RecommendedVNetList struct {
	Description    string            `json:"description"`
	Count          int               `json:"count"`
	TargetVNetList []RecommendedVNet `json:"targetVNetList"`
}

// RecommendedSecurityGroup represents the recommended security group information.
type RecommendedSecurityGroup struct {
	Status              string           `json:"status"`
	SourceServers       []string         `json:"sourceServers"`
	Description         string           `json:"description"`
	TargetSecurityGroup SecurityGroupReq `json:"targetSecurityGroup"`
}

// RecommendedSecurityGroupList represents a list of recommended security group information.
type RecommendedSecurityGroupList struct {
	Status                  string                     `json:"status"`
	Description             string                     `json:"description"`
	Count                   int                        `json:"count"`
	TargetSecurityGroupList []RecommendedSecurityGroup `json:"targetSecurityGroupList"`
}

// RecommendedVmSpec represents the recommended virtual machine specification information for a single server.
type RecommendedVmSpec struct {
	Status        string   `json:"status"`
	SourceServers []string `json:"sourceServers"`
	Description   string   `json:"description"`
	TargetVmSpec  SpecInfo `json:"targetVmSpec"`
}

// RecommendedVmSpecList represents a collection of recommended VM specifications across multiple source servers.
type RecommendedVmSpecList struct {
	Status                string              `json:"status"`
	Description           string              `json:"description"`
	Count                 int                 `json:"count"`
	RecommendedVmSpecList []RecommendedVmSpec `json:"recommendedVmSpecList"`
}

// RecommendedVmOsImage represents the recommended virtual machine OS image information for a single server.
type RecommendedVmOsImage struct {
	Status          string    `json:"status"`
	SourceServers   []string  `json:"sourceServers"`
	Description     string    `json:"description"`
	TargetVmOsImage ImageInfo `json:"targetVmOsImage"`
	// Count            int                   `json:"count"`
	// TargetVmOsImages []ImageInfo `json:"targetVmOsImages"`
}

// RecommendedVmOsImageList represents a collection of recommended VM OS images across multiple source servers.
type RecommendedVmOsImageList struct {
	Status                   string                 `json:"status"`
	Description              string                 `json:"description"`
	Count                    int                    `json:"count"`
	RecommendedVmOsImageList []RecommendedVmOsImage `json:"recommendedVmOsImageList"`
}
