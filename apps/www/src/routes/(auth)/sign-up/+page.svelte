<script>
	import Icon from '@iconify/svelte';
	import { superForm } from 'sveltekit-superforms';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Field from '$lib/components/ui/field/index.js';
	import { Input } from '$lib/components/ui/input/index.js';

	const { data } = $props();

	const { form, errors, constraints, enhance } = superForm(data.form, {
		resetForm: true
	});
</script>

<div class="grid min-h-svh lg:grid-cols-2">
	<div class="flex flex-col gap-4 p-6 md:p-10">
		<div class="flex justify-center gap-2 md:justify-start">
			<a href="#/" class="flex items-center gap-2 text-xl font-bold">
				<div
					class="flex size-10 items-center justify-center rounded-md bg-primary text-primary-foreground"
				>
					<Icon icon="hugeicons:cowboy-hat" class="size-10" />
				</div>
				Artemis
			</a>
		</div>
		<div class="flex flex-1 items-center justify-center">
			<div class="w-full max-w-md">
				<form method="POST" use:enhance class="flex flex-col gap-6">
					<Field.Group>
						<div class="flex flex-col items-center gap-1 text-center">
							<h1 class="text-2xl font-bold">Create your account</h1>
							<p class="text-sm text-balance text-muted-foreground">
								Fill in the form below to create your account
							</p>
						</div>

						<Field.Field>
							<Field.Label for="name">Full Name</Field.Label>
							<Input
								id="name"
								name="name"
								type="text"
								placeholder="John Doe"
								bind:value={$form.name}
								{...$constraints.name}
							/>
							{#if $errors.name}
								<p class="text-[0.8rem] font-medium text-destructive">{$errors.name}</p>
							{/if}
						</Field.Field>

						<Field.Field>
							<Field.Label for="email">Email</Field.Label>
							<Input
								id="email"
								name="email"
								type="email"
								placeholder="m@example.com"
								bind:value={$form.email}
								{...$constraints.email}
							/>
							{#if $errors.email}
								<p class="text-[0.8rem] font-medium text-destructive">{$errors.email}</p>
							{:else}
								<Field.Description>We'll use this to contact you.</Field.Description>
							{/if}
						</Field.Field>

						<Field.Field>
							<Field.Label for="password">Password</Field.Label>
							<Input
								id="password"
								name="password"
								type="password"
								bind:value={$form.password}
								{...$constraints.password}
							/>
							{#if $errors.password}
								<p class="text-[0.8rem] font-medium text-destructive">{$errors.password}</p>
							{:else}
								<Field.Description>Must be at least 8 characters long.</Field.Description>
							{/if}
						</Field.Field>

						<Field.Field>
							<Field.Label for="confirm-password">Confirm Password</Field.Label>
							<Input
								id="confirm-password"
								name="confirmPassword"
								type="password"
								bind:value={$form.confirmPassword}
								{...$constraints.confirmPassword}
							/>
							{#if $errors.confirmPassword}
								<p class="text-[0.8rem] font-medium text-destructive">{$errors.confirmPassword}</p>
							{:else}
								<Field.Description>Please confirm your password.</Field.Description>
							{/if}
						</Field.Field>

						<Field.Field>
							<Button type="submit">Create Account</Button>
						</Field.Field>
					</Field.Group>
				</form>
			</div>
		</div>
	</div>
	<div class="relative hidden bg-muted lg:block">
		<img
			src="/placeholder.svg"
			alt=""
			class="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
		/>
	</div>
</div>
