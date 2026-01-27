<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';

	const stats = [
		{ label: 'Active Projects', value: 8, change: '+2', trend: 'up', icon: 'lucide:folder-kanban' },
		{ label: 'Total Clients', value: 24, change: '+3', trend: 'up', icon: 'lucide:users' },
		{ label: 'Hours Tracked', value: '164h', change: '+12%', trend: 'up', icon: 'lucide:clock' },
		{
			label: 'Outstanding',
			value: '$12,450',
			change: '+$2,100',
			trend: 'up',
			icon: 'lucide:dollar-sign'
		}
	];

	const recentProjects = [
		{
			id: 1,
			name: 'Website Redesign',
			client: 'Acme Corp',
			progress: 75,
			dueDate: 'Feb 15',
			status: 'in-progress'
		},
		{
			id: 2,
			name: 'Mobile App Development',
			client: 'TechStart Inc',
			progress: 30,
			dueDate: 'Mar 1',
			status: 'in-progress'
		},
		{
			id: 3,
			name: 'Brand Identity',
			client: 'Green Leaf Co',
			progress: 90,
			dueDate: 'Jan 30',
			status: 'review'
		},
		{
			id: 4,
			name: 'Marketing Campaign',
			client: 'Global Dynamics',
			progress: 0,
			dueDate: 'Mar 15',
			status: 'planning'
		}
	];

	const recentActivity = [
		{
			id: 1,
			type: 'invoice_paid',
			title: 'Invoice #INV-001 paid',
			description: 'Acme Corp paid $2,500',
			time: '2 hours ago',
			icon: 'lucide:check-circle',
			color: 'text-emerald-500'
		},
		{
			id: 2,
			type: 'task_completed',
			title: 'Task completed',
			description: 'Homepage design - Website Redesign',
			time: '4 hours ago',
			icon: 'lucide:check-square',
			color: 'text-blue-500'
		},
		{
			id: 3,
			type: 'client_added',
			title: 'New client added',
			description: 'Green Leaf Co was added',
			time: 'Yesterday',
			icon: 'lucide:user-plus',
			color: 'text-violet-500'
		},
		{
			id: 4,
			type: 'project_created',
			title: 'New project created',
			description: 'Mobile App Development for TechStart',
			time: '2 days ago',
			icon: 'lucide:folder-plus',
			color: 'text-amber-500'
		},
		{
			id: 5,
			type: 'comment',
			title: 'New comment',
			description: 'Sarah commented on Brand Identity project',
			time: '3 days ago',
			icon: 'lucide:message-square',
			color: 'text-pink-500'
		}
	];

	const upcomingTasks = [
		{
			id: 1,
			title: 'Review logo concepts',
			project: 'Brand Identity',
			due: 'Today',
			priority: 'high'
		},
		{
			id: 2,
			title: 'Client meeting - Acme Corp',
			project: 'Website Redesign',
			due: 'Tomorrow',
			priority: 'high'
		},
		{ id: 3, title: 'Submit proposal', project: 'New Project', due: 'Jan 29', priority: 'medium' },
		{ id: 4, title: 'Update documentation', project: 'Mobile App', due: 'Jan 30', priority: 'low' }
	];

	const quickActions = [
		{
			label: 'New Project',
			icon: 'lucide:folder-plus',
			href: '/projects/new',
			color: 'bg-blue-500/10 text-blue-600'
		},
		{
			label: 'New Client',
			icon: 'lucide:user-plus',
			href: '/clients/new',
			color: 'bg-violet-500/10 text-violet-600'
		},
		{
			label: 'Create Invoice',
			icon: 'lucide:file-plus',
			href: '/invoices/new',
			color: 'bg-emerald-500/10 text-emerald-600'
		},
		{
			label: 'Track Time',
			icon: 'lucide:clock',
			href: '/time',
			color: 'bg-amber-500/10 text-amber-600'
		},
		{
			label: 'Send Email',
			icon: 'lucide:mail',
			href: '/emails/new',
			color: 'bg-pink-500/10 text-pink-600'
		},
		{
			label: 'New Template',
			icon: 'lucide:file-code',
			href: '/templates/new',
			color: 'bg-cyan-500/10 text-cyan-600'
		}
	];
</script>

<svelte:head>
	<title>Dashboard - Artemis</title>
</svelte:head>

<div class="min-h-[calc(100vh-3.5rem)] bg-background">
	<div class="border-b border-border bg-gradient-to-r from-primary/5 via-primary/3 to-transparent">
		<div class="px-6 py-8">
			<div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
				<div>
					<h1 class="text-2xl font-semibold tracking-tight">Good afternoon, Alex! 👋</h1>
					<p class="text-muted-foreground">Here's what's happening with your projects today.</p>
				</div>
				<div class="flex gap-2">
					<Button variant="outline" class="gap-2">
						<Icon icon="lucide:calendar" class="size-4" />
						<span class="hidden sm:inline">Schedule</span>
					</Button>
					<Button class="gap-2">
						<Icon icon="lucide:plus" class="size-4" />
						New Project
					</Button>
				</div>
			</div>
		</div>
	</div>

	<div class="p-6">
		<div class="mb-8 grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
			{#each stats as stat}
				<div
					class="rounded-xl border border-border bg-card p-4 transition-all hover:border-primary/20 hover:shadow-sm"
				>
					<div class="flex items-center justify-between">
						<div class="flex items-center gap-3">
							<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
								<Icon icon={stat.icon} class="size-5 text-primary" />
							</div>
							<div>
								<p class="text-sm text-muted-foreground">{stat.label}</p>
								<p class="text-2xl font-semibold">{stat.value}</p>
							</div>
						</div>
						<div class="flex items-center gap-1 text-sm text-emerald-600">
							<Icon icon="lucide:trending-up" class="size-4" />
							{stat.change}
						</div>
					</div>
				</div>
			{/each}
		</div>

		<div class="grid gap-6 lg:grid-cols-3">
			<div class="space-y-6 lg:col-span-2">
				<div class="rounded-xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Quick Actions</h3>
					<div class="grid grid-cols-3 gap-3 sm:grid-cols-6">
						{#each quickActions as action}
							<a
								href={action.href}
								class="group flex flex-col items-center gap-2 rounded-lg border border-border p-3 transition-all hover:border-primary/30 hover:bg-muted"
							>
								<div
									class="flex size-10 items-center justify-center rounded-lg {action.color} transition-transform group-hover:scale-110"
								>
									<Icon icon={action.icon} class="size-5" />
								</div>
								<span class="text-xs font-medium">{action.label}</span>
							</a>
						{/each}
					</div>
				</div>

				<div class="rounded-xl border border-border bg-card">
					<div class="flex items-center justify-between border-b border-border p-5">
						<h3 class="font-semibold">Active Projects</h3>
						<Button variant="ghost" size="sm" href="/projects">
							View all
							<Icon icon="lucide:arrow-right" class="ml-1 size-4" />
						</Button>
					</div>
					<div class="divide-y divide-border">
						{#each recentProjects as project}
							<div class="flex items-center gap-4 p-4 transition-colors hover:bg-muted/50">
								<div
									class="flex size-10 shrink-0 items-center justify-center rounded-lg bg-primary/10"
								>
									<Icon icon="lucide:folder" class="size-5 text-primary" />
								</div>
								<div class="min-w-0 flex-1">
									<div class="flex items-center gap-2">
										<p class="truncate font-medium">{project.name}</p>
										{#if project.status === 'review'}
											<span
												class="rounded-full bg-amber-100 px-2 py-0.5 text-xs font-medium text-amber-700 dark:bg-amber-900/30 dark:text-amber-400"
											>
												Review
											</span>
										{/if}
									</div>
									<p class="text-sm text-muted-foreground">
										{project.client} · Due {project.dueDate}
									</p>
								</div>
								<div class="hidden w-32 sm:block">
									<div class="mb-1 flex justify-between text-xs">
										<span class="text-muted-foreground">{project.progress}%</span>
									</div>
									<div class="h-1.5 w-full overflow-hidden rounded-full bg-muted">
										<div
											class="h-full rounded-full bg-primary transition-all"
											style="width: {project.progress}%"
										></div>
									</div>
								</div>
								<Button variant="ghost" size="icon" class="shrink-0">
									<Icon icon="lucide:more-vertical" class="size-4" />
								</Button>
							</div>
						{/each}
					</div>
				</div>

				<div class="rounded-xl border border-border bg-card">
					<div class="flex items-center justify-between border-b border-border p-5">
						<h3 class="font-semibold">Recent Activity</h3>
						<Button variant="ghost" size="sm">View all</Button>
					</div>
					<div class="divide-y divide-border">
						{#each recentActivity as activity}
							<div class="flex items-start gap-3 p-4 transition-colors hover:bg-muted/50">
								<div class="mt-0.5 shrink-0">
									<Icon icon={activity.icon} class="size-5 {activity.color}" />
								</div>
								<div class="min-w-0 flex-1">
									<p class="font-medium">{activity.title}</p>
									<p class="text-sm text-muted-foreground">{activity.description}</p>
								</div>
								<span class="shrink-0 text-xs text-muted-foreground">{activity.time}</span>
							</div>
						{/each}
					</div>
				</div>
			</div>

			<div class="space-y-6">
				<div class="rounded-xl border border-border bg-card p-5">
					<div class="flex items-center justify-between">
						<h3 class="font-semibold">Time Tracker</h3>
						<Button variant="ghost" size="icon">
							<Icon icon="lucide:history" class="size-4" />
						</Button>
					</div>
					<div class="mt-4 text-center">
						<div class="font-mono text-4xl font-semibold tracking-tight">02:34:12</div>
						<p class="mt-1 text-sm text-muted-foreground">Website Redesign</p>
					</div>
					<div class="mt-4 flex gap-2">
						<Button class="flex-1 gap-2">
							<Icon icon="lucide:pause" class="size-4" />
							Pause
						</Button>
						<Button variant="outline" class="shrink-0">
							<Icon icon="lucide:check" class="size-4" />
						</Button>
					</div>
				</div>

				<div class="rounded-xl border border-border bg-card">
					<div class="flex items-center justify-between border-b border-border p-5">
						<h3 class="font-semibold">Up Next</h3>
						<Button variant="ghost" size="sm" href="/tasks">View all</Button>
					</div>
					<div class="divide-y divide-border">
						{#each upcomingTasks as task}
							<div class="group flex items-start gap-3 p-4 transition-colors hover:bg-muted/50">
								<button
									class="mt-0.5 shrink-0 rounded border border-border p-1 hover:border-primary hover:bg-primary/10"
								>
									<Icon
										icon="lucide:check"
										class="size-3 text-transparent group-hover:text-primary"
									/>
								</button>
								<div class="min-w-0 flex-1">
									<p class="font-medium">{task.title}</p>
									<p class="text-sm text-muted-foreground">{task.project}</p>
								</div>
								<div class="shrink-0 text-right">
									<span
										class="text-xs {task.priority === 'high'
											? 'text-destructive'
											: 'text-muted-foreground'}">{task.due}</span
									>
								</div>
							</div>
						{/each}
					</div>
					<div class="border-t border-border p-4">
						<Button variant="outline" class="w-full gap-2">
							<Icon icon="lucide:plus" class="size-4" />
							Add Task
						</Button>
					</div>
				</div>

				<div class="rounded-xl border border-border bg-card p-5">
					<div class="flex items-center justify-between">
						<h3 class="font-semibold">Recent Clients</h3>
						<Button variant="ghost" size="sm" href="/clients">View all</Button>
					</div>
					<div class="mt-4 space-y-3">
						{#each ['Acme Corp', 'TechStart Inc', 'Green Leaf Co'] as client, i}
							<div class="flex items-center gap-3">
								<div
									class="flex size-9 items-center justify-center rounded-full bg-primary/10 text-sm font-medium text-primary"
								>
									{client.charAt(0)}
								</div>
								<div class="min-w-0 flex-1">
									<p class="truncate font-medium">{client}</p>
									<p class="text-xs text-muted-foreground">
										{['2 active projects', '1 active project', 'Brand identity'][i]}
									</p>
								</div>
							</div>
						{/each}
					</div>
				</div>

				<div
					class="rounded-xl border border-border bg-gradient-to-br from-primary/5 to-primary/10 p-5"
				>
					<div class="flex items-start gap-3">
						<div class="flex size-8 items-center justify-center rounded-full bg-primary/20">
							<Icon icon="lucide:lightbulb" class="size-4 text-primary" />
						</div>
						<div>
							<h4 class="font-medium">Pro Tip</h4>
							<p class="mt-1 text-sm text-muted-foreground">
								Use templates to save time on repetitive emails and proposals.
							</p>
							<Button variant="link" class="h-auto px-0 py-1 text-sm" href="/templates">
								Explore templates
							</Button>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
