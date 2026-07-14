'use client';

import React, { useMemo } from 'react';
import { ReactFlow, Background, Controls, Edge, Node } from '@xyflow/react';
import { RecommendedInfra } from '../../types/migration';
import '@xyflow/react/dist/style.css';

interface TopologyMapProps {
  data: RecommendedInfra | null;
}

export const TopologyMap: React.FC<TopologyMapProps> = ({ data }) => {
  const { nodes, edges } = useMemo(() => {
    const nodesList: Node[] = [];
    const edgesList: Edge[] = [];

    if (!data) return { nodes: nodesList, edges: edgesList };

    const { targetVNet, targetSshKey, targetSecurityGroupList, targetInfra, targetNlbList } = data;

    // 1. Render VNet parent box
    const vNetId = 'vnet-node';
    nodesList.push({
      id: vNetId,
      type: 'group',
      data: { label: `${targetVNet.name} (${targetVNet.cidrBlock})` },
      position: { x: 50, y: 150 },
      style: {
        width: 600,
        height: 380,
        backgroundColor: 'rgba(6, 182, 212, 0.05)',
        border: '1px dashed var(--color-brand-cyan)',
        borderRadius: '16px',
      },
    });

    // 2. Render Subnet nested parent box inside VNet
    const subnet = targetVNet.subnetInfoList[0];
    const subnetId = 'subnet-node';
    if (subnet) {
      nodesList.push({
        id: subnetId,
        parentId: vNetId,
        type: 'group',
        data: { label: `${subnet.name} (${subnet.ipv4_CIDR})` },
        position: { x: 30, y: 50 },
        style: {
          width: 540,
          height: 300,
          backgroundColor: 'rgba(14, 165, 233, 0.08)',
          border: '1px solid var(--color-brand-blue)',
          borderRadius: '12px',
        },
      });
    }

    // 3. Render SSH Key node (outside VNet, usually top)
    const sshNodeId = 'ssh-node';
    nodesList.push({
      id: sshNodeId,
      data: { label: `🔑 SSH Key:\n${targetSshKey.name}` },
      position: { x: 300, y: 30 },
      style: {
        background: 'var(--color-bg-panel)',
        color: 'var(--color-brand-purple)',
        border: '1px solid var(--color-brand-purple)',
        borderRadius: '8px',
        padding: '10px',
        fontSize: '11px',
        width: 140,
        textAlign: 'center',
        whiteSpace: 'pre-wrap',
        boxShadow: '0 0 10px rgba(129, 140, 248, 0.15)',
      },
    });

    // 4. Render VM NodeGroups inside Subnet container
    const nodeGroups = targetInfra.nodeGroups;
    nodeGroups.forEach((ng, index) => {
      const vmNodeId = `vm-ng-${index}`;

      // Calculate layout coordinates inside Subnet parent
      const xPos = 40 + index * 260;
      const yPos = 80;

      nodesList.push({
        id: vmNodeId,
        parentId: subnetId,
        data: {
          label: `🖥️ NodeGroup: ${ng.name}\nSize: ${ng.nodeGroupSize} | Spec: ${ng.specId}\nImg: ${ng.imageId.slice(0, 18)}...`,
        },
        position: { x: xPos, y: yPos },
        style: {
          background: 'var(--color-bg-input)',
          color: 'var(--color-text-main)',
          border: '1px solid var(--color-border-main)',
          borderRadius: '10px',
          padding: '12px',
          fontSize: '11px',
          width: 200,
          lineHeight: '1.4',
          boxShadow: '0 4px 6px -1px rgba(0, 0, 0, 0.1)',
        },
      });

      // Edge from SSH Key to VM
      edgesList.push({
        id: `edge-ssh-vm-${index}`,
        source: sshNodeId,
        target: vmNodeId,
        type: 'smoothstep',
        style: { stroke: 'var(--color-brand-purple)', strokeWidth: 1.5, strokeDasharray: '4 4' },
      });
    });

    // 5. Render Security Groups nodes (placed side-by-side)
    targetSecurityGroupList.forEach((sg, idx) => {
      const sgNodeId = `sg-node-${idx}`;

      // Position SG nodes to the left/right of VNet
      const xPos = 700;
      const yPos = idx === 0 ? 180 : 320;

      // Extract ports list for rendering
      const ports = sg.firewallRules.map(rule => `${rule.protocol}:${rule.dstPorts}`).join(', ');

      nodesList.push({
        id: sgNodeId,
        data: { label: `🛡️ SG: ${sg.name}\nRules: ${ports}` },
        position: { x: xPos, y: yPos },
        style: {
          background: 'var(--color-bg-panel)',
          color: 'var(--color-brand-cyan)',
          border: '1px solid var(--color-brand-cyan)',
          borderRadius: '8px',
          padding: '10px',
          fontSize: '11px',
          width: 160,
          textAlign: 'center',
          whiteSpace: 'pre-wrap',
          boxShadow: '0 0 10px rgba(6, 182, 212, 0.1)',
        },
      });

      // Draw connection edges from Security Groups to matching VM NodeGroups
      nodeGroups.forEach((ng, ngIdx) => {
        if (ng.securityGroupIds.includes(sg.name) || ng.securityGroupIds.some(id => id.includes(sg.name))) {
          edgesList.push({
            id: `edge-sg-${idx}-vm-${ngIdx}`,
            source: sgNodeId,
            target: `vm-ng-${ngIdx}`,
            type: 'smoothstep',
            animated: true,
            style: { stroke: 'var(--color-brand-cyan)', strokeWidth: 2 },
          });
        }
      });
    });

    // 6. Render Public/Private NLB (placed left of VNet container)
    if (targetNlbList && targetNlbList.length > 0) {
      targetNlbList.forEach((nlb, nlbIdx) => {
        const nlbNodeId = `nlb-node-${nlbIdx}`;

        // Position NLB left of the subnet VNet container
        const xPos = 700;
        const yPos = 50;

        nodesList.push({
          id: nlbNodeId,
          data: {
            label: `⚖️ NLB: ${nlb.type}\nListens: ${nlb.listener.protocol}:${nlb.listener.port}\nTargets: ${nlb.targetGroup.nodeGroupId}`,
          },
          position: { x: xPos, y: yPos },
          style: {
            background: 'var(--color-bg-panel)',
            color: 'var(--color-brand-purple)',
            border: '1px solid var(--color-brand-purple)',
            borderRadius: '8px',
            padding: '10px',
            fontSize: '11px',
            width: 160,
            textAlign: 'center',
            whiteSpace: 'pre-wrap',
            boxShadow: '0 0 10px rgba(217, 70, 239, 0.15)',
          },
        });

        // Edge from Load Balancer target group to VM NodeGroups
        nodeGroups.forEach((ng, ngIdx) => {
          if (ng.name.includes(nlb.targetGroup.nodeGroupId) || nlb.targetGroup.nodeGroupId.includes(ng.name)) {
            edgesList.push({
              id: `edge-nlb-${nlbIdx}-vm-${ngIdx}`,
              source: nlbNodeId,
              target: `vm-ng-${ngIdx}`,
              type: 'smoothstep',
              animated: true,
              style: { stroke: 'var(--color-brand-purple)', strokeWidth: 2 },
            });
          }
        });
      });
    }

    return { nodes: nodesList, edges: edgesList };
  }, [data]);

  return (
    <div className="w-full h-full bg-bg-panel border border-border-main rounded-xl overflow-hidden relative">
      {data ? (
        <ReactFlow
          nodes={nodes}
          edges={edges}
          fitView
          attributionPosition="bottom-left"
          minZoom={0.5}
          maxZoom={1.5}
        >
          <Background color="var(--color-border-main)" gap={16} size={1} />
          <Controls className="bg-bg-input border border-border-main text-text-muted [&_button]:border-border-main [&_button]:bg-bg-input" />
        </ReactFlow>
      ) : (
        <div className="absolute inset-0 flex items-center justify-center text-xs text-text-muted">
          No cloud architecture recommendation loaded.
        </div>
      )}
    </div>
  );
};
