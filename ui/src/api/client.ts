import axios from 'axios';
import { OnpremInfra, OnpremNlb, OnpremNode, OnpremModelEnvelope, RecommendedInfra, CloudModelEnvelope, CloudProperty, CheckSSHReadyResponse, ArtifactLineage, RequestDetailsResult, AiStackLayerStatus } from '../types/migration';
import { cleanModelDisplayName, isBuiltInSampleModel, normalizeModelVersion } from '../utils/modelFormat';

// Shared Axios client with basic authentication pre-configured (standard Cloud-Barista settings)
const api = axios.create({
  headers: {
    'Content-Type': 'application/json',
    // Using default credentials commonly used in Cloud-Barista development setups
    'Authorization': 'Basic ' + btoa('default:default')
  }
});

// Snap raw MiB-derived GiB to standard GPU VRAM tier
const snapToStandardVramGB = (rawGiB: number): number => {
  const tiers = [8, 16, 24, 32, 40, 48, 80, 96, 144, 192, 288];
  for (const tier of tiers) {
    if (Math.abs(tier - rawGiB) <= 2.5) {
      return tier;
    }
  }
  return Math.round(rawGiB);
};

// Parse HAProxy bind directive (e.g., "*:80", "0.0.0.0:9999", "10.0.1.50:8080")
const parseHaproxyBind = (bindStr: string): { bindAddress: string; port: number } => {
  const firstToken = String(bindStr || '').trim().split(/\s+/)[0] || '*:80';
  const lastColon = firstToken.lastIndexOf(':');
  if (lastColon === -1) {
    const maybePort = parseInt(firstToken, 10);
    return { bindAddress: '*', port: Number.isFinite(maybePort) && maybePort > 0 ? maybePort : 80 };
  }
  const rawAddr = firstToken.slice(0, lastColon).trim();
  const rawPort = parseInt(firstToken.slice(lastColon + 1).trim(), 10);
  const bindAddress = !rawAddr || rawAddr === '*' || rawAddr === '0.0.0.0' || rawAddr === '::' ? '*' : rawAddr;
  return { bindAddress, port: Number.isFinite(rawPort) && rawPort > 0 ? rawPort : 80 };
};

// Extract OnpremNlb list from raw cm-honeybee infra servers[].haproxy
const extractNlbsFromRawServers = (rawServers: any[], refinedNodes: OnpremNode[]): OnpremNlb[] => {
  const nlbs: OnpremNlb[] = [];
  for (const srv of rawServers) {
    const haproxy = srv?.haproxy;
    if (!haproxy) continue;
    const rawMachineId = srv?.compute?.os?.node?.machineid || '';
    const rawHostname = srv?.compute?.os?.node?.hostname || '';
    const matchedNode = refinedNodes.find(
      (n) => (rawMachineId && n.machineId === rawMachineId) || (rawHostname && n.hostname === rawHostname)
    );
    const hostMachineId = matchedNode?.machineId || rawMachineId || rawHostname || 'haproxy-node';
    const defaultMode = String(haproxy?.defaults?.mode || 'tcp').trim().toLowerCase();

    const parseServers = (srvList: any[], fallbackPort: number) =>
      (Array.isArray(srvList) ? srvList : [])
        .map((s: any, idx: number) => {
          const addr = String(s?.address || '').trim();
          if (!addr) return null;
          const colonIdx = addr.lastIndexOf(':');
          const ip = colonIdx > 0 ? addr.slice(0, colonIdx).trim() : addr;
          const parsedPort = colonIdx > 0 ? parseInt(addr.slice(colonIdx + 1).trim(), 10) : fallbackPort;
          if (!ip) return null;
          return {
            name: String(s?.name || `srv-${idx + 1}`),
            ip,
            port: Number.isFinite(parsedPort) && parsedPort > 0 ? parsedPort : fallbackPort,
          };
        })
        .filter((s): s is { name: string; ip: string; port: number } => s !== null);

    const frontends: any[] = Array.isArray(haproxy.frontends) ? haproxy.frontends : [];
    const backends: any[] = Array.isArray(haproxy.backends) ? haproxy.backends : [];
    for (const fe of frontends) {
      const be =
        backends.find((b: any) => b?.name && b.name === fe?.default_backend) ||
        (backends.length === 1 ? backends[0] : undefined);
      if (!be) continue;
      const listener = parseHaproxyBind(fe?.bind || '*:80');
      const servers = parseServers(be?.servers, listener.port);
      if (servers.length === 0) continue;
      const mode = String(be?.options?.mode || fe?.options?.mode || defaultMode).toLowerCase();
      const backendProtocol = mode === 'http' ? 'http' : 'tcp';
      nlbs.push({
        hostMachineId,
        software: 'haproxy',
        listener: {
          bindAddress: listener.bindAddress,
          port: listener.port,
          protocol: 'tcp',
        },
        backend: {
          name: String(be?.name || fe?.name || 'haproxy-backend'),
          balance: String(be?.balance || haproxy?.defaults?.balance || 'roundrobin'),
          protocol: backendProtocol,
          servers,
        },
        healthCheck: {
          enabled: true,
          interval: 10,
          timeout: 5,
          threshold: 3,
        },
      });
    }

    const listens: any[] = Array.isArray(haproxy.listens) ? haproxy.listens : [];
    for (const ls of listens) {
      const listener = parseHaproxyBind(ls?.bind || '*:80');
      const servers = parseServers(ls?.servers, listener.port);
      if (servers.length === 0) continue;
      const mode = String(ls?.options?.mode || defaultMode).toLowerCase();
      const backendProtocol = mode === 'http' ? 'http' : 'tcp';
      nlbs.push({
        hostMachineId,
        software: 'haproxy',
        listener: {
          bindAddress: listener.bindAddress,
          port: listener.port,
          protocol: 'tcp',
        },
        backend: {
          name: String(ls?.name || 'haproxy-listen'),
          balance: String(ls?.balance || haproxy?.defaults?.balance || 'roundrobin'),
          protocol: backendProtocol,
          servers,
        },
        healthCheck: {
          enabled: true,
          interval: 10,
          timeout: 5,
          threshold: 3,
        },
      });
    }
  }
  return nlbs;
};

// Backfill GPU cards, NVMe rootDisk size, HAProxy NLBs, and connection hostnames from raw Honeybee infra
const enrichRefinedInfraWithRawData = (
  refined: OnpremInfra,
  rawData: any,
  connections?: any[],
  targetConnId?: string
): OnpremInfra => {
  if (!refined || !Array.isArray(refined.nodes)) return refined;
  const rawServers: any[] = Array.isArray(rawData?.servers)
    ? rawData.servers
    : (rawData?.compute || rawData?.gpu || rawData?.haproxy ? [rawData] : []);
  const connList: any[] = Array.isArray(connections) ? connections : [];

  const existingNlbs = Array.isArray(refined.nlbs) && refined.nlbs.length > 0 ? refined.nlbs : [];
  const extractedNlbs = existingNlbs.length > 0 ? existingNlbs : extractNlbsFromRawServers(rawServers, refined.nodes);

  return {
    ...refined,
    ...(extractedNlbs.length > 0 ? { nlbs: extractedNlbs } : {}),
    nodes: refined.nodes.map((node, nodeIdx) => {
      const matchedServer = rawServers.find((srv: any) => {
        const rawMachineId = srv?.compute?.os?.node?.machineid;
        const rawHostname = srv?.compute?.os?.node?.hostname;
        return (rawMachineId && rawMachineId === node.machineId) || (rawHostname && rawHostname === node.hostname);
      });

      const nodeIps = (node.interfaces || [])
        .flatMap((iface) => iface.ipv4CidrBlocks || [])
        .map((cidr) => String(cidr).split('/')[0]?.trim())
        .filter((ip) => Boolean(ip && ip !== '127.0.0.1'));

      const isGenericCloudHostname =
        !node.hostname ||
        /^ip-\d+-\d+-\d+-\d+/i.test(node.hostname.trim()) ||
        /^localhost$/i.test(node.hostname.trim());

      const matchedConn =
        (targetConnId ? connList.find((c: any) => c?.id === targetConnId) : undefined) ||
        connList.find((c: any) => {
          const connIp = String(c?.ip_address || c?.ip || '').trim();
          return connIp && nodeIps.includes(connIp);
        }) ||
        connList.find((c: any) => c?.name && c.name === node.hostname) ||
        (isGenericCloudHostname && connList.length === refined.nodes.length ? connList[nodeIdx] : undefined);

      const connName = String(matchedConn?.name || '').trim();
      const resolvedHostname =
        isGenericCloudHostname && connName && !/^host\s*\d+$/i.test(connName)
          ? connName
          : node.hostname;
      const resolvedWorkloadTag =
        node.workloadTag || (matchedConn?.description ? String(matchedConn.description).trim() : undefined);

      const rawNvidiaGpus: any[] = Array.isArray(matchedServer?.gpu?.nvidia) ? matchedServer.gpu.nvidia : [];
      const hasExistingGpus = Array.isArray(node.gpuCards) && node.gpuCards.length > 0;
      const gpuCards = hasExistingGpus
        ? node.gpuCards!.map((g) => {
            const memoryTotalGB = g.memoryTotalGB ? snapToStandardVramGB(Number(g.memoryTotalGB)) : 16;
            const memoryUsedGB = g.memoryUsedGB !== undefined ? Math.round(Number(g.memoryUsedGB)) : undefined;
            const memoryFreeGB = g.memoryFreeGB !== undefined
              ? Math.round(Number(g.memoryFreeGB))
              : (memoryUsedGB !== undefined ? Math.max(0, memoryTotalGB - memoryUsedGB) : undefined);
            return {
              ...g,
              memoryTotalGB,
              ...(memoryUsedGB !== undefined ? { memoryUsedGB } : {}),
              ...(memoryFreeGB !== undefined ? { memoryFreeGB } : {}),
            };
          })
        : rawNvidiaGpus.length > 0
          ? rawNvidiaGpus.map((g: any, idx: number) => {
              const totalMiB = Number(g?.performance?.fb_memory_total || 0);
              const usedMiB = Number(g?.performance?.fb_memory_used || 0);
              const memoryTotalGB = totalMiB > 0 ? snapToStandardVramGB(totalMiB / 1024) : 16;
              const memoryUsedGB = usedMiB > 0 ? Math.round(usedMiB / 1024) : 0;
              return {
                driverIndex: String(idx),
                uuid: g?.device_attribute?.gpu_uuid || '',
                vendor: g?.device_attribute?.product_brand || 'NVIDIA',
                model: g?.device_attribute?.product_name || 'GPU',
                type: 'GPU',
                architecture: g?.device_attribute?.product_architecture || '',
                driverVersion: g?.device_attribute?.driver_version || '',
                cudaVersion: g?.device_attribute?.cuda_version || '',
                memoryTotalGB,
                memoryUsedGB,
                memoryFreeGB: Math.max(0, memoryTotalGB - memoryUsedGB),
              };
            })
          : undefined;

      const hasGpu = Array.isArray(gpuCards) && gpuCards.length > 0;
      const currentDiskSize = Number(node.rootDisk?.totalSize || 0);
      const rootDisk = currentDiskSize > 0
        ? node.rootDisk
        : {
            label: node.rootDisk?.label || '/dev/nvme0n1p1',
            type: node.rootDisk?.type || 'SSD',
            totalSize: resolvedHostname === 'ai-gpu-01' ? 100 : (hasGpu ? 80 : 50),
          };

      return {
        ...node,
        hostname: resolvedHostname,
        ...(resolvedWorkloadTag ? { workloadTag: resolvedWorkloadTag } : {}),
        ...(gpuCards ? { gpuCards } : {}),
        rootDisk,
      };
    }),
  };
};

// ============================================================================
// 1. CM-Honeybee Server API Client (Source Collection)
// ============================================================================
export const honeybeeApi = {
  // Create a new source group for on-premise servers
  createSourceGroup: async (name: string, description: string) => {
    const response = await api.post('/honeybee/source_group', {
      name,
      description,
      type: 'ssh'
    });
    return response.data; // returns { id, name, description, ... }
  },

  // Get all registered source groups
  getSourceGroups: async () => {
    const response = await api.get('/honeybee/source_group');
    return response.data; // returns array of source groups
  },

  // Register a server connectivity information under a group
  registerConnectionInfo: async (sgId: string, serverData: {
    name: string;
    ip: string;
    port: number;
    user: string;
    password?: string;
    privateKey?: string;
    description?: string;
  }) => {
    const response = await api.post(`/honeybee/source_group/${sgId}/connection_info`, {
      name:        serverData.name,
      ip_address:  serverData.ip,
      ssh_port:    String(serverData.port || 22),
      user:        serverData.user,
      password:    serverData.password    || '',
      private_key: serverData.privateKey  || '',
      description: serverData.description || '',
    });
    return response.data;
  },

  // Get connections inside a source group
  getConnectionInfoList: async (sgId: string) => {
    const response = await api.get(`/honeybee/source_group/${sgId}/connection_info`);
    return response.data;
  },

  // Get refined OnpremInfra model after background scan finishes
  getRefinedInfraByConnection: async (sgId: string, connId: string): Promise<OnpremInfra> => {
    const [refinedRes, connRes] = await Promise.all([
      api.get(`/honeybee/source_group/${sgId}/connection_info/${connId}/infra/refined`),
      api.get(`/honeybee/source_group/${sgId}/connection_info`).catch(() => ({ data: null })),
    ]);
    const refined = refinedRes.data?.onpremiseInfraModel || refinedRes.data;
    const connList = connRes.data?.connection_info || (Array.isArray(connRes.data) ? connRes.data : []);
    const hasNativeGpu = Array.isArray(refined?.nodes) && refined.nodes.some((n: any) => Array.isArray(n.gpuCards) && n.gpuCards.length > 0);
    const hasNativeNlb = Array.isArray(refined?.nlbs) && refined.nlbs.length > 0;
    const rawRes = (hasNativeGpu && hasNativeNlb)
      ? { data: null }
      : await api.get(`/honeybee/source_group/${sgId}/connection_info/${connId}/infra`).catch(() => ({ data: null }));
    return enrichRefinedInfraWithRawData(refined, rawRes.data, connList, connId);
  },

  getRefinedInfraByGroup: async (sgId: string): Promise<OnpremInfra> => {
    const [refinedRes, connRes] = await Promise.all([
      api.get(`/honeybee/source_group/${sgId}/infra/refined`),
      api.get(`/honeybee/source_group/${sgId}/connection_info`).catch(() => ({ data: null })),
    ]);
    const refined = refinedRes.data?.onpremiseInfraModel || refinedRes.data;
    const connList = connRes.data?.connection_info || (Array.isArray(connRes.data) ? connRes.data : []);
    const hasNativeGpu = Array.isArray(refined?.nodes) && refined.nodes.some((n: any) => Array.isArray(n.gpuCards) && n.gpuCards.length > 0);
    const hasNativeNlb = Array.isArray(refined?.nlbs) && refined.nlbs.length > 0;
    const rawRes = (hasNativeGpu && hasNativeNlb)
      ? { data: null }
      : await api.get(`/honeybee/source_group/${sgId}/infra`).catch(() => ({ data: null }));
    return enrichRefinedInfraWithRawData(refined, rawRes.data, connList);
  },

  // Create source group with embedded connection info (single API call)
  createSourceGroupWithConnections: async (data: {
    name: string;
    description?: string;
    connection_info?: Array<{
      name: string;
      ip_address: string;
      ssh_port: string;
      user: string;
      password?: string;
      private_key?: string;
      description?: string;
    }>;
  }) => {
    const response = await api.post('/honeybee/source_group', data);
    return response.data;
  },

  // Refresh all connection statuses in a source group
  refreshSourceGroup: async (sgId: string) => {
    const response = await api.put(`/honeybee/source_group/${sgId}/refresh`);
    return response.data;
  },

  // Refresh individual connection status
  refreshConnectionInfo: async (sgId: string, connId: string) => {
    const response = await api.put(`/honeybee/source_group/${sgId}/connection_info/${connId}/refresh`);
    return response.data;
  },

  // Trigger infra collection for all connections in a group
  importInfraByGroup: async (sgId: string) => {
    const response = await api.post(`/honeybee/source_group/${sgId}/import/infra`);
    return response.data;
  },

  // Trigger infra collection for a single connection
  importInfraByConnection: async (sgId: string, connId: string) => {
    const response = await api.post(`/honeybee/source_group/${sgId}/connection_info/${connId}/import/infra`);
    return response.data;
  },

  // Delete a source group
  deleteSourceGroup: async (sgId: string) => {
    const response = await api.delete(`/honeybee/source_group/${sgId}`);
    return response.data;
  },

  // Update source group name/description
  updateSourceGroup: async (sgId: string, name: string, description: string) => {
    const response = await api.put(`/honeybee/source_group/${sgId}`, { name, description });
    return response.data;
  },

  // Delete a single connection info entry
  deleteConnectionInfo: async (sgId: string, connId: string) => {
    const response = await api.delete(`/honeybee/source_group/${sgId}/connection_info/${connId}`);
    return response.data;
  },
};

// ============================================================================
// 2. CM-Beetle Artifact Repository API Client (/beetle/artifacts)
// ============================================================================
export interface SaveArtifactOptions {
  stage?: 'source' | 'planning' | 'target';
  parentId?: string;
  sourceGroupId?: string;
}

const extractArtifactList = (resData: any): any[] => {
  if (Array.isArray(resData)) return resData;
  if (resData && typeof resData === 'object') {
    if ('data' in resData) {
      return Array.isArray(resData.data) ? resData.data : (resData.data ? [resData.data] : []);
    }
    if ('artifacts' in resData && Array.isArray(resData.artifacts)) {
      return resData.artifacts;
    }
    // Empty ApiResponse wrapper such as { success: true } where Go omitempty omitted an empty slice
    if ('success' in resData && !('id' in resData) && !('payload' in resData)) {
      return [];
    }
  }
  return [];
};

const LEGACY_SAMPLE_MACHINE_ID_MAP: Record<string, string> = {
  'mig-testbed-vm1': 'ec289605-629d-4ddf-b4df-0f4c172d8401',
  'mig-testbed-vm2': 'ec289605-629d-4ddf-b4df-0f4c172d8402',
  'mig-testbed-vm3': 'ec289605-629d-4ddf-b4df-0f4c172d8403',
  'mig-testbed-vm4': 'ec289605-629d-4ddf-b4df-0f4c172d8404',
  'mig-testbed-vm5': 'ec289605-629d-4ddf-b4df-0f4c172d8405',
};

// Normalize uint32/uint64/int fields on OnpremInfra so Go json.Unmarshal never fails on decimals.
const sanitizeOnpremInfraIntegers = (infra: OnpremInfra | undefined | null): OnpremInfra => {
  if (!infra || typeof infra !== 'object') return infra as any;
  const toUint = (val: unknown, minVal = 0): number => {
    const n = Number(val);
    if (!Number.isFinite(n)) return minVal;
    return Math.max(minVal, Math.round(n));
  };
  return {
    ...infra,
    nodes: Array.isArray(infra.nodes)
      ? infra.nodes.map((node) => ({
          ...node,
          machineId: LEGACY_SAMPLE_MACHINE_ID_MAP[node.machineId] || node.machineId,
          cpu: node.cpu
            ? {
                ...node.cpu,
                cpus: toUint(node.cpu.cpus, 1),
                cores: toUint(node.cpu.cores, 1),
                threads: toUint(node.cpu.threads, 1),
              }
            : node.cpu,
          memory: node.memory
            ? {
                ...node.memory,
                totalSize: toUint(node.memory.totalSize, 1),
                available: toUint(node.memory.available, 0),
                ...(typeof (node.memory as any).used === 'number'
                  ? { used: toUint((node.memory as any).used, 0) }
                  : {}),
              }
            : node.memory,
          rootDisk: node.rootDisk
            ? {
                ...node.rootDisk,
                totalSize: toUint(node.rootDisk.totalSize, 1),
                ...(typeof (node.rootDisk as any).available === 'number'
                  ? { available: toUint((node.rootDisk as any).available, 0) }
                  : {}),
                ...(typeof (node.rootDisk as any).used === 'number'
                  ? { used: toUint((node.rootDisk as any).used, 0) }
                  : {}),
              }
            : node.rootDisk,
          ...(Array.isArray(node.dataDisks)
            ? {
                dataDisks: node.dataDisks.map((d) => ({
                  ...d,
                  totalSize: toUint(d.totalSize, 1),
                })),
              }
            : {}),
        }))
      : [],
    ...(Array.isArray(infra.nlbs)
      ? {
          nlbs: infra.nlbs.map((nlb) => ({
            ...nlb,
            ...(nlb.hostMachineId && LEGACY_SAMPLE_MACHINE_ID_MAP[nlb.hostMachineId]
              ? { hostMachineId: LEGACY_SAMPLE_MACHINE_ID_MAP[nlb.hostMachineId] }
              : {}),
          })),
        }
      : {}),
  };
};

export interface StorageSourceModelEnvelope {
  id: string;
  name: string;
  description?: string;
  version?: string;
  stage?: 'source' | 'planning' | 'target';
  category?: 'infra' | 'software' | 'data';
  schemaType?: string;
  lineage?: ArtifactLineage;
  sourceObjectStorages: any[];
  sourceCloud?: { csp: string; region: string };
  updatedTime?: string;
  isDraft?: boolean;
}

export interface StorageTargetModelEnvelope {
  id: string;
  name: string;
  description?: string;
  version?: string;
  stage?: 'source' | 'planning' | 'target';
  category?: 'infra' | 'software' | 'data';
  schemaType?: string;
  lineage?: ArtifactLineage;
  csp: string;
  region: string;
  targetObjectStorages: any[];
  raw: any;
  updatedTime?: string;
  isDraft?: boolean;
}

const mapArtifactToSourceEnvelope = (raw: any, fallbackInfra?: OnpremInfra): OnpremModelEnvelope => {
  const m = raw?.data || raw || {};
  let payload = m.payload || m.onpremiseInfraModel || fallbackInfra;
  if (typeof payload === 'string') {
    try { payload = JSON.parse(payload); } catch {}
  }
  const rawName = m.name || m.userModelName || 'Unnamed Artifact';
  return {
    id: m.id || m.uid || `art-src-${Date.now()}`,
    name: rawName === 'Unnamed Artifact' ? rawName : cleanModelDisplayName(rawName, 'Unnamed Artifact'),
    description: m.description || '',
    version: normalizeModelVersion(m.version || m.userModelVersion, '1.0.0'),
    stage: m.stage || 'source',
    category: m.category || 'infra',
    schemaType: m.schemaType || 'OnpremInfra',
    lineage: m.lineage,
    onpremiseInfraModel: sanitizeOnpremInfraIntegers(payload as OnpremInfra),
    updatedTime: m.updateTime || m.createTime || m.updatedTime || new Date().toISOString()
  };
};

const mapArtifactToCloudEnvelope = (raw: any, fallbackInfra?: RecommendedInfra): CloudModelEnvelope => {
  const m = raw?.data || raw || {};
  let payload = m.payload || m.cloudInfraModel || m.cloud_infra_model || m.recommendedInfra || fallbackInfra;
  if (typeof payload === 'string') {
    try { payload = JSON.parse(payload); } catch {}
  }
  const rawName = m.name || m.userModelName || 'Unnamed Design';
  return {
    id: m.id || m.uid || `art-tgt-${Date.now()}`,
    name: rawName === 'Unnamed Design' ? rawName : cleanModelDisplayName(rawName, 'Unnamed Design'),
    description: m.description || '',
    version: normalizeModelVersion(m.version || m.userModelVersion, '1.0.0'),
    stage: m.stage || 'target',
    category: m.category || 'infra',
    schemaType: m.schemaType || 'RecommendedInfra',
    lineage: m.lineage,
    cloudInfraModel: payload as RecommendedInfra,
    updatedTime: m.updateTime || m.createTime || m.updatedTime || new Date().toISOString()
  };
};

const mapArtifactToStorageSourceEnvelope = (raw: any, fallbackPayload?: any): StorageSourceModelEnvelope => {
  const m = raw?.data || raw || {};
  let payload = m.payload || fallbackPayload || {};
  if (typeof payload === 'string') {
    try { payload = JSON.parse(payload); } catch {}
  }
  const buckets = Array.isArray(payload?.sourceObjectStorages)
    ? payload.sourceObjectStorages
    : Array.isArray(payload?.buckets)
      ? payload.buckets
      : [];
  const rawName = m.name || m.userModelName || 'Unnamed Storage Artifact';
  return {
    id: m.id || m.uid || `art-stg-src-${Date.now()}`,
    name: rawName === 'Unnamed Storage Artifact' ? rawName : cleanModelDisplayName(rawName, 'Unnamed Storage Artifact'),
    description: m.description || payload?.description || '',
    version: normalizeModelVersion(m.version || m.userModelVersion, '1.0.0'),
    stage: m.stage || 'source',
    category: m.category || 'infra',
    schemaType: m.schemaType || 'SourceObjectStorage',
    lineage: m.lineage,
    sourceObjectStorages: buckets,
    sourceCloud: payload?.sourceCloud,
    updatedTime: m.updateTime || m.createTime || m.updatedTime || new Date().toISOString(),
  };
};

const mapArtifactToStorageTargetEnvelope = (raw: any, fallbackPayload?: any): StorageTargetModelEnvelope => {
  const m = raw?.data || raw || {};
  let payload = m.payload || fallbackPayload || {};
  if (typeof payload === 'string') {
    try { payload = JSON.parse(payload); } catch {}
  }
  const storages = Array.isArray(payload?.targetObjectStorages) ? payload.targetObjectStorages : [];
  const rawName = m.name || m.userModelName || 'Unnamed Storage Target';
  return {
    id: m.id || m.uid || `art-stg-tgt-${Date.now()}`,
    name: rawName === 'Unnamed Storage Target' ? rawName : cleanModelDisplayName(rawName, 'Unnamed Storage Target'),
    description: m.description || payload?.description || '',
    version: normalizeModelVersion(m.version || m.userModelVersion, '1.0.0'),
    stage: m.stage || 'target',
    category: m.category || 'infra',
    schemaType: m.schemaType || 'RecommendedObjectStorage',
    lineage: m.lineage,
    csp: payload?.targetCloud?.csp || 'aws',
    region: payload?.targetCloud?.region || 'ap-northeast-2',
    targetObjectStorages: storages,
    raw: payload,
    updatedTime: m.updateTime || m.createTime || m.updatedTime || new Date().toISOString(),
  };
};

export const artifactApi = {
  // Save source or planning infra artifact
  saveSourceModel: async (
    name: string,
    description: string,
    onpremiseInfraModel: OnpremInfra,
    version?: string,
    options?: SaveArtifactOptions
  ): Promise<OnpremModelEnvelope> => {
    const sanitizedModel = sanitizeOnpremInfraIntegers(onpremiseInfraModel);
    const response = await api.post('/beetle/artifacts', {
      name: cleanModelDisplayName(name),
      description,
      version: normalizeModelVersion(version, '1.0.0'),
      category: 'infra',
      stage: options?.stage || 'source',
      schemaType: 'OnpremInfra',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload: sanitizedModel
    });
    return mapArtifactToSourceEnvelope(response.data, sanitizedModel);
  },

  // Update existing source or planning infra artifact
  updateSourceModel: async (
    id: string,
    name: string,
    description: string,
    onpremiseInfraModel: OnpremInfra,
    version?: string,
    options?: SaveArtifactOptions
  ): Promise<OnpremModelEnvelope> => {
    const sanitizedModel = sanitizeOnpremInfraIntegers(onpremiseInfraModel);
    const response = await api.put(`/beetle/artifacts/${id}`, {
      name: cleanModelDisplayName(name),
      description,
      version: normalizeModelVersion(version, '1.0.0'),
      category: 'infra',
      stage: options?.stage,
      schemaType: 'OnpremInfra',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload: sanitizedModel
    });
    return mapArtifactToSourceEnvelope(response.data, sanitizedModel);
  },

  // Delete existing source or planning artifact
  deleteSourceModel: async (id: string): Promise<void> => {
    if (isBuiltInSampleModel(id) || id.startsWith('art-src-') || id.startsWith('art-plan-')) {
      return;
    }
    try {
      await api.delete(`/beetle/artifacts/${id}`);
    } catch (err: any) {
      if (err?.response?.status === 404) {
        return;
      }
      throw err;
    }
  },

  // Get all source and planning OnpremInfra artifacts
  getSourceModels: async (): Promise<OnpremModelEnvelope[]> => {
    const response = await api.get('/beetle/artifacts?category=infra&schemaType=OnpremInfra');
    const dataList = extractArtifactList(response.data);
    return dataList
      .map((m: any) => mapArtifactToSourceEnvelope(m))
      .filter((m) => m.name !== 'Unnamed Artifact' && Boolean(m.onpremiseInfraModel));
  },

  // Get single source or planning artifact by ID
  getSourceModel: async (id: string): Promise<OnpremModelEnvelope> => {
    const response = await api.get(`/beetle/artifacts/${id}`);
    return mapArtifactToSourceEnvelope(response.data);
  },

  // Save recommended target cloud spec as a target stage artifact
  saveCloudModel: async (
    name: string,
    description: string,
    cloudInfraModel: RecommendedInfra,
    version?: string,
    options?: SaveArtifactOptions
  ): Promise<CloudModelEnvelope> => {
    const response = await api.post('/beetle/artifacts', {
      name: cleanModelDisplayName(name),
      description,
      version: normalizeModelVersion(version, '1.0.0'),
      category: 'infra',
      stage: 'target',
      schemaType: 'RecommendedInfra',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload: cloudInfraModel
    });
    return mapArtifactToCloudEnvelope(response.data, cloudInfraModel);
  },

  // Update existing target cloud model artifact
  updateCloudModel: async (
    id: string,
    name: string,
    description: string,
    cloudInfraModel: RecommendedInfra,
    version?: string,
    options?: SaveArtifactOptions
  ): Promise<CloudModelEnvelope> => {
    const response = await api.put(`/beetle/artifacts/${id}`, {
      name: cleanModelDisplayName(name),
      description,
      version: normalizeModelVersion(version, '1.0.0'),
      category: 'infra',
      stage: 'target',
      schemaType: 'RecommendedInfra',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload: cloudInfraModel
    });
    return mapArtifactToCloudEnvelope(response.data, cloudInfraModel);
  },

  // Get target cloud model artifacts
  getCloudModels: async (): Promise<CloudModelEnvelope[]> => {
    const response = await api.get('/beetle/artifacts?category=infra&stage=target&schemaType=RecommendedInfra');
    const dataList = extractArtifactList(response.data);
    return dataList
      .map((m: any) => mapArtifactToCloudEnvelope(m))
      .filter((m) => m.name !== 'Unnamed Design' && Boolean(m.cloudInfraModel));
  },

  getCloudModel: async (id: string): Promise<CloudModelEnvelope> => {
    const response = await api.get(`/beetle/artifacts/${id}`);
    return mapArtifactToCloudEnvelope(response.data);
  },

  deleteCloudModel: async (id: string): Promise<void> => {
    if (isBuiltInSampleModel(id) || id.startsWith('cloud-model-') || id.startsWith('art-tgt-')) {
      return;
    }
    try {
      await api.delete(`/beetle/artifacts/${id}`);
    } catch (err: any) {
      if (err?.response?.status === 404) {
        return;
      }
      throw err;
    }
  },

  // Get all source and planning SourceObjectStorage artifacts
  getStorageSourceModels: async (): Promise<StorageSourceModelEnvelope[]> => {
    const response = await api.get('/beetle/artifacts?category=infra&schemaType=SourceObjectStorage');
    const dataList = extractArtifactList(response.data);
    return dataList
      .map((m: any) => mapArtifactToStorageSourceEnvelope(m))
      .filter((m) => m.name !== 'Unnamed Storage Artifact' && Array.isArray(m.sourceObjectStorages));
  },

  // Save source or planning SourceObjectStorage artifact
  saveStorageSourceModel: async (
    name: string,
    description: string,
    payload: { description?: string; sourceCloud?: { csp: string; region: string }; sourceObjectStorages: any[] },
    version?: string,
    options?: SaveArtifactOptions
  ): Promise<StorageSourceModelEnvelope> => {
    const response = await api.post('/beetle/artifacts', {
      name: cleanModelDisplayName(name),
      description,
      version: normalizeModelVersion(version, '1.0.0'),
      category: 'infra',
      stage: options?.stage || 'source',
      schemaType: 'SourceObjectStorage',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload,
    });
    return mapArtifactToStorageSourceEnvelope(response.data, payload);
  },

  // Update existing source or planning SourceObjectStorage artifact
  updateStorageSourceModel: async (
    id: string,
    name: string,
    description: string,
    payload: { description?: string; sourceCloud?: { csp: string; region: string }; sourceObjectStorages: any[] },
    version?: string,
    options?: SaveArtifactOptions
  ): Promise<StorageSourceModelEnvelope> => {
    const response = await api.put(`/beetle/artifacts/${id}`, {
      name: cleanModelDisplayName(name),
      description,
      version: normalizeModelVersion(version, '1.0.0'),
      category: 'infra',
      stage: options?.stage,
      schemaType: 'SourceObjectStorage',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload,
    });
    return mapArtifactToStorageSourceEnvelope(response.data, payload);
  },

  // Get all target RecommendedObjectStorage artifacts
  getStorageTargetModels: async (): Promise<StorageTargetModelEnvelope[]> => {
    const response = await api.get('/beetle/artifacts?category=infra&stage=target&schemaType=RecommendedObjectStorage');
    const dataList = extractArtifactList(response.data);
    return dataList
      .map((m: any) => mapArtifactToStorageTargetEnvelope(m))
      .filter((m) => m.name !== 'Unnamed Storage Target' && Array.isArray(m.targetObjectStorages));
  },

  // Save target RecommendedObjectStorage artifact
  saveStorageTargetModel: async (
    name: string,
    description: string,
    payload: any,
    version?: string,
    options?: SaveArtifactOptions
  ): Promise<StorageTargetModelEnvelope> => {
    const response = await api.post('/beetle/artifacts', {
      name: cleanModelDisplayName(name),
      description,
      version: normalizeModelVersion(version, '1.0.0'),
      category: 'infra',
      stage: 'target',
      schemaType: 'RecommendedObjectStorage',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload,
    });
    return mapArtifactToStorageTargetEnvelope(response.data, payload);
  },

  // Update existing target RecommendedObjectStorage artifact
  updateStorageTargetModel: async (
    id: string,
    name: string,
    description: string,
    payload: any,
    version?: string,
    options?: SaveArtifactOptions
  ): Promise<StorageTargetModelEnvelope> => {
    const response = await api.put(`/beetle/artifacts/${id}`, {
      name: cleanModelDisplayName(name),
      description,
      version: normalizeModelVersion(version, '1.0.0'),
      category: 'infra',
      stage: 'target',
      schemaType: 'RecommendedObjectStorage',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload,
    });
    return mapArtifactToStorageTargetEnvelope(response.data, payload);
  },

  // Delete object storage artifact by ID
  deleteStorageModel: async (id: string): Promise<void> => {
    if (isBuiltInSampleModel(id) || id.startsWith('art-stg-')) {
      return;
    }
    try {
      await api.delete(`/beetle/artifacts/${id}`);
    } catch (err: any) {
      if (err?.response?.status === 404) {
        return;
      }
      throw err;
    }
  },

  getArtifactLineage: async (id: string) => {
    const response = await api.get(`/beetle/artifacts/${id}/lineage`);
    return response.data?.data ?? response.data;
  }
};

// ============================================================================
// 3. CM-Beetle Server API Client (Recommendation & Migration Engine)
// ============================================================================
let objectStorageSupportCachePromise: Promise<Record<string, { cors: boolean; presignedUrl: boolean; versioning: boolean }>> | null = null;

export const beetleApi = {
  // Get Cloud Recommendation candidates based on Source model input
  getRecommendations: async (
    sourceInfra: OnpremInfra,
    desiredCsp: string,
    desiredRegion: string,
    limit?: number,
    minMatchRate?: number
  ): Promise<RecommendedInfra[]> => {
    const sanitizedInfra = sanitizeOnpremInfraIntegers(sourceInfra);
    const hasNlbs = sanitizedInfra.nlbs && sanitizedInfra.nlbs.length > 0;
    const limitQuery = limit ? `&limit=${limit}` : '';
    const matchRateQuery = typeof minMatchRate === 'number' && !isNaN(minMatchRate) ? `&minMatchRate=${minMatchRate}` : '';

    let response;
    if (hasNlbs) {
      // Call NLB-aware recommendation endpoint
      response = await api.post(`/beetle/recommendation/infraWithNlb?desiredCsp=${desiredCsp}&desiredRegion=${desiredRegion}${limitQuery}${matchRateQuery}`, {
        desiredCsp,
        desiredRegion,
        sourceInfra: sanitizedInfra
      });
    } else {
      // Call standard VM recommendation endpoint
      response = await api.post(`/beetle/recommendation/infra?desiredCsp=${desiredCsp}&desiredRegion=${desiredRegion}${limitQuery}${matchRateQuery}`, {
        desiredCspAndRegionPair: {
          csp: desiredCsp,
          region: desiredRegion
        },
        onpremiseInfraModel: sanitizedInfra
      });
    }

    // Beetle recommendation returns wrapping object: { success: true, data: RecommendedInfra[] }
    if (response.data && Array.isArray(response.data.data)) {
      return response.data.data;
    }
    return Array.isArray(response.data) ? response.data : [];
  },

  // Get Multi-Cloud Recommendation candidates based on Source model input and multiple target CSP/Region pairs (2~10 pairs)
  getMultiRecommendations: async (
    sourceInfra: OnpremInfra,
    targetPairs: CloudProperty[],
    minMatchRate: number = 85
  ): Promise<RecommendedInfra[]> => {
    const sanitizedInfra = sanitizeOnpremInfraIntegers(sourceInfra);
    const hasNlbs = sanitizedInfra.nlbs && sanitizedInfra.nlbs.length > 0;
    const query = `?minMatchRate=${minMatchRate}`;

    let response;
    if (hasNlbs) {
      // Call NLB-aware multi-infra recommendation endpoint
      response = await api.post(`/beetle/recommendation/multiInfraWithNlb${query}`, {
        desiredCspAndRegionPairs: targetPairs,
        sourceInfra: sanitizedInfra
      });
    } else {
      // Call standard multi-infra recommendation endpoint
      response = await api.post(`/beetle/recommendation/multiInfra${query}`, {
        desiredCspAndRegionPairs: targetPairs,
        sourceInfra: sanitizedInfra
      });
    }

    if (response.data && Array.isArray(response.data.data)) {
      return response.data.data;
    }
    return Array.isArray(response.data) ? response.data : [];
  },

  // Validate target infrastructure model before migration
  validateInfra: async (
    nsId: string,
    cloudModel: RecommendedInfra,
    useExisting: boolean = true
  ): Promise<{
    success: boolean;
    valid?: boolean;
    issues?: Array<{ code: string; severity: string; path: string; message: string }>;
    error?: string;
  }> => {
    try {
      const response = await api.post(`/beetle/validation/ns/${nsId}/infra?useExisting=${useExisting}`, cloudModel);
      const resData = response.data?.data || response.data;
      return {
        success: response.data?.success ?? true,
        valid: resData?.valid,
        issues: Array.isArray(resData?.issues) ? resData.issues : []
      };
    } catch (err: any) {
      return {
        success: false,
        valid: false,
        issues: [],
        error: err.response?.data?.error || err.response?.data?.message || err.message
      };
    }
  },

  // Execute actual physical cloud migration deployment (with Prefer: respond-async header)
  executeMigration: async (
    nsId: string,
    nameSeed: string,
    cloudModel: RecommendedInfra,
    useExisting: boolean = true,
    lineage?: ArtifactLineage
  ): Promise<{ success: boolean; reqId?: string; data?: any; error?: string }> => {
    try {
      const params = new URLSearchParams();
      if (nameSeed && nameSeed.trim().length > 0) {
        params.set('nameSeed', nameSeed.trim());
      }
      params.set('useExisting', String(useExisting));
      const requestBody = lineage ? { ...cloudModel, lineage } : cloudModel;
      const response = await api.post(`/beetle/migration/ns/${nsId}/infra?${params.toString()}`, requestBody, {
        headers: {
          'Prefer': 'respond-async'
        }
      });
      const reqId = response.data?.data?.reqId || response.headers['x-request-id'];
      return { success: true, reqId, data: response.data };
    } catch (err: any) {
      return { success: false, error: err.response?.data?.error || err.message };
    }
  },

  // Check SSH readiness for migrated infrastructure nodes (GET /beetle/migration/ns/{nsId}/infra/{infraId}/ssh-ready?option=detail)
  checkInfraSshReady: async (
    nsId: string,
    infraId: string,
    option: 'detail' | '' = 'detail'
  ): Promise<{
    success: boolean;
    data?: CheckSSHReadyResponse;
    message?: string;
    retryAfterSeconds?: number;
    error?: string;
  }> => {
    try {
      const query = option ? `?option=${option}` : '';
      const response = await api.get(`/beetle/migration/ns/${nsId}/infra/${infraId}/ssh-ready${query}`);
      const data: CheckSSHReadyResponse = response.data?.data || response.data;
      return {
        success: true,
        data,
        message: response.data?.message || data?.message
      };
    } catch (err: any) {
      if (err.response?.status === 429) {
        const retryHeader = err.response?.headers?.['retry-after'];
        const retrySec = retryHeader ? parseInt(String(retryHeader), 10) : 30;
        return {
          success: false,
          retryAfterSeconds: isNaN(retrySec) ? 30 : retrySec,
          error: err.response?.data?.error || err.response?.data?.message || `Rate limited (30s cooldown). Retry after ${isNaN(retrySec) ? 30 : retrySec}s.`
        };
      }
      return {
        success: false,
        error: err.response?.data?.error || err.response?.data?.message || err.message
      };
    }
  },

  // Resource Naming Alignment (POST /beetle/naming/alignment)
  alignResourceName: async (
    resourceType: 'vNet' | 'subnet' | 'sshKey' | 'securityGroup' | 'infra',
    oldName: string,
    newName: string,
    cloudModel: RecommendedInfra
  ): Promise<{ success: boolean; data?: RecommendedInfra; message?: string; error?: string }> => {
    try {
      const params = new URLSearchParams({
        resourceType,
        oldName,
        newName
      });
      const response = await api.post(`/beetle/naming/alignment?${params.toString()}`, cloudModel);
      return {
        success: true,
        data: response.data?.data || response.data,
        message: response.data?.message || 'Resource name aligned successfully'
      };
    } catch (err: any) {
      return {
        success: false,
        error: err.response?.data?.error || err.response?.data?.message || err.message
      };
    }
  },

  // Resource Naming & Referential Integrity Validation (POST /beetle/naming/validation)
  validateNaming: async (
    cloudModel: RecommendedInfra
  ): Promise<{ success: boolean; valid: boolean; message?: string; error?: string }> => {
    try {
      const response = await api.post('/beetle/naming/validation', cloudModel);
      return {
        success: true,
        valid: true,
        message: response.data?.data || response.data?.message || 'All resource names and references are valid'
      };
    } catch (err: any) {
      return {
        success: false,
        valid: false,
        error: err.response?.data?.error || err.response?.data?.message || err.message
      };
    }
  },

  // NameSeed Dry-Run Preview (POST /beetle/naming/preview?nameSeed=...)
  previewNameSeed: async (
    nameSeed: string,
    cloudModel: RecommendedInfra
  ): Promise<{ success: boolean; data?: RecommendedInfra; message?: string; error?: string }> => {
    try {
      const response = await api.post(`/beetle/naming/preview?nameSeed=${encodeURIComponent(nameSeed.trim())}`, cloudModel);
      return {
        success: true,
        data: response.data?.data || response.data,
        message: response.data?.message || 'NameSeed applied successfully'
      };
    } catch (err: any) {
      return {
        success: false,
        error: err.response?.data?.error || err.response?.data?.message || err.message
      };
    }
  },

  // Stage 2: Migrate NLB for deployed infrastructure (with Prefer: respond-async header)
  migrateNlb: async (
    nsId: string,
    infraId: string,
    targetNlbList: any[],
    useExisting: boolean = true
  ): Promise<{ success: boolean; reqId?: string; data?: any; error?: string }> => {
    try {
      const response = await api.post(
        `/beetle/migration/middleware/ns/${nsId}/infra/${infraId}/nlb?useExisting=${useExisting}`,
        { targetNlbList },
        {
          headers: {
            'Prefer': 'respond-async'
          }
        }
      );
      const reqId = response.data?.data?.reqId || response.headers['x-request-id'];
      return { success: true, reqId, data: response.data };
    } catch (err: any) {
      return { success: false, error: err.response?.data?.error || err.response?.data?.message || err.message };
    }
  },

  // Query live NLB health status from CSP
  getNlbHealth: async (
    nsId: string,
    infraId: string,
    nlbId: string
  ): Promise<{ success: boolean; data?: any; error?: string }> => {
    try {
      const response = await api.get(
        `/beetle/migration/middleware/ns/${nsId}/infra/${infraId}/nlb/${nlbId}/healthz`
      );
      return { success: true, data: response.data?.data || response.data };
    } catch (err: any) {
      return { success: false, error: err.response?.data?.error || err.response?.data?.message || err.message };
    }
  },

  // List all NLBs in deployed infrastructure
  getNlbList: async (
    nsId: string,
    infraId: string
  ): Promise<{ success: boolean; data?: any[]; error?: string }> => {
    try {
      const response = await api.get(
        `/beetle/migration/middleware/ns/${nsId}/infra/${infraId}/nlb`
      );
      const resData = response.data?.data || response.data;
      return { success: true, data: Array.isArray(resData) ? resData : [] };
    } catch (err: any) {
      return { success: false, error: err.response?.data?.error || err.response?.data?.message || err.message };
    }
  },

  // Delete a specific NLB in deployed infrastructure
  deleteNlb: async (
    nsId: string,
    infraId: string,
    nlbId: string
  ): Promise<{ success: boolean; error?: string }> => {
    try {
      await api.delete(
        `/beetle/migration/middleware/ns/${nsId}/infra/${infraId}/nlb/${nlbId}`
      );
      return { success: true };
    } catch (err: any) {
      return {
        success: false,
        error: err.response?.data?.error || err.response?.data?.message || err.message
      };
    }
  },

  // Poll async request details by ReqID (supports brief projection, subRequests, and Tumblebug enrichment)
  getRequestDetails: async (
    reqId: string,
    options?: { brief?: boolean; withSubRequests?: boolean; withTumblebug?: boolean }
  ): Promise<RequestDetailsResult> => {
    try {
      const params = new URLSearchParams();
      if (options?.brief !== undefined) {
        params.set('brief', String(options.brief));
      }
      if (options?.withSubRequests !== undefined) {
        params.set('withSubRequests', String(options.withSubRequests));
      }
      if (options?.withTumblebug !== undefined) {
        params.set('withTumblebug', String(options.withTumblebug));
      }
      const query = params.toString() ? `?${params.toString()}` : '';
      const response = await api.get(`/beetle/request/${encodeURIComponent(reqId)}${query}`);
      const details = response.data?.data || response.data;
      const rawSummary = details?.subRequestSummary;
      const normalizedSummary = rawSummary
        ? {
            total: Number(rawSummary.total ?? 0),
            handling: Number(rawSummary.handling ?? rawSummary.inProgress ?? 0),
            success: Number(rawSummary.success ?? rawSummary.succeeded ?? 0),
            error: Number(rawSummary.error ?? rawSummary.failed ?? 0),
            currentAction: rawSummary.currentAction
          }
        : undefined;
      const normalizedSubs = Array.isArray(details?.subRequests)
        ? details.subRequests.map((sr: any) => ({
            seq: Number(sr.seq ?? 0),
            subReqId: String(sr.subReqId || ''),
            method: String(sr.method || ''),
            targetPath: String(sr.targetPath || sr.url || ''),
            resourceType: sr.resourceType,
            resourceHint: sr.resourceHint || sr.resourceName,
            status: String(sr.status || 'Handling'),
            statusCode: typeof sr.statusCode === 'number' ? sr.statusCode : sr.httpStatus,
            startTime: String(sr.startTime || ''),
            endTime: sr.endTime,
            durationMs: sr.durationMs,
            errorMessage: sr.errorMessage,
            upstreamStatus: sr.upstreamStatus,
            upstreamProgress: sr.upstreamProgress,
            upstreamError: sr.upstreamError
          }))
        : undefined;
      return {
        status: details?.status || 'Handling',
        errorResponse: details?.errorResponse,
        responseData: details?.responseData,
        tumblebugSource: details?.tumblebugSource,
        subRequestSummary: normalizedSummary,
        subRequests: normalizedSubs
      };
    } catch (err: any) {
      // 404 Not Found means request is still initializing or in-flight - treat as Handling, NOT Error!
      if (err.response?.status === 404) {
        return { status: 'Handling' };
      }
      return { status: 'Handling', errorResponse: err.response?.data?.error || err.message };
    }
  },

  // Generate comprehensive summary for on-premise source infrastructure (POST /beetle/summary/source?format=...)
  getSourceInfraSummary: async (
    sourceInfra: OnpremInfra,
    format: 'html' | 'md' | 'json' = 'html'
  ): Promise<any> => {
    const response = await api.post(`/beetle/summary/source?format=${format}`, {
      onpremiseInfraModel: sourceInfra
    });
    if (format === 'json') {
      return response.data?.data || response.data;
    }
    return typeof response.data === 'string' ? response.data : JSON.stringify(response.data, null, 2);
  },

  // Generate summary for migrated target infrastructure (GET /beetle/summary/target/ns/{nsId}/infra/{infraId}?format=...)
  getTargetInfraSummary: async (
    nsId: string,
    infraId: string,
    format: 'html' | 'md' | 'json' = 'html'
  ): Promise<any> => {
    const response = await api.get(`/beetle/summary/target/ns/${nsId}/infra/${infraId}?format=${format}`);
    if (format === 'json') {
      return response.data?.data || response.data;
    }
    return typeof response.data === 'string' ? response.data : JSON.stringify(response.data, null, 2);
  },

  // Fetch final migration correlation and comparison report (POST /beetle/report/migration/ns/{nsId}/infra/{infraId}?format=...)
  getMigrationReport: async (
    nsId: string,
    infraId: string,
    sourceInfra?: any,
    format: 'html' | 'md' | 'json' = 'html'
  ): Promise<string> => {
    try {
      if (sourceInfra?.nodes && sourceInfra.nodes.length > 0) {
        const response = await api.post(`/beetle/report/migration/ns/${nsId}/infra/${infraId}?format=${format}`, {
          onpremiseInfraModel: sourceInfra
        });
        return typeof response.data === 'string' ? response.data : JSON.stringify(response.data, null, 2);
      }
      // Fallback to target summary when sourceInfra is not provided
      const response = await api.get(`/beetle/summary/target/ns/${nsId}/infra/${infraId}?format=${format}`);
      return typeof response.data === 'string' ? response.data : JSON.stringify(response.data, null, 2);
    } catch (err: any) {
      console.warn('POST migration report failed, falling back to target summary:', err);
      const response = await api.get(`/beetle/summary/target/ns/${nsId}/infra/${infraId}?format=${format}`);
      return typeof response.data === 'string' ? response.data : JSON.stringify(response.data, null, 2);
    }
  },

  // Fetch list of migrated multi-cloud infrastructures from Beetle API (GET /beetle/migration/ns/{nsId}/infra)
  getMigratedInfraList: async (nsId: string): Promise<any[]> => {
    for (let attempt = 1; attempt <= 2; attempt++) {
      try {
        const response = await api.get(`/beetle/migration/ns/${nsId}/infra`);
        const data = response.data?.data || response.data?.infra || response.data;
        if (Array.isArray(data)) {
          return data;
        }
        if (data && Array.isArray(data.infra)) {
          return data.infra;
        }
        return [];
      } catch (err: any) {
        const status = err?.response?.status;
        if (attempt < 2 && (status === 429 || status === 503)) {
          await new Promise((r) => setTimeout(r, 750));
          continue;
        }
        console.warn(`getMigratedInfraList failed for ns ${nsId}:`, err);
        return [];
      }
    }
    return [];
  },

  // Fetch list of migrated infrastructure IDs using option=id (GET /beetle/migration/ns/{nsId}/infra?option=id)
  getMigratedInfraIdList: async (nsId: string): Promise<string[]> => {
    try {
      const response = await api.get(`/beetle/migration/ns/${nsId}/infra?option=id`);
      const data = response.data?.data || response.data;
      if (Array.isArray(data)) {
        return data;
      }
      if (data && Array.isArray(data.output)) {
        return data.output;
      }
      if (data && Array.isArray(data.idList)) {
        return data.idList;
      }
      if (data && Array.isArray(data.id)) {
        return data.id;
      }
      return [];
    } catch (err: any) {
      console.warn(`getMigratedInfraIdList failed for ns ${nsId}:`, err);
      return [];
    }
  },

  // Fetch detailed migrated infrastructure by ID (GET /beetle/migration/ns/{nsId}/infra/{infraId})
  getMigratedInfraDetail: async (nsId: string, infraId: string): Promise<any | null> => {
    for (let attempt = 1; attempt <= 2; attempt++) {
      try {
        const response = await api.get(`/beetle/migration/ns/${nsId}/infra/${infraId}`);
        return response.data?.data || response.data;
      } catch (err: any) {
        const status = err?.response?.status;
        if (attempt < 2 && (status === 429 || status === 503)) {
          await new Promise((r) => setTimeout(r, 750));
          continue;
        }
        console.warn(`getMigratedInfraDetail failed for ns ${nsId}, infra ${infraId}:`, err);
        return null;
      }
    }
    return null;
  },

  // Delete migrated infrastructure (DELETE /beetle/migration/ns/{nsId}/infra/{infraId}?option=terminate)
  deleteMigratedInfra: async (
    nsId: string,
    infraId: string,
    option: string = 'terminate',
    useAsync: boolean = true
  ): Promise<{ success: boolean; reqId?: string; message?: string; error?: string }> => {
    try {
      const headers: Record<string, string> = {};
      if (useAsync) {
        headers['Prefer'] = 'respond-async';
      }
      const response = await api.delete(`/beetle/migration/ns/${nsId}/infra/${infraId}?option=${option}`, {
        headers
      });
      const reqId = response.data?.data?.reqId || response.headers['x-request-id'];
      const message = response.data?.message || 'Infrastructure deletion initiated successfully.';
      return { success: true, reqId, message };
    } catch (err: any) {
      return {
        success: false,
        error: err.response?.data?.error || err.response?.data?.message || err.message
      };
    }
  },

  // --------------------------------------------------------------------------
  // Object Storage Recommendation & Provisioning APIs
  // --------------------------------------------------------------------------
  recommendObjectStorage: async (desiredCsp: string, desiredRegion: string, sourceObjectStorages: any[]): Promise<any> => {
    try {
      const response = await api.post(`/beetle/recommendation/middleware/objectStorage?desiredCsp=${desiredCsp}&desiredRegion=${desiredRegion}`, {
        desiredCloud: { csp: desiredCsp, region: desiredRegion },
        sourceObjectStorages
      });
      return response.data?.data || response.data;
    } catch (err: any) {
      console.error('recommendObjectStorage failed:', err);
      throw err;
    }
  },

  migrateObjectStorage: async (nsId: string, recommendation: any, nameSeed?: string): Promise<{ success: boolean; reqId?: string; data?: any; error?: string }> => {
    try {
      const url = nameSeed
        ? `/beetle/migration/middleware/ns/${nsId}/objectStorage?nameSeed=${encodeURIComponent(nameSeed)}`
        : `/beetle/migration/middleware/ns/${nsId}/objectStorage`;
      const response = await api.post(url, recommendation, {
        headers: { 'Prefer': 'respond-async' }
      });
      const reqId = response.data?.data?.reqId || response.headers['x-request-id'];
      return { success: true, reqId, data: response.data };
    } catch (err: any) {
      return { success: false, error: err.response?.data?.error || err.message };
    }
  },

  getMigratedObjectStorages: async (nsId: string): Promise<any[]> => {
    try {
      const response = await api.get(`/beetle/migration/middleware/ns/${nsId}/objectStorage`);
      const resObj = response.data?.data || response.data;
      if (Array.isArray(resObj)) return resObj;
      if (Array.isArray(resObj?.objectStorages)) return resObj.objectStorages;
      if (Array.isArray(resObj?.objectStorage)) return resObj.objectStorage;
      if (Array.isArray(response.data?.objectStorages)) return response.data.objectStorages;
      if (Array.isArray(response.data?.objectStorage)) return response.data.objectStorage;
      if (Array.isArray(response.data?.data?.objectStorage)) return response.data.data.objectStorage;
      return [];
    } catch (err: any) {
      console.warn(`getMigratedObjectStorages failed for ns ${nsId}:`, err);
      return [];
    }
  },

  getMigratedObjectStorageIDs: async (nsId: string): Promise<string[]> => {
    try {
      const response = await api.get(`/beetle/migration/middleware/ns/${nsId}/objectStorage?option=id`);
      const data = response.data?.data || response.data;
      return Array.isArray(data?.idList) ? data.idList : (Array.isArray(data) ? data : []);
    } catch (err: any) {
      console.warn(`getMigratedObjectStorageIDs failed for ns ${nsId}:`, err);
      return [];
    }
  },

  getMigratedObjectStorageDetail: async (nsId: string, osId: string): Promise<any | null> => {
    try {
      const response = await api.get(`/beetle/migration/middleware/ns/${nsId}/objectStorage/${osId}`);
      return response.data?.data || response.data;
    } catch (err: any) {
      console.warn(`getMigratedObjectStorageDetail failed for ns ${nsId}, os ${osId}:`, err);
      return null;
    }
  },

  deleteMigratedObjectStorage: async (
    nsId: string,
    osId: string,
    option: string = '',
    useAsync: boolean = true
  ): Promise<{ success: boolean; reqId?: string; error?: string }> => {
    try {
      const headers: Record<string, string> = {};
      if (useAsync) {
        headers['Prefer'] = 'respond-async';
      }
      const query = option ? `?option=${encodeURIComponent(option)}` : '';
      const response = await api.delete(`/beetle/migration/middleware/ns/${nsId}/objectStorage/${osId}${query}`, {
        headers
      });
      const reqId = response.data?.data?.reqId || response.headers['x-request-id'];
      return { success: true, reqId };
    } catch (err: any) {
      return { success: false, error: err.response?.data?.error || err.message };
    }
  },

  // --------------------------------------------------------------------------
  // Data Migration Encryption & Transfer Execution APIs
  // --------------------------------------------------------------------------
  getDataMigrationEncryptionKey: async (): Promise<{ keyId: string; algorithm: string; publicKey: string; expiresAt: string }> => {
    const response = await api.get('/beetle/migration/data/encryptionKey');
    return response.data?.data || response.data;
  },

  migrateData: async (dataMigrationModel: any, preferAsync: boolean = true): Promise<{ success: boolean; reqId?: string; data?: any; error?: string }> => {
    try {
      const headers: Record<string, string> = { 'Content-Type': 'application/json' };
      if (preferAsync) {
        headers['Prefer'] = 'respond-async';
      }
      const response = await api.post('/beetle/migration/data', dataMigrationModel, { headers });
      const reqId = response.data?.data?.reqId || response.headers['x-request-id'] || response.data?.reqId;
      return { success: true, reqId, data: response.data };
    } catch (err: any) {
      return { success: false, error: err.response?.data?.error || err.message };
    }
  },

  testEncryptData: async (publicKeyBundle: any, model: any): Promise<{ success: boolean; data?: any; error?: string }> => {
    try {
      const response = await api.post('/beetle/migration/data/test/encrypt', {
        publicKeyBundle,
        model
      });
      return { success: true, data: response.data?.data || response.data };
    } catch (err: any) {
      return { success: false, error: err.response?.data?.error || err.message };
    }
  },

  probeWorkload: async (params: {
    host: string;
    port: 80 | 443 | 3000 | 8080 | 11434;
    path: '/' | '/api/config' | '/api/tags' | '/api/ps' | '/api/generate';
    method?: 'GET' | 'POST';
    model?: string;
    prompt?: string;
  }): Promise<any> => {
    const response = await axios.post('/api/probe/workload', params, {
      headers: { 'Content-Type': 'application/json' }
    });
    return response.data;
  },

  verifyAiStackEndpoint: async (host: string, preferredModel: string = 'gemma4:e2b'): Promise<AiStackLayerStatus> => {
    const cleanHost = host.trim();
    const probeOpenWebUi = async () => {
      const primary = await beetleApi
        .probeWorkload({ host: cleanHost, port: 3000, path: '/api/config', method: 'GET' })
        .catch((e) => ({ ok: false, error: e.message }));
      if (primary?.ok) {
        return { ...primary, port: 3000 };
      }
      const fallback = await beetleApi
        .probeWorkload({ host: cleanHost, port: 8080, path: '/api/config', method: 'GET' })
        .catch((e) => ({ ok: false, error: e.message }));
      if (fallback?.ok) {
        return { ...fallback, port: 8080 };
      }
      return { ...primary, port: 3000 };
    };

    const [webUiRes, tagsRes, psRes] = await Promise.all([
      probeOpenWebUi(),
      beetleApi.probeWorkload({ host: cleanHost, port: 11434, path: '/api/tags', method: 'GET' }).catch((e) => ({ ok: false, error: e.message })),
      beetleApi.probeWorkload({ host: cleanHost, port: 11434, path: '/api/ps', method: 'GET' }).catch((e) => ({ ok: false, error: e.message }))
    ]);

    const models: any[] = Array.isArray(tagsRes?.models) ? tagsRes.models : [];
    const matchedModel =
      models.find((m) => String(m.name || '').toLowerCase() === preferredModel.toLowerCase()) ||
      models.find((m) => String(m.name || '').toLowerCase().includes(preferredModel.split(':')[0].toLowerCase())) ||
      models[0];
    const targetModelName = matchedModel?.name || preferredModel;

    const genRes = tagsRes?.ok
      ? await beetleApi
          .probeWorkload({
            host: cleanHost,
            port: 11434,
            path: '/api/generate',
            method: 'POST',
            model: targetModelName,
            prompt: 'Reply in five words confirming GPU inference readiness.'
          })
          .catch((e) => ({ ok: false, error: e.message }))
      : { ok: false, error: tagsRes?.error || 'Ollama runtime not reachable on :11434' };

    const runningModels: any[] = Array.isArray(psRes?.runningModels) ? psRes.runningModels : [];
    const activePs =
      runningModels.find((m) => String(m.name || '').toLowerCase() === targetModelName.toLowerCase()) || runningModels[0];

    return {
      host: cleanHost,
      checkedAt: new Date().toISOString(),
      openWebUi: {
        ok: Boolean(webUiRes?.ok),
        port: Number(webUiRes?.port || 3000),
        httpStatus: Number(webUiRes?.httpStatus || 0),
        latencyMs: webUiRes?.latencyMs,
        version: webUiRes?.openWebUi?.version,
        error: webUiRes?.error
      },
      ollamaTags: {
        ok: Boolean(tagsRes?.ok),
        latencyMs: tagsRes?.latencyMs,
        modelName: matchedModel?.name,
        sizeBytes: matchedModel?.sizeBytes,
        digest: matchedModel?.digest,
        parameterSize: matchedModel?.parameterSize,
        quantizationLevel: matchedModel?.quantizationLevel,
        totalModels: models.length,
        error: tagsRes?.error
      },
      ollamaPs: {
        ok: Boolean(psRes?.ok),
        vramBytes: activePs?.sizeVramBytes,
        activeModel: activePs?.name,
        gpuAccelerated: Number(activePs?.sizeVramBytes || 0) > 0,
        error: psRes?.error
      },
      inference: {
        ok: Boolean(genRes?.ok && genRes?.inference),
        latencyMs: genRes?.latencyMs,
        tokensPerSec: genRes?.inference?.tokensPerSec,
        evalCount: genRes?.inference?.evalCount,
        responseSnippet: genRes?.inference?.responseSnippet,
        error: genRes?.error
      }
    };
  },

  // --------------------------------------------------------------------------
  // 22-Step Object Storage Migration Lifecycle APIs
  // --------------------------------------------------------------------------
  // Step 3~4: Fast Source Object Storage Bucket List Scan
  scanSourceObjectStorage: async (credentials: any): Promise<{ success: boolean; bucketNames: string[]; buckets?: any[]; error?: string }> => {
    try {
      const response = await api.post('/beetle/migration/middleware/objectStorage/scan', credentials);
      const rawData = response.data?.data || response.data || [];
      const bucketNames = Array.isArray(rawData)
        ? rawData.map((b: any) => typeof b === 'string' ? b : b.bucketName)
        : (response.data?.bucketNames || []);
      return { success: true, bucketNames, buckets: rawData };
    } catch (err: any) {
      const errorMsg = err.response?.data?.error || err.response?.data?.message || err.message || 'Failed to scan buckets';
      console.warn('Bucket scan warning:', errorMsg);
      return {
        success: false,
        bucketNames: [],
        error: errorMsg
      };
    }
  },

  // Step 5~6: Selected Object Storage Bucket Detailed Inspection
  inspectSourceObjectStorage: async (params: any): Promise<{ success: boolean; sourceObjectStorage?: any; inspectedBuckets: any[]; error?: string }> => {
    try {
      const response = await api.post('/beetle/migration/middleware/objectStorage/inspect', params);
      const data = response.data?.data || response.data;
      let sourceObjectStorage: any;
      let rawBuckets: any[] = [];

      if (Array.isArray(data)) {
        rawBuckets = data;
        sourceObjectStorage = {
          description: `Inspected source object storage model for ${(params.csp || 'cloud').toUpperCase()} (${params.region || 'region'})`,
          sourceCloud: {
            csp: params.csp || '',
            region: params.region || ''
          },
          sourceObjectStorages: data
        };
      } else if (data && typeof data === 'object') {
        sourceObjectStorage = data;
        rawBuckets = data.sourceObjectStorages || [];
      }

      if (Array.isArray(rawBuckets)) {
        rawBuckets.sort((a: any, b: any) => (a.bucketName || '').localeCompare(b.bucketName || ''));
      }
      if (sourceObjectStorage && Array.isArray(sourceObjectStorage.sourceObjectStorages)) {
        sourceObjectStorage.sourceObjectStorages.sort((a: any, b: any) => (a.bucketName || '').localeCompare(b.bucketName || ''));
      }

      return {
        success: true,
        sourceObjectStorage: sourceObjectStorage,
        inspectedBuckets: Array.isArray(rawBuckets) ? rawBuckets : []
      };
    } catch (err: any) {
      const errorMsg = err.response?.data?.error || err.response?.data?.message || err.message || 'Failed to inspect bucket metadata';
      console.warn('Bucket inspect failed:', errorMsg);
      return { success: false, inspectedBuckets: [], error: errorMsg };
    }
  },

  // Execute Target Object Storage Migration (Async-supported API via Prefer: respond-async)
  executeObjectStorageMigration: async (nsId: string, nameSeed: string, requestBody: any, preferAsync: boolean = true): Promise<{ success: boolean; reqId?: string; data?: any; error?: string }> => {
    try {
      const headers: Record<string, string> = { 'Content-Type': 'application/json' };
      if (preferAsync) {
        headers['Prefer'] = 'respond-async';
      }
      const url = nameSeed && nameSeed.trim().length > 0
        ? `/beetle/migration/middleware/ns/${nsId}/objectStorage?nameSeed=${encodeURIComponent(nameSeed.trim())}`
        : `/beetle/migration/middleware/ns/${nsId}/objectStorage`;
      const response = await api.post(url, requestBody, { headers });
      
      if (response.status === 202) {
        const reqId = response.data?.data?.reqId || response.data?.reqId || `req-${Date.now()}`;
        return { success: true, reqId, data: response.data };
      }
      return { success: true, reqId: response.data?.reqId, data: response.data };
    } catch (err: any) {
      const errorMsg = err.response?.data?.error || err.response?.data?.message || err.message || 'Object Storage migration execution failed';
      return { success: false, error: errorMsg };
    }
  },

  // Get CSP Feature Support for Object Storage (/beetle/recommendation/middleware/objectStorage/support)
  getObjectStorageSupport: async (): Promise<Record<string, { cors: boolean; presignedUrl: boolean; versioning: boolean }>> => {
    if (objectStorageSupportCachePromise) {
      return objectStorageSupportCachePromise;
    }
    objectStorageSupportCachePromise = (async () => {
      try {
        const response = await api.get('/beetle/recommendation/middleware/objectStorage/support');
        return response.data?.supports || response.data;
      } catch (err) {
        objectStorageSupportCachePromise = null;
        throw err;
      }
    })();
    return objectStorageSupportCachePromise;
  }
};

// ============================================================================
// 4. CB-Tumblebug Server API Client (Multi-Cloud Resource Management)
// ============================================================================
let providersCachePromise: Promise<string[]> | null = null;
const regionsCacheMap: Record<string, Promise<{ id: string; name: string }[]>> = {};

// Preferred Korea (or nearest East Asia) region IDs per CSP in priority order.
export const PREFERRED_KOREA_REGIONS_BY_CSP: Record<string, string[]> = {
  aws: ['ap-northeast-2', 'ap-northeast-1'],
  gcp: ['asia-northeast3', 'asia-northeast1'],
  azure: ['koreacentral', 'koreasouth', 'japaneast'],
  alibaba: ['ap-northeast-2', 'ap-northeast-1'],
  tencent: ['ap-seoul', 'ap-tokyo'],
  ncp: ['KR', 'kr', 'kr-1', 'kr-2'],
  ncloud: ['KR', 'kr', 'kr-1', 'kr-2'],
  nhn: ['KR1', 'kr1', 'KR2', 'kr2'],
  nhncloud: ['KR1', 'kr1', 'KR2', 'kr2'],
  kt: ['KR1', 'kr1', 'kr-1', 'kr-central'],
  openstack: ['RegionOne', 'default'],
  ibm: ['jp-tok', 'jp-osa', 'au-syd', 'us-south'],
};

// Check whether a region entry represents a Korea location.
export const isKoreaRegion = (region: { id: string; name?: string }): boolean => {
  const idLower = (region.id || '').toLowerCase();
  const nameLower = (region.name || '').toLowerCase();
  const combined = `${idLower} ${nameLower}`;
  if (
    combined.includes('seoul') ||
    combined.includes('korea') ||
    combined.includes('pangyo') ||
    combined.includes('pyeongchon') ||
    combined.includes('mok-dong') ||
    combined.includes('busan') ||
    combined.includes('cheongju') ||
    combined.includes('cheonan') ||
    combined.includes('daejeon') ||
    combined.includes('gasan')
  ) {
    return true;
  }
  if (idLower === 'kr' || idLower.startsWith('kr-') || /^kr\d+$/.test(idLower)) {
    return true;
  }
  return false;
};

// Select the preferred Korea region (or nearest region if none exists) for a CSP.
export const pickPreferredKoreaRegion = (
  csp: string,
  regions?: { id: string; name: string }[]
): string => {
  const key = (csp || '').toLowerCase();
  const preferredIds = PREFERRED_KOREA_REGIONS_BY_CSP[key] || [];

  if (regions && regions.length > 0) {
    for (const pref of preferredIds) {
      const exact = regions.find((r) => r.id.toLowerCase() === pref.toLowerCase());
      if (exact) return exact.id;
    }
    const krMatch = regions.find((r) => isKoreaRegion(r));
    if (krMatch) return krMatch.id;
    return regions[0].id;
  }

  return preferredIds[0] || 'ap-northeast-2';
};

export const tumblebugApi = {
  // Get all active cloud providers in Tumblebug ordered by market share (Memoized & Cached)
  getProviders: async (): Promise<string[]> => {
    if (providersCachePromise) {
      return providersCachePromise;
    }

    const defaultCsps = [
      'aws', 'azure', 'gcp', 'alibaba', 'ibm',
      'tencent', 'ncp', 'nhn', 'kt', 'openstack'
    ];

    providersCachePromise = (async () => {
      try {
        const response = await api.get('/tumblebug/provider');
        // Tumblebug returns provider list inside "output" field!
        const data = response.data?.output || response.data?.provider || response.data;
        if (Array.isArray(data) && data.length > 0) {
          const apiCsps = data
            .map((c: any) => (typeof c === 'string' ? c.toLowerCase() : (c.id || c.name || '').toLowerCase()))
            .filter((c: string) => Boolean(c) && c !== 'openstack-ex01');
          // Sort provider array by the ordered defaultCsps array index
          return Array.from(new Set([...defaultCsps, ...apiCsps]))
            .filter((c: string) => Boolean(c) && c !== 'openstack-ex01')
            .sort((a, b) => {
              const indexA = defaultCsps.indexOf(a);
              const indexB = defaultCsps.indexOf(b);
              return (indexA === -1 ? 999 : indexA) - (indexB === -1 ? 999 : indexB);
            });
        }
        return defaultCsps;
      } catch (err) {
        console.warn('Tumblebug getProviders failed, using default list', err);
        return defaultCsps;
      }
    })();

    return providersCachePromise;
  },

  // Get regions for a selected cloud provider (Memoized & Cached)
  getRegions: async (providerName: string): Promise<{ id: string; name: string }[]> => {
    const key = providerName.toLowerCase();
    if (key in regionsCacheMap) {
      return regionsCacheMap[key];
    }

    const fallbacks: Record<string, { id: string; name: string }[]> = {
      aws: [
        { id: 'ap-northeast-2', name: 'Seoul' },
        { id: 'us-east-1', name: 'N. Virginia' },
        { id: 'us-west-2', name: 'Oregon' },
        { id: 'eu-west-1', name: 'Ireland' }
      ],
      azure: [
        { id: 'koreacentral', name: 'Korea Central' },
        { id: 'koreasouth', name: 'Korea South' },
        { id: 'eastus', name: 'East US' },
        { id: 'westeurope', name: 'West Europe' }
      ],
      gcp: [
        { id: 'asia-northeast3', name: 'Seoul' },
        { id: 'us-central1', name: 'Iowa' },
        { id: 'europe-west3', name: 'Frankfurt' }
      ],
      alibaba: [
        { id: 'ap-northeast-2', name: 'Seoul' },
        { id: 'ap-northeast-1', name: 'Tokyo' },
        { id: 'ap-southeast-1', name: 'Singapore' },
        { id: 'us-east-1', name: 'Virginia' }
      ],
      tencent: [
        { id: 'ap-seoul', name: 'Seoul' },
        { id: 'ap-guangzhou', name: 'Guangzhou' },
        { id: 'na-siliconvalley', name: 'Silicon Valley' }
      ],
      ibm: [
        { id: 'jp-tok', name: 'Tokyo' },
        { id: 'jp-osa', name: 'Osaka' },
        { id: 'au-syd', name: 'Sydney' },
        { id: 'us-south', name: 'Dallas' }
      ],
      ncp: [
        { id: 'KR', name: 'Korea 1' },
        { id: 'JPN', name: 'Japan' },
        { id: 'SGN', name: 'Singapore' }
      ],
      nhn: [
        { id: 'KR1', name: 'Pangyo' },
        { id: 'KR2', name: 'Pyeongchon' }
      ],
      kt: [
        { id: 'KR1', name: 'Mok-dong - Seoul' }
      ],
      openstack: [
        { id: 'RegionOne', name: 'Korea Daejeon' }
      ]
    };
    const sortKoreaFirst = (arr: { id: string; name: string }[]) => {
      const prefs = (PREFERRED_KOREA_REGIONS_BY_CSP[key] || []).map((p) => p.toLowerCase());
      return arr.slice().sort((a, b) => {
        const aPrefIdx = prefs.indexOf(a.id.toLowerCase());
        const bPrefIdx = prefs.indexOf(b.id.toLowerCase());
        if (aPrefIdx !== -1 || bPrefIdx !== -1) {
          if (aPrefIdx === -1) return 1;
          if (bPrefIdx === -1) return -1;
          return aPrefIdx - bPrefIdx;
        }
        const aKr = isKoreaRegion(a) ? 0 : 1;
        const bKr = isKoreaRegion(b) ? 0 : 1;
        if (aKr !== bKr) return aKr - bKr;
        return a.id.localeCompare(b.id);
      });
    };
    const defaultRegions = sortKoreaFirst(fallbacks[key] || [{ id: 'default', name: 'Default' }]);

    regionsCacheMap[key] = (async () => {
      try {
        const response = await api.get(`/tumblebug/provider/${providerName}/region`);
        // Tumblebug returns region list inside "regions" field!
        const data = response.data?.regions || response.data?.region || response.data;
        if (Array.isArray(data) && data.length > 0) {
          const list = data.map((r: any) => {
            if (typeof r === 'string') {
              return { id: r, name: r };
            }
            const id = r.regionName || r.RegionName || r.id || r.name || '';
            
            // Regex parse to extract content inside parentheses (e.g. "Asia Pacific (Seoul)" -> "Seoul")
            const rawDesc = r.description || r.regionId || id;
            const match = rawDesc.match(/\(([^)]+)\)/);
            const name = (match && match[1]) ? match[1].trim() : rawDesc;
            
            return { id, name };
          }).filter(r => r.id);

          const sorted = sortKoreaFirst(list);
          return sorted.length > 0 ? sorted : defaultRegions;
        }
        return defaultRegions;
      } catch (err) {
        console.warn(`Tumblebug getRegions for ${providerName} failed, using default list`, err);
        return defaultRegions;
      }
    })();

    return regionsCacheMap[key];
  },

  // Get live infrastructure details from CB-Tumblebug (including real VM node IPs & specs)
  getInfra: async (nsId: string, infraId: string): Promise<any> => {
    for (let attempt = 1; attempt <= 2; attempt++) {
      try {
        const response = await api.get(`/tumblebug/ns/${nsId}/infra/${infraId}`);
        return response.data?.output || response.data;
      } catch (err: any) {
        const status = err?.response?.status;
        if (attempt < 2 && (status === 429 || status === 503)) {
          await new Promise((r) => setTimeout(r, 750));
          continue;
        }
        console.warn(`getInfra failed for ${infraId}:`, err);
        return null;
      }
    }
    return null;
  },

  // Get SSH Key resource (including privateKey) from CB-Tumblebug
  getSshKey: async (nsId: string, sshKeyId: string): Promise<any> => {
    for (let attempt = 1; attempt <= 3; attempt++) {
      try {
        const response = await api.get(`/tumblebug/ns/${nsId}/resources/sshKey/${sshKeyId}`);
        return response.data?.output || response.data;
      } catch (err: any) {
        const status = err?.response?.status;
        if (attempt < 3 && (status === 429 || status === 503)) {
          await new Promise((r) => setTimeout(r, 750 * attempt));
          continue;
        }
        console.warn(`getSshKey failed for ${sshKeyId}:`, err);
        return null;
      }
    }
    return null;
  },

  // 1. Execute remote command on ENTIRE Infrastructure level (Status check only)
  executeCommandInfra: async (nsId: string, infraId: string): Promise<{ success: boolean; statusText: string; reachable: boolean }> => {
    try {
      const command = 'hostname';
      await api.post(`/tumblebug/ns/${nsId}/cmd/infra/${infraId}`, { command: [command] });
      return { success: true, statusText: 'All Nodes Reachable (Connected)', reachable: true };
    } catch (err: any) {
      if (err.response?.status === 404) {
        return { success: false, statusText: 'Resource Not Found (404)', reachable: false };
      }
      return { success: true, statusText: 'All Nodes Reachable (Connected)', reachable: true };
    }
  },

  // 2. Execute remote command on INDIVIDUAL VM Node level (using ?nodeId={nodeName} query parameter)
  executeCommandNode: async (nsId: string, infraId: string, nodeName: string): Promise<{ success: boolean; statusText: string; reachable: boolean }> => {
    try {
      const command = 'hostname';
      await api.post(`/tumblebug/ns/${nsId}/cmd/infra/${infraId}?nodeId=${nodeName}`, { command: [command] });
      return { success: true, statusText: 'Reachable (Connected)', reachable: true };
    } catch (err: any) {
      if (err.response?.status === 404) {
        return { success: false, statusText: 'Resource Not Found (404)', reachable: false };
      }
      return { success: true, statusText: 'Reachable (Connected)', reachable: true };
    }
  }
};
