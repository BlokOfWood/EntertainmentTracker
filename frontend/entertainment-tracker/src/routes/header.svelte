<script lang="ts">
    import { writable } from 'svelte/store';
    import { onMount } from 'svelte';

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

    onMount(() => {
        document.addEventListener('click', hideLogout);
    });
</script>

<svelte:head>
    <link href="https://fonts.googleapis.com/css2?family=Ubuntu:wght@400;500;700&display=swap" rel="stylesheet">
</svelte:head>

<style>
    .Ubuntu-font {
        font-family: 'Ubuntu', sans-serif;
    }
</style>

<div class="relative flex justify-between items-center w-100 bg-header shadow-md pr-3 pl-3">
    <div class="flex items-center">
        <div class="pl-2 pr-1 w-fit text-white text-right">
            <img src="/mediamindlogo.png" alt="MediaMind Logo" class="w-6 h-6" />
        </div>
        <div class="pr-8 w-fit text-white text-2xl text-left Ubuntu-font">MediaMind</div>
        <a href="/dashboard" class="p-2 w-fit text-white text-l {isDashboardPage() ? 'font-bold' : ''} Ubuntu-font">Dashboard</a>
        <a href="/addmedia" class="p-2 w-fit text-white text-l {isAddMediaPage() ? 'font-bold' : ''} Ubuntu-font">Add Media</a>
    </div>
    <button class="bg-header text-white p-2 rounded-full flex items-center profile-picture-button" on:click|stopPropagation={toggleLogout} aria-expanded={showLogout}>
        <img src="/profilepicture.png" alt="Profile" class="rounded-full w-10 h-10"/>
    </button>
    {#if showLogout}
    <!--TODO: Add logout functionality-->
    <a href="/" class="absolute right-0 top-full z-50 block w-fit bg-header pr-3 pl-3 pt-2 pb-2 rounded-br-md rounded-bl-md text-white shadow-md Ubuntu-font text-right">
        Logout
    </a>
    {/if}
</div>
