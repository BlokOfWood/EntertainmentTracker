<script lang="ts">
	import { goto } from '$app/navigation';
	import { getBookByISBN, searchBooksByTitle } from '$lib/addmedia.api';
	import type { BookSearchResponse } from '$lib/api.model';
	import { createWork } from '$lib/works.api';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	let query = '';
	let displayedBooks = writable(undefined as BookSearchResponse[] | null | undefined);
	const isbnRegex =
		/^(?:ISBN(?:-13)?:?\ )?(?=[0-9]{13}$|(?=(?:[0-9]+[-\ ]){4})[-\ 0-9]{17}$)97[89][-\ ]?[0-9]{1,5}[-\ ]?[0-9]+[-\ ]?[0-9]+[-\ ]?[0-9]$/;

	let disabled = false;

	async function searchBooks() {
		let response: BookSearchResponse[] = [];

		if (isbnRegex.test(query)) {
			response = [(await getBookByISBN(query)).body.book];
		} else response = (await searchBooksByTitle(query)).body.books;

		displayedBooks.set(response);
	}

	onMount(async () => {
		disabled = false;
	});

	async function addBook(idx: number) {
		const books = $displayedBooks;
		if (books === null) return;

		disabled = true;
		const addResult = await createWork({
			title: books[idx].title,
			type: 'book',
			current_progress: 0,
			target_progress: books[idx].page_count,
			third_party_id: books[idx].id,
			status: 'watching'
		});

		if (addResult.statusCode < 400) {
			await goto('/works/dashboard');
		}
	}
</script>

<div>
	<form on:submit={searchBooks} class="flex flex-col items-center gap-6">
		<div class="w-fit text-left">
			<label class="Inter-font block text-sm" for="query-field">Title or ISBN</label>
			<input
				bind:value={query}
				id="query-field"
				class="mt-1"
				type="text"
				placeholder="Book Title or ISBN"
			/>
		</div>
		<button class="fancy-button mb-6">Search</button>
	</form>

	<div
		class="grid grid-cols-[auto_1fr_1fr_1fr_1fr_auto] items-center gap-y-4 overflow-auto text-center text-sm"
	>
		{#if $displayedBooks !== undefined}
			<div>Thumbnail</div>
			<div>Title</div>
			<div>Author</div>
			<div>ISBN</div>
			<div>Page count</div>
			<div></div>
			{#if $displayedBooks === null}
				<div class="col-span-6">No books found</div>
			{:else}
				{#each $displayedBooks as book, idx}
					<div class="text-center">
						<img src={book.thumbnail} alt={book.title} class="h-24 w-16" />
					</div>
					<div>{book.title}</div>
					<div>{book.author}</div>
					<div>{book.isbn}</div>
					<div>{book.page_count}</div>
					<div class="text-center">
						<button
							class="h-fit w-fit text-center disabled:opacity-35"
							{disabled}
							on:click={() => addBook(idx)}
							aria-label="Add"
							><img src="/plus-square.svg" alt="plus sign" />
						</button>
					</div>
				{/each}
			{/if}
		{/if}
	</div>
</div>
