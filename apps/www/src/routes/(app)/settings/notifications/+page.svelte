<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';

	let emailNotifications = $state({
		marketing: true,
		social: true,
		updates: true,
		security: true,
		newsletter: false
	});

	let pushNotifications = $state({
		mentions: true,
		messages: true,
		comments: true,
		reminders: false
	});

	let notificationFrequency = $state('realtime');
</script>

<svelte:head>
	<title>Notification Settings - Artemis</title>
</svelte:head>

<div class="space-y-8">
	<div>
		<h1 class="text-2xl font-semibold tracking-tight">Notifications</h1>
		<p class="text-sm text-muted-foreground">
			Choose what notifications you want to receive and how you want to receive them.
		</p>
	</div>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-center gap-3">
			<div class="flex size-10 items-center justify-center rounded-lg bg-primary/10">
				<Icon icon="lucide:mail" class="size-5 text-primary" />
			</div>
			<div>
				<h2 class="text-lg font-medium">Email Notifications</h2>
				<p class="text-sm text-muted-foreground">Control which emails you receive from us.</p>
			</div>
		</div>

		<div class="space-y-3">
			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Marketing & Promotions</h3>
					<p class="text-sm text-muted-foreground">
						Receive emails about new features, offers, and promotions.
					</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input type="checkbox" bind:checked={emailNotifications.marketing} class="peer sr-only" />
					<div
						class="h-6 w-11 rounded-full bg-muted peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-[''] peer-checked:after:translate-x-5 dark:bg-input"
					></div>
				</label>
			</div>

			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Social Activity</h3>
					<p class="text-sm text-muted-foreground">
						Get notified when someone follows you, mentions you, or interacts with your content.
					</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input type="checkbox" bind:checked={emailNotifications.social} class="peer sr-only" />
					<div
						class="h-6 w-11 rounded-full bg-muted peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-[''] peer-checked:after:translate-x-5 dark:bg-input"
					></div>
				</label>
			</div>

			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Product Updates</h3>
					<p class="text-sm text-muted-foreground">
						Receive emails about new features and improvements.
					</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input type="checkbox" bind:checked={emailNotifications.updates} class="peer sr-only" />
					<div
						class="h-6 w-11 rounded-full bg-muted peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-[''] peer-checked:after:translate-x-5 dark:bg-input"
					></div>
				</label>
			</div>

			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Security Alerts</h3>
					<p class="text-sm text-muted-foreground">
						Get important security notifications about your account.
					</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input
						type="checkbox"
						bind:checked={emailNotifications.security}
						class="peer sr-only"
						disabled
					/>
					<div
						class="h-6 w-11 rounded-full bg-primary opacity-50 after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:translate-x-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-['']"
					></div>
				</label>
			</div>

			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Newsletter</h3>
					<p class="text-sm text-muted-foreground">
						Receive our weekly newsletter with industry news and tips.
					</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input
						type="checkbox"
						bind:checked={emailNotifications.newsletter}
						class="peer sr-only"
					/>
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
				<Icon icon="lucide:bell-ring" class="size-5 text-primary" />
			</div>
			<div>
				<h2 class="text-lg font-medium">Push Notifications</h2>
				<p class="text-sm text-muted-foreground">
					Manage your browser and mobile push notifications.
				</p>
			</div>
		</div>

		<div class="space-y-3">
			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Mentions & Tags</h3>
					<p class="text-sm text-muted-foreground">
						Get notified when someone mentions or tags you.
					</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input type="checkbox" bind:checked={pushNotifications.mentions} class="peer sr-only" />
					<div
						class="h-6 w-11 rounded-full bg-muted peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-[''] peer-checked:after:translate-x-5 dark:bg-input"
					></div>
				</label>
			</div>

			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Direct Messages</h3>
					<p class="text-sm text-muted-foreground">
						Receive notifications for new direct messages.
					</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input type="checkbox" bind:checked={pushNotifications.messages} class="peer sr-only" />
					<div
						class="h-6 w-11 rounded-full bg-muted peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-[''] peer-checked:after:translate-x-5 dark:bg-input"
					></div>
				</label>
			</div>

			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Comments</h3>
					<p class="text-sm text-muted-foreground">
						Get notified when someone comments on your posts.
					</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input type="checkbox" bind:checked={pushNotifications.comments} class="peer sr-only" />
					<div
						class="h-6 w-11 rounded-full bg-muted peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-[''] peer-checked:after:translate-x-5 dark:bg-input"
					></div>
				</label>
			</div>

			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Reminders</h3>
					<p class="text-sm text-muted-foreground">
						Receive reminder notifications for upcoming events.
					</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input type="checkbox" bind:checked={pushNotifications.reminders} class="peer sr-only" />
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
				<Icon icon="lucide:clock" class="size-5 text-primary" />
			</div>
			<div>
				<h2 class="text-lg font-medium">Email Digest Frequency</h2>
				<p class="text-sm text-muted-foreground">
					Choose how often you want to receive email digests.
				</p>
			</div>
		</div>

		<div class="grid gap-3 sm:grid-cols-3">
			<label class="cursor-pointer">
				<input
					type="radio"
					name="frequency"
					value="realtime"
					bind:group={notificationFrequency}
					class="peer sr-only"
				/>
				<div
					class="rounded-lg border-2 border-border p-4 transition-all peer-checked:border-primary peer-checked:bg-primary/5"
				>
					<div class="flex items-center gap-2">
						<Icon
							icon="lucide:zap"
							class="size-5 text-muted-foreground peer-checked:text-primary"
						/>
						<span class="font-medium">Real-time</span>
					</div>
					<p class="mt-1 text-xs text-muted-foreground">Receive emails immediately</p>
				</div>
			</label>

			<label class="cursor-pointer">
				<input
					type="radio"
					name="frequency"
					value="daily"
					bind:group={notificationFrequency}
					class="peer sr-only"
				/>
				<div
					class="rounded-lg border-2 border-border p-4 transition-all peer-checked:border-primary peer-checked:bg-primary/5"
				>
					<div class="flex items-center gap-2">
						<Icon
							icon="lucide:sun"
							class="size-5 text-muted-foreground peer-checked:text-primary"
						/>
						<span class="font-medium">Daily</span>
					</div>
					<p class="mt-1 text-xs text-muted-foreground">One email per day</p>
				</div>
			</label>

			<label class="cursor-pointer">
				<input
					type="radio"
					name="frequency"
					value="weekly"
					bind:group={notificationFrequency}
					class="peer sr-only"
				/>
				<div
					class="rounded-lg border-2 border-border p-4 transition-all peer-checked:border-primary peer-checked:bg-primary/5"
				>
					<div class="flex items-center gap-2">
						<Icon
							icon="lucide:calendar"
							class="size-5 text-muted-foreground peer-checked:text-primary"
						/>
						<span class="font-medium">Weekly</span>
					</div>
					<p class="mt-1 text-xs text-muted-foreground">One email per week</p>
				</div>
			</label>
		</div>
	</section>

	<Separator />

	<section class="space-y-4">
		<div class="flex items-start justify-between">
			<div class="space-y-1">
				<h2 class="text-lg font-medium">Quiet Hours</h2>
				<p class="text-sm text-muted-foreground">Pause notifications during specific hours.</p>
			</div>
			<label class="relative inline-flex cursor-pointer items-center">
				<input type="checkbox" class="peer sr-only" />
				<div
					class="h-6 w-11 rounded-full bg-muted peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-[''] peer-checked:after:translate-x-5 dark:bg-input"
				></div>
			</label>
		</div>

		<div class="grid gap-4 sm:grid-cols-2">
			<div class="space-y-2">
				<label for="startTime" class="text-sm font-medium">Start time</label>
				<input
					id="startTime"
					type="time"
					value="22:00"
					class="h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-xs transition-colors focus-visible:border-ring focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				/>
			</div>
			<div class="space-y-2">
				<label for="endTime" class="text-sm font-medium">End time</label>
				<input
					id="endTime"
					type="time"
					value="08:00"
					class="h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-xs transition-colors focus-visible:border-ring focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				/>
			</div>
		</div>
	</section>

	<div class="flex items-center justify-end gap-4 pt-4">
		<Button type="button" variant="outline">Reset to defaults</Button>
		<Button type="submit">Save preferences</Button>
	</div>
</div>
