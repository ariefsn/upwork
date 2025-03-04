<script lang="ts">
	import { page } from '$app/state';
	import Alert from '$lib/components/atomic/atoms/Alert.svelte';
	import Loading from '$lib/components/atomic/atoms/Loading.svelte';
	import AlertVaul from '$lib/components/atomic/molecules/AlertVaul.svelte';
	import Footer from '$lib/components/atomic/molecules/Footer.svelte';
	import DeleteAccount from '$lib/components/atomic/organisms/DeleteAccount.svelte';
	import UploadHistory from '$lib/components/atomic/organisms/UploadHistory.svelte';
	import { Toaster } from '$lib/components/ui/sonner';
	import { deleteAccountState, errorState, loadingState, uploadHistoryState } from '$lib/stores';
	import { ModeWatcher } from 'mode-watcher';
	import { deepMerge, MetaTags } from 'svelte-meta-tags';
	import '../app.css';

	let { children, data } = $props();

	let metatags = $derived.by(() => deepMerge({}, page.data.pageMetaTags));

	const gTagId = data.gTagId;
</script>

<svelte:head>
	<!-- Google tag (gtag.js) -->
	{@html `
		<script async src="https://www.googletagmanager.com/gtag/js?id=${gTagId}"></script>
		<script>
			window.dataLayer = window.dataLayer || [];
			function gtag() {
				dataLayer.push(arguments);
			}
			gtag('js', new Date());

			gtag('config', '${gTagId}');
		</script>
	`}
</svelte:head>

<MetaTags {...metatags} />

<ModeWatcher />

<Toaster />

<Alert
	class="z-[80]"
	open={$errorState.message !== ''}
	onClose={() => {
		errorState.set({ title: '', message: '' });
	}}
	title={$errorState.title}
	description={$errorState.message}
/>

{#if $loadingState}
	<Loading />
{/if}

{@render children()}

<AlertVaul
	open={$uploadHistoryState}
	closeOnOutsideClick={$errorState.message ? false : true}
	onClose={() => {
		uploadHistoryState.set(false);
	}}
>
	{#snippet content()}
		<UploadHistory faqLink={data?.faqLink} />
	{/snippet}
</AlertVaul>

<AlertVaul
	open={$deleteAccountState}
	closeOnOutsideClick={$errorState.message ? false : true}
	onClose={() => {
		deleteAccountState.set(false);
	}}
>
	{#snippet content()}
		<DeleteAccount />
	{/snippet}
</AlertVaul>

<Footer />
