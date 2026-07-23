'use client';

import React, { useState, useEffect } from 'react';
import { tumblebugApi } from '@/api/client';
import { Eye, EyeOff, RefreshCw } from 'lucide-react';

export interface CspCredentialFormProps {
  csp: string;
  onCspChange: (csp: string) => void;
  region: string;
  onRegionChange: (region: string) => void;
  accessKey: string;
  onAccessKeyChange: (val: string) => void;
  secretKey: string;
  onSecretKeyChange: (val: string) => void;
  tenantId?: string;
  onTenantIdChange?: (val: string) => void;
  subscriptionId?: string;
  onSubscriptionIdChange?: (val: string) => void;
  s3AccessKey?: string;
  onS3AccessKeyChange?: (val: string) => void;
  s3SecretKey?: string;
  onS3SecretKeyChange?: (val: string) => void;
  endpoint?: string;
  onEndpointChange?: (val: string) => void;
  showEndpoint?: boolean;
  isEncryptedLabel?: boolean;
  className?: string;
}

const DEFAULT_CSPS = ['aws', 'azure', 'gcp', 'alibaba', 'ibm', 'tencent', 'ncp', 'nhn', 'kt', 'openstack'];

export const CspCredentialForm: React.FC<CspCredentialFormProps> = ({
  csp,
  onCspChange,
  region,
  onRegionChange,
  accessKey,
  onAccessKeyChange,
  secretKey,
  onSecretKeyChange,
  tenantId = '',
  onTenantIdChange,
  subscriptionId = '',
  onSubscriptionIdChange,
  s3AccessKey = '',
  onS3AccessKeyChange,
  s3SecretKey = '',
  onS3SecretKeyChange,
  endpoint = '',
  onEndpointChange,
  showEndpoint = false,
  isEncryptedLabel = false,
  className = ''
}) => {
  const [cspList, setCspList] = useState<string[]>(DEFAULT_CSPS);
  const [regions, setRegions] = useState<{ id: string; name: string }[]>([]);
  const [isLoadingRegions, setIsLoadingRegions] = useState(false);
  const [showSecret, setShowSecret] = useState(false);
  const [showS3Secret, setShowS3Secret] = useState(false);

  // Fetch Providers on Initial Mount
  useEffect(() => {
    const fetchProviders = async () => {
      try {
        const list = await tumblebugApi.getProviders();
        if (list && list.length > 0) {
          setCspList(list);
        }
      } catch (err) {
        console.warn('Failed to load tumblebug providers:', err);
      }
    };
    fetchProviders();
  }, []);

  // Fetch Regions whenever CSP changes
  useEffect(() => {
    let isMounted = true;
    const fetchRegions = async () => {
      setIsLoadingRegions(true);
      try {
        const fetched = await tumblebugApi.getRegions(csp || 'aws');
        const sorted = [...fetched].sort((a, b) => a.id.localeCompare(b.id));
        if (isMounted) {
          setRegions(sorted);
          if (sorted.length > 0 && (!region || !sorted.some(r => r.id === region))) {
            onRegionChange(sorted[0].id);
          }
        }
      } catch (err) {
        console.warn(`Failed to load regions for ${csp}:`, err);
      } finally {
        if (isMounted) setIsLoadingRegions(false);
      }
    };
    fetchRegions();
    return () => {
      isMounted = false;
    };
  }, [csp]);

  const handleCspSelect = (newCsp: string) => {
    onCspChange(newCsp);
    onAccessKeyChange('');
    onSecretKeyChange('');
    if (onTenantIdChange) onTenantIdChange('');
    if (onSubscriptionIdChange) onSubscriptionIdChange('');
    if (onS3AccessKeyChange) onS3AccessKeyChange('');
    if (onS3SecretKeyChange) onS3SecretKeyChange('');
  };

  const activeCsp = (csp || 'aws').toLowerCase();
  const encSuffix = isEncryptedLabel ? ' (Encrypted)' : '';

  return (
    <div className={`p-4 rounded-xl border border-border-main bg-bg-main/40 space-y-4 font-sans text-xs ${className}`}>
      
      {/* Row 1: Cloud Provider (CSP) & Region (Top Section) */}
      <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 pb-3 border-b border-border-main/50">
        <div>
          <label className="block text-text-muted font-bold mb-1.5">
            Cloud Provider (CSP)
          </label>
          <select
            value={activeCsp}
            onChange={(e) => handleCspSelect(e.target.value)}
            className="w-full px-3.5 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-extrabold uppercase focus:outline-none focus:border-emerald-500 cursor-pointer text-xs"
          >
            {cspList.map((p) => (
              <option key={p} value={p}>
                {p.toUpperCase()}
              </option>
            ))}
          </select>
        </div>

        <div>
          <label className="block text-text-muted font-bold mb-1.5 flex items-center justify-between">
            <span>Region</span>
            {isLoadingRegions && <RefreshCw className="w-3 h-3 animate-spin text-emerald-500" />}
          </label>
          <select
            value={region}
            onChange={(e) => onRegionChange(e.target.value)}
            className="w-full px-3.5 py-2 bg-bg-input border border-border-main rounded-xl text-text-main font-mono focus:outline-none focus:border-emerald-500 cursor-pointer text-xs"
          >
            {regions.length > 0 ? (
              regions.map((r) => (
                <option key={r.id} value={r.id}>
                  {r.id} {r.name && r.name !== r.id ? `(${r.name})` : ''}
                </option>
              ))
            ) : (
              <option value={region || 'ap-northeast-2'}>{region || 'ap-northeast-2'}</option>
            )}
          </select>
        </div>
      </div>

      {/* Row 2: Dynamic Credential Input Fields per CSP (Bottom Section) */}
      <div>
        {activeCsp === 'azure' ? (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-3.5">
            <div>
              <label className="block text-text-muted font-medium mb-1">Application (Client) ID{encSuffix}</label>
              <input
                type="text"
                placeholder="00000000-0000-0000-0000-000000000000"
                value={accessKey}
                onChange={(e) => onAccessKeyChange(e.target.value)}
                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
              />
            </div>
            <div>
              <label className="block text-text-muted font-medium mb-1">Client Secret Value{encSuffix}</label>
              <div className="relative">
                <input
                  type={showSecret ? 'text' : 'password'}
                  placeholder="••••••••••••••••"
                  value={secretKey}
                  onChange={(e) => onSecretKeyChange(e.target.value)}
                  className="w-full px-3 py-2 pr-9 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
                />
                <button
                  type="button"
                  onClick={() => setShowSecret(!showSecret)}
                  className="absolute right-2.5 top-1/2 -translate-y-1/2 text-text-muted hover:text-text-main transition cursor-pointer"
                >
                  {showSecret ? <EyeOff className="w-4 h-4" /> : <Eye className="w-4 h-4" />}
                </button>
              </div>
            </div>
            <div>
              <label className="block text-text-muted font-medium mb-1">Directory (Tenant) ID</label>
              <input
                type="text"
                placeholder="Tenant GUID"
                value={tenantId}
                onChange={(e) => onTenantIdChange && onTenantIdChange(e.target.value)}
                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
              />
            </div>
            <div>
              <label className="block text-text-muted font-medium mb-1">Subscription ID</label>
              <input
                type="text"
                placeholder="Subscription GUID"
                value={subscriptionId}
                onChange={(e) => onSubscriptionIdChange && onSubscriptionIdChange(e.target.value)}
                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
              />
            </div>
          </div>
        ) : activeCsp === 'gcp' ? (
          <div className="grid grid-cols-1 sm:grid-cols-3 gap-3.5">
            <div>
              <label className="block text-text-muted font-medium mb-1">Project ID</label>
              <input
                type="text"
                placeholder="my-gcp-project-id"
                value={tenantId}
                onChange={(e) => onTenantIdChange && onTenantIdChange(e.target.value)}
                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
              />
            </div>
            <div>
              <label className="block text-text-muted font-medium mb-1">Client Email (Service Account){encSuffix}</label>
              <input
                type="text"
                placeholder="sa-name@project-id.iam.gserviceaccount.com"
                value={accessKey}
                onChange={(e) => onAccessKeyChange(e.target.value)}
                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
              />
            </div>
            <div>
              <label className="block text-text-muted font-medium mb-1">Private Key (JSON / Key Content){encSuffix}</label>
              <div className="relative">
                <input
                  type={showSecret ? 'text' : 'password'}
                  placeholder="••••••••••••••••"
                  value={secretKey}
                  onChange={(e) => onSecretKeyChange(e.target.value)}
                  className="w-full px-3 py-2 pr-9 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
                />
                <button
                  type="button"
                  onClick={() => setShowSecret(!showSecret)}
                  className="absolute right-2.5 top-1/2 -translate-y-1/2 text-text-muted hover:text-text-main transition cursor-pointer"
                >
                  {showSecret ? <EyeOff className="w-4 h-4" /> : <Eye className="w-4 h-4" />}
                </button>
              </div>
            </div>
          </div>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-3.5">
            <div>
              <label className="block text-text-muted font-medium mb-1">
                {activeCsp === 'tencent' ? 'SecretId' : activeCsp === 'alibaba' ? 'AccessKey ID' : activeCsp === 'ncp' ? 'Access Key ID' : 'Access Key ID'}{encSuffix}
              </label>
              <input
                type="text"
                placeholder={activeCsp === 'ncp' ? 'ncloud_access_key_...' : activeCsp === 'alibaba' ? 'LTAI...' : activeCsp === 'tencent' ? 'AKID...' : 'AKIA...'}
                value={accessKey}
                onChange={(e) => onAccessKeyChange(e.target.value)}
                className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
              />
            </div>
            <div>
              <label className="block text-text-muted font-medium mb-1">
                {activeCsp === 'tencent' ? 'SecretKey' : activeCsp === 'alibaba' ? 'AccessKey Secret' : activeCsp === 'ncp' ? 'Secret Key' : 'Secret Access Key'}{encSuffix}
              </label>
              <div className="relative">
                <input
                  type={showSecret ? 'text' : 'password'}
                  placeholder="••••••••••••••••••••••••"
                  value={secretKey}
                  onChange={(e) => onSecretKeyChange(e.target.value)}
                  className="w-full px-3 py-2 pr-9 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
                />
                <button
                  type="button"
                  onClick={() => setShowSecret(!showSecret)}
                  className="absolute right-2.5 top-1/2 -translate-y-1/2 text-text-muted hover:text-text-main transition cursor-pointer"
                >
                  {showSecret ? <EyeOff className="w-4 h-4" /> : <Eye className="w-4 h-4" />}
                </button>
              </div>
            </div>
          </div>
        )}

        {/* S3 Interoperability Credentials (AWS S3-Compatible Interoperability) */}
        {['gcp', 'azure', 'ibm', 'openstack', 'nhn', 'nhncloud', 'kt', 'ktcloud'].includes(activeCsp) && (
          <div className="pt-3.5 border-t border-border-main/40 mt-3.5 space-y-2">
            <div className="flex flex-wrap items-center gap-1.5">
              <span className="text-xs font-extrabold text-teal-400">
                S3 Interoperability Credentials (AWS S3-Compatible Interoperability)
              </span>
              <span className="text-[11px] text-text-muted font-mono">
                (S3AccessKey &amp; S3SecretKey for Object Storage API Control)
              </span>
            </div>
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-3.5">
              <div>
                <label className="block text-text-muted font-medium mb-1">S3 Access Key (S3AccessKey)</label>
                <input
                  type="text"
                  placeholder="S3 Interoperability Access Key"
                  value={s3AccessKey}
                  onChange={(e) => onS3AccessKeyChange && onS3AccessKeyChange(e.target.value)}
                  className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
                />
              </div>
              <div>
                <label className="block text-text-muted font-medium mb-1">S3 Secret Key (S3SecretKey)</label>
                <div className="relative">
                  <input
                    type={showS3Secret ? 'text' : 'password'}
                    placeholder="••••••••••••••••"
                    value={s3SecretKey}
                    onChange={(e) => onS3SecretKeyChange && onS3SecretKeyChange(e.target.value)}
                    className="w-full px-3 py-2 pr-9 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
                  />
                  <button
                    type="button"
                    onClick={() => setShowS3Secret(!showS3Secret)}
                    className="absolute right-2.5 top-1/2 -translate-y-1/2 text-text-muted hover:text-text-main transition cursor-pointer"
                  >
                    {showS3Secret ? <EyeOff className="w-4 h-4" /> : <Eye className="w-4 h-4" />}
                  </button>
                </div>
              </div>
            </div>
          </div>
        )}

        {/* Optional Custom S3 Endpoint URL */}
        {showEndpoint && onEndpointChange && (
          <div className="mt-3">
            <label className="block text-text-muted font-medium mb-1">Custom S3 Endpoint URL (Optional)</label>
            <input
              type="text"
              placeholder={`https://s3.${region || 'ap-northeast-2'}.amazonaws.com`}
              value={endpoint}
              onChange={(e) => onEndpointChange(e.target.value)}
              className="w-full px-3 py-2 bg-bg-input border border-border-main rounded-lg text-text-main font-mono focus:outline-none focus:border-emerald-500 text-xs"
            />
          </div>
        )}
      </div>

    </div>
  );
};
