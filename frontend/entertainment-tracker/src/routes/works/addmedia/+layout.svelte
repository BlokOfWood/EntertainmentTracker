<script>
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	let { children } = $props();

	onMount(async () => {
		if ($page.url.pathname === '/works/addmedia') {
			await goto('/works/addmedia/books');
		}
	});
</script>

<div class="mx-auto w-full h-full max-w-screen-md py-3">
	<div class="rounded bg-white py-4 h-full grid grid-rows-[auto_1fr] gap-6">
		<div class="flex justify-center gap-6 text-xl">
			<a
				href="/works/addmedia/books"
				class="Ubuntu-font active:font-bold
		{$page.url.pathname.endsWith('addmedia/books') ? 'font-bold' : ''}
		">Books</a
			>
			<a
				href="/works/addmedia/tvseries"
				class="Ubuntu-font active:font-bold
		{$page.url.pathname.endsWith('addmedia/tvseries') ? 'font-bold' : ''}
		">TV Series</a
			>
			<a
				href="/works/addmedia/movies"
				class="Ubuntu-font active:font-bold
		{$page.url.pathname.endsWith('addmedia/movies') ? 'font-bold' : ''}
		">Movies</a
			>

			<a
				href="/works/addmedia/youtube"
				class="Ubuntu-font active:font-bold
		{$page.url.pathname.endsWith('addmedia/youtube') ? 'font-bold' : ''}
		">YouTube</a
			>
		</div>

		<div class="Inter-font overflow-auto h-full px-6">
			{@render children()}
		</div>
	</div>
</div>
