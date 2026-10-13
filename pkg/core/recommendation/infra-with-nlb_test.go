package recommendation

import (
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/stretchr/testify/assert"
)

func TestSupplementFirewallForRemoteAndNlb_NhnCloud_DifferentPorts(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "node-01",
		FirewallTable: []onpremmodel.FirewallRuleProperty{
			{
				Action:    "allow",
				Direction: "inbound",
				Protocol:  "tcp",
				SrcCIDR:   "0.0.0.0/0",
				DstPorts:  "22",
			},
		},
	}

	rnlb := resolvedNlb{
		sourceNlb: onpremmodel.NlbProperty{
			Listener: onpremmodel.NlbListenerProperty{
				Port:     9999,
				Protocol: "tcp",
			},
		},
		backendPort: 8086,
	}

	skeleton := cloudmodel.RecommendedInfra{
		TargetVNet: cloudmodel.VNetReq{
			CidrBlock: "10.0.0.0/16",
			SubnetInfoList: []cloudmodel.SubnetReq{
				{
					IPv4_CIDR: "10.0.1.0/24",
				},
			},
		},
	}

	// First supplement remote client rules, then supplement NLB/subnet rules.
	result := supplementFirewallIncomingFromRemote(node, rnlb, "nhn")
	result = supplementFirewallNlbToNode(result, rnlb, "nhn", skeleton)

	hasRule := func(port string, cidr string) bool {
		for _, r := range result.FirewallTable {
			if r.Direction == "inbound" && r.Action == "allow" && r.Protocol == "tcp" && r.DstPorts == port && r.SrcCIDR == cidr {
				return true
			}
		}
		return false
	}

	// Backend port 8086 rules (both remote pass-through and NLB-to-VM subnet traffic)
	assert.True(t, hasRule("8086", "0.0.0.0/0"))
	assert.True(t, hasRule("8086", "10.0.0.0/16"))
	assert.True(t, hasRule("8086", "10.0.1.0/24"))
	// Incoming listener port 9999 rules (remote client access)
	assert.True(t, hasRule("9999", "0.0.0.0/0"))
	assert.False(t, hasRule("9999", "10.0.0.0/16"))
	assert.False(t, hasRule("9999", "10.0.1.0/24"))
}

func TestSupplementFirewallForRemoteAndNlb_NhnCloud_IdenticalPorts(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "node-01",
		FirewallTable: []onpremmodel.FirewallRuleProperty{
			{
				Action:    "allow",
				Direction: "inbound",
				Protocol:  "tcp",
				SrcCIDR:   "0.0.0.0/0",
				DstPorts:  "22",
			},
		},
	}

	rnlb := resolvedNlb{
		sourceNlb: onpremmodel.NlbProperty{
			Listener: onpremmodel.NlbListenerProperty{
				Port:     8086,
				Protocol: "tcp",
			},
		},
		backendPort: 8086,
	}

	skeleton := cloudmodel.RecommendedInfra{
		TargetVNet: cloudmodel.VNetReq{
			CidrBlock: "10.0.0.0/16",
			SubnetInfoList: []cloudmodel.SubnetReq{
				{
					IPv4_CIDR: "10.0.1.0/24",
				},
			},
		},
	}

	// Incoming port and backend port are identical (8086).
	result := supplementFirewallIncomingFromRemote(node, rnlb, "nhn")
	result = supplementFirewallNlbToNode(result, rnlb, "nhn", skeleton)

	hasRule := func(port string, cidr string) bool {
		for _, r := range result.FirewallTable {
			if r.Direction == "inbound" && r.Action == "allow" && r.Protocol == "tcp" && r.DstPorts == port && r.SrcCIDR == cidr {
				return true
			}
		}
		return false
	}

	assert.True(t, hasRule("8086", "0.0.0.0/0"))
	assert.True(t, hasRule("8086", "10.0.0.0/16"))
	assert.True(t, hasRule("8086", "10.0.1.0/24"))
	assert.False(t, hasRule("9999", "0.0.0.0/0"))
}

func TestSupplementFirewallForRemoteAndNlb_OtherCsp(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "node-01",
		FirewallTable: []onpremmodel.FirewallRuleProperty{
			{
				Action:    "allow",
				Direction: "inbound",
				Protocol:  "tcp",
				SrcCIDR:   "0.0.0.0/0",
				DstPorts:  "22",
			},
		},
	}

	rnlb := resolvedNlb{
		sourceNlb: onpremmodel.NlbProperty{
			Listener: onpremmodel.NlbListenerProperty{
				Port:     9999,
				Protocol: "tcp",
			},
		},
		backendPort: 8086,
	}

	skeleton := cloudmodel.RecommendedInfra{
		TargetVNet: cloudmodel.VNetReq{
			CidrBlock: "10.0.0.0/16",
			SubnetInfoList: []cloudmodel.SubnetReq{
				{
					IPv4_CIDR: "10.0.1.0/24",
				},
			},
		},
	}

	// For AWS, public backend port is added and NHN-specific NLB rule is a no-op.
	result := supplementFirewallIncomingFromRemote(node, rnlb, "aws")
	result = supplementFirewallNlbToNode(result, rnlb, "aws", skeleton)

	hasRule := func(port string, cidr string) bool {
		for _, r := range result.FirewallTable {
			if r.Direction == "inbound" && r.Action == "allow" && r.Protocol == "tcp" && r.DstPorts == port && r.SrcCIDR == cidr {
				return true
			}
		}
		return false
	}

	assert.True(t, hasRule("8086", "0.0.0.0/0"))
	assert.False(t, hasRule("8086", "10.0.0.0/16"))
	assert.True(t, hasRule("9999", "0.0.0.0/0"))
}
