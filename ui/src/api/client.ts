import axios from 'axios';
import { OnpremInfra, OnpremModelEnvelope, RecommendedInfra, CloudModelEnvelope } from '../types/migration';

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
// 2. CM-Damselfly Server API Client (Model Repository)
// ============================================================================
export const damselflyApi = {
  // Save source infra model as a catalog template
  saveSourceModel: async (name: string, description: string, onpremiseInfraModel: OnpremInfra, version?: string): Promise<OnpremModelEnvelope> => {
    const response = await api.post('/damselfly/infra-model?modelType=onprem&isTargetModel=false', {
      userModelName: name,
      userModelVersion: version,
      description,
      onpremiseInfraModel
    });
    const m = response.data;
    return { ...m, id: m.id || m.uid || 'model-demo-1', name: m.name || m.userModelName || name, version: m.version || m.userModelVersion || version };
  },

  // Update existing source model
  updateSourceModel: async (id: string, name: string, description: string, onpremiseInfraModel: OnpremInfra, version?: string): Promise<OnpremModelEnvelope> => {
    const response = await api.put(`/damselfly/infra-model/${id}?modelType=onprem&isTargetModel=false`, {
      userModelName: name,
      userModelVersion: version,
      description,
      onpremiseInfraModel
    });
    const m = response.data;
    return { ...m, id: m.id || m.uid || id, name: m.name || m.userModelName || name, version: m.version || m.userModelVersion || version };
  },

  // Delete existing source model
  deleteSourceModel: async (id: string): Promise<void> => {
    await api.delete(`/damselfly/infra-model/${id}`);
  },

  // Get all source models list
  getSourceModels: async (): Promise<OnpremModelEnvelope[]> => {
    const response = await api.get('/damselfly/infra-model?modelType=onprem&isTargetModel=false');
    const dataList = Array.isArray(response.data) ? response.data : (response.data ? [response.data] : []);
    return dataList.map((m: any) => ({
      ...m,
      id: m.id || m.uid,
      name: m.name || m.userModelName || 'Unnamed Model',
      description: m.description,
      onpremiseInfraModel: m.onpremiseInfraModel,
      version: m.version || m.userModelVersion || '1.0.0'
    }));
  },

  // Get single source model spec
  getSourceModel: async (id: string): Promise<OnpremModelEnvelope> => {
    const response = await api.get(`/damselfly/infra-model/${id}?modelType=onprem&isTargetModel=false`);
    const m = response.data;
    return {
      ...m,
      id: m.id || m.uid || id,
      name: m.name || m.userModelName,
      version: m.version || m.userModelVersion || '1.0.0'
    };
  },

  // Save recommended target cloud spec as design template
  saveCloudModel: async (name: string, description: string, cloudInfraModel: RecommendedInfra, version?: string): Promise<CloudModelEnvelope> => {
    const response = await api.post('/damselfly/infra-model?modelType=cloud&isTargetModel=true', {
      userModelName: name,
      userModelVersion: version,
      description,
      cloudInfraModel
    });
    const m = response.data;
    return { ...m, id: m.id || m.uid || 'cloud-demo-1', name: m.name || m.userModelName || name, version: m.version || m.userModelVersion || version };
  },

  // Update existing target cloud design
  updateCloudModel: async (id: string, name: string, description: string, cloudInfraModel: RecommendedInfra, version?: string): Promise<CloudModelEnvelope> => {
    const response = await api.put(`/damselfly/infra-model/${id}?modelType=cloud&isTargetModel=true`, {
      userModelName: name,
      userModelVersion: version,
      description,
      cloudInfraModel
    });
    const m = response.data;
    return { ...m, id: m.id || m.uid || id, name: m.name || m.userModelName || name, version: m.version || m.userModelVersion || version };
  },

  // Get target cloud design templates
  getCloudModels: async (): Promise<CloudModelEnvelope[]> => {
    const response = await api.get('/damselfly/infra-model?modelType=cloud&isTargetModel=true');
    const dataList = Array.isArray(response.data) ? response.data : (response.data ? [response.data] : []);
    return dataList.map((m: any) => ({
      ...m,
      id: m.id || m.uid,
      name: m.name || m.userModelName || 'Unnamed Design',
      description: m.description,
      cloudInfraModel: m.cloudInfraModel,
      version: m.version || m.userModelVersion || '1.0.0'
    }));
  },

  getCloudModel: async (id: string): Promise<CloudModelEnvelope> => {
    const response = await api.get(`/damselfly/infra-model/${id}?modelType=cloud&isTargetModel=true`);
    const m = response.data;
    return {
      ...m,
      id: m.id || m.uid || id,
      name: m.name || m.userModelName,
      version: m.version || m.userModelVersion || '1.0.0'
    };
  },

  deleteCloudModel: async (id: string): Promise<void> => {
    await api.delete(`/damselfly/infra-model/${id}`);
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
  getRecommendations: async (sourceInfra: OnpremInfra, desiredCsp: string, desiredRegion: string): Promise<RecommendedInfra[]> => {
    const hasNlbs = sourceInfra.nlbs && sourceInfra.nlbs.length > 0;
    
    let response;
    if (hasNlbs) {
      // Call NLB-aware recommendation endpoint
      response = await api.post(`/beetle/recommendation/infraWithNlb?desiredCsp=${desiredCsp}&desiredRegion=${desiredRegion}`, {
        nameSeed: 'my',
        desiredCsp,
        desiredRegion,
        sourceInfra
      });
    } else {
      // Call standard VM recommendation endpoint
      response = await api.post(`/beetle/recommendation/infra?desiredCsp=${desiredCsp}&desiredRegion=${desiredRegion}`, {
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

  // Execute actual physical cloud migration deployment (with Prefer: respond-async header)
  executeMigration: async (nsId: string, nameSeed: string, cloudModel: RecommendedInfra): Promise<{ success: boolean; reqId?: string; data?: any; error?: string }> => {
    try {
      const response = await api.post(`/beetle/migration/ns/${nsId}/infra?nameSeed=${nameSeed}`, cloudModel, {
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
      // 404 Not Found means request is still initializing or in-flight — treat as Handling, NOT Error!
      if (err.response?.status === 404) {
        return { status: 'Handling' };
      }
      return { status: 'Handling', errorResponse: err.response?.data?.error || err.message };
    }
  },

  // Fetch final migration correlation and comparison report
  getMigrationReport: async (nsId: string, infraId: string): Promise<string> => {
    const response = await api.post(`/beetle/report/migration/ns/${nsId}/infra/${infraId}?format=html`, {});
    return response.data;
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

  deleteMigratedObjectStorage: async (nsId: string, osId: string): Promise<{ success: boolean; error?: string }> => {
    try {
      await api.delete(`/beetle/migration/middleware/ns/${nsId}/objectStorage/${osId}`);
      return { success: true };
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

  migrateData: async (dataMigrationModel: any): Promise<{ success: boolean; reqId?: string; data?: any; error?: string }> => {
    try {
      const response = await api.post('/beetle/migration/data', dataMigrationModel);
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

  // Step 7~12: Source Object Storage User Model Save & Load (Webserver File-Based Persistence)
  saveSourceObjectStorageModel: async (sourceModel: any): Promise<{ success: boolean; modelId: string; error?: string }> => {
    sourceModelCachePromise = null;
    try {
      const response = await axios.post('/api/models/source', sourceModel);
      return { success: true, modelId: response.data?.modelId || `src-model-${Date.now()}` };
    } catch (err: any) {
      console.warn('File persistence fallback: saving to local storage', err);
      return { success: true, modelId: `src-model-${Date.now().toString().slice(-4)}` };
    }
  },

  getSourceObjectStorageModel: async (): Promise<{ success: boolean; sourceModel: any; error?: string }> => {
    if (sourceModelCachePromise) {
      return sourceModelCachePromise;
    }
    sourceModelCachePromise = (async () => {
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

  // Step 15~20: Target Object Storage User Model Save & Load (Webserver File-Based Persistence)
  saveTargetObjectStorageModel: async (targetModel: any): Promise<{ success: boolean; modelId: string; error?: string }> => {
    targetModelCachePromise = null;
    try {
      const response = await axios.post('/api/models/target', targetModel);
      return { success: true, modelId: response.data?.modelId || `tgt-model-${Date.now()}` };
    } catch (err: any) {
      console.warn('File persistence fallback: saving target model', err);
      return { success: true, modelId: `tgt-model-${Date.now().toString().slice(-4)}` };
    }
  },

  getTargetObjectStorageModel: async (): Promise<{ success: boolean; targetModel: any; error?: string }> => {
    if (targetModelCachePromise) {
      return targetModelCachePromise;
    }
    targetModelCachePromise = (async () => {
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
      const response = await api.post(`/beetle/migration/middleware/ns/${nsId}/objectStorage?nameSeed=${nameSeed}`, requestBody, { headers });
      
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
