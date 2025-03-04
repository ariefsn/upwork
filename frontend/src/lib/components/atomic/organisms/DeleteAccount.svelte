<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import Button from '$lib/components/ui/button/button.svelte';
	import { userDelete, userResendDeleteToken } from '$lib/graphql';
	import { capitalize } from '$lib/helper';
	import { deleteAccountState, loadingState } from '$lib/stores';
	import { t } from '$lib/translations';
	import { toast } from 'svelte-sonner';
	import { z } from 'zod';
	import Input from '../molecules/Input.svelte';

	const schema = z.object({
		userID: z.string().min(1),
		code: z.string().min(1)
	});

	let userID = $state('');
	let code = $state('');
	let data = $derived({ userID, code });
	let errors: z.infer<typeof schema> = $state({ userID: '', code: '' });

	const isValid = $derived.by(() => {
		return schema.safeParse(data).success && errors?.userID === '' && errors?.code === '';
	});

	const onSubmit = async () => {
		const { error, success } = schema.safeParse(data);
		if (!success) {
			const errs = error?.flatten().fieldErrors;
			errors = {
				userID: errs?.userID?.[0] || '',
				code: errs?.code?.[0] || ''
			};

			return;
		}

		loadingState.set(true);

		const res = await userDelete({ id: userID, code });

		if (res.error) {
			const message = capitalize(
				res.error.graphQLErrors?.[0]?.message || $t('app.somethingWentWrong')
			);
			toast.error(message);
			loadingState.set(false);
			return;
		}

		toast.success($t('user.deleteForm.toast.delete.success.message'));

		userID = '';
		code = '';
		errors = { userID: '', code: '' };
		deleteAccountState.set(false);
		invalidateAll();
		loadingState.set(false);
	};

	const onResendCode = async () => {
		const res = await userResendDeleteToken(userID);

		if (res.error) {
			const message = capitalize(
				res.error.graphQLErrors?.[0]?.message || $t('app.somethingWentWrong')
			);
			errors = { userID: message, code: '' };
			loadingState.set(false);
			return;
		}

		toast.success($t('user.deleteForm.toast.resend.success.message'));

		code = '';
		errors = { userID: '', code: '' };
		loadingState.set(false);
	};
</script>

<div class="grid grid-cols-12 gap-4">
	<div class="col-span-12 flex flex-col gap-2">
		<Input
			label={$t('user.deleteForm.inputs.userID.label')}
			placeholder={$t('user.deleteForm.inputs.userID.placeholder')}
			bind:value={userID}
			error={errors.userID}
			on:input={() => (errors = { ...errors, userID: '' })}
		>
			{#snippet hint()}
				{#if userID}
					<Button
						size="sm"
						variant="ghost"
						class="px-0 hover:bg-transparent"
						onclick={onResendCode}
					>
						{$t('user.deleteForm.buttons.resend.title')}
					</Button>
				{/if}
			{/snippet}
		</Input>
		<Input
			label={$t('user.deleteForm.inputs.code.label')}
			placeholder={$t('user.deleteForm.inputs.code.placeholder')}
			bind:value={code}
			error={errors?.code}
		></Input>
	</div>
</div>

<hr class="my-2" />

{#if isValid}
	<p class="rounded-sm border border-red-500 p-2 text-sm text-red-500">
		{$t('user.deleteForm.buttons.submit.warn')}
	</p>
{/if}

<Button class="w-full" disabled={!isValid} onclick={onSubmit}
	>{$t('user.deleteForm.buttons.submit.title')}</Button
>
