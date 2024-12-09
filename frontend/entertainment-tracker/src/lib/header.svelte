<script lang="ts">
	import { logout as logoutRequest, getCurrentUser } from '$lib/user.api';
	import { goto } from '$app/navigation';
	import api from './api';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	$: currentUrlValue = $page.url.pathname;

	const isDashboardPage = () => currentUrlValue.includes('dashboard');
	const isAddMediaPage = () => currentUrlValue.includes('addmedia');

	let showLogout = false;

	let username = writable('User');

	onMount(() => {
		getCurrentUser().then((response) => {
			if (response.ok) {
				console.log(response);
				$username = response.body.user.name;
			} else {
				console.log('Failed to get current user');
			}
		});
	});

	function toggleLogout() {
		showLogout = !showLogout;
	}

	function hideLogout(event: MouseEvent) {
		if (typeof window === 'undefined') return;

		const profileButton = document.querySelector('.profile-picture-button')!;
		if (!profileButton.contains(event.target as Node)) {
			showLogout = false;
		}
	}

	function logout() {
		logoutRequest().then(async (response) => {
			api.setToken(null);
			if (response.ok) {
				console.log('Logout successful');
			} else {
				console.log('Logout failed');
			}
			await goto('/user/login');
		});
	}
</script>

<svelte:document on:click={hideLogout} />
<div class="w-100 bg-header relative flex items-center justify-between pl-3 pr-3 shadow-md">
	<div class="flex items-center">
		<a class="flex items-center" href="/works/dashboard">
			<div class="w-fit pl-2 pr-1 text-right text-white">
				<img src="/mediamindlogo.png" alt="MediaMind Logo" class="h-6 w-6" />
			</div>
			<div class="Ubuntu-font w-fit pr-8 text-left text-2xl text-white">MediaMind</div>
		</a>
		<a
			href="/works/dashboard"
			class="text-l w-fit p-2 text-white {isDashboardPage() ? 'font-bold' : ''} Ubuntu-font"
			>Dashboard</a
		>
		<a
			href="/works/addmedia"
			class="text-l w-fit p-2 text-white {isAddMediaPage() ? 'font-bold' : ''} Ubuntu-font"
			>Add Media</a
		>
	</div>
	<button
		class="bg-header profile-picture-button flex items-center rounded-full p-2 text-white"
		on:click|stopPropagation={toggleLogout}
		aria-expanded={showLogout}
	>
		<div class="flex items-center gap-4">
			<div>{$username}</div>
			<div class="leading-10 h-10 w-10 align-middle content-center rounded-full bg-gray-600">
				<div class="mb-0.5 text-lg uppercase">
					{$username
						.split(' ')
						.map((part) => (part.length > 0 ? part[0] : ''))
						.slice(0, 2)
						.join('')}
				</div>
			</div>
		</div>
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
