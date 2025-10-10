package recommendation

import (
	"fmt"
	"math/big"
	"net"

	"github.com/cloud-barista/cb-tumblebug/src/core/common/netutil"
	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
	"github.com/rs/zerolog/log"
)

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

	if len(srcInfra.Network.IPv4Networks.CidrBlocks) == 0 && len(srcInfra.Network.IPv4Networks.DefaultGateways) == 0 {
		err := fmt.Errorf("no network information found in the source computing infrastructure")
		log.Error().Err(err).Msg("failed to recommend a virtual network")
		return emptyRes, err
	}

	var srcNetworks []string
	// * Note: srcInfra.Network.IPv4Networks.CidrBlocks is the input from the user (e.g., network admin)
	if len(srcInfra.Network.IPv4Networks.CidrBlocks) != 0 {
		srcNetworks = srcInfra.Network.IPv4Networks.CidrBlocks
	} else if len(srcInfra.Network.IPv4Networks.DefaultGateways) != 0 {
		// * Note: To estimate the network address space of the source computing infrastructure,
		// * Source networks are derived by combining the default gateway and network interface information of each server.
		srcNetworks, err = deriveSourceNetworksFromDefaultGateways(srcInfra)
		if err != nil {
			log.Error().Err(err).Msg("failed to derive CIDR blocks from default gateways")
			return emptyRes, err
		}
	} else {
		log.Warn().Msg("no network information found in the source computing infrastructure")
		return emptyRes, fmt.Errorf("no network information found in the source computing infrastructure")
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

	// Estimate the more :D super network for each private network
	// TODO: Set the number of networks to be included in the super network
	estimateNumNetworks := 4
	if len(supernet10) > 0 {
		supernet10, err = estimateSupernet(supernet10, estimateNumNetworks)
		if err != nil {
			log.Warn().Err(err).Msg("failed to estimate supernet for 10.x.x.x")
		}
	}
	if len(supernet172) > 0 {
		supernet172, err = estimateSupernet(supernet172, estimateNumNetworks)
		if err != nil {
			log.Warn().Err(err).Msg("failed to estimate supernet for 172.x.x.x")
		}
	}
	if len(supernet192) > 0 {
		supernet192, err = estimateSupernet(supernet192, estimateNumNetworks)
		if err != nil {
			log.Warn().Err(err).Msg("failed to estimate supernet for 192.x.x.x")
		}
	}

	// Select a super network for the vNet
	// ? But how to select the super network?
	// * Currrently, a list of recommended networks is returned.

	if supernet10 != "" {
		// Set tempSubnets by the CIDR blocks from the source computing infra
		tempSubnets := []cloudmodel.SubnetReq{}
		for _, cidr := range cidrs10 {
			networkAddr, err := toNetworkAddress(cidr)
			if err != nil {
				log.Warn().Err(err).Msgf("failed to parse CIDR block %s", cidr)
				continue
			}

			tempSubnets = append(tempSubnets, cloudmodel.SubnetReq{
				Name:        "INSERT_YOUR_SUBNET_NAME", // TODO: Set a name for the subnet
				Description: "subnet from source computing infra",
				IPv4_CIDR:   networkAddr,
			})
		}

		// Set the calculated supernet as the tempVNet
		tempVNet := cloudmodel.VNetReq{
			Name:           "INSERT_YOUR_VNET_NAME", // TODO: Set a name for the vNet
			ConnectionName: fmt.Sprintf("%s-%s", csp, region),
			Description:    "Recommended vNet for " + netutil.PrivateNetwork10Dot,
			CidrBlock:      supernet10,
			SubnetInfoList: tempSubnets,
		}

		// Append recommended virtual network info to the list
		recommendedVNets = append(recommendedVNets, tempVNet)
	}

	if supernet172 != "" {

		// Set tempSubnets by the CIDR blocks from the source computing infra
		tempSubnets := []cloudmodel.SubnetReq{}
		for _, cidr := range cidrs172 {
			networkAddr, err := toNetworkAddress(cidr)
			if err != nil {
				log.Warn().Err(err).Msgf("failed to parse CIDR block %s", cidr)
				continue
			}

			tempSubnets = append(tempSubnets, cloudmodel.SubnetReq{
				Name:        "INSERT_YOUR_SUBNET_NAME", // TODO: Set a name for the subnet
				Description: "subnet from source computing infra",
				IPv4_CIDR:   networkAddr,
			})
		}

		tempVNet := cloudmodel.VNetReq{
			Name:           "INSERT_YOUR_VNET_NAME", // TODO: Set a name for the vNet
			ConnectionName: fmt.Sprintf("%s-%s", csp, region),
			Description:    "Recommended vNet for " + netutil.PrivateNetwork172Dot,
			CidrBlock:      supernet172,
			SubnetInfoList: tempSubnets,
		}

		// Append recommended virtual network info to the list
		recommendedVNets = append(recommendedVNets, tempVNet)
	}

	if supernet192 != "" {

		// Set tempSubnets by the CIDR blocks from the source computing infra
		tempSubnets := []cloudmodel.SubnetReq{}
		for _, cidr := range cidrs192 {

			networkAddr, err := toNetworkAddress(cidr)
			if err != nil {
				log.Warn().Err(err).Msgf("failed to parse CIDR block %s", cidr)
				continue
			}

			tempSubnets = append(tempSubnets, cloudmodel.SubnetReq{
				Name:        "INSERT_YOUR_SUBNET_NAME", // TODO: Set a name for the subnet
				Description: "subnet from source computing infra",
				IPv4_CIDR:   networkAddr,
			})
		}

		// Set the calculated supernet as the vNet
		tempVNet := cloudmodel.VNetReq{
			Name:           "INSERT_YOUR_VNET_NAME", // TODO: Set a name for the vNet
			ConnectionName: fmt.Sprintf("%s-%s", csp, region),
			Description:    "Recommended vNet for " + netutil.PrivateNetwork192Dot,
			CidrBlock:      supernet192,
			SubnetInfoList: tempSubnets,
		}

		// Append recommended virtual network info to the list
		recommendedVNets = append(recommendedVNets, tempVNet)
	}

	// [Output]
	if len(recommendedVNets) == 0 {
		return emptyRes, fmt.Errorf("no recommended virtual network found for the source computing infra")
	}

	return recommendedVNets, nil
}

func deriveSourceNetworksFromDefaultGateways(srcInfra onpremmodel.OnpremInfra) ([]string, error) {
	if len(srcInfra.Network.IPv4Networks.DefaultGateways) == 0 {
		return nil, fmt.Errorf("no network information found in the source computing infrastructure")
	}

	var sourceNetworks []string
	// 1. Find the server that has the same "machine ID" as the gateway
	for _, gateway := range srcInfra.Network.IPv4Networks.DefaultGateways {
		for _, server := range srcInfra.Servers {
			if server.MachineId == gateway.MachineId {

				// 2. Find the network interface that has the same network "name" as the gateway
				for _, nic := range server.Interfaces {
					if nic.Name == gateway.InterfaceName {

						// 3. Get "prefix length" from the network interface
						if nic.IPv4CidrBlocks == nil && len(nic.IPv4CidrBlocks) == 0 {
							log.Warn().Msgf("no IPv4 CIDR blocks found in the network interface %s of the server %s", nic.Name, server.MachineId)
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
