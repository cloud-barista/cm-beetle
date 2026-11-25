import { NextRequest } from 'next/server';
import { proxy } from '@/lib/proxy';

async function handler(
  req: NextRequest,
  { params }: { params: Promise<{ path?: string[] }> },
) {
  const { path = [] } = await params;
  const targetBase = process.env.BEETLE_ENDPOINT || 'http://localhost:8056';
  return proxy(req, targetBase, 'beetle', path, {
    auth: {
      username: process.env.BEETLE_API_USERNAME,
      password: process.env.BEETLE_API_PASSWORD,
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
