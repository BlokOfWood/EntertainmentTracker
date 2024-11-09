<script lang="ts">
	import { writable } from 'svelte/store';
	import { onMount } from 'svelte';
	import { logout as logoutRequest } from '$lib/user.api';
	import { redirect } from '@sveltejs/kit';

	const initialUrl = writable(typeof window !== 'undefined' ? window.location.href : '');

	$: currentUrlValue = $initialUrl;

	const isDashboardPage = () => currentUrlValue.includes('dashboard');
	const isAddMediaPage = () => currentUrlValue.includes('addmedia');

	let showLogout = false;

	function toggleLogout() {
		showLogout = !showLogout;
	}

	function hideLogout(event: MouseEvent) {
		const profileButton = document.querySelector('.profile-picture-button')!;
		if (!profileButton.contains(event.target as Node)) {
			showLogout = false;
		}
	}

	function logout() {
		logoutRequest().then((response) => {
			if (response.ok) {
				console.log('Logout successful');
				redirect(302, '/login');
			} else {
				console.log('Logout failed');
			}
		});
	}

	onMount(() => {
		document.addEventListener('click', hideLogout);
	});
</script>

<div class="w-100 bg-header relative flex items-center justify-between pl-3 pr-3 shadow-md">
	<div class="flex items-center">
		<div class="w-fit pl-2 pr-1 text-right text-white">
			<img src="/mediamindlogo.png" alt="MediaMind Logo" class="h-6 w-6" />
		</div>
		<div class="Ubuntu-font w-fit pr-8 text-left text-2xl text-white">MediaMind</div>
		<a
			href="/dashboard"
			class="text-l w-fit p-2 text-white {isDashboardPage() ? 'font-bold' : ''} Ubuntu-font"
			>Dashboard</a
		>
		<a
			href="/addmedia"
			class="text-l w-fit p-2 text-white {isAddMediaPage() ? 'font-bold' : ''} Ubuntu-font"
			>Add Media</a
		>
	</div>
	<button
		class="bg-header profile-picture-button flex items-center rounded-full p-2 text-white"
		on:click|stopPropagation={toggleLogout}
		aria-expanded={showLogout}
	>
		<img src="/profilepicture.png" alt="Profile" class="h-10 w-10 rounded-full" />
	</button>
	{#if showLogout}
		<button
			on:click={logout}
			class="bg-header Ubuntu-font absolute right-0 top-full z-50 block w-fit rounded-bl-md rounded-br-md pb-2 pl-3 pr-3 pt-2 text-right text-white shadow-md"
		>
			Logout
		</button>
	{/if}
</div>
