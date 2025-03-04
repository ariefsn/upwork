<script lang="ts">
	import * as Select from '$lib/components/ui/select/index.js';
	import { changeLocale, i18nLocales, locale } from '$lib/translations';

	let open = $state(false);
</script>

<Select.Root
	selected={{ value: $locale, label: $locale.toUpperCase() }}
	onSelectedChange={(v) => {
		changeLocale(v?.value || 'en');
		open = false;
	}}
>
	<Select.Trigger
		class="w-auto border-0 bg-transparent p-0 text-lg duration-300"
		aria-label="Select language"
		onclick={() => {
			open = !open;
		}}
	>
		{$locale.toUpperCase()}
	</Select.Trigger>
	<Select.Content class="!w-12">
		{#each $i18nLocales.sort() as l}
			<Select.Item value={l} label={l.toUpperCase()} class="px-2">
				{l.toUpperCase()}
			</Select.Item>
		{/each}
	</Select.Content>
	<Select.Input name="favoriteFruit" />
</Select.Root>
