<script lang="ts">
	import { t } from '$lib/translations';

	let {
		q,
		onClick
	}: {
		q?: number;
		onClick?: (e: MouseEvent) => void;
	} = $props();
</script>

<a href={`/faqs${q ? `?q=${q}` : ''}`} class="underline" onclick={onClick}>
	{$t('faq.faq', { faqCount: 2 })}
</a>
