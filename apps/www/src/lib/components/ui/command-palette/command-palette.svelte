<script lang="ts">
	import { cn } from '$lib/utils.js';
	import Icon from '@iconify/svelte';

	interface Props {
		open?: boolean;
		onClose?: () => void;
	}

	let { open = $bindable(false), onClose }: Props = $props();

	let searchQuery = $state('');
	let selectedIndex = $state(0);
	let inputRef: HTMLInputElement | null = $state(null);

	interface CommandItem {
		id: string;
		label: string;
		description?: string;
		icon: string;
		shortcut?: string;
		action: () => void;
		category: string;
	}

	const recentItems: CommandItem[] = [
		{ id: 'r1', label: 'Website Redesign', description: 'Project', icon: 'hugeicons:folder-kanban', action: () => { window.location.href = '/projects'; close(); }, category: 'Recent' },
		{ id: 'r2', label: 'Acme Corporation', description: 'Client', icon: 'hugeicons:building-02', action: () => { window.location.href = '/clients'; close(); }, category: 'Recent' },
		{ id: 'r3', label: 'Invoice #2026-001', description: '$2,500 • Paid', icon: 'hugeicons:invoice-01', action: () => { window.location.href = '/invoices'; close(); }, category: 'Recent' }
	];

	const navigationItems: CommandItem[] = [
		{ id: 'n1', label: 'Dashboard', icon: 'hugeicons:layout-dashboard', shortcut: 'G D', action: () => { window.location.href = '/'; close(); }, category: 'Navigation' },
		{ id: 'n2', label: 'Projects', icon: 'hugeicons:folder-kanban', shortcut: 'G P', action: () => { window.location.href = '/projects'; close(); }, category: 'Navigation' },
		{ id: 'n3', label: 'Clients', icon: 'hugeicons:users', shortcut: 'G C', action: () => { window.location.href = '/clients'; close(); }, category: 'Navigation' },
		{ id: 'n4', label: 'Tasks', icon: 'hugeicons:check-square', shortcut: 'G T', action: () => { window.location.href = '/tasks'; close(); }, category: 'Navigation' },
		{ id: 'n5', label: 'Time Tracking', icon: 'hugeicons:clock', shortcut: 'G M', action: () => { window.location.href = '/time'; close(); }, category: 'Navigation' },
		{ id: 'n6', label: 'Invoices', icon: 'hugeicons:invoice-01', shortcut: 'G I', action: () => { window.location.href = '/invoices'; close(); }, category: 'Navigation' },
		{ id: 'n7', label: 'Files', icon: 'hugeicons:folder-open', shortcut: 'G F', action: () => { window.location.href = '/files'; close(); }, category: 'Navigation' },
		{ id: 'n8', label: 'Templates', icon: 'hugeicons:layout-template', shortcut: 'G L', action: () => { window.location.href = '/templates'; close(); }, category: 'Navigation' },
		{ id: 'n9', label: 'Emails', icon: 'hugeicons:mail-01', shortcut: 'G E', action: () => { window.location.href = '/emails'; close(); }, category: 'Navigation' },
		{ id: 'n10', label: 'Settings', icon: 'hugeicons:settings-01', shortcut: 'G S', action: () => { window.location.href = '/settings'; close(); }, category: 'Navigation' }
	];

	const actionItems: CommandItem[] = [
		{ id: 'a1', label: 'New Project', icon: 'hugeicons:plus-sign', shortcut: 'C P', action: () => { window.location.href = '/projects'; close(); }, category: 'Actions' },
		{ id: 'a2', label: 'New Task', icon: 'hugeicons:check-square', shortcut: 'C T', action: () => { window.location.href = '/tasks'; close(); }, category: 'Actions' },
		{ id: 'a3', label: 'New Invoice', icon: 'hugeicons:invoice-01', shortcut: 'C I', action: () => { window.location.href = '/invoices'; close(); }, category: 'Actions' },
		{ id: 'a4', label: 'Upload File', icon: 'hugeicons:upload-01', action: () => { window.location.href = '/files'; close(); }, category: 'Actions' },
		{ id: 'a5', label: 'Create Template', icon: 'hugeicons:layout-template', action: () => { window.location.href = '/templates'; close(); }, category: 'Actions' },
		{ id: 'a6', label: 'Compose Email', icon: 'hugeicons:pencil-edit-01', action: () => { window.location.href = '/emails'; close(); }, category: 'Actions' },
		{ id: 'a7', label: 'Start Timer', icon: 'hugeicons:play', action: () => { window.location.href = '/time'; close(); }, category: 'Actions' }
	];

	const allItems = [...recentItems, ...navigationItems, ...actionItems];

	const filteredItems = $derived(() => {
		if (!searchQuery.trim()) return [...recentItems.slice(0, 3), ...navigationItems, ...actionItems];
		
		const q = searchQuery.toLowerCase();
		return allItems.filter(item => 
			item.label.toLowerCase().includes(q) ||
			item.description?.toLowerCase().includes(q) ||
			item.category.toLowerCase().includes(q)
		);
	});

	const groupedItems = $derived(() => {
		const items = filteredItems();
		const groups: Record<string, CommandItem[]> = {};
		
		items.forEach(item => {
			if (!groups[item.category]) groups[item.category] = [];
			groups[item.category].push(item);
		});
		
		return groups;
	});

	function close() {
		open = false;
		searchQuery = '';
		selectedIndex = 0;
		onClose?.();
	}

	function handleKeydown(e: KeyboardEvent) {
		if (!open) return;
		
		const items = filteredItems();
		
		switch (e.key) {
			case 'Escape':
				e.preventDefault();
				close();
				break;
			case 'ArrowDown':
				e.preventDefault();
				selectedIndex = Math.min(selectedIndex + 1, items.length - 1);
				scrollIntoView();
				break;
			case 'ArrowUp':
				e.preventDefault();
				selectedIndex = Math.max(selectedIndex - 1, 0);
				scrollIntoView();
				break;
			case 'Enter':
				e.preventDefault();
				if (items[selectedIndex]) {
					items[selectedIndex].action();
				}
				break;
		}
	}

	function scrollIntoView() {
		setTimeout(() => {
			const element = document.querySelector(`[data-index="${selectedIndex}"]`);
			element?.scrollIntoView({ block: 'nearest', behavior: 'smooth' });
		}, 0);
	}

	function handleOverlayClick(e: MouseEvent) {
		if (e.target === e.currentTarget) close();
	}

	$effect(() => {
		if (open && inputRef) {
			setTimeout(() => inputRef?.focus(), 50);
		}
	});
</script>

<svelte:window onkeydown={handleKeydown} />

{#if open}
	<div 
		class="fixed inset-0 z-50 flex items-start justify-center pt-[20vh]"
		onclick={handleOverlayClick}
	>
		<div class="absolute inset-0 bg-background/80 backdrop-blur-sm"></div>
		
		<div class="relative w-full max-w-2xl mx-4 rounded-xl border border-border bg-card shadow-2xl overflow-hidden flex flex-col max-h-[60vh]">
			<div class="flex items-center gap-3 px-4 py-4 border-b border-border">
				<Icon icon="hugeicons:search-01" class="size-5 text-muted-foreground shrink-0" />
				<input
					bind:this={inputRef}
					type="text"
					placeholder="Search pages, actions, or type a command..."
					class="flex-1 bg-transparent text-lg outline-none placeholder:text-muted-foreground"
					bind:value={searchQuery}
					onkeydown={handleKeydown}
				/>
				<div class="flex items-center gap-1 shrink-0">
					<kbd class="hidden sm:inline-flex h-7 items-center gap-1 rounded border bg-muted px-2 text-xs font-medium">
						<Icon icon="hugeicons:arrow-up" class="size-3" />
					</kbd>
					<kbd class="hidden sm:inline-flex h-7 items-center gap-1 rounded border bg-muted px-2 text-xs font-medium">
						<Icon icon="hugeicons:arrow-down" class="size-3" />
					</kbd>
					<kbd class="hidden sm:inline-flex h-7 items-center gap-1 rounded border bg-muted px-2 text-xs font-medium ml-2">
						ESC
					</kbd>
				</div>
			</div>
			
			<div class="flex-1 overflow-y-auto p-2">
				{#if filteredItems().length === 0}
					<div class="flex flex-col items-center justify-center py-12 text-center">
						<div class="size-12 rounded-full bg-muted flex items-center justify-center mb-3">
							<Icon icon="hugeicons:search-01" class="size-6 text-muted-foreground" />
						</div>
						<p class="text-sm text-muted-foreground">No results found for "{searchQuery}"</p>
						<p class="text-xs text-muted-foreground mt-1">Try searching for pages, actions, or recent items</p>
					</div>
				{:else}
					{#each Object.entries(groupedItems()) as [category, items], catIndex}
						{#if items.length > 0}
							<div class="px-2 py-2">
								<p class="text-xs font-semibold text-muted-foreground uppercase tracking-wider px-2 py-1">{category}</p>
								<div class="space-y-0.5">
									{#each items as item, itemIndex}
										{@const globalIndex = catIndex * 100 + itemIndex}
										<button
											data-index={globalIndex}
											class={cn(
												'w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-left transition-colors',
												selectedIndex === globalIndex ? 'bg-primary text-primary-foreground' : 'hover:bg-muted'
											)}
											onmouseenter={() => selectedIndex = globalIndex}
											onclick={() => item.action()}
										>
											<div class="size-8 rounded-lg {selectedIndex === globalIndex ? 'bg-primary-foreground/20' : 'bg-muted'} flex items-center justify-center shrink-0">
												<Icon icon={item.icon} class="size-4" />
											</div>
											<div class="flex-1 min-w-0">
												<p class="font-medium truncate">{item.label}</p>
												{#if item.description}
													<p class={cn('text-xs truncate', selectedIndex === globalIndex ? 'text-primary-foreground/70' : 'text-muted-foreground')}>{item.description}</p>
												{/if}
											</div>
											{#if item.shortcut}
												<div class="flex items-center gap-0.5 shrink-0">
													{#each item.shortcut.split(' ') as key}
														<kbd class={cn(
															'inline-flex h-6 min-w-6 items-center justify-center rounded px-1.5 text-xs font-medium',
															selectedIndex === globalIndex ? 'bg-primary-foreground/20 text-primary-foreground' : 'bg-muted text-muted-foreground'
														)}>
															{key}
														</kbd>
													{/each}
												</div>
											{/if}
										</button>
									{/each}
								</div>
							</div>
						{/if}
					{/each}
				{/if}
			</div>
			
			<div class="border-t border-border px-4 py-2 bg-muted/50 flex items-center justify-between text-xs text-muted-foreground">
				<div class="flex items-center gap-3">
					<span class="flex items-center gap-1">
						<kbd class="inline-flex h-5 items-center rounded border bg-background px-1.5 text-[10px]">↵</kbd>
						to select
					</span>
				</div>
				<p class="hidden sm:block">Pro tip: Press <kbd class="inline-flex h-5 items-center rounded border bg-background px-1.5 text-[10px]">⌘</kbd> <kbd class="inline-flex h-5 items-center rounded border bg-background px-1.5 text-[10px]">K</kbd> from anywhere</p>
			</div>
		</div>
	</div>
{/if}
