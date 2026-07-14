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

// Damselfly model envelopes
export interface OnpremModelEnvelope {
  id: string;
  name: string;
  description?: string;
  onpremiseInfraModel: OnpremInfra;
  version?: string;
  updatedTime?: string;
}

// Target Cloud recommendations & deployments (Beetle / Tumblebug)
export interface TargetCloud {
  csp: string;
  region: string;
}

export interface NodeGroup {
  name: string;
  nodeGroupSize: number;
  label: {
    nlbBackend?: string;
    sourceMachineIds: string; // Comma separated list of source machine UIDs
  };
  description: string;
  connectionName: string;
  specId: string;
  imageId: string;
  vNetId: string;
  subnetId: string;
  securityGroupIds: string[];
  sshKeyId: string;
  rootDiskSize: number;
  dataDiskIds?: string[] | null;
}

export interface TargetInfra {
  name: string;
  description: string;
  nodeGroups: NodeGroup[];
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
  status: 'highly-matched' | 'partially-matched' | 'unacceptable';
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

// Damselfly cloud model envelop
export interface CloudModelEnvelope {
  id: string;
  name: string;
  description?: string;
  cloudInfraModel: RecommendedInfra;
  version?: string;
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
