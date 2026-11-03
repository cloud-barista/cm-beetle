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

// Package summary provides source infrastructure NLB summary collection logic
package summary

import (
	"strings"

	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
)

type sourceNodeRef struct {
	machineId string
	hostname  string
}

// collectSourceNlbResources extracts on-premise load balancer configurations and resolves server mappings
func collectSourceNlbResources(infraData *onpremmodel.OnpremInfra) []SourceNlbInfo {
	if infraData == nil || len(infraData.NLBs) == 0 {
		return nil
	}

	// Build lookup maps from MachineId and IPv4 address to source node identity.
	machineIdToHost := make(map[string]string, len(infraData.Nodes))
	ipToNode := make(map[string]sourceNodeRef)
	for _, node := range infraData.Nodes {
		machineIdToHost[node.MachineId] = node.Hostname
		for _, iface := range node.Interfaces {
			for _, cidr := range iface.IPv4CidrBlocks {
				ip := strings.TrimSpace(strings.Split(cidr, "/")[0])
				if ip != "" {
					ipToNode[ip] = sourceNodeRef{
						machineId: node.MachineId,
						hostname:  node.Hostname,
					}
				}
			}
		}
	}

	nlbs := make([]SourceNlbInfo, 0, len(infraData.NLBs))
	for _, nlb := range infraData.NLBs {
		backendServers := make([]SourceNlbBackendServerInfo, 0, len(nlb.Backend.Servers))
		for _, srv := range nlb.Backend.Servers {
			cleanIP := strings.TrimSpace(strings.Split(srv.IP, "/")[0])
			ref := ipToNode[cleanIP]
			backendServers = append(backendServers, SourceNlbBackendServerInfo{
				Name:      srv.Name,
				IP:        srv.IP,
				Port:      srv.Port,
				Weight:    srv.Weight,
				MachineId: ref.machineId,
				Hostname:  ref.hostname,
			})
		}

		nlbs = append(nlbs, SourceNlbInfo{
			Software:       nlb.Software,
			HostMachineId:  nlb.HostMachineId,
			HostHostname:   machineIdToHost[nlb.HostMachineId],
			BindAddress:    nlb.Listener.BindAddress,
			ListenerPort:   nlb.Listener.Port,
			ListenerProto:  nlb.Listener.Protocol,
			BackendName:    nlb.Backend.Name,
			BackendProto:   nlb.Backend.Protocol,
			BackendPort:    resolveSourceBackendPort(nlb.Backend.Servers),
			Balance:        nlb.Backend.Balance,
			BackendServers: backendServers,
			HealthCheck: SourceNlbHealthCheckInfo{
				Enabled:   nlb.HealthCheck.Enabled,
				Protocol:  nlb.HealthCheck.Protocol,
				Port:      nlb.HealthCheck.Port,
				Interval:  nlb.HealthCheck.Interval,
				Timeout:   nlb.HealthCheck.Timeout,
				Threshold: nlb.HealthCheck.Threshold,
			},
		})
	}

	return nlbs
}

// resolveSourceBackendPort selects a representative backend port using majority vote, then highest weight, then first server
func resolveSourceBackendPort(servers []onpremmodel.NlbServerProperty) int {
	if len(servers) == 0 {
		return 0
	}

	type portStats struct {
		count     int
		maxWeight int
		firstIdx  int
	}

	stats := make(map[int]*portStats)
	for idx, srv := range servers {
		if srv.Port == 0 {
			continue
		}
		st, exists := stats[srv.Port]
		if !exists {
			stats[srv.Port] = &portStats{
				count:     1,
				maxWeight: srv.Weight,
				firstIdx:  idx,
			}
			continue
		}
		st.count++
		if srv.Weight > st.maxWeight {
			st.maxWeight = srv.Weight
		}
	}

	var bestPort int
	var best *portStats
	for port, st := range stats {
		if best == nil ||
			st.count > best.count ||
			(st.count == best.count && st.maxWeight > best.maxWeight) ||
			(st.count == best.count && st.maxWeight == best.maxWeight && st.firstIdx < best.firstIdx) {
			bestPort = port
			best = st
		}
	}

	return bestPort
}
