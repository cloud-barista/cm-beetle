import { NextRequest, NextResponse } from 'next/server';

const ALLOWED_PORTS = new Set<number>([80, 443, 3000, 8080, 11434]);
const ALLOWED_PATHS = new Set<string>(['/', '/api/config', '/api/tags', '/api/ps', '/api/generate']);
const MAX_RESPONSE_BYTES = 65536;
const PROBE_TIMEOUT_MS = 15000;

// Validate IPv4 literal and reject loopback, unspecified, link-local/IMDS, CGNAT metadata, and multicast ranges.
function parseSafeIpv4Target(rawHost: string): { ip?: string; error?: string } {
  const trimmed = rawHost.trim();
  const ipv4StrictRegex = /^(?:0|[1-9]\d{0,2})\.(?:0|[1-9]\d{0,2})\.(?:0|[1-9]\d{0,2})\.(?:0|[1-9]\d{0,2})$/;
  if (!ipv4StrictRegex.test(trimmed)) {
    return { error: 'Host must be a canonical IPv4 address (e.g., 10.0.1.11 or public IPv4).' };
  }

  const octets = trimmed.split('.').map((part) => Number(part));
  if (octets.some((n) => !Number.isInteger(n) || n < 0 || n > 255)) {
    return { error: 'Host IPv4 octets must be between 0 and 255.' };
  }

  const [a, b, c] = octets;
  if (a === 0) {
    return { error: 'Unspecified IPv4 range (0.0.0.0/8) is blocked.' };
  }
  if (a === 127) {
    return { error: 'Loopback IPv4 range (127.0.0.0/8) is blocked.' };
  }
  if (a === 169 && b === 254) {
    return { error: 'Link-local and cloud metadata range (169.254.0.0/16) is blocked.' };
  }
  if (a === 100 && b >= 64 && b <= 127) {
    return { error: 'Carrier-grade NAT / metadata range (100.64.0.0/10) is blocked.' };
  }
  if (a === 192 && b === 0 && c === 0) {
    return { error: 'Reserved IETF protocol range (192.0.0.0/24) is blocked.' };
  }
  if (a >= 224) {
    return { error: 'Multicast and broadcast IPv4 ranges (224.0.0.0/4) are blocked.' };
  }

  return { ip: `${a}.${b}.${c}.${octets[3]}` };
}

async function readBoundedText(response: Response, maxBytes: number): Promise<string> {
  if (!response.body) {
    return '';
  }
  const reader = response.body.getReader();
  const decoder = new TextDecoder('utf-8');
  let total = 0;
  let output = '';

  while (true) {
    const { done, value } = await reader.read();
    if (done || !value) {
      break;
    }
    total += value.byteLength;
    if (total > maxBytes) {
      const allowedSlice = value.subarray(0, Math.max(0, value.byteLength - (total - maxBytes)));
      output += decoder.decode(allowedSlice, { stream: false });
      await reader.cancel();
      break;
    }
    output += decoder.decode(value, { stream: true });
  }
  output += decoder.decode();
  return output;
}

export async function POST(req: NextRequest) {
  try {
    const body = await req.json().catch(() => null);
    if (!body || typeof body !== 'object') {
      return NextResponse.json({ ok: false, error: 'Invalid JSON request body.' }, { status: 400 });
    }

    const rawHost = typeof body.host === 'string' ? body.host : '';
    const parsedHost = parseSafeIpv4Target(rawHost);
    if (!parsedHost.ip) {
      return NextResponse.json({ ok: false, error: parsedHost.error || 'Invalid target host.' }, { status: 400 });
    }

    const numericPort = Number(body.port);
    if (!Number.isInteger(numericPort) || !ALLOWED_PORTS.has(numericPort)) {
      return NextResponse.json(
        { ok: false, error: 'Port is not in allowlist (allowed: 80, 443, 3000, 8080, 11434).' },
        { status: 400 }
      );
    }

    const rawPath = typeof body.path === 'string' ? body.path.trim() : '/';
    if (!ALLOWED_PATHS.has(rawPath)) {
      return NextResponse.json(
        { ok: false, error: 'Path is not in allowlist (allowed: /, /api/config, /api/tags, /api/ps, /api/generate).' },
        { status: 400 }
      );
    }

    const isGenerate = rawPath === '/api/generate';
    const method = isGenerate && String(body.method || 'POST').toUpperCase() === 'POST' ? 'POST' : 'GET';

    let serializedPayload: string | undefined;
    if (method === 'POST' && isGenerate) {
      const rawModel = typeof body.model === 'string' && body.model.trim() ? body.model.trim() : 'gemma4:e2b';
      if (!/^[a-zA-Z0-9._:-]{1,64}$/.test(rawModel)) {
        return NextResponse.json({ ok: false, error: 'Invalid model identifier format.' }, { status: 400 });
      }
      const rawPrompt =
        typeof body.prompt === 'string' && body.prompt.trim()
          ? body.prompt.trim().slice(0, 200)
          : 'Reply in five words confirming GPU inference readiness.';
      serializedPayload = JSON.stringify({
        model: rawModel,
        prompt: rawPrompt,
        stream: false
      });
    }

    const protocol = numericPort === 443 ? 'https:' : 'http:';
    const targetUrl = new URL(rawPath, `${protocol}//${parsedHost.ip}:${numericPort}`);

    const startedAt = Date.now();
    const upstream = await fetch(targetUrl.toString(), {
      method,
      headers: serializedPayload
        ? { 'Content-Type': 'application/json', Accept: 'application/json' }
        : { Accept: 'application/json, text/plain;q=0.9, */*;q=0.8' },
      body: serializedPayload,
      redirect: 'error',
      cache: 'no-store',
      signal: AbortSignal.timeout(PROBE_TIMEOUT_MS)
    });
    const latencyMs = Math.max(1, Date.now() - startedAt);
    const textBody = await readBoundedText(upstream, MAX_RESPONSE_BYTES);

    let parsedJson: Record<string, any> | null = null;
    try {
      parsedJson = textBody ? JSON.parse(textBody) : null;
    } catch {
      parsedJson = null;
    }

    const result: Record<string, any> = {
      ok: upstream.ok,
      httpStatus: upstream.status,
      latencyMs,
      checkedAt: new Date().toISOString()
    };

    if (rawPath === '/api/config' && parsedJson) {
      result.openWebUi = {
        name: String(parsedJson.name || 'Open WebUI'),
        version: String(parsedJson.version || 'active'),
        status: Boolean(parsedJson.status ?? upstream.ok)
      };
    } else if (rawPath === '/api/tags' && parsedJson && Array.isArray(parsedJson.models)) {
      result.models = parsedJson.models.map((m: any) => ({
        name: String(m.name || m.model || 'unknown'),
        sizeBytes: Number(m.size || 0),
        digest: String(m.digest || ''),
        parameterSize: String(m.details?.parameter_size || ''),
        quantizationLevel: String(m.details?.quantization_level || '')
      }));
    } else if (rawPath === '/api/ps' && parsedJson && Array.isArray(parsedJson.models)) {
      result.runningModels = parsedJson.models.map((m: any) => ({
        name: String(m.name || m.model || 'unknown'),
        sizeBytes: Number(m.size || 0),
        sizeVramBytes: Number(m.size_vram || 0),
        digest: String(m.digest || '')
      }));
    } else if (rawPath === '/api/generate' && parsedJson) {
      const evalCount = Number(parsedJson.eval_count || 0);
      const evalDurationNs = Number(parsedJson.eval_duration || 0);
      const tokensPerSec =
        evalCount > 0 && evalDurationNs > 0
          ? Math.round((evalCount / (evalDurationNs / 1e9)) * 10) / 10
          : undefined;
      result.inference = {
        model: String(parsedJson.model || ''),
        responseSnippet: String(parsedJson.response || '').trim().slice(0, 160),
        evalCount,
        tokensPerSec
      };
    }

    return NextResponse.json(result, { status: 200 });
  } catch (err: any) {
    const message =
      err?.name === 'TimeoutError' || err?.name === 'AbortError'
        ? `Probe timed out after ${PROBE_TIMEOUT_MS / 1000}s`
        : err?.message || 'Target endpoint unreachable';
    return NextResponse.json(
      {
        ok: false,
        httpStatus: 0,
        error: message,
        checkedAt: new Date().toISOString()
      },
      { status: 200 }
    );
  }
}
