<script lang="ts">
	import type { TableCellProps } from '$lib/models';
	import { cn } from '$lib/utils';
	import { tick } from 'svelte';
	import Dropzone from 'svelte-file-dropzone';

	let {
		onChange,
		placeholder = 'Drop your CSV file here, or click to select.'
	}: {
		onChange?: (value: { headers: TableCellProps[]; rows: TableCellProps[][]; file: File }) => void;
		placeholder?: string;
	} = $props();

	let files: File[] = $state([]);
	let fileData: string[][] = $state([]);

	const processRawCSV = (data: string) => {
		const output = [];
		const rows = data.split('\n');
		for (let i = 0; i < rows.length; i++) {
			const row = rows[i].replaceAll(', ', '#*|').replaceAll(',', '|||');
			const cells = row.split('|||').map((e) => e.replaceAll('#*|', ', ').replaceAll('"', ''));
			if (cells.length === 1 && !cells[0]) {
				continue;
			}
			output.push(cells.map((e) => e.replaceAll('\r', '')));
		}

		return output;
	};

	const headers = $derived(
		(fileData.length > 0 ? fileData[0] : []).map((e, i) => ({
			label: e,
			class: cn({
				'font-bold': true,
				'w-32': i === 0,
				'text-right': [9, 10, 11, 12].includes(i)
			})
		}))
	);
	const rows = $derived(
		(fileData.length > 1 ? fileData.slice(1) : []).map((e) =>
			e.map((c, i) => ({
				label: c,
				class: cn({
					'text-right': [9, 10, 11, 12].includes(i)
				})
			}))
		)
	);

	const handleFilesSelect = async (e: CustomEvent<{ acceptedFiles: File[] }>) => {
		files = e.detail.acceptedFiles;

		for (let i = 0; i < files.length; i++) {
			const reader = new FileReader();
			reader.onload = () => {
				const binaryStr = reader.result;
				fileData = processRawCSV(binaryStr?.toString() ?? '');
			};
			reader.readAsText(files[i]);
		}

		await tick();

		setTimeout(() => {
			if (files.length) {
				onChange?.({
					file: files[0],
					headers,
					rows
				});
			}
		}, 250);
	};
</script>

<Dropzone
	on:drop={handleFilesSelect}
	multiple={false}
	accept=".csv"
	disableDefaultStyles
	containerClasses="text-center py-8 border border-dashed outline-none rounded-md"
>
	{#if files.length}
		{files[0].name}
	{:else}
		{placeholder}
	{/if}
</Dropzone>
