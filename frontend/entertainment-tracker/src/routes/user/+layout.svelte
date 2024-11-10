<script>
	import { goto } from '$app/navigation';
	import { api } from '$lib/api';
	import { onMount } from 'svelte';

	let { children } = $props();

	onMount(async () => {
		if (api.validToken) {
			await goto('/works/dashboard');
		}
	});
</script>

{@render children()}
