<script lang="ts">
	import { page } from '$app/state';
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';
	import { Label } from '$lib/components/ui/label/index.js';

	let showDeleteModal = $state(false);
	let showEditModal = $state(false);
	let newComment = $state('');
	let replyToComment: string | null = $state(null);
	let replyText = $state('');

	const taskId = page.params.id;

	const projects = [
		{ id: '1', name: 'Website Redesign', client: 'Acme Corp', color: 'bg-blue-500' },
		{ id: '2', name: 'Mobile App', client: 'TechStart', color: 'bg-violet-500' },
		{ id: '3', name: 'Brand Identity', client: 'Green Leaf', color: 'bg-emerald-500' },
		{ id: '4', name: 'Marketing Campaign', client: 'Global Dynamics', color: 'bg-amber-500' },
		{ id: '5', name: 'E-commerce Platform', client: 'Retail Plus', color: 'bg-rose-500' }
	];

	const teamMembers = [
		{ name: 'Sarah Chen', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah', initials: 'SC' },
		{ name: 'Mike Ross', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Mike', initials: 'MR' },
		{ name: 'David Kim', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=David', initials: 'DK' },
		{ name: 'Lisa Wang', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Lisa', initials: 'LW' },
		{ name: 'Emma Wilson', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Emma', initials: 'EW' },
		{ name: 'Jane Doe', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Jane', initials: 'JD' }
	];

	const task = {
		id: taskId,
		title: 'Design system setup',
		description: 'Create the foundation design system with colors, typography, and component library. This includes defining the color palette, typography scale, spacing system, and creating reusable components that will be used across all pages of the website.\n\nThe design system should follow modern best practices and ensure consistency throughout the entire project. Make sure to document all decisions and create guidelines for future reference.',
		status: 'done' as const,
		priority: 'high' as const,
		project: projects[0],
		assignee: teamMembers[0],
		reporter: teamMembers[1],
		dueDate: 'Jan 10, 2026',
		createdAt: 'Jan 5, 2026',
		updatedAt: 'Jan 10, 2026',
		completedAt: 'Jan 10, 2026',
		estimatedHours: 16,
		loggedHours: 14,
		subtasks: [
			{ id: 's1', title: 'Define color palette with primary, secondary, and neutral colors', completed: true, assignee: teamMembers[0] },
			{ id: 's2', title: 'Create typography scale with headings and body text', completed: true, assignee: teamMembers[0] },
			{ id: 's3', title: 'Build component library in Figma with variants', completed: true, assignee: teamMembers[0] },
			{ id: 's4', title: 'Document usage guidelines and best practices', completed: true, assignee: teamMembers[0] }
		],
		tags: ['Design', 'UI/UX', 'Figma', 'Foundation'],
		attachments: [
			{ 
				id: 'a1', 
				name: 'Design-System-v1.fig', 
				size: '12.4 MB', 
				type: 'figma',
				uploadedBy: teamMembers[0],
				uploadedAt: 'Jan 8, 2026',
				thumbnail: 'https://placehold.co/400x300/f3f4f6/666?text=Figma+File'
			},
			{ 
				id: 'a2', 
				name: 'Color-Palette.pdf', 
				size: '2.1 MB', 
				type: 'pdf',
				uploadedBy: teamMembers[0],
				uploadedAt: 'Jan 7, 2026',
				thumbnail: null
			},
			{ 
				id: 'a3', 
				name: 'Typography-Scale.png', 
				size: '856 KB', 
				type: 'image',
				uploadedBy: teamMembers[0],
				uploadedAt: 'Jan 7, 2026',
				thumbnail: 'https://placehold.co/400x300/e0e7ff/4f46e5?text=Typography'
			},
			{ 
				id: 'a4', 
				name: 'Component-Library.fig', 
				size: '8.7 MB', 
				type: 'figma',
				uploadedBy: teamMembers[0],
				uploadedAt: 'Jan 9, 2026',
				thumbnail: 'https://placehold.co/400x300/f3f4f6/666?text=Components'
			}
		],
		comments: [
			{
				id: 'c1',
				author: teamMembers[1],
				text: 'Great work on the color palette! The primary blue really captures the brand essence. I think we should consider adding a darker shade for better contrast on light backgrounds.',
				createdAt: 'Jan 8, 2026 at 2:30 PM',
				likes: 3,
				replies: [
					{
						id: 'r1',
						author: teamMembers[0],
						text: 'Thanks Mike! I added a 900 shade to the palette. Check the updated file.',
						createdAt: 'Jan 8, 2026 at 3:15 PM',
						likes: 1
					}
				]
			},
			{
				id: 'c2',
				author: teamMembers[3],
				text: 'The typography scale looks good. Are we using Inter for the body text or should we consider something else?',
				createdAt: 'Jan 9, 2026 at 10:20 AM',
				likes: 2,
				replies: []
			},
			{
				id: 'c3',
				author: teamMembers[4],
				text: 'Component library is comprehensive! I noticed the button component is missing a loading state. Could you add that before we mark this complete?',
				createdAt: 'Jan 9, 2026 at 4:45 PM',
				likes: 1,
				replies: [
					{
						id: 'r2',
						author: teamMembers[0],
						text: 'Good catch! Added the loading state with spinner variant.',
						createdAt: 'Jan 10, 2026 at 9:00 AM',
						likes: 2
					},
					{
						id: 'r3',
						author: teamMembers[4],
						text: 'Perfect, thanks!',
						createdAt: 'Jan 10, 2026 at 9:30 AM',
						likes: 0
					}
				]
			}
		],
		activity: [
			{ id: 'act1', type: 'created', user: teamMembers[1], timestamp: 'Jan 5, 2026 at 9:00 AM', text: 'created this task' },
			{ id: 'act2', type: 'assigned', user: teamMembers[1], timestamp: 'Jan 5, 2026 at 9:05 AM', text: 'assigned to Sarah Chen' },
			{ id: 'act3', type: 'status', user: teamMembers[0], timestamp: 'Jan 5, 2026 at 2:00 PM', text: 'changed status to In Progress' },
			{ id: 'act4', type: 'attachment', user: teamMembers[0], timestamp: 'Jan 7, 2026 at 11:30 AM', text: 'attached Color-Palette.pdf' },
			{ id: 'act5', type: 'attachment', user: teamMembers[0], timestamp: 'Jan 8, 2026 at 10:00 AM', text: 'attached Design-System-v1.fig' },
			{ id: 'act6', type: 'comment', user: teamMembers[1], timestamp: 'Jan 8, 2026 at 2:30 PM', text: 'added a comment' },
			{ id: 'act7', type: 'subtask', user: teamMembers[0], timestamp: 'Jan 10, 2026 at 3:00 PM', text: 'completed all subtasks' },
			{ id: 'act8', type: 'status', user: teamMembers[0], timestamp: 'Jan 10, 2026 at 4:00 PM', text: 'changed status to Done' }
		],
		relatedTasks: [
			{ id: 't2', title: 'Homepage mockups', status: 'done', project: projects[0] },
			{ id: 't3', title: 'About page development', status: 'in-progress', project: projects[0] },
			{ id: 't5', title: 'Mobile responsiveness review', status: 'review', project: projects[0] }
		],
		watchers: [teamMembers[1], teamMembers[3], teamMembers[4]]
	};

	const kanbanColumns = [
		{ id: 'todo', title: 'To Do', color: 'bg-slate-500', icon: 'hugeicons:circle' },
		{ id: 'in-progress', title: 'In Progress', color: 'bg-blue-500', icon: 'hugeicons:loading-02' },
		{ id: 'review', title: 'Review', color: 'bg-amber-500', icon: 'hugeicons:search-01' },
		{ id: 'done', title: 'Done', color: 'bg-emerald-500', icon: 'hugeicons:tick-01' }
	];

	function getStatusColor(status: string): string {
		switch (status) {
			case 'todo':
				return 'bg-slate-500/10 text-slate-600 border-slate-500/20';
			case 'in-progress':
				return 'bg-blue-500/10 text-blue-600 border-blue-500/20';
			case 'review':
				return 'bg-amber-500/10 text-amber-600 border-amber-500/20';
			case 'done':
				return 'bg-emerald-500/10 text-emerald-600 border-emerald-500/20';
			default:
				return 'bg-muted';
		}
	}

	function getPriorityColor(priority: string): string {
		switch (priority) {
			case 'urgent':
				return 'bg-rose-500 text-white';
			case 'high':
				return 'bg-orange-500 text-white';
			case 'medium':
				return 'bg-amber-400 text-white';
			case 'low':
				return 'bg-blue-400 text-white';
			default:
				return 'bg-muted';
		}
	}

	function getPriorityIcon(priority: string): string {
		switch (priority) {
			case 'urgent':
				return 'hugeicons:alert-02';
			case 'high':
				return 'hugeicons:arrow-up-02';
			case 'medium':
				return 'hugeicons:arrow-right-02';
			case 'low':
				return 'hugeicons:arrow-down-02';
			default:
				return 'hugeicons:minus-sign';
		}
	}

	function getFileIcon(type: string): string {
		switch (type) {
			case 'pdf':
				return 'hugeicons:pdf-01';
			case 'image':
				return 'hugeicons:image-01';
			case 'figma':
				return 'hugeicons:figma';
			default:
				return 'hugeicons:file-01';
		}
	}

	function getActivityIcon(type: string): string {
		switch (type) {
			case 'created':
				return 'hugeicons:plus-sign';
			case 'assigned':
				return 'hugeicons:user-add-01';
			case 'status':
				return 'hugeicons:loading-02';
			case 'attachment':
				return 'hugeicons:attachment-02';
			case 'comment':
				return 'hugeicons:message-01';
			case 'subtask':
				return 'hugeicons:checkmark-square-02';
			default:
				return 'hugeicons:activity-01';
		}
	}

	function handleAddComment() {
		if (!newComment.trim()) return;
		console.log('Adding comment:', newComment);
		newComment = '';
	}

	function handleReply(commentId: string) {
		if (!replyText.trim()) return;
		console.log('Replying to comment', commentId, ':', replyText);
		replyText = '';
		replyToComment = null;
	}

	const completedSubtasks = task.subtasks.filter(s => s.completed).length;
	const subtasksProgress = (completedSubtasks / task.subtasks.length) * 100;
</script>

<svelte:head>
	<title>{task.title} - Task - Artemis</title>
</svelte:head>

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="px-6 py-8">
		<div class="mb-6 flex flex-col gap-4 lg:flex-row lg:items-start lg:justify-between">
			<div class="flex items-start gap-4">
				<Button variant="outline" size="icon" href="/tasks">
					<Icon icon="hugeicons:arrow-left-01" class="size-5" />
				</Button>
				<div>
					<div class="mb-2 flex items-center gap-2">
						<div class="size-2 rounded-full {task.project.color}"></div>
						<span class="text-sm text-muted-foreground">{task.project.name}</span>
						<span class="text-muted-foreground">·</span>
						<span class="text-sm text-muted-foreground">{task.id}</span>
					</div>
					<h1 class="text-2xl font-bold">{task.title}</h1>
				</div>
			</div>
			<div class="flex flex-wrap gap-2">
				<Button variant="outline" class="gap-2">
					<Icon icon="hugeicons:clock-01" class="size-4" />
					Track Time
				</Button>
				<Button variant="outline" class="gap-2">
					<Icon icon="hugeicons:link-01" class="size-4" />
					Copy Link
				</Button>
				<Button variant="outline" class="gap-2" onclick={() => showEditModal = true}>
					<Icon icon="hugeicons:pencil-edit-01" class="size-4" />
					Edit
				</Button>
				<Button variant="outline" class="gap-2 text-destructive hover:text-destructive" onclick={() => showDeleteModal = true}>
					<Icon icon="hugeicons:delete-01" class="size-4" />
					Delete
				</Button>
			</div>
		</div>

		<div class="grid gap-6 lg:grid-cols-3">
			<div class="space-y-6 lg:col-span-2">
				<div class="rounded-2xl border border-border bg-card p-6">
					<div class="mb-4 flex flex-wrap items-center gap-3">
						<span class="flex items-center gap-1 rounded-full border px-3 py-1 text-sm font-medium {getStatusColor(task.status)}">
							<Icon icon={kanbanColumns.find(c => c.id === task.status)?.icon || 'hugeicons:circle'} class="size-4" />
							{task.status.replace('-', ' ')}
						</span>
						<span class="flex items-center gap-1 rounded-full px-3 py-1 text-sm font-medium {getPriorityColor(task.priority)}">
							<Icon icon={getPriorityIcon(task.priority)} class="size-4" />
							{task.priority}
						</span>
						<div class="flex gap-1">
							{#each task.tags as tag}
								<span class="rounded-md bg-muted px-2 py-1 text-xs">{tag}</span>
							{/each}
						</div>
					</div>

					<div class="prose prose-sm max-w-none">
						<p class="whitespace-pre-wrap text-muted-foreground">{task.description}</p>
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-6">
					<div class="mb-4 flex items-center justify-between">
						<h3 class="font-semibold">Subtasks</h3>
						<span class="text-sm text-muted-foreground">{completedSubtasks} of {task.subtasks.length}</span>
					</div>
					<div class="mb-4 h-2 overflow-hidden rounded-full bg-muted">
						<div class="h-full rounded-full bg-primary transition-all" style="width: {subtasksProgress}%"></div>
					</div>
					<div class="space-y-2">
						{#each task.subtasks as subtask}
							<label class="flex cursor-pointer items-start gap-3 rounded-lg p-2 transition-colors hover:bg-muted">
								<input type="checkbox" checked={subtask.completed} class="mt-0.5 size-4 rounded border-input" />
								<div class="flex-1">
									<span class="text-sm {subtask.completed ? 'text-muted-foreground line-through' : ''}">
										{subtask.title}
									</span>
									<div class="mt-1 flex items-center gap-2">
										<img src={subtask.assignee.avatar} alt={subtask.assignee.name} class="size-5 rounded-full" />
										<span class="text-xs text-muted-foreground">{subtask.assignee.name}</span>
									</div>
								</div>
							</label>
						{/each}
					</div>
					<Button variant="outline" class="mt-4 w-full gap-2">
						<Icon icon="hugeicons:plus-sign" class="size-4" />
						Add Subtask
					</Button>
				</div>

				<div class="rounded-2xl border border-border bg-card p-6">
					<h3 class="mb-4 font-semibold">Attachments ({task.attachments.length})</h3>
					<div class="grid gap-4 sm:grid-cols-2">
						{#each task.attachments as attachment}
							<div class="group relative overflow-hidden rounded-xl border border-border bg-muted/30 p-3 transition-all hover:border-primary/20">
								{#if attachment.thumbnail}
									<div class="mb-3 aspect-video overflow-hidden rounded-lg bg-muted">
										<img src={attachment.thumbnail} alt={attachment.name} class="size-full object-cover" />
									</div>
								{:else}
									<div class="mb-3 flex aspect-video items-center justify-center rounded-lg bg-muted">
										<Icon icon={getFileIcon(attachment.type)} class="size-12 text-muted-foreground" />
									</div>
								{/if}
								<div class="flex items-start justify-between">
									<div class="min-w-0 flex-1">
										<p class="truncate text-sm font-medium">{attachment.name}</p>
										<p class="text-xs text-muted-foreground">{attachment.size}</p>
									</div>
									<Button variant="ghost" size="icon" class="size-7 opacity-0 transition-opacity group-hover:opacity-100">
										<Icon icon="hugeicons:download-01" class="size-4" />
									</Button>
								</div>
								<div class="mt-2 flex items-center gap-2 text-xs text-muted-foreground">
									<img src={attachment.uploadedBy.avatar} alt={attachment.uploadedBy.name} class="size-4 rounded-full" />
									<span>{attachment.uploadedBy.name}</span>
									<span>·</span>
									<span>{attachment.uploadedAt}</span>
								</div>
							</div>
						{/each}
					</div>
					<Button variant="outline" class="mt-4 w-full gap-2">
						<Icon icon="hugeicons:upload-01" class="size-4" />
						Upload File
					</Button>
				</div>

				<div class="rounded-2xl border border-border bg-card p-6">
					<h3 class="mb-4 font-semibold">Comments ({task.comments.length})</h3>
					
					<div class="mb-6 flex gap-3">
						<img src={teamMembers[0].avatar} alt="You" class="size-9 rounded-full" />
						<div class="flex-1">
							<textarea
								placeholder="Add a comment..."
								bind:value={newComment}
								rows={3}
								class="w-full resize-none rounded-lg border border-input bg-background px-3 py-2 text-sm transition-colors placeholder:text-muted-foreground focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-none"
							></textarea>
							<div class="mt-2 flex justify-end">
								<Button onclick={handleAddComment} disabled={!newComment.trim()}>
									<Icon icon="hugeicons:send-01" class="mr-2 size-4" />
									Comment
								</Button>
							</div>
						</div>
					</div>

					<div class="space-y-6">
						{#each task.comments as comment}
							<div class="flex gap-3">
								<img src={comment.author.avatar} alt={comment.author.name} class="size-9 rounded-full" />
								<div class="flex-1">
									<div class="rounded-xl bg-muted p-3">
										<div class="mb-1 flex items-center gap-2">
											<span class="font-medium">{comment.author.name}</span>
											<span class="text-xs text-muted-foreground">{comment.createdAt}</span>
										</div>
										<p class="text-sm">{comment.text}</p>
									</div>
									<div class="mt-1 flex items-center gap-4">
										<button class="flex items-center gap-1 text-xs text-muted-foreground hover:text-foreground">
											<Icon icon="hugeicons:thumbs-up" class="size-3" />
											{comment.likes} likes
										</button>
										<button 
											class="flex items-center gap-1 text-xs text-muted-foreground hover:text-foreground"
											onclick={() => replyToComment = replyToComment === comment.id ? null : comment.id}
										>
											<Icon icon="hugeicons:reply" class="size-3" />
											Reply
										</button>
									</div>

									{#if comment.replies.length > 0}
										<div class="mt-3 space-y-3">
											{#each comment.replies as reply}
												<div class="flex gap-3">
													<img src={reply.author.avatar} alt={reply.author.name} class="size-7 rounded-full" />
													<div class="flex-1">
														<div class="rounded-xl bg-muted/50 p-3">
															<div class="mb-1 flex items-center gap-2">
																<span class="font-medium text-sm">{reply.author.name}</span>
																<span class="text-xs text-muted-foreground">{reply.createdAt}</span>
															</div>
															<p class="text-sm">{reply.text}</p>
														</div>
														<div class="mt-1 flex items-center gap-4">
															<button class="flex items-center gap-1 text-xs text-muted-foreground hover:text-foreground">
																<Icon icon="hugeicons:thumbs-up" class="size-3" />
																{reply.likes} likes
															</button>
														</div>
													</div>
												</div>
											{/each}
										</div>
									{/if}

									{#if replyToComment === comment.id}
										<div class="mt-3 flex gap-3">
											<img src={teamMembers[0].avatar} alt="You" class="size-7 rounded-full" />
											<div class="flex-1">
												<textarea
													placeholder="Write a reply..."
													bind:value={replyText}
													rows={2}
													class="w-full resize-none rounded-lg border border-input bg-background px-3 py-2 text-sm transition-colors placeholder:text-muted-foreground focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-none"
												></textarea>
												<div class="mt-2 flex justify-end gap-2">
													<Button variant="outline" size="sm" onclick={() => replyToComment = null}>
														Cancel
													</Button>
													<Button size="sm" onclick={() => handleReply(comment.id)} disabled={!replyText.trim()}>
														Reply
													</Button>
												</div>
											</div>
										</div>
									{/if}
								</div>
							</div>
						{/each}
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-6">
					<h3 class="mb-4 font-semibold">Activity</h3>
					<div class="space-y-4">
						{#each task.activity as activity}
							<div class="flex gap-3">
								<div class="flex size-8 shrink-0 items-center justify-center rounded-lg bg-muted">
									<Icon icon={getActivityIcon(activity.type)} class="size-4 text-muted-foreground" />
								</div>
								<div class="flex-1">
									<p class="text-sm">
										<span class="font-medium">{activity.user.name}</span>
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
					<h3 class="mb-4 font-semibold">Details</h3>
					<div class="space-y-4">
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Assignee</span>
							<div class="flex items-center gap-2">
								<img src={task.assignee.avatar} alt={task.assignee.name} class="size-6 rounded-full" />
								<span class="text-sm">{task.assignee.name}</span>
							</div>
						</div>
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Reporter</span>
							<div class="flex items-center gap-2">
								<img src={task.reporter.avatar} alt={task.reporter.name} class="size-6 rounded-full" />
								<span class="text-sm">{task.reporter.name}</span>
							</div>
						</div>
						<Separator />
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Due Date</span>
							<div class="flex items-center gap-2">
								<Icon icon="hugeicons:calendar-01" class="size-4 text-muted-foreground" />
								<span class="text-sm">{task.dueDate}</span>
							</div>
						</div>
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Created</span>
							<span class="text-sm">{task.createdAt}</span>
						</div>
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Updated</span>
							<span class="text-sm">{task.updatedAt}</span>
						</div>
						{#if task.completedAt}
							<div class="flex items-center justify-between">
								<span class="text-sm text-muted-foreground">Completed</span>
								<span class="text-sm text-emerald-600">{task.completedAt}</span>
							</div>
						{/if}
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Time Tracking</h3>
					<div class="space-y-4">
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Estimated</span>
							<span class="text-sm font-medium">{task.estimatedHours}h</span>
						</div>
						<div class="flex items-center justify-between">
							<span class="text-sm text-muted-foreground">Logged</span>
							<span class="text-sm font-medium">{task.loggedHours}h</span>
						</div>
						<div class="h-2 overflow-hidden rounded-full bg-muted">
							<div 
								class="h-full rounded-full bg-primary transition-all" 
								style="width: {Math.min((task.loggedHours / task.estimatedHours) * 100, 100)}%"
							></div>
						</div>
						<div class="flex items-center justify-between text-xs">
							<span class="text-muted-foreground">{Math.round((task.loggedHours / task.estimatedHours) * 100)}% used</span>
							<span class="{task.loggedHours > task.estimatedHours ? 'text-rose-500' : 'text-emerald-600'}">
								{task.loggedHours > task.estimatedHours ? '+' : ''}{task.loggedHours - task.estimatedHours}h
							</span>
						</div>
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Watchers ({task.watchers.length})</h3>
					<div class="flex -space-x-2">
						{#each task.watchers as watcher}
							<img src={watcher.avatar} alt={watcher.name} class="size-8 rounded-full border-2 border-background" title={watcher.name} />
						{/each}
						<button class="flex size-8 items-center justify-center rounded-full border-2 border-background bg-muted text-xs font-medium hover:bg-muted/80">
							<Icon icon="hugeicons:plus-sign" class="size-4" />
						</button>
					</div>
				</div>

				<div class="rounded-2xl border border-border bg-card p-5">
					<h3 class="mb-4 font-semibold">Related Tasks</h3>
					<div class="space-y-3">
						{#each task.relatedTasks as related}
							<a href="/tasks/{related.id}" class="group flex items-center gap-3 rounded-lg p-2 transition-colors hover:bg-muted">
								<div class="size-2 rounded-full {related.project.color}"></div>
								<div class="flex-1 min-w-0">
									<p class="truncate text-sm font-medium transition-colors group-hover:text-primary">{related.title}</p>
									<p class="text-xs text-muted-foreground">{related.project.name}</p>
								</div>
								<span class="flex shrink-0 items-center gap-1 rounded-full px-2 py-0.5 text-xs font-medium {getStatusColor(related.status)}">
									{related.status.replace('-', ' ')}
								</span>
							</a>
						{/each}
					</div>
				</div>
			</div>
		</div>
	</div>
</div>

<Modal
	bind:open={showDeleteModal}
	title="Delete Task"
	description="Are you sure you want to delete this task? This action cannot be undone."
	size="sm"
>
	<div class="flex justify-end gap-2">
		<Button variant="outline" onclick={() => showDeleteModal = false}>Cancel</Button>
		<Button variant="destructive" onclick={() => { showDeleteModal = false; window.location.href = '/tasks'; }}>
			<Icon icon="hugeicons:delete-01" class="mr-2 size-4" />
			Delete Task
		</Button>
	</div>
</Modal>
