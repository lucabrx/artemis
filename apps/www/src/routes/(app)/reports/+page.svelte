<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';

	let dateRange = $state('this-month');
	let reportType = $state('overview');

	const revenueData = [
		{ month: 'Aug', revenue: 18500, expenses: 8200, profit: 10300 },
		{ month: 'Sep', revenue: 22100, expenses: 9500, profit: 12600 },
		{ month: 'Oct', revenue: 19800, expenses: 8800, profit: 11000 },
		{ month: 'Nov', revenue: 25400, expenses: 10200, profit: 15200 },
		{ month: 'Dec', revenue: 28900, expenses: 11500, profit: 17400 },
		{ month: 'Jan', revenue: 24750, expenses: 9860, profit: 14890 }
	];

	const clientRevenue = [
		{ name: 'Acme Corp', revenue: 48500, hours: 320, color: 'bg-blue-500' },
		{ name: 'TechStart', revenue: 32100, hours: 210, color: 'bg-violet-500' },
		{ name: 'Green Leaf', revenue: 18600, hours: 145, color: 'bg-emerald-500' },
		{ name: 'Global Dynamics', revenue: 41200, hours: 285, color: 'bg-amber-500' },
		{ name: 'Others', revenue: 19500, hours: 165, color: 'bg-slate-400' }
	];

	const projectProfitability = [
		{ name: 'Website Redesign', revenue: 28500, costs: 12400, profit: 16100, margin: 56 },
		{ name: 'Mobile App', revenue: 34200, costs: 15800, profit: 18400, margin: 54 },
		{ name: 'Brand Identity', revenue: 12800, costs: 4200, profit: 8600, margin: 67 },
		{ name: 'Marketing Campaign', revenue: 18600, costs: 8900, profit: 9700, margin: 52 },
		{ name: 'E-commerce Platform', revenue: 45800, costs: 22100, profit: 23700, margin: 52 }
	];

	const timeByCategory = [
		{ category: 'Design', hours: 420, percentage: 35, color: 'bg-pink-500' },
		{ category: 'Development', hours: 380, percentage: 32, color: 'bg-blue-500' },
		{ category: 'Meetings', hours: 180, percentage: 15, color: 'bg-amber-500' },
		{ category: 'Research', hours: 120, percentage: 10, color: 'bg-violet-500' },
		{ category: 'Admin', hours: 96, percentage: 8, color: 'bg-slate-400' }
	];

	const upcomingMilestones = [
		{ project: 'Website Redesign', milestone: 'Launch Phase 1', date: 'Feb 15', status: 'on-track' },
		{ project: 'Mobile App', milestone: 'Beta Release', date: 'Feb 28', status: 'at-risk' },
		{ project: 'Brand Identity', milestone: 'Final Delivery', date: 'Jan 30', status: 'on-track' }
	];

	const recentActivity = [
		{ type: 'invoice', text: 'Invoice INV-2026-001 paid by Acme Corp', amount: 13750, time: '2 hours ago' },
		{ type: 'time', text: '160 hours logged this week', amount: null, time: 'Yesterday' },
		{ type: 'expense', text: 'New expense: Figma subscription', amount: 144, time: 'Yesterday' },
		{ type: 'project', text: 'Brand Identity project completed', amount: null, time: '2 days ago' }
	];

	const totalRevenue = revenueData.reduce((acc, m) => acc + m.revenue, 0);
	const totalExpenses = revenueData.reduce((acc, m) => acc + m.expenses, 0);
	const totalProfit = revenueData.reduce((acc, m) => acc + m.profit, 0);
	const avgMargin = Math.round((totalProfit / totalRevenue) * 100);

	const maxRevenue = Math.max(...revenueData.map(d => d.revenue));

	function formatCurrency(amount: number): string {
		return new Intl.NumberFormat('en-US', {
			style: 'currency',
			currency: 'USD',
			minimumFractionDigits: 0,
			maximumFractionDigits: 0
		}).format(amount);
	}

	function getTrendIcon(trend: 'up' | 'down' | 'neutral') {
		switch (trend) {
			case 'up':
				return 'hugeicons:trend-up';
			case 'down':
				return 'hugeicons:trend-down';
			default:
				return 'hugeicons:minus-sign';
		}
	}

	const dateRanges = [
		{ id: 'this-week', label: 'This Week' },
		{ id: 'last-week', label: 'Last Week' },
		{ id: 'this-month', label: 'This Month' },
		{ id: 'last-month', label: 'Last Month' },
		{ id: 'this-quarter', label: 'This Quarter' },
		{ id: 'this-year', label: 'This Year' },
		{ id: 'custom', label: 'Custom Range' }
	];

	const reportTypes = [
		{ id: 'overview', label: 'Overview', icon: 'hugeicons:dashboard' },
		{ id: 'financial', label: 'Financial', icon: 'hugeicons:chart' },
		{ id: 'time', label: 'Time Reports', icon: 'hugeicons:clock-01' },
		{ id: 'projects', label: 'Projects', icon: 'hugeicons:folder-01' },
		{ id: 'clients', label: 'Clients', icon: 'hugeicons:users-01' }
	];
</script>

<svelte:head>
	<title>Reports - Artemis</title>
</svelte:head>

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="px-6 py-4">
		<div class="mb-6 flex items-center justify-end gap-2">
			<select
				bind:value={dateRange}
				class="h-10 rounded-md border border-input bg-background px-3 text-sm"
			>
				{#each dateRanges as range}
					<option value={range.id}>{range.label}</option>
				{/each}
			</select>
			<Button variant="outline" class="gap-2">
				<Icon icon="hugeicons:download-01" class="size-4" />
				Export
			</Button>
		</div>

		<div class="mb-8 flex flex-wrap gap-2">
			{#each reportTypes as type}
				<button
					class="flex items-center gap-2 rounded-lg px-4 py-2 text-sm font-medium transition-colors {reportType === type.id ? 'bg-primary text-primary-foreground' : 'border border-border text-muted-foreground hover:text-foreground'}"
					onclick={() => reportType = type.id}
				>
					<Icon icon={type.icon} class="size-4" />
					{type.label}
				</button>
			{/each}
		</div>

		<div class="grid gap-4 md:grid-cols-4">
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">Total Revenue</p>
						<p class="text-2xl font-bold">{formatCurrency(totalRevenue)}</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-emerald-500/10">
						<Icon icon="hugeicons:dollar-circle" class="size-5 text-emerald-600" />
					</div>
				</div>
				<div class="mt-2 flex items-center gap-1 text-sm text-emerald-600">
					<Icon icon="hugeicons:trend-up" class="size-4" />
					<span>+12.5%</span>
					<span class="text-muted-foreground">vs last period</span>
				</div>
			</div>
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">Total Expenses</p>
						<p class="text-2xl font-bold">{formatCurrency(totalExpenses)}</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-rose-500/10">
						<Icon icon="hugeicons:receipt" class="size-5 text-rose-600" />
					</div>
				</div>
				<div class="mt-2 flex items-center gap-1 text-sm text-rose-600">
					<Icon icon="hugeicons:trend-up" class="size-4" />
					<span>+8.2%</span>
					<span class="text-muted-foreground">vs last period</span>
				</div>
			</div>
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">Net Profit</p>
						<p class="text-2xl font-bold">{formatCurrency(totalProfit)}</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-blue-500/10">
						<Icon icon="hugeicons:piggy-bank" class="size-5 text-blue-600" />
					</div>
				</div>
				<div class="mt-2 flex items-center gap-1 text-sm text-blue-600">
					<Icon icon="hugeicons:trend-up" class="size-4" />
					<span>+15.3%</span>
					<span class="text-muted-foreground">vs last period</span>
				</div>
			</div>
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">Profit Margin</p>
						<p class="text-2xl font-bold">{avgMargin}%</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-violet-500/10">
						<Icon icon="hugeicons:percent-circle" class="size-5 text-violet-600" />
					</div>
				</div>
				<div class="mt-2 flex items-center gap-1 text-sm text-violet-600">
					<Icon icon="hugeicons:trend-up" class="size-4" />
					<span>+2.1%</span>
					<span class="text-muted-foreground">vs last period</span>
				</div>
			</div>
		</div>

		<div class="mt-8 grid gap-6 lg:grid-cols-3">
			<div class="space-y-6 lg:col-span-2">
				<div class="rounded-2xl border border-border bg-card p-6">
					<div class="mb-6 flex items-center justify-between">
						<div>
							<h2 class="text-lg font-semibold">Revenue Overview</h2>
							<p class="text-sm text-muted-foreground">Revenue, expenses, and profit over time</p>
						</div>
						<div class="flex items-center gap-4">
							<div class="flex items-center gap-2">
								<div class="size-3 rounded-full bg-emerald-500"></div>
								<span class="text-xs text-muted-foreground">Revenue</span>
							</div>
							<div class="flex items-center gap-2">
								<div class="size-3 rounded-full bg-rose-400"></div>
								<span class="text-xs text-muted-foreground">Expenses</span>
							</div>
							<div class="flex items-center gap-2">
								<div class="size-3 rounded-full bg-blue-500"></div>
								<span class="text-xs text-muted-foreground">Profit</span>
							</div>
						</div>
					</div>

					<div class="flex h-64 items-end gap-2">
						{#each revenueData as month}
							<div class="flex flex-1 flex-col gap-1">
								<div class="relative flex flex-col gap-0.5">
									<div 
										class="rounded-t bg-emerald-500 transition-all hover:bg-emerald-600"
										style="height: {(month.revenue / maxRevenue) * 150}px"
									></div>
									<div 
										class="rounded-t bg-rose-400 transition-all hover:bg-rose-500"
										style="height: {(month.expenses / maxRevenue) * 150}px"
									></div>
								</div>
								<span class="text-center text-xs text-muted-foreground">{month.month}</span>
							</div>
						{/each}
					</div>
				</div>

				<div class="grid gap-6 md:grid-cols-2">
					<div class="rounded-2xl border border-border bg-card p-6">
						<h2 class="mb-4 text-lg font-semibold">Revenue by Client</h2>
						<div class="space-y-4">
							{#each clientRevenue as client}
								<div>
									<div class="mb-1 flex items-center justify-between">
										<span class="text-sm font-medium">{client.name}</span>
										<span class="text-sm">{formatCurrency(client.revenue)}</span>
									</div>
									<div class="h-2 overflow-hidden rounded-full bg-muted">
										<div class="h-full rounded-full {client.color}" style="width: {(client.revenue / 50000) * 100}%"></div>
									</div>
									<p class="mt-1 text-xs text-muted-foreground">{client.hours} hours · ${Math.round(client.revenue / client.hours)}/hr avg</p>
								</div>
							{/each}
						</div>
					</div>

					<div class="rounded-2xl border border-border bg-card p-6">
						<h2 class="mb-4 text-lg font-semibold">Time by Category</h2>
						<div class="space-y-3">
							{#each timeByCategory as cat}
								<div class="flex items-center gap-3">
									<div class="flex size-8 items-center justify-center rounded-lg {cat.color}">
										<Icon icon="hugeicons:clock-01" class="size-4 text-white" />
									</div>
									<div class="flex-1">
										<div class="flex items-center justify-between">
											<span class="text-sm font-medium">{cat.category}</span>
											<span class="text-sm">{cat.hours}h</span>
										</div>
										<div class="h-1.5 overflow-hidden rounded-full bg-muted">
											<div class="h-full rounded-full {cat.color}" style="width: {cat.percentage}%"></div>
										</div>
									</div>
									<span class="w-10 text-right text-xs text-muted-foreground">{cat.percentage}%</span>
								</div>
							{/each}
						</div>
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-6">
					<div class="mb-4 flex items-center justify-between">
						<h2 class="text-lg font-semibold">Project Profitability</h2>
						<Button variant="ghost" size="sm" class="gap-1">
							View All
							<Icon icon="hugeicons:arrow-right-01" class="size-4" />
						</Button>
					</div>
					<div class="overflow-x-auto">
						<table class="w-full">
							<thead>
								<tr class="border-b text-left text-sm text-muted-foreground">
									<th class="pb-3 font-medium">Project</th>
									<th class="pb-3 text-right font-medium">Revenue</th>
									<th class="pb-3 text-right font-medium">Costs</th>
									<th class="pb-3 text-right font-medium">Profit</th>
									<th class="pb-3 text-right font-medium">Margin</th>
								</tr>
							</thead>
							<tbody class="text-sm">
								{#each projectProfitability as project}
									<tr class="border-b">
										<td class="py-3 font-medium">{project.name}</td>
										<td class="py-3 text-right">{formatCurrency(project.revenue)}</td>
										<td class="py-3 text-right">{formatCurrency(project.costs)}</td>
										<td class="py-3 text-right font-medium text-emerald-600">{formatCurrency(project.profit)}</td>
										<td class="py-3 text-right">
											<span class="rounded-full bg-emerald-500/10 px-2 py-0.5 text-xs font-medium text-emerald-600">
												{project.margin}%
											</span>
										</td>
									</tr>
								{/each}
							</tbody>
						</table>
					</div>
				</div>
			</div>

			<div class="space-y-6">
				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Quick Stats</h3>
					<div class="space-y-4">
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Total Hours</span>
							<span class="font-semibold">1,196h</span>
						</div>
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Billable Hours</span>
							<span class="font-semibold">982h (82%)</span>
						</div>
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Avg Hourly Rate</span>
							<span class="font-semibold">$133</span>
						</div>
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Active Projects</span>
							<span class="font-semibold">12</span>
						</div>
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Invoices Sent</span>
							<span class="font-semibold">28</span>
						</div>
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Upcoming Milestones</h3>
					<div class="space-y-3">
						{#each upcomingMilestones as milestone}
							<div class="rounded-lg bg-muted/50 p-3">
								<div class="mb-1 flex items-center justify-between">
									<span class="text-sm font-medium">{milestone.milestone}</span>
									<span class="text-xs {milestone.status === 'on-track' ? 'text-emerald-600' : 'text-amber-600'}">
										{milestone.status === 'on-track' ? 'On Track' : 'At Risk'}
									</span>
								</div>
								<p class="text-xs text-muted-foreground">{milestone.project}</p>
								<p class="mt-1 text-xs text-muted-foreground">Due {milestone.date}</p>
							</div>
						{/each}
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Recent Activity</h3>
					<div class="space-y-3">
						{#each recentActivity as activity}
							<div class="flex gap-3">
								<div class="flex size-8 shrink-0 items-center justify-center rounded-lg bg-muted">
									{#if activity.type === 'invoice'}
										<Icon icon="hugeicons:invoice-01" class="size-4 text-emerald-500" />
									{:else if activity.type === 'time'}
										<Icon icon="hugeicons:clock-01" class="size-4 text-blue-500" />
									{:else if activity.type === 'expense'}
										<Icon icon="hugeicons:receipt" class="size-4 text-rose-500" />
									{:else}
										<Icon icon="hugeicons:checkmark-circle-01" class="size-4 text-violet-500" />
									{/if}
								</div>
								<div class="flex-1">
									<p class="text-sm">{activity.text}</p>
									{#if activity.amount}
										<p class="text-sm font-medium">{formatCurrency(activity.amount)}</p>
									{/if}
									<p class="text-xs text-muted-foreground">{activity.time}</p>
								</div>
							</div>
						{/each}
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-gradient-to-br from-primary/5 to-primary/10 p-5">
					<div class="flex items-start gap-3">
						<div class="flex size-10 items-center justify-center rounded-xl bg-primary/10">
							<Icon icon="hugeicons:file-export" class="size-5 text-primary" />
						</div>
						<div>
							<h3 class="font-semibold">Need a custom report?</h3>
							<p class="mt-1 text-sm text-muted-foreground">
								Generate detailed PDF reports with custom date ranges and filters.
							</p>
							<Button variant="link" class="mt-2 h-auto px-0">
								Create Custom Report
								<Icon icon="hugeicons:arrow-right-01" class="ml-1 size-4" />
							</Button>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
