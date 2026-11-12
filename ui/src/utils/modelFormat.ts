// Shared formatting, stage badge metadata, sample detection, and naming helpers for migration models.

export const BUILTIN_SOURCE_INFRA_SAMPLE_IDS = new Set([
  'sample-source-migration-testbed-1',
  'sample-source-9node-testbed-1',
  'sample-source-15node-dc-1',
  'sample-source-gpu-infra-1',
  'sample-source-infra-1',
  'sample-9-node-testbed',
  'sample-15-node-enterprise',
  'demo-model-1',
  'demo-gpu-model-1',
]);

export const BUILTIN_CLOUD_INFRA_SAMPLE_IDS = new Set([
  'sample-cloud-migration-testbed-1',
  'sample-cloud-ai-stack-testbed-1',
  'sample-mig-testbed-cloud-1',
  'cloud-demo-1',
]);

export const BUILTIN_STORAGE_SAMPLE_IDS = new Set([
  'sample-source-storage-01',
  'sample-target-storage-01',
]);

export type ModelStageKind = 'sample' | 'source' | 'planning' | 'target' | 'draft';

// Check whether a model envelope or ID belongs to a built-in sample model.
export function isBuiltInSampleModel(
  modelOrId: string | { id?: string | null; name?: string | null } | null | undefined
): boolean {
  if (!modelOrId) return false;
  if (typeof modelOrId === 'string') {
    return (
      BUILTIN_SOURCE_INFRA_SAMPLE_IDS.has(modelOrId) ||
      BUILTIN_CLOUD_INFRA_SAMPLE_IDS.has(modelOrId) ||
      BUILTIN_STORAGE_SAMPLE_IDS.has(modelOrId)
    );
  }
  const id = String(modelOrId.id || '').trim();
  if (
    id &&
    (BUILTIN_SOURCE_INFRA_SAMPLE_IDS.has(id) ||
      BUILTIN_CLOUD_INFRA_SAMPLE_IDS.has(id) ||
      BUILTIN_STORAGE_SAMPLE_IDS.has(id))
  ) {
    return true;
  }
  const rawName = String(modelOrId.name || '').trim();
  return /^\[sample\]/i.test(rawName);
}

// Strip leading bracket tags such as [Sample], [Source], [Planning], [Target], or [Draft].
export function cleanModelDisplayName(name?: string | null, fallback = 'unnamed-model'): string {
  const raw = String(name || '').trim();
  if (!raw) return fallback;
  const cleaned = raw.replace(/^(\s*\[(sample|source|planning|target|draft)\]\s*)+/i, '').trim();
  return cleaned || fallback;
}

// Normalize version string by stripping leading v/V characters.
export function normalizeModelVersion(version?: string | null, fallback = '1.0.0'): string {
  const raw = String(version || '').trim().replace(/^[vV]+/, '').trim();
  return raw || fallback;
}

// Format source or planning model option label without stage brackets.
export function formatSourceOptionText(
  model: { name?: string | null; version?: string | null; isDraft?: boolean },
  fallbackVersion = '1.0',
  isDraft = false
): string {
  const cleanName = cleanModelDisplayName(model.name);
  if (isDraft || model.isDraft) {
    return `${cleanName} (Unsaved Draft)`;
  }
  const ver = normalizeModelVersion(model.version, fallbackVersion);
  return `${cleanName} (v${ver})`;
}

// Format target cloud or object storage model option label with CSP and version.
export function formatTargetOptionText(
  model: { name?: string | null; version?: string | null },
  csp?: string | null,
  isDraft = false
): string {
  const cleanName = cleanModelDisplayName(model.name);
  const cspLabel = String(csp || 'CLOUD').trim().toUpperCase() || 'CLOUD';
  if (isDraft) {
    return `${cleanName} (${cspLabel} / Unsaved Draft)`;
  }
  const ver = normalizeModelVersion(model.version, '1.0');
  return `${cleanName} (${cspLabel} / v${ver})`;
}

// Resolve canonical stage kind for badge rendering.
export function resolveModelStageKind(
  model: { id?: string | null; name?: string | null; stage?: string | null; isDraft?: boolean } | null | undefined,
  fallbackStage: 'source' | 'planning' | 'target' = 'source'
): ModelStageKind {
  if (!model || model.isDraft) return 'draft';
  if (isBuiltInSampleModel(model)) return 'sample';
  if (model.stage === 'planning') return 'planning';
  if (model.stage === 'target') return 'target';
  if (model.stage === 'source') return 'source';
  return fallbackStage;
}

// Return label and brand-aligned Tailwind badge classes for a model stage kind.
export function getStageBadgeMeta(kind: ModelStageKind): { label: string; className: string } {
  switch (kind) {
    case 'sample':
      return {
        label: 'Sample',
        className: 'bg-amber-500/15 text-amber-600 dark:text-amber-400 border-amber-500/30',
      };
    case 'source':
      return {
        label: 'Source',
        className: 'bg-emerald-500/15 text-emerald-600 dark:text-emerald-400 border-emerald-500/30',
      };
    case 'planning':
      return {
        label: 'Planning',
        className: 'bg-teal-500/15 text-teal-600 dark:text-teal-400 border-teal-500/30',
      };
    case 'target':
      return {
        label: 'Target',
        className: 'bg-emerald-500/15 text-emerald-600 dark:text-emerald-400 border-emerald-500/30',
      };
    case 'draft':
    default:
      return {
        label: 'Draft',
        className: 'bg-amber-500/15 text-amber-600 dark:text-amber-400 border-amber-500/30',
      };
  }
}

// Build clean default model name for Source, Planning, or Target save modals.
export function buildDefaultModelName(
  stage: 'source' | 'planning' | 'target',
  baseHint?: string | null,
  csp?: string | null
): string {
  const cleanedBase = cleanModelDisplayName(baseHint, '')
    .replace(/\.(json|yaml|yml)$/i, '')
    .replace(/-(source|planning|plan|target)(-model)?(-rev\d+)?$/i, '')
    .trim();

  const root = cleanedBase || (stage === 'target' ? 'cloud' : 'infra');
  if (stage === 'source') {
    return `${root}-source`;
  }
  if (stage === 'planning') {
    return `${root}-plan`;
  }
  const cleanCsp = String(csp || '').trim().toLowerCase();
  if (cleanCsp && !root.toLowerCase().endsWith(`-${cleanCsp}`)) {
    return `${root}-${cleanCsp}-target`;
  }
  return `${root}-target`;
}

// Validate cloud resource identifier against pkg/core/common/naming.go IsValidName rules.
export function isValidCloudResourceName(name?: string | null): boolean {
  const s = String(name || '').trim();
  if (s.length < 3 || s.length > 63) return false;
  return /^[a-zA-Z0-9][a-zA-Z0-9-]*$/.test(s);
}

// Preserve valid targetInfra.name or generate a compliant slug from artifact name.
export function resolveValidTargetInfraName(
  existingInfraName?: string | null,
  artifactName?: string | null,
  csp?: string | null
): string {
  const trimmedExisting = String(existingInfraName || '').trim();
  if (isValidCloudResourceName(trimmedExisting)) {
    return trimmedExisting;
  }
  const trimmedArtifact = cleanModelDisplayName(artifactName, '');
  if (isValidCloudResourceName(trimmedArtifact)) {
    return trimmedArtifact;
  }
  const slug = trimmedArtifact
    .toLowerCase()
    .replace(/[^a-z0-9-]+/g, '-')
    .replace(/^-+|-+$/g, '')
    .replace(/-{2,}/g, '-')
    .slice(0, 55);
  if (slug.length >= 3 && /^[a-z0-9]/.test(slug)) {
    return slug;
  }
  const cspSlug = String(csp || 'cloud').toLowerCase().replace(/[^a-z0-9]+/g, '') || 'cloud';
  return `${cspSlug}-infra-01`;
}
