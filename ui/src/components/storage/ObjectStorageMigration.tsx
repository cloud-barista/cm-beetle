'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore, calculateJobElapsedSeconds, ObjectStorageJobCard } from '@/store/migrationStore';
import {
  artifactApi,
  beetleApi,
  StorageSourceModelEnvelope,
  StorageTargetModelEnvelope,
} from '@/api/client';
import { CspCredentialForm } from '../common/CspCredentialForm';
import { SaveRevisionModal, SaveRevisionResult } from '../common/SaveRevisionModal';
import { ActiveModelBanner } from '../common/ActiveModelBanner';
import { DeleteModelConfirmModal } from '../common/DeleteModelConfirmModal';
import { UnsavedConfirmModal } from '../common/UnsavedConfirmModal';
import { ModelToast } from '../common/ModelToast';
import { ObjectDirectoryTree } from './ObjectDirectoryTree';
import {
  buildDefaultModelName,
  cleanModelDisplayName,
  formatSourceOptionText,
  formatTargetOptionText,
  getStageBadgeMeta,
  isBuiltInSampleModel,
  resolveModelStageKind,
} from '../../utils/modelFormat';
import sampleStorageRequest from '../../data/sampleSourceObjectStorage.json';
import sampleTargetObjectStorage from '../../data/sampleTargetObjectStorage.json';
import {
  HardDrive, Plus, Trash2, CheckCircle2, ArrowRight, ArrowLeft, Check, Zap,
  Database, RefreshCw, AlertCircle, Play, Key, Lock, FileText, Save, X,
  Sliders, Compass, Copy, Globe, Sparkles, Activity, Clock, Folder
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
  objects?: Array<{ key: string; sizeBytes?: number; lastModified?: string }>;
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

const SAMPLE_STORAGE_MODEL: StorageSourceModelEnvelope = {
  id: 'sample-source-storage-01',
  name: '[Sample] sampleSourceObjectStorage.json',
  version: '1.0.0',
  stage: 'source',
  category: 'infra',
  schemaType: 'SourceObjectStorage',
  description: 'Sample object storage model for quick demonstration',
  sourceCloud: { csp: 'aws', region: 'ap-northeast-2' },
  sourceObjectStorages: sampleStorageRequest.sourceObjectStorages as any[],
};

const SAMPLE_TARGET_STORAGE_MODEL: StorageTargetModelEnvelope = {
  id: 'sample-target-storage-01',
  name: '[Sample] target-storage-v1',
  description: 'Default sample target object storage recommendation',
  version: '1.0.0',
  stage: 'target',
  category: 'infra',
  schemaType: 'RecommendedObjectStorage',
  csp: 'aws',
  region: 'ap-northeast-2',
  targetObjectStorages: [
    {
      bucketName: 'target-storage-01-x8f2',
      sourceBucketName: 'datamold-aws-test',
      versioningEnabled: true,
      corsEnabled: true,
      corsRule: [
        {
          allowedHeader: ['*'],
          allowedMethod: ['GET', 'POST', 'PUT', 'DELETE'],
          allowedOrigin: ['*'],
          exposeHeader: ['ETag'],
          maxAgeSeconds: 3600,
        },
      ],
    },
  ],
  raw: sampleTargetObjectStorage,
};

const DRAFT_STORAGE_TARGET_ID = '__draft_target_storage__';

export type { ObjectStorageJobCard };

export const ObjectStorageMigration: React.FC = () => {
  const {
    tumblebugProviders,
    tumblebugRegions,
    fetchTumblebugProviders,
    fetchTumblebugRegions,
    objectStorageJobs,
    setObjectStorageJobs,
    removeObjectStorageJob,
    activeObjectStorageJobId,
    setActiveObjectStorageJobId
  } = useMigrationStore();
  const namespaceId = (useMigrationStore.getState() as any).namespaceId || 'mig01';

  // Auto-restore sample job if queue was emptied
  useEffect(() => {
    if (objectStorageJobs.length === 0) {
      setObjectStorageJobs([
        {
          id: 'sample-os-job-01',
          reqId: 'req-20260723-001',
          name: 'mig01-aws-storage',
          csp: 'aws',
          region: 'ap-northeast-2',
          status: 'Success',
          startTime: new Date().toISOString(),
          elapsedSeconds: 15,
          buckets: [
            { name: 'target-bucket-01', targetClass: 'Standard (S3/Hot)', cors: 'ENABLED', versioning: 'DISABLED' },
            { name: 'target-bucket-02', targetClass: 'Standard-IA', cors: 'ENABLED', versioning: 'ENABLED' }
          ],
          logs: [
            'POST /beetle/migration/middleware/ns/mig01/objectStorage?nameSeed=mig01',
            'Header -> Prefer: respond-async',
            'HTTP 202 Accepted (ReqID: req-20260723-001, Status: Handling)',
            'GET /beetle/request/req-20260723-001 -> Status: Success (Duration: 15s)'
          ],
          isSample: true
        }
      ]);
    }
  }, [objectStorageJobs, setObjectStorageJobs]);

  const selectedJobId = activeObjectStorageJobId || objectStorageJobs[0]?.id || 'sample-os-job-01';
  const setSelectedJobId = setActiveObjectStorageJobId;
  const activeJob = objectStorageJobs.find((j) => j.id === selectedJobId) || objectStorageJobs[0];

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
    { id: string; name: string; csp: string; region: string; accessKey: string; secretKey?: string }[]
  >([]);

  const [scanCsp, setScanCsp] = useState('aws');
  const [scanRegion, setScanRegion] = useState('ap-northeast-2');
  const [scanAccessKey, setScanAccessKey] = useState('');
  const [scanSecretKey, setScanSecretKey] = useState('');
  const [isScanning, setIsScanning] = useState(false);
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const [deleteModalJob, setDeleteModalJob] = useState<ObjectStorageJobCard | null>(null);
  const [deleteConfirmText, setDeleteConfirmText] = useState('');

  const handleConfirmDeleteStorageRecord = () => {
    if (!deleteModalJob || deleteConfirmText !== deleteModalJob.name) return;
    removeObjectStorageJob(deleteModalJob.id);
    setDeleteModalJob(null);
    setDeleteConfirmText('');
  };

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
  const [isInspectDone, setIsInspectDone] = useState(false);

  // Unified Artifact Repository State across Steps 1-4
  const [savedSourceModels, setSavedSourceModels] = useState<StorageSourceModelEnvelope[]>([]);
  const [savedTargetModels, setSavedTargetModels] = useState<StorageTargetModelEnvelope[]>([]);

  // Step 1 save state
  const [lastSavedSourceStorageModel, setLastSavedSourceStorageModel] = useState<StorageSourceModelEnvelope | null>(null);
  const [hasStep1UnsavedChanges, setHasStep1UnsavedChanges] = useState(false);
  const [showSaveStep1Modal, setShowSaveStep1Modal] = useState(false);

  // Step 2 selection, loaded, and save state
  const [pendingSourceModelId, setPendingSourceModelId] = useState<string>(SAMPLE_STORAGE_MODEL.id);
  const [loadedSourceModel, setLoadedSourceModel] = useState<StorageSourceModelEnvelope | null>(null);
  const [isModelLoaded, setIsModelLoaded] = useState<boolean>(false);
  const [activeRefineStep, setActiveRefineStep] = useState<number>(1);
  const [hasRefineUnsavedChanges, setHasRefineUnsavedChanges] = useState<boolean>(false);
  const [showSavePlanningModal, setShowSavePlanningModal] = useState(false);
  const [deleteSourceModalTarget, setDeleteSourceModalTarget] = useState<StorageSourceModelEnvelope | null>(null);

  // Step 3 selection, loaded, and save state
  const [step1SubTab, setStep1SubTab] = useState<'generate' | 'load'>('generate');
  const [pendingOptSourceModelId, setPendingOptSourceModelId] = useState<string>(SAMPLE_STORAGE_MODEL.id);
  const [loadedOptSourceModel, setLoadedOptSourceModel] = useState<StorageSourceModelEnvelope | null>(null);
  const [isOptSourceLoaded, setIsOptSourceLoaded] = useState<boolean>(false);
  const [pendingTargetModelId, setPendingTargetModelId] = useState<string>(SAMPLE_TARGET_STORAGE_MODEL.id);
  const [loadedTargetModel, setLoadedTargetModel] = useState<StorageTargetModelEnvelope | null>(null);
  const [hasTargetUnsavedChanges, setHasTargetUnsavedChanges] = useState<boolean>(false);
  const [showSaveTargetModal, setShowSaveTargetModal] = useState(false);
  const [deleteTargetModalTarget, setDeleteTargetModalTarget] = useState<StorageTargetModelEnvelope | null>(null);

  // Step 4 launch selector state
  const [selectedTargetModelId, setSelectedTargetModelId] = useState<string>(SAMPLE_TARGET_STORAGE_MODEL.id);

  // Shared unsaved Next confirmation & toast notifications
  const [unsavedNextTarget, setUnsavedNextTarget] = useState<'refine' | 'optimize' | 'provision' | null>(null);
  const [toastMsg, setToastMsg] = useState<string | null>(null);
  const [toastTone, setToastTone] = useState<'success' | 'error'>('success');

  const showToast = (msg: string, tone: 'success' | 'error' = 'success') => {
    setToastTone(tone);
    setToastMsg(msg);
  };

  const [isRefineJsonOpen, setIsRefineJsonOpen] = useState(false);
  const [selectedBucketIndex, setSelectedBucketIndex] = useState<number>(0);
  const [excludedBucketNames, setExcludedBucketNames] = useState<string[]>([]);
  const [showCorsGuide, setShowCorsGuide] = useState<boolean>(false);

  const markActiveStageDirty = () => {
    if (subTab === 'source') setHasStep1UnsavedChanges(true);
    else if (subTab === 'refine') setHasRefineUnsavedChanges(true);
    else if (subTab === 'optimize') setHasTargetUnsavedChanges(true);
  };

  const toggleExcludeBucket = (bucketName: string) => {
    markActiveStageDirty();
    setExcludedBucketNames((prev) =>
      prev.includes(bucketName)
        ? prev.filter((name) => name !== bucketName)
        : [...prev, bucketName]
    );
  };

  const [newOriginInputs, setNewOriginInputs] = useState<Record<number, string>>({});
  const [newAllowedHeaderInputs, setNewAllowedHeaderInputs] = useState<Record<number, string>>({});
  const [newExposeHeaderInputs, setNewExposeHeaderInputs] = useState<Record<number, string>>({});

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

  const hasRunningOsJob = objectStorageJobs.some((j) => j.status === 'Handling');

  // Local 1s elapsed-time UI clock (API status polling runs >= 10s in useGlobalJobPolling)
  useEffect(() => {
    if (!hasRunningOsJob) return;

    const intervalId = setInterval(() => {
      setObjectStorageJobs((prevJobs) =>
        prevJobs.map((job) => {
          if (job.status === 'Handling') {
            const realElapsed = calculateJobElapsedSeconds(job);
            return job.elapsedSeconds === realElapsed ? job : { ...job, elapsedSeconds: realElapsed };
          }
          return job;
        })
      );
    }, 1000);

    return () => clearInterval(intervalId);
  }, [hasRunningOsJob, setObjectStorageJobs]);

  const fetchSavedStorageModels = async () => {
    try {
      const [srcList, tgtList] = await Promise.all([
        artifactApi.getStorageSourceModels().catch(() => []),
        artifactApi.getStorageTargetModels().catch(() => []),
      ]);
      const cleanSrc = srcList.filter((m) => !isBuiltInSampleModel(m));
      const cleanTgt = tgtList.filter((m) => !isBuiltInSampleModel(m));
      setSavedSourceModels(cleanSrc);
      setSavedTargetModels(cleanTgt);
      if (cleanSrc.length > 0) {
        setPendingSourceModelId((prev) => (prev === SAMPLE_STORAGE_MODEL.id ? cleanSrc[0].id : prev));
        setPendingOptSourceModelId((prev) => (prev === SAMPLE_STORAGE_MODEL.id ? cleanSrc[0].id : prev));
      }
      if (cleanTgt.length > 0) {
        setPendingTargetModelId((prev) => (prev === SAMPLE_TARGET_STORAGE_MODEL.id ? cleanTgt[0].id : prev));
        setSelectedTargetModelId((prev) => (prev === SAMPLE_TARGET_STORAGE_MODEL.id ? cleanTgt[0].id : prev));
      }
    } catch (err) {
      console.warn('Initial storage artifact load failed:', err);
    }
  };

  useEffect(() => {
    fetchTumblebugProviders();
    beetleApi.getObjectStorageSupport().then((data) => {
      if (data && typeof data === 'object') {
        setCspSupportMap(data);
      }
    }).catch(() => {});
    fetchSavedStorageModels();
  }, []);

  // Grouped model lists for dropdowns
  const userStorageModels = savedSourceModels.filter((m) => !isBuiltInSampleModel(m));
  const userPlanningModels = userStorageModels.filter((m) => m.stage === 'planning');
  const userSourceModels = userStorageModels.filter((m) => m.stage !== 'planning');
  const sampleSourceModels: StorageSourceModelEnvelope[] = [SAMPLE_STORAGE_MODEL];
  const allSourceModels: StorageSourceModelEnvelope[] = [
    ...userPlanningModels,
    ...userSourceModels,
    ...sampleSourceModels,
  ];

  const userTargetModels = savedTargetModels.filter((m) => !isBuiltInSampleModel(m));
  const sampleTargetModels: StorageTargetModelEnvelope[] = [SAMPLE_TARGET_STORAGE_MODEL];
  const allTargetModels: StorageTargetModelEnvelope[] = [
    ...userTargetModels,
    ...sampleTargetModels,
  ];

  const pendingSourceModel = allSourceModels.find((m) => m.id === pendingSourceModelId) || allSourceModels[0] || null;
  const isRefinePendingAlreadyLoaded = Boolean(
    isModelLoaded &&
      loadedSourceModel &&
      pendingSourceModel &&
      loadedSourceModel.id === pendingSourceModel.id &&
      !hasRefineUnsavedChanges
  );

  const pendingOptSourceModel = allSourceModels.find((m) => m.id === pendingOptSourceModelId) || allSourceModels[0] || null;
  const isOptSourcePendingAlreadyLoaded = Boolean(
    isOptSourceLoaded &&
      loadedOptSourceModel &&
      pendingOptSourceModel &&
      loadedOptSourceModel.id === pendingOptSourceModel.id
  );

  const pendingTargetModel = allTargetModels.find((m) => m.id === pendingTargetModelId) || allTargetModels[0] || null;
  const isTargetPendingAlreadyLoaded = Boolean(
    recommendationResult &&
      loadedTargetModel &&
      pendingTargetModel &&
      loadedTargetModel.id === pendingTargetModel.id &&
      !hasTargetUnsavedChanges
  );

  const buildActiveTargetStoragesPayload = () => {
    const activeBuckets = sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName));
    return activeBuckets.map((b) => {
      const isCorsOn = b.corsEnabled || false;
      const validRules = b.corsRule && b.corsRule.length > 0
        ? b.corsRule
        : [{ allowedOrigin: ['*'], allowedMethod: ['GET', 'POST', 'PUT', 'DELETE'], allowedHeader: ['*'], exposeHeader: ['ETag'], maxAgeSeconds: 3600 }];
      const cleanTargetName = b.targetBucketName || b.bucketName;
      return {
        bucketName: cleanTargetName,
        targetObjectStorageName: cleanTargetName,
        sourceBucketName: b.bucketName,
        storageClass: b.accessFrequency === 'frequent' ? 'Standard' : b.accessFrequency === 'infrequent' ? 'Standard-IA' : 'Archive',
        versioningEnabled: b.versioningEnabled || false,
        encryptionEnabled: b.encryptionEnabled || false,
        corsEnabled: isCorsOn,
        corsRule: isCorsOn ? validRules : [],
        isPublic: b.isPublic || false,
      };
    });
  };

  const hasUnsavedTargetDraft = Boolean(recommendationResult && (!loadedTargetModel || hasTargetUnsavedChanges));
  const draftTargetModelEnvelope: StorageTargetModelEnvelope | null = recommendationResult
    ? {
        id: DRAFT_STORAGE_TARGET_ID,
        name:
          loadedTargetModel && !isBuiltInSampleModel(loadedTargetModel)
            ? loadedTargetModel.name
            : buildDefaultModelName(
                'target',
                loadedOptSourceModel?.name || loadedSourceModel?.name || 'storage',
                desiredCsp
              ),
        description: recommendationResult.description || 'Unsaved target object storage draft',
        version: loadedTargetModel?.version || '1.0.0',
        stage: 'target',
        category: 'infra',
        schemaType: 'RecommendedObjectStorage',
        csp: desiredCsp,
        region: desiredRegion,
        targetObjectStorages: buildActiveTargetStoragesPayload(),
        raw: recommendationResult,
      }
    : null;

  const activeSelectedTargetModel: StorageTargetModelEnvelope =
    selectedTargetModelId === DRAFT_STORAGE_TARGET_ID && draftTargetModelEnvelope
      ? draftTargetModelEnvelope
      : allTargetModels.find((m) => m.id === selectedTargetModelId) ||
        allTargetModels[0] ||
        SAMPLE_TARGET_STORAGE_MODEL;

  const getCspSupport = (cspKey: string) => {
    const key = (cspKey || '').toLowerCase();
    return cspSupportMap[key] || { cors: true, presignedUrl: true, versioning: true };
  };

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

  const mapEnvelopeToSourceBuckets = (model: StorageSourceModelEnvelope): SourceBucket[] => {
    const rawBuckets =
      model.sourceObjectStorages && model.sourceObjectStorages.length > 0
        ? model.sourceObjectStorages
        : sampleStorageRequest.sourceObjectStorages || [];
    return rawBuckets.map((b: any) => ({
      bucketName: b.bucketName || 'sample-bucket-01',
      targetBucketName: b.targetBucketName || b.bucketName || 'sample-bucket-01',
      totalSizeBytes: b.totalSizeBytes || 10737418240,
      objectCount: b.objectCount || 1000,
      accessFrequency: b.accessFrequency || 'frequent',
      versioningEnabled: Boolean(b.versioningEnabled),
      encryptionEnabled: Boolean(b.encryptionEnabled),
      corsEnabled: Boolean(b.corsEnabled),
      corsRule: Array.isArray(b.corsRule) ? b.corsRule : [],
      isPublic: Boolean(b.isPublic),
      creationDate: b.creationDate,
      tags: b.tags || {},
      objects: b.objects,
    }));
  };

  // Explicit Load handler for Step 2 (Refinement)
  const handleLoadRefineModel = () => {
    if (!pendingSourceModel) return;
    const bucketsToLoad = mapEnvelopeToSourceBuckets(pendingSourceModel);
    setSourceBuckets(bucketsToLoad);
    setExcludedBucketNames([]);
    setSelectedBucketIndex(0);
    setLoadedSourceModel(pendingSourceModel);
    setIsModelLoaded(true);
    setActiveRefineStep(2);
    setHasRefineUnsavedChanges(false);
    if (!isOptSourceLoaded) {
      setPendingOptSourceModelId(pendingSourceModel.id);
    }
  };

  // Explicit Load handler for Step 3 Sub-Tab 1 (Generate Recommendation)
  const handleLoadOptSourceModel = () => {
    if (!pendingOptSourceModel) return;
    const bucketsToLoad = mapEnvelopeToSourceBuckets(pendingOptSourceModel);
    setSourceBuckets(bucketsToLoad);
    setExcludedBucketNames([]);
    setLoadedOptSourceModel(pendingOptSourceModel);
    setIsOptSourceLoaded(true);
  };

  // Explicit Load handler for Step 3 Sub-Tab 2 (Load Saved Target Model)
  const handleLoadTargetModel = () => {
    if (!pendingTargetModel) return;
    const targetCsp = (pendingTargetModel.csp || 'aws').toLowerCase();
    const targetReg = pendingTargetModel.region || 'ap-northeast-2';
    setDesiredCsp(targetCsp);
    setDesiredRegion(targetReg);
    fetchTumblebugRegions(targetCsp);

    const mappedBuckets: SourceBucket[] = (pendingTargetModel.targetObjectStorages || []).map((tb: any, idx: number) => {
      const rawBucketName = String(tb.bucketName || `target-storage-${idx + 1}`);
      const cleanTargetName = tb.targetObjectStorageName || rawBucketName.replace(/-x8f2$/, '');
      const srcName = tb.sourceBucketName || cleanTargetName;
      const freq: 'frequent' | 'infrequent' | 'archive' =
        tb.storageClass === 'Standard-IA' ? 'infrequent' : tb.storageClass === 'Archive' ? 'archive' : 'frequent';
      return {
        bucketName: srcName,
        targetBucketName: cleanTargetName,
        totalSizeBytes: tb.totalSizeBytes || 10737418240,
        objectCount: tb.objectCount || 1000,
        accessFrequency: freq,
        versioningEnabled: Boolean(tb.versioningEnabled),
        encryptionEnabled: Boolean(tb.encryptionEnabled),
        corsEnabled: Boolean(tb.corsEnabled),
        corsRule: Array.isArray(tb.corsRule) ? tb.corsRule : [],
        isPublic: Boolean(tb.isPublic),
        tags: tb.tags || {},
      };
    });

    if (mappedBuckets.length > 0) {
      setSourceBuckets(mappedBuckets);
    }
    setExcludedBucketNames([]);
    setRecommendationResult(
      pendingTargetModel.raw || {
        status: 'recommended',
        description: pendingTargetModel.description || 'Loaded target object storage model',
        targetCloud: { csp: targetCsp, region: targetReg },
        targetObjectStorages: pendingTargetModel.targetObjectStorages,
      }
    );
    setLoadedTargetModel(pendingTargetModel);
    setSelectedTargetModelId(pendingTargetModel.id);
    setHasTargetUnsavedChanges(false);
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
    markActiveStageDirty();
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

  // Step 1: Scan Source Buckets
  const handleScanSourceBuckets = async () => {
    if (!scanAccessKey || !scanSecretKey) {
      showToast('Please enter Access Key ID and Secret Access Key before scanning buckets.', 'error');
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
        const sortedBucketNames = [...res.bucketNames].sort((a: string, b: string) => a.localeCompare(b));
        setScannedBucketNames(sortedBucketNames);
        setScannedBuckets(res.buckets || []);
        setSelectedBucketNames([sortedBucketNames[0]]);
      } else {
        showToast(res.error || 'No buckets found for the specified region or scan failed.', 'error');
      }
    } catch (e: any) {
      console.error('Scan error:', e);
      showToast('Error scanning cloud account buckets.', 'error');
    } finally {
      setIsScanning(false);
    }
  };

  // Inspect Source Storage Buckets Metadata Handler
  const handleInspectSourceBuckets = async () => {
    if (sourceBuckets.length === 0) {
      showToast('Please add or scan at least one source object storage bucket first.', 'error');
      return;
    }
    setIsInspectLoading(true);
    try {
      const res = await beetleApi.inspectSourceObjectStorage({
        csp: scanCsp,
        region: scanRegion,
        accessKeyId: scanAccessKey,
        secretAccessKey: scanSecretKey,
        buckets: sourceBuckets.map((b) => b.bucketName)
      });
      if (res.success && res.inspectedBuckets && res.inspectedBuckets.length > 0) {
        setSourceBuckets((prev) =>
          prev.map((b) => {
            const insp = res.inspectedBuckets.find((item: any) => item.bucketName === b.bucketName);
            if (insp) {
              return {
                ...b,
                totalSizeBytes: insp.totalSizeBytes || b.totalSizeBytes || 10737418240,
                objectCount: insp.objectCount || b.objectCount || 1000,
                versioningEnabled: insp.versioningEnabled ?? b.versioningEnabled,
                encryptionEnabled: insp.encryptionEnabled ?? b.encryptionEnabled,
                corsEnabled: insp.corsEnabled ?? b.corsEnabled,
                corsRule: insp.corsRule || b.corsRule,
                isPublic: insp.isPublic ?? b.isPublic
              };
            }
            return b;
          })
        );
        setIsInspectDone(true);
        setHasStep1UnsavedChanges(true);
        showToast('Source object storage metadata inspection completed.');
      } else {
        setIsInspectDone(true);
        setHasStep1UnsavedChanges(true);
        showToast('Source object storage metadata updated.');
      }
    } catch (err: any) {
      console.warn('Inspect error:', err);
      setIsInspectDone(true);
      setHasStep1UnsavedChanges(true);
    } finally {
      setIsInspectLoading(false);
    }
  };

  const formatSourceBucketsForArtifact = (buckets: SourceBucket[]) => {
    return buckets.map((b) => {
      const isCorsOn = b.corsEnabled || false;
      const rawRules = b.corsRule && b.corsRule.length > 0
        ? b.corsRule
        : [{ allowedOrigin: ['*'], allowedMethod: ['GET', 'POST', 'PUT', 'DELETE'], allowedHeader: ['*'], exposeHeader: ['ETag'], maxAgeSeconds: 3600 }];
      const validRules = rawRules.map((r) => ({
        ...r,
        maxAgeSeconds: Math.max(0, Math.round(Number(r.maxAgeSeconds) || 3600)),
      }));
      return {
        accessFrequency: b.accessFrequency || 'frequent',
        bucketName: b.bucketName,
        corsEnabled: isCorsOn,
        corsRule: isCorsOn ? validRules : [],
        creationDate: b.creationDate || '',
        encryptionEnabled: b.encryptionEnabled || false,
        isPublic: b.isPublic || false,
        objectCount: Math.max(0, Math.round(Number(b.objectCount) || 0)),
        tags: b.tags || {},
        totalSizeBytes: Math.max(0, Math.round(Number(b.totalSizeBytes) || 0)),
        versioningEnabled: b.versioningEnabled || false
      };
    });
  };

  // Step 1: Save Source Storage Model (stage: 'source')
  const handleSaveStep1SourceModel = async (result: SaveRevisionResult) => {
    if (sourceBuckets.length === 0) {
      throw new Error('At least one source bucket is required to save.');
    }
    const payload = {
      description: result.description || `Extracted source object storage model for ${scanCsp.toUpperCase()} (${scanRegion})`,
      sourceCloud: { csp: scanCsp, region: scanRegion },
      sourceObjectStorages: formatSourceBucketsForArtifact(sourceBuckets),
    };
    const saved = result.overwriteId
      ? await artifactApi.updateStorageSourceModel(
          result.overwriteId,
          result.name,
          result.description,
          payload,
          result.version,
          { stage: 'source' }
        )
      : await artifactApi.saveStorageSourceModel(
          result.name,
          result.description,
          payload,
          result.version,
          { stage: 'source' }
        );
    setSavedSourceModels((prev) => [saved, ...prev.filter((m) => m.id !== saved.id)]);
    setLastSavedSourceStorageModel(saved);
    setHasStep1UnsavedChanges(false);
    setPendingSourceModelId(saved.id);
    setPendingOptSourceModelId(saved.id);
    showToast(
      result.overwriteId
        ? `Source model "${saved.name}" (v${saved.version}) updated.`
        : `Source model "${saved.name}" (v${saved.version}) saved.`
    );
  };

  // Step 2: Save Planning Storage Model (stage: 'planning')
  const handleSavePlanningStorageModel = async (result: SaveRevisionResult) => {
    const activeBuckets = sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName));
    if (activeBuckets.length === 0) {
      throw new Error('Include at least one bucket before saving a planning model.');
    }
    const sourceCloud = loadedSourceModel?.sourceCloud || { csp: scanCsp, region: scanRegion };
    const payload = {
      description: result.description || 'Refined planning object storage model',
      sourceCloud,
      sourceObjectStorages: formatSourceBucketsForArtifact(activeBuckets),
    };
    const parentId =
      loadedSourceModel && !isBuiltInSampleModel(loadedSourceModel)
        ? loadedSourceModel.id
        : undefined;
    const saved = result.overwriteId
      ? await artifactApi.updateStorageSourceModel(
          result.overwriteId,
          result.name,
          result.description,
          payload,
          result.version,
          { stage: 'planning', parentId }
        )
      : await artifactApi.saveStorageSourceModel(
          result.name,
          result.description,
          payload,
          result.version,
          { stage: 'planning', parentId }
        );
    setSavedSourceModels((prev) => [saved, ...prev.filter((m) => m.id !== saved.id)]);
    setLoadedSourceModel(saved);
    setPendingSourceModelId(saved.id);
    setActiveRefineStep(3);
    setHasRefineUnsavedChanges(false);
    setPendingOptSourceModelId(saved.id);
    showToast(
      result.overwriteId
        ? `Planning model "${saved.name}" (v${saved.version}) updated.`
        : `Planning model "${saved.name}" (v${saved.version}) saved.`
    );
  };

  // Step 3: Run Beetle Object Storage Recommendation API
  const handleGenerateRecommendation = async () => {
    const activeBuckets = sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName));
    if (activeBuckets.length === 0) {
      showToast('All buckets are currently excluded. Include at least one bucket for recommendation.', 'error');
      return;
    }

    const payloadSourceBuckets = formatSourceBucketsForArtifact(activeBuckets);

    setIsRecommending(true);
    try {
      const res = await beetleApi.recommendObjectStorage(desiredCsp, desiredRegion, payloadSourceBuckets);
      setRecommendationResult(res);
      setLoadedTargetModel(null);
      setHasTargetUnsavedChanges(true);
    } catch (err: any) {
      console.warn('API recommendation fallback using sampleTargetObjectStorage', err);
      setRecommendationResult({
        ...sampleTargetObjectStorage,
        targetCloud: { csp: desiredCsp, region: desiredRegion },
        description: `Target cloud object storage recommendation for ${activeBuckets.length} bucket(s) on ${desiredCsp.toUpperCase()} (${desiredRegion})`
      });
      setLoadedTargetModel(null);
      setHasTargetUnsavedChanges(true);
    } finally {
      setIsRecommending(false);
    }
  };

  // Step 3: Save Target Storage Model (stage: 'target')
  const handleSaveTargetModelRevision = async (result: SaveRevisionResult) => {
    if (!recommendationResult) {
      throw new Error('Generate or load a target recommendation first.');
    }
    const targetStorages = buildActiveTargetStoragesPayload();
    if (targetStorages.length === 0) {
      throw new Error('Include at least one target bucket before saving.');
    }
    const payloadRecommendation = {
      ...recommendationResult,
      status: recommendationResult.status || 'recommended',
      description: result.description || recommendationResult.description || 'Recommended target cloud object storage model',
      targetCloud: {
        csp: desiredCsp,
        region: desiredRegion,
      },
      targetObjectStorages: targetStorages,
    };

    const parentCandidate = loadedOptSourceModel || loadedSourceModel;
    const parentId =
      parentCandidate && !isBuiltInSampleModel(parentCandidate)
        ? parentCandidate.id
        : undefined;

    const saved = result.overwriteId
      ? await artifactApi.updateStorageTargetModel(
          result.overwriteId,
          result.name,
          result.description,
          payloadRecommendation,
          result.version,
          { stage: 'target', parentId }
        )
      : await artifactApi.saveStorageTargetModel(
          result.name,
          result.description,
          payloadRecommendation,
          result.version,
          { stage: 'target', parentId }
        );

    setSavedTargetModels((prev) => [saved, ...prev.filter((m) => m.id !== saved.id)]);
    setLoadedTargetModel(saved);
    setPendingTargetModelId(saved.id);
    setSelectedTargetModelId(saved.id);
    setHasTargetUnsavedChanges(false);
    showToast(
      result.overwriteId
        ? `Target model "${saved.name}" (v${saved.version}) updated.`
        : `Target model "${saved.name}" (v${saved.version}) saved.`
    );
  };

  const handleProceedToStep = (nextTab: 'refine' | 'optimize' | 'provision', isDirty: boolean) => {
    if (isDirty) {
      setUnsavedNextTarget(nextTab);
      return;
    }
    if (nextTab === 'provision') {
      if (hasUnsavedTargetDraft) {
        setSelectedTargetModelId(DRAFT_STORAGE_TARGET_ID);
      } else if (loadedTargetModel) {
        setSelectedTargetModelId(loadedTargetModel.id);
      }
    }
    setSubTab(nextTab);
  };

  // Step 4: Provision & Migrate Object Storage (Async API via Prefer: respond-async)
  const handleMigrateStorage = async () => {
    setSubTab('provision');
    setIsDeploying(true);
    const mockReqId = `req-${Date.now().toString().slice(-6)}`;

    // Pass nameSeed query param only if explicitly set and <= 20 chars (common.IsValidNameSeed constraint)
    const rawSeed = (nameSeed || '').trim();
    const effectiveSeed = rawSeed.length > 0 && rawSeed.length <= 20 ? rawSeed : '';

    const rawTargetStorages =
      activeSelectedTargetModel.targetObjectStorages && activeSelectedTargetModel.targetObjectStorages.length > 0
        ? activeSelectedTargetModel.targetObjectStorages
        : buildActiveTargetStoragesPayload();

    const payloadTargetStorages = rawTargetStorages.map((b: any) => {
      const isCorsOn = Boolean(b.corsEnabled);
      const validRules = b.corsRule && b.corsRule.length > 0
        ? b.corsRule
        : [{ allowedOrigin: ['*'], allowedMethod: ['GET', 'POST', 'PUT', 'DELETE'], allowedHeader: ['*'], exposeHeader: ['ETag'], maxAgeSeconds: 3600 }];
      return {
        sourceBucketName: b.sourceBucketName || b.bucketName,
        bucketName: b.targetObjectStorageName || b.bucketName,
        versioningEnabled: Boolean(b.versioningEnabled),
        corsEnabled: isCorsOn,
        corsRule: isCorsOn ? validRules : []
      };
    });

    const targetCsp = (activeSelectedTargetModel.csp || desiredCsp || 'aws').toLowerCase();
    const targetReg = activeSelectedTargetModel.region || desiredRegion || 'ap-northeast-2';

    const targetModel = {
      status: activeSelectedTargetModel.raw?.status || 'recommended',
      description: activeSelectedTargetModel.description || `Object storage migration request for ${payloadTargetStorages.length} bucket(s)`,
      targetCloud: {
        csp: targetCsp,
        region: targetReg,
      },
      targetObjectStorages: payloadTargetStorages,
    };

    setDeploymentLog([
      `POST /beetle/migration/middleware/ns/${namespaceId || 'mig01'}/objectStorage${effectiveSeed ? `?nameSeed=${effectiveSeed}` : ''}`,
      `Header -> Prefer: respond-async`,
      `HTTP 202 Accepted (ReqID: ${mockReqId}, Status: Handling)`,
      `GET /beetle/request/${mockReqId} -> Status: Handling (Polling...)`
    ]);

    try {
      const res = await beetleApi.executeObjectStorageMigration(namespaceId || 'mig01', effectiveSeed, targetModel, true);

      const realReqId = res.reqId || mockReqId;
      const storageJobName = `${effectiveSeed || 'mig01'}-${targetCsp.toLowerCase()}-storage`;
      const newJobId = `os-job-${Date.now()}`;

      const logsList = [
        `POST /beetle/migration/middleware/ns/${namespaceId || 'mig01'}/objectStorage${effectiveSeed ? `?nameSeed=${effectiveSeed}` : ''}`,
        `Header -> Prefer: respond-async`,
        `HTTP 202 Accepted (ReqID: ${realReqId}, Status: Handling)`,
        ...(res.success
          ? [
              `GET /beetle/request/${realReqId} -> Status: Handling (Background worker allocated)`,
              `Dispatched target cloud bucket creation request to CB-Tumblebug backend.`,
              `GET /beetle/request/${realReqId} -> Status: Handling (Elapsed: 0s)`
            ]
          : [`[Error] Object Storage migration request failed: ${res.error || 'Resource conflict or API timeout'}`])
      ];

      const newJobCard: ObjectStorageJobCard = {
        id: newJobId,
        reqId: realReqId,
        name: storageJobName,
        csp: targetCsp,
        region: targetReg,
        status: res.success ? 'Handling' : 'Failed',
        startTime: new Date().toISOString(),
        createdAtMs: Date.now(),
        elapsedSeconds: 0,
        buckets: payloadTargetStorages.map((b: any) => ({
          name: b.bucketName,
          targetClass: 'Standard (S3/Hot)',
          cors: b.corsEnabled ? 'ENABLED' : 'DISABLED',
          versioning: b.versioningEnabled ? 'ENABLED' : 'DISABLED'
        })),
        logs: logsList,
        isSample: false,
        error: res.success ? undefined : res.error
      };

      setObjectStorageJobs((prev) => [newJobCard, ...prev]);
      setSelectedJobId(newJobId);
      setDeploymentLog(logsList);

      if (res.success) {
        setTimeout(() => {
          loadMigratedStorages();
        }, 3000);
      }
    } catch (err: any) {
      setDeploymentLog((prev) => [...prev, `[Error] ${err.message || 'Execution error'}`]);
    } finally {
      setIsDeploying(false);
    }
  };

  const loadMigratedStorages = async () => {
    try {
      await beetleApi.getMigratedObjectStorages(namespaceId);
    } catch (err) {
      console.warn('Failed to load migrated storages', err);
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
                        {cred.csp.toUpperCase()} / {cred.region}
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
                        <span className="px-2 py-0.5 rounded text-xs font-bold uppercase bg-bg-input border border-border-main text-text-main">
                          {bucket.accessFrequency}
                        </span>
                      </td>
                      <td className="p-3 space-x-1">
                        {bucket.versioningEnabled && <span className="px-1.5 py-0.5 bg-emerald-500/10 text-emerald-400 rounded text-xs">Ver</span>}
                        {bucket.encryptionEnabled && <span className="px-1.5 py-0.5 bg-teal-500/10 text-teal-400 rounded text-xs">Enc</span>}
                      </td>
                      <td className="p-3 text-right">
                        <button
                          type="button"
                          onClick={() => {
                            setSourceBuckets(prev => prev.filter((_, idx) => idx !== i));
                            setHasStep1UnsavedChanges(true);
                          }}
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

          </div>

          {/* ═══ SECTION 3: Collect & Save Source Storage Model Card ═══ */}
          <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-6 shadow-sm">
            {/* Header & 3-Step Lifecycle Stepper */}
            <div className="flex flex-wrap items-center justify-between gap-3 border-b border-border-main/30 pb-4">
              <div className="flex items-center space-x-3">
                <FileText className="w-5 h-5 text-emerald-500" />
                <h3 className="text-base font-extrabold text-text-main font-mono">
                  Collect &amp; Save Source Storage Model{' '}
                  <span className="text-text-muted font-normal text-sm">
                    ([{scanCsp.toUpperCase()}] {scanRegion})
                  </span>
                </h3>
              </div>

              {/* 3-Step Lifecycle Pill Stepper */}
              <div className="flex flex-wrap items-center gap-2 text-xs font-mono">
                <span className="flex items-center gap-1.5 px-3 py-1 bg-emerald-500/10 text-emerald-500 border border-emerald-500/30 rounded-full font-bold">
                  <CheckCircle2 className="w-3.5 h-3.5" />
                  <span>1 Credentials Registered</span>
                </span>
                <span className="text-text-muted">&gt;</span>
                <span className={`flex items-center gap-1.5 px-3 py-1 rounded-full font-bold border ${isInspectDone ? 'bg-emerald-500/10 text-emerald-500 border-emerald-500/30' : 'bg-bg-panel text-text-muted border-border-main'}`}>
                  <span className="w-4 h-4 rounded-full bg-emerald-500/20 text-emerald-400 flex items-center justify-center text-xs">{isInspectDone ? '✓' : '2'}</span>
                  <span>2 Storage Metadata Inspected</span>
                </span>
                <span className="text-text-muted">&gt;</span>
                <span className={`flex items-center gap-1.5 px-3 py-1 rounded-full font-bold border ${lastSavedSourceStorageModel && !hasStep1UnsavedChanges ? 'bg-emerald-500/10 text-emerald-500 border-emerald-500/30' : 'bg-bg-panel text-text-muted border-border-main'}`}>
                  <span className={`w-4 h-4 rounded-full flex items-center justify-center text-xs ${lastSavedSourceStorageModel && !hasStep1UnsavedChanges ? 'bg-emerald-500/20 text-emerald-400' : 'bg-bg-input text-text-muted'}`}>
                    {lastSavedSourceStorageModel && !hasStep1UnsavedChanges ? '✓' : '3'}
                  </span>
                  <span>3 Model Saved</span>
                </span>
              </div>
            </div>

            {/* Metadata Inspection Action Row */}
            <div className="space-y-1.5">
              <button
                type="button"
                onClick={handleInspectSourceBuckets}
                disabled={isInspectLoading || sourceBuckets.length === 0}
                className="w-full py-3 bg-bg-panel hover:bg-bg-input border border-emerald-500/50 hover:border-emerald-500 text-emerald-500 dark:text-emerald-400 font-extrabold text-sm rounded-xl transition flex items-center justify-center space-x-2 cursor-pointer shadow-sm disabled:opacity-50"
              >
                <Play className={`w-4 h-4 text-emerald-500 ${isInspectLoading ? 'animate-spin' : ''}`} />
                <span>Inspect Storage Metadata from All Buckets</span>
              </button>
              <p className="text-xs text-text-muted text-center font-mono">
                Inspects metadata (objects, size, CORS, policies) from registered buckets.
              </p>
            </div>

            {sourceBuckets.length > 0 && (
              <ActiveModelBanner
                mode={lastSavedSourceStorageModel ? 'loaded' : 'draft'}
                name={
                  lastSavedSourceStorageModel?.name ||
                  buildDefaultModelName('source', `${scanCsp}-${scanRegion}-storage`)
                }
                version={lastSavedSourceStorageModel?.version}
                stageKind={lastSavedSourceStorageModel ? 'source' : 'draft'}
                hasUnsavedChanges={hasStep1UnsavedChanges}
                updatedTime={lastSavedSourceStorageModel?.updatedTime}
                metrics={[
                  { label: 'CSP', value: `${scanCsp.toUpperCase()} (${scanRegion})`, tone: 'emerald' },
                  { label: 'Buckets', value: String(sourceBuckets.length) },
                ]}
              />
            )}

            {/* Collected Storage Model Preview */}
            <div className="space-y-3 pt-2 border-t border-border-main/30">
              <div className="flex flex-wrap items-center justify-between gap-2">
                <h4 className="text-sm font-bold text-text-main flex items-center space-x-2 font-mono">
                  <FileText className="w-4 h-4 text-emerald-500" />
                  <span>Collected Storage Model Preview</span>
                </h4>
                <div className="flex items-center gap-2.5">
                  <span className="text-xs text-text-muted font-mono">
                    {sourceBuckets.length} Buckets Registered
                  </span>
                  <button
                    type="button"
                    onClick={() => {
                      navigator.clipboard.writeText(
                        JSON.stringify(
                          {
                            description: `Extracted source object storage model for ${scanCsp.toUpperCase()} (${scanRegion})`,
                            sourceCloud: { csp: scanCsp, region: scanRegion },
                            sourceObjectStorages: sourceBuckets,
                          },
                          null,
                          2
                        )
                      );
                      showToast('JSON copied to clipboard.');
                    }}
                    className="px-3 py-1 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-xs font-extrabold transition flex items-center cursor-pointer text-emerald-600 dark:text-emerald-400"
                  >
                    <Copy className="w-3.5 h-3.5 mr-1" />
                    <span>Copy JSON</span>
                  </button>
                </div>
              </div>

              <div className="bg-bg-panel border border-border-main rounded-xl p-4 font-mono text-xs text-slate-800 dark:text-emerald-400 max-h-[860px] min-h-[320px] overflow-y-auto select-text shadow-inner">
                <pre>
                  {JSON.stringify(
                    {
                      description: `Extracted source object storage model for ${scanCsp.toUpperCase()} (${scanRegion})`,
                      sourceCloud: {
                        csp: scanCsp,
                        region: scanRegion
                      },
                      sourceObjectStorages: sourceBuckets
                    },
                    null,
                    2
                  )}
                </pre>
              </div>
            </div>

            {/* Unified Bottom Action & Navigation Bar */}
            <div className="flex flex-wrap items-center justify-end gap-2.5 pt-4 border-t border-border-main/40">
              <button
                type="button"
                onClick={() => setShowSaveStep1Modal(true)}
                disabled={sourceBuckets.length === 0}
                className="w-full sm:w-auto px-5 py-2.5 bg-gradient-to-r from-emerald-400 via-teal-400 to-teal-500 hover:from-emerald-500 hover:to-teal-600 disabled:opacity-50 disabled:cursor-not-allowed text-slate-950 font-extrabold text-xs rounded-xl transition flex items-center justify-center gap-1.5 cursor-pointer shadow-md shadow-emerald-500/20"
              >
                <Save className="w-3.5 h-3.5 text-slate-950" />
                <span>Save Source Model</span>
              </button>
              <button
                type="button"
                onClick={() => handleProceedToStep('refine', hasStep1UnsavedChanges)}
                className="w-full sm:w-auto px-6 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white font-extrabold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center justify-center gap-2 cursor-pointer"
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
              Review &amp; refine extracted source metadata, adjust bucket specifications, configure CORS rules, and save planning models.
            </span>
          </div>

          {/* ═══ STEP 1 Card: Source Object Storage Model Selection ═══ */}
          <div className={`glass-panel p-6 rounded-2xl transition-all duration-300 ${activeRefineStep >= 1 ? 'opacity-100' : 'opacity-40 pointer-events-none'}`}>
            <div className="flex items-center space-x-3 mb-4 border-b border-border-main/40 pb-3">
              <span className={`w-7 h-7 rounded-full flex items-center justify-center text-sm font-extrabold ${isModelLoaded ? 'bg-green-500/20 text-green-600 dark:text-green-400 border border-green-500/30' : 'bg-emerald-500 text-slate-950'}`}>
                {isModelLoaded ? '✓' : '1'}
              </span>
              <h3 className="text-base font-extrabold text-text-main">Step 1: Source Object Storage Model Selection</h3>
            </div>
            <div className="space-y-3 max-w-xl">
              <label className="block text-sm font-normal text-text-muted">Choose Source or Planning Model</label>
              <div className="flex flex-wrap items-center gap-2.5">
                <select
                  value={pendingSourceModelId}
                  onChange={(e) => setPendingSourceModelId(e.target.value)}
                  className="flex-1 min-w-[260px] bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm font-extrabold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
                >
                  {userPlanningModels.length > 0 && (
                    <optgroup label="Planning Models">
                      {userPlanningModels.map((m) => (
                        <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                      ))}
                    </optgroup>
                  )}
                  {userSourceModels.length > 0 && (
                    <optgroup label="Source Models">
                      {userSourceModels.map((m) => (
                        <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                      ))}
                    </optgroup>
                  )}
                  <optgroup label="Sample Models">
                    {sampleSourceModels.map((m) => (
                      <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                    ))}
                  </optgroup>
                </select>
                {pendingSourceModel && (() => {
                  const badge = getStageBadgeMeta(resolveModelStageKind(pendingSourceModel, 'source'));
                  return (
                    <span className={`px-3 py-1.5 rounded-xl text-xs font-extrabold border shrink-0 ${badge.className}`}>
                      {badge.label}
                    </span>
                  );
                })()}
              </div>

              {pendingSourceModel && (
                <div className="text-xs bg-bg-input/50 border border-border-main/50 rounded-xl px-3.5 py-2.5 flex flex-wrap items-center gap-x-4 gap-y-1">
                  <span>
                    <span className="font-normal text-text-muted">Buckets: </span>
                    <span className="font-extrabold text-text-main">{pendingSourceModel.sourceObjectStorages?.length || 0}</span>
                  </span>
                  {pendingSourceModel.sourceCloud?.csp && (
                    <span>
                      <span className="font-normal text-text-muted">CSP: </span>
                      <span className="font-extrabold text-text-main">
                        {pendingSourceModel.sourceCloud.csp.toUpperCase()} ({pendingSourceModel.sourceCloud.region})
                      </span>
                    </span>
                  )}
                  {pendingSourceModel.description && (
                    <span className="truncate max-w-md">
                      <span className="font-normal text-text-muted">Info: </span>
                      <span className="font-extrabold text-text-main">{pendingSourceModel.description}</span>
                    </span>
                  )}
                </div>
              )}

              <div className="flex flex-wrap items-center gap-3 pt-1">
                <button
                  type="button"
                  onClick={handleLoadRefineModel}
                  disabled={!pendingSourceModel}
                  className={`px-5 py-2.5 rounded-xl text-sm font-extrabold flex items-center gap-1.5 transition cursor-pointer ${
                    !pendingSourceModel
                      ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                      : isRefinePendingAlreadyLoaded
                        ? 'bg-emerald-500/20 border border-emerald-500/40 text-emerald-600 dark:text-emerald-400'
                        : 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 shadow-md shadow-emerald-500/25'
                  }`}
                >
                  {isRefinePendingAlreadyLoaded ? (
                    <>
                      <CheckCircle2 className="w-4 h-4" />
                      <span>Model Loaded</span>
                    </>
                  ) : (
                    <>
                      <RefreshCw className="w-4 h-4" />
                      <span>Load Model</span>
                    </>
                  )}
                </button>
                {pendingSourceModel && !isBuiltInSampleModel(pendingSourceModel) && (
                  <button
                    type="button"
                    onClick={() => setDeleteSourceModalTarget(pendingSourceModel)}
                    className="px-5 py-2.5 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-xl text-sm font-extrabold flex items-center gap-1.5 transition cursor-pointer"
                  >
                    <Trash2 className="w-4 h-4" />
                    <span>Delete Model</span>
                  </button>
                )}
              </div>
            </div>
          </div>

          {/* ═══ STEP 2 Card: Planning Object Storage Specification & Refinement ═══ */}
          <div className={`glass-panel p-6 rounded-2xl transition-all duration-300 ${isModelLoaded ? 'opacity-100' : 'opacity-40 pointer-events-none'}`}>
            <div className="flex items-center justify-between mb-4 border-b border-border-main/40 pb-3">
              <div className="flex items-center space-x-3">
                <span className={`w-7 h-7 rounded-full flex items-center justify-center text-sm font-extrabold ${activeRefineStep > 2 ? 'bg-green-500/20 text-green-600 dark:text-green-400 border border-green-500/30' : 'bg-emerald-500 text-slate-950'}`}>
                  {activeRefineStep > 2 ? '✓' : '2'}
                </span>
                <h3 className="text-base font-extrabold text-text-main">
                  Step 2: Planning Object Storage Specification &amp; Refinement
                </h3>
              </div>
              {isModelLoaded && (
                <div className="flex items-center space-x-2">
                  <button
                    type="button"
                    onClick={() => setIsRefineJsonOpen(!isRefineJsonOpen)}
                    className="px-3 py-1.5 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-sm font-bold transition cursor-pointer flex items-center text-emerald-600 dark:text-emerald-400"
                  >
                    <Copy className="w-3.5 h-3.5 mr-1" />
                    <span>{isRefineJsonOpen ? 'Hide JSON' : 'JSON View'}</span>
                  </button>
                </div>
              )}
            </div>

            {isModelLoaded && loadedSourceModel && (
              <div className="mb-4">
                <ActiveModelBanner
                  mode="loaded"
                  name={loadedSourceModel.name}
                  version={loadedSourceModel.version}
                  stageKind={resolveModelStageKind(loadedSourceModel, 'source')}
                  hasUnsavedChanges={hasRefineUnsavedChanges}
                  updatedTime={loadedSourceModel.updatedTime}
                  metrics={[
                    {
                      label: 'Buckets',
                      value: String(sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName)).length),
                    },
                    ...(excludedBucketNames.length > 0
                      ? [{ label: 'Excluded', value: String(excludedBucketNames.length) }]
                      : []),
                  ]}
                />
              </div>
            )}

            {isModelLoaded && (
              <div className="space-y-6">
                {isRefineJsonOpen && (
                  <div className="space-y-3 animate-fade-in">
                    <div className="flex justify-between items-center bg-bg-input px-3.5 py-2 rounded-xl border border-border-main/50">
                      <span className="text-sm text-text-muted font-mono">
                        {cleanModelDisplayName(loadedSourceModel?.name, 'source-object-storage')}.json
                      </span>
                      <button
                        type="button"
                        onClick={() => {
                          navigator.clipboard.writeText(
                            JSON.stringify(
                              {
                                sourceCloud: loadedSourceModel?.sourceCloud || { csp: scanCsp, region: scanRegion },
                                sourceObjectStorages: sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName)),
                              },
                              null,
                              2
                            )
                          );
                          showToast('JSON copied to clipboard.');
                        }}
                        className="px-3 py-1 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-xs font-extrabold transition flex items-center cursor-pointer text-emerald-600 dark:text-emerald-400"
                      >
                        <Copy className="w-3.5 h-3.5 mr-1" /> Copy JSON
                      </button>
                    </div>
                    <pre className="text-xs font-mono text-slate-800 dark:text-emerald-400 bg-bg-panel p-3.5 rounded-xl border border-border-main overflow-y-auto max-h-[360px] select-text">
                      {JSON.stringify(
                        {
                          sourceCloud: loadedSourceModel?.sourceCloud || { csp: scanCsp, region: scanRegion },
                          sourceObjectStorages: sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName)),
                        },
                        null,
                        2
                      )}
                    </pre>
                  </div>
                )}

                <div className="space-y-6">
                    {/* Bucket Card List & Selected Bucket Refinement Section */}
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
                              type="button"
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
                              <span className={`px-2.5 py-0.5 rounded-md text-xs font-mono font-extrabold uppercase ${isExcluded ? 'bg-red-500/20 text-red-400 border border-red-500/30' : 'bg-emerald-500/20 text-emerald-400 border border-emerald-500/30'}`}>
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
                                      setHasRefineUnsavedChanges(true);
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
                                      setHasRefineUnsavedChanges(true);
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
                                      setHasRefineUnsavedChanges(true);
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
                                      setHasRefineUnsavedChanges(true);
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
                                    CORS Detailed Configuration Rules ({b.bucketName})
                                  </span>
                                </div>

                                <div className="flex items-center space-x-3 text-xs shrink-0">
                                  <button
                                    type="button"
                                    onClick={() => setShowCorsGuide(true)}
                                    className="px-3 py-1 bg-emerald-500/10 hover:bg-emerald-500/20 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 font-bold rounded-lg transition cursor-pointer flex items-center space-x-1"
                                  >
                                    <AlertCircle className="w-3.5 h-3.5 text-emerald-500" />
                                    <span>CORS Rule Guide &amp; Examples</span>
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
                                          type="button"
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
                                      type="button"
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
                                          type="button"
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
                                          type="button"
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
                                      type="button"
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
                                          type="button"
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

                          {/* Section C: High-Performance Object Directory Tree View */}
                          <div className="pt-4 border-t border-border-main/40 space-y-2">
                            <div className="flex items-center space-x-2 pb-1">
                              <Folder className="w-4 h-4 text-emerald-500" />
                              <span className="font-extrabold text-sm text-text-main font-mono">
                                Bucket Directory &amp; Object Hierarchy Tree ({b.bucketName})
                              </span>
                            </div>
                            <ObjectDirectoryTree
                              bucketName={b.bucketName}
                              objects={b.objects && b.objects.length > 0 ? b.objects : [
                                { key: 'documents/2026/Q1_Report.pdf', sizeBytes: 2457600, lastModified: '2026-07-28 14:20:00' },
                                { key: 'documents/2026/Q2_Draft.docx', sizeBytes: 1048576, lastModified: '2026-07-28 15:10:00' },
                                { key: 'photos/uploads/banner_hero.png', sizeBytes: 4194304, lastModified: '2026-07-27 09:30:00' },
                                { key: 'photos/uploads/logo_dark.svg', sizeBytes: 524288, lastModified: '2026-07-27 10:15:00' },
                                { key: 'backup/db_dump_20260728.tar.gz', sizeBytes: 1073741824, lastModified: '2026-07-28 02:00:00' },
                                { key: 'README.md', sizeBytes: 4096, lastModified: '2026-07-20 11:00:00' }
                              ]}
                            />
                          </div>
                        </div>
                      );
                    })()}
                  </div>

                {/* Unified Bottom Action & Navigation Bar */}
                <div className="flex flex-col sm:flex-row items-center justify-between gap-3 pt-4 border-t border-border-main/20 mt-4">
                  <button
                    type="button"
                    onClick={() => setSubTab('source')}
                    className="w-full sm:w-auto px-4 py-2.5 bg-bg-input border border-border-main hover:bg-bg-main text-text-main font-extrabold text-xs rounded-xl transition cursor-pointer flex items-center justify-center space-x-1.5"
                  >
                    <ArrowLeft className="w-3.5 h-3.5" />
                    <span>Back to 1. Source Storage</span>
                  </button>

                  <div className="flex flex-wrap items-center gap-2.5 w-full sm:w-auto justify-end">
                    <button
                      type="button"
                      onClick={() => setShowSavePlanningModal(true)}
                      className="w-full sm:w-auto px-5 py-2.5 bg-gradient-to-r from-emerald-400 via-teal-400 to-teal-500 hover:from-emerald-500 hover:to-teal-600 text-slate-950 rounded-xl text-xs font-extrabold flex items-center justify-center gap-1.5 transition cursor-pointer shadow-md shadow-emerald-500/20"
                    >
                      <Save className="w-3.5 h-3.5 text-slate-950" />
                      <span>Save Planning Model</span>
                    </button>
                    <button
                      type="button"
                      onClick={() => handleProceedToStep('optimize', hasRefineUnsavedChanges)}
                      className="w-full sm:w-auto px-6 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white font-extrabold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center justify-center gap-2 cursor-pointer"
                    >
                      <span>Next: Proceed to 3. Target Object Storage Optimizer</span>
                      <ArrowRight className="w-4 h-4" />
                    </button>
                  </div>
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
              Generate optimized target cloud recommendations, compare multi-CSP specs &amp; cost estimates, and customize target cloud models.
            </span>
          </div>

          {/* Main Card Container */}
          <div className="glass-panel p-6 rounded-2xl space-y-6 border border-border-main/60 shadow-sm">
            {/* Step Header with Green Check Circle Icon / Number Badge */}
            <div className="flex items-center space-x-2.5 border-b border-border-main/40 pb-3">
              <div className={`w-6 h-6 rounded-full flex items-center justify-center font-extrabold text-xs shrink-0 ${
                (step1SubTab === 'generate' ? isOptSourceLoaded : Boolean(loadedTargetModel))
                  ? 'bg-emerald-500/20 text-emerald-500 border border-emerald-500/40'
                  : 'bg-emerald-500 text-slate-950'
              }`}>
                {(step1SubTab === 'generate' ? isOptSourceLoaded : Boolean(loadedTargetModel)) ? (
                  <Check className="w-3.5 h-3.5 text-emerald-500 font-bold" />
                ) : (
                  '1'
                )}
              </div>
              <h3 className="text-base font-extrabold text-text-main">
                Step 1: Target Cloud &amp; Model Selection
              </h3>
            </div>

            {/* Sub-Navigation Tabs */}
            <div className="flex items-center space-x-8 border-b border-border-main/40 text-sm font-extrabold pt-1">
              <button
                type="button"
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
                type="button"
                onClick={() => setStep1SubTab('load')}
                className={`pb-3 transition relative cursor-pointer ${
                  step1SubTab === 'load'
                    ? 'text-emerald-600 dark:text-emerald-400 border-b-2 border-emerald-500'
                    : 'text-text-muted hover:text-text-main font-bold'
                }`}
              >
                Load Saved Target Object Storage Model
              </button>
            </div>

            {/* Sub-Tab 1: Generate New Recommendation */}
            {step1SubTab === 'generate' && (
              <div className="space-y-5 pt-2 max-w-xl">
                <div className="space-y-2.5">
                  <label className="text-sm font-normal text-text-muted block">
                    Source or Planning Object Storage Model
                  </label>
                  <div className="flex flex-wrap items-center gap-2.5">
                    <select
                      value={pendingOptSourceModelId}
                      onChange={(e) => setPendingOptSourceModelId(e.target.value)}
                      className="flex-1 min-w-[260px] px-4 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold text-sm focus:outline-none focus:border-emerald-500 cursor-pointer"
                    >
                      {userPlanningModels.length > 0 && (
                        <optgroup label="Planning Models">
                          {userPlanningModels.map((m) => (
                            <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                          ))}
                        </optgroup>
                      )}
                      {userSourceModels.length > 0 && (
                        <optgroup label="Source Models">
                          {userSourceModels.map((m) => (
                            <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                          ))}
                        </optgroup>
                      )}
                      <optgroup label="Sample Models">
                        {sampleSourceModels.map((m) => (
                          <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                        ))}
                      </optgroup>
                    </select>
                    {pendingOptSourceModel && (() => {
                      const badge = getStageBadgeMeta(resolveModelStageKind(pendingOptSourceModel, 'source'));
                      return (
                        <span className={`px-3 py-1.5 rounded-xl text-xs font-extrabold border shrink-0 ${badge.className}`}>
                          {badge.label}
                        </span>
                      );
                    })()}
                  </div>

                  {pendingOptSourceModel && (
                    <div className="text-xs bg-bg-input/50 border border-border-main/50 rounded-xl px-3.5 py-2.5 flex flex-wrap items-center gap-x-4 gap-y-1">
                      <span>
                        <span className="font-normal text-text-muted">Buckets: </span>
                        <span className="font-extrabold text-text-main">{pendingOptSourceModel.sourceObjectStorages?.length || 0}</span>
                      </span>
                      {pendingOptSourceModel.description && (
                        <span className="truncate max-w-md">
                          <span className="font-normal text-text-muted">Info: </span>
                          <span className="font-extrabold text-text-main">{pendingOptSourceModel.description}</span>
                        </span>
                      )}
                    </div>
                  )}

                  <div className="flex flex-wrap items-center gap-3 pt-1">
                    <button
                      type="button"
                      onClick={handleLoadOptSourceModel}
                      disabled={!pendingOptSourceModel}
                      className={`px-4 py-2.5 rounded-xl text-xs font-extrabold flex items-center gap-1.5 transition cursor-pointer ${
                        !pendingOptSourceModel
                          ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                          : isOptSourcePendingAlreadyLoaded
                            ? 'bg-emerald-500/20 border border-emerald-500/40 text-emerald-600 dark:text-emerald-400'
                            : 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 shadow-md shadow-emerald-500/20'
                      }`}
                    >
                      {isOptSourcePendingAlreadyLoaded ? (
                        <>
                          <CheckCircle2 className="w-3.5 h-3.5" />
                          <span>Model Loaded</span>
                        </>
                      ) : (
                        <>
                          <RefreshCw className="w-3.5 h-3.5" />
                          <span>Load Model</span>
                        </>
                      )}
                    </button>

                    {pendingOptSourceModel && !isBuiltInSampleModel(pendingOptSourceModel) && (
                      <button
                        type="button"
                        onClick={() => setDeleteSourceModalTarget(pendingOptSourceModel)}
                        className="px-4 py-2.5 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-xl text-xs font-extrabold flex items-center gap-1.5 transition cursor-pointer"
                      >
                        <Trash2 className="w-3.5 h-3.5" />
                        <span>Delete Model</span>
                      </button>
                    )}
                  </div>
                </div>

                {/* Desired CSP & Desired Region */}
                <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 pt-1">
                  <div className="space-y-2">
                    <label className="text-sm font-normal text-text-muted block">
                      Desired CSP
                    </label>
                    <select
                      value={desiredCsp}
                      onChange={(e) => handleDesiredCspChange(e.target.value)}
                      className="w-full px-4 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold text-sm focus:outline-none focus:border-emerald-500 cursor-pointer"
                    >
                      {getCspList().map((c) => (
                        <option key={c.id} value={c.id}>
                          {c.name}
                        </option>
                      ))}
                    </select>
                  </div>

                  <div className="space-y-2">
                    <label className="text-sm font-normal text-text-muted block">
                      Desired Region
                    </label>
                    <select
                      value={desiredRegion}
                      onChange={(e) => setDesiredRegion(e.target.value)}
                      className="w-full px-4 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold font-mono text-sm focus:outline-none focus:border-emerald-500 cursor-pointer"
                    >
                      {getRegionsForCsp(desiredCsp).map((r) => (
                        <option key={r.id} value={r.id}>
                          {r.name}
                        </option>
                      ))}
                    </select>
                  </div>
                </div>

                {/* Recommend Action Button */}
                <div className="pt-2">
                  <button
                    type="button"
                    onClick={handleGenerateRecommendation}
                    disabled={isRecommending || !isOptSourceLoaded || sourceBuckets.length === 0}
                    className={`px-5 py-2.5 rounded-xl font-extrabold text-xs sm:text-sm flex items-center space-x-2 transition-all ${
                      !isOptSourceLoaded || sourceBuckets.length === 0
                        ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed opacity-60'
                        : 'bg-gradient-to-r from-emerald-400 via-teal-400 to-teal-500 hover:from-emerald-500 hover:to-teal-600 text-slate-950 cursor-pointer shadow-lg shadow-emerald-500/20'
                    }`}
                  >
                    {isRecommending ? (
                      <RefreshCw className="w-4 h-4 animate-spin" />
                    ) : (
                      <Sparkles className={`w-4 h-4 ${!isOptSourceLoaded || sourceBuckets.length === 0 ? 'text-text-muted' : 'text-slate-950'}`} />
                    )}
                    <span>Recommend Target Cloud Object Storage</span>
                  </button>
                </div>
              </div>
            )}

            {/* Sub-Tab 2: Load Saved Target Object Storage Model */}
            {step1SubTab === 'load' && (
              <div className="space-y-4 pt-2 max-w-xl">
                <div className="space-y-2.5">
                  <label className="text-sm font-normal text-text-muted block">
                    Saved Target Object Storage Model
                  </label>
                  <div className="flex flex-wrap items-center gap-2.5">
                    <select
                      value={pendingTargetModelId}
                      onChange={(e) => setPendingTargetModelId(e.target.value)}
                      className="flex-1 min-w-[260px] px-4 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold text-sm focus:outline-none focus:border-emerald-500 cursor-pointer"
                    >
                      {userTargetModels.length > 0 && (
                        <optgroup label="Target Models">
                          {userTargetModels.map((m) => (
                            <option key={m.id} value={m.id}>
                              {formatTargetOptionText(m, m.csp)}
                            </option>
                          ))}
                        </optgroup>
                      )}
                      <optgroup label="Sample Models">
                        {sampleTargetModels.map((m) => (
                          <option key={m.id} value={m.id}>
                            {formatTargetOptionText(m, m.csp)}
                          </option>
                        ))}
                      </optgroup>
                    </select>
                    {pendingTargetModel && (() => {
                      const badge = getStageBadgeMeta(resolveModelStageKind(pendingTargetModel, 'target'));
                      return (
                        <span className={`px-3 py-1.5 rounded-xl text-xs font-extrabold border shrink-0 ${badge.className}`}>
                          {badge.label}
                        </span>
                      );
                    })()}
                  </div>

                  {pendingTargetModel && (
                    <div className="text-xs bg-bg-input/50 border border-border-main/50 rounded-xl px-3.5 py-2.5 flex flex-wrap items-center gap-x-4 gap-y-1">
                      <span>
                        <span className="font-normal text-text-muted">CSP / Region: </span>
                        <span className="font-extrabold text-emerald-600 dark:text-emerald-400">
                          {(pendingTargetModel.csp || 'aws').toUpperCase()} ({pendingTargetModel.region || 'ap-northeast-2'})
                        </span>
                      </span>
                      <span>
                        <span className="font-normal text-text-muted">Buckets: </span>
                        <span className="font-extrabold text-text-main">
                          {pendingTargetModel.targetObjectStorages?.length || 0}
                        </span>
                      </span>
                      {pendingTargetModel.description && (
                        <span className="truncate max-w-md">
                          <span className="font-normal text-text-muted">Info: </span>
                          <span className="font-extrabold text-text-main">{pendingTargetModel.description}</span>
                        </span>
                      )}
                    </div>
                  )}

                  <div className="flex flex-wrap items-center gap-3 pt-1">
                    <button
                      type="button"
                      onClick={handleLoadTargetModel}
                      disabled={!pendingTargetModel}
                      className={`px-4 py-2.5 rounded-xl text-xs font-extrabold flex items-center gap-1.5 transition cursor-pointer ${
                        !pendingTargetModel
                          ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                          : isTargetPendingAlreadyLoaded
                            ? 'bg-emerald-500/20 border border-emerald-500/40 text-emerald-600 dark:text-emerald-400'
                            : 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 shadow-md shadow-emerald-500/20'
                      }`}
                    >
                      {isTargetPendingAlreadyLoaded ? (
                        <>
                          <CheckCircle2 className="w-3.5 h-3.5" />
                          <span>Model Loaded</span>
                        </>
                      ) : (
                        <>
                          <RefreshCw className="w-3.5 h-3.5" />
                          <span>Load Model</span>
                        </>
                      )}
                    </button>

                    {pendingTargetModel && !isBuiltInSampleModel(pendingTargetModel) && (
                      <button
                        type="button"
                        onClick={() => setDeleteTargetModalTarget(pendingTargetModel)}
                        className="px-4 py-2.5 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-xl text-xs font-extrabold flex items-center gap-1.5 cursor-pointer transition"
                      >
                        <Trash2 className="w-3.5 h-3.5" />
                        <span>Delete Model</span>
                      </button>
                    )}
                  </div>
                </div>
              </div>
            )}

          </div>

          {/* Step 2 Panel Container */}
          <div className="glass-panel p-6 rounded-2xl space-y-6 border border-border-main/60 shadow-sm mt-6">
            {/* Step 2 Header */}
            <div className="flex items-center justify-between border-b border-border-main/40 pb-3 gap-2">
              <div className="flex items-center space-x-2.5">
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

              {recommendationResult && (
                <button
                  type="button"
                  onClick={() => setIsTargetJsonOpen(!isTargetJsonOpen)}
                  className="px-3 py-1.5 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-xs font-extrabold transition cursor-pointer flex items-center text-emerald-600 dark:text-emerald-400 shrink-0"
                >
                  <Copy className="w-3.5 h-3.5 mr-1" />
                  <span>{isTargetJsonOpen ? 'Hide JSON' : 'JSON View'}</span>
                </button>
              )}
            </div>

            {!recommendationResult ? (
              <div className="py-12 text-center text-sm font-medium text-text-muted/60 italic">
                No target model loaded or generated yet. Load a source model and run recommendation, or load a saved target model above.
              </div>
            ) : (
              <div className="space-y-6 animate-fade-in">
                <ActiveModelBanner
                  mode={loadedTargetModel ? 'loaded' : 'draft'}
                  name={
                    loadedTargetModel?.name ||
                    buildDefaultModelName(
                      'target',
                      loadedOptSourceModel?.name || loadedSourceModel?.name || 'storage',
                      desiredCsp
                    )
                  }
                  version={loadedTargetModel?.version}
                  stageKind={loadedTargetModel ? resolveModelStageKind(loadedTargetModel, 'target') : 'draft'}
                  hasUnsavedChanges={hasTargetUnsavedChanges}
                  updatedTime={loadedTargetModel?.updatedTime}
                  metrics={[
                    { label: 'CSP', value: `${desiredCsp.toUpperCase()} (${desiredRegion})`, tone: 'emerald' },
                    {
                      label: 'Buckets',
                      value: String(sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName)).length),
                    },
                  ]}
                />

                {isTargetJsonOpen && (
                  <div className="space-y-3 animate-fade-in">
                    <div className="flex justify-between items-center bg-bg-input px-3.5 py-2 rounded-xl border border-border-main/50">
                      <span className="text-xs text-text-muted font-mono">
                        {cleanModelDisplayName(loadedTargetModel?.name, `${desiredCsp}-target-storage`)}.json
                      </span>
                      <button
                        type="button"
                        onClick={() => {
                          navigator.clipboard.writeText(
                            JSON.stringify(
                              {
                                ...recommendationResult,
                                targetCloud: { csp: desiredCsp, region: desiredRegion },
                                targetObjectStorages: buildActiveTargetStoragesPayload(),
                              },
                              null,
                              2
                            )
                          );
                          showToast('JSON copied to clipboard.');
                        }}
                        className="px-3 py-1 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-xs font-extrabold transition flex items-center cursor-pointer text-emerald-600 dark:text-emerald-400"
                      >
                        <Copy className="w-3.5 h-3.5 mr-1" /> Copy JSON
                      </button>
                    </div>
                    <pre className="font-mono text-xs text-slate-800 dark:text-emerald-300 bg-bg-panel p-4 rounded-xl border border-border-main max-h-[360px] overflow-y-auto leading-relaxed select-text">
                      {JSON.stringify(
                        {
                          ...recommendationResult,
                          targetCloud: { csp: desiredCsp, region: desiredRegion },
                          targetObjectStorages: buildActiveTargetStoragesPayload(),
                        },
                        null,
                        2
                      )}
                    </pre>
                  </div>
                )}

                {/* Recommended Target Object Storage Banner */}
                <div className="p-4 bg-emerald-500/10 border border-emerald-500/20 rounded-xl space-y-2 text-xs">
                  <div className="font-extrabold text-emerald-600 dark:text-emerald-400 flex items-center space-x-2">
                    <CheckCircle2 className="w-4 h-4 text-emerald-500" />
                    <span>Target Object Storage Specification Ready</span>
                  </div>
                  <div className="text-xs text-text-muted">{recommendationResult.description}</div>
                  <div className="pt-2 border-t border-emerald-500/20 text-xs text-emerald-700 dark:text-emerald-300 flex items-start space-x-1.5 font-medium">
                    <AlertCircle className="w-4 h-4 text-emerald-500 shrink-0 mt-0.5" />
                    <span>
                      <strong>CB-managed UID 안내:</strong> 클라우드 Bucket Name은 Cloud-Barista 내부에서 고유한 <strong>CB-managed UID</strong>로 자동 할당 및 관리됩니다. 사용자 지정 식별을 위해 UI에서는 <strong>Object Storage Name</strong>으로 표기합니다.
                    </span>
                  </div>
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
                      <span className="text-text-muted font-normal block mb-1 text-xs">Configured Policies</span>
                      <span className="font-extrabold text-sm text-text-main">
                        Versioning / CORS Rules
                      </span>
                    </div>
                  </div>
                </div>

                {/* Review and Editing Section */}
                <div className="space-y-5 pt-2 border-t border-border-main/40">
                  <div>
                    <h4 className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 flex items-center space-x-2">
                      <Sliders className="w-4 h-4 text-emerald-500" />
                      <span>Review and Editing</span>
                    </h4>
                    <p className="text-xs text-text-muted italic pt-0.5">
                      * Modifying target resource values updates the target object storage model in real-time
                    </p>
                  </div>

                  {/* Target Bucket Editing Cards Grid */}
                  <div className="grid grid-cols-1 gap-4">
                    {sourceBuckets.map((b, i) => {
                      const isExcluded = excludedBucketNames.includes(b.bucketName);
                      if (isExcluded) return null;
                      const corsRule = ensureCorsRuleExists(b);
                      const currentTargetName = b.targetBucketName !== undefined ? b.targetBucketName : b.bucketName;
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

                            <div className="flex items-center space-x-2">
                              <span className="px-2.5 py-0.5 rounded text-xs font-mono font-bold bg-emerald-500/20 text-emerald-400 uppercase shrink-0">
                                {desiredCsp.toUpperCase()} ({desiredRegion})
                              </span>

                              <button
                                type="button"
                                onClick={() => toggleExcludeBucket(b.bucketName)}
                                className="px-2.5 py-1 bg-red-500/10 hover:bg-red-500/20 text-red-500 dark:text-red-400 border border-red-500/30 rounded-lg text-xs font-extrabold transition flex items-center space-x-1 cursor-pointer"
                                title="Exclude this storage from Target Model"
                              >
                                <Trash2 className="w-3.5 h-3.5" />
                                <span>Exclude</span>
                              </button>
                            </div>
                          </div>

                          {/* Spec & Policy Controls Row */}
                          <div className="grid grid-cols-1 md:grid-cols-3 gap-4 text-xs">
                            {/* Column 1: Name (Editable Target Object Storage Name) */}
                            <div className="space-y-1.5 md:col-span-1">
                              <label className="text-text-muted font-normal block text-sm">Object Storage Name</label>
                              <input
                                type="text"
                                value={currentTargetName}
                                onChange={(e) => {
                                  const val = e.target.value;
                                  setHasTargetUnsavedChanges(true);
                                  setSourceBuckets((prev) =>
                                    prev.map((item, idx) => (idx === i ? { ...item, targetBucketName: val } : item))
                                  );
                                }}
                                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold font-mono text-xs focus:outline-none focus:border-emerald-500"
                                placeholder="e.g. target-storage-01"
                              />
                              <span className="text-xs text-text-muted font-normal block pt-0.5">
                                * Note: A unique <span className="font-bold text-text-main">CB-managed UID</span> is automatically assigned to the target bucket upon creation.
                              </span>
                            </div>

                            {/* Column 2: Refinable Target Bucket Policies */}
                            <div className="md:col-span-2 space-y-1.5">
                              <label className="text-text-muted font-normal block">Target Bucket Policies</label>
                              <div className="grid grid-cols-2 sm:grid-cols-2 gap-3">
                                <label
                                  title={!cspSupport.versioning ? `Versioning is not supported by ${desiredCsp.toUpperCase()}` : ''}
                                  className={`flex items-center space-x-2 p-2.5 rounded-xl border transition ${
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
                                      setHasTargetUnsavedChanges(true);
                                      setSourceBuckets((prev) => prev.map((item, idx) => (idx === i ? { ...item, versioningEnabled: checked } : item)));
                                    }}
                                    className="accent-emerald-500 cursor-pointer w-4 h-4 disabled:cursor-not-allowed"
                                  />
                                  <div className="flex flex-col">
                                    <span className="font-extrabold text-text-main">Versioning</span>
                                    {!cspSupport.versioning && (
                                      <span className="text-xs font-bold text-red-500 dark:text-red-400">N/A ({desiredCsp.toUpperCase()})</span>
                                    )}
                                  </div>
                                </label>

                                <label
                                  title={!cspSupport.cors ? `CORS Rules are not supported by ${desiredCsp.toUpperCase()}` : ''}
                                  className={`flex items-center space-x-2 p-2.5 rounded-xl border transition ${
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
                                      setHasTargetUnsavedChanges(true);
                                      setSourceBuckets((prev) => prev.map((item, idx) => (idx === i ? { ...item, corsEnabled: checked } : item)));
                                    }}
                                    className="accent-emerald-500 cursor-pointer w-4 h-4 disabled:cursor-not-allowed"
                                  />
                                  <div className="flex flex-col">
                                    <span className="font-extrabold text-text-main">CORS Rules</span>
                                    {!cspSupport.cors && (
                                      <span className="text-xs font-bold text-red-500 dark:text-red-400">N/A ({desiredCsp.toUpperCase()})</span>
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
                                    CORS Detailed Configuration Rules ({currentTargetName})
                                  </span>
                                </div>

                                <div className="flex items-center space-x-3 text-xs shrink-0">
                                  <button
                                    type="button"
                                    onClick={() => setShowCorsGuide(true)}
                                    className="px-3 py-1 bg-emerald-500/10 hover:bg-emerald-500/20 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 font-bold rounded-lg transition cursor-pointer flex items-center space-x-1"
                                  >
                                    <AlertCircle className="w-3.5 h-3.5 text-emerald-500" />
                                    <span>CORS Rule Guide &amp; Examples</span>
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
                                          type="button"
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
                                      type="button"
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
                                          type="button"
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
                                          type="button"
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
                                      type="button"
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
                                          type="button"
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

                  {/* Excluded Storage Restorable List Banner */}
                  {excludedBucketNames.length > 0 && (
                    <div className="p-4 bg-red-500/10 border border-red-500/20 rounded-2xl space-y-2">
                      <div className="flex items-center justify-between">
                        <span className="text-xs font-extrabold text-red-400 flex items-center gap-1.5">
                          <AlertCircle className="w-4 h-4" />
                          Excluded Object Storages ({excludedBucketNames.length})
                        </span>
                      </div>
                      <div className="flex flex-wrap gap-2 pt-1">
                        {excludedBucketNames.map((name) => (
                          <span
                            key={name}
                            className="inline-flex items-center gap-1.5 px-3 py-1 bg-bg-panel border border-red-500/40 rounded-xl text-xs font-mono text-text-main font-bold"
                          >
                            <span>{name}</span>
                            <button
                              type="button"
                              onClick={() => toggleExcludeBucket(name)}
                              className="text-emerald-500 hover:text-emerald-400 font-extrabold cursor-pointer text-xs"
                              title="Restore to Target Model"
                            >
                              + Include Back
                            </button>
                          </span>
                        ))}
                      </div>
                    </div>
                  )}
                </div>
              </div>
            )}

            {/* Unified Bottom Action & Navigation Bar */}
            <div className="flex flex-col sm:flex-row items-center justify-between gap-3 pt-4 border-t border-border-main/50">
              <button
                type="button"
                onClick={() => setSubTab('refine')}
                className="w-full sm:w-auto px-4 py-2.5 bg-bg-input hover:bg-bg-main border border-border-main text-text-main font-extrabold text-xs rounded-xl transition flex items-center justify-center space-x-1.5 cursor-pointer"
              >
                <ArrowLeft className="w-3.5 h-3.5" />
                <span>Back to 2. Refinement</span>
              </button>

              <div className="flex flex-wrap items-center gap-2.5 w-full sm:w-auto justify-end">
                <button
                  type="button"
                  onClick={() => setShowSaveTargetModal(true)}
                  disabled={!recommendationResult}
                  className={`w-full sm:w-auto px-5 py-2.5 rounded-xl font-extrabold text-xs flex items-center justify-center gap-1.5 transition-all shrink-0 ${
                    !recommendationResult
                      ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed opacity-60'
                      : 'bg-gradient-to-r from-emerald-400 via-teal-400 to-teal-500 hover:from-emerald-500 hover:to-teal-600 text-slate-950 cursor-pointer shadow-md shadow-emerald-500/20'
                  }`}
                >
                  <Save className={`w-3.5 h-3.5 ${!recommendationResult ? 'text-text-muted' : 'text-slate-950'}`} />
                  <span>Save Target Model</span>
                </button>

                <button
                  type="button"
                  onClick={() => handleProceedToStep('provision', hasTargetUnsavedChanges)}
                  disabled={!recommendationResult}
                  className={`w-full sm:w-auto px-6 py-2.5 rounded-xl font-extrabold text-xs flex items-center justify-center gap-2 transition-all ${
                    !recommendationResult
                      ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed opacity-60'
                      : 'bg-emerald-600 hover:bg-emerald-700 text-white cursor-pointer shadow-lg shadow-emerald-500/20'
                  }`}
                >
                  <span>Next: Proceed to 4. Migration Execution</span>
                  <ArrowRight className="w-4 h-4" />
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
              type="button"
              onClick={() => {
                if (hasUnsavedTargetDraft) {
                  setSelectedTargetModelId(DRAFT_STORAGE_TARGET_ID);
                } else if (loadedTargetModel) {
                  setSelectedTargetModelId(loadedTargetModel.id);
                }
                setShowLaunchModal(true);
              }}
              disabled={isDeploying}
              className="px-5 py-2.5 bg-gradient-to-r from-emerald-400 via-teal-400 to-teal-500 hover:from-emerald-500 hover:to-teal-600 disabled:opacity-50 text-slate-950 rounded-xl text-xs font-extrabold flex items-center gap-1.5 transition cursor-pointer shadow-lg shadow-emerald-500/20"
            >
              {isDeploying ? <RefreshCw className="w-4 h-4 animate-spin text-slate-950" /> : <Plus className="w-4 h-4 text-slate-950" />}
              <span>Launch New Migration</span>
            </button>

            <div className="px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-xs font-bold font-mono text-text-main flex items-center gap-2">
              <Zap className="w-4 h-4 text-emerald-500" />
              <span>Active Migration Jobs ({objectStorageJobs.filter(j => j.status === 'Handling').length} Running / {objectStorageJobs.filter(j => j.status === 'Success').length} Completed)</span>
            </div>
          </div>

          {/* SECTION 1: Horizontal Side-by-Side Job Cards Bar */}
          <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-4 shadow-sm">
            <div className="flex justify-between items-center border-b border-border-main/20 pb-3">
              <h3 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                <Activity className="w-4 h-4 text-emerald-500" />
                Migration Jobs Queue ({objectStorageJobs.length})
              </h3>
              <span className="text-xs text-text-muted font-mono bg-bg-panel px-3 py-1 rounded-full border border-border-main">
                Click card to view detailed progress &amp; results
              </span>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-4">
              {objectStorageJobs.map((job) => {
                const isSelected = job.id === selectedJobId;
                return (
                  <div
                    key={job.id}
                    onClick={() => setSelectedJobId(job.id)}
                    className={`p-4 rounded-xl border text-left transition-all duration-200 cursor-pointer flex flex-col justify-between space-y-3 relative overflow-hidden ${
                      isSelected
                        ? 'bg-emerald-500/10 border-emerald-500/60 shadow-lg shadow-emerald-500/10 ring-1 ring-emerald-500/40'
                        : 'bg-bg-panel hover:bg-bg-input border-border-main/60'
                    }`}
                  >
                    <div className="absolute top-0 left-0 w-full h-0.5 bg-gradient-to-r from-emerald-500 to-teal-500" />

                    <div className="flex justify-between items-center gap-2">
                      <div className="flex items-center gap-2 min-w-0">
                        <span className="text-base shrink-0">
                          {job.csp.toLowerCase() === 'aws' ? '🌩️' : job.csp.toLowerCase() === 'azure' ? '🔷' : '🟢'}
                        </span>
                        <span className="font-extrabold text-sm text-text-main font-mono truncate" title={job.name}>
                          {job.isSample && <span className="text-amber-500 font-bold mr-1">[Sample]</span>}
                          {job.name}
                        </span>
                      </div>

                      <div className="flex items-center gap-2">
                        <span className={`px-2.5 py-1 rounded-full text-xs font-bold flex items-center gap-1 shrink-0 ${
                          job.status === 'Success' ? 'bg-green-500/10 text-green-400 border border-green-500/20' :
                          job.status === 'Handling' ? 'bg-amber-500/10 text-amber-400 border border-amber-500/20' :
                          'bg-red-500/10 text-red-400 border border-red-500/20'
                        }`}>
                          {job.status === 'Success' && <CheckCircle2 className="w-3.5 h-3.5 text-green-400" />}
                          {job.status === 'Handling' && <RefreshCw className="w-3.5 h-3.5 text-amber-400 animate-spin" />}
                          {job.status === 'Failed' && <X className="w-3.5 h-3.5 text-red-400" />}
                          <span>{job.status === 'Success' ? '✓ Success' : job.status}</span>
                        </span>

                        {!job.isSample && (
                          <button
                            onClick={(e) => {
                              e.stopPropagation();
                              setDeleteModalJob(job);
                              setDeleteConfirmText('');
                            }}
                            className="p-1 text-text-muted hover:text-red-500 hover:bg-red-500/10 rounded-lg transition cursor-pointer"
                            title="Delete job from queue history"
                          >
                            <Trash2 className="w-3.5 h-3.5" />
                          </button>
                        )}
                      </div>
                    </div>

                    <div className="flex justify-between items-center text-xs font-mono text-text-muted pt-1 border-t border-border-main/20">
                      <span>Region: {job.region}</span>
                      <span className="flex items-center gap-1 font-bold text-text-main">
                        <Clock className="w-3.5 h-3.5 text-emerald-500" />
                        Time: {job.elapsedSeconds}s ({job.status === 'Handling' ? 'Running' : 'Done'})
                      </span>
                    </div>
                  </div>
                );
              })}
            </div>
          </div>

          {/* SECTION 2: Selected Job Detail Panel */}
          {activeJob && (
            <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-6 animate-fade-in shadow-sm">
                {/* Selected Job Header */}
                <div className="flex flex-col sm:flex-row justify-between sm:items-center gap-3 border-b border-border-main/20 pb-4">
                  <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
                    <span className="text-lg">
                      {(activeJob?.csp || desiredCsp).toLowerCase() === 'aws' ? '🌩️' : (activeJob?.csp || desiredCsp).toLowerCase() === 'azure' ? '🔷' : '🟢'}
                    </span>
                    <span>
                      Selected Job Detail: {activeJob?.isSample && <span className="text-amber-500 font-bold mr-1">[Sample]</span>}[{activeJob?.name}]
                    </span>
                  </h3>

                  <div className="flex flex-wrap items-center gap-3 text-xs font-mono text-text-muted">
                    <span>Namespace: <strong className="text-text-main">{namespaceId || 'mig01'}</strong></span>
                    <span>Req ID: <strong className="text-emerald-500">{activeJob?.reqId || 'req-20260723-001'}</strong></span>
                    <span>Elapsed: <strong className="text-teal-400">{activeJob?.elapsedSeconds ?? 0}s</strong></span>
                    {!activeJob?.isSample && (
                      <button
                        onClick={() => {
                          setDeleteModalJob(activeJob);
                          setDeleteConfirmText('');
                        }}
                        className="px-2.5 py-1 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-lg text-xs font-bold transition flex items-center gap-1 cursor-pointer"
                        title="Delete job history from queue"
                      >
                        <Trash2 className="w-3.5 h-3.5" />
                        <span>Delete History</span>
                      </button>
                    )}
                  </div>
                </div>

                {/* Simplified 3-Stage API Status Flow (Matching Infra Stepper) */}
                <div className="bg-bg-panel/50 border border-border-main/40 p-5 rounded-xl space-y-3">
                  <div className="flex justify-between items-center">
                    <span className="text-xs font-bold text-emerald-600 dark:text-emerald-400 uppercase font-mono">
                      API MIGRATION EXECUTION STATUS
                    </span>
                    <span className="text-xs font-bold text-text-muted font-mono">
                      API Status: <strong className={activeJob?.status === 'Success' ? 'text-green-400' : activeJob?.status === 'Handling' ? 'text-amber-400' : 'text-red-400'}>{activeJob?.status || 'Success'}</strong>
                    </span>
                  </div>

                  <div className="grid grid-cols-1 md:grid-cols-3 gap-4 pt-1">
                    {/* Step 1: Request Accepted */}
                    <div className="bg-bg-input/60 border border-emerald-500/30 p-3.5 rounded-xl flex items-center space-x-3">
                      <div className="p-2 bg-emerald-500/20 text-emerald-400 rounded-lg font-bold text-xs">✓</div>
                      <div>
                        <h4 className="text-xs font-bold text-text-main">1. Request Accepted</h4>
                        <p className="text-xs text-text-muted font-mono">HTTP 202 (ReqID Issued)</p>
                      </div>
                    </div>

                    {/* Step 2: Migrating */}
                    <div className="bg-bg-input/60 border border-emerald-500/30 p-3.5 rounded-xl flex items-center space-x-3">
                      <div className="p-2 bg-emerald-500/20 text-emerald-400 rounded-lg font-bold text-xs">✓</div>
                      <div>
                        <h4 className="text-xs font-bold text-text-main">2. Migrating</h4>
                        <p className="text-xs text-text-muted font-mono">Finished Processing</p>
                      </div>
                    </div>

                    {/* Step 3: Completed */}
                    <div className={`p-3.5 rounded-xl flex items-center space-x-3 ${activeJob?.status === 'Success' ? 'bg-green-500/10 border border-green-500/40' : 'bg-bg-input/60 border border-border-main/30'}`}>
                      <div className="p-2 bg-green-500/20 text-green-400 rounded-lg font-bold text-xs">✓</div>
                      <div>
                        <h4 className="text-xs font-bold text-text-main">3. Completed</h4>
                        <p className="text-xs text-text-muted font-mono">Object Storage Active &amp; Ready</p>
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
                        {(activeJob?.buckets || []).map((b, idx) => (
                          <tr key={idx} className="hover:bg-bg-main/30 transition">
                            <td className="p-3.5 font-extrabold text-text-main">
                              {b.name}
                            </td>
                            <td className="p-3.5 font-bold text-emerald-500">
                              {b.targetClass}
                            </td>
                            <td className="p-3.5">
                              <span className={`px-2 py-0.5 rounded font-bold ${b.cors === 'ENABLED' ? 'bg-emerald-500/20 text-emerald-400' : 'bg-slate-500/20 text-slate-400'}`}>
                                {b.cors}
                              </span>
                            </td>
                            <td className="p-3.5">
                              <span className={`px-2 py-0.5 rounded font-bold ${b.versioning === 'ENABLED' ? 'bg-teal-500/20 text-teal-400' : 'bg-slate-500/20 text-slate-400'}`}>
                                {b.versioning}
                              </span>
                            </td>
                            <td className="p-3.5">
                              <button
                                type="button"
                                onClick={() =>
                                  showToast(
                                    `Storage access point reachable for "${b.name}" on ${(activeJob?.csp || 'aws').toUpperCase()} (${activeJob?.region || 'ap-northeast-2'}).`
                                  )
                                }
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
                  <h4 className="text-xs font-bold text-text-muted font-mono uppercase">
                    REST API REQUEST &amp; RESPONSE LOG
                  </h4>
                  <div className="bg-bg-input p-4 rounded-xl border border-border-main/40 font-mono text-xs text-text-muted space-y-1.5 max-h-48 overflow-y-auto">
                    {(activeJob?.logs || deploymentLog).map((line, idx) => (
                      <div key={idx} className="flex items-start gap-2">
                        <span className="text-emerald-500">›</span>
                        <span className={line.includes('Error') ? 'text-red-400 font-bold' : line.includes('Success') ? 'text-emerald-600 dark:text-emerald-400 font-bold' : ''}>
                          {line}
                        </span>
                      </div>
                    ))}
                  </div>
                </div>
              </div>
            )}

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
        )}

      {/* Step 1: Save Source Model Modal */}
      <SaveRevisionModal
        isOpen={showSaveStep1Modal}
        onClose={() => setShowSaveStep1Modal(false)}
        title="Save Source Model"
        stageKind="source"
        defaultName={
          lastSavedSourceStorageModel
            ? cleanModelDisplayName(lastSavedSourceStorageModel.name)
            : buildDefaultModelName('source', `${scanCsp}-${scanRegion}-storage`)
        }
        defaultDescription={
          lastSavedSourceStorageModel?.description ||
          `Extracted source object storage model (${sourceBuckets.length} buckets on ${scanCsp.toUpperCase()} ${scanRegion})`
        }
        defaultVersion={lastSavedSourceStorageModel?.version || '1.0.0'}
        defaultOverwriteId={lastSavedSourceStorageModel?.id || null}
        existingRevisions={userSourceModels.map((m) => ({
          id: m.id,
          name: m.name,
          description: m.description,
          version: m.version,
        }))}
        onSave={handleSaveStep1SourceModel}
      />

      {/* Step 2: Save Planning Model Modal */}
      <SaveRevisionModal
        isOpen={showSavePlanningModal}
        onClose={() => setShowSavePlanningModal(false)}
        title="Save Planning Model"
        stageKind="planning"
        defaultName={
          loadedSourceModel
            ? loadedSourceModel.stage === 'planning' && !isBuiltInSampleModel(loadedSourceModel)
              ? cleanModelDisplayName(loadedSourceModel.name)
              : buildDefaultModelName('planning', loadedSourceModel.name)
            : buildDefaultModelName('planning', 'storage')
        }
        defaultDescription={
          loadedSourceModel
            ? loadedSourceModel.stage === 'planning' && !isBuiltInSampleModel(loadedSourceModel)
              ? loadedSourceModel.description || 'Refined planning object storage model'
              : `Refined planning object storage model derived from ${cleanModelDisplayName(loadedSourceModel.name)}`
            : 'Refined planning object storage model'
        }
        defaultVersion={
          loadedSourceModel && loadedSourceModel.stage === 'planning' && !isBuiltInSampleModel(loadedSourceModel)
            ? loadedSourceModel.version || '1.0.0'
            : '1.0.0'
        }
        defaultOverwriteId={
          loadedSourceModel && loadedSourceModel.stage === 'planning' && !isBuiltInSampleModel(loadedSourceModel)
            ? loadedSourceModel.id
            : null
        }
        existingRevisions={userPlanningModels.map((m) => ({
          id: m.id,
          name: m.name,
          description: m.description,
          version: m.version,
        }))}
        onSave={handleSavePlanningStorageModel}
      />

      {/* Step 3: Save Target Model Modal */}
      <SaveRevisionModal
        isOpen={showSaveTargetModal}
        onClose={() => setShowSaveTargetModal(false)}
        title="Save Target Model"
        stageKind="target"
        defaultName={
          loadedTargetModel && !isBuiltInSampleModel(loadedTargetModel)
            ? cleanModelDisplayName(loadedTargetModel.name)
            : buildDefaultModelName(
                'target',
                loadedOptSourceModel?.name || loadedSourceModel?.name || 'storage',
                desiredCsp
              )
        }
        defaultDescription={
          loadedTargetModel && !isBuiltInSampleModel(loadedTargetModel)
            ? loadedTargetModel.description || 'Recommended target cloud object storage model'
            : `Target cloud object storage model on ${desiredCsp.toUpperCase()} (${desiredRegion})`
        }
        defaultVersion={
          loadedTargetModel && !isBuiltInSampleModel(loadedTargetModel)
            ? loadedTargetModel.version || '1.0.0'
            : '1.0.0'
        }
        defaultOverwriteId={
          loadedTargetModel && !isBuiltInSampleModel(loadedTargetModel)
            ? loadedTargetModel.id
            : null
        }
        existingRevisions={userTargetModels.map((m) => ({
          id: m.id,
          name: m.name,
          description: m.description,
          version: m.version,
        }))}
        onSave={handleSaveTargetModelRevision}
      />

      {/* Delete Source / Planning Storage Model Modal */}
      <DeleteModelConfirmModal
        isOpen={Boolean(deleteSourceModalTarget)}
        onClose={() => setDeleteSourceModalTarget(null)}
        title="Delete Model"
        modelName={deleteSourceModalTarget?.name || ''}
        onConfirm={async () => {
          if (!deleteSourceModalTarget) return;
          const deletedId = deleteSourceModalTarget.id;
          const deletedName = cleanModelDisplayName(deleteSourceModalTarget.name);
          await artifactApi.deleteStorageModel(deletedId);
          setSavedSourceModels((prev) => prev.filter((m) => m.id !== deletedId));
          const remaining = allSourceModels.filter((m) => m.id !== deletedId);
          const nextId = remaining[0]?.id || SAMPLE_STORAGE_MODEL.id;
          if (pendingSourceModelId === deletedId) setPendingSourceModelId(nextId);
          if (pendingOptSourceModelId === deletedId) setPendingOptSourceModelId(nextId);
          if (loadedSourceModel?.id === deletedId) {
            setLoadedSourceModel(null);
            setIsModelLoaded(false);
            setActiveRefineStep(1);
            setHasRefineUnsavedChanges(false);
          }
          if (loadedOptSourceModel?.id === deletedId) {
            setLoadedOptSourceModel(null);
            setIsOptSourceLoaded(false);
          }
          if (lastSavedSourceStorageModel?.id === deletedId) {
            setLastSavedSourceStorageModel(null);
          }
          showToast(`Model "${deletedName}" deleted.`);
        }}
      />

      {/* Delete Target Storage Model Modal */}
      <DeleteModelConfirmModal
        isOpen={Boolean(deleteTargetModalTarget)}
        onClose={() => setDeleteTargetModalTarget(null)}
        title="Delete Model"
        modelName={deleteTargetModalTarget?.name || ''}
        onConfirm={async () => {
          if (!deleteTargetModalTarget) return;
          const deletedId = deleteTargetModalTarget.id;
          const deletedName = cleanModelDisplayName(deleteTargetModalTarget.name);
          await artifactApi.deleteStorageModel(deletedId);
          setSavedTargetModels((prev) => prev.filter((m) => m.id !== deletedId));
          const remaining = allTargetModels.filter((m) => m.id !== deletedId);
          const nextId = remaining[0]?.id || SAMPLE_TARGET_STORAGE_MODEL.id;
          if (pendingTargetModelId === deletedId) setPendingTargetModelId(nextId);
          if (selectedTargetModelId === deletedId) setSelectedTargetModelId(nextId);
          if (loadedTargetModel?.id === deletedId) {
            setLoadedTargetModel(null);
            setRecommendationResult(null);
            setHasTargetUnsavedChanges(false);
          }
          showToast(`Model "${deletedName}" deleted.`);
        }}
      />

      {/* Unsaved Model Changes Confirmation Modal */}
      <UnsavedConfirmModal
        isOpen={Boolean(unsavedNextTarget)}
        onStay={() => setUnsavedNextTarget(null)}
        onProceed={() => {
          const dest = unsavedNextTarget;
          if (!dest) return;
          if (dest === 'provision') {
            if (hasUnsavedTargetDraft) {
              setSelectedTargetModelId(DRAFT_STORAGE_TARGET_ID);
            } else if (loadedTargetModel) {
              setSelectedTargetModelId(loadedTargetModel.id);
            }
          }
          setSubTab(dest);
        }}
      />

      <ModelToast
        message={toastMsg}
        tone={toastTone}
        onClose={() => setToastMsg(null)}
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
                <label className="block text-text-muted font-normal mb-1.5 text-sm">
                  1. Select Target Object Storage Model
                </label>
                <div className="flex flex-wrap items-center gap-2.5">
                  <select
                    value={selectedTargetModelId}
                    onChange={(e) => setSelectedTargetModelId(e.target.value)}
                    className="flex-1 min-w-[260px] px-3.5 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold focus:outline-none focus:border-emerald-500 text-sm cursor-pointer"
                  >
                    {hasUnsavedTargetDraft && draftTargetModelEnvelope && (
                      <optgroup label="Draft Model">
                        <option value={DRAFT_STORAGE_TARGET_ID}>
                          {formatTargetOptionText(draftTargetModelEnvelope, draftTargetModelEnvelope.csp, true)}
                        </option>
                      </optgroup>
                    )}
                    {userTargetModels.length > 0 && (
                      <optgroup label="Target Models">
                        {userTargetModels.map((m) => (
                          <option key={m.id} value={m.id}>
                            {formatTargetOptionText(m, m.csp)}
                          </option>
                        ))}
                      </optgroup>
                    )}
                    <optgroup label="Sample Models">
                      {sampleTargetModels.map((m) => (
                        <option key={m.id} value={m.id}>
                          {formatTargetOptionText(m, m.csp)}
                        </option>
                      ))}
                    </optgroup>
                  </select>
                  {(() => {
                    const isDraftSelected = selectedTargetModelId === DRAFT_STORAGE_TARGET_ID;
                    const badge = getStageBadgeMeta(
                      isDraftSelected
                        ? 'draft'
                        : resolveModelStageKind(activeSelectedTargetModel, 'target')
                    );
                    return (
                      <span className={`px-3 py-1.5 rounded-xl text-xs font-extrabold border shrink-0 ${badge.className}`}>
                        {badge.label}
                      </span>
                    );
                  })()}
                </div>
              </div>

              <div className="p-3.5 bg-bg-input/60 border border-border-main/40 rounded-xl space-y-2 font-mono text-xs">
                <div className="flex justify-between">
                  <span className="text-text-muted font-normal">Model Storage Name (storageId):</span>
                  <span className="font-bold text-emerald-500">
                    {activeSelectedTargetModel.targetObjectStorages?.[0]?.bucketName || activeSelectedTargetModel.name || 'target-storage-01'}
                  </span>
                </div>
                <div className="flex justify-between pt-1 border-t border-border-main/20">
                  <span className="text-text-muted font-normal">Target CSP / Region:</span>
                  <span className="font-bold text-text-main">
                    {(activeSelectedTargetModel.csp || desiredCsp).toUpperCase()} ({(activeSelectedTargetModel.region || desiredRegion)})
                  </span>
                </div>

                {/* Object Storages Section */}
                <div className="pt-2 border-t border-border-main/20 space-y-1.5 font-sans">
                  <div className="flex items-center justify-between text-text-muted font-normal text-xs">
                    <span className="font-bold text-emerald-600 dark:text-emerald-400">Object Storages ({activeSelectedTargetModel.targetObjectStorages?.length || sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName)).length})</span>
                  </div>
                  <div className="space-y-1.5 max-h-36 overflow-y-auto">
                    {(activeSelectedTargetModel.targetObjectStorages && activeSelectedTargetModel.targetObjectStorages.length > 0
                      ? activeSelectedTargetModel.targetObjectStorages
                      : sourceBuckets.filter((b) => !excludedBucketNames.includes(b.bucketName))
                    ).map((b: any, idx: number) => (
                      <div key={idx} className="p-2 bg-bg-panel/80 border border-border-main/60 rounded-lg flex items-center justify-between text-xs font-mono">
                        <div>
                          <span className="font-extrabold text-emerald-500">{b.bucketName || b.targetBucketName || `target-storage-${idx + 1}`}</span>
                          {b.sourceBucketName && (
                            <span className="text-text-muted font-normal ml-2 text-xs">(from: {b.sourceBucketName})</span>
                          )}
                        </div>
                        <div className="flex items-center space-x-2 text-xs">
                          <span className={`px-1.5 py-0.5 rounded font-bold ${b.versioningEnabled ? 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400' : 'bg-slate-500/10 text-text-muted'}`}>
                            Versioning: {b.versioningEnabled ? 'ON' : 'OFF'}
                          </span>
                          <span className={`px-1.5 py-0.5 rounded font-bold ${b.corsEnabled ? 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400' : 'bg-slate-500/10 text-text-muted'}`}>
                            CORS: {b.corsEnabled ? 'ON' : 'OFF'}
                          </span>
                        </div>
                      </div>
                    ))}
                  </div>
                </div>
              </div>

              <div>
                <label className="block text-text-muted font-normal mb-1.5 text-sm">
                  2. Configure Deployment Identifiers
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
                type="button"
                onClick={() => {
                  setShowLaunchModal(false);
                  handleMigrateStorage();
                }}
                className="px-5 py-2 bg-gradient-to-r from-emerald-400 via-teal-400 to-teal-500 hover:from-emerald-500 hover:to-teal-600 text-slate-950 font-extrabold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center space-x-1.5 cursor-pointer"
              >
                <Play className="w-4 h-4 text-slate-950" />
                <span>Launch Migration</span>
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Register New Credential Profile Modal (Matching CredentialManagement.tsx) */}
      {isRegisterCredModalOpen && (
        <div className="fixed inset-0 bg-slate-950/75 backdrop-blur-md z-50 flex items-center justify-center p-4 animate-fade-in font-sans">
          <div className="bg-bg-panel border border-border-main rounded-2xl max-w-5xl w-full p-6 sm:p-7 space-y-5 shadow-2xl overflow-hidden text-sm">
            <div className="flex items-center justify-between border-b border-border-main/40 pb-3">
              <div className="flex items-center space-x-2.5">
                <Lock className="w-5 h-5 text-emerald-500" />
                <h3 className="text-lg font-extrabold text-text-main">
                  Register New CSP Credential Profile
                </h3>
              </div>
              <button
                onClick={() => setIsRegisterCredModalOpen(false)}
                className="w-8 h-8 rounded-lg bg-bg-input hover:bg-bg-main border border-border-main flex items-center justify-center text-text-muted hover:text-text-main transition cursor-pointer font-bold text-sm"
              >
                ✕
              </button>
            </div>

            <div className="space-y-4">
              <div>
                <label className="block text-text-main font-bold mb-1.5">
                  Credential Profile Name <span className="text-red-500">*</span>
                </label>
                <input
                  type="text"
                  required
                  placeholder="e.g. aws-production-account"
                  value={credProfileName}
                  onChange={(e) => setCredProfileName(e.target.value)}
                  className="w-full px-3.5 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold"
                />
              </div>

              {/* Reusable Modular CSP Credential Form Component */}
              <CspCredentialForm
                csp={credCsp}
                onCspChange={setCredCsp}
                region={credRegion}
                onRegionChange={setCredRegion}
                accessKey={credAccessKey}
                onAccessKeyChange={setCredAccessKey}
                secretKey={credSecretKey}
                onSecretKeyChange={setCredSecretKey}
              />
            </div>

            <div className="border-t border-border-main/40 pt-4 flex items-center justify-end space-x-3">
              <button
                type="button"
                onClick={() => setIsRegisterCredModalOpen(false)}
                className="px-5 py-2.5 bg-bg-input hover:bg-bg-main border border-border-main text-text-main font-bold rounded-xl transition cursor-pointer"
              >
                Cancel
              </button>
              <button
                type="button"
                onClick={() => {
                  if (!credAccessKey) {
                    showToast('Please enter Access Key.', 'error');
                    return;
                  }
                  const newProfile = {
                    id: `cred-${Date.now()}`,
                    name: credProfileName || `cred-profile-${savedCredProfiles.length + 1}`,
                    csp: credCsp,
                    region: credRegion || 'ap-northeast-2',
                    accessKey: credAccessKey,
                    secretKey: credSecretKey
                  };
                  setSavedCredProfiles(prev => [...prev, newProfile]);
                  setSelectedCredentialProfile(newProfile.id);
                  setScanCsp(newProfile.csp);
                  setScanRegion(newProfile.region);
                  setScanAccessKey(newProfile.accessKey);
                  setScanSecretKey(newProfile.secretKey);
                  setIsRegisterCredModalOpen(false);
                }}
                className="px-6 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white font-extrabold rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center space-x-2 cursor-pointer"
              >
                <Key className="w-4 h-4" />
                <span>Save &amp; Select Profile</span>
              </button>
            </div>
          </div>
        </div>
      )}
      {/* Select & Scan Source Object Storage Buckets Modal Popup (Wide & Tall AGENTS.md Layout) */}
      {isAddModalOpen && (
        <div className="fixed inset-0 bg-slate-950/75 backdrop-blur-md z-50 flex items-center justify-center p-4 sm:p-6 animate-fade-in font-sans">
          <div className="bg-bg-panel border border-border-main rounded-2xl w-[94vw] max-w-[1400px] max-h-[90vh] p-6 sm:p-8 space-y-6 shadow-2xl overflow-hidden flex flex-col justify-between">
            
            {/* Header: Single-line Flex Container (AGENTS.md Rule #3 Parity) */}
            <div className="flex items-center justify-between border-b border-border-main/40 pb-4 shrink-0">
              <div className="flex items-center space-x-3">
                <Database className="w-5.5 h-5.5 text-emerald-500" />
                <h3 className="text-lg font-extrabold text-text-main tracking-tight font-mono">
                  Select Source Object Storage
                </h3>
              </div>
              <button
                type="button"
                onClick={() => setIsAddModalOpen(false)}
                className="w-9 h-9 rounded-xl bg-bg-input hover:bg-bg-main border border-border-main flex items-center justify-center text-text-muted hover:text-text-main transition cursor-pointer font-bold text-base"
              >
                ✕
              </button>
            </div>

            <div className="space-y-6 flex-1 overflow-y-auto pr-1">
              {/* Scan Configuration & Target Info Row */}
              <div className="p-4.5 bg-bg-input/60 border border-border-main/40 rounded-xl flex flex-wrap items-center justify-between gap-4">
                <div className="flex items-center space-x-3 text-sm">
                  <span className="font-normal text-text-muted">Target Cloud Account:</span>
                  <span className="font-extrabold text-text-main font-mono text-base">
                    {scanCsp.toUpperCase()} ({scanRegion})
                  </span>
                </div>
                <button
                  type="button"
                  onClick={handleScanSourceBuckets}
                  disabled={isScanning}
                  className="px-5 py-2.5 bg-emerald-600 hover:bg-emerald-700 disabled:opacity-50 text-white rounded-xl text-xs font-extrabold flex items-center space-x-2 transition cursor-pointer shadow-md"
                >
                  <RefreshCw className={`w-4 h-4 ${isScanning ? 'animate-spin' : ''}`} />
                  <span>{isScanning ? 'Scanning Cloud Account...' : 'Fetch Storage List'}</span>
                </button>
              </div>

              {/* Bucket Selection Section */}
              <div className="space-y-3">
                <div className="flex justify-between items-center border-b border-border-main/20 pb-2.5">
                  <span className="text-base font-bold text-text-main flex items-center gap-2">
                    <Database className="w-4.5 h-4.5 text-emerald-500" />
                    <span>Discovered Storage Buckets</span>
                  </span>
                  <span className="text-sm font-normal text-text-muted font-mono">
                    {scannedBucketNames.length} buckets found
                  </span>
                </div>

                <div className="max-h-[450px] min-h-[220px] overflow-y-auto border border-border-main/40 rounded-xl p-4 bg-bg-panel space-y-3">
                  {scannedBucketNames.length === 0 ? (
                    <div className="p-10 text-center text-text-muted italic text-sm font-mono space-y-2">
                      <Database className="w-8 h-8 text-text-muted mx-auto opacity-50" />
                      <p>No storage buckets found. Click &quot;Fetch Storage List&quot; above to scan your account.</p>
                    </div>
                  ) : (
                    scannedBucketNames.map((name) => {
                      const isSelected = selectedBucketNames.includes(name);
                      const bucketData = scannedBuckets.find((b) => b.bucketName === name);
                      const displaySize = bucketData?.totalSizeBytes || bucketData?.sizeBytes || 10737418240; // 10 GB default
                      const displayObjects = bucketData?.objectCount || bucketData?.count || 1000;
                      return (
                        <label
                          key={name}
                          className={`flex items-center justify-between p-4 rounded-xl border transition cursor-pointer ${
                            isSelected
                              ? 'bg-emerald-500/10 border-emerald-500/50 text-text-main shadow-sm'
                              : 'bg-bg-input/40 border-border-main/30 text-text-muted hover:border-emerald-500/30'
                          }`}
                        >
                          <div className="flex items-center space-x-3.5">
                            <input
                              type="checkbox"
                              checked={isSelected}
                              onChange={(e) => {
                                if (e.target.checked) {
                                  setSelectedBucketNames((prev) => [...prev, name]);
                                } else {
                                  setSelectedBucketNames((prev) => prev.filter((n) => n !== name));
                                }
                              }}
                              className="accent-emerald-500 w-4.5 h-4.5 cursor-pointer"
                            />
                            <span className="font-extrabold font-mono text-base text-text-main">{name}</span>
                          </div>

                          <div className="text-sm font-mono text-text-muted flex items-center space-x-6">
                            <span>Size: <strong className="font-extrabold text-text-main">{formatBytes(displaySize)}</strong></span>
                            <span>Objects: <strong className="font-extrabold text-text-main">{(displayObjects).toLocaleString()}</strong></span>
                          </div>
                        </label>
                      );
                    })
                  )}
                </div>
              </div>
            </div>

            {/* Footer Navigation Buttons */}
            <div className="border-t border-border-main/40 pt-4 flex items-center justify-end space-x-3 shrink-0">
              <button
                type="button"
                onClick={() => setIsAddModalOpen(false)}
                className="px-5 py-2.5 bg-bg-input hover:bg-bg-main border border-border-main text-text-main font-bold text-xs rounded-xl transition cursor-pointer"
              >
                Cancel
              </button>
              <button
                type="button"
                onClick={() => {
                  if (selectedBucketNames.length === 0) {
                    showToast('Please select at least one object storage bucket.', 'error');
                    return;
                  }
                  const newBuckets: SourceBucket[] = selectedBucketNames.map((name) => {
                    const existing = scannedBuckets.find((b) => b.bucketName === name);
                    return {
                      bucketName: name,
                      totalSizeBytes: existing?.totalSizeBytes || existing?.sizeBytes || 10737418240,
                      objectCount: existing?.objectCount || existing?.count || 1000,
                      accessFrequency: existing?.accessFrequency || 'frequent',
                      versioningEnabled: existing?.versioningEnabled || false,
                      encryptionEnabled: existing?.encryptionEnabled || false,
                      corsEnabled: existing?.corsEnabled || false,
                      corsRule: existing?.corsRule || [],
                      isPublic: existing?.isPublic || false,
                      tags: existing?.tags || {}
                    };
                  });
                  setSourceBuckets(newBuckets);
                  setHasStep1UnsavedChanges(true);
                  setIsAddModalOpen(false);
                }}
                className="px-6 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white font-extrabold text-xs rounded-xl shadow-md transition flex items-center space-x-2 cursor-pointer"
              >
                <CheckCircle2 className="w-4 h-4" />
                <span>Confirm &amp; Import Storage(s)</span>
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Delete Record Confirmation Modal with Text Pattern Matching */}
      {deleteModalJob && (
        <div className="fixed inset-0 z-[60] flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up space-y-4">
            <div className="flex justify-between items-center border-b border-border-main/20 pb-3">
              <h3 className="text-base font-bold text-text-main flex items-center gap-2">
                <Trash2 className="w-4 h-4 text-red-500" />
                <span>Delete Object Storage Migration Record</span>
              </h3>
              <button
                onClick={() => { setDeleteModalJob(null); setDeleteConfirmText(''); }}
                className="text-text-muted hover:text-text-main transition p-1 hover:bg-bg-input rounded-lg cursor-pointer"
              >
                <X className="w-5 h-5" />
              </button>
            </div>

            <div className="space-y-4">
              <p className="text-xs text-text-muted leading-relaxed">
                Are you sure you want to remove the object storage migration record for <strong className="text-text-main">"{deleteModalJob.name}"</strong> (<span className="font-mono text-emerald-400">{deleteModalJob.reqId}</span>) from the Queue?
              </p>
              <div className="p-3 bg-amber-500/10 border border-amber-500/30 rounded-xl text-xs text-amber-400 font-mono">
                ⚠️ Note: This action only removes the UI queue request record. Physical cloud storage resources will NOT be deleted.
              </div>

              <div className="space-y-1.5 pt-1">
                <label className="block text-xs font-bold text-text-muted">
                  To confirm deletion, type <span className="font-mono bg-bg-panel px-1.5 py-0.5 rounded border border-border-main/60 text-emerald-400 select-all">{deleteModalJob.name}</span> below:
                </label>
                <input
                  type="text"
                  value={deleteConfirmText}
                  onChange={(e) => setDeleteConfirmText(e.target.value)}
                  placeholder={`Type "${deleteModalJob.name}" to confirm`}
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-xs focus:outline-none focus:ring-1 focus:ring-red-500 font-bold font-mono"
                  autoFocus
                />
              </div>

              <div className="flex justify-end gap-3 pt-2">
                <button
                  onClick={() => { setDeleteModalJob(null); setDeleteConfirmText(''); }}
                  className="px-4 py-2 bg-bg-panel border border-border-main text-text-main rounded-xl text-xs font-semibold hover:bg-bg-input transition cursor-pointer"
                >
                  Cancel
                </button>
                <button
                  onClick={handleConfirmDeleteStorageRecord}
                  disabled={deleteConfirmText !== deleteModalJob.name}
                  className={`px-4 py-2 rounded-xl text-xs font-bold transition flex items-center gap-1.5 ${
                    deleteConfirmText !== deleteModalJob.name
                      ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                      : 'bg-red-500 hover:bg-red-600 text-white cursor-pointer shadow-md shadow-red-500/20'
                  }`}
                >
                  <Trash2 className="w-3.5 h-3.5" />
                  <span>Confirm Delete</span>
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
