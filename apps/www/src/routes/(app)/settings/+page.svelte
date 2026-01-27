<script lang="ts">
	import { goto } from '$app/navigation';
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { toggleMode, mode } from 'mode-watcher';

	const user = {
		firstName: 'Alex',
		lastName: 'Morgan',
		email: 'alex.morgan@example.com',
		username: 'alexmorgan',
		bio: 'Product designer and developer based in San Francisco. Passionate about creating beautiful and functional user experiences.',
		location: 'San Francisco, CA',
		website: 'https://alexmorgan.dev',
		company: 'Artemis Inc.',
		role: 'Senior Product Designer',
		phone: '+1 (555) 123-4567',
		timezone: 'America/Los_Angeles',
		language: 'en',
		avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex'
	};

	const workspace = {
		name: 'Design Studio',
		slug: 'design-studio',
		plan: 'Pro',
		seats: 5,
		usedSeats: 3,
		billingEmail: 'billing@designstudio.com'
	};

	const sessions = [
		{
			id: '1',
			device: 'MacBook Pro',
			browser: 'Chrome',
			os: 'macOS',
			location: 'San Francisco, CA',
			ip: '192.168.1.1',
			current: true,
			lastActive: 'Active now'
		},
		{
			id: '2',
			device: 'iPhone 15 Pro',
			browser: 'Safari',
			os: 'iOS',
			location: 'San Francisco, CA',
			ip: '192.168.1.2',
			current: false,
			lastActive: '2 hours ago'
		},
		{
			id: '3',
			device: 'Windows PC',
			browser: 'Edge',
			os: 'Windows 11',
			location: 'New York, NY',
			ip: '203.0.113.1',
			current: false,
			lastActive: '3 days ago'
		}
	];

	const teamMembers = [
		{
			id: 1,
			name: 'Alex Morgan',
			email: 'alex@designstudio.com',
			role: 'owner',
			status: 'active',
			avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex',
			joined: 'Jan 2023'
		},
		{
			id: 2,
			name: 'Sarah Chen',
			email: 'sarah@designstudio.com',
			role: 'admin',
			status: 'active',
			avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah',
			joined: 'Mar 2023'
		},
		{
			id: 3,
			name: 'Mike Ross',
			email: 'mike@designstudio.com',
			role: 'member',
			status: 'active',
			avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Mike',
			joined: 'Jun 2023'
		}
	];

	const invoices = [
		{
			id: 'INV-2026-001',
			date: 'Jan 27, 2026',
			amount: 29.0,
			status: 'paid',
			description: 'Pro Plan - Monthly'
		},
		{
			id: 'INV-2025-012',
			date: 'Dec 27, 2025',
			amount: 29.0,
			status: 'paid',
			description: 'Pro Plan - Monthly'
		},
		{
			id: 'INV-2025-011',
			date: 'Nov 27, 2025',
			amount: 29.0,
			status: 'paid',
			description: 'Pro Plan - Monthly'
		},
		{
			id: 'INV-2025-010',
			date: 'Oct 27, 2025',
			amount: 29.0,
			status: 'paid',
			description: 'Pro Plan - Monthly'
		}
	];

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

	let activeTab = $state('profile');
	let saving = $state<string | null>(null);
	let showPasswordForm = $state(false);
	let show2FASetup = $state(false);
	let showInviteModal = $state(false);
	let quietHours = $state({ enabled: false, start: '22:00', end: '08:00' });
	let notificationPrefs = $state({
		email: {
			marketing: true,
			social: true,
			updates: true,
			security: true,
			newsletter: false,
			digest: 'realtime'
		},
		push: { mentions: true, messages: true, comments: true, reminders: false, tasks: true }
	});

	const tabs = [
		{ id: 'profile', label: 'Profile', icon: 'lucide:user' },
		{ id: 'security', label: 'Security', icon: 'lucide:shield' },
		{ id: 'notifications', label: 'Notifications', icon: 'lucide:bell' },
		{ id: 'appearance', label: 'Appearance', icon: 'lucide:palette' },
		{ id: 'billing', label: 'Billing', icon: 'lucide:credit-card' },
		{ id: 'workspace', label: 'Workspace', icon: 'lucide:building-2' }
	];

	const accentColors = [
		{ name: 'Emerald', value: 'emerald', class: 'bg-emerald-500' },
		{ name: 'Blue', value: 'blue', class: 'bg-blue-500' },
		{ name: 'Violet', value: 'violet', class: 'bg-violet-500' },
		{ name: 'Rose', value: 'rose', class: 'bg-rose-500' },
		{ name: 'Amber', value: 'amber', class: 'bg-amber-500' },
		{ name: 'Cyan', value: 'cyan', class: 'bg-cyan-500' }
	];

	async function saveSection(section: string) {
		saving = section;
		await new Promise((resolve) => setTimeout(resolve, 800));
		saving = null;
	}

	function closeSettings() {
		goto('/');
	}

	function revokeSession(sessionId: string) {
		console.log('Revoking session:', sessionId);
	}

	function downloadInvoice(invoiceId: string) {
		console.log('Downloading invoice:', invoiceId);
	}

	function getRoleBadgeColor(role: string) {
		switch (role) {
			case 'owner':
				return 'bg-amber-100 text-amber-700 dark:bg-amber-900/30 dark:text-amber-400';
			case 'admin':
				return 'bg-blue-100 text-blue-700 dark:bg-blue-900/30 dark:text-blue-400';
			default:
				return 'bg-muted text-muted-foreground';
		}
	}

	function getDeviceIcon(device: string) {
		if (device.includes('iPhone') || device.includes('Android')) return 'lucide:smartphone';
		if (device.includes('iPad') || device.includes('Tablet')) return 'lucide:tablet';
		return 'lucide:laptop';
	}
</script>

<svelte:head>
	<title>Settings - Artemis</title>
</svelte:head>

<div class="fixed inset-0 z-50 bg-background">
	<header class="flex h-14 items-center border-b border-border px-4">
		<div class="flex items-center gap-4">
			<Button variant="ghost" size="icon" onclick={closeSettings}>
				<Icon icon="lucide:x" class="size-5" />
			</Button>
			<h1 class="text-lg font-semibold">Settings</h1>
		</div>
		<div class="flex-1"></div>
	</header>

	<div class="flex h-[calc(100vh-3.5rem)]">
		<aside class="hidden w-64 shrink-0 border-r border-border bg-muted/30 md:block">
			<nav class="space-y-1 p-3">
				{#each tabs as tab}
					<button
						class="flex w-full items-center gap-3 rounded-lg px-3 py-2 text-sm transition-colors {activeTab ===
						tab.id
							? 'bg-primary/10 font-medium text-primary'
							: 'text-muted-foreground hover:bg-muted hover:text-foreground'}"
						onclick={() => (activeTab = tab.id)}
					>
						<Icon icon={tab.icon} class="size-4" />
						{tab.label}
					</button>
				{/each}
			</nav>
		</aside>

		<div class="absolute right-0 bottom-0 left-0 border-t border-border bg-background md:hidden">
			<div class="flex gap-1 overflow-x-auto p-2">
				{#each tabs as tab}
					<button
						class="flex shrink-0 items-center gap-2 rounded-lg px-3 py-2 text-sm whitespace-nowrap transition-colors {activeTab ===
						tab.id
							? 'bg-primary/10 font-medium text-primary'
							: 'text-muted-foreground hover:bg-muted'}"
						onclick={() => (activeTab = tab.id)}
					>
						<Icon icon={tab.icon} class="size-4" />
						{tab.label}
					</button>
				{/each}
			</div>
		</div>

		<main class="flex-1 overflow-y-auto p-4 pb-24 md:p-8 md:pb-8">
			<div class="mx-auto max-w-3xl">
				{#if activeTab === 'profile'}
					<div class="space-y-8">
						<div>
							<h2 class="text-2xl font-semibold tracking-tight">Profile</h2>
							<p class="text-muted-foreground">
								Manage your personal information and how others see you.
							</p>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<div class="flex flex-col gap-6 sm:flex-row">
								<div class="mx-auto shrink-0 sm:mx-0">
									<div class="relative">
										<img
											src={user.avatar}
											alt="Profile"
											class="size-28 rounded-full bg-muted object-cover ring-4 ring-background"
										/>
										<button
											class="absolute -right-1 -bottom-1 flex size-9 items-center justify-center rounded-full bg-primary text-primary-foreground shadow-md transition-colors hover:bg-primary/90"
										>
											<Icon icon="lucide:camera" class="size-4" />
										</button>
									</div>
								</div>
								<div class="flex-1 text-center sm:text-left">
									<h3 class="text-lg font-semibold">Profile Photo</h3>
									<p class="mt-1 text-sm text-muted-foreground">
										This will be displayed on your profile and throughout the platform. Recommended:
										500x500px or larger.
									</p>
									<div class="mt-4 flex justify-center gap-2 sm:justify-start">
										<Button variant="outline" size="sm">
											<Icon icon="lucide:upload" class="mr-2 size-4" />
											Upload new
										</Button>
										<Button
											variant="ghost"
											size="sm"
											class="text-destructive hover:bg-destructive/10 hover:text-destructive"
										>
											<Icon icon="lucide:trash-2" class="mr-2 size-4" />
											Remove
										</Button>
									</div>
								</div>
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Personal Information</h3>
							<div class="grid gap-4 sm:grid-cols-2">
								<div class="space-y-2">
									<label for="firstName" class="text-sm font-medium">First name</label>
									<Input id="firstName" value={user.firstName} />
								</div>
								<div class="space-y-2">
									<label for="lastName" class="text-sm font-medium">Last name</label>
									<Input id="lastName" value={user.lastName} />
								</div>
							</div>
							<div class="mt-4 grid gap-4 sm:grid-cols-2">
								<div class="space-y-2">
									<label for="email" class="text-sm font-medium">Email address</label>
									<Input id="email" type="email" value={user.email} />
									<p class="text-xs text-muted-foreground">
										This email is used for notifications and sign-in.
									</p>
								</div>
								<div class="space-y-2">
									<label for="phone" class="text-sm font-medium">Phone number</label>
									<Input id="phone" type="tel" value={user.phone} placeholder="+1 (555) 000-0000" />
								</div>
							</div>
							<div class="mt-4 space-y-2">
								<label for="username" class="text-sm font-medium">Username</label>
								<div class="relative">
									<span class="absolute top-1/2 left-3 -translate-y-1/2 text-muted-foreground"
										>@</span
									>
									<Input id="username" value={user.username} class="pl-7" />
								</div>
								<p class="text-xs text-muted-foreground">Your public username for profile URL.</p>
							</div>
							<div class="mt-4 space-y-2">
								<label for="bio" class="text-sm font-medium">Bio</label>
								<textarea
									id="bio"
									rows="4"
									class="w-full resize-none rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none"
									>{user.bio}</textarea
								>
								<p class="text-xs text-muted-foreground">
									Brief description for your profile. Maximum 500 characters.
								</p>
							</div>
							<div class="mt-6 flex justify-end">
								<Button onclick={() => saveSection('personal')} disabled={saving === 'personal'}>
									{#if saving === 'personal'}
										<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
									{/if}
									Save changes
								</Button>
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Professional</h3>
							<div class="grid gap-4 sm:grid-cols-2">
								<div class="space-y-2">
									<label for="company" class="text-sm font-medium">Company</label>
									<Input id="company" value={user.company} />
								</div>
								<div class="space-y-2">
									<label for="role" class="text-sm font-medium">Job title</label>
									<Input id="role" value={user.role} />
								</div>
							</div>
							<div class="mt-4 space-y-2">
								<label for="location" class="text-sm font-medium">Location</label>
								<Input id="location" value={user.location} placeholder="City, Country" />
							</div>
							<div class="mt-4 space-y-2">
								<label for="website" class="text-sm font-medium">Website</label>
								<div class="relative">
									<span class="absolute top-1/2 left-3 -translate-y-1/2 text-muted-foreground"
										><Icon icon="lucide:globe" class="size-4" /></span
									>
									<Input
										id="website"
										type="url"
										value={user.website}
										class="pl-10"
										placeholder="https://your-website.com"
									/>
								</div>
							</div>
							<div class="mt-6 flex justify-end">
								<Button
									onclick={() => saveSection('professional')}
									disabled={saving === 'professional'}
								>
									{#if saving === 'professional'}
										<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
									{/if}
									Save changes
								</Button>
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
								<div>
									<h3 class="font-semibold">Public Profile</h3>
									<p class="mt-1 text-sm text-muted-foreground">
										Allow others to see your profile and activity on the platform.
									</p>
								</div>
								<label class="relative inline-flex shrink-0 cursor-pointer items-center">
									<input type="checkbox" checked class="peer sr-only" />
									<div
										class="h-6 w-11 rounded-full bg-muted transition-colors peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-transform after:content-[''] peer-checked:after:translate-x-5"
									></div>
								</label>
							</div>
						</div>
					</div>
				{:else if activeTab === 'security'}
					<div class="space-y-8">
						<div>
							<h2 class="text-2xl font-semibold tracking-tight">Security</h2>
							<p class="text-muted-foreground">
								Manage your password, two-factor authentication, and active sessions.
							</p>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							{#if !showPasswordForm}
								<div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
									<div class="flex items-start gap-4">
										<div
											class="flex size-10 shrink-0 items-center justify-center rounded-lg bg-muted"
										>
											<Icon icon="lucide:key" class="size-5 text-muted-foreground" />
										</div>
										<div>
											<h3 class="font-semibold">Password</h3>
											<p class="text-sm text-muted-foreground">
												Last changed 3 months ago. We recommend changing your password regularly.
											</p>
										</div>
									</div>
									<Button variant="outline" onclick={() => (showPasswordForm = true)}
										>Change password</Button
									>
								</div>
							{:else}
								<div class="space-y-4">
									<div class="mb-4 flex items-center gap-2">
										<Button variant="ghost" size="icon" onclick={() => (showPasswordForm = false)}>
											<Icon icon="lucide:arrow-left" class="size-4" />
										</Button>
										<h3 class="font-semibold">Change Password</h3>
									</div>
									<div class="space-y-2">
										<label for="currentPassword" class="text-sm font-medium">Current password</label
										>
										<Input
											id="currentPassword"
											type="password"
											placeholder="Enter your current password"
										/>
									</div>
									<div class="space-y-2">
										<label for="newPassword" class="text-sm font-medium">New password</label>
										<Input id="newPassword" type="password" placeholder="Enter new password" />
										<p class="text-xs text-muted-foreground">
											Must be at least 8 characters with uppercase, number, and special character.
										</p>
									</div>
									<div class="space-y-2">
										<label for="confirmPassword" class="text-sm font-medium"
											>Confirm new password</label
										>
										<Input
											id="confirmPassword"
											type="password"
											placeholder="Confirm new password"
										/>
									</div>
									<div class="flex gap-2 pt-2">
										<Button
											onclick={() => {
												saveSection('password');
												showPasswordForm = false;
											}}
											disabled={saving === 'password'}
										>
											{#if saving === 'password'}
												<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
											{/if}
											Update password
										</Button>
										<Button variant="ghost" onclick={() => (showPasswordForm = false)}
											>Cancel</Button
										>
									</div>
								</div>
							{/if}
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							{#if !show2FASetup}
								<div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
									<div class="flex items-start gap-4">
										<div
											class="flex size-10 shrink-0 items-center justify-center rounded-lg bg-muted"
										>
											<Icon icon="lucide:shield-check" class="size-5 text-muted-foreground" />
										</div>
										<div>
											<h3 class="font-semibold">Two-Factor Authentication</h3>
											<p class="text-sm text-muted-foreground">
												Add an extra layer of security to your account by requiring a verification
												code.
											</p>
										</div>
									</div>
									<Button variant="outline" onclick={() => (show2FASetup = true)}>Enable 2FA</Button
									>
								</div>
							{:else}
								<div class="space-y-6">
									<div class="flex items-center gap-2">
										<Button variant="ghost" size="icon" onclick={() => (show2FASetup = false)}>
											<Icon icon="lucide:arrow-left" class="size-4" />
										</Button>
										<h3 class="font-semibold">Set Up Two-Factor Authentication</h3>
									</div>
									<div class="flex flex-col gap-6 md:flex-row">
										<div class="mx-auto size-40 shrink-0 rounded-lg bg-muted p-2 md:mx-0">
											<div class="grid h-full w-full grid-cols-5 grid-rows-5 gap-1">
												{#each Array(25) as _, i}
													<div
														class="{Math.random() > 0.5
															? 'bg-foreground'
															: 'bg-transparent'} rounded-sm"
													></div>
												{/each}
											</div>
										</div>
										<div class="flex-1 space-y-4">
											<p class="text-sm text-muted-foreground">
												Scan this QR code with your authenticator app (Google Authenticator, Authy,
												or 1Password).
											</p>
											<div class="flex items-center gap-2">
												<code class="flex-1 rounded bg-muted px-3 py-1.5 font-mono text-sm"
													>XXXX-XXXX-XXXX-XXXX</code
												>
												<Button variant="ghost" size="sm">
													<Icon icon="lucide:copy" class="size-4" />
												</Button>
											</div>
											<div class="space-y-2">
												<label for="verifyCode" class="text-sm font-medium">Verification code</label
												>
												<div class="flex gap-2">
													<Input
														id="verifyCode"
														placeholder="000000"
														class="w-28 text-center text-lg tracking-widest"
														maxlength={6}
													/>
													<Button onclick={() => saveSection('2fa')}>Verify</Button>
												</div>
											</div>
										</div>
									</div>
								</div>
							{/if}
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<div class="flex items-start gap-4">
								<div
									class="flex size-10 shrink-0 items-center justify-center rounded-lg bg-amber-500/10"
								>
									<Icon icon="lucide:shield-alert" class="size-5 text-amber-500" />
								</div>
								<div class="flex-1">
									<h3 class="font-semibold">Backup Codes</h3>
									<p class="mt-1 text-sm text-muted-foreground">
										Save these codes in a secure place. You can use them to access your account if
										you lose your authenticator device.
									</p>
									<Button variant="outline" size="sm" class="mt-3">
										<Icon icon="lucide:download" class="mr-2 size-4" />
										Download backup codes
									</Button>
								</div>
							</div>
						</div>

						<div class="overflow-hidden rounded-xl border border-border bg-card">
							<div class="border-b border-border p-4">
								<h3 class="font-semibold">Active Sessions</h3>
								<p class="text-sm text-muted-foreground">
									Manage devices that are currently signed in to your account.
								</p>
							</div>
							<div class="divide-y divide-border">
								{#each sessions as session}
									<div
										class="flex flex-col gap-4 p-4 sm:flex-row sm:items-center sm:justify-between"
									>
										<div class="flex items-center gap-4">
											<div
												class="flex size-10 shrink-0 items-center justify-center rounded-lg bg-muted"
											>
												<Icon
													icon={getDeviceIcon(session.device)}
													class="size-5 text-muted-foreground"
												/>
											</div>
											<div class="min-w-0">
												<div class="flex flex-wrap items-center gap-2">
													<span class="text-sm font-medium">{session.device}</span>
													{#if session.current}
														<span
															class="rounded-full bg-primary/10 px-2 py-0.5 text-xs text-primary"
															>Current</span
														>
													{/if}
												</div>
												<p class="truncate text-xs text-muted-foreground">
													{session.browser} on {session.os} · {session.location}
												</p>
												<p class="text-xs text-muted-foreground">{session.lastActive}</p>
											</div>
										</div>
										{#if !session.current}
											<Button
												variant="ghost"
												size="sm"
												class="self-start text-destructive sm:self-center"
												onclick={() => revokeSession(session.id)}
											>
												Revoke
											</Button>
										{/if}
									</div>
								{/each}
							</div>
							<div class="border-t border-border bg-muted/30 p-4">
								<Button variant="outline" size="sm" class="w-full">
									<Icon icon="lucide:log-out" class="mr-2 size-4" />
									Sign out all other devices
								</Button>
							</div>
						</div>

						<div class="overflow-hidden rounded-xl border border-border bg-card">
							<div class="border-b border-border p-4">
								<h3 class="font-semibold">Recent Security Activity</h3>
							</div>
							<div class="divide-y divide-border">
								{#each [{ action: 'Successful sign-in', location: 'San Francisco, CA', device: 'Chrome on macOS', time: '2 hours ago', status: 'success' }, { action: 'Password changed', location: 'San Francisco, CA', device: 'Safari on macOS', time: '3 months ago', status: 'success' }, { action: 'Failed sign-in attempt', location: 'New York, NY', device: 'Firefox on Windows', time: '3 days ago', status: 'failed' }] as activity}
									<div class="flex items-start gap-4 p-4">
										<div
											class="size-8 rounded-full {activity.status === 'failed'
												? 'bg-destructive/10'
												: 'bg-emerald-500/10'} mt-0.5 flex shrink-0 items-center justify-center"
										>
											<Icon
												icon={activity.status === 'failed' ? 'lucide:x' : 'lucide:check'}
												class="size-4 {activity.status === 'failed'
													? 'text-destructive'
													: 'text-emerald-500'}"
											/>
										</div>
										<div class="min-w-0 flex-1">
											<div
												class="flex flex-col gap-1 sm:flex-row sm:items-center sm:justify-between"
											>
												<p class="text-sm font-medium">{activity.action}</p>
												<span class="text-xs text-muted-foreground">{activity.time}</span>
											</div>
											<p class="text-xs text-muted-foreground">
												{activity.device} · {activity.location}
											</p>
										</div>
									</div>
								{/each}
							</div>
						</div>
					</div>
				{:else if activeTab === 'notifications'}
					<div class="space-y-8">
						<div>
							<h2 class="text-2xl font-semibold tracking-tight">Notifications</h2>
							<p class="text-muted-foreground">
								Choose what notifications you receive and how you're notified.
							</p>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Email Notifications</h3>
							<div class="space-y-4">
								<label
									class="-mx-3 flex cursor-pointer items-start justify-between gap-4 rounded-lg p-3 transition-colors hover:bg-muted/50"
								>
									<div>
										<p class="text-sm font-medium">Marketing & Promotions</p>
										<p class="text-xs text-muted-foreground">
											New features, offers, and product announcements
										</p>
									</div>
									<input
										type="checkbox"
										bind:checked={notificationPrefs.email.marketing}
										class="mt-0.5 size-4 shrink-0 rounded border-input text-primary focus:ring-primary"
									/>
								</label>
								<label
									class="-mx-3 flex cursor-pointer items-start justify-between gap-4 rounded-lg p-3 transition-colors hover:bg-muted/50"
								>
									<div>
										<p class="text-sm font-medium">Social Activity</p>
										<p class="text-xs text-muted-foreground">
											When someone follows you, mentions you, or interacts with your content
										</p>
									</div>
									<input
										type="checkbox"
										bind:checked={notificationPrefs.email.social}
										class="mt-0.5 size-4 shrink-0 rounded border-input text-primary focus:ring-primary"
									/>
								</label>
								<label
									class="-mx-3 flex cursor-pointer items-start justify-between gap-4 rounded-lg p-3 transition-colors hover:bg-muted/50"
								>
									<div>
										<p class="text-sm font-medium">Product Updates</p>
										<p class="text-xs text-muted-foreground">
											New features, improvements, and changelog updates
										</p>
									</div>
									<input
										type="checkbox"
										bind:checked={notificationPrefs.email.updates}
										class="mt-0.5 size-4 shrink-0 rounded border-input text-primary focus:ring-primary"
									/>
								</label>
								<label
									class="-mx-3 flex cursor-pointer items-start justify-between gap-4 rounded-lg p-3 transition-colors hover:bg-muted/50"
								>
									<div>
										<p class="text-sm font-medium">Weekly Newsletter</p>
										<p class="text-xs text-muted-foreground">
											Industry news, tips, and best practices
										</p>
									</div>
									<input
										type="checkbox"
										bind:checked={notificationPrefs.email.newsletter}
										class="mt-0.5 size-4 shrink-0 rounded border-input text-primary focus:ring-primary"
									/>
								</label>
								<Separator />
								<label
									class="-mx-3 flex cursor-pointer items-start justify-between gap-4 rounded-lg p-3 opacity-75 transition-colors hover:bg-muted/50"
								>
									<div>
										<p class="text-sm font-medium">Security Alerts</p>
										<p class="text-xs text-muted-foreground">
											Important security notifications about your account
										</p>
									</div>
									<input
										type="checkbox"
										checked
										disabled
										class="mt-0.5 size-4 shrink-0 rounded border-input"
									/>
								</label>
							</div>
							<div class="mt-6 border-t border-border pt-4">
								<Button
									onclick={() => saveSection('email-notifications')}
									disabled={saving === 'email-notifications'}
								>
									{#if saving === 'email-notifications'}
										<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
									{/if}
									Save email preferences
								</Button>
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Push Notifications</h3>
							<div class="space-y-4">
								<label
									class="-mx-3 flex cursor-pointer items-start justify-between gap-4 rounded-lg p-3 transition-colors hover:bg-muted/50"
								>
									<div>
										<p class="text-sm font-medium">Mentions & Tags</p>
										<p class="text-xs text-muted-foreground">When someone mentions or tags you</p>
									</div>
									<input
										type="checkbox"
										bind:checked={notificationPrefs.push.mentions}
										class="mt-0.5 size-4 shrink-0 rounded border-input text-primary focus:ring-primary"
									/>
								</label>
								<label
									class="-mx-3 flex cursor-pointer items-start justify-between gap-4 rounded-lg p-3 transition-colors hover:bg-muted/50"
								>
									<div>
										<p class="text-sm font-medium">Direct Messages</p>
										<p class="text-xs text-muted-foreground">
											New messages from team members and clients
										</p>
									</div>
									<input
										type="checkbox"
										bind:checked={notificationPrefs.push.messages}
										class="mt-0.5 size-4 shrink-0 rounded border-input text-primary focus:ring-primary"
									/>
								</label>
								<label
									class="-mx-3 flex cursor-pointer items-start justify-between gap-4 rounded-lg p-3 transition-colors hover:bg-muted/50"
								>
									<div>
										<p class="text-sm font-medium">Task Updates</p>
										<p class="text-xs text-muted-foreground">
											Assignments, due dates, and completions
										</p>
									</div>
									<input
										type="checkbox"
										bind:checked={notificationPrefs.push.tasks}
										class="mt-0.5 size-4 shrink-0 rounded border-input text-primary focus:ring-primary"
									/>
								</label>
								<label
									class="-mx-3 flex cursor-pointer items-start justify-between gap-4 rounded-lg p-3 transition-colors hover:bg-muted/50"
								>
									<div>
										<p class="text-sm font-medium">Comments</p>
										<p class="text-xs text-muted-foreground">New comments on projects and tasks</p>
									</div>
									<input
										type="checkbox"
										bind:checked={notificationPrefs.push.comments}
										class="mt-0.5 size-4 shrink-0 rounded border-input text-primary focus:ring-primary"
									/>
								</label>
								<label
									class="-mx-3 flex cursor-pointer items-start justify-between gap-4 rounded-lg p-3 transition-colors hover:bg-muted/50"
								>
									<div>
										<p class="text-sm font-medium">Reminders</p>
										<p class="text-xs text-muted-foreground">
											Upcoming deadlines and scheduled events
										</p>
									</div>
									<input
										type="checkbox"
										bind:checked={notificationPrefs.push.reminders}
										class="mt-0.5 size-4 shrink-0 rounded border-input text-primary focus:ring-primary"
									/>
								</label>
							</div>
							<div class="mt-6 border-t border-border pt-4">
								<Button
									onclick={() => saveSection('push-notifications')}
									disabled={saving === 'push-notifications'}
								>
									{#if saving === 'push-notifications'}
										<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
									{/if}
									Save push preferences
								</Button>
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Email Digest Frequency</h3>
							<div class="grid gap-3 sm:grid-cols-3">
								{#each [{ value: 'realtime', label: 'Real-time', icon: 'lucide:zap', desc: 'Send immediately' }, { value: 'daily', label: 'Daily Digest', icon: 'lucide:sun', desc: 'Once per day at 9am' }, { value: 'weekly', label: 'Weekly', icon: 'lucide:calendar', desc: 'Every Monday' }] as option}
									<label class="cursor-pointer">
										<input
											type="radio"
											name="digest"
											value={option.value}
											bind:group={notificationPrefs.email.digest}
											class="peer sr-only"
										/>
										<div
											class="h-full rounded-lg border-2 border-border p-4 transition-all peer-checked:border-primary peer-checked:bg-primary/5"
										>
											<Icon
												icon={option.icon}
												class="mb-2 size-5 text-muted-foreground peer-checked:text-primary"
											/>
											<p class="text-sm font-medium">{option.label}</p>
											<p class="text-xs text-muted-foreground">{option.desc}</p>
										</div>
									</label>
								{/each}
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<div class="mb-4 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
								<div>
									<h3 class="font-semibold">Quiet Hours</h3>
									<p class="text-sm text-muted-foreground">
										Pause notifications during specific hours.
									</p>
								</div>
								<label class="relative inline-flex shrink-0 cursor-pointer items-center">
									<input type="checkbox" bind:checked={quietHours.enabled} class="peer sr-only" />
									<div
										class="h-6 w-11 rounded-full bg-muted transition-colors peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-transform after:content-[''] peer-checked:after:translate-x-5"
									></div>
								</label>
							</div>
							{#if quietHours.enabled}
								<div class="grid gap-4 border-t border-border pt-4 sm:grid-cols-2">
									<div class="space-y-2">
										<label for="quietStart" class="text-sm font-medium">Start time</label>
										<input
											id="quietStart"
											type="time"
											bind:value={quietHours.start}
											class="h-10 w-full rounded-md border border-input bg-background px-3 text-sm focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none"
										/>
									</div>
									<div class="space-y-2">
										<label for="quietEnd" class="text-sm font-medium">End time</label>
										<input
											id="quietEnd"
											type="time"
											bind:value={quietHours.end}
											class="h-10 w-full rounded-md border border-input bg-background px-3 text-sm focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none"
										/>
									</div>
								</div>
								<div class="mt-4">
									<Button
										size="sm"
										onclick={() => saveSection('quiet-hours')}
										disabled={saving === 'quiet-hours'}
									>
										{#if saving === 'quiet-hours'}
											<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
										{/if}
										Save quiet hours
									</Button>
								</div>
							{/if}
						</div>
					</div>
				{:else if activeTab === 'appearance'}
					<div class="space-y-8">
						<div>
							<h2 class="text-2xl font-semibold tracking-tight">Appearance</h2>
							<p class="text-muted-foreground">Customize how Artemis looks and feels for you.</p>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Theme</h3>
							<div class="grid gap-4 sm:grid-cols-3">
								<button
									class="group relative rounded-xl border-2 p-4 transition-all {mode.current ===
									'light'
										? 'border-primary bg-primary/5'
										: 'border-border hover:border-primary/50'}"
									onclick={() => {
										if (mode.current === 'dark') toggleMode();
									}}
								>
									<div class="mb-3 aspect-video overflow-hidden rounded-lg bg-slate-100 shadow-sm">
										<div class="h-full w-full bg-gradient-to-br from-slate-50 to-slate-200 p-3">
											<div class="h-2 w-12 rounded bg-slate-300"></div>
											<div class="mt-2 h-2 w-8 rounded bg-slate-300"></div>
											<div class="mt-4 flex gap-2">
												<div class="h-6 w-6 rounded bg-slate-300"></div>
												<div class="h-6 flex-1 rounded bg-slate-300"></div>
											</div>
										</div>
									</div>
									<div class="flex items-center justify-center gap-2">
										{#if mode.current === 'light'}
											<Icon icon="lucide:check-circle" class="size-4 text-primary" />
										{/if}
										<span class="font-medium">Light</span>
									</div>
								</button>

								<button
									class="group relative rounded-xl border-2 p-4 transition-all {mode.current ===
									'dark'
										? 'border-primary bg-primary/5'
										: 'border-border hover:border-primary/50'}"
									onclick={() => {
										if (mode.current !== 'dark') toggleMode();
									}}
								>
									<div class="mb-3 aspect-video overflow-hidden rounded-lg bg-slate-900 shadow-sm">
										<div class="h-full w-full bg-gradient-to-br from-slate-800 to-slate-950 p-3">
											<div class="h-2 w-12 rounded bg-slate-600"></div>
											<div class="mt-2 h-2 w-8 rounded bg-slate-600"></div>
											<div class="mt-4 flex gap-2">
												<div class="h-6 w-6 rounded bg-slate-600"></div>
												<div class="h-6 flex-1 rounded bg-slate-600"></div>
											</div>
										</div>
									</div>
									<div class="flex items-center justify-center gap-2">
										{#if mode.current === 'dark'}
											<Icon icon="lucide:check-circle" class="size-4 text-primary" />
										{/if}
										<span class="font-medium">Dark</span>
									</div>
								</button>

								<button
									class="group relative rounded-xl border-2 border-border p-4 transition-all hover:border-primary/50"
								>
									<div class="mb-3 aspect-video overflow-hidden rounded-lg shadow-sm">
										<div class="grid h-full w-full grid-cols-2">
											<div class="bg-gradient-to-br from-slate-50 to-slate-200 p-3">
												<div class="h-1.5 w-6 rounded bg-slate-300"></div>
											</div>
											<div class="bg-gradient-to-br from-slate-800 to-slate-950 p-3">
												<div class="h-1.5 w-6 rounded bg-slate-600"></div>
											</div>
										</div>
									</div>
									<div class="flex items-center justify-center">
										<span class="font-medium">System</span>
									</div>
								</button>
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Accent Color</h3>
							<div class="flex flex-wrap gap-3">
								{#each accentColors as color}
									<button
										class="group relative flex size-12 items-center justify-center rounded-full {color.class} transition-all hover:scale-110 focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 {color.value ===
										'emerald'
											? 'ring-2 ring-current ring-offset-2'
											: ''}"
										onclick={() => {}}
										aria-label={`Select ${color.name} accent color`}
									>
										{#if color.value === 'emerald'}
											<Icon icon="lucide:check" class="size-5 text-white" />
										{/if}
									</button>
								{/each}
							</div>
							<div class="mt-4">
								<Button
									onclick={() => saveSection('appearance')}
									disabled={saving === 'appearance'}
								>
									{#if saving === 'appearance'}
										<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
									{/if}
									Save appearance
								</Button>
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Language & Region</h3>
							<div class="grid gap-4 sm:grid-cols-2">
								<div class="space-y-2">
									<label for="language" class="text-sm font-medium">Language</label>
									<select
										id="language"
										class="h-10 w-full rounded-md border border-input bg-background px-3 text-sm focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none"
									>
										<option value="en">English (US)</option>
										<option value="en-gb">English (UK)</option>
										<option value="es">Español</option>
										<option value="fr">Français</option>
										<option value="de">Deutsch</option>
									</select>
								</div>
								<div class="space-y-2">
									<label for="timezone" class="text-sm font-medium">Timezone</label>
									<select
										id="timezone"
										class="h-10 w-full rounded-md border border-input bg-background px-3 text-sm focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none"
									>
										<option value="America/Los_Angeles">Pacific Time (PT)</option>
										<option value="America/Denver">Mountain Time (MT)</option>
										<option value="America/Chicago">Central Time (CT)</option>
										<option value="America/New_York">Eastern Time (ET)</option>
										<option value="Europe/London">London (GMT)</option>
									</select>
								</div>
							</div>
							<div class="mt-4">
								<Button onclick={() => saveSection('locale')} disabled={saving === 'locale'}>
									{#if saving === 'locale'}
										<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
									{/if}
									Save preferences
								</Button>
							</div>
						</div>
					</div>
				{:else if activeTab === 'billing'}
					<div class="space-y-8">
						<div>
							<h2 class="text-2xl font-semibold tracking-tight">Billing</h2>
							<p class="text-muted-foreground">
								Manage your subscription, payment methods, and billing history.
							</p>
						</div>

						<div
							class="rounded-xl border border-primary/20 bg-gradient-to-br from-primary/5 to-primary/10 p-6"
						>
							<div class="flex flex-col gap-4 md:flex-row md:items-start md:justify-between">
								<div>
									<div class="flex items-center gap-2">
										<h3 class="text-xl font-bold">Pro Plan</h3>
										<span
											class="rounded-full bg-primary/10 px-2.5 py-0.5 text-xs font-medium text-primary"
											>Active</span
										>
									</div>
									<p class="mt-2 text-3xl font-bold">
										$29<span class="text-base font-normal text-muted-foreground">/month</span>
									</p>
									<p class="mt-1 text-sm text-muted-foreground">
										Next billing date: February 27, 2026
									</p>
								</div>
								<div class="flex gap-2">
									<Button variant="outline">Change plan</Button>
									<Button
										variant="ghost"
										class="text-destructive hover:bg-destructive/10 hover:text-destructive"
										>Cancel</Button
									>
								</div>
							</div>
							<Separator class="my-4 bg-primary/20" />
							<div class="grid gap-2 sm:grid-cols-2">
								{#each ['Unlimited projects', '100GB storage', 'Priority support', 'Advanced analytics', 'Custom domains', 'Team collaboration'] as feature}
									<div class="flex items-center gap-2 text-sm">
										<Icon icon="lucide:check" class="size-4 text-primary" />
										{feature}
									</div>
								{/each}
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<div class="mb-4 flex items-center justify-between">
								<h3 class="font-semibold">Payment Methods</h3>
								<Button size="sm" variant="outline">
									<Icon icon="lucide:plus" class="mr-2 size-4" />
									Add method
								</Button>
							</div>
							<div class="space-y-3">
								{#each paymentMethods as method}
									<div
										class="flex items-center justify-between rounded-lg border border-border p-4"
									>
										<div class="flex items-center gap-4">
											<div class="flex size-10 items-center justify-center rounded-lg bg-muted">
												<Icon icon="lucide:credit-card" class="size-5 text-muted-foreground" />
											</div>
											<div>
												<div class="flex items-center gap-2">
													<span class="font-medium">{method.brand} ending in {method.last4}</span>
													{#if method.isDefault}
														<span class="rounded bg-muted px-2 py-0.5 text-xs">Default</span>
													{/if}
												</div>
												<p class="text-sm text-muted-foreground">
													Expires {method.expMonth.toString().padStart(2, '0')}/{method.expYear}
												</p>
											</div>
										</div>
										<div class="flex items-center gap-1">
											{#if !method.isDefault}
												<Button variant="ghost" size="sm">Set default</Button>
											{/if}
											<Button variant="ghost" size="icon" class="text-destructive">
												<Icon icon="lucide:trash-2" class="size-4" />
											</Button>
										</div>
									</div>
								{/each}
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Billing Information</h3>
							<div class="grid gap-4 sm:grid-cols-2">
								<div class="space-y-2">
									<label for="billingCompany" class="text-sm font-medium">Company name</label>
									<Input id="billingCompany" value={workspace.name} />
								</div>
								<div class="space-y-2">
									<label for="billingEmail" class="text-sm font-medium">Billing email</label>
									<Input id="billingEmail" type="email" value={workspace.billingEmail} />
								</div>
							</div>
							<div class="mt-4 space-y-2">
								<label for="billingAddress" class="text-sm font-medium">Address</label>
								<Input id="billingAddress" placeholder="123 Main St" />
							</div>
							<div class="mt-4 grid gap-4 sm:grid-cols-3">
								<div class="space-y-2">
									<label for="billingCity" class="text-sm font-medium">City</label>
									<Input id="billingCity" placeholder="San Francisco" />
								</div>
								<div class="space-y-2">
									<label for="billingState" class="text-sm font-medium">State</label>
									<Input id="billingState" placeholder="CA" />
								</div>
								<div class="space-y-2">
									<label for="billingZip" class="text-sm font-medium">Postal code</label>
									<Input id="billingZip" placeholder="94102" />
								</div>
							</div>
							<div class="mt-6">
								<Button
									onclick={() => saveSection('billing-info')}
									disabled={saving === 'billing-info'}
								>
									{#if saving === 'billing-info'}
										<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
									{/if}
									Update billing info
								</Button>
							</div>
						</div>

						<div class="overflow-hidden rounded-xl border border-border bg-card">
							<div class="flex items-center justify-between border-b border-border p-4">
								<h3 class="font-semibold">Invoice History</h3>
								<Button variant="ghost" size="sm">
									<Icon icon="lucide:download" class="mr-2 size-4" />
									Download all
								</Button>
							</div>
							<div class="overflow-x-auto">
								<table class="w-full text-sm">
									<thead class="border-b border-border bg-muted/50">
										<tr>
											<th class="px-4 py-3 text-left font-medium">Invoice</th>
											<th class="px-4 py-3 text-left font-medium">Date</th>
											<th class="px-4 py-3 text-left font-medium">Amount</th>
											<th class="px-4 py-3 text-left font-medium">Status</th>
											<th class="px-4 py-3 text-right"></th>
										</tr>
									</thead>
									<tbody class="divide-y divide-border">
										{#each invoices as invoice}
											<tr>
												<td class="px-4 py-3 font-medium">{invoice.id}</td>
												<td class="px-4 py-3 text-muted-foreground">{invoice.date}</td>
												<td class="px-4 py-3">${invoice.amount.toFixed(2)}</td>
												<td class="px-4 py-3">
													<span
														class="inline-flex items-center gap-1 text-xs font-medium text-emerald-600"
													>
														<Icon icon="lucide:check-circle" class="size-3" />
														{invoice.status}
													</span>
												</td>
												<td class="px-4 py-3 text-right">
													<Button
														variant="ghost"
														size="icon"
														onclick={() => downloadInvoice(invoice.id)}
													>
														<Icon icon="lucide:download" class="size-4" />
													</Button>
												</td>
											</tr>
										{/each}
									</tbody>
								</table>
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Usage</h3>
							<div class="grid gap-4 sm:grid-cols-2">
								<div class="space-y-2">
									<div class="flex items-center justify-between text-sm">
										<span class="font-medium">Storage</span>
										<span class="text-muted-foreground">45GB / 100GB</span>
									</div>
									<div class="h-2 w-full overflow-hidden rounded-full bg-muted">
										<div class="h-full w-[45%] rounded-full bg-primary"></div>
									</div>
								</div>
								<div class="space-y-2">
									<div class="flex items-center justify-between text-sm">
										<span class="font-medium">Team Members</span>
										<span class="text-muted-foreground"
											>{workspace.usedSeats} / {workspace.seats}</span
										>
									</div>
									<div class="h-2 w-full overflow-hidden rounded-full bg-muted">
										<div class="h-full w-[60%] rounded-full bg-primary"></div>
									</div>
								</div>
							</div>
						</div>
					</div>
				{:else if activeTab === 'workspace'}
					<div class="space-y-8">
						<div>
							<h2 class="text-2xl font-semibold tracking-tight">Workspace</h2>
							<p class="text-muted-foreground">
								Manage your team, organization settings, and access controls.
							</p>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Workspace Details</h3>
							<div class="space-y-2">
								<label for="workspaceName" class="text-sm font-medium">Workspace name</label>
								<Input id="workspaceName" value={workspace.name} />
							</div>
							<div class="mt-4 space-y-2">
								<label for="workspaceSlug" class="text-sm font-medium">Workspace URL</label>
								<div class="relative">
									<span
										class="absolute top-1/2 left-3 -translate-y-1/2 text-sm text-muted-foreground"
										>artemis.app/</span
									>
									<Input id="workspaceSlug" value={workspace.slug} class="pl-24" />
								</div>
								<p class="text-xs text-muted-foreground">
									This is your workspace's unique URL on Artemis.
								</p>
							</div>
							<div class="mt-6">
								<Button
									onclick={() => saveSection('workspace-details')}
									disabled={saving === 'workspace-details'}
								>
									{#if saving === 'workspace-details'}
										<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
									{/if}
									Save changes
								</Button>
							</div>
						</div>

						<div class="overflow-hidden rounded-xl border border-border bg-card">
							<div
								class="flex flex-col gap-4 border-b border-border p-4 sm:flex-row sm:items-center sm:justify-between"
							>
								<div>
									<h3 class="font-semibold">Team Members</h3>
									<p class="text-sm text-muted-foreground">
										{workspace.usedSeats} of {workspace.seats} seats used
									</p>
								</div>
								<Button onclick={() => (showInviteModal = true)}>
									<Icon icon="lucide:user-plus" class="mr-2 size-4" />
									Invite member
								</Button>
							</div>
							<div class="divide-y divide-border">
								{#each teamMembers as member}
									<div
										class="flex flex-col gap-4 p-4 sm:flex-row sm:items-center sm:justify-between"
									>
										<div class="flex items-center gap-4">
											<img
												src={member.avatar}
												alt={member.name}
												class="size-10 rounded-full bg-muted"
											/>
											<div>
												<p class="font-medium">{member.name}</p>
												<p class="text-sm text-muted-foreground">{member.email}</p>
											</div>
										</div>
										<div class="flex items-center gap-2">
											<span class="rounded-full px-2 py-1 text-xs {getRoleBadgeColor(member.role)}">
												{member.role.charAt(0).toUpperCase() + member.role.slice(1)}
											</span>
											<Button variant="ghost" size="icon">
												<Icon icon="lucide:more-vertical" class="size-4" />
											</Button>
										</div>
									</div>
								{/each}
							</div>
						</div>

						<div class="rounded-xl border border-border bg-card p-6">
							<h3 class="mb-4 font-semibold">Invite Link</h3>
							<p class="mb-3 text-sm text-muted-foreground">
								Share this link to invite people directly to your workspace.
							</p>
							<div class="flex gap-2">
								<Input
									value="https://artemis.app/invite/{workspace.slug}/abc123xyz"
									readonly
									class="bg-muted font-mono text-sm"
								/>
								<Button variant="outline">
									<Icon icon="lucide:copy" class="size-4" />
								</Button>
								<Button variant="outline">
									<Icon icon="lucide:refresh-cw" class="size-4" />
								</Button>
							</div>
						</div>

						<div class="rounded-xl border border-destructive/20 bg-destructive/5 p-6">
							<div class="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
								<div>
									<h3 class="font-semibold text-destructive">Danger Zone</h3>
									<p class="text-sm text-muted-foreground">
										Once you delete a workspace, there is no going back. Please be certain.
									</p>
								</div>
								<Button variant="destructive">Delete workspace</Button>
							</div>
						</div>
					</div>
				{/if}
			</div>
		</main>
	</div>
</div>
