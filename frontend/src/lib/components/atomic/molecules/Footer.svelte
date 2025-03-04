<script lang="ts">
	import { deleteAccountState, uploadHistoryState } from '$lib/stores';
	import { cn } from '$lib/utils';
	import { FilePlus2, Home, MessageCircleQuestion, Moon, Sun, Trash2 } from 'lucide-svelte';
	import { mode, toggleMode } from 'mode-watcher';
	import FooterItem from '../atoms/FooterItem.svelte';
	import LanguageSelect from '../atoms/LanguageSelect.svelte';
</script>

<footer
	class={cn(
		'fixed bottom-0 z-50 m-auto flex w-full items-center justify-center',
		'motion-scale-in-[0.5] motion-translate-x-in-[0%] motion-translate-y-in-[102%] motion-opacity-in-[0%] motion-rotate-in-[-1deg] motion-blur-in-[5px] motion-duration-[0.35s] motion-duration-[0.53s]/scale motion-duration-[0.53s]/translate motion-duration-[0.63s]/rotate'
	)}
>
	<div
		class="mb-3 flex items-center gap-4 rounded-md p-3 shadow-md dark:shadow-md dark:shadow-gray-100"
	>
		<FooterItem href="/" title="Home">
			<Home />
		</FooterItem>
		<FooterItem
			title="Upload History"
			onClick={() => {
				uploadHistoryState.set(true);
			}}
		>
			<FilePlus2 />
		</FooterItem>
		<FooterItem href="/faqs" title="FAQs">
			<MessageCircleQuestion />
		</FooterItem>
		<FooterItem onClick={toggleMode} title="Toggle Theme">
			{#if $mode == 'dark'}
				<Sun />
			{:else}
				<Moon />
			{/if}
		</FooterItem>
		<FooterItem
			title="Delete Account"
			onClick={() => {
				deleteAccountState.set(true);
			}}
		>
			<Trash2 />
		</FooterItem>
		<LanguageSelect />
	</div>
</footer>
