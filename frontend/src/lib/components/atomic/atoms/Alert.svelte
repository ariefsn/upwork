<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog/index.js';
	import { cn } from '$lib/utils';
	import type { Snippet } from 'svelte';

	let {
		open = $bindable(true),
		title,
		description,
		content,
		footer,
		onClose,
		forever,
		class: classess
	}: {
		open?: boolean;
		title?: string;
		description?: string;
		content?: Snippet;
		footer?: Snippet;
		onClose?: () => void;
		forever?: boolean;
		class?: string;
	} = $props();
</script>

<Dialog.Root
	bind:open
	onOpenChange={(e) => !e && setTimeout(() => onClose?.(), 250)}
	closeOnEscape={!forever}
	closeOnOutsideClick={!forever}
>
	<Dialog.Content
		class={cn(
			'z-[77] sm:max-w-[425px]',
			'motion-scale-in-[0.5] motion-translate-x-in-[-37%] motion-translate-y-in-[68%] motion-opacity-in-[0%] motion-rotate-in-[-10deg] motion-blur-in-[5px] motion-duration-[0.35s] motion-duration-[0.53s]/scale motion-duration-[0.53s]/translate motion-duration-[0.63s]/rotate w-64',
			classess
		)}
	>
		{#if title || description}
			<Dialog.Header>
				{#if title}
					<Dialog.Title>{title}</Dialog.Title>
				{/if}
				{#if description}
					<Dialog.Description>
						{description}
					</Dialog.Description>
				{/if}
			</Dialog.Header>
		{/if}

		{#if content}
			{@render content()}
		{/if}

		{#if footer}
			<Dialog.Footer>
				{@render footer()}
			</Dialog.Footer>
		{/if}
	</Dialog.Content>
</Dialog.Root>
