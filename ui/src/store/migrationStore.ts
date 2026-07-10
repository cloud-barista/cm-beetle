import { create } from 'zustand';
import { OnpremInfra, OnpremModelEnvelope, RecommendedInfra, CloudModelEnvelope } from '../types/migration';
import { honeybeeApi, damselflyApi, beetleApi, tumblebugApi } from '../api/client';

// ----------------------------------------------------------------------------
// HIGH QUALITY DEMO / SOURCE INFRA MODEL DATA
// Cloned exactly from: cmd/test-cli/infra-with-nlb/testconf/recommendation-request.json
// ----------------------------------------------------------------------------
const DEMO_SOURCE_INFRA: OnpremInfra = {
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
  id: 'model-demo-1',
  name: 'onpremise-web-db-v1',
  description: 'Production server cluster containing 1 Web and 2 InfluxDB DB nodes',
  onpremiseInfraModel: DEMO_SOURCE_INFRA,
  version: '1.0',
  updatedTime: new Date().toISOString()
};

interface MigrationState {
  activeTab: 'source' | 'design' | 'migrate';
  setActiveTab: (tab: 'source' | 'design' | 'migrate') => void;
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
  saveSourceModel: (name: string, description: string) => Promise<void>;
  updateSourceModel: (id: string, name: string, description: string, updatedInfra: OnpremInfra) => Promise<void>;

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
  saveCloudModel: (name: string, description: string) => Promise<void>;
  fetchTumblebugProviders: () => Promise<void>;
  fetchTumblebugRegions: (providerName: string) => Promise<void>;

  // Page 3: Migration Execution State
  namespaceId: string;
  nameSeed: string;
  isDeploying: boolean;
  activeDeploymentId: string;
  deploymentStatus: any | null;
  liveReportHtml: string;
  
  setNamespaceId: (nsId: string) => void;
  setNameSeed: (seed: string) => void;
  startMigration: (cloudModel: RecommendedInfra) => Promise<void>;
  fetchDeploymentStatus: () => Promise<void>;
  fetchMigrationReport: () => Promise<void>;
}

export const useMigrationStore = create<MigrationState>((set, get) => ({
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
    { id: 'ap-northeast-2', name: 'Asia Pacific (Seoul)' },
    { id: 'us-east-1', name: 'US East (N. Virginia)' },
    { id: 'us-west-2', name: 'US West (Oregon)' },
    { id: 'eu-west-1', name: 'Europe (Ireland)' }
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

  fetchRefinedInfra: async () => {
    set({ refinedSourceInfra: DEMO_SOURCE_INFRA });
  },

  fetchRefinedInfraByGroup: async () => {
    set({ refinedSourceInfra: DEMO_SOURCE_INFRA });
  },

  fetchSavedSourceModels: async () => {
    // Damselfly disconnected. Pure local offline flow to prevent empty select box.
    set({ savedSourceModels: [DEFAULT_FALLBACK_SOURCE_MODEL] });
  },

  selectSourceModel: (model) => {
    set({ 
      selectedSourceModel: model,
      refinedSourceInfra: model ? model.onpremiseInfraModel : null
    });
  },

  saveSourceModel: async (name, description) => {
    const infra = get().refinedSourceInfra || DEMO_SOURCE_INFRA;
    const demoModel: OnpremModelEnvelope = {
      id: `model-demo-${Date.now()}`,
      name,
      description,
      onpremiseInfraModel: infra,
      version: '1.0',
      updatedTime: new Date().toISOString()
    };
    set({ savedSourceModels: [...get().savedSourceModels, demoModel] });
  },

  updateSourceModel: async (id, name, description, updatedInfra) => {
    const updatedList = get().savedSourceModels.map(m => m.id === id ? { ...m, name, description, onpremiseInfraModel: updatedInfra } : m);
    set({ savedSourceModels: updatedList });
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
    // Damselfly disconnected. Pure local offline flow to prevent empty select box.
    const fallbackCloudModels: CloudModelEnvelope[] = [
      {
        id: 'cloud-demo-1',
        name: 'cloud-target-v1',
        description: 'Optimized Cloud architecture generated for onpremise cluster.',
        cloudInfraModel: {
          status: 'highly-matched',
          description: 'Optimal spec scenario matching source cores & memory properties exactly.',
          targetCloud: { csp: 'aws', region: 'ap-northeast-2' },
          targetInfra: {
            name: 'optimal-infra',
            description: 'AWS recommended compute spec infrastructure',
            nodeGroups: [
              { name: 'web-group', nodeGroupSize: 1, label: { sourceMachineIds: 'ec268ed7-821e-9d73-e79f-961262161624' }, description: 'Web node spec mapping', connectionName: 'conn-aws', specId: 't3.small', imageId: 'ami-ubuntu-jammy-22.04', vNetId: 'vnet-demo', subnetId: 'subnet-1', securityGroupIds: ['sg-web'], sshKeyId: 'demo-key', rootDiskSize: 10 },
              { name: 'db-group', nodeGroupSize: 2, label: { sourceMachineIds: 'ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf' }, description: 'InfluxDB nodes mapping', connectionName: 'conn-aws', specId: 't3.large', imageId: 'ami-ubuntu-jammy-22.04', vNetId: 'vnet-demo', subnetId: 'subnet-1', securityGroupIds: ['sg-db'], sshKeyId: 'demo-key', rootDiskSize: 30 }
            ]
          },
          targetVNet: {
            name: 'vnet-demo',
            connectionName: 'conn-aws',
            cidrBlock: '10.0.0.0/16',
            subnetInfoList: [{ name: 'subnet-1', ipv4_CIDR: '10.0.1.0/24', description: 'Primary Subnet Block' }],
            description: 'Primary Virtual Network'
          },
          targetSshKey: { name: 'demo-key', connectionName: 'conn-aws', description: 'Admin cluster SSH key' },
          targetSecurityGroupList: [
            { name: 'sg-web', connectionName: 'conn-aws', vNetId: 'vnet-demo', description: 'Security rules for web access', firewallRules: [{ action: 'allow', direction: 'inbound', dstCIDR: '0.0.0.0/0', dstPorts: '22,9999', protocol: 'tcp', srcCIDR: '0.0.0.0/0', srcPorts: '*' }] }
          ]
        },
        version: '1.0',
        updatedTime: new Date().toISOString()
      }
    ];
    set({ savedCloudModels: fallbackCloudModels });
  },

  selectCloudModel: (model) => {
    set({ 
      selectedCloudModel: model,
      editedCandidate: model ? model.cloudInfraModel : null
    });
  },

  saveCloudModel: async (name, description) => {
    const candidate = get().editedCandidate;
    if (!candidate) throw new Error('No edited cloud configuration to save');
    const demoModel: CloudModelEnvelope = {
      id: `cloud-demo-${Date.now()}`,
      name,
      description,
      cloudInfraModel: candidate,
      version: '1.0',
      updatedTime: new Date().toISOString()
    };
    set({ savedCloudModels: [...get().savedCloudModels, demoModel] });
  },

  fetchTumblebugProviders: async () => {
    const list = await tumblebugApi.getProviders();
    set({ tumblebugProviders: list });
    // Dynamically pull regions for the currently set desiredCsp or the first returned provider
    const activeCsp = get().desiredCsp || (list.length > 0 ? list[0] : 'aws');
    const regions = await tumblebugApi.getRegions(activeCsp);
    set({ tumblebugRegions: regions });
  },

  fetchTumblebugRegions: async (providerName: string) => {
    const regions = await tumblebugApi.getRegions(providerName);
    set({ tumblebugRegions: regions });
  },

  // --------------------------------------------------------------------------
  // Page 3: Migration Execution - Live Connected to Beetle API
  // --------------------------------------------------------------------------
  namespaceId: 'mig01',
  nameSeed: 'name',
  isDeploying: false,
  activeDeploymentId: '',
  deploymentStatus: null,
  liveReportHtml: '',

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
    const nsId = get().namespaceId;
    const infraId = get().activeDeploymentId;
    if (!nsId || !infraId) return;

    try {
      const statusData = await beetleApi.getMigrationStatus(nsId, infraId);
      set({ deploymentStatus: statusData });
      
      if (statusData && (statusData.status === 'Succeeded' || statusData.status === 'Failed' || statusData.status === 'Completed')) {
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
  }
}));
