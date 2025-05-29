package recommendation

import tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"

// RecommendationStatus represents the status of a recommendation.
type RecommendationStatus string

const (
	NothingRecommended   RecommendationStatus = "none"
	PartiallyRecommended RecommendationStatus = "partial"
	FullyRecommended     RecommendationStatus = "ok"
)

// CspRegionPair represents a pair of cloud service provider (CSP) and region.
type CspRegionPair struct {
	Csp    string `json:"csp" example:"aws"`
	Region string `json:"region" example:"ap-northeast-2"`
}

// RecommendedVmInfra represents the recommended virtual machine infrastructure information.
type RecommendedVmInfra struct {
	Status                  string                       `json:"status"`
	Description             string                       `json:"description"`
	TargetVmInfra           tbmodel.TbMciReq             `json:"targetVmInfra"`
	TargetNetwork           tbmodel.TbVNetReq            `json:"targetNetwork"`
	TargetVmSpecList        []tbmodel.TbSpecReq          `json:"targetVmSpecList"`
	TargetVmOsImageList     []tbmodel.TbImageReq         `json:"targetVmOsImageList"`
	TargetSecurityGroupList []tbmodel.TbSecurityGroupReq `json:"targetSecurityGroupList"`
}

// RecommendedVmInfraDynamic represents the recommended virtual machine infrastructure information.
type RecommendedVmInfraDynamic struct {
	Status        string                  `json:"status"`
	Description   string                  `json:"description"`
	TargetVmInfra tbmodel.TbMciDynamicReq `json:"targetVmInfra"`
}

// RecommendedVmInfraDynamicList represents a list of recommended virtual machine infrastructure information.
type RecommendedVmInfraDynamicList struct {
	Description       string                      `json:"description"`
	Count             int                         `json:"count"`
	TargetVmInfraList []RecommendedVmInfraDynamic `json:"targetVmInfraList"`
}

// RecommendedNetwork represents the recommended network information.
// * May be mainly used this object
type RecommendedNetwork struct {
	Status        string            `json:"status"`
	Description   string            `json:"description"`
	TargetNetwork tbmodel.TbVNetReq `json:"targetNetwork"`
}

// RecommendedNetworkList represents a list of recommended network information.
type RecommendedNetworkList struct {
	Description       string               `json:"description"`
	Count             int                  `json:"count"`
	TargetNetworkList []RecommendedNetwork `json:"targetNetworkList"`
}

// RecommendedSecurityGroup represents the recommended security group information.
type RecommendedSecurityGroup struct {
	Status              string                     `json:"status"`
	Description         string                     `json:"description"`
	TargetSecurityGroup tbmodel.TbSecurityGroupReq `json:"targetSecurityGroup"`
}

// RecommendedSecurityGroupList represents a list of recommended security group information.
type RecommendedSecurityGroupList struct {
	Description             string                     `json:"description"`
	Count                   int                        `json:"count"`
	TargetSecurityGroupList []RecommendedSecurityGroup `json:"targetSecurityGroupList"`
}

// RecommendedVmSpec represents the recommended virtual machine specification information.
type RecommendedVmSpec struct {
	Status       string            `json:"status"`
	Description  string            `json:"description"`
	TargetVmSpec tbmodel.TbSpecReq `json:"targetVmSpec"`
}

// RecommendedVmSpecList represents a list of recommended virtual machine specification information.
type RecommendedVmSpecList struct {
	Description      string              `json:"description"`
	Count            int                 `json:"count"`
	TargetVmSpecList []RecommendedVmSpec `json:"targetVmSpecList"`
}

// RecommendedVmOsImage represents the recommended virtual machine OS image information.
type RecommendedVmOsImage struct {
	Status          string             `json:"status"`
	Description     string             `json:"description"`
	TargetVmOsImage tbmodel.TbImageReq `json:"targetVmOsImage"`
}

// RecommendedVmOsImageList represents a list of recommended virtual machine OS image information.
type RecommendedVmOsImageList struct {
	Description         string                 `json:"description"`
	Count               int                    `json:"count"`
	TargetVmOsImageList []RecommendedVmOsImage `json:"targetVmOsImageList"`
}

/*
 * Models for Container Infrastructure (i.e., an infrastructure for Kubernetes)
 */

type RecommendedInfraInfo struct {
	Status      string                  `json:"status"`
	Description string                  `json:"description"`
	TargetInfra tbmodel.TbMciDynamicReq `json:"targetInfra"`
}
