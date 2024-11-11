<script>
	import { goto } from '$app/navigation';
	import { api } from '$lib/api';
	import { onMount } from 'svelte';
	import Header from '$lib/header.svelte';

	let { children } = $props();

	onMount(async () => {
		if (api.validToken) {
			await goto('/works/dashboard');
		}
	});
</script>

<div class="bg-background relative flex h-screen flex-col">
	<Header />
	{@render children()}
</div>
