<script lang="ts">
	import Icon from '@iconify/svelte';
	import { mode, toggleMode } from 'mode-watcher';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';

	const isLight = $derived(mode.current === 'light');
	const isDark = $derived(mode.current === 'dark');

	const accentColor = $state('emerald');
	const fontSize = $state('default');
	const sidebarCollapsed = $state(false);
	const reducedMotion = $state(false);
	const highContrast = $state(false);

	const accentColors = [
		{ name: 'Emerald', value: 'emerald', class: 'bg-emerald-500' },
		{ name: 'Blue', value: 'blue', class: 'bg-blue-500' },
		{ name: 'Purple', value: 'purple', class: 'bg-purple-500' },
		{ name: 'Rose', value: 'rose', class: 'bg-rose-500' },
		{ name: 'Orange', value: 'orange', class: 'bg-orange-500' },
		{ name: 'Cyan', value: 'cyan', class: 'bg-cyan-500' }
	];

	const fontSizes = [
		{ name: 'Small', value: 'small', description: '90% of default size' },
		{ name: 'Default', value: 'default', description: '100% of default size' },
		{ name: 'Large', value: 'large', description: '110% of default size' },
		{ name: 'Extra Large', value: 'xl', description: '125% of default size' }
	];
</script>

<svelte:head>
	<title>Appearance Settings - Artemis</title>
</svelte:head>

<div class="space-y-8">
	<div>
		<h1 class="text-2xl font-semibold tracking-tight">Appearance</h1>
		<p class="text-sm text-muted-foreground">Customize how Artemis looks and feels for you.</p>
	</div>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-center gap-3">
			<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
				<Icon icon="lucide:palette" class="size-5 text-primary" />
			</div>
			<div>
				<h2 class="text-lg font-medium">Theme</h2>
				<p class="text-sm text-muted-foreground">
					Choose between light and dark mode, or sync with your system.
				</p>
			</div>
		</div>

		<div class="grid gap-4 sm:grid-cols-3">
			<button
				class="group relative flex flex-col items-center gap-3 rounded-lg border-2 {isLight
					? 'border-primary bg-primary/5'
					: 'border-border'} p-4 transition-all hover:border-primary/50 focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none"
				onclick={() => {
					if (mode.current === 'dark') toggleMode();
				}}
			>
				<div class="aspect-video w-full overflow-hidden rounded-md bg-slate-100 shadow-sm">
					<div class="h-full w-full bg-gradient-to-br from-slate-50 to-slate-200 p-3">
						<div class="h-2 w-12 rounded bg-slate-300"></div>
						<div class="mt-2 h-2 w-8 rounded bg-slate-300"></div>
						<div class="mt-4 flex gap-2">
							<div class="h-8 w-8 rounded bg-slate-300"></div>
							<div class="h-8 flex-1 rounded bg-slate-300"></div>
						</div>
					</div>
				</div>
				<div class="flex items-center gap-2">
					{#if isLight}
						<Icon icon="lucide:check-circle" class="size-4 text-primary" />
					{/if}
					<span class="font-medium">Light</span>
				</div>
			</button>

			<button
				class="group relative flex flex-col items-center gap-3 rounded-lg border-2 {isDark
					? 'border-primary bg-primary/5'
					: 'border-border'} p-4 transition-all hover:border-primary/50 focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none"
				onclick={() => {
					if (mode.current !== 'dark') toggleMode();
				}}
			>
				<div class="aspect-video w-full overflow-hidden rounded-md bg-slate-900 shadow-sm">
					<div class="h-full w-full bg-gradient-to-br from-slate-800 to-slate-950 p-3">
						<div class="h-2 w-12 rounded bg-slate-600"></div>
						<div class="mt-2 h-2 w-8 rounded bg-slate-600"></div>
						<div class="mt-4 flex gap-2">
							<div class="h-8 w-8 rounded bg-slate-600"></div>
							<div class="h-8 flex-1 rounded bg-slate-600"></div>
						</div>
					</div>
				</div>
				<div class="flex items-center gap-2">
					{#if isDark}
						<Icon icon="lucide:check-circle" class="size-4 text-primary" />
					{/if}
					<span class="font-medium">Dark</span>
				</div>
			</button>

			<button
				class="group relative flex flex-col items-center gap-3 rounded-lg border-2 border-border p-4 transition-all hover:border-primary/50 focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none"
			>
				<div class="aspect-video w-full overflow-hidden rounded-md shadow-sm">
					<div class="grid h-full w-full grid-cols-2">
						<div class="bg-gradient-to-br from-slate-50 to-slate-200 p-2">
							<div class="h-1.5 w-8 rounded bg-slate-300"></div>
							<div class="mt-1 h-1.5 w-5 rounded bg-slate-300"></div>
						</div>
						<div class="bg-gradient-to-br from-slate-800 to-slate-950 p-2">
							<div class="h-1.5 w-8 rounded bg-slate-600"></div>
							<div class="mt-1 h-1.5 w-5 rounded bg-slate-600"></div>
						</div>
					</div>
				</div>
				<div class="flex items-center gap-2">
					<span class="font-medium">System</span>
				</div>
			</button>
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-center gap-3">
			<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
				<Icon icon="lucide:paintbrush" class="size-5 text-primary" />
			</div>
			<div>
				<h2 class="text-lg font-medium">Accent Color</h2>
				<p class="text-sm text-muted-foreground">
					Choose your preferred accent color for buttons and highlights.
				</p>
			</div>
		</div>

		<div class="flex flex-wrap gap-3">
			{#each accentColors as color}
				<button
					class="group relative flex size-12 items-center justify-center rounded-full {color.class} transition-all hover:scale-110 focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 focus-visible:outline-none {accentColor ===
					color.value
						? 'ring-2 ring-current ring-offset-2'
						: ''}"
					onclick={() => (accentColor = color.value)}
					aria-label={`Select ${color.name} accent color`}
				>
					{#if accentColor === color.value}
						<Icon icon="lucide:check" class="size-5 text-white" />
					{/if}
				</button>
			{/each}
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-center gap-3">
			<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
				<Icon icon="lucide:type" class="size-5 text-primary" />
			</div>
			<div>
				<h2 class="text-lg font-medium">Font Size</h2>
				<p class="text-sm text-muted-foreground">
					Adjust the text size throughout the application.
				</p>
			</div>
		</div>

		<div class="space-y-2">
			{#each fontSizes as size}
				<label
					class="flex cursor-pointer items-center justify-between rounded-lg border border-border p-4 transition-colors hover:bg-muted/50"
				>
					<div class="flex items-center gap-3">
						<input
							type="radio"
							name="fontSize"
							value={size.value}
							bind:group={fontSize}
							class="size-4 text-primary focus:ring-primary"
						/>
						<div>
							<span class="font-medium">{size.name}</span>
							<p class="text-sm text-muted-foreground">{size.description}</p>
						</div>
					</div>
					<span
						class="text-muted-foreground"
						style="font-size: {size.value === 'small'
							? '12px'
							: size.value === 'large'
								? '18px'
								: size.value === 'xl'
									? '20px'
									: '16px'}"
					>
						Aa
					</span>
				</label>
			{/each}
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-center gap-3">
			<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
				<Icon icon="lucide:layout" class="size-5 text-primary" />
			</div>
			<div>
				<h2 class="text-lg font-medium">Interface</h2>
				<p class="text-sm text-muted-foreground">Customize your interface layout and behavior.</p>
			</div>
		</div>

		<div class="space-y-3">
			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Compact Sidebar</h3>
					<p class="text-sm text-muted-foreground">Use a more compact sidebar layout.</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input type="checkbox" bind:checked={sidebarCollapsed} class="peer sr-only" />
					<div
						class="h-6 w-11 rounded-full bg-muted peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-[''] peer-checked:after:translate-x-5 dark:bg-input"
					></div>
				</label>
			</div>

			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Reduced Motion</h3>
					<p class="text-sm text-muted-foreground">Minimize animations throughout the interface.</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input type="checkbox" bind:checked={reducedMotion} class="peer sr-only" />
					<div
						class="h-6 w-11 rounded-full bg-muted peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-[''] peer-checked:after:translate-x-5 dark:bg-input"
					></div>
				</label>
			</div>

			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">High Contrast</h3>
					<p class="text-sm text-muted-foreground">Increase contrast for better visibility.</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input type="checkbox" bind:checked={highContrast} class="peer sr-only" />
					<div
						class="h-6 w-11 rounded-full bg-muted peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-[''] peer-checked:after:translate-x-5 dark:bg-input"
					></div>
				</label>
			</div>
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-center gap-3">
			<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
				<Icon icon="lucide:globe" class="size-5 text-primary" />
			</div>
			<div>
				<h2 class="text-lg font-medium">Language & Region</h2>
				<p class="text-sm text-muted-foreground">
					Choose your preferred language and regional settings.
				</p>
			</div>
		</div>

		<div class="grid gap-4 sm:grid-cols-2">
			<div class="space-y-2">
				<label class="text-sm font-medium" for="language">Language</label>
				<select
					id="language"
					class="h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-xs transition-colors focus-visible:border-ring focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				>
					<option value="en">English (US)</option>
					<option value="en-gb">English (UK)</option>
					<option value="es">Español</option>
					<option value="fr">Français</option>
					<option value="de">Deutsch</option>
					<option value="it">Italiano</option>
					<option value="pt">Português</option>
					<option value="ja">日本語</option>
					<option value="ko">한국어</option>
					<option value="zh">中文</option>
				</select>
			</div>

			<div class="space-y-2">
				<label class="text-sm font-medium" for="timezone">Timezone</label>
				<select
					id="timezone"
					class="h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-xs transition-colors focus-visible:border-ring focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				>
					<option value="utc">UTC</option>
					<option value="est">Eastern Time (ET)</option>
					<option value="cst">Central Time (CT)</option>
					<option value="mst">Mountain Time (MT)</option>
					<option value="pst" selected>Pacific Time (PT)</option>
					<option value="gmt">Greenwich Mean Time (GMT)</option>
					<option value="cet">Central European Time (CET)</option>
					<option value="jst">Japan Standard Time (JST)</option>
				</select>
			</div>
		</div>

		<div class="grid gap-4 sm:grid-cols-2">
			<div class="space-y-2">
				<label class="text-sm font-medium" for="dateFormat">Date Format</label>
				<select
					id="dateFormat"
					class="h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-xs transition-colors focus-visible:border-ring focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				>
					<option value="mdy">MM/DD/YYYY</option>
					<option value="dmy">DD/MM/YYYY</option>
					<option value="ymd">YYYY/MM/DD</option>
				</select>
			</div>

			<div class="space-y-2">
				<label class="text-sm font-medium" for="timeFormat">Time Format</label>
				<select
					id="timeFormat"
					class="h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-xs transition-colors focus-visible:border-ring focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				>
					<option value="12h">12-hour (AM/PM)</option>
					<option value="24h">24-hour</option>
				</select>
			</div>
		</div>
	</section>

	<div class="flex items-center justify-end gap-4 pt-4">
		<Button type="button" variant="outline">Reset to defaults</Button>
		<Button type="submit">Save changes</Button>
	</div>
</div>
