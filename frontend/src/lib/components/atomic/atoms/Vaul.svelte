<script lang="ts">
	import { cn } from '$lib/utils';
	import type { Snippet } from 'svelte';
	import { Drawer } from 'vaul-svelte';

	let {
		open = $bindable(false),
		class: classess,
		body,
		content,
		onClose,
		trigger,
		closeOnOutsideClick = $bindable(true)
	}: {
		open?: boolean;
		class?: string;
		body?: Snippet;
		content?: Snippet;
		trigger?: Snippet;
		onClose?: () => void;
		closeOnOutsideClick?: boolean;
	} = $props();
</script>

<div data-vaul-drawer-wrapper class={cn('min-h-[100vh] bg-white dark:bg-black/95', classess)}>
	{@render body?.()}
	<Drawer.Root
		bind:open
		shouldScaleBackground
		onOpenChange={(e) => !e && setTimeout(() => onClose?.(), 250)}
		{closeOnOutsideClick}
	>
		{#if trigger}
			<Drawer.Trigger asChild>
				{@render trigger?.()}
			</Drawer.Trigger>
		{/if}
		<Drawer.Portal>
			<Drawer.Overlay class="fixed inset-0 bg-black/40 dark:bg-black/10" />
			<Drawer.Content
				class="fixed bottom-0 left-0 right-0 z-[75] mt-24 flex h-[372px] rounded-t-[10px] bg-white outline-none dark:bg-zinc-800"
			>
				<div class="flex-1 rounded-t-[10px] bg-white p-4 dark:bg-black/90">
					<div class="mx-auto mb-8 h-1.5 w-12 flex-shrink-0 rounded-full bg-zinc-300"></div>
					{@render content?.()}
				</div>
			</Drawer.Content>
		</Drawer.Portal>
	</Drawer.Root>
</div>
