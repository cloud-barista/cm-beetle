import { NextRequest } from 'next/server';
import { proxy } from '@/lib/proxy';

const TARGET = process.env.HONEYBEE_ENDPOINT || 'http://localhost:8081';

async function handler(
  req: NextRequest,
  { params }: { params: Promise<{ path?: string[] }> },
) {
  const { path = [] } = await params;
  return proxy(req, TARGET, 'honeybee', path);
}

export {
  handler as GET,
  handler as POST,
  handler as PUT,
  handler as DELETE,
  handler as PATCH,
  handler as HEAD,
};
