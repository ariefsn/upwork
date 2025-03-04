<script lang="ts">
	import { page } from '$app/state';
	import { t } from '$lib/translations';
	import { ChevronLeft } from 'lucide-svelte';

	const message = $derived.by(() => {
		switch (page.error?.message) {
			case 'NO_EARNINGS':
				return $t('upwork.noEarnings');

			default:
				return page.error?.message;
		}
	});
</script>

<div class="flex h-screen items-center justify-center">
	<div class="h-fit text-center">
		<div style="padding-bottom: 12px;">
			<p class="mb-2 text-3xl font-bold">
				{page.status}
			</p>
			<p class="text-xl">
				{message}
			</p>
		</div>
		<a href="/" class="animate-border-bottom mt-5">
			<div class="flex items-center justify-center">
				<ChevronLeft class="mr-1 h-6 w-6" />
				<span>{$t('app.back')}</span>
			</div>
		</a>
	</div>
</div>
