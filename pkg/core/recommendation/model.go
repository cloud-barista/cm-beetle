package recommendation

import tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"

// RecommendationStatus represents the status of a recommendation.
type RecommendationStatus string

const (
	NothingRecommended   RecommendationStatus = "none"
	PartiallyRecommended RecommendationStatus = "partial"
	FullyRecommended     RecommendationStatus = "ok"
)

type CspRegionPair struct {
	Csp    string `json:"csp" example:"aws"`
	Region string `json:"region" example:"ap-northeast-2"`
}

type RecommendedVmInfraInfo struct {
	Status        string                  `json:"status"`
	Description   string                  `json:"description"`
	TargetVmInfra tbmodel.TbMciDynamicReq `json:"targetVmInfra"`
}

type RecommendedVmInfraInfoList struct {
	Description       string                   `json:"description"`
	Count             int                      `json:"count"`
	TargetVmInfraList []RecommendedVmInfraInfo `json:"targetVmInfraList"`
}

type RecommendedNetworkInfo struct {
	Status        string            `json:"status"`
	Description   string            `json:"description"`
	TargetNetwork tbmodel.TbVNetReq `json:"targetNetwork"`
}

type RecommendedNetworkInfoList struct {
	Description       string                   `json:"description"`
	Count             int                      `json:"count"`
	TargetNetworkList []RecommendedNetworkInfo `json:"targetNetworkList"`
}

type RecommendedSecurityGroupInfo struct {
	Status              string                     `json:"status"`
	Description         string                     `json:"description"`
	TargetSecurityGroup tbmodel.TbSecurityGroupReq `json:"targetSecurityGroup"`
}

type RecommendedSecurityGroupInfoList struct {
	Description             string                         `json:"description"`
	Count                   int                            `json:"count"`
	TargetSecurityGroupList []RecommendedSecurityGroupInfo `json:"targetSecurityGroupList"`
}

type RecommendedInfraInfo struct {
	Status      string                  `json:"status"`
	Description string                  `json:"description"`
	TargetInfra tbmodel.TbMciDynamicReq `json:"targetInfra"`
}
