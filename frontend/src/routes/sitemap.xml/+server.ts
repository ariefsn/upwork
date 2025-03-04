import { env } from '$env/dynamic/public'
import { userIds } from '$lib/graphql'
import { error, type RequestHandler } from '@sveltejs/kit'
import * as sitemap from 'super-sitemap'

export const GET: RequestHandler = async ({ cookies }) => {
  let freelancerIds: string[]

  try {
    const [resUserIds] = await Promise.all([
      userIds(),
    ])

    freelancerIds = (resUserIds?.data?.getUserIds || []).map((id) => id)
  } catch (err) {
    return error(500, err as Error)
  }

  return await sitemap.response({
    origin: env.PUBLIC_ORIGIN,
    paramValues: {
      '/freelancers/[id]': freelancerIds,
    },
    defaultChangefreq: 'daily',
    defaultPriority: 0.7,
    sort: 'alpha',
  })
}