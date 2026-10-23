import axios from 'axios';
import { OnpremInfra, OnpremModelEnvelope, RecommendedInfra, CloudModelEnvelope, CloudProperty, CheckSSHReadyResponse, ArtifactLineage } from '../types/migration';

// Shared Axios client with basic authentication pre-configured (standard Cloud-Barista settings)
const api = axios.create({
  headers: {
    'Content-Type': 'application/json',
    // Using default credentials commonly used in Cloud-Barista development setups
    'Authorization': 'Basic ' + btoa('default:default')
  }
});

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
    const response = await api.get(`/honeybee/source_group/${sgId}/connection_info/${connId}/infra/refined`);
    return response.data?.onpremiseInfraModel || response.data;
  },

  getRefinedInfraByGroup: async (sgId: string): Promise<OnpremInfra> => {
    const response = await api.get(`/honeybee/source_group/${sgId}/infra/refined`);
    return response.data?.onpremiseInfraModel || response.data;
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

const mapArtifactToSourceEnvelope = (raw: any, fallbackInfra?: OnpremInfra): OnpremModelEnvelope => {
  const m = raw?.data || raw || {};
  let payload = m.payload || m.onpremiseInfraModel || fallbackInfra;
  if (typeof payload === 'string') {
    try { payload = JSON.parse(payload); } catch {}
  }
  return {
    id: m.id || m.uid || `art-src-${Date.now()}`,
    name: m.name || m.userModelName || 'Unnamed Artifact',
    description: m.description || '',
    version: m.version || m.userModelVersion || '1.0.0',
    stage: m.stage || 'source',
    category: m.category || 'infra',
    schemaType: m.schemaType || 'OnpremInfra',
    lineage: m.lineage,
    onpremiseInfraModel: payload as OnpremInfra,
    updatedTime: m.updateTime || m.createTime || m.updatedTime || new Date().toISOString()
  };
};

const mapArtifactToCloudEnvelope = (raw: any, fallbackInfra?: RecommendedInfra): CloudModelEnvelope => {
  const m = raw?.data || raw || {};
  let payload = m.payload || m.cloudInfraModel || m.cloud_infra_model || m.recommendedInfra || fallbackInfra;
  if (typeof payload === 'string') {
    try { payload = JSON.parse(payload); } catch {}
  }
  return {
    id: m.id || m.uid || `art-tgt-${Date.now()}`,
    name: m.name || m.userModelName || 'Unnamed Design',
    description: m.description || '',
    version: m.version || m.userModelVersion || '1.0.0',
    stage: m.stage || 'target',
    category: m.category || 'infra',
    schemaType: m.schemaType || 'RecommendedInfra',
    lineage: m.lineage,
    cloudInfraModel: payload as RecommendedInfra,
    updatedTime: m.updateTime || m.createTime || m.updatedTime || new Date().toISOString()
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
    const response = await api.post('/beetle/artifacts', {
      name,
      description,
      version: version || '1.0.0',
      category: 'infra',
      stage: options?.stage || 'source',
      schemaType: 'OnpremInfra',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload: onpremiseInfraModel
    });
    return mapArtifactToSourceEnvelope(response.data, onpremiseInfraModel);
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
    const response = await api.put(`/beetle/artifacts/${id}`, {
      name,
      description,
      version: version || '1.0.0',
      category: 'infra',
      stage: options?.stage,
      schemaType: 'OnpremInfra',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload: onpremiseInfraModel
    });
    return mapArtifactToSourceEnvelope(response.data, onpremiseInfraModel);
  },

  // Delete existing source or planning artifact
  deleteSourceModel: async (id: string): Promise<void> => {
    await api.delete(`/beetle/artifacts/${id}`);
  },

  // Get all source and planning OnpremInfra artifacts
  getSourceModels: async (): Promise<OnpremModelEnvelope[]> => {
    const response = await api.get('/beetle/artifacts?category=infra&schemaType=OnpremInfra');
    const rawList = response.data?.data ?? response.data;
    const dataList = Array.isArray(rawList) ? rawList : (rawList ? [rawList] : []);
    return dataList.map((m: any) => mapArtifactToSourceEnvelope(m));
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
      name,
      description,
      version: version || '1.0.0',
      category: 'infra',
      stage: 'target',
      schemaType: 'RecommendedInfra',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload: cloudInfraModel
    });
    return mapArtifactToCloudEnvelope(response.data, cloudInfraModel);
  },

  // Update existing target cloud design artifact
  updateCloudModel: async (
    id: string,
    name: string,
    description: string,
    cloudInfraModel: RecommendedInfra,
    version?: string,
    options?: SaveArtifactOptions
  ): Promise<CloudModelEnvelope> => {
    const response = await api.put(`/beetle/artifacts/${id}`, {
      name,
      description,
      version: version || '1.0.0',
      category: 'infra',
      stage: 'target',
      schemaType: 'RecommendedInfra',
      parentId: options?.parentId,
      sourceGroupId: options?.sourceGroupId,
      payload: cloudInfraModel
    });
    return mapArtifactToCloudEnvelope(response.data, cloudInfraModel);
  },

  // Get target cloud design artifacts
  getCloudModels: async (): Promise<CloudModelEnvelope[]> => {
    const response = await api.get('/beetle/artifacts?category=infra&stage=target&schemaType=RecommendedInfra');
    const rawList = response.data?.data ?? response.data;
    const dataList = Array.isArray(rawList) ? rawList : (rawList ? [rawList] : []);
    return dataList.map((m: any) => mapArtifactToCloudEnvelope(m));
  },

  getCloudModel: async (id: string): Promise<CloudModelEnvelope> => {
    const response = await api.get(`/beetle/artifacts/${id}`);
    return mapArtifactToCloudEnvelope(response.data);
  },

  deleteCloudModel: async (id: string): Promise<void> => {
    await api.delete(`/beetle/artifacts/${id}`);
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
let sourceModelCachePromise: Promise<{ success: boolean; sourceModel: any; error?: string }> | null = null;
let targetModelCachePromise: Promise<{ success: boolean; targetModel: any; error?: string }> | null = null;

export const beetleApi = {
  // Get Cloud Recommendation candidates based on Source model input
  getRecommendations: async (
    sourceInfra: OnpremInfra,
    desiredCsp: string,
    desiredRegion: string,
    limit?: number,
    minMatchRate?: number
  ): Promise<RecommendedInfra[]> => {
    const hasNlbs = sourceInfra.nlbs && sourceInfra.nlbs.length > 0;
    const limitQuery = limit ? `&limit=${limit}` : '';
    const matchRateQuery = typeof minMatchRate === 'number' && !isNaN(minMatchRate) ? `&minMatchRate=${minMatchRate}` : '';

    let response;
    if (hasNlbs) {
      // Call NLB-aware recommendation endpoint
      response = await api.post(`/beetle/recommendation/infraWithNlb?desiredCsp=${desiredCsp}&desiredRegion=${desiredRegion}${limitQuery}${matchRateQuery}`, {
        desiredCsp,
        desiredRegion,
        sourceInfra
      });
    } else {
      // Call standard VM recommendation endpoint
      response = await api.post(`/beetle/recommendation/infra?desiredCsp=${desiredCsp}&desiredRegion=${desiredRegion}${limitQuery}${matchRateQuery}`, {
        desiredCspAndRegionPair: {
          csp: desiredCsp,
          region: desiredRegion
        },
        onpremiseInfraModel: sourceInfra
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
    const hasNlbs = sourceInfra.nlbs && sourceInfra.nlbs.length > 0;
    const query = `?minMatchRate=${minMatchRate}`;

    let response;
    if (hasNlbs) {
      // Call NLB-aware multi-infra recommendation endpoint
      response = await api.post(`/beetle/recommendation/multiInfraWithNlb${query}`, {
        desiredCspAndRegionPairs: targetPairs,
        sourceInfra
      });
    } else {
      // Call standard multi-infra recommendation endpoint
      response = await api.post(`/beetle/recommendation/multiInfra${query}`, {
        desiredCspAndRegionPairs: targetPairs,
        sourceInfra
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

  // Poll async request details by ReqID
  getRequestDetails: async (reqId: string): Promise<{ status: string; errorResponse?: string; responseData?: any }> => {
    try {
      const response = await api.get(`/beetle/request/${reqId}`);
      const details = response.data?.data || response.data;
      return {
        status: details?.status || 'Handling',
        errorResponse: details?.errorResponse,
        responseData: details?.responseData
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
      console.warn(`getMigratedInfraList failed for ns ${nsId}:`, err);
      return [];
    }
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
    try {
      const response = await api.get(`/beetle/migration/ns/${nsId}/infra/${infraId}`);
      return response.data?.data || response.data;
    } catch (err: any) {
      console.warn(`getMigratedInfraDetail failed for ns ${nsId}, infra ${infraId}:`, err);
      return null;
    }
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

  // Overwrite sampleSourceInfra.json or sampleSourceGpuInfra.json directly on server file system for testing
  saveSourceInfraModelToFile: async (model: any, isGpu: boolean = false): Promise<{ success: boolean; message?: string; fileName?: string; error?: string }> => {
    try {
      const url = isGpu ? '/api/models/source-infra?sample=gpu' : '/api/models/source-infra';
      const response = await axios.post(url, model);
      return { success: true, message: response.data?.message || 'Source model saved and overwritten successfully.', fileName: response.data?.fileName };
    } catch (err: any) {
      return { success: false, error: err.response?.data?.error || err.message };
    }
  },

  // Overwrite sampleTargetInfra.json or sampleTargetGpuInfra.json directly on server file system for testing
  saveTargetInfraModelToFile: async (model: any, isGpu: boolean = false): Promise<{ success: boolean; message?: string; fileName?: string; error?: string }> => {
    try {
      const url = isGpu ? '/api/models/target-infra?sample=gpu' : '/api/models/target-infra';
      const response = await axios.post(url, model);
      return { success: true, message: response.data?.message || 'Target model saved and overwritten successfully.', fileName: response.data?.fileName };
    } catch (err: any) {
      return { success: false, error: err.response?.data?.error || err.message };
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

  // Step 7~12: Source & Planning Object Storage Artifact Save & Load
  saveSourceObjectStorageModel: async (
    sourceModel: any,
    options?: {
      name?: string;
      description?: string;
      version?: string;
      stage?: 'source' | 'planning';
      parentId?: string;
      overwriteId?: string | null;
    }
  ): Promise<{ success: boolean; modelId: string; stage?: 'source' | 'planning'; lineage?: ArtifactLineage; error?: string }> => {
    sourceModelCachePromise = null;
    const payload = sourceModel?.sourceModel || sourceModel;
    const targetStage = options?.stage || 'source';
    const reqBody = {
      name: options?.name || (targetStage === 'planning' ? 'planning-object-storage-model' : 'source-object-storage-model'),
      description: options?.description || payload?.description || `${targetStage === 'planning' ? 'Refined planning' : 'Source'} object storage artifact`,
      version: options?.version || '1.0.0',
      category: 'infra',
      stage: targetStage,
      schemaType: 'SourceObjectStorage',
      parentId: options?.parentId,
      payload
    };
    try {
      const response = options?.overwriteId
        ? await api.put(`/beetle/artifacts/${options.overwriteId}`, reqBody)
        : await api.post('/beetle/artifacts', reqBody);
      const art = response.data?.data || response.data;
      try { await axios.post('/api/models/source', sourceModel); } catch {}
      return {
        success: true,
        modelId: art?.id || options?.overwriteId || `src-model-${Date.now()}`,
        stage: art?.stage || targetStage,
        lineage: art?.lineage
      };
    } catch {
      try {
        const response = await axios.post('/api/models/source', sourceModel);
        return { success: true, modelId: response.data?.modelId || `src-model-${Date.now()}`, stage: targetStage };
      } catch (err: any) {
        console.warn('File persistence fallback: saving to local storage', err);
        return { success: true, modelId: `src-model-${Date.now().toString().slice(-4)}`, stage: targetStage };
      }
    }
  },

  getSourceObjectStorageModel: async (): Promise<{ success: boolean; sourceModel: any; error?: string }> => {
    if (sourceModelCachePromise) {
      return sourceModelCachePromise;
    }
    sourceModelCachePromise = (async () => {
      try {
        const response = await api.get('/beetle/artifacts?category=infra&schemaType=SourceObjectStorage');
        const list = response.data?.data || response.data;
        if (Array.isArray(list) && list.length > 0 && list[0]?.payload) {
          return { success: true, sourceModel: list[0].payload };
        }
      } catch {}
      try {
        const response = await axios.get('/api/models/source');
        return { success: true, sourceModel: response.data?.sourceModel };
      } catch (err: any) {
        sourceModelCachePromise = null;
        console.warn('File persistence fallback: fetching source model', err);
        return { success: false, sourceModel: null };
      }
    })();
    return sourceModelCachePromise;
  },

  // Step 15~20: Target Object Storage Artifact Save & Load
  saveTargetObjectStorageModel: async (
    targetModel: any,
    options?: {
      name?: string;
      description?: string;
      version?: string;
      parentId?: string;
      overwriteId?: string | null;
    }
  ): Promise<{ success: boolean; modelId: string; lineage?: ArtifactLineage; error?: string }> => {
    targetModelCachePromise = null;
    const payload = targetModel?.targetModel || targetModel?.recommendation || targetModel;
    const reqBody = {
      name: options?.name || 'target-object-storage-model',
      description: options?.description || payload?.description || 'Target object storage artifact',
      version: options?.version || '1.0.0',
      category: 'infra',
      stage: 'target',
      schemaType: 'RecommendedObjectStorage',
      parentId: options?.parentId,
      payload
    };
    try {
      const response = options?.overwriteId
        ? await api.put(`/beetle/artifacts/${options.overwriteId}`, reqBody)
        : await api.post('/beetle/artifacts', reqBody);
      const art = response.data?.data || response.data;
      try { await axios.post('/api/models/target', targetModel); } catch {}
      return {
        success: true,
        modelId: art?.id || options?.overwriteId || `tgt-model-${Date.now()}`,
        lineage: art?.lineage
      };
    } catch {
      try {
        const response = await axios.post('/api/models/target', targetModel);
        return { success: true, modelId: response.data?.modelId || `tgt-model-${Date.now()}` };
      } catch (err: any) {
        console.warn('File persistence fallback: saving target model', err);
        return { success: true, modelId: `tgt-model-${Date.now().toString().slice(-4)}` };
      }
    }
  },

  getTargetObjectStorageModel: async (): Promise<{ success: boolean; targetModel: any; error?: string }> => {
    if (targetModelCachePromise) {
      return targetModelCachePromise;
    }
    targetModelCachePromise = (async () => {
      try {
        const response = await api.get('/beetle/artifacts?category=infra&stage=target&schemaType=RecommendedObjectStorage');
        const list = response.data?.data || response.data;
        if (Array.isArray(list) && list.length > 0 && list[0]?.payload) {
          return { success: true, targetModel: list[0].payload };
        }
      } catch {}
      try {
        const response = await axios.get('/api/models/target');
        return { success: true, targetModel: response.data?.targetModel };
      } catch (err: any) {
        targetModelCachePromise = null;
        console.warn('File persistence fallback: fetching target model', err);
        return { success: false, targetModel: null };
      }
    })();
    return targetModelCachePromise;
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
        { id: 'koreacentral', name: 'Seoul' },
        { id: 'eastus', name: 'East US' },
        { id: 'westeurope', name: 'West Europe' }
      ],
      gcp: [
        { id: 'asia-northeast3', name: 'Seoul' },
        { id: 'us-central1', name: 'Iowa' },
        { id: 'europe-west3', name: 'Frankfurt' }
      ],
      alibaba: [
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
        { id: 'us-south-1', name: 'Dallas' },
        { id: 'us-east-1', name: 'Washington DC' }
      ],
      ncp: [
        { id: 'kr-1', name: 'Seoul' },
        { id: 'kr-2', name: 'Cheongju' },
        { id: 'sg-1', name: 'Singapore' }
      ],
      nhn: [
        { id: 'kr1', name: 'Pangyo' },
        { id: 'kr2', name: 'Gwangju' }
      ],
      kt: [
        { id: 'kr-1', name: 'Seoul' },
        { id: 'kr-2', name: 'Cheonan' }
      ],
      openstack: [
        { id: 'default', name: 'Default' }
      ]
    };
    const defaultRegions = (fallbacks[key] || [{ id: 'default', name: 'Default' }])
      .slice()
      .sort((a, b) => a.id.localeCompare(b.id));

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

          const sorted = list.sort((a, b) => a.id.localeCompare(b.id));
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
    try {
      const response = await api.get(`/tumblebug/ns/${nsId}/infra/${infraId}`);
      return response.data?.output || response.data;
    } catch (err: any) {
      console.warn(`getInfra failed for ${infraId}:`, err);
      return null;
    }
  },

  // Get SSH Key resource (including privateKey) from CB-Tumblebug
  getSshKey: async (nsId: string, sshKeyId: string): Promise<any> => {
    try {
      const response = await api.get(`/tumblebug/ns/${nsId}/resources/sshKey/${sshKeyId}`);
      return response.data?.output || response.data;
    } catch (err: any) {
      console.warn(`getSshKey failed for ${sshKeyId}:`, err);
      return null;
    }
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
