import { create, StateCreator } from 'zustand';
import { persist } from 'zustand/middleware';
import { OnpremInfra, OnpremModelEnvelope, RecommendedInfra, CloudModelEnvelope } from '../types/migration';
import { honeybeeApi, damselflyApi, beetleApi, tumblebugApi } from '../api/client';
import recommendedInfraSample from '../data/sampleTargetInfra.json';

// ----------------------------------------------------------------------------
// HIGH QUALITY DEMO / SOURCE INFRA MODEL DATA
// Cloned exactly from: cmd/test-cli/infra-with-nlb/testconf/recommendation-request.json
// ----------------------------------------------------------------------------
export const DEMO_SOURCE_INFRA: OnpremInfra = {
  network: {
    ipv4Networks: {
      defaultGateways: [
        {
          interfaceName: "ens5",
          ip: "10.0.1.1",
          machineId: "ec268ed7-821e-9d73-e79f-961262161624"
        },
        {
          interfaceName: "ens5",
          ip: "10.0.1.1",
          machineId: "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
        },
        {
          interfaceName: "ens5",
          ip: "10.0.1.1",
          machineId: "ec288dd0-c6fa-8a49-2f60-bc898311febf"
        }
      ],
      cidrBlocks: []
    },
    ipv6Networks: {}
  },
  nodes: [
    {
      cpu: {
        architecture: "x86_64",
        cores: 1,
        cpus: 1,
        maxSpeed: 2.499,
        model: "Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz",
        threads: 2,
        vendor: "GenuineIntel"
      },
      firewallTable: [
        { action: "allow", direction: "inbound", dstCIDR: "0.0.0.0/0", dstPorts: "22", protocol: "tcp", srcCIDR: "0.0.0.0/0", srcPorts: "*" },
        { action: "allow", direction: "inbound", dstCIDR: "0.0.0.0/0", dstPorts: "9999", protocol: "tcp", srcCIDR: "0.0.0.0/0", srcPorts: "*" },
        { action: "allow", direction: "inbound", dstCIDR: "0.0.0.0/0", dstPorts: "*", protocol: "*", srcCIDR: "10.0.0.0/16", srcPorts: "*" },
        { action: "allow", direction: "outbound", dstCIDR: "0.0.0.0/0", dstPorts: "*", protocol: "*", srcCIDR: "0.0.0.0/0", srcPorts: "*" }
      ],
      hostname: "ip-10-0-1-30",
      interfaces: [
        { name: "lo", state: "up", mtu: 65536, ipv4CidrBlocks: ["127.0.0.1/8"], ipv6CidrBlocks: ["::1/128"] },
        { name: "ens5", state: "up", macAddress: "02:6f:de:fc:71:b1", mtu: 9001, ipv4CidrBlocks: ["10.0.1.30/24"], ipv6CidrBlocks: ["fe80::6f:deff:fefc:71b1/64"] }
      ],
      machineId: "ec268ed7-821e-9d73-e79f-961262161624",
      memory: { available: 1, totalSize: 2, type: "DDR4" },
      os: { id: "ubuntu", idLike: "debian", name: "Ubuntu", prettyName: "Ubuntu 22.04.3 LTS", version: "22.04.3 LTS (Jammy Jellyfish)", versionCodename: "jammy", versionId: "22.04" },
      rootDisk: { label: "", totalSize: 8, type: "SSD" },
      routingTable: [
        { destination: "0.0.0.0/0", gateway: "10.0.1.1", interface: "ens5", linkState: "up", metric: 100, protocol: "kernel", scope: "universe" },
        { destination: "10.0.1.0/24", gateway: "10.0.1.1", interface: "ens5", linkState: "up", metric: 100, protocol: "kernel", scope: "universe" }
      ]
    },
    {
      cpu: {
        architecture: "x86_64",
        cores: 2,
        cpus: 1,
        maxSpeed: 2.499,
        model: "Intel(R) Xeon(R) Platinum 8175M CPU @ 2.50GHz",
        threads: 4,
        vendor: "GenuineIntel"
      },
      firewallTable: [
        { action: "allow", direction: "inbound", dstCIDR: "0.0.0.0/0", dstPorts: "22", protocol: "tcp", srcCIDR: "0.0.0.0/0", srcPorts: "*" },
        { action: "allow", direction: "inbound", dstCIDR: "0.0.0.0/0", dstPorts: "8086", protocol: "tcp", srcCIDR: "10.0.0.0/16", srcPorts: "*" },
        { action: "allow", direction: "inbound", dstCIDR: "0.0.0.0/0", dstPorts: "*", protocol: "*", srcCIDR: "10.0.0.0/16", srcPorts: "*" },
        { action: "allow", direction: "outbound", dstCIDR: "0.0.0.0/0", dstPorts: "*", protocol: "*", srcCIDR: "0.0.0.0/0", srcPorts: "*" }
      ],
      hostname: "ip-10-0-1-221",
      interfaces: [
        { name: "lo", state: "up", mtu: 65536, ipv4CidrBlocks: ["127.0.0.1/8"], ipv6CidrBlocks: ["::1/128"] },
        { name: "ens5", state: "up", macAddress: "02:08:96:7d:f4:17", mtu: 9001, ipv4CidrBlocks: ["10.0.1.221/24"], ipv6CidrBlocks: ["fe80::8:96ff:fe7d:f417/64"] }
      ],
      machineId: "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      memory: { available: 15, totalSize: 16, type: "DDR4" },
      os: { id: "ubuntu", idLike: "debian", name: "Ubuntu", prettyName: "Ubuntu 22.04.3 LTS", version: "22.04.3 LTS (Jammy Jellyfish)", versionCodename: "jammy", versionId: "22.04" },
      rootDisk: { label: "", totalSize: 30, type: "SSD" },
      routingTable: [
        { destination: "0.0.0.0/0", gateway: "10.0.1.1", interface: "ens5", linkState: "up", metric: 100, protocol: "kernel", scope: "universe" },
        { destination: "10.0.1.0/24", gateway: "10.0.1.1", interface: "ens5", linkState: "up", metric: 100, protocol: "kernel", scope: "universe" }
      ]
    },
    {
      cpu: {
        architecture: "x86_64",
        cores: 2,
        cpus: 1,
        maxSpeed: 2.499,
        model: "Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz",
        threads: 4,
        vendor: "GenuineIntel"
      },
      firewallTable: [
        { action: "allow", direction: "inbound", dstCIDR: "0.0.0.0/0", dstPorts: "22", protocol: "tcp", srcCIDR: "0.0.0.0/0", srcPorts: "*" },
        { action: "allow", direction: "inbound", dstCIDR: "0.0.0.0/0", dstPorts: "8086", protocol: "tcp", srcCIDR: "10.0.0.0/16", srcPorts: "*" },
        { action: "allow", direction: "inbound", dstCIDR: "0.0.0.0/0", dstPorts: "*", protocol: "*", srcCIDR: "10.0.0.0/16", srcPorts: "*" },
        { action: "allow", direction: "outbound", dstCIDR: "0.0.0.0/0", dstPorts: "*", protocol: "*", srcCIDR: "0.0.0.0/0", srcPorts: "*" }
      ],
      hostname: "ip-10-0-1-138",
      interfaces: [
        { name: "lo", state: "up", mtu: 65536, ipv4CidrBlocks: ["127.0.0.1/8"], ipv6CidrBlocks: ["::1/128"] },
        { name: "ens5", state: "up", macAddress: "02:bf:6e:6c:6e:31", mtu: 9001, ipv4CidrBlocks: ["10.0.1.138/24"], ipv6CidrBlocks: ["fe80::bf:6eff:fe6c:6e31/64"] }
      ],
      machineId: "ec288dd0-c6fa-8a49-2f60-bc898311febf",
      memory: { available: 7, totalSize: 8, type: "DDR4" },
      os: { id: "ubuntu", idLike: "debian", name: "Ubuntu", prettyName: "Ubuntu 22.04.3 LTS", version: "22.04.3 LTS (Jammy Jellyfish)", versionCodename: "jammy", versionId: "22.04" },
      rootDisk: { label: "", totalSize: 30, type: "SSD" },
      routingTable: [
        { destination: "0.0.0.0/0", gateway: "10.0.1.1", interface: "ens5", linkState: "up", metric: 100, protocol: "kernel", scope: "universe" },
        { destination: "10.0.1.0/24", gateway: "10.0.1.1", interface: "ens5", linkState: "up", metric: 100, protocol: "kernel", scope: "universe" }
      ]
    }
  ],
  nlbs: [
    {
      hostMachineId: "ec268ed7-821e-9d73-e79f-961262161624",
      software: "haproxy",
      listener: { bindAddress: "*", port: 9999, protocol: "tcp" },
      backend: {
        name: "influxdb_back",
        balance: "roundrobin",
        protocol: "tcp",
        servers: [
          { name: "influx1", ip: "10.0.1.221", port: 8086 },
          { name: "influx2", ip: "10.0.1.138", port: 8086 }
        ]
      },
      healthCheck: { enabled: true, interval: 10, timeout: 10, threshold: 3 }
    }
  ]
};

const DEFAULT_FALLBACK_SOURCE_MODEL: OnpremModelEnvelope = {
  id: 'sample-source-infra-1',
  name: '[Sample] web-haproxy-influxdb',
  description: '1 HAProxy/App node + 2 InfluxDB nodes with NLB (sample)',
  onpremiseInfraModel: DEMO_SOURCE_INFRA,
  version: '1.0',
  updatedTime: new Date().toISOString()
};

export type TabType = 'infra' | 'storage' | 'data' | 'credential' | 'overview' | 'source' | 'refine' | 'design' | 'migrate' | 'operations';

interface MigrationState {
  activeTab: TabType;
  setActiveTab: (tab: TabType) => void;
  themeMode: 'dark' | 'light';
  toggleTheme: () => void;

  // Page 1: Source Center State
  sourceGroups: any[];
  activeSgId: string;
  connections: any[];
  isLoadingSource: boolean;
  refinedSourceInfra: OnpremInfra | null;
  savedSourceModels: OnpremModelEnvelope[];
  selectedSourceModel: OnpremModelEnvelope | null;
  
  fetchSourceGroups: () => Promise<void>;
  createSourceGroup: (name: string, description: string) => Promise<string>;
  deleteSourceGroup: (sgId: string) => Promise<void>;
  refreshSourceGroup: (sgId: string) => Promise<void>;
  registerConnection: (serverData: any) => Promise<void>;
  fetchRefinedInfra: (connId: string) => Promise<void>;
  fetchRefinedInfraByGroup: (sgId: string) => Promise<void>;
  fetchSavedSourceModels: () => Promise<void>;
  selectSourceModel: (model: OnpremModelEnvelope | null) => void;
  saveSourceModel: (name: string, description: string, version: string, infra: OnpremInfra) => Promise<OnpremModelEnvelope>;
  updateSourceModel: (id: string, name: string, description: string, version: string, infra: OnpremInfra) => Promise<OnpremModelEnvelope>;
  deleteSourceModel: (id: string) => Promise<void>;

  // Page 2: Target Cloud Optimizer State
  desiredCsp: string;
  desiredRegion: string;
  recommendationCandidates: RecommendedInfra[];
  selectedCandidateIndex: number;
  editedCandidate: RecommendedInfra | null;
  savedCloudModels: CloudModelEnvelope[];
  selectedCloudModel: CloudModelEnvelope | null;
  isRecommending: boolean;
  tumblebugProviders: string[];
  tumblebugRegions: { id: string; name: string }[];
  
  setDesiredCsp: (csp: string) => Promise<void>;
  setDesiredRegion: (region: string) => void;
  triggerRecommendation: (sourceInfra: OnpremInfra) => Promise<void>;
  selectCandidate: (index: number) => void;
  updateEditedCandidate: (updated: RecommendedInfra) => void;
  fetchSavedCloudModels: () => Promise<void>;
  selectCloudModel: (model: CloudModelEnvelope | null) => void;
  saveCloudModel: (name: string, description: string, version: string, cloudInfra: RecommendedInfra) => Promise<CloudModelEnvelope>;
  updateCloudModel: (id: string, name: string, description: string, version: string, cloudInfra: RecommendedInfra) => Promise<CloudModelEnvelope>;
  deleteCloudModel: (id: string) => Promise<void>;
  fetchTumblebugProviders: () => Promise<void>;
  fetchTumblebugRegions: (providerName: string) => Promise<void>;

  // Page 3: Migration Execution State & Job History (Persisted across tab navigation)
  jobs: MigrationJob[];
  activeJobId: string;
  namespaceId: string;
  nameSeed: string;
  isDeploying: boolean;
  activeDeploymentId: string;
  deploymentStatus: any | null;
  liveReportHtml: string;
  
  // Page 5: Infrastructure Deletion Request Tracking (Persisted across tab navigation)
  deletingInfrasMap: Record<string, { reqId: string; status: string }>;
  startInfraTeardown: (infraId: string, reqId: string) => void;
  removeInfraTeardown: (infraId: string) => void;
  pollInfraTeardownStatus: (nsId: string, infraId: string, reqId: string) => Promise<{ completed: boolean; success: boolean; error?: string }>;

  setJobs: (jobs: MigrationJob[] | ((prev: MigrationJob[]) => MigrationJob[])) => void;
  setActiveJobId: (id: string) => void;
  addJob: (job: MigrationJob) => void;
  removeJob: (id: string) => void;
  setNamespaceId: (nsId: string) => void;
  setNameSeed: (seed: string) => void;
  startMigration: (cloudModel: RecommendedInfra) => Promise<void>;
  fetchDeploymentStatus: () => Promise<void>;
  fetchMigrationReport: () => Promise<void>;
}

export interface MigrationJob {
  id: string;
  reqId: string;
  infraId: string;
  nsId: string;
  nameSeed: string;
  csp: string;
  region: string;
  status: 'Handling' | 'Success' | 'Failed';
  startTime: string;
  elapsedSeconds: number;
  nodeGroupsCount: number;
  totalVms: number;
  logs: string[];
  vms?: { publicIp: string; privateIp: string; specId: string; name: string }[];
  error?: string;
  isSample?: boolean;
}

const storeInitializer: StateCreator<MigrationState> = (set, get) => ({
  activeTab: 'source',
  setActiveTab: (tab) => set({ activeTab: tab }),
  themeMode: 'light',
  toggleTheme: () => {
    const nextTheme = get().themeMode === 'dark' ? 'light' : 'dark';
    set({ themeMode: nextTheme });
    const root = window.document.documentElement;
    if (nextTheme === 'light') {
      root.classList.add('light');
      root.classList.remove('dark');
    } else {
      root.classList.add('dark');
      root.classList.remove('light');
    }
  },

  // --------------------------------------------------------------------------
  // Page 1: Source Infrastructures - Offline demo data aligned with Honeybee API schema
  // --------------------------------------------------------------------------
  sourceGroups: [],
  activeSgId: '',
  connections: [],
  isLoadingSource: false,
  refinedSourceInfra: DEMO_SOURCE_INFRA,
  savedSourceModels: [DEFAULT_FALLBACK_SOURCE_MODEL],
  selectedSourceModel: DEFAULT_FALLBACK_SOURCE_MODEL,
  tumblebugProviders: [
    'aws', 'azure', 'gcp', 'alibaba', 'tencent',
    'ibm', 'ncp', 'nhn', 'kt', 'openstack'
  ],
  tumblebugRegions: [
    { id: 'ap-northeast-2', name: 'Seoul' },
    { id: 'us-east-1', name: 'N. Virginia' },
    { id: 'us-west-2', name: 'Oregon' },
    { id: 'eu-west-1', name: 'Ireland' }
  ],

  fetchSourceGroups: async () => {
    try {
      const response = await honeybeeApi.getSourceGroups();
      // ListSourceGroupRes: { source_group: [], connection_info_status_count: {} }
      const list = response?.source_group || (Array.isArray(response) ? response : []);
      if (list.length > 0) {
        // Real data from Honeybee — clear any stale selection
        set({ sourceGroups: list, activeSgId: '', connections: [] });
        return;
      }
      // Honeybee is up but has no groups yet — show empty state
      set({ sourceGroups: [], activeSgId: '', connections: [] });
    } catch {
      // Honeybee unreachable — stay with current state (no crash)
      set({ sourceGroups: [], activeSgId: '', connections: [] });
    }
  },

  createSourceGroup: async (name, description) => {
    try {
      const result = await honeybeeApi.createSourceGroup(name, description);
      const newId = result?.id || `sg-${Date.now()}`;
      const newGroup = { id: newId, name, description, connection_info_status_count: { connection_info_total: 0, count_connection_success: 0, count_connection_failed: 0, count_agent_success: 0, count_agent_failed: 0 } };
      set({ sourceGroups: [...get().sourceGroups, newGroup], activeSgId: newId, connections: [] });
      return newId;
    } catch {
      const demoId = `sg-demo-${Date.now()}`;
      const newGroup = { id: demoId, name, description, connection_info_status_count: { connection_info_total: 0, count_connection_success: 0, count_connection_failed: 0, count_agent_success: 0, count_agent_failed: 0 } };
      set({ sourceGroups: [...get().sourceGroups, newGroup], activeSgId: demoId, connections: [] });
      return demoId;
    }
  },

  deleteSourceGroup: async (sgId) => {
    try {
      await honeybeeApi.deleteSourceGroup(sgId);
    } catch {
      // proceed with local removal even if API call fails
    }
    const updated = get().sourceGroups.filter((g: any) => g.id !== sgId);
    const nextActiveId = get().activeSgId === sgId ? (updated[0]?.id || '') : get().activeSgId;
    set({ sourceGroups: updated, activeSgId: nextActiveId, connections: nextActiveId ? get().connections : [] });
  },

  refreshSourceGroup: async (sgId) => {
    try {
      await honeybeeApi.refreshSourceGroup(sgId);
    } catch {
      // noop in offline mode
    }
  },

  registerConnection: async (serverData) => {
    const sgId = get().activeSgId;
    try {
      const result = await honeybeeApi.registerConnectionInfo(sgId, serverData);
      const newConn = result || {
        id: `conn-${Date.now()}`,
        name: serverData.name,
        ip_address: serverData.ip,
        ssh_port: String(serverData.port),
        user: serverData.user,
        connection_status: '',
        agent_status: '',
      };
      set({ connections: [...get().connections, newConn] });
    } catch {
      const newConn = {
        id: `conn-demo-${Date.now()}`,
        name: serverData.name,
        ip_address: serverData.ip,
        ssh_port: String(serverData.port),
        user: serverData.user,
        connection_status: '',
        agent_status: '',
      };
      set({ connections: [...get().connections, newConn] });
    }
  },

  fetchRefinedInfra: async (connId: string) => {
    const activeSgId = get().activeSgId;
    if (!activeSgId) return;
    try {
      const data = await honeybeeApi.getRefinedInfraByConnection(activeSgId, connId);
      set({ refinedSourceInfra: data });
    } catch (err) {
      console.warn('Failed to fetch refined infra from Honeybee by connection, falling back to demo:', err);
      set({ refinedSourceInfra: DEMO_SOURCE_INFRA });
    }
  },

  fetchRefinedInfraByGroup: async (sgId: string) => {
    try {
      const data = await honeybeeApi.getRefinedInfraByGroup(sgId);
      set({ refinedSourceInfra: data });
    } catch (err) {
      console.warn('Failed to fetch refined infra from Honeybee by group:', err);
      // For real groups, set to null if fetching fails (i.e. not yet collected or error)
      set({ refinedSourceInfra: null });
    }
  },

  fetchSavedSourceModels: async () => {
    try {
      const models = await damselflyApi.getSourceModels();
      // Always include the built-in sample model at the top of the list
      const withoutSample = models.filter((m: OnpremModelEnvelope) => m.id !== DEFAULT_FALLBACK_SOURCE_MODEL.id);
      set({ savedSourceModels: [DEFAULT_FALLBACK_SOURCE_MODEL, ...withoutSample] });
    } catch {
      // Damselfly unreachable — show sample model only
      set({ savedSourceModels: [DEFAULT_FALLBACK_SOURCE_MODEL] });
    }
  },

  selectSourceModel: (model) => {
    set({ 
      selectedSourceModel: model,
      refinedSourceInfra: model ? model.onpremiseInfraModel : null
    });
  },

  saveSourceModel: async (name, description, version, infra) => {
    const saved = await damselflyApi.saveSourceModel(name, description, infra, version);
    set({
      savedSourceModels: [...get().savedSourceModels, saved],
      selectedSourceModel: saved
    });
    return saved;
  },

  updateSourceModel: async (id, name, description, version, updatedInfra) => {
    const saved = await damselflyApi.updateSourceModel(id, name, description, updatedInfra, version);
    const updatedList = get().savedSourceModels.map(m => m.id === id ? saved : m);
    set({ savedSourceModels: updatedList, selectedSourceModel: saved });
    return saved;
  },

  deleteSourceModel: async (id) => {
    await damselflyApi.deleteSourceModel(id);
    const updatedList = get().savedSourceModels.filter(m => m.id !== id);
    set({
      savedSourceModels: updatedList,
      selectedSourceModel: get().selectedSourceModel?.id === id ? null : get().selectedSourceModel
    });
  },

  // --------------------------------------------------------------------------
  // Page 2: Target Cloud Optimizer - Live Connected to Beetle API
  // --------------------------------------------------------------------------
  desiredCsp: 'aws',
  desiredRegion: 'ap-northeast-2',
  recommendationCandidates: [],
  selectedCandidateIndex: 0,
  editedCandidate: null,
  savedCloudModels: [],
  selectedCloudModel: null,
  isRecommending: false,

  setDesiredCsp: async (csp) => {
    set({ desiredCsp: csp });
    try {
      const regions = await tumblebugApi.getRegions(csp);
      set({ tumblebugRegions: regions });
      if (regions.length > 0) {
        set({ desiredRegion: regions[0].id });
      }
    } catch (err) {
      console.error('Failed to sync regions on CSP change', err);
    }
  },
  setDesiredRegion: (region) => set({ desiredRegion: region }),

  triggerRecommendation: async (sourceInfra) => {
    set({ isRecommending: true, recommendationCandidates: [], editedCandidate: null });
    try {
      // Direct live REST API request to Beetle with fully qualified JSON mapping
      const candidates = await beetleApi.getRecommendations(sourceInfra, get().desiredCsp, get().desiredRegion);
      set({ 
        recommendationCandidates: candidates, 
        selectedCandidateIndex: 0,
        editedCandidate: candidates.length > 0 ? candidates[0] : null,
        isRecommending: false 
      });
    } catch (err) {
      console.error('Beetle Recommendation API Error:', err);
      set({ isRecommending: false });
      alert('Failed to fetch recommendations from Beetle. Check server logs.');
    }
  },

  selectCandidate: (index) => {
    const list = get().recommendationCandidates;
    if (index >= 0 && index < list.length) {
      set({ 
        selectedCandidateIndex: index,
        editedCandidate: list[index]
      });
    }
  },

  updateEditedCandidate: (updated) => {
    set({ editedCandidate: updated });
    const list = [...get().recommendationCandidates];
    const index = get().selectedCandidateIndex;
    if (index >= 0 && index < list.length) {
      list[index] = updated;
      set({ recommendationCandidates: list });
    }
  },

  fetchSavedCloudModels: async () => {
    // Fallback demo model shown when Damselfly is unreachable, to prevent an empty select box.
    const fallbackCloudModels: CloudModelEnvelope[] = [
      {
        id: 'cloud-demo-1',
        name: '[Sample] cloud-target-v1',
        description: 'Optimized Cloud architecture generated for onpremise cluster.',
        cloudInfraModel: recommendedInfraSample.data[0] as any,
        version: '1.0',
        updatedTime: new Date().toISOString()
      }
    ];
    try {
      const models = await damselflyApi.getCloudModels();
      set({ savedCloudModels: [fallbackCloudModels[0], ...models] });
    } catch {
      // Damselfly unreachable — show fallback demo model only
      set({ savedCloudModels: fallbackCloudModels });
    }
  },

  selectCloudModel: (model) => {
    set({ 
      selectedCloudModel: model,
      editedCandidate: model ? model.cloudInfraModel : null,
      recommendationCandidates: model ? [model.cloudInfraModel] : []
    });
  },

  saveCloudModel: async (name, description, version, cloudInfra) => {
    const saved = await damselflyApi.saveCloudModel(name, description, cloudInfra, version);
    set({
      savedCloudModels: [...get().savedCloudModels, saved],
      selectedCloudModel: saved
    });
    return saved;
  },

  updateCloudModel: async (id, name, description, version, cloudInfra) => {
    const saved = await damselflyApi.updateCloudModel(id, name, description, cloudInfra, version);
    const updatedList = get().savedCloudModels.map(m => m.id === id ? saved : m);
    set({ savedCloudModels: updatedList, selectedCloudModel: saved });
    return saved;
  },

  deleteCloudModel: async (id) => {
    try {
      await damselflyApi.deleteCloudModel(id);
      const updated = get().savedCloudModels.filter(m => m.id !== id);
      set({
        savedCloudModels: updated,
        selectedCloudModel: get().selectedCloudModel?.id === id ? null : get().selectedCloudModel,
        editedCandidate: get().selectedCloudModel?.id === id ? null : get().editedCandidate,
        recommendationCandidates: get().selectedCloudModel?.id === id ? [] : get().recommendationCandidates
      });
    } catch (err) {
      console.error(err);
      throw err;
    }
  },

  fetchTumblebugProviders: async () => {
    const rawList = await tumblebugApi.getProviders();
    const list = rawList.filter((p: string) => p.toLowerCase() !== 'openstack-ex01');
    set({ tumblebugProviders: list });
    // Dynamically pull regions for the currently set desiredCsp or the first returned provider
    const activeCsp = get().desiredCsp || (list.length > 0 ? list[0] : 'aws');
    const regions = await tumblebugApi.getRegions(activeCsp);
    set({ tumblebugRegions: regions });
    
    // Sync desiredRegion if the current one is not in the loaded regions list
    if (regions.length > 0) {
      const exists = regions.some(r => r.id === get().desiredRegion);
      if (!exists) {
        set({ desiredRegion: regions[0].id });
      }
    }
  },

  fetchTumblebugRegions: async (providerName: string) => {
    const regions = await tumblebugApi.getRegions(providerName);
    set({ tumblebugRegions: regions });
    if (regions.length > 0) {
      const exists = regions.some(r => r.id === get().desiredRegion);
      if (!exists) {
        set({ desiredRegion: regions[0].id });
      }
    }
  },

  // --------------------------------------------------------------------------
  // Page 3: Migration Execution - Live Connected to Beetle API
  // --------------------------------------------------------------------------
  jobs: [
    {
      id: 'req-aws-01',
      reqId: 'req-20260721-001',
      infraId: 'mig01-aws-infra',
      nsId: 'mig01',
      nameSeed: '',
      csp: 'AWS',
      region: 'ap-northeast-2',
      status: 'Handling',
      startTime: '19:25:10',
      elapsedSeconds: 15,
      nodeGroupsCount: 2,
      totalVms: 3,
      isSample: true,
      logs: [
        'POST /beetle/migration/ns/mig01/infra?nameSeed=',
        'HTTP 202 Accepted (ReqID: req-20260721-001, Status: Handling)',
        'GET /beetle/request/req-20260721-001 -> Status: Handling (Elapsed: 15s)'
      ]
    }
  ],
  activeJobId: 'req-aws-01',
  namespaceId: 'mig01',
  nameSeed: '',
  isDeploying: false,
  activeDeploymentId: '',
  deploymentStatus: null,
  liveReportHtml: '',

  setJobs: (fn) => set((state) => ({ jobs: typeof fn === 'function' ? fn(state.jobs) : fn })),
  setActiveJobId: (id) => set({ activeJobId: id }),
  addJob: (job) => set((state) => ({ jobs: [job, ...state.jobs], activeJobId: job.id })),
  removeJob: (id) => set((state) => {
    const updatedJobs = state.jobs.filter(j => j.id !== id);
    const nextActiveId = state.activeJobId === id ? (updatedJobs[0]?.id || '') : state.activeJobId;
    return { jobs: updatedJobs, activeJobId: nextActiveId };
  }),
  setNamespaceId: (nsId) => set({ namespaceId: nsId }),
  setNameSeed: (seed) => set({ nameSeed: seed }),

  startMigration: async (cloudModel) => {
    set({ isDeploying: true, deploymentStatus: null, liveReportHtml: '' });
    const nsId = get().namespaceId;
    const nameSeed = get().nameSeed;
    
    try {
      // Direct live REST API request to Beetle for physical provisioning
      const res = await beetleApi.executeMigration(nsId, nameSeed, cloudModel);
      if (res.success && res.data) {
        const infraId = res.data.name || cloudModel.targetInfra.name;
        set({ activeDeploymentId: infraId });
        await get().fetchDeploymentStatus();
      } else {
        set({ isDeploying: false, deploymentStatus: { status: 'Failed', error: res.error } });
      }
    } catch (err: any) {
      console.error('Beetle Execution API Error:', err);
      set({ isDeploying: false });
      alert('Failed to launch deployment on Beetle. Check server logs.');
    }
  },

  fetchDeploymentStatus: async () => {
    const reqId = get().activeDeploymentId;
    if (!reqId) return;

    try {
      const statusData = await beetleApi.getRequestDetails(reqId);
      set({ deploymentStatus: statusData });
      
      if (statusData && (statusData.status === 'Success' || statusData.status === 'Error')) {
        set({ isDeploying: false });
        await get().fetchMigrationReport();
      }
    } catch (err) {
      console.error('Failed to poll deployment status:', err);
    }
  },

  fetchMigrationReport: async () => {
    const nsId = get().namespaceId;
    const infraId = get().activeDeploymentId;
    if (!nsId || !infraId) return;

    try {
      const reportHtml = await beetleApi.getMigrationReport(nsId, infraId);
      set({ liveReportHtml: reportHtml });
    } catch (err) {
      console.error('Failed to load migration report:', err);
    }
  },

  // Page 5: Infrastructure Deletion Request Tracking
  deletingInfrasMap: {},
  startInfraTeardown: (infraId: string, reqId: string) => {
    set((state) => ({
      deletingInfrasMap: {
        ...state.deletingInfrasMap,
        [infraId]: { reqId, status: 'Terminating' }
      }
    }));
  },
  removeInfraTeardown: (infraId: string) => {
    set((state) => {
      const updated = { ...state.deletingInfrasMap };
      delete updated[infraId];
      return { deletingInfrasMap: updated };
    });
  },
  pollInfraTeardownStatus: async (nsId: string, infraId: string, reqId: string) => {
    try {
      const reqDetails = await beetleApi.getRequestDetails(reqId);
      const reqStatus = reqDetails.status;
      if (reqStatus === 'Completed' || reqStatus === 'Succeeded' || reqStatus === 'Success') {
        get().removeInfraTeardown(infraId);
        return { completed: true, success: true };
      } else if (reqStatus === 'Failed' || reqStatus === 'Error') {
        get().removeInfraTeardown(infraId);
        return { completed: true, success: false, error: reqDetails.errorResponse || 'Teardown failed' };
      }
      return { completed: false, success: true };
    } catch (err: any) {
      console.warn('Poll error for teardown request', reqId, err);
      return { completed: false, success: false };
    }
  }
});

export const useMigrationStore = create<MigrationState>()(
  persist(storeInitializer, {
    name: 'cm-beetle-migration-store-v1',
    partialize: (state) => ({
      jobs: state.jobs,
      activeJobId: state.activeJobId,
      savedCloudModels: state.savedCloudModels,
      selectedCloudModel: state.selectedCloudModel,
      savedSourceModels: state.savedSourceModels,
      selectedSourceModel: state.selectedSourceModel,
      namespaceId: state.namespaceId,
      nameSeed: state.nameSeed,
      themeMode: state.themeMode,
      deletingInfrasMap: state.deletingInfrasMap
    })
  })
);
