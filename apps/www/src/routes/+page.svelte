<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import Icon from '@iconify/svelte';
	import { fly, fade } from 'svelte/transition';
	import { onMount } from 'svelte';

	let mounted = $state(false);

	onMount(() => {
		mounted = true;
	});
</script>

<svelte:head>
	<title>Artemis — Currently in Beta</title>
	<meta name="description" content="The next-generation platform for modern teams. Currently in beta." />
</svelte:head>

<div class="flex h-[100dvh] w-full flex-col bg-slate-950">
	<nav class="flex items-center justify-between px-6 py-5 lg:px-12">
		{#if mounted}
			<div class="flex items-center gap-3" in:fly={{ y: -10, duration: 400, delay: 100 }}>
				<div class="flex h-9 w-9 items-center justify-center rounded-lg bg-emerald-500">
					<Icon icon="lucide:layers" class="h-4 w-4 text-white" />
				</div>
				<span class="text-lg font-semibold text-white">Artemis</span>
			</div>
			
			<div class="flex items-center gap-3" in:fly={{ y: -10, duration: 400, delay: 200 }}>
				<Button variant="ghost" href="/sign-in" class="text-sm text-slate-400 hover:bg-slate-900 hover:text-white">
					Sign in
				</Button>
				<Button href="/sign-up" class="bg-white text-slate-950 hover:bg-slate-200">
					Get Started
				</Button>
			</div>
		{/if}
	</nav>

	<div class="flex flex-1 flex-col items-center justify-center px-6 text-center">
		{#if mounted}
			<div in:fly={{ y: -10, duration: 400, delay: 300 }}>
				<div class="inline-flex items-center gap-2 rounded-full border border-slate-800 bg-slate-900 px-4 py-1.5">
					<span class="relative flex h-1.5 w-1.5">
						<span class="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-500 opacity-75"></span>
						<span class="relative inline-flex h-1.5 w-1.5 rounded-full bg-emerald-500"></span>
					</span>
					<span class="text-xs font-medium text-emerald-500">Currently in Beta</span>
				</div>
			</div>

			<h1 class="mt-8 text-4xl font-semibold tracking-tight text-white sm:text-5xl md:text-6xl" in:fly={{ y: 20, duration: 500, delay: 400 }}>
				Business management,<br />
				<span class="text-slate-500">reimagined.</span>
			</h1>

			<p class="mt-5 max-w-md text-base text-slate-400" in:fly={{ y: 15, duration: 500, delay: 500 }}>
				A clean, focused platform for teams that value simplicity and speed.
			</p>

			<div class="mt-8 flex flex-col gap-3 sm:flex-row" in:fly={{ y: 15, duration: 500, delay: 600 }}>
				<Button size="lg" href="/sign-up" class="gap-2 bg-white px-8 text-slate-950 hover:bg-slate-200">
					Join the Beta
					<Icon icon="lucide:arrow-right" class="h-4 w-4" />
				</Button>
				<Button size="lg" variant="outline" href="/sign-in" class="border-slate-800 bg-transparent px-8 text-slate-300 hover:bg-slate-900 hover:text-white">
					Sign In
				</Button>
			</div>

			<div class="mt-12 flex flex-wrap items-center justify-center gap-8" in:fade={{ duration: 400, delay: 700 }}>
				<div class="flex items-center gap-2 text-sm text-slate-500">
					<Icon icon="lucide:check" class="h-4 w-4 text-emerald-500" />
					<span>Free during beta</span>
				</div>
				<div class="flex items-center gap-2 text-sm text-slate-500">
					<Icon icon="lucide:check" class="h-4 w-4 text-emerald-500" />
					<span>No credit card required</span>
				</div>
				<div class="flex items-center gap-2 text-sm text-slate-500">
					<Icon icon="lucide:check" class="h-4 w-4 text-emerald-500" />
					<span>Cancel anytime</span>
				</div>
			</div>
		{/if}
	</div>

	<footer class="px-6 py-5 lg:px-12">
		{#if mounted}
			<div class="flex items-center justify-center gap-6 text-xs text-slate-600" in:fade={{ duration: 400, delay: 800 }}>
				<span>© 2026 Artemis</span>
			</div>
		{/if}
	</footer>
</div>
