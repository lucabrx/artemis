<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';
	import { Label } from '$lib/components/ui/label/index.js';

	let viewMode = $state<'board' | 'list' | 'kanban'>('board');
	let searchQuery = $state('');
	let statusFilter = $state('all');
	let showNewProjectModal = $state(false);

	interface Task {
		id: string;
		title: string;
		status: 'todo' | 'in-progress' | 'review' | 'done';
		assignee?: string;
		dueDate?: string;
		priority: 'low' | 'medium' | 'high';
	}

	interface Project {
		id: string;
		name: string;
		description: string;
		client: string;
		clientAvatar: string;
		status: 'planning' | 'in-progress' | 'review' | 'completed' | 'on-hold';
		progress: number;
		priority: 'low' | 'medium' | 'high';
		startDate: string;
		dueDate: string;
		budget?: number;
		spent?: number;
		tasks: Task[];
		team: { name: string; avatar: string }[];
		tags: string[];
		color: string;
	}

	const projects: Project[] = [
		{
			id: '1',
			name: 'Website Redesign',
			description: 'Complete overhaul of the company website with modern design and improved UX.',
			client: 'Acme Corporation',
			clientAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Acme',
			status: 'in-progress',
			progress: 75,
			priority: 'high',
			startDate: 'Jan 1, 2026',
			dueDate: 'Feb 15, 2026',
			budget: 25000,
			spent: 18750,
			tasks: [
				{ id: 't1', title: 'Design system setup', status: 'done', assignee: 'Sarah Chen', priority: 'high' },
				{ id: 't2', title: 'Homepage mockups', status: 'done', assignee: 'Sarah Chen', priority: 'high' },
				{ id: 't3', title: 'About page', status: 'in-progress', assignee: 'Mike Ross', priority: 'medium' },
				{ id: 't4', title: 'Contact form', status: 'todo', assignee: 'Mike Ross', priority: 'medium' },
				{ id: 't5', title: 'Mobile responsiveness', status: 'review', assignee: 'Sarah Chen', priority: 'high' }
			],
			team: [
				{ name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah' },
				{ name: 'Mike Ross', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Mike' }
			],
			tags: ['Web Design', 'UI/UX'],
			color: 'bg-blue-500'
		},
		{
			id: '2',
			name: 'Mobile App Development',
			description: 'Native iOS and Android app for customer engagement and loyalty program.',
			client: 'TechStart Inc',
			clientAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=TechStart',
			status: 'in-progress',
			progress: 30,
			priority: 'high',
			startDate: 'Jan 15, 2026',
			dueDate: 'Mar 1, 2026',
			budget: 45000,
			spent: 13500,
			tasks: [
				{ id: 't6', title: 'Requirements gathering', status: 'done', assignee: 'Alex Morgan', priority: 'high' },
				{ id: 't7', title: 'Wireframes', status: 'done', assignee: 'Alex Morgan', priority: 'high' },
				{ id: 't8', title: 'API design', status: 'in-progress', assignee: 'David Kim', priority: 'high' },
				{ id: 't9', title: 'iOS prototype', status: 'todo', assignee: 'Lisa Wang', priority: 'medium' },
				{ id: 't10', title: 'Android prototype', status: 'todo', assignee: 'Tom Johnson', priority: 'medium' }
			],
			team: [
				{ name: 'Alex Morgan', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex' },
				{ name: 'David Kim', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=David' },
				{ name: 'Lisa Wang', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Lisa' }
			],
			tags: ['Mobile', 'iOS', 'Android'],
			color: 'bg-violet-500'
		},
		{
			id: '3',
			name: 'Brand Identity',
			description: 'Complete brand refresh including logo, colors, typography, and brand guidelines.',
			client: 'Green Leaf Co',
			clientAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=GreenLeaf',
			status: 'review',
			progress: 90,
			priority: 'medium',
			startDate: 'Jan 5, 2026',
			dueDate: 'Jan 30, 2026',
			budget: 15000,
			spent: 13500,
			tasks: [
				{ id: 't11', title: 'Brand discovery', status: 'done', assignee: 'Emma Wilson', priority: 'high' },
				{ id: 't12', title: 'Logo concepts', status: 'done', assignee: 'Emma Wilson', priority: 'high' },
				{ id: 't13', title: 'Color palette', status: 'done', assignee: 'Emma Wilson', priority: 'medium' },
				{ id: 't14', title: 'Brand guidelines', status: 'review', assignee: 'Emma Wilson', priority: 'high' }
			],
			team: [
				{ name: 'Emma Wilson', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Emma' }
			],
			tags: ['Branding', 'Design'],
			color: 'bg-emerald-500'
		},
		{
			id: '4',
			name: 'Marketing Campaign',
			description: 'Q1 digital marketing campaign across social media, email, and paid ads.',
			client: 'Global Dynamics',
			clientAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Global',
			status: 'planning',
			progress: 10,
			priority: 'medium',
			startDate: 'Feb 1, 2026',
			dueDate: 'Mar 15, 2026',
			budget: 35000,
			spent: 3500,
			tasks: [
				{ id: 't15', title: 'Strategy document', status: 'done', assignee: 'John Smith', priority: 'high' },
				{ id: 't16', title: 'Content calendar', status: 'in-progress', assignee: 'Jane Doe', priority: 'high' },
				{ id: 't17', title: 'Ad creatives', status: 'todo', assignee: 'Mike Ross', priority: 'medium' },
				{ id: 't18', title: 'Landing page', status: 'todo', assignee: 'Sarah Chen', priority: 'medium' }
			],
			team: [
				{ name: 'John Smith', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=John' },
				{ name: 'Jane Doe', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Jane' }
			],
			tags: ['Marketing', 'Social Media'],
			color: 'bg-amber-500'
		},
		{
			id: '5',
			name: 'E-commerce Platform',
			description: 'Custom e-commerce solution with inventory management and payment processing.',
			client: 'Retail Plus',
			clientAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Retail',
			status: 'on-hold',
			progress: 45,
			priority: 'low',
			startDate: 'Dec 1, 2025',
			dueDate: 'Apr 1, 2026',
			budget: 60000,
			spent: 27000,
			tasks: [
				{ id: 't19', title: 'Architecture design', status: 'done', assignee: 'David Kim', priority: 'high' },
				{ id: 't20', title: 'Database schema', status: 'done', assignee: 'David Kim', priority: 'high' },
				{ id: 't21', title: 'User authentication', status: 'in-progress', assignee: 'Tom Johnson', priority: 'high' },
				{ id: 't22', title: 'Product catalog', status: 'todo', assignee: 'Lisa Wang', priority: 'medium' }
			],
			team: [
				{ name: 'David Kim', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=David' },
				{ name: 'Tom Johnson', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Tom' }
			],
			tags: ['E-commerce', 'Web Development'],
			color: 'bg-rose-500'
		},
		{
			id: '6',
			name: 'Dashboard Analytics',
			description: 'Internal analytics dashboard for tracking KPIs and business metrics.',
			client: 'Internal',
			clientAvatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Internal',
			status: 'completed',
			progress: 100,
			priority: 'high',
			startDate: 'Dec 15, 2025',
			dueDate: 'Jan 20, 2026',
			budget: 20000,
			spent: 19500,
			tasks: [
				{ id: 't23', title: 'Requirements', status: 'done', assignee: 'Alex Morgan', priority: 'high' },
				{ id: 't24', title: 'Data pipeline', status: 'done', assignee: 'David Kim', priority: 'high' },
				{ id: 't25', title: 'Chart components', status: 'done', assignee: 'Sarah Chen', priority: 'medium' },
				{ id: 't26', title: 'Dashboard UI', status: 'done', assignee: 'Sarah Chen', priority: 'medium' }
			],
			team: [
				{ name: 'Alex Morgan', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex' },
				{ name: 'David Kim', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=David' },
				{ name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah' }
			],
			tags: ['Internal', 'Analytics'],
			color: 'bg-cyan-500'
		}
	];

	const statusOptions = [
		{ value: 'all', label: 'All Projects', count: projects.length },
		{ value: 'in-progress', label: 'In Progress', count: projects.filter(p => p.status === 'in-progress').length },
		{ value: 'planning', label: 'Planning', count: projects.filter(p => p.status === 'planning').length },
		{ value: 'review', label: 'Review', count: projects.filter(p => p.status === 'review').length },
		{ value: 'completed', label: 'Completed', count: projects.filter(p => p.status === 'completed').length },
		{ value: 'on-hold', label: 'On Hold', count: projects.filter(p => p.status === 'on-hold').length }
	];

	function formatCurrency(amount: number): string {
		return new Intl.NumberFormat('en-US', {
			style: 'currency',
			currency: 'USD',
			maximumFractionDigits: 0
		}).format(amount);
	}

	function getStatusColor(status: string): string {
		switch (status) {
			case 'in-progress':
				return 'bg-emerald-500/10 text-emerald-600 border-emerald-500/20';
			case 'planning':
				return 'bg-blue-500/10 text-blue-600 border-blue-500/20';
			case 'review':
				return 'bg-amber-500/10 text-amber-600 border-amber-500/20';
			case 'completed':
				return 'bg-violet-500/10 text-violet-600 border-violet-500/20';
			case 'on-hold':
				return 'bg-slate-500/10 text-slate-600 border-slate-500/20';
			default:
				return 'bg-muted';
		}
	}

	function getStatusLabel(status: string): string {
		return status.split('-').map(w => w.charAt(0).toUpperCase() + w.slice(1)).join(' ');
	}

	function getPriorityColor(priority: string): string {
		switch (priority) {
			case 'high':
				return 'text-rose-500 bg-rose-500/10';
			case 'medium':
				return 'text-amber-500 bg-amber-500/10';
			case 'low':
				return 'text-blue-500 bg-blue-500/10';
			default:
				return 'text-muted-foreground bg-muted';
		}
	}

	const filteredProjects = $derived(() => {
		let filtered = projects;
		if (statusFilter !== 'all') {
			filtered = filtered.filter(p => p.status === statusFilter);
		}
		if (searchQuery) {
			const q = searchQuery.toLowerCase();
			filtered = filtered.filter(
				p =>
					p.name.toLowerCase().includes(q) ||
					p.client.toLowerCase().includes(q) ||
					p.description.toLowerCase().includes(q)
			);
		}
		return filtered;
	});

	const kanbanColumns = [
		{ id: 'todo', title: 'To Do', color: 'bg-slate-500' },
		{ id: 'in-progress', title: 'In Progress', color: 'bg-blue-500' },
		{ id: 'review', title: 'Review', color: 'bg-amber-500' },
		{ id: 'done', title: 'Done', color: 'bg-emerald-500' }
	];

	let newProject = $state({
		name: '',
		description: '',
		client: '',
		status: 'planning' as const,
		priority: 'medium' as const,
		dueDate: '',
		budget: ''
	});

	function handleCreateProject(e: Event) {
		e.preventDefault();
		console.log('Creating project:', newProject);
		newProject = {
			name: '',
			description: '',
			client: '',
			status: 'planning',
			priority: 'medium',
			dueDate: '',
			budget: ''
		};
		showNewProjectModal = false;
	}
</script>

<svelte:head>
	<title>Projects - Artemis</title>
</svelte:head>

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="px-6 py-4">
		<div class="mb-6 flex items-center justify-end gap-2">
			<Button variant="outline" class="gap-2">
				<Icon icon="hugeicons:filter" class="size-5" />
				<span class="hidden sm:inline">Filter</span>
			</Button>
			<Button class="gap-2" onclick={() => (showNewProjectModal = true)}>
				<Icon icon="hugeicons:plus-sign" class="size-5" />
				New Project
			</Button>
		</div>

		<div class="overflow-hidden rounded-2xl border border-border bg-card">
			<div class="flex flex-col gap-4 border-b border-border p-4 lg:flex-row lg:items-center">
				<div class="relative flex-1">
					<Icon
						icon="hugeicons:search-01"
						class="absolute top-1/2 left-3 size-5 -translate-y-1/2 text-muted-foreground"
					/>
					<Input
						placeholder="Search projects by name, client, or description..."
						class="pl-11"
						bind:value={searchQuery}
					/>
				</div>

				<div class="flex flex-wrap items-center gap-2">
					<div class="flex rounded-lg border border-border p-1">
						{#each statusOptions as option}
							<button
								class="rounded-md px-3 py-1.5 text-sm font-medium transition-colors {statusFilter ===
								option.value
									? 'bg-primary text-primary-foreground'
									: 'text-muted-foreground hover:text-foreground'}"
								onclick={() => (statusFilter = option.value)}
							>
								{option.label}
								<span class="ml-1 text-xs opacity-70">({option.count})</span>
							</button>
						{/each}
					</div>

					<Separator orientation="vertical" class="hidden h-9 lg:block" />

					<div class="flex rounded-lg border border-border p-1">
						<button
							class="rounded p-1.5 transition-colors {viewMode === 'board'
								? 'bg-muted text-foreground'
								: 'text-muted-foreground hover:text-foreground'}"
							onclick={() => (viewMode = 'board')}
							title="Board View"
						>
							<Icon icon="hugeicons:grid" class="size-5" />
						</button>
						<button
							class="rounded p-1.5 transition-colors {viewMode === 'list'
								? 'bg-muted text-foreground'
								: 'text-muted-foreground hover:text-foreground'}"
							onclick={() => (viewMode = 'list')}
							title="List View"
						>
							<Icon icon="hugeicons:menu-square" class="size-5" />
						</button>
						<button
							class="rounded p-1.5 transition-colors {viewMode === 'kanban'
								? 'bg-muted text-foreground'
								: 'text-muted-foreground hover:text-foreground'}"
							onclick={() => (viewMode = 'kanban')}
							title="Kanban View"
						>
							<Icon icon="hugeicons:columns-01" class="size-5" />
						</button>
					</div>
				</div>
			</div>

			{#if filteredProjects().length === 0}
				<div class="p-12 text-center">
					<div class="mb-4 inline-flex size-16 items-center justify-center rounded-full bg-muted">
						<Icon icon="hugeicons:folder-01" class="size-8 text-muted-foreground" />
					</div>
					<h3 class="text-lg font-semibold">No projects found</h3>
					<p class="mt-1 text-muted-foreground">Try adjusting your search or filters.</p>
				</div>

			{:else if viewMode === 'board'}
				<div class="grid gap-4 p-4 md:grid-cols-2 xl:grid-cols-3">
					{#each filteredProjects() as project}
						<div
							class="group flex flex-col rounded-xl border border-border bg-card p-5 transition-all hover:border-primary/20 hover:shadow-lg"
						>
							<div class="mb-4 flex items-start justify-between">
								<div class="flex items-center gap-3">
									<div class="size-3 rounded-full {project.color}"></div>
									<h3 class="font-semibold transition-colors group-hover:text-primary">
										{project.name}
									</h3>
								</div>
								<span
									class="rounded-full border px-2 py-1 text-xs font-medium capitalize {getStatusColor(
										project.status
									)}"
								>
									{getStatusLabel(project.status)}
								</span>
							</div>

							<p class="mb-4 line-clamp-2 text-sm text-muted-foreground">
								{project.description}
							</p>

							<div class="mb-4 flex items-center gap-2">
								<img
									src={project.clientAvatar}
									alt={project.client}
									class="size-6 rounded-full bg-muted"
								/>
								<span class="text-sm text-muted-foreground">{project.client}</span>
							</div>

							<div class="mb-4">
								<div class="mb-2 flex items-center justify-between">
									<span class="text-sm font-medium">{project.progress}%</span>
									<span class="text-xs text-muted-foreground">
										{project.tasks.filter(t => t.status === 'done').length}/{project.tasks.length} tasks
									</span>
								</div>
								<div class="h-2 overflow-hidden rounded-full bg-muted">
									<div
										class="h-full {project.color} rounded-full transition-all"
										style="width: {project.progress}%"
									></div>
								</div>
							</div>

							<div class="mb-4 flex items-center justify-between text-sm">
								<div class="flex items-center gap-1 text-muted-foreground">
									<Icon icon="hugeicons:calendar-01" class="size-4" />
									<span>Due {project.dueDate}</span>
								</div>
								<span class="rounded px-2 py-0.5 text-xs font-medium capitalize {getPriorityColor(project.priority)}">
									{project.priority}
								</span>
							</div>

							<div class="mb-4 flex items-center justify-between">
								<div class="flex -space-x-2">
									{#each project.team.slice(0, 4) as member, i}
										<img
											src={member.avatar}
											alt={member.name}
											class="size-8 rounded-full border-2 border-background bg-muted"
											title={member.name}
										/>
										{/each}
									{#if project.team.length > 4}
										<div class="flex size-8 items-center justify-center rounded-full border-2 border-background bg-muted text-xs font-medium">
											+{project.team.length - 4}
										</div>
									{/if}
								</div>
								{#if project.budget}
									<span class="text-sm font-medium">{formatCurrency(project.spent || 0)} / {formatCurrency(project.budget)}</span>
								{/if}
							</div>

							<div class="mb-4 flex flex-wrap gap-2">
								{#each project.tags as tag}
									<span class="rounded-md bg-muted px-2 py-1 text-xs">{tag}</span>
								{/each}
							</div>

							<div class="mt-auto flex gap-2 border-t border-border pt-4">
								<Button variant="outline" size="sm" class="flex-1">
									<Icon icon="hugeids:view" class="mr-2 size-4" />
									View
								</Button>
								<Button variant="outline" size="sm" class="flex-1">
									<Icon icon="hugeicons:pencil-edit-01" class="mr-2 size-4" />
									Edit
								</Button>
							</div>
						</div>
					{/each}
				</div>

			{:else if viewMode === 'list'}
				<div class="divide-y divide-border">
					{#each filteredProjects() as project}
						<div
							class="group flex cursor-pointer flex-col gap-4 p-4 transition-colors hover:bg-muted/30 sm:flex-row sm:items-center"
						>
							<div class="flex min-w-0 flex-1 items-center gap-3">
								<div class="size-3 shrink-0 rounded-full {project.color}"></div>
								<div class="min-w-0">
									<p class="truncate font-semibold transition-colors group-hover:text-primary">
										{project.name}
									</p>
									<p class="truncate text-sm text-muted-foreground">{project.client}</p>
								</div>
							</div>

							<div class="hidden sm:block">
								<span
									class="rounded-full border px-2 py-1 text-xs font-medium capitalize {getStatusColor(
										project.status
									)}"
								>
									{getStatusLabel(project.status)}
								</span>
							</div>

							<div class="hidden w-32 md:block">
								<div class="mb-1 flex justify-between text-xs">
									<span class="font-medium">{project.progress}%</span>
								</div>
								<div class="h-1.5 overflow-hidden rounded-full bg-muted">
									<div
										class="h-full {project.color} rounded-full"
										style="width: {project.progress}%"
									></div>
								</div>
							</div>

							<div class="hidden text-right lg:block">
								<p class="text-sm font-medium">
									{project.tasks.filter(t => t.status === 'done').length}/{project.tasks.length}
								</p>
								<p class="text-xs text-muted-foreground">Tasks done</p>
							</div>

							<div class="hidden text-right lg:block">
								<p class="text-sm font-medium">{project.dueDate}</p>
								<p class="text-xs text-muted-foreground">Due date</p>
							</div>

							{#if project.budget}
								<div class="hidden text-right xl:block">
									<p class="text-sm font-medium">{formatCurrency(project.budget)}</p>
									<p class="text-xs text-muted-foreground">Budget</p>
								</div>
							{/if}

							<div class="hidden items-center justify-end gap-2 md:flex">
								<div class="flex -space-x-2">
									{#each project.team.slice(0, 3) as member}
										<img
											src={member.avatar}
											alt={member.name}
											class="size-7 rounded-full border-2 border-background bg-muted"
										/>
									{/each}
								</div>
								<Button variant="ghost" size="icon" class="opacity-0 transition-opacity group-hover:opacity-100">
									<Icon icon="hugeicons:more-vertical" class="size-4" />
								</Button>
							</div>
						</div>
					{/each}
				</div>

			{:else if viewMode === 'kanban'}
				<div class="flex h-[calc(100vh-24rem)] gap-4 overflow-x-auto p-4">
					{#each kanbanColumns as column}
						<div class="flex w-80 shrink-0 flex-col rounded-xl border border-border bg-muted/30">
							<div class="flex items-center justify-between border-b border-border p-3">
								<div class="flex items-center gap-2">
									<div class="size-2 rounded-full {column.color}"></div>
									<h3 class="font-semibold">{column.title}</h3>
									<span class="rounded-full bg-muted px-2 py-0.5 text-xs">
										{projects.flatMap(p => p.tasks).filter(t => t.status === column.id).length}
									</span>
								</div>
								<Button variant="ghost" size="icon" class="size-7">
									<Icon icon="hugeicons:plus-sign" class="size-4" />
								</Button>
							</div>

							<div class="flex-1 space-y-2 overflow-y-auto p-2">
								{#each projects.flatMap(p => p.tasks.filter(t => t.status === column.id).map(t => ({ ...t, project: p }))) as task}
									<div
										class="group cursor-pointer rounded-lg border border-border bg-card p-3 transition-all hover:border-primary/20 hover:shadow-md"
									>
										<div class="mb-2 flex items-center gap-2">
											<div class="size-2 rounded-full {task.project.color}"></div>
											<span class="text-xs text-muted-foreground">{task.project.name}</span>
										</div>

										<h4 class="mb-3 text-sm font-medium">{task.title}</h4>

										<div class="flex items-center justify-between">
											{#if task.assignee}
												<div class="flex items-center gap-1.5">
													<div class="flex size-6 items-center justify-center rounded-full bg-primary/10 text-xs font-medium text-primary">
														{task.assignee.split(' ').map(n => n[0]).join('')}
													</div>
													<span class="text-xs text-muted-foreground">{task.assignee}</span>
												</div>
											{:else}
												<span class="text-xs text-muted-foreground">Unassigned</span>
											{/if}

											<span class="rounded px-1.5 py-0.5 text-xs capitalize {getPriorityColor(task.priority)}">
												{task.priority}
											</span>
										</div>
									</div>
								{/each}
							</div>
						</div>
					{/each}
				</div>
			{/if}
		</div>
	</div>
</div>

<Modal
	bind:open={showNewProjectModal}
	title="Create New Project"
	description="Fill in the details below to create a new project."
	size="lg"
>
	<form id="new-project-form" onsubmit={handleCreateProject} class="space-y-4">
		<div class="space-y-2">
			<Label for="project-name">Project Name *</Label>
			<Input id="project-name" placeholder="Enter project name" bind:value={newProject.name} required />
		</div>

		<div class="space-y-2">
			<Label for="project-description">Description</Label>
			<textarea
				id="project-description"
				placeholder="Describe the project..."
				bind:value={newProject.description}
				rows={3}
				class="flex w-full resize-none rounded-md border border-input bg-background px-3 py-2 text-sm shadow-sm transition-colors placeholder:text-muted-foreground focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
			></textarea>
		</div>

		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="project-client">Client</Label>
				<Input id="project-client" placeholder="Client name" bind:value={newProject.client} />
			</div>
			<div class="space-y-2">
				<Label for="project-due-date">Due Date</Label>
				<Input id="project-due-date" type="date" bind:value={newProject.dueDate} />
			</div>
		</div>

		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="project-status">Status</Label>
				<select
					id="project-status"
					bind:value={newProject.status}
					class="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-sm transition-colors focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				>
					<option value="planning">Planning</option>
					<option value="in-progress">In Progress</option>
					<option value="review">Review</option>
					<option value="completed">Completed</option>
					<option value="on-hold">On Hold</option>
				</select>
			</div>
			<div class="space-y-2">
				<Label for="project-priority">Priority</Label>
				<select
					id="project-priority"
					bind:value={newProject.priority}
					class="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-sm transition-colors focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				>
					<option value="low">Low</option>
					<option value="medium">Medium</option>
					<option value="high">High</option>
				</select>
			</div>
		</div>

		<div class="space-y-2">
			<Label for="project-budget">Budget (USD)</Label>
			<Input
				id="project-budget"
				type="number"
				placeholder="0.00"
				bind:value={newProject.budget}
			/>
		</div>
	</form>

	{#snippet footer()}
		<Button variant="outline" onclick={() => (showNewProjectModal = false)}>Cancel</Button>
		<Button type="submit" form="new-project-form">
			<Icon icon="hugeicons:plus-sign" class="mr-2 size-4" />
			Create Project
		</Button>
	{/snippet}
</Modal>
