<script>
    import { Button } from "$lib/components/ui/button";
</script>

<h1>Welcome to SvelteKit</h1>
<p>
    Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the
    documentation
</p>
<div class="grid gap-6 mt-4">
    <div class="flex gap-4">
        <Button variant="destructive">Click me</Button>
        <Button variant="outline">Click me</Button>
        <Button variant="link">Click me</Button>
        <Button variant="secondary">Click me</Button>
        <Button>Click me</Button>
        <Button variant="link">Click me</Button>
    </div>

    <div class="flex gap-4">
        <Button size="sm" variant="destructive">Click me</Button>
        <Button size="sm" variant="outline">Click me</Button>
        <Button size="sm" variant="link">Click me</Button>
        <Button size="sm" variant="secondary">Click me</Button>
        <Button size="sm">Click me</Button>
        <Button size="sm" variant="link">Click me</Button>
    </div>
</div>
