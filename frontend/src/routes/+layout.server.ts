import { PUBLIC_GOOGLE_TAG_ID } from "$env/static/public"
import LinkFaq from "$lib/components/atomic/atoms/LinkFaq.svelte"
import { render } from "svelte/server"
import type { LayoutServerLoad } from "./$types"

export const load: LayoutServerLoad = async ({ url }) => {
  const { body: faqLink } = render(LinkFaq, { props: { q: 3 } })

  return {
    faqLink,
    gTagId: PUBLIC_GOOGLE_TAG_ID,
  }
}