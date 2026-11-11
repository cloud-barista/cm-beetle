// Source infrastructure types (Onpremise / Honeybee)
export interface CpuProperty {
  architecture: string;
  cores: number;
  cpus: number;
  maxSpeed: number;
  model: string;
  threads: number;
  vendor: string;
}

export interface MemoryProperty {
  available: number;
  totalSize: number;
  type: string;
}

export interface DiskProperty {
  label: string;
  totalSize: number;
  type: string;
}

export interface InterfaceProperty {
  name: string;
  state: string;
  macAddress?: string;
  mtu: number;
  ipv4CidrBlocks: string[];
  ipv6CidrBlocks: string[];
}

export interface RouteProperty {
  destination: string;
  gateway: string;
  interface: string;
  linkState: string;
  metric: number;
  protocol: string;
  scope: string;
}

export interface FirewallRule {
  // Lowercase fields — on-premise source infra (onpremmodel.FirewallRuleProperty)
  action?: string;
  direction?: string;
  dstCIDR?: string;
  dstPorts?: string;
  protocol?: string;
  srcCIDR?: string;
  srcPorts?: string;
  // Capitalized fields — CB-Tumblebug target recommendation (cloudmodel.FirewallRuleReq)
  Direction?: string;
  Protocol?: string;
  CIDR?: string;
  Ports?: string;
}

export interface OsProperty {
  id: string;
  idLike: string;
  name: string;
  prettyName: string;
  version: string;
  versionCodename: string;
  versionId: string;
}

export interface GpuCardProperty {
  driverIndex?: string;
  uuid?: string;
  vendor?: string;
  model?: string;
  type?: string;
  architecture?: string;
  driverVersion?: string;
  cudaVersion?: string;
  slot?: string;
  pciBusId?: string;
  eccEnabled?: boolean;
  memoryTotalGB?: number;
  memoryReservedGB?: number;
  memoryFreeGB?: number;
  memoryUsedGB?: number;
}

export interface OnpremNode {
  machineId: string;
  hostname: string;
  cpu: CpuProperty;
  memory: MemoryProperty;
  rootDisk: DiskProperty;
  dataDisks?: DiskProperty[] | null;
  interfaces: InterfaceProperty[];
  routingTable: RouteProperty[];
  firewallTable: FirewallRule[];
  os: OsProperty;
  gpuCards?: GpuCardProperty[] | null;
  workloadTag?: string;
  clusterGroup?: 'ai-inference' | 'analytics-compute' | 'backend-db';
  serviceUrl?: string;
}

export interface GatewayProperty {
  interfaceName: string;
  ip: string;
  machineId: string;
}

export interface OnpremNetwork {
  ipv4Networks: {
    defaultGateways?: GatewayProperty[];
    cidrBlocks?: string[];
  };
  ipv6Networks: Record<string, unknown>;
}

export interface OnpremNlbBackend {
  name: string;
  balance: string;
  protocol: string;
  servers: Array<{
    name: string;
    ip: string;
    port: number;
  }>;
}

export interface OnpremNlb {
  hostMachineId: string;
  software: string;
  listener: {
    bindAddress: string;
    port: number;
    protocol: string;
  };
  backend: OnpremNlbBackend;
  healthCheck: {
    enabled: boolean;
    interval: number;
    timeout: number;
    threshold: number;
  };
}

export interface OnpremInfra {
  network: OnpremNetwork;
  nodes: OnpremNode[];
  nlbs?: OnpremNlb[];
}

// Migration artifact envelopes
export interface ArtifactLineage {
  sourceGroupId?: string;
  sourceModelId?: string;
  planningModelId?: string;
  targetModelId?: string;
}

export interface OnpremModelEnvelope {
  id: string;
  name: string;
  description?: string;
  onpremiseInfraModel: OnpremInfra;
  version?: string;
  stage?: 'source' | 'planning' | 'target';
  category?: 'infra' | 'software' | 'data';
  schemaType?: string;
  lineage?: ArtifactLineage;
  updatedTime?: string;
}

// Target Cloud recommendations & deployments (Beetle / Tumblebug)
export interface TargetCloud {
  csp: string;
  region: string;
}

export type CloudProperty = TargetCloud;

export interface NodeGroup {
  name: string;
  nodeGroupSize: number;
  label: {
    nlbBackend?: string;
    sourceMachineIds?: string;
    'cm-source-machine-ids'?: string;
    [key: string]: string | undefined;
  };
  description: string;
  connectionName: string;
  specId: string;
  imageId: string;
  cspImageName?: string;
  vNetId: string;
  subnetId: string;
  securityGroupIds: string[];
  sshKeyId: string;
  rootDiskSize: number;
  rootDiskType?: string;
  dataDiskIds?: string[] | null;
}

export interface PostCommandReq {
  commands?: string[];
  cmd?: string[];
  nodeGroupName?: string;
  nodeName?: string;
  userName?: string;
}

export interface TargetInfra {
  name: string;
  description: string;
  nodeGroups: NodeGroup[];
  postCommands?: PostCommandReq[];
  postCommandAsync?: boolean;
  label?: Record<string, string>;
}

export interface SubnetInfo {
  name: string;
  ipv4_CIDR: string;
  description: string;
  zone?: string;
}

export interface TargetVNet {
  name: string;
  connectionName: string;
  cidrBlock: string;
  subnetInfoList: SubnetInfo[];
  description: string;
}

export interface TargetSshKey {
  name: string;
  connectionName: string;
  description: string;
  publicKey?: string;
  privateKey?: string;
}

export interface TargetSecurityGroup {
  name: string;
  connectionName: string;
  vNetId: string;
  description: string;
  firewallRules: FirewallRule[];
}

export interface TargetNlb {
  description: string;
  type: string; // PUBLIC | INTERNAL
  scope: string; // REGION | GLOBAL
  listener: {
    protocol: string;
    port: string;
  };
  targetGroup: {
    protocol: string;
    port: string;
    nodeGroupId: string;
  };
  healthChecker: {
    interval: number;
    threshold: number;
    timeout: number;
  };
}

// Recommended Candidate item
export interface RecommendedInfra {
  status: 'highly-matched' | 'partially-matched' | 'unacceptable' | 'failed' | 'nothing-to-recommend' | string;
  description: string;
  targetCloud: TargetCloud;
  targetInfra: TargetInfra;
  targetVNet: TargetVNet;
  targetSshKey: TargetSshKey;
  targetSecurityGroupList: TargetSecurityGroup[];
  targetNlbList?: TargetNlb[];
  targetSpecList?: any[];
  targetOsImageList?: any[];
}

// Target cloud artifact envelope
export interface CloudModelEnvelope {
  id: string;
  name: string;
  description?: string;
  cloudInfraModel: RecommendedInfra;
  version?: string;
  stage?: 'source' | 'planning' | 'target';
  category?: 'infra' | 'software' | 'data';
  schemaType?: string;
  lineage?: ArtifactLineage;
  updatedTime?: string;
}

// Deployment Live status types
export interface VmStatusInfo {
  id: string;
  name: string;
  status: string; // Running, Creating, Failed
  publicIp: string;
  privateIp: string;
  specId: string;
  imageId: string;
}

export interface NodeGroupStatus {
  name: string;
  size: number;
  vms: VmStatusInfo[];
}

export interface DeploymentStatus {
  infraId: string;
  status: string; // Succeeded, Deploying, Failed, PartialFailed
  progress: number; // 0 to 100
  nodeGroups: NodeGroupStatus[];
  nlbStatus?: {
    id: string;
    dnsName?: string;
    publicIp?: string;
    status: string;
  };
}

// SSH Readiness Check types (GET /beetle/migration/ns/{nsId}/infra/{infraId}/ssh-ready)
export interface NodeSSHStatus {
  id: string;
  name: string;
  publicIP: string;
  privateIP: string;
  username: string;
  status: string;
  sshReady: boolean;
  sshPort: number;
  error?: string;
}

export interface CheckSSHReadyResponse {
  ready: boolean;
  totalNodes: number;
  readyNodes: number;
  nodeStatus?: NodeSSHStatus[];
  message: string;
  checkedAt: string;
  nextAllowedTime?: string;
}

// Outbound CB-Tumblebug sub-request tracking types (GET /beetle/request/{reqId})
export interface TumblebugRequestSnapshot {
  reqId: string;
  status: string;
  method: string;
  url: string;
  startTime?: string;
  endTime?: string;
  elapsedTime?: number;
  errorResponse?: string;
}

export interface SubRequestItem {
  seq: number;
  subReqId: string;
  method: string;
  targetPath: string;
  resourceType?: string;
  resourceHint?: string;
  status: string;
  statusCode?: number;
  startTime: string;
  endTime?: string;
  durationMs?: number;
  errorMessage?: string;
  upstreamStatus?: string;
  upstreamProgress?: any;
  upstreamError?: string;
}

export interface SubRequestSummary {
  total: number;
  handling: number;
  success: number;
  error: number;
  currentAction?: string;
}

export interface DataHookOutputs {
  sourcePreCmd?: string;
  destinationPreCmd?: string;
  sourcePostCmd?: string;
  destinationPostCmd?: string;
  [key: string]: string | undefined;
}

export interface RequestDetailsResult {
  status: string;
  errorResponse?: string;
  responseData?: any;
  tumblebugSource?: string;
  subRequestSummary?: SubRequestSummary;
  subRequests?: SubRequestItem[];
}

export interface AiStackLayerStatus {
  host: string;
  checkedAt: string;
  isSample?: boolean;
  openWebUi: {
    ok: boolean;
    port?: number;
    httpStatus: number;
    latencyMs?: number;
    version?: string;
    error?: string;
  };
  ollamaTags: {
    ok: boolean;
    latencyMs?: number;
    modelName?: string;
    sizeBytes?: number;
    digest?: string;
    parameterSize?: string;
    quantizationLevel?: string;
    totalModels?: number;
    error?: string;
  };
  ollamaPs: {
    ok: boolean;
    vramBytes?: number;
    activeModel?: string;
    gpuAccelerated?: boolean;
    error?: string;
  };
  inference: {
    ok: boolean;
    latencyMs?: number;
    tokensPerSec?: number;
    evalCount?: number;
    responseSnippet?: string;
    error?: string;
  };
}


