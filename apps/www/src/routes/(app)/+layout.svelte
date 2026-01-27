<script lang="ts">
	import { page } from '$app/state';
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { toggleMode, mode } from 'mode-watcher';

	let { children } = $props();

	let sidebarCollapsed = $state(false);
	let mobileMenuOpen = $state(false);

	const mainNavigation = [
		{ href: '/', label: 'Dashboard', icon: 'lucide:layout-dashboard' },
		{ href: '/clients', label: 'Clients', icon: 'lucide:users' },
		{ href: '/projects', label: 'Projects', icon: 'lucide:folder-kanban' },
		{ href: '/tasks', label: 'Tasks', icon: 'lucide:check-square' },
		{ href: '/time', label: 'Time Tracking', icon: 'lucide:clock' }
	];

	const financialNavigation = [
		{ href: '/invoices', label: 'Invoices', icon: 'lucide:file-text' },
		{ href: '/expenses', label: 'Expenses', icon: 'lucide:receipt' },
		{ href: '/reports', label: 'Reports', icon: 'lucide:bar-chart-3' }
	];

	const resourceNavigation = [
		{ href: '/templates', label: 'Templates', icon: 'lucide:layout-template' },
		{ href: '/files', label: 'Files', icon: 'lucide:folder-open' },
		{ href: '/emails', label: 'Emails', icon: 'lucide:mail' }
	];

	const settingsNavigation = [
		{ href: '/settings', label: 'Settings', icon: 'lucide:settings' },
		{ href: '/help', label: 'Help & Support', icon: 'lucide:help-circle' }
	];

	const isActive = (href: string) => {
		if (href === '/') {
			return page.url.pathname === '/';
		}
		return page.url.pathname.startsWith(href);
	};

	const user = {
		name: 'Alex Morgan',
		email: 'alex@example.com',
		avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex',
		workspace: {
			name: 'Design Studio',
			plan: 'Pro'
		}
	};

	const notifications = [
		{
			id: 1,
			title: 'New invoice paid',
			message: 'Acme Corp paid $2,500',
			time: '2 min ago',
			read: false
		},
		{
			id: 2,
			title: 'Task due soon',
			message: 'Website redesign - Homepage',
			time: '1 hour ago',
			read: false
		},
		{
			id: 3,
			title: 'Client message',
			message: 'Sarah from TechStart commented',
			time: '3 hours ago',
			read: true
		}
	];

	let showNotifications = $state(false);
	let showUserMenu = $state(false);

	function toggleSidebar() {
		sidebarCollapsed = !sidebarCollapsed;
	}
</script>

<div class="min-h-screen bg-background">
	<header
		class="sticky top-0 z-50 flex h-14 items-center justify-between border-b border-border bg-background px-4 lg:hidden"
	>
		<div class="flex items-center gap-3">
			<Button variant="ghost" size="icon" onclick={() => (mobileMenuOpen = !mobileMenuOpen)}>
				<Icon icon="lucide:menu" class="size-5" />
			</Button>
			<a href="/" class="flex items-center gap-2">
				<div
					class="flex size-8 items-center justify-center rounded-md bg-primary text-primary-foreground"
				>
					<Icon icon="hugeicons:cowboy-hat" class="size-5" />
				</div>
				<span class="font-semibold">Artemis</span>
			</a>
		</div>
		<div class="flex items-center gap-2">
			<Button variant="ghost" size="icon" class="relative">
				<Icon icon="lucide:bell" class="size-5" />
				<span class="absolute top-1.5 right-1.5 size-2 rounded-full bg-destructive"></span>
			</Button>
			<img src={user.avatar} alt={user.name} class="size-8 rounded-full bg-muted" />
		</div>
	</header>

	{#if mobileMenuOpen}
		<div
			class="fixed inset-0 z-40 bg-background/80 backdrop-blur-sm lg:hidden"
			onclick={() => (mobileMenuOpen = false)}
			role="button"
			tabindex="0"
		></div>
		<div class="fixed inset-y-0 left-0 z-50 w-80 bg-background shadow-xl lg:hidden">
			<div class="flex h-full flex-col p-4">
				<div class="mb-6 flex items-center justify-between">
					<span class="text-lg font-semibold">Menu</span>
					<Button variant="ghost" size="icon" onclick={() => (mobileMenuOpen = false)}>
						<Icon icon="lucide:x" class="size-5" />
					</Button>
				</div>
				<nav class="flex-1 space-y-6 overflow-auto">
					<div class="space-y-1">
						{#each mainNavigation as item}
							<Button
								variant={isActive(item.href) ? 'secondary' : 'ghost'}
								class="w-full justify-start gap-3"
								href={item.href}
								onclick={() => (mobileMenuOpen = false)}
							>
								<Icon icon={item.icon} class="size-5" />
								{item.label}
							</Button>
						{/each}
					</div>
					<Separator />
					<div class="space-y-1">
						{#each [...financialNavigation, ...resourceNavigation] as item}
							<Button
								variant={isActive(item.href) ? 'secondary' : 'ghost'}
								class="w-full justify-start gap-3"
								href={item.href}
								onclick={() => (mobileMenuOpen = false)}
							>
								<Icon icon={item.icon} class="size-5" />
								{item.label}
							</Button>
						{/each}
					</div>
				</nav>
			</div>
		</div>
	{/if}

	<div class="flex">
		<aside
			class="{sidebarCollapsed
				? 'w-16'
				: 'w-64'} sticky top-0 hidden h-screen flex-col border-r border-border bg-muted/30 transition-all duration-300 lg:flex"
		>
			<div
				class="flex h-14 items-center {sidebarCollapsed
					? 'justify-center px-2'
					: 'px-4'} gap-3 border-b border-border"
			>
				<a href="/" class="flex items-center gap-2">
					<div
						class="flex size-8 shrink-0 items-center justify-center rounded-md bg-primary text-primary-foreground"
					>
						<Icon icon="hugeicons:cowboy-hat" class="size-5" />
					</div>
					{#if !sidebarCollapsed}
						<span class="font-semibold">Artemis</span>
					{/if}
				</a>
			</div>

			{#if !sidebarCollapsed}
				<div class="border-b border-border p-3">
					<button
						class="flex w-full items-center justify-between rounded-md border border-border bg-background px-3 py-2 text-sm hover:bg-muted"
					>
						<div class="flex items-center gap-2">
							<div
								class="flex size-6 items-center justify-center rounded bg-primary/10 text-xs font-medium text-primary"
							>
								{user.workspace.name.charAt(0)}
							</div>
							<div class="text-left">
								<div class="font-medium">{user.workspace.name}</div>
								<div class="text-xs text-muted-foreground">{user.workspace.plan} Plan</div>
							</div>
						</div>
						<Icon icon="lucide:chevrons-up-down" class="size-4 text-muted-foreground" />
					</button>
				</div>
			{/if}

			<nav class="flex-1 overflow-y-auto py-4">
				<div class="space-y-1 {sidebarCollapsed ? 'px-2' : 'px-3'}">
					{#each mainNavigation as item}
						<Button
							variant={isActive(item.href) ? 'secondary' : 'ghost'}
							class={sidebarCollapsed ? 'w-full justify-center px-2' : 'w-full justify-start gap-3'}
							href={item.href}
							title={sidebarCollapsed ? item.label : undefined}
						>
							<Icon icon={item.icon} class="size-5 shrink-0" />
							{#if !sidebarCollapsed}
								<span>{item.label}</span>
							{/if}
						</Button>
					{/each}
				</div>

				{#if !sidebarCollapsed}
					<div class="mt-6 px-3">
						<div class="mb-2 px-2 text-xs font-medium text-muted-foreground">Financial</div>
					</div>
				{:else}
					<div class="mt-6 flex justify-center">
						<Separator class="w-8" />
					</div>
				{/if}
				<div class="space-y-1 {sidebarCollapsed ? 'px-2' : 'px-3'}">
					{#each financialNavigation as item}
						<Button
							variant={isActive(item.href) ? 'secondary' : 'ghost'}
							class={sidebarCollapsed ? 'w-full justify-center px-2' : 'w-full justify-start gap-3'}
							href={item.href}
							title={sidebarCollapsed ? item.label : undefined}
						>
							<Icon icon={item.icon} class="size-5 shrink-0" />
							{#if !sidebarCollapsed}
								<span>{item.label}</span>
							{/if}
						</Button>
					{/each}
				</div>

				{#if !sidebarCollapsed}
					<div class="mt-6 px-3">
						<div class="mb-2 px-2 text-xs font-medium text-muted-foreground">Resources</div>
					</div>
				{:else}
					<div class="mt-6 flex justify-center">
						<Separator class="w-8" />
					</div>
				{/if}
				<div class="space-y-1 {sidebarCollapsed ? 'px-2' : 'px-3'}">
					{#each resourceNavigation as item}
						<Button
							variant={isActive(item.href) ? 'secondary' : 'ghost'}
							class={sidebarCollapsed ? 'w-full justify-center px-2' : 'w-full justify-start gap-3'}
							href={item.href}
							title={sidebarCollapsed ? item.label : undefined}
						>
							<Icon icon={item.icon} class="size-5 shrink-0" />
							{#if !sidebarCollapsed}
								<span>{item.label}</span>
							{/if}
						</Button>
					{/each}
				</div>
			</nav>

			<div class="border-t border-border p-3">
				<div class="space-y-1">
					{#each settingsNavigation as item}
						<Button
							variant={isActive(item.href) ? 'secondary' : 'ghost'}
							class={sidebarCollapsed ? 'w-full justify-center px-2' : 'w-full justify-start gap-3'}
							href={item.href}
							title={sidebarCollapsed ? item.label : undefined}
						>
							<Icon icon={item.icon} class="size-5 shrink-0" />
							{#if !sidebarCollapsed}
								<span>{item.label}</span>
							{/if}
						</Button>
					{/each}
				</div>

				{#if !sidebarCollapsed}
					<Separator class="my-3" />
					<button
						class="flex w-full items-center gap-3 rounded-md p-2 hover:bg-muted"
						onclick={() => (showUserMenu = !showUserMenu)}
					>
						<img src={user.avatar} alt={user.name} class="size-8 rounded-full bg-muted" />
						<div class="flex-1 text-left">
							<div class="text-sm font-medium">{user.name}</div>
							<div class="text-xs text-muted-foreground">{user.email}</div>
						</div>
						<Icon icon="lucide:more-vertical" class="size-4 text-muted-foreground" />
					</button>
				{:else}
					<Separator class="my-3" />
					<Button variant="ghost" size="icon" class="w-full">
						<img src={user.avatar} alt={user.name} class="size-8 rounded-full bg-muted" />
					</Button>
				{/if}
			</div>
		</aside>

		<div class="flex min-h-screen flex-1 flex-col">
			<header
				class="sticky top-0 z-30 hidden h-14 items-center justify-between border-b border-border bg-background/95 px-6 backdrop-blur supports-[backdrop-filter]:bg-background/60 lg:flex"
			>
				<div class="flex items-center gap-4">
					<Button variant="ghost" size="icon" onclick={toggleSidebar}>
						<Icon
							icon={sidebarCollapsed ? 'lucide:panel-left-open' : 'lucide:panel-left-close'}
							class="size-5"
						/>
					</Button>
					<span class="text-sm text-muted-foreground"
						>{page.url.pathname === '/' ? 'Dashboard' : ''}</span
					>
				</div>

				<div class="flex items-center gap-2">
					<Button variant="outline" class="h-8 gap-2 text-muted-foreground" onclick={() => {}}>
						<Icon icon="lucide:search" class="size-4" />
						<span class="text-sm">Search...</span>
						<kbd class="hidden rounded border bg-muted px-1.5 text-xs font-medium sm:inline-block"
							>⌘K</kbd
						>
					</Button>

					<Separator orientation="vertical" class="mx-1 h-6" />

					<Button size="sm" class="gap-2">
						<Icon icon="lucide:plus" class="size-4" />
						<span class="hidden sm:inline">Create</span>
					</Button>

					<Separator orientation="vertical" class="mx-1 h-6" />

					<Button variant="ghost" size="icon" onclick={toggleMode}>
						{#if mode.current === 'dark'}
							<Icon icon="lucide:moon" class="size-5" />
						{:else}
							<Icon icon="lucide:sun" class="size-5" />
						{/if}
					</Button>

					<div class="relative">
						<Button
							variant="ghost"
							size="icon"
							class="relative"
							onclick={() => (showNotifications = !showNotifications)}
						>
							<Icon icon="lucide:bell" class="size-5" />
							<span class="absolute top-1.5 right-1.5 size-2 rounded-full bg-destructive"></span>
						</Button>

						{#if showNotifications}
							<div
								class="absolute top-full right-0 mt-2 w-80 rounded-lg border border-border bg-popover p-2 shadow-lg"
							>
								<div class="flex items-center justify-between border-b border-border px-2 pb-2">
									<span class="font-medium">Notifications</span>
									<Button variant="ghost" size="sm" class="h-auto px-2 py-1 text-xs">
										Mark all read
									</Button>
								</div>
								<div class="max-h-80 overflow-y-auto py-1">
									{#each notifications as notification}
										<button
											class="w-full rounded-md px-2 py-2 text-left hover:bg-muted {notification.read
												? ''
												: 'bg-muted/50'}"
										>
											<div class="flex items-start gap-2">
												{#if !notification.read}
													<span class="mt-1.5 size-2 shrink-0 rounded-full bg-primary"></span>
												{:else}
													<span class="mt-1.5 size-2 shrink-0 rounded-full bg-transparent"></span>
												{/if}
												<div class="flex-1">
													<div class="text-sm font-medium">{notification.title}</div>
													<div class="text-xs text-muted-foreground">{notification.message}</div>
													<div class="mt-1 text-xs text-muted-foreground">{notification.time}</div>
												</div>
											</div>
										</button>
									{/each}
								</div>
								<div class="border-t border-border px-2 pt-2">
									<Button variant="ghost" size="sm" class="w-full" href="/notifications">
										View all notifications
									</Button>
								</div>
							</div>
						{/if}
					</div>
				</div>
			</header>

			<main class="flex-1">
				{@render children()}
			</main>
		</div>
	</div>
</div>
