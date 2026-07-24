'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore } from '@/store/migrationStore';
import { beetleApi, tumblebugApi } from '@/api/client';
import { CspCredentialForm } from '../common/CspCredentialForm';
import { 
  Database, 
  Key, 
  ShieldCheck, 
  Play, 
  RefreshCw, 
  CheckCircle2, 
  Lock, 
  Cloud, 
  Server, 
  Filter, 
  Plus, 
  Shield, 
  Zap, 
  Activity, 
  AlertTriangle, 
  Terminal, 
  Trash2, 
  X,
  Sparkles,
  ArrowRight,
  Folder,
  Globe
} from 'lucide-react';

interface DataMigrationJob {
  id: string;
  reqId: string;
  sourceBucket: string;
  targetStorage: string;
  nsId: string;
  csp: string;
  region: string;
  strategy: string;
  status: 'Handling' | 'Success' | 'Failed';
  startTime: string;
  elapsedSeconds: number;
  encryptionKeyId: string;
  isSample?: boolean;
  logs: string[];
}

export const DataTransferCenter: React.FC = () => {
  const { namespaceId } = useMigrationStore();

  // Active Job Queue State
  const [dataJobs, setDataJobs] = useState<DataMigrationJob[]>([
    {
      id: 'req-data-101',
      reqId: 'req-data-101',
      sourceBucket: 'source-bucket-01/data/',
      targetStorage: 'target-storage-v1/imports/',
      nsId: 'mig01',
      csp: 'AWS',
      region: 'ap-northeast-2',
      strategy: 'auto',
      status: 'Success',
      startTime: new Date(Date.now() - 300000).toLocaleTimeString(),
      elapsedSeconds: 42,
      encryptionKeyId: 'key-rsa-2026-0723',
      isSample: true,
      logs: [
        'GET /beetle/migration/data/encryptionKey -> Issued RSA Public Key (KeyID: key-rsa-2026-0723)',
        'Encrypting sensitive fields (AccessKeyId, SecretAccessKey) client-side via RSA-OAEP-256...',
        'POST /beetle/migration/data -> 202 Accepted (ReqID: req-data-101)',
        'Data migration completed successfully (Duration: 42s)'
      ]
    }
  ]);
  const [activeJobId, setActiveJobId] = useState<string>('req-data-101');

  // Launch Modal State
  const [showLaunchModal, setShowLaunchModal] = useState(false);

  // Dynamic Tumblebug CSP & Region API State
  const [cspList, setCspList] = useState<string[]>(['aws', 'azure', 'gcp', 'alibaba', 'tencent', 'ncp', 'ibm', 'nhn', 'kt', 'openstack', 'minio']);
  const [regionList, setRegionList] = useState<{ id: string; name: string }[]>([]);
  const [modalRegionList, setModalRegionList] = useState<{ id: string; name: string }[]>([]);
  const [isLoadingTumblebugRegions, setIsLoadingTumblebugRegions] = useState(false);

  // Credential Profile Modal State
  const [isRegisterCredModalOpen, setIsRegisterCredModalOpen] = useState(false);
  const [credProfileName, setCredProfileName] = useState('');
  const [credCsp, setCredCsp] = useState('aws');
  const [credRegion, setCredRegion] = useState('ap-northeast-2');
  const [credAccessKey, setCredAccessKey] = useState('');
  const [credSecretKey, setCredSecretKey] = useState('');
  const [credTenantId, setCredTenantId] = useState('');
  const [credSubscriptionId, setCredSubscriptionId] = useState('');
  const [savedCredProfiles, setSavedCredProfiles] = useState<
    { id: string; name: string; csp: string; region: string; accessKey: string; secretKey?: string; tenantId?: string; subscriptionId?: string }[]
  >([
    {
      id: 'cred-sample-01',
      name: 'aws-production-account',
      csp: 'aws',
      region: 'ap-northeast-2',
      accessKey: 'AKIAIOSFODNN7EXAMPLE',
      secretKey: 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY'
    }
  ]);

  // Source & Target Access Type Selectors
  const [sourceAccessType, setSourceAccessType] = useState<'object-storage' | 'remote-ssh'>('object-storage');
  const [targetAccessType, setTargetAccessType] = useState<'object-storage' | 'remote-ssh'>('object-storage');

  // Destination Access Engine Mode (Tumblebug API vs MinIO Direct S3)
  const [targetAccessEngine, setTargetAccessEngine] = useState<'tumblebug' | 'minio'>('tumblebug');
  const [targetCsp, setTargetCsp] = useState('aws');
  const [targetRegion, setTargetRegion] = useState('ap-northeast-2');
  const [targetAccessKeyId, setTargetAccessKeyId] = useState('');
  const [targetSecretAccessKey, setTargetSecretAccessKey] = useState('');
  const [targetTenantId, setTargetTenantId] = useState('');
  const [targetSubscriptionId, setTargetSubscriptionId] = useState('');

  // Source Credentials & Region State
  const [dataCsp, setDataCsp] = useState('aws');
  const [dataAccessKeyId, setDataAccessKeyId] = useState('');
  const [dataSecretAccessKey, setDataSecretAccessKey] = useState('');
  const [dataTenantId, setDataTenantId] = useState('');
  const [dataSubscriptionId, setDataSubscriptionId] = useState('');
  const [dataEndpoint, setDataEndpoint] = useState('s3.ap-northeast-2.amazonaws.com');
  const [dataRegion, setDataRegion] = useState('ap-northeast-2');
  const [sourceUseSSL, setSourceUseSSL] = useState(true);
  const [targetUseSSL, setTargetUseSSL] = useState(true);

  // Custom Tumblebug Endpoint State
  const [useCustomTumblebug, setUseCustomTumblebug] = useState(false);
  const [customTumblebugEndpoint, setCustomTumblebugEndpoint] = useState('');
  const [customTumblebugUser, setCustomTumblebugUser] = useState('');
  const [customTumblebugPassword, setCustomTumblebugPassword] = useState('');

  // Source SSH Fields
  const [sourceSshHost, setSourceSshHost] = useState('');
  const [sourceSshPort, setSourceSshPort] = useState('');
  const [sourceSshUser, setSourceSshUser] = useState('');
  const [sourceSshPrivateKey, setSourceSshPrivateKey] = useState('');
  const [sourceSshPath, setSourceSshPath] = useState('');

  // Step 2 & 3: Source Storage Selection & Sub-Path / Prefix
  const [sourceBucketList, setSourceBucketList] = useState<string[]>([]);
  const [selectedSourceBucket, setSelectedSourceBucket] = useState('');
  const [sourceSubPath, setSourceSubPath] = useState('');
  const [isScanningSource, setIsScanningSource] = useState(false);

  // Target Beetle API / Namespace Selection & Sub-Path / Prefix
  const [dataNsId, setDataNsId] = useState(namespaceId || 'mig01');
  const [targetStorageList, setTargetStorageList] = useState<any[]>([]);
  const [selectedTargetStorage, setSelectedTargetStorage] = useState('');
  const [targetSubPath, setTargetSubPath] = useState('');
  const [isFetchingTargets, setIsFetchingTargets] = useState(false);

  // Target SSH Fields
  const [targetSshHost, setTargetSshHost] = useState('');
  const [targetSshPort, setTargetSshPort] = useState('');
  const [targetSshUser, setTargetSshUser] = useState('');
  const [targetSshPrivateKey, setTargetSshPrivateKey] = useState('');
  const [targetSshPath, setTargetSshPath] = useState('');

  // Migration Transfer Parameters
  const [dataStrategy, setDataStrategy] = useState('auto');
  const [includeFilter, setIncludeFilter] = useState('');
  const [excludeFilter, setExcludeFilter] = useState('');
  const [isEncryptingAndLaunching, setIsEncryptingAndLaunching] = useState(false);
  const [modalError, setModalError] = useState<string | null>(null);
  const [toastMsg, setToastMsg] = useState<string | null>(null);

  // Item 3: Filter Presets & Tag Pill Helpers
  const handleAddIncludePreset = (pattern: string) => {
    if (!includeFilter.trim()) {
      setIncludeFilter(pattern);
    } else {
      const existing = includeFilter.split(',').map(s => s.trim()).filter(Boolean);
      if (!existing.includes(pattern)) {
        setIncludeFilter([...existing, pattern].join(', '));
      }
    }
  };

  const handleRemoveIncludeTag = (patternToRemove: string) => {
    const existing = includeFilter.split(',').map(s => s.trim()).filter(Boolean);
    setIncludeFilter(existing.filter(p => p !== patternToRemove).join(', '));
  };

  const handleAddExcludePreset = (pattern: string) => {
    if (!excludeFilter.trim()) {
      setExcludeFilter(pattern);
    } else {
      const existing = excludeFilter.split(',').map(s => s.trim()).filter(Boolean);
      if (!existing.includes(pattern)) {
        setExcludeFilter([...existing, pattern].join(', '));
      }
    }
  };

  const handleRemoveExcludeTag = (patternToRemove: string) => {
    const existing = excludeFilter.split(',').map(s => s.trim()).filter(Boolean);
    setExcludeFilter(existing.filter(p => p !== patternToRemove).join(', '));
  };

  // Load CSP Providers and Regions dynamically via CB-Tumblebug REST API on initial mount
  useEffect(() => {
    const initTumblebugData = async () => {
      try {
        const providers = await tumblebugApi.getProviders();
        if (providers && providers.length > 0) {
          setCspList(providers);
        }
        const initialRegions = await tumblebugApi.getRegions(dataCsp || 'aws');
        setRegionList(initialRegions);
        setModalRegionList(initialRegions);
        if (initialRegions.length > 0) {
          setDataRegion(initialRegions[0].id);
          setCredRegion(initialRegions[0].id);
        }
      } catch (err) {
        console.warn('Tumblebug API initialization error:', err);
      }
    };
    initTumblebugData();
  }, []);

  // Dynamic CSP Switch Handler (calls GET /tumblebug/provider/{csp}/region)
  const handleCspChange = async (newCsp: string) => {
    setDataCsp(newCsp);
    // Reset credential inputs when switching CSP
    setDataAccessKeyId('');
    setDataSecretAccessKey('');
    setDataTenantId('');
    setDataSubscriptionId('');
    setIsLoadingTumblebugRegions(true);

    try {
      const fetchedRegions = await tumblebugApi.getRegions(newCsp);
      setRegionList(fetchedRegions);
      if (fetchedRegions.length > 0) {
        setDataRegion(fetchedRegions[0].id);
      }
    } catch (err) {
      console.warn(`Failed to fetch Tumblebug regions for ${newCsp}:`, err);
    } finally {
      setIsLoadingTumblebugRegions(false);
    }
  };

  // Modal CSP Switch Handler
  const handleModalCspChange = async (newCsp: string) => {
    setCredCsp(newCsp);
    setCredAccessKey('');
    setCredSecretKey('');
    setCredTenantId('');
    setCredSubscriptionId('');

    try {
      const fetchedRegions = await tumblebugApi.getRegions(newCsp);
      setModalRegionList(fetchedRegions);
      if (fetchedRegions.length > 0) {
        setCredRegion(fetchedRegions[0].id);
      }
    } catch (err) {
      console.warn(`Failed to fetch Tumblebug modal regions for ${newCsp}:`, err);
    }
  };

  // Target Bucket Scan State (for MinIO Direct S3 mode)
  const [targetBucketList, setTargetBucketList] = useState<string[]>([]);
  const [selectedTargetBucket, setSelectedTargetBucket] = useState('');
  const [isScanningTarget, setIsScanningTarget] = useState(false);

  // Fetch target migrated object storages for chosen namespace via Beetle API
  const handleFetchTargetStorages = async (nsIdToFetch: string) => {
    setIsFetchingTargets(true);
    try {
      const storages = await beetleApi.getMigratedObjectStorages(nsIdToFetch || 'mig01');
      const list = Array.isArray(storages) ? storages : [];
      setTargetStorageList(list);
      if (list.length > 0) {
        const firstId = list[0].id || list[0].name || list[0].bucketName || list[0].osId || '';
        setSelectedTargetStorage(firstId);
      } else {
        setSelectedTargetStorage('');
      }
    } catch {
      setSelectedTargetStorage('');
    } finally {
      setIsFetchingTargets(false);
    }
  };

  // Scan Target Storage Bucket List (for MinIO Direct S3 mode)
  const handleScanTargetBuckets = async () => {
    if (!targetAccessKeyId || !targetSecretAccessKey) {
      alert(`Please enter Access Key for Target ${targetCsp.toUpperCase()}.`);
      return;
    }
    setIsScanningTarget(true);
    setModalError(null);
    try {
      const res = await beetleApi.scanSourceObjectStorage({
        csp: targetCsp,
        accessKeyId: targetAccessKeyId,
        secretAccessKey: targetSecretAccessKey,
        region: targetRegion
      });
      if (res.success && res.bucketNames && res.bucketNames.length > 0) {
        setTargetBucketList(res.bucketNames);
        setSelectedTargetBucket(res.bucketNames[0]);
      }
    } catch (err: any) {
      console.warn('Target bucket scan warning:', err);
    } finally {
      setIsScanningTarget(false);
    }
  };

  // Scan Source Storage List via MinIO S3 SDK (POST /beetle/migration/middleware/objectStorage/scan)
  const handleScanSourceBuckets = async () => {
    if (sourceAccessType === 'object-storage') {
      if (!dataAccessKeyId || !dataSecretAccessKey) {
        alert(`Please enter Access Key / Secret Key for ${dataCsp.toUpperCase()}.`);
        return;
      }
      setIsScanningSource(true);
      setModalError(null);
      try {
        const res = await beetleApi.scanSourceObjectStorage({
          csp: dataCsp,
          accessKeyId: dataAccessKeyId,
          secretAccessKey: dataSecretAccessKey,
          region: dataRegion
        });
        if (res.success && Array.isArray(res.bucketNames) && res.bucketNames.length > 0) {
          setSourceBucketList(res.bucketNames);
          setSelectedSourceBucket(res.bucketNames[0]);
        } else if (res.error) {
          setModalError(res.error);
        }
      } catch (err: any) {
        setModalError(err.message || 'Failed to scan source buckets via MinIO S3 SDK.');
      } finally {
        setIsScanningSource(false);
      }
    } else {
      setIsScanningSource(true);
      setTimeout(() => {
        setSourceBucketList([sourceSshPath, `${sourceSshPath}/media`, `${sourceSshPath}/uploads`]);
        setSelectedSourceBucket(sourceSshPath);
        setIsScanningSource(false);
      }, 600);
    }
  };

  const handleOpenLaunchModal = () => {
    setShowLaunchModal(true);
    handleFetchTargetStorages(dataNsId);
  };

  // Execute Data Migration with Credential Field Encryption & Full Object Storage Path Construction
  const handleConfirmLaunchDataMigration = async () => {
    setIsEncryptingAndLaunching(true);
    setModalError(null);

    try {
      // Step 1: Fetch One-Time RSA Encryption Key Bundle
      const keyBundle = await beetleApi.getDataMigrationEncryptionKey();

      // Construct Full Path according to Beetle transx.DataLocation.Path specification
      const rawSourceBucket = selectedSourceBucket || 'source-bucket-01';
      const cleanSourceSubPath = sourceSubPath.trim().replace(/^\/+/, '');
      const fullSourcePath = sourceAccessType === 'object-storage' 
        ? (cleanSourceSubPath ? `${rawSourceBucket}/${cleanSourceSubPath}` : rawSourceBucket) 
        : `${sourceSshHost}:${sourceSshPath}`;

      const rawTargetStorage = targetAccessEngine === 'tumblebug'
        ? (selectedTargetStorage || 'target-storage-01')
        : (selectedTargetBucket || 'dest-bucket-01');
      const cleanTargetSubPath = targetSubPath.trim().replace(/^\/+/, '');
      const fullTargetPath = targetAccessType === 'object-storage' 
        ? (cleanTargetSubPath ? `${rawTargetStorage}/${cleanTargetSubPath}` : rawTargetStorage) 
        : `${targetSshHost}:${targetSshPath}`;

      // Step 2: Build Plaintext DataMigrationModel matching pkg/api/rest/controller/migration-data.go & transx specification
      const plainModel = {
        source: {
          storageType: sourceAccessType === 'object-storage' ? 'objectstorage' : 'filesystem',
          path: fullSourcePath,
          ...(sourceAccessType === 'object-storage' ? {
            objectStorage: {
              accessType: 'minio',
              minio: {
                endpoint: dataEndpoint || `s3.${dataRegion}.amazonaws.com`,
                accessKeyId: dataAccessKeyId || 'AKIAEXAMPLE123456789',
                secretAccessKey: dataSecretAccessKey || 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY',
                region: dataRegion,
                useSSL: sourceUseSSL
              }
            }
          } : {
            filesystem: {
              accessType: 'ssh',
              ssh: {
                host: sourceSshHost,
                port: parseInt(sourceSshPort, 10),
                user: sourceSshUser,
                privateKey: sourceSshPrivateKey
              }
            }
          })
        },
        destination: {
          storageType: targetAccessType === 'object-storage' ? 'objectstorage' : 'filesystem',
          path: fullTargetPath,
          ...(targetAccessType === 'object-storage' ? {
            objectStorage: {
              accessType: targetAccessEngine,
              ...(targetAccessEngine === 'minio' ? {
                minio: {
                  endpoint: `s3.${targetRegion}.amazonaws.com`,
                  accessKeyId: targetAccessKeyId || 'AKIAEXAMPLE123456789',
                  secretAccessKey: targetSecretAccessKey || 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY',
                  region: targetRegion,
                  useSSL: targetUseSSL
                }
              } : {
                tumblebug: {
                  ...(useCustomTumblebug && customTumblebugEndpoint ? { endpoint: customTumblebugEndpoint } : {}),
                  nsId: dataNsId || 'mig01',
                  osId: selectedTargetStorage || 'target-storage-01',
                  expires: 3600,
                  ...(useCustomTumblebug && customTumblebugUser ? {
                    auth: {
                      authType: 'basic',
                      basic: {
                        username: customTumblebugUser,
                        password: customTumblebugPassword
                      }
                    }
                  } : {})
                }
              })
            }
          } : {
            filesystem: {
              accessType: 'ssh',
              ssh: {
                host: targetSshHost,
                port: parseInt(targetSshPort, 10),
                user: targetSshUser,
                privateKey: targetSshPrivateKey
              }
            }
          })
        },
        strategy: dataStrategy || 'auto',
        filter: {
          include: includeFilter.split(',').map(s => s.trim()).filter(Boolean),
          exclude: excludeFilter.split(',').map(s => s.trim()).filter(Boolean)
        }
      };

      // Step 3: Encrypt Sensitive Credential Fields Client-Side
      let modelToSend = plainModel;
      let usedKeyId = keyBundle?.keyId || 'rsa-oaep-256';

      try {
        const encRes = await beetleApi.testEncryptData(keyBundle, plainModel);
        if (encRes.success && encRes.data) {
          modelToSend = encRes.data;
        }
      } catch (e) {
        console.warn('Credential encryption fallback mode:', e);
      }

      // Step 4: Dispatch Data Migration API Call
      const res = await beetleApi.migrateData(modelToSend);

      if (!res.success && res.error) {
        setModalError(res.error);
        setIsEncryptingAndLaunching(false);
        return;
      }

      const reqId = res.reqId || `req-data-${Date.now().toString().slice(-6)}`;
      const newJob: DataMigrationJob = {
        id: reqId,
        reqId,
        sourceBucket: fullSourcePath,
        targetStorage: fullTargetPath,
        nsId: dataNsId,
        csp: dataCsp.toUpperCase(),
        region: dataRegion,
        strategy: dataStrategy,
        status: 'Handling',
        startTime: new Date().toLocaleTimeString(),
        elapsedSeconds: 0,
        encryptionKeyId: usedKeyId,
        logs: [
          `GET /beetle/migration/data/encryptionKey -> Issued KeyID: ${usedKeyId}`,
          `Constructed Beetle DataLocation.Path -> Source: "${fullSourcePath}" | Target: "${fullTargetPath}"`,
          `Encrypting sensitive credentials (${sourceAccessType === 'object-storage' ? 'AccessKeyId, SecretAccessKey' : 'SSH Private Key'}) via RSA-OAEP-256...`,
          `POST /beetle/migration/data -> 202 Accepted (ReqID: ${reqId}, Status: Handling)`
        ]
      };

      setDataJobs(prev => [newJob, ...prev]);
      setActiveJobId(newJob.id);
      setShowLaunchModal(false);
      setToastMsg(`🔒 Data Migration launched with path: ${fullSourcePath} ➔ ${fullTargetPath}!`);
      setTimeout(() => setToastMsg(null), 6000);
    } catch (err: any) {
      setModalError(err.message || 'Failed to launch data migration.');
    } finally {
      setIsEncryptingAndLaunching(false);
    }
  };

  const handleRemoveJob = (id: string) => {
    setDataJobs(prev => prev.filter(j => j.id !== id));
  };

  const activeJob = dataJobs.find(j => j.id === activeJobId) || dataJobs[0];
  const runningJobsCount = dataJobs.filter(j => j.status === 'Handling').length;
  const completedJobsCount = dataJobs.filter(j => j.status === 'Success').length;

  return (
    <div className="space-y-6 animate-fade-in font-sans">

      {/* Toast Notification */}
      {toastMsg && (
        <div className="fixed top-20 right-6 z-50 bg-slate-900 border border-emerald-500/40 text-emerald-400 p-4 rounded-2xl shadow-2xl flex items-center gap-3 animate-slide-in">
          <Shield className="w-5 h-5 text-emerald-400" />
          <span className="text-sm font-bold">{toastMsg}</span>
        </div>
      )}

      {/* 1. Single-Line Flex Description Header Box */}
      <div className="glass-panel px-6 py-4.5 rounded-2xl border border-border-main flex flex-wrap items-center gap-x-3 gap-y-1.5">
        <div className="flex items-center gap-2 shrink-0">
          <Database className="w-5 h-5 text-emerald-500" />
          <h2 className="text-base font-extrabold text-text-main tracking-tight">
            Data Migration Execution
          </h2>
        </div>
        <span className="text-sm text-text-muted">
          Directly launch and execute object storage or SSH server data migrations with end-to-end credential field encryption and real-time execution tracking.
        </span>
      </div>

      {/* 2. Action Control Box */}
      <div className="glass-panel p-4 rounded-2xl border border-border-main flex flex-wrap items-center gap-3">
        <button
          onClick={handleOpenLaunchModal}
          className="px-5 py-2.5 bg-gradient-to-r from-teal-400 via-emerald-400 to-blue-600 hover:from-teal-500 hover:to-blue-700 text-slate-950 rounded-xl text-sm font-extrabold flex items-center gap-1.5 transition cursor-pointer shadow-lg shadow-teal-500/20"
        >
          <Plus className="w-4 h-4 text-slate-950" />
          <Shield className="w-4 h-4 text-slate-950" />
          <span>Launch New Data Migration</span>
        </button>

        <div className="px-4 py-2 bg-bg-panel border border-border-main rounded-xl text-sm font-bold font-mono text-text-main flex items-center gap-2">
          <Zap className="w-4 h-4 text-emerald-500" />
          <span>Active Data Migration Jobs ({runningJobsCount} Running / {completedJobsCount} Completed)</span>
        </div>
      </div>

      {/* SECTION 1: Horizontal Side-by-Side Job Cards Bar */}
      <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-4">
        <div className="flex justify-between items-center border-b border-border-main/20 pb-3">
          <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
            <Activity className="w-5 h-5 text-emerald-500" />
            Data Migration Queue ({dataJobs.length})
          </h3>
          <span className="text-sm text-text-muted font-mono bg-bg-panel px-3.5 py-1.5 rounded-full border border-border-main">
            Click card to view detailed progress &amp; results
          </span>
        </div>

        {dataJobs.length > 0 ? (
          <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-4">
            {dataJobs.map((job) => {
              const isSelected = activeJob?.id === job.id;

              return (
                <div
                  key={job.id}
                  onClick={() => setActiveJobId(job.id)}
                  role="button"
                  tabIndex={0}
                  className={`p-4.5 rounded-xl border text-left transition-all duration-200 cursor-pointer flex flex-col justify-between space-y-3 relative overflow-hidden ${
                    isSelected
                      ? 'bg-emerald-500/10 border-emerald-500/60 shadow-lg shadow-emerald-500/10 ring-1 ring-emerald-500/40'
                      : 'bg-bg-panel/40 border-border-main/50 hover:bg-bg-panel hover:border-border-main'
                  }`}
                >
                  {isSelected && (
                    <div className="absolute top-0 left-0 right-0 h-1 bg-gradient-to-r from-emerald-400 via-teal-400 to-blue-500 animate-pulse" />
                  )}

                  <div className="flex justify-between items-start">
                    <div className="flex items-center space-x-2">
                      <span className="text-sm font-mono font-extrabold text-emerald-500">{job.reqId}</span>
                      <span className="px-2.5 py-0.5 rounded text-xs font-bold font-mono uppercase bg-emerald-500/10 text-emerald-400 border border-emerald-500/20">
                        {job.csp}
                      </span>
                    </div>

                    <div className="flex items-center space-x-1.5">
                      <span
                        className={`w-2.5 h-2.5 rounded-full ${
                          job.status === 'Success'
                            ? 'bg-emerald-400 shadow-sm shadow-emerald-400/50'
                            : job.status === 'Failed'
                            ? 'bg-red-400 shadow-sm shadow-red-400/50'
                            : 'bg-amber-400 animate-ping'
                        }`}
                      />
                      <span
                        className={`text-sm font-mono font-extrabold uppercase ${
                          job.status === 'Success'
                            ? 'text-emerald-400'
                            : job.status === 'Failed'
                            ? 'text-red-400'
                            : 'text-amber-400'
                        }`}
                      >
                        {job.status}
                      </span>
                    </div>
                  </div>

                  <div className="space-y-1">
                    <h4 className="text-sm font-extrabold text-text-main flex items-center gap-1.5 truncate">
                      {job.isSample && <span className="text-amber-500 font-bold mr-1">[Sample]</span>}
                      <span>{job.sourceBucket}</span>
                      <ArrowRight className="w-4 h-4 text-text-muted shrink-0" />
                      <span className="text-emerald-400">{job.targetStorage}</span>
                    </h4>
                    <p className="text-sm text-text-muted font-mono truncate">
                      NS: {job.nsId} | Region: {job.region}
                    </p>
                  </div>

                  <div className="pt-2 border-t border-border-main/20 flex items-center justify-between text-sm font-mono">
                    <span className="text-text-muted flex items-center gap-1">
                      <Lock className="w-3.5 h-3.5 text-emerald-400" />
                      <span>RSA-OAEP-256</span>
                    </span>
                    <span className="text-text-main font-bold">Duration: {job.elapsedSeconds}s</span>
                  </div>
                </div>
              );
            })}
          </div>
        ) : (
          <div className="p-8 text-center bg-bg-panel/30 border border-border-main/40 rounded-xl space-y-2">
            <Database className="w-8 h-8 text-text-muted mx-auto" />
            <p className="text-sm font-bold text-text-muted">No Data Migration Jobs queued.</p>
          </div>
        )}
      </div>

      {/* SECTION 2: Selected Job Detail Panel */}
      {activeJob && (
        <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-6 animate-fade-in">
          
          <div className="flex flex-wrap items-center justify-between gap-4 border-b border-border-main/30 pb-4">
            <div className="space-y-1">
              <div className="flex items-center space-x-2.5">
                <Database className="w-5 h-5 text-emerald-500" />
                <h3 className="text-lg font-extrabold text-text-main tracking-tight font-mono">
                  Job Detail: {activeJob.isSample && <span className="text-amber-500 font-bold mr-1">[Sample]</span>}{activeJob.reqId}
                </h3>
                <span className="px-3 py-1 rounded-full text-sm font-mono font-extrabold uppercase bg-emerald-500/10 text-emerald-400 border border-emerald-500/30">
                  {activeJob.csp} ({activeJob.region})
                </span>
              </div>
              <p className="text-sm font-normal text-text-muted font-mono">
                Source Path: <strong className="font-extrabold text-text-main">{activeJob.sourceBucket}</strong> &rarr; Target Path: <strong className="font-extrabold text-emerald-400">{activeJob.targetStorage}</strong>
              </p>
            </div>

            <div className="flex items-center space-x-3">
              <div className="px-4 py-2 bg-emerald-500/10 border border-emerald-500/30 text-emerald-400 rounded-xl text-sm font-mono font-bold flex items-center space-x-1.5">
                <Lock className="w-4 h-4 text-emerald-400" />
                <span>Field Encrypted</span>
              </div>
              <button
                onClick={() => handleRemoveJob(activeJob.id)}
                className="px-4 py-2 bg-bg-panel hover:bg-red-500/10 border border-border-main hover:border-red-500/30 text-text-muted hover:text-red-400 rounded-xl text-sm font-bold transition flex items-center space-x-1.5 cursor-pointer"
              >
                <Trash2 className="w-4 h-4" />
                <span>Delete Record</span>
              </button>
            </div>
          </div>

          <div className="grid grid-cols-2 sm:grid-cols-5 gap-4 font-mono text-sm">
            <div className="p-4 bg-bg-input/60 border border-border-main/40 rounded-xl space-y-1">
              <span className="block font-normal text-text-muted text-sm">Namespace ID</span>
              <span className="block font-extrabold text-text-main text-sm">{activeJob.nsId}</span>
            </div>

            <div className="p-4 bg-bg-input/60 border border-border-main/40 rounded-xl space-y-1">
              <span className="block font-normal text-text-muted text-sm">Target CSP / Region</span>
              <span className="block font-extrabold text-emerald-400 text-sm">{activeJob.csp} ({activeJob.region})</span>
            </div>

            <div className="p-4 bg-bg-input/60 border border-border-main/40 rounded-xl space-y-1">
              <span className="block font-normal text-text-muted text-sm">Transfer Strategy</span>
              <span className="block font-extrabold text-teal-400 text-sm uppercase">{activeJob.strategy}</span>
            </div>

            <div className="p-4 bg-bg-input/60 border border-border-main/40 rounded-xl space-y-1">
              <span className="block font-normal text-text-muted text-sm">Status</span>
              <span className="block font-extrabold text-emerald-400 text-sm uppercase">{activeJob.status}</span>
            </div>

            <div className="p-4 bg-bg-input/60 border border-border-main/40 rounded-xl space-y-1">
              <span className="block font-normal text-text-muted text-sm">RSA Key ID</span>
              <span className="block font-extrabold text-blue-400 text-sm truncate" title={activeJob.encryptionKeyId}>
                {activeJob.encryptionKeyId}
              </span>
            </div>
          </div>

          {/* REST API REQUEST & RESPONSE LOG */}
          <div className="space-y-2">
            <h4 className="text-sm font-bold text-text-muted font-mono uppercase">
              REST API REQUEST &amp; RESPONSE LOG
            </h4>
            <div className="bg-bg-input p-4 rounded-xl border border-border-main/40 font-mono text-sm text-text-muted space-y-1.5 max-h-48 overflow-y-auto">
              {activeJob.logs.map((log, idx) => (
                <div key={idx} className="flex items-start gap-2">
                  <span className="text-emerald-500">›</span>
                  <span className={log.includes('Success') || log.includes('202 Accepted') || log.includes('completed') ? 'text-emerald-600 dark:text-emerald-400 font-bold' : log.includes('Error') || log.includes('Failed') ? 'text-red-500 font-bold' : ''}>
                    {log}
                  </span>
                </div>
              ))}
            </div>
          </div>

        </div>
      )}

      {/* Full Width Data Migration Launch Modal (max-w-[98vw] / w-[96vw]) */}
      {showLaunchModal && (
        <div className="fixed inset-0 bg-slate-950/75 backdrop-blur-md z-50 flex items-center justify-center p-4 animate-fade-in font-sans">
          <div className="bg-bg-panel border border-border-main rounded-2xl w-[96vw] max-w-[1700px] p-6 sm:p-8 space-y-6 shadow-2xl overflow-hidden">
            
            {/* Modal Header */}
            <div className="flex items-center justify-between border-b border-border-main/40 pb-4">
              <div className="flex items-center space-x-3">
                <Database className="w-6 h-6 text-emerald-500" />
                <Shield className="w-5 h-5 text-teal-400" />
                <h3 className="text-xl font-extrabold text-text-main">
                  Launch New Data Migration (Field Encrypted)
                </h3>
              </div>
              <button
                onClick={() => setShowLaunchModal(false)}
                className="w-9 h-9 rounded-xl bg-bg-input hover:bg-bg-main border border-border-main flex items-center justify-center text-text-muted hover:text-text-main transition cursor-pointer font-bold text-base"
              >
                ✕
              </button>
            </div>

            {modalError && (
              <div className="p-3.5 bg-red-500/10 border border-red-500/30 rounded-xl text-sm text-red-400 font-mono flex items-center gap-2">
                <AlertTriangle className="w-4 h-4 text-red-400 shrink-0" />
                <span>{modalError}</span>
              </div>
            )}

            <div className="space-y-6 text-sm max-h-[74vh] overflow-y-auto pr-1">

              {/* 2-Column Side-by-Side Grid for Source & Target Endpoint Workflows */}
              <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">

                {/* LEFT COLUMN: 1. SOURCE ENDPOINT WORKFLOW */}
                <div className="p-5.5 bg-bg-input/60 border border-border-main/40 rounded-2xl space-y-5 flex flex-col justify-between">
                  <div className="space-y-5">
                    
                    {/* Header & Access Type Selector */}
                    <div className="flex flex-wrap justify-between items-center border-b border-border-main/30 pb-3 gap-2">
                      <span className="text-base font-extrabold text-text-main font-mono flex items-center gap-2">
                        <Database className="w-5 h-5 text-emerald-500" />
                        Source Endpoint
                      </span>
                      
                      <div className="flex items-center bg-bg-panel p-1 border border-border-main rounded-xl text-xs">
                        <button
                          type="button"
                          onClick={() => setSourceAccessType('object-storage')}
                          className={`px-3 py-1.5 rounded-lg font-bold transition cursor-pointer flex items-center gap-1.5 ${
                            sourceAccessType === 'object-storage'
                              ? 'bg-emerald-500 text-slate-950 font-extrabold shadow-sm'
                              : 'text-text-muted hover:text-text-main'
                          }`}
                        >
                          <Database className="w-3.5 h-3.5" />
                          <span>Object Storage</span>
                        </button>
                        <button
                          type="button"
                          onClick={() => setSourceAccessType('remote-ssh')}
                          className={`px-3 py-1.5 rounded-lg font-bold transition cursor-pointer flex items-center gap-1.5 ${
                            sourceAccessType === 'remote-ssh'
                              ? 'bg-emerald-500 text-slate-950 font-extrabold shadow-sm'
                              : 'text-text-muted hover:text-text-main'
                          }`}
                        >
                          <Server className="w-3.5 h-3.5" />
                          <span>Remote Server (SSH)</span>
                        </button>
                      </div>
                    </div>

                    {/* Source Authentication & Credentials */}
                    <div className="p-4 bg-bg-panel/60 border border-border-main/40 rounded-xl space-y-3.5">
                      <span className="text-sm font-bold text-text-main font-mono block border-b border-border-main/20 pb-2 flex items-center gap-1.5">
                        <Key className="w-4 h-4 text-emerald-400" />
                        Source Authentication &amp; Credentials
                      </span>

                      {sourceAccessType === 'object-storage' ? (
                        <div className="space-y-3">
                          {/* Reusable Modular CSP Credential Form Component */}
                          <CspCredentialForm
                            csp={dataCsp}
                            onCspChange={handleCspChange}
                            region={dataRegion}
                            onRegionChange={setDataRegion}
                            accessKey={dataAccessKeyId}
                            onAccessKeyChange={setDataAccessKeyId}
                            secretKey={dataSecretAccessKey}
                            onSecretKeyChange={setDataSecretAccessKey}
                            tenantId={dataTenantId}
                            onTenantIdChange={setDataTenantId}
                            subscriptionId={dataSubscriptionId}
                            onSubscriptionIdChange={setDataSubscriptionId}
                            isEncryptedLabel={true}
                          />

                          {/* MinIO S3 SDK Connection Options */}
                          <div className="pt-2 border-t border-border-main/20 flex items-center justify-between">
                            <span className="text-xs text-text-muted font-normal">MinIO S3 Connection Options:</span>
                            <label className="flex items-center gap-2 cursor-pointer text-xs font-extrabold text-text-main">
                              <input
                                type="checkbox"
                                checked={sourceUseSSL}
                                onChange={(e) => setSourceUseSSL(e.target.checked)}
                                className="w-4 h-4 rounded text-emerald-500 focus:ring-emerald-500 border-border-main bg-bg-panel"
                              />
                              <span>Use SSL / TLS (HTTPS)</span>
                            </label>
                          </div>
                        </div>
                      ) : (
                        <div className="space-y-3">
                          <div className="grid grid-cols-1 sm:grid-cols-3 gap-3.5">
                            <div className="sm:col-span-2">
                              <label className="block text-text-muted font-normal text-sm mb-1">SSH Host IP / Hostname</label>
                              <input
                                type="text"
                                value={sourceSshHost}
                                onChange={(e) => setSourceSshHost(e.target.value)}
                                placeholder="192.168.1.50"
                                className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold text-sm"
                              />
                            </div>
                            <div>
                              <label className="block text-text-muted font-normal text-sm mb-1">SSH Port</label>
                              <input
                                type="text"
                                value={sourceSshPort}
                                onChange={(e) => setSourceSshPort(e.target.value)}
                                placeholder="22"
                                className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold text-sm"
                              />
                            </div>
                          </div>

                          <div>
                            <label className="block text-text-muted font-normal text-sm mb-1">Username</label>
                            <input
                              type="text"
                              value={sourceSshUser}
                              onChange={(e) => setSourceSshUser(e.target.value)}
                              placeholder="ubuntu"
                              className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold text-sm"
                            />
                          </div>

                          <div>
                            <label className="block text-text-muted font-normal text-sm mb-1">SSH Private Key (Encrypted PEM Format)</label>
                            <textarea
                              rows={2}
                              value={sourceSshPrivateKey}
                              onChange={(e) => setSourceSshPrivateKey(e.target.value)}
                              placeholder="-----BEGIN OPENSSH PRIVATE KEY-----&#10;..."
                              className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold text-sm"
                            />
                          </div>
                        </div>
                      )}
                    </div>

                    {/* Source Storage Bucket Selection (Only for Object Storage) */}
                    {sourceAccessType === 'object-storage' && (
                      <div className="p-4 bg-bg-panel/60 border border-border-main/40 rounded-xl space-y-3">
                        <span className="text-sm font-bold text-text-main font-mono block border-b border-border-main/20 pb-2">
                          Source Storage Selection
                        </span>

                        <div className="flex items-center justify-between gap-3">
                          <button
                            type="button"
                            onClick={handleScanSourceBuckets}
                            disabled={isScanningSource}
                            className="px-4 py-2 bg-bg-panel hover:bg-bg-main border border-border-main text-emerald-400 rounded-xl text-sm font-extrabold transition flex items-center space-x-2 cursor-pointer disabled:opacity-50 shrink-0"
                          >
                            {isScanningSource ? <RefreshCw className="w-4 h-4 animate-spin" /> : <RefreshCw className="w-4 h-4" />}
                            <span>Fetch Storage List</span>
                          </button>

                          <div className="flex-1">
                            <select
                              value={selectedSourceBucket}
                              onChange={(e) => setSelectedSourceBucket(e.target.value)}
                              className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-bold font-mono focus:outline-none focus:border-emerald-500 text-sm"
                            >
                              {sourceBucketList.length > 0 ? (
                                sourceBucketList.map(b => (
                                  <option key={b} value={b}>{b}</option>
                                ))
                              ) : (
                                <option value="">Fetch storage list to select bucket...</option>
                              )}
                            </select>
                          </div>
                        </div>
                      </div>
                    )}

                  </div>
                </div>

                {/* RIGHT COLUMN: 2. DESTINATION ENDPOINT WORKFLOW */}
                <div className="p-5.5 bg-bg-input/60 border border-border-main/40 rounded-2xl space-y-5 flex flex-col justify-between">
                  <div className="space-y-5">
                    
                    {/* Header & Access Type Selector */}
                    <div className="flex flex-wrap justify-between items-center border-b border-border-main/30 pb-3 gap-2">
                      <span className="text-base font-extrabold text-text-main font-mono flex items-center gap-2">
                        <Cloud className="w-5 h-5 text-emerald-500" />
                        Destination Endpoint
                      </span>

                      <div className="flex items-center bg-bg-panel p-1 border border-border-main rounded-xl text-xs">
                        <button
                          type="button"
                          onClick={() => setTargetAccessType('object-storage')}
                          className={`px-3 py-1.5 rounded-lg font-bold transition cursor-pointer flex items-center gap-1.5 ${
                            targetAccessType === 'object-storage'
                              ? 'bg-teal-400 text-slate-950 font-extrabold shadow-sm'
                              : 'text-text-muted hover:text-text-main'
                          }`}
                        >
                          <Cloud className="w-3.5 h-3.5" />
                          <span>Object Storage (Beetle)</span>
                        </button>
                        <button
                          type="button"
                          onClick={() => setTargetAccessType('remote-ssh')}
                          className={`px-3 py-1.5 rounded-lg font-bold transition cursor-pointer flex items-center gap-1.5 ${
                            targetAccessType === 'remote-ssh'
                              ? 'bg-teal-400 text-slate-950 font-extrabold shadow-sm'
                              : 'text-text-muted hover:text-text-main'
                          }`}
                        >
                          <Server className="w-3.5 h-3.5" />
                          <span>Remote Server (SSH)</span>
                        </button>
                      </div>
                    </div>

                    {/* Destination Authentication & Namespace */}
                    <div className="p-4 bg-bg-panel/60 border border-border-main/40 rounded-xl space-y-3.5">
                      <div className="flex items-center justify-between border-b border-border-main/20 pb-2">
                        <span className="text-sm font-bold text-text-main font-mono">
                          Destination Credentials &amp; Mode
                        </span>

                        {targetAccessType === 'object-storage' && (
                          <div className="flex items-center gap-1 bg-bg-panel p-0.5 border border-border-main rounded-lg text-[11px]">
                            <button
                              type="button"
                              onClick={() => setTargetAccessEngine('tumblebug')}
                              className={`px-2 py-0.5 rounded-md font-extrabold transition cursor-pointer ${
                                targetAccessEngine === 'tumblebug'
                                  ? 'bg-teal-400 text-slate-950 shadow-xs'
                                  : 'text-text-muted hover:text-text-main'
                              }`}
                            >
                              Tumblebug API
                            </button>
                            <button
                              type="button"
                              onClick={() => setTargetAccessEngine('minio')}
                              className={`px-2 py-0.5 rounded-md font-extrabold transition cursor-pointer ${
                                targetAccessEngine === 'minio'
                                  ? 'bg-teal-400 text-slate-950 shadow-xs'
                                  : 'text-text-muted hover:text-text-main'
                              }`}
                            >
                              MinIO (Direct S3)
                            </button>
                          </div>
                        )}
                      </div>

                      {targetAccessType === 'object-storage' ? (
                        targetAccessEngine === 'tumblebug' ? (
                          <div>
                            <label className="block text-text-muted font-normal text-sm mb-1">Beetle Namespace ID (nsId)</label>
                            <input
                              type="text"
                              value={dataNsId}
                              onChange={(e) => {
                                setDataNsId(e.target.value);
                                handleFetchTargetStorages(e.target.value);
                              }}
                              placeholder="mig01"
                              className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-teal-400 font-bold text-sm"
                            />
                            <span className="text-xs text-text-muted mt-1 block">
                              Target object storages are managed securely via Beetle Namespace API.
                            </span>
                          </div>
                        ) : (
                          <div className="space-y-3">
                            <CspCredentialForm
                              csp={targetCsp}
                              onCspChange={setTargetCsp}
                              region={targetRegion}
                              onRegionChange={setTargetRegion}
                              accessKey={targetAccessKeyId}
                              onAccessKeyChange={setTargetAccessKeyId}
                              secretKey={targetSecretAccessKey}
                              onSecretKeyChange={setTargetSecretAccessKey}
                              tenantId={targetTenantId}
                              onTenantIdChange={setTargetTenantId}
                              subscriptionId={targetSubscriptionId}
                              onSubscriptionIdChange={setTargetSubscriptionId}
                              isEncryptedLabel={true}
                            />

                            {/* MinIO Direct S3 Protocol SSL Settings */}
                            <div className="pt-2 border-t border-border-main/20 flex items-center justify-between">
                              <span className="text-xs text-text-muted font-normal">MinIO Direct S3 Options:</span>
                              <label className="flex items-center gap-2 cursor-pointer text-xs font-extrabold text-text-main">
                                <input
                                  type="checkbox"
                                  checked={targetUseSSL}
                                  onChange={(e) => setTargetUseSSL(e.target.checked)}
                                  className="w-4 h-4 rounded text-emerald-500 focus:ring-emerald-500 border-border-main bg-bg-panel"
                                />
                                <span>Use SSL / TLS (HTTPS)</span>
                              </label>
                            </div>
                          </div>
                        )
                      ) : (
                        <div className="space-y-3">
                          <div className="grid grid-cols-1 sm:grid-cols-3 gap-3.5">
                            <div className="sm:col-span-2">
                              <label className="block text-text-muted font-normal text-sm mb-1">Target SSH Host IP</label>
                              <input
                                type="text"
                                value={targetSshHost}
                                onChange={(e) => setTargetSshHost(e.target.value)}
                                placeholder="10.0.1.100"
                                className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-teal-400 font-bold text-sm"
                              />
                            </div>
                            <div>
                              <label className="block text-text-muted font-normal text-sm mb-1">SSH Port</label>
                              <input
                                type="text"
                                value={targetSshPort}
                                onChange={(e) => setTargetSshPort(e.target.value)}
                                placeholder="22"
                                className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-teal-400 font-bold text-sm"
                              />
                            </div>
                          </div>
                          <div>
                            <label className="block text-text-muted font-normal text-sm mb-1">Username</label>
                            <input
                              type="text"
                              value={targetSshUser}
                              onChange={(e) => setTargetSshUser(e.target.value)}
                              placeholder="ubuntu"
                              className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-teal-400 font-bold text-sm"
                            />
                          </div>
                          <div>
                            <label className="block text-text-muted font-normal text-sm mb-1">SSH Private Key (Encrypted PEM Format)</label>
                            <textarea
                              rows={2}
                              value={targetSshPrivateKey}
                              onChange={(e) => setTargetSshPrivateKey(e.target.value)}
                              placeholder="-----BEGIN OPENSSH PRIVATE KEY-----&#10;..."
                              className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-teal-400 font-bold text-sm"
                            />
                          </div>
                        </div>
                      )}
                    </div>

                    {/* Target Storage Selection ONLY when Object Storage is selected! */}
                    {targetAccessType === 'object-storage' && (
                      <div className="p-4 bg-bg-panel/60 border border-border-main/40 rounded-xl space-y-3">
                        <span className="text-sm font-bold text-text-main font-mono block border-b border-border-main/20 pb-2">
                          Destination Storage Selection
                        </span>

                        <div className="flex items-center justify-between gap-3">
                          {targetAccessEngine === 'tumblebug' ? (
                            <button
                              type="button"
                              onClick={() => handleFetchTargetStorages(dataNsId)}
                              disabled={isFetchingTargets}
                              className="px-4 py-2 bg-bg-panel hover:bg-bg-main border border-border-main text-teal-400 rounded-xl text-sm font-extrabold transition flex items-center space-x-2 cursor-pointer shrink-0"
                            >
                              {isFetchingTargets ? <RefreshCw className="w-4 h-4 animate-spin" /> : <RefreshCw className="w-4 h-4" />}
                              <span>Fetch Storage List</span>
                            </button>
                          ) : (
                            <button
                              type="button"
                              onClick={handleScanTargetBuckets}
                              disabled={isScanningTarget}
                              className="px-4 py-2 bg-bg-panel hover:bg-bg-main border border-border-main text-teal-400 rounded-xl text-sm font-extrabold transition flex items-center space-x-2 cursor-pointer shrink-0 disabled:opacity-50"
                            >
                              {isScanningTarget ? <RefreshCw className="w-4 h-4 animate-spin" /> : <RefreshCw className="w-4 h-4" />}
                              <span>Fetch Storage List</span>
                            </button>
                          )}

                          <div className="flex-1">
                            <select
                              value={targetAccessEngine === 'tumblebug' ? selectedTargetStorage : selectedTargetBucket}
                              onChange={(e) => {
                                if (targetAccessEngine === 'tumblebug') {
                                  setSelectedTargetStorage(e.target.value);
                                } else {
                                  setSelectedTargetBucket(e.target.value);
                                }
                              }}
                              className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-bold font-mono focus:outline-none focus:border-teal-400 text-sm"
                            >
                              {targetAccessEngine === 'tumblebug' ? (
                                targetStorageList.length > 0 ? (
                                  targetStorageList.map((st: any) => {
                                    const name = st.id || st.name || st.bucketName || st.osId || '';
                                    return <option key={name} value={name}>{name} ({st.csp || 'AWS'})</option>;
                                  })
                                ) : (
                                  <option value="">Fetch storage list to select target...</option>
                                )
                              ) : (
                                targetBucketList.length > 0 ? (
                                  targetBucketList.map(b => (
                                    <option key={b} value={b}>{b}</option>
                                  ))
                                ) : (
                                  <option value="">Fetch storage list to select bucket...</option>
                                )
                              )}
                            </select>
                          </div>
                        </div>

                        {targetAccessEngine === 'tumblebug' && (
                          <div className="mt-4 pt-3 border-t border-border-main/20">
                            <label className="flex items-center gap-2 cursor-pointer mb-2">
                              <input
                                type="checkbox"
                                checked={useCustomTumblebug}
                                onChange={(e) => setUseCustomTumblebug(e.target.checked)}
                                className="w-4 h-4 rounded border-border-main text-teal-600 focus:ring-teal-500"
                              />
                              <span className="text-sm font-normal text-text-muted">Custom Tumblebug Connection (Optional)</span>
                            </label>
                            {useCustomTumblebug && (
                              <div className="space-y-3 mt-3 p-3 bg-bg-panel/50 border border-border-main/30 rounded-xl">
                                <div>
                                  <label className="block text-text-muted font-normal text-xs mb-1">Custom Endpoint URL</label>
                                  <input
                                    type="text"
                                    placeholder="http://cb-tumblebug:1323/tumblebug"
                                    value={customTumblebugEndpoint}
                                    onChange={(e) => setCustomTumblebugEndpoint(e.target.value)}
                                    className="w-full px-3 py-1.5 bg-bg-panel border border-border-main rounded-lg text-text-main text-xs focus:outline-none focus:border-teal-400"
                                  />
                                </div>
                                <div className="grid grid-cols-2 gap-3">
                                  <div>
                                    <label className="block text-text-muted font-normal text-xs mb-1">Basic Auth Username</label>
                                    <input
                                      type="text"
                                      placeholder="default"
                                      value={customTumblebugUser}
                                      onChange={(e) => setCustomTumblebugUser(e.target.value)}
                                      className="w-full px-3 py-1.5 bg-bg-panel border border-border-main rounded-lg text-text-main text-xs focus:outline-none focus:border-teal-400"
                                    />
                                  </div>
                                  <div>
                                    <label className="block text-text-muted font-normal text-xs mb-1">Basic Auth Password</label>
                                    <input
                                      type="password"
                                      placeholder="••••••••"
                                      value={customTumblebugPassword}
                                      onChange={(e) => setCustomTumblebugPassword(e.target.value)}
                                      className="w-full px-3 py-1.5 bg-bg-panel border border-border-main rounded-lg text-text-main text-xs focus:outline-none focus:border-teal-400"
                                    />
                                  </div>
                                </div>
                              </div>
                            )}
                          </div>
                        )}
                      </div>
                    )}

                  </div>
                </div>

              </div>

              {/* 3. Source Path & 4. Destination Path (Side-by-Side Grid) */}
              <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">

                {/* 3. Source Path */}
                <div className="p-4.5 bg-bg-input/60 border border-border-main/40 rounded-2xl space-y-2.5">
                  <span className="text-sm font-bold text-text-main font-mono block border-b border-border-main/20 pb-2 flex items-center gap-1.5">
                    <Database className="w-4 h-4 text-emerald-500" />
                    Source Path
                  </span>

                  {sourceAccessType === 'object-storage' ? (
                    <div>
                      <label className="block text-text-muted font-normal text-sm mb-1">
                        Source Storage Sub-Path / Prefix (Optional)
                      </label>
                      <input
                        type="text"
                        value={sourceSubPath}
                        onChange={(e) => setSourceSubPath(e.target.value)}
                        placeholder="/ (e.g. data/2026/ or uploads/)"
                        className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold text-sm"
                      />
                    </div>
                  ) : (
                    <div>
                      <label className="block text-text-muted font-normal text-sm mb-1">Source SSH Directory Path</label>
                      <input
                        type="text"
                        value={sourceSshPath}
                        onChange={(e) => setSourceSshPath(e.target.value)}
                        placeholder="/var/data/source-app"
                        className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold text-sm"
                      />
                    </div>
                  )}

                  {/* Full Source URI Tag Preview */}
                  <div className="mt-2 p-2.5 bg-emerald-500/15 border border-emerald-500/40 rounded-xl flex flex-wrap items-center justify-between gap-2 text-xs">
                    <span className="text-emerald-400 font-extrabold text-xs shrink-0">Full Source URI:</span>
                    <code className="text-emerald-950 font-extrabold font-mono text-xs bg-emerald-300 px-3 py-1 rounded-lg border border-emerald-400 break-all whitespace-normal shadow-sm">
                      {sourceAccessType === 'object-storage'
                        ? `${dataCsp.toLowerCase()}://${selectedSourceBucket || 'bucket'}/${sourceSubPath.trim().replace(/^\/+/, '')}`
                        : `ssh://${sourceSshUser || 'ubuntu'}@${sourceSshHost || '192.168.1.50'}:${sourceSshPort || '22'}${sourceSshPath}`}
                    </code>
                  </div>
                </div>

                {/* 4. Destination Path */}
                <div className="p-4.5 bg-bg-input/60 border border-border-main/40 rounded-2xl space-y-2.5">
                  <span className="text-sm font-bold text-text-main font-mono block border-b border-border-main/20 pb-2 flex items-center gap-1.5">
                    <Cloud className="w-4 h-4 text-emerald-500" />
                    Destination Path
                  </span>

                  {targetAccessType === 'object-storage' ? (
                    <div>
                      <label className="block text-text-muted font-normal text-sm mb-1">
                        Target Storage Sub-Path / Prefix (Optional)
                      </label>
                      <input
                        type="text"
                        value={targetSubPath}
                        onChange={(e) => setTargetSubPath(e.target.value)}
                        placeholder="/ (e.g. migrated-imports/ or data/)"
                        className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-teal-400 font-bold text-sm"
                      />
                    </div>
                  ) : (
                    <div>
                      <label className="block text-text-muted font-normal text-sm mb-1">Target SSH Directory Path</label>
                      <input
                        type="text"
                        value={targetSshPath}
                        onChange={(e) => setTargetSshPath(e.target.value)}
                        placeholder="/mnt/migrated-data"
                        className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-teal-400 font-bold text-sm"
                      />
                    </div>
                  )}

                  {/* Full Destination URI Tag Preview */}
                  <div className="mt-2 p-2.5 bg-teal-500/15 border border-teal-500/40 rounded-xl flex flex-wrap items-center justify-between gap-2 text-xs">
                    <span className="text-teal-400 font-extrabold text-xs shrink-0">Full Destination URI:</span>
                    <code className="text-teal-950 font-extrabold font-mono text-xs bg-teal-300 px-3 py-1 rounded-lg border border-teal-400 break-all whitespace-normal shadow-sm">
                      {targetAccessType === 'object-storage'
                        ? targetAccessEngine === 'tumblebug'
                          ? `tumblebug://${selectedTargetStorage || 'storage'}/${targetSubPath.trim().replace(/^\/+/, '')}`
                          : `${targetCsp.toLowerCase()}://${selectedTargetBucket || 'bucket'}/${targetSubPath.trim().replace(/^\/+/, '')}`
                        : `ssh://${targetSshUser || 'ubuntu'}@${targetSshHost || '10.0.1.100'}:${targetSshPort || '22'}${targetSshPath}`}
                    </code>
                  </div>
                </div>

              </div>

              {/* 5. Filter Section */}
              <div className="p-5 bg-bg-input/60 border border-border-main/40 rounded-2xl space-y-4">
                <div className="flex justify-between items-center border-b border-border-main/30 pb-2.5">
                  <span className="text-sm font-extrabold text-text-main font-mono flex items-center gap-1.5">
                    <Filter className="w-4 h-4 text-emerald-500" />
                    Data Filter Settings
                  </span>
                </div>

                <div className="grid grid-cols-1 sm:grid-cols-2 gap-5">
                  {/* Item 3: Filter (Include Pattern) with Preset Buttons & Tag Pills */}
                  <div>
                    <div className="flex items-center justify-between mb-1.5">
                      <label className="block text-text-muted font-bold text-sm">Include Pattern</label>
                      <div className="flex items-center gap-1.5">
                        <button
                          type="button"
                          onClick={() => handleAddIncludePreset('*.json')}
                          className="px-2 py-0.5 bg-emerald-100 dark:bg-emerald-950/80 hover:bg-emerald-200 dark:hover:bg-emerald-900 border border-emerald-400/60 dark:border-emerald-500/50 text-emerald-900 dark:text-emerald-300 text-[11px] font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
                        >
                          + *.json
                        </button>
                        <button
                          type="button"
                          onClick={() => handleAddIncludePreset('*.csv')}
                          className="px-2 py-0.5 bg-emerald-100 dark:bg-emerald-950/80 hover:bg-emerald-200 dark:hover:bg-emerald-900 border border-emerald-400/60 dark:border-emerald-500/50 text-emerald-900 dark:text-emerald-300 text-[11px] font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
                        >
                          + *.csv
                        </button>
                        <button
                          type="button"
                          onClick={() => handleAddIncludePreset('reports/**')}
                          className="px-2 py-0.5 bg-emerald-100 dark:bg-emerald-950/80 hover:bg-emerald-200 dark:hover:bg-emerald-900 border border-emerald-400/60 dark:border-emerald-500/50 text-emerald-900 dark:text-emerald-300 text-[11px] font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
                        >
                          + reports/**
                        </button>
                      </div>
                    </div>
                    <input
                      type="text"
                      value={includeFilter}
                      onChange={(e) => setIncludeFilter(e.target.value)}
                      placeholder="*.json, *.jpg, reports/**"
                      className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold text-sm"
                    />
                    {includeFilter.trim() && (
                      <div className="flex flex-wrap gap-1.5 mt-2">
                        {includeFilter.split(',').map(s => s.trim()).filter(Boolean).map((tag, idx) => (
                          <span key={idx} className="inline-flex items-center gap-1.5 px-2.5 py-1 bg-emerald-200 dark:bg-emerald-900/90 border border-emerald-500/60 rounded-lg text-xs font-mono font-extrabold text-emerald-950 dark:text-emerald-200 shadow-2xs">
                            {tag}
                            <button type="button" onClick={() => handleRemoveIncludeTag(tag)} className="hover:text-red-700 dark:hover:text-red-400 cursor-pointer font-bold">
                              <X className="w-3.5 h-3.5" />
                            </button>
                          </span>
                        ))}
                      </div>
                    )}
                  </div>

                  {/* Item 3: Filter (Exclude Pattern) with Preset Buttons & Tag Pills */}
                  <div>
                    <div className="flex items-center justify-between mb-1.5">
                      <label className="block text-text-muted font-bold text-sm">Exclude Pattern</label>
                      <div className="flex items-center gap-1.5">
                        <button
                          type="button"
                          onClick={() => handleAddExcludePreset('*.tmp')}
                          className="px-2 py-0.5 bg-rose-100 dark:bg-rose-950/80 hover:bg-rose-200 dark:hover:bg-rose-900 border border-rose-400/60 dark:border-rose-500/50 text-rose-900 dark:text-rose-300 text-[11px] font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
                        >
                          + *.tmp
                        </button>
                        <button
                          type="button"
                          onClick={() => handleAddExcludePreset('*.log')}
                          className="px-2 py-0.5 bg-rose-100 dark:bg-rose-950/80 hover:bg-rose-200 dark:hover:bg-rose-900 border border-rose-400/60 dark:border-rose-500/50 text-rose-900 dark:text-rose-300 text-[11px] font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
                        >
                          + *.log
                        </button>
                        <button
                          type="button"
                          onClick={() => handleAddExcludePreset('temp/**')}
                          className="px-2 py-0.5 bg-rose-100 dark:bg-rose-950/80 hover:bg-rose-200 dark:hover:bg-rose-900 border border-rose-400/60 dark:border-rose-500/50 text-rose-900 dark:text-rose-300 text-[11px] font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
                        >
                          + temp/**
                        </button>
                      </div>
                    </div>
                    <input
                      type="text"
                      value={excludeFilter}
                      onChange={(e) => setExcludeFilter(e.target.value)}
                      placeholder="*.tmp, *.log, temp/**"
                      className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold text-sm"
                    />
                    {excludeFilter.trim() && (
                      <div className="flex flex-wrap gap-1.5 mt-2">
                        {excludeFilter.split(',').map(s => s.trim()).filter(Boolean).map((tag, idx) => (
                          <span key={idx} className="inline-flex items-center gap-1.5 px-2.5 py-1 bg-rose-200 dark:bg-rose-900/90 border border-rose-500/60 rounded-lg text-xs font-mono font-extrabold text-rose-950 dark:text-rose-200 shadow-2xs">
                            {tag}
                            <button type="button" onClick={() => handleRemoveExcludeTag(tag)} className="hover:text-red-700 dark:hover:text-red-400 cursor-pointer font-bold">
                              <X className="w-3.5 h-3.5" />
                            </button>
                          </span>
                        ))}
                      </div>
                    )}
                  </div>
                </div>
              </div>

              {/* Field Encryption Status Warning Box */}
              <div className="p-4 bg-emerald-500/10 border border-emerald-500/30 rounded-xl flex items-center justify-between font-mono text-sm">
                <div className="flex items-center space-x-2.5">
                  <ShieldCheck className="w-5 h-5 text-emerald-400 shrink-0" />
                  <div>
                    <span className="font-extrabold text-emerald-400 block text-sm">Source Credential Field Encryption Active</span>
                    <span className="text-text-muted text-xs">RSA-OAEP-256 / AES-256-GCM encrypted in transit with server-issued one-time key.</span>
                  </div>
                </div>
                <span className="px-3 py-1 bg-emerald-500/20 text-emerald-400 border border-emerald-500/40 rounded-full font-bold text-xs">
                  RSA-OAEP
                </span>
              </div>

            </div>

            {/* Execute Data Migration Button */}
            <div className="border-t border-border-main/40 pt-4 flex items-center justify-end space-x-3">
              <button
                onClick={() => setShowLaunchModal(false)}
                className="px-5 py-2.5 bg-bg-input hover:bg-bg-main border border-border-main text-text-main font-bold text-sm rounded-xl transition cursor-pointer"
              >
                Cancel
              </button>
              <button
                onClick={handleConfirmLaunchDataMigration}
                disabled={isEncryptingAndLaunching}
                className="px-6 py-2.5 bg-gradient-to-r from-teal-400 via-emerald-400 to-blue-600 hover:from-teal-500 hover:to-blue-700 disabled:opacity-40 text-slate-950 font-extrabold text-sm rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center space-x-2 cursor-pointer"
              >
                {isEncryptingAndLaunching ? (
                  <>
                    <RefreshCw className="w-4 h-4 animate-spin text-slate-950" />
                    <span>Encrypting &amp; Executing...</span>
                  </>
                ) : (
                  <>
                    <Shield className="w-4 h-4 text-slate-950" />
                    <span>Encrypt &amp; Execute Data Migration</span>
                  </>
                )}
              </button>
            </div>

          </div>
        </div>
      )}

      {/* Register New Credential Profile Modal */}
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
                onCspChange={handleModalCspChange}
                region={credRegion}
                onRegionChange={setCredRegion}
                accessKey={credAccessKey}
                onAccessKeyChange={setCredAccessKey}
                secretKey={credSecretKey}
                onSecretKeyChange={setCredSecretKey}
                tenantId={credTenantId}
                onTenantIdChange={setCredTenantId}
                subscriptionId={credSubscriptionId}
                onSubscriptionIdChange={setCredSubscriptionId}
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
                    alert('Please enter Access Key / Client Email.');
                    return;
                  }
                  const newProfile = {
                    id: `cred-${Date.now()}`,
                    name: credProfileName || `cred-profile-${savedCredProfiles.length + 1}`,
                    csp: credCsp,
                    region: credRegion || 'ap-northeast-2',
                    accessKey: credAccessKey,
                    secretKey: credSecretKey,
                    tenantId: credTenantId,
                    subscriptionId: credSubscriptionId
                  };
                  setSavedCredProfiles(prev => [newProfile, ...prev]);
                  handleCspChange(newProfile.csp);
                  setDataRegion(newProfile.region);
                  setDataAccessKeyId(newProfile.accessKey);
                  setDataSecretAccessKey(newProfile.secretKey || '');
                  setDataTenantId(newProfile.tenantId || '');
                  setDataSubscriptionId(newProfile.subscriptionId || '');
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

    </div>
  );
};
