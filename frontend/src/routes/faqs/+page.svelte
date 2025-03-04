<script lang="ts">
	import { page } from '$app/state';
	import Accordion from '$lib/components/atomic/atoms/Accordion.svelte';
	import Container from '$lib/components/atomic/atoms/Container.svelte';
	import type { FaqItem } from '$lib/models';
	import { uploadHistoryState } from '$lib/stores';
	import { t } from '$lib/translations';
	import { onMount } from 'svelte';

	const faqKeys = [
		'whatIsThis',
		'howDoesItWork',
		'howCanIGetMyUpworkID',
		'canIUseMyUsernameInsteadOfMyUpworkID',
		'isTheAppFree',
		'isTheAppNeedsToAccessMyUpworkAccount',
		'whereICanExportMyEarningsFromUpwork',
		'isItWillReadAllOfTheTransactionTypes',
		'isMyUploadedCSVFileWillBeStored',
		'canIDeleteMyHistory',
		'canISeeTheSourceCode'
	];

	const q = +(page.url.searchParams.get('q') ?? '-1');

	const sourceCodeLink = 'https://github.com/ariefsn/upwork';

	const faqs = $derived.by<FaqItem[]>(() => {
		return faqKeys.map((key) => {
			let answer = $t(`faq.${key}.answer`);
			switch (key) {
				case 'canISeeTheSourceCode':
					answer = $t(`faq.${key}.answer`, {
						faqSourceCodeLink: `<a href="${sourceCodeLink}" class="underline" target="_blank" rel="noopener noreferrer">${sourceCodeLink}</a>`
					});
					break;
			}
			return {
				id: key,
				question: $t(`faq.${key}.question`),
				answer
			};
		});
	});

	const selected = $derived.by(() => {
		if (q > -1) {
			return faqKeys[q - 1];
		}

		return undefined;
	});

	onMount(() => {
		uploadHistoryState.set(false);
	});
</script>

<div class="flex h-screen flex-col items-center justify-center">
	<div class="mb-2 text-lg font-bold">{$t('faq.title')}</div>
	<Container class="h-4/6 overflow-y-auto">
		<Accordion items={faqs} value={selected} />
	</Container>
</div>
