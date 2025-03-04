<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import type { Snippet } from 'svelte';

	let {
		trigger,
		content
	}: {
		trigger: Snippet;
		content?: Snippet;
	} = $props();
</script>

<Popover.Root portal={null}>
	<Popover.Trigger asChild let:builder>
		<Button
			builders={[builder]}
			variant="ghost"
			class="m-0 p-0 outline-none hover:bg-transparent focus-visible:ring-0 focus-visible:ring-offset-0"
		>
			{@render trigger()}
		</Button>
	</Popover.Trigger>
	<Popover.Content class="w-80">
		{@render content?.()}
	</Popover.Content>
</Popover.Root>
