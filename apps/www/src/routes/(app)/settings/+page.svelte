<script lang="ts">
	import Icon from '@iconify/svelte';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Field from '$lib/components/ui/field/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';

	const user = {
		firstName: 'Alex',
		lastName: 'Morgan',
		email: 'alex.morgan@example.com',
		username: 'alexmorgan',
		bio: 'Product designer and developer based in San Francisco. Passionate about creating beautiful and functional user experiences.',
		location: 'San Francisco, CA',
		website: 'https://alexmorgan.dev',
		company: 'Artemis Inc.',
		role: 'Senior Product Designer',
		avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Alex'
	};

	let isLoading = $state(false);

	async function handleSubmit() {
		isLoading = true;
		await new Promise((resolve) => setTimeout(resolve, 1000));
		isLoading = false;
	}
</script>

<svelte:head>
	<title>Profile Settings - Artemis</title>
</svelte:head>

<div class="space-y-8">
	<div>
		<h1 class="text-2xl font-semibold tracking-tight">Profile</h1>
		<p class="text-sm text-muted-foreground">
			Manage your public profile and personal information.
		</p>
	</div>

	<Separator />

	<section class="space-y-4">
		<h2 class="text-lg font-medium">Avatar</h2>
		<div class="flex items-start gap-6">
			<div class="relative">
				<img
					src={user.avatar}
					alt="Profile"
					class="size-24 rounded-full bg-muted object-cover ring-2 ring-border"
				/>
				<button
					class="absolute -right-1 -bottom-1 flex size-8 items-center justify-center rounded-full bg-primary text-primary-foreground shadow-sm hover:bg-primary/90"
				>
					<Icon icon="lucide:camera" class="size-4" />
				</button>
			</div>
			<div class="space-y-2">
				<div class="flex gap-2">
					<Button variant="outline" size="sm">
						<Icon icon="lucide:upload" class="mr-2 size-4" />
						Upload new
					</Button>
					<Button
						variant="outline"
						size="sm"
						class="text-destructive hover:bg-destructive/10 hover:text-destructive"
					>
						<Icon icon="lucide:trash-2" class="mr-2 size-4" />
						Remove
					</Button>
				</div>
				<p class="text-xs text-muted-foreground">
					Recommended: Square JPG, PNG, or GIF, at least 1000x1000px.
				</p>
			</div>
		</div>
	</section>

	<Separator />

	<form
		onsubmit={(e) => {
			e.preventDefault();
			handleSubmit();
		}}
		class="space-y-6"
	>
		<section class="space-y-4">
			<h2 class="text-lg font-medium">Personal Information</h2>
			<div class="grid gap-4 sm:grid-cols-2">
				<Field.Field>
					<Field.Label for="firstName">First name</Field.Label>
					<Input id="firstName" value={user.firstName} placeholder="Enter your first name" />
				</Field.Field>

				<Field.Field>
					<Field.Label for="lastName">Last name</Field.Label>
					<Input id="lastName" value={user.lastName} placeholder="Enter your last name" />
				</Field.Field>
			</div>

			<Field.Field>
				<Field.Label for="username">Username</Field.Label>
				<div class="relative">
					<span class="absolute top-1/2 left-3 -translate-y-1/2 text-muted-foreground">@</span>
					<Input id="username" value={user.username} class="pl-7" placeholder="username" />
				</div>
				<Field.Description>This will be your public handle on Artemis.</Field.Description>
			</Field.Field>

			<Field.Field>
				<Field.Label for="email">Email address</Field.Label>
				<Input id="email" type="email" value={user.email} placeholder="you@example.com" />
				<Field.Description>We'll never share your email with anyone else.</Field.Description>
			</Field.Field>

			<Field.Field>
				<Field.Label for="bio">Bio</Field.Label>
				<textarea
					id="bio"
					rows="4"
					class="min-h-[100px] w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
					placeholder="Tell us a little about yourself...">{user.bio}</textarea
				>
				<Field.Description
					>Brief description for your profile. Maximum 500 characters.</Field.Description
				>
			</Field.Field>
		</section>

		<Separator />

		<section class="space-y-4">
			<h2 class="text-lg font-medium">Professional</h2>
			<div class="grid gap-4 sm:grid-cols-2">
				<Field.Field>
					<Field.Label for="company">Company</Field.Label>
					<Input id="company" value={user.company} placeholder="Your company name" />
				</Field.Field>

				<Field.Field>
					<Field.Label for="role">Job title</Field.Label>
					<Input id="role" value={user.role} placeholder="Your job title" />
				</Field.Field>
			</div>

			<Field.Field>
				<Field.Label for="location">Location</Field.Label>
				<Input id="location" value={user.location} placeholder="City, Country" />
			</Field.Field>

			<Field.Field>
				<Field.Label for="website">Website</Field.Label>
				<div class="relative">
					<span class="absolute top-1/2 left-3 -translate-y-1/2 text-muted-foreground">
						<Icon icon="lucide:globe" class="size-4" />
					</span>
					<Input
						id="website"
						type="url"
						value={user.website}
						class="pl-10"
						placeholder="https://your-website.com"
					/>
				</div>
			</Field.Field>
		</section>

		<Separator />

		<section class="space-y-4">
			<h2 class="text-lg font-medium">Social Links</h2>
			<div class="space-y-4">
				<Field.Field>
					<Field.Label for="twitter">Twitter</Field.Label>
					<div class="relative">
						<span class="absolute top-1/2 left-3 -translate-y-1/2 text-muted-foreground">
							<Icon icon="lucide:twitter" class="size-4" />
						</span>
						<Input id="twitter" placeholder="@username" class="pl-10" />
					</div>
				</Field.Field>

				<Field.Field>
					<Field.Label for="github">GitHub</Field.Label>
					<div class="relative">
						<span class="absolute top-1/2 left-3 -translate-y-1/2 text-muted-foreground">
							<Icon icon="lucide:github" class="size-4" />
						</span>
						<Input id="github" placeholder="username" class="pl-10" />
					</div>
				</Field.Field>

				<Field.Field>
					<Field.Label for="linkedin">LinkedIn</Field.Label>
					<div class="relative">
						<span class="absolute top-1/2 left-3 -translate-y-1/2 text-muted-foreground">
							<Icon icon="lucide:linkedin" class="size-4" />
						</span>
						<Input id="linkedin" placeholder="linkedin.com/in/username" class="pl-10" />
					</div>
				</Field.Field>
			</div>
		</section>

		<Separator />

		<section class="space-y-4">
			<div class="flex items-center justify-between rounded-lg border border-border p-4">
				<div class="space-y-0.5">
					<h3 class="font-medium">Public Profile</h3>
					<p class="text-sm text-muted-foreground">
						Allow others to see your profile and activity.
					</p>
				</div>
				<label class="relative inline-flex cursor-pointer items-center">
					<input type="checkbox" checked class="peer sr-only" />
					<div
						class="h-6 w-11 rounded-full bg-muted peer-checked:bg-primary peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:outline-none after:absolute after:top-[2px] after:left-[2px] after:h-5 after:w-5 after:rounded-full after:bg-white after:shadow after:transition-all after:content-[''] peer-checked:after:translate-x-5 dark:bg-input"
					></div>
				</label>
			</div>
		</section>

		<div class="flex items-center justify-end gap-4 pt-4">
			<Button type="button" variant="outline">Cancel</Button>
			<Button type="submit" disabled={isLoading}>
				{#if isLoading}
					<Icon icon="lucide:loader-2" class="mr-2 size-4 animate-spin" />
				{/if}
				Save changes
			</Button>
		</div>
	</form>
</div>
