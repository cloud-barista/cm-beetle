export interface ResolvedNodeLineage {
  sourceMachineId: string;
  sourceHostname: string;
  serviceUrl?: string;
  servicePorts?: number[];
}

export function parseDiscreteTcpPorts(rawPorts: string | number | undefined, toPort?: string | number): number[] {
  if (rawPorts === undefined || rawPorts === null) return [];
  const rawStr = String(rawPorts).trim();
  if (!rawStr || rawStr === '*' || rawStr.toLowerCase() === 'all' || rawStr === '-1') return [];

  if (toPort !== undefined && toPort !== null) {
    const toStr = String(toPort).trim();
    if (toStr && toStr !== rawStr && !rawStr.includes(',')) {
      return [];
    }
  }

  const result: number[] = [];
  for (const part of rawStr.split(',')) {
    const token = part.trim();
    if (!token || token.includes('-') || token.includes(':')) continue;
    const port = Number(token);
    if (Number.isInteger(port) && port >= 1 && port <= 65535 && port !== 22) {
      result.push(port);
    }
  }
  return result;
}

function extractServicePortsFromServer(server: any): number[] {
  if (!server) return [];
  const ports = new Set<number>();

  if (typeof server.serviceUrl === 'string' && server.serviceUrl.trim()) {
    const match = server.serviceUrl.match(/:(\d{2,5})(?:\/|$)/);
    if (match) {
      const p = Number(match[1]);
      if (Number.isInteger(p) && p >= 1 && p <= 65535 && p !== 22) {
        ports.add(p);
      }
    }
  }

  if (Array.isArray(server.firewallTable)) {
    for (const rule of server.firewallTable) {
      const dir = String(rule?.direction || 'inbound').toLowerCase();
      const proto = String(rule?.protocol || 'tcp').toLowerCase();
      if (dir === 'outbound' || (proto !== 'tcp' && proto !== 'all' && proto !== '*')) continue;
      for (const p of parseDiscreteTcpPorts(rule?.dstPorts ?? rule?.fromPort, rule?.toPort)) {
        ports.add(p);
      }
    }
  }

  return Array.from(ports).sort((a, b) => a - b);
}

function stripLegacyDecomposedSuffix(machineId: string): string {
  for (const sep of ['-primary-', '-gpu-']) {
    const idx = machineId.lastIndexOf(sep);
    if (idx > 0) {
      return machineId.slice(0, idx);
    }
  }
  return machineId;
}

function parseVmOrdinalIndex(vmName: string): number {
  const match = vmName.match(/-(\d+)$/);
  if (!match) return -1;
  const n = parseInt(match[1], 10);
  return Number.isNaN(n) || n < 1 ? -1 : n - 1;
}

function extractSourceServers(refinedSourceInfra: any): any[] {
  if (!refinedSourceInfra) return [];
  const root = refinedSourceInfra.onpremiseInfraModel || refinedSourceInfra;
  if (Array.isArray(root.servers)) return root.servers;
  if (Array.isArray(root.nodes)) return root.nodes;
  return [];
}

export function resolveTargetNodeLineage(
  node: any,
  nodeGroups?: any[],
  refinedSourceInfra?: any
): ResolvedNodeLineage | null {
  if (!node) return null;

  const nodeName = String(node.name || node.id || '');
  const subGroupId = String(node.subGroupId || node.nodeGroupId || '');

  let rawLabelCsv =
    node.label?.['cm-source-machine-ids'] ||
    node.label?.sourceMachineIds ||
    '';

  if (!rawLabelCsv && Array.isArray(nodeGroups)) {
    const matchedNg = nodeGroups.find(
      (ng: any) =>
        ng?.name &&
        (ng.name === subGroupId ||
          ng.name === nodeName ||
          nodeName.startsWith(`${ng.name}-`))
    );
    rawLabelCsv =
      matchedNg?.label?.['cm-source-machine-ids'] ||
      matchedNg?.label?.sourceMachineIds ||
      '';
  }

  let sourceMachineId = '';
  if (rawLabelCsv) {
    const ids = String(rawLabelCsv)
      .split(',')
      .map((s) => s.trim())
      .filter(Boolean);
    if (ids.length === 1) {
      sourceMachineId = stripLegacyDecomposedSuffix(ids[0]);
    } else if (ids.length > 1) {
      const idx = parseVmOrdinalIndex(nodeName);
      const chosen = idx >= 0 && idx < ids.length ? ids[idx] : ids[0];
      sourceMachineId = stripLegacyDecomposedSuffix(chosen);
    }
  }

  const servers = extractSourceServers(refinedSourceInfra);

  if (sourceMachineId) {
    const matchedServer = servers.find(
      (s: any) =>
        String(s.machineId || '').toLowerCase() === sourceMachineId.toLowerCase()
    );
    return {
      sourceMachineId,
      sourceHostname: matchedServer?.hostname || '',
      serviceUrl: matchedServer?.serviceUrl || undefined,
      servicePorts: extractServicePortsFromServer(matchedServer),
    };
  }

  // Fallback for nodes without labels: match against source servers by machineId or hostname token
  for (const s of servers) {
    const mid = String(s.machineId || '');
    const cleanAlnum = mid.replace(/[^a-zA-Z0-9]/g, '').toLowerCase();
    const short8 = cleanAlnum.length > 8 ? cleanAlnum.slice(-8) : cleanAlnum;
    if (
      (mid && nodeName.toLowerCase().includes(mid.toLowerCase())) ||
      (short8.length >= 6 && nodeName.toLowerCase().includes(short8))
    ) {
      return {
        sourceMachineId: mid,
        sourceHostname: s.hostname || '',
        serviceUrl: s.serviceUrl || undefined,
        servicePorts: extractServicePortsFromServer(s),
      };
    }
  }

  return null;
}
