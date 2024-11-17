<script>
	import { goto } from '$app/navigation';
	import { api } from '$lib/api';
	import { onMount } from 'svelte';
	import Header from '$lib/header.svelte';
	import { page } from '$app/stores';

	let { children } = $props();

	onMount(async () => {
		if (api.validToken && $page.url.pathname === '/') {
			await goto('/works/dashboard');
		} else {
			await goto('/user/login');
		}
	});
</script>

<div class="bg-background relative flex h-screen flex-col">
	<Header />
	<div class="flex-grow overflow-auto">
		{@render children()}
	</div>
</div>
