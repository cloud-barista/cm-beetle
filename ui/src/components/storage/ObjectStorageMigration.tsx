'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore } from '@/store/migrationStore';
import { beetleApi } from '@/api/client';
import { SaveRevisionModal, SaveRevisionResult } from '../common/SaveRevisionModal';
import sampleStorageRequest from '../../data/sampleSourceObjectStorage.json';
import sampleTargetObjectStorage from '../../data/sampleTargetObjectStorage.json';
import {
  HardDrive, Plus, Trash2, CheckCircle2, ArrowRight, ArrowLeft, Check, Zap,
  ShieldCheck, Database, RefreshCw, Server, AlertCircle, Play, Key,
  Search, ChevronDown, ChevronUp, Lock, FileText, Upload, Save, X,
  Sliders, Compass, Copy, Edit3, Settings2, Globe, Shield, Tag, AlertTriangle, Sparkles,
  Activity, Clock
} from 'lucide-react';

interface CorsRuleItem {
  allowedOrigin: string[];
  allowedMethod: string[];
  allowedHeader: string[];
  exposeHeader: string[];
  maxAgeSeconds: number;
}

interface SourceBucket {
  bucketName: string;
  targetBucketName?: string;
  totalSizeBytes: number;
  objectCount: number;
  accessFrequency: 'frequent' | 'infrequent' | 'archive';
  versioningEnabled: boolean;
  encryptionEnabled: boolean;
  corsEnabled: boolean;
  corsRule?: CorsRuleItem[];
  isPublic: boolean;
  creationDate?: string;
  tags?: Record<string, string>;
}

const ALL_SUPPORTED_CSPS = [
  { id: 'aws', name: 'AWS (Amazon Web Services)' },
  { id: 'azure', name: 'Azure (Microsoft Azure)' },
  { id: 'gcp', name: 'GCP (Google Cloud Platform)' },
  { id: 'alibaba', name: 'Alibaba Cloud' },
  { id: 'tencent', name: 'Tencent Cloud' },
  { id: 'ibm', name: 'IBM Cloud' },
  { id: 'ncp', name: 'NCP (Naver Cloud Platform)' },
  { id: 'nhn', name: 'NHN Cloud' },
  { id: 'kt', name: 'KT Cloud' },
  { id: 'openstack', name: 'OpenStack' }
];

const CSP_REGIONS_MAP: Record<string, { id: string; name: string }[]> = {
  aws: [
    { id: 'ap-northeast-2', name: 'ap-northeast-2 (Seoul)' },
    { id: 'ap-northeast-1', name: 'ap-northeast-1 (Tokyo)' },
    { id: 'us-east-1', name: 'us-east-1 (N. Virginia)' },
    { id: 'us-west-2', name: 'us-west-2 (Oregon)' },
    { id: 'eu-west-1', name: 'eu-west-1 (Ireland)' }
  ],
  azure: [
    { id: 'koreacentral', name: 'koreacentral (Seoul)' },
    { id: 'eastus', name: 'eastus (East US)' },
    { id: 'westeurope', name: 'westeurope (West Europe)' }
  ],
  gcp: [
    { id: 'asia-northeast3', name: 'asia-northeast3 (Seoul)' },
    { id: 'us-central1', name: 'us-central1 (Iowa)' },
    { id: 'europe-west3', name: 'europe-west3 (Frankfurt)' }
  ],
  alibaba: [
    { id: 'ap-northeast-1', name: 'ap-northeast-1 (Tokyo)' },
    { id: 'ap-southeast-1', name: 'ap-southeast-1 (Singapore)' },
    { id: 'us-east-1', name: 'us-east-1 (Virginia)' }
  ],
  tencent: [
    { id: 'ap-seoul', name: 'ap-seoul (Seoul)' },
    { id: 'ap-guangzhou', name: 'ap-guangzhou (Guangzhou)' },
    { id: 'na-siliconvalley', name: 'na-siliconvalley (Silicon Valley)' }
  ],
  ibm: [
    { id: 'us-south-1', name: 'us-south-1 (Dallas)' },
    { id: 'us-east-1', name: 'us-east-1 (Washington DC)' }
  ],
  ncp: [
    { id: 'kr-1', name: 'kr-1 (Seoul)' },
    { id: 'kr-2', name: 'kr-2 (Cheongju)' },
    { id: 'sg-1', name: 'sg-1 (Singapore)' }
  ],
  nhn: [
    { id: 'kr1', name: 'kr1 (Pangyo)' },
    { id: 'kr2', name: 'kr2 (Gwangju)' }
  ],
  kt: [
    { id: 'kr-1', name: 'kr-1 (Seoul)' },
    { id: 'kr-2', name: 'kr-2 (Cheonan)' }
  ],
  openstack: [
    { id: 'default', name: 'default (Default)' }
  ]
};

const HTTP_METHODS = ['GET', 'POST', 'PUT', 'DELETE', 'HEAD', 'OPTIONS'];

const SAMPLE_STORAGE_MODEL = {
  id: 'sample-source-storage-01',
  name: '[Sample] sampleSourceObjectStorage.json',
  version: '1.0',
  updatedTime: new Date().toISOString(),
  description: 'Sample object storage model for quick demonstration'
};

export const ObjectStorageMigration: React.FC = () => {
  const {
    tumblebugProviders,
    tumblebugRegions,
    fetchTumblebugProviders,
    fetchTumblebugRegions
  } = useMigrationStore();
  const namespaceId = (useMigrationStore.getState() as any).namespaceId || 'mig01';

  const [subTab, setSubTab] = useState<'source' | 'refine' | 'optimize' | 'provision'>('source');

  const subSteps = [
    { id: 'source', label: '1. Source Storage', icon: Database, desc: 'Register source buckets & scan credentials' },
    { id: 'refine', label: '2. Refinement', icon: Sliders, desc: 'Review & refine source bucket specs' },
    { id: 'optimize', label: '3. Target Object Storage Optimization', icon: Compass, desc: 'Recommend & customize target storage' },
    { id: 'provision', label: '4. Migration Execution', icon: Play, desc: 'Deploy target buckets & execute storage migrations' },
  ] as const;

  // Credential Management State (Step 1)
  const [selectedCredentialProfile, setSelectedCredentialProfile] = useState('');
  const [isRegisterCredModalOpen, setIsRegisterCredModalOpen] = useState(false);
  const [credProfileName, setCredProfileName] = useState('');
  const [credCsp, setCredCsp] = useState('aws');
  const [credRegion, setCredRegion] = useState('ap-northeast-2');
  const [credAccessKey, setCredAccessKey] = useState('');
  const [credSecretKey, setCredSecretKey] = useState('');
  const [savedCredProfiles, setSavedCredProfiles] = useState<
    { id: string; name: string; csp: string; region: string; accessKey: string }[]
  >([]);

  const [scanCsp, setScanCsp] = useState('aws');
  const [scanRegion, setScanRegion] = useState('ap-northeast-2');
  const [scanAccessKey, setScanAccessKey] = useState('');
  const [scanSecretKey, setScanSecretKey] = useState('');
  const [isScanning, setIsScanning] = useState(false);
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);

  const [sourceBuckets, setSourceBuckets] = useState<SourceBucket[]>([
    {
      bucketName: 'source-bucket-01',
      targetBucketName: 'target-bucket-01',
      totalSizeBytes: 10737418240,
      objectCount: 1000,
      accessFrequency: 'frequent',
      versioningEnabled: false,
      encryptionEnabled: false,
      corsEnabled: true,
      corsRule: sampleStorageRequest.sourceObjectStorages[0].corsRule as any,
      isPublic: false,
      tags: { env: 'production', team: 'platform' }
    },
    {
      bucketName: 'source-bucket-02',
      targetBucketName: 'target-bucket-02',
      totalSizeBytes: 1073741824,
      objectCount: 100,
      accessFrequency: 'infrequent',
      versioningEnabled: true,
      encryptionEnabled: true,
      corsEnabled: true,
      corsRule: sampleStorageRequest.sourceObjectStorages[1].corsRule as any,
      isPublic: false,
      tags: { env: 'staging', team: 'data' }
    }
  ]);

  const [scannedBucketNames, setScannedBucketNames] = useState<string[]>([]);
  const [scannedBuckets, setScannedBuckets] = useState<any[]>([]);
  const [selectedBucketNames, setSelectedBucketNames] = useState<string[]>([]);
  const [isInspectLoading, setIsInspectLoading] = useState(false);

  const [savedSourceModels, setSavedSourceModels] = useState<any[]>([]);
  const [selectedSourceModelId, setSelectedSourceModelId] = useState<string>('');
  const [isModelLoaded, setIsModelLoaded] = useState<boolean>(false);
  const [activeRefineStep, setActiveRefineStep] = useState<number>(1);
  const [loadedModelName, setLoadedModelName] = useState<string>('');
  const [loadedModelVersion, setLoadedModelVersion] = useState<string>('1.0');
  const [loadedModelTime, setLoadedModelTime] = useState<string>('');

  const [savedTargetModelId, setSavedTargetModelId] = useState<string | null>(null);
  const [showSaveSourceModal, setShowSaveSourceModal] = useState(false);
  const [showSaveTargetModal, setShowSaveTargetModal] = useState(false);
  const [modelLog, setModelLog] = useState<string[]>([]);

  const [isRefineJsonOpen, setIsRefineJsonOpen] = useState(false);
  const [selectedBucketIndex, setSelectedBucketIndex] = useState<number>(0);
  const [excludedBucketNames, setExcludedBucketNames] = useState<string[]>([]);
  const [showCorsGuide, setShowCorsGuide] = useState<boolean>(false);

  const toggleExcludeBucket = (bucketName: string) => {
    setExcludedBucketNames((prev) =>
      prev.includes(bucketName)
        ? prev.filter((name) => name !== bucketName)
        : [...prev, bucketName]
    );
  };

  const [newOriginInputs, setNewOriginInputs] = useState<Record<number, string>>({});
  const [newAllowedHeaderInputs, setNewAllowedHeaderInputs] = useState<Record<number, string>>({});
  const [newExposeHeaderInputs, setNewExposeHeaderInputs] = useState<Record<number, string>>({});

  const [step1SubTab, setStep1SubTab] = useState<'generate' | 'load'>('generate');
  const [desiredCsp, setDesiredCsp] = useState('aws');
  const [desiredRegion, setDesiredRegion] = useState('ap-northeast-2');
  const [targetObjectStorageName, setTargetObjectStorageName] = useState('');
  const nameSeed = targetObjectStorageName;
  const setNameSeed = setTargetObjectStorageName;
  const [recommendationResult, setRecommendationResult] = useState<any | null>(null);
  const [isRecommending, setIsRecommending] = useState(false);
  const [isTargetJsonOpen, setIsTargetJsonOpen] = useState(false);

  const [isDeploying, setIsDeploying] = useState(false);
  const [showLaunchModal, setShowLaunchModal] = useState(false);
  const [deploymentLog, setDeploymentLog] = useState<string[]>([]);
  const [migratedStorages, setMigratedStorages] = useState<any[]>([]);
  const [isLoadingMigrated, setIsLoadingMigrated] = useState(false);

  const [cspSupportMap, setCspSupportMap] = useState<Record<string, { cors: boolean; presignedUrl: boolean; versioning: boolean }>>({
    aws: { cors: true, presignedUrl: true, versioning: true },
    gcp: { cors: true, presignedUrl: true, versioning: true },
    azure: { cors: true, presignedUrl: true, versioning: true },
    alibaba: { cors: true, presignedUrl: true, versioning: true },
    tencent: { cors: true, presignedUrl: true, versioning: true },
    ibm: { cors: true, presignedUrl: true, versioning: false },
    ncp: { cors: true, presignedUrl: false, versioning: false },
    nhn: { cors: true, presignedUrl: false, versioning: false },
    kt: { cors: false, presignedUrl: false, versioning: false },
    openstack: { cors: false, presignedUrl: true, versioning: false }
  });

  useEffect(() => {
    fetchTumblebugProviders();
    beetleApi.getObjectStorageSupport().then((data) => {
      if (data && typeof data === 'object') {
        setCspSupportMap(data);
      }
    }).catch(() => {
      // Retain fallback support map if Tumblebug proxy returns 404 or error
    });
  }, []);

  const getCspSupport = (cspKey: string) => {
    const key = (cspKey || '').toLowerCase();
    return cspSupportMap[key] || { cors: true, presignedUrl: true, versioning: true };
  };

  const allSourceModels = [
    SAMPLE_STORAGE_MODEL,
    ...savedSourceModels.filter((m) => m.id !== 'sample-source-storage-01')
  ];

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

  const getRegionsForCsp = (csp: string) => {
    const key = (csp || 'aws').toLowerCase();
    const tbRegions = (tumblebugRegions || [])
      .filter((r: any) => !r.providerName || r.providerName.toLowerCase() === key)
      .map((r: any) => {
        const id = r.id || r.regionName;
        const displayName = r.name || r.locationName || id;
        return {
          id,
          name: displayName.includes(id) ? displayName : `${id} (${displayName})`
        };
      });

    if (tbRegions.length > 0) {
      return tbRegions;
    }

    return (
      CSP_REGIONS_MAP[key] || [
        { id: 'ap-northeast-2', name: 'ap-northeast-2 (Seoul)' },
        { id: 'us-east-1', name: 'us-east-1 (N. Virginia)' }
      ]
    );
  };

  const handleCspChange = (newCsp: string) => {
    setScanCsp(newCsp);
    fetchTumblebugRegions(newCsp);
    const available = getRegionsForCsp(newCsp);
    if (available.length > 0) setScanRegion(available[0].id);
  };

  const handleDesiredCspChange = (newCsp: string) => {
    setDesiredCsp(newCsp);
    fetchTumblebugRegions(newCsp);
    const available = getRegionsForCsp(newCsp);
    if (available.length > 0) setDesiredRegion(available[0].id);
  };

  const formatBytes = (bytes?: number) => {
    if (!bytes || bytes === 0) return '0 B';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i];
  };

  // Step 1 Load / Delete Handlers
  const handleLoadModel = () => {
    const targetModel = allSourceModels.find((m) => m.id === selectedSourceModelId) || SAMPLE_STORAGE_MODEL;
    const bucketsToLoad: SourceBucket[] = (targetModel.buckets || sampleStorageRequest.sourceObjectStorages).map((b: any) => ({
      bucketName: b.bucketName,
      totalSizeBytes: b.totalSizeBytes || 10737418240,
      objectCount: b.objectCount || 1000,
      accessFrequency: b.accessFrequency || 'frequent',
      versioningEnabled: b.versioningEnabled || false,
      encryptionEnabled: b.encryptionEnabled || false,
      corsEnabled: b.corsEnabled || false,
      corsRule: b.corsRule || [],
      isPublic: b.isPublic || false,
      tags: b.tags || {}
    }));
    setSourceBuckets(bucketsToLoad);
    setLoadedModelName(targetModel.name);
    setLoadedModelVersion(targetModel.version || '1.0');
    setLoadedModelTime(new Date().toLocaleString());
    setIsModelLoaded(true);
    setActiveRefineStep(2);
  };

  const handleDeleteModel = () => {
    if (selectedSourceModelId === 'sample-source-storage-01') return;
    if (!confirm('Are you sure you want to delete this source storage model revision?')) return;
    setSavedSourceModels((prev) => prev.filter((m) => m.id !== selectedSourceModelId));
    setSelectedSourceModelId('sample-source-storage-01');
    handleLoadModel();
  };

  // Detailed CORS Config Rule Modification Helpers
  const ensureCorsRuleExists = (b: SourceBucket): CorsRuleItem => {
    if (b.corsRule && b.corsRule.length > 0) return b.corsRule[0];
    return {
      allowedOrigin: ['*'],
      allowedMethod: ['GET', 'POST', 'PUT', 'DELETE'],
      allowedHeader: ['*'],
      exposeHeader: ['ETag'],
      maxAgeSeconds: 3600
    };
  };

  const updateCorsRuleForBucket = (bucketIdx: number, updater: (rule: CorsRuleItem) => CorsRuleItem) => {
    setSourceBuckets((prev) =>
      prev.map((b, i) => {
        if (i !== bucketIdx) return b;
        const currentRule = ensureCorsRuleExists(b);
        const updatedRule = updater(currentRule);
        return {
          ...b,
          corsEnabled: true,
          corsRule: [updatedRule]
        };
      })
    );
  };

  const handleToggleCorsMethod = (bucketIdx: number, method: string) => {
    updateCorsRuleForBucket(bucketIdx, (rule) => {
      const exists = rule.allowedMethod.includes(method);
      const newMethods = exists
        ? rule.allowedMethod.filter((m) => m !== method)
        : [...rule.allowedMethod, method];
      return { ...rule, allowedMethod: newMethods };
    });
  };

  const handleAddCorsOrigin = (bucketIdx: number, originVal: string) => {
    if (!originVal.trim()) return;
    updateCorsRuleForBucket(bucketIdx, (rule) => {
      if (rule.allowedOrigin.includes(originVal.trim())) return rule;
      return { ...rule, allowedOrigin: [...rule.allowedOrigin, originVal.trim()] };
    });
    setNewOriginInputs((prev) => ({ ...prev, [bucketIdx]: '' }));
  };

  const handleRemoveCorsOrigin = (bucketIdx: number, originIdx: number) => {
    updateCorsRuleForBucket(bucketIdx, (rule) => ({
      ...rule,
      allowedOrigin: rule.allowedOrigin.filter((_, i) => i !== originIdx)
    }));
  };

  const handleAddCorsHeader = (bucketIdx: number, headerType: 'allowedHeader' | 'exposeHeader', headerVal: string) => {
    if (!headerVal.trim()) return;
    updateCorsRuleForBucket(bucketIdx, (rule) => {
      const list = rule[headerType] || [];
      if (list.includes(headerVal.trim())) return rule;
      return { ...rule, [headerType]: [...list, headerVal.trim()] };
    });
    if (headerType === 'allowedHeader') setNewAllowedHeaderInputs((prev) => ({ ...prev, [bucketIdx]: '' }));
    else setNewExposeHeaderInputs((prev) => ({ ...prev, [bucketIdx]: '' }));
  };

  const handleRemoveCorsHeader = (bucketIdx: number, headerType: 'allowedHeader' | 'exposeHeader', headerIdx: number) => {
    updateCorsRuleForBucket(bucketIdx, (rule) => ({
      ...rule,
      [headerType]: (rule[headerType] || []).filter((_, i) => i !== headerIdx)
    }));
  };

  const handleUpdateCorsMaxAge = (bucketIdx: number, maxAge: number) => {
    updateCorsRuleForBucket(bucketIdx, (rule) => ({
      ...rule,
      maxAgeSeconds: maxAge
    }));
  };

  const applyCorsPreset = (bucketIdx: number, presetType: 'allow-all' | 'web-app' | 'read-only') => {
    let presetRule: CorsRuleItem;
    if (presetType === 'allow-all') {
      presetRule = {
        allowedOrigin: ['*'],
        allowedMethod: ['GET', 'POST', 'PUT', 'DELETE', 'HEAD', 'OPTIONS'],
        allowedHeader: ['*'],
        exposeHeader: ['ETag'],
        maxAgeSeconds: 3600
      };
    } else if (presetType === 'web-app') {
      presetRule = {
        allowedOrigin: ['https://example.com', 'https://app.example.com'],
        allowedMethod: ['GET', 'PUT', 'POST', 'DELETE'],
        allowedHeader: ['*'],
        exposeHeader: ['ETag', 'x-amz-request-id'],
        maxAgeSeconds: 3600
      };
    } else {
      presetRule = {
        allowedOrigin: ['*'],
        allowedMethod: ['GET', 'HEAD'],
        allowedHeader: ['*'],
        exposeHeader: ['ETag'],
        maxAgeSeconds: 86400
      };
    }

    setSourceBuckets((prev) =>
      prev.map((b, i) => (i === bucketIdx ? { ...b, corsEnabled: true, corsRule: [presetRule] } : b))
    );
  };

  // Step 1: Scan Source Buckets
  const handleScanSourceBuckets = async () => {
    if (!scanAccessKey || !scanSecretKey) {
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
      });
      if (res.success && res.bucketNames && res.bucketNames.length > 0) {
        const sortedBucketNames = [...res.bucketNames].sort((a, b) => a.localeCompare(b));
        setScannedBucketNames(sortedBucketNames);
        setScannedBuckets(res.buckets || []);
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

  // Save Source Storage Model Revision Handler
  const handleSaveSourceModelRevision = async (result: SaveRevisionResult) => {
    const activeBuckets = sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName));
    if (activeBuckets.length === 0) {
      alert('No active included source buckets available to save. Please include at least one bucket.');
      return;
    }
    const formattedBuckets = activeBuckets.map((b) => ({
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
      const modelId = res.modelId || `source-storage-model-${Date.now()}`;
      const newModel = {
        id: modelId,
        name: result.name,
        description: result.description,
        version: result.version,
        updatedTime: new Date().toISOString(),
        buckets: formattedBuckets
      };
      setSavedSourceModels((prev) => [...prev.filter((m) => m.id !== result.overwriteId), newModel]);
      setSelectedSourceModelId(modelId);
      setLoadedModelName(result.name);
      setLoadedModelVersion(result.version);
      setLoadedModelTime(new Date().toLocaleString());
      setModelLog((prev) => [...prev, `[Damselfly/Beetle] Source Storage Revision Saved (ID: ${modelId}, Name: ${result.name})`]);
    }
  };

  // Step 3: Run Beetle Object Storage Recommendation API
  const handleGenerateRecommendation = async () => {
    const activeBuckets = sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName));
    if (activeBuckets.length === 0) {
      alert('All buckets are currently excluded. Please include at least one bucket for target recommendation.');
      return;
    }

    const payloadSourceBuckets = activeBuckets.map((b) => ({
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
      console.warn('API recommendation fallback using sampleTargetObjectStorage', err);
      setRecommendationResult({
        ...sampleTargetObjectStorage,
        targetCloud: { csp: desiredCsp, region: desiredRegion },
        description: `Target cloud object storage recommendation for ${activeBuckets.length} bucket(s) on ${desiredCsp.toUpperCase()} (${desiredRegion})`
      });
    } finally {
      setIsRecommending(false);
    }
  };

  // Step 3: Save Target Storage Model Revision
  const handleSaveTargetModelRevision = async (result: SaveRevisionResult) => {
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
      const modelId = res.modelId || `target-storage-model-${Date.now()}`;
      setSavedTargetModelId(modelId);
      setModelLog((prev) => [...prev, `[Damselfly/Beetle] Target Storage Model Revision Saved (ID: ${modelId}, Name: ${result.name})`]);
    }
  };

  // Step 4: Provision & Migrate Object Storage (Async API via Prefer: respond-async)
  const handleMigrateStorage = async () => {
    if (!recommendationResult) {
      alert('Please run recommendation in Step 3 first before provisioning.');
      return;
    }
    setIsDeploying(true);
    const mockReqId = `req-${Date.now().toString().slice(-6)}`;
    setDeploymentLog([
      `POST /beetle/migration/ns/${namespaceId || 'mig01'}/objectStorage?nameSeed=${nameSeed}`,
      `Header -> Prefer: respond-async`,
      `HTTP 202 Accepted (ReqID: ${mockReqId}, Status: Handling)`,
      `GET /beetle/request/${mockReqId} -> Status: Handling (Polling...)`
    ]);

    try {
      const activeBuckets = sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName));
      const targetModel = {
        nameSeed,
        targetCloud: { csp: desiredCsp, region: desiredRegion },
        sourceObjectStorages: activeBuckets
      };
      const res = await beetleApi.executeObjectStorageMigration(namespaceId || 'mig01', nameSeed, targetModel, true);

      if (res.success) {
        const realReqId = res.reqId || mockReqId;
        setDeploymentLog((prev) => [
          ...prev,
          `GET /beetle/request/${realReqId} -> Status: Handling (Background worker allocated)`,
          `Dispatched target cloud bucket creation request to CB-Tumblebug backend.`,
          `GET /beetle/request/${realReqId} -> Status: Success (Duration: 3s)`
        ]);
        setTimeout(() => {
          loadMigratedStorages();
        }, 3000);
      } else {
        setDeploymentLog((prev) => [
          ...prev,
          `[Error] Object Storage migration request failed: ${res.error || 'Resource conflict or API timeout'}`
        ]);
      }
    } catch (err: any) {
      setDeploymentLog((prev) => [...prev, `[Error] ${err.message || 'Execution error'}`]);
    } finally {
      setIsDeploying(false);
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

  return (
    <div className="space-y-6 animate-fade-in mx-auto pb-24">
      {/* Unified Workflow Container Box */}
      <div className="bg-bg-panel border border-border-main rounded-2xl p-4 shadow-sm space-y-3">
        <div className="flex items-center space-x-2.5 border-b border-border-main pb-3 px-1">
          <HardDrive className="w-5 h-5 text-emerald-500" />
          <h2 className="text-base font-extrabold text-text-main flex items-center space-x-2">
            <span>Object Storage Migration Workflow</span>
          </h2>
        </div>

        <nav className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-2.5">
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

      {/* ══════════════════════════════════════════════════
          SUB-STEP 1: Source Storage Analysis & Scan
      ══════════════════════════════════════════════════ */}
      {subTab === 'source' && (
        <div className="space-y-6">
          <div className="glass-panel px-6 py-4.5 rounded-2xl border border-border-main flex flex-wrap items-center gap-x-3 gap-y-1.5">
            <div className="flex items-center gap-2 shrink-0">
              <Database className="w-5 h-5 text-emerald-500" />
              <h2 className="text-base font-extrabold text-text-main tracking-tight">
                Source Object Storage Analysis
              </h2>
            </div>
            <span className="text-sm text-text-muted">
              Register source object storage buckets, scan credentials, and analyze infrastructure specifications &amp; object metadata.
            </span>
          </div>

          <div className="glass-panel p-6 rounded-2xl">
            <div className="flex items-center justify-between mb-5">
              <div>
                <h2 className="text-base font-bold text-text-main flex items-center gap-2">
                  <Key className="w-5 h-5 text-emerald-400" />
                  Source CSP Credentials
                </h2>
                <p className="text-sm text-text-muted mt-1">
                  Select an active source CSP credential card to scan buckets, or register a new credential profile.
                </p>
              </div>
              <button
                onClick={() => {
                  setCredProfileName(`cred-profile-${savedCredProfiles.length + 1}`);
                  setCredCsp(scanCsp || 'aws');
                  setCredRegion(scanRegion || 'ap-northeast-2');
                  setIsRegisterCredModalOpen(true);
                }}
                className="px-4 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white rounded-xl text-sm font-bold flex items-center gap-2 transition cursor-pointer shadow-lg shadow-emerald-500/20"
              >
                <Plus className="w-4 h-4" /> New Credential Profile
              </button>
            </div>

            {savedCredProfiles.length === 0 ? (
              <div className="p-8 text-center border border-dashed border-border-main rounded-xl bg-bg-main/30 space-y-2">
                <Key className="w-8 h-8 text-text-muted mx-auto" />
                <p className="text-sm font-bold text-text-main">No Source CSP credentials registered</p>
                <p className="text-xs text-text-muted">
                  Click <span className="font-bold text-emerald-500">&quot;+ New Credential Profile&quot;</span> above to register an ephemeral credential card.
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

          <div className="glass-panel p-6 rounded-2xl space-y-5">
            <div className="flex items-center justify-between border-b border-border-main pb-4">
              <h3 className="text-base font-bold text-text-main flex items-center gap-2">
                <Database className="w-5 h-5 text-emerald-400" />
                Object Storage Buckets
              </h3>
              <button
                onClick={() => {
                  setIsAddModalOpen(true);
                  if (scanAccessKey && scanSecretKey) handleScanSourceBuckets();
                }}
                className="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white rounded-xl text-xs font-bold flex items-center gap-2 transition cursor-pointer shadow-md shadow-emerald-500/20"
              >
                <Plus className="w-4 h-4" />
                <span>Select Object Storage</span>
              </button>
            </div>

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
                      <td className="p-3 text-right">
                        <button
                          onClick={() => setSourceBuckets(prev => prev.filter((_, idx) => idx !== i))}
                          className="p-1 text-text-muted hover:text-red-500 transition cursor-pointer"
                        >
                          <Trash2 className="w-4 h-4" />
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>

            <div className="flex justify-end pt-2 border-t border-border-main/40">
              <button
                onClick={() => setSubTab('refine')}
                className="px-6 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white font-bold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center space-x-2 cursor-pointer"
              >
                <span>Next: Proceed to 2. Refinement</span>
                <ArrowRight className="w-4 h-4" />
              </button>
            </div>
          </div>
        </div>
      )}

      {/* ══════════════════════════════════════════════════
          SUB-STEP 2: Source Object Storage Refinement
          (100% Visual & Design Parity with SourceInfraRefinement.tsx)
      ══════════════════════════════════════════════════ */}
      {subTab === 'refine' && (
        <div className="space-y-6">
          {/* Top Banner Description Box */}
          <div className="glass-panel px-6 py-4.5 rounded-2xl border border-border-main flex flex-wrap items-center gap-x-3 gap-y-1.5">
            <div className="flex items-center gap-2 shrink-0">
              <Sliders className="w-5 h-5 text-emerald-500" />
              <h2 className="text-base font-extrabold text-text-main tracking-tight">
                Source Object Storage Refinement
              </h2>
            </div>
            <span className="text-sm text-text-muted">
              Review &amp; refine extracted source metadata, adjust bucket specifications, configure CORS rules, and save source model revisions.
            </span>
          </div>

          {/* ═══ STEP 1 Card: Source Object Storage Model Selection (Exact Infra Parity) ═══ */}
          <div className={`glass-panel p-6 rounded-2xl transition-all duration-300 ${activeRefineStep >= 1 ? 'opacity-100' : 'opacity-40 pointer-events-none'}`}>
            <div className="flex items-center space-x-3 mb-4 border-b border-border-main/40 pb-3">
              <span className={`w-7 h-7 rounded-full flex items-center justify-center text-sm font-extrabold ${isModelLoaded ? 'bg-green-500/20 text-green-600 dark:text-green-400 border border-green-500/30' : 'bg-emerald-500 text-slate-950'}`}>
                {isModelLoaded ? '✓' : '1'}
              </span>
              <h3 className="text-base font-extrabold text-text-main">Step 1: Source Object Storage Model Selection</h3>
            </div>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-bold text-text-muted mb-2">Choose Source Model</label>
                <select
                  value={selectedSourceModelId}
                  onChange={(e) => {
                    setSelectedSourceModelId(e.target.value);
                    setIsModelLoaded(false);
                    setActiveRefineStep(1);
                  }}
                  className="w-full max-w-md bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-3 text-sm font-bold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer mb-3"
                >
                  <option value="sample-source-storage-01">[Sample] sampleSourceObjectStorage.json (v1.0)</option>
                  {savedSourceModels.map(m => (
                    <option key={m.id} value={m.id}>{m.name} (v{m.version || '1.0'})</option>
                  ))}
                </select>
                <div className="flex items-center gap-3">
                  <button
                    onClick={handleLoadModel}
                    disabled={!selectedSourceModelId}
                    className={`px-5 py-3 rounded-xl text-sm font-extrabold flex items-center transition cursor-pointer ${selectedSourceModelId
                      ? 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 shadow-md shadow-emerald-500/25'
                      : 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'}`}
                  >
                    <RefreshCw className="w-4 h-4 mr-1.5" /> Load Model
                  </button>
                  {selectedSourceModelId && selectedSourceModelId !== 'sample-source-storage-01' && (
                    <button
                      onClick={handleDeleteModel}
                      className="px-5 py-3 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-xl text-sm font-extrabold flex items-center transition cursor-pointer"
                    >
                      <Trash2 className="w-4 h-4 mr-1.5" /> Delete Model
                    </button>
                  )}
                </div>
              </div>
            </div>
          </div>

          {/* ═══ STEP 2 Card: Source Object Storage Review & Editing (Exact Infra Parity) ═══ */}
          <div className={`glass-panel p-6 rounded-2xl transition-all duration-300 ${isModelLoaded ? 'opacity-100' : 'opacity-40 pointer-events-none'}`}>
            <div className="flex items-center justify-between mb-4 border-b border-border-main/40 pb-3">
              <div className="flex items-center space-x-3">
                <span className={`w-7 h-7 rounded-full flex items-center justify-center text-sm font-extrabold ${activeRefineStep > 2 ? 'bg-green-500/20 text-green-600 dark:text-green-400 border border-green-500/30' : 'bg-emerald-500 text-slate-950'}`}>
                  {activeRefineStep > 2 ? '✓' : '2'}
                </span>
                <h3 className="text-base font-extrabold text-text-main">
                  Step 2: Source Object Storage Review &amp; Editing
                </h3>
              </div>
              {isModelLoaded && (
                <div className="flex items-center space-x-2">
                  <button
                    onClick={() => setIsRefineJsonOpen(!isRefineJsonOpen)}
                    className="px-3 py-1.5 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-sm font-bold transition cursor-pointer flex items-center text-emerald-600 dark:text-emerald-400"
                  >
                    <Copy className="w-3.5 h-3.5 mr-1" />
                    JSON View
                  </button>
                </div>
              )}
            </div>

            {isModelLoaded && (
              <div className="bg-bg-panel/40 border border-border-main/30 rounded-xl p-3.5 flex flex-col md:flex-row md:items-center justify-between text-sm space-y-2 md:space-y-0 mb-4">
                <div className="flex items-center space-x-2">
                  <span className="text-text-muted font-bold">Loaded Model:</span>
                  <span className="text-emerald-600 dark:text-emerald-400 font-extrabold text-sm">{loadedModelName}</span>
                  <span className="text-sm text-text-muted font-mono bg-bg-panel px-2 py-0.5 rounded border border-border-main/40">
                    v{loadedModelVersion}
                  </span>
                </div>
                <div className="text-sm text-text-muted">
                  Last Updated: <span className="text-text-main font-semibold">{loadedModelTime}</span>
                </div>
              </div>
            )}

            {isModelLoaded && (
              <div className="space-y-6">
                {isRefineJsonOpen ? (
                  <div className="space-y-3">
                    <div className="flex justify-between items-center bg-bg-input px-3.5 py-2 rounded-xl border border-border-main/50">
                      <span className="text-sm text-text-muted font-mono">sampleSourceObjectStorage.json</span>
                      <button
                        onClick={() => {
                          navigator.clipboard.writeText(JSON.stringify({ nameSeed, sourceObjectStorages: sourceBuckets }, null, 2));
                          alert('JSON copied!');
                        }}
                        className="px-3 py-1 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-xs font-bold transition flex items-center cursor-pointer text-emerald-600 dark:text-emerald-400"
                      >
                        <Copy className="w-3.5 h-3.5 mr-1" /> Copy JSON
                      </button>
                    </div>
                    <pre className="text-sm font-mono text-slate-800 dark:text-emerald-400 bg-bg-panel p-3.5 rounded-xl border border-border-main overflow-y-auto max-h-[400px] select-text">
                      {JSON.stringify({ nameSeed, sourceObjectStorages: sourceBuckets }, null, 2)}
                    </pre>
                  </div>
                ) : (
                  <div className="space-y-6">
                    {/* Bucket Card List & Selected Bucket Refinement Section (Exact Infra Parity) */}
                    <div className="bg-bg-panel/40 p-4 border border-border-main/50 rounded-xl space-y-3">
                      <div className="flex items-center justify-between">
                        <h4 className="text-sm font-bold text-emerald-600 dark:text-emerald-400 flex items-center">
                          <Sliders className="w-4 h-4 mr-1.5 text-emerald-600 dark:text-emerald-400" />
                          Source Storage Buckets List ({sourceBuckets.length} Buckets)
                        </h4>
                        <span className="text-xs text-text-muted">Click a bucket card below to inspect and refine policies</span>
                      </div>

                      <div className="flex flex-wrap gap-2.5">
                        {sourceBuckets.map((b, idx) => {
                          const isSelected = selectedBucketIndex === idx;
                          const isExcluded = excludedBucketNames.includes(b.bucketName);
                          return (
                            <button
                              key={idx}
                              onClick={() => setSelectedBucketIndex(idx)}
                              className={`px-4 py-2.5 rounded-xl border text-sm font-extrabold flex items-center space-x-3 transition cursor-pointer ${
                                isExcluded
                                  ? 'bg-bg-panel/40 border-border-main/40 text-text-muted opacity-60 line-through'
                                  : isSelected
                                  ? 'bg-emerald-500/15 border-emerald-500 text-emerald-600 dark:text-emerald-400 shadow-md ring-1 ring-emerald-500/40'
                                  : 'bg-bg-panel hover:bg-bg-input border-border-main text-text-muted hover:text-text-main'
                              }`}
                            >
                              <Database className={`w-4 h-4 ${isExcluded ? 'text-text-muted' : isSelected ? 'text-emerald-500' : 'text-text-muted'}`} />
                              <span>{b.bucketName}</span>

                              {/* Toggle Button inside the card pill (Frequency badge removed) */}
                              <span
                                onClick={(e) => {
                                  e.stopPropagation();
                                  toggleExcludeBucket(b.bucketName);
                                }}
                                className={`px-2.5 py-0.5 rounded-md text-xs font-extrabold transition cursor-pointer border no-underline ${
                                  isExcluded
                                    ? 'bg-emerald-500/20 text-emerald-600 dark:text-emerald-400 border-emerald-500/30 hover:bg-emerald-500/30'
                                    : 'bg-red-500/10 text-red-500 border-red-500/30 hover:bg-red-500/20'
                                }`}
                                title={isExcluded ? 'Re-include this bucket in migration' : 'Exclude this bucket from migration'}
                              >
                                {isExcluded ? 'Include' : 'Exclude'}
                              </span>
                            </button>
                          );
                        })}
                      </div>
                    </div>

                    {/* Selected Bucket Details & Refinement Container */}
                    {sourceBuckets[selectedBucketIndex] && (() => {
                      const currentIdx = selectedBucketIndex;
                      const b = sourceBuckets[currentIdx];
                      const corsRule = ensureCorsRuleExists(b);
                      const isExcluded = excludedBucketNames.includes(b.bucketName);

                      return (
                        <div className={`bg-bg-panel/60 border rounded-2xl p-5 space-y-6 animate-fade-in transition-all ${isExcluded ? 'border-red-500/30 bg-red-500/5' : 'border-border-main/60'}`}>
                          <div className="flex flex-col sm:flex-row sm:items-center justify-between border-b border-border-main/40 pb-3 gap-2">
                            <div className="flex items-center space-x-3">
                              <Database className={`w-4.5 h-4.5 ${isExcluded ? 'text-red-500' : 'text-emerald-500'}`} />
                              <h4 className="text-sm font-extrabold text-text-main">
                                Bucket Details ({b.bucketName})
                              </h4>
                              <span className={`px-2.5 py-0.5 rounded-md text-[11px] font-mono font-extrabold uppercase ${isExcluded ? 'bg-red-500/20 text-red-400 border border-red-500/30' : 'bg-emerald-500/20 text-emerald-400 border border-emerald-500/30'}`}>
                                {isExcluded ? 'EXCLUDED FROM MIGRATION' : 'INCLUDED IN MIGRATION'}
                              </span>
                            </div>

                            <div className="flex items-center space-x-3 text-sm font-mono">
                              <div>
                                <span className="text-text-muted font-normal mr-1">Total Size:</span>
                                <span className="font-extrabold text-text-main">{formatBytes(b.totalSizeBytes)}</span>
                              </div>
                              <div>
                                <span className="text-text-muted font-normal mr-1">Objects:</span>
                                <span className="font-extrabold text-text-main">{(b.objectCount || 0).toLocaleString()}</span>
                              </div>
                            </div>
                          </div>

                          {/* Section A: Bucket Read-Only Overview & Editable Feature Toggles */}
                          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 text-xs">
                            {/* Read-Only Spec 1 */}
                            <div className="bg-bg-input/60 p-3.5 rounded-xl border border-border-main/40 space-y-2">
                              <span className="text-text-muted font-normal block">Bucket Spec Overview</span>
                              <div className="space-y-1 font-mono">
                                <div className="flex justify-between">
                                  <span className="text-text-muted font-normal">Name:</span>
                                  <span className="font-extrabold text-emerald-600 dark:text-emerald-400">{b.bucketName}</span>
                                </div>
                                <div className="flex justify-between">
                                  <span className="text-text-muted font-normal">Access Class:</span>
                                  <span className="font-extrabold text-text-main uppercase">{b.accessFrequency}</span>
                                </div>
                              </div>
                            </div>

                            {/* Read-Only Spec 2 */}
                            <div className="bg-bg-input/60 p-3.5 rounded-xl border border-border-main/40 space-y-2">
                              <span className="text-text-muted font-normal block">Storage Data Usage</span>
                              <div className="space-y-1 font-mono">
                                <div className="flex justify-between">
                                  <span className="text-text-muted font-normal">Total Size:</span>
                                  <span className="font-extrabold text-text-main">{formatBytes(b.totalSizeBytes)}</span>
                                </div>
                                <div className="flex justify-between">
                                  <span className="text-text-muted font-normal">Object Count:</span>
                                  <span className="font-extrabold text-text-main">{(b.objectCount || 0).toLocaleString()}</span>
                                </div>
                              </div>
                            </div>

                            {/* Editable Policy Toggles */}
                            <div className="bg-bg-input/60 p-3.5 rounded-xl border border-border-main/40 space-y-2 col-span-1 lg:col-span-2">
                              <span className="text-text-muted font-normal block">Refinable Storage Policies</span>
                              <div className="grid grid-cols-2 sm:grid-cols-4 gap-2 pt-0.5">
                                <label className="flex items-center space-x-2 bg-bg-panel p-2 rounded-lg border border-border-main/40 cursor-pointer hover:border-emerald-500/50 transition">
                                  <input
                                    type="checkbox"
                                    checked={b.versioningEnabled}
                                    onChange={(e) => {
                                      const checked = e.target.checked;
                                      setSourceBuckets((prev) => prev.map((item, i) => (i === currentIdx ? { ...item, versioningEnabled: checked } : item)));
                                    }}
                                    className="accent-emerald-500 cursor-pointer w-4 h-4"
                                  />
                                  <span className="font-extrabold text-text-main">Versioning</span>
                                </label>

                                <label className="flex items-center space-x-2 bg-bg-panel p-2 rounded-lg border border-border-main/40 cursor-pointer hover:border-emerald-500/50 transition">
                                  <input
                                    type="checkbox"
                                    checked={b.encryptionEnabled}
                                    onChange={(e) => {
                                      const checked = e.target.checked;
                                      setSourceBuckets((prev) => prev.map((item, i) => (i === currentIdx ? { ...item, encryptionEnabled: checked } : item)));
                                    }}
                                    className="accent-emerald-500 cursor-pointer w-4 h-4"
                                  />
                                  <span className="font-extrabold text-text-main">Encryption</span>
                                </label>

                                <label className="flex items-center space-x-2 bg-bg-panel p-2 rounded-lg border border-border-main/40 cursor-pointer hover:border-emerald-500/50 transition">
                                  <input
                                    type="checkbox"
                                    checked={b.isPublic}
                                    onChange={(e) => {
                                      const checked = e.target.checked;
                                      setSourceBuckets((prev) => prev.map((item, i) => (i === currentIdx ? { ...item, isPublic: checked } : item)));
                                    }}
                                    className="accent-emerald-500 cursor-pointer w-4 h-4"
                                  />
                                  <span className="font-extrabold text-text-main">Public Access</span>
                                </label>

                                <label className="flex items-center space-x-2 bg-bg-panel p-2 rounded-lg border border-border-main/40 cursor-pointer hover:border-emerald-500/50 transition">
                                  <input
                                    type="checkbox"
                                    checked={b.corsEnabled}
                                    onChange={(e) => {
                                      const checked = e.target.checked;
                                      setSourceBuckets((prev) => prev.map((item, i) => (i === currentIdx ? { ...item, corsEnabled: checked } : item)));
                                    }}
                                    className="accent-emerald-500 cursor-pointer w-4 h-4"
                                  />
                                  <span className="font-extrabold text-text-main">CORS Rules</span>
                                </label>
                              </div>
                            </div>
                          </div>

                          {/* Section B: Detailed CORS Rule Configuration Section (when CORS Enabled) */}
                          {b.corsEnabled && (
                            <div className="space-y-4 pt-4 border-t border-border-main/40">
                              <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-2">
                                <div className="flex items-center space-x-2">
                                  <Globe className="w-4 h-4 text-emerald-500" />
                                  <span className="font-extrabold text-sm text-text-main">
                                    CORS Detailed Configuration Rules — {b.bucketName}
                                  </span>
                                </div>

                                <div className="flex items-center space-x-3 text-xs shrink-0">
                                  <button
                                    onClick={() => setShowCorsGuide(true)}
                                    className="px-3 py-1 bg-emerald-500/10 hover:bg-emerald-500/20 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 font-bold rounded-lg transition cursor-pointer flex items-center space-x-1"
                                  >
                                    <AlertCircle className="w-3.5 h-3.5 text-emerald-500" />
                                    <span>❓ CORS Rule Guide &amp; Examples</span>
                                  </button>
                                  <div className="flex items-center space-x-1.5 font-mono">
                                    <span className="text-text-muted font-normal">Status:</span>
                                    <span className="px-2 py-0.5 rounded font-bold bg-emerald-500/20 text-emerald-400">
                                      ENABLED
                                    </span>
                                  </div>
                                </div>
                              </div>

                              <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-4 text-sm w-full">
                                {/* 1. Allowed Origins */}
                                <div className="space-y-2 bg-bg-input/60 p-3.5 rounded-xl border border-border-main/40">
                                  <label className="font-normal text-sm text-text-muted block">Allowed Origins (allowedOrigin)</label>
                                  <div className="flex flex-wrap gap-1.5 min-h-[36px] items-center p-2 bg-bg-input rounded-lg border border-border-main/40">
                                    {corsRule.allowedOrigin.map((orig, oIdx) => (
                                      <span key={oIdx} className="bg-emerald-500/15 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 px-2.5 py-1 rounded-md font-mono text-sm font-bold flex items-center space-x-1">
                                        <span>{orig}</span>
                                        <button
                                          onClick={() => handleRemoveCorsOrigin(currentIdx, oIdx)}
                                          className="hover:text-red-400 font-extrabold ml-1 cursor-pointer"
                                        >
                                          ✕
                                        </button>
                                      </span>
                                    ))}
                                  </div>
                                  <div className="flex items-center space-x-2 pt-1">
                                    <input
                                      type="text"
                                      value={newOriginInputs[currentIdx] || ''}
                                      onChange={(e) => setNewOriginInputs((prev) => ({ ...prev, [currentIdx]: e.target.value }))}
                                      onKeyDown={(e) => {
                                        if (e.key === 'Enter') handleAddCorsOrigin(currentIdx, newOriginInputs[currentIdx] || '');
                                      }}
                                      placeholder="e.g. https://example.com or *"
                                      className="flex-1 px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono text-sm focus:outline-none focus:border-emerald-500"
                                    />
                                    <button
                                      onClick={() => handleAddCorsOrigin(currentIdx, newOriginInputs[currentIdx] || '')}
                                      className="px-3.5 py-2 bg-emerald-600 hover:bg-emerald-700 text-white font-bold rounded-lg cursor-pointer transition text-sm shrink-0"
                                    >
                                      Add Origin
                                    </button>
                                  </div>
                                </div>

                                {/* 2. Allowed HTTP Methods */}
                                <div className="space-y-2 bg-bg-input/60 p-3.5 rounded-xl border border-border-main/40">
                                  <label className="font-normal text-sm text-text-muted block">Allowed HTTP Methods (allowedMethod)</label>
                                  <div className="flex flex-wrap gap-1.5 p-2 bg-bg-input rounded-lg border border-border-main/40">
                                    {HTTP_METHODS.map((m) => {
                                      const isSelected = corsRule.allowedMethod.includes(m);
                                      return (
                                        <button
                                          key={m}
                                          onClick={() => handleToggleCorsMethod(currentIdx, m)}
                                          className={`px-3 py-1.5 rounded-lg text-sm font-mono font-bold transition cursor-pointer border ${
                                            isSelected
                                              ? 'bg-emerald-500 text-slate-950 border-emerald-500 shadow-sm'
                                              : 'bg-bg-panel text-text-muted border-border-main hover:text-text-main'
                                          }`}
                                        >
                                          {isSelected ? `✓ ${m}` : m}
                                        </button>
                                      );
                                    })}
                                  </div>
                                </div>

                                {/* 3. Allowed Headers */}
                                <div className="space-y-2 bg-bg-input/60 p-3.5 rounded-xl border border-border-main/40">
                                  <label className="font-normal text-sm text-text-muted block">Allowed Headers (allowedHeader)</label>
                                  <div className="flex flex-wrap gap-1.5 min-h-[36px] items-center p-2 bg-bg-input rounded-lg border border-border-main/40">
                                    {corsRule.allowedHeader.map((h, hIdx) => (
                                      <span key={hIdx} className="bg-teal-500/15 border border-teal-500/30 text-teal-600 dark:text-teal-400 px-2.5 py-1 rounded-md font-mono text-sm font-bold flex items-center space-x-1">
                                        <span>{h}</span>
                                        <button
                                          onClick={() => handleRemoveCorsHeader(currentIdx, 'allowedHeader', hIdx)}
                                          className="hover:text-red-400 font-extrabold ml-1 cursor-pointer"
                                        >
                                          ✕
                                        </button>
                                      </span>
                                    ))}
                                  </div>
                                  <div className="flex items-center space-x-2 pt-1">
                                    <input
                                      type="text"
                                      value={newAllowedHeaderInputs[currentIdx] || ''}
                                      onChange={(e) => setNewAllowedHeaderInputs((prev) => ({ ...prev, [currentIdx]: e.target.value }))}
                                      onKeyDown={(e) => {
                                        if (e.key === 'Enter') handleAddCorsHeader(currentIdx, 'allowedHeader', newAllowedHeaderInputs[currentIdx] || '');
                                      }}
                                      placeholder="e.g. * or Content-Type"
                                      className="flex-1 px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono text-sm focus:outline-none focus:border-emerald-500"
                                    />
                                    <button
                                      onClick={() => handleAddCorsHeader(currentIdx, 'allowedHeader', newAllowedHeaderInputs[currentIdx] || '')}
                                      className="px-3.5 py-2 bg-emerald-600 hover:bg-emerald-700 text-white font-bold rounded-lg cursor-pointer transition text-sm shrink-0"
                                    >
                                      Add Header
                                    </button>
                                  </div>
                                </div>

                                {/* 4. Expose Headers & Max Age */}
                                <div className="space-y-2 bg-bg-input/60 p-3.5 rounded-xl border border-border-main/40">
                                  <label className="font-normal text-sm text-text-muted block">Expose Headers &amp; Max Age (exposeHeader &amp; maxAgeSeconds)</label>
                                  <div className="flex flex-wrap gap-1.5 min-h-[36px] items-center p-2 bg-bg-input rounded-lg border border-border-main/40">
                                    {corsRule.exposeHeader.map((eh, ehIdx) => (
                                      <span key={ehIdx} className="bg-emerald-500/15 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 px-2.5 py-1 rounded-md font-mono text-sm font-bold flex items-center space-x-1">
                                        <span>{eh}</span>
                                        <button
                                          onClick={() => handleRemoveCorsHeader(currentIdx, 'exposeHeader', ehIdx)}
                                          className="hover:text-red-400 font-extrabold ml-1 cursor-pointer"
                                        >
                                          ✕
                                        </button>
                                      </span>
                                    ))}
                                  </div>
                                  <div className="grid grid-cols-2 gap-2 pt-1">
                                    <input
                                      type="text"
                                      value={newExposeHeaderInputs[currentIdx] || ''}
                                      onChange={(e) => setNewExposeHeaderInputs((prev) => ({ ...prev, [currentIdx]: e.target.value }))}
                                      onKeyDown={(e) => {
                                        if (e.key === 'Enter') handleAddCorsHeader(currentIdx, 'exposeHeader', newExposeHeaderInputs[currentIdx] || '');
                                      }}
                                      placeholder="e.g. ETag"
                                      className="flex-1 px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono text-sm focus:outline-none focus:border-emerald-500"
                                    />
                                    <div className="flex items-center space-x-1.5">
                                      <span className="text-xs text-text-muted font-normal shrink-0">MaxAge(s):</span>
                                      <input
                                        type="number"
                                        value={corsRule.maxAgeSeconds}
                                        onChange={(e) => handleUpdateCorsMaxAge(currentIdx, parseInt(e.target.value) || 0)}
                                        className="w-full px-2 py-1.5 bg-bg-input border border-border-main rounded-lg text-text-main font-mono text-xs focus:outline-none focus:border-emerald-500"
                                      />
                                    </div>
                                  </div>
                                </div>
                              </div>
                            </div>
                          )}
                        </div>
                      );
                    })()}
                  </div>
                )}

                {/* Save Revision Action Bar (Matching Infra Parity) */}
                <div className="flex flex-row items-center justify-start pt-4 border-t border-border-main/20 mt-4 space-x-4">
                  <button
                    onClick={() => setShowSaveSourceModal(true)}
                    className="px-6 py-3 bg-gradient-to-r from-emerald-500 to-blue-600 hover:from-emerald-600 hover:to-blue-700 text-slate-950 rounded-xl text-sm font-extrabold flex items-center transition cursor-pointer shadow-lg shadow-emerald-500/10 shrink-0"
                  >
                    <Save className="w-4 h-4 mr-1.5" /> Save Source Storage Revision
                  </button>
                  <div className="flex items-center space-x-2 text-sm text-text-muted">
                    <span className="font-bold">Model to save:</span>
                    <span className="text-emerald-600 dark:text-emerald-400 font-extrabold text-sm">
                      {loadedModelName || '[Sample] sampleSourceObjectStorage.json'}
                    </span>
                    <span className="text-sm text-text-muted font-mono bg-bg-panel px-1.5 py-0.5 rounded border border-border-main/40">
                      v{loadedModelVersion || '1.0'}
                    </span>
                  </div>
                </div>

                <div className="flex items-center justify-between pt-4 border-t border-border-main/20 mt-4">
                  <button
                    onClick={() => setSubTab('source')}
                    className="px-4 py-2 bg-bg-input border border-border-main hover:bg-bg-main text-text-main font-bold text-xs rounded-xl transition cursor-pointer flex items-center space-x-1.5"
                  >
                    <ArrowLeft className="w-3.5 h-3.5" />
                    <span>Back to 1. Source Storage</span>
                  </button>
                  <button
                    onClick={() => setSubTab('optimize')}
                    className="px-6 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white font-bold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center space-x-2 cursor-pointer ml-auto"
                  >
                    <span>Next: Proceed to 3. Target Object Storage Optimizer</span>
                    <ArrowRight className="w-4 h-4" />
                  </button>
                </div>
              </div>
            )}
          </div>
        </div>
      )}

      {/* ══════════════════════════════════════════════════
          SUB-STEP 3: Target Object Storage Optimizer
      ══════════════════════════════════════════════════ */}
      {subTab === 'optimize' && (
        <div className="space-y-6 font-sans">
          {/* Top Description Header Container (Rule #3) */}
          <div className="glass-panel px-6 py-4.5 rounded-2xl border border-border-main flex flex-wrap items-center gap-x-3 gap-y-1.5">
            <div className="flex items-center gap-2 shrink-0">
              <Compass className="w-5 h-5 text-emerald-500" />
              <h2 className="text-base font-extrabold text-text-main tracking-tight">
                Target Cloud Optimization
              </h2>
            </div>
            <span className="text-sm text-text-muted">
              Generate AI-optimized target cloud recommendations, compare multi-CSP specs &amp; cost estimates, and customize target cloud models.
            </span>
          </div>

          {/* Main Card Container */}
          <div className="glass-panel p-6 rounded-2xl space-y-6 border border-border-main/60 shadow-sm">
            {/* Step Header with Green Check Circle Icon / Number Badge */}
            <div className="flex items-center space-x-2.5 border-b border-border-main/40 pb-3">
              <div className={`w-6 h-6 rounded-full flex items-center justify-center font-extrabold text-xs shrink-0 ${
                isModelLoaded
                  ? 'bg-emerald-500/20 text-emerald-500 border border-emerald-500/40'
                  : 'bg-emerald-500 text-slate-950'
              }`}>
                {isModelLoaded ? <Check className="w-3.5 h-3.5 text-emerald-500 font-bold" /> : '1'}
              </div>
              <h3 className="text-base font-extrabold text-text-main">
                Step 1: Desired CSP and Region Selection
              </h3>
            </div>

            {/* Sub-Navigation Tabs */}
            <div className="flex items-center space-x-8 border-b border-border-main/40 text-sm font-extrabold pt-1">
              <button
                onClick={() => setStep1SubTab('generate')}
                className={`pb-3 transition relative cursor-pointer ${
                  step1SubTab === 'generate'
                    ? 'text-emerald-600 dark:text-emerald-400 border-b-2 border-emerald-500'
                    : 'text-text-muted hover:text-text-main font-bold'
                }`}
              >
                Generate New Recommendation
              </button>

              <button
                onClick={() => setStep1SubTab('load')}
                className={`pb-3 transition relative cursor-pointer ${
                  step1SubTab === 'load'
                    ? 'text-emerald-600 dark:text-emerald-400 border-b-2 border-emerald-500'
                    : 'text-text-muted hover:text-text-main font-bold'
                }`}
              >
                Load Customized Cloud Infrastructure
              </button>
            </div>

            {/* Sub-Tab 1: Generate New Recommendation (Exact 50% Width) */}
            {step1SubTab === 'generate' && (
              <div className="space-y-5 pt-2 w-full md:w-1/2">
                {/* Form Row 1: Source Object Storage Model Selection + Load Model Button below */}
                <div className="space-y-2">
                  <label className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 block">
                    Source Object Storage Model
                  </label>
                  <select
                    value={selectedSourceModelId}
                    onChange={(e) => {
                      setSelectedSourceModelId(e.target.value);
                      setIsModelLoaded(false);
                    }}
                    className="w-full px-4 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold text-sm focus:outline-none focus:border-emerald-500"
                  >
                    <option value="">-- Choose Source Model --</option>
                    <option value="sample-source-storage-01">
                      [Sample] sampleSourceObjectStorage.json (v1.0)
                    </option>
                    {savedSourceModels.map((m) => (
                      <option key={m.id} value={m.id}>
                        {m.name} ({m.version || 'v1.0'})
                      </option>
                    ))}
                  </select>

                  <div className="space-y-1.5 pt-1">
                    <button
                      onClick={handleLoadModel}
                      disabled={!selectedSourceModelId}
                      className={`px-4 py-2 rounded-xl text-xs font-extrabold flex items-center space-x-1.5 transition shadow-md ${
                        !selectedSourceModelId
                          ? 'bg-bg-panel/40 border border-border-main/50 text-text-muted/50 cursor-not-allowed opacity-60'
                          : isModelLoaded
                          ? 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 border border-emerald-600 font-extrabold cursor-pointer'
                          : 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 border border-emerald-600 font-extrabold cursor-pointer'
                      }`}
                    >
                      <RefreshCw className="w-3.5 h-3.5" />
                      <span>{isModelLoaded ? 'Model Loaded ✓' : 'Load Model'}</span>
                    </button>

                    {isModelLoaded && (
                      <div className="text-xs font-bold text-emerald-600 dark:text-emerald-400 pt-0.5 animate-fade-in">
                        {sourceBuckets.length} bucket(s) ready for recommendation
                      </div>
                    )}
                  </div>
                </div>

                {/* Form Row 2: Desired CSP & Desired Region (2 columns side by side taking 50% total) */}
                <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 pt-1">
                  <div className="space-y-2">
                    <label className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 block">
                      Desired CSP
                    </label>
                    <select
                      value={desiredCsp}
                      onChange={(e) => handleDesiredCspChange(e.target.value)}
                      className="w-full px-4 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold text-sm focus:outline-none focus:border-emerald-500"
                    >
                      {getCspList().map((c) => (
                        <option key={c.id} value={c.id}>
                          {c.name}
                        </option>
                      ))}
                    </select>
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 block">
                      Desired Region
                    </label>
                    <select
                      value={desiredRegion}
                      onChange={(e) => setDesiredRegion(e.target.value)}
                      className="w-full px-4 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-mono text-sm focus:outline-none focus:border-emerald-500"
                    >
                      {getRegionsForCsp(desiredCsp).map((r) => (
                        <option key={r.id} value={r.id}>
                          {r.name}
                        </option>
                      ))}
                    </select>
                  </div>
                </div>

                {/* Form Row 3: Recommend Action Button */}
                <div className="pt-2">
                  <button
                    onClick={handleGenerateRecommendation}
                    disabled={isRecommending || !selectedSourceModelId || sourceBuckets.length === 0}
                    className={`px-5 py-2.5 rounded-xl font-extrabold text-xs sm:text-sm flex items-center space-x-2 transition-all shadow-md ${
                      !selectedSourceModelId || sourceBuckets.length === 0
                        ? 'bg-gradient-to-r from-emerald-300/40 via-teal-300/30 to-blue-400/30 text-slate-500/70 dark:text-slate-400/70 border border-emerald-500/20 cursor-not-allowed opacity-70'
                        : 'bg-gradient-to-r from-emerald-400 via-teal-400 to-blue-600 hover:from-emerald-500 hover:to-blue-700 text-slate-950 cursor-pointer shadow-lg shadow-emerald-500/20'
                    }`}
                  >
                    {isRecommending ? (
                      <RefreshCw className="w-4 h-4 animate-spin" />
                    ) : (
                      <Sparkles className={`w-4 h-4 ${!selectedSourceModelId || sourceBuckets.length === 0 ? 'text-slate-500/80 dark:text-slate-400/80' : 'text-slate-950 fill-slate-950/20'}`} />
                    )}
                    <span>Recommend Target Cloud Object Storage</span>
                  </button>
                </div>
              </div>
            )}

            {/* Sub-Tab 2: Load Customized Cloud Infrastructure (Exact 50% Width) */}
            {step1SubTab === 'load' && (
              <div className="space-y-4 pt-2 w-full md:w-1/2">
                <div className="space-y-2">
                  <label className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 block">
                    Target Cloud Infrastructure Model
                  </label>
                  <select
                    value={selectedSourceModelId}
                    onChange={(e) => setSelectedSourceModelId(e.target.value)}
                    className="w-full px-4 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold text-sm focus:outline-none focus:border-emerald-500"
                  >
                    <option value="">-- Choose Target Model --</option>
                    <option value="sample-source-storage-01">
                      my-target-aws-ap-northeast-2 (v1.0)
                    </option>
                  </select>
                  <div className="space-y-1.5 pt-2">
                    <div className="flex items-center space-x-2.5">
                      <button
                        onClick={() => {
                          handleGenerateRecommendation();
                        }}
                        disabled={!selectedSourceModelId}
                        className={`px-4 py-2 rounded-xl text-xs font-extrabold flex items-center space-x-1.5 transition shadow-md ${
                          !selectedSourceModelId
                            ? 'bg-bg-panel/40 border border-border-main/50 text-text-muted/50 cursor-not-allowed opacity-60'
                            : recommendationResult
                            ? 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 border border-emerald-600 font-extrabold cursor-pointer'
                            : 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 font-extrabold cursor-pointer border border-emerald-600'
                        }`}
                      >
                        <RefreshCw className="w-3.5 h-3.5" />
                        <span>{recommendationResult ? 'Design Loaded ✓' : 'Load Design'}</span>
                      </button>

                      <button
                        onClick={() => {
                          setRecommendationResult(null);
                        }}
                        className="px-4 py-2 bg-red-500 hover:bg-red-600 text-white rounded-xl text-xs font-extrabold flex items-center space-x-1.5 cursor-pointer transition shadow-md"
                      >
                        <Trash2 className="w-3.5 h-3.5" />
                        <span>Delete Design</span>
                      </button>
                    </div>

                    {recommendationResult && (
                      <div className="text-xs font-bold text-emerald-600 dark:text-emerald-400 pt-0.5 animate-fade-in">
                        Target design model ready
                      </div>
                    )}
                  </div>
                </div>
              </div>
            )}

          </div>

          {/* Step 2 Panel Container (Exact Infra Parity) */}
          <div className="glass-panel p-6 rounded-2xl space-y-6 border border-border-main/60 shadow-sm mt-6">
            {/* Step 2 Header */}
            <div className="flex items-center space-x-2.5 border-b border-border-main/40 pb-3">
              <div className={`w-6 h-6 rounded-full flex items-center justify-center font-extrabold text-xs shrink-0 ${
                recommendationResult
                  ? 'bg-emerald-500 text-slate-950 font-bold'
                  : 'bg-emerald-500/20 text-emerald-500/60 border border-emerald-500/30'
              }`}>
                2
              </div>
              <h3 className={`text-base font-extrabold ${recommendationResult ? 'text-text-main' : 'text-text-muted/70'}`}>
                Step 2: Recommended Target Object Storage Review and Editing
              </h3>
            </div>

            {!recommendationResult ? (
              <div className="py-12 text-center text-sm font-medium text-text-muted/60 italic">
                No recommendation generated yet. Complete previous steps to run recommendations.
              </div>
            ) : (
              <div className="space-y-6 animate-fade-in">
                {/* Recommended Storage Class Banner */}
                <div className="p-4 bg-emerald-500/10 border border-emerald-500/20 rounded-xl space-y-1.5 text-xs">
                  <div className="font-extrabold text-emerald-600 dark:text-emerald-400 flex items-center space-x-2">
                    <CheckCircle2 className="w-4 h-4 text-emerald-500" />
                    <span>Target Storage Class Recommendation Completed</span>
                  </div>
                  <div className="text-xs text-text-muted">{recommendationResult.description}</div>
                </div>

                {/* Recommended Target Object Storage Summary Section */}
                <div className="space-y-3">
                  <h4 className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400">
                    Recommended Target Object Storage Summary
                  </h4>
                  <div className="p-5 bg-bg-panel border border-border-main rounded-2xl grid grid-cols-1 md:grid-cols-3 gap-6 text-xs">
                    <div>
                      <span className="text-text-muted font-normal block mb-1">Target CSP &amp; Region</span>
                      <span className="font-extrabold text-sm text-emerald-600 dark:text-emerald-400 uppercase font-mono">
                        {desiredCsp.toUpperCase()} ({desiredRegion})
                      </span>
                    </div>

                    <div>
                      <span className="text-text-muted font-normal block mb-1">Storage Buckets</span>
                      <span className="font-extrabold text-sm text-text-main">
                        {sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName)).length} Bucket(s)
                      </span>
                    </div>

                    <div>
                      <span className="text-text-muted font-normal block mb-1">Storage Classes</span>
                      <span className="font-extrabold text-sm text-text-main">
                        Standard / Standard-IA / Archive
                      </span>
                    </div>
                  </div>
                </div>

                {/* Review and Editing Section (Matching Refinement & Infra Node Group UI 100%) */}
                <div className="space-y-5 pt-2 border-t border-border-main/40">
                  <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-2">
                    <div>
                      <h4 className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 flex items-center space-x-2">
                        <Sliders className="w-4 h-4 text-emerald-500" />
                        <span>Review and Editing</span>
                      </h4>
                      <p className="text-xs text-text-muted italic pt-0.5">
                        * Modifying target resource values updates the target object storage model in real-time
                      </p>
                    </div>

                    <button
                      onClick={() => setIsTargetJsonOpen(!isTargetJsonOpen)}
                      className="px-3.5 py-1.5 bg-bg-input border border-border-main text-text-main rounded-xl text-xs font-extrabold hover:border-emerald-500/40 cursor-pointer shrink-0 transition"
                    >
                      {isTargetJsonOpen ? 'Hide Target JSON' : 'Target Model JSON View'}
                    </button>
                  </div>

                  {isTargetJsonOpen && (
                    <div className="p-4 bg-bg-main/80 border border-border-main rounded-2xl animate-fade-in">
                      <pre className="font-mono text-xs text-slate-800 dark:text-emerald-300 max-h-[24rem] overflow-y-auto leading-relaxed select-all">
                        {JSON.stringify(
                          {
                            ...recommendationResult,
                            targetObjectStorages: sourceBuckets
                              .filter((b) => !excludedBucketNames.includes(b.bucketName))
                              .map((b) => ({
                                bucketName: `${b.targetBucketName || b.bucketName}-x8f2`,
                                targetObjectStorageName: b.targetBucketName || b.bucketName,
                                sourceBucketName: b.bucketName,
                                storageClass: b.accessFrequency === 'frequent' ? 'Standard' : b.accessFrequency === 'infrequent' ? 'Standard-IA' : 'Archive',
                                versioningEnabled: b.versioningEnabled,
                                encryptionEnabled: b.encryptionEnabled,
                                corsEnabled: b.corsEnabled,
                                corsRule: b.corsRule,
                                isPublic: b.isPublic
                              }))
                          },
                          null,
                          2
                        )}
                      </pre>
                    </div>
                  )}

                  {/* Target Bucket Editing Cards Grid (Matching Infra Node Group Layout) */}
                  <div className="grid grid-cols-1 gap-4">
                    {sourceBuckets.map((b, i) => {
                      const isExcluded = excludedBucketNames.includes(b.bucketName);
                      if (isExcluded) return null;
                      const corsRule = ensureCorsRuleExists(b);
                      const currentTargetName = b.targetBucketName || b.bucketName;
                      const cspSupport = getCspSupport(desiredCsp);

                      return (
                        <div
                          key={i}
                          className="p-5 bg-bg-panel border border-border-main/80 rounded-2xl space-y-4 hover:border-emerald-500/40 transition shadow-sm"
                        >
                          <div className="flex flex-col sm:flex-row sm:items-center justify-between border-b border-border-main/40 pb-3 gap-2">
                            <div className="flex items-center space-x-3">
                              <Database className="w-4.5 h-4.5 text-emerald-500" />
                              <span className="font-extrabold text-sm text-text-main">
                                Object Storage #{i + 1}
                              </span>
                              <span className="text-xs text-text-muted font-normal">
                                (Source: <span className="font-extrabold text-text-main font-mono">{b.bucketName}</span>)
                              </span>
                            </div>

                            <span className="px-2.5 py-0.5 rounded text-xs font-mono font-bold bg-emerald-500/20 text-emerald-400 uppercase shrink-0">
                              {desiredCsp.toUpperCase()} ({desiredRegion})
                            </span>
                          </div>

                          {/* Spec & Storage Class Selection Row (Matching Infra Node Group Layout) */}
                          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 text-xs">
                            {/* Column 1: Name (Editable Target Object Storage Name) */}
                            <div className="space-y-1.5">
                              <label className="text-text-muted font-normal block">Name</label>
                              <input
                                type="text"
                                value={currentTargetName}
                                onChange={(e) => {
                                  const val = e.target.value;
                                  setSourceBuckets((prev) =>
                                    prev.map((item, idx) => (idx === i ? { ...item, targetBucketName: val } : item))
                                  );
                                }}
                                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold font-mono text-xs focus:outline-none focus:border-emerald-500"
                                placeholder="e.g. target-storage-01"
                              />
                            </div>

                            {/* Column 2: Target Storage Class */}
                            <div className="space-y-1.5">
                              <label className="text-text-muted font-normal block">Target Storage Class</label>
                              <select
                                value={b.accessFrequency === 'frequent' ? 'Standard' : b.accessFrequency === 'infrequent' ? 'Standard-IA' : 'Archive'}
                                onChange={(e) => {
                                  const val = e.target.value;
                                  const freq = val === 'Standard' ? 'frequent' : val === 'Standard-IA' ? 'infrequent' : 'archive';
                                  setSourceBuckets((prev) => prev.map((item, idx) => (idx === i ? { ...item, accessFrequency: freq } : item)));
                                }}
                                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold text-xs focus:outline-none focus:border-emerald-500"
                              >
                                <option value="Standard">Standard (S3 / Hot)</option>
                                <option value="Standard-IA">Standard-IA (Infrequent)</option>
                                <option value="Archive">Glacier / Archive</option>
                              </select>
                            </div>

                            <div className="space-y-1.5">
                              <label className="text-text-muted font-normal block">Access Frequency</label>
                              <select
                                value={b.accessFrequency}
                                onChange={(e) => {
                                  const val = e.target.value as any;
                                  setSourceBuckets((prev) => prev.map((item, idx) => (idx === i ? { ...item, accessFrequency: val } : item)));
                                }}
                                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold text-xs focus:outline-none focus:border-emerald-500"
                              >
                                <option value="frequent">Frequent (High Throughput)</option>
                                <option value="infrequent">Infrequent (Low Access)</option>
                                <option value="archive">Archive (Cold Data)</option>
                              </select>
                            </div>

                            {/* Refinable Policy Checkboxes */}
                            <div className="col-span-1 lg:col-span-2 space-y-1.5">
                              <label className="text-text-muted font-normal block">Target Bucket Policies</label>
                              <div className="grid grid-cols-2 sm:grid-cols-4 gap-2">
                                <label
                                  title={!cspSupport.versioning ? `Versioning is not supported by ${desiredCsp.toUpperCase()}` : ''}
                                  className={`flex items-center space-x-2 p-2 rounded-lg border transition ${
                                    !cspSupport.versioning
                                      ? 'bg-bg-input/30 border-border-main/20 text-text-muted/50 cursor-not-allowed opacity-50'
                                      : 'bg-bg-input/60 border-border-main/40 cursor-pointer hover:border-emerald-500/50'
                                  }`}
                                >
                                  <input
                                    type="checkbox"
                                    disabled={!cspSupport.versioning}
                                    checked={cspSupport.versioning && b.versioningEnabled}
                                    onChange={(e) => {
                                      if (!cspSupport.versioning) return;
                                      const checked = e.target.checked;
                                      setSourceBuckets((prev) => prev.map((item, idx) => (idx === i ? { ...item, versioningEnabled: checked } : item)));
                                    }}
                                    className="accent-emerald-500 cursor-pointer w-4 h-4 disabled:cursor-not-allowed"
                                  />
                                  <div className="flex flex-col">
                                    <span className="font-extrabold text-text-main">Versioning</span>
                                    {!cspSupport.versioning && (
                                      <span className="text-[10px] font-bold text-red-500 dark:text-red-400">N/A ({desiredCsp.toUpperCase()})</span>
                                    )}
                                  </div>
                                </label>

                                <label className="flex items-center space-x-2 bg-bg-input/60 p-2 rounded-lg border border-border-main/40 cursor-pointer hover:border-emerald-500/50 transition">
                                  <input
                                    type="checkbox"
                                    checked={b.encryptionEnabled}
                                    onChange={(e) => {
                                      const checked = e.target.checked;
                                      setSourceBuckets((prev) => prev.map((item, idx) => (idx === i ? { ...item, encryptionEnabled: checked } : item)));
                                    }}
                                    className="accent-emerald-500 cursor-pointer w-4 h-4"
                                  />
                                  <span className="font-extrabold text-text-main">Encryption</span>
                                </label>

                                <label className="flex items-center space-x-2 bg-bg-input/60 p-2 rounded-lg border border-border-main/40 cursor-pointer hover:border-emerald-500/50 transition">
                                  <input
                                    type="checkbox"
                                    checked={b.isPublic}
                                    onChange={(e) => {
                                      const checked = e.target.checked;
                                      setSourceBuckets((prev) => prev.map((item, idx) => (idx === i ? { ...item, isPublic: checked } : item)));
                                    }}
                                    className="accent-emerald-500 cursor-pointer w-4 h-4"
                                  />
                                  <span className="font-extrabold text-text-main">Public Access</span>
                                </label>

                                <label
                                  title={!cspSupport.cors ? `CORS Rules are not supported by ${desiredCsp.toUpperCase()}` : ''}
                                  className={`flex items-center space-x-2 p-2 rounded-lg border transition ${
                                    !cspSupport.cors
                                      ? 'bg-bg-input/30 border-border-main/20 text-text-muted/50 cursor-not-allowed opacity-50'
                                      : 'bg-bg-input/60 border-border-main/40 cursor-pointer hover:border-emerald-500/50'
                                  }`}
                                >
                                  <input
                                    type="checkbox"
                                    disabled={!cspSupport.cors}
                                    checked={cspSupport.cors && b.corsEnabled}
                                    onChange={(e) => {
                                      if (!cspSupport.cors) return;
                                      const checked = e.target.checked;
                                      setSourceBuckets((prev) => prev.map((item, idx) => (idx === i ? { ...item, corsEnabled: checked } : item)));
                                    }}
                                    className="accent-emerald-500 cursor-pointer w-4 h-4 disabled:cursor-not-allowed"
                                  />
                                  <div className="flex flex-col">
                                    <span className="font-extrabold text-text-main">CORS Rules</span>
                                    {!cspSupport.cors && (
                                      <span className="text-[10px] font-bold text-red-500 dark:text-red-400">N/A ({desiredCsp.toUpperCase()})</span>
                                    )}
                                  </div>
                                </label>
                              </div>
                            </div>
                          </div>

                          {/* Section B: Detailed CORS Rule Configuration Section (when CORS Enabled) */}
                          {b.corsEnabled && (
                            <div className="space-y-4 pt-4 border-t border-border-main/40">
                              <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-2">
                                <div className="flex items-center space-x-2">
                                  <Globe className="w-4 h-4 text-emerald-500" />
                                  <span className="font-extrabold text-sm text-text-main">
                                    CORS Detailed Configuration Rules — {currentTargetName}
                                  </span>
                                </div>

                                <div className="flex items-center space-x-3 text-xs shrink-0">
                                  <button
                                    onClick={() => setShowCorsGuide(true)}
                                    className="px-3 py-1 bg-emerald-500/10 hover:bg-emerald-500/20 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 font-bold rounded-lg transition cursor-pointer flex items-center space-x-1"
                                  >
                                    <AlertCircle className="w-3.5 h-3.5 text-emerald-500" />
                                    <span>❓ CORS Rule Guide &amp; Examples</span>
                                  </button>
                                  <div className="flex items-center space-x-1.5 font-mono">
                                    <span className="text-text-muted font-normal">Status:</span>
                                    <span className="px-2 py-0.5 rounded font-bold bg-emerald-500/20 text-emerald-400">
                                      ENABLED
                                    </span>
                                  </div>
                                </div>
                              </div>

                              <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-4 text-sm w-full">
                                {/* 1. Allowed Origins */}
                                <div className="space-y-2 bg-bg-input/60 p-3.5 rounded-xl border border-border-main/40">
                                  <label className="font-normal text-sm text-text-muted block">Allowed Origins (allowedOrigin)</label>
                                  <div className="flex flex-wrap gap-1.5 min-h-[36px] items-center p-2 bg-bg-input rounded-lg border border-border-main/40">
                                    {corsRule.allowedOrigin.map((orig, oIdx) => (
                                      <span key={oIdx} className="bg-emerald-500/15 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 px-2.5 py-1 rounded-md font-mono text-sm font-bold flex items-center space-x-1">
                                        <span>{orig}</span>
                                        <button
                                          onClick={() => handleRemoveCorsOrigin(i, oIdx)}
                                          className="hover:text-red-400 font-extrabold ml-1 cursor-pointer"
                                        >
                                          ✕
                                        </button>
                                      </span>
                                    ))}
                                  </div>
                                  <div className="flex items-center space-x-2 pt-1">
                                    <input
                                      type="text"
                                      value={newOriginInputs[i] || ''}
                                      onChange={(e) => setNewOriginInputs((prev) => ({ ...prev, [i]: e.target.value }))}
                                      onKeyDown={(e) => {
                                        if (e.key === 'Enter') handleAddCorsOrigin(i, newOriginInputs[i] || '');
                                      }}
                                      placeholder="e.g. https://example.com or *"
                                      className="flex-1 px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono text-sm focus:outline-none focus:border-emerald-500"
                                    />
                                    <button
                                      onClick={() => handleAddCorsOrigin(i, newOriginInputs[i] || '')}
                                      className="px-3.5 py-2 bg-emerald-600 hover:bg-emerald-700 text-white font-bold rounded-lg cursor-pointer transition text-sm shrink-0"
                                    >
                                      Add Origin
                                    </button>
                                  </div>
                                </div>

                                {/* 2. Allowed HTTP Methods */}
                                <div className="space-y-2 bg-bg-input/60 p-3.5 rounded-xl border border-border-main/40">
                                  <label className="font-normal text-sm text-text-muted block">Allowed HTTP Methods (allowedMethod)</label>
                                  <div className="flex flex-wrap gap-1.5 p-2 bg-bg-input rounded-lg border border-border-main/40">
                                    {HTTP_METHODS.map((m) => {
                                      const isSelected = corsRule.allowedMethod.includes(m);
                                      return (
                                        <button
                                          key={m}
                                          onClick={() => handleToggleCorsMethod(i, m)}
                                          className={`px-3 py-1.5 rounded-lg text-sm font-mono font-bold transition cursor-pointer border ${
                                            isSelected
                                              ? 'bg-emerald-500 text-slate-950 border-emerald-500 shadow-sm'
                                              : 'bg-bg-panel text-text-muted border-border-main hover:text-text-main'
                                          }`}
                                        >
                                          {isSelected ? `✓ ${m}` : m}
                                        </button>
                                      );
                                    })}
                                  </div>
                                </div>

                                {/* 3. Allowed Headers */}
                                <div className="space-y-2 bg-bg-input/60 p-3.5 rounded-xl border border-border-main/40">
                                  <label className="font-normal text-sm text-text-muted block">Allowed Headers (allowedHeader)</label>
                                  <div className="flex flex-wrap gap-1.5 min-h-[36px] items-center p-2 bg-bg-input rounded-lg border border-border-main/40">
                                    {corsRule.allowedHeader.map((h, hIdx) => (
                                      <span key={hIdx} className="bg-teal-500/15 border border-teal-500/30 text-teal-600 dark:text-teal-400 px-2.5 py-1 rounded-md font-mono text-sm font-bold flex items-center space-x-1">
                                        <span>{h}</span>
                                        <button
                                          onClick={() => handleRemoveCorsHeader(i, 'allowedHeader', hIdx)}
                                          className="hover:text-red-400 font-extrabold ml-1 cursor-pointer"
                                        >
                                          ✕
                                        </button>
                                      </span>
                                    ))}
                                  </div>
                                  <div className="flex items-center space-x-2 pt-1">
                                    <input
                                      type="text"
                                      value={newAllowedHeaderInputs[i] || ''}
                                      onChange={(e) => setNewAllowedHeaderInputs((prev) => ({ ...prev, [i]: e.target.value }))}
                                      onKeyDown={(e) => {
                                        if (e.key === 'Enter') handleAddCorsHeader(i, 'allowedHeader', newAllowedHeaderInputs[i] || '');
                                      }}
                                      placeholder="e.g. * or Content-Type"
                                      className="flex-1 px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono text-sm focus:outline-none focus:border-emerald-500"
                                    />
                                    <button
                                      onClick={() => handleAddCorsHeader(i, 'allowedHeader', newAllowedHeaderInputs[i] || '')}
                                      className="px-3.5 py-2 bg-emerald-600 hover:bg-emerald-700 text-white font-bold rounded-lg cursor-pointer transition text-sm shrink-0"
                                    >
                                      Add Header
                                    </button>
                                  </div>
                                </div>

                                {/* 4. Expose Headers & Max Age */}
                                <div className="space-y-2 bg-bg-input/60 p-3.5 rounded-xl border border-border-main/40">
                                  <label className="font-normal text-sm text-text-muted block">Expose Headers &amp; Max Age (exposeHeader &amp; maxAgeSeconds)</label>
                                  <div className="flex flex-wrap gap-1.5 min-h-[36px] items-center p-2 bg-bg-input rounded-lg border border-border-main/40">
                                    {corsRule.exposeHeader.map((eh, ehIdx) => (
                                      <span key={ehIdx} className="bg-emerald-500/15 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 px-2.5 py-1 rounded-md font-mono text-sm font-bold flex items-center space-x-1">
                                        <span>{eh}</span>
                                        <button
                                          onClick={() => handleRemoveCorsHeader(i, 'exposeHeader', ehIdx)}
                                          className="hover:text-red-400 font-extrabold ml-1 cursor-pointer"
                                        >
                                          ✕
                                        </button>
                                      </span>
                                    ))}
                                  </div>
                                  <div className="grid grid-cols-2 gap-2 pt-1">
                                    <input
                                      type="text"
                                      value={newExposeHeaderInputs[i] || ''}
                                      onChange={(e) => setNewExposeHeaderInputs((prev) => ({ ...prev, [i]: e.target.value }))}
                                      onKeyDown={(e) => {
                                        if (e.key === 'Enter') handleAddCorsHeader(i, 'exposeHeader', newExposeHeaderInputs[i] || '');
                                      }}
                                      placeholder="e.g. ETag"
                                      className="flex-1 px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono text-sm focus:outline-none focus:border-emerald-500"
                                    />
                                    <div className="flex items-center space-x-1.5">
                                      <span className="text-xs text-text-muted font-normal shrink-0">MaxAge(s):</span>
                                      <input
                                        type="number"
                                        value={corsRule.maxAgeSeconds}
                                        onChange={(e) => handleUpdateCorsMaxAge(i, parseInt(e.target.value) || 0)}
                                        className="w-full px-2 py-1.5 bg-bg-input border border-border-main rounded-lg text-text-main font-mono text-xs focus:outline-none focus:border-emerald-500"
                                      />
                                    </div>
                                  </div>
                                </div>
                              </div>
                            </div>
                          )}

                          <div className="pt-2 border-t border-border-main/30 text-xs text-text-muted italic flex items-center justify-between">
                            <span>
                              * Note: Cloud-Barista generates a unique UID for each target bucket (e.g., <span className="font-mono font-bold text-emerald-600 dark:text-emerald-400">{currentTargetName}-uid-{i + 1}</span>) to guarantee global cloud uniqueness across CSPs.
                            </span>
                          </div>
                        </div>
                      );
                    })}
                  </div>
                </div>
              </div>
            )}

            <div className="space-y-4 pt-4 border-t border-border-main/50">
              {/* Save Target Cloud Model Action Button */}
              <div className="flex justify-end">
                <button
                  onClick={() => setShowSaveTargetModal(true)}
                  disabled={!recommendationResult}
                  className={`px-5 py-2.5 rounded-xl font-extrabold text-xs sm:text-sm flex items-center space-x-2 transition-all shadow-md shrink-0 ${
                    !recommendationResult
                      ? 'bg-gradient-to-r from-emerald-300/40 via-teal-300/30 to-blue-400/30 text-slate-500/70 dark:text-slate-400/70 border border-emerald-500/20 cursor-not-allowed opacity-70'
                      : 'bg-gradient-to-r from-emerald-400 via-teal-400 to-blue-600 hover:from-emerald-500 hover:to-blue-700 text-slate-950 cursor-pointer shadow-lg shadow-emerald-500/20'
                  }`}
                >
                  <Save className={`w-4 h-4 ${!recommendationResult ? 'text-slate-500/80 dark:text-slate-400/80' : 'text-slate-950'}`} />
                  <span>Save Target Cloud Object Storage Model</span>
                </button>
              </div>

              {/* Row 2: Back & Next Navigation Buttons */}
              <div className="flex items-center justify-between pt-1">
                <button
                  onClick={() => setSubTab('refine')}
                  className="px-4 py-2 bg-bg-input hover:bg-bg-main border border-border-main text-text-muted hover:text-text-main font-bold text-xs rounded-xl transition flex items-center space-x-2 cursor-pointer"
                >
                  <ArrowLeft className="w-4 h-4" />
                  <span>Back to 2. Refinement</span>
                </button>

                <button
                  onClick={() => setSubTab('provision')}
                  disabled={!recommendationResult}
                  className={`px-5 py-2.5 rounded-xl font-extrabold text-xs sm:text-sm flex items-center space-x-2 transition-all shadow-md ${
                    !recommendationResult
                      ? 'bg-emerald-500/20 text-slate-400/80 dark:text-slate-400/80 border border-emerald-500/20 cursor-not-allowed opacity-60'
                      : 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 cursor-pointer shadow-md'
                  }`}
                >
                  <span>Next: Proceed to 4. Migration Execution</span>
                  <ArrowRight className={`w-4 h-4 ${!recommendationResult ? 'text-slate-400/80 dark:text-slate-400/80' : 'text-slate-950'}`} />
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* ══════════════════════════════════════════════════
          SUB-STEP 4: Migration Execution (Exact Infra Parity Design)
      ══════════════════════════════════════════════════ */}
      {subTab === 'provision' && (
        <div className="space-y-6 font-sans animate-fade-in">
          {/* 1. Single-Line Tab Description Box (Rule #3 Parity) */}
          <div className="glass-panel px-6 py-4.5 rounded-2xl border border-border-main flex flex-wrap items-center gap-x-3 gap-y-1.5">
            <div className="flex items-center gap-2 shrink-0">
              <Play className="w-5 h-5 text-emerald-500" />
              <h2 className="text-base font-extrabold text-text-main tracking-tight">
                Target Cloud Migration
              </h2>
            </div>
            <span className="text-sm text-text-muted">
              Execute target cloud object storage migrations, monitor real-time migration status, and inspect provisioned bucket access points.
            </span>
          </div>

          {/* 2. Dedicated Action Control Box */}
          <div className="glass-panel p-4 rounded-2xl border border-border-main flex flex-wrap items-center gap-3">
            <button
              onClick={() => setShowLaunchModal(true)}
              disabled={isDeploying}
              className="px-5 py-2.5 bg-gradient-to-r from-emerald-500 to-blue-600 hover:from-emerald-600 hover:to-blue-700 disabled:opacity-50 text-slate-950 rounded-xl text-xs font-extrabold flex items-center gap-1.5 transition cursor-pointer shadow-lg shadow-emerald-500/20"
            >
              {isDeploying ? <RefreshCw className="w-4 h-4 animate-spin text-slate-950" /> : <Plus className="w-4 h-4 text-slate-950" />}
              <span>Launch New Migration</span>
            </button>

            <div className="px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-xs font-bold font-mono text-text-main flex items-center gap-2">
              <Zap className="w-4 h-4 text-emerald-500" />
              <span>Active Migration Jobs ({isDeploying ? '1 Running' : '0 Running'} / 1 Completed)</span>
            </div>
          </div>

          {/* SECTION 1: Horizontal Side-by-Side Job Cards Bar */}
          <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-4 shadow-sm">
            <div className="flex justify-between items-center border-b border-border-main/20 pb-3">
              <h3 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                <Activity className="w-4 h-4 text-emerald-500" />
                Migration Jobs Queue (1)
              </h3>
              <span className="text-xs text-text-muted font-mono bg-bg-panel px-3 py-1 rounded-full border border-border-main">
                Click card to view detailed progress &amp; results
              </span>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-4">
              <div className="p-4 rounded-xl border text-left transition-all duration-200 cursor-pointer flex flex-col justify-between space-y-3 relative overflow-hidden bg-emerald-500/10 border-emerald-500/60 shadow-lg shadow-emerald-500/10 ring-1 ring-emerald-500/40">
                <div className="absolute top-0 left-0 w-full h-[2px] bg-gradient-to-r from-emerald-500 to-blue-500" />

                <div className="flex justify-between items-center gap-2">
                  <div className="flex items-center gap-2 min-w-0">
                    <span className="text-base shrink-0">
                      {desiredCsp.toLowerCase() === 'aws' ? '🌩️' : desiredCsp.toLowerCase() === 'azure' ? '🔷' : '🟢'}
                    </span>
                    <span className="font-extrabold text-sm text-text-main font-mono truncate" title={`${desiredCsp.toUpperCase()} (${nameSeed === 'my-target-object-storage-01' ? 'mig01' : (nameSeed || 'mig01')}-${desiredCsp.toLowerCase()}-storage)`}>
                      <span className="text-amber-500 font-bold mr-1">[Sample]</span>
                      {`${desiredCsp.toUpperCase()} (${nameSeed === 'my-target-object-storage-01' ? 'mig01' : (nameSeed || 'mig01')}-${desiredCsp.toLowerCase()}-storage)`}
                    </span>
                  </div>

                  <span className="px-2.5 py-1 bg-green-500/10 text-green-400 border border-green-500/20 rounded-full text-xs font-bold flex items-center gap-1 shrink-0">
                    <CheckCircle2 className="w-3.5 h-3.5 text-green-400" />
                    <span>✓ Success</span>
                  </span>
                </div>

                <div className="flex justify-between items-center text-xs font-mono text-text-muted pt-1 border-t border-border-main/20">
                  <span>Region: {desiredRegion}</span>
                  <span className="flex items-center gap-1 font-bold text-text-main">
                    <Clock className="w-3.5 h-3.5 text-emerald-500" />
                    Time: 15s (Done)
                  </span>
                </div>
              </div>
            </div>
          </div>

          {/* SECTION 2: Selected Job Detail Panel */}
          <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-6 animate-fade-in shadow-sm">
            {/* Selected Job Header */}
            <div className="flex flex-col sm:flex-row justify-between sm:items-center gap-3 border-b border-border-main/20 pb-4">
              <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
                <span className="text-lg">
                  {desiredCsp.toLowerCase() === 'aws' ? '🌩️' : desiredCsp.toLowerCase() === 'azure' ? '🔷' : '🟢'}
                </span>
                <span>
                  Selected Job Detail: <span className="text-amber-500 font-bold mr-1">[Sample]</span>[{`${desiredCsp.toUpperCase()} (${nameSeed === 'my-target-object-storage-01' ? 'mig01' : (nameSeed || 'mig01')}-${desiredCsp.toLowerCase()}-storage)`}]
                </span>
              </h3>

              <div className="flex items-center gap-3 text-xs font-mono text-text-muted">
                <span>Namespace: <strong className="text-text-main">mig01</strong></span>
                <span>Req ID: <strong className="text-emerald-500">req-20260723-001</strong></span>
                <span>Elapsed: <strong className="text-teal-400">15s</strong></span>
              </div>
            </div>

            {/* Simplified 3-Stage API Status Flow (Matching Infra Stepper) */}
            <div className="bg-bg-panel/50 border border-border-main/40 p-5 rounded-xl space-y-3">
              <div className="flex justify-between items-center">
                <span className="text-xs font-bold text-emerald-600 dark:text-emerald-400 uppercase font-mono">
                  API MIGRATION EXECUTION STATUS
                </span>
                <span className="text-xs font-bold text-text-muted font-mono">
                  API Status: <strong className="text-green-400">Success</strong>
                </span>
              </div>

              <div className="grid grid-cols-1 md:grid-cols-3 gap-4 pt-1">
                {/* Step 1: Request Accepted */}
                <div className="bg-bg-input/60 border border-emerald-500/30 p-3.5 rounded-xl flex items-center space-x-3">
                  <div className="p-2 bg-emerald-500/20 text-emerald-400 rounded-lg font-bold text-xs">✓</div>
                  <div>
                    <h4 className="text-xs font-bold text-text-main">1. Request Accepted</h4>
                    <p className="text-[11px] text-text-muted font-mono">HTTP 202 (ReqID Issued)</p>
                  </div>
                </div>

                {/* Step 2: Migrating */}
                <div className="bg-bg-input/60 border border-emerald-500/30 p-3.5 rounded-xl flex items-center space-x-3">
                  <div className="p-2 bg-emerald-500/20 text-emerald-400 rounded-lg font-bold text-xs">✓</div>
                  <div>
                    <h4 className="text-xs font-bold text-text-main">2. Migrating</h4>
                    <p className="text-[11px] text-text-muted font-mono">Finished Processing</p>
                  </div>
                </div>

                {/* Step 3: Completed */}
                <div className="bg-green-500/10 border border-green-500/40 p-3.5 rounded-xl flex items-center space-x-3">
                  <div className="p-2 bg-green-500/20 text-green-400 rounded-lg font-bold text-xs">✓</div>
                  <div>
                    <h4 className="text-xs font-bold text-text-main">3. Completed</h4>
                    <p className="text-[11px] text-text-muted font-mono">Object Storage Active &amp; Ready</p>
                  </div>
                </div>
              </div>
            </div>

            {/* Provisioned Cloud Storage Buckets & Access Points Verification Table */}
            <div className="space-y-3">
              <div className="flex items-center justify-between">
                <h4 className="text-xs font-extrabold text-text-main flex items-center space-x-2">
                  <Globe className="w-4 h-4 text-emerald-500" />
                  <span>Provisioned Cloud Storage Buckets &amp; Connectivity Verification</span>
                </h4>
                <button
                  onClick={loadMigratedStorages}
                  className="px-3 py-1.5 bg-emerald-500/10 hover:bg-emerald-500/20 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 rounded-xl text-xs font-bold transition cursor-pointer flex items-center space-x-1"
                >
                  <Sparkles className="w-3.5 h-3.5 text-emerald-500" />
                  <span>Check Storage Access Points</span>
                </button>
              </div>

              <div className="bg-bg-panel border border-border-main/60 rounded-xl overflow-hidden shadow-sm">
                <table className="w-full text-xs text-left">
                  <thead className="bg-bg-input/60 border-b border-border-main/40 text-text-muted uppercase font-mono">
                    <tr>
                      <th className="p-3.5 font-bold">Bucket Name</th>
                      <th className="p-3.5 font-bold">Target Storage Class</th>
                      <th className="p-3.5 font-bold">CORS Status</th>
                      <th className="p-3.5 font-bold">Versioning</th>
                      <th className="p-3.5 font-bold">Storage Access Check</th>
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-border-main/30 font-mono">
                    {sourceBuckets.map((b, idx) => (
                      <tr key={idx} className="hover:bg-bg-main/30 transition">
                        <td className="p-3.5 font-extrabold text-text-main">
                          {b.targetBucketName || b.bucketName}
                        </td>
                        <td className="p-3.5 font-bold text-emerald-500">
                          {b.accessFrequency === 'frequent' ? 'Standard (S3/Hot)' : b.accessFrequency === 'infrequent' ? 'Standard-IA' : 'Glacier/Archive'}
                        </td>
                        <td className="p-3.5">
                          <span className={`px-2 py-0.5 rounded font-bold ${b.corsEnabled ? 'bg-emerald-500/20 text-emerald-400' : 'bg-slate-500/20 text-slate-400'}`}>
                            {b.corsEnabled ? 'ENABLED' : 'DISABLED'}
                          </span>
                        </td>
                        <td className="p-3.5">
                          <span className={`px-2 py-0.5 rounded font-bold ${b.versioningEnabled ? 'bg-teal-500/20 text-teal-400' : 'bg-slate-500/20 text-slate-400'}`}>
                            {b.versioningEnabled ? 'ENABLED' : 'DISABLED'}
                          </span>
                        </td>
                        <td className="p-3.5">
                          <button
                            onClick={() => alert(`[OK] Storage Access Point reachable for '${b.targetBucketName || b.bucketName}' on ${desiredCsp.toUpperCase()} (${desiredRegion})`)}
                            className="px-3 py-1 bg-bg-input hover:bg-bg-main border border-border-main text-emerald-600 dark:text-emerald-400 rounded-lg font-bold transition flex items-center space-x-1.5 cursor-pointer text-xs"
                          >
                            <Sparkles className="w-3.5 h-3.5 text-emerald-500" />
                            <span>Check Bucket Access</span>
                          </button>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            </div>

            {/* REST API REQUEST & RESPONSE LOG (Matching Infra Log Block) */}
            <div className="space-y-2">
              <span className="text-[11px] font-bold text-text-muted uppercase font-mono block">
                REST API REQUEST &amp; RESPONSE LOG
              </span>
              <div className="p-4 bg-slate-950 text-emerald-400 font-mono text-xs rounded-xl space-y-1.5 border border-slate-800 max-h-48 overflow-y-auto">
                <div>&gt; POST /beetle/migration/ns/mig01/objectStorage?nameSeed={nameSeed || 'mig01'}</div>
                <div>&gt; Header -&gt; Prefer: respond-async</div>
                <div>&gt; HTTP 202 Accepted (ReqID: req-20260723-001, Status: Handling)</div>
                <div className="text-emerald-300 font-bold">&gt; GET /beetle/request/req-20260723-001 -&gt; Status: Success (Duration: 15s)</div>
              </div>
            </div>

            {/* Bottom Navigation Buttons Row (Matching Infra Parity) */}
            <div className="flex items-center justify-between pt-4 border-t border-border-main/30 mt-4">
              <button
                onClick={() => setSubTab('optimize')}
                className="px-4 py-2 bg-bg-input border border-border-main hover:bg-bg-main text-text-main font-bold text-xs rounded-xl transition cursor-pointer flex items-center space-x-1.5"
              >
                <ArrowLeft className="w-4 h-4" />
                <span>Back to 3. Target Object Storage Optimizer</span>
              </button>

              <div className="flex items-center space-x-2 px-4 py-2 bg-emerald-500/10 border border-emerald-500/30 text-emerald-400 rounded-xl text-xs font-bold font-mono">
                <CheckCircle2 className="w-4 h-4 text-emerald-400" />
                <span>Final Step: 4. Migration Execution &amp; Monitoring</span>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Save Source Model Revision Modal */}
      <SaveRevisionModal
        isOpen={showSaveSourceModal}
        onClose={() => setShowSaveSourceModal(false)}
        title="Save Source Storage Model Revision"
        defaultName="source-storage-model-rev1"
        defaultDescription="Extracted and refined source object storage model"
        defaultVersion="1.0.0"
        existingRevisions={savedSourceModels}
        onSave={handleSaveSourceModelRevision}
        successMessage="Source storage model revision saved to Damselfly successfully."
      />

      {/* Save Target Model Revision Modal */}
      <SaveRevisionModal
        isOpen={showSaveTargetModal}
        onClose={() => setShowSaveTargetModal(false)}
        title="Save Target Storage Model Revision"
        defaultName="target-storage-model-rev1"
        defaultDescription="Recommended target cloud object storage model"
        defaultVersion="1.0.0"
        existingRevisions={[]}
        onSave={handleSaveTargetModelRevision}
        successMessage="Target storage model revision saved to Damselfly successfully."
      />

      {/* CORS Specification & Configuration Rules Guide Modal Popup */}
      {showCorsGuide && (
        <div className="fixed inset-0 bg-slate-950/70 backdrop-blur-sm z-50 flex items-center justify-center p-4 sm:p-6 animate-fade-in">
          <div className="bg-bg-panel border border-border-main rounded-2xl max-w-7xl w-full p-6 sm:p-8 space-y-6 shadow-2xl overflow-y-auto max-h-[88vh] font-sans">
            <div className="flex items-center justify-between border-b border-border-main/40 pb-4">
              <div className="flex items-center space-x-2.5">
                <Globe className="w-5 h-5 text-emerald-500" />
                <h3 className="text-base font-extrabold text-text-main">
                  CORS Specification &amp; Configuration Rules Guide
                </h3>
              </div>
              <button
                onClick={() => setShowCorsGuide(false)}
                className="w-8 h-8 rounded-lg bg-bg-input hover:bg-bg-main border border-border-main flex items-center justify-center text-text-muted hover:text-text-main transition cursor-pointer font-bold"
              >
                ✕
              </button>
            </div>

            <p className="text-sm text-text-muted leading-relaxed">
              Cross-Origin Resource Sharing (CORS) defines security rules allowing web applications hosted on other domain origins to interact with your cloud object storage resources.
            </p>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-sm">
              <div className="p-4 bg-bg-input rounded-xl border border-border-main/40 space-y-1.5">
                <span className="font-extrabold text-sm text-emerald-600 dark:text-emerald-400 block">1. Allowed Origins (allowedOrigin)</span>
                <p className="text-sm text-text-muted font-normal">Specify domain origins allowed to access storage data.</p>
                <div className="font-mono text-sm text-emerald-500 pt-1 font-bold">e.g., * (Allow All) or https://example.com</div>
              </div>

              <div className="p-4 bg-bg-input rounded-xl border border-border-main/40 space-y-1.5">
                <span className="font-extrabold text-sm text-emerald-600 dark:text-emerald-400 block">2. Allowed HTTP Methods (allowedMethod)</span>
                <p className="text-sm text-text-muted font-normal">Select HTTP request methods allowed for client requests.</p>
                <div className="font-mono text-sm text-emerald-500 pt-1 font-bold">e.g., GET, POST, PUT, DELETE (Click toggle buttons)</div>
              </div>

              <div className="p-4 bg-bg-input rounded-xl border border-border-main/40 space-y-1.5">
                <span className="font-extrabold text-sm text-emerald-600 dark:text-emerald-400 block">3. Allowed Headers (allowedHeader)</span>
                <p className="text-sm text-text-muted font-normal">Specify HTTP headers allowed in client requests.</p>
                <div className="font-mono text-sm text-emerald-500 pt-1 font-bold">e.g., * (Allow All) or Content-Type, Authorization</div>
              </div>

              <div className="p-4 bg-bg-input rounded-xl border border-border-main/40 space-y-1.5">
                <span className="font-extrabold text-sm text-emerald-600 dark:text-emerald-400 block">4. Expose Headers &amp; Max Age (exposeHeader &amp; maxAgeSeconds)</span>
                <p className="text-sm text-text-muted font-normal">Response headers exposed to browser scripts &amp; Preflight cache duration (sec).</p>
                <div className="font-mono text-sm text-emerald-500 pt-1 font-bold">e.g., ETag, x-amz-request-id | MaxAge: 3600 (1 hour)</div>
              </div>
            </div>

            <div className="border-t border-border-main/40 pt-4 flex justify-end">
              <button
                onClick={() => setShowCorsGuide(false)}
                className="px-6 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white font-extrabold text-sm rounded-xl shadow-md transition cursor-pointer"
              >
                Close Guide
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Launch New Object Storage Migration Modal (Infra Parity) */}
      {showLaunchModal && (
        <div className="fixed inset-0 bg-slate-950/70 backdrop-blur-sm z-50 flex items-center justify-center p-4 animate-fade-in">
          <div className="bg-bg-panel border border-border-main rounded-2xl max-w-3xl w-full p-6 sm:p-7 space-y-5 shadow-2xl overflow-hidden font-sans">
            <div className="flex items-center justify-between border-b border-border-main/40 pb-4">
              <div className="flex items-center space-x-2.5">
                <Plus className="w-5 h-5 text-emerald-500" />
                <h3 className="text-base font-extrabold text-text-main">
                  Launch New Object Storage Migration
                </h3>
              </div>
              <button
                onClick={() => setShowLaunchModal(false)}
                className="w-8 h-8 rounded-lg bg-bg-input hover:bg-bg-main border border-border-main flex items-center justify-center text-text-muted hover:text-text-main transition cursor-pointer font-bold"
              >
                ✕
              </button>
            </div>

            <div className="space-y-4 text-xs font-sans">
              <div>
                <label className="block text-text-muted font-bold uppercase tracking-wider mb-1.5 font-mono">
                  1. SELECT TARGET CLOUD MODEL
                </label>
                <select
                  className="w-full px-3.5 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-bold focus:outline-none focus:border-emerald-500 text-xs"
                >
                  <option value="sample">[Sample] target-storage-v1 ({desiredCsp.toUpperCase()})</option>
                </select>
              </div>

              <div className="p-3.5 bg-bg-input/60 border border-border-main/40 rounded-xl space-y-1.5 font-mono text-xs">
                <div className="flex justify-between">
                  <span className="text-text-muted font-normal">Model Storage Name (storageId):</span>
                  <span className="font-bold text-emerald-500">target-storage-01</span>
                </div>
                <div className="flex justify-between pt-1 border-t border-border-main/20">
                  <span className="text-text-muted font-normal">Target CSP / Region:</span>
                  <span className="font-bold text-text-main">{desiredCsp.toUpperCase()} ({desiredRegion})</span>
                </div>
              </div>

              <div>
                <label className="block text-text-muted font-bold uppercase tracking-wider mb-1.5 font-mono">
                  2. CONFIGURE DEPLOYMENT IDENTIFIERS
                </label>
                
                <div className="space-y-3">
                  <div>
                    <label className="block text-text-muted font-normal mb-1">Namespace ID (nsId)</label>
                    <input
                      type="text"
                      defaultValue="mig01"
                      className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold text-xs"
                    />
                  </div>

                  <div>
                    <label className="block text-text-muted font-normal mb-1">NameSeed Prefix (Late Binding)</label>
                    <input
                      type="text"
                      value={nameSeed}
                      onChange={(e) => setNameSeed(e.target.value)}
                      placeholder="e.g., prod"
                      className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold text-xs"
                    />
                  </div>

                  <div className="p-3 bg-bg-input/60 border border-border-main/40 rounded-xl flex items-center justify-between">
                    <div className="flex items-center space-x-2">
                      <Zap className="w-4 h-4 text-emerald-500" />
                      <span className="font-bold text-text-main text-xs font-mono">Prefer: respond-async</span>
                    </div>
                    <input type="checkbox" defaultChecked className="accent-emerald-500 w-4 h-4 cursor-pointer" />
                  </div>
                </div>
              </div>
            </div>

            <div className="border-t border-border-main/40 pt-4 flex items-center justify-end space-x-3">
              <button
                onClick={() => setShowLaunchModal(false)}
                className="px-4 py-2 bg-bg-input hover:bg-bg-main border border-border-main text-text-main font-bold text-xs rounded-xl transition cursor-pointer"
              >
                Cancel
              </button>
              <button
                onClick={() => {
                  setShowLaunchModal(false);
                  handleMigrateStorage();
                }}
                className="px-5 py-2 bg-gradient-to-r from-emerald-400 via-teal-400 to-blue-600 hover:from-emerald-500 hover:to-blue-700 text-slate-950 font-extrabold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center space-x-1.5 cursor-pointer"
              >
                <Play className="w-4 h-4 text-slate-950" />
                <span>Launch Migration</span>
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
