import { NextRequest } from 'next/server';
import { proxy } from '@/lib/proxy';

const TARGET = process.env.SPIDER_ENDPOINT || 'http://localhost:1024';

async function handler(
  req: NextRequest,
  { params }: { params: Promise<{ path?: string[] }> },
) {
  const { path = [] } = await params;
  return proxy(req, TARGET, 'spider', path);
}

export {
  handler as GET,
  handler as POST,
  handler as PUT,
  handler as DELETE,
  handler as PATCH,
  handler as HEAD,
};
