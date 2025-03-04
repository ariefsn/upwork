import { env } from "$env/dynamic/public";
import { formatCurrency, generatePageMetaData } from "$lib";
import { earningsUserYearly } from "$lib/graphql";
import type { EarningsDataMonthly } from "$lib/graphql/generated";
import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params, url }) => {
  const userID = params.id
  const searchParams = url.searchParams
  const qYear = searchParams.get('year')
  const date = new Date()
  const year = qYear ? +qYear : date.getFullYear()

  const res = await earningsUserYearly(userID, year)
  const user = res.data?.getUser
  const itemsLimited = (res.data?.getEarnings ?? []).filter(e => e !== null)
  const years = (res.data?.getEarningsYears ?? []).filter(e => e !== null)

  if (itemsLimited.length === 0 || !user) {
    error(404, 'NO_EARNINGS')
  }

  const itemsMonths = itemsLimited.map(e => e.month)
  const items: EarningsDataMonthly[] = []
  for (let i = 0; i < 12; i++) {
    const month = i + 1;
    const monthIndex = itemsMonths.indexOf(month)
    if (monthIndex > -1) {
      items.push(itemsLimited[monthIndex])
      continue
    }
    items.push({
      userID,
      month,
      year,
      totalAmount: 0,
      totalFee: 0,
      items: []
    })
  }

  const pageMetaTags = generatePageMetaData({
    url,
    appName: env.PUBLIC_APP_NAME,
    title: `${user.fullName}`,
    description: `${user.fullName} has made ${formatCurrency(items.reduce((total, e) => total + (e.totalAmount - e.totalFee), 0))} in ${year}.`,
    tags: ['upwork', 'salary', 'earnings', 'transactions', 'invoices', user.fullName, user.id],
  })

  return {
    items,
    user,
    year,
    years,
    pageMetaTags,
  }
}