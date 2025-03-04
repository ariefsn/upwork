import { env } from '$env/dynamic/public';
import { type RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async () => {
  const body = [
    'User-agent: *',
    'Allow: /',
    '',
    `Sitemap: ${env.PUBLIC_ORIGIN}/sitemap.xml`
  ].join('\n').trim();

  const headers = {
    'Content-Type': 'text/plain',
  };

  return new Response(body, { headers });
}