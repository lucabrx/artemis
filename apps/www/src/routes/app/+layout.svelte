<script lang="ts">
	import { page } from '$app/state';
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';
	import { toggleMode, mode } from 'mode-watcher';
	import { CommandPalette } from '$lib/components/ui/command-palette/index.js';

	let { children } = $props();

	let sidebarCollapsed = $state(false);
	let mobileMenuOpen = $state(false);
	let commandPaletteOpen = $state(false);

	function handleKeydown(e: KeyboardEvent) {
		if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
			e.preventDefault();
			commandPaletteOpen = true;
		}
	}

	const mainNavigation = [
		{ href: '/app', label: 'Dashboard', icon: 'lucide:layout-dashboard' },
		{ href: '/app/clients', label: 'Clients', icon: 'lucide:users' },
		{ href: '/app/projects', label: 'Projects', icon: 'lucide:folder-kanban' },
		{ href: '/app/tasks', label: 'Tasks', icon: 'lucide:check-square' },
		{ href: '/app/time', label: 'Time', icon: 'lucide:clock' }
	];

	const financialNavigation = [
		{ href: '/app/invoices', label: 'Invoices', icon: 'lucide:file-text' },
		{ href: '/app/expenses', label: 'Expenses', icon: 'lucide:receipt' },
		{ href: '/app/reports', label: 'Reports', icon: 'lucide:bar-chart-3' }
	];

	const resourceNavigation = [
		{ href: '/app/templates', label: 'Templates', icon: 'lucide:layout-template' },
		{ href: '/app/files', label: 'Files', icon: 'lucide:folder-open' },
		{ href: '/app/emails', label: 'Emails', icon: 'lucide:mail', badge: 3 }
	];

	const settingsNavigation = [
		{ href: '/app/settings', label: 'Settings', icon: 'lucide:settings' },
		{ href: '/app/help', label: 'Help', icon: 'lucide:help-circle' }
	];

	const isActive = (href: string) => {
		if (href === '/app') {
			return page.url.pathname === '/app' || page.url.pathname === '/app/';
		}
		return page.url.pathname.startsWith(href);
	};

	const pageTitle = $derived(() => {
		const path = page.url.pathname;
		if (path === '/app' || path === '/app/')
			return { title: 'Dashboard', subtitle: "Here's what's happening today" };
		if (path.startsWith('/app/clients'))
			return { title: 'Clients', subtitle: 'Manage your relationships' };
		if (path.startsWith('/app/projects'))
			return { title: 'Projects', subtitle: 'Track progress and collaborate' };
		if (path.startsWith('/app/tasks'))
			return { title: 'Tasks', subtitle: 'Stay organized and productive' };
		if (path.startsWith('/app/time'))
			return { title: 'Time Tracking', subtitle: 'Log hours and manage entries' };
		if (path.startsWith('/app/invoices'))
			return { title: 'Invoices', subtitle: 'Create and manage billing' };
		if (path.startsWith('/app/expenses'))
			return { title: 'Expenses', subtitle: 'Track business spending' };
		if (path.startsWith('/app/reports'))
			return { title: 'Reports', subtitle: 'Insights and analytics' };
		if (path.startsWith('/app/templates'))
			return { title: 'Templates', subtitle: 'Reusable content' };
		if (path.startsWith('/app/files')) return { title: 'Files', subtitle: 'Documents and assets' };
		if (path.startsWith('/app/emails')) return { title: 'Emails', subtitle: 'Inbox and templates' };
		if (path.startsWith('/app/settings'))
			return { title: 'Settings', subtitle: 'Configure preferences' };
		return { title: 'Artemis', subtitle: '' };
	});

	const user = {
		name: 'Alex Morgan',
		email: 'alex@designstudio.com',
		avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex',
		workspace: {
			name: 'Design Studio',
			plan: 'Pro',
			color: 'bg-emerald-500'
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
	let showWorkspaceMenu = $state(false);
	let showUpgradeModal = $state(false);
	let hoveredNavItem = $state<string | null>(null);

	function toggleSidebar() {
		sidebarCollapsed = !sidebarCollapsed;
	}

	function handleWindowClick() {
		showWorkspaceMenu = false;
		showNotifications = false;
		showUserMenu = false;
	}
</script>

<svelte:head>
	<title>Artemis</title>
</svelte:head>

<svelte:window onkeydown={handleKeydown} onclick={handleWindowClick} />

<CommandPalette bind:open={commandPaletteOpen} />

<div class="min-h-screen bg-background">
	<header
		class="sticky top-0 z-50 flex h-14 items-center justify-between border-b border-border bg-background/95 px-4 backdrop-blur supports-[backdrop-filter]:bg-background/60 lg:hidden"
	>
		<div class="flex items-center gap-3">
			<Button variant="ghost" size="icon" onclick={() => (mobileMenuOpen = !mobileMenuOpen)}>
				<Icon icon="lucide:menu" class="size-5" />
			</Button>
			<a href="/" class="flex items-center gap-2">
				<div
					class="flex size-8 items-center justify-center rounded-lg bg-primary text-primary-foreground"
				>
					<Icon icon="lucide:hexagon" class="size-5" />
				</div>
				<span class="font-semibold">Artemis</span>
			</a>
		</div>
		<div class="flex items-center gap-2">
			<Button variant="ghost" size="icon" class="relative">
				<Icon icon="lucide:bell" class="size-5" />
				<span
					class="absolute top-1.5 right-1.5 size-2 rounded-full bg-destructive ring-2 ring-background"
				></span>
			</Button>
			<img
				src={user.avatar}
				alt={user.name}
				class="size-8 rounded-full bg-muted ring-2 ring-border"
			/>
		</div>
	</header>

	{#if mobileMenuOpen}
		<div
			class="fixed inset-0 z-40 bg-background/80 backdrop-blur-sm lg:hidden"
			onclick={() => (mobileMenuOpen = false)}
			role="button"
			tabindex="0"
		></div>
		<div class="fixed inset-y-0 left-0 z-50 w-72 bg-background shadow-2xl lg:hidden">
			<div class="flex h-full flex-col">
				<div class="flex h-14 items-center justify-between border-b border-border px-4">
					<span class="font-semibold">Menu</span>
					<Button variant="ghost" size="icon" onclick={() => (mobileMenuOpen = false)}>
						<Icon icon="lucide:x" class="size-5" />
					</Button>
				</div>
				<nav class="flex-1 space-y-1 overflow-auto p-3">
					{#each [...mainNavigation, ...financialNavigation, ...resourceNavigation, ...settingsNavigation] as item}
						<a
							href={item.href}
							class="flex items-center gap-3 rounded-lg px-3 py-2.5 text-sm transition-colors {isActive(
								item.href
							)
								? 'bg-primary/10 font-medium text-primary'
								: 'text-muted-foreground hover:bg-muted hover:text-foreground'}"
							onclick={() => (mobileMenuOpen = false)}
						>
							<Icon icon={item.icon} class="size-5" />
							{item.label}
						</a>
					{/each}
				</nav>
			</div>
		</div>
	{/if}

	<div class="flex">
		<aside
			class="{sidebarCollapsed
				? 'w-[72px]'
				: 'w-64'} sticky top-0 hidden h-screen flex-col border-r border-border bg-sidebar transition-all duration-300 ease-out lg:flex"
		>
			<div
				class="flex h-16 items-center {sidebarCollapsed
					? 'justify-center px-3'
					: 'px-4'} gap-3 border-b border-sidebar-border"
			>
				<a href="/" class="flex items-center gap-3">
					<div
						class="flex size-9 items-center justify-center rounded-xl bg-primary text-primary-foreground shadow-lg shadow-primary/20"
					>
						<Icon icon="lucide:hexagon" class="size-5" />
					</div>
					{#if !sidebarCollapsed}
						<span class="text-lg font-semibold tracking-tight">Artemis</span>
					{/if}
				</a>
			</div>

			{#if !sidebarCollapsed}
				<div class="border-b border-sidebar-border p-3">
					<button
						class="group flex w-full items-center gap-3 rounded-xl border border-sidebar-border bg-sidebar-accent/50 px-3 py-2.5 text-left transition-all hover:border-sidebar-border/80 hover:bg-sidebar-accent"
						onclick={(e) => {
							e.stopPropagation();
							showWorkspaceMenu = !showWorkspaceMenu;
						}}
					>
						<div
							class="flex size-8 items-center justify-center rounded-lg {user.workspace
								.color} text-sm font-semibold text-white shadow-sm"
						>
							{user.workspace.name.charAt(0)}
						</div>
						<div class="min-w-0 flex-1">
							<div class="truncate text-sm font-medium text-sidebar-foreground">
								{user.workspace.name}
							</div>
							<div class="text-xs text-muted-foreground">{user.workspace.plan}</div>
						</div>
						<Icon
							icon="lucide:chevrons-up-down"
							class="size-4 text-muted-foreground opacity-60 transition-opacity group-hover:opacity-100"
						/>
					</button>

					{#if showWorkspaceMenu}
						<div
							class="absolute top-[88px] left-4 z-50 w-64 rounded-xl border border-border bg-popover p-2 shadow-2xl"
							onclick={(e) => e.stopPropagation()}
						>
							<div class="mb-1 flex items-center gap-3 rounded-lg bg-muted/50 p-2">
								<div
									class="flex size-9 items-center justify-center rounded-lg {user.workspace
										.color} text-sm font-semibold text-white"
								>
									{user.workspace.name.charAt(0)}
								</div>
								<div class="min-w-0 flex-1">
									<div class="truncate text-sm font-medium">{user.workspace.name}</div>
									<div class="text-xs text-muted-foreground">{user.workspace.plan} Plan</div>
								</div>
								<div class="size-2 rounded-full bg-emerald-500 ring-2 ring-emerald-500/20"></div>
							</div>

							<button
								class="flex w-full items-center gap-3 rounded-lg px-3 py-2 text-sm transition-colors hover:bg-muted"
							>
								<Icon icon="lucide:plus" class="size-4 text-muted-foreground" />
								Create workspace
							</button>
							<button
								class="flex w-full items-center gap-3 rounded-lg px-3 py-2 text-sm transition-colors hover:bg-muted"
							>
								<Icon icon="lucide:settings" class="size-4 text-muted-foreground" />
								Workspace settings
							</button>
							<button
								class="flex w-full items-center gap-3 rounded-lg px-3 py-2 text-sm transition-colors hover:bg-muted"
							>
								<Icon icon="lucide:user-plus" class="size-4 text-muted-foreground" />
								Invite members
							</button>

							<Separator class="my-2" />

							<button
								class="flex w-full items-center justify-center gap-2 rounded-lg bg-primary px-3 py-2 text-sm font-medium text-primary-foreground transition-colors hover:bg-primary/90"
								onclick={() => {
									showWorkspaceMenu = false;
									showUpgradeModal = true;
								}}
							>
								<Icon icon="lucide:sparkles" class="size-4" />
								Upgrade to Team
							</button>
						</div>
					{/if}
				</div>
			{:else}
				<div class="flex justify-center border-b border-sidebar-border py-3">
					<button
						class="flex size-9 items-center justify-center rounded-xl {user.workspace
							.color} text-sm font-semibold text-white shadow-sm transition-transform hover:scale-105"
						onclick={(e) => {
							e.stopPropagation();
							showWorkspaceMenu = !showWorkspaceMenu;
						}}
					>
						{user.workspace.name.charAt(0)}
					</button>
				</div>
			{/if}

			<nav class="flex-1 overflow-y-auto py-3">
				<div class="space-y-0.5 {sidebarCollapsed ? 'px-2' : 'px-3'}">
					{#each mainNavigation as item}
						{@const active = isActive(item.href)}
						<a
							href={item.href}
							class="group relative flex items-center {sidebarCollapsed
								? 'justify-center'
								: 'gap-3'} rounded-xl px-3 py-2.5 text-sm transition-all duration-200 {active
								? 'bg-white/[0.08] font-medium text-sidebar-foreground'
								: 'text-sidebar-foreground/60 hover:bg-white/[0.04] hover:text-sidebar-foreground/90'}"
							title={sidebarCollapsed ? item.label : undefined}
							onmouseenter={() => (hoveredNavItem = item.label)}
							onmouseleave={() => (hoveredNavItem = null)}
						>
							{#if active}
								<div
									class="absolute top-1/2 left-0 h-4 w-[3px] -translate-y-1/2 rounded-r-full bg-sidebar-foreground"
								></div>
							{/if}
							<Icon
								icon={item.icon}
								class="size-5 shrink-0 {active
									? 'text-sidebar-foreground'
									: 'text-muted-foreground group-hover:text-sidebar-foreground/90'} transition-colors"
							/>
							{#if !sidebarCollapsed}
								<span class="truncate">{item.label}</span>
							{/if}
						</a>
					{/each}
				</div>

				{#if !sidebarCollapsed}
					<div class="mt-6 mb-2 px-6">
						<div class="flex items-center gap-2">
							<div class="h-px flex-1 bg-sidebar-border/60"></div>
							<span
								class="text-[10px] font-semibold tracking-wider text-muted-foreground/60 uppercase"
								>Finance</span
							>
							<div class="h-px flex-1 bg-sidebar-border/60"></div>
						</div>
					</div>
				{:else}
					<div class="mt-6 flex justify-center">
						<div class="h-px w-8 bg-sidebar-border/60"></div>
					</div>
				{/if}
				<div class="space-y-0.5 {sidebarCollapsed ? 'px-2' : 'px-3'}">
					{#each financialNavigation as item}
						{@const active = isActive(item.href)}
						<a
							href={item.href}
							class="group relative flex items-center {sidebarCollapsed
								? 'justify-center'
								: 'gap-3'} rounded-xl px-3 py-2.5 text-sm transition-all duration-200 {active
								? 'bg-primary/10 font-medium text-primary'
								: 'text-sidebar-foreground/80 hover:bg-sidebar-accent hover:text-sidebar-foreground'}"
							title={sidebarCollapsed ? item.label : undefined}
						>
							{#if active}
								<div
									class="absolute top-1/2 left-0 h-6 w-1 -translate-y-1/2 rounded-r-full bg-primary"
								></div>
							{/if}
							<Icon
								icon={item.icon}
								class="size-5 shrink-0 {active
									? 'text-primary'
									: 'text-muted-foreground group-hover:text-sidebar-foreground'} transition-colors"
							/>
							{#if !sidebarCollapsed}
								<span class="truncate">{item.label}</span>
							{/if}
						</a>
					{/each}
				</div>

				{#if !sidebarCollapsed}
					<div class="mt-6 mb-2 px-6">
						<div class="flex items-center gap-2">
							<div class="h-px flex-1 bg-sidebar-border/60"></div>
							<span
								class="text-[10px] font-semibold tracking-wider text-muted-foreground/60 uppercase"
								>Resources</span
							>
							<div class="h-px flex-1 bg-sidebar-border/60"></div>
						</div>
					</div>
				{:else}
					<div class="mt-6 flex justify-center">
						<div class="h-px w-8 bg-sidebar-border/60"></div>
					</div>
				{/if}
				<div class="space-y-0.5 {sidebarCollapsed ? 'px-2' : 'px-3'}">
					{#each resourceNavigation as item}
						{@const active = isActive(item.href)}
						<a
							href={item.href}
							class="group relative flex items-center {sidebarCollapsed
								? 'justify-center'
								: 'gap-3'} rounded-xl px-3 py-2.5 text-sm transition-all duration-200 {active
								? 'bg-primary/10 font-medium text-primary'
								: 'text-sidebar-foreground/80 hover:bg-sidebar-accent hover:text-sidebar-foreground'}"
							title={sidebarCollapsed ? item.label : undefined}
						>
							{#if active}
								<div
									class="absolute top-1/2 left-0 h-6 w-1 -translate-y-1/2 rounded-r-full bg-primary"
								></div>
							{/if}
							<div class="relative">
								<Icon
									icon={item.icon}
									class="size-5 shrink-0 {active
										? 'text-primary'
										: 'text-muted-foreground group-hover:text-sidebar-foreground'} transition-colors"
								/>
								{#if item.badge && !sidebarCollapsed}
									<span
										class="absolute -top-1.5 -right-1.5 flex h-4 min-w-4 items-center justify-center rounded-full bg-destructive px-1 text-[10px] font-bold text-destructive-foreground"
										>{item.badge}</span
									>
								{/if}
							</div>
							{#if !sidebarCollapsed}
								<span class="truncate">{item.label}</span>
								{#if item.badge}
									<span
										class="ml-auto flex h-5 min-w-5 items-center justify-center rounded-full bg-destructive/10 px-1.5 text-[10px] font-bold text-destructive"
										>{item.badge}</span
									>
								{/if}
							{/if}
						</a>
					{/each}
				</div>
			</nav>

			<div class="border-t border-sidebar-border p-3">
				<div class="space-y-0.5 {sidebarCollapsed ? 'px-0' : 'px-0'}">
					{#each settingsNavigation as item}
						{@const active = isActive(item.href)}
						<a
							href={item.href}
							class="group relative flex items-center {sidebarCollapsed
								? 'justify-center'
								: 'gap-3'} rounded-xl px-3 py-2 text-sm transition-all duration-200 {active
								? 'bg-primary/10 font-medium text-primary'
								: 'text-sidebar-foreground/80 hover:bg-sidebar-accent hover:text-sidebar-foreground'}"
							title={sidebarCollapsed ? item.label : undefined}
						>
							<Icon
								icon={item.icon}
								class="size-[18px] shrink-0 {active
									? 'text-sidebar-foreground'
									: 'text-muted-foreground group-hover:text-sidebar-foreground/90'} transition-colors"
							/>
							{#if !sidebarCollapsed}
								<span class="truncate">{item.label}</span>
							{/if}
						</a>
					{/each}
				</div>

				{#if !sidebarCollapsed}
					<Separator class="my-3 bg-sidebar-border/60" />
					<button
						class="group flex w-full items-center gap-3 rounded-xl p-2.5 transition-colors hover:bg-sidebar-accent"
						onclick={(e) => {
							e.stopPropagation();
							showUserMenu = !showUserMenu;
						}}
					>
						<img
							src={user.avatar}
							alt={user.name}
							class="size-9 rounded-full bg-muted ring-2 ring-transparent transition-all group-hover:ring-border"
						/>
						<div class="min-w-0 flex-1 text-left">
							<div class="truncate text-sm font-medium text-sidebar-foreground">{user.name}</div>
							<div class="truncate text-xs text-muted-foreground">{user.email}</div>
						</div>
						<Icon
							icon="lucide:more-vertical"
							class="size-4 text-muted-foreground opacity-0 transition-opacity group-hover:opacity-100"
						/>
					</button>
				{:else}
					<div class="flex justify-center pt-2">
						<button
							class="group relative"
							onclick={(e) => {
								e.stopPropagation();
								showUserMenu = !showUserMenu;
							}}
						>
							<img
								src={user.avatar}
								alt={user.name}
								class="size-10 rounded-full bg-muted ring-2 ring-transparent transition-all group-hover:ring-border"
							/>
							<span
								class="absolute right-0 bottom-0 size-3 rounded-full bg-emerald-500 ring-2 ring-sidebar"
							></span>
						</button>
					</div>
				{/if}
			</div>
		</aside>

		<div class="flex min-h-screen flex-1 flex-col">
			<header
				class="sticky top-0 z-30 hidden h-16 items-center justify-between border-b border-border bg-background/95 px-6 backdrop-blur supports-[backdrop-filter]:bg-background/60 lg:flex"
			>
				<div class="flex min-w-0 items-center gap-4">
					<Button variant="ghost" size="icon" class="shrink-0" onclick={toggleSidebar}>
						<Icon
							icon={sidebarCollapsed ? 'lucide:panel-left-open' : 'lucide:panel-left-close'}
							class="size-5 text-muted-foreground"
						/>
					</Button>
					<div class="h-6 w-px bg-border"></div>
					<div class="min-w-0">
						<h1 class="truncate text-lg font-semibold tracking-tight">{pageTitle().title}</h1>
						<p class="hidden truncate text-xs text-muted-foreground xl:block">
							{pageTitle().subtitle}
						</p>
					</div>
				</div>

				<div class="flex shrink-0 items-center gap-2">
					<Button
						variant="outline"
						class="h-9 gap-2 border-transparent bg-muted/50 px-3 text-muted-foreground hover:border-border hover:bg-muted hover:text-foreground"
						onclick={() => (commandPaletteOpen = true)}
					>
						<Icon icon="lucide:search" class="size-4" />
						<span class="hidden text-sm md:inline">Search</span>
						<kbd
							class="hidden rounded border bg-background px-1.5 text-[10px] font-medium text-muted-foreground xl:inline-block"
							>⌘K</kbd
						>
					</Button>

					<Separator orientation="vertical" class="mx-1 h-6" />

					<Button
						size="sm"
						class="h-9 gap-2 bg-primary shadow-lg shadow-primary/20 transition-all hover:-translate-y-0.5 hover:shadow-xl hover:shadow-primary/30"
					>
						<Icon icon="lucide:plus" class="size-4" />
						<span class="hidden sm:inline">Create</span>
					</Button>

					<Separator orientation="vertical" class="mx-1 h-6" />

					<Button variant="ghost" size="icon" class="size-9" onclick={toggleMode}>
						{#if mode.current === 'dark'}
							<Icon icon="lucide:moon" class="size-[18px]" />
						{:else}
							<Icon icon="lucide:sun" class="size-[18px]" />
						{/if}
					</Button>

					<div class="relative">
						<Button
							variant="ghost"
							size="icon"
							class="relative size-9"
							onclick={(e) => {
								e.stopPropagation();
								showNotifications = !showNotifications;
							}}
						>
							<Icon icon="lucide:bell" class="size-[18px]" />
							<span
								class="absolute top-2 right-2 size-2 rounded-full bg-destructive ring-2 ring-background"
							></span>
						</Button>

						{#if showNotifications}
							<div
								class="absolute top-full right-0 mt-2 w-80 rounded-2xl border border-border bg-popover p-2 shadow-2xl"
							>
								<div class="flex items-center justify-between px-3 py-2">
									<span class="text-sm font-semibold">Notifications</span>
									<Button
										variant="ghost"
										size="sm"
										class="h-auto px-2 py-1 text-xs text-muted-foreground hover:text-foreground"
									>
										Mark all read
									</Button>
								</div>
								<div class="max-h-80 space-y-0.5 overflow-y-auto py-1">
									{#each notifications as notification}
										<button
											class="w-full rounded-xl px-3 py-2.5 text-left transition-colors hover:bg-muted {notification.read
												? ''
												: 'bg-primary/5'}"
										>
											<div class="flex items-start gap-3">
												{#if !notification.read}
													<span class="mt-1.5 size-2 shrink-0 rounded-full bg-primary"></span>
												{:else}
													<span class="mt-1.5 size-2 shrink-0 rounded-full bg-transparent"></span>
												{/if}
												<div class="min-w-0 flex-1">
													<div class="text-sm font-medium">{notification.title}</div>
													<div class="line-clamp-1 text-xs text-muted-foreground">
														{notification.message}
													</div>
													<div class="mt-1 text-[10px] text-muted-foreground/70">
														{notification.time}
													</div>
												</div>
											</div>
										</button>
									{/each}
								</div>
								<div class="mt-2 border-t border-border px-1 pt-2">
									<Button
										variant="ghost"
										size="sm"
										class="w-full text-xs"
										href="/app/notifications"
									>
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

	<Modal bind:open={showUpgradeModal} title="" description="" size="lg" showCloseButton={true}>
		<div class="py-6">
			<div class="mb-8 text-center">
				<div class="mb-4 inline-flex size-12 items-center justify-center rounded-2xl bg-primary/10">
					<Icon icon="lucide:sparkles" class="size-6 text-primary" />
				</div>
				<h2 class="mb-2 text-2xl font-bold">Upgrade your workspace</h2>
				<p class="mx-auto max-w-md text-sm text-muted-foreground">
					Unlock more features and collaborate with your entire team.
				</p>
			</div>

			<div class="grid gap-4 md:grid-cols-3">
				<div class="rounded-2xl border border-border bg-card p-5">
					<div class="mb-4">
						<h3 class="mb-1 font-semibold">Free</h3>
						<p class="text-sm text-muted-foreground">For individuals</p>
					</div>
					<div class="mb-4">
						<span class="text-3xl font-bold">$0</span>
						<span class="text-sm text-muted-foreground">/mo</span>
					</div>
					<ul class="mb-6 space-y-2.5">
						<li class="flex items-center gap-2 text-sm">
							<Icon icon="lucide:check" class="size-4 text-primary" />
							<span>3 projects</span>
						</li>
						<li class="flex items-center gap-2 text-sm">
							<Icon icon="lucide:check" class="size-4 text-primary" />
							<span>1GB storage</span>
						</li>
					</ul>
				</div>

				<div class="relative rounded-2xl border-2 border-primary bg-primary/5 p-5">
					<div class="absolute -top-3 left-1/2 -translate-x-1/2">
						<span
							class="rounded-full bg-primary px-2.5 py-1 text-[10px] font-bold tracking-wide text-primary-foreground uppercase"
							>Popular</span
						>
					</div>
					<div class="mb-4">
						<h3 class="mb-1 font-semibold">Pro</h3>
						<p class="text-sm text-muted-foreground">For professionals</p>
					</div>
					<div class="mb-4">
						<span class="text-3xl font-bold">$12</span>
						<span class="text-sm text-muted-foreground">/mo</span>
					</div>
					<ul class="mb-6 space-y-2.5">
						<li class="flex items-center gap-2 text-sm">
							<Icon icon="lucide:check" class="size-4 text-primary" />
							<span>Unlimited projects</span>
						</li>
						<li class="flex items-center gap-2 text-sm">
							<Icon icon="lucide:check" class="size-4 text-primary" />
							<span>50GB storage</span>
						</li>
						<li class="flex items-center gap-2 text-sm">
							<Icon icon="lucide:check" class="size-4 text-primary" />
							<span>Priority support</span>
						</li>
					</ul>
					<Button class="w-full">Upgrade Now</Button>
				</div>

				<div class="rounded-2xl border border-border bg-card p-5">
					<div class="mb-4">
						<h3 class="mb-1 font-semibold">Team</h3>
						<p class="text-sm text-muted-foreground">For teams</p>
					</div>
					<div class="mb-4">
						<span class="text-3xl font-bold">$29</span>
						<span class="text-sm text-muted-foreground">/mo</span>
					</div>
					<ul class="mb-6 space-y-2.5">
						<li class="flex items-center gap-2 text-sm">
							<Icon icon="lucide:check" class="size-4 text-primary" />
							<span>Everything in Pro</span>
						</li>
						<li class="flex items-center gap-2 text-sm">
							<Icon icon="lucide:check" class="size-4 text-primary" />
							<span>Team collaboration</span>
						</li>
						<li class="flex items-center gap-2 text-sm">
							<Icon icon="lucide:check" class="size-4 text-primary" />
							<span>Admin controls</span>
						</li>
					</ul>
				</div>
			</div>
		</div>
	</Modal>
</div>
