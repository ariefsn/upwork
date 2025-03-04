<script lang="ts">
	import { formatCurrency } from '$lib';
	import type { UserData } from '$lib/graphql/generated';
	import { t } from '$lib/translations';
	import { cn } from '$lib/utils';
	import Badge from '../atoms/Badge.svelte';

	let {
		user,
		amount = 0,
		fee = 0,
		class: classess
	}: {
		user: UserData;
		amount?: number;
		fee?: number;
		class?: string;
	} = $props();
</script>

<div
	class={cn(
		'group relative flex flex-col items-start justify-between gap-2 pb-2 hover:cursor-pointer md:flex-row md:items-center',
		'motion-translate-x-in-[0%] motion-translate-y-in-[-200%]',
		classess
	)}
>
	<div class="absolute bottom-1 h-[2px] w-full bg-gray-100 sm:bottom-1 dark:bg-gray-800"></div>
	<div
		class="absolute bottom-1 h-[2px] w-0 bg-black duration-300 group-hover:w-full sm:bottom-1 dark:bg-white"
	></div>

	<a href={`/freelancers/${user.id}`} class="flex-1">
		<div class="text-primary text-sm font-bold md:text-xl">{user.fullName}</div>
		<div class="text-primary md:text-md mt-1 text-xs sm:text-sm">{user.title}</div>
	</a>

	<div
		class="mb-1 flex h-full flex-row items-center justify-between gap-2 md:mb-0 md:flex-col md:items-end"
	>
		<Badge href={user.url}>{$t('upwork.profile')}</Badge>
		<span class="text-md font-bold md:text-xl">
			{formatCurrency(amount - fee)}
		</span>
	</div>
</div>
