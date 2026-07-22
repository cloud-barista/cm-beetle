'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore } from '@/store/migrationStore';
import { beetleApi, tumblebugApi } from '@/api/client';
import { HardDrive, Plus, Trash2, CheckCircle2, ArrowRight, ArrowLeft, ShieldCheck, Database, RefreshCw, Server, AlertCircle, Play, Key, Search, ChevronDown, ChevronUp, Lock, FileText, Upload, Save, X } from 'lucide-react';

interface SourceBucket {
  bucketName: string;
  totalSizeBytes: number;
  objectCount: number;
  accessFrequency: 'frequent' | 'infrequent' | 'archive';
  versioningEnabled: boolean;
  encryptionEnabled: boolean;
  corsEnabled: boolean;
  corsRule?: any[];
  isPublic: boolean;
  creationDate?: string;
  tags?: Record<string, string>;
}

interface StorageGroup {
  id: string;
  name: string;
  csp: string;
  region: string;
  buckets: SourceBucket[];
}

const SAMPLE_STORAGE_GROUPS: StorageGroup[] = [
  {
    id: 'sg-sample-aws',
    name: '[Sample] aws-prod-storage',
    csp: 'aws',
    region: 'ap-northeast-2',
    buckets: [
      {
        bucketName: 'legacy-app-media-bucket',
        totalSizeBytes: 53687091200, // 50 GB
        objectCount: 14250,
        accessFrequency: 'frequent',
        versioningEnabled: true,
        encryptionEnabled: true,
        corsEnabled: false,
        isPublic: false
      }
    ]
  }
];

const ALL_SUPPORTED_CSPS = [
  { id: 'aws', name: 'AWS (Amazon Web Services)' },
  { id: 'gcp', name: 'GCP (Google Cloud Platform)' },
  { id: 'azure', name: 'Azure (Microsoft Azure)' },
  { id: 'ncp', name: 'NCP (Naver Cloud Platform)' },
  { id: 'alibaba', name: 'Alibaba Cloud' },
  { id: 'tencent', name: 'Tencent Cloud' },
  { id: 'openstack', name: 'OpenStack' },
  { id: 'ibm', name: 'IBM Cloud' }
];

const CSP_REGIONS_MAP: Record<string, { id: string; name: string }[]> = {
  aws: [
    { id: 'ap-northeast-2', name: 'ap-northeast-2 (Seoul)' },
    { id: 'ap-northeast-1', name: 'ap-northeast-1 (Tokyo)' },
    { id: 'us-east-1', name: 'us-east-1 (N. Virginia)' },
    { id: 'us-west-2', name: 'us-west-2 (Oregon)' },
    { id: 'eu-central-1', name: 'eu-central-1 (Frankfurt)' },
    { id: 'eu-west-1', name: 'eu-west-1 (Ireland)' },
    { id: 'eu-north-1', name: 'eu-north-1 (Stockholm)' },
    { id: 'ap-southeast-1', name: 'ap-southeast-1 (Singapore)' },
    { id: 'ap-southeast-2', name: 'ap-southeast-2 (Sydney)' }
  ],
  gcp: [
    { id: 'asia-northeast3', name: 'asia-northeast3 (Seoul)' },
    { id: 'asia-northeast1', name: 'asia-northeast1 (Tokyo)' },
    { id: 'us-central1', name: 'us-central1 (Iowa)' },
    { id: 'europe-west1', name: 'europe-west1 (Belgium)' },
    { id: 'asia-southeast1', name: 'asia-southeast1 (Singapore)' }
  ],
  azure: [
    { id: 'koreacentral', name: 'koreacentral (Korea Central)' },
    { id: 'japaneast', name: 'japaneast (Japan East)' },
    { id: 'eastus', name: 'eastus (East US)' },
    { id: 'westeurope', name: 'westeurope (West Europe)' },
    { id: 'southeastasia', name: 'southeastasia (Southeast Asia)' }
  ],
  ncp: [
    { id: 'ap-northeast-2', name: 'ap-northeast-2 (Korea)' },
    { id: 'kr-1', name: 'kr-1 (Korea Region 1)' },
    { id: 'kr-2', name: 'kr-2 (Korea Region 2)' },
    { id: 'sng-1', name: 'sng-1 (Singapore)' }
  ],
  alibaba: [
    { id: 'cn-hongkong', name: 'cn-hongkong (Hong Kong)' },
    { id: 'ap-northeast-1', name: 'ap-northeast-1 (Japan)' },
    { id: 'ap-southeast-1', name: 'ap-southeast-1 (Singapore)' },
    { id: 'us-west-1', name: 'us-west-1 (Silicon Valley)' }
  ],
  tencent: [
    { id: 'ap-seoul', name: 'ap-seoul (Seoul)' },
    { id: 'ap-tokyo', name: 'ap-tokyo (Tokyo)' },
    { id: 'ap-singapore', name: 'ap-singapore (Singapore)' },
    { id: 'na-siliconvalley', name: 'na-siliconvalley (Silicon Valley)' }
  ],
  openstack: [
    { id: 'RegionOne', name: 'RegionOne' },
    { id: 'region-1', name: 'region-1' }
  ],
  ibm: [
    { id: 'us-south', name: 'us-south (Dallas)' },
    { id: 'us-east', name: 'us-east (Washington DC)' },
    { id: 'eu-gb', name: 'eu-gb (London)' },
    { id: 'eu-de', name: 'eu-de (Frankfurt)' },
    { id: 'jp-tok', name: 'jp-tok (Tokyo)' }
  ]
};

export const ObjectStorageMigration: React.FC = () => {
  const { namespaceId, tumblebugProviders, tumblebugRegions, fetchTumblebugProviders, fetchTumblebugRegions } = useMigrationStore();

  // Storage Groups State
  const [storageGroups, setStorageGroups] = useState<StorageGroup[]>([]);
  const [activeGroupId, setActiveGroupId] = useState<string | null>(null);
  const [isNewGroupModalOpen, setIsNewGroupModalOpen] = useState(false);
  const [newGroupName, setNewGroupName] = useState('');

  // Combined CSP List (Static Master List + Tumblebug Providers fallback)
  const getCspList = () => {
    const map = new Map<string, string>();
    ALL_SUPPORTED_CSPS.forEach((c) => map.set(c.id, c.name));
    if (tumblebugProviders && tumblebugProviders.length > 0) {
      tumblebugProviders.forEach((p: string) => {
        const key = p.toLowerCase();
        if (!map.has(key)) map.set(key, `${p.toUpperCase()} (${p})`);
      });
    }
    return Array.from(map.entries()).map(([id, name]) => ({ id, name }));
  };

  // Dynamic Region Selector for CSP (Static Master Map + Tumblebug Regions fallback)
  const getRegionsForCsp = (csp: string) => {
    const key = (csp || 'aws').toLowerCase();
    const staticRegions = CSP_REGIONS_MAP[key] || [];
    const tbRegions = (tumblebugRegions || [])
      .filter((r: any) => !r.providerName || r.providerName.toLowerCase() === key)
      .map((r: any) => ({ id: r.id, name: `${r.id} (${r.locationName || r.name || r.id})` }));

    const regionMap = new Map<string, string>();
    staticRegions.forEach((r) => regionMap.set(r.id, r.name));
    tbRegions.forEach((r: any) => {
      if (!regionMap.has(r.id)) regionMap.set(r.id, r.name);
    });

    const result = Array.from(regionMap.entries()).map(([id, name]) => ({ id, name }));
    return result.length > 0
      ? result
      : [
          { id: 'ap-northeast-2', name: 'ap-northeast-2 (Seoul)' },
          { id: 'us-east-1', name: 'us-east-1 (N. Virginia)' }
        ];
  };

  const handleCspChange = (newCsp: string) => {
    setScanCsp(newCsp);
    fetchTumblebugRegions(newCsp);
    const available = getRegionsForCsp(newCsp);
    if (available.length > 0) {
      setScanRegion(available[0].id);
    }
  };

  // Collapsible Credential Management & Modal State
  const [isCredentialOpen, setIsCredentialOpen] = useState(true);
  const [credentialSourceMode, setCredentialSourceMode] = useState<'saved' | 'new'>('saved');
  const [selectedCredentialProfile, setSelectedCredentialProfile] = useState('');
  const [isRegisterCredModalOpen, setIsRegisterCredModalOpen] = useState(false);

  const [credProfileName, setCredProfileName] = useState('');
  const [credCsp, setCredCsp] = useState('aws');
  const [credRegion, setCredRegion] = useState('ap-northeast-2');
  const [credAccessKey, setCredAccessKey] = useState('');
  const [credSecretKey, setCredSecretKey] = useState('');
  const [credTenantId, setCredTenantId] = useState('');
  const [credSubId, setCredSubId] = useState('');
  const [credS3AccessKey, setCredS3AccessKey] = useState('');
  const [credS3SecretKey, setCredS3SecretKey] = useState('');

  const [savedCredProfiles, setSavedCredProfiles] = useState<
    { id: string; name: string; csp: string; region: string; accessKey: string }[]
  >([]);

  // Registration Mode State: 'manual' vs 'scan'
  const [registrationMode, setRegistrationMode] = useState<'manual' | 'scan'>('manual');

  // Source Bucket Scan State (Credential-based Auto Discovery)
  const [scanCsp, setScanCsp] = useState('aws');
  const [scanRegion, setScanRegion] = useState('ap-northeast-2');
  const [scanAccessKey, setScanAccessKey] = useState('');
  const [scanSecretKey, setScanSecretKey] = useState('');
  const [scanTenantId, setScanTenantId] = useState('');
  const [scanSubId, setScanSubId] = useState('');
  const [scanS3AccessKey, setScanS3AccessKey] = useState('');
  const [scanS3SecretKey, setScanS3SecretKey] = useState('');
  const [isScanning, setIsScanning] = useState(false);
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);

  // Pixel-perfect Top/Bottom split: Top (CSP & Region) | Bottom (Dynamic Credential Fields)
  const renderCspCredentialBlock = (
    csp: string,
    onCspChange: (csp: string) => void,
    region: string,
    onRegionChange: (r: string) => void,
    keyVal: string, setKeyVal: (v: string) => void,
    secretVal: string, setSecretVal: (v: string) => void,
    tenantVal: string, setTenantVal: (v: string) => void,
    subVal: string, setSubVal: (v: string) => void,
    s3KeyVal?: string, setS3KeyVal?: (v: string) => void,
    s3SecretVal?: string, setS3SecretVal?: (v: string) => void
  ) => {
    const c = csp.toLowerCase();

    return (
      <div className="space-y-4">
        {/* Top Section: Target Cloud Provider & Region Selection */}
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 pb-3 border-b border-border-main/50">
          <div>
            <label className="block text-text-muted font-bold mb-1">Source CSP</label>
            <select
              value={csp}
              onChange={(e) => onCspChange(e.target.value)}
              className="w-full px-3.5 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold uppercase focus:outline-none focus:border-emerald-500 cursor-pointer"
            >
              {getCspList().map((c) => (
                <option key={c.id} value={c.id}>
                  {c.name}
                </option>
              ))}
            </select>
          </div>
          <div>
            <label className="block text-text-muted font-bold mb-1">Region</label>
            <select
              value={region}
              onChange={(e) => onRegionChange(e.target.value)}
              className="w-full px-3.5 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 cursor-pointer"
            >
              {getRegionsForCsp(csp).map((r) => (
                <option key={r.id} value={r.id}>
                  {r.name}
                </option>
              ))}
            </select>
          </div>
        </div>

        {/* Bottom Section: Dynamic Authentication Credentials per CSP (Exact CSP Native Labels) */}
        <div>
          {c === 'azure' && (
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-3.5">
              <div>
                <label className="block text-text-muted font-medium mb-1">Application (Client) ID</label>
                <input
                  type="text"
                  placeholder="00000000-0000-0000-0000-000000000000"
                  value={keyVal}
                  onChange={(e) => setKeyVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
              <div>
                <label className="block text-text-muted font-medium mb-1">Client Secret Value</label>
                <input
                  type="password"
                  placeholder="••••••••••••••••"
                  value={secretVal}
                  onChange={(e) => setSecretVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
              <div>
                <label className="block text-text-muted font-medium mb-1">Directory (Tenant) ID</label>
                <input
                  type="text"
                  placeholder="Tenant GUID"
                  value={tenantVal}
                  onChange={(e) => setTenantVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
              <div>
                <label className="block text-text-muted font-medium mb-1">Subscription ID</label>
                <input
                  type="text"
                  placeholder="Subscription GUID"
                  value={subVal}
                  onChange={(e) => setSubVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
            </div>
          )}

          {c === 'gcp' && (
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-3.5">
              <div>
                <label className="block text-text-muted font-medium mb-1">Project ID</label>
                <input
                  type="text"
                  placeholder="my-gcp-project-id"
                  value={keyVal}
                  onChange={(e) => setKeyVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
              <div>
                <label className="block text-text-muted font-medium mb-1">Client Email (Service Account Email)</label>
                <input
                  type="text"
                  placeholder="sa-name@project-id.iam.gserviceaccount.com"
                  value={tenantVal}
                  onChange={(e) => setTenantVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
              <div>
                <label className="block text-text-muted font-medium mb-1">Private Key (JSON / Key Content)</label>
                <input
                  type="password"
                  placeholder="••••••••••••••••"
                  value={secretVal}
                  onChange={(e) => setSecretVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
            </div>
          )}

          {c === 'openstack' && (
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-3.5">
              <div>
                <label className="block text-text-muted font-medium mb-1">Username</label>
                <input
                  type="text"
                  placeholder="admin"
                  value={keyVal}
                  onChange={(e) => setKeyVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
              <div>
                <label className="block text-text-muted font-medium mb-1">Password</label>
                <input
                  type="password"
                  placeholder="••••••••••••••••"
                  value={secretVal}
                  onChange={(e) => setSecretVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
              <div>
                <label className="block text-text-muted font-medium mb-1">Project (Tenant) Name</label>
                <input
                  type="text"
                  placeholder="my-project"
                  value={tenantVal}
                  onChange={(e) => setTenantVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
            </div>
          )}

          {c !== 'azure' && c !== 'gcp' && c !== 'openstack' && (
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-3.5">
              <div>
                <label className="block text-text-muted font-medium mb-1">
                  {c === 'tencent' ? 'SecretId' : c === 'alibaba' ? 'AccessKey ID' : c === 'ncp' ? 'Access Key ID' : 'Access Key ID'}
                </label>
                <input
                  type="text"
                  placeholder={c === 'ncp' ? 'NCP Access Key' : c === 'alibaba' ? 'LTAI...' : c === 'tencent' ? 'AKID...' : 'AKIA...'}
                  value={keyVal}
                  onChange={(e) => setKeyVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
              <div>
                <label className="block text-text-muted font-medium mb-1">
                  {c === 'tencent' ? 'SecretKey' : c === 'alibaba' ? 'AccessKey Secret' : c === 'ncp' ? 'Secret Key' : 'Secret Access Key'}
                </label>
                <input
                  type="password"
                  placeholder="••••••••••••••••"
                  value={secretVal}
                  onChange={(e) => setSecretVal(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                />
              </div>
            </div>
          )}
          {/* S3 Interoperability Section for GCP, Azure, IBM, OpenStack, NHN, KT */}
          {['gcp', 'azure', 'ibm', 'openstack', 'nhn', 'nhncloud', 'kt', 'ktcloud'].includes(c) && (
            <div className="pt-3 border-t border-border-main/40 mt-3.5 space-y-2">
              <div className="flex items-center space-x-1.5 text-text-muted">
                <span className="text-[11px] font-extrabold text-teal-400">S3 Interoperability Credentials (AWS S3-Compatible Interoperability)</span>
                <span className="text-[10px] text-text-muted">(S3AccessKey &amp; S3SecretKey for Object Storage API Control)</span>
              </div>
              <div className="grid grid-cols-1 sm:grid-cols-2 gap-3.5">
                <div>
                  <label className="block text-text-muted font-medium mb-1">S3 Access Key (S3AccessKey)</label>
                  <input
                    type="text"
                    placeholder="S3 Interoperability Access Key"
                    value={s3KeyVal || ''}
                    onChange={(e) => setS3KeyVal && setS3KeyVal(e.target.value)}
                    className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                  />
                </div>
                <div>
                  <label className="block text-text-muted font-medium mb-1">S3 Secret Key (S3SecretKey)</label>
                  <input
                    type="password"
                    placeholder="••••••••••••••••"
                    value={s3SecretVal || ''}
                    onChange={(e) => setS3SecretVal && setS3SecretVal(e.target.value)}
                    className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                  />
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    );
  };

  // Source Buckets Form State (Initial State: Empty)
  const [sourceBuckets, setSourceBuckets] = useState<SourceBucket[]>([]);

  const [newBucket, setNewBucket] = useState<SourceBucket>({
    bucketName: '',
    totalSizeBytes: 10737418240, // 10 GB
    objectCount: 1000,
    accessFrequency: 'frequent',
    versioningEnabled: false,
    encryptionEnabled: true,
    corsEnabled: false,
    isPublic: false
  });

  // Fast Scan & Selection State (Steps 3~6)
  const [scannedBucketNames, setScannedBucketNames] = useState<string[]>([]);
  const [scannedBuckets, setScannedBuckets] = useState<any[]>([]);
  const [selectedBucketNames, setSelectedBucketNames] = useState<string[]>([]);
  const [isInspectLoading, setIsInspectLoading] = useState(false);

  const formatBytes = (bytes?: number) => {
    if (!bytes || bytes === 0) return '0 B';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i];
  };

  // Model Persistence State (Steps 7~12, 15~20)
  const [collectedSourceModel, setCollectedSourceModel] = useState<any | null>(null);
  const [savedSourceModelId, setSavedSourceModelId] = useState<string | null>(null);
  const [savedTargetModelId, setSavedTargetModelId] = useState<string | null>(null);
  const [modelLog, setModelLog] = useState<string[]>([]);

  // Step 3~4: Fast Bucket List Scan Handler
  const handleScanSourceBuckets = async () => {
    const effectiveAccessKey = scanS3AccessKey || scanAccessKey;
    const effectiveSecretKey = scanS3SecretKey || scanSecretKey;
    if (!effectiveAccessKey || !effectiveSecretKey) {
      alert('Please enter Access Key ID and Secret Access Key before scanning buckets.');
      return;
    }

    setIsScanning(true);
    try {
      const res = await beetleApi.scanSourceObjectStorage({
        csp: scanCsp,
        region: scanRegion,
        accessKeyId: scanAccessKey,
        secretAccessKey: scanSecretKey,
        tenantId: scanTenantId,
        subscriptionId: scanSubId,
        s3AccessKey: scanS3AccessKey,
        s3SecretKey: scanS3SecretKey,
      });
      if (res.success && res.bucketNames && res.bucketNames.length > 0) {
        const sortedBucketNames = [...res.bucketNames].sort((a, b) => a.localeCompare(b));
        const sortedBuckets = [...(res.buckets || [])].sort((a: any, b: any) =>
          (typeof a === 'string' ? a : a.bucketName || '').localeCompare(typeof b === 'string' ? b : b.bucketName || '')
        );
        setScannedBucketNames(sortedBucketNames);
        setScannedBuckets(sortedBuckets);
        setSelectedBucketNames([sortedBucketNames[0]]);
      } else {
        alert(res.error || 'No buckets found for the specified region or scan failed.');
      }
    } catch (e: any) {
      console.error('Scan error:', e);
      alert('Error scanning cloud account buckets.');
    } finally {
      setIsScanning(false);
    }
  };

  // Step 5~6: Selected Bucket Metadata Inspection Handler
  const handleInspectSelectedBuckets = async () => {
    const bucketsToInspect = sourceBuckets.length > 0
      ? sourceBuckets.map((b) => b.bucketName)
      : (selectedBucketNames.length > 0 ? selectedBucketNames : scannedBucketNames);

    if (bucketsToInspect.length === 0) {
      alert('Please select or add at least one object storage bucket to collect metadata.');
      return;
    }

    const activeCred = savedCredProfiles.find((c) => c.id === selectedCredentialProfile || c.name === selectedCredentialProfile) || savedCredProfiles[0];
    const keyToScan = activeCred ? (activeCred.accessKey || (activeCred as any).s3AccessKey) : (scanS3AccessKey || scanAccessKey || credAccessKey);
    const secretToScan = activeCred ? ((activeCred as any).secretKey || (activeCred as any).s3SecretKey) : (scanS3SecretKey || scanSecretKey || credSecretKey);

    if (!keyToScan || !secretToScan) {
      alert('Access Key ID and Secret Access Key are required to collect bucket metadata. Please register or select a valid credential profile.');
      return;
    }

    setIsInspectLoading(true);
    try {
      const res = await beetleApi.inspectSourceObjectStorage({
        csp: activeCred ? activeCred.csp : scanCsp,
        region: activeCred ? activeCred.region : scanRegion,
        accessKeyId: keyToScan,
        secretAccessKey: secretToScan,
        s3AccessKey: activeCred ? (activeCred as any).s3AccessKey : (scanS3AccessKey || credS3AccessKey),
        s3SecretKey: activeCred ? (activeCred as any).s3SecretKey : (scanS3SecretKey || credS3SecretKey),
        tenantId: activeCred ? (activeCred as any).tenantId : (scanTenantId || credTenantId),
        subscriptionId: activeCred ? (activeCred as any).subId : (scanSubId || credSubId),
        selectedBucketNames: bucketsToInspect
      });
      if (res.success && res.inspectedBuckets && res.inspectedBuckets.length > 0) {
        setSourceBuckets(res.inspectedBuckets);
        if (res.sourceObjectStorage) {
          setCollectedSourceModel(res.sourceObjectStorage);
        }
        setRegistrationMode('manual');
        alert(`Successfully collected deep metadata for ${res.inspectedBuckets.length} object storage bucket(s)!`);
      } else {
        alert(res.error || 'Failed to inspect bucket metadata.');
      }
    } catch (e: any) {
      console.error('Inspect error:', e);
      alert('Error collecting object storage metadata.');
    } finally {
      setIsInspectLoading(false);
    }
  };

  // Step 7~8: Save Source User Model (Damselfly / Beetle API)
  const handleSaveSourceModel = async () => {
    if (sourceBuckets.length === 0) {
      alert('No source buckets registered to save.');
      return;
    }
    const formattedBuckets = sourceBuckets.map((b) => ({
      accessFrequency: b.accessFrequency || 'frequent',
      bucketName: b.bucketName,
      corsEnabled: b.corsEnabled || false,
      corsRule: b.corsRule || [],
      creationDate: b.creationDate || '',
      encryptionEnabled: b.encryptionEnabled || false,
      isPublic: b.isPublic || false,
      objectCount: b.objectCount || 0,
      tags: b.tags || {},
      totalSizeBytes: b.totalSizeBytes || 0,
      versioningEnabled: b.versioningEnabled || false
    }));

    const res = await beetleApi.saveSourceObjectStorageModel({
      namespaceId,
      sourceObjectStorages: formattedBuckets
    });
    if (res.success) {
      setSavedSourceModelId(res.modelId);
      setModelLog((prev) => [...prev, `[Damselfly/Beetle] Source User Model Saved (ID: ${res.modelId})`]);
    }
  };

  // Step 15~16: Save Target User Model (Damselfly / Beetle Mock API)
  const handleSaveTargetModel = async () => {
    if (!recommendationResult) {
      alert('Generate target recommendation first.');
      return;
    }
    const res = await beetleApi.saveTargetObjectStorageModel({
      namespaceId,
      recommendation: recommendationResult,
      nameSeed
    });
    if (res.success) {
      setSavedTargetModelId(res.modelId);
      setModelLog((prev) => [...prev, `[Damselfly/Beetle] Target User Model Saved (ID: ${res.modelId})`]);
    }
  };

  // Recommendation & Provisioning State
  const [desiredCsp, setDesiredCsp] = useState('aws');
  const [desiredRegion, setDesiredRegion] = useState('ap-northeast-2');
  const [nameSeed, setNameSeed] = useState('mig');
  const [recommendationResult, setRecommendationResult] = useState<any | null>(null);
  const [isRecommending, setIsRecommending] = useState(false);
  const [isDeploying, setIsDeploying] = useState(false);
  const [deploymentLog, setDeploymentLog] = useState<string[]>([]);
  const [activeReqId, setActiveReqId] = useState<string | null>(null);

  // Existing Migrated Storage List
  const [migratedStorages, setMigratedStorages] = useState<any[]>([]);
  const [isLoadingMigrated, setIsLoadingMigrated] = useState(false);

  useEffect(() => {
    // Zero automatic background API calls on page mount to prevent unnecessary HTTP traffic
  }, [namespaceId]);

  const handleDesiredCspChange = async (csp: string) => {
    setDesiredCsp(csp);
    await fetchTumblebugRegions(csp);
  };

  const addSourceBucket = () => {
    if (!newBucket.bucketName.trim()) {
      alert('Please enter a valid Bucket Name.');
      return;
    }
    setSourceBuckets((prev) => [...prev, { ...newBucket }]);
    setNewBucket({
      bucketName: '',
      totalSizeBytes: 10737418240,
      objectCount: 1000,
      accessFrequency: 'frequent',
      versioningEnabled: false,
      encryptionEnabled: true,
      corsEnabled: false,
      isPublic: false
    });
  };

  const removeSourceBucket = (index: number) => {
    setSourceBuckets((prev) => prev.filter((_, i) => i !== index));
  };

  const handleGetRecommendation = async () => {
    if (sourceBuckets.length === 0) {
      alert('Please register at least one source bucket.');
      return;
    }

    const payloadSourceBuckets = sourceBuckets.map((b) => ({
      accessFrequency: b.accessFrequency || 'frequent',
      bucketName: b.bucketName,
      corsEnabled: b.corsEnabled || false,
      corsRule: b.corsRule || [],
      creationDate: b.creationDate || '',
      encryptionEnabled: b.encryptionEnabled || false,
      isPublic: b.isPublic || false,
      objectCount: b.objectCount || 0,
      tags: b.tags || {},
      totalSizeBytes: b.totalSizeBytes || 0,
      versioningEnabled: b.versioningEnabled || false
    }));

    setIsRecommending(true);
    try {
      const res = await beetleApi.recommendObjectStorage(desiredCsp, desiredRegion, payloadSourceBuckets);
      setRecommendationResult(res);
    } catch (err: any) {
      console.error('Recommendation failed', err);
      // Fallback demo recommendation envelope if API server returns mock/offline error
      setRecommendationResult({
        status: 'Recommended',
        description: `Recommended target storage for ${sourceBuckets.length} bucket(s) on ${desiredCsp.toUpperCase()} (${desiredRegion})`,
        targetCloud: { csp: desiredCsp, region: desiredRegion },
        targetObjectStorages: sourceBuckets.map((b) => ({
          bucketName: b.bucketName,
          sourceBucketName: b.bucketName,
          versioningEnabled: b.versioningEnabled,
          corsEnabled: b.corsEnabled,
          storageClass: b.accessFrequency === 'frequent' ? 'Standard' : 'Cool/Glacier'
        }))
      });
    } finally {
      setIsRecommending(false);
    }
  };

  const handleMigrateStorage = async () => {
    if (!recommendationResult) {
      alert('Please run recommendation first before provisioning.');
      return;
    }
    setIsDeploying(true);
    setDeploymentLog([
      `Initiating Object Storage Migration in namespace '${namespaceId}'...`,
      `Target CSP: ${desiredCsp.toUpperCase()}, Region: ${desiredRegion}`,
      `Applying Late-Binding Seed: '${nameSeed}'`
    ]);

    try {
      const res = await beetleApi.migrateObjectStorage(namespaceId, recommendationResult, nameSeed);
      if (res.success) {
        setDeploymentLog((prev) => [
          ...prev,
          `API Response: 202 Accepted (ReqID: ${res.reqId || 'req-async-001'})`,
          `Target Buckets Creation Request Dispatched to CB-Tumblebug.`,
          `Migration job handling in background.`
        ]);
        setActiveReqId(res.reqId || 'req-async-001');
        setTimeout(() => {
          loadMigratedStorages();
        }, 3000);
      } else {
        setDeploymentLog((prev) => [...prev, `Error: ${res.error || 'Failed to migrate object storage'}`]);
      }
    } catch (err: any) {
      setDeploymentLog((prev) => [...prev, `Failed to dispatch migration request: ${err.message}`]);
    } finally {
      setIsDeploying(false);
    }
  };

  const loadMigratedStorages = async () => {
    setIsLoadingMigrated(true);
    try {
      const list = await beetleApi.getMigratedObjectStorages(namespaceId);
      setMigratedStorages(list);
    } catch (err) {
      console.warn('Failed to load migrated storages', err);
    } finally {
      setIsLoadingMigrated(false);
    }
  };

  const handleDeleteStorage = async (osId: string) => {
    if (!confirm(`Are you sure you want to delete object storage '${osId}'?`)) return;
    try {
      await beetleApi.deleteMigratedObjectStorage(namespaceId, osId);
      await loadMigratedStorages();
    } catch (err) {
      alert('Failed to delete storage');
    }
  };
  // Workflow Sub-Tab State
  const [subTab, setSubTab] = useState<'source' | 'recommend' | 'provision'>('source');

  const subSteps = [
    { id: 'source', label: '1. Source Storage', icon: Database, desc: 'Register source buckets & scan credentials' },
    { id: 'recommend', label: '2. Target Recommendation', icon: ShieldCheck, desc: 'Recommend target storage classes & specs' },
    { id: 'provision', label: '3. Bucket Provisioning', icon: HardDrive, desc: 'Deploy target buckets & manage storages' },
  ] as const;

  return (
    <div className="space-y-6 animate-fade-in">
      {/* Unified Workflow Container Box */}
      <div className="bg-bg-panel border border-border-main rounded-2xl p-4 shadow-sm space-y-3">
        {/* Row 1: Workflow Title Line */}
        <div className="flex items-center space-x-2.5 border-b border-border-main pb-3 px-1">
          <HardDrive className="w-5 h-5 text-emerald-500" />
          <h2 className="text-base font-extrabold text-text-main flex items-center space-x-2">
            <span>Object Storage Migration Workflow</span>
            <span className="px-2 py-0.5 text-xs font-mono font-extrabold bg-amber-500/10 text-amber-600 dark:text-amber-400 border border-amber-500/20 rounded-md">
              WIP
            </span>
          </h2>
        </div>

        {/* Row 2: Workflow Tab Cards */}
        <nav className="grid grid-cols-1 sm:grid-cols-3 gap-2.5">
          {subSteps.map((step) => {
            const Icon = step.icon;
            const isActive = subTab === step.id;
            return (
              <button
                key={step.id}
                onClick={() => setSubTab(step.id)}
                className={`flex items-center p-3.5 rounded-xl border text-left transition-all duration-200 cursor-pointer ${
                  isActive
                    ? 'border-emerald-500 bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 font-extrabold shadow-sm'
                    : 'border-border-main bg-bg-main/20 text-text-muted hover:border-emerald-500/30 hover:bg-bg-panel'
                }`}
              >
                <Icon className={`w-5 h-5 mr-3 shrink-0 ${isActive ? 'text-emerald-500' : 'text-text-muted'}`} />
                <div className="min-w-0">
                  <div className="text-sm font-bold truncate">{step.label}</div>
                  <div className="text-xs text-text-muted truncate mt-0.5">{step.desc}</div>
                </div>
              </button>
            );
          })}
        </nav>
      </div>

      {/* Sub-step 1: Source Storage Registration & Credential Scan */}
      {subTab === 'source' && (
        <div className="space-y-6">
          {/* Top Banner Description Box (Rule 1: Tab Description Header Box) */}
          <div className="glass-panel px-6 py-4.5 rounded-2xl border border-border-main flex flex-wrap items-center gap-x-3 gap-y-1.5">
            <div className="flex items-center gap-2 shrink-0">
              <Database className="w-5 h-5 text-emerald-500" />
              <h2 className="text-base font-extrabold text-text-main tracking-tight">
                Source Object Storage Analysis
              </h2>
            </div>
            <span className="text-sm text-text-muted">
              Register source object storage buckets, scan credentials, and analyze infrastructure specifications & object metadata.
            </span>
          </div>

          {/* ══════════════════════════════════════════════════
              SECTION 1 — Source CSP Credentials (Ephemeral Profile Management)
          ══════════════════════════════════════════════════ */}
          <div className="glass-panel p-6 rounded-2xl">
            <div className="flex items-center justify-between mb-5">
              <div>
                <h2 className="text-base font-bold text-text-main flex items-center gap-2">
                  <Key className="w-5 h-5 text-emerald-400" />
                  Source CSP Credentials
                </h2>
                <p className="text-sm text-text-muted mt-1">
                  Select an active source CSP credential card to scan buckets, or register a new credential.
                </p>
              </div>
              <div className="flex items-center gap-3">
                <button
                  onClick={() => {
                    setCredProfileName(`cred-profile-${savedCredProfiles.length + 1}`);
                    setCredCsp(scanCsp || 'aws');
                    setCredRegion(scanRegion || 'ap-northeast-2');
                    setIsRegisterCredModalOpen(true);
                  }}
                  className="px-4 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white rounded-xl text-sm font-bold flex items-center gap-2 transition cursor-pointer shadow-lg shadow-emerald-500/20"
                >
                  <Plus className="w-4 h-4" /> New Credential
                </button>
              </div>
            </div>

            {/* Credential Cards Grid */}
            {savedCredProfiles.length === 0 ? (
              <div className="p-8 text-center border border-dashed border-border-main rounded-xl bg-bg-main/30 space-y-2">
                <Key className="w-8 h-8 text-text-muted mx-auto" />
                <p className="text-sm font-bold text-text-main">No Source CSP credentials registered</p>
                <p className="text-xs text-text-muted">
                  Click <span className="font-bold text-emerald-500">&quot;+ New Credential&quot;</span> above to register an ephemeral credential card.
                </p>
              </div>
            ) : (
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
                {savedCredProfiles.map((cred) => {
                  const isActive = cred.id === selectedCredentialProfile || cred.name === selectedCredentialProfile;
                  return (
                    <div
                      key={cred.id}
                      onClick={() => {
                        setSelectedCredentialProfile(cred.id);
                        setScanCsp(cred.csp);
                        setScanRegion(cred.region);
                        setScanAccessKey(cred.accessKey);
                        setScanSecretKey((cred as any).secretKey || '');
                        setScanTenantId((cred as any).tenantId || '');
                        setScanSubId((cred as any).subId || '');
                        setScanS3AccessKey((cred as any).s3AccessKey || '');
                        setScanS3SecretKey((cred as any).s3SecretKey || '');
                      }}
                      className={`p-4 rounded-xl border transition cursor-pointer relative ${
                        isActive
                          ? 'border-emerald-500 bg-emerald-500/10 shadow-sm'
                          : 'border-border-main bg-bg-main/40 hover:border-emerald-500/40'
                      }`}
                    >
                      <div className="flex items-center justify-between mb-2">
                        <span className="font-extrabold text-sm text-text-main">{cred.name}</span>
                        {isActive ? (
                          <span className="w-2.5 h-2.5 rounded-full bg-emerald-500 animate-pulse" />
                        ) : (
                          <span className="w-2.5 h-2.5 rounded-full bg-text-muted/40" />
                        )}
                      </div>
                      <p className="text-xs text-text-muted mb-3 font-mono">
                        {cred.csp.toUpperCase()} • {cred.region}
                      </p>
                      <div className="flex items-center justify-between text-xs">
                        <span className="px-2 py-0.5 rounded bg-emerald-500/20 text-emerald-400 font-mono font-bold">
                          {cred.accessKey ? (cred.accessKey.length > 12 ? `${cred.accessKey.substring(0, 8)}...` : cred.accessKey) : 'Credential Set'}
                        </span>
                        {isActive ? (
                          <span className="font-bold text-emerald-500 flex items-center gap-1">
                            <CheckCircle2 className="w-3.5 h-3.5" /> Active Credential
                          </span>
                        ) : (
                          <span className="text-text-muted hover:text-emerald-400">Use Credential</span>
                        )}
                      </div>
                    </div>
                  );
                })}
              </div>
            )}
          </div>

          {/* ══════════════════════════════════════════════════
              SECTION 2 — Object Storage Buckets Table & Select Modal
          ══════════════════════════════════════════════════ */}
          <div className="glass-panel p-6 rounded-2xl space-y-5">
            <div className="flex items-center justify-between border-b border-border-main pb-4">
              <div className="flex items-center space-x-2">
                <h3 className="text-base font-bold text-text-main flex items-center gap-2">
                  <Database className="w-5 h-5 text-emerald-400" />
                  Object Storage Buckets
                  {selectedCredentialProfile && (
                    <span className="text-sm font-mono text-text-muted">
                      — [{savedCredProfiles.find(c => c.id === selectedCredentialProfile || c.name === selectedCredentialProfile)?.name || selectedCredentialProfile}]
                    </span>
                  )}
                </h3>
              </div>
              <div className="flex items-center gap-2 text-xs font-bold">
                <button
                  onClick={() => {
                    setIsAddModalOpen(true);
                    const activeCred = savedCredProfiles.find((c) => c.id === selectedCredentialProfile || c.name === selectedCredentialProfile);
                    const keyToScan = activeCred ? activeCred.accessKey : (scanS3AccessKey || scanAccessKey);
                    const secretToScan = activeCred ? (activeCred as any).secretKey : (scanS3SecretKey || scanSecretKey);
                    if (keyToScan && secretToScan) {
                      handleScanSourceBuckets();
                    }
                  }}
                  className="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white rounded-xl text-xs font-bold flex items-center gap-2 transition cursor-pointer shadow-md shadow-emerald-500/20"
                >
                  <Plus className="w-4 h-4" />
                  <span>Select Object Storage</span>
                </button>
              </div>
            </div>

            {/* Registered Buckets Table */}
            <div className="overflow-x-auto border border-border-main rounded-xl">
              <table className="w-full text-xs text-left">
                <thead className="bg-bg-main/50 text-text-muted font-bold border-b border-border-main">
                  <tr>
                    <th className="p-3 text-center w-12">#</th>
                    <th className="p-3">Bucket Name</th>
                    <th className="p-3">Region / CSP</th>
                    <th className="p-3">Total Size</th>
                    <th className="p-3">Object Count</th>
                    <th className="p-3">Storage Class</th>
                    <th className="p-3">Flags</th>
                    <th className="p-3">Status</th>
                    <th className="p-3 text-right">Actions</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-border-main font-mono">
                  {sourceBuckets.map((bucket, i) => (
                    <tr key={i} className="hover:bg-bg-main/30">
                      <td className="p-3 text-center text-text-muted">{i + 1}</td>
                      <td className="p-3 font-bold text-emerald-600 dark:text-emerald-400">{bucket.bucketName}</td>
                      <td className="p-3 text-text-muted">{scanRegion} ({scanCsp.toUpperCase()})</td>
                      <td className="p-3 font-bold">{formatBytes(bucket.totalSizeBytes)}</td>
                      <td className="p-3">{(bucket.objectCount ?? 0).toLocaleString()}</td>
                      <td className="p-3">
                        <span className="px-2 py-0.5 rounded text-[11px] font-bold uppercase bg-bg-input border border-border-main text-text-main">
                          {bucket.accessFrequency}
                        </span>
                      </td>
                      <td className="p-3 space-x-1">
                        {bucket.versioningEnabled && <span className="px-1.5 py-0.5 bg-emerald-500/10 text-emerald-400 rounded text-[10px]">Ver</span>}
                        {bucket.encryptionEnabled && <span className="px-1.5 py-0.5 bg-teal-500/10 text-teal-400 rounded text-[10px]">Enc</span>}
                      </td>
                      <td className="p-3 space-x-1.5">
                        <span className="px-2 py-0.5 bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 rounded font-bold">
                          Selected
                        </span>
                      </td>
                      <td className="p-3 text-right">
                        <button
                          onClick={() => removeSourceBucket(i)}
                          className="p-1 text-text-muted hover:text-red-500 transition cursor-pointer"
                          title="Remove bucket"
                        >
                          <Trash2 className="w-4 h-4" />
                        </button>
                      </td>
                    </tr>
                  ))}
                  {sourceBuckets.length === 0 && (
                    <tr>
                      <td colSpan={9} className="p-8 text-center text-text-muted">
                        No object storage buckets selected yet. Click <span className="font-bold text-emerald-500">&quot;+ Select Object Storage&quot;</span> above to choose buckets.
                      </td>
                    </tr>
                  )}
                </tbody>
              </table>
            </div>

            {/* Table Bottom Status */}
            <div className="flex items-center justify-between pt-1">
              <span className="text-xs text-text-muted font-mono">
                {sourceBuckets.length} Bucket(s) Selected for Migration
              </span>
            </div>
          </div>

          {/* Modal Popup: Select Object Storage Buckets */}
          {isAddModalOpen && (
            <div className="fixed inset-0 z-50 bg-black/75 backdrop-blur-sm flex items-center justify-center p-4">
              <div className="bg-bg-panel border border-border-main rounded-2xl max-w-6xl w-full p-6 shadow-2xl space-y-5 animate-fade-in text-xs max-h-[90vh] overflow-y-auto">
                {/* Modal Header */}
                <div className="flex items-center justify-between border-b border-border-main pb-3">
                  <div className="flex items-center space-x-2">
                    <Database className="w-5 h-5 text-emerald-500" />
                    <h3 className="text-base font-extrabold text-text-main">
                      Select Object Storage Buckets
                    </h3>
                  </div>
                  <button
                    onClick={() => setIsAddModalOpen(false)}
                    className="p-1 text-text-muted hover:text-text-main transition cursor-pointer"
                  >
                    <X className="w-5 h-5" />
                  </button>
                </div>

                {/* Step 1: Credential Profile & CSP Region Settings (Compact Layout) */}
                <div className="p-4 rounded-xl border border-border-main bg-bg-main/40 space-y-4">
                  <div className="flex items-center justify-between">
                    <h4 className="font-extrabold text-text-main flex items-center space-x-2">
                      <Key className="w-4 h-4 text-emerald-500" />
                      <span>1. Select Credential Profile &amp; Target Region</span>
                    </h4>
                  </div>

                  {/* Line 1: Credential Profile (Full Width Row) */}
                  <div>
                    <label className="block text-text-main font-bold mb-1">
                      Credential Profile <span className="text-red-500">*</span>
                    </label>
                    {savedCredProfiles.length > 0 ? (
                      <select
                        value={selectedCredentialProfile}
                        onChange={(e) => {
                          const selectedId = e.target.value;
                          setSelectedCredentialProfile(selectedId);
                          const cred = savedCredProfiles.find((c) => c.id === selectedId || c.name === selectedId);
                          if (cred) {
                            setScanCsp(cred.csp);
                            setScanRegion(cred.region);
                            setScanAccessKey(cred.accessKey);
                            setScanSecretKey((cred as any).secretKey || '');
                          }
                        }}
                        className="w-full px-3.5 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-bold text-xs focus:outline-none focus:border-emerald-500"
                      >
                        {savedCredProfiles.map((c) => (
                          <option key={c.id} value={c.id}>
                            {c.name} ({c.csp.toUpperCase()} • {c.region})
                          </option>
                        ))}
                      </select>
                    ) : (
                      <div className="p-2.5 bg-bg-input/50 border border-border-main rounded-xl text-text-muted text-xs font-mono">
                        No credential profile registered yet. Please click &quot;+ New Credential&quot; on the main screen first.
                      </div>
                    )}
                  </div>

                  {/* Line 2: Source CSP and Region (Side-by-Side 2-Column Row) */}
                  <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div>
                      <label className="block text-text-main font-bold mb-1">Source CSP</label>
                      <select
                        value={scanCsp}
                        onChange={(e) => handleCspChange(e.target.value)}
                        className="w-full px-3.5 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-bold text-xs focus:outline-none focus:border-emerald-500"
                      >
                        {getCspList().map((c) => (
                          <option key={c.id} value={c.id}>
                            {c.name}
                          </option>
                        ))}
                      </select>
                    </div>
                    <div>
                      <label className="block text-text-main font-bold mb-1">Region</label>
                      <select
                        value={scanRegion}
                        onChange={(e) => setScanRegion(e.target.value)}
                        className="w-full px-3.5 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-bold text-xs focus:outline-none focus:border-emerald-500"
                      >
                        {getRegionsForCsp(scanCsp).map((r) => (
                          <option key={r.id} value={r.id}>
                            {r.name}
                          </option>
                        ))}
                      </select>
                    </div>
                  </div>

                  <div className="flex justify-end pt-1">
                    <button
                      onClick={handleScanSourceBuckets}
                      disabled={isScanning}
                      className="px-5 py-2 bg-emerald-600 hover:bg-emerald-700 disabled:opacity-50 text-white font-bold text-xs rounded-xl transition flex items-center space-x-2 cursor-pointer shadow-md shadow-emerald-500/20"
                    >
                      {isScanning ? <RefreshCw className="w-3.5 h-3.5 animate-spin" /> : <Search className="w-3.5 h-3.5" />}
                      <span>Scan Cloud Account Buckets</span>
                    </button>
                  </div>
                </div>

                {/* Step 2: Discovered Buckets Selection Table */}
                <div className="p-4 rounded-xl border border-border-main bg-bg-main/40 space-y-3">
                  <div className="flex items-center justify-between">
                    <h4 className="font-extrabold text-text-main flex items-center space-x-2">
                      <CheckCircle2 className="w-4 h-4 text-emerald-500" />
                      <span>2. Discovered Cloud Buckets ({scannedBucketNames.length})</span>
                    </h4>
                    {isScanning && <span className="text-xs text-emerald-400 font-bold flex items-center gap-1.5"><RefreshCw className="w-3.5 h-3.5 animate-spin" /> Scanning buckets...</span>}
                  </div>

                  <div className="overflow-x-auto border border-border-main rounded-xl">
                    <table className="w-full text-xs text-left">
                      <thead className="bg-bg-main/50 text-text-muted font-bold border-b border-border-main">
                        <tr>
                          <th className="p-2.5 text-center w-10">Select</th>
                          <th className="p-2.5 text-center w-10">#</th>
                          <th className="p-2.5">Bucket Name</th>
                          <th className="p-2.5">Region / CSP</th>
                          <th className="p-2.5">Total Size</th>
                          <th className="p-2.5">Object Count</th>
                          <th className="p-2.5 text-right">Status</th>
                        </tr>
                      </thead>
                      <tbody className="divide-y divide-border-main font-mono">
                        {scannedBucketNames.length > 0 ? (
                          [...scannedBucketNames].sort((a, b) => a.localeCompare(b)).map((bName, idx) => {
                            const isAlreadyRegistered = sourceBuckets.some((b) => b.bucketName === bName);
                            const isChecked = selectedBucketNames.includes(bName);
                            const bInfo = scannedBuckets.find((b: any) => (typeof b === 'string' ? b : b.bucketName) === bName) || {};
                            const displayRegion = bInfo.region || scanRegion;
                            const displaySize = formatBytes(bInfo.sizeBytes);
                            const displayCount = bInfo.objectCount !== undefined ? bInfo.objectCount.toLocaleString() : '0';

                            return (
                              <tr
                                key={idx}
                                className={`transition ${
                                  isAlreadyRegistered
                                    ? 'bg-bg-main/20 opacity-70'
                                    : isChecked
                                    ? 'bg-emerald-500/10 cursor-pointer'
                                    : 'hover:bg-bg-main/30 cursor-pointer'
                                }`}
                                onClick={() => {
                                  if (isAlreadyRegistered) return;
                                  if (isChecked) {
                                    setSelectedBucketNames((prev) => prev.filter((n) => n !== bName));
                                  } else {
                                    setSelectedBucketNames((prev) => [...prev, bName]);
                                  }
                                }}
                              >
                                <td className="p-2.5 text-center">
                                  {isAlreadyRegistered ? (
                                    <span className="text-text-muted">—</span>
                                  ) : (
                                    <input
                                      type="checkbox"
                                      checked={isChecked}
                                      onChange={() => {}}
                                      className="accent-emerald-500 cursor-pointer"
                                    />
                                  )}
                                </td>
                                <td className="p-2.5 text-center text-text-muted">{idx + 1}</td>
                                <td className="p-2.5 font-bold text-emerald-400">{bName}</td>
                                <td className="p-2.5 text-text-muted">{displayRegion} ({scanCsp.toUpperCase()})</td>
                                <td className="p-2.5 font-bold">{displaySize}</td>
                                <td className="p-2.5">{displayCount}</td>
                                <td className="p-2.5 text-right">
                                  {isAlreadyRegistered ? (
                                    <span className="px-2.5 py-0.5 rounded text-[10px] font-mono font-bold bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 inline-flex items-center space-x-1">
                                      <CheckCircle2 className="w-3 h-3" />
                                      <span>Selected</span>
                                    </span>
                                  ) : (
                                    <span className="text-text-muted text-[10px]">Ready to Select</span>
                                  )}
                                </td>
                              </tr>
                            );
                          })
                        ) : (
                          <tr>
                            <td colSpan={7} className="p-6 text-center text-text-muted">
                              {isScanning ? 'Scanning cloud account buckets...' : 'No buckets discovered.'}
                            </td>
                          </tr>
                        )}
                      </tbody>
                    </table>
                  </div>
                </div>

                {/* Modal Footer Actions */}
                <div className="flex justify-end space-x-3 pt-3 border-t border-border-main">
                  <button
                    onClick={() => setIsAddModalOpen(false)}
                    className="px-5 py-2.5 bg-bg-input border border-border-main hover:bg-bg-main text-text-main rounded-xl font-bold transition cursor-pointer"
                  >
                    Cancel
                  </button>
                  <button
                    onClick={() => {
                      const newBucketsToSelect = selectedBucketNames.map((bName) => {
                        const bInfo = scannedBuckets.find((b: any) => (typeof b === 'string' ? b : b.bucketName) === bName) || {};
                        return {
                          bucketName: bName,
                          totalSizeBytes: bInfo.sizeBytes || 0,
                          objectCount: bInfo.objectCount || 0,
                          accessFrequency: 'frequent' as const,
                          versioningEnabled: bInfo.versioningEnabled || false,
                          encryptionEnabled: true,
                          corsEnabled: false,
                          isPublic: false
                        };
                      });

                      setSourceBuckets((prev) => {
                        const existingNames = new Set(prev.map((b) => b.bucketName));
                        const filtered = newBucketsToSelect.filter((b) => !existingNames.has(b.bucketName));
                        return [...prev, ...filtered];
                      });
                      setIsAddModalOpen(false);
                    }}
                    disabled={selectedBucketNames.length === 0}
                    className="px-6 py-2.5 bg-emerald-500 hover:bg-emerald-600 disabled:opacity-50 text-white font-extrabold text-xs rounded-xl transition cursor-pointer shadow-lg shadow-emerald-500/20"
                  >
                    Add Selected Buckets ({selectedBucketNames.length})
                  </button>
                </div>
              </div>
            </div>
          )}

          {/* Modal Popup: Register New CSP Credential Profile */}
          {isRegisterCredModalOpen && (
            <div className="fixed inset-0 z-50 bg-black/75 backdrop-blur-sm flex items-center justify-center p-4">
              <div className="bg-bg-panel border border-border-main rounded-2xl max-w-2xl w-full p-6 shadow-2xl space-y-5 animate-fade-in text-xs max-h-[90vh] overflow-y-auto">
                {/* Modal Header */}
                <div className="flex items-center justify-between border-b border-border-main pb-3">
                  <div className="flex items-center space-x-2">
                    <Lock className="w-5 h-5 text-emerald-500" />
                    <h3 className="text-base font-extrabold text-text-main">
                      Register New Source CSP Credential Profile
                    </h3>
                  </div>
                  <button
                    onClick={() => setIsRegisterCredModalOpen(false)}
                    className="p-1 text-text-muted hover:text-text-main transition cursor-pointer"
                  >
                    <X className="w-5 h-5" />
                  </button>
                </div>

                {/* Row 1: Credential Profile Name */}
                <div>
                  <label className="block text-text-main font-bold mb-1">
                    Credential Profile Name <span className="text-red-500">*</span>
                  </label>
                  <input
                    type="text"
                    value={credProfileName}
                    onChange={(e) => setCredProfileName(e.target.value)}
                    placeholder="e.g. aws-prod-credential"
                    className="w-full px-3.5 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-bold focus:outline-none focus:border-emerald-500"
                  />
                </div>

                {/* Row 2 & Row 3+: Top/Bottom Split (Row 2 = CSP & Region | Row 3+ = Dynamic Credential Inputs) */}
                <div className="p-4 rounded-xl border border-border-main bg-bg-main/40 space-y-4">
                  {renderCspCredentialBlock(
                    credCsp,
                    (newCsp) => {
                      setCredCsp(newCsp);
                      const available = getRegionsForCsp(newCsp);
                      if (available.length > 0) setCredRegion(available[0].id);
                    },
                    credRegion,
                    setCredRegion,
                    credAccessKey, setCredAccessKey,
                    credSecretKey, setCredSecretKey,
                    credTenantId, setCredTenantId,
                    credSubId, setCredSubId
                  )}
                </div>

                {/* Modal Footer */}
                <div className="flex justify-end space-x-3 pt-3 border-t border-border-main">
                  <button
                    onClick={() => setIsRegisterCredModalOpen(false)}
                    className="px-5 py-2.5 bg-bg-input border border-border-main hover:bg-bg-main text-text-main rounded-xl font-bold transition cursor-pointer"
                  >
                    Cancel
                  </button>
                  <button
                    onClick={() => {
                      if (!credProfileName.trim()) {
                        alert('Please enter a Credential Profile Name.');
                        return;
                      }
                      const newProfile = {
                        id: `cred-${Date.now()}`,
                        name: credProfileName,
                        csp: credCsp,
                        region: credRegion,
                        accessKey: credAccessKey,
                        secretKey: credSecretKey,
                        tenantId: credTenantId,
                        subId: credSubId,
                        s3AccessKey: credS3AccessKey,
                        s3SecretKey: credS3SecretKey
                      };
                      setSavedCredProfiles((prev) => [...prev, newProfile]);
                      setSelectedCredentialProfile(newProfile.id);
                      setScanCsp(credCsp);
                      setScanRegion(credRegion);
                      setScanAccessKey(credAccessKey);
                      setScanSecretKey(credSecretKey);
                      setScanTenantId(credTenantId);
                      setScanSubId(credSubId);
                      setScanS3AccessKey(credS3AccessKey);
                      setScanS3SecretKey(credS3SecretKey);
                      setIsRegisterCredModalOpen(false);
                    }}
                    className="px-6 py-2.5 bg-emerald-500 hover:bg-emerald-600 text-white font-extrabold text-xs rounded-xl transition cursor-pointer shadow-lg shadow-emerald-500/20"
                  >
                    Save Credential Profile
                  </button>
                </div>
              </div>
            </div>
          )}

          {/* ══════════════════════════════════════════════════
              SECTION 3 — Collect & Save Source Model
          ══════════════════════════════════════════════════ */}
          <div className="glass-panel p-6 rounded-2xl space-y-5">
            <div className="flex items-center justify-between border-b border-border-main pb-4">
              <div>
                <h3 className="text-base font-bold text-text-main flex items-center gap-2">
                  <Database className="w-5 h-5 text-emerald-500" />
                  Collect &amp; Save Source Model
                  <span className="text-sm font-mono text-text-muted"> — [Sample] aws-prod-storage</span>
                </h3>
              </div>
              <div className="flex items-center space-x-2 text-xs font-mono">
                <span className="px-2.5 py-1 bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 rounded-md font-bold flex items-center gap-1.5">
                  <CheckCircle2 className="w-3.5 h-3.5" />
                  <span>Buckets Registered</span>
                </span>
                <span className="text-text-muted">&gt;</span>
                <span className={`px-2.5 py-1 rounded-md font-bold flex items-center gap-1.5 ${
                  sourceBuckets.length > 0
                    ? 'bg-emerald-500/10 text-emerald-400 border border-emerald-500/20'
                    : 'bg-bg-input text-text-muted border border-border-main'
                }`}>
                  <span>2 Storage Collected</span>
                </span>
                <span className="text-text-muted">&gt;</span>
                <span className={`px-2.5 py-1 rounded-md font-bold flex items-center gap-1.5 ${
                  savedSourceModelId
                    ? 'bg-emerald-500/10 text-emerald-400 border border-emerald-500/20'
                    : 'bg-bg-input text-text-muted border border-border-main'
                }`}>
                  <span>3 {savedSourceModelId ? `Model Saved (${savedSourceModelId})` : 'Model Saved'}</span>
                </span>
              </div>
            </div>

            {/* Action Buttons Row (Exact Screenshot Alignment) */}
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
              <div>
                <button
                  onClick={handleInspectSelectedBuckets}
                  disabled={isScanning || isInspectLoading}
                  className="w-full py-3 px-6 bg-bg-panel border border-emerald-500 hover:bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 font-extrabold text-sm rounded-xl transition flex items-center justify-center space-x-2 cursor-pointer disabled:opacity-50"
                >
                  {isInspectLoading ? (
                    <RefreshCw className="w-4 h-4 animate-spin text-emerald-500" />
                  ) : (
                    <Play className="w-4 h-4 text-emerald-500 fill-emerald-500/20" />
                  )}
                  <span>{isInspectLoading ? 'Collecting Metadata...' : 'Collect Object Storage Metadata'}</span>
                </button>
                <p className="text-[11px] text-text-muted text-center mt-2">
                  Triggers the MinIO / CSP API to inspect and extract metadata from all registered buckets.
                </p>
              </div>

              <div>
                <button
                  onClick={handleSaveSourceModel}
                  disabled={sourceBuckets.length === 0}
                  className="w-full py-3 px-6 bg-emerald-600 hover:bg-emerald-700 disabled:opacity-50 text-white font-extrabold text-sm rounded-xl transition flex items-center justify-center space-x-2 cursor-pointer shadow-lg shadow-emerald-500/20"
                >
                  <Save className="w-4 h-4" />
                  <span>Save Source Storage Revision</span>
                </button>
                <p className="text-[11px] text-text-muted text-center mt-2">
                  Opens a popup to name, version, and save this collected model to Damselfly.
                </p>
              </div>
            </div>

            {/* Collected Storage Model Preview (Screenshot Alignment) */}
            <div className="space-y-2 pt-2">
              <h4 className="text-sm font-bold text-text-main flex items-center space-x-2">
                <Database className="w-4 h-4 text-emerald-500" />
                <span>Collected Storage Model Preview</span>
              </h4>
              <div className="p-4 bg-bg-main/60 border border-border-main rounded-2xl shadow-inner overflow-hidden">
                <pre className="font-mono text-xs text-text-main dark:text-emerald-400/90 max-h-[48rem] min-h-[28rem] overflow-y-auto leading-relaxed select-all">
                  {JSON.stringify(
                    collectedSourceModel || {
                      description: `Inspected source object storage model for ${scanCsp.toUpperCase()} (${scanRegion})`,
                      sourceCloud: {
                        csp: scanCsp,
                        region: scanRegion
                      },
                      sourceObjectStorages: sourceBuckets.map((b) => ({
                        accessFrequency: b.accessFrequency || 'frequent',
                        bucketName: b.bucketName,
                        corsEnabled: b.corsEnabled || false,
                        corsRule: b.corsRule || [],
                        creationDate: b.creationDate || '',
                        encryptionEnabled: b.encryptionEnabled || false,
                        isPublic: b.isPublic || false,
                        objectCount: b.objectCount || 0,
                        tags: b.tags || {},
                        totalSizeBytes: b.totalSizeBytes || 0,
                        versioningEnabled: b.versioningEnabled || false
                      }))
                    },
                    null,
                    2
                  )}
                </pre>
              </div>
            </div>

            <div className="flex justify-end pt-2 border-t border-border-main/40">
              <button
                onClick={() => setSubTab('recommend')}
                disabled={sourceBuckets.length === 0}
                className="px-6 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white font-bold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center space-x-2 cursor-pointer"
              >
                <span>Next: Proceed to 2. Target Recommendation</span>
                <ArrowRight className="w-4 h-4" />
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Sub-step 2: Target Recommendation & Sizing */}
      {(subTab as string) === 'recommend' && (
        <div className="bg-bg-panel border border-border-main rounded-xl p-6 space-y-6">
          <div className="flex items-center justify-between border-b border-border-main pb-3">
            <div className="flex items-center space-x-2">
              <ShieldCheck className="w-5 h-5 text-emerald-500" />
              <h3 className="text-sm font-extrabold text-text-main">2. Target Cloud Recommendation & Sizing</h3>
            </div>
          </div>

          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-xs">
            <div>
              <label className="block text-text-muted font-medium mb-1">Target Cloud Provider</label>
              <select
                value={desiredCsp}
                onChange={(e) => handleCspChange(e.target.value)}
                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-bold uppercase focus:outline-none focus:border-emerald-500"
              >
                {tumblebugProviders.map((p) => (
                  <option key={p} value={p}>
                    {p.toUpperCase()}
                  </option>
                ))}
              </select>
            </div>

            <div>
              <label className="block text-text-muted font-medium mb-1">Target Cloud Region</label>
              <select
                value={desiredRegion}
                onChange={(e) => setDesiredRegion(e.target.value)}
                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main focus:outline-none focus:border-emerald-500"
              >
                {tumblebugRegions.map((r) => (
                  <option key={r.id} value={r.id}>
                    {r.id} ({r.name})
                  </option>
                ))}
              </select>
            </div>
          </div>

          <button
            onClick={handleGetRecommendation}
            disabled={isRecommending || sourceBuckets.length === 0}
            className="w-full py-3 bg-emerald-600 hover:bg-emerald-700 disabled:opacity-50 text-white font-bold text-xs rounded-xl transition flex items-center justify-center space-x-2 cursor-pointer shadow-lg shadow-emerald-500/20"
          >
            {isRecommending ? <RefreshCw className="w-4 h-4 animate-spin" /> : <Play className="w-4 h-4" />}
            <span>Generate Target Storage Recommendation</span>
          </button>

          {recommendationResult && (
            <div className="p-4 bg-emerald-500/10 border border-emerald-500/20 rounded-xl space-y-2 text-xs">
              <div className="font-bold text-emerald-600 dark:text-emerald-400 flex items-center space-x-2">
                <CheckCircle2 className="w-4 h-4" />
                <span>Target Storage Class Recommendation Result</span>
              </div>
              <div className="text-xs text-text-muted">{recommendationResult.description}</div>
            </div>
          )}

          <div className="flex items-center justify-between pt-2 border-t border-border-main/50">
            <button
              onClick={() => setSubTab('source')}
              className="px-5 py-2.5 bg-bg-input/60 hover:bg-bg-main border border-border-main text-text-main font-bold text-xs rounded-xl transition flex items-center space-x-2 cursor-pointer"
            >
              <ArrowLeft className="w-4 h-4" />
              <span>Back to 1. Source Storage Analysis</span>
            </button>
            <div className="flex items-center space-x-2">
              <button
                onClick={handleSaveTargetModel}
                disabled={!recommendationResult}
                className={`px-4 py-2.5 rounded-xl text-xs font-bold border transition flex items-center space-x-1.5 cursor-pointer ${
                  savedTargetModelId
                    ? 'border-emerald-500 bg-emerald-500/10 text-emerald-400'
                    : 'border-border-main bg-bg-input text-text-muted hover:border-emerald-500/40'
                }`}
              >
                <Database className="w-3.5 h-3.5" />
                <span>{savedTargetModelId ? `Target Model Saved (${savedTargetModelId})` : 'Save Target Model (Damselfly)'}</span>
              </button>
              <button
                onClick={() => setSubTab('provision')}
                disabled={!recommendationResult}
                className="px-6 py-2.5 bg-emerald-600 hover:bg-emerald-700 disabled:opacity-50 text-white font-bold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center space-x-2 cursor-pointer"
              >
                <span>Next: Proceed to 3. Bucket Provisioning</span>
                <ArrowRight className="w-4 h-4" />
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Sub-step 3: Bucket Provisioning & Migrated Storage List */}
      {(subTab as string) === 'provision' && (
        <div className="grid grid-cols-1 lg:grid-cols-12 gap-6">
          <div className="lg:col-span-6 space-y-6">
            <div className="bg-bg-panel border border-border-main rounded-xl p-6 space-y-4">
              <h3 className="text-sm font-extrabold text-text-main border-b border-border-main pb-3">
                3. Provision Target Cloud Object Storage
              </h3>

              <div className="space-y-4 text-xs">
                <div>
                  <label className="block text-text-muted font-medium mb-1">Late Binding Prefix Seed (nameSeed)</label>
                  <input
                    type="text"
                    value={nameSeed}
                    onChange={(e) => setNameSeed(e.target.value)}
                    placeholder="e.g. prod"
                    className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500"
                  />
                  <span className="text-xs text-text-muted mt-1 block">
                    Bucket names will be prefixed on creation: <code className="text-emerald-500">{nameSeed ? `${nameSeed}-` : ''}bucketName</code>
                  </span>
                </div>

                <button
                  onClick={handleMigrateStorage}
                  disabled={isDeploying}
                  className="w-full py-3 bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-700 hover:to-teal-700 text-white font-bold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center justify-center space-x-2 cursor-pointer"
                >
                  {isDeploying ? <RefreshCw className="w-4 h-4 animate-spin" /> : <ArrowRight className="w-4 h-4" />}
                  <span>Provision Target Cloud Object Storage</span>
                </button>

                {deploymentLog.length > 0 && (
                  <div className="p-3 bg-slate-950 border border-border-main rounded-lg font-mono text-xs text-emerald-400 space-y-1 max-h-40 overflow-y-auto">
                    {deploymentLog.map((log, idx) => (
                      <div key={idx}>&gt; {log}</div>
                    ))}
                  </div>
                )}
              </div>
            </div>
          </div>

          <div className="lg:col-span-6 space-y-6">
            <div className="bg-bg-panel border border-border-main rounded-xl p-6 space-y-4">
              <div className="flex items-center justify-between border-b border-border-main pb-3">
                <div className="flex items-center space-x-2">
                  <Server className="w-4 h-4 text-emerald-500" />
                  <h3 className="text-sm font-extrabold text-text-main">Migrated Storage Buckets ({namespaceId})</h3>
                </div>
                <button
                  onClick={loadMigratedStorages}
                  className="p-1.5 text-text-muted hover:text-emerald-500 transition cursor-pointer"
                  title="Refresh list"
                >
                  <RefreshCw className={`w-3.5 h-3.5 ${isLoadingMigrated ? 'animate-spin' : ''}`} />
                </button>
              </div>

              {migratedStorages.length === 0 ? (
                <div className="py-6 text-center text-text-muted text-xs">
                  No active object storages found in namespace <code className="text-emerald-500">{namespaceId}</code>.
                </div>
              ) : (
                <div className="space-y-2 max-h-60 overflow-y-auto font-mono text-xs">
                  {migratedStorages.map((storage: any, idx: number) => (
                    <div key={idx} className="p-3 bg-bg-main/40 border border-border-main rounded-lg flex items-center justify-between">
                      <div>
                        <div className="font-bold text-emerald-600 dark:text-emerald-400">{storage.name || storage.id}</div>
                        <div className="text-xs text-text-muted">Connection: {storage.connectionName}</div>
                      </div>
                      <button
                        onClick={() => handleDeleteStorage(storage.id || storage.name)}
                        className="p-1 text-text-muted hover:text-red-500 transition cursor-pointer"
                        title="Delete storage"
                      >
                        <Trash2 className="w-4 h-4" />
                      </button>
                    </div>
                  ))}
                </div>
              )}
            </div>
          </div>

          <div className="lg:col-span-12">
            <div className="flex items-center justify-between pt-4 border-t border-border-main/50 mt-2">
              <button
                onClick={() => setSubTab('recommend')}
                className="px-5 py-2.5 bg-bg-input/60 hover:bg-bg-main border border-border-main text-text-main font-bold text-xs rounded-xl transition flex items-center space-x-2 cursor-pointer"
              >
                <ArrowLeft className="w-4 h-4" />
                <span>Back to 2. Target Recommendation</span>
              </button>

              <div className="flex items-center space-x-2 px-4 py-2 bg-emerald-500/10 border border-emerald-500/30 text-emerald-400 rounded-xl text-xs font-bold font-mono">
                <CheckCircle2 className="w-4 h-4 text-emerald-400" />
                <span>Final Step: 3. Bucket Provisioning &amp; Monitoring</span>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
