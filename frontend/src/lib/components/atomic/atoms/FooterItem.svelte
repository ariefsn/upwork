<script lang="ts">
	export let href: string = '';
	export let title: string = '';
	export let onClick: ((e: MouseEvent) => void) | undefined = undefined;
</script>

{#if href}
	<a
		{href}
		class="flex items-center duration-300 hover:rotate-45"
		aria-label={`navigate to ${title}`}
	>
		<slot />
	</a>
{:else}
	<button class="duration-300 hover:rotate-45" onclick={onClick} aria-label={title}>
		<slot />
	</button>
{/if}
