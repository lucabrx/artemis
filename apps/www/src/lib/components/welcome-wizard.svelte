<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';

	let { onComplete }: { onComplete: () => void } = $props();

	let step = $state(1);
	let workspaceName = $state('');
	let workspaceType = $state('');
	let teamSize = $state('');
	let useCases = $state<string[]>([]);

	const totalSteps = 4;

	const workspaceTypes = [
		{ id: 'agency', label: 'Creative Agency', icon: 'lucide:palette' },
		{ id: 'freelance', label: 'Freelance', icon: 'lucide:user' },
		{ id: 'product', label: 'Product Team', icon: 'lucide:box' },
		{ id: 'consulting', label: 'Consulting', icon: 'lucide:briefcase' }
	];

	const teamSizes = [
		{ id: 'solo', label: 'Just me', desc: 'Freelancer or solo founder', icon: 'lucide:user' },
		{ id: 'small', label: '2-5 people', desc: 'Small team or startup', icon: 'lucide:users' },
		{ id: 'medium', label: '6-15 people', desc: 'Growing agency or team', icon: 'lucide:building' },
		{
			id: 'large',
			label: '16+ people',
			desc: 'Established organization',
			icon: 'lucide:building-2'
		}
	];

	const useCaseOptions = [
		{ id: 'projects', label: 'Project Management', icon: 'lucide:folder-kanban' },
		{ id: 'clients', label: 'Client Management', icon: 'lucide:users' },
		{ id: 'invoicing', label: 'Invoicing', icon: 'lucide:file-text' },
		{ id: 'time', label: 'Time Tracking', icon: 'lucide:clock' },
		{ id: 'collab', label: 'Team Collaboration', icon: 'lucide:message-square' },
		{ id: 'reports', label: 'Reporting', icon: 'lucide:bar-chart-3' }
	];

	function next() {
		if (step < totalSteps) step++;
		else onComplete();
	}

	function back() {
		if (step > 1) step--;
	}

	function skip() {
		onComplete();
	}

	function toggleUseCase(id: string) {
		if (useCases.includes(id)) {
			useCases = useCases.filter((u) => u !== id);
		} else {
			useCases = [...useCases, id];
		}
	}
</script>

<div class="fixed inset-0 z-50 bg-background">
	<div class="flex h-full flex-col">
		<header class="flex items-center justify-between border-b border-border px-6 py-4">
			<div class="flex items-center gap-3">
				<div
					class="flex size-10 items-center justify-center rounded-xl bg-primary text-primary-foreground"
				>
					<Icon icon="hugeicons:cowboy-hat" class="size-6" />
				</div>
				<span class="text-xl font-bold">Artemis</span>
			</div>
			<div class="flex items-center gap-4">
				<div class="hidden gap-1 sm:flex">
					{#each Array(totalSteps) as _, i}
						<div class="h-1.5 w-8 rounded-full {i + 1 <= step ? 'bg-primary' : 'bg-muted'}"></div>
					{/each}
				</div>
				<Button variant="ghost" onclick={skip}>Skip</Button>
			</div>
		</header>

		<div class="flex-1 overflow-y-auto">
			<div class="mx-auto max-w-2xl px-6 py-12">
				{#if step === 1}
					<div class="mb-10 space-y-4 text-center">
						<div
							class="mb-4 inline-flex size-20 items-center justify-center rounded-2xl bg-primary/10"
						>
							<Icon icon="lucide:sparkles" class="size-10 text-primary" />
						</div>
						<h1 class="text-3xl font-bold tracking-tight">Welcome to Artemis!</h1>
						<p class="text-lg text-muted-foreground">
							Let's get your workspace set up in just a few steps.
						</p>
					</div>

					<div class="space-y-6">
						<div class="space-y-2">
							<label for="workspaceName" class="text-sm font-medium"
								>What should we call your workspace?</label
							>
							<input
								id="workspaceName"
								type="text"
								bind:value={workspaceName}
								placeholder="e.g., Design Studio, Acme Agency"
								class="h-12 w-full rounded-xl border border-input bg-background px-4 text-lg focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none"
							/>
						</div>

						<div class="space-y-2">
							<span class="text-sm font-medium">What type of work do you do?</span>
							<div class="grid grid-cols-2 gap-3">
								{#each workspaceTypes as type}
									<button
										class="flex flex-col items-center gap-3 rounded-xl border-2 p-6 transition-all {workspaceType ===
										type.id
											? 'border-primary bg-primary/5'
											: 'border-border hover:border-primary/30'}"
										onclick={() => (workspaceType = type.id)}
									>
										<Icon
											icon={type.icon}
											class="size-8 {workspaceType === type.id
												? 'text-primary'
												: 'text-muted-foreground'}"
										/>
										<span class="font-medium">{type.label}</span>
									</button>
								{/each}
							</div>
						</div>
					</div>
				{:else if step === 2}
					<div class="mb-10 space-y-4 text-center">
						<div
							class="mb-4 inline-flex size-20 items-center justify-center rounded-2xl bg-blue-500/10"
						>
							<Icon icon="lucide:users" class="size-10 text-blue-500" />
						</div>
						<h1 class="text-3xl font-bold tracking-tight">How big is your team?</h1>
						<p class="text-lg text-muted-foreground">
							This helps us tailor the experience for your workflow.
						</p>
					</div>

					<div class="grid gap-3">
						{#each teamSizes as size}
							<button
								class="flex items-center gap-4 rounded-xl border-2 p-5 text-left transition-all {teamSize ===
								size.id
									? 'border-primary bg-primary/5'
									: 'border-border hover:border-primary/30'}"
								onclick={() => (teamSize = size.id)}
							>
								<div
									class="flex size-12 items-center justify-center rounded-xl {teamSize === size.id
										? 'bg-primary/10'
										: 'bg-muted'}"
								>
									<Icon
										icon={size.icon}
										class="size-6 {teamSize === size.id ? 'text-primary' : 'text-muted-foreground'}"
									/>
								</div>
								<div class="flex-1">
									<p class="font-semibold">{size.label}</p>
									<p class="text-sm text-muted-foreground">{size.desc}</p>
								</div>
								{#if teamSize === size.id}
									<Icon icon="lucide:check-circle" class="size-6 text-primary" />
								{/if}
							</button>
						{/each}
					</div>
				{:else if step === 3}
					<div class="mb-10 space-y-4 text-center">
						<div
							class="mb-4 inline-flex size-20 items-center justify-center rounded-2xl bg-violet-500/10"
						>
							<Icon icon="lucide:target" class="size-10 text-violet-500" />
						</div>
						<h1 class="text-3xl font-bold tracking-tight">What will you use Artemis for?</h1>
						<p class="text-lg text-muted-foreground">
							Select all that apply. We'll customize your dashboard.
						</p>
					</div>

					<div class="grid grid-cols-2 gap-3">
						{#each useCaseOptions as option}
							<button
								class="flex flex-col items-center gap-3 rounded-xl border-2 p-5 transition-all {useCases.includes(
									option.id
								)
									? 'border-primary bg-primary/5'
									: 'border-border hover:border-primary/30'}"
								onclick={() => toggleUseCase(option.id)}
							>
								<Icon
									icon={option.icon}
									class="size-8 {useCases.includes(option.id)
										? 'text-primary'
										: 'text-muted-foreground'}"
								/>
								<span class="text-center text-sm font-medium">{option.label}</span>
								{#if useCases.includes(option.id)}
									<Icon icon="lucide:check" class="size-4 text-primary" />
								{/if}
							</button>
						{/each}
					</div>
				{:else if step === 4}
					<div class="space-y-6 py-8 text-center">
						<div
							class="mb-6 inline-flex size-24 items-center justify-center rounded-3xl bg-emerald-500/10"
						>
							<Icon icon="lucide:check" class="size-12 text-emerald-500" />
						</div>
						<h1 class="text-3xl font-bold tracking-tight">You're all set!</h1>
						<p class="mx-auto max-w-md text-lg text-muted-foreground">
							Your workspace is ready. Start managing projects, tracking time, and invoicing
							clients.
						</p>

						<div class="mx-auto mt-8 grid max-w-md grid-cols-3 gap-4">
							<button
								class="rounded-xl bg-muted/50 p-4 text-center transition-colors hover:bg-muted"
							>
								<Icon icon="lucide:plus" class="mx-auto mb-2 size-6 text-primary" />
								<p class="text-sm font-medium">Create Project</p>
							</button>
							<button
								class="rounded-xl bg-muted/50 p-4 text-center transition-colors hover:bg-muted"
							>
								<Icon icon="lucide:user-plus" class="mx-auto mb-2 size-6 text-primary" />
								<p class="text-sm font-medium">Add Client</p>
							</button>
							<button
								class="rounded-xl bg-muted/50 p-4 text-center transition-colors hover:bg-muted"
							>
								<Icon icon="lucide:settings" class="mx-auto mb-2 size-6 text-primary" />
								<p class="text-sm font-medium">Explore</p>
							</button>
						</div>
					</div>
				{/if}
			</div>
		</div>

		<div class="border-t border-border px-6 py-4">
			<div class="mx-auto flex max-w-2xl justify-between">
				<Button variant="ghost" onclick={back} disabled={step === 1}>
					<Icon icon="lucide:arrow-left" class="mr-2 size-4" />
					Back
				</Button>
				<Button onclick={next} size="lg" class="gap-2">
					{step === totalSteps ? 'Get Started' : 'Continue'}
					<Icon
						icon={step === totalSteps ? 'lucide:rocket' : 'lucide:arrow-right'}
						class="size-4"
					/>
				</Button>
			</div>
		</div>
	</div>
</div>
