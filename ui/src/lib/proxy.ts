import { NextRequest } from 'next/server';

export interface ProxyOptions {
  auth?: {
    username?: string;
    password?: string;
    required?: boolean;
  };
}

/**
 * Generic reverse-proxy helper for Next.js API route handlers.
 *
 * Forwards an incoming request to a backend service, preserving the method,
 * headers, query string, and body. Backend endpoints and Basic Auth credentials
 * are resolved from server-side environment variables, so browser code never
 * needs to know service URLs or credentials.
 */
export async function proxy(
  req: NextRequest,
  targetBase: string,
  prefix: string,
  path: string[],
  options?: ProxyOptions,
): Promise<Response> {
  const suffix = path.length > 0 ? `/${path.join('/')}` : '';
  const targetUrl = `${targetBase.replace(/\/$/, '')}/${prefix}${suffix}${req.nextUrl.search}`;

  const headers = new Headers(req.headers);
  // Strip hop-by-hop / host-specific headers before forwarding
  headers.delete('host');
  headers.delete('connection');
  headers.delete('content-length');

  if (options?.auth) {
    headers.delete('authorization');
    const username = options.auth.username?.trim() || '';
    const password = options.auth.password || '';
    if (options.auth.required && (!username || !password)) {
      return new Response(
        JSON.stringify({
          error: `Backend API credentials are not configured for ${prefix}. Set the required server environment variables.`,
        }),
        { status: 500, headers: { 'content-type': 'application/json' } },
      );
    }
    if (username && password) {
      headers.set('authorization', `Basic ${Buffer.from(`${username}:${password}`).toString('base64')}`);
    }
  }

  const init: RequestInit & { duplex?: 'half' } = {
    method: req.method,
    headers,
    redirect: 'manual',
  };

  if (req.method !== 'GET' && req.method !== 'HEAD') {
    init.body = await req.arrayBuffer();
    init.duplex = 'half';
  }

  let upstream: Response;
  try {
    upstream = await fetch(targetUrl, init);
  } catch {
    return new Response(
      JSON.stringify({ error: `Upstream service unavailable: ${prefix}` }),
      { status: 502, headers: { 'content-type': 'application/json' } },
    );
  }

  const respHeaders = new Headers(upstream.headers);
  // These are recomputed by the runtime and must not be forwarded verbatim
  respHeaders.delete('content-encoding');
  respHeaders.delete('content-length');
  respHeaders.delete('transfer-encoding');

  return new Response(upstream.body, {
    status: upstream.status,
    statusText: upstream.statusText,
    headers: respHeaders,
  });
}
