<script lang="ts">
	import * as Table from '$lib/components/ui/table';
	import type { TableCellProps } from '$lib/models';

	let {
		headers = [],
		caption,
		rows = []
	}: {
		headers?: TableCellProps[];
		rows?: TableCellProps[][];
		caption?: string;
	} = $props();
</script>

<Table.Root>
	{#if caption}
		<Table.Caption>A list of your recent invoices.</Table.Caption>
	{/if}
	{#if headers.length}
		<Table.Header>
			<Table.Row>
				{#each headers as h}
					<Table.Head class={h.class}>{h.label}</Table.Head>
				{/each}
			</Table.Row>
		</Table.Header>
	{/if}
	<Table.Body>
		{#each rows as r}
			<Table.Row>
				{#each r as cell}
					<Table.Cell class={cell.class}>{cell.label}</Table.Cell>
				{/each}
			</Table.Row>
		{/each}
	</Table.Body>
</Table.Root>
