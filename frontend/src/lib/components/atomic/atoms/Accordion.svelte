<script lang="ts">
	import * as Accordion from '$lib/components/ui/accordion/index.js';
	import type { FaqItem } from '$lib/models';
	import { cn } from '$lib/utils';
	import {} from '@sveltelegos-blue/svelte-legos';

	let {
		value = $bindable(),
		items = []
	}: {
		value?: string;
		items?: FaqItem[];
	} = $props();
</script>

<Accordion.Root
	bind:value
	class="flex flex-col gap-2"
	onValueChange={(v) => {
		if (!v) return;
		// NOTE: uncomment this if want to scroll to viewport
		// const id = ['answer', v].join('_');
		// setTimeout(() => {
		// 	const el = document.querySelector(`[itemid="${id}"]`);
		// 	if (el) {
		// 		el.scrollIntoView({ behavior: 'smooth' });
		// 	}
		// }, 100);
	}}
>
	{#each items as { question, answer, id }, i}
		{#key id}
			<Accordion.Item
				value={id}
				style={`--delay: ${i * 200}ms`}
				class={cn(
					'motion-translate-x-in-[-102%] motion-translate-y-in-[0%] border-0 focus-visible:outline-0',
					`motion-delay-[var(--delay)]`
				)}
			>
				<Accordion.Trigger
					class={cn(
						'border-primary rounded-md border p-3 hover:no-underline',
						'focus-visible:outline-0',
						{
							'rounded-b-none': value === id
						}
					)}>{question}</Accordion.Trigger
				>
				<Accordion.Content
					itemid={['answer', id].join('_')}
					class={cn('border-primary rounded-b-md border border-t-0 p-3 pb-0', {})}
					>{@html answer}</Accordion.Content
				>
			</Accordion.Item>
		{/key}
	{/each}
</Accordion.Root>
