<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Field from '$lib/components/ui/field/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';

	const currentPlan = {
		name: 'Pro',
		price: 29,
		interval: 'month',
		features: [
			'Unlimited projects',
			'100GB storage',
			'Priority support',
			'Advanced analytics',
			'Custom domains',
			'Team collaboration'
		],
		nextBilling: 'Feb 27, 2026',
		status: 'active'
	};

	const paymentMethods = [
		{
			id: '1',
			type: 'card',
			brand: 'Visa',
			last4: '4242',
			expMonth: 12,
			expYear: 2027,
			isDefault: true
		},
		{
			id: '2',
			type: 'card',
			brand: 'Mastercard',
			last4: '8888',
			expMonth: 8,
			expYear: 2026,
			isDefault: false
		}
	];

	const invoices = [
		{ id: 'INV-001', date: 'Jan 27, 2026', amount: 29, status: 'paid' },
		{ id: 'INV-002', date: 'Dec 27, 2025', amount: 29, status: 'paid' },
		{ id: 'INV-003', date: 'Nov 27, 2025', amount: 29, status: 'paid' },
		{ id: 'INV-004', date: 'Oct 27, 2025', amount: 29, status: 'paid' }
	];

	const billingHistory = [
		{
			date: 'Jan 27, 2026',
			description: 'Pro Plan - Monthly',
			amount: -29,
			type: 'charge'
		},
		{
			date: 'Jan 15, 2026',
			description: 'Credit applied',
			amount: 10,
			type: 'credit'
		},
		{
			date: 'Dec 27, 2025',
			description: 'Pro Plan - Monthly',
			amount: -29,
			type: 'charge'
		}
	];

	const showCancelDialog = $state(false);
</script>

<svelte:head>
	<title>Billing Settings - Artemis</title>
</svelte:head>

<div class="space-y-8">
	<div>
		<h1 class="text-2xl font-semibold tracking-tight">Billing</h1>
		<p class="text-sm text-muted-foreground">
			Manage your subscription, payment methods, and billing history.
		</p>
	</div>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-center gap-3">
			<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
				<Icon icon="lucide:crown" class="size-5 text-primary" />
			</div>
			<div>
				<h2 class="text-lg font-medium">Current Plan</h2>
				<p class="text-sm text-muted-foreground">
					You are currently on the {currentPlan.name} plan.
				</p>
			</div>
		</div>

		<div class="rounded-xl border border-border bg-gradient-to-br from-primary/5 to-primary/10 p-6">
			<div class="flex flex-col gap-6 sm:flex-row sm:items-start sm:justify-between">
				<div class="space-y-2">
					<div class="flex items-center gap-2">
						<h3 class="text-2xl font-bold">{currentPlan.name}</h3>
						<span class="rounded-full bg-primary/10 px-2.5 py-0.5 text-xs font-medium text-primary">
							{currentPlan.status}
						</span>
					</div>
					<div class="flex items-baseline gap-1">
						<span class="text-3xl font-bold">${currentPlan.price}</span>
						<span class="text-muted-foreground">/{currentPlan.interval}</span>
					</div>
					<p class="text-sm text-muted-foreground">
						Next billing date: {currentPlan.nextBilling}
					</p>
				</div>
				<div class="flex gap-2">
					<Button variant="outline">Change plan</Button>
					<Button
						variant="outline"
						class="text-destructive hover:bg-destructive/10 hover:text-destructive"
					>
						Cancel
					</Button>
				</div>
			</div>

			<Separator class="my-6" />

			<div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
				{#each currentPlan.features as feature}
					<div class="flex items-center gap-2">
						<Icon icon="lucide:check" class="size-4 text-primary" />
						<span class="text-sm">{feature}</span>
					</div>
				{/each}
			</div>
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-start justify-between">
			<div class="flex items-center gap-3">
				<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
					<Icon icon="lucide:credit-card" class="size-5 text-primary" />
				</div>
				<div>
					<h2 class="text-lg font-medium">Payment Methods</h2>
					<p class="text-sm text-muted-foreground">
						Manage your payment methods and billing preferences.
					</p>
				</div>
			</div>
			<Button variant="outline" size="sm">
				<Icon icon="lucide:plus" class="mr-2 size-4" />
				Add method
			</Button>
		</div>

		<div class="space-y-3">
			{#each paymentMethods as method}
				<div class="flex items-center justify-between rounded-lg border border-border p-4">
					<div class="flex items-center gap-4">
						<div class="flex size-10 items-center justify-center rounded-md bg-muted">
							<Icon
								icon={method.brand === 'Visa' ? 'lucide:credit-card' : 'lucide:credit-card'}
								class="size-5 text-muted-foreground"
							/>
						</div>
						<div>
							<div class="flex items-center gap-2">
								<span class="font-medium">{method.brand} ending in {method.last4}</span>
								{#if method.isDefault}
									<span class="rounded-full bg-muted px-2 py-0.5 text-xs font-medium">
										Default
									</span>
								{/if}
							</div>
							<p class="text-sm text-muted-foreground">
								Expires {method.expMonth.toString().padStart(2, '0')}/{method.expYear}
							</p>
						</div>
					</div>
					<div class="flex items-center gap-2">
						{#if !method.isDefault}
							<Button variant="ghost" size="sm">Set default</Button>
						{/if}
						<Button variant="ghost" size="sm" class="text-destructive hover:text-destructive">
							<Icon icon="lucide:trash-2" class="size-4" />
						</Button>
					</div>
				</div>
			{/each}
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-center gap-3">
			<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
				<Icon icon="lucide:building-2" class="size-5 text-primary" />
			</div>
			<div>
				<h2 class="text-lg font-medium">Billing Information</h2>
				<p class="text-sm text-muted-foreground">
					Update your billing address and tax information.
				</p>
			</div>
		</div>

		<form class="grid gap-4 sm:grid-cols-2">
			<Field.Field>
				<Field.Label for="companyName">Company name (optional)</Field.Label>
				<Input id="companyName" placeholder="Your company" />
			</Field.Field>

			<Field.Field>
				<Field.Label for="taxId">Tax ID / VAT</Field.Label>
				<Input id="taxId" placeholder="XX-XXXXXXX" />
			</Field.Field>

			<Field.Field class="sm:col-span-2">
				<Field.Label for="address">Street address</Field.Label>
				<Input id="address" placeholder="123 Main St" />
			</Field.Field>

			<Field.Field>
				<Field.Label for="city">City</Field.Label>
				<Input id="city" placeholder="San Francisco" />
			</Field.Field>

			<Field.Field>
				<Field.Label for="postal">Postal code</Field.Label>
				<Input id="postal" placeholder="94102" />
			</Field.Field>

			<Field.Field>
				<Field.Label for="country">Country</Field.Label>
				<select
					id="country"
					class="h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-xs transition-colors focus-visible:border-ring focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				>
					<option value="us">United States</option>
					<option value="ca">Canada</option>
					<option value="uk">United Kingdom</option>
					<option value="de">Germany</option>
					<option value="fr">France</option>
					<option value="au">Australia</option>
				</select>
			</Field.Field>

			<Field.Field>
				<Field.Label for="state">State / Province</Field.Label>
				<Input id="state" placeholder="California" />
			</Field.Field>
		</form>

		<div class="flex justify-end pt-2">
			<Button>Update billing info</Button>
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-start justify-between">
			<div class="flex items-center gap-3">
				<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
					<Icon icon="lucide:file-text" class="size-5 text-primary" />
				</div>
				<div>
					<h2 class="text-lg font-medium">Invoices</h2>
					<p class="text-sm text-muted-foreground">Download your past invoices and receipts.</p>
				</div>
			</div>
			<Button variant="ghost" size="sm">
				<Icon icon="lucide:download" class="mr-2 size-4" />
				Download all
			</Button>
		</div>

		<div class="rounded-lg border">
			<table class="w-full text-sm">
				<thead class="border-b bg-muted/50">
					<tr>
						<th class="px-4 py-3 text-left font-medium">Invoice</th>
						<th class="px-4 py-3 text-left font-medium">Date</th>
						<th class="px-4 py-3 text-left font-medium">Amount</th>
						<th class="px-4 py-3 text-left font-medium">Status</th>
						<th class="px-4 py-3 text-right font-medium"></th>
					</tr>
				</thead>
				<tbody>
					{#each invoices as invoice}
						<tr class="border-b last:border-b-0">
							<td class="px-4 py-3 font-medium">{invoice.id}</td>
							<td class="px-4 py-3">{invoice.date}</td>
							<td class="px-4 py-3">${invoice.amount.toFixed(2)}</td>
							<td class="px-4 py-3">
								<span
									class="inline-flex items-center gap-1 rounded-full bg-emerald-100 px-2 py-0.5 text-xs font-medium text-emerald-700 dark:bg-emerald-900/30 dark:text-emerald-400"
								>
									<Icon icon="lucide:check" class="size-3" />
									{invoice.status}
								</span>
							</td>
							<td class="px-4 py-3 text-right">
								<Button variant="ghost" size="sm">
									<Icon icon="lucide:download" class="size-4" />
								</Button>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-center gap-3">
			<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
				<Icon icon="lucide:bar-chart-3" class="size-5 text-primary" />
			</div>
			<div>
				<h2 class="text-lg font-medium">Usage</h2>
				<p class="text-sm text-muted-foreground">
					Track your resource usage for this billing period.
				</p>
			</div>
		</div>

		<div class="grid gap-4 sm:grid-cols-2">
			<div class="rounded-lg border border-border p-4">
				<div class="mb-2 flex items-center justify-between">
					<span class="text-sm font-medium">Storage</span>
					<span class="text-sm text-muted-foreground">45GB / 100GB</span>
				</div>
				<div class="h-2 w-full overflow-hidden rounded-full bg-muted">
					<div class="h-full w-[45%] rounded-full bg-primary transition-all"></div>
				</div>
			</div>

			<div class="rounded-lg border border-border p-4">
				<div class="mb-2 flex items-center justify-between">
					<span class="text-sm font-medium">API Calls</span>
					<span class="text-sm text-muted-foreground">12,450 / 50,000</span>
				</div>
				<div class="h-2 w-full overflow-hidden rounded-full bg-muted">
					<div class="h-full w-[25%] rounded-full bg-primary transition-all"></div>
				</div>
			</div>

			<div class="rounded-lg border border-border p-4">
				<div class="mb-2 flex items-center justify-between">
					<span class="text-sm font-medium">Team Members</span>
					<span class="text-sm text-muted-foreground">3 / 10</span>
				</div>
				<div class="h-2 w-full overflow-hidden rounded-full bg-muted">
					<div class="h-full w-[30%] rounded-full bg-primary transition-all"></div>
				</div>
			</div>

			<div class="rounded-lg border border-border p-4">
				<div class="mb-2 flex items-center justify-between">
					<span class="text-sm font-medium">Projects</span>
					<span class="text-sm text-muted-foreground">8 / Unlimited</span>
				</div>
				<div class="h-2 w-full overflow-hidden rounded-full bg-muted">
					<div class="h-full w-[8%] rounded-full bg-primary transition-all"></div>
				</div>
			</div>
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-center gap-3">
			<div class="flex size-10 items-center justify-center rounded-lg bg-destructive/10">
				<Icon icon="lucide:alert-triangle" class="size-5 text-destructive" />
			</div>
			<div>
				<h2 class="text-lg font-medium text-destructive">Danger Zone</h2>
				<p class="text-sm text-muted-foreground">Irreversible and destructive actions.</p>
			</div>
		</div>

		<div class="rounded-lg border border-destructive/20 bg-destructive/5 p-4">
			<div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
				<div>
					<h3 class="font-medium text-destructive">Cancel Subscription</h3>
					<p class="text-sm text-muted-foreground">
						Your subscription will be canceled at the end of the billing period.
					</p>
				</div>
				<Button variant="destructive">Cancel subscription</Button>
			</div>
		</div>
	</section>
</div>
