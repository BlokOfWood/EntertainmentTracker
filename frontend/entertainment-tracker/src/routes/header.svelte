<script lang="ts">
    import { writable } from 'svelte/store';
    import { onMount } from 'svelte';

    const initialUrl = writable(typeof window !== 'undefined' ? window.location.href : '');
    
    // Reactive statement
    $: currentUrlValue = $initialUrl;

    const isDashboardPage = () => {
        return currentUrlValue.length > 0 && currentUrlValue.includes('dashboard');
    };

    const isAddMediaPage = () => {
        return currentUrlValue.length > 0 && currentUrlValue.includes('addmedia');
    };

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

<div class="flex justify-between items-center w-100 bg-header shadow-md pr-3 pl-3">
    <div class="flex items-center">
        <div class="border-1 border-solid border-red-600 pl-2 pr-1 w-fit text-white text-right">
            <img src="/mediamindlogo.png" alt="MediaMind Logo" class="w-6 h-6" />
        </div>
        <div class="border-1 border-solid border-red-600 pr-8 w-fit text-white text-2xl text-left Ubuntu-font">MediaMind</div>
        <a href="/dashboard" class="border-1 border-solid border-red-600 p-2 w-fit text-white text-l {isDashboardPage() ? 'font-bold' : ''} Ubuntu-font">Dashboard</a>
        <a href="/addmedia" class="border-1 border-solid border-red-600 p-2 w-fit text-white text-l {isAddMediaPage() ? 'font-bold' : ''} Ubuntu-font">Add Media</a>
    </div>
    <button class="bg-header text-white p-2 rounded-full flex items-center profile-picture-button" on:click|stopPropagation={toggleLogout} aria-expanded={showLogout}>
        <img src="/profilepicture.png" alt="Profile" class="rounded-full w-10 h-10"/>
    </button>
</div>
{#if showLogout}
<div class="flex justify-end">
    <a href="/" class="w-fit bg-header pr-3 pl-3 pt-2 pb-2 rounded-br-md rounded-bl-md text-white shadow-md Ubuntu-font">
        Logout
    </a>
</div>
{/if}