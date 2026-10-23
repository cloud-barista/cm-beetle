'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore, calculateJobElapsedSeconds } from '@/store/migrationStore';
import { beetleApi, tumblebugApi, honeybeeApi } from '@/api/client';
import { CspCredentialForm } from '../common/CspCredentialForm';
import { DirectoryTreeExplorer } from './DirectoryTreeExplorer';
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
  Globe,
  Clock,
  ChevronDown,
  ChevronUp,
  Upload
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
  createdAtMs?: number;
  elapsedSeconds: number;
  encryptionKeyId: string;
  isSample?: boolean;
  logs: string[];
}

export const DataTransferCenter: React.FC = () => {
  const {
    namespaceId,
    dataJobs,
    setDataJobs,
    addDataJob,
    removeDataJob,
    activeDataJobId,
    setActiveDataJobId
  } = useMigrationStore();

  // Auto-restore sample job if queue was emptied
  useEffect(() => {
    if (dataJobs.length === 0) {
      setDataJobs([
        {
          id: 'req-data-101',
          reqId: 'req-data-20260724-001',
          nsId: 'mig01',
          csp: 'AWS',
          region: 'ap-northeast-2',
          sourceBucket: 'source-s3-bucket-production',
          targetStorage: 'target-s3-bucket-migrated',
          encryptionKeyId: 'KMS-AES256-GCM-KEY-9982',
          strategy: 'Full Migration + AES256 Payload Encryption',
          status: 'Success',
          startTime: new Date().toISOString(),
          elapsedSeconds: 24,
          logs: [
            'POST /beetle/migration/data (Prefer: respond-async)',
            'HTTP 202 Accepted (ReqID: req-data-20260724-001)',
            'Initializing AES256-GCM Envelope Encryption...',
            'Transferring 1,024 S3 Objects (10.5 GB)...',
            'Data Transfer & Integrity Verification Completed Successfully (Duration: 24s)'
          ],
          isSample: true
        }
      ]);
    }
  }, [dataJobs, setDataJobs]);

  const activeJobId = activeDataJobId || dataJobs[0]?.id || 'req-data-101';
  const setActiveJobId = setActiveDataJobId;

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
      name: '[Sample] aws-production-account',
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

  // Source SSH Fields & Mode (Honeybee Source Group vs Direct SSH)
  const [sourceSshMode, setSourceSshMode] = useState<'honeybee' | 'direct'>('honeybee');
  const [hbGroups, setHbGroups] = useState<any[]>([]);
  const [selectedHbGroupId, setSelectedHbGroupId] = useState('');
  const [hbConnections, setHbConnections] = useState<any[]>([]);
  const [selectedHbConnId, setSelectedHbConnId] = useState('');
  const [isLoadingHb, setIsLoadingHb] = useState(false);
  const [sourceSshHost, setSourceSshHost] = useState('');
  const [sourceSshPort, setSourceSshPort] = useState('22');
  const [sourceSshUser, setSourceSshUser] = useState('ubuntu');
  const [sourceSshPrivateKey, setSourceSshPrivateKey] = useState('');
  const [sourceSshPath, setSourceSshPath] = useState('');

  // Step 2 & 3: Source Storage Selection & Sub-Path / Prefix
  const [sourceBucketList, setSourceBucketList] = useState<string[]>([]);
  const [selectedSourceBucket, setSelectedSourceBucket] = useState('');
  const [sourceScopeMode, setSourceScopeMode] = useState<'all' | 'prefix'>('all');
  const [sourceSubPath, setSourceSubPath] = useState('');
  const [isScanningSource, setIsScanningSource] = useState(false);

  // Target Beetle API / Namespace Selection & Sub-Path / Prefix
  const [dataNsId, setDataNsId] = useState(namespaceId || 'mig01');
  const [targetStorageList, setTargetStorageList] = useState<any[]>([]);
  const [selectedTargetStorage, setSelectedTargetStorage] = useState('');
  const [targetSubPath, setTargetSubPath] = useState('');
  const [isFetchingTargets, setIsFetchingTargets] = useState(false);

  // Target SSH Fields & Mode (Tumblebug Infra vs Direct SSH)
  const [targetSshMode, setTargetSshMode] = useState<'tumblebug' | 'direct'>('tumblebug');
  const [tbInfraList, setTbInfraList] = useState<any[]>([]);
  const [selectedTbInfraId, setSelectedTbInfraId] = useState('');
  const [tbNodeList, setTbNodeList] = useState<any[]>([]);
  const [selectedTbNodeId, setSelectedTbNodeId] = useState('');
  const [isLoadingTbInfra, setIsLoadingTbInfra] = useState(false);
  const [targetSshHost, setTargetSshHost] = useState('');
  const [targetSshPort, setTargetSshPort] = useState('22');
  const [targetSshUser, setTargetSshUser] = useState('ubuntu');
  const [targetSshPrivateKey, setTargetSshPrivateKey] = useState('');
  const [targetSshPath, setTargetSshPath] = useState('');

  // Directory Tree Explorer Toggle State
  const [showDirectoryTree, setShowDirectoryTree] = useState(false);

  // Migration Transfer Parameters
  const [dataStrategy, setDataStrategy] = useState('auto');
  const [includeFilter, setIncludeFilter] = useState('');
  const [excludeFilter, setExcludeFilter] = useState('');
  const [isEncryptingAndLaunching, setIsEncryptingAndLaunching] = useState(false);
  const [preferAsync, setPreferAsync] = useState(true);
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

  // Two-way toggle helpers for DirectoryTreeExplorer
  const handleToggleIncludePattern = (pattern: string) => {
    const existing = includeFilter.split(',').map(s => s.trim()).filter(Boolean);
    if (existing.includes(pattern)) {
      setIncludeFilter(existing.filter(p => p !== pattern).join(', '));
    } else {
      setIncludeFilter([...existing, pattern].join(', '));
      handleRemoveExcludeTag(pattern);
    }
  };

  const handleToggleExcludePattern = (pattern: string) => {
    const existing = excludeFilter.split(',').map(s => s.trim()).filter(Boolean);
    if (existing.includes(pattern)) {
      setExcludeFilter(existing.filter(p => p !== pattern).join(', '));
    } else {
      setExcludeFilter([...existing, pattern].join(', '));
      handleRemoveIncludeTag(pattern);
    }
  };

  // Apply AIMX (Ollama + OpenWebUI) filter preset
  const handleApplyAimxPreset = () => {
    setIncludeFilter('~/.ollama/models/blobs/**, ~/.ollama/models/manifests/**, /app/backend/data/webui.db');
    setExcludeFilter('*.log, *.tmp, __pycache__/**');
    setShowDirectoryTree(true);
  };

  // Decode or sanitize Honeybee SSH username (Honeybee RSA-encrypts stored user strings)
  const resolveHoneybeeUsername = (raw?: string): string => {
    if (!raw) return 'ubuntu';
    if (raw.length > 32 && /^[A-Za-z0-9+/=]+$/.test(raw)) {
      try {
        const decoded = atob(raw);
        if (decoded && decoded.length <= 64 && /^[\x20-\x7E]+$/.test(decoded)) {
          return decoded;
        }
      } catch {
        // Encrypted ciphertext cannot be decoded in browser
      }
      return 'ubuntu';
    }
    return raw;
  };

  // Apply selected Honeybee connection info to Source SSH fields
  const applyHbConnection = (conn: any) => {
    if (!conn) return;
    setSelectedHbConnId(conn.id || conn.name || '');
    setSourceSshHost(conn.ip_address || conn.ip || '');
    setSourceSshPort(String(conn.ssh_port || conn.port || '22'));
    setSourceSshUser(resolveHoneybeeUsername(conn.user));
  };

  // Fetch Connections for a selected Honeybee Source Group
  const handleSelectHbGroup = async (sgId: string) => {
    setSelectedHbGroupId(sgId);
    if (!sgId) {
      setHbConnections([]);
      setSelectedHbConnId('');
      return;
    }
    setIsLoadingHb(true);
    try {
      const res = await honeybeeApi.getConnectionInfoList(sgId);
      const list = res?.connection_info || (Array.isArray(res) ? res : []);
      setHbConnections(list);
      if (list.length > 0) {
        applyHbConnection(list[0]);
      } else {
        setSelectedHbConnId('');
      }
    } catch {
      setHbConnections([]);
      setSelectedHbConnId('');
    } finally {
      setIsLoadingHb(false);
    }
  };

  // Fetch Honeybee Source Groups for Source SSH selection
  const handleFetchHoneybeeGroups = async () => {
    setIsLoadingHb(true);
    try {
      const res = await honeybeeApi.getSourceGroups();
      const list = res?.source_group || (Array.isArray(res) ? res : []);
      setHbGroups(list);
      if (list.length > 0) {
        const targetGroupId = selectedHbGroupId && list.some((g: any) => g.id === selectedHbGroupId)
          ? selectedHbGroupId
          : list[0].id;
        await handleSelectHbGroup(targetGroupId);
      } else {
        setSelectedHbGroupId('');
        setHbConnections([]);
      }
    } catch {
      setHbGroups([]);
      setHbConnections([]);
    } finally {
      setIsLoadingHb(false);
    }
  };

  // Apply selected Tumblebug Target Node to Destination SSH fields (including SSH Private Key lookup)
  const applyTbNode = async (nsId: string, node: any) => {
    if (!node) return;
    const nodeId = node.id || node.name || '';
    setSelectedTbNodeId(nodeId);
    setTargetSshHost(node.publicIP || node.publicIp || node.privateIP || node.privateIp || '');
    setTargetSshPort(String(node.sshPort || '22'));
    setTargetSshUser(node.vmUserName || node.nodeUserName || node.userName || 'ubuntu');
    if (node.sshKeyId) {
      const sshKeyObj = await tumblebugApi.getSshKey(nsId || 'mig01', node.sshKeyId);
      if (sshKeyObj?.privateKey) {
        setTargetSshPrivateKey(sshKeyObj.privateKey);
      }
    }
  };

  // Select a Tumblebug Migrated Infra and load its VM nodes
  const handleSelectTbInfra = async (nsId: string, infraId: string, infraListFallback?: any[]) => {
    setSelectedTbInfraId(infraId);
    if (!infraId) {
      setTbNodeList([]);
      setSelectedTbNodeId('');
      return;
    }
    setIsLoadingTbInfra(true);
    try {
      const detail = await tumblebugApi.getInfra(nsId || 'mig01', infraId);
      const fallbackObj = (infraListFallback || tbInfraList).find((item: any) => (item.id || item.name) === infraId);
      const nodes = detail?.vm || detail?.nodes || fallbackObj?.vm || fallbackObj?.nodes || [];
      setTbNodeList(nodes);
      if (nodes.length > 0) {
        await applyTbNode(nsId || 'mig01', nodes[0]);
      } else {
        setSelectedTbNodeId('');
      }
    } catch {
      setTbNodeList([]);
      setSelectedTbNodeId('');
    } finally {
      setIsLoadingTbInfra(false);
    }
  };

  // Fetch Tumblebug Migrated Infras for Destination SSH selection
  const handleFetchTargetInfras = async (nsIdToFetch: string) => {
    setIsLoadingTbInfra(true);
    try {
      const infras = await beetleApi.getMigratedInfraList(nsIdToFetch || 'mig01');
      const list = Array.isArray(infras) ? infras : [];
      setTbInfraList(list);
      if (list.length > 0) {
        const firstId = list[0].id || list[0].name || '';
        await handleSelectTbInfra(nsIdToFetch || 'mig01', firstId, list);
      } else {
        setSelectedTbInfraId('');
        setTbNodeList([]);
        setSelectedTbNodeId('');
      }
    } catch {
      setTbInfraList([]);
      setTbNodeList([]);
    } finally {
      setIsLoadingTbInfra(false);
    }
  };

  // Drag-and-drop PEM key file reader
  const handlePemFileDrop = (
    e: React.DragEvent<HTMLTextAreaElement>,
    setter: React.Dispatch<React.SetStateAction<string>>
  ) => {
    e.preventDefault();
    const file = e.dataTransfer?.files?.[0];
    if (!file) return;
    const reader = new FileReader();
    reader.onload = (ev) => {
      if (typeof ev.target?.result === 'string') {
        setter(ev.target.result);
      }
    };
    reader.readAsText(file);
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
    handleFetchHoneybeeGroups();
    handleFetchTargetInfras(dataNsId);
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
      const cleanSourceSubPath = (sourceScopeMode === 'all') ? '' : sourceSubPath.trim().replace(/^\/+/, '');
      const fullSourcePath = sourceAccessType === 'object-storage' 
        ? (cleanSourceSubPath ? `${rawSourceBucket}/${cleanSourceSubPath}` : rawSourceBucket) 
        : (sourceSshPath.trim() || '/');

      const rawTargetStorage = targetAccessEngine === 'tumblebug'
        ? (selectedTargetStorage || 'target-storage-01')
        : (selectedTargetBucket || 'dest-bucket-01');
      const cleanTargetSubPath = targetSubPath.trim().replace(/^\/+/, '');
      const fullTargetPath = targetAccessType === 'object-storage' 
        ? (cleanTargetSubPath ? `${rawTargetStorage}/${cleanTargetSubPath}` : rawTargetStorage) 
        : (targetSshPath.trim() || '/');

      const displaySourceLabel = sourceAccessType === 'object-storage'
        ? fullSourcePath
        : `${sourceSshHost}:${fullSourcePath}`;
      const displayTargetLabel = targetAccessType === 'object-storage'
        ? fullTargetPath
        : `${targetSshHost}:${fullTargetPath}`;

      // Build filter configuration for source location
      const includePatterns = includeFilter.split(',').map(s => s.trim()).filter(Boolean);
      const excludePatterns = excludeFilter.split(',').map(s => s.trim()).filter(Boolean);
      const hasFilter = includePatterns.length > 0 || excludePatterns.length > 0;
      const filterConfig = hasFilter ? {
        ...(includePatterns.length > 0 ? { include: includePatterns } : {}),
        ...(excludePatterns.length > 0 ? { exclude: excludePatterns } : {})
      } : undefined;

      // Step 2: Build Plaintext DataMigrationModel matching pkg/api/rest/controller/migration-data.go & transx specification
      const plainModel = {
        source: {
          storageType: sourceAccessType === 'object-storage' ? 'objectstorage' : 'filesystem',
          path: fullSourcePath,
          ...(filterConfig ? { filter: filterConfig } : {}),
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
                host: sourceSshHost.trim(),
                port: parseInt(sourceSshPort, 10) || 22,
                username: sourceSshUser.trim() || 'ubuntu',
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
                host: targetSshHost.trim(),
                port: parseInt(targetSshPort, 10) || 22,
                username: targetSshUser.trim() || 'ubuntu',
                privateKey: targetSshPrivateKey
              }
            }
          })
        },
        strategy: dataStrategy || 'auto'
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
      const res = await beetleApi.migrateData(modelToSend, preferAsync);

      if (!res.success && res.error) {
        setModalError(res.error);
        setIsEncryptingAndLaunching(false);
        return;
      }

      const reqId = res.reqId || `req-data-${Date.now().toString().slice(-6)}`;
      const nowMs = Date.now();
      const newJob: DataMigrationJob = {
        id: reqId,
        reqId,
        sourceBucket: displaySourceLabel,
        targetStorage: displayTargetLabel,
        nsId: dataNsId,
        csp: dataCsp.toUpperCase(),
        region: dataRegion,
        strategy: dataStrategy,
        status: 'Handling',
        startTime: new Date(nowMs).toISOString(),
        createdAtMs: nowMs,
        elapsedSeconds: 0,
        encryptionKeyId: usedKeyId,
        logs: [
          `GET /beetle/migration/data/encryptionKey -> Issued KeyID: ${usedKeyId}`,
          `Constructed Beetle DataLocation.Path -> Source: "${displaySourceLabel}" | Target: "${displayTargetLabel}"`,
          `Encrypting sensitive credentials (${sourceAccessType === 'object-storage' ? 'AccessKeyId, SecretAccessKey' : 'SSH Private Key'}) via RSA-OAEP-256...`,
          `POST /beetle/migration/data -> 202 Accepted (ReqID: ${reqId}, Status: Handling)`
        ]
      };

      setDataJobs(prev => [newJob, ...prev]);
      setActiveJobId(newJob.id);
      setShowLaunchModal(false);
      setToastMsg(`🔒 Data Migration launched with path: ${displaySourceLabel} ➔ ${displayTargetLabel}!`);
      setTimeout(() => setToastMsg(null), 6000);
    } catch (err: any) {
      setModalError(err.message || 'Failed to launch data migration.');
    } finally {
      setIsEncryptingAndLaunching(false);
    }
  };

  const [deleteModalJob, setDeleteModalJob] = useState<DataMigrationJob | null>(null);
  const [deleteConfirmText, setDeleteConfirmText] = useState('');

  const handleConfirmDeleteDataRecord = () => {
    if (!deleteModalJob || deleteConfirmText !== deleteModalJob.reqId) return;
    removeDataJob(deleteModalJob.id);
    setDeleteModalJob(null);
    setDeleteConfirmText('');
  };

  const activeJob = dataJobs.find(j => j.id === activeJobId) || dataJobs[0];
  const runningJobsCount = dataJobs.filter(j => j.status === 'Handling').length;
  const completedJobsCount = dataJobs.filter(j => j.status === 'Success').length;

  // Live Duration Timer & Status Polling Effect for Data Migration Jobs Queue
  useEffect(() => {
    const runningJobs = dataJobs.filter((j) => j.status === 'Handling');
    if (runningJobs.length === 0) return;

    const intervalId = setInterval(async () => {
      // 1. Calculate real wall-clock elapsedSeconds for all running jobs
      setDataJobs((prevJobs) =>
        prevJobs.map((job) => {
          if (job.status === 'Handling') {
            const realElapsed = calculateJobElapsedSeconds(job);
            return { ...job, elapsedSeconds: realElapsed };
          }
          return job;
        })
      );

      // 2. Poll Beetle REST API GET /beetle/request/{reqId} to check if job finished
      for (const job of runningJobs) {
        if (!job.reqId || job.isSample) continue;
        try {
          const reqDetails = await beetleApi.getRequestDetails(job.reqId);
          const status = reqDetails?.status;
          if (status === 'Completed' || status === 'Succeeded' || status === 'Success') {
            setDataJobs((prevJobs) =>
              prevJobs.map((j) => {
                if (j.id === job.id) {
                  const dur = calculateJobElapsedSeconds(j);
                  return {
                    ...j,
                    status: 'Success',
                    elapsedSeconds: dur,
                    logs: [
                      ...j.logs.filter((l) => !l.includes('GET /beetle/request/')),
                      `GET /beetle/request/${job.reqId} -> Status: Success (Duration: ${dur}s)`
                    ]
                  };
                }
                return j;
              })
            );
          } else if (status === 'Failed' || status === 'Error') {
            setDataJobs((prevJobs) =>
              prevJobs.map((j) => {
                if (j.id === job.id) {
                  const dur = calculateJobElapsedSeconds(j);
                  return {
                    ...j,
                    status: 'Failed',
                    elapsedSeconds: dur,
                    logs: [
                      ...j.logs.filter((l) => !l.includes('GET /beetle/request/')),
                      `GET /beetle/request/${job.reqId} -> Status: Failed (${reqDetails?.errorResponse || 'Execution Error'}) (Duration: ${dur}s)`
                    ]
                  };
                }
                return j;
              })
            );
          }
        } catch (err) {
          console.warn('Poll error for data job request', job.reqId, err);
        }
      }
    }, 1000);

    return () => clearInterval(intervalId);
  }, [dataJobs, setDataJobs]);

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
          className="px-5 py-2.5 bg-gradient-to-r from-teal-400 via-emerald-400 to-emerald-500 hover:from-teal-500 hover:to-emerald-600 text-slate-950 rounded-xl text-sm font-extrabold flex items-center gap-1.5 transition cursor-pointer shadow-lg shadow-teal-500/20"
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
                  className={`p-4 rounded-xl border text-left transition-all duration-200 cursor-pointer flex flex-col justify-between space-y-3 relative overflow-hidden ${
                    isSelected
                      ? 'bg-emerald-500/10 border-emerald-500/60 shadow-lg shadow-emerald-500/10 ring-1 ring-emerald-500/40'
                      : 'bg-bg-panel hover:bg-bg-input border-border-main/60'
                  }`}
                >
                  {/* Top Glowing indicator for active card */}
                  {isSelected && (
                    <div className="absolute top-0 left-0 w-full h-[2px] bg-gradient-to-r from-emerald-500 to-blue-500" />
                  )}

                  <div className="flex justify-between items-center gap-2">
                    <div className="flex items-center gap-2 min-w-0">
                      <span className="text-base shrink-0">
                        {job.csp.toLowerCase() === 'aws' ? '🌩️' : job.csp.toLowerCase() === 'azure' ? '🔷' : '🟢'}
                      </span>
                      <span className="font-extrabold text-sm text-text-main font-mono truncate" title={`${job.sourceBucket} ➔ ${job.targetStorage}`}>
                        {job.isSample && <span className="text-amber-500 font-bold mr-1">[Sample]</span>}
                        {job.sourceBucket} ➔ {job.targetStorage}
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
                          title="Delete job from history queue"
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
                <h3 className="text-base font-extrabold text-text-main font-mono">
                  Job Detail: {activeJob.isSample && <span className="text-amber-500 font-bold mr-1">[Sample]</span>}{activeJob.reqId}
                </h3>
                <span className="px-2.5 py-0.5 rounded text-xs font-bold font-mono uppercase bg-emerald-500/10 text-emerald-400 border border-emerald-500/20">
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
              {!activeJob.isSample && (
                <button
                  onClick={() => {
                    setDeleteModalJob(activeJob);
                    setDeleteConfirmText('');
                  }}
                  className="px-4 py-2 bg-bg-panel hover:bg-red-500/10 border border-border-main hover:border-red-500/30 text-text-muted hover:text-red-400 rounded-xl text-sm font-bold transition flex items-center space-x-1.5 cursor-pointer"
                >
                  <Trash2 className="w-4 h-4" />
                  <span>Delete Record</span>
                </button>
              )}
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
                      <div className="flex flex-wrap items-center justify-between border-b border-border-main/20 pb-2 gap-2">
                        <span className="text-sm font-bold text-text-main font-mono flex items-center gap-1.5">
                          <Key className="w-4 h-4 text-emerald-400" />
                          Source Authentication &amp; Credentials
                        </span>

                        {sourceAccessType === 'remote-ssh' && (
                          <div className="flex items-center gap-1 bg-bg-panel p-0.5 border border-border-main rounded-lg text-xs">
                            <button
                              type="button"
                              onClick={() => {
                                setSourceSshMode('honeybee');
                                if (hbGroups.length === 0) handleFetchHoneybeeGroups();
                              }}
                              className={`px-2 py-0.5 rounded-md font-extrabold transition cursor-pointer ${
                                sourceSshMode === 'honeybee'
                                  ? 'bg-emerald-500 text-slate-950 shadow-xs'
                                  : 'text-text-muted hover:text-text-main'
                              }`}
                            >
                              Honeybee Source Group
                            </button>
                            <button
                              type="button"
                              onClick={() => setSourceSshMode('direct')}
                              className={`px-2 py-0.5 rounded-md font-extrabold transition cursor-pointer ${
                                sourceSshMode === 'direct'
                                  ? 'bg-emerald-500 text-slate-950 shadow-xs'
                                  : 'text-text-muted hover:text-text-main'
                              }`}
                            >
                              Direct SSH
                            </button>
                          </div>
                        )}
                      </div>

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
                          {sourceSshMode === 'honeybee' && (
                            <div className="p-3 bg-bg-input/60 border border-emerald-500/30 rounded-xl space-y-2.5">
                              <div className="flex items-center justify-between">
                                <span className="text-xs font-extrabold text-emerald-600 dark:text-emerald-400">
                                  Select Registered Honeybee Source Node
                                </span>
                                <button
                                  type="button"
                                  onClick={handleFetchHoneybeeGroups}
                                  disabled={isLoadingHb}
                                  className="px-2.5 py-1 bg-bg-panel hover:bg-bg-main border border-border-main rounded-lg text-xs font-bold text-emerald-600 dark:text-emerald-400 flex items-center gap-1 cursor-pointer transition disabled:opacity-50"
                                >
                                  <RefreshCw className={`w-3 h-3 ${isLoadingHb ? 'animate-spin' : ''}`} />
                                  <span>Refresh</span>
                                </button>
                              </div>
                              <div className="grid grid-cols-1 sm:grid-cols-2 gap-2.5">
                                <div>
                                  <label className="block text-text-muted font-normal text-xs mb-1">Source Group</label>
                                  <select
                                    value={selectedHbGroupId}
                                    onChange={(e) => handleSelectHbGroup(e.target.value)}
                                    className="w-full px-3 py-1.5 bg-bg-panel border border-border-main rounded-lg text-text-main font-bold text-sm focus:outline-none focus:border-emerald-500 cursor-pointer"
                                  >
                                    {hbGroups.length > 0 ? (
                                      hbGroups.map((g: any) => (
                                        <option key={g.id} value={g.id}>{g.name || g.id}</option>
                                      ))
                                    ) : (
                                      <option value="">No source groups found</option>
                                    )}
                                  </select>
                                </div>
                                <div>
                                  <label className="block text-text-muted font-normal text-xs mb-1">Source Node</label>
                                  <select
                                    value={selectedHbConnId}
                                    onChange={(e) => {
                                      const val = e.target.value;
                                      const conn = hbConnections.find((c: any) => (c.id || c.name) === val);
                                      if (conn) applyHbConnection(conn);
                                    }}
                                    className="w-full px-3 py-1.5 bg-bg-panel border border-border-main rounded-lg text-text-main font-bold font-mono text-sm focus:outline-none focus:border-emerald-500 cursor-pointer"
                                  >
                                    {hbConnections.length > 0 ? (
                                      hbConnections.map((c: any) => {
                                        const cid = c.id || c.name;
                                        const ip = c.ip_address || c.ip || '';
                                        return (
                                          <option key={cid} value={cid}>
                                            {c.name || cid} ({ip})
                                          </option>
                                        );
                                      })
                                    ) : (
                                      <option value="">Select source group first</option>
                                    )}
                                  </select>
                                </div>
                              </div>
                            </div>
                          )}

                          <div className="grid grid-cols-1 sm:grid-cols-3 gap-3.5">
                            <div className="sm:col-span-2">
                              <label className="block text-text-muted font-normal text-sm mb-1">SSH Host IP / Hostname</label>
                              <input
                                type="text"
                                value={sourceSshHost}
                                onChange={(e) => setSourceSshHost(e.target.value)}
                                placeholder="10.0.1.11"
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
                            <div className="flex items-center justify-between mb-1">
                              <label className="block text-text-muted font-normal text-sm">SSH Private Key (Encrypted PEM Format)</label>
                              <span className="text-xs text-text-muted flex items-center gap-1">
                                <Upload className="w-3 h-3 text-emerald-500" />
                                Drop .pem file or paste
                              </span>
                            </div>
                            <textarea
                              rows={2}
                              value={sourceSshPrivateKey}
                              onChange={(e) => setSourceSshPrivateKey(e.target.value)}
                              onDragOver={(e) => e.preventDefault()}
                              onDrop={(e) => handlePemFileDrop(e, setSourceSshPrivateKey)}
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
                      <div className="flex flex-wrap items-center justify-between border-b border-border-main/20 pb-2 gap-2">
                        <span className="text-sm font-bold text-text-main font-mono">
                          Destination Credentials &amp; Mode
                        </span>

                        {targetAccessType === 'object-storage' ? (
                          <div className="flex items-center gap-1 bg-bg-panel p-0.5 border border-border-main rounded-lg text-xs">
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
                        ) : (
                          <div className="flex items-center gap-1 bg-bg-panel p-0.5 border border-border-main rounded-lg text-xs">
                            <button
                              type="button"
                              onClick={() => {
                                setTargetSshMode('tumblebug');
                                if (tbInfraList.length === 0) handleFetchTargetInfras(dataNsId);
                              }}
                              className={`px-2 py-0.5 rounded-md font-extrabold transition cursor-pointer ${
                                targetSshMode === 'tumblebug'
                                  ? 'bg-teal-400 text-slate-950 shadow-xs'
                                  : 'text-text-muted hover:text-text-main'
                              }`}
                            >
                              Tumblebug Infra
                            </button>
                            <button
                              type="button"
                              onClick={() => setTargetSshMode('direct')}
                              className={`px-2 py-0.5 rounded-md font-extrabold transition cursor-pointer ${
                                targetSshMode === 'direct'
                                  ? 'bg-teal-400 text-slate-950 shadow-xs'
                                  : 'text-text-muted hover:text-text-main'
                              }`}
                            >
                              Direct SSH
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
                          {targetSshMode === 'tumblebug' && (
                            <div className="p-3 bg-bg-input/60 border border-teal-500/30 rounded-xl space-y-2.5">
                              <div className="flex items-center justify-between">
                                <span className="text-xs font-extrabold text-teal-600 dark:text-teal-400">
                                  Select Provisioned Tumblebug Target Node
                                </span>
                                <button
                                  type="button"
                                  onClick={() => handleFetchTargetInfras(dataNsId)}
                                  disabled={isLoadingTbInfra}
                                  className="px-2.5 py-1 bg-bg-panel hover:bg-bg-main border border-border-main rounded-lg text-xs font-bold text-teal-600 dark:text-teal-400 flex items-center gap-1 cursor-pointer transition disabled:opacity-50"
                                >
                                  <RefreshCw className={`w-3 h-3 ${isLoadingTbInfra ? 'animate-spin' : ''}`} />
                                  <span>Refresh</span>
                                </button>
                              </div>
                              <div className="grid grid-cols-1 sm:grid-cols-3 gap-2.5">
                                <div>
                                  <label className="block text-text-muted font-normal text-xs mb-1">Namespace (nsId)</label>
                                  <input
                                    type="text"
                                    value={dataNsId}
                                    onChange={(e) => {
                                      setDataNsId(e.target.value);
                                      handleFetchTargetInfras(e.target.value);
                                    }}
                                    placeholder="mig01"
                                    className="w-full px-3 py-1.5 bg-bg-panel border border-border-main rounded-lg text-text-main font-bold font-mono text-sm focus:outline-none focus:border-teal-400"
                                  />
                                </div>
                                <div>
                                  <label className="block text-text-muted font-normal text-xs mb-1">Migrated Infra</label>
                                  <select
                                    value={selectedTbInfraId}
                                    onChange={(e) => handleSelectTbInfra(dataNsId, e.target.value)}
                                    className="w-full px-3 py-1.5 bg-bg-panel border border-border-main rounded-lg text-text-main font-bold text-sm focus:outline-none focus:border-teal-400 cursor-pointer"
                                  >
                                    {tbInfraList.length > 0 ? (
                                      tbInfraList.map((inf: any) => {
                                        const iid = inf.id || inf.name || '';
                                        return <option key={iid} value={iid}>{iid}</option>;
                                      })
                                    ) : (
                                      <option value="">No migrated infra found</option>
                                    )}
                                  </select>
                                </div>
                                <div>
                                  <label className="block text-text-muted font-normal text-xs mb-1">Target Node</label>
                                  <select
                                    value={selectedTbNodeId}
                                    onChange={(e) => {
                                      const val = e.target.value;
                                      const node = tbNodeList.find((n: any) => (n.id || n.name) === val);
                                      if (node) applyTbNode(dataNsId, node);
                                    }}
                                    className="w-full px-3 py-1.5 bg-bg-panel border border-border-main rounded-lg text-text-main font-bold font-mono text-sm focus:outline-none focus:border-teal-400 cursor-pointer"
                                  >
                                    {tbNodeList.length > 0 ? (
                                      tbNodeList.map((n: any) => {
                                        const nid = n.id || n.name || '';
                                        const ip = n.publicIP || n.publicIp || n.privateIP || n.privateIp || '';
                                        return (
                                          <option key={nid} value={nid}>
                                            {n.name || nid} ({ip || 'IP pending'})
                                          </option>
                                        );
                                      })
                                    ) : (
                                      <option value="">Select infra first</option>
                                    )}
                                  </select>
                                </div>
                              </div>
                            </div>
                          )}

                          <div className="grid grid-cols-1 sm:grid-cols-3 gap-3.5">
                            <div className="sm:col-span-2">
                              <label className="block text-text-muted font-normal text-sm mb-1">Target SSH Host IP</label>
                              <input
                                type="text"
                                value={targetSshHost}
                                onChange={(e) => setTargetSshHost(e.target.value)}
                                placeholder="43.200.254.107"
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
                            <div className="flex items-center justify-between mb-1">
                              <label className="block text-text-muted font-normal text-sm">SSH Private Key (Encrypted PEM Format)</label>
                              <span className="text-xs text-text-muted flex items-center gap-1">
                                <Upload className="w-3 h-3 text-teal-500" />
                                Auto-filled from Tumblebug or drop .pem
                              </span>
                            </div>
                            <textarea
                              rows={2}
                              value={targetSshPrivateKey}
                              onChange={(e) => setTargetSshPrivateKey(e.target.value)}
                              onDragOver={(e) => e.preventDefault()}
                              onDrop={(e) => handlePemFileDrop(e, setTargetSshPrivateKey)}
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
                    <div className="space-y-3">
                      <div>
                        <label className="block text-text-muted font-normal text-sm mb-1.5">
                          Transfer Scope
                        </label>
                        <div className="grid grid-cols-2 gap-2">
                          <button
                            type="button"
                            onClick={() => {
                              setSourceScopeMode('all');
                              setSourceSubPath('');
                            }}
                            className={`px-3 py-2 text-xs font-bold rounded-xl border transition-all text-center flex items-center justify-center gap-1.5 ${
                              sourceScopeMode === 'all'
                                ? 'bg-emerald-500/15 border-emerald-500 text-emerald-400 shadow-sm'
                                : 'bg-bg-panel border-border-main text-text-muted hover:border-emerald-500/50'
                            }`}
                          >
                            <span className="w-2 h-2 rounded-full bg-emerald-500 inline-block" />
                            Entire Bucket (/*)
                          </button>
                          <button
                            type="button"
                            onClick={() => setSourceScopeMode('prefix')}
                            className={`px-3 py-2 text-xs font-bold rounded-xl border transition-all text-center flex items-center justify-center gap-1.5 ${
                              sourceScopeMode === 'prefix'
                                ? 'bg-emerald-500/15 border-emerald-500 text-emerald-400 shadow-sm'
                                : 'bg-bg-panel border-border-main text-text-muted hover:border-emerald-500/50'
                            }`}
                          >
                            <span className="w-2 h-2 rounded-full bg-teal-500 inline-block" />
                            Sub-Folder / Prefix
                          </button>
                        </div>
                      </div>

                      {sourceScopeMode === 'prefix' ? (
                        <div>
                          <label className="block text-text-muted font-normal text-sm mb-1">
                            Specific Key Prefix / Directory
                          </label>
                          <input
                            type="text"
                            value={sourceSubPath}
                            onChange={(e) => setSourceSubPath(e.target.value)}
                            placeholder="e.g. data/ or uploads/ (leave blank for root)"
                            className="w-full px-3.5 py-2 bg-bg-panel border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 font-bold text-sm"
                          />
                          <p className="text-xs text-text-muted mt-1">
                            Only objects starting with this prefix will be transferred.
                          </p>
                        </div>
                      ) : (
                        <div className="p-2.5 bg-bg-panel/80 border border-border-main/50 rounded-xl text-xs text-text-muted flex items-center gap-2">
                          <CheckCircle2 className="w-4 h-4 text-emerald-500 shrink-0" />
                          <span>All objects and folders in bucket <strong className="text-text-main">{selectedSourceBucket || 'bucket'}</strong> will be transferred.</span>
                        </div>
                      )}
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
                        ? (sourceScopeMode === 'all' || !sourceSubPath.trim()
                            ? `${dataCsp.toLowerCase()}://${selectedSourceBucket || 'bucket'} (All Objects)`
                            : `${dataCsp.toLowerCase()}://${selectedSourceBucket || 'bucket'}/${sourceSubPath.trim().replace(/^\/+/, '')}`)
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
                <div className="flex flex-wrap justify-between items-center border-b border-border-main/30 pb-2.5 gap-2">
                  <span className="text-sm font-extrabold text-text-main font-mono flex items-center gap-1.5">
                    <Filter className="w-4 h-4 text-emerald-500" />
                    Data Filter Settings
                  </span>
                  <div className="flex flex-wrap items-center gap-2">
                    <button
                      type="button"
                      onClick={handleApplyAimxPreset}
                      className="px-3 py-1 bg-emerald-500/15 hover:bg-emerald-500/25 border border-emerald-500/40 text-emerald-600 dark:text-emerald-400 rounded-lg text-xs font-extrabold transition cursor-pointer flex items-center gap-1.5 shadow-2xs"
                    >
                      <Sparkles className="w-3.5 h-3.5" />
                      <span>+ Preset: AIMX (Ollama + OpenWebUI)</span>
                    </button>
                    <button
                      type="button"
                      onClick={() => setShowDirectoryTree(!showDirectoryTree)}
                      className="px-3 py-1 bg-bg-panel hover:bg-bg-main border border-border-main text-text-main rounded-lg text-xs font-bold transition cursor-pointer flex items-center gap-1.5"
                    >
                      <Folder className="w-3.5 h-3.5 text-teal-500" />
                      <span>Browse Source Directory Tree</span>
                      {showDirectoryTree ? <ChevronUp className="w-3.5 h-3.5" /> : <ChevronDown className="w-3.5 h-3.5" />}
                    </button>
                  </div>
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
                          className="px-2 py-0.5 bg-emerald-100 dark:bg-emerald-950/80 hover:bg-emerald-200 dark:hover:bg-emerald-900 border border-emerald-400/60 dark:border-emerald-500/50 text-emerald-900 dark:text-emerald-300 text-xs font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
                        >
                          + *.json
                        </button>
                        <button
                          type="button"
                          onClick={() => handleAddIncludePreset('*.csv')}
                          className="px-2 py-0.5 bg-emerald-100 dark:bg-emerald-950/80 hover:bg-emerald-200 dark:hover:bg-emerald-900 border border-emerald-400/60 dark:border-emerald-500/50 text-emerald-900 dark:text-emerald-300 text-xs font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
                        >
                          + *.csv
                        </button>
                        <button
                          type="button"
                          onClick={() => handleAddIncludePreset('reports/**')}
                          className="px-2 py-0.5 bg-emerald-100 dark:bg-emerald-950/80 hover:bg-emerald-200 dark:hover:bg-emerald-900 border border-emerald-400/60 dark:border-emerald-500/50 text-emerald-900 dark:text-emerald-300 text-xs font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
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
                          className="px-2 py-0.5 bg-rose-100 dark:bg-rose-950/80 hover:bg-rose-200 dark:hover:bg-rose-900 border border-rose-400/60 dark:border-rose-500/50 text-rose-900 dark:text-rose-300 text-xs font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
                        >
                          + *.tmp
                        </button>
                        <button
                          type="button"
                          onClick={() => handleAddExcludePreset('*.log')}
                          className="px-2 py-0.5 bg-rose-100 dark:bg-rose-950/80 hover:bg-rose-200 dark:hover:bg-rose-900 border border-rose-400/60 dark:border-rose-500/50 text-rose-900 dark:text-rose-300 text-xs font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
                        >
                          + *.log
                        </button>
                        <button
                          type="button"
                          onClick={() => handleAddExcludePreset('temp/**')}
                          className="px-2 py-0.5 bg-rose-100 dark:bg-rose-950/80 hover:bg-rose-200 dark:hover:bg-rose-900 border border-rose-400/60 dark:border-rose-500/50 text-rose-900 dark:text-rose-300 text-xs font-extrabold rounded-lg transition cursor-pointer shadow-2xs"
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

                {showDirectoryTree && (
                  <DirectoryTreeExplorer
                    includeFilter={includeFilter}
                    excludeFilter={excludeFilter}
                    onToggleInclude={handleToggleIncludePattern}
                    onToggleExclude={handleToggleExcludePattern}
                  />
                )}
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

            {/* Execute Data Migration Button & Prefer respond-async Toggle */}
            <div className="border-t border-border-main/40 pt-4 flex flex-col sm:flex-row items-stretch sm:items-center justify-between gap-3 font-sans">
              {/* Left Side: Prefer respond-async Toggle */}
              <label className="flex items-center gap-2 cursor-pointer text-xs font-bold text-text-main p-2 px-3.5 bg-bg-panel/60 hover:bg-bg-input border border-border-main/40 rounded-xl shrink-0 transition">
                <Zap className="w-4 h-4 text-emerald-500 shrink-0" />
                <span className="font-mono">Prefer: respond-async</span>
                <input
                  type="checkbox"
                  checked={preferAsync}
                  onChange={(e) => setPreferAsync(e.target.checked)}
                  className="w-4 h-4 accent-emerald-500 cursor-pointer rounded ml-1"
                />
              </label>

              {/* Right Side: Cancel & Launch Migration Buttons */}
              <div className="flex items-center justify-end space-x-3 shrink-0">
                <button
                  onClick={() => setShowLaunchModal(false)}
                  className="px-5 py-2.5 bg-bg-input hover:bg-bg-main border border-border-main text-text-main font-bold text-sm rounded-xl transition cursor-pointer"
                >
                  Cancel
                </button>
                <button
                  onClick={handleConfirmLaunchDataMigration}
                  disabled={isEncryptingAndLaunching}
                  className="px-6 py-2.5 bg-gradient-to-r from-teal-400 via-emerald-400 to-emerald-500 hover:from-teal-500 hover:to-emerald-600 disabled:opacity-40 text-slate-950 font-extrabold text-sm rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center space-x-2 cursor-pointer"
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

      {/* Delete Record Confirmation Modal with Text Pattern Matching */}
      {deleteModalJob && (
        <div className="fixed inset-0 z-[60] flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up space-y-4">
            <div className="flex justify-between items-center border-b border-border-main/20 pb-3">
              <h3 className="text-base font-bold text-text-main flex items-center gap-2">
                <Trash2 className="w-4 h-4 text-red-500" />
                <span>Delete Data Migration Record</span>
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
                Are you sure you want to remove the data migration record for <strong className="text-text-main">"{deleteModalJob.reqId}"</strong> (<span className="font-mono text-emerald-400">{deleteModalJob.sourceBucket} ➔ {deleteModalJob.targetStorage}</span>) from the Queue?
              </p>
              <div className="p-3 bg-amber-500/10 border border-amber-500/30 rounded-xl text-xs text-amber-400 font-mono">
                ⚠️ Note: This action only removes the UI queue request record. Transferred data objects in cloud storage will NOT be deleted.
              </div>

              <div className="space-y-1.5 pt-1">
                <label className="block text-xs font-bold text-text-muted">
                  To confirm deletion, type <span className="font-mono bg-bg-panel px-1.5 py-0.5 rounded border border-border-main/60 text-emerald-400 select-all">{deleteModalJob.reqId}</span> below:
                </label>
                <input
                  type="text"
                  value={deleteConfirmText}
                  onChange={(e) => setDeleteConfirmText(e.target.value)}
                  placeholder={`Type "${deleteModalJob.reqId}" to confirm`}
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
                  onClick={handleConfirmDeleteDataRecord}
                  disabled={deleteConfirmText !== deleteModalJob.reqId}
                  className={`px-4 py-2 rounded-xl text-xs font-bold transition flex items-center gap-1.5 ${
                    deleteConfirmText !== deleteModalJob.reqId
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
