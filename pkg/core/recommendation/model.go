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

// * Note: The models for VM infrastructure is moved to cm-model repo for sharing and use with other subsystems.

/*
 * Models for VM infrastructure
 */

// // RecommendedVmInfra represents the recommended virtual machine infrastructure information.
// type RecommendedVmInfra struct {
// 	Status                  string                       `json:"status"`
// 	Description             string                       `json:"description"`
// 	TargetVmInfra           tbmodel.TbMciReq             `json:"targetVmInfra"`
// 	TargetVNet              tbmodel.TbVNetReq            `json:"targetVNet"`
// 	TargetSshKey            tbmodel.TbSshKeyReq          `json:"targetSshKey"`
// 	TargetVmSpecList        []tbmodel.TbSpecInfo         `json:"targetVmSpecList"`
// 	TargetVmOsImageList     []tbmodel.TbImageInfo        `json:"targetVmOsImageList"`
// 	TargetSecurityGroupList []tbmodel.TbSecurityGroupReq `json:"targetSecurityGroupList"`
// }

// // RecommendedVmInfraDynamic represents the recommended virtual machine infrastructure information.
// type RecommendedVmInfraDynamic struct {
// 	Status        string                  `json:"status"`
// 	Description   string                  `json:"description"`
// 	TargetVmInfra tbmodel.TbMciDynamicReq `json:"targetVmInfra"`
// }

// // RecommendedVmInfraDynamicList represents a list of recommended virtual machine infrastructure information.
// type RecommendedVmInfraDynamicList struct {
// 	Description       string                      `json:"description"`
// 	Count             int                         `json:"count"`
// 	TargetVmInfraList []RecommendedVmInfraDynamic `json:"targetVmInfraList"`
// }

// // RecommendedVNet represents the recommended virtual network information.
// // * May be mainly used this object
// type RecommendedVNet struct {
// 	Status      string            `json:"status"`
// 	Description string            `json:"description"`
// 	TargetVNet  tbmodel.TbVNetReq `json:"targetVNet"`
// }

// // RecommendedVNetList represents a list of recommended virtual network information.
// type RecommendedVNetList struct {
// 	Description    string            `json:"description"`
// 	Count          int               `json:"count"`
// 	TargetVNetList []RecommendedVNet `json:"targetVNetList"`
// }

// // RecommendedSecurityGroup represents the recommended security group information.
// type RecommendedSecurityGroup struct {
// 	Status              string                     `json:"status"`
// 	SourceServers       []string                   `json:"sourceServers"`
// 	Description         string                     `json:"description"`
// 	TargetSecurityGroup tbmodel.TbSecurityGroupReq `json:"targetSecurityGroup"`
// }

// // RecommendedSecurityGroupList represents a list of recommended security group information.
// type RecommendedSecurityGroupList struct {
// 	Status                  string                     `json:"status"`
// 	Description             string                     `json:"description"`
// 	Count                   int                        `json:"count"`
// 	TargetSecurityGroupList []RecommendedSecurityGroup `json:"targetSecurityGroupList"`
// }

// // RecommendedVmSpec represents the recommended virtual machine specification information for a single server.
// type RecommendedVmSpec struct {
// 	Status        string             `json:"status"`
// 	SourceServers []string           `json:"sourceServers"`
// 	Description   string             `json:"description"`
// 	TargetVmSpec  tbmodel.TbSpecInfo `json:"targetVmSpec"`
// }

// // RecommendedVmSpecList represents a collection of recommended VM specifications across multiple source servers.
// type RecommendedVmSpecList struct {
// 	Status                string              `json:"status"`
// 	Description           string              `json:"description"`
// 	Count                 int                 `json:"count"`
// 	RecommendedVmSpecList []RecommendedVmSpec `json:"recommendedVmSpecList"`
// }

// // RecommendedVmOsImage represents the recommended virtual machine OS image information for a single server.
// type RecommendedVmOsImage struct {
// 	Status          string              `json:"status"`
// 	SourceServers   []string            `json:"sourceServers"`
// 	Description     string              `json:"description"`
// 	TargetVmOsImage tbmodel.TbImageInfo `json:"targetVmOsImage"`
// 	// Count            int                   `json:"count"`
// 	// TargetVmOsImages []tbmodel.TbImageInfo `json:"targetVmOsImages"`
// }

// // RecommendedVmOsImageList represents a collection of recommended VM OS images across multiple source servers.
// type RecommendedVmOsImageList struct {
// 	Status                   string                 `json:"status"`
// 	Description              string                 `json:"description"`
// 	Count                    int                    `json:"count"`
// 	RecommendedVmOsImageList []RecommendedVmOsImage `json:"recommendedVmOsImageList"`
// }

/*
 * Models for Container Infrastructure (i.e., an infrastructure for Kubernetes)
 */
type RecommendedInfraInfo struct {
	Status      string                  `json:"status"`
	Description string                  `json:"description"`
	TargetInfra tbmodel.TbMciDynamicReq `json:"targetInfra"`
}
