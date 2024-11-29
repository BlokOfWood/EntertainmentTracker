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
	<div class="rounded bg-white p-6 h-full">
		<div class="mb-6 flex justify-center gap-6 text-xl">
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

		<div class="Inter-font">
			{@render children()}
		</div>
	</div>
</div>
