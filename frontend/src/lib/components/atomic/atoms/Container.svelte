<script lang="ts">
	import { cn } from '$lib/utils';
	import type { Snippet } from 'svelte';

	let {
		class: classess,
		children
	}: {
		class?: string;
		children: Snippet;
	} = $props();
</script>

<div class={cn('w-5/6 md:w-4/5 lg:w-3/6 xl:w-3/6 2xl:w-3/6', classess)}>
	{@render children()}
</div>
