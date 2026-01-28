<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';
	import { Label } from '$lib/components/ui/label/index.js';

	let viewMode = $state<'list' | 'grid'>('list');
	let statusFilter = $state('all');
	let searchQuery = $state('');
	let showNewInvoiceModal = $state(false);
	let selectedInvoice = $state<Invoice | null>(null);
	let showDeleteModal = $state(false);

	interface Invoice {
		id: string;
		number: string;
		client: {
			name: string;
			company: string;
			avatar: string;
			email: string;
		};
		project?: {
			name: string;
			color: string;
		};
		issueDate: string;
		dueDate: string;
		subtotal: number;
		tax: number;
		total: number;
		paid: number;
		balance: number;
		status: 'draft' | 'sent' | 'viewed' | 'paid' | 'overdue' | 'cancelled';
		items: {
			description: string;
			quantity: number;
			rate: number;
			amount: number;
		}[];
		notes?: string;
		terms?: string;
	}

	const clients = [
		{ name: 'Sarah Chen', company: 'Acme Corporation', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Acme', email: 'sarah@acmecorp.com' },
		{ name: 'Michael Ross', company: 'TechStart Inc', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=TechStart', email: 'mike@techstart.io' },
		{ name: 'Emily Watson', company: 'Green Leaf Co', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=GreenLeaf', email: 'emily@greenleaf.co' },
		{ name: 'David Park', company: 'Global Dynamics', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Global', email: 'david@globaldynamics.com' },
		{ name: 'Lisa Thompson', company: 'InnovateLabs', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Innovate', email: 'lisa@innovatelabs.com' }
	];

	const projects = [
		{ name: 'Website Redesign', color: 'bg-blue-500' },
		{ name: 'Mobile App', color: 'bg-violet-500' },
		{ name: 'Brand Identity', color: 'bg-emerald-500' },
		{ name: 'Marketing Campaign', color: 'bg-amber-500' },
		{ name: 'E-commerce Platform', color: 'bg-rose-500' }
	];

	const invoices: Invoice[] = [
		{
			id: 'inv1',
			number: 'INV-2026-001',
			client: clients[0],
			project: projects[0],
			issueDate: 'Jan 15, 2026',
			dueDate: 'Feb 15, 2026',
			subtotal: 12500,
			tax: 1250,
			total: 13750,
			paid: 13750,
			balance: 0,
			status: 'paid',
			items: [
				{ description: 'Website Design - Phase 1', quantity: 40, rate: 150, amount: 6000 },
				{ description: 'UI/UX Consulting', quantity: 10, rate: 200, amount: 2000 },
				{ description: 'Development Hours', quantity: 30, rate: 150, amount: 4500 }
			],
			notes: 'Thank you for your business!',
			terms: 'Net 30 days'
		},
		{
			id: 'inv2',
			number: 'INV-2026-002',
			client: clients[1],
			project: projects[1],
			issueDate: 'Jan 20, 2026',
			dueDate: 'Feb 20, 2026',
			subtotal: 8750,
			tax: 0,
			total: 8750,
			paid: 4375,
			balance: 4375,
			status: 'sent',
			items: [
				{ description: 'Mobile App UI Design', quantity: 25, rate: 175, amount: 4375 },
				{ description: 'Prototype Development', quantity: 25, rate: 175, amount: 4375 }
			],
			notes: '50% deposit received',
			terms: 'Net 15 days'
		},
		{
			id: 'inv3',
			number: 'INV-2026-003',
			client: clients[2],
			project: projects[2],
			issueDate: 'Jan 25, 2026',
			dueDate: 'Feb 25, 2026',
			subtotal: 5200,
			tax: 520,
			total: 5720,
			paid: 0,
			balance: 5720,
			status: 'viewed',
			items: [
				{ description: 'Logo Design Package', quantity: 1, rate: 2500, amount: 2500 },
				{ description: 'Brand Guidelines', quantity: 1, rate: 1500, amount: 1500 },
				{ description: 'Business Card Design', quantity: 3, rate: 400, amount: 1200 }
			],
			notes: '',
			terms: 'Net 30 days'
		},
		{
			id: 'inv4',
			number: 'INV-2026-004',
			client: clients[3],
			project: projects[3],
			issueDate: 'Jan 10, 2026',
			dueDate: 'Jan 25, 2026',
			subtotal: 15000,
			tax: 1500,
			total: 16500,
			paid: 0,
			balance: 16500,
			status: 'overdue',
			items: [
				{ description: 'Marketing Strategy', quantity: 20, rate: 250, amount: 5000 },
				{ description: 'Ad Creative Design', quantity: 40, rate: 150, amount: 6000 },
				{ description: 'Campaign Management', quantity: 20, rate: 200, amount: 4000 }
			],
			notes: 'Payment reminder sent',
			terms: 'Net 15 days'
		},
		{
			id: 'inv5',
			number: 'INV-2026-005',
			client: clients[4],
			project: projects[4],
			issueDate: 'Jan 28, 2026',
			dueDate: 'Feb 28, 2026',
			subtotal: 3200,
			tax: 0,
			total: 3200,
			paid: 0,
			balance: 3200,
			status: 'draft',
			items: [
				{ description: 'Database Schema Design', quantity: 16, rate: 200, amount: 3200 }
			],
			notes: 'Draft - awaiting approval',
			terms: 'Net 30 days'
		},
		{
			id: 'inv6',
			number: 'INV-2025-089',
			client: clients[0],
			issueDate: 'Dec 15, 2025',
			dueDate: 'Jan 15, 2026',
			subtotal: 8500,
			tax: 850,
			total: 9350,
			paid: 9350,
			balance: 0,
			status: 'paid',
			items: [
				{ description: 'Q4 Consulting Services', quantity: 50, rate: 170, amount: 8500 }
			],
			notes: '',
			terms: 'Net 30 days'
		}
	];

	const statusOptions = [
		{ value: 'all', label: 'All Invoices', count: invoices.length },
		{ value: 'draft', label: 'Draft', count: invoices.filter(i => i.status === 'draft').length },
		{ value: 'sent', label: 'Sent', count: invoices.filter(i => i.status === 'sent').length },
		{ value: 'viewed', label: 'Viewed', count: invoices.filter(i => i.status === 'viewed').length },
		{ value: 'paid', label: 'Paid', count: invoices.filter(i => i.status === 'paid').length },
		{ value: 'overdue', label: 'Overdue', count: invoices.filter(i => i.status === 'overdue').length }
	];

	let newInvoice = $state({
		client: '',
		project: '',
		issueDate: new Date().toISOString().split('T')[0],
		dueDate: '',
		items: [{ description: '', quantity: 1, rate: 0 }],
		notes: '',
		terms: 'Net 30 days'
	});

	function getStatusColor(status: string): string {
		switch (status) {
			case 'draft':
				return 'bg-slate-500/10 text-slate-600 border-slate-500/20';
			case 'sent':
				return 'bg-blue-500/10 text-blue-600 border-blue-500/20';
			case 'viewed':
				return 'bg-violet-500/10 text-violet-600 border-violet-500/20';
			case 'paid':
				return 'bg-emerald-500/10 text-emerald-600 border-emerald-500/20';
			case 'overdue':
				return 'bg-rose-500/10 text-rose-600 border-rose-500/20';
			case 'cancelled':
				return 'bg-gray-500/10 text-gray-600 border-gray-500/20';
			default:
				return 'bg-muted';
		}
	}

	function getStatusIcon(status: string): string {
		switch (status) {
			case 'draft':
				return 'hugeicons:file-01';
			case 'sent':
				return 'hugeicons:send-01';
			case 'viewed':
				return 'hugeicons:eye';
			case 'paid':
				return 'hugeicons:tick-double-01';
			case 'overdue':
				return 'hugeicons:alert-02';
			default:
				return 'hugeicons:file-01';
		}
	}

	function formatCurrency(amount: number): string {
		return new Intl.NumberFormat('en-US', {
			style: 'currency',
			currency: 'USD',
			minimumFractionDigits: 0,
			maximumFractionDigits: 0
		}).format(amount);
	}

	const filteredInvoices = $derived(() => {
		let filtered = invoices;
		if (statusFilter !== 'all') {
			filtered = filtered.filter(i => i.status === statusFilter);
		}
		if (searchQuery) {
			const q = searchQuery.toLowerCase();
			filtered = filtered.filter(
				i =>
					i.number.toLowerCase().includes(q) ||
					i.client.name.toLowerCase().includes(q) ||
					i.client.company.toLowerCase().includes(q)
			);
		}
		return filtered;
	});

	const stats = {
		totalOutstanding: invoices.filter(i => i.status !== 'paid' && i.status !== 'draft' && i.status !== 'cancelled').reduce((acc, i) => acc + i.balance, 0),
		totalOverdue: invoices.filter(i => i.status === 'overdue').reduce((acc, i) => acc + i.balance, 0),
		totalPaid: invoices.filter(i => i.status === 'paid').reduce((acc, i) => acc + i.total, 0),
		paidThisMonth: 24850
	};

	function handleCreateInvoice(e: Event) {
		e.preventDefault();
		console.log('Creating invoice:', newInvoice);
		showNewInvoiceModal = false;
	}

	function addItem() {
		newInvoice.items = [...newInvoice.items, { description: '', quantity: 1, rate: 0 }];
	}

	function removeItem(index: number) {
		newInvoice.items = newInvoice.items.filter((_, i) => i !== index);
	}
</script>

<svelte:head>
	<title>Invoices - Artemis</title>
</svelte:head>

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="px-6 py-4">
		<div class="mb-6 flex items-center justify-end gap-2">
			<Button variant="outline" class="gap-2">
				<Icon icon="hugeicons:download-01" class="size-4" />
				Export
			</Button>
			<Button class="gap-2" onclick={() => showNewInvoiceModal = true}>
				<Icon icon="hugeicons:plus-sign" class="size-4" />
				New Invoice
			</Button>
		</div>

		<div class="grid gap-4 md:grid-cols-4">
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">Outstanding</p>
						<p class="text-2xl font-bold">{formatCurrency(stats.totalOutstanding)}</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-amber-500/10">
						<Icon icon="hugeicons:invoice-01" class="size-5 text-amber-600" />
					</div>
				</div>
			</div>
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">Overdue</p>
						<p class="text-2xl font-bold text-rose-600">{formatCurrency(stats.totalOverdue)}</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-rose-500/10">
						<Icon icon="hugeicons:alert-02" class="size-5 text-rose-600" />
					</div>
				</div>
			</div>
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">Total Paid</p>
						<p class="text-2xl font-bold text-emerald-600">{formatCurrency(stats.totalPaid)}</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-emerald-500/10">
						<Icon icon="hugeicons:tick-double-01" class="size-5 text-emerald-600" />
					</div>
				</div>
			</div>
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">This Month</p>
						<p class="text-2xl font-bold text-blue-600">{formatCurrency(stats.paidThisMonth)}</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-blue-500/10">
						<Icon icon="hugeicons:calendar-01" class="size-5 text-blue-600" />
					</div>
				</div>
			</div>
		</div>

		<div class="mt-8 overflow-hidden rounded-2xl border border-border bg-card">
			<div class="flex flex-col gap-4 border-b border-border p-4">
				<div class="flex flex-col gap-4 lg:flex-row lg:items-center">
					<div class="relative flex-1">
						<Icon
							icon="hugeicons:search-01"
							class="absolute top-1/2 left-3 size-5 -translate-y-1/2 text-muted-foreground"
						/>
						<Input
							placeholder="Search invoices by number, client, or company..."
							class="pl-11"
							bind:value={searchQuery}
						/>
					</div>

					<div class="flex flex-wrap items-center gap-2">
						<Separator orientation="vertical" class="hidden h-9 lg:block" />

						<div class="flex rounded-lg border border-border p-1">
							<button
								class="rounded p-1.5 transition-colors {viewMode === 'list'
									? 'bg-muted text-foreground'
									: 'text-muted-foreground hover:text-foreground'}"
								onclick={() => viewMode = 'list'}
							>
								<Icon icon="hugeicons:menu-square" class="size-5" />
							</button>
							<button
								class="rounded p-1.5 transition-colors {viewMode === 'grid'
									? 'bg-muted text-foreground'
									: 'text-muted-foreground hover:text-foreground'}"
								onclick={() => viewMode = 'grid'}
							>
								<Icon icon="hugeicons:grid" class="size-5" />
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
							onclick={() => statusFilter = option.value}
						>
							{option.label}
							<span class="ml-1 text-xs opacity-70">({option.count})</span>
						</button>
					{/each}
				</div>
			</div>

			{#if filteredInvoices().length === 0}
				<div class="p-12 text-center">
					<div class="mb-4 inline-flex size-16 items-center justify-center rounded-full bg-muted">
						<Icon icon="hugeicons:invoice-01" class="size-8 text-muted-foreground" />
					</div>
					<h3 class="text-lg font-semibold">No invoices found</h3>
					<p class="mt-1 text-muted-foreground">Try adjusting your search or filters.</p>
				</div>

			{:else if viewMode === 'list'}
				<div class="divide-y divide-border">
					{#each filteredInvoices() as invoice}
						<div class="group flex flex-col gap-4 p-4 transition-colors hover:bg-muted/30 sm:flex-row sm:items-center">
							<div class="flex min-w-0 flex-1 items-center gap-3">
								<div class="flex size-10 shrink-0 items-center justify-center rounded-lg bg-primary/10">
									<Icon icon="hugeicons:invoice-01" class="size-5 text-primary" />
								</div>
								<div class="min-w-0">
									<div class="flex items-center gap-2">
										<p class="font-semibold">{invoice.number}</p>
										{#if invoice.project}
											<div class="hidden items-center gap-1 sm:flex">
												<div class="size-2 rounded-full {invoice.project.color}"></div>
												<span class="text-xs text-muted-foreground">{invoice.project.name}</span>
											</div>
										{/if}
									</div>
									<div class="flex items-center gap-2 text-sm text-muted-foreground">
										<span>{invoice.client.company}</span>
										<span>·</span>
										<span>Due {invoice.dueDate}</span>
									</div>
								</div>
							</div>

							<div class="flex flex-wrap items-center gap-4 sm:flex-nowrap">
								<span class="flex items-center gap-1 rounded-full border px-2 py-1 text-xs font-medium capitalize {getStatusColor(invoice.status)}">
									<Icon icon={getStatusIcon(invoice.status)} class="size-3" />
									{invoice.status}
								</span>

								<div class="text-right">
									<p class="font-semibold">{formatCurrency(invoice.total)}</p>
									{#if invoice.balance > 0}
										<p class="text-xs text-rose-600">{formatCurrency(invoice.balance)} due</p>
									{:else}
										<p class="text-xs text-emerald-600">Paid</p>
									{/if}
								</div>

								<div class="flex items-center gap-1">
									<Button variant="ghost" size="icon" class="size-8" href="/app/invoices/{invoice.id}">
										<Icon icon="hugeicons:view" class="size-4" />
									</Button>
									<Button
										variant="ghost"
										size="icon"
										class="size-8 opacity-0 transition-opacity group-hover:opacity-100"
										onclick={() => { selectedInvoice = invoice; showDeleteModal = true; }}
									>
										<Icon icon="hugeicons:more-vertical" class="size-4" />
									</Button>
								</div>
							</div>
						</div>
					{/each}
				</div>

			{:else if viewMode === 'grid'}
				<div class="grid gap-4 p-4 md:grid-cols-2 xl:grid-cols-3">
					{#each filteredInvoices() as invoice}
						<div class="group rounded-xl border border-border bg-card p-5 transition-all hover:border-primary/20 hover:shadow-lg">
							<div class="mb-4 flex items-start justify-between">
								<div class="flex items-center gap-3">
									<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
										<Icon icon="hugeicons:invoice-01" class="size-5 text-primary" />
									</div>
									<div>
										<p class="font-semibold">{invoice.number}</p>
										<p class="text-xs text-muted-foreground">{invoice.issueDate}</p>
									</div>
								</div>
								<span class="flex items-center gap-1 rounded-full border px-2 py-1 text-xs font-medium capitalize {getStatusColor(invoice.status)}">
									<Icon icon={getStatusIcon(invoice.status)} class="size-3" />
									{invoice.status}
								</span>
							</div>

							<div class="mb-4 flex items-center gap-3">
								<img src={invoice.client.avatar} alt={invoice.client.name} class="size-10 rounded-full" />
								<div>
									<p class="font-medium">{invoice.client.company}</p>
									<p class="text-sm text-muted-foreground">{invoice.client.name}</p>
								</div>
							</div>

							{#if invoice.project}
								<div class="mb-4 flex items-center gap-2">
									<div class="size-2 rounded-full {invoice.project.color}"></div>
									<span class="text-sm text-muted-foreground">{invoice.project.name}</span>
								</div>
							{/if}

							<div class="mb-4 flex items-center justify-between rounded-lg bg-muted p-3">
								<div>
									<p class="text-xs text-muted-foreground">Total</p>
									<p class="text-xl font-bold">{formatCurrency(invoice.total)}</p>
								</div>
								<div class="text-right">
									<p class="text-xs text-muted-foreground">Due {invoice.dueDate}</p>
									{#if invoice.balance > 0}
										<p class="text-sm font-medium text-rose-600">{formatCurrency(invoice.balance)} due</p>
									{:else}
										<p class="text-sm font-medium text-emerald-600">Paid in full</p>
									{/if}
								</div>
							</div>

							<div class="flex gap-2">
								<Button variant="outline" size="sm" class="flex-1" href="/app/invoices/{invoice.id}">
									<Icon icon="hugeicons:view" class="mr-2 size-4" />
									View
								</Button>
								{#if invoice.status === 'draft'}
									<Button size="sm" class="flex-1">
										<Icon icon="hugeicons:send-01" class="mr-2 size-4" />
										Send
									</Button>
								{:else if invoice.status === 'overdue'}
									<Button size="sm" variant="secondary" class="flex-1">
										<Icon icon="hugeicons:mail-01" class="mr-2 size-4" />
										Remind
									</Button>
								{:else}
									<Button size="sm" variant="outline" class="flex-1">
										<Icon icon="hugeicons:download-01" class="mr-2 size-4" />
										PDF
									</Button>
								{/if}
							</div>
						</div>
					{/each}
				</div>
			{/if}
		</div>
	</div>
</div>

<Modal
	bind:open={showNewInvoiceModal}
	title="Create New Invoice"
	description="Create a new invoice for your client."
	size="lg"
>
	<form id="new-invoice-form" onsubmit={handleCreateInvoice} class="space-y-4">
		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="invoice-client">Client *</Label>
				<select
					id="invoice-client"
					bind:value={newInvoice.client}
					required
					class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-1 text-sm"
				>
					<option value="">Select client...</option>
					{#each clients as client}
						<option value={client.email}>{client.company}</option>
					{/each}
				</select>
			</div>
			<div class="space-y-2">
				<Label for="invoice-project">Project</Label>
				<select
					id="invoice-project"
					bind:value={newInvoice.project}
					class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-1 text-sm"
				>
					<option value="">Select project...</option>
					{#each projects as project}
						<option value={project.name}>{project.name}</option>
					{/each}
				</select>
			</div>
		</div>

		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="invoice-issue">Issue Date *</Label>
				<Input id="invoice-issue" type="date" bind:value={newInvoice.issueDate} required />
			</div>
			<div class="space-y-2">
				<Label for="invoice-due">Due Date *</Label>
				<Input id="invoice-due" type="date" bind:value={newInvoice.dueDate} required />
			</div>
		</div>

		<div class="space-y-2">
			<Label>Line Items</Label>
			<div class="space-y-2">
				{#each newInvoice.items as item, i}
					<div class="flex items-start gap-2">
						<Input placeholder="Description" bind:value={item.description} class="flex-1" />
						<Input type="number" placeholder="Qty" bind:value={item.quantity} class="w-20" />
						<Input type="number" placeholder="Rate" bind:value={item.rate} class="w-28" />
						<Button type="button" variant="ghost" size="icon" onclick={() => removeItem(i)} disabled={newInvoice.items.length === 1}>
							<Icon icon="hugeicons:delete-01" class="size-4" />
						</Button>
					</div>
				{/each}
			</div>
			<Button type="button" variant="outline" size="sm" class="gap-2" onclick={addItem}>
				<Icon icon="hugeicons:plus-sign" class="size-4" />
				Add Item
			</Button>
		</div>

		<div class="space-y-2">
			<Label for="invoice-notes">Notes</Label>
			<textarea
				id="invoice-notes"
				placeholder="Additional notes for the client..."
				bind:value={newInvoice.notes}
				rows={2}
				class="flex w-full resize-none rounded-md border border-input bg-background px-3 py-2 text-sm"
			></textarea>
		</div>

		<div class="space-y-2">
			<Label for="invoice-terms">Terms</Label>
			<Input id="invoice-terms" bind:value={newInvoice.terms} />
		</div>
	</form>

	{#snippet footer()}
		<Button variant="outline" onclick={() => showNewInvoiceModal = false}>Cancel</Button>
		<Button type="submit" form="new-invoice-form">
			<Icon icon="hugeicons:plus-sign" class="mr-2 size-4" />
			Create Invoice
		</Button>
	{/snippet}
</Modal>

<Modal
	bind:open={showDeleteModal}
	title="Delete Invoice"
	description="Are you sure you want to delete invoice {selectedInvoice?.number}? This action cannot be undone."
	size="sm"
>
	<div class="flex justify-end gap-2">
		<Button variant="outline" onclick={() => showDeleteModal = false}>Cancel</Button>
		<Button variant="destructive">
			<Icon icon="hugeicons:delete-01" class="mr-2 size-4" />
			Delete
		</Button>
	</div>
</Modal>
