'use client';

import React, { useState } from 'react';
import { useMigrationStore } from '@/store/migrationStore';
import { CspCredentialForm } from '../common/CspCredentialForm';
import { Key, Plus, Trash2, RefreshCw, CheckCircle2, Lock, Eye, EyeOff, AlertCircle, X } from 'lucide-react';

interface CredentialProfile {
  id: string;
  profileName: string;
  csp: string;
  region: string;
  accessKeyMasked: string;
  status: 'valid' | 'invalid' | 'unverified';
  createdAt: string;
}

const CSP_REGIONS_MAP: Record<string, { id: string; name: string }[]> = {
  aws: [
    { id: 'ap-northeast-2', name: 'ap-northeast-2 (Seoul)' },
    { id: 'ap-northeast-1', name: 'ap-northeast-1 (Tokyo)' },
    { id: 'us-east-1', name: 'us-east-1 (N. Virginia)' },
    { id: 'us-west-2', name: 'us-west-2 (Oregon)' },
    { id: 'eu-central-1', name: 'eu-central-1 (Frankfurt)' },
    { id: 'ap-southeast-1', name: 'ap-southeast-1 (Singapore)' }
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
  ]
};

export const CredentialManagement: React.FC = () => {
  const { tumblebugProviders, tumblebugRegions, fetchTumblebugRegions } = useMigrationStore();

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [showSecret, setShowSecret] = useState(false);
  const [isVerifying, setIsVerifying] = useState<string | null>(null);

  // Form State
  const [profileName, setProfileName] = useState('');
  const [csp, setCsp] = useState('aws');
  const [region, setRegion] = useState('ap-northeast-2');
  const [accessKey, setAccessKey] = useState('');
  const [secretKey, setSecretKey] = useState('');
  const [tenantId, setTenantId] = useState('');
  const [subscriptionId, setSubscriptionId] = useState('');
  const [s3AccessKey, setS3AccessKey] = useState('');
  const [s3SecretKey, setS3SecretKey] = useState('');

  // Sample Saved Credentials
  const [credentials, setCredentials] = useState<CredentialProfile[]>([
    {
      id: 'cred-001',
      profileName: 'aws-production-account',
      csp: 'aws',
      region: 'ap-northeast-2',
      accessKeyMasked: 'AKIAIOSFODNN7*******',
      status: 'valid',
      createdAt: '2026-07-22 10:30'
    },
    {
      id: 'cred-002',
      profileName: 'gcp-service-account-prod',
      csp: 'gcp',
      region: 'asia-northeast3',
      accessKeyMasked: 'gcp-sa-prod@cloud-barista.iam.gserviceaccount.com',
      status: 'valid',
      createdAt: '2026-07-21 14:15'
    },
    {
      id: 'cred-003',
      profileName: 'ncloud-kr-account',
      csp: 'ncp',
      region: 'kr-1',
      accessKeyMasked: 'ncloud_access_key_*******',
      status: 'unverified',
      createdAt: '2026-07-20 09:00'
    }
  ]);

  const getRegionsForCsp = (cspKey: string) => {
    const key = cspKey.toLowerCase();
    if (CSP_REGIONS_MAP[key]) return CSP_REGIONS_MAP[key];
    if (tumblebugRegions && tumblebugRegions.length > 0) {
      return tumblebugRegions.map((r: any) => ({ id: r.id, name: `${r.id} (${r.locationName || r.name || r.id})` }));
    }
    return [
      { id: 'ap-northeast-2', name: 'ap-northeast-2 (Seoul)' },
      { id: 'us-east-1', name: 'us-east-1 (N. Virginia)' }
    ];
  };

  const handleCspChange = (newCsp: string) => {
    setCsp(newCsp);
    fetchTumblebugRegions(newCsp);
    const available = getRegionsForCsp(newCsp);
    if (available.length > 0) {
      setRegion(available[0].id);
    }
  };

  const handleVerifyCredential = async (id: string) => {
    setIsVerifying(id);
    try {
      await new Promise((resolve) => setTimeout(resolve, 1000));
      setCredentials((prev) =>
        prev.map((c) => (c.id === id ? { ...c, status: 'valid' } : c))
      );
    } finally {
      setIsVerifying(null);
    }
  };

  const handleDeleteCredential = (id: string) => {
    if (!confirm('Are you sure you want to delete this credential profile?')) return;
    setCredentials((prev) => prev.filter((c) => c.id !== id));
  };

  const openCreateModal = () => {
    setProfileName('');
    setCsp('aws');
    setRegion('ap-northeast-2');
    setAccessKey('');
    setSecretKey('');
    setTenantId('');
    setSubscriptionId('');
    setIsModalOpen(true);
  };

  const handleSaveCredential = (e: React.FormEvent) => {
    e.preventDefault();
    if (!profileName.trim() || !accessKey.trim()) {
      alert('Please fill in profile name and access key.');
      return;
    }
    const newCred: CredentialProfile = {
      id: `cred-${Date.now().toString().slice(-4)}`,
      profileName: profileName.trim(),
      csp,
      region,
      accessKeyMasked: accessKey.length > 8 ? `${accessKey.slice(0, 8)}*******` : accessKey,
      status: 'valid',
      createdAt: new Date().toISOString().replace('T', ' ').slice(0, 16)
    };
    setCredentials((prev) => [newCred, ...prev]);
    setIsModalOpen(false);
  };

  return (
    <div className="space-y-6 animate-fade-in">
      {/* Tab Description Header Box (Rule-compliant) */}
      <div className="bg-bg-panel border border-border-main rounded-2xl p-4 shadow-sm space-y-3">
        <div className="flex flex-wrap items-center gap-x-3 gap-y-1.5 px-2 py-1">
          <Key className="w-5 h-5 text-emerald-500 shrink-0" />
          <h2 className="text-base font-extrabold text-text-main flex items-center space-x-2">
            <span>Cloud Credential Management</span>
            <span className="px-2 py-0.5 text-xs font-mono font-extrabold bg-amber-500/10 text-amber-600 dark:text-amber-400 border border-amber-500/20 rounded-md">
              WIP
            </span>
          </h2>
          <span className="text-sm text-text-muted">
            Register and manage CSP Access Credentials and Vault secret keys for Multi-Cloud discovery &amp; migration
          </span>
        </div>
      </div>

      {/* Main Panel */}
      <div className="bg-bg-panel border border-border-main rounded-2xl p-6 space-y-6 shadow-sm">
        {/* Navigation / Actions Bar */}
        <div className="flex items-center justify-between border-b border-border-main pb-4">
          <div>
            <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
              <Lock className="w-5 h-5 text-emerald-400" />
              Registered Credentials ({credentials.length})
            </h3>
            <p className="text-sm text-text-muted mt-0.5">
              Select or manage credentials for infrastructure &amp; object storage discovery.
            </p>
          </div>
          <button
            onClick={openCreateModal}
            className="px-4 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white rounded-xl text-sm font-bold flex items-center gap-2 transition cursor-pointer shadow-md shadow-emerald-500/20"
          >
            <Plus className="w-4 h-4" />
            <span>Register New Credential</span>
          </button>
        </div>

        {/* Credentials Table View */}
        <div className="overflow-x-auto border border-border-main rounded-xl">
          <table className="w-full text-xs text-left">
            <thead className="bg-bg-main/50 text-text-muted font-bold border-b border-border-main">
              <tr>
                <th className="p-3.5">Profile Name</th>
                <th className="p-3.5">CSP</th>
                <th className="p-3.5">Region</th>
                <th className="p-3.5">Access Key ID / Identity</th>
                <th className="p-3.5">Status</th>
                <th className="p-3.5">Created At</th>
                <th className="p-3.5 text-right">Actions</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-border-main font-mono">
              {credentials.map((cred) => (
                <tr key={cred.id} className="hover:bg-bg-main/30 transition">
                  <td className="p-3.5 font-bold text-emerald-600 dark:text-emerald-400">{cred.profileName}</td>
                  <td className="p-3.5">
                    <span className="px-2 py-0.5 rounded text-[11px] font-bold uppercase bg-bg-input border border-border-main">
                      {cred.csp}
                    </span>
                  </td>
                  <td className="p-3.5">{cred.region}</td>
                  <td className="p-3.5 text-text-main">{cred.accessKeyMasked}</td>
                  <td className="p-3.5">
                    {cred.status === 'valid' ? (
                      <span className="px-2 py-0.5 rounded text-[10px] font-bold bg-emerald-500/10 text-emerald-500 border border-emerald-500/20 flex items-center space-x-1 w-max">
                        <CheckCircle2 className="w-3 h-3" />
                        <span>Valid</span>
                      </span>
                    ) : (
                      <span className="px-2 py-0.5 rounded text-[10px] font-bold bg-amber-500/10 text-amber-500 border border-amber-500/20 flex items-center space-x-1 w-max">
                        <AlertCircle className="w-3 h-3" />
                        <span>Unverified</span>
                      </span>
                    )}
                  </td>
                  <td className="p-3.5 text-text-muted">{cred.createdAt}</td>
                  <td className="p-3.5 text-right space-x-2">
                    <button
                      onClick={() => handleVerifyCredential(cred.id)}
                      disabled={isVerifying === cred.id}
                      className="px-2.5 py-1 text-xs bg-bg-input hover:bg-bg-main border border-border-main text-text-main rounded-md transition cursor-pointer"
                      title="Verify API connection"
                    >
                      {isVerifying === cred.id ? <RefreshCw className="w-3.5 h-3.5 animate-spin" /> : 'Verify'}
                    </button>
                    <button
                      onClick={() => handleDeleteCredential(cred.id)}
                      className="p-1 text-text-muted hover:text-red-500 transition cursor-pointer inline-block"
                      title="Delete profile"
                    >
                      <Trash2 className="w-3.5 h-3.5" />
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>

      {/* Modal Popup: Register New CSP Credential Profile (Structured 3-Row Layout matching New Storage Group Popup Header) */}
      {isModalOpen && (
        <div className="fixed inset-0 z-50 bg-black/75 backdrop-blur-sm flex items-center justify-center p-4">
          <div className="bg-bg-panel border border-border-main rounded-2xl max-w-6xl w-full p-6 shadow-2xl space-y-5 animate-fade-in text-xs max-h-[90vh] overflow-y-auto">
            {/* Modal Header */}
            <div className="flex items-center justify-between border-b border-border-main pb-3">
              <div className="flex items-center space-x-2">
                <Lock className="w-5 h-5 text-emerald-500" />
                <h3 className="text-base font-extrabold text-text-main">
                  Register New CSP Credential Profile
                </h3>
              </div>
              <button
                onClick={() => setIsModalOpen(false)}
                className="p-1 text-text-muted hover:text-text-main transition cursor-pointer"
              >
                <X className="w-5 h-5" />
              </button>
            </div>

            <form onSubmit={handleSaveCredential} className="space-y-4">
              {/* Row 1: Credential Profile Name */}
              <div>
                <label className="block text-text-main font-bold mb-1">
                  Credential Profile Name <span className="text-red-500">*</span>
                </label>
                <input
                  type="text"
                  required
                  placeholder="e.g. aws-prod-credential"
                  value={profileName}
                  onChange={(e) => setProfileName(e.target.value)}
                  className="w-full px-3.5 py-2.5 bg-bg-input border border-border-main rounded-xl text-text-main font-bold focus:outline-none focus:border-emerald-500"
                />
              </div>

              {/* Modular CSP Credential Form Component */}
              <CspCredentialForm
                csp={csp}
                onCspChange={handleCspChange}
                region={region}
                onRegionChange={setRegion}
                accessKey={accessKey}
                onAccessKeyChange={setAccessKey}
                secretKey={secretKey}
                onSecretKeyChange={setSecretKey}
                tenantId={tenantId}
                onTenantIdChange={setTenantId}
                subscriptionId={subscriptionId}
                onSubscriptionIdChange={setSubscriptionId}
                s3AccessKey={s3AccessKey}
                onS3AccessKeyChange={setS3AccessKey}
                s3SecretKey={s3SecretKey}
                onS3SecretKeyChange={setS3SecretKey}
              />

              {/* Modal Footer */}
              <div className="flex justify-end space-x-3 pt-3 border-t border-border-main">
                <button
                  type="button"
                  onClick={() => setIsModalOpen(false)}
                  className="px-5 py-2.5 bg-bg-input border border-border-main hover:bg-bg-main text-text-main rounded-xl font-bold transition cursor-pointer"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  className="px-6 py-2.5 bg-emerald-500 hover:bg-emerald-600 text-white font-extrabold text-xs rounded-xl transition cursor-pointer shadow-lg shadow-emerald-500/20"
                >
                  Save Credential Profile
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
};
