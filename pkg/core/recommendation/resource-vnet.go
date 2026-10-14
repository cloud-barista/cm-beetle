package recommendation

import (
	"fmt"
	"math/big"
	"net"

	"github.com/cloud-barista/cb-tumblebug/src/core/common/netutil"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/rs/zerolog/log"
)

const (
	// DefaultTargetVNetCidr is the default target VNet CIDR block recommended when source network info is absent
	DefaultTargetVNetCidr = "10.0.0.0/16"
	// DefaultTargetSubnetCidr is the default target Subnet CIDR block recommended when source network info is absent
	DefaultTargetSubnetCidr = "10.0.1.0/24"
)

// buildDefaultVNet creates a standard default target virtual network (10.0.0.0/16, subnet 10.0.1.0/24).
func buildDefaultVNet(csp, region string) cloudmodel.VNetReq {
	return cloudmodel.VNetReq{
		Name:           "INSERT_YOUR_VNET_NAME",
		ConnectionName: fmt.Sprintf("%s-%s", csp, region),
		Description:    fmt.Sprintf("Default target virtual network (%s) for migration", DefaultTargetVNetCidr),
		CidrBlock:      DefaultTargetVNetCidr,
		SubnetInfoList: []cloudmodel.SubnetReq{
			{
				Name:        "INSERT_YOUR_SUBNET_NAME",
				Description: fmt.Sprintf("Default target subnet (%s) for migration", DefaultTargetSubnetCidr),
				IPv4_CIDR:   DefaultTargetSubnetCidr,
			},
		},
	}
}

// RecommendVNet recommends virtual network configurations based on source infrastructure
func RecommendVNet(csp string, region string, srcInfra onpremmodel.OnpremInfra) ([]cloudmodel.VNetReq, error) {

	var emptyRes []cloudmodel.VNetReq
	var recommendedVNets []cloudmodel.VNetReq

	// [Input]
	ok, err := IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msgf("invalid csp (%s) or region (%s)", csp, region)
		return emptyRes, err
	}

	// When source network info is absent, recommend a standard default target VNet (10.0.0.0/16, subnet 10.0.1.0/24)
	if len(srcInfra.Network.IPv4Networks.CidrBlocks) == 0 && len(srcInfra.Network.IPv4Networks.DefaultGateways) == 0 {
		log.Info().Msgf("no network information found in source infrastructure; recommending default target network (%s, subnet %s)", DefaultTargetVNetCidr, DefaultTargetSubnetCidr)
		return []cloudmodel.VNetReq{buildDefaultVNet(csp, region)}, nil
	}

	var srcNetworks []string
	if len(srcInfra.Network.IPv4Networks.CidrBlocks) != 0 {
		srcNetworks = srcInfra.Network.IPv4Networks.CidrBlocks
	} else if len(srcInfra.Network.IPv4Networks.DefaultGateways) != 0 {
		// Source networks are derived by combining default gateway and network interfaces
		var deriveErr error
		srcNetworks, deriveErr = deriveSourceNetworksFromDefaultGateways(srcInfra)
		if deriveErr != nil || len(srcNetworks) == 0 {
			log.Warn().Err(deriveErr).Msgf("failed to derive CIDR from gateways; recommending default target network (%s, subnet %s)", DefaultTargetVNetCidr, DefaultTargetSubnetCidr)
			return []cloudmodel.VNetReq{buildDefaultVNet(csp, region)}, nil
		}
	}
	log.Debug().Msgf("Source networks (CIDR blocks): %v", srcNetworks)

	// [Process] Recommend the vNet and subnets
	// * Note:
	// * At least 1 subnet is required.
	// * Derive a super network that includes user's all networks and set it as a vNet
	// * Set user's networks as subnets

	// ? Assumption: a network in on-premise infrastructure is designed and configured with various network segments or types.
	// * Thus, it must be selected which of these network segments will be the vNet.
	// ? If so, is grouping the network segments required?

	// Categorizes the entered CIDR blocks by private network (i.e., 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16)
	var cidrs10 []string
	var cidrs172 []string
	var cidrs192 []string

	for _, srcNetwork := range srcNetworks {
		identifiedNet, err := netutil.WhichPrivateNetworkByCidr(srcNetwork)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to identify the network %s", srcNetwork)
			continue
		}
		log.Debug().Msgf("identified network: %s", identifiedNet)

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
	log.Debug().Msgf("CIDR blocks for %s: %v", netutil.PrivateNetwork10Dot, cidrs10)
	log.Debug().Msgf("CIDR blocks for %s: %v", netutil.PrivateNetwork172Dot, cidrs172)
	log.Debug().Msgf("CIDR blocks for %s: %v", netutil.PrivateNetwork192Dot, cidrs192)

	// Build recommended VNet for each private network group using CSP network constraints
	estimateNumNetworks := 4
	groups := []struct {
		name  string
		cidrs []string
	}{
		{netutil.PrivateNetwork10Dot, cidrs10},
		{netutil.PrivateNetwork172Dot, cidrs172},
		{netutil.PrivateNetwork192Dot, cidrs192},
	}

	for _, g := range groups {
		if len(g.cidrs) == 0 {
			continue
		}
		vnet, err := buildRecommendedVNet(csp, region, g.name, g.cidrs, estimateNumNetworks)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to build recommended vNet for %s in %s", g.name, csp)
			continue
		}
		if vnet != nil {
			recommendedVNets = append(recommendedVNets, *vnet)
		}
	}

	// [Output]
	if len(recommendedVNets) == 0 {
		return emptyRes, fmt.Errorf("no recommended virtual network found for %s: check if source networks conform to CSP network constraints", csp)
	}

	return recommendedVNets, nil
}

// buildRecommendedVNet creates a VNet recommendation for a single private network group under CSP constraints.
func buildRecommendedVNet(csp, region, privateNetName string, cidrs []string, estimateNumNetworks int) (*cloudmodel.VNetReq, error) {
	if len(cidrs) == 0 {
		return nil, nil
	}

	supernet, err := netutil.CalculateSupernet(cidrs)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate supernet for %s: %w", privateNetName, err)
	}

	adjustedVNetCidr, note, err := AdjustVNetCidrForCsp(csp, supernet, estimateNumNetworks)
	if err != nil {
		return nil, fmt.Errorf("failed to adjust vNet CIDR for %s (%s): %w", csp, privateNetName, err)
	}

	tempSubnets := make([]cloudmodel.SubnetReq, 0, len(cidrs))
	for _, cidr := range cidrs {
		networkAddr, err := toNetworkAddress(cidr)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to parse CIDR block %s", cidr)
			continue
		}

		adjustedSubnet, err := AdjustSubnetCidrForCsp(csp, adjustedVNetCidr, networkAddr)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to adjust subnet CIDR %s for %s; skipping", networkAddr, csp)
			continue
		}

		tempSubnets = append(tempSubnets, cloudmodel.SubnetReq{
			Name:        "INSERT_YOUR_SUBNET_NAME",
			Description: "subnet from source computing infra",
			IPv4_CIDR:   adjustedSubnet,
		})
	}

	if len(tempSubnets) == 0 {
		return nil, fmt.Errorf("no valid subnets could be recommended for %s in %s", privateNetName, csp)
	}

	vnetDesc := "Recommended vNet for " + privateNetName
	if note != "" {
		vnetDesc += " (" + note + ")"
	}

	vnet := &cloudmodel.VNetReq{
		Name:           "INSERT_YOUR_VNET_NAME",
		ConnectionName: fmt.Sprintf("%s-%s", csp, region),
		Description:    vnetDesc,
		CidrBlock:      adjustedVNetCidr,
		SubnetInfoList: tempSubnets,
	}

	return vnet, nil
}

func deriveSourceNetworksFromDefaultGateways(srcInfra onpremmodel.OnpremInfra) ([]string, error) {
	if len(srcInfra.Network.IPv4Networks.DefaultGateways) == 0 {
		return nil, fmt.Errorf("no network information found in the source computing infrastructure")
	}

	var sourceNetworks []string
	// 1. Find the node that has the same "machine ID" as the gateway
	for _, gateway := range srcInfra.Network.IPv4Networks.DefaultGateways {
		for _, node := range srcInfra.Nodes {
			if node.MachineId == gateway.MachineId {

				// 2. Find the network interface that has the same network "name" as the gateway
				for _, nic := range node.Interfaces {
					if nic.Name == gateway.InterfaceName {

						// 3. Get "prefix length" from the network interface
						if nic.IPv4CidrBlocks == nil && len(nic.IPv4CidrBlocks) == 0 {
							log.Warn().Msgf("no IPv4 CIDR blocks found in the network interface %s of the node %s", nic.Name, node.MachineId)
							continue
						}

						cidrBlock := nic.IPv4CidrBlocks[0]
						_, ipNet, err := net.ParseCIDR(cidrBlock)
						if err != nil {
							log.Warn().Err(err).Msgf("failed to parse CIDR block %s", cidrBlock)
							continue
						}

						prefixLen, _ := ipNet.Mask.Size()

						// 4. Derive the CIDR block from the gateway IP and prefix length
						gatewayCIDR := fmt.Sprintf("%s/%d", gateway.IP, prefixLen)

						// 5. Append the derived CIDR block to the list
						sourceNetworks = append(sourceNetworks, gatewayCIDR)
					}
				}
			}
		}
	}

	// Deduplicate the source networks
	sourceNetworks = deduplicateSlice(sourceNetworks)

	return sourceNetworks, nil
}

func deduplicateSlice[T comparable](slice []T) []T {
	// Create a map to track unique elements
	uniqueMap := make(map[T]struct{})
	for _, item := range slice {
		uniqueMap[item] = struct{}{}
	}

	// Convert the map keys back to a slice
	result := make([]T, 0, len(uniqueMap))
	for item := range uniqueMap {
		result = append(result, item)
	}
	return result
}

// estimateSupernet finds the smallest supernet that contains a given number
// of consecutive networks, starting from a given CIDR.
func estimateSupernet(startCIDR string, numNetworks int) (string, error) {
	// 1. Parse the starting CIDR.
	ip, ipNet, err := net.ParseCIDR(startCIDR)
	if err != nil {
		return "", fmt.Errorf("invalid CIDR: %v", err)
	}

	// Ensure it's an IPv4 address.
	ipv4 := ip.To4()
	if ipv4 == nil {
		return "", fmt.Errorf("only IPv4 addresses are supported")
	}

	// 2. Calculate the total IP range.
	// Number of addresses in the start network (e.g., /24 -> 256).
	prefixLen, bits := ipNet.Mask.Size()
	numAddressesPerNet := 1 << (bits - prefixLen)

	// Total number of addresses to cover.
	totalAddresses := numAddressesPerNet * numNetworks

	// Convert the starting IP to an integer for calculation.
	startIPint := big.NewInt(0)
	startIPint.SetBytes(ipv4)

	// Calculate the last IP address in the entire range.
	// Last IP = Start IP + Total Addresses - 1.
	offset := big.NewInt(int64(totalAddresses - 1))
	endIPint := big.NewInt(0)
	endIPint.Add(startIPint, offset)

	// Convert the integer back to a net.IP.
	firstIP := ipv4
	lastIP := net.IP(endIPint.Bytes())

	// 3. Find the common supernet.
	// Iterate from the initial prefix down to 0, finding the first prefix
	// length where both the first and last IPs belong to the same network.
	for newPrefixLen := prefixLen; newPrefixLen >= 0; newPrefixLen-- {
		mask := net.CIDRMask(newPrefixLen, bits)
		network1 := firstIP.Mask(mask)
		network2 := lastIP.Mask(mask)

		// If both IPs belong to the same network, we've found our supernet.
		if network1.Equal(network2) {
			return (&net.IPNet{IP: network1, Mask: mask}).String(), nil
		}
	}

	return "", fmt.Errorf("could not find a common supernet")
}

func toNetworkAddress(cidr string) (string, error) {
	_, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", fmt.Errorf("failed to parse CIDR block %s: %v", cidr, err)
	}
	return subnet.String(), nil
}
