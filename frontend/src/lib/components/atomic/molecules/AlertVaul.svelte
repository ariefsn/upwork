<script lang="ts">
	import { mediaQuery } from '@sveltelegos-blue/svelte-legos';
	import type { Snippet } from 'svelte';
	import Alert from '../atoms/Alert.svelte';
	import Vaul from '../atoms/Vaul.svelte';

	let {
		open = $bindable(true),
		onClose = $bindable(),
		forever = $bindable(false),
		content,
		closeOnOutsideClick
	}: {
		open?: boolean;
		content?: Snippet;
		onClose?: () => void;
		forever?: boolean;
		closeOnOutsideClick?: boolean;
	} = $props();

	const isLargeScreen = mediaQuery('(min-width: 48rem)');
</script>

{#if $isLargeScreen}
	<Alert bind:open {onClose}>
		{#snippet content()}
			{@render content?.()}
		{/snippet}
	</Alert>
{:else}
	{#key open}
		<Vaul bind:open {closeOnOutsideClick} {onClose}>
			{#snippet content()}
				{@render content?.()}
			{/snippet}
		</Vaul>
	{/key}
{/if}
