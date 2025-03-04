import { env } from "$env/dynamic/public";
import { generatePageMetaData } from "$lib";
import { earningsUsers } from "$lib/graphql";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ url }) => {
  const date = new Date()
  const res = await earningsUsers(date.getFullYear())
  const items = (res.data?.getEarningsUsers ?? []).filter(e => e !== null)

  const pageMetaTags = generatePageMetaData({
    url,
    appName: env.PUBLIC_APP_NAME,
    title: 'Home',
    description: 'Upload and share your Upwork earnings with your friends.',
    tags: ['upwork', 'salary', 'earnings', 'transactions', 'invoices'],
  })

  return {
    items,
    pageMetaTags
  }
}