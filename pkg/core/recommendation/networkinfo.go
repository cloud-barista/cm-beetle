/*
Copyright 2024 The Cloud-Barista Authors.
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

package recommendation

import (
	_ "embed"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

//go:embed assets/networkinfo.yaml
var embeddedNetworkInfoYAML []byte

// NetworkInfoConfig holds the top-level network configuration across CSPs.
type NetworkInfoConfig struct {
	Network map[string]CspNetworkConfig `yaml:"network"`
}

// CspNetworkConfig describes network characteristics and constraints for a single CSP.
type CspNetworkConfig struct {
	Description         string            `yaml:"description"`
	Links               []string          `yaml:"link"`
	AvailableCidrBlocks []CidrBlockConfig `yaml:"available-cidr-blocks"`
	ReservedCidrBlocks  []CidrBlockConfig `yaml:"reserved-cidr-blocks"`
	VNet                *VNetConfig       `yaml:"vnet"`
	Subnet              *SubnetConfig     `yaml:"subnet"`
	VPN                 *VpnConfig        `yaml:"vpn"`
}

// CidrBlockConfig holds an IP range and its description.
type CidrBlockConfig struct {
	CidrBlock   string `yaml:"cidr-block"`
	Description string `yaml:"description"`
}

// VNetConfig describes VNet CIDR prefix length limits for a CSP.
type VNetConfig struct {
	PrefixLength PrefixLengthConfig `yaml:"prefix-length"`
}

// SubnetConfig describes Subnet CIDR prefix length limits and reserved IPs for a CSP.
type SubnetConfig struct {
	PrefixLength PrefixLengthConfig `yaml:"prefix-length"`
	ReservedIPs  ReservedIPsConfig  `yaml:"reserved-ips"`
}

// PrefixLengthConfig defines min and max CIDR prefix lengths.
type PrefixLengthConfig struct {
	Min         int    `yaml:"min"`
	Max         int    `yaml:"max"`
	Description string `yaml:"description"`
}

// ReservedIPsConfig captures the count or value of reserved IP addresses.
type ReservedIPsConfig struct {
	Count       int    `yaml:"count"`
	Value       int    `yaml:"value"`
	Description string `yaml:"description"`
}

// VpnConfig describes VPN characteristics such as gateway subnet requirements.
type VpnConfig struct {
	GatewaySubnet *GatewaySubnetConfig `yaml:"gateway-subnet"`
}

// GatewaySubnetConfig describes requirements for gateway subnets.
type GatewaySubnetConfig struct {
	Required     bool               `yaml:"required"`
	Name         string             `yaml:"name"`
	PrefixLength PrefixLengthConfig `yaml:"prefix-length"`
}

var (
	networkInfoOnce   sync.Once
	cachedNetworkInfo *NetworkInfoConfig
	loadNetworkErr    error
)

// getNetworkInfoConfig returns the cached CSP network configuration loaded from embedded assets.
func getNetworkInfoConfig() (*NetworkInfoConfig, error) {
	networkInfoOnce.Do(func() {
		var cfg NetworkInfoConfig
		if err := yaml.Unmarshal(embeddedNetworkInfoYAML, &cfg); err != nil {
			loadNetworkErr = fmt.Errorf("failed to unmarshal networkinfo.yaml: %w", err)
			return
		}
		cachedNetworkInfo = &cfg
	})
	return cachedNetworkInfo, loadNetworkErr
}

// GetCspNetworkConfig retrieves network constraints for a given CSP name.
func GetCspNetworkConfig(cspName string) (CspNetworkConfig, bool) {
	cfg, err := getNetworkInfoConfig()
	if err != nil {
		log.Error().Err(err).Msg("failed to load network info configuration")
		return CspNetworkConfig{}, false
	}
	item, ok := cfg.Network[strings.ToLower(strings.TrimSpace(cspName))]
	return item, ok
}

// AdjustVNetCidrForCsp validates and clamps a calculated supernet CIDR to satisfy CSP VNet constraints.
// It returns the adjusted CIDR, an informational note if clamped, and any validation error.
func AdjustVNetCidrForCsp(cspName string, rawSupernet string, estimateNetworks int) (string, string, error) {
	ip, ipNet, err := net.ParseCIDR(rawSupernet)
	if err != nil {
		return "", "", fmt.Errorf("invalid CIDR %s: %w", rawSupernet, err)
	}

	cspCfg, ok := GetCspNetworkConfig(cspName)
	if !ok {
		return "", "", fmt.Errorf("unknown CSP %s for network configuration", cspName)
	}

	prefixLen, bits := ipNet.Mask.Size()

	// CSP has no VNet prefix limit (e.g. GCP defines CIDR at subnet level only)
	if cspCfg.VNet == nil {
		if estimateNetworks > 1 {
			est, err := estimateSupernet(rawSupernet, estimateNetworks)
			return est, "", err
		}
		return rawSupernet, "", nil
	}

	minPrefix := cspCfg.VNet.PrefixLength.Min
	maxPrefix := cspCfg.VNet.PrefixLength.Max
	var note string

	// Clamp prefix length when supernet is wider than CSP allows (e.g., /14 clamped to /16)
	targetPrefix := prefixLen
	if minPrefix > 0 && prefixLen < minPrefix {
		targetPrefix = minPrefix
		note = fmt.Sprintf("Note: Clamped from /%d to /%d to comply with %s maximum VPC capacity", prefixLen, minPrefix, cspName)
		log.Info().Msgf("Clamped supernet %s (prefix /%d) to /%d for %s", rawSupernet, prefixLen, minPrefix, cspName)
	} else if estimateNetworks > 1 {
		estimated, err := estimateSupernet(rawSupernet, estimateNetworks)
		if err == nil {
			if _, estNet, err := net.ParseCIDR(estimated); err == nil {
				targetPrefix, _ = estNet.Mask.Size()
			}
		}
		if minPrefix > 0 && targetPrefix < minPrefix {
			targetPrefix = minPrefix
			note = fmt.Sprintf("Note: Clamped expanded VNet to /%d to comply with %s maximum VPC capacity", minPrefix, cspName)
		}
	}

	// Clamp prefix length to CSP allowable maxPrefix
	if maxPrefix > 0 && targetPrefix > maxPrefix {
		targetPrefix = maxPrefix
		note = fmt.Sprintf("Note: Adjusted VNet to /%d to comply with %s minimum VPC capacity", maxPrefix, cspName)
	}

	// Construct clamped VNet network CIDR
	mask := net.CIDRMask(targetPrefix, bits)
	adjustedVNet := (&net.IPNet{IP: ip.Mask(mask), Mask: mask}).String()

	// Verify adjusted VNet is contained in available CIDR blocks
	if err := validateAvailableCidrBlocks(cspName, cspCfg, adjustedVNet); err != nil {
		return "", "", err
	}

	// Warn if VNet overlaps with CSP reserved CIDR blocks
	checkReservedCidrOverlap(cspName, cspCfg, adjustedVNet, "VNet")

	return adjustedVNet, note, nil
}

// AdjustSubnetCidrForCsp validates and adjusts a subnet CIDR to conform to CSP subnet constraints.
func AdjustSubnetCidrForCsp(cspName string, vnetCidr string, rawSubnetCidr string) (string, error) {
	subIP, subNet, err := net.ParseCIDR(rawSubnetCidr)
	if err != nil {
		return "", fmt.Errorf("invalid subnet CIDR %s: %w", rawSubnetCidr, err)
	}

	cspCfg, ok := GetCspNetworkConfig(cspName)
	if !ok {
		return "", fmt.Errorf("unknown CSP %s for network configuration", cspName)
	}

	subPrefix, bits := subNet.Mask.Size()
	adjustedPrefix := subPrefix

	if cspCfg.Subnet != nil {
		subMin := cspCfg.Subnet.PrefixLength.Min
		subMax := cspCfg.Subnet.PrefixLength.Max

		// Fail fast when subnet is larger than CSP allows
		if subMin > 0 && subPrefix < subMin {
			return "", fmt.Errorf("subnet CIDR %s (prefix /%d) exceeds %s maximum subnet size (minimum prefix length is /%d)",
				rawSubnetCidr, subPrefix, cspName, subMin)
		}

		// Adjust subnets smaller than CSP allows (e.g., /29 -> /28 in NCP/AWS)
		if subMax > 0 && subPrefix > subMax {
			adjustedPrefix = subMax
			log.Info().Msgf("Adjusted subnet %s from /%d to /%d to meet %s constraint (max prefix /%d)",
				rawSubnetCidr, subPrefix, adjustedPrefix, cspName, subMax)
		}
	}

	// Compute adjusted subnet network address
	mask := net.CIDRMask(adjustedPrefix, bits)
	adjustedSubnet := (&net.IPNet{IP: subIP.Mask(mask), Mask: mask}).String()

	// Verify subnet falls entirely inside target VNet (when VNet CIDR is non-empty)
	if vnetCidr != "" {
		_, vNetIpNet, err := net.ParseCIDR(vnetCidr)
		if err == nil {
			vNetPrefix, _ := vNetIpNet.Mask.Size()
			if !vNetIpNet.Contains(subIP) || adjustedPrefix < vNetPrefix {
				return "", fmt.Errorf("subnet CIDR %s is not within vNet CIDR %s", adjustedSubnet, vnetCidr)
			}
		}
	}

	// Warn if subnet overlaps with CSP reserved CIDR blocks
	checkReservedCidrOverlap(cspName, cspCfg, adjustedSubnet, "Subnet")

	return adjustedSubnet, nil
}

// validateAvailableCidrBlocks ensures the CIDR belongs to the CSP's permitted address blocks.
func validateAvailableCidrBlocks(cspName string, cfg CspNetworkConfig, cidr string) error {
	if len(cfg.AvailableCidrBlocks) == 0 {
		return nil
	}

	_, targetNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return fmt.Errorf("invalid CIDR for validation %s: %w", cidr, err)
	}
	targetPrefix, _ := targetNet.Mask.Size()

	for _, block := range cfg.AvailableCidrBlocks {
		_, availNet, err := net.ParseCIDR(block.CidrBlock)
		if err != nil {
			continue
		}
		availPrefix, _ := availNet.Mask.Size()
		if availNet.Contains(targetNet.IP) && availPrefix <= targetPrefix {
			return nil
		}
	}

	return fmt.Errorf("CIDR %s is outside permitted available blocks for %s", cidr, cspName)
}

// checkReservedCidrOverlap logs a warning when a recommended CIDR overlaps with CSP reserved ranges.
func checkReservedCidrOverlap(cspName string, cfg CspNetworkConfig, cidr string, scope string) {
	if len(cfg.ReservedCidrBlocks) == 0 {
		return
	}

	_, checkNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return
	}

	for _, reserved := range cfg.ReservedCidrBlocks {
		_, resNet, err := net.ParseCIDR(reserved.CidrBlock)
		if err != nil {
			continue
		}
		// Overlap occurs if either network contains the other's base IP
		if checkNet.Contains(resNet.IP) || resNet.Contains(checkNet.IP) {
			log.Warn().Msgf("%s CIDR %s in %s overlaps with reserved range %s (%s)",
				scope, cidr, cspName, reserved.CidrBlock, reserved.Description)
		}
	}
}
