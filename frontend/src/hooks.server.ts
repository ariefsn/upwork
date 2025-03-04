import { env } from '$env/dynamic/public';
import { type Handle } from '@sveltejs/kit';
import { handleProxy } from 'sveltekit-proxy';

export const handle: Handle = async ({ event, resolve }) => {
  const req = event.request;

  event.request.headers.set('X-Client-IP', req.headers.get('X-Forwarded-For') ?? "");

  if (event.url.pathname.includes('/api')) {
    return handleProxy({
      target: env.PUBLIC_API_URL ?? '',
      rewrite: (path) => path.replace('/api', ''),
      origin: env.PUBLIC_ORIGIN
    })({ event, resolve });
  }

  const response = await resolve(event);

  return response;
};
