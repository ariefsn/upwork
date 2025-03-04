<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { page } from '$app/state';
	import Button from '$lib/components/ui/button/button.svelte';
	import { earningsUpload } from '$lib/graphql';
	import { capitalize } from '$lib/helper';
	import type { TableCellProps } from '$lib/models';
	import { loadingState, uploadHistoryState } from '$lib/stores';
	import { t } from '$lib/translations';
	import { Info } from 'lucide-svelte';
	import { toast } from 'svelte-sonner';
	import { z } from 'zod';
	import Popup from '../atoms/Popup.svelte';
	import Table from '../atoms/Table.svelte';
	import DropCsv from '../molecules/DropCsv.svelte';
	import Input from '../molecules/Input.svelte';

	const schema = z.object({
		userID: z.string(),
		email: z.string().email()
	});

	let {
		preview,
		faqLink
	}: {
		preview?: boolean;
		faqLink: string;
	} = $props();

	page;

	let userID = $state('');
	let email = $state('');
	let data = $derived({ userID, email });
	let errors: z.infer<typeof schema> = $state({ userID: '', email: '' });
	let headers: TableCellProps[] = $state([]);
	let rows: TableCellProps[][] = $state([]);
	let file: File | null = $state(null);

	const isValid = $derived.by(() => {
		return schema.safeParse(data).success && !errors?.userID && !errors?.email && file;
	});

	const onSubmit = async () => {
		const { error, success } = schema.safeParse(data);
		if (!success) {
			const errs = error?.flatten().fieldErrors;
			errors = {
				userID: errs?.userID?.[0] || '',
				email: errs?.email?.[0] || ''
			};

			return;
		}

		loadingState.set(true);

		const res = await earningsUpload({
			email,
			userID,
			file
		});

		if (res.error) {
			const message = capitalize(
				res.error.graphQLErrors?.[0]?.message || $t('app.somethingWentWrong')
			);
			toast.error(message);
			loadingState.set(false);
			return;
		}

		toast.success($t('upwork.uploadForm.toast.upload.success.message'));

		userID = '';
		email = '';
		errors = { userID: '', email: '' };
		uploadHistoryState.set(false);
		invalidateAll();
		loadingState.set(false);
	};
</script>

<div class="grid grid-cols-12 gap-4">
	<div class="col-span-12 flex flex-col gap-2">
		<Input
			label={$t('upwork.uploadForm.inputs.userID.label')}
			placeholder={$t('upwork.uploadForm.inputs.userID.placeholder')}
			bind:value={userID}
			error={errors?.userID}
			on:input={() => (errors = { ...errors, userID: '' })}
		>
			{#snippet hint()}
				<Popup>
					{#snippet trigger()}
						<Info />
					{/snippet}

					{#snippet content()}
						{@html $t('upwork.uploadForm.inputs.userID.hint', { faq: faqLink })}
					{/snippet}
				</Popup>
			{/snippet}
		</Input>
		<Input
			label={$t('upwork.uploadForm.inputs.email.label')}
			placeholder={$t('upwork.uploadForm.inputs.email.placeholder')}
			bind:value={email}
			error={errors?.email}
			on:input={() => (errors = { ...errors, email: '' })}
		>
			{#snippet hint()}
				<Popup>
					{#snippet trigger()}
						<Info />
					{/snippet}

					{#snippet content()}
						{$t('upwork.uploadForm.inputs.email.hint')}
					{/snippet}
				</Popup>
			{/snippet}
		</Input>
	</div>
	<div class="col-span-12">
		<DropCsv
			placeholder={$t('upwork.uploadForm.inputs.file.placeholder')}
			onChange={(value) => {
				headers = [...value.headers];
				rows = [...value.rows];
				file = value.file;
			}}
		/>
	</div>
</div>

<hr class="my-2" />

<Button class="w-full" disabled={!isValid} onclick={onSubmit}
	>{$t('upwork.uploadForm.buttons.submit.title')}</Button
>

{#if preview}
	<div class="max-h-96 overflow-auto">
		<Table {headers} {rows} />
	</div>
{/if}
