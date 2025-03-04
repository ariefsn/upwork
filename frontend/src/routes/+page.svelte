<script lang="ts">
	import Container from '$lib/components/atomic/atoms/Container.svelte';
	import Freelancer from '$lib/components/atomic/molecules/Freelancer.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { uploadHistoryState } from '$lib/stores';
	import { t } from '$lib/translations';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<div class="flex h-screen flex-col items-center justify-center">
	<Container class="max-h-5/6 overflow-y-auto">
		{#if data.items.length === 0}
			<div class="text-center">
				<h1 class="text-3xl font-bold">{$t('upwork.noFreelancers')}</h1>
				<Button size="sm" class="mt-2" onclick={() => uploadHistoryState.set(true)}>
					<p>{$t('upwork.uploadEarnings')}</p>
				</Button>
			</div>
		{/if}

		{#each data.items as { user, amount, fee }}
			<Freelancer {user} {amount} {fee} />
		{/each}
	</Container>
</div>
