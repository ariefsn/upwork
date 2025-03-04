import { browser } from "$app/environment";
import { getLocale, loadTranslations } from "$lib/translations";
import type { LayoutLoad } from "./$types";

export const load: LayoutLoad = async ({ url, data }) => {
  const { pathname } = url

  let initLocale = 'en'

  if (browser) {
    initLocale = getLocale()
  }

  await loadTranslations(initLocale, pathname)

  return {
    ...data
  }
}