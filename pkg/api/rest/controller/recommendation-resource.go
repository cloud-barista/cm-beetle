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

// Package common is to handle REST API for common funcitonalities
package controller

import (
	"fmt"
	"net/http"

	"github.com/cloud-barista/cb-tumblebug/src/core/common/netutil"
	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type RecommendNetworkRequest struct {
	DesiredProvider string   `json:"desiredProvider" example:"aws"`
	DesiredRegion   string   `json:"desiredRegion" example:"ap-northeast-2"`
	CidrBlocks      []string `json:"cidrBlocks" example:""`
}

type RecommendNetworkResponse struct {
	recommendation.RecommendedNetworkList
}

// RecommendNetwork godoc
// @ID RecommendNetwork
// @Summary Recommend an appropriate network for cloud migration
// @Description Recommend an appropriate network for cloud migration
// @Description
// @Description [Note] `desiredProvider` and `desiredRegion` are required.
// @Description - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredProvider and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Tags [Recommendation] Resource
// @Accept json
// @Produce	json
// @Param UserNetwork body RecommendNetworkRequest true "Specify the your network to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendNetworkResponse "The result of recommended network"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/mci/network [post]
func RecommendNetwork(c echo.Context) error {

	res, err := recommendNetwork(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend network")
		res := common.SimpleMsg{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, res)
}

func recommendNetwork(c echo.Context) (RecommendNetworkResponse, error) {

	emptyRes := RecommendNetworkResponse{}
	recommendedNetworkInfoList := RecommendNetworkResponse{}

	// [Input]
	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	var req RecommendNetworkRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request")
		return emptyRes, err
	}

	log.Trace().Msgf("req: %v\n", req)

	// Validate the input
	// if req.DesiredProvider == "" && desiredProvider == "" {
	if desiredProvider == "" {
		err := fmt.Errorf("invalid request: 'desiredProvider' is required")
		log.Warn().Msg(err.Error())
		return emptyRes, err
	}
	// if req.DesiredRegion == "" && desiredRegion == "" {
	if desiredRegion == "" {
		err := fmt.Errorf("invalid request: 'desiredRegion' is required")
		log.Warn().Msg(err.Error())
		return emptyRes, err
	}

	provider := desiredProvider
	region := desiredRegion

	ok, err := recommendation.IsValidCspAndRegion(provider, region)
	if !ok {
		log.Error().Err(err).Msgf("invalid provider (%s) or region (%s)", provider, region)
		return emptyRes, err
	}

	// ! TBD: Validate req if needed

	// ! It's a dummy data. It should be replaced with the actual model.
	req.CidrBlocks = []string{
		"192.168.0.0/24",
		"192.168.1.0/24",
	}
	srcNetworks := req.CidrBlocks

	// [Process] Recommend the network vNet and subnets
	// * Note:
	// * At least 1 subnet is required.
	// * Derive a super network that includes user's all networks and set it as a vNet
	// * Set user's networks as subnets

	// ? Assumption: a network in on-premise infrastructure is designed and configured with various network segments or types.
	// * Thus, it must be selected which of these network segments will be the vNet.
	// ? If so, is grouping the network segments required?

	// Categorizes the entered CIDR blocks by private network (i.e., 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16)
	cidrs10 := []string{}
	cidrs172 := []string{}
	cidrs192 := []string{}

	for _, srcNetwork := range srcNetworks {
		identifiedNet, err := netutil.WhichPrivateNetworkByCidr(srcNetwork)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to identify the network %s", srcNetwork)
			continue
		}

		switch identifiedNet {
		case netutil.PrivateNetwork10Dot:
			cidrs10 = append(cidrs10, srcNetwork)
		case netutil.PrivateNetwork172Dot:
			cidrs172 = append(cidrs172, srcNetwork)
		case netutil.PrivateNetwork192Dot:
			cidrs192 = append(cidrs192, srcNetwork)
		default:
			log.Warn().Msgf("skipped because CIDR block (%s) is not a private network", srcNetwork)
			continue
		}
	}

	// Calculate the super network of each group
	var supernet10, supernet172, supernet192 string = "", "", ""

	if len(cidrs10) > 0 {
		supernet10, err = netutil.CalculateSupernet(cidrs10)
		if err != nil {
			log.Warn().Err(err).Msg("failed to calculate supernet")
		}
		log.Debug().Msgf("supernet10: %s\n", supernet10)
	}

	if len(cidrs172) > 0 {
		supernet172, err = netutil.CalculateSupernet(cidrs172)
		if err != nil {
			log.Warn().Err(err).Msg("failed to calculate supernet")
		}
		log.Debug().Msgf("supernet172: %s\n", supernet172)
	}

	if len(cidrs192) > 0 {
		supernet192, err = netutil.CalculateSupernet(cidrs192)
		if err != nil {
			log.Warn().Err(err).Msg("failed to calculate supernet")
		}
		log.Debug().Msgf("supernet192: %s\n", supernet192)
	}

	// Select a super network for the vNet
	// ? But how to select the super network?
	// * Currrently, a list of recommended networks is returned.

	recommendedNetworks := []recommendation.RecommendedNetwork{}
	if supernet10 != "" {

		// Set subnets by the CIDR blocks from the source computing infra
		subnets := []tbmodel.TbSubnetReq{}
		for i, cidr := range cidrs10 {
			subnets = append(subnets, tbmodel.TbSubnetReq{
				Name:        fmt.Sprintf("subnet-%02d", i+1),
				Description: "subnet from source computing infra",
				IPv4_CIDR:   cidr,
			})
		}

		// Set the calculated supernet as the vNet
		tempNetworkInfo := recommendation.RecommendedNetwork{
			Status:      "",
			Description: "Recommended network information",
			TargetNetwork: tbmodel.TbVNetReq{
				Name:           "vnet-01",
				Description:    "Recommended vNet for " + netutil.PrivateNetwork10Dot,
				CidrBlock:      supernet10,
				SubnetInfoList: subnets,
			},
		}

		// Append recommended network info to the list
		recommendedNetworks = append(recommendedNetworks, tempNetworkInfo)
	}

	if supernet172 != "" {

		// Set subnets by the CIDR blocks from the source computing infra
		subnets := []tbmodel.TbSubnetReq{}
		for i, cidr := range cidrs172 {
			subnets = append(subnets, tbmodel.TbSubnetReq{
				Name:        fmt.Sprintf("subnet-%02d", i+1),
				Description: "subnet from source computing infra",
				IPv4_CIDR:   cidr,
			})
		}

		tempNetworkInfo := recommendation.RecommendedNetwork{
			Status:      "",
			Description: "Recommended network information",
			TargetNetwork: tbmodel.TbVNetReq{
				Name:           "vnet-01",
				Description:    "Recommended vNet for " + netutil.PrivateNetwork172Dot,
				CidrBlock:      supernet172,
				SubnetInfoList: subnets,
			},
		}
		// Append recommended network info to the list
		recommendedNetworks = append(recommendedNetworks, tempNetworkInfo)
	}

	if supernet192 != "" {

		// Set subnets by the CIDR blocks from the source computing infra
		subnets := []tbmodel.TbSubnetReq{}
		for i, cidr := range cidrs192 {
			subnets = append(subnets, tbmodel.TbSubnetReq{
				Name:        fmt.Sprintf("subnet-%02d", i+1),
				Description: "subnet from source computing infra",
				IPv4_CIDR:   cidr,
			})
		}

		// Set the calculated supernet as the vNet
		tempNetworkInfo := recommendation.RecommendedNetwork{
			Status:      "",
			Description: "Recommended network information",
			TargetNetwork: tbmodel.TbVNetReq{
				Name:           "vnet-01",
				Description:    "Recommended vNet for " + netutil.PrivateNetwork192Dot,
				CidrBlock:      supernet192,
				SubnetInfoList: subnets,
			},
		}

		// Append recommended network info to the list
		recommendedNetworks = append(recommendedNetworks, tempNetworkInfo)
	}

	// * TBD: Consider the number of hosts

	// [Output]
	recommendedNetworkInfoList.Description = "Recommended network information list"
	recommendedNetworkInfoList.Count = len(recommendedNetworks)
	recommendedNetworkInfoList.TargetNetworkList = recommendedNetworks

	log.Debug().Msgf("recommendedNetworkInfoList: %v\n", recommendedNetworkInfoList)

	return recommendedNetworkInfoList, nil
}

type RecommendSecurityGroupRequest struct {
	// ! To be replaced with the actual model
	// FirewallRules []inframodel.FirewallRuleProperty `json:"firewallRules" example:""`
	FirewallRules []FirewallRuleProperty `json:"firewallRules" example:""`
}

// To be replaced with the actual model
type FirewallRuleProperty struct { // note: reference command `sudo ufw status verbose`
	SrcCIDR   string `json:"srcCIDR,omitempty"`
	DstCIDR   string `json:"dstCIDR,omitempty"`
	SrcPorts  string `json:"srcPorts,omitempty"`
	DstPorts  string `json:"dstPorts,omitempty"`
	Protocol  string `json:"protocol,omitempty"`  // TCP, UDP, ICMP
	Direction string `json:"direction,omitempty"` // inbound, outbound
	Action    string `json:"action,omitempty"`    // allow, deny
}

type RecommendSecurityGroupResponse struct {
	recommendation.RecommendedSecurityGroupList
}

// RecommendSecurityGroup godoc
// @ID RecommendSecurityGroup
// @Summary Recommend an appropriate security group for cloud migration
// @Description Recommend an appropriate security group for cloud migration
// @Description
// @Description [Note] `desiredProvider` and `desiredRegion` are required.
// @Description - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredProvider and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Tags [Recommendation] Resource
// @Accept  json
// @Produce  json
// @Param UserNetwork body RecommendSecurityGroupRequest true "Specify the your network to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendSecurityGroupResponse "The result of recommended network"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/mci/securityGroup [post]
func RecommendSecurityGroup(c echo.Context) error {
	res, err := recommendSecurityGroup(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend security group")
		res := common.SimpleMsg{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, res.(RecommendSecurityGroupResponse))
}

func recommendSecurityGroup(c echo.Context) (any, error) {
	emptyRes := RecommendSecurityGroupResponse{}
	recommendedSecurityGroupInfoList := RecommendSecurityGroupResponse{}

	// [Input]
	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	var req RecommendSecurityGroupRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request")
		return emptyRes, err
	}

	log.Trace().Msgf("req: %v\n", req)

	// Validate the input
	if desiredProvider == "" {
		err := fmt.Errorf("invalid request: 'desiredProvider' is required")
		log.Warn().Msg(err.Error())
		return emptyRes, err
	}
	if desiredRegion == "" {
		err := fmt.Errorf("invalid request: 'desiredRegion' is required")
		log.Warn().Msg(err.Error())
		return emptyRes, err
	}

	provider := desiredProvider
	region := desiredRegion

	ok, err := recommendation.IsValidCspAndRegion(provider, region)
	if !ok {
		log.Error().Err(err).Msgf("invalid provider (%s) or region (%s)", provider, region)
		return emptyRes, err
	}

	// Use the provided firewall rules or fall back to dummy data if empty
	if len(req.FirewallRules) == 0 {
		log.Warn().Msg("no firewall rules provided, using sample data")
		req.FirewallRules = dummyFirewallRules
	}

	// [Process] Recommend the security group

	// Create security group recommendations
	recommendedSecurityGroups := []recommendation.RecommendedSecurityGroup{}

	// ! To be updated with the actual model and real data
	// ! A list of firewall rules(i.e., firewall table) will be entered (currently, it's a dummy single firewall table)

	// Create a security group for all rules
	recommendedSecurityGroup := recommendation.RecommendedSecurityGroup{
		Status:      string(recommendation.FullyRecommended),
		Description: "Security group containing all firewall rules",
		TargetSecurityGroup: tbmodel.TbSecurityGroupReq{
			Name:          "recommended-sg-all-rules",
			Description:   "Recommended security group for all firewall rules",
			VNetId:        "subnet-to-be-defined", // This should be defined based on network recommendation or user input
			FirewallRules: generateSecurityGroupRules(req.FirewallRules),
		},
	}
	recommendedSecurityGroups = append(recommendedSecurityGroups, recommendedSecurityGroup)

	// *TBD: Remove duplicate Security Groups

	// *TBD: Consider how to assign a security group to a host

	// [Output]
	recommendedSecurityGroupInfoList.Description = "Recommended security group information list"
	recommendedSecurityGroupInfoList.Count = len(recommendedSecurityGroups)
	recommendedSecurityGroupInfoList.TargetSecurityGroupList = recommendedSecurityGroups

	log.Debug().Msgf("recommendedSecurityGroupInfoList: %+v\n", recommendedSecurityGroupInfoList)

	return recommendedSecurityGroupInfoList, nil
}

// generateSecurityGroupRules converts FirewallRuleProperty to tbmodel.TbFirewallRuleInfo
func generateSecurityGroupRules(rules []FirewallRuleProperty) *[]tbmodel.TbFirewallRuleInfo {
	var tbRules []tbmodel.TbFirewallRuleInfo

	for _, rule := range rules {

		// Skip 'deny' rules (note: SecurityGroup does not support adding 'deny' rules)
		if rule.Action == "deny" {
			continue
		}

		// ? To be handle multiple ports comma-separated
		// TBD

		tbRule := tbmodel.TbFirewallRuleInfo{
			Direction:  rule.Direction,
			IPProtocol: rule.Protocol,
		}

		// ! To be updated with the actual model and real data
		// Set port information based on direction
		if rule.Direction == "inbound" {
			// Set port or port range (Note: use destination ports for inbound traffic)
			// ! Handle multiple ports comma-seprated (AWS dosen't allow comma-seperated)
			if rule.DstPorts != "" {
				tbRule.FromPort = rule.DstPorts
				tbRule.ToPort = rule.DstPorts
			}
		} else if rule.Direction == "outbound" {
			// For outbound traffic, we use source ports if available, otherwise destination ports
			if rule.SrcPorts != "" {
				tbRule.FromPort = rule.SrcPorts
				tbRule.ToPort = rule.SrcPorts
			} else if rule.DstPorts != "" {
				tbRule.FromPort = rule.DstPorts
				tbRule.ToPort = rule.DstPorts
			}
		}

		// Set CIDR based on direction with fallbacks
		if rule.Direction == "inbound" {
			// For inbound, use source CIDR (where traffic comes from)
			if rule.SrcCIDR != "" {
				tbRule.CIDR = rule.SrcCIDR
			} else if rule.DstCIDR != "" {
				// Fallback to destination CIDR if source is not specified
				tbRule.CIDR = rule.DstCIDR
				log.Warn().Msgf("Using DstCIDR for inbound rule because SrcCIDR is not specified")
			} else {
				// Default to any address if no CIDR specified
				tbRule.CIDR = "0.0.0.0/0"
				log.Warn().Msgf("Using 0.0.0.0/0 for inbound rule because no CIDR is specified")
			}
		} else if rule.Direction == "outbound" {
			// For outbound, use destination CIDR (where traffic goes to)
			if rule.DstCIDR != "" {
				tbRule.CIDR = rule.DstCIDR
			} else if rule.SrcCIDR != "" {
				// Fallback to source CIDR if destination is not specified
				tbRule.CIDR = rule.SrcCIDR
				log.Warn().Msgf("Using SrcCIDR for outbound rule because DstCIDR is not specified")
			} else {
				// Default to any address if no CIDR specified
				tbRule.CIDR = "0.0.0.0/0"
				log.Warn().Msgf("Using 0.0.0.0/0 for outbound rule because no CIDR is specified")
			}
		}

		tbRules = append(tbRules, tbRule)

		log.Debug().Msgf("FirwewallRule: %+v", rule)
		log.Debug().Msgf("SecurityGroupRule: %+v", tbRule)
	}

	return &tbRules
}

var dummyFirewallRules = []FirewallRuleProperty{
	{
		SrcCIDR:   "0.0.0.0/0",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "22",
		Protocol:  "TCP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "0.0.0.0/0",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "80,443",
		Protocol:  "TCP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "10.0.0.0/16",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "3306",
		Protocol:  "TCP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "172.16.0.0/12",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "5432",
		Protocol:  "TCP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "0.0.0.0/0",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "25",
		Protocol:  "TCP",
		Direction: "inbound",
		Action:    "deny",
	},
	{
		SrcCIDR:   "0.0.0.0/0",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "53",
		Protocol:  "UDP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "0.0.0.0/0",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "1194",
		Protocol:  "UDP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "192.168.1.10/32",
		DstCIDR:   "0.0.0.0/0",
		SrcPorts:  "32768-60999",
		Protocol:  "TCP",
		Direction: "outbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "192.168.1.10/32",
		DstCIDR:   "8.8.8.8/32",
		DstPorts:  "53",
		Protocol:  "UDP",
		Direction: "outbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "192.168.1.10/32",
		DstCIDR:   "0.0.0.0/0",
		Protocol:  "ICMP",
		Direction: "outbound",
		Action:    "allow",
	},
}
