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
    return response.data;
  },

  getRefinedInfraByGroup: async (sgId: string): Promise<OnpremInfra> => {
    const response = await api.get(`/honeybee/source_group/${sgId}/infra/refined`);
    return response.data;
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
  saveSourceModel: async (name: string, description: string, onpremiseInfraModel: OnpremInfra): Promise<OnpremModelEnvelope> => {
    const response = await api.post('/damselfly/infra-model?modelType=onprem&isTargetModel=false', {
      userModelName: name,
      description,
      onpremiseInfraModel
    });
    const m = response.data;
    return { ...m, id: m.id || m.uid || 'model-demo-1', name: m.name || m.userModelName || name };
  },

  // Update existing source model
  updateSourceModel: async (id: string, name: string, description: string, onpremiseInfraModel: OnpremInfra): Promise<OnpremModelEnvelope> => {
    const response = await api.put(`/damselfly/infra-model/${id}?modelType=onprem&isTargetModel=false`, {
      userModelName: name,
      description,
      onpremiseInfraModel
    });
    const m = response.data;
    return { ...m, id: m.id || m.uid || id, name: m.name || m.userModelName || name };
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
      onpremiseInfraModel: m.onpremiseInfraModel
    }));
  },

  // Get single source model spec
  getSourceModel: async (id: string): Promise<OnpremModelEnvelope> => {
    const response = await api.get(`/damselfly/infra-model/${id}?modelType=onprem&isTargetModel=false`);
    const m = response.data;
    return { ...m, id: m.id || m.uid || id, name: m.name || m.userModelName };
  },

  // Save recommended target cloud spec as design template
  saveCloudModel: async (name: string, description: string, cloudInfraModel: RecommendedInfra): Promise<CloudModelEnvelope> => {
    const response = await api.post('/damselfly/infra-model?modelType=cloud&isTargetModel=true', {
      userModelName: name,
      description,
      cloudInfraModel
    });
    const m = response.data;
    return { ...m, id: m.id || m.uid || 'cloud-demo-1', name: m.name || m.userModelName || name };
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
      cloudInfraModel: m.cloudInfraModel
    }));
  },

  getCloudModel: async (id: string): Promise<CloudModelEnvelope> => {
    const response = await api.get(`/damselfly/infra-model/${id}?modelType=cloud&isTargetModel=true`);
    const m = response.data;
    return { ...m, id: m.id || m.uid || id, name: m.name || m.userModelName };
  }
};

// ============================================================================
// 3. CM-Beetle Server API Client (Recommendation & Migration Engine)
// ============================================================================
export const beetleApi = {
  // Get Cloud Recommendation candidates based on Source model input
  getRecommendations: async (sourceInfra: OnpremInfra, desiredCsp: string, desiredRegion: string): Promise<RecommendedInfra[]> => {
    // Unconditionally call /recommendation/infraWithNlb as requested
    const endpoint = '/beetle/recommendation/infraWithNlb';
      
    const response = await api.post(`${endpoint}?desiredCsp=${desiredCsp}&desiredRegion=${desiredRegion}`, {
      nameSeed: 'my',
      desiredCsp,
      desiredRegion,
      sourceInfra
    });
    
    // Beetle recommendation returns wrapping object: { success: true, data: RecommendedInfra[] }
    if (response.data && Array.isArray(response.data.data)) {
      return response.data.data;
    }
    return Array.isArray(response.data) ? response.data : [];
  },

  // Execute actual physical cloud migration deployment
  executeMigration: async (nsId: string, nameSeed: string, cloudModel: RecommendedInfra): Promise<{ success: boolean; data?: any; error?: string }> => {
    try {
      const response = await api.post(`/beetle/migration/ns/${nsId}/infra?nameSeed=${nameSeed}`, cloudModel);
      return { success: true, data: response.data };
    } catch (err: any) {
      return { success: false, error: err.response?.data?.error || err.message };
    }
  },

  // Poll migration setup status
  getMigrationStatus: async (nsId: string, infraId: string): Promise<any> => {
    const response = await api.get(`/beetle/migration/ns/${nsId}/infra/${infraId}`);
    return response.data;
  },

  // Fetch final migration correlation and comparison report
  getMigrationReport: async (nsId: string, infraId: string): Promise<string> => {
    const response = await api.post(`/beetle/report/migration/ns/${nsId}/infra/${infraId}?format=html`, {});
    return response.data;
  }
};

// ============================================================================
// 4. CB-Tumblebug Server API Client (Multi-Cloud Resource Management)
// ============================================================================
export const tumblebugApi = {
  // Get all active cloud providers in Tumblebug ordered by market share
  getProviders: async (): Promise<string[]> => {
    const defaultCsps = [
      'aws', 'azure', 'gcp', 'alibaba', 'tencent',
      'ibm', 'ncp', 'nhn', 'kt', 'openstack'
    ];
    try {
      const response = await api.get('/tumblebug/provider');
      // Tumblebug returns provider list inside "output" field!
      const data = response.data?.output || response.data?.provider || response.data;
      if (Array.isArray(data) && data.length > 0) {
        const apiCsps = data.map((c: any) => typeof c === 'string' ? c.toLowerCase() : (c.id || c.name || '').toLowerCase()).filter(Boolean);
        // Sort provider array by the ordered defaultCsps array index
        return Array.from(new Set([...defaultCsps, ...apiCsps]))
          .filter(Boolean)
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
  },

  // Get regions for a selected cloud provider (use fallback regions only if API fails)
  getRegions: async (providerName: string): Promise<{ id: string; name: string }[]> => {
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
    const defaultRegions = fallbacks[providerName.toLowerCase()] || [{ id: 'default', name: 'Default' }];

    try {
      const response = await api.get(`/tumblebug/provider/${providerName}/region`);
      // Tumblebug returns region list inside "regions" field!
      const data = response.data?.regions || response.data?.region || response.data;
      if (Array.isArray(data) && data.length > 0) {
        return data.map((r: any) => {
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
      }
      return defaultRegions;
    } catch (err) {
      console.warn(`Tumblebug getRegions for ${providerName} failed, using default regions`, err);
      return defaultRegions;
    }
  }
};
