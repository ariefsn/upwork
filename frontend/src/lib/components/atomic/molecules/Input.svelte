<script lang="ts">
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { cn } from '$lib/utils';
	import type { Snippet } from 'svelte';

	let {
		value = $bindable(),
		label,
		placeholder,
		type = 'text',
		class: classess = 'w-full',
		error = $bindable(),
		hint
	}: {
		label?: string;
		value?: string;
		placeholder?: string;
		type?: 'text' | 'email';
		class?: string;
		error?: string;
		hint?: Snippet;
	} = $props();
</script>

<div>
	{#if label}
		<Label for={label}>{label}</Label>
	{/if}
	<div class="relative">
		<Input
			{type}
			{placeholder}
			class={cn(classess, {
				'mt-1': label
			})}
			bind:value
			id={label}
			on:change
			on:input
		/>

		{#if hint}
			<div class="absolute inset-y-0 right-0 mr-3 flex items-center">
				{@render hint?.()}
			</div>
		{/if}
	</div>
	{#if error}
		<p class="ml-1 mt-1 text-xs text-red-500">{error}</p>
	{/if}
</div>
