<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Field from '$lib/components/ui/field/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';

	const sessions = [
		{
			id: '1',
			device: 'MacBook Pro',
			location: 'San Francisco, CA',
			ip: '192.168.1.1',
			current: true,
			lastActive: 'Active now'
		},
		{
			id: '2',
			device: 'iPhone 15 Pro',
			location: 'San Francisco, CA',
			ip: '192.168.1.2',
			current: false,
			lastActive: '2 hours ago'
		},
		{
			id: '3',
			device: 'Chrome on Windows',
			location: 'New York, NY',
			ip: '203.0.113.1',
			current: false,
			lastActive: '3 days ago'
		}
	];

	let showPasswordForm = $state(false);
	const showTwoFactorSetup = $state(false);
	let isLoading = $state(false);

	async function handlePasswordChange() {
		isLoading = true;
		await new Promise((resolve) => setTimeout(resolve, 1000));
		isLoading = false;
		showPasswordForm = false;
	}
</script>

<svelte:head>
	<title>Security Settings - Artemis</title>
</svelte:head>

<div class="space-y-8">
	<div>
		<h1 class="text-2xl font-semibold tracking-tight">Security</h1>
		<p class="text-sm text-muted-foreground">
			Manage your password, two-factor authentication, and active sessions.
		</p>
	</div>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-start justify-between">
			<div class="space-y-1">
				<h2 class="text-lg font-medium">Password</h2>
				<p class="text-sm text-muted-foreground">
					Last changed 3 months ago. We recommend changing your password regularly.
				</p>
			</div>
			{#if !showPasswordForm}
				<Button variant="outline" onclick={() => (showPasswordForm = true)}>
					<Icon icon="lucide:key" class="mr-2 size-4" />
					Change password
				</Button>
			{/if}
		</div>

		{#if showPasswordForm}
			<form
				onsubmit={(e) => {
					e.preventDefault();
					handlePasswordChange();
				}}
				class="rounded-lg border border-border bg-card p-4"
			>
				<div class="space-y-4">
					<Field.Field>
						<Field.Label for="currentPassword">Current password</Field.Label>
						<Input id="currentPassword" type="password" placeholder="Enter your current password" />
					</Field.Field>

					<Field.Field>
						<Field.Label for="newPassword">New password</Field.Label>
						<Input id="newPassword" type="password" placeholder="Enter new password" />
						<Field.Description
							>Must be at least 8 characters with numbers and symbols.</Field.Description
						>
					</Field.Field>

					<Field.Field>
						<Field.Label for="confirmPassword">Confirm new password</Field.Label>
						<Input id="confirmPassword" type="password" placeholder="Confirm new password" />
					</Field.Field>
				</div>
				<div class="mt-4 flex items-center gap-2">
					<Button type="submit" disabled={isLoading}>
						{#if isLoading}
							<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
						{/if}
						Update password
					</Button>
					<Button type="button" variant="ghost" onclick={() => (showPasswordForm = false)}>
						Cancel
					</Button>
				</div>
			</form>
		{/if}
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-start justify-between">
			<div class="space-y-1">
				<h2 class="text-lg font-medium">Two-Factor Authentication</h2>
				<p class="text-sm text-muted-foreground">
					Add an extra layer of security to your account by requiring a verification code.
				</p>
			</div>
			<Button variant="outline" onclick={() => (showTwoFactorSetup = !showTwoFactorSetup)}>
				<Icon icon="lucide:shield-check" class="mr-2 size-4" />
				Enable 2FA
			</Button>
		</div>

		{#if showTwoFactorSetup}
			<div class="rounded-lg border border-border bg-card p-6">
				<div class="flex flex-col items-center gap-6 text-center sm:flex-row sm:text-left">
					<div class="size-32 shrink-0 rounded-lg bg-muted p-2">
						<div class="grid h-full w-full grid-cols-5 grid-rows-5 gap-1">
							{#each Array(25) as _, i}
								<div
									class="{Math.random() > 0.5 ? 'bg-foreground' : 'bg-transparent'} rounded-sm"
								></div>
							{/each}
						</div>
					</div>
					<div class="space-y-4">
						<h3 class="font-medium">Scan with your authenticator app</h3>
						<p class="text-sm text-muted-foreground">
							Scan the QR code above with your preferred authenticator app (Google Authenticator,
							Authy, etc.) or enter the setup key manually.
						</p>
						<div class="flex items-center gap-2">
							<code class="rounded bg-muted px-3 py-1.5 font-mono text-sm">XXXX-XXXX-XXXX-XXXX</code
							>
							<Button variant="ghost" size="sm" class="h-8">
								<Icon icon="lucide:copy" class="mr-1 size-4" />
								Copy
							</Button>
						</div>
						<div class="flex gap-2">
							<Input
								placeholder="Enter 6-digit code"
								class="w-40 text-center text-lg tracking-widest"
								maxlength={6}
							/>
							<Button>Verify</Button>
						</div>
					</div>
				</div>
			</div>
		{/if}

		<div class="rounded-lg border border-border/60 bg-muted/30 p-4">
			<div class="flex items-start gap-4">
				<div class="mt-1 shrink-0">
					<Icon icon="lucide:shield-alert" class="size-5 text-amber-500" />
				</div>
				<div class="flex-1 space-y-2">
					<h3 class="font-medium">Backup Codes</h3>
					<p class="text-sm text-muted-foreground">
						Save these codes in a secure place. You can use them to access your account if you lose
						your authenticator device.
					</p>
					<Button variant="outline" size="sm">
						<Icon icon="lucide:download" class="mr-2 size-4" />
						Download backup codes
					</Button>
				</div>
			</div>
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-start justify-between">
			<div class="space-y-1">
				<h2 class="text-lg font-medium">Active Sessions</h2>
				<p class="text-sm text-muted-foreground">Manage your active sessions across all devices.</p>
			</div>
			<Button
				variant="outline"
				size="sm"
				class="text-destructive hover:bg-destructive/10 hover:text-destructive"
			>
				<Icon icon="lucide:log-out" class="mr-2 size-4" />
				Sign out all
			</Button>
		</div>

		<div class="space-y-3">
			{#each sessions as session}
				<div class="flex items-center justify-between rounded-lg border border-border p-4">
					<div class="flex items-center gap-4">
						<div class="flex size-10 items-center justify-center rounded-full bg-muted">
							<Icon
								icon={session.device.includes('iPhone')
									? 'lucide:smartphone'
									: session.device.includes('Mac')
										? 'lucide:laptop'
										: 'lucide:monitor'}
								class="size-5 text-muted-foreground"
							/>
						</div>
						<div>
							<div class="flex items-center gap-2">
								<span class="font-medium">{session.device}</span>
								{#if session.current}
									<span
										class="rounded-full bg-primary/10 px-2 py-0.5 text-xs font-medium text-primary"
									>
										Current
									</span>
								{/if}
							</div>
							<p class="text-sm text-muted-foreground">
								{session.location} · {session.ip} · {session.lastActive}
							</p>
						</div>
					</div>
					{#if !session.current}
						<Button variant="ghost" size="sm">
							<Icon icon="lucide:x" class="size-4" />
							<span class="sr-only">Revoke</span>
						</Button>
					{/if}
				</div>
			{/each}
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-start justify-between">
			<div class="space-y-1">
				<h2 class="text-lg font-medium">Login History</h2>
				<p class="text-sm text-muted-foreground">Recent login activity on your account.</p>
			</div>
			<Button variant="ghost" size="sm">
				View all
				<Icon icon="lucide:arrow-right" class="ml-1 size-4" />
			</Button>
		</div>

		<div class="rounded-lg border">
			<table class="w-full text-sm">
				<thead class="border-b bg-muted/50">
					<tr>
						<th class="px-4 py-3 text-left font-medium">Date & Time</th>
						<th class="px-4 py-3 text-left font-medium">Device</th>
						<th class="px-4 py-3 text-left font-medium">Location</th>
						<th class="px-4 py-3 text-left font-medium">Status</th>
					</tr>
				</thead>
				<tbody>
					<tr class="border-b last:border-b-0">
						<td class="px-4 py-3">Jan 27, 2026 · 9:42 AM</td>
						<td class="px-4 py-3">Chrome on macOS</td>
						<td class="px-4 py-3">San Francisco, CA</td>
						<td class="px-4 py-3">
							<span class="inline-flex items-center gap-1 text-emerald-600">
								<Icon icon="lucide:check-circle" class="size-4" />
								Success
							</span>
						</td>
					</tr>
					<tr class="border-b last:border-b-0">
						<td class="px-4 py-3">Jan 26, 2026 · 3:15 PM</td>
						<td class="px-4 py-3">Safari on iOS</td>
						<td class="px-4 py-3">San Francisco, CA</td>
						<td class="px-4 py-3">
							<span class="inline-flex items-center gap-1 text-emerald-600">
								<Icon icon="lucide:check-circle" class="size-4" />
								Success
							</span>
						</td>
					</tr>
					<tr>
						<td class="px-4 py-3">Jan 25, 2026 · 11:30 AM</td>
						<td class="px-4 py-3">Firefox on Windows</td>
						<td class="px-4 py-3">New York, NY</td>
						<td class="px-4 py-3">
							<span class="inline-flex items-center gap-1 text-amber-600">
								<Icon icon="lucide:alert-circle" class="size-4" />
								Failed
							</span>
						</td>
					</tr>
				</tbody>
			</table>
		</div>
	</section>
</div>
