package cloudmodel

// RecommendedInfraModel represents the recommended virtual machine infrastructure model.
type RecommendedInfraModel struct {
	RecommendedInfraModel RecommendedInfra `json:"recommendedInfraModel" validate:"required"`
}

// RecommendedInfra represents the recommended virtual machine infrastructure information.
type RecommendedInfra struct {
	NameSeed                string             `json:"nameSeed"`
	Status                  string             `json:"status"`
	Description             string             `json:"description"`
	TargetCloud             CloudProperty      `json:"targetCloud"`
	TargetInfra             InfraReq           `json:"targetInfra"`
	TargetVNet              VNetReq            `json:"targetVNet"`
	TargetSshKey            SshKeyReq          `json:"targetSshKey"`
	TargetSpecList          []SpecInfo         `json:"targetSpecList"`
	TargetOsImageList       []ImageInfo        `json:"targetOsImageList"`
	TargetSecurityGroupList []SecurityGroupReq `json:"targetSecurityGroupList"`
}

// CloudProperty represents the cloud service provider (CSP) information.
type CloudProperty struct {
	Csp    string `json:"csp" example:"aws"`
	Region string `json:"region" example:"ap-northeast-2"`
}

// RecommendedInfraDynamic represents the recommended virtual machine infrastructure information.
type RecommendedInfraDynamic struct {
	Status      string          `json:"status"`
	Description string          `json:"description"`
	TargetInfra InfraDynamicReq `json:"targetInfra"`
}

// RecommendedInfraDynamicList represents a list of recommended virtual machine infrastructure information.
type RecommendedInfraDynamicList struct {
	Description     string                    `json:"description"`
	Count           int                       `json:"count"`
	TargetInfraList []RecommendedInfraDynamic `json:"targetInfraList"`
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

// RecommendedSpec represents the recommended virtual machine specification information for a single server.
type RecommendedSpec struct {
	Status        string   `json:"status"`
	SourceServers []string `json:"sourceServers"`
	Description   string   `json:"description"`
	TargetSpec    SpecInfo `json:"targetSpec"`
}

// RecommendedSpecList represents a collection of recommended VM specifications across multiple source servers.
type RecommendedSpecList struct {
	Status              string            `json:"status"`
	Description         string            `json:"description"`
	Count               int               `json:"count"`
	RecommendedSpecList []RecommendedSpec `json:"recommendedSpecList"`
}

// RecommendedOsImage represents the recommended virtual machine OS image information for a single server.
type RecommendedOsImage struct {
	Status        string    `json:"status"`
	SourceServers []string  `json:"sourceServers"`
	Description   string    `json:"description"`
	TargetOsImage ImageInfo `json:"targetOsImage"`
	// Count            int                   `json:"count"`
	// TargetVmOsImages []ImageInfo `json:"targetVmOsImages"`
}

// RecommendedOsImageList represents a collection of recommended VM OS images across multiple source servers.
type RecommendedOsImageList struct {
	Status                 string               `json:"status"`
	Description            string               `json:"description"`
	Count                  int                  `json:"count"`
	RecommendedOsImageList []RecommendedOsImage `json:"recommendedOsImageList"`
}
