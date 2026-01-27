<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';

	let searchQuery = $state('');
	let statusFilter = $state('all');
	let viewMode = $state<'grid' | 'list'>('grid');
	let selectedClient = $state<Client | null>(null);
	let showClientDrawer = $state(false);
	let showNewClientModal = $state(false);

	let newClient = $state({
		name: '',
		email: '',
		company: '',
		phone: '',
		website: '',
		status: 'lead' as const,
		notes: ''
	});

	function handleCreateClient(e: Event) {
		e.preventDefault();
		console.log('Creating client:', newClient);
		newClient = {
			name: '',
			email: '',
			company: '',
			phone: '',
			website: '',
			status: 'lead',
			notes: ''
		};
		showNewClientModal = false;
	}

	interface Client {
		id: string;
		name: string;
		email: string;
		company: string;
		avatar: string;
		status: 'active' | 'lead' | 'inactive' | 'churned';
		totalRevenue: number;
		outstanding: number;
		projects: { active: number; total: number };
		lastContact: string;
		tags: string[];
		phone?: string;
		website?: string;
		address?: string;
		notes?: string;
	}

	const clients: Client[] = [
		{
			id: '1',
			name: 'Sarah Chen',
			email: 'sarah@acmecorp.com',
			company: 'Acme Corporation',
			avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Acme',
			status: 'active',
			totalRevenue: 125000,
			outstanding: 15000,
			projects: { active: 3, total: 8 },
			lastContact: '2 hours ago',
			tags: ['Enterprise', 'Tech'],
			phone: '+1 (555) 123-4567',
			website: 'https://acmecorp.com',
			address: '123 Market St, San Francisco, CA',
			notes: 'Key decision maker. Prefers email communication. Renewal coming up in Q2.'
		},
		{
			id: '2',
			name: 'Michael Ross',
			email: 'mike@techstart.io',
			company: 'TechStart Inc',
			avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=TechStart',
			status: 'active',
			totalRevenue: 85000,
			outstanding: 0,
			projects: { active: 2, total: 5 },
			lastContact: '1 day ago',
			tags: ['Startup', 'SaaS'],
			phone: '+1 (555) 987-6543',
			website: 'https://techstart.io',
			notes: 'Fast growing startup. Budget conscious but values quality.'
		},
		{
			id: '3',
			name: 'Emily Watson',
			email: 'emily@greenleaf.co',
			company: 'Green Leaf Co',
			avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=GreenLeaf',
			status: 'lead',
			totalRevenue: 0,
			outstanding: 0,
			projects: { active: 0, total: 0 },
			lastContact: '3 days ago',
			tags: ['Sustainability', 'Retail'],
			phone: '+1 (555) 456-7890',
			website: 'https://greenleaf.co',
			notes: 'Interested in brand redesign. Proposal sent, awaiting feedback.'
		},
		{
			id: '4',
			name: 'David Park',
			email: 'david@globaldynamics.com',
			company: 'Global Dynamics',
			avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Global',
			status: 'active',
			totalRevenue: 210000,
			outstanding: 45000,
			projects: { active: 1, total: 12 },
			lastContact: '5 hours ago',
			tags: ['Enterprise', 'Consulting'],
			phone: '+1 (555) 234-5678',
			website: 'https://globaldynamics.com',
			notes: 'Long-term client. Multiple departments, complex approvals.'
		},
		{
			id: '5',
			name: 'Lisa Thompson',
			email: 'lisa@innovateLabs.com',
			company: 'InnovateLabs',
			avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Innovate',
			status: 'inactive',
			totalRevenue: 45000,
			outstanding: 0,
			projects: { active: 0, total: 3 },
			lastContact: '3 months ago',
			tags: ['Agency', 'Marketing'],
			phone: '+1 (555) 876-5432',
			notes: 'Project completed. Following up for maintenance contract.'
		},
		{
			id: '6',
			name: 'James Wilson',
			email: 'james@nexgen.tech',
			company: 'NexGen Technologies',
			avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=NexGen',
			status: 'lead',
			totalRevenue: 0,
			outstanding: 0,
			projects: { active: 0, total: 0 },
			lastContact: '1 week ago',
			tags: ['AI', 'Enterprise'],
			phone: '+1 (555) 345-6789',
			website: 'https://nexgen.tech',
			notes: 'Referred by Sarah Chen. Initial call scheduled for next week.'
		}
	];

	const stats = [
		{
			label: 'Total Clients',
			value: '24',
			icon: 'lucide:users',
			color: 'bg-blue-500/10 text-blue-600'
		},
		{
			label: 'Active',
			value: '18',
			icon: 'lucide:user-check',
			color: 'bg-emerald-500/10 text-emerald-600'
		},
		{
			label: 'Leads',
			value: '4',
			icon: 'lucide:user-plus',
			color: 'bg-amber-500/10 text-amber-600'
		},
		{
			label: 'Outstanding',
			value: '$60,000',
			icon: 'lucide:dollar-sign',
			color: 'bg-violet-500/10 text-violet-600'
		}
	];

	const statusOptions = [
		{ value: 'all', label: 'All Clients', count: clients.length },
		{
			value: 'active',
			label: 'Active',
			count: clients.filter((c) => c.status === 'active').length
		},
		{ value: 'lead', label: 'Leads', count: clients.filter((c) => c.status === 'lead').length },
		{
			value: 'inactive',
			label: 'Inactive',
			count: clients.filter((c) => c.status === 'inactive').length
		}
	];

	function getStatusColor(status: string) {
		switch (status) {
			case 'active':
				return 'bg-emerald-500/10 text-emerald-600 border-emerald-500/20';
			case 'lead':
				return 'bg-amber-500/10 text-amber-600 border-amber-500/20';
			case 'inactive':
				return 'bg-slate-500/10 text-slate-600 border-slate-500/20';
			case 'churned':
				return 'bg-rose-500/10 text-rose-600 border-rose-500/20';
			default:
				return 'bg-muted';
		}
	}

	function formatCurrency(amount: number) {
		return new Intl.NumberFormat('en-US', {
			style: 'currency',
			currency: 'USD',
			maximumFractionDigits: 0
		}).format(amount);
	}

	function openClientDrawer(client: Client) {
		selectedClient = client;
		showClientDrawer = true;
	}

	function closeClientDrawer() {
		showClientDrawer = false;
		setTimeout(() => (selectedClient = null), 200);
	}

	const filteredClients = $derived(() => {
		let filtered = clients;
		if (statusFilter !== 'all') {
			filtered = filtered.filter((c) => c.status === statusFilter);
		}
		if (searchQuery) {
			const q = searchQuery.toLowerCase();
			filtered = filtered.filter(
				(c) =>
					c.name.toLowerCase().includes(q) ||
					c.company.toLowerCase().includes(q) ||
					c.email.toLowerCase().includes(q)
			);
		}
		return filtered;
	});
</script>

<svelte:head>
	<title>Clients - Artemis</title>
</svelte:head>

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="px-6 py-8">
		<div class="mb-8 flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
			<div>
				<h1 class="text-3xl font-bold tracking-tight">Clients</h1>
				<p class="mt-1 text-muted-foreground">
					Manage your clients, track communications, and view project history.
				</p>
			</div>
			<Button class="gap-2" onclick={() => (showNewClientModal = true)}>
				<Icon icon="lucide:user-plus" class="size-4" />
				Add Client
			</Button>
		</div>

		<div class="mb-8 grid gap-4 md:grid-cols-2 lg:grid-cols-4">
			{#each stats as stat}
				<div class="rounded-2xl border border-border bg-card p-5 transition-shadow hover:shadow-lg">
					<div class="flex items-center gap-4">
						<div class="flex size-12 items-center justify-center rounded-xl {stat.color}">
							<Icon icon={stat.icon} class="size-6" />
						</div>
						<div>
							<p class="text-sm text-muted-foreground">{stat.label}</p>
							<p class="text-2xl font-bold">{stat.value}</p>
						</div>
					</div>
				</div>
			{/each}
		</div>

		<div class="overflow-hidden rounded-2xl border border-border bg-card">
			<div class="flex flex-col gap-4 border-b border-border p-4 sm:flex-row">
				<div class="relative flex-1">
					<Icon
						icon="lucide:search"
						class="absolute top-1/2 left-3 size-4 -translate-y-1/2 text-muted-foreground"
					/>
					<Input
						placeholder="Search clients by name, company, or email..."
						class="pl-10"
						bind:value={searchQuery}
					/>
				</div>
				<div class="flex gap-2">
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
					<Separator orientation="vertical" class="h-9" />
					<div class="flex rounded-lg border border-border p-1">
						<button
							class="rounded p-1.5 transition-colors {viewMode === 'grid'
								? 'bg-muted'
								: 'text-muted-foreground hover:text-foreground'}"
							onclick={() => (viewMode = 'grid')}
						>
							<Icon icon="lucide:grid-3x3" class="size-4" />
						</button>
						<button
							class="rounded p-1.5 transition-colors {viewMode === 'list'
								? 'bg-muted'
								: 'text-muted-foreground hover:text-foreground'}"
							onclick={() => (viewMode = 'list')}
						>
							<Icon icon="lucide:list" class="size-4" />
						</button>
					</div>
				</div>
			</div>

			{#if filteredClients().length === 0}
				<div class="p-12 text-center">
					<div class="mb-4 inline-flex size-16 items-center justify-center rounded-full bg-muted">
						<Icon icon="lucide:users" class="size-8 text-muted-foreground" />
					</div>
					<h3 class="text-lg font-semibold">No clients found</h3>
					<p class="mt-1 text-muted-foreground">Try adjusting your search or filters.</p>
				</div>
			{:else if viewMode === 'grid'}
				<div class="grid gap-4 p-4 md:grid-cols-2 xl:grid-cols-3">
					{#each filteredClients() as client}
						<div
							class="group cursor-pointer rounded-xl border border-border bg-card p-5 transition-all hover:border-primary/20 hover:shadow-lg"
							onclick={() => openClientDrawer(client)}
						>
							<div class="mb-4 flex items-start justify-between">
								<div class="flex items-center gap-3">
									<img
										src={client.avatar}
										alt={client.name}
										class="size-12 rounded-full bg-muted"
									/>
									<div>
										<h3 class="font-semibold transition-colors group-hover:text-primary">
											{client.name}
										</h3>
										<p class="text-sm text-muted-foreground">{client.company}</p>
									</div>
								</div>
								<span
									class="rounded-full border px-2 py-1 text-xs font-medium capitalize {getStatusColor(
										client.status
									)}"
								>
									{client.status}
								</span>
							</div>

							<div class="mb-4 grid grid-cols-2 gap-4">
								<div>
									<p class="text-xs text-muted-foreground">Total Revenue</p>
									<p class="font-semibold">{formatCurrency(client.totalRevenue)}</p>
								</div>
								<div>
									<p class="text-xs text-muted-foreground">Outstanding</p>
									<p class="font-semibold {client.outstanding > 0 ? 'text-amber-600' : ''}">
										{formatCurrency(client.outstanding)}
									</p>
								</div>
							</div>

							<div class="flex items-center justify-between border-t border-border pt-4">
								<div class="flex items-center gap-1 text-sm text-muted-foreground">
									<Icon icon="lucide:folder" class="size-4" />
									<span>{client.projects.active} active / {client.projects.total} total</span>
								</div>
								<span class="text-xs text-muted-foreground">{client.lastContact}</span>
							</div>

							<div class="mt-4 flex gap-2">
								{#each client.tags as tag}
									<span class="rounded-md bg-muted px-2 py-1 text-xs">{tag}</span>
								{/each}
							</div>

							<div
								class="mt-4 flex gap-2 border-t border-border pt-4 opacity-0 transition-opacity group-hover:opacity-100"
							>
								<Button
									variant="outline"
									size="sm"
									class="flex-1"
									onclick={(e) => {
										e.stopPropagation();
									}}
								>
									<Icon icon="lucide:mail" class="mr-2 size-4" />
									Email
								</Button>
								<Button
									variant="outline"
									size="sm"
									class="flex-1"
									onclick={(e) => {
										e.stopPropagation();
									}}
								>
									<Icon icon="lucide:plus" class="mr-2 size-4" />
									Project
								</Button>
							</div>
						</div>
					{/each}
				</div>
			{:else}
				<div class="divide-y divide-border">
					{#each filteredClients() as client}
						<div
							class="group flex cursor-pointer items-center gap-4 p-4 transition-colors hover:bg-muted/30"
							onclick={() => openClientDrawer(client)}
						>
							<img
								src={client.avatar}
								alt={client.name}
								class="size-10 shrink-0 rounded-full bg-muted"
							/>
							<div class="grid min-w-0 flex-1 grid-cols-2 items-center gap-4 md:grid-cols-4">
								<div>
									<p class="truncate font-semibold transition-colors group-hover:text-primary">
										{client.name}
									</p>
									<p class="truncate text-sm text-muted-foreground">{client.company}</p>
								</div>
								<div class="hidden md:block">
									<p class="text-sm font-medium">{formatCurrency(client.totalRevenue)}</p>
									<p class="text-xs text-muted-foreground">Total Revenue</p>
								</div>
								<div class="hidden md:block">
									<p class="text-sm font-medium">{client.projects.active} active</p>
									<p class="text-xs text-muted-foreground">Projects</p>
								</div>
								<div class="flex items-center justify-end gap-2">
									<span
										class="rounded-full border px-2 py-1 text-xs font-medium capitalize {getStatusColor(
											client.status
										)}"
									>
										{client.status}
									</span>
									<Button
										variant="ghost"
										size="icon"
										class="opacity-0 transition-opacity group-hover:opacity-100"
										onclick={(e) => {
											e.stopPropagation();
										}}
									>
										<Icon icon="lucide:more-vertical" class="size-4" />
									</Button>
								</div>
							</div>
						</div>
					{/each}
				</div>
			{/if}
		</div>
	</div>
</div>

{#if showClientDrawer && selectedClient}
	<div class="fixed inset-0 z-50">
		<div
			class="absolute inset-0 bg-background/80 backdrop-blur-sm"
			onclick={closeClientDrawer}
		></div>
		<div
			class="absolute top-0 right-0 bottom-0 w-full max-w-2xl overflow-y-auto border-l border-border bg-background shadow-2xl"
		>
			<div
				class="sticky top-0 flex items-center justify-between border-b border-border bg-background/95 p-4 backdrop-blur"
			>
				<div class="flex items-center gap-3">
					<Button variant="ghost" size="icon" onclick={closeClientDrawer}>
						<Icon icon="lucide:x" class="size-5" />
					</Button>
					<span class="font-semibold">Client Details</span>
				</div>
				<div class="flex gap-2">
					<Button variant="outline" size="sm">
						<Icon icon="lucide:edit" class="mr-2 size-4" />
						Edit
					</Button>
					<Button variant="outline" size="sm" class="text-destructive hover:text-destructive">
						<Icon icon="lucide:trash-2" class="mr-2 size-4" />
					</Button>
				</div>
			</div>

			<div class="space-y-8 p-6">
				<div class="flex items-start gap-4">
					<img
						src={selectedClient.avatar}
						alt={selectedClient.name}
						class="size-20 rounded-2xl bg-muted"
					/>
					<div class="flex-1">
						<div class="flex items-start justify-between">
							<div>
								<h2 class="text-2xl font-bold">{selectedClient.name}</h2>
								<p class="text-lg text-muted-foreground">{selectedClient.company}</p>
							</div>
							<span
								class="rounded-full border px-3 py-1 text-sm font-medium capitalize {getStatusColor(
									selectedClient.status
								)}"
							>
								{selectedClient.status}
							</span>
						</div>
						<div class="mt-3 flex gap-2">
							{#each selectedClient.tags as tag}
								<span class="rounded-full bg-muted px-3 py-1 text-sm">{tag}</span>
							{/each}
						</div>
					</div>
				</div>

				<div class="grid grid-cols-3 gap-4">
					<div class="rounded-xl border border-border p-4 text-center">
						<p class="text-2xl font-bold">{formatCurrency(selectedClient.totalRevenue)}</p>
						<p class="text-sm text-muted-foreground">Total Revenue</p>
					</div>
					<div class="rounded-xl border border-border p-4 text-center">
						<p class="text-2xl font-bold {selectedClient.outstanding > 0 ? 'text-amber-600' : ''}">
							{formatCurrency(selectedClient.outstanding)}
						</p>
						<p class="text-sm text-muted-foreground">Outstanding</p>
					</div>
					<div class="rounded-xl border border-border p-4 text-center">
						<p class="text-2xl font-bold">{selectedClient.projects.total}</p>
						<p class="text-sm text-muted-foreground">Projects</p>
					</div>
				</div>

				<div class="space-y-4 rounded-xl border border-border p-6">
					<h3 class="font-semibold">Contact Information</h3>
					<div class="space-y-3">
						<div class="flex items-center gap-3">
							<Icon icon="lucide:mail" class="size-5 text-muted-foreground" />
							<a href="mailto:{selectedClient.email}" class="text-primary hover:underline"
								>{selectedClient.email}</a
							>
						</div>
						{#if selectedClient.phone}
							<div class="flex items-center gap-3">
								<Icon icon="lucide:phone" class="size-5 text-muted-foreground" />
								<a href="tel:{selectedClient.phone}" class="hover:underline"
									>{selectedClient.phone}</a
								>
							</div>
						{/if}
						{#if selectedClient.website}
							<div class="flex items-center gap-3">
								<Icon icon="lucide:globe" class="size-5 text-muted-foreground" />
								<a
									href={selectedClient.website}
									target="_blank"
									class="text-primary hover:underline">{selectedClient.website}</a
								>
							</div>
						{/if}
						{#if selectedClient.address}
							<div class="flex items-center gap-3">
								<Icon icon="lucide:map-pin" class="size-5 text-muted-foreground" />
								<span class="text-muted-foreground">{selectedClient.address}</span>
							</div>
						{/if}
					</div>
				</div>

				{#if selectedClient.notes}
					<div class="rounded-xl border border-border p-6">
						<h3 class="mb-3 font-semibold">Notes</h3>
						<p class="text-muted-foreground">{selectedClient.notes}</p>
					</div>
				{/if}

				<div class="overflow-hidden rounded-xl border border-border">
					<div class="flex items-center justify-between border-b border-border p-4">
						<h3 class="font-semibold">Recent Projects</h3>
						<Button variant="ghost" size="sm">
							<Icon icon="lucide:plus" class="mr-2 size-4" />
							New Project
						</Button>
					</div>
					<div class="divide-y divide-border">
						{#each [{ name: 'Website Redesign', status: 'In Progress', budget: '$15,000', progress: 75 }, { name: 'Brand Identity', status: 'Completed', budget: '$8,500', progress: 100 }, { name: 'Marketing Campaign', status: 'Planning', budget: '$25,000', progress: 10 }] as project}
							<div
								class="flex items-center justify-between p-4 transition-colors hover:bg-muted/30"
							>
								<div>
									<p class="font-medium">{project.name}</p>
									<p class="text-sm text-muted-foreground">Budget: {project.budget}</p>
								</div>
								<div class="text-right">
									<span class="text-sm">{project.status}</span>
									<div class="mt-1 h-1.5 w-24 overflow-hidden rounded-full bg-muted">
										<div
											class="h-full rounded-full bg-primary"
											style="width: {project.progress}%"
										></div>
									</div>
								</div>
							</div>
						{/each}
					</div>
				</div>

				<div class="overflow-hidden rounded-xl border border-border">
					<div class="border-b border-border p-4">
						<h3 class="font-semibold">Recent Activity</h3>
					</div>
					<div class="divide-y divide-border">
						{#each [{ action: 'Invoice sent', date: '2 days ago', amount: '$5,000' }, { action: 'Meeting completed', date: '1 week ago' }, { action: 'Proposal accepted', date: '2 weeks ago' }] as activity}
							<div class="flex items-center justify-between p-4">
								<div class="flex items-center gap-3">
									<div class="flex size-8 items-center justify-center rounded-lg bg-muted">
										<Icon icon="lucide:activity" class="size-4 text-muted-foreground" />
									</div>
									<span>{activity.action}</span>
								</div>
								<div class="text-right">
									{#if activity.amount}
										<p class="font-medium">{activity.amount}</p>
									{/if}
									<p class="text-sm text-muted-foreground">{activity.date}</p>
								</div>
							</div>
						{/each}
					</div>
				</div>
			</div>
		</div>
	</div>
{/if}

<Modal
	bind:open={showNewClientModal}
	title="Add New Client"
	description="Fill in the details below to create a new client."
	size="lg"
>
	<form id="new-client-form" onsubmit={handleCreateClient} class="space-y-4">
		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="client-name">Full Name *</Label>
				<Input id="client-name" placeholder="John Doe" bind:value={newClient.name} required />
			</div>
			<div class="space-y-2">
				<Label for="client-email">Email *</Label>
				<Input
					id="client-email"
					type="email"
					placeholder="john@example.com"
					bind:value={newClient.email}
					required
				/>
			</div>
		</div>

		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="client-company">Company</Label>
				<Input id="client-company" placeholder="Acme Inc." bind:value={newClient.company} />
			</div>
			<div class="space-y-2">
				<Label for="client-phone">Phone</Label>
				<Input
					id="client-phone"
					type="tel"
					placeholder="+1 (555) 123-4567"
					bind:value={newClient.phone}
				/>
			</div>
		</div>

		<div class="space-y-2">
			<Label for="client-website">Website</Label>
			<Input
				id="client-website"
				type="url"
				placeholder="https://example.com"
				bind:value={newClient.website}
			/>
		</div>

		<div class="space-y-2">
			<Label for="client-status">Status</Label>
			<select
				id="client-status"
				bind:value={newClient.status}
				class="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-sm transition-colors focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
			>
				<option value="lead">Lead</option>
				<option value="active">Active</option>
				<option value="inactive">Inactive</option>
			</select>
		</div>

		<div class="space-y-2">
			<Label for="client-notes">Notes</Label>
			<textarea
				id="client-notes"
				placeholder="Add any additional information about this client..."
				bind:value={newClient.notes}
				rows={3}
				class="flex w-full resize-none rounded-md border border-input bg-background px-3 py-2 text-sm shadow-sm transition-colors placeholder:text-muted-foreground focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
			></textarea>
		</div>
	</form>

	{#snippet footer()}
		<Button variant="outline" onclick={() => (showNewClientModal = false)}>Cancel</Button>
		<Button type="submit" form="new-client-form">
			<Icon icon="lucide:user-plus" class="mr-2 size-4" />
			Create Client
		</Button>
	{/snippet}
</Modal>
