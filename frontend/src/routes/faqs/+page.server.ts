import { env } from "$env/dynamic/public";
import { generatePageMetaData } from "$lib";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params, url }) => {
  const pageMetaTags = generatePageMetaData({
    url,
    appName: env.PUBLIC_APP_NAME,
    title: `FAQs`,
    description: `Have a question? Find the answer here.`,
    tags: ['upwork', 'salary', 'earnings', 'transactions', 'invoices', 'faq', 'question', 'answer'],
  })

  return {
    pageMetaTags,
  }
}