<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Modal } from '$lib/components/ui/modal/index.js';
	import { Label } from '$lib/components/ui/label/index.js';

	let activeTab = $state<'inbox' | 'templates'>('inbox');
	let inboxFilter = $state('all');
	let templateFilter = $state('all');
	let searchQuery = $state('');
	let selectedEmail = $state<Email | null>(null);
	let selectedTemplate = $state<Template | null>(null);
	let showComposeModal = $state(false);
	let showTemplateModal = $state(false);
	let showPreviewModal = $state(false);
	
	let sidebarWidth = $state(480);
	let isResizing = $state(false);
	let resizeContainer: HTMLElement | null = $state(null);
	
	function startResize(e: MouseEvent) {
		isResizing = true;
		document.body.style.cursor = 'col-resize';
		document.body.style.userSelect = 'none';
	}
	
	function handleResize(e: MouseEvent) {
		if (!isResizing || !resizeContainer) return;
		const rect = resizeContainer.getBoundingClientRect();
		const newWidth = e.clientX - rect.left;
		sidebarWidth = Math.max(300, Math.min(600, newWidth));
	}
	
	function stopResize() {
		isResizing = false;
		document.body.style.cursor = '';
		document.body.style.userSelect = '';
	}

	interface Email {
		id: string;
		sender: {
			name: string;
			email: string;
			avatar: string;
		};
		subject: string;
		preview: string;
		body: string;
		date: string;
		read: boolean;
		starred: boolean;
		folder: 'inbox' | 'sent' | 'drafts' | 'archive';
		labels: string[];
		hasAttachments: boolean;
		attachments?: { name: string; size: string }[];
	}

	interface Template {
		id: string;
		name: string;
		subject: string;
		category: 'onboarding' | 'proposals' | 'follow-ups' | 'invoices' | 'updates' | 'general';
		content: string;
		variables: string[];
		lastUsed: string;
		usageCount: number;
		favorite: boolean;
	}

	const mockEmails: Email[] = [
		{
			id: 'e1',
			sender: {
				name: 'Sarah Chen',
				email: 'sarah@acmecorp.com',
				avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah'
			},
			subject: 'Website redesign project - Timeline update',
			preview: 'Hi Alex, I wanted to reach out regarding the timeline for the website redesign project. We\'re making good progress...',
			body: `Hi Alex,<br><br>
				I wanted to reach out regarding the timeline for the website redesign project. We're making good progress on the initial mockups and should have something to show you by end of week.<br><br>
				<strong>Key updates:</strong><br>
				• Homepage design - 80% complete<br>
				• About page wireframes - Ready for review<br>
				• Mobile responsive layouts - In progress<br><br>
				Could we schedule a quick call on Friday to discuss the direction? Let me know what time works best for you.<br><br>
				Best regards,<br>
				Sarah`,
			date: '10:30 AM',
			read: false,
			starred: true,
			folder: 'inbox',
			labels: ['Work', 'Important'],
			hasAttachments: true,
			attachments: [{ name: 'mockups-v2.fig', size: '12.4 MB' }]
		},
		{
			id: 'e2',
			sender: {
				name: 'Michael Ross',
				email: 'mike@techstart.io',
				avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Mike'
			},
			subject: 'Invoice #INV-2026-002 - Payment received',
			preview: 'Thank you for your payment! We have received your transfer of $4,375 for Invoice #INV-2026-002...',
			body: `Hi Alex,<br><br>
				Thank you for your payment! We have received your transfer of $4,375 for Invoice #INV-2026-002.<br><br>
				<strong>Payment Details:</strong><br>
				Invoice: INV-2026-002<br>
				Amount: $4,375.00<br>
				Date Received: Jan 25, 2026<br>
				Payment Method: Bank Transfer<br><br>
				Your remaining balance is $4,375.00, due on Feb 20, 2026.<br><br>
				Thanks again for your business!<br><br>
				Best,<br>
				Michael`,
			date: 'Yesterday',
			read: true,
			starred: false,
			folder: 'inbox',
			labels: ['Finance'],
			hasAttachments: false
		},
		{
			id: 'e3',
			sender: {
				name: 'Emily Watson',
				email: 'emily@greenleaf.co',
				avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Emily'
			},
			subject: 'Brand guidelines feedback',
			preview: 'Hi there! I\'ve reviewed the brand guidelines document and have some feedback to share...',
			body: `Hi Alex,<br><br>
				I've reviewed the brand guidelines document and have some feedback to share. Overall, I love the direction we're taking!<br><br>
				<strong>Feedback:</strong><br>
				• The color palette is perfect - exactly what we envisioned<br>
				• Typography choices are clean and modern<br>
				• Could we explore one more variation of the logo mark?<br>
				• The spacing guidelines are very helpful<br><br>
				I've attached my detailed notes. Let me know when you'd like to discuss!<br><br>
				Cheers,<br>
				Emily`,
			date: 'Jan 26',
			read: true,
			starred: false,
			folder: 'inbox',
			labels: ['Work'],
			hasAttachments: true,
			attachments: [{ name: 'brand-feedback.pdf', size: '2.1 MB' }]
		},
		{
			id: 'e4',
			sender: {
				name: 'David Park',
				email: 'david@globaldynamics.com',
				avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=David'
			},
			subject: 'Marketing campaign kickoff meeting',
			preview: 'Looking forward to our kickoff meeting next Monday. I\'ve prepared the initial brief...',
			body: `Hi Alex,<br><br>
				Looking forward to our kickoff meeting next Monday. I've prepared the initial brief and would love to walk you through our goals for Q1.<br><br>
				<strong>Agenda:</strong><br>
				1. Campaign objectives and KPIs<br>
				2. Target audience analysis<br>
				3. Creative direction discussion<br>
				4. Timeline and milestones<br>
				5. Budget allocation<br><br>
				Meeting link: [Zoom Link]<br><br>
				See you Monday!<br><br>
				David`,
			date: 'Jan 25',
			read: true,
			starred: true,
			folder: 'inbox',
			labels: ['Meeting', 'Important'],
			hasAttachments: false
		},
		{
			id: 'e5',
			sender: {
				name: 'Lisa Thompson',
				email: 'lisa@innovatelabs.com',
				avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Lisa'
			},
			subject: 'Q1 Project requirements document',
			preview: 'Attached is the detailed requirements document for the Q1 projects we discussed...',
			body: `Hi Alex,<br><br>
				Attached is the detailed requirements document for the Q1 projects we discussed in our last call.<br><br>
				<strong>Key highlights:</strong><br>
				• 3 new features to implement<br>
				• Performance improvements (target: 40% faster load times)<br>
				• Mobile app enhancements<br>
				• API integrations with third-party services<br><br>
				Please review and let me know if you have any questions. We're aiming to start development by Feb 1st.<br><br>
				Best,<br>
				Lisa`,
			date: 'Jan 24',
			read: false,
			starred: false,
			folder: 'inbox',
			labels: ['Work', 'Urgent'],
			hasAttachments: true,
			attachments: [
				{ name: 'Q1-Requirements.pdf', size: '4.8 MB' },
				{ name: 'technical-specs.docx', size: '1.2 MB' }
			]
		},
		{
			id: 'e6',
			sender: {
				name: 'You',
				email: 'alex@artemis.app',
				avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex'
			},
			subject: 'Re: Proposal for new project',
			preview: 'Thanks for sending over the proposal. I\'ve reviewed it and have a few questions...',
			body: `Hi Team,<br><br>
				Thanks for sending over the proposal. I've reviewed it and have a few questions before we move forward.<br><br>
				1. Can we adjust the timeline to accommodate the holiday period?<br>
				2. The budget seems a bit high - can we discuss scope adjustments?<br>
				3. Who will be the main point of contact on your end?<br><br>
				Looking forward to hearing from you.<br><br>
				Best,<br>
				Alex`,
			date: 'Jan 23',
			read: true,
			starred: false,
			folder: 'sent',
			labels: [],
			hasAttachments: false
		}
	];

	const mockTemplates: Template[] = [
		{
			id: 't1',
			name: 'New Client Onboarding',
			subject: 'Welcome to {{company_name}} - Let\'s get started!',
			category: 'onboarding',
			content: `Hi {{client_name}},<br><br>
				Welcome to {{company_name}}! I'm thrilled to have you on board.<br><br>
				<strong>Next steps:</strong><br>
				1. We'll schedule a kickoff call within the next 48 hours<br>
				2. You'll receive access to our project management tool<br>
				3. I'll send over the initial questionnaire<br><br>
				If you have any questions in the meantime, feel free to reach out.<br><br>
				Best regards,<br>
				{{sender_name}}`,
			variables: ['client_name', 'company_name', 'sender_name'],
			lastUsed: '2 days ago',
			usageCount: 24,
			favorite: true
		},
		{
			id: 't2',
			name: 'Project Proposal',
			subject: 'Project Proposal: {{project_name}}',
			category: 'proposals',
			content: `Hi {{client_name}},<br><br>
				Thank you for considering {{company_name}} for your {{project_name}} project.<br><br>
				<strong>Proposal Summary:</strong><br>
				• Scope: {{project_scope}}<br>
				• Timeline: {{timeline}}<br>
				• Investment: {{price}}<br><br>
				I've attached the detailed proposal for your review. Please let me know if you have any questions or would like to schedule a call to discuss.<br><br>
				Looking forward to working together!<br><br>
				Best,<br>
				{{sender_name}}`,
			variables: ['client_name', 'company_name', 'project_name', 'project_scope', 'timeline', 'price', 'sender_name'],
			lastUsed: '1 week ago',
			usageCount: 56,
			favorite: true
		},
		{
			id: 't3',
			name: 'Follow-up After Meeting',
			subject: 'Great speaking with you - Next steps',
			category: 'follow-ups',
			content: `Hi {{client_name}},<br><br>
				It was great speaking with you today about {{project_name}}.<br><br>
				<strong>As discussed, here are the next steps:</strong><br>
				1. {{next_step_1}}<br>
				2. {{next_step_2}}<br>
				3. {{next_step_3}}<br><br>
				I'll send over {{deliverable}} by {{date}}.<br><br>
				Thanks again for your time!<br><br>
				Best,<br>
				{{sender_name}}`,
			variables: ['client_name', 'project_name', 'next_step_1', 'next_step_2', 'next_step_3', 'deliverable', 'date', 'sender_name'],
			lastUsed: '3 days ago',
			usageCount: 38,
			favorite: false
		},
		{
			id: 't4',
			name: 'Invoice Reminder',
			subject: 'Friendly reminder: Invoice {{invoice_number}} due {{due_date}}',
			category: 'invoices',
			content: `Hi {{client_name}},<br><br>
				I hope you're doing well.<br><br>
				This is a friendly reminder that Invoice {{invoice_number}} for {{amount}} is due on {{due_date}}.<br><br>
				You can view and pay the invoice here: {{invoice_link}}<br><br>
				If you have any questions or concerns, please don't hesitate to reach out.<br><br>
				Thank you for your business!<br><br>
				Best regards,<br>
				{{sender_name}}`,
			variables: ['client_name', 'invoice_number', 'amount', 'due_date', 'invoice_link', 'sender_name'],
			lastUsed: '5 days ago',
			usageCount: 42,
			favorite: false
		},
		{
			id: 't5',
			name: 'Project Update',
			subject: '{{project_name}} - Weekly Update',
			category: 'updates',
			content: `Hi {{client_name}},<br><br>
				Here's your weekly update on {{project_name}}:<br><br>
				<strong>Completed this week:</strong><br>
				• {{completed_1}}<br>
				• {{completed_2}}<br>
				• {{completed_3}}<br><br>
				<strong>Up next:</strong><br>
				• {{upcoming_1}}<br>
				• {{upcoming_2}}<br><br>
				<strong>Questions/Notes:</strong><br>
				{{notes}}<br><br>
				Please let me know if you have any questions!<br><br>
				Best,<br>
				{{sender_name}}`,
			variables: ['client_name', 'project_name', 'completed_1', 'completed_2', 'completed_3', 'upcoming_1', 'upcoming_2', 'notes', 'sender_name'],
			lastUsed: '1 day ago',
			usageCount: 31,
			favorite: true
		},
		{
			id: 't6',
			name: 'Thank You Email',
			subject: 'Thank you for your business!',
			category: 'general',
			content: `Hi {{client_name}},<br><br>
				I wanted to take a moment to thank you for choosing {{company_name}} for {{project_name}}.<br><br>
				It's been a pleasure working with you, and I'm thrilled with the results we've achieved together.<br><br>
				If you know anyone who might benefit from our services, I'd greatly appreciate a referral.<br><br>
				Thanks again, and I look forward to working with you again in the future!<br><br>
				Warm regards,<br>
				{{sender_name}}`,
			variables: ['client_name', 'company_name', 'project_name', 'sender_name'],
			lastUsed: '2 weeks ago',
			usageCount: 18,
			favorite: false
		}
	];

	const categoryColors: Record<string, string> = {
		onboarding: 'bg-emerald-500/10 text-emerald-600 border-emerald-500/20',
		proposals: 'bg-blue-500/10 text-blue-600 border-blue-500/20',
		'follow-ups': 'bg-amber-500/10 text-amber-600 border-amber-500/20',
		invoices: 'bg-violet-500/10 text-violet-600 border-violet-500/20',
		updates: 'bg-cyan-500/10 text-cyan-600 border-cyan-500/20',
		general: 'bg-slate-500/10 text-slate-600 border-slate-500/20'
	};

	const labelColors: Record<string, string> = {
		Work: 'bg-blue-500/10 text-blue-600',
		Important: 'bg-rose-500/10 text-rose-600',
		Finance: 'bg-emerald-500/10 text-emerald-600',
		Meeting: 'bg-violet-500/10 text-violet-600',
		Urgent: 'bg-amber-500/10 text-amber-600'
	};

	const inboxStats = {
		unread: mockEmails.filter(e => e.folder === 'inbox' && !e.read).length,
		starred: mockEmails.filter(e => e.starred).length,
		sent: mockEmails.filter(e => e.folder === 'sent').length,
		drafts: 3
	};

	const templateStats = {
		total: mockTemplates.length,
		favorites: mockTemplates.filter(t => t.favorite).length,
		mostUsed: mockTemplates.reduce((max, t) => t.usageCount > max.usageCount ? t : max, mockTemplates[0])
	};

	const inboxFilters = [
		{ value: 'all', label: 'All Mail', count: mockEmails.filter(e => e.folder === 'inbox').length },
		{ value: 'unread', label: 'Unread', count: inboxStats.unread },
		{ value: 'starred', label: 'Starred', count: inboxStats.starred },
		{ value: 'sent', label: 'Sent', count: inboxStats.sent }
	];

	const templateCategories = [
		{ value: 'all', label: 'All Templates', count: mockTemplates.length },
		{ value: 'favorites', label: 'Favorites', count: templateStats.favorites },
		{ value: 'onboarding', label: 'Onboarding', count: mockTemplates.filter(t => t.category === 'onboarding').length },
		{ value: 'proposals', label: 'Proposals', count: mockTemplates.filter(t => t.category === 'proposals').length },
		{ value: 'follow-ups', label: 'Follow-ups', count: mockTemplates.filter(t => t.category === 'follow-ups').length },
		{ value: 'invoices', label: 'Invoices', count: mockTemplates.filter(t => t.category === 'invoices').length },
		{ value: 'updates', label: 'Updates', count: mockTemplates.filter(t => t.category === 'updates').length }
	];

	let filteredEmails = $derived(() => {
		let filtered = mockEmails.filter(e => e.folder === 'inbox' || inboxFilter === 'sent');
		if (inboxFilter === 'unread') filtered = filtered.filter(e => !e.read);
		if (inboxFilter === 'starred') filtered = filtered.filter(e => e.starred);
		if (inboxFilter === 'sent') filtered = mockEmails.filter(e => e.folder === 'sent');
		if (searchQuery) {
			const q = searchQuery.toLowerCase();
			filtered = filtered.filter(e => 
				e.subject.toLowerCase().includes(q) ||
				e.sender.name.toLowerCase().includes(q) ||
				e.preview.toLowerCase().includes(q)
			);
		}
		return filtered;
	});

	let filteredTemplates = $derived(() => {
		let filtered = mockTemplates;
		if (templateFilter === 'favorites') filtered = filtered.filter(t => t.favorite);
		if (templateFilter !== 'all' && templateFilter !== 'favorites') {
			filtered = filtered.filter(t => t.category === templateFilter);
		}
		if (searchQuery) {
			const q = searchQuery.toLowerCase();
			filtered = filtered.filter(t => 
				t.name.toLowerCase().includes(q) ||
				t.subject.toLowerCase().includes(q)
			);
		}
		return filtered;
	});

	function toggleStar(email: Email) {
		email.starred = !email.starred;
	}

	let composeData = $state({
		to: '',
		subject: '',
		body: ''
	});

	let newTemplate = $state({
		name: '',
		subject: '',
		category: 'general' as Template['category'],
		content: ''
	});

	function handleSendEmail(e: Event) {
		e.preventDefault();
		console.log('Sending email:', composeData);
		showComposeModal = false;
		composeData = { to: '', subject: '', body: '' };
	}

	function handleSaveTemplate(e: Event) {
		e.preventDefault();
		console.log('Saving template:', newTemplate);
		showTemplateModal = false;
		newTemplate = { name: '', subject: '', category: 'general', content: '' };
	}

	function useTemplate(template: Template) {
		composeData.subject = template.subject;
		composeData.body = template.content.replace(/<br>/g, '\n');
		showPreviewModal = false;
		showComposeModal = true;
	}
</script>

<svelte:head>
	<title>Emails - Artemis</title>
</svelte:head>

<svelte:window onmousemove={handleResize} onmouseup={stopResize} />

<div class="min-h-screen bg-gradient-to-b from-background to-muted/20">
	<div class="px-6 py-4">
		
		<div class="mb-6 flex items-center justify-end gap-2">
			{#if activeTab === 'inbox'}
				<Button variant="outline" class="gap-2">
					<Icon icon="hugeicons:settings-01" class="size-4" />
					Settings
				</Button>
				<Button class="gap-2" onclick={() => showComposeModal = true}>
					<Icon icon="hugeicons:pencil-edit-01" class="size-4" />
					Compose
				</Button>
			{:else}
				<Button variant="outline" class="gap-2">
					<Icon icon="hugeicons:upload-01" class="size-4" />
					Import
				</Button>
				<Button class="gap-2" onclick={() => showTemplateModal = true}>
					<Icon icon="hugeicons:plus-sign" class="size-4" />
					New Template
				</Button>
			{/if}
		</div>

		
		<div class="mt-8">
			<div class="flex items-center gap-1 rounded-lg border border-border bg-muted/50 p-1 w-fit">
				<button
					class="flex items-center gap-2 rounded-md px-4 py-2 text-sm font-medium transition-colors {activeTab === 'inbox' ? 'bg-background text-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground'}"
					onclick={() => { activeTab = 'inbox'; searchQuery = ''; }}
				>
					<Icon icon="hugeicons:mail-01" class="size-4" />
					Inbox
				</button>
				<button
					class="flex items-center gap-2 rounded-md px-4 py-2 text-sm font-medium transition-colors {activeTab === 'templates' ? 'bg-background text-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground'}"
					onclick={() => { activeTab = 'templates'; searchQuery = ''; }}
				>
					<Icon icon="hugeicons:layout-template" class="size-4" />
					Templates
				</button>
			</div>
		</div>

		
		{#if activeTab === 'inbox'}
			<div class="mb-4 flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between mt-4">
				<div class="relative flex-1 max-w-md">
					<Icon icon="hugeicons:search-01" class="absolute top-1/2 left-3 size-5 -translate-y-1/2 text-muted-foreground" />
					<Input placeholder="Search emails..." class="pl-10" bind:value={searchQuery} />
				</div>
				<div class="flex flex-wrap gap-2">
					{#each inboxFilters as filter}
						<button
							class="rounded-md px-3 py-1.5 text-sm font-medium transition-colors {inboxFilter === filter.value ? 'bg-primary text-primary-foreground' : 'border border-border text-muted-foreground hover:text-foreground'}"
							onclick={() => inboxFilter = filter.value}
						>
							{filter.label}
							<span class="ml-1 text-xs opacity-70">({filter.count})</span>
						</button>
					{/each}
				</div>
			</div>
		{/if}

		
		<div class="overflow-hidden rounded-2xl border border-border bg-card" bind:this={resizeContainer}>
			{#if activeTab === 'inbox'}
				
				<div class="flex flex-col lg:flex-row h-[calc(100vh-13rem)]">
					
					<div class="w-full lg:w-auto border-r border-border flex flex-col" style="width: {sidebarWidth}px; min-width: 350px; max-width: 700px;">
						
						<div class="divide-y divide-border flex-1 overflow-y-auto">
							{#each filteredEmails() as email}
								<div
									class="group cursor-pointer w-full text-left p-4 transition-colors hover:bg-muted/50 {selectedEmail?.id === email.id ? 'bg-primary/5 border-l-2 border-l-primary' : 'border-l-2 border-l-transparent'} {!email.read ? 'bg-muted/20' : ''}"
									onclick={() => { selectedEmail = email; email.read = true; }}
									role="button"
									tabindex="0"
									onkeydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { selectedEmail = email; email.read = true; } }}
								>
									<div class="flex items-start gap-3">
										<img src={email.sender.avatar} alt={email.sender.name} class="size-10 rounded-full bg-muted shrink-0" />
										<div class="flex-1 min-w-0">
											<div class="flex items-center justify-between gap-2">
												<span class="font-medium truncate {!email.read ? 'font-semibold' : ''}">{email.sender.name}</span>
												<span class="text-xs text-muted-foreground shrink-0">{email.date}</span>
											</div>
											<p class="text-sm font-medium truncate mt-0.5 {!email.read ? 'text-foreground' : 'text-muted-foreground'}">{email.subject}</p>
											<p class="text-xs text-muted-foreground truncate mt-0.5">{email.preview}</p>
											<div class="flex items-center gap-2 mt-2">
												{#each email.labels as label}
													<span class="text-[10px] px-1.5 py-0.5 rounded {labelColors[label] || 'bg-muted'}">{label}</span>
												{/each}
												{#if email.hasAttachments}
													<Icon icon="hugeicons:attachment-01" class="size-3 text-muted-foreground" />
												{/if}
											</div>
										</div>
										<div
											class="shrink-0 p-1 transition-colors {email.starred ? 'text-amber-500' : 'text-muted-foreground group-hover:text-amber-500'}"
											role="button"
											tabindex="0"
											onclick={(e) => { e.stopPropagation(); toggleStar(email); }}
											onkeydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.stopPropagation(); toggleStar(email); } }}
										>
											<Icon icon={email.starred ? "hugeicons:star" : "hugeicons:star-off"} class="size-4" />
										</div>
									</div>
								</div>
							{/each}
						</div>
					</div>

					
					<div
						class="hidden lg:flex w-4 -ml-2 z-10 cursor-col-resize items-center justify-center group"
						onmousedown={startResize}
						role="separator"
						aria-label="Resize sidebar"
						title="Drag to resize"
					>
						<div class="w-1 h-full bg-border group-hover:bg-primary/50 transition-colors flex items-center justify-center">
							<div class="w-0.5 h-8 rounded-full bg-muted-foreground/30 group-hover:bg-primary"></div>
						</div>
					</div>

					
					<div class="flex-1 bg-background min-h-0 overflow-auto">
						{#if selectedEmail}
							<div class="h-full flex flex-col">
								
								<div class="border-b border-border p-6">
									<div class="flex items-start justify-between gap-4">
										<h2 class="text-xl font-semibold">{selectedEmail.subject}</h2>
										<div class="flex items-center gap-1 shrink-0">
											<Button variant="ghost" size="icon" class="size-8">
												<Icon icon="hugeicons:reply" class="size-4" />
											</Button>
											<Button variant="ghost" size="icon" class="size-8">
												<Icon icon="hugeicons:forward" class="size-4" />
											</Button>
											<Button variant="ghost" size="icon" class="size-8">
												<Icon icon="hugeicons:more-vertical" class="size-4" />
											</Button>
										</div>
									</div>
									<div class="flex items-center gap-3 mt-4">
										<img src={selectedEmail.sender.avatar} alt={selectedEmail.sender.name} class="size-10 rounded-full bg-muted" />
										<div class="flex-1">
											<div class="flex items-center gap-2">
												<span class="font-medium">{selectedEmail.sender.name}</span>
												<span class="text-sm text-muted-foreground">&lt;{selectedEmail.sender.email}&gt;</span>
											</div>
											<div class="text-sm text-muted-foreground">To: alex@artemis.app</div>
										</div>
										<span class="text-sm text-muted-foreground">{selectedEmail.date}</span>
									</div>
								</div>

								
								<div class="flex-1 p-6 overflow-y-auto">
									<div class="prose prose-sm max-w-none dark:prose-invert">
										{@html selectedEmail.body}
									</div>
								</div>

								
								{#if selectedEmail.hasAttachments && selectedEmail.attachments}
									<div class="border-t border-border p-6">
										<p class="text-sm font-medium mb-3">Attachments ({selectedEmail.attachments.length})</p>
										<div class="flex flex-wrap gap-3">
											{#each selectedEmail.attachments as attachment}
												<div class="flex items-center gap-3 rounded-lg border border-border bg-muted/50 px-3 py-2">
													<Icon icon="hugeicons:file-01" class="size-8 text-muted-foreground" />
													<div>
														<p class="text-sm font-medium">{attachment.name}</p>
														<p class="text-xs text-muted-foreground">{attachment.size}</p>
													</div>
													<Button variant="ghost" size="icon" class="size-8 ml-2">
														<Icon icon="hugeicons:download-01" class="size-4" />
													</Button>
												</div>
											{/each}
										</div>
									</div>
								{/if}

								
								<div class="border-t border-border p-6">
									<div class="flex items-center gap-3">
										<Button class="gap-2" onclick={() => showComposeModal = true}>
											<Icon icon="hugeicons:reply" class="size-4" />
											Reply
										</Button>
										<Button variant="outline" class="gap-2">
											<Icon icon="hugeicons:reply-all" class="size-4" />
											Reply All
										</Button>
										<Button variant="outline" class="gap-2">
											<Icon icon="hugeicons:forward" class="size-4" />
											Forward
										</Button>
									</div>
								</div>
							</div>
						{:else}
							<div class="h-full flex flex-col items-center justify-center p-12 text-center">
								<div class="mb-4 inline-flex size-16 items-center justify-center rounded-full bg-muted">
									<Icon icon="hugeicons:mail-01" class="size-8 text-muted-foreground" />
								</div>
								<h3 class="text-lg font-semibold">Select an email</h3>
								<p class="mt-1 text-muted-foreground">Choose an email from the list to view its contents.</p>
							</div>
						{/if}
					</div>
				</div>
			{:else}
				
				<div class="p-4 space-y-4">
					
					<div class="flex flex-col gap-3 lg:flex-row lg:items-center">
						<div class="relative flex-1">
							<Icon icon="hugeicons:search-01" class="absolute top-1/2 left-3 size-5 -translate-y-1/2 text-muted-foreground" />
							<Input placeholder="Search templates..." class="pl-10" bind:value={searchQuery} />
						</div>
						<div class="flex flex-wrap gap-2">
							{#each templateCategories as cat}
								<button
									class="rounded-md px-3 py-1.5 text-sm font-medium transition-colors {templateFilter === cat.value ? 'bg-primary text-primary-foreground' : 'border border-border text-muted-foreground hover:text-foreground'}"
									onclick={() => templateFilter = cat.value}
								>
									{cat.label}
									<span class="ml-1 text-xs opacity-70">({cat.count})</span>
								</button>
							{/each}
						</div>
					</div>

					
					{#if filteredTemplates().length === 0}
						<div class="p-12 text-center">
							<div class="mb-4 inline-flex size-16 items-center justify-center rounded-full bg-muted">
								<Icon icon="hugeicons:layout-template" class="size-8 text-muted-foreground" />
							</div>
							<h3 class="text-lg font-semibold">No templates found</h3>
							<p class="mt-1 text-muted-foreground">Try adjusting your search or filters.</p>
						</div>
					{:else}
						<div class="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
							{#each filteredTemplates() as template}
								<div class="group rounded-xl border border-border bg-card p-5 transition-all hover:border-primary/20 hover:shadow-lg">
									<div class="flex items-start justify-between">
										<div class="flex-1 min-w-0">
											<div class="flex items-center gap-2">
												<h3 class="font-semibold truncate">{template.name}</h3>
												{#if template.favorite}
													<Icon icon="hugeicons:star" class="size-4 text-amber-500 shrink-0" />
												{/if}
											</div>
											<span class="inline-flex items-center gap-1 rounded-full border px-2 py-0.5 text-xs font-medium capitalize mt-2 {categoryColors[template.category]}">
												{template.category.replace('-', ' ')}
											</span>
										</div>
										<Button variant="ghost" size="icon" class="size-8 opacity-0 transition-opacity group-hover:opacity-100">
											<Icon icon="hugeicons:more-vertical" class="size-4" />
										</Button>
									</div>

									<div class="mt-4 space-y-2">
										<p class="text-sm text-muted-foreground truncate">
											<span class="font-medium text-foreground">Subject:</span> {template.subject}
										</p>
										<div class="flex flex-wrap gap-1">
											{#each template.variables.slice(0, 4) as variable}
												<code class="text-xs px-1.5 py-0.5 rounded bg-muted">{`{{${variable}}}`}</code>
											{/each}
											{#if template.variables.length > 4}
												<code class="text-xs px-1.5 py-0.5 rounded bg-muted">+{template.variables.length - 4}</code>
											{/if}
										</div>
									</div>

									<div class="mt-4 pt-4 border-t border-border">
										<div class="flex items-center justify-between text-sm text-muted-foreground mb-3">
											<span class="flex items-center gap-1">
												<Icon icon="hugeicons:clock" class="size-4" />
												{template.lastUsed}
											</span>
											<span class="flex items-center gap-1">
												<Icon icon="hugeicons:send-01" class="size-4" />
												{template.usageCount} uses
											</span>
										</div>
										<div class="flex gap-2">
											<Button variant="outline" size="sm" class="flex-1 gap-1" onclick={() => { selectedTemplate = template; showPreviewModal = true; }}>
												<Icon icon="hugeicons:view" class="size-4" />
												Preview
											</Button>
											<Button size="sm" class="flex-1 gap-1" onclick={() => useTemplate(template)}>
												<Icon icon="hugeicons:pencil-edit-01" class="size-4" />
												Use
											</Button>
										</div>
									</div>
								</div>
							{/each}
						</div>
					{/if}
				</div>
			{/if}
		</div>
	</div>
</div>


<Modal
	bind:open={showComposeModal}
	title="Compose Email"
	description="Write a new email message."
	size="xl"
>
	<form id="compose-form" onsubmit={handleSendEmail} class="space-y-4">
		<div class="space-y-2">
			<Label for="email-to">To</Label>
			<Input id="email-to" placeholder="recipient@example.com" bind:value={composeData.to} />
		</div>
		<div class="space-y-2">
			<Label for="email-subject">Subject</Label>
			<Input id="email-subject" placeholder="Email subject..." bind:value={composeData.subject} />
		</div>
		<div class="space-y-2">
			<Label for="email-body">Message</Label>
			<div class="rounded-md border border-input">
				<div class="flex items-center gap-1 border-b border-input bg-muted/50 p-2">
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:text-bold" class="size-4" />
					</Button>
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:text-italic" class="size-4" />
					</Button>
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:text-underline" class="size-4" />
					</Button>
					<Separator orientation="vertical" class="mx-1 h-5" />
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:left-to-right-list-bullet" class="size-4" />
					</Button>
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:number-01" class="size-4" />
					</Button>
					<Separator orientation="vertical" class="mx-1 h-5" />
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:link-01" class="size-4" />
					</Button>
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:image-01" class="size-4" />
					</Button>
				</div>
				<textarea
					id="email-body"
					placeholder="Write your message..."
					bind:value={composeData.body}
					rows={10}
					class="w-full resize-none rounded-b-md bg-background px-3 py-2 text-sm outline-none"
				></textarea>
			</div>
		</div>
		<div class="flex items-center gap-2 rounded-lg border border-dashed border-border p-3">
			<Icon icon="hugeicons:attachment-01" class="size-5 text-muted-foreground" />
			<span class="text-sm text-muted-foreground">Drop files here or click to attach</span>
		</div>
	</form>

	{#snippet footer()}
		<Button variant="outline" onclick={() => showComposeModal = false}>Discard</Button>
		<div class="flex gap-2">
			<Button variant="outline" class="gap-2">
				<Icon icon="hugeicons:save-01" class="size-4" />
				Save Draft
			</Button>
			<Button type="submit" form="compose-form" class="gap-2">
				<Icon icon="hugeicons:send-01" class="size-4" />
				Send Email
			</Button>
		</div>
	{/snippet}
</Modal>


<Modal
	bind:open={showTemplateModal}
	title="Create New Template"
	description="Create a reusable email template with variables."
	size="xl"
>
	<form id="template-form" onsubmit={handleSaveTemplate} class="space-y-4">
		<div class="grid gap-4 md:grid-cols-2">
			<div class="space-y-2">
				<Label for="template-name">Template Name *</Label>
				<Input id="template-name" placeholder="e.g., Client Onboarding" bind:value={newTemplate.name} required />
			</div>
			<div class="space-y-2">
				<Label for="template-category">Category</Label>
				<select
					id="template-category"
					bind:value={newTemplate.category}
					class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-1 text-sm"
				>
					<option value="onboarding">Onboarding</option>
					<option value="proposals">Proposals</option>
					<option value="follow-ups">Follow-ups</option>
					<option value="invoices">Invoices</option>
					<option value="updates">Updates</option>
					<option value="general">General</option>
				</select>
			</div>
		</div>
		<div class="space-y-2">
			<Label for="template-subject">Subject Line *</Label>
			<Input id="template-subject" placeholder="e.g., Welcome to {{company_name}}!" bind:value={newTemplate.subject} required />
			<p class="text-xs text-muted-foreground">Use {'{{variable_name}}'} for dynamic content</p>
		</div>
		<div class="space-y-2">
			<Label for="template-content">Template Content *</Label>
			<div class="rounded-md border border-input">
				<div class="flex items-center gap-1 border-b border-input bg-muted/50 p-2">
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:text-bold" class="size-4" />
					</Button>
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:text-italic" class="size-4" />
					</Button>
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:text-underline" class="size-4" />
					</Button>
					<Separator orientation="vertical" class="mx-1 h-5" />
					<Button type="button" variant="ghost" size="icon" class="size-8">
						<Icon icon="hugeicons:insert-variable" class="size-4" />
					</Button>
				</div>
				<textarea
					id="template-content"
					placeholder="Write your template content... Use {{variable_name}} for dynamic fields."
					bind:value={newTemplate.content}
					rows={10}
					required
					class="w-full resize-none rounded-b-md bg-background px-3 py-2 text-sm outline-none"
				></textarea>
			</div>
		</div>
		<div class="rounded-lg bg-muted/50 p-3">
			<p class="text-sm font-medium mb-2">Available Variables</p>
			<div class="flex flex-wrap gap-2">
				{#each ['client_name', 'company_name', 'project_name', 'sender_name', 'date', 'invoice_number', 'amount'] as varName}
					<button
						type="button"
						class="text-xs px-2 py-1 rounded bg-background border border-border hover:border-primary transition-colors"
						onclick={() => newTemplate.content += ` {{${varName}}}`}
					>
						{`{{${varName}}}`}
					</button>
				{/each}
			</div>
		</div>
	</form>

	{#snippet footer()}
		<Button variant="outline" onclick={() => showTemplateModal = false}>Cancel</Button>
		<Button type="submit" form="template-form" class="gap-2">
			<Icon icon="hugeicons:save-01" class="size-4" />
			Save Template
		</Button>
	{/snippet}
</Modal>


<Modal
	bind:open={showPreviewModal}
	title={selectedTemplate?.name || 'Template Preview'}
	description="Preview how this template will look when sent."
	size="lg"
>
	{#if selectedTemplate}
		<div class="space-y-4">
			<div class="flex items-center gap-2">
				<span class="inline-flex items-center gap-1 rounded-full border px-2 py-0.5 text-xs font-medium capitalize {categoryColors[selectedTemplate.category]}">
					{selectedTemplate.category.replace('-', ' ')}
				</span>
				{#if selectedTemplate.favorite}
					<Icon icon="hugeicons:star" class="size-4 text-amber-500" />
				{/if}
			</div>
			<div class="rounded-lg border border-border bg-muted/30 p-4">
				<p class="text-sm text-muted-foreground mb-1">Subject:</p>
				<p class="font-medium">{selectedTemplate.subject}</p>
			</div>
			<div class="rounded-lg border border-border p-4">
				<p class="text-sm text-muted-foreground mb-2">Preview:</p>
				<div class="prose prose-sm max-w-none dark:prose-invert">
					{@html selectedTemplate.content}
				</div>
			</div>
			<div class="rounded-lg bg-muted/50 p-3">
				<p class="text-sm font-medium mb-2">Variables used in this template:</p>
				<div class="flex flex-wrap gap-2">
					{#each selectedTemplate.variables as variable}
						<code class="text-xs px-2 py-1 rounded bg-background border border-border">{`{{${variable}}}`}</code>
					{/each}
				</div>
			</div>
		</div>
	{/if}

	{#snippet footer()}
		<Button variant="outline" onclick={() => showPreviewModal = false}>Close</Button>
		<Button class="gap-2" onclick={() => useTemplate(selectedTemplate!)}>
			<Icon icon="hugeicons:pencil-edit-01" class="size-4" />
			Use Template
		</Button>
	{/snippet}
</Modal>
