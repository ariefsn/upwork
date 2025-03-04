<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Command from '$lib/components/ui/command/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import type { DropdownItem } from '$lib/models';
	import { cn } from '$lib/utils.js';
	import Check from 'lucide-svelte/icons/check';
	import ChevronsUpDown from 'lucide-svelte/icons/chevrons-up-down';
	import { tick } from 'svelte';

	let {
		items = [],
		searchPlaceholder = 'Search item...',
		emptyPlaceholder = 'No item found.',
		value,
		onChange
	}: {
		items?: DropdownItem[];
		searchPlaceholder?: string;
		emptyPlaceholder?: string;
		value?: string;
		onChange?: (value: string) => void;
	} = $props();

	let open = $state(false);

	const selectedValue = $derived(items.find((f) => f.value === value)?.label ?? 'Select a item...');

	function closeAndFocusTrigger(triggerId: string) {
		open = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
</script>

<Popover.Root bind:open let:ids>
	<Popover.Trigger asChild let:builder>
		<Button
			builders={[builder]}
			variant="outline"
			role="combobox"
			aria-expanded={open}
			class="w-[100px] justify-between"
			aria-label={selectedValue}
		>
			{selectedValue}
			<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
		</Button>
	</Popover.Trigger>
	<Popover.Content class="w-[100px] p-0">
		<Command.Root {value}>
			<Command.Input placeholder={searchPlaceholder} />
			<Command.Empty>{emptyPlaceholder}</Command.Empty>
			<Command.Group>
				{#each items as item}
					<Command.Item
						value={item.value}
						onSelect={(currentValue: string) => {
							value = currentValue;
							onChange?.(value);
							closeAndFocusTrigger(ids.trigger);
						}}
					>
						<Check class={cn('mr-2 h-4 w-4', value !== item.value && 'text-transparent')} />
						{item.label}
					</Command.Item>
				{/each}
			</Command.Group>
		</Command.Root>
	</Popover.Content>
</Popover.Root>
