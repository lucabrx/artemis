<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';
	import { Label } from '$lib/components/ui/label/index.js';

	let viewMode = $state<'day' | 'week' | 'month'>('week');
	let showNewEntryModal = $state(false);
	let runningTimer = $state(false);
	let elapsedSeconds = $state(0);
	let timerInterval: ReturnType<typeof setInterval> | null = null;
	let selectedEntry = $state<TimeEntry | null>(null);
	let showEntryDetails = $state(false);

	interface TimeEntry {
		id: string;
		description: string;
		project: {
			id: string;
			name: string;
			color: string;
			client: string;
		};
		task?: {
			id: string;
			title: string;
		};
		startTime: string;
		endTime?: string;
		duration: number;
		billable: boolean;
		billed: boolean;
		date: string;
		tags: string[];
	}

	const projects = [
		{ id: '1', name: 'Website Redesign', client: 'Acme Corp', color: 'bg-blue-500' },
		{ id: '2', name: 'Mobile App', client: 'TechStart', color: 'bg-violet-500' },
		{ id: '3', name: 'Brand Identity', client: 'Green Leaf', color: 'bg-emerald-500' },
		{ id: '4', name: 'Marketing Campaign', client: 'Global Dynamics', color: 'bg-amber-500' },
		{ id: '5', name: 'E-commerce Platform', client: 'Retail Plus', color: 'bg-rose-500' }
	];

	const tasks = [
		{ id: 't1', title: 'Design system setup' },
		{ id: 't2', title: 'Homepage mockups' },
		{ id: 't3', title: 'API architecture' },
		{ id: 't4', title: 'Logo concepts' },
		{ id: 't5', title: 'Content calendar' }
	];

	const timeEntries: TimeEntry[] = [
		{
			id: 'e1',
			description: 'Design system color palette refinement',
			project: projects[0],
			task: tasks[0],
			startTime: '09:00 AM',
			endTime: '11:30 AM',
			duration: 9000,
			billable: true,
			billed: false,
			date: 'Today',
			tags: ['Design', 'UI']
		},
		{
			id: 'e2',
			description: 'Client meeting - design review',
			project: projects[0],
			startTime: '02:00 PM',
			endTime: '03:00 PM',
			duration: 3600,
			billable: true,
			billed: false,
			date: 'Today',
			tags: ['Meeting']
		},
		{
			id: 'e3',
			description: 'Mobile app wireframes',
			project: projects[1],
			task: tasks[2],
			startTime: '10:00 AM',
			endTime: '01:00 PM',
			duration: 10800,
			billable: true,
			billed: true,
			date: 'Yesterday',
			tags: ['Design', 'Mobile']
		},
		{
			id: 'e4',
			description: 'Brand guidelines documentation',
			project: projects[2],
			task: tasks[3],
			startTime: '03:00 PM',
			endTime: '05:30 PM',
			duration: 9000,
			billable: true,
			billed: false,
			date: 'Yesterday',
			tags: ['Documentation']
		},
		{
			id: 'e5',
			description: 'Internal team sync',
			project: projects[0],
			startTime: '09:00 AM',
			endTime: '09:30 AM',
			duration: 1800,
			billable: false,
			billed: false,
			date: 'Yesterday',
			tags: ['Internal']
		},
		{
			id: 'e6',
			description: 'Marketing campaign assets',
			project: projects[3],
			task: tasks[4],
			startTime: '11:00 AM',
			endTime: '03:00 PM',
			duration: 14400,
			billable: true,
			billed: true,
			date: 'Jan 26',
			tags: ['Design', 'Marketing']
		},
		{
			id: 'e7',
			description: 'Code review and bug fixes',
			project: projects[4],
			startTime: '09:00 AM',
			endTime: '12:00 PM',
			duration: 10800,
			billable: true,
			billed: false,
			date: 'Jan 25',
			tags: ['Development']
		},
		{
			id: 'e8',
			description: 'Research and planning',
			project: projects[1],
			startTime: '02:00 PM',
			endTime: '04:00 PM',
			duration: 7200,
			billable: false,
			billed: false,
			date: 'Jan 25',
			tags: ['Research']
		}
	];

	const weeklyData = [
		{ day: 'Mon', date: 'Jan 20', hours: 6.5, billable: 5.5 },
		{ day: 'Tue', date: 'Jan 21', hours: 8.0, billable: 7.0 },
		{ day: 'Wed', date: 'Jan 22', hours: 7.5, billable: 6.5 },
		{ day: 'Thu', date: 'Jan 23', hours: 8.5, billable: 8.0 },
		{ day: 'Fri', date: 'Jan 24', hours: 6.0, billable: 5.0 },
		{ day: 'Sat', date: 'Jan 25', hours: 4.0, billable: 3.0 },
		{ day: 'Sun', date: 'Jan 26', hours: 0, billable: 0 },
		{ day: 'Mon', date: 'Jan 27', hours: 5.5, billable: 4.5 },
		{ day: 'Tue', date: 'Jan 28', hours: 7.0, billable: 6.0 }
	];

	const projectBreakdown = [
		{ project: projects[0], hours: 24.5, percentage: 45, amount: 2940 },
		{ project: projects[1], hours: 12.0, percentage: 22, amount: 1440 },
		{ project: projects[2], hours: 8.5, percentage: 16, amount: 1020 },
		{ project: projects[3], hours: 6.0, percentage: 11, amount: 720 },
		{ project: projects[4], hours: 3.0, percentage: 6, amount: 360 }
	];

	let newEntry = $state({
		description: '',
		project: '',
		task: '',
		date: new Date().toISOString().split('T')[0],
		startTime: '',
		endTime: '',
		duration: '',
		billable: true,
		tags: ''
	});

	function formatDuration(seconds: number): string {
		const hours = Math.floor(seconds / 3600);
		const minutes = Math.floor((seconds % 3600) / 60);
		const secs = seconds % 60;
		return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`;
	}

	function formatHours(seconds: number): string {
		const hours = Math.floor(seconds / 3600);
		const minutes = Math.floor((seconds % 3600) / 60);
		return `${hours}h ${minutes}m`;
	}

	function toggleTimer() {
		if (runningTimer) {
			runningTimer = false;
			if (timerInterval) clearInterval(timerInterval);
		} else {
			runningTimer = true;
			timerInterval = setInterval(() => {
				elapsedSeconds++;
			}, 1000);
		}
	}

	function resetTimer() {
		runningTimer = false;
		if (timerInterval) clearInterval(timerInterval);
		elapsedSeconds = 0;
	}

	function handleCreateEntry(e: Event) {
		e.preventDefault();
		console.log('Creating entry:', newEntry);
		newEntry = {
			description: '',
			project: '',
			task: '',
			date: new Date().toISOString().split('T')[0],
			startTime: '',
			endTime: '',
			duration: '',
			billable: true,
			tags: ''
		};
		showNewEntryModal = false;
	}

	const totalThisWeek = weeklyData.reduce((acc, day) => acc + day.hours, 0);
	const billableThisWeek = weeklyData.reduce((acc, day) => acc + day.billable, 0);
	const totalAmount = projectBreakdown.reduce((acc, p) => acc + p.amount, 0);

	const groupedEntries = $derived(() => {
		const groups: Record<string, TimeEntry[]> = {};
		timeEntries.forEach(entry => {
			if (!groups[entry.date]) groups[entry.date] = [];
			groups[entry.date].push(entry);
		});
		return groups;
	});
</script>

<svelte:head>
	<title>Time Tracking - Artemis</title>
</svelte:head>

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="px-6 py-4">
		<div class="mb-6 flex items-center justify-end gap-2">
			<Button variant="outline" class="gap-2">
				<Icon icon="hugeicons:download-01" class="size-4" />
				Export
			</Button>
			<Button class="gap-2" onclick={() => showNewEntryModal = true}>
				<Icon icon="hugeicons:plus-sign" class="size-4" />
				Add Entry
			</Button>
		</div>

		<div class="grid gap-6 lg:grid-cols-3">
			<div class="space-y-6 lg:col-span-2">
				<div class="rounded-2xl border border-border bg-card p-6">
					<div class="mb-6 flex items-center justify-between">
						<div>
							<h2 class="text-lg font-semibold">Timer</h2>
							<p class="text-sm text-muted-foreground">Track your work in real-time</p>
						</div>
						<div class="flex items-center gap-2">
							<div class="size-2 rounded-full {runningTimer ? 'animate-pulse bg-emerald-500' : 'bg-slate-400'}"></div>
							<span class="text-sm text-muted-foreground">{runningTimer ? 'Running' : 'Stopped'}</span>
						</div>
					</div>

					<div class="mb-6 text-center">
						<div class="font-mono text-6xl font-bold tracking-tight">
							{formatDuration(elapsedSeconds)}
						</div>
					</div>

					<div class="mb-6 grid gap-4 md:grid-cols-3">
						<div>
							<Label for="timer-project">Project</Label>
							<select id="timer-project" class="mt-1.5 h-10 w-full rounded-md border border-input bg-background px-3 text-sm">
								<option value="">Select project...</option>
								{#each projects as project}
									<option value={project.id}>{project.name}</option>
								{/each}
							</select>
						</div>
						<div>
							<Label for="timer-task">Task</Label>
							<select id="timer-task" class="mt-1.5 h-10 w-full rounded-md border border-input bg-background px-3 text-sm">
								<option value="">Select task...</option>
								{#each tasks as task}
									<option value={task.id}>{task.title}</option>
								{/each}
							</select>
						</div>
						<div>
							<Label for="timer-desc">Description</Label>
							<Input id="timer-desc" placeholder="What are you working on?" class="mt-1.5" />
						</div>
					</div>

					<div class="flex items-center justify-between">
						<label class="flex items-center gap-2">
							<input type="checkbox" checked class="size-4 rounded border-input" />
							<span class="text-sm">Billable</span>
						</label>
						<div class="flex gap-2">
							{#if elapsedSeconds > 0}
								<Button variant="outline" onclick={resetTimer}>
									<Icon icon="hugeicons:cancel-01" class="mr-2 size-4" />
									Reset
								</Button>
							{/if}
							<Button class="gap-2" variant={runningTimer ? 'secondary' : 'default'} onclick={toggleTimer}>
								{#if runningTimer}
									<Icon icon="hugeicons:pause" class="size-4" />
									Stop
								{:else}
									<Icon icon="hugeicons:play" class="size-4" />
									Start Timer
								{/if}
							</Button>
						</div>
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-6">
					<div class="mb-4 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
						<h2 class="text-lg font-semibold">Time Entries</h2>
						<div class="flex rounded-lg border border-border p-1">
							{#each [{id: 'day', label: 'Day'}, {id: 'week', label: 'Week'}, {id: 'month', label: 'Month'}] as mode}
								<button
									class="rounded px-3 py-1 text-sm font-medium transition-colors {viewMode === mode.id ? 'bg-primary text-primary-foreground' : 'text-muted-foreground hover:text-foreground'}"
									onclick={() => viewMode = mode.id as typeof viewMode}
								>
									{mode.label}
								</button>
							{/each}
						</div>
					</div>

					<div class="space-y-6">
						{#each Object.entries(groupedEntries()) as [date, entries]}
							{@const dayTotal = entries.reduce((acc, e) => acc + e.duration, 0)}
							{@const dayBillable = entries.filter(e => e.billable).reduce((acc, e) => acc + e.duration, 0)}
							<div>
								<div class="mb-3 flex items-center justify-between">
									<div class="flex items-center gap-2">
										<h3 class="font-semibold">{date}</h3>
										<span class="text-sm text-muted-foreground">{entries.length} entries</span>
									</div>
									<div class="text-right text-sm">
										<span class="font-medium">{formatHours(dayTotal)}</span>
										<span class="text-muted-foreground">({formatHours(dayBillable)} billable)</span>
									</div>
								</div>

								<div class="space-y-2">
									{#each entries as entry}
										<div class="group flex items-center gap-4 rounded-xl border border-border bg-muted/30 p-3 transition-all hover:border-primary/20 hover:bg-muted/50">
											<div class="size-2 shrink-0 rounded-full {entry.project.color}"></div>
											<div class="min-w-0 flex-1">
												<p class="truncate font-medium">{entry.description}</p>
												<div class="flex flex-wrap items-center gap-2 text-sm text-muted-foreground">
													<span>{entry.project.name}</span>
													{#if entry.task}
														<span>·</span>
														<span>{entry.task.title}</span>
													{/if}
													<span>·</span>
													<span>{entry.startTime} - {entry.endTime || 'Running'}</span>
												</div>
												<div class="mt-1 flex flex-wrap gap-1">
													{#each entry.tags as tag}
														<span class="rounded bg-muted px-1.5 py-0.5 text-[10px]">{tag}</span>
													{/each}
													{#if entry.billable}
														<span class="rounded bg-emerald-500/10 px-1.5 py-0.5 text-[10px] text-emerald-600">
															<Icon icon="hugeicons:dollar-circle" class="inline size-3" />
															Billable
														</span>
													{/if}
													{#if entry.billed}
														<span class="rounded bg-blue-500/10 px-1.5 py-0.5 text-[10px] text-blue-600">
															<Icon icon="hugeicons:invoice-01" class="inline size-3" />
															Billed
														</span>
													{/if}
												</div>
											</div>
											<div class="hidden shrink-0 text-right sm:block">
												<p class="font-mono font-medium">{formatHours(entry.duration)}</p>
												<p class="text-xs text-muted-foreground">${(entry.duration / 3600 * 120).toFixed(0)}</p>
											</div>
											<Button
												variant="ghost"
												size="icon"
												class="shrink-0 opacity-0 transition-opacity group-hover:opacity-100"
												onclick={() => { selectedEntry = entry; showEntryDetails = true; }}
											>
												<Icon icon="hugeicons:more-vertical" class="size-4" />
											</Button>
										</div>
									{/each}
								</div>
							</div>
						{/each}
					</div>
				</div>
			</div>

			<div class="space-y-6">
				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">This Week</h3>
					<div class="space-y-4">
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Total Hours</span>
							<span class="text-2xl font-bold">{totalThisWeek.toFixed(1)}h</span>
						</div>
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Billable</span>
							<span class="text-lg font-semibold text-emerald-600">{billableThisWeek.toFixed(1)}h</span>
						</div>
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Non-Billable</span>
							<span class="text-sm font-medium">{(totalThisWeek - billableThisWeek).toFixed(1)}h</span>
						</div>
						<div class="h-2 overflow-hidden rounded-full bg-muted">
							<div class="h-full rounded-full bg-emerald-500 transition-all" style="width: {(billableThisWeek / totalThisWeek) * 100}%"></div>
						</div>
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Weekly Overview</h3>
					<div class="space-y-3">
						{#each weeklyData.slice(-7) as day}
							<div class="flex items-center gap-3">
								<div class="w-8 shrink-0 text-xs text-muted-foreground">{day.day}</div>
								<div class="flex-1">
									<div class="h-6 overflow-hidden rounded-md bg-muted">
										<div class="flex h-full">
											<div class="h-full bg-emerald-500" style="width: {(day.billable / 10) * 100}%"></div>
											<div class="h-full bg-slate-300" style="width: {((day.hours - day.billable) / 10) * 100}%"></div>
										</div>
									</div>
								</div>
								<div class="w-12 shrink-0 text-right text-xs font-medium">{day.hours}h</div>
							</div>
						{/each}
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Project Breakdown</h3>
					<div class="space-y-4">
						{#each projectBreakdown as item}
							<div>
								<div class="mb-1 flex items-center justify-between">
									<div class="flex items-center gap-2">
										<div class="size-2 rounded-full {item.project.color}"></div>
										<span class="text-sm">{item.project.name}</span>
									</div>
									<span class="text-sm font-medium">{item.hours}h</span>
								</div>
								<div class="flex items-center justify-between text-xs text-muted-foreground">
									<span>{item.percentage}%</span>
									<span>${item.amount}</span>
								</div>
								<div class="mt-1 h-1.5 overflow-hidden rounded-full bg-muted">
									<div class="h-full rounded-full {item.project.color}" style="width: {item.percentage}%"></div>
								</div>
							</div>
						{/each}
					</div>
					<Separator class="my-4" />
					<div class="flex items-center justify-between">
						<span class="font-semibold">Total Value</span>
						<span class="text-lg font-bold">${totalAmount.toLocaleString()}</span>
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-gradient-to-br from-primary/5 to-primary/10 p-5">
					<div class="flex items-start gap-3">
						<div class="flex size-10 items-center justify-center rounded-xl bg-primary/10">
							<Icon icon="hugeicons:bulb" class="size-5 text-primary" />
						</div>
						<div>
							<h3 class="font-semibold">Pro Tip</h3>
							<p class="mt-1 text-sm text-muted-foreground">
								Use keyboard shortcut <kbd class="rounded bg-muted px-1.5 py-0.5 text-xs">Space</kbd> to quickly start/stop the timer.
							</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<Modal
	bind:open={showNewEntryModal}
	title="Add Time Entry"
	description="Manually add a time entry."
	size="lg"
>
	<form id="new-entry-form" onsubmit={handleCreateEntry} class="space-y-4">
		<div class="space-y-2">
			<Label for="entry-description">Description *</Label>
			<Input id="entry-description" placeholder="What did you work on?" bind:value={newEntry.description} required />
		</div>

		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="entry-project">Project *</Label>
				<select
					id="entry-project"
					bind:value={newEntry.project}
					required
					class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-sm transition-colors"
				>
					<option value="">Select project...</option>
					{#each projects as project}
						<option value={project.id}>{project.name}</option>
					{/each}
				</select>
			</div>
			<div class="space-y-2">
				<Label for="entry-task">Task</Label>
				<select
					id="entry-task"
					bind:value={newEntry.task}
					class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-sm transition-colors"
				>
					<option value="">Select task...</option>
					{#each tasks as task}
						<option value={task.id}>{task.title}</option>
					{/each}
				</select>
			</div>
		</div>

		<div class="grid gap-4 md:grid-cols-3">
			<div class="space-y-2">
				<Label for="entry-date">Date *</Label>
				<Input id="entry-date" type="date" bind:value={newEntry.date} required />
			</div>
			<div class="space-y-2">
				<Label for="entry-start">Start Time</Label>
				<Input id="entry-start" type="time" bind:value={newEntry.startTime} />
			</div>
			<div class="space-y-2">
				<Label for="entry-end">End Time</Label>
				<Input id="entry-end" type="time" bind:value={newEntry.endTime} />
			</div>
		</div>

		<div class="space-y-2">
			<Label for="entry-duration">Duration (HH:MM)</Label>
			<Input id="entry-duration" placeholder="01:30" bind:value={newEntry.duration} />
		</div>

		<div class="space-y-2">
			<Label for="entry-tags">Tags</Label>
			<Input id="entry-tags" placeholder="Design, Meeting, Development..." bind:value={newEntry.tags} />
		</div>

		<label class="flex items-center gap-2">
			<input type="checkbox" bind:checked={newEntry.billable} class="size-4 rounded border-input" />
			<span class="text-sm">Billable</span>
		</label>
	</form>

	{#snippet footer()}
		<Button variant="outline" onclick={() => showNewEntryModal = false}>Cancel</Button>
		<Button type="submit" form="new-entry-form">
			<Icon icon="hugeicons:plus-sign" class="mr-2 size-4" />
			Add Entry
		</Button>
	{/snippet}
</Modal>

<Modal
	bind:open={showEntryDetails}
	title="Time Entry Details"
	description="View and edit time entry."
	size="md"
>
	{#if selectedEntry}
		<div class="space-y-4">
			<div class="flex items-center gap-3">
				<div class="size-3 rounded-full {selectedEntry.project.color}"></div>
				<span class="font-medium">{selectedEntry.project.name}</span>
			</div>
			<p class="text-lg">{selectedEntry.description}</p>
			<div class="grid grid-cols-2 gap-4 rounded-xl bg-muted p-4">
				<div>
					<p class="text-xs text-muted-foreground">Duration</p>
					<p class="font-mono text-lg font-semibold">{formatHours(selectedEntry.duration)}</p>
				</div>
				<div>
					<p class="text-xs text-muted-foreground">Amount</p>
					<p class="text-lg font-semibold">${(selectedEntry.duration / 3600 * 120).toFixed(2)}</p>
				</div>
			</div>
		</div>
	{/if}

	{#snippet footer()}
		<Button variant="outline" onclick={() => showEntryDetails = false}>Close</Button>
		<Button variant="destructive">
			<Icon icon="hugeicons:delete-01" class="mr-2 size-4" />
			Delete
		</Button>
	{/snippet}
</Modal>
