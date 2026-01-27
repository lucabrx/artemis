<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import WelcomeWizard from '$lib/components/welcome-wizard.svelte';

	let showWizard = $state(false);

	const stats = [
		{
			label: 'Revenue',
			value: '$24,500',
			change: '+12.5%',
			icon: 'lucide:dollar-sign',
			color: 'bg-emerald-500/10 text-emerald-600'
		},
		{
			label: 'Active Projects',
			value: '12',
			change: '+3',
			icon: 'lucide:folder-kanban',
			color: 'bg-blue-500/10 text-blue-600'
		},
		{
			label: 'Hours Tracked',
			value: '164h',
			change: '+8.2%',
			icon: 'lucide:clock',
			color: 'bg-violet-500/10 text-violet-600'
		},
		{
			label: 'Invoices Paid',
			value: '18',
			change: '+5',
			icon: 'lucide:file-check',
			color: 'bg-amber-500/10 text-amber-600'
		}
	];

	const activities = [
		{
			id: 1,
			title: 'Invoice paid',
			desc: 'Acme Corp paid $2,500 for Website Redesign',
			time: '2 min ago',
			icon: 'lucide:check-circle',
			color: 'text-emerald-500',
			bg: 'bg-emerald-500/10'
		},
		{
			id: 2,
			title: 'Project completed',
			desc: 'Brand Identity project marked as done',
			time: '1 hour ago',
			icon: 'lucide:check-square',
			color: 'text-blue-500',
			bg: 'bg-blue-500/10'
		},
		{
			id: 3,
			title: 'New client',
			desc: 'TechStart Inc joined your workspace',
			time: '3 hours ago',
			icon: 'lucide:user-plus',
			color: 'text-violet-500',
			bg: 'bg-violet-500/10'
		},
		{
			id: 4,
			title: 'Task assigned',
			desc: 'Homepage design assigned to Sarah Chen',
			time: '5 hours ago',
			icon: 'lucide:plus',
			color: 'text-amber-500',
			bg: 'bg-amber-500/10'
		},
		{
			id: 5,
			title: 'New comment',
			desc: 'Mike commented on Mobile App project',
			time: 'Yesterday',
			icon: 'lucide:message-square',
			color: 'text-pink-500',
			bg: 'bg-pink-500/10'
		}
	];

	const projects = [
		{
			id: 1,
			name: 'Website Redesign',
			client: 'Acme Corp',
			progress: 75,
			tasks: { done: 24, total: 32 },
			due: 'Feb 15',
			color: 'bg-blue-500'
		},
		{
			id: 2,
			name: 'Mobile App Development',
			client: 'TechStart Inc',
			progress: 30,
			tasks: { done: 8, total: 45 },
			due: 'Mar 1',
			color: 'bg-violet-500'
		},
		{
			id: 3,
			name: 'Brand Identity',
			client: 'Green Leaf Co',
			progress: 90,
			tasks: { done: 18, total: 20 },
			due: 'Jan 30',
			color: 'bg-amber-500'
		},
		{
			id: 4,
			name: 'Marketing Campaign',
			client: 'Global Dynamics',
			progress: 10,
			tasks: { done: 2, total: 30 },
			due: 'Mar 15',
			color: 'bg-emerald-500'
		}
	];

	const weeklyData = [45, 52, 48, 65, 72, 58, 68];
	const weekDays = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
	const maxValue = Math.max(...weeklyData);
</script>

<svelte:head>
	<title>Dashboard - Artemis</title>
</svelte:head>

{#if showWizard}
	<WelcomeWizard onComplete={() => (showWizard = false)} />
{:else}
	<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
		<div class="px-6 py-8">
			<div class="mb-8 flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
				<div>
					<h1 class="text-3xl font-bold tracking-tight">Dashboard</h1>
					<p class="mt-1 text-muted-foreground">Welcome back! Here's what's happening today.</p>
				</div>
				<div class="flex gap-2">
					<Button variant="outline" class="gap-2">
						<Icon icon="lucide:calendar" class="size-4" />
						<span class="hidden sm:inline">This Week</span>
					</Button>
					<Button class="gap-2">
						<Icon icon="lucide:plus" class="size-4" />
						New Project
					</Button>
				</div>
			</div>

			<div class="mb-8 grid gap-4 md:grid-cols-2 lg:grid-cols-4">
				{#each stats as stat}
					<div
						class="rounded-2xl border border-border bg-card p-5 transition-shadow hover:shadow-lg"
					>
						<div class="flex items-start justify-between">
							<div class="flex items-center gap-3">
								<div class="flex size-10 items-center justify-center rounded-xl {stat.color}">
									<Icon icon={stat.icon} class="size-5" />
								</div>
								<div>
									<p class="text-sm text-muted-foreground">{stat.label}</p>
									<p class="text-2xl font-bold">{stat.value}</p>
								</div>
							</div>
							<div class="flex items-center gap-1 text-sm font-medium text-emerald-600">
								<Icon icon="lucide:trending-up" class="size-4" />
								{stat.change}
							</div>
						</div>
					</div>
				{/each}
			</div>

			<div class="grid gap-6 lg:grid-cols-3">
				<div class="space-y-6 lg:col-span-2">
					<div class="rounded-2xl border border-border bg-card p-6">
						<div class="mb-6 flex items-center justify-between">
							<h2 class="text-lg font-semibold">Activity Overview</h2>
							<div class="flex gap-2">
								{#each weekDays as day, i}
									<div class="text-center">
										<div class="mb-1 text-xs text-muted-foreground">{day}</div>
										<div
											class="relative w-8 overflow-hidden rounded-full bg-muted"
											style="height: 80px;"
										>
											<div
												class="absolute bottom-0 w-full rounded-full bg-primary transition-all"
												style="height: {(weeklyData[i] / maxValue) * 100}%"
											></div>
										</div>
										<div class="mt-1 text-xs font-medium">{weeklyData[i]}h</div>
									</div>
								{/each}
							</div>
						</div>
						<div class="flex items-center justify-between border-t border-border pt-4">
							<div>
								<p class="text-2xl font-bold">48h</p>
								<p class="text-sm text-muted-foreground">Total hours this week</p>
							</div>
							<div class="text-right">
								<p class="text-2xl font-bold">$3,240</p>
								<p class="text-sm text-muted-foreground">Billable amount</p>
							</div>
						</div>
					</div>

					<div class="overflow-hidden rounded-2xl border border-border bg-card">
						<div class="flex items-center justify-between border-b border-border p-6">
							<h2 class="text-lg font-semibold">Active Projects</h2>
							<Button variant="ghost" size="sm" class="gap-1">
								View all
								<Icon icon="lucide:arrow-right" class="size-4" />
							</Button>
						</div>
						<div class="divide-y divide-border">
							{#each projects as project}
								<div class="p-5 transition-colors hover:bg-muted/30">
									<div class="flex flex-col gap-4 sm:flex-row sm:items-center">
										<div class="min-w-0 flex-1">
											<div class="flex items-center gap-3">
												<div class="size-3 rounded-full {project.color}"></div>
												<h3 class="truncate font-semibold">{project.name}</h3>
											</div>
											<p class="mt-1 text-sm text-muted-foreground">
												{project.client} · Due {project.due}
											</p>
										</div>
										<div class="flex items-center gap-6">
											<div class="text-right">
												<p class="text-sm font-medium">{project.progress}%</p>
												<div class="mt-1 h-2 w-24 overflow-hidden rounded-full bg-muted">
													<div
														class="h-full {project.color} rounded-full transition-all"
														style="width: {project.progress}%"
													></div>
												</div>
											</div>
											<div class="hidden text-right sm:block">
												<p class="text-sm font-medium">
													{project.tasks.done}/{project.tasks.total}
												</p>
												<p class="text-xs text-muted-foreground">Tasks</p>
											</div>
										</div>
									</div>
								</div>
							{/each}
						</div>
					</div>
				</div>

				<div class="space-y-6">
					<div class="rounded-2xl border border-border bg-card p-6">
						<h2 class="mb-4 text-lg font-semibold">Time Tracker</h2>
						<div class="py-6 text-center">
							<div class="font-mono text-5xl font-bold tracking-tight">02:34:18</div>
							<p class="mt-2 text-muted-foreground">Website Redesign</p>
						</div>
						<div class="grid grid-cols-2 gap-2">
							<Button class="gap-2">
								<Icon icon="lucide:pause" class="size-4" />
								Pause
							</Button>
							<Button variant="outline">
								<Icon icon="lucide:check" class="size-4" />
							</Button>
						</div>
					</div>

					<div class="rounded-2xl border border-border bg-card p-6">
						<div class="mb-4 flex items-center justify-between">
							<h2 class="text-lg font-semibold">Recent Activity</h2>
							<Button variant="ghost" size="sm">View all</Button>
						</div>
						<div class="space-y-4">
							{#each activities as activity}
								<div class="flex gap-3">
									<div
										class="flex size-9 items-center justify-center rounded-lg {activity.bg} shrink-0"
									>
										<Icon icon={activity.icon} class="size-4 {activity.color}" />
									</div>
									<div class="min-w-0 flex-1">
										<p class="text-sm font-medium">{activity.title}</p>
										<p class="truncate text-sm text-muted-foreground">{activity.desc}</p>
										<p class="mt-1 text-xs text-muted-foreground">{activity.time}</p>
									</div>
								</div>
							{/each}
						</div>
					</div>

					<div
						class="rounded-2xl border border-border bg-gradient-to-br from-primary/5 to-primary/10 p-6"
					>
						<div class="flex items-start gap-3">
							<div class="flex size-10 items-center justify-center rounded-xl bg-primary/10">
								<Icon icon="lucide:sparkles" class="size-5 text-primary" />
							</div>
							<div>
								<h3 class="font-semibold">Pro Tip</h3>
								<p class="mt-1 text-sm text-muted-foreground">
									Use templates to create projects faster. Set up your first template now.
								</p>
								<Button variant="link" class="mt-2 h-auto px-0" href="/templates">
									Explore templates
									<Icon icon="lucide:arrow-right" class="ml-1 size-4" />
								</Button>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
{/if}
