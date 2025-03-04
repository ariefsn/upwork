<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import { formatCurrency, formatDate } from '$lib';
	import Container from '$lib/components/atomic/atoms/Container.svelte';
	import Dropdown from '$lib/components/atomic/atoms/Dropdown.svelte';
	import Freelancer from '$lib/components/atomic/molecules/Freelancer.svelte';
	import { EarningType, type EarningsDataMonthly } from '$lib/graphql/generated';
	import { t } from '$lib/translations';
	import { cn } from '$lib/utils';
	import { scaleBand } from 'd3-scale';
	import { Axis, Bars, Chart, Highlight, Labels, LinearGradient, Svg, Tooltip } from 'layerchart';
	import { mode } from 'mode-watcher';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();

	const total = $derived.by(() =>
		data.items.reduce(
			(prev, curr) => ({
				amount: prev.amount + curr.totalAmount,
				fee: prev.fee + curr.totalFee
			}),
			{ amount: 0, fee: 0 }
		)
	);

	const isDark = $derived($mode === 'dark');

	const getMonthlyDetails = (d: EarningsDataMonthly) => {
		const items = d.items;

		return items.reduce(
			(prev, curr) => {
				const isHourly = curr.type === EarningType.Hourly;

				return {
					fixed: {
						amount: prev.fixed.amount + (!isHourly ? curr.amount : 0),
						fee: prev.fixed.fee + (!isHourly ? curr.fee : 0),
						total:
							prev.fixed.amount +
							(!isHourly ? curr.amount : 0) -
							(prev.fixed.fee + (!isHourly ? curr.fee : 0))
					},
					hourly: {
						amount: prev.hourly.amount + (isHourly ? curr.amount : 0),
						fee: prev.hourly.fee + (isHourly ? curr.fee : 0),
						total:
							prev.hourly.amount +
							(isHourly ? curr.amount : 0) -
							(prev.hourly.fee + (isHourly ? curr.fee : 0))
					}
				};
			},
			{
				fixed: { amount: 0, fee: 0, total: 0 },
				hourly: { amount: 0, fee: 0, total: 0 }
			}
		);
	};
</script>

<div class="flex h-screen flex-col items-center justify-center">
	<Container>
		<Freelancer user={data.user} amount={total.amount} fee={total.fee} class="mb-4 !pb-5" />
		<div class="mb-4 w-full text-right">
			<Dropdown
				searchPlaceholder={$t('app.year', { yearCount: data.years.length })}
				value={data.year + ''}
				items={data.years.map((e) => ({ value: e + '', label: e + '' }))}
				onChange={(val) => {
					const thisYear = new Date().getFullYear() === +val;
					let url = page.url.pathname;
					if (!thisYear) {
						url += `?year=${val}`;
					}
					goto(url);
				}}
			/>
		</div>
	</Container>

	<Container
		class={cn(
			'h-3/6',
			'motion-scale-in-[0.5] motion-translate-x-in-[103%] motion-translate-y-in-[-38%] motion-opacity-in-[0%] motion-rotate-in-[-10deg] motion-blur-in-[5px] motion-duration-[0.35s] motion-duration-[0.53s]/scale motion-duration-[0.53s]/translate motion-duration-[0.63s]/rotate'
		)}
	>
		{#key [isDark, data.year]}
			<Chart
				data={data.items}
				x={(d: EarningsDataMonthly) => formatDate(new Date(d.year, d.month - 1), 'MMM YY')}
				xScale={scaleBand().padding(0.4)}
				y={(d: EarningsDataMonthly) => d.totalAmount - d.totalFee}
				yDomain={[0, null]}
				yNice={4}
				padding={{ left: 16, bottom: 24 }}
				tooltip={{ mode: 'band' }}
				height={100}
			>
				<Svg>
					<Axis
						placement="left"
						grid={{
							class: 'stroke-primary/10'
						}}
						rule
						classes={{
							root: 'fill-black dark:fill-white'
						}}
					/>
					<Axis
						placement="bottom"
						rule={{ class: 'fill-black dark:fill-white' }}
						grid
						tickLabelProps={{
							rotate: 315,
							textAnchor: 'end',
							class: 'fill-black dark:fill-white font-semibold'
						}}
					/>
					<LinearGradient class="from-[#0061ff] to-[#60efff]" vertical let:gradient>
						<Bars strokeWidth={1} fill={gradient} class="stroke-primary" />
					</LinearGradient>

					<LinearGradient class="from-[#0061ff] to-[#60efff]" vertical let:gradient>
						<Highlight
							area={{
								class: 'fill-primary opacity-10'
							}}
							bar={{
								fill: gradient,
								stroke: isDark ? 'white' : 'black',
								strokeWidth: 1
							}}
						/>
					</LinearGradient>
					<Labels format={formatCurrency} class="fill-black dark:fill-white" />
				</Svg>
				<Tooltip.Root let:data>
					<Tooltip.Header>
						<div>
							{formatDate(new Date(data.year, data.month - 1), 'MMM YYYY')}
							<div class="mt-1 text-left">
								{formatCurrency(data.totalAmount - data.totalFee)}
							</div>
						</div>
					</Tooltip.Header>
					<Tooltip.List>
						<Tooltip.Item
							label="Fixed"
							value={formatCurrency(getMonthlyDetails(data).fixed.total)}
						/>
						<Tooltip.Item
							label="Hourly"
							value={formatCurrency(getMonthlyDetails(data).hourly.total)}
						/>
					</Tooltip.List>
				</Tooltip.Root>
			</Chart>
		{/key}
	</Container>
</div>
