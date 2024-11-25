<script>
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { api } from '$lib/api';
	import { onMount } from 'svelte';

	onMount(async () => {
		if ($page.status === 404) {
			if (api.validToken) {
				await goto('/works/dashboard');
			} else {
				await goto('/user/login');
			}
		}
	});
</script>
