import { NextRequest } from 'next/server';

/**
 * Generic reverse-proxy helper for Next.js API route handlers.
 *
 * Forwards an incoming request to a backend service, preserving the method,
 * headers, query string, and body. Backend endpoints are resolved from
 * server-side environment variables, so browser code never needs to know the
 * real service URLs (which differ between local dev and Docker Compose).
 *
 * @param req        The incoming Next.js request.
 * @param targetBase Base URL of the backend service (e.g. http://cm-honeybee:8081).
 * @param prefix     Path prefix the backend serves under (e.g. "honeybee").
 * @param path       Catch-all path segments after the prefix.
 */
export async function proxy(
  req: NextRequest,
  targetBase: string,
  prefix: string,
  path: string[],
): Promise<Response> {
  const suffix = path.length > 0 ? `/${path.join('/')}` : '';
  const targetUrl = `${targetBase.replace(/\/$/, '')}/${prefix}${suffix}${req.nextUrl.search}`;

  const headers = new Headers(req.headers);
  // Strip hop-by-hop / host-specific headers before forwarding
  headers.delete('host');
  headers.delete('connection');
  headers.delete('content-length');

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
