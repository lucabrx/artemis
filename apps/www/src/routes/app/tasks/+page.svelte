<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';
	import { Label } from '$lib/components/ui/label/index.js';

	let viewMode = $state<'kanban' | 'list'>('kanban');
	let searchQuery = $state('');
	let projectFilter = $state('all');
	let statusFilter = $state('all');
	let priorityFilter = $state('all');
	let showNewTaskModal = $state(false);

	interface Task {
		id: string;
		title: string;
		description: string;
		status: 'todo' | 'in-progress' | 'review' | 'done';
		priority: 'low' | 'medium' | 'high' | 'urgent';
		project: {
			id: string;
			name: string;
			color: string;
			client: string;
		};
		assignee: {
			name: string;
			avatar: string;
			initials: string;
		} | null;
		dueDate: string;
		estimatedHours: number;
		loggedHours: number;
		subtasks: { id: string; title: string; completed: boolean }[];
		tags: string[];
		attachments: number;
		comments: number;
		createdAt: string;
	}

	const projects = [
		{ id: '1', name: 'Website Redesign', client: 'Acme Corp', color: 'bg-blue-500' },
		{ id: '2', name: 'Mobile App', client: 'TechStart', color: 'bg-violet-500' },
		{ id: '3', name: 'Brand Identity', client: 'Green Leaf', color: 'bg-emerald-500' },
		{ id: '4', name: 'Marketing Campaign', client: 'Global Dynamics', color: 'bg-amber-500' },
		{ id: '5', name: 'E-commerce Platform', client: 'Retail Plus', color: 'bg-rose-500' }
	];

	const tasks: Task[] = [
		{
			id: 't1',
			title: 'Design system setup',
			description: 'Create the foundation design system with colors, typography, and component library.',
			status: 'done',
			priority: 'high',
			project: projects[0],
			assignee: { name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah', initials: 'SC' },
			dueDate: 'Jan 10',
			estimatedHours: 16,
			loggedHours: 14,
			subtasks: [
				{ id: 's1', title: 'Color palette', completed: true },
				{ id: 's2', title: 'Typography scale', completed: true },
				{ id: 's3', title: 'Component library', completed: true }
			],
			tags: ['Design', 'UI/UX'],
			attachments: 3,
			comments: 5,
			createdAt: 'Jan 5'
		},
		{
			id: 't2',
			title: 'Homepage mockups',
			description: 'Design high-fidelity mockups for the homepage with all responsive breakpoints.',
			status: 'done',
			priority: 'high',
			project: projects[0],
			assignee: { name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah', initials: 'SC' },
			dueDate: 'Jan 15',
			estimatedHours: 24,
			loggedHours: 22,
			subtasks: [
				{ id: 's4', title: 'Desktop version', completed: true },
				{ id: 's5', title: 'Mobile version', completed: true },
				{ id: 's6', title: 'Tablet version', completed: true }
			],
			tags: ['Design', 'Mockups'],
			attachments: 8,
			comments: 12,
			createdAt: 'Jan 8'
		},
		{
			id: 't3',
			title: 'About page development',
			description: 'Implement the about page with team section and company history.',
			status: 'in-progress',
			priority: 'medium',
			project: projects[0],
			assignee: { name: 'Mike Ross', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Mike', initials: 'MR' },
			dueDate: 'Jan 25',
			estimatedHours: 12,
			loggedHours: 6,
			subtasks: [
				{ id: 's7', title: 'HTML structure', completed: true },
				{ id: 's8', title: 'Styling', completed: true },
				{ id: 's9', title: 'Animations', completed: false }
			],
			tags: ['Development', 'Frontend'],
			attachments: 2,
			comments: 3,
			createdAt: 'Jan 12'
		},
		{
			id: 't4',
			title: 'Contact form integration',
			description: 'Build and integrate the contact form with validation and email notifications.',
			status: 'todo',
			priority: 'medium',
			project: projects[0],
			assignee: { name: 'Mike Ross', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Mike', initials: 'MR' },
			dueDate: 'Jan 30',
			estimatedHours: 8,
			loggedHours: 0,
			subtasks: [
				{ id: 's10', title: 'Form UI', completed: false },
				{ id: 's11', title: 'Validation', completed: false },
				{ id: 's12', title: 'Email integration', completed: false }
			],
			tags: ['Development', 'Forms'],
			attachments: 1,
			comments: 0,
			createdAt: 'Jan 14'
		},
		{
			id: 't5',
			title: 'Mobile responsiveness review',
			description: 'Test and fix responsive issues across all pages on various devices.',
			status: 'review',
			priority: 'high',
			project: projects[0],
			assignee: { name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah', initials: 'SC' },
			dueDate: 'Feb 1',
			estimatedHours: 10,
			loggedHours: 8,
			subtasks: [
				{ id: 's13', title: 'iPhone testing', completed: true },
				{ id: 's14', title: 'Android testing', completed: true },
				{ id: 's15', title: 'Bug fixes', completed: false }
			],
			tags: ['QA', 'Mobile'],
			attachments: 5,
			comments: 7,
			createdAt: 'Jan 18'
		},
		{
			id: 't6',
			title: 'API architecture design',
			description: 'Design the REST API architecture and define all endpoints.',
			status: 'in-progress',
			priority: 'urgent',
			project: projects[1],
			assignee: { name: 'David Kim', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=David', initials: 'DK' },
			dueDate: 'Jan 20',
			estimatedHours: 20,
			loggedHours: 15,
			subtasks: [
				{ id: 's16', title: 'Endpoint definitions', completed: true },
				{ id: 's17', title: 'Data models', completed: true },
				{ id: 's18', title: 'Authentication', completed: false }
			],
			tags: ['API', 'Backend'],
			attachments: 4,
			comments: 8,
			createdAt: 'Jan 10'
		},
		{
			id: 't7',
			title: 'iOS prototype',
			description: 'Build interactive prototype for iOS app using SwiftUI.',
			status: 'todo',
			priority: 'high',
			project: projects[1],
			assignee: { name: 'Lisa Wang', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Lisa', initials: 'LW' },
			dueDate: 'Feb 5',
			estimatedHours: 40,
			loggedHours: 0,
			subtasks: [
				{ id: 's19', title: 'Setup project', completed: false },
				{ id: 's20', title: 'Main screens', completed: false },
				{ id: 's21', title: 'Navigation', completed: false }
			],
			tags: ['iOS', 'Mobile'],
			attachments: 2,
			comments: 2,
			createdAt: 'Jan 15'
		},
		{
			id: 't8',
			title: 'Logo concepts',
			description: 'Create 3 different logo concepts for brand refresh.',
			status: 'review',
			priority: 'high',
			project: projects[2],
			assignee: { name: 'Emma Wilson', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Emma', initials: 'EW' },
			dueDate: 'Jan 22',
			estimatedHours: 16,
			loggedHours: 14,
			subtasks: [
				{ id: 's22', title: 'Concept 1', completed: true },
				{ id: 's23', title: 'Concept 2', completed: true },
				{ id: 's24', title: 'Concept 3', completed: true }
			],
			tags: ['Branding', 'Logo'],
			attachments: 6,
			comments: 15,
			createdAt: 'Jan 8'
		},
		{
			id: 't9',
			title: 'Brand guidelines document',
			description: 'Compile comprehensive brand guidelines with usage rules.',
			status: 'in-progress',
			priority: 'medium',
			project: projects[2],
			assignee: { name: 'Emma Wilson', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Emma', initials: 'EW' },
			dueDate: 'Jan 28',
			estimatedHours: 12,
			loggedHours: 8,
			subtasks: [
				{ id: 's25', title: 'Logo usage', completed: true },
				{ id: 's26', title: 'Color specifications', completed: false },
				{ id: 's27', title: 'Typography rules', completed: false }
			],
			tags: ['Branding', 'Documentation'],
			attachments: 3,
			comments: 4,
			createdAt: 'Jan 12'
		},
		{
			id: 't10',
			title: 'Content calendar',
			description: 'Plan content calendar for Q1 marketing campaign.',
			status: 'done',
			priority: 'medium',
			project: projects[3],
			assignee: { name: 'Jane Doe', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Jane', initials: 'JD' },
			dueDate: 'Jan 18',
			estimatedHours: 8,
			loggedHours: 7,
			subtasks: [
				{ id: 's28', title: 'Research topics', completed: true },
				{ id: 's29', title: 'Schedule posts', completed: true },
				{ id: 's30', title: 'Assign writers', completed: true }
			],
			tags: ['Marketing', 'Content'],
			attachments: 2,
			comments: 6,
			createdAt: 'Jan 5'
		},
		{
			id: 't11',
			title: 'Ad creative design',
			description: 'Design banner ads for Google Ads and social media.',
			status: 'todo',
			priority: 'low',
			project: projects[3],
			assignee: null,
			dueDate: 'Feb 10',
			estimatedHours: 12,
			loggedHours: 0,
			subtasks: [
				{ id: 's31', title: 'Google Ads sizes', completed: false },
				{ id: 's32', title: 'Facebook/Instagram', completed: false },
				{ id: 's33', title: 'LinkedIn ads', completed: false }
			],
			tags: ['Marketing', 'Design'],
			attachments: 0,
			comments: 1,
			createdAt: 'Jan 16'
		},
		{
			id: 't12',
			title: 'Database schema design',
			description: 'Design PostgreSQL schema for e-commerce platform.',
			status: 'done',
			priority: 'high',
			project: projects[4],
			assignee: { name: 'David Kim', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=David', initials: 'DK' },
			dueDate: 'Jan 15',
			estimatedHours: 16,
			loggedHours: 14,
			subtasks: [
				{ id: 's34', title: 'User tables', completed: true },
				{ id: 's35', title: 'Product tables', completed: true },
				{ id: 's36', title: 'Order tables', completed: true }
			],
			tags: ['Database', 'Backend'],
			attachments: 3,
			comments: 5,
			createdAt: 'Jan 3'
		}
	];

	const kanbanColumns = [
		{ id: 'todo', title: 'To Do', color: 'bg-slate-500', icon: 'hugeicons:circle' },
		{ id: 'in-progress', title: 'In Progress', color: 'bg-blue-500', icon: 'hugeicons:loading-02' },
		{ id: 'review', title: 'Review', color: 'bg-amber-500', icon: 'hugeicons:search-01' },
		{ id: 'done', title: 'Done', color: 'bg-emerald-500', icon: 'hugeicons:tick-01' }
	];

	const statusOptions = [
		{ value: 'all', label: 'All Tasks', count: tasks.length },
		{ value: 'todo', label: 'To Do', count: tasks.filter(t => t.status === 'todo').length },
		{ value: 'in-progress', label: 'In Progress', count: tasks.filter(t => t.status === 'in-progress').length },
		{ value: 'review', label: 'In Review', count: tasks.filter(t => t.status === 'review').length },
		{ value: 'done', label: 'Done', count: tasks.filter(t => t.status === 'done').length }
	];

	const projectOptions = [
		{ value: 'all', label: 'All Projects', count: tasks.length },
		...projects.map(p => ({
			value: p.id,
			label: p.name,
			count: tasks.filter(t => t.project.id === p.id).length
		}))
	];

	const priorityOptions = [
		{ value: 'all', label: 'All Priorities' },
		{ value: 'urgent', label: 'Urgent' },
		{ value: 'high', label: 'High' },
		{ value: 'medium', label: 'Medium' },
		{ value: 'low', label: 'Low' }
	];

	let newTask = $state({
		title: '',
		description: '',
		project: '',
		status: 'todo' as const,
		priority: 'medium' as const,
		assignee: '',
		dueDate: '',
		estimatedHours: ''
	});

	function getStatusColor(status: string): string {
		switch (status) {
			case 'todo':
				return 'bg-slate-500/10 text-slate-600 border-slate-500/20';
			case 'in-progress':
				return 'bg-blue-500/10 text-blue-600 border-blue-500/20';
			case 'review':
				return 'bg-amber-500/10 text-amber-600 border-amber-500/20';
			case 'done':
				return 'bg-emerald-500/10 text-emerald-600 border-emerald-500/20';
			default:
				return 'bg-muted';
		}
	}

	function getPriorityColor(priority: string): string {
		switch (priority) {
			case 'urgent':
				return 'bg-rose-500 text-white';
			case 'high':
				return 'bg-orange-500 text-white';
			case 'medium':
				return 'bg-amber-400 text-white';
			case 'low':
				return 'bg-blue-400 text-white';
			default:
				return 'bg-muted';
		}
	}

	function getPriorityIcon(priority: string): string {
		switch (priority) {
			case 'urgent':
				return 'hugeicons:alert-02';
			case 'high':
				return 'hugeicons:arrow-up-02';
			case 'medium':
				return 'hugeicons:arrow-right-02';
			case 'low':
				return 'hugeicons:arrow-down-02';
			default:
				return 'hugeicons:minus-sign';
		}
	}

	const filteredTasks = $derived(() => {
		let filtered = tasks;
		if (projectFilter !== 'all') {
			filtered = filtered.filter(t => t.project.id === projectFilter);
		}
		if (statusFilter !== 'all') {
			filtered = filtered.filter(t => t.status === statusFilter);
		}
		if (priorityFilter !== 'all') {
			filtered = filtered.filter(t => t.priority === priorityFilter);
		}
		if (searchQuery) {
			const q = searchQuery.toLowerCase();
			filtered = filtered.filter(
				t =>
					t.title.toLowerCase().includes(q) ||
					t.description.toLowerCase().includes(q) ||
					t.project.name.toLowerCase().includes(q)
			);
		}
		return filtered;
	});

	function handleCreateTask(e: Event) {
		e.preventDefault();
		console.log('Creating task:', newTask);
		newTask = {
			title: '',
			description: '',
			project: '',
			status: 'todo',
			priority: 'medium',
			assignee: '',
			dueDate: '',
			estimatedHours: ''
		};
		showNewTaskModal = false;
	}
</script>

<svelte:head>
	<title>Tasks - Artemis</title>
</svelte:head>

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="px-6 py-4">
		<div class="mb-6 flex items-center justify-end gap-2">
			<Button variant="outline" class="gap-2">
				<Icon icon="hugeicons:filter" class="size-5" />
				<span class="hidden sm:inline">Filter</span>
			</Button>
			<Button class="gap-2" onclick={() => (showNewTaskModal = true)}>
				<Icon icon="hugeicons:plus-sign" class="size-5" />
				New Task
			</Button>
		</div>

		<div class="overflow-hidden rounded-2xl border border-border bg-card">
			<div class="flex flex-col gap-4 border-b border-border p-4">
				<div class="flex flex-col gap-4 lg:flex-row lg:items-center">
					<div class="relative flex-1">
						<Icon
							icon="hugeicons:search-01"
							class="absolute top-1/2 left-3 size-5 -translate-y-1/2 text-muted-foreground"
						/>
						<Input
							placeholder="Search tasks by title, description, or project..."
							class="pl-11"
							bind:value={searchQuery}
						/>
					</div>

					<div class="flex flex-wrap items-center gap-2">
						<select
							bind:value={projectFilter}
							class="h-9 rounded-md border border-input bg-background px-3 text-sm"
						>
							{#each projectOptions as option}
								<option value={option.value}>{option.label} ({option.count})</option>
							{/each}
						</select>

						<select
							bind:value={priorityFilter}
							class="h-9 rounded-md border border-input bg-background px-3 text-sm"
						>
							{#each priorityOptions as option}
								<option value={option.value}>{option.label}</option>
							{/each}
						</select>

						<Separator orientation="vertical" class="hidden h-9 lg:block" />

						<div class="flex rounded-lg border border-border p-1">
							<button
								class="rounded p-1.5 transition-colors {viewMode === 'kanban'
									? 'bg-muted text-foreground'
									: 'text-muted-foreground hover:text-foreground'}"
								onclick={() => (viewMode = 'kanban')}
								title="Kanban View"
							>
								<Icon icon="hugeicons:columns-01" class="size-5" />
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
						</div>
					</div>
				</div>

				<div class="flex flex-wrap gap-2">
					{#each statusOptions as option}
						<button
							class="rounded-md px-3 py-1.5 text-sm font-medium transition-colors {statusFilter === option.value
								? 'bg-primary text-primary-foreground'
								: 'border border-border text-muted-foreground hover:text-foreground'}"
							onclick={() => (statusFilter = option.value)}
						>
							{option.label}
							<span class="ml-1 text-xs opacity-70">({option.count})</span>
						</button>
					{/each}
				</div>
			</div>

			{#if filteredTasks().length === 0}
				<div class="p-12 text-center">
					<div class="mb-4 inline-flex size-16 items-center justify-center rounded-full bg-muted">
						<Icon icon="hugeicons:task-01" class="size-8 text-muted-foreground" />
					</div>
					<h3 class="text-lg font-semibold">No tasks found</h3>
					<p class="mt-1 text-muted-foreground">Try adjusting your search or filters.</p>
				</div>

			{:else if viewMode === 'kanban'}
				<div class="flex h-[calc(100vh-20rem)] gap-4 overflow-x-auto p-4">
					{#each kanbanColumns as column}
						{@const columnTasks = filteredTasks().filter(t => t.status === column.id)}
						<div class="flex w-80 shrink-0 flex-col rounded-xl border border-border bg-muted/30">
							<div class="flex items-center justify-between border-b border-border p-3">
								<div class="flex items-center gap-2">
									<div class="size-2 rounded-full {column.color}"></div>
									<h3 class="font-semibold">{column.title}</h3>
									<span class="rounded-full bg-muted px-2 py-0.5 text-xs">
										{columnTasks.length}
									</span>
								</div>
								<Button variant="ghost" size="icon" class="size-7">
									<Icon icon="hugeicons:plus-sign" class="size-4" />
								</Button>
							</div>

							<div class="flex-1 space-y-2 overflow-y-auto p-2">
								{#each columnTasks as task}
									<a
										href="/app/tasks/{task.id}"
										class="group block rounded-lg border border-border bg-card p-3 transition-all hover:border-primary/20 hover:shadow-md"
									>
										<div class="mb-2 flex items-start justify-between gap-2">
											<div class="flex items-center gap-1.5">
												<div class="size-2 rounded-full {task.project.color}"></div>
												<span class="text-xs text-muted-foreground">{task.project.name}</span>
											</div>
											<div class="flex items-center gap-1">
												<Icon icon={getPriorityIcon(task.priority)} class="size-3 {task.priority === 'urgent' ? 'text-rose-500' : task.priority === 'high' ? 'text-orange-500' : 'text-muted-foreground'}" />
											</div>
										</div>

										<h4 class="mb-2 text-sm font-medium line-clamp-2">{task.title}</h4>

										{#if task.tags.length > 0}
											<div class="mb-2 flex flex-wrap gap-1">
												{#each task.tags.slice(0, 2) as tag}
													<span class="rounded bg-muted px-1.5 py-0.5 text-[10px]">{tag}</span>
												{/each}
												{#if task.tags.length > 2}
													<span class="rounded bg-muted px-1.5 py-0.5 text-[10px]">+{task.tags.length - 2}</span>
												{/if}
											</div>
										{/if}

										<div class="flex items-center justify-between">
											{#if task.assignee}
												<img
													src={task.assignee.avatar}
													alt={task.assignee.name}
													class="size-6 rounded-full border border-background"
													title={task.assignee.name}
												/>
											{:else}
												<div class="flex size-6 items-center justify-center rounded-full border border-dashed border-muted-foreground/30">
													<Icon icon="hugeicons:user" class="size-3 text-muted-foreground/50" />
												</div>
											{/if}

											<div class="flex items-center gap-3 text-xs text-muted-foreground">
												{#if task.subtasks.length > 0}
													<div class="flex items-center gap-1">
														<Icon icon="hugeicons:checkmark-square-02" class="size-3" />
														<span>{task.subtasks.filter(s => s.completed).length}/{task.subtasks.length}</span>
													</div>
												{/if}
												<div class="flex items-center gap-1">
													<Icon icon="hugeicons:calendar-01" class="size-3" />
													<span>{task.dueDate}</span>
												</div>
											</div>
										</div>
									</a>
								{/each}
							</div>
						</div>
					{/each}
				</div>

			{:else if viewMode === 'list'}
				<div class="divide-y divide-border">
					{#each filteredTasks() as task}
						<a
							href="/app/tasks/{task.id}"
							class="group flex cursor-pointer flex-col gap-4 p-4 transition-colors hover:bg-muted/30 sm:flex-row sm:items-center"
						>
							<div class="flex min-w-0 flex-1 items-center gap-3">
								<div class="size-2 shrink-0 rounded-full {task.project.color}"></div>
								<div class="min-w-0">
									<p class="truncate font-semibold transition-colors group-hover:text-primary">
										{task.title}
									</p>
									<p class="truncate text-sm text-muted-foreground">
										{task.project.name} · {task.project.client}
									</p>
								</div>
							</div>

							<div class="flex flex-wrap items-center gap-4 sm:flex-nowrap">
								<div class="flex items-center gap-2">
									<span class="flex items-center gap-1 rounded-full px-2 py-0.5 text-xs font-medium {getStatusColor(task.status)}">
										<Icon icon={kanbanColumns.find(c => c.id === task.status)?.icon || 'hugeicons:circle'} class="size-3" />
										{task.status.replace('-', ' ')}
									</span>
									<span class="flex items-center gap-1 rounded px-1.5 py-0.5 text-xs {getPriorityColor(task.priority)}">
										<Icon icon={getPriorityIcon(task.priority)} class="size-3" />
									</span>
								</div>

								<div class="hidden items-center gap-3 text-sm text-muted-foreground md:flex">
									{#if task.subtasks.length > 0}
										<div class="flex items-center gap-1">
											<Icon icon="hugeicons:checkmark-square-02" class="size-4" />
											<span>{task.subtasks.filter(s => s.completed).length}/{task.subtasks.length}</span>
										</div>
									{/if}
									<div class="flex items-center gap-1">
										<Icon icon="hugeicons:calendar-01" class="size-4" />
										<span>{task.dueDate}</span>
									</div>
								</div>

								<div class="flex items-center gap-3">
									{#if task.assignee}
										<img
											src={task.assignee.avatar}
											alt={task.assignee.name}
											class="size-8 rounded-full border-2 border-background"
										/>
									{:else}
										<div class="flex size-8 items-center justify-center rounded-full border-2 border-dashed border-muted-foreground/30">
											<Icon icon="hugeicons:user" class="size-4 text-muted-foreground/50" />
										</div>
									{/if}
									<Button
										variant="ghost"
										size="icon"
										class="opacity-0 transition-opacity group-hover:opacity-100"
										onclick={(e) => e.stopPropagation()}
									>
										<Icon icon="hugeicons:more-vertical" class="size-4" />
									</Button>
								</div>
							</div>
						</a>
					{/each}
				</div>
			{/if}
		</div>
	</div>
</div>

<Modal
	bind:open={showNewTaskModal}
	title="Create New Task"
	description="Add a new task to your project."
	size="lg"
>
	<form id="new-task-form" onsubmit={handleCreateTask} class="space-y-4">
		<div class="space-y-2">
			<Label for="task-title">Task Title *</Label>
			<Input id="task-title" placeholder="Enter task title" bind:value={newTask.title} required />
		</div>

		<div class="space-y-2">
			<Label for="task-description">Description</Label>
			<textarea
				id="task-description"
				placeholder="Describe the task..."
				bind:value={newTask.description}
				rows={3}
				class="flex w-full resize-none rounded-md border border-input bg-background px-3 py-2 text-sm shadow-sm transition-colors placeholder:text-muted-foreground focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-none"
			></textarea>
		</div>

		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="task-project">Project *</Label>
				<select
					id="task-project"
					bind:value={newTask.project}
					required
					class="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-sm transition-colors focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-none"
				>
					<option value="">Select a project</option>
					{#each projects as project}
						<option value={project.id}>{project.name}</option>
					{/each}
				</select>
			</div>
			<div class="space-y-2">
				<Label for="task-assignee">Assignee</Label>
				<Input id="task-assignee" placeholder="Assign to..." bind:value={newTask.assignee} />
			</div>
		</div>

		<div class="grid gap-4 md:grid-cols-3">
			<div class="space-y-2">
				<Label for="task-status">Status</Label>
				<select
					id="task-status"
					bind:value={newTask.status}
					class="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-sm transition-colors focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-none"
				>
					<option value="todo">To Do</option>
					<option value="in-progress">In Progress</option>
					<option value="review">Review</option>
					<option value="done">Done</option>
				</select>
			</div>
			<div class="space-y-2">
				<Label for="task-priority">Priority</Label>
				<select
					id="task-priority"
					bind:value={newTask.priority}
					class="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-sm transition-colors focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-none"
				>
					<option value="low">Low</option>
					<option value="medium">Medium</option>
					<option value="high">High</option>
					<option value="urgent">Urgent</option>
				</select>
			</div>
			<div class="space-y-2">
				<Label for="task-due-date">Due Date</Label>
				<Input id="task-due-date" type="date" bind:value={newTask.dueDate} />
			</div>
		</div>

		<div class="space-y-2">
			<Label for="task-estimated-hours">Estimated Hours</Label>
			<Input
				id="task-estimated-hours"
				type="number"
				placeholder="0"
				bind:value={newTask.estimatedHours}
			/>
		</div>
	</form>

	{#snippet footer()}
		<Button variant="outline" onclick={() => (showNewTaskModal = false)}>Cancel</Button>
		<Button type="submit" form="new-task-form">
			<Icon icon="hugeicons:plus-sign" class="mr-2 size-4" />
			Create Task
		</Button>
	{/snippet}
</Modal>
