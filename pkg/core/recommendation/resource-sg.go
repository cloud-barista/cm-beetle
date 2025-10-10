package recommendation

import (
	"fmt"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
	"github.com/rs/zerolog/log"
)

func RecommendSecurityGroup(csp string, region string, server onpremmodel.ServerProperty) (cloudmodel.SecurityGroupReq, error) {

	var emptyRes = cloudmodel.SecurityGroupReq{}
	var recommendedSecurityGroup = cloudmodel.SecurityGroupReq{}

	// [Input]
	ok, err := IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msgf("invalid provider (%s) or region (%s)", csp, region)
		return emptyRes, err
	}

	firewallRules := server.FirewallTable
	log.Debug().Msgf("firewallRules: %+v", firewallRules)

	// Default rules
	// * Note: Spider supports this rule. Do not set this rule to avoid duplication error.
	// ruleToAllowAllOutboundTraffic := cloudmodel.FirewallRuleReq{
	// 	Direction:  "outbound",
	// 	IPProtocol: "all",
	// 	CIDR:       "0.0.0.0/0",
	// 	FromPort:   "0",
	// 	ToPort:     "0",

	// [Process] Recommend the security group based on server.FirewallTable
	// Create security group recommendations
	var sgRules []cloudmodel.FirewallRuleReq
	var firewallRulesPtr *[]cloudmodel.FirewallRuleReq

	// Generate security group rules based on server firewall configuration
	if len(firewallRules) == 0 {
		log.Warn().Msg("No firewall rules provided from server.FirewallTable - security group will be created without predefined rules")
		// Note: SSH access rule will be added during migration phase if needed
		firewallRulesPtr = nil // Use nil to indicate no rules defined
	} else {
		log.Info().Msgf("Generating security group rules based on %d firewall rules from server configuration", len(firewallRules))
		sgRules = generateSecurityGroupRules(firewallRules)
		firewallRulesPtr = &sgRules // Point to the generated rules
	}

	log.Debug().Msgf("sgRules: %+v", sgRules)

	// [Output]
	// Create a security group for all rules
	recommendedSecurityGroup = cloudmodel.SecurityGroupReq{
		Name:           "INSERT_YOUR_SECURITY_GROUP_NAME",
		VNetId:         "INSERT_YOUR_VNET_ID",
		ConnectionName: fmt.Sprintf("%s-%s", csp, region),
		Description:    fmt.Sprintf("Recommended security group for %s", server.MachineId), // Set MachineId to identify the source server
		FirewallRules:  firewallRulesPtr,
	}

	log.Debug().Msgf("recommendedSecurityGroup: %+v", recommendedSecurityGroup)

	return recommendedSecurityGroup, nil
}

func RecommendSecurityGroups(csp string, region string, servers []onpremmodel.ServerProperty) (cloudmodel.RecommendedSecurityGroupList, error) {

	var emptyRet = cloudmodel.RecommendedSecurityGroupList{}
	var recommendedSecurityGroupList = cloudmodel.RecommendedSecurityGroupList{}

	// [Input]
	ok, err := IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msgf("invalid provider (%s) or region (%s)", csp, region)
		return emptyRet, err
	}

	// [Process] Recommend the security group for each server
	var tempRecSgList []cloudmodel.SecurityGroupReq
	var targetSecurityGroupList []cloudmodel.RecommendedSecurityGroup

	for _, server := range servers {
		// Recommend a security group for the server
		recommendedTargetSg, err := RecommendSecurityGroup(csp, region, server)
		if err != nil {
			log.Error().Err(err).Msgf("failed to recommend security group for server: %+v", server)
			recommendedTargetSg.Description = fmt.Sprintf("Failed to recommend security group for %s", server.MachineId) // Set MachineId to identify the source server
			recommendedTargetSg.FirewallRules = nil                                                                      // No rules if recommendation fails
		}

		// Check duplicates and append the recommended security group
		exists, idx, _ := containSg(tempRecSgList, recommendedTargetSg)

		// If not exists, append the recommended security group
		// If exists, just append the MachineId to the existing security group
		if !exists {
			// Note: This is a temporary list for checking duplicates
			tempRecSgList = append(tempRecSgList, recommendedTargetSg)

			// Create a temporary recommended security group
			tempRecommendedSecurityGroup := cloudmodel.RecommendedSecurityGroup{
				SourceServers:       []string{server.MachineId}, // Set MachineId to identify the source server
				Description:         "Recommended security group",
				TargetSecurityGroup: recommendedTargetSg,
			}

			// Set status
			if recommendedTargetSg.FirewallRules != nil {
				tempRecommendedSecurityGroup.Status = string(FullyRecommended)
			} else {
				tempRecommendedSecurityGroup.Status = string(NothingRecommended)
			}

			// Append the recommended security group to the list
			targetSecurityGroupList = append(targetSecurityGroupList, tempRecommendedSecurityGroup)
		} else {
			// Just append the MachineId to the existing security group
			targetSecurityGroupList[idx].SourceServers = append(targetSecurityGroupList[idx].SourceServers, server.MachineId)
		}
	}

	// [Output]
	countFailed := 0
	for _, recSg := range targetSecurityGroupList {
		if recSg.Status == string(NothingRecommended) {
			countFailed++
		}
	}

	recommendedSecurityGroupList.Count = len(targetSecurityGroupList)
	switch countFailed {
	case 0:
		recommendedSecurityGroupList.Status = string(FullyRecommended)
		recommendedSecurityGroupList.Description = "Successfully recommended and deduplicated security groups for all servers"
	case recommendedSecurityGroupList.Count:
		recommendedSecurityGroupList.Status = string(NothingRecommended)
		recommendedSecurityGroupList.Description = "Unable to recommend any security groups for the servers in the source infrastructure"
	default:
		recommendedSecurityGroupList.Status = string(PartiallyRecommended)
		recommendedSecurityGroupList.Description = fmt.Sprintf("Partially recommended security groups: %d of %d server groups have recommendations",
			recommendedSecurityGroupList.Count-countFailed, recommendedSecurityGroupList.Count)
	}

	recommendedSecurityGroupList.TargetSecurityGroupList = targetSecurityGroupList

	log.Debug().Msgf("recommendedSecurityGroupList: %+v", recommendedSecurityGroupList)

	return recommendedSecurityGroupList, nil
}

func containSg(sgList []cloudmodel.SecurityGroupReq, sg cloudmodel.SecurityGroupReq) (bool, int, cloudmodel.SecurityGroupReq) {

	log.Debug().Msgf("Checking for duplicate security group: %+v", sg)
	log.Debug().Msgf("Firewall rules: %+v", sg.FirewallRules)

	// Check duplicates and append the recommended security group
	temp := cloudmodel.SecurityGroupReq{}
	exists := false
	idx := -1
	for i, sgItem := range sgList {
		// Both SGs have rules defined
		if sgItem.FirewallRules != nil && sg.FirewallRules != nil {
			// Quick check if they have the same number of rules
			if len(*sgItem.FirewallRules) == len(*sg.FirewallRules) {
				areAllRulesSame := true

				// Create maps for each rule in both security groups for comparison
				sgRulesMap := make(map[string]bool)
				for _, rule := range *sg.FirewallRules {
					// Create a unique key for each rule
					key := fmt.Sprintf("%s-%s-%s-%s",
						rule.Direction, rule.Protocol, rule.CIDR, rule.Ports)
					sgRulesMap[key] = true
				}

				// Check if all rules in the recommended SG exist in the current SG
				for _, rule := range *sgItem.FirewallRules {
					key := fmt.Sprintf("%s-%s-%s-%s",
						rule.Direction, rule.Protocol, rule.CIDR, rule.Ports)
					if !sgRulesMap[key] {
						areAllRulesSame = false
						break
					}
				}

				if areAllRulesSame {
					exists = true
					temp = sgItem
					idx = i
					break
				}
			}
		}
	}

	return exists, idx, temp
}

// formatCIDR formats the CIDR string:
// - If it's "anywhere", return "0.0.0.0/0"
// - If it doesn't have a prefix (like "/24"), add "/32"
// - Otherwise return as is
func formatCIDR(cidr string) string {
	if cidr == "anywhere" {
		return "0.0.0.0/0"
	}

	// Check if the CIDR has a prefix
	if !strings.Contains(cidr, "/") {
		// If it's a valid IP without prefix, add "/32"
		return cidr + "/32"
	}

	return cidr
}

// generateSecurityGroupRules converts FirewallRuleProperty to tbmodel.TbFirewallRuleInfo
func generateSecurityGroupRules(rules []onpremmodel.FirewallRuleProperty) []cloudmodel.FirewallRuleReq {
	var tbRules []cloudmodel.FirewallRuleReq

	for _, rule := range rules {
		// Skip 'deny' rules (note: SecurityGroup does not support adding 'deny' rules)
		if rule.Action == "deny" {
			continue
		}

		// Skip rules with no protocol specified
		if rule.Protocol == "" {
			log.Warn().Msgf("Protocol is not specified in rule: %+v - skipping rule", rule)
			continue
		}

		// Skip IPv6 rules (currently not supported)
		if isIPv6Rule(rule) {
			log.Warn().Msgf("IPv6 rule detected but not currently supported: %+v - skipping rule", rule)
			continue
		}

		// Handle protocol wildcard
		protocol := rule.Protocol
		if protocol == "*" {
			protocol = "ALL"
		}

		switch rule.Direction {
		case "inbound":
			// Set CIDR block for source - For inbound, use source CIDR (where traffic comes from)
			if rule.SrcCIDR == "" {
				log.Warn().Msgf("SrcCIDR is not specified in rule: %+v - skipping rule", rule)
				continue
			}

			// Format the CIDR correctly
			srcCIDR := formatCIDR(rule.SrcCIDR)
			log.Debug().Msgf("Formatted SrcCIDR from '%s' to '%s'", rule.SrcCIDR, srcCIDR)

			// ! Skip default outbound rule that allows all traffic because it is automatically created by cloud providers, CB-Spider, or CB-Tumblebug.
			// TODO: To be updated if the default rule is needed in the future.
			if strings.ToLower(protocol) == "all" && srcCIDR == "0.0.0.0/0" {
				log.Debug().Msgf("Skipping default inbound ALL traffic rule (may conflict with existing rules): %+v", rule)
				continue
			}

			// Handle destination ports based on format
			if rule.DstPorts == "" {
				// Skip rules without port information for non-ICMP/ALL protocols
				log.Debug().Msgf("Skipping inbound rule without port information: %+v", rule)
				continue
			}

			// * NOTE: Handle destination ports (where traffic is going to)
			// Special cases based on CB-Spider specification:
			// - (protocol: TCP/UDP/ALL) "*" port from the source is converted to "1-65535" for the target ports.
			// - (protocol: ICMP) "*" ports from the source is omitted in the target ports.
			// - Comma-separated ports (e.g., 22,23,24)
			// - Port range with colon notation (e.g., 30000:40000)
			// - Single port (e.g., 22)

			protocolLower := strings.ToLower(protocol)
			switch protocolLower {
			case "icmp":
				tbRule := cloudmodel.FirewallRuleReq{
					Direction: rule.Direction,
					Protocol:  protocol,
					CIDR:      srcCIDR,
				}
				tbRules = append(tbRules, tbRule)
				log.Debug().Msgf("Created inbound rule for 'ICMP' protocol: %+v", tbRule)

			case "tcp", "udp", "all":
				var dstPorts string
				// Handle wildcard ports based on protocol
				if rule.DstPorts == "*" {
					dstPorts = "1:65535" // TCP/UDP use 1-65535 range
				} else {
					dstPorts = rule.DstPorts
				}

				// * Note: CB-Tumblebug will handle comma-separated ports and port ranges.
				// In here, just convert : to - for port ranges.
				if strings.Contains(dstPorts, ":") {
					// Handle port range with colon notation (e.g., 30000:40000)
					portRange := strings.Split(dstPorts, ":")
					if len(portRange) == 2 {
						dstPorts = strings.TrimSpace(portRange[0]) + "-" + strings.TrimSpace(portRange[1])
					} else {
						log.Warn().Msgf("Invalid port range format in rule.DstPorts: %s - skipping rule", dstPorts)
						continue
					}
				}

				tbRule := cloudmodel.FirewallRuleReq{
					Direction: rule.Direction,
					Protocol:  protocol,
					CIDR:      srcCIDR,
					Ports:     dstPorts,
				}

				tbRules = append(tbRules, tbRule)
				log.Debug().Msgf("Created inbound rule for '%s' protocol: %+v", protocol, tbRule)
			default:
				log.Warn().Msgf("Unsupported protocol '%s' in inbound rule: %+v - skipping rule", protocol, rule)
				continue
			}

		case "outbound":
			// Set CIDR block for destination
			if rule.DstCIDR == "" {
				// Skip rule if no CIDR is specified
				log.Warn().Msgf("No CIDR specified for outbound rule: %+v - skipping rule", rule)
				continue
			}

			// Format the CIDR correctly
			dstCIDR := formatCIDR(rule.DstCIDR)
			log.Debug().Msgf("Formatted outbound CIDR from '%s' to '%s'", rule.DstCIDR, dstCIDR)

			// ! Skip default outbound rule that allows all traffic because it is automatically created by cloud providers, CB-Spider, or CB-Tumblebug.
			// TODO: To be updated if the default rule is needed in the future.
			if strings.ToLower(protocol) == "all" && dstCIDR == "0.0.0.0/0" {
				log.Debug().Msgf("Skipping default outbound ALL traffic rule (may conflict with existing rules): %+v", rule)
				continue
			}

			// Handle destination ports based on format
			if rule.DstPorts == "" {
				// Skip rules without port information for non-ICMP/ALL protocols
				log.Debug().Msgf("Skipping inbound rule without port information: %+v", rule)
				continue
			}

			// * NOTE: Handle destination ports (where traffic is going to)
			// Special cases based on CB-Spider specification:
			// - (protocol: TCP/UDP/ALL) "*" port from the source is converted to "1-65535" for the target ports.
			// - (protocol: ICMP) "*" ports from the source is omitted in the target ports.
			// - Comma-separated ports (e.g., 22,23,24)
			// - Port range with colon notation (e.g., 30000:40000)
			// - Single port (e.g., 22)

			protocolLower := strings.ToLower(protocol)
			switch protocolLower {
			case "icmp":
				// Special case for ICMP protocol - no ports needed, just CIDR
				tbRule := cloudmodel.FirewallRuleReq{
					Direction: rule.Direction,
					Protocol:  protocol,
					CIDR:      dstCIDR,
				}
				tbRules = append(tbRules, tbRule)
				log.Debug().Msgf("Created outbound rule for 'ICMP' protocol: %+v", tbRule)

			case "tcp", "udp", "all": // Handle destination ports with wildcard support based on CB-Spider specification

				var dstPorts string
				// Handle wildcard ports based on protocol
				if rule.DstPorts == "*" {
					dstPorts = "1:65535" // TCP/UDP use 1-65535 range
				} else {
					dstPorts = rule.DstPorts
				}

				// * Note: CB-Tumblebug will handle comma-separated ports and port ranges.
				// In here, just convert : to - for port ranges.
				if strings.Contains(dstPorts, ":") {
					// Handle port range with colon notation
					portRange := strings.Split(dstPorts, ":")
					if len(portRange) == 2 {
						dstPorts = strings.TrimSpace(portRange[0]) + "-" + strings.TrimSpace(portRange[1])
					} else {
						log.Warn().Msgf("Invalid port range format: %s - skipping rule", dstPorts)
						continue
					}
				}

				tbRule := cloudmodel.FirewallRuleReq{
					Direction: rule.Direction,
					Protocol:  protocol,
					CIDR:      dstCIDR,
					Ports:     dstPorts,
				}
				tbRules = append(tbRules, tbRule)
				log.Debug().Msgf("Created outbound rule for '%s' protocol: %+v", protocol, tbRule)
			default:
				log.Warn().Msgf("Unsupported protocol '%s' in outbound rule: %+v - skipping rule", protocol, rule)
				continue
			}

		default:
			log.Warn().Msgf("Unknown direction '%s' in rule: %+v", rule.Direction, rule)
		}

		log.Debug().Msgf("Original FirewallRule: %+v", rule)
	}

	return tbRules
}

// isIPv6Rule checks if the firewall rule contains IPv6 elements
func isIPv6Rule(rule onpremmodel.FirewallRuleProperty) bool {
	// Check for IPv6 CIDR blocks (contains ":")
	if strings.Contains(rule.SrcCIDR, ":") || strings.Contains(rule.DstCIDR, ":") {
		return true
	}

	// Check for IPv6-specific protocols
	protocol := strings.ToLower(rule.Protocol)
	if protocol == "icmpv6" || protocol == "ipv6" {
		return true
	}

	return false
}
