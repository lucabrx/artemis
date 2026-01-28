<script lang="ts">
	import { page } from '$app/state';
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Input } from '$lib/components/ui/input/index.js';

	let showRecordPaymentModal = $state(false);
	let showSendReminderModal = $state(false);
	let showEditModal = $state(false);
	let showDeleteModal = $state(false);
	let paymentAmount = $state('');

	const invoiceId = page.params.id;

	const invoice = {
		id: invoiceId,
		number: 'INV-2026-001',
		status: 'sent' as const,
		client: {
			name: 'Sarah Chen',
			company: 'Acme Corporation',
			email: 'sarah@acmecorp.com',
			phone: '+1 (555) 123-4567',
			address: '123 Market Street, Suite 400',
			city: 'San Francisco, CA 94105',
			avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Acme'
		},
		project: {
			name: 'Website Redesign',
			color: 'bg-blue-500'
		},
		issueDate: 'January 15, 2026',
		dueDate: 'February 15, 2026',
		subtotal: 12500,
		taxRate: 10,
		tax: 1250,
		total: 13750,
		paid: 6875,
		balance: 6875,
		items: [
			{
				description: 'Website Design - Phase 1',
				details: 'Initial design concepts, wireframes, and high-fidelity mockups for homepage and key interior pages. Includes 3 rounds of revisions.',
				quantity: 40,
				rate: 150,
				amount: 6000
			},
			{
				description: 'UI/UX Consulting',
				details: 'User research, usability testing, and consultation on best practices for conversion optimization.',
				quantity: 10,
				rate: 200,
				amount: 2000
			},
			{
				description: 'Development Hours',
				details: 'Frontend development using React and Tailwind CSS. Responsive implementation and component library creation.',
				quantity: 30,
				rate: 150,
				amount: 4500
			}
		],
		notes: 'Thank you for your business! We appreciate the opportunity to work with Acme Corporation on this exciting project. Please don\'t hesitate to reach out if you have any questions about this invoice.',
		terms: 'Net 30 days. A 1.5% monthly late fee will be applied to overdue balances.',
		paymentHistory: [
			{
				date: 'January 20, 2026',
				method: 'Bank Transfer',
				amount: 6875,
				note: '50% deposit as agreed'
			}
		],
		activity: [
			{ id: 'a1', type: 'created', user: 'You', timestamp: 'Jan 15, 2026 at 9:00 AM', text: 'created the invoice' },
			{ id: 'a2', type: 'sent', user: 'You', timestamp: 'Jan 15, 2026 at 9:05 AM', text: 'sent to sarah@acmecorp.com' },
			{ id: 'a3', type: 'viewed', user: 'System', timestamp: 'Jan 15, 2026 at 2:30 PM', text: 'invoice was viewed by client' },
			{ id: 'a4', type: 'payment', user: 'System', timestamp: 'Jan 20, 2026 at 10:15 AM', text: 'payment of $6,875 received' }
		],
		reminders: [
			{ date: 'Jan 25, 2026', type: 'automatic', sent: true },
			{ date: 'Feb 1, 2026', type: 'scheduled', sent: false }
		]
	};

	const company = {
		name: 'Your Design Studio',
		address: '456 Creative Avenue',
		city: 'New York, NY 10001',
		email: 'billing@yourstudio.com',
		phone: '+1 (555) 987-6543',
		logo: 'https://placehold.co/200x80/f3f4f6/666?text=YOUR+STUDIO'
	};

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
			default:
				return 'bg-muted';
		}
	}

	function getActivityIcon(type: string): string {
		switch (type) {
			case 'created':
				return 'hugeicons:plus-sign';
			case 'sent':
				return 'hugeicons:send-01';
			case 'viewed':
				return 'hugeicons:eye';
			case 'payment':
				return 'hugeicons:dollar-circle';
			default:
				return 'hugeicons:activity-01';
		}
	}

	function formatCurrency(amount: number): string {
		return new Intl.NumberFormat('en-US', {
			style: 'currency',
			currency: 'USD',
			minimumFractionDigits: 2
		}).format(amount);
	}

	function handleRecordPayment() {
		console.log('Recording payment:', paymentAmount);
		showRecordPaymentModal = false;
		paymentAmount = '';
	}

	const progressPercent = (invoice.paid / invoice.total) * 100;
</script>

<svelte:head>
	<title>Invoice {invoice.number} - Artemis</title>
</svelte:head>

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="px-6 py-8">
		<div class="mb-6 flex flex-col gap-4 lg:flex-row lg:items-start lg:justify-between">
			<div class="flex items-start gap-4">
				<Button variant="outline" size="icon" href="/app/invoices">
					<Icon icon="hugeicons:arrow-left-01" class="size-5" />
				</Button>
				<div>
					<div class="mb-2 flex items-center gap-3">
						<h1 class="text-2xl font-bold">{invoice.number}</h1>
						<span class="flex items-center gap-1 rounded-full border px-3 py-1 text-sm font-medium capitalize {getStatusColor(invoice.status)}">
							{invoice.status}
						</span>
					</div>
					<p class="text-muted-foreground">
						Issued {invoice.issueDate} · Due {invoice.dueDate}
					</p>
				</div>
			</div>
			<div class="flex flex-wrap gap-2">
				<Button variant="outline" class="gap-2">
					<Icon icon="hugeicons:download-01" class="size-4" />
					Download PDF
				</Button>
				<Button variant="outline" class="gap-2">
					<Icon icon="hugeicons:printer" class="size-4" />
					Print
				</Button>
				<Button variant="outline" class="gap-2" onclick={() => showSendReminderModal = true}>
					<Icon icon="hugeicons:mail-01" class="size-4" />
					Send Reminder
				</Button>
				<Button variant="outline" class="gap-2" onclick={() => showEditModal = true}>
					<Icon icon="hugeicons:pencil-edit-01" class="size-4" />
					Edit
				</Button>
				{#if invoice.balance > 0}
					<Button class="gap-2" onclick={() => showRecordPaymentModal = true}>
						<Icon icon="hugeicons:dollar-circle" class="size-4" />
						Record Payment
					</Button>
				{/if}
			</div>
		</div>

		<div class="grid gap-6 lg:grid-cols-3">
			<div class="space-y-6 lg:col-span-2">
				<div class="overflow-hidden rounded-2xl border border-border bg-card">
					<div class="border-b border-border p-6">
						<div class="mb-8 flex flex-col gap-6 sm:flex-row sm:items-start sm:justify-between">
							<div>
								<img src={company.logo} alt={company.name} class="mb-4 h-12 object-contain" />
								<h2 class="font-semibold">{company.name}</h2>
								<p class="text-sm text-muted-foreground">{company.address}</p>
								<p class="text-sm text-muted-foreground">{company.city}</p>
								<p class="text-sm text-muted-foreground">{company.email}</p>
							</div>
							<div class="text-left sm:text-right">
								<h3 class="text-sm font-medium text-muted-foreground">Bill To</h3>
								<p class="font-semibold">{invoice.client.company}</p>
								<p class="text-sm">{invoice.client.name}</p>
								<p class="text-sm text-muted-foreground">{invoice.client.address}</p>
								<p class="text-sm text-muted-foreground">{invoice.client.city}</p>
								<p class="text-sm text-muted-foreground">{invoice.client.email}</p>
							</div>
						</div>

						<div class="mb-6 grid grid-cols-2 gap-4 md:grid-cols-4">
							<div>
								<p class="text-xs text-muted-foreground">Invoice Number</p>
								<p class="font-medium">{invoice.number}</p>
							</div>
							<div>
								<p class="text-xs text-muted-foreground">Issue Date</p>
								<p class="font-medium">{invoice.issueDate}</p>
							</div>
							<div>
								<p class="text-xs text-muted-foreground">Due Date</p>
								<p class="font-medium">{invoice.dueDate}</p>
							</div>
							<div>
								<p class="text-xs text-muted-foreground">Project</p>
								<div class="flex items-center gap-2">
									<div class="size-2 rounded-full {invoice.project.color}"></div>
									<p class="font-medium">{invoice.project.name}</p>
								</div>
							</div>
						</div>
					</div>

					<div class="p-6">
						<table class="w-full">
							<thead>
								<tr class="border-b text-left text-sm text-muted-foreground">
									<th class="pb-3 font-medium">Description</th>
									<th class="pb-3 font-medium">Qty</th>
									<th class="pb-3 font-medium">Rate</th>
									<th class="pb-3 text-right font-medium">Amount</th>
								</tr>
							</thead>
							<tbody class="text-sm">
								{#each invoice.items as item}
									<tr class="border-b">
										<td class="py-4">
											<p class="font-medium">{item.description}</p>
											<p class="text-xs text-muted-foreground">{item.details}</p>
										</td>
										<td class="py-4">{item.quantity}</td>
										<td class="py-4">{formatCurrency(item.rate)}</td>
										<td class="py-4 text-right font-medium">{formatCurrency(item.amount)}</td>
									</tr>
								{/each}
							</tbody>
						</table>

						<div class="mt-6 flex flex-col gap-4 sm:flex-row sm:justify-between">
							<div class="sm:max-w-xs">
								{#if invoice.notes}
									<div class="mb-4">
										<p class="text-xs font-medium text-muted-foreground">Notes</p>
										<p class="text-sm">{invoice.notes}</p>
									</div>
								{/if}
								{#if invoice.terms}
									<div>
										<p class="text-xs font-medium text-muted-foreground">Terms</p>
										<p class="text-sm text-muted-foreground">{invoice.terms}</p>
									</div>
								{/if}
							</div>

							<div class="min-w-[200px] space-y-2 text-sm">
								<div class="flex justify-between">
									<span class="text-muted-foreground">Subtotal</span>
									<span class="font-medium">{formatCurrency(invoice.subtotal)}</span>
								</div>
								<div class="flex justify-between">
									<span class="text-muted-foreground">Tax ({invoice.taxRate}%)</span>
									<span class="font-medium">{formatCurrency(invoice.tax)}</span>
								</div>
								<Separator />
								<div class="flex justify-between text-base">
									<span class="font-semibold">Total</span>
									<span class="font-bold">{formatCurrency(invoice.total)}</span>
								</div>
								<div class="flex justify-between text-emerald-600">
									<span>Paid</span>
									<span class="font-medium">-{formatCurrency(invoice.paid)}</span>
								</div>
								<div class="flex justify-between text-lg font-bold">
									<span>Balance Due</span>
									<span class={invoice.balance > 0 ? 'text-rose-600' : 'text-emerald-600'}>
										{formatCurrency(invoice.balance)}
									</span>
								</div>
							</div>
						</div>
					</div>
				</div>

				{#if invoice.paymentHistory.length > 0}
					<div class="rounded-2xl border border-border bg-card p-6">
						<h3 class="mb-4 font-semibold">Payment History</h3>
						<div class="space-y-3">
							{#each invoice.paymentHistory as payment}
								<div class="flex items-center justify-between rounded-lg bg-muted/50 p-3">
									<div class="flex items-center gap-3">
										<div class="flex size-10 items-center justify-center rounded-full bg-emerald-500/10">
											<Icon icon="hugeicons:tick-double-01" class="size-5 text-emerald-600" />
										</div>
										<div>
											<p class="font-medium">{formatCurrency(payment.amount)}</p>
											<p class="text-xs text-muted-foreground">{payment.method}</p>
										</div>
									</div>
									<div class="text-right">
										<p class="text-sm text-muted-foreground">{payment.date}</p>
										{#if payment.note}
											<p class="text-xs text-muted-foreground">{payment.note}</p>
										{/if}
									</div>
								</div>
							{/each}
						</div>
					</div>
				{/if}

				<div class="rounded-2xl border border-border bg-card p-6">
					<h3 class="mb-4 font-semibold">Activity</h3>
					<div class="space-y-4">
						{#each invoice.activity as activity}
							<div class="flex gap-3">
								<div class="flex size-8 shrink-0 items-center justify-center rounded-lg bg-muted">
									<Icon icon={getActivityIcon(activity.type)} class="size-4 text-muted-foreground" />
								</div>
								<div class="flex-1">
									<p class="text-sm">
										<span class="font-medium">{activity.user}</span>
										<span class="text-muted-foreground">{activity.text}</span>
									</p>
									<p class="text-xs text-muted-foreground">{activity.timestamp}</p>
								</div>
							</div>
						{/each}
					</div>
				</div>
			</div>

			<div class="space-y-6">
				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Payment Status</h3>
					<div class="mb-4 text-center">
						<div class="relative mx-auto mb-2 inline-flex size-24 items-center justify-center">
							<svg class="size-full -rotate-90" viewBox="0 0 36 36">
								<path
									class="fill-none stroke-muted"
									d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"
									stroke-width="3"
								/>
								<path
									class="fill-none stroke-emerald-500"
									stroke-dasharray="{progressPercent}, 100"
									d="M18 2.0845 a 15.9155 15.9155 0 0 1 0 31.831 a 15.9155 15.9155 0 0 1 0 -31.831"
									stroke-width="3"
								/>
							</svg>
							<span class="absolute text-lg font-bold">{Math.round(progressPercent)}%</span>
						</div>
						<p class="text-sm text-muted-foreground">{formatCurrency(invoice.paid)} of {formatCurrency(invoice.total)}</p>
					</div>

					<div class="space-y-3">
						<div class="flex justify-between text-sm">
							<span class="text-muted-foreground">Total Amount</span>
							<span class="font-medium">{formatCurrency(invoice.total)}</span>
						</div>
						<div class="flex justify-between text-sm">
							<span class="text-muted-foreground">Amount Paid</span>
							<span class="font-medium text-emerald-600">{formatCurrency(invoice.paid)}</span>
						</div>
						<Separator />
						<div class="flex justify-between">
							<span class="font-medium">Balance Due</span>
							<span class="font-bold {invoice.balance > 0 ? 'text-rose-600' : 'text-emerald-600'}">
								{formatCurrency(invoice.balance)}
							</span>
						</div>
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Client</h3>
					<div class="mb-4 flex items-center gap-3">
						<img src={invoice.client.avatar} alt={invoice.client.name} class="size-12 rounded-full" />
						<div>
							<p class="font-medium">{invoice.client.company}</p>
							<p class="text-sm text-muted-foreground">{invoice.client.name}</p>
						</div>
					</div>
					<div class="space-y-2 text-sm">
						<div class="flex items-center gap-2">
							<Icon icon="hugeicons:mail-01" class="size-4 text-muted-foreground" />
							<span>{invoice.client.email}</span>
						</div>
						<div class="flex items-center gap-2">
							<Icon icon="hugeicons:call" class="size-4 text-muted-foreground" />
							<span>{invoice.client.phone}</span>
						</div>
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Reminders</h3>
					<div class="space-y-3">
						{#each invoice.reminders as reminder}
							<div class="flex items-center justify-between text-sm">
								<div class="flex items-center gap-2">
									<Icon 
										icon={reminder.sent ? 'hugeicons:tick-double-01' : 'hugeicons:clock-01'} 
										class="size-4 {reminder.sent ? 'text-emerald-600' : 'text-muted-foreground'}" 
									/>
									<span>{reminder.date}</span>
								</div>
								<span class="text-xs text-muted-foreground capitalize">{reminder.type}</span>
							</div>
						{/each}
					</div>
					<Button variant="outline" class="mt-4 w-full gap-2" onclick={() => showSendReminderModal = true}>
						<Icon icon="hugeicons:mail-01" class="size-4" />
						Send Manual Reminder
					</Button>
				</div>

				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Quick Actions</h3>
					<div class="space-y-2">
						<Button variant="outline" class="w-full justify-start gap-2">
							<Icon icon="hugeicons:duplicate" class="size-4" />
							Duplicate Invoice
						</Button>
						<Button variant="outline" class="w-full justify-start gap-2">
							<Icon icon="hugeicons:link-01" class="size-4" />
							Copy Share Link
						</Button>
						<Button variant="outline" class="w-full justify-start gap-2 text-destructive hover:text-destructive" onclick={() => showDeleteModal = true}>
							<Icon icon="hugeicons:delete-01" class="size-4" />
							Delete Invoice
						</Button>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<Modal
	bind:open={showRecordPaymentModal}
	title="Record Payment"
	description="Record a payment for this invoice."
	size="sm"
>
	<div class="space-y-4">
		<div class="rounded-xl bg-muted p-4 text-center">
			<p class="text-sm text-muted-foreground">Balance Due</p>
			<p class="text-2xl font-bold">{formatCurrency(invoice.balance)}</p>
		</div>
		<div class="space-y-2">
			<Label for="payment-amount">Payment Amount</Label>
			<Input 
				id="payment-amount" 
				type="number" 
				placeholder="0.00" 
				bind:value={paymentAmount}
			/>
		</div>
		<div class="space-y-2">
			<Label for="payment-method">Payment Method</Label>
			<select id="payment-method" class="h-10 w-full rounded-md border border-input bg-background px-3 text-sm">
				<option value="">Select method...</option>
				<option value="bank">Bank Transfer</option>
				<option value="card">Credit Card</option>
				<option value="check">Check</option>
				<option value="cash">Cash</option>
				<option value="other">Other</option>
			</select>
		</div>
		<div class="space-y-2">
			<Label for="payment-date">Payment Date</Label>
			<Input id="payment-date" type="date" value={new Date().toISOString().split('T')[0]} />
		</div>
		<div class="space-y-2">
			<Label for="payment-note">Note (Optional)</Label>
			<Input id="payment-note" placeholder="Add a note..." />
		</div>
	</div>

	{#snippet footer()}
		<Button variant="outline" onclick={() => showRecordPaymentModal = false}>Cancel</Button>
		<Button onclick={handleRecordPayment} disabled={!paymentAmount}>
			<Icon icon="hugeicons:tick-double-01" class="mr-2 size-4" />
			Record Payment
		</Button>
	{/snippet}
</Modal>

<Modal
	bind:open={showSendReminderModal}
	title="Send Payment Reminder"
	description="Send a payment reminder email to the client."
	size="md"
>
	<div class="space-y-4">
		<div class="rounded-xl bg-muted p-4">
			<p class="text-sm font-medium">To: {invoice.client.email}</p>
			<p class="text-sm text-muted-foreground">Subject: Payment Reminder - Invoice {invoice.number}</p>
		</div>
		<div class="space-y-2">
			<Label for="reminder-message">Message</Label>
			<textarea
				id="reminder-message"
				rows={5}
				class="w-full resize-none rounded-md border border-input bg-background px-3 py-2 text-sm"
			>Dear {invoice.client.name},

I hope this email finds you well. This is a friendly reminder that invoice {invoice.number} for {formatCurrency(invoice.balance)} is due on {invoice.dueDate}.

Please let me know if you have any questions or need assistance with the payment.

Thank you for your business!

Best regards,
{company.name}</textarea>
		</div>
	</div>

	{#snippet footer()}
		<Button variant="outline" onclick={() => showSendReminderModal = false}>Cancel</Button>
		<Button>
			<Icon icon="hugeicons:send-01" class="mr-2 size-4" />
			Send Reminder
		</Button>
	{/snippet}
</Modal>

<Modal
	bind:open={showDeleteModal}
	title="Delete Invoice"
	description="Are you sure you want to delete invoice {invoice.number}? This action cannot be undone."
	size="sm"
>
	<div class="flex justify-end gap-2">
		<Button variant="outline" onclick={() => showDeleteModal = false}>Cancel</Button>
		<Button variant="destructive" href="/app/invoices">
			<Icon icon="hugeicons:delete-01" class="mr-2 size-4" />
			Delete Invoice
		</Button>
	</div>
</Modal>
