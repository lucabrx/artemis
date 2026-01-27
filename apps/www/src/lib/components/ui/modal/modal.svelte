<script lang="ts">
	import { cn } from '$lib/utils.js';
	import type { Snippet } from 'svelte';

	interface Props {
		open?: boolean;
		onOpenChange?: (open: boolean) => void;
		title?: string;
		description?: string;
		children?: Snippet;
		footer?: Snippet;
		class?: string;
		size?: 'sm' | 'md' | 'lg' | 'xl' | 'full';
		showCloseButton?: boolean;
		closeOnOverlayClick?: boolean;
	}

	let {
		open = $bindable(false),
		onOpenChange,
		title,
		description,
		children,
		footer,
		class: className,
		size = 'md',
		showCloseButton = true,
		closeOnOverlayClick = true
	}: Props = $props();

	const sizeClasses = {
		sm: 'max-w-sm',
		md: 'max-w-lg',
		lg: 'max-w-2xl',
		xl: 'max-w-4xl',
		full: 'max-w-[calc(100vw-2rem)]'
	};

	function handleClose() {
		open = false;
		onOpenChange?.(false);
	}

	function handleOverlayClick(e: MouseEvent) {
		if (closeOnOverlayClick && e.target === e.currentTarget) {
			handleClose();
		}
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Escape' && open) {
			handleClose();
		}
	}

	$effect(() => {
		if (open) {
			document.body.style.overflow = 'hidden';
		} else {
			document.body.style.overflow = '';
		}
		return () => {
			document.body.style.overflow = '';
		};
	});
</script>

<svelte:window onkeydown={handleKeydown} />

{#if open}
	<div class="fixed inset-0 z-50 flex items-center justify-center">
		<div
			class="absolute inset-0 bg-background/80 backdrop-blur-sm transition-opacity"
			onclick={handleOverlayClick}
			aria-hidden="true"
		></div>

		<div
			class={cn(
				'relative z-50 w-full scale-100 transform rounded-xl border border-border bg-card p-0 shadow-2xl transition-all',
				sizeClasses[size],
				className
			)}
			role="dialog"
			aria-modal="true"
			aria-labelledby={title ? 'modal-title' : undefined}
			aria-describedby={description ? 'modal-description' : undefined}
		>
			{#if title || showCloseButton}
				<div class="flex items-center justify-between border-b border-border px-6 py-4">
					<div class="space-y-1">
						{#if title}
							<h2 id="modal-title" class="text-lg font-semibold">{title}</h2>
						{/if}
						{#if description}
							<p id="modal-description" class="text-sm text-muted-foreground">{description}</p>
						{/if}
					</div>
					{#if showCloseButton}
						<button
							type="button"
							class="rounded-md p-1 text-muted-foreground transition-colors hover:bg-muted hover:text-foreground"
							onclick={handleClose}
							aria-label="Close modal"
						>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="20"
								height="20"
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
							>
								<path d="M18 6 6 18" />
								<path d="m6 6 12 12" />
							</svg>
						</button>
					{/if}
				</div>
			{/if}

			<div class={cn('px-6 py-4', !title && !showCloseButton && 'pt-6')}>
				{@render children?.()}
			</div>

			{#if footer}
				<div class="flex items-center justify-end gap-2 border-t border-border px-6 py-4">
					{@render footer()}
				</div>
			{/if}
		</div>
	</div>
{/if}
