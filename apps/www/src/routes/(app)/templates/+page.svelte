<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';
	import { Label } from '$lib/components/ui/label/index.js';

	let viewMode = $state<'grid' | 'list'>('grid');
	let searchQuery = $state('');
	let categoryFilter = $state<string>('all');
	let typeFilter = $state<'all' | 'email' | 'document' | 'project'>('all');
	let showCreateModal = $state(false);
	let showPreviewModal = $state(false);
	let showEditModal = $state(false);
	let showDeleteModal = $state(false);
	let selectedTemplate = $state<Template | null>(null);

	interface Template {
		id: string;
		name: string;
		description: string;
		type: 'email' | 'document' | 'project';
		category: string;
		icon: string;
		color: string;
		content: string;
		variables: string[];
		lastUsed: string;
		usageCount: number;
		favorite: boolean;
		createdAt: string;
		author: {
			name: string;
			avatar: string;
		};
	}

	const templates: Template[] = [
		{
			id: 't1',
			name: 'New Client Onboarding',
			description: 'Welcome email for new clients with next steps',
			type: 'email',
			category: 'onboarding',
			icon: 'hugeicons:mail-01',
			color: 'bg-emerald-500',
			content: `Hi {{client_name}},<br><br>Welcome to {{company_name}}! I'm thrilled to have you on board.<br><br><strong>Next steps:</strong><br>1. We'll schedule a kickoff call<br>2. You'll receive access to our project tool<br>3. Initial questionnaire<br><br>Best regards,<br>{{sender_name}}`,
			variables: ['client_name', 'company_name', 'sender_name'],
			lastUsed: '2 days ago',
			usageCount: 24,
			favorite: true,
			createdAt: 'Dec 15, 2025',
			author: { name: 'Alex Morgan', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex' }
		},
		{
			id: 't2',
			name: 'Project Proposal',
			description: 'Professional project proposal with scope and pricing',
			type: 'document',
			category: 'proposals',
			icon: 'hugeicons:file-01',
			color: 'bg-blue-500',
			content: `PROJECT PROPOSAL<br><br>Client: {{client_name}}<br>Project: {{project_name}}<br><br>SCOPE:<br>{{project_scope}}<br><br>TIMELINE: {{timeline}}<br>INVESTMENT: {{price}}`,
			variables: ['client_name', 'project_name', 'project_scope', 'timeline', 'price'],
			lastUsed: '1 week ago',
			usageCount: 56,
			favorite: true,
			createdAt: 'Nov 20, 2025',
			author: { name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah' }
		},
		{
			id: 't3',
			name: 'Follow-up After Meeting',
			description: 'Meeting recap with action items',
			type: 'email',
			category: 'follow-ups',
			icon: 'hugeicons:mail-01',
			color: 'bg-amber-500',
			content: `Hi {{client_name}},<br><br>Great speaking with you about {{project_name}}.<br><br><strong>Next steps:</strong><br>• {{action_item_1}}<br>• {{action_item_2}}<br><br>Thanks,<br>{{sender_name}}`,
			variables: ['client_name', 'project_name', 'action_item_1', 'action_item_2', 'sender_name'],
			lastUsed: '3 days ago',
			usageCount: 38,
			favorite: false,
			createdAt: 'Dec 1, 2025',
			author: { name: 'Alex Morgan', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex' }
		},
		{
			id: 't4',
			name: 'Invoice Reminder',
			description: 'Friendly payment reminder email',
			type: 'email',
			category: 'invoices',
			icon: 'hugeicons:mail-01',
			color: 'bg-violet-500',
			content: `Hi {{client_name}},<br><br>This is a friendly reminder that Invoice {{invoice_number}} for {{amount}} is due on {{due_date}}.<br><br>Thank you for your business!`,
			variables: ['client_name', 'invoice_number', 'amount', 'due_date'],
			lastUsed: '5 days ago',
			usageCount: 42,
			favorite: false,
			createdAt: 'Nov 15, 2025',
			author: { name: 'Mike Ross', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Mike' }
		},
		{
			id: 't5',
			name: 'Web Design Project',
			description: 'Complete web design project structure',
			type: 'project',
			category: 'projects',
			icon: 'hugeicons:folder-kanban',
			color: 'bg-cyan-500',
			content: `Project Template with tasks: Discovery, Wireframes, Design, Development, Testing`,
			variables: ['client_name', 'project_name', 'budget'],
			lastUsed: '1 day ago',
			usageCount: 12,
			favorite: true,
			createdAt: 'Jan 10, 2026',
			author: { name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah' }
		},
		{
			id: 't6',
			name: 'Service Agreement',
			description: 'Standard service contract template',
			type: 'document',
			category: 'contracts',
			icon: 'hugeicons:file-01',
			color: 'bg-rose-500',
			content: `SERVICE AGREEMENT<br><br>Between {{company_name}} and {{client_name}}<br><br>Services: {{services}}<br>Terms: {{terms}}<br>Payment: {{payment_terms}}`,
			variables: ['company_name', 'client_name', 'services', 'terms', 'payment_terms'],
			lastUsed: '2 weeks ago',
			usageCount: 18,
			favorite: true,
			createdAt: 'Oct 5, 2025',
			author: { name: 'Alex Morgan', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex' }
		},
		{
			id: 't7',
			name: 'Weekly Update',
			description: 'Project status update email',
			type: 'email',
			category: 'updates',
			icon: 'hugeicons:mail-01',
			color: 'bg-indigo-500',
			content: `Hi {{client_name}},<br><br>Weekly update on {{project_name}}:<br><br><strong>Completed:</strong><br>• {{completed_1}}<br>• {{completed_2}}<br><br><strong>Up next:</strong><br>• {{upcoming_1}}`,
			variables: ['client_name', 'project_name', 'completed_1', 'completed_2', 'upcoming_1'],
			lastUsed: '4 days ago',
			usageCount: 31,
			favorite: false,
			createdAt: 'Dec 20, 2025',
			author: { name: 'Emma Wilson', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Emma' }
		},
		{
			id: 't8',
			name: 'Thank You Note',
			description: 'Appreciation email after project completion',
			type: 'email',
			category: 'general',
			icon: 'hugeicons:mail-01',
			color: 'bg-pink-500',
			content: `Hi {{client_name}},<br><br>Thank you for choosing {{company_name}} for {{project_name}}. It's been a pleasure working with you!<br><br>Warm regards,<br>{{sender_name}}`,
			variables: ['client_name', 'company_name', 'project_name', 'sender_name'],
			lastUsed: '1 month ago',
			usageCount: 15,
			favorite: false,
			createdAt: 'Sep 15, 2025',
			author: { name: 'Lisa Thompson', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Lisa' }
		}
	];

	const categories = [
		{ value: 'all', label: 'All Categories', count: templates.length },
		{ value: 'onboarding', label: 'Onboarding', count: templates.filter(t => t.category === 'onboarding').length },
		{ value: 'proposals', label: 'Proposals', count: templates.filter(t => t.category === 'proposals').length },
		{ value: 'follow-ups', label: 'Follow-ups', count: templates.filter(t => t.category === 'follow-ups').length },
		{ value: 'invoices', label: 'Invoices', count: templates.filter(t => t.category === 'invoices').length },
		{ value: 'contracts', label: 'Contracts', count: templates.filter(t => t.category === 'contracts').length },
		{ value: 'updates', label: 'Updates', count: templates.filter(t => t.category === 'updates').length },
		{ value: 'projects', label: 'Projects', count: templates.filter(t => t.category === 'projects').length },
		{ value: 'general', label: 'General', count: templates.filter(t => t.category === 'general').length }
	];

	const filteredTemplates = $derived(() => {
		let filtered = templates;
		
		if (typeFilter !== 'all') {
			filtered = filtered.filter(t => t.type === typeFilter);
		}
		
		if (categoryFilter !== 'all') {
			filtered = filtered.filter(t => t.category === categoryFilter);
		}
		
		if (searchQuery) {
			const q = searchQuery.toLowerCase();
			filtered = filtered.filter(t => 
				t.name.toLowerCase().includes(q) ||
				t.description.toLowerCase().includes(q) ||
				t.category.toLowerCase().includes(q)
			);
		}
		
		return filtered;
	});

	const favorites = $derived(() => templates.filter(t => t.favorite));
	const recentTemplates = $derived(() => [...templates]
		.sort((a, b) => {
			const timeA = a.lastUsed.includes('day') ? 1 : a.lastUsed.includes('week') ? 7 : 30;
			const timeB = b.lastUsed.includes('day') ? 1 : b.lastUsed.includes('week') ? 7 : 30;
			return timeA - timeB;
		})
		.slice(0, 3)
	);

	function toggleFavorite(e: Event, template: Template) {
		e.stopPropagation();
		template.favorite = !template.favorite;
	}

	let newTemplate = $state({
		name: '',
		description: '',
		type: 'email' as const,
		category: 'general',
		content: ''
	});

	function handleCreateTemplate(e: Event) {
		e.preventDefault();
		showCreateModal = false;
		newTemplate = { name: '', description: '', type: 'email', category: 'general', content: '' };
	}

	function handleDeleteTemplate() {
		showDeleteModal = false;
		selectedTemplate = null;
	}

	function useTemplate(template: Template) {
		selectedTemplate = template;
	}
</script>

<svelte:head>
	<title>Templates - Artemis</title>
</svelte:head>

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="flex h-[calc(100vh-3.5rem)]">
		<div class="w-64 border-r border-border bg-muted/30 flex flex-col">
			<div class="p-4">
				<Button class="w-full gap-2" onclick={() => showCreateModal = true}>
					<Icon icon="hugeicons:plus-sign" class="size-4" />
					New Template
				</Button>
			</div>
			
			<nav class="flex-1 px-2 space-y-0.5 overflow-y-auto">
				<button
					class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-all {typeFilter === 'all' && categoryFilter === 'all' ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:bg-muted hover:text-foreground'}"
					onclick={() => { typeFilter = 'all'; categoryFilter = 'all'; }}
				>
					<Icon icon="hugeicons:layout-template" class="size-5" />
					All Templates
				</button>
				<button
					class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-all {typeFilter === 'email' ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:bg-muted hover:text-foreground'}"
					onclick={() => { typeFilter = 'email'; categoryFilter = 'all'; }}
				>
					<Icon icon="hugeicons:mail-01" class="size-5" />
					Email Templates
				</button>
				<button
					class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-all {typeFilter === 'document' ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:bg-muted hover:text-foreground'}"
					onclick={() => { typeFilter = 'document'; categoryFilter = 'all'; }}
				>
					<Icon icon="hugeicons:file-01" class="size-5" />
					Documents
				</button>
				<button
					class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-all {typeFilter === 'project' ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:bg-muted hover:text-foreground'}"
					onclick={() => { typeFilter = 'project'; categoryFilter = 'all'; }}
				>
					<Icon icon="hugeicons:folder-kanban" class="size-5" />
					Project Templates
				</button>
				
				<Separator class="my-3" />
				
				<p class="px-3 py-2 text-xs font-semibold text-muted-foreground uppercase tracking-wider">Categories</p>
				{#each categories.filter(c => c.value !== 'all') as cat}
					<button
						class="w-full flex items-center justify-between px-3 py-2 rounded-lg text-sm font-medium transition-all {categoryFilter === cat.value ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:bg-muted hover:text-foreground'}"
						onclick={() => { categoryFilter = cat.value; typeFilter = 'all'; }}
					>
						<span class="capitalize">{cat.label}</span>
						<span class="text-xs opacity-70">{cat.count}</span>
					</button>
				{/each}
			</nav>
		</div>

		<div class="flex-1 flex flex-col min-w-0 overflow-hidden">
			<div class="px-6 py-4 border-b border-border bg-background/50 backdrop-blur">
				<div class="flex items-center justify-between gap-4">
					<div>
						<h1 class="text-2xl font-bold">Templates</h1>
						<p class="text-sm text-muted-foreground">Manage reusable templates for emails, documents, and projects</p>
					</div>
					
					<div class="flex items-center gap-2 shrink-0">
						<div class="relative">
							<Icon icon="hugeicons:search-01" class="absolute left-3 top-1/2 -translate-y-1/2 size-4 text-muted-foreground" />
							<Input placeholder="Search templates..." class="pl-9 w-64" bind:value={searchQuery} />
						</div>
						
						<Separator orientation="vertical" class="h-6" />
						
						<div class="flex rounded-lg border border-border p-1 bg-muted/30">
							<button
								class="rounded p-1.5 transition-all {viewMode === 'grid' ? 'bg-background text-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground'}"
								onclick={() => viewMode = 'grid'}
							>
								<Icon icon="hugeicons:grid" class="size-4" />
							</button>
							<button
								class="rounded p-1.5 transition-all {viewMode === 'list' ? 'bg-background text-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground'}"
								onclick={() => viewMode = 'list'}
							>
								<Icon icon="hugeicons:list" class="size-4" />
							</button>
						</div>
					</div>
				</div>
			</div>

			<div class="flex-1 overflow-y-auto p-6">
				{#if favorites().length > 0 && !searchQuery && categoryFilter === 'all' && typeFilter === 'all'}
					<div class="mb-8">
						<div class="flex items-center gap-2 mb-4">
							<Icon icon="hugeicons:star" class="size-5 text-amber-500" />
							<h2 class="text-lg font-semibold">Favorites</h2>
						</div>
						<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
							{#each favorites() as template (template.id)}
								<div
									class="group relative rounded-xl border border-border bg-card p-4 transition-all hover:border-primary/20 hover:shadow-md hover:-translate-y-0.5 cursor-pointer"
									role="button"
									tabindex="0"
									onclick={() => { selectedTemplate = template; showPreviewModal = true; }}
									onkeydown={(e) => { if (e.key === 'Enter') { selectedTemplate = template; showPreviewModal = true; } }}
								>
									<div class="absolute top-3 right-3">
										<button 
											class="p-1.5 rounded-md hover:bg-muted transition-colors"
											onclick={(e) => toggleFavorite(e, template)}
											type="button"
										>
											<Icon icon="hugeicons:star" class="size-4 text-amber-500 fill-amber-500" />
										</button>
									</div>
									
									<div class="flex items-start gap-3 mb-3">
										<div class="size-12 rounded-xl {template.color} flex items-center justify-center shrink-0">
											<Icon icon={template.icon} class="size-6 text-white" />
										</div>
										<div class="flex-1 min-w-0 pt-1">
											<h3 class="font-semibold truncate pr-6">{template.name}</h3>
											<p class="text-xs text-muted-foreground capitalize">{template.category}</p>
										</div>
									</div>
									
									<p class="text-sm text-muted-foreground line-clamp-2 mb-3">{template.description}</p>
									
									<div class="flex items-center justify-between text-xs text-muted-foreground">
										<span class="flex items-center gap-1">
											<Icon icon="hugeicons:send-01" class="size-3" />
											{template.usageCount} uses
										</span>
										<span>{template.lastUsed}</span>
									</div>
								</div>
							{/each}
						</div>
					</div>
				{/if}

				<div>
					<div class="flex items-center justify-between mb-4">
						<h2 class="text-lg font-semibold">
							{#if searchQuery}
								Search Results
							{:else if categoryFilter !== 'all'}
								{categories.find(c => c.value === categoryFilter)?.label}
							{:else if typeFilter !== 'all'}
								{typeFilter.charAt(0).toUpperCase() + typeFilter.slice(1)} Templates
							{:else}
								All Templates
							{/if}
						</h2>
						<span class="text-sm text-muted-foreground">{filteredTemplates().length} templates</span>
					</div>
					
					{#if filteredTemplates().length === 0}
						<div class="flex flex-col items-center justify-center py-20 text-center">
							<div class="size-20 rounded-full bg-muted flex items-center justify-center mb-4">
								<Icon icon="hugeicons:layout-template" class="size-10 text-muted-foreground" />
							</div>
							<h3 class="text-lg font-medium">
								{#if searchQuery}
									No templates match your search
								{:else}
									No templates in this category
								{/if}
							</h3>
							<p class="text-muted-foreground mt-1 max-w-sm">
								{#if searchQuery}
									Try adjusting your search terms
								{:else}
									Create a new template to get started
								{/if}
							</p>
							<Button class="mt-4 gap-2" onclick={() => showCreateModal = true}>
								<Icon icon="hugeicons:plus-sign" class="size-4" />
								Create Template
							</Button>
						</div>
					{:else if viewMode === 'grid'}
						<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
							{#each filteredTemplates() as template (template.id)}
								<div
									class="group relative rounded-xl border border-border bg-card p-4 transition-all hover:border-primary/20 hover:shadow-md hover:-translate-y-0.5 cursor-pointer"
									role="button"
									tabindex="0"
									onclick={() => { selectedTemplate = template; showPreviewModal = true; }}
									onkeydown={(e) => { if (e.key === 'Enter') { selectedTemplate = template; showPreviewModal = true; } }}
								>
									<div class="absolute top-3 right-3 opacity-0 group-hover:opacity-100 transition-opacity">
										<button 
											class="p-1.5 rounded-md hover:bg-muted transition-colors"
											onclick={(e) => toggleFavorite(e, template)}
											type="button"
										>
											<Icon icon="hugeicons:star" class="size-4 {template.favorite ? 'text-amber-500 fill-amber-500' : 'text-muted-foreground'}" />
										</button>
									</div>
									
									<div class="flex items-start gap-3 mb-3">
										<div class="size-12 rounded-xl {template.color} flex items-center justify-center shrink-0">
											<Icon icon={template.icon} class="size-6 text-white" />
										</div>
										<div class="flex-1 min-w-0 pt-1">
											<h3 class="font-semibold truncate pr-6">{template.name}</h3>
											<p class="text-xs text-muted-foreground capitalize">{template.category}</p>
										</div>
									</div>
									
									<p class="text-sm text-muted-foreground line-clamp-2 mb-3">{template.description}</p>
									
									<div class="flex flex-wrap gap-1 mb-3">
										{#each template.variables.slice(0, 3) as variable}
											<code class="text-[10px] px-1.5 py-0.5 rounded bg-muted">{`{{${variable}}}`}</code>
										{/each}
										{#if template.variables.length > 3}
											<code class="text-[10px] px-1.5 py-0.5 rounded bg-muted">+{template.variables.length - 3}</code>
										{/if}
									</div>
									
									<div class="flex items-center justify-between text-xs text-muted-foreground">
										<span class="flex items-center gap-1">
											<Icon icon="hugeicons:send-01" class="size-3" />
											{template.usageCount} uses
										</span>
										<span>{template.lastUsed}</span>
									</div>
								</div>
							{/each}
						</div>
					{:else}
						<div class="rounded-xl border border-border bg-card overflow-hidden">
							<div class="grid grid-cols-12 gap-4 px-4 py-3 text-sm font-medium text-muted-foreground border-b border-border bg-muted/50">
								<div class="col-span-4">Name</div>
								<div class="col-span-2">Type</div>
								<div class="col-span-2">Category</div>
								<div class="col-span-2">Usage</div>
								<div class="col-span-2">Last Used</div>
							</div>
							{#each filteredTemplates() as template (template.id)}
								<div
									class="group grid grid-cols-12 gap-4 px-4 py-3 items-center border-b border-border last:border-b-0 transition-colors hover:bg-muted/30 cursor-pointer"
									role="button"
									tabindex="0"
									onclick={() => { selectedTemplate = template; showPreviewModal = true; }}
									onkeydown={(e) => { if (e.key === 'Enter') { selectedTemplate = template; showPreviewModal = true; } }}
								>
									<div class="col-span-4 flex items-center gap-3">
										<div class="size-10 rounded-lg {template.color} flex items-center justify-center shrink-0">
											<Icon icon={template.icon} class="size-5 text-white" />
										</div>
										<div class="min-w-0">
											<div class="flex items-center gap-2">
												<p class="font-medium truncate">{template.name}</p>
												{#if template.favorite}
													<Icon icon="hugeicons:star" class="size-3 text-amber-500 fill-amber-500" />
												{/if}
											</div>
											<p class="text-xs text-muted-foreground truncate">{template.description}</p>
										</div>
									</div>
									<div class="col-span-2">
										<span class="inline-flex items-center gap-1.5 px-2 py-1 rounded-md bg-muted text-xs capitalize">
											<Icon icon={template.type === 'email' ? 'hugeicons:mail-01' : template.type === 'document' ? 'hugeicons:file-01' : 'hugeicons:folder-kanban'} class="size-3" />
											{template.type}
										</span>
									</div>
									<div class="col-span-2 text-sm capitalize">{template.category}</div>
									<div class="col-span-2 text-sm text-muted-foreground">{template.usageCount} times</div>
									<div class="col-span-2 flex items-center justify-between">
										<span class="text-sm text-muted-foreground">{template.lastUsed}</span>
										<button 
											class="opacity-0 group-hover:opacity-100 transition-opacity p-1 hover:bg-muted rounded"
											onclick={(e) => { e.stopPropagation(); toggleFavorite(e, template); }}
											type="button"
										>
											<Icon icon="hugeicons:star" class="size-4 {template.favorite ? 'text-amber-500 fill-amber-500' : 'text-muted-foreground'}" />
										</button>
									</div>
								</div>
							{/each}
						</div>
					{/if}
				</div>
			</div>
		</div>
	</div>
</div>

<Modal
	bind:open={showCreateModal}
	title="Create New Template"
	description="Create a reusable template"
	size="lg"
>
	<form onsubmit={handleCreateTemplate} class="space-y-4">
		<div class="grid grid-cols-2 gap-4">
			<div class="space-y-2">
				<Label for="template-name">Template Name</Label>
				<Input id="template-name" placeholder="e.g., Client Onboarding" bind:value={newTemplate.name} />
			</div>
			<div class="space-y-2">
				<Label for="template-type">Type</Label>
				<select id="template-type" bind:value={newTemplate.type} class="w-full h-10 rounded-md border border-input bg-background px-3 py-1 text-sm">
					<option value="email">Email</option>
					<option value="document">Document</option>
					<option value="project">Project</option>
				</select>
			</div>
		</div>
		<div class="space-y-2">
			<Label for="template-category">Category</Label>
			<select id="template-category" bind:value={newTemplate.category} class="w-full h-10 rounded-md border border-input bg-background px-3 py-1 text-sm">
				<option value="onboarding">Onboarding</option>
				<option value="proposals">Proposals</option>
				<option value="follow-ups">Follow-ups</option>
				<option value="invoices">Invoices</option>
				<option value="contracts">Contracts</option>
				<option value="updates">Updates</option>
				<option value="projects">Projects</option>
				<option value="general">General</option>
			</select>
		</div>
		<div class="space-y-2">
			<Label for="template-description">Description</Label>
			<Input id="template-description" placeholder="Brief description of this template..." bind:value={newTemplate.description} />
		</div>
		<div class="space-y-2">
			<Label for="template-content">Content</Label>
			<div class="rounded-md border border-input overflow-hidden">
				<div class="flex items-center gap-1 border-b border-input bg-muted/50 p-2">
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:text-bold" class="size-4" />
					</Button>
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:text-italic" class="size-4" />
					</Button>
					<Separator orientation="vertical" class="h-5 mx-1" />
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:variable" class="size-4" />
					</Button>
				</div>
				<textarea
					id="template-content"
					placeholder="Template content... Use {{variable_name}} for dynamic fields."
					bind:value={newTemplate.content}
					rows={8}
					class="w-full resize-none bg-background px-3 py-2 text-sm outline-none"
				></textarea>
			</div>
		</div>
		<div class="rounded-lg bg-muted/50 p-3">
			<p class="text-sm font-medium mb-2">Quick Variables</p>
			<div class="flex flex-wrap gap-2">
				{#each ['client_name', 'project_name', 'company_name', 'sender_name', 'date'] as varName}
					<button
						type="button"
						class="text-xs px-2 py-1 rounded bg-background border border-border hover:border-primary transition-colors"
						onclick={() => newTemplate.content += ` {{${varName}}}`}
					>
						{`{{${varName}}}`}
					</button>
				{/each}
			</div>
		</div>
	</form>
	
	{#snippet footer()}
		<Button variant="outline" onclick={() => showCreateModal = false}>Cancel</Button>
		<Button type="submit" onclick={handleCreateTemplate} disabled={!newTemplate.name.trim()}>
			<Icon icon="hugeicons:plus-sign" class="size-4 mr-2" />
			Create Template
		</Button>
	{/snippet}
</Modal>

<Modal
	bind:open={showPreviewModal}
	title={selectedTemplate?.name || 'Preview'}
	description={selectedTemplate?.description}
	size="xl"
>
	{#if selectedTemplate}
		<div class="flex items-start gap-4 p-4 border-b border-border">
			<div class="size-16 rounded-xl {selectedTemplate.color} flex items-center justify-center shrink-0">
				<Icon icon={selectedTemplate.icon} class="size-8 text-white" />
			</div>
			<div class="flex-1 min-w-0">
				<div class="flex items-center gap-2 mb-1">
					<span class="inline-flex items-center gap-1.5 px-2 py-0.5 rounded-md bg-muted text-xs capitalize">
						{selectedTemplate.type}
					</span>
					<span class="text-xs text-muted-foreground">•</span>
					<span class="text-xs text-muted-foreground capitalize">{selectedTemplate.category}</span>
				</div>
				<h3 class="text-lg font-semibold">{selectedTemplate.name}</h3>
				<p class="text-sm text-muted-foreground">{selectedTemplate.description}</p>
			</div>
			<div class="flex gap-2">
				<Button variant="outline" size="icon" onclick={() => { selectedTemplate && toggleFavorite(new Event('click'), selectedTemplate); }}>
					<Icon icon="hugeicons:star" class="size-4 {selectedTemplate.favorite ? 'text-amber-500 fill-amber-500' : ''}" />
				</Button>
				<Button variant="outline" size="icon" onclick={() => { showPreviewModal = false; showEditModal = true; }}>
					<Icon icon="hugeicons:pencil-edit-01" class="size-4" />
				</Button>
				<Button variant="outline" size="icon" class="text-destructive hover:text-destructive" onclick={() => { showPreviewModal = false; showDeleteModal = true; }}>
					<Icon icon="hugeicons:delete-01" class="size-4" />
				</Button>
			</div>
		</div>
		
		<div class="p-6">
			<div class="rounded-lg border border-border bg-muted/30 p-6 mb-6">
				<p class="text-sm text-muted-foreground mb-2">Preview:</p>
				<div class="prose prose-sm max-w-none dark:prose-invert">
					{@html selectedTemplate.content}
				</div>
			</div>
			
			<div class="space-y-4">
				<div>
					<p class="text-sm font-medium mb-2">Variables</p>
					<div class="flex flex-wrap gap-2">
						{#each selectedTemplate.variables as variable}
							<code class="text-sm px-2 py-1 rounded bg-muted">{`{{${variable}}}`}</code>
						{/each}
					</div>
				</div>
				
				<div class="flex items-center gap-4 text-sm text-muted-foreground">
					<span class="flex items-center gap-1">
						<Icon icon="hugeicons:send-01" class="size-4" />
						Used {selectedTemplate.usageCount} times
					</span>
					<span>•</span>
					<span>Last used {selectedTemplate.lastUsed}</span>
					<span>•</span>
					<span>Created {selectedTemplate.createdAt}</span>
				</div>
			</div>
		</div>
	{/if}
	
	{#snippet footer()}
		<Button variant="outline" onclick={() => showPreviewModal = false}>Close</Button>
		<div class="flex gap-2">
			<Button variant="outline" class="gap-2">
				<Icon icon="hugeicons:copy-01" class="size-4" />
				Duplicate
			</Button>
			<Button class="gap-2">
				<Icon icon="hugeicons:pencil-edit-01" class="size-4" />
				Use Template
			</Button>
		</div>
	{/snippet}
</Modal>

<Modal
	bind:open={showDeleteModal}
	title="Delete Template"
	description="Are you sure you want to delete this template? This action cannot be undone."
	size="sm"
>
	<p class="text-sm text-muted-foreground">
		Template "{selectedTemplate?.name}" will be permanently deleted.
	</p>
	
	{#snippet footer()}
		<Button variant="outline" onclick={() => showDeleteModal = false}>Cancel</Button>
		<Button variant="destructive" onclick={handleDeleteTemplate}>
			<Icon icon="hugeicons:delete-01" class="size-4 mr-2" />
			Delete
		</Button>
	{/snippet}
</Modal>
