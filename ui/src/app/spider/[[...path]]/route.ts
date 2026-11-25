import { NextRequest } from 'next/server';
import { proxy } from '@/lib/proxy';

async function handler(
  req: NextRequest,
  { params }: { params: Promise<{ path?: string[] }> },
) {
  const { path = [] } = await params;
  const targetBase = process.env.SPIDER_ENDPOINT || 'http://localhost:1024';
  return proxy(req, targetBase, 'spider', path, {
    auth: {
      username: process.env.SPIDER_API_USERNAME,
      password: process.env.SPIDER_API_PASSWORD,
      required: true,
    },
  });
}

export {
  handler as GET,
  handler as POST,
  handler as PUT,
  handler as DELETE,
  handler as PATCH,
  handler as HEAD,
};
