<script lang="ts">
	import Icon from '@iconify/svelte';
	import { page } from '$app/state';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';

	const { children } = $props();

	const navigation = [
		{ href: '/settings', label: 'Profile', icon: 'lucide:user' },
		{ href: '/settings/security', label: 'Security', icon: 'lucide:shield' },
		{
			href: '/settings/notifications',
			label: 'Notifications',
			icon: 'lucide:bell'
		},
		{ href: '/settings/appearance', label: 'Appearance', icon: 'lucide:palette' },
		{ href: '/settings/billing', label: 'Billing', icon: 'lucide:credit-card' }
	];

	const isActive = (href: string) => {
		if (href === '/settings') {
			return page.url.pathname === '/settings';
		}
		return page.url.pathname.startsWith(href);
	};
</script>

<div class="min-h-screen bg-background">
	<header
		class="sticky top-0 z-50 w-full border-b border-border/40 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60"
	>
		<div class="flex h-14 items-center gap-4 px-4 md:px-8">
			<a href="/" class="flex items-center gap-2 text-lg font-semibold">
				<div
					class="flex size-8 items-center justify-center rounded-md bg-primary text-primary-foreground"
				>
					<Icon icon="hugeicons:cowboy-hat" class="size-5" />
				</div>
				<span class="hidden sm:inline">Artemis</span>
			</a>
			<Separator orientation="vertical" class="h-6" />
			<span class="text-sm text-muted-foreground">Settings</span>
		</div>
	</header>

	<div class="flex flex-col lg:flex-row">
		<aside
			class="w-full border-b border-border/40 lg:w-64 lg:border-r lg:border-b-0 lg:bg-muted/30"
		>
			<nav class="flex gap-1 overflow-x-auto p-2 lg:flex-col lg:p-4">
				{#each navigation as item}
					<Button
						variant={isActive(item.href) ? 'secondary' : 'ghost'}
						class="justify-start gap-3 whitespace-nowrap"
						href={item.href}
					>
						<Icon icon={item.icon} class="size-4" />
						<span class="hidden sm:inline">{item.label}</span>
					</Button>
				{/each}
			</nav>
		</aside>

		<main class="flex-1 p-4 md:p-8">
			<div class="mx-auto max-w-3xl">
				{@render children()}
			</div>
		</main>
	</div>
</div>
