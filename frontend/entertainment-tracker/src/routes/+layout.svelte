<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import '../app.css';
	import { api } from '$lib/api';
	import { goto } from '$app/navigation';
	import { get } from 'svelte/store';

	let { children } = $props();

	onMount(async () => {
		if (!get(api).validToken) {
			await goto('/user/login');
		} else if ($page.url.pathname === '/') {
			await goto('/works/dashboard');
		}
	});
</script>

{@render children()}
