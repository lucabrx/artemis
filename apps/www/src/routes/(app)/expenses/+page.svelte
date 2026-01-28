<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';
	import { Label } from '$lib/components/ui/label/index.js';

	let viewMode = $state<'list' | 'grid'>('list');
	let categoryFilter = $state('all');
	let statusFilter = $state('all');
	let searchQuery = $state('');
	let showNewExpenseModal = $state(false);
	let showExpenseDetails = $state(false);
	let selectedExpense = $state<Expense | null>(null);

	interface Expense {
		id: string;
		description: string;
		merchant: string;
		category: string;
		amount: number;
		date: string;
		project?: {
			name: string;
			color: string;
		};
		client?: {
			name: string;
			company: string;
			avatar: string;
		};
		receipt?: {
			url: string;
			name: string;
		};
		billable: boolean;
		billed: boolean;
		notes?: string;
		paymentMethod: string;
		status: 'pending' | 'approved' | 'reimbursed';
		submittedBy: {
			name: string;
			avatar: string;
		};
	}

	const projects = [
		{ name: 'Website Redesign', color: 'bg-blue-500' },
		{ name: 'Mobile App', color: 'bg-violet-500' },
		{ name: 'Brand Identity', color: 'bg-emerald-500' },
		{ name: 'Marketing Campaign', color: 'bg-amber-500' },
		{ name: 'E-commerce Platform', color: 'bg-rose-500' }
	];

	const clients = [
		{ name: 'Sarah Chen', company: 'Acme Corp', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Acme' },
		{ name: 'Michael Ross', company: 'TechStart', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=TechStart' },
		{ name: 'Emily Watson', company: 'Green Leaf', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=GreenLeaf' }
	];

	const teamMembers = [
		{ name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah' },
		{ name: 'Mike Ross', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Mike' },
		{ name: 'David Kim', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=David' }
	];

	const categories = [
		{ id: 'software', name: 'Software', icon: 'hugeicons:computer', color: 'bg-blue-500', budget: 500 },
		{ id: 'travel', name: 'Travel', icon: 'hugeicons:airplane', color: 'bg-violet-500', budget: 1000 },
		{ id: 'meals', name: 'Meals', icon: 'hugeicons:utensils', color: 'bg-amber-500', budget: 400 },
		{ id: 'office', name: 'Office', icon: 'hugeicons:folder-01', color: 'bg-emerald-500', budget: 200 },
		{ id: 'equipment', name: 'Equipment', icon: 'hugeicons:camera-01', color: 'bg-rose-500', budget: 800 },
		{ id: 'marketing', name: 'Marketing', icon: 'hugeicons:megaphone', color: 'bg-cyan-500', budget: 600 },
		{ id: 'training', name: 'Training', icon: 'hugeicons:book-01', color: 'bg-indigo-500', budget: 300 },
		{ id: 'other', name: 'Other', icon: 'hugeicons:more-horizontal', color: 'bg-slate-500', budget: 200 }
	];

	const expenses: Expense[] = [
		{
			id: 'exp1',
			description: 'Figma Pro subscription - Annual',
			merchant: 'Figma',
			category: 'software',
			amount: 144,
			date: 'Jan 28, 2026',
			paymentMethod: 'Credit Card',
			billable: false,
			billed: false,
			status: 'approved',
			submittedBy: teamMembers[0],
			receipt: { url: 'https://placehold.co/400x600/f3f4f6/666?text=Receipt', name: 'figma_receipt.pdf' },
			notes: 'Annual subscription renewal for design team'
		},
		{
			id: 'exp2',
			description: 'Client meeting lunch',
			merchant: 'The Bistro',
			category: 'meals',
			amount: 87.50,
			date: 'Jan 27, 2026',
			project: projects[0],
			client: clients[0],
			paymentMethod: 'Credit Card',
			billable: true,
			billed: false,
			status: 'pending',
			submittedBy: teamMembers[1],
			receipt: { url: 'https://placehold.co/400x600/f3f4f6/666?text=Receipt', name: 'bistro_receipt.pdf' },
			notes: 'Meeting with Sarah to discuss design revisions'
		},
		{
			id: 'exp3',
			description: 'Flight to SF - Client Workshop',
			merchant: 'Delta Airlines',
			category: 'travel',
			amount: 450,
			date: 'Jan 25, 2026',
			project: projects[3],
			client: clients[2],
			paymentMethod: 'Credit Card',
			billable: true,
			billed: true,
			status: 'reimbursed',
			submittedBy: teamMembers[0],
			receipt: { url: 'https://placehold.co/400x600/f3f4f6/666?text=Receipt', name: 'delta_ticket.pdf' },
			notes: 'Round trip for 2-day workshop'
		},
		{
			id: 'exp4',
			description: 'MacBook Pro M3 - Development',
			merchant: 'Apple Store',
			category: 'equipment',
			amount: 2499,
			date: 'Jan 22, 2026',
			paymentMethod: 'Credit Card',
			billable: false,
			billed: false,
			status: 'approved',
			submittedBy: teamMembers[2],
			receipt: { url: 'https://placehold.co/400x600/f3f4f6/666?text=Receipt', name: 'apple_invoice.pdf' },
			notes: 'New development machine for David'
		},
		{
			id: 'exp5',
			description: 'Google Ads - Campaign',
			merchant: 'Google',
			category: 'marketing',
			amount: 500,
			date: 'Jan 20, 2026',
			project: projects[3],
			paymentMethod: 'Credit Card',
			billable: true,
			billed: true,
			status: 'reimbursed',
			submittedBy: teamMembers[0],
			receipt: { url: 'https://placehold.co/400x600/f3f4f6/666?text=Receipt', name: 'google_ads.pdf' }
		},
		{
			id: 'exp6',
			description: 'Stationery supplies',
			merchant: 'Office Depot',
			category: 'office',
			amount: 45.20,
			date: 'Jan 18, 2026',
			paymentMethod: 'Debit Card',
			billable: false,
			billed: false,
			status: 'approved',
			submittedBy: teamMembers[1],
			receipt: { url: 'https://placehold.co/400x600/f3f4f6/666?text=Receipt', name: 'office_depot.pdf' }
		},
		{
			id: 'exp7',
			description: 'Uber - Client site visit',
			merchant: 'Uber',
			category: 'travel',
			amount: 34.50,
			date: 'Jan 15, 2026',
			project: projects[1],
			client: clients[1],
			paymentMethod: 'Credit Card',
			billable: true,
			billed: false,
			status: 'pending',
			submittedBy: teamMembers[2],
			receipt: { url: 'https://placehold.co/400x600/f3f4f6/666?text=Receipt', name: 'uber_receipt.pdf' }
		},
		{
			id: 'exp8',
			description: 'AWS Hosting - January',
			merchant: 'Amazon Web Services',
			category: 'software',
			amount: 127.80,
			date: 'Jan 10, 2026',
			paymentMethod: 'Credit Card',
			billable: false,
			billed: false,
			status: 'approved',
			submittedBy: teamMembers[0],
			receipt: { url: 'https://placehold.co/400x600/f3f4f6/666?text=Receipt', name: 'aws_invoice.pdf' }
		},
		{
			id: 'exp9',
			description: 'Design conference ticket',
			merchant: 'DesignConf',
			category: 'training',
			amount: 299,
			date: 'Jan 5, 2026',
			paymentMethod: 'Credit Card',
			billable: false,
			billed: false,
			status: 'approved',
			submittedBy: teamMembers[0],
			receipt: { url: 'https://placehold.co/400x600/f3f4f6/666?text=Receipt', name: 'conference_ticket.pdf' },
			notes: 'Annual design conference - March 2026'
		}
	];

	const statusOptions = [
		{ value: 'all', label: 'All', count: expenses.length },
		{ value: 'pending', label: 'Pending', count: expenses.filter(e => e.status === 'pending').length },
		{ value: 'approved', label: 'Approved', count: expenses.filter(e => e.status === 'approved').length },
		{ value: 'reimbursed', label: 'Reimbursed', count: expenses.filter(e => e.status === 'reimbursed').length }
	];

	let newExpense = $state({
		description: '',
		merchant: '',
		category: '',
		amount: '',
		date: new Date().toISOString().split('T')[0],
		project: '',
		billable: false,
		paymentMethod: 'Credit Card',
		notes: ''
	});

	function getCategoryIcon(categoryId: string): string {
		return categories.find(c => c.id === categoryId)?.icon || 'hugeicons:more-horizontal';
	}

	function getCategoryColor(categoryId: string): string {
		return categories.find(c => c.id === categoryId)?.color || 'bg-slate-500';
	}

	function getCategoryName(categoryId: string): string {
		return categories.find(c => c.id === categoryId)?.name || categoryId;
	}

	function getStatusColor(status: string): string {
		switch (status) {
			case 'pending':
				return 'bg-amber-500/10 text-amber-600 border-amber-500/20';
			case 'approved':
				return 'bg-blue-500/10 text-blue-600 border-blue-500/20';
			case 'reimbursed':
				return 'bg-emerald-500/10 text-emerald-600 border-emerald-500/20';
			default:
				return 'bg-muted';
		}
	}

	function formatCurrency(amount: number): string {
		return new Intl.NumberFormat('en-US', {
			style: 'currency',
			currency: 'USD'
		}).format(amount);
	}

	const filteredExpenses = $derived(() => {
		let filtered = expenses;
		if (categoryFilter !== 'all') {
			filtered = filtered.filter(e => e.category === categoryFilter);
		}
		if (statusFilter !== 'all') {
			filtered = filtered.filter(e => e.status === statusFilter);
		}
		if (searchQuery) {
			const q = searchQuery.toLowerCase();
			filtered = filtered.filter(
				e =>
					e.description.toLowerCase().includes(q) ||
					e.merchant.toLowerCase().includes(q) ||
					getCategoryName(e.category).toLowerCase().includes(q)
			);
		}
		return filtered;
	});

	const stats = {
		totalThisMonth: expenses.reduce((acc, e) => acc + e.amount, 0),
		billable: expenses.filter(e => e.billable).reduce((acc, e) => acc + e.amount, 0),
		reimbursable: expenses.filter(e => !e.billable).reduce((acc, e) => acc + e.amount, 0),
		pending: expenses.filter(e => e.status === 'pending').reduce((acc, e) => acc + e.amount, 0)
	};

	const categorySpending = categories.map(cat => {
		const spent = expenses.filter(e => e.category === cat.id).reduce((acc, e) => acc + e.amount, 0);
		return { ...cat, spent, remaining: cat.budget - spent };
	});

	function handleCreateExpense(e: Event) {
		e.preventDefault();
		console.log('Creating expense:', newExpense);
		showNewExpenseModal = false;
	}

	function openExpenseDetails(expense: Expense) {
		selectedExpense = expense;
		showExpenseDetails = true;
	}
</script>

<svelte:head>
	<title>Expenses - Artemis</title>
</svelte:head>

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="px-6 py-4">
		<div class="mb-6 flex items-center justify-end gap-2">
			<Button variant="outline" class="gap-2">
				<Icon icon="hugeicons:download-01" class="size-4" />
				Export
			</Button>
			<Button class="gap-2" onclick={() => showNewExpenseModal = true}>
				<Icon icon="hugeicons:plus-sign" class="size-4" />
				Add Expense
			</Button>
		</div>

		<div class="grid gap-4 md:grid-cols-4">
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">This Month</p>
						<p class="text-2xl font-bold">{formatCurrency(stats.totalThisMonth)}</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
						<Icon icon="hugeicons:dollar-circle" class="size-5 text-primary" />
					</div>
				</div>
			</div>
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">Billable</p>
						<p class="text-2xl font-bold text-blue-600">{formatCurrency(stats.billable)}</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-blue-500/10">
						<Icon icon="hugeicons:invoice-01" class="size-5 text-blue-600" />
					</div>
				</div>
			</div>
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">Reimbursable</p>
						<p class="text-2xl font-bold text-emerald-600">{formatCurrency(stats.reimbursable)}</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-emerald-500/10">
						<Icon icon="hugeicons:refund" class="size-5 text-emerald-600" />
					</div>
				</div>
			</div>
			<div class="rounded-xl border border-border bg-card p-4">
				<div class="flex items-center justify-between">
					<div>
						<p class="text-sm text-muted-foreground">Pending</p>
						<p class="text-2xl font-bold text-amber-600">{formatCurrency(stats.pending)}</p>
					</div>
					<div class="flex size-10 items-center justify-center rounded-lg bg-amber-500/10">
						<Icon icon="hugeicons:clock-01" class="size-5 text-amber-600" />
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
							placeholder="Search expenses by description, merchant, or category..."
							class="pl-11"
							bind:value={searchQuery}
						/>
					</div>

					<div class="flex flex-wrap items-center gap-2">
						<div class="flex rounded-lg border border-border p-1">
							{#each statusOptions as option}
								<button
									class="rounded px-3 py-1 text-sm font-medium transition-colors {statusFilter === option.value ? 'bg-primary text-primary-foreground' : 'text-muted-foreground hover:text-foreground'}"
									onclick={() => statusFilter = option.value}
								>
									{option.label}
									<span class="ml-1 text-xs opacity-70">({option.count})</span>
								</button>
							{/each}
						</div>

						<Separator orientation="vertical" class="hidden h-9 lg:block" />

						<div class="flex rounded-lg border border-border p-1">
							<button
								class="rounded p-1.5 transition-colors {viewMode === 'list' ? 'bg-muted text-foreground' : 'text-muted-foreground hover:text-foreground'}"
								onclick={() => viewMode = 'list'}
							>
								<Icon icon="hugeicons:menu-square" class="size-5" />
							</button>
							<button
								class="rounded p-1.5 transition-colors {viewMode === 'grid' ? 'bg-muted text-foreground' : 'text-muted-foreground hover:text-foreground'}"
								onclick={() => viewMode = 'grid'}
							>
								<Icon icon="hugeicons:grid" class="size-5" />
							</button>
						</div>
					</div>
				</div>

				<div class="flex flex-wrap gap-2">
					<button
						class="flex items-center gap-2 rounded-full border px-3 py-1.5 text-sm font-medium transition-colors {categoryFilter === 'all' ? 'bg-primary text-primary-foreground border-primary' : 'border-border text-muted-foreground hover:text-foreground'}"
						onclick={() => categoryFilter = 'all'}
					>
						<Icon icon="hugeicons:layers" class="size-4" />
						All Categories
					</button>
					{#each categories as category}
						{@const count = expenses.filter(e => e.category === category.id).length}
						<button
							class="flex items-center gap-2 rounded-full border px-3 py-1.5 text-sm font-medium transition-colors {categoryFilter === category.id ? 'bg-primary text-primary-foreground border-primary' : 'border-border text-muted-foreground hover:text-foreground'}"
							onclick={() => categoryFilter = category.id}
						>
							<div class="size-2 rounded-full {category.color}"></div>
							{category.name}
							<span class="text-xs opacity-70">({count})</span>
						</button>
					{/each}
				</div>
			</div>

			{#if filteredExpenses().length === 0}
				<div class="p-12 text-center">
					<div class="mb-4 inline-flex size-16 items-center justify-center rounded-full bg-muted">
						<Icon icon="hugeicons:receipt" class="size-8 text-muted-foreground" />
					</div>
					<h3 class="text-lg font-semibold">No expenses found</h3>
					<p class="mt-1 text-muted-foreground">Try adjusting your search or filters.</p>
				</div>

			{:else if viewMode === 'list'}
				<div class="divide-y divide-border">
					{#each filteredExpenses() as expense}
						<div class="group flex flex-col gap-4 p-4 transition-colors hover:bg-muted/30 sm:flex-row sm:items-center">
							<div class="flex min-w-0 flex-1 items-center gap-3">
								<div class="flex size-10 shrink-0 items-center justify-center rounded-lg {getCategoryColor(expense.category)}">
									<Icon icon={getCategoryIcon(expense.category)} class="size-5 text-white" />
								</div>
								<div class="min-w-0">
									<p class="truncate font-semibold">{expense.description}</p>
									<div class="flex flex-wrap items-center gap-2 text-sm text-muted-foreground">
										<span>{expense.merchant}</span>
										<span>·</span>
										<span>{expense.date}</span>
										{#if expense.project}
											<span>·</span>
											<div class="flex items-center gap-1">
												<div class="size-1.5 rounded-full {expense.project.color}"></div>
												<span>{expense.project.name}</span>
											</div>
										{/if}
									</div>
								</div>
							</div>

							<div class="flex flex-wrap items-center gap-4 sm:flex-nowrap">
								<div class="flex items-center gap-2">
									{#if expense.billable}
										<span class="flex items-center gap-1 rounded-full bg-blue-500/10 px-2 py-0.5 text-xs font-medium text-blue-600">
											<Icon icon="hugeicons:invoice-01" class="size-3" />
											Billable
										</span>
									{/if}
									{#if expense.receipt}
										<span class="flex items-center gap-1 rounded-full bg-muted px-2 py-0.5 text-xs">
											<Icon icon="hugeicons:receipt" class="size-3" />
											Receipt
										</span>
									{/if}
									<span class="flex items-center gap-1 rounded-full border px-2 py-0.5 text-xs font-medium capitalize {getStatusColor(expense.status)}">
										{expense.status}
									</span>
								</div>

								<div class="flex items-center gap-3">
									<p class="font-semibold">{formatCurrency(expense.amount)}</p>
									<img src={expense.submittedBy.avatar} alt={expense.submittedBy.name} class="size-8 rounded-full" />
									<Button
										variant="ghost"
										size="icon"
										class="opacity-0 transition-opacity group-hover:opacity-100"
										onclick={() => openExpenseDetails(expense)}
									>
										<Icon icon="hugeicons:more-vertical" class="size-4" />
									</Button>
								</div>
							</div>
						</div>
					{/each}
				</div>

			{:else if viewMode === 'grid'}
				<div class="grid gap-4 p-4 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
					{#each filteredExpenses() as expense}
						<div class="group rounded-xl border border-border bg-card p-4 transition-all hover:border-primary/20 hover:shadow-lg">
							<div class="mb-3 flex items-start justify-between">
								<div class="flex size-10 items-center justify-center rounded-lg {getCategoryColor(expense.category)}">
									<Icon icon={getCategoryIcon(expense.category)} class="size-5 text-white" />
								</div>
								<span class="flex items-center gap-1 rounded-full border px-2 py-0.5 text-xs font-medium capitalize {getStatusColor(expense.status)}">
									{expense.status}
								</span>
							</div>

							<h4 class="mb-1 font-semibold">{expense.description}</h4>
							<p class="mb-3 text-sm text-muted-foreground">{expense.merchant}</p>

							<div class="mb-3 flex items-center gap-2 text-sm text-muted-foreground">
								<Icon icon="hugeicons:calendar-01" class="size-4" />
								<span>{expense.date}</span>
							</div>

							{#if expense.project}
								<div class="mb-3 flex items-center gap-2">
									<div class="size-2 rounded-full {expense.project.color}"></div>
									<span class="text-sm text-muted-foreground">{expense.project.name}</span>
								</div>
							{/if}

							<div class="mb-4 flex flex-wrap gap-2">
								{#if expense.billable}
									<span class="flex items-center gap-1 rounded-full bg-blue-500/10 px-2 py-0.5 text-xs font-medium text-blue-600">
										<Icon icon="hugeicons:invoice-01" class="size-3" />
										Billable
									</span>
								{/if}
								{#if expense.receipt}
									<span class="flex items-center gap-1 rounded-full bg-muted px-2 py-0.5 text-xs">
										<Icon icon="hugeicons:receipt" class="size-3" />
										Receipt
									</span>
								{/if}
							</div>

							<div class="flex items-center justify-between border-t border-border pt-3">
								<div class="flex items-center gap-2">
									<img src={expense.submittedBy.avatar} alt={expense.submittedBy.name} class="size-6 rounded-full" />
									<span class="text-sm text-muted-foreground">{expense.submittedBy.name}</span>
								</div>
								<p class="font-semibold">{formatCurrency(expense.amount)}</p>
							</div>
						</div>
					{/each}
				</div>
			{/if}
		</div>

		<div class="mt-8 grid gap-6 lg:grid-cols-2">
			<div class="rounded-2xl border border-border bg-card p-6">
				<h3 class="mb-4 font-semibold">Budget Overview</h3>
				<div class="space-y-4">
					{#each categorySpending as cat}
						<div>
							<div class="mb-1 flex items-center justify-between">
								<div class="flex items-center gap-2">
									<div class="flex size-6 items-center justify-center rounded {cat.color}">
										<Icon icon={cat.icon} class="size-3 text-white" />
									</div>
									<span class="text-sm font-medium">{cat.name}</span>
								</div>
								<span class="text-xs text-muted-foreground">{formatCurrency(cat.spent)} / {formatCurrency(cat.budget)}</span>
							</div>
							<div class="h-2 overflow-hidden rounded-full bg-muted">
								<div class="h-full rounded-full {cat.color}" style="width: {Math.min((cat.spent / cat.budget) * 100, 100)}%"></div>
							</div>
							<p class="mt-1 text-xs text-muted-foreground">{Math.round((cat.spent / cat.budget) * 100)}% used · {formatCurrency(cat.remaining)} remaining</p>
						</div>
					{/each}
				</div>
			</div>

			<div class="rounded-2xl border border-border bg-card p-6">
				<h3 class="mb-4 font-semibold">Spending by Category</h3>
				<div class="grid grid-cols-2 gap-4 sm:grid-cols-4">
					{#each categorySpending.slice(0, 4) as cat}
						<div class="rounded-xl bg-muted/50 p-4 text-center">
							<div class="mx-auto mb-2 flex size-10 items-center justify-center rounded-lg {cat.color}">
								<Icon icon={cat.icon} class="size-5 text-white" />
							</div>
							<p class="text-lg font-bold">{formatCurrency(cat.spent)}</p>
							<p class="text-xs text-muted-foreground">{cat.name}</p>
						</div>
					{/each}
				</div>
			</div>
		</div>
	</div>
</div>

<Modal
	bind:open={showNewExpenseModal}
	title="Add New Expense"
	description="Record a new business expense."
	size="lg"
>
	<form id="new-expense-form" onsubmit={handleCreateExpense} class="space-y-4">
		<div class="space-y-2">
			<Label for="expense-description">Description *</Label>
			<Input id="expense-description" placeholder="What was this expense for?" bind:value={newExpense.description} required />
		</div>

		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="expense-merchant">Merchant *</Label>
				<Input id="expense-merchant" placeholder="Store or vendor name" bind:value={newExpense.merchant} required />
			</div>
			<div class="space-y-2">
				<Label for="expense-category">Category *</Label>
				<select
					id="expense-category"
					bind:value={newExpense.category}
					required
					class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-1 text-sm"
				>
					<option value="">Select category...</option>
					{#each categories as cat}
						<option value={cat.id}>{cat.name}</option>
					{/each}
				</select>
			</div>
		</div>

		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="expense-amount">Amount *</Label>
				<Input id="expense-amount" type="number" step="0.01" placeholder="0.00" bind:value={newExpense.amount} required />
			</div>
			<div class="space-y-2">
				<Label for="expense-date">Date *</Label>
				<Input id="expense-date" type="date" bind:value={newExpense.date} required />
			</div>
		</div>

		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="expense-project">Project</Label>
				<select
					id="expense-project"
					bind:value={newExpense.project}
					class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-1 text-sm"
				>
					<option value="">Select project...</option>
					{#each projects as project}
						<option value={project.name}>{project.name}</option>
					{/each}
				</select>
			</div>
			<div class="space-y-2">
				<Label for="expense-payment">Payment Method</Label>
				<select
					id="expense-payment"
					bind:value={newExpense.paymentMethod}
					class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-1 text-sm"
				>
					<option value="Credit Card">Credit Card</option>
					<option value="Debit Card">Debit Card</option>
					<option value="Cash">Cash</option>
					<option value="Bank Transfer">Bank Transfer</option>
				</select>
			</div>
		</div>

		<div class="space-y-2">
			<Label for="expense-receipt">Receipt</Label>
			<div class="flex items-center gap-2 rounded-lg border border-dashed border-border p-4">
				<Icon icon="hugeicons:upload-01" class="size-5 text-muted-foreground" />
				<span class="text-sm text-muted-foreground">Drop receipt here or click to upload</span>
			</div>
		</div>

		<div class="space-y-2">
			<Label for="expense-notes">Notes</Label>
			<textarea
				id="expense-notes"
				placeholder="Additional details..."
				bind:value={newExpense.notes}
				rows={2}
				class="flex w-full resize-none rounded-md border border-input bg-background px-3 py-2 text-sm"
			></textarea>
		</div>

		<label class="flex items-center gap-2">
			<input type="checkbox" bind:checked={newExpense.billable} class="size-4 rounded border-input" />
			<span class="text-sm">Billable to client</span>
		</label>
	</form>

	{#snippet footer()}
		<Button variant="outline" onclick={() => showNewExpenseModal = false}>Cancel</Button>
		<Button type="submit" form="new-expense-form">
			<Icon icon="hugeicons:plus-sign" class="mr-2 size-4" />
			Add Expense
		</Button>
	{/snippet}
</Modal>

<Modal
	bind:open={showExpenseDetails}
	title="Expense Details"
	description="View and manage expense."
	size="md"
>
	{#if selectedExpense}
		<div class="space-y-4">
			<div class="flex items-start gap-4">
				<div class="flex size-12 items-center justify-center rounded-xl {getCategoryColor(selectedExpense.category)}">
					<Icon icon={getCategoryIcon(selectedExpense.category)} class="size-6 text-white" />
				</div>
				<div>
					<h3 class="font-semibold">{selectedExpense.description}</h3>
					<p class="text-sm text-muted-foreground">{selectedExpense.merchant}</p>
				</div>
			</div>

			<div class="grid grid-cols-2 gap-4 rounded-xl bg-muted p-4">
				<div>
					<p class="text-xs text-muted-foreground">Amount</p>
					<p class="text-lg font-bold">{formatCurrency(selectedExpense.amount)}</p>
				</div>
				<div>
					<p class="text-xs text-muted-foreground">Date</p>
					<p class="text-sm font-medium">{selectedExpense.date}</p>
				</div>
				<div>
					<p class="text-xs text-muted-foreground">Category</p>
					<p class="text-sm font-medium">{getCategoryName(selectedExpense.category)}</p>
				</div>
				<div>
					<p class="text-xs text-muted-foreground">Status</p>
					<span class="inline-flex items-center gap-1 rounded-full border px-2 py-0.5 text-xs font-medium capitalize {getStatusColor(selectedExpense.status)}">
						{selectedExpense.status}
					</span>
				</div>
			</div>

			{#if selectedExpense.receipt}
				<div>
					<p class="mb-2 text-sm font-medium">Receipt</p>
					<div class="aspect-video overflow-hidden rounded-xl border border-border">
						<img src={selectedExpense.receipt.url} alt="Receipt" class="size-full object-cover" />
					</div>
					<p class="mt-1 text-xs text-muted-foreground">{selectedExpense.receipt.name}</p>
				</div>
			{/if}
		</div>
	{/if}

	{#snippet footer()}
		<Button variant="outline" onclick={() => showExpenseDetails = false}>Close</Button>
		<Button variant="destructive">
			<Icon icon="hugeicons:delete-01" class="mr-2 size-4" />
			Delete
		</Button>
	{/snippet}
</Modal>
