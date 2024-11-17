<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import '../app.css';
	import { api } from '$lib/api';
	import { goto } from '$app/navigation';

	let { children } = $props();

	onMount(async () => {
		if ($page.url.pathname === '/') {
			if (!api.validToken) {
				await goto('/user/login');
			} else {
				await goto('/works/dashboard');
			}
		}
	});
</script>

{@render children()}
