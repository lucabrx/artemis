<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';
	import { Label } from '$lib/components/ui/label/index.js';

	let viewMode = $state<'grid' | 'list'>('grid');
	let currentFolder = $state<string | null>(null);
	let selectedFiles = $state<string[]>([]);
	let searchQuery = $state('');
	let filterType = $state<'all' | 'project' | 'general'>('all');
	let showUploadModal = $state(false);
	let showNewFolderModal = $state(false);
	let showPreviewModal = $state(false);
	let previewFile = $state<FileItem | null>(null);
	let contextMenuFile = $state<FileItem | null>(null);
	let showContextMenu = $state(false);
	let contextMenuPos = $state({ x: 0, y: 0 });
	let isDragging = $state(false);

	interface FileItem {
		id: string;
		name: string;
		type: 'file' | 'folder';
		extension?: string;
		size?: string;
		project?: {
			id: string;
			name: string;
			color: string;
		};
		uploadedBy: {
			name: string;
			avatar: string;
		};
		uploadedAt: string;
		updatedAt: string;
		thumbnail?: string;
		parentId: string | null;
		starred?: boolean;
	}

	const folders: FileItem[] = [
		{
			id: 'f1',
			name: 'Website Redesign',
			type: 'folder',
			uploadedBy: { name: 'Alex Morgan', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex' },
			uploadedAt: 'Jan 15, 2026',
			updatedAt: '2 hours ago',
			parentId: null,
			starred: true
		},
		{
			id: 'f2',
			name: 'Brand Assets',
			type: 'folder',
			uploadedBy: { name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah' },
			uploadedAt: 'Jan 10, 2026',
			updatedAt: '1 day ago',
			parentId: null
		},
		{
			id: 'f3',
			name: 'Invoices 2026',
			type: 'folder',
			uploadedBy: { name: 'Alex Morgan', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex' },
			uploadedAt: 'Jan 5, 2026',
			updatedAt: '3 days ago',
			parentId: null
		},
		{
			id: 'f4',
			name: 'Mobile App Designs',
			type: 'folder',
			project: { id: 'p2', name: 'Mobile App', color: 'bg-violet-500' },
			uploadedBy: { name: 'Mike Ross', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Mike' },
			uploadedAt: 'Jan 20, 2026',
			updatedAt: '5 hours ago',
			parentId: null,
			starred: true
		}
	];

	const files: FileItem[] = [
		{
			id: 'file1',
			name: 'Homepage_V3_Final',
			type: 'file',
			extension: 'fig',
			size: '24.5 MB',
			project: { id: 'p1', name: 'Website Redesign', color: 'bg-blue-500' },
			uploadedBy: { name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah' },
			uploadedAt: 'Jan 28, 2026',
			updatedAt: '2 hours ago',
			thumbnail: 'https://placehold.co/300x200/e0e7ff/6366f1?text=Figma',
			parentId: 'f1',
			starred: true
		},
		{
			id: 'file2',
			name: 'Logo_Primary_Vector',
			type: 'file',
			extension: 'svg',
			size: '156 KB',
			project: { id: 'p3', name: 'Brand Identity', color: 'bg-emerald-500' },
			uploadedBy: { name: 'Emma Wilson', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Emma' },
			uploadedAt: 'Jan 25, 2026',
			updatedAt: '1 day ago',
			thumbnail: 'https://placehold.co/300x200/dcfce7/22c55e?text=SVG',
			parentId: 'f2'
		},
		{
			id: 'file3',
			name: 'Q1_Financial_Report',
			type: 'file',
			extension: 'pdf',
			size: '3.2 MB',
			uploadedBy: { name: 'Alex Morgan', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex' },
			uploadedAt: 'Jan 26, 2026',
			updatedAt: '2 days ago',
			thumbnail: 'https://placehold.co/300x200/fee2e2/ef4444?text=PDF',
			parentId: null
		},
		{
			id: 'file4',
			name: 'User_Research_Summary',
			type: 'file',
			extension: 'docx',
			size: '1.8 MB',
			project: { id: 'p1', name: 'Website Redesign', color: 'bg-blue-500' },
			uploadedBy: { name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah' },
			uploadedAt: 'Jan 24, 2026',
			updatedAt: '4 days ago',
			thumbnail: 'https://placehold.co/300x200/dbeafe/3b82f6?text=DOCX',
			parentId: 'f1'
		},
		{
			id: 'file5',
			name: 'App_Icons_All_Sizes',
			type: 'file',
			extension: 'png',
			size: '8.4 MB',
			project: { id: 'p2', name: 'Mobile App', color: 'bg-violet-500' },
			uploadedBy: { name: 'Lisa Wang', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Lisa' },
			uploadedAt: 'Jan 27, 2026',
			updatedAt: '12 hours ago',
			thumbnail: 'https://placehold.co/300x200/f3e8ff/a855f7?text=PNG',
			parentId: 'f4',
			starred: true
		},
		{
			id: 'file6',
			name: 'Contract_AcmeCorp_2026',
			type: 'file',
			extension: 'pdf',
			size: '2.1 MB',
			project: { id: 'p1', name: 'Website Redesign', color: 'bg-blue-500' },
			uploadedBy: { name: 'Alex Morgan', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex' },
			uploadedAt: 'Jan 22, 2026',
			updatedAt: '1 week ago',
			thumbnail: 'https://placehold.co/300x200/fee2e2/ef4444?text=PDF',
			parentId: null
		},
		{
			id: 'file7',
			name: 'Marketing_Banner_Set',
			type: 'file',
			extension: 'psd',
			size: '45.2 MB',
			project: { id: 'p4', name: 'Marketing Campaign', color: 'bg-amber-500' },
			uploadedBy: { name: 'Jane Doe', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Jane' },
			uploadedAt: 'Jan 23, 2026',
			updatedAt: '5 days ago',
			thumbnail: 'https://placehold.co/300x200/fef3c7/f59e0b?text=PSD',
			parentId: null
		},
		{
			id: 'file8',
			name: 'Database_Schema_v2',
			type: 'file',
			extension: 'sql',
			size: '24 KB',
			project: { id: 'p5', name: 'E-commerce Platform', color: 'bg-rose-500' },
			uploadedBy: { name: 'David Kim', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=David' },
			uploadedAt: 'Jan 21, 2026',
			updatedAt: '1 week ago',
			thumbnail: 'https://placehold.co/300x200/ffe4e6/f43f5e?text=SQL',
			parentId: null
		}
	];

	const allItems = [...folders, ...files];

	const currentFolderData = $derived(folders.find(f => f.id === currentFolder));

	const currentFolderName = $derived(() => {
		if (!currentFolder) return 'All Files';
		return currentFolderData?.name || 'All Files';
	});

	const breadcrumb = $derived(() => {
		if (!currentFolder) return [{ id: null, name: 'Files' }];
		return [{ id: null, name: 'Files' }, { id: currentFolder, name: currentFolderData?.name }];
	});

	const filteredItems = $derived(() => {
		let items = allItems.filter(item => item.parentId === currentFolder);
		
		if (filterType === 'project') {
			items = items.filter(item => item.project);
		} else if (filterType === 'general') {
			items = items.filter(item => !item.project);
		}
		
		if (searchQuery) {
			const q = searchQuery.toLowerCase();
			items = items.filter(item => 
				item.name.toLowerCase().includes(q) ||
				item.extension?.toLowerCase().includes(q)
			);
		}
		
		return items;
	});

	const recentFiles = $derived(() => {
		return [...files].sort((a, b) => {
			const dateA = new Date(a.updatedAt.includes('hour') ? Date.now() : a.updatedAt).getTime();
			const dateB = new Date(b.updatedAt.includes('hour') ? Date.now() : b.updatedAt).getTime();
			return dateB - dateA;
		}).slice(0, 4);
	});

	const starredCount = $derived(() => {
		return allItems.filter(item => item.starred).length;
	});

	function getFileIcon(extension?: string): string {
		const icons: Record<string, string> = {
			fig: 'hugeicons:figma',
			svg: 'hugeicons:svg-01',
			pdf: 'hugeicons:pdf-01',
			docx: 'hugeicons:doc-01',
			png: 'hugeicons:png-01',
			jpg: 'hugeicons:jpg-01',
			psd: 'hugeicons:psd',
			sql: 'hugeicons:database',
			xlsx: 'hugeicons:xls-01',
			zip: 'hugeicons:zip-01'
		};
		return icons[extension || ''] || 'hugeicons:file-01';
	}

	function getFileColor(extension?: string): string {
		const colors: Record<string, string> = {
			fig: 'text-purple-600 bg-purple-500/10',
			svg: 'text-green-600 bg-green-500/10',
			pdf: 'text-red-600 bg-red-500/10',
			docx: 'text-blue-600 bg-blue-500/10',
			png: 'text-purple-600 bg-purple-500/10',
			jpg: 'text-purple-600 bg-purple-500/10',
			psd: 'text-amber-600 bg-amber-500/10',
			sql: 'text-rose-600 bg-rose-500/10'
		};
		return colors[extension || ''] || 'text-slate-600 bg-slate-500/10';
	}

	function handleFileClick(item: FileItem) {
		if (item.type === 'folder') {
			currentFolder = item.id;
			selectedFiles = [];
			searchQuery = '';
		} else {
			previewFile = item;
			showPreviewModal = true;
		}
	}

	function navigateToFolder(folderId: string | null) {
		currentFolder = folderId;
		selectedFiles = [];
		searchQuery = '';
	}

	function toggleSelect(id: string, e: Event) {
		e.stopPropagation();
		if (selectedFiles.includes(id)) {
			selectedFiles = selectedFiles.filter(f => f !== id);
		} else {
			selectedFiles = [...selectedFiles, id];
		}
	}

	function toggleStar(e: Event, item: FileItem) {
		e.stopPropagation();
		item.starred = !item.starred;
	}

	function handleContextMenu(e: MouseEvent, item: FileItem) {
		e.preventDefault();
		const rect = (e.currentTarget as HTMLElement).getBoundingClientRect();
		const x = Math.min(e.clientX, window.innerWidth - 200);
		const y = Math.min(e.clientY, window.innerHeight - 250);
		contextMenuFile = item;
		contextMenuPos = { x, y };
		showContextMenu = true;
	}

	function closeContextMenu() {
		showContextMenu = false;
		setTimeout(() => { contextMenuFile = null; }, 100);
	}

	function handleDragOver(e: DragEvent) {
		e.preventDefault();
		isDragging = true;
	}

	function handleDragLeave(e: DragEvent) {
		e.preventDefault();
		isDragging = false;
	}

	function handleDrop(e: DragEvent) {
		e.preventDefault();
		isDragging = false;
	}

	let newFolderName = $state('');

	function handleCreateFolder(e: Event) {
		e.preventDefault();
		if (!newFolderName.trim()) return;
		showNewFolderModal = false;
		newFolderName = '';
	}

	function handleUpload() {
		showUploadModal = false;
	}

	function clearSelection() {
		selectedFiles = [];
	}
</script>

<svelte:head>
	<title>Files - Artemis</title>
</svelte:head>

<svelte:window onclick={closeContextMenu} />

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="flex h-[calc(100vh-3.5rem)]">
		<div class="w-64 border-r border-border bg-muted/30 flex flex-col">
			<div class="p-4">
				<Button class="w-full gap-2" onclick={() => showUploadModal = true}>
					<Icon icon="hugeicons:upload-01" class="size-4" />
					Upload
				</Button>
			</div>
			
			<nav class="flex-1 px-2 space-y-0.5">
				<button
					class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-all {currentFolder === null && filterType === 'all' ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:bg-muted hover:text-foreground'}"
					onclick={() => { navigateToFolder(null); filterType = 'all'; }}
				>
					<Icon icon="hugeicons:folder-library" class="size-5" />
					All Files
				</button>
				<button
					class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-all {filterType === 'project' ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:bg-muted hover:text-foreground'}"
					onclick={() => { filterType = 'project'; navigateToFolder(null); }}
				>
					<Icon icon="hugeicons:folder-check" class="size-5" />
					Project Files
				</button>
				<button
					class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium transition-all {filterType === 'general' ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:bg-muted hover:text-foreground'}"
					onclick={() => { filterType = 'general'; navigateToFolder(null); }}
				>
					<Icon icon="hugeicons:folder" class="size-5" />
					General Files
				</button>
				
				<Separator class="my-3" />
				
				<p class="px-3 py-2 text-xs font-semibold text-muted-foreground uppercase tracking-wider">Quick Access</p>
				<button class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium text-muted-foreground hover:bg-muted hover:text-foreground transition-colors">
					<Icon icon="hugeicons:star" class="size-5 text-amber-500" />
					<span class="flex-1 text-left">Starred</span>
					{#if starredCount() > 0}
						<span class="text-xs bg-muted px-2 py-0.5 rounded-full">{starredCount()}</span>
					{/if}
				</button>
				<button class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium text-muted-foreground hover:bg-muted hover:text-foreground transition-colors">
					<Icon icon="hugeicons:clock" class="size-5" />
					Recent
				</button>
				<button class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium text-muted-foreground hover:bg-muted hover:text-foreground transition-colors">
					<Icon icon="hugeicons:share-01" class="size-5" />
					Shared
				</button>
				<button class="w-full flex items-center gap-3 px-3 py-2 rounded-lg text-sm font-medium text-muted-foreground hover:bg-muted hover:text-foreground transition-colors">
					<Icon icon="hugeicons:trash-01" class="size-5" />
					Trash
				</button>
			</nav>
			
			<div class="p-4 border-t border-border mt-auto">
				<div class="space-y-2">
					<div class="flex items-center justify-between text-sm">
						<span class="text-muted-foreground">Storage</span>
						<span class="font-medium">75%</span>
					</div>
					<div class="h-2 rounded-full bg-muted overflow-hidden">
						<div class="h-full w-3/4 bg-gradient-to-r from-primary to-primary/80 rounded-full"></div>
					</div>
					<p class="text-xs text-muted-foreground">15 GB of 20 GB used</p>
				</div>
			</div>
		</div>

		<div class="flex-1 flex flex-col min-w-0">
			<div class="px-6 py-4 border-b border-border bg-background/50 backdrop-blur">
				<div class="flex items-center justify-between gap-4">
					<div class="flex items-center gap-2 min-w-0">
						{#if currentFolder}
							<Button variant="ghost" size="icon" class="size-8 shrink-0" onclick={() => navigateToFolder(null)}>
								<Icon icon="hugeicons:arrow-left-01" class="size-5" />
							</Button>
						{/if}
						<div class="flex items-center gap-2 min-w-0">
							{#each breadcrumb() as crumb, i}
								{#if i > 0}
									<Icon icon="hugeicons:arrow-right-01" class="size-4 text-muted-foreground shrink-0" />
								{/if}
								{#if i === breadcrumb().length - 1}
									<span class="text-sm font-medium truncate">{crumb.name}</span>
								{:else}
									<button
										class="text-sm font-medium text-muted-foreground hover:text-foreground transition-colors"
										onclick={() => navigateToFolder(crumb.id)}
									>
										{crumb.name}
									</button>
								{/if}
							{/each}
						</div>
					</div>
					
					<div class="flex items-center gap-2 shrink-0">
						{#if selectedFiles.length > 0}
							<div class="flex items-center gap-2 bg-muted/50 rounded-lg px-3 py-1.5">
								<span class="text-sm font-medium">{selectedFiles.length} selected</span>
								<Separator orientation="vertical" class="h-4" />
								<button class="p-1 hover:bg-muted rounded" onclick={clearSelection}>
									<Icon icon="hugeicons:cancel-01" class="size-4" />
								</button>
							</div>
							<Button variant="outline" size="sm" class="gap-2">
								<Icon icon="hugeicons:share-01" class="size-4" />
								Share
								</Button>
							<Button variant="outline" size="sm" class="gap-2 text-destructive hover:text-destructive hover:border-destructive">
								<Icon icon="hugeicons:delete-01" class="size-4" />
								Delete
							</Button>
							<Separator orientation="vertical" class="h-6" />
						{/if}
						
						<div class="relative">
							<Icon icon="hugeicons:search-01" class="absolute left-3 top-1/2 -translate-y-1/2 size-4 text-muted-foreground" />
							<Input placeholder="Search files..." class="pl-9 w-56" bind:value={searchQuery} />
						</div>
						
						<Separator orientation="vertical" class="h-6" />
						
						<div class="flex rounded-lg border border-border p-1 bg-muted/30">
							<button
								class="rounded p-1.5 transition-all {viewMode === 'grid' ? 'bg-background text-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground'}"
								onclick={() => viewMode = 'grid'}
								title="Grid view"
							>
								<Icon icon="hugeicons:grid" class="size-4" />
							</button>
							<button
								class="rounded p-1.5 transition-all {viewMode === 'list' ? 'bg-background text-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground'}"
								onclick={() => viewMode = 'list'}
								title="List view"
							>
								<Icon icon="hugeicons:list" class="size-4" />
							</button>
						</div>
						
						<Button variant="outline" size="icon" onclick={() => showNewFolderModal = true} title="New folder">
							<Icon icon="hugeicons:folder-add" class="size-4" />
						</Button>
					</div>
				</div>
			</div>

			<div 
				class="flex-1 overflow-y-auto p-6 transition-colors {isDragging ? 'bg-primary/5' : ''}"
				ondragover={handleDragOver}
				ondragleave={handleDragLeave}
				ondrop={handleDrop}
			>
				{#if currentFolder === null && filterType === 'all' && !searchQuery}
					<div class="mb-8">
						<div class="flex items-center justify-between mb-4">
							<h2 class="text-lg font-semibold">Recent Files</h2>
							<button class="text-sm text-primary hover:underline">View all</button>
						</div>
						<div class="grid grid-cols-2 md:grid-cols-4 gap-4">
							{#each recentFiles() as file}
								<button
									class="group relative rounded-xl border border-border bg-card p-3 text-left transition-all hover:border-primary/20 hover:shadow-md hover:-translate-y-0.5"
									onclick={() => handleFileClick(file)}
									oncontextmenu={(e) => handleContextMenu(e, file)}
								>
									{#if file.starred}
										<div class="absolute top-2 right-2">
											<Icon icon="hugeicons:star" class="size-4 text-amber-500 fill-amber-500" />
										</div>
									{/if}
									<div class="aspect-video rounded-lg bg-muted mb-3 overflow-hidden">
										{#if file.thumbnail}
											<img src={file.thumbnail} alt={file.name} class="w-full h-full object-cover" />
										{:else}
											<div class="w-full h-full flex items-center justify-center {getFileColor(file.extension)}">
												<Icon icon={getFileIcon(file.extension)} class="size-12" />
											</div>
										{/if}
									</div>
									<p class="font-medium truncate">{file.name}</p>
									<p class="text-xs text-muted-foreground">{file.updatedAt}</p>
								</button>
							{/each}
						</div>
					</div>
				{/if}

				<div>
					<div class="flex items-center justify-between mb-4">
						<h2 class="text-lg font-semibold">{currentFolderName()}</h2>
						{#if filteredItems().length > 0}
							<p class="text-sm text-muted-foreground">{filteredItems().length} items</p>
						{/if}
					</div>
					
					{#if filteredItems().length === 0}
						<div class="flex flex-col items-center justify-center py-20 text-center">
							<div class="size-20 rounded-full bg-muted flex items-center justify-center mb-4">
								<Icon icon="hugeicons:folder-open" class="size-10 text-muted-foreground" />
							</div>
							<h3 class="text-lg font-medium">
								{#if searchQuery}
									No files match your search
								{:else if currentFolder}
									This folder is empty
								{:else}
									No files yet
								{/if}
							</h3>
							<p class="text-muted-foreground mt-1 max-w-sm">
								{#if searchQuery}
									Try adjusting your search terms
								{:else if currentFolder}
									Drag files here or use the upload button to add files
								{:else}
									Upload files or create a folder to get started
								{/if}
							</p>
							{#if !searchQuery}
								<Button class="mt-4 gap-2" onclick={() => showUploadModal = true}>
									<Icon icon="hugeicons:upload-01" class="size-4" />
									Upload Files
								</Button>
							{/if}
						</div>
					{:else if viewMode === 'grid'}
						<div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
							{#each filteredItems() as item (item.id)}
								<div
									class="group relative rounded-xl border border-border bg-card p-3 transition-all hover:border-primary/20 hover:shadow-md hover:-translate-y-0.5 {selectedFiles.includes(item.id) ? 'border-primary ring-2 ring-primary/20' : ''}"
									role="button"
									tabindex="0"
									onclick={() => handleFileClick(item)}
									onkeydown={(e) => { if (e.key === 'Enter') handleFileClick(item); }}
									oncontextmenu={(e) => handleContextMenu(e, item)}
								>
									<div class="absolute top-2 left-2 z-10">
										<button
											class="size-6 rounded-md border shadow-sm flex items-center justify-center transition-all {selectedFiles.includes(item.id) ? 'bg-primary border-primary text-primary-foreground' : 'bg-background/90 backdrop-blur border-border opacity-0 group-hover:opacity-100'}"
											onclick={(e) => toggleSelect(item.id, e)}
											type="button"
										>
											{#if selectedFiles.includes(item.id)}
												<Icon icon="hugeicons:tick-01" class="size-4" />
											{/if}
										</button>
									</div>
									
									<div class="absolute top-2 right-2 z-10">
										<button 
											class="size-6 rounded-md border border-border bg-background/90 backdrop-blur shadow-sm flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity hover:bg-muted"
											onclick={(e) => toggleStar(e, item)}
											type="button"
										>
											<Icon icon="hugeicons:star" class="size-4 {item.starred ? 'text-amber-500 fill-amber-500' : 'text-muted-foreground'}" />
										</button>
									</div>
									
									<div class="aspect-square rounded-lg bg-muted mb-3 flex items-center justify-center overflow-hidden">
										{#if item.type === 'folder'}
											<div class="flex flex-col items-center gap-2">
												<Icon icon="hugeicons:folder-01" class="size-16 text-amber-500" />
												{#if item.project}
													<span class="text-xs px-2 py-0.5 rounded-full bg-slate-800 text-white">{item.project.name}</span>
												{/if}
											</div>
										{:else if item.thumbnail}
											<img src={item.thumbnail} alt={item.name} class="w-full h-full object-cover" />
										{:else}
											<div class="flex flex-col items-center gap-2 {getFileColor(item.extension)}">
												<Icon icon={getFileIcon(item.extension)} class="size-16" />
											</div>
										{/if}
									</div>
									
									<p class="font-medium truncate" title={item.name}>{item.name}</p>
									<div class="flex items-center justify-between mt-1">
										<span class="text-xs text-muted-foreground">{item.type === 'file' ? item.size : item.updatedAt}</span>
										{#if item.project && item.type === 'file'}
											<div class="size-2 rounded-full {item.project.color}" title={item.project.name}></div>
										{/if}
									</div>
								</div>
							{/each}
						</div>
					{:else}
						<div class="rounded-xl border border-border bg-card overflow-hidden">
							<div class="grid grid-cols-12 gap-4 px-4 py-3 text-sm font-medium text-muted-foreground border-b border-border bg-muted/50">
								<div class="col-span-6 flex items-center gap-3">
									<div class="size-5"></div>
									<span>Name</span>
								</div>
								<div class="col-span-2">Project</div>
								<div class="col-span-2">Size</div>
								<div class="col-span-2">Modified</div>
							</div>
							{#each filteredItems() as item (item.id)}
								<div
									class="group grid grid-cols-12 gap-4 px-4 py-3 items-center border-b border-border last:border-b-0 transition-colors hover:bg-muted/30 {selectedFiles.includes(item.id) ? 'bg-primary/5' : ''}"
									role="button"
									tabindex="0"
									onclick={() => handleFileClick(item)}
									onkeydown={(e) => { if (e.key === 'Enter') handleFileClick(item); }}
									oncontextmenu={(e) => handleContextMenu(e, item)}
								>
									<div class="col-span-6 flex items-center gap-3">
										<button
											class="size-5 rounded border flex items-center justify-center transition-all {selectedFiles.includes(item.id) ? 'bg-primary border-primary text-primary-foreground' : 'border-border hover:border-primary bg-background'}"
											onclick={(e) => toggleSelect(item.id, e)}
											type="button"
										>
											{#if selectedFiles.includes(item.id)}
												<Icon icon="hugeicons:tick-01" class="size-3" />
											{/if}
										</button>
										<div class="size-10 rounded-lg bg-muted flex items-center justify-center shrink-0 {item.type === 'file' ? getFileColor(item.extension) : ''}">
											<Icon icon={item.type === 'folder' ? 'hugeicons:folder-01' : getFileIcon(item.extension)} class="size-5 {item.type === 'folder' ? 'text-amber-500' : ''}" />
										</div>
										<div class="min-w-0">
											<div class="flex items-center gap-2">
												<p class="font-medium truncate">{item.name}</p>
												{#if item.starred}
													<Icon icon="hugeicons:star" class="size-3 text-amber-500 fill-amber-500 shrink-0" />
												{/if}
											</div>
											<p class="text-xs text-muted-foreground">{item.extension?.toUpperCase() || 'Folder'}</p>
										</div>
									</div>
									<div class="col-span-2">
										{#if item.project}
											<div class="flex items-center gap-2">
												<div class="size-2 rounded-full {item.project.color}"></div>
												<span class="text-sm truncate">{item.project.name}</span>
											</div>
										{:else}
											<span class="text-sm text-muted-foreground">-</span>
										{/if}
									</div>
									<div class="col-span-2 text-sm text-muted-foreground">{item.size || '-'}</div>
									<div class="col-span-2 flex items-center gap-2">
										<span class="text-sm text-muted-foreground">{item.updatedAt}</span>
										<button 
											class="opacity-0 group-hover:opacity-100 transition-opacity p-1 hover:bg-muted rounded"
											onclick={(e) => { e.stopPropagation(); toggleStar(e, item); }}
											type="button"
										>
											<Icon icon="hugeicons:star" class="size-4 {item.starred ? 'text-amber-500 fill-amber-500' : 'text-muted-foreground'}" />
										</button>
									</div>
								</div>
							{/each}
						</div>
					{/if}
				</div>
			</div>
		</div>
	</div>
</div>

{#if showContextMenu}
	<div
		class="fixed z-50 min-w-48 rounded-lg border border-border bg-popover shadow-lg py-1"
		style="left: {contextMenuPos.x}px; top: {contextMenuPos.y}px;"
		role="menu"
	>
		<button class="w-full px-3 py-2 text-sm text-left hover:bg-muted flex items-center gap-3 transition-colors">
			<Icon icon="hugeicons:view" class="size-4" />
			Preview
		</button>
		<button class="w-full px-3 py-2 text-sm text-left hover:bg-muted flex items-center gap-3 transition-colors">
			<Icon icon="hugeicons:download-01" class="size-4" />
			Download
		</button>
		<button class="w-full px-3 py-2 text-sm text-left hover:bg-muted flex items-center gap-3 transition-colors">
			<Icon icon="hugeicons:share-01" class="size-4" />
			Share
		</button>
		<button 
			class="w-full px-3 py-2 text-sm text-left hover:bg-muted flex items-center gap-3 transition-colors"
			onclick={(e) => { contextMenuFile && toggleStar(e, contextMenuFile); closeContextMenu(); }}
		>
			<Icon icon="hugeicons:star" class="size-4" />
			{contextMenuFile?.starred ? 'Unstar' : 'Star'}
		</button>
		<Separator class="my-1" />
		<button class="w-full px-3 py-2 text-sm text-left hover:bg-muted flex items-center gap-3 transition-colors">
			<Icon icon="hugeicons:pencil-edit-01" class="size-4" />
			Rename
		</button>
		<button class="w-full px-3 py-2 text-sm text-left hover:bg-muted text-destructive flex items-center gap-3 transition-colors">
			<Icon icon="hugeicons:delete-01" class="size-4" />
			Delete
		</button>
	</div>
{/if}

<Modal
	bind:open={showUploadModal}
	title="Upload Files"
	description="Drag and drop files or click to browse"
	size="lg"
>
	<div 
		class="border-2 border-dashed rounded-xl p-12 text-center transition-all {isDragging ? 'border-primary bg-primary/5' : 'border-border hover:border-primary/50'}"
		ondragover={handleDragOver}
		ondragleave={handleDragLeave}
		ondrop={handleDrop}
	>
		<div class="size-16 rounded-full bg-muted flex items-center justify-center mx-auto mb-4">
			<Icon icon="hugeicons:upload-01" class="size-8 text-muted-foreground" />
		</div>
		<p class="font-medium">Drag and drop files here</p>
		<p class="text-sm text-muted-foreground mt-1">or click to browse from your computer</p>
		<Button class="mt-4" variant="outline" onclick={() => {}}>Browse Files</Button>
	</div>
	
	<div class="mt-4 space-y-2">
		<p class="text-sm font-medium">Upload to</p>
		<select class="w-full h-10 rounded-md border border-input bg-background px-3 py-1 text-sm">
			<option>All Files</option>
			<option>Website Redesign</option>
			<option>Brand Assets</option>
			<option>Mobile App Designs</option>
		</select>
	</div>
	
	{#snippet footer()}
		<Button variant="outline" onclick={() => showUploadModal = false}>Cancel</Button>
		<Button class="gap-2" onclick={handleUpload}>
			<Icon icon="hugeicons:upload-01" class="size-4" />
			Upload
		</Button>
	{/snippet}
</Modal>

<Modal
	bind:open={showNewFolderModal}
	title="Create New Folder"
	description="Enter a name for the new folder"
	size="sm"
>
	<form onsubmit={handleCreateFolder} class="space-y-4">
		<div class="space-y-2">
			<Label for="folder-name">Folder Name</Label>
			<Input id="folder-name" placeholder="New Folder" bind:value={newFolderName} />
		</div>
	</form>
	
	{#snippet footer()}
		<Button variant="outline" onclick={() => showNewFolderModal = false}>Cancel</Button>
		<Button type="submit" onclick={handleCreateFolder} disabled={!newFolderName.trim()}>
			Create Folder
		</Button>
	{/snippet}
</Modal>

<Modal
	bind:open={showPreviewModal}
	title={previewFile?.name || 'Preview'}
	description="File preview"
	size="xl"
>
	{#if previewFile}
		<div class="flex flex-col">
			<div class="flex items-center gap-4 p-4 border-b border-border">
				<div class="size-12 rounded-lg bg-muted flex items-center justify-center {getFileColor(previewFile.extension)}">
					<Icon icon={getFileIcon(previewFile.extension)} class="size-6" />
				</div>
				<div class="flex-1 min-w-0">
					<p class="font-semibold truncate">{previewFile.name}</p>
					<p class="text-sm text-muted-foreground">{previewFile.size} • {previewFile.extension?.toUpperCase()}</p>
				</div>
				<Button variant="outline" size="icon" class="shrink-0" onclick={() => previewFile && toggleStar(new Event('click'), previewFile)}>
					<Icon icon="hugeicons:star" class="size-4 {previewFile.starred ? 'text-amber-500 fill-amber-500' : ''}" />
				</Button>
			</div>
			
			<div class="flex items-center justify-center bg-muted/30 min-h-[400px]">
				{#if previewFile.thumbnail}
					<img src={previewFile.thumbnail} alt={previewFile.name} class="max-w-full max-h-[400px] object-contain rounded-lg shadow-lg" />
				{:else}
					<div class="text-center py-20">
						<Icon icon={getFileIcon(previewFile.extension)} class="size-24 mx-auto mb-4 opacity-50" />
						<p class="text-muted-foreground">Preview not available for this file type</p>
					</div>
				{/if}
			</div>
			
			<div class="flex items-center justify-between p-4 border-t border-border gap-4">
				<div class="flex items-center gap-3">
					<img src={previewFile.uploadedBy.avatar} alt={previewFile.uploadedBy.name} class="size-8 rounded-full" />
					<div class="text-sm">
						<p class="truncate">Uploaded by {previewFile.uploadedBy.name}</p>
						<p class="text-muted-foreground">{previewFile.uploadedAt}</p>
					</div>
				</div>
				<div class="flex gap-2 shrink-0">
					<Button variant="outline" class="gap-2">
						<Icon icon="hugeicons:share-01" class="size-4" />
						Share
					</Button>
					<Button class="gap-2">
						<Icon icon="hugeicons:download-01" class="size-4" />
						Download
					</Button>
				</div>
			</div>
		</div>
	{/if}
	
	{#snippet footer()}
		<Button variant="outline" onclick={() => showPreviewModal = false}>Close</Button>
	{/snippet}
</Modal>
