<script lang=ts>
	import type { Work } from '$lib/api.model';
	import Header from '$lib/header.svelte';
	import { getWorks } from '$lib/works.api';
	import { onMount } from 'svelte';

	let works: Work[] = [];
	let originalWorks: Work[] = []; // To store the original order of works

	onMount(async () => {
		getWorks().then((response) => {
			if (response.ok) {
				console.log('Works fetched successfully');
				console.log(response.body);
				works=response.body.mediaEntries; //<-itt szerinted mit hagyok ki? Az is lehet, hogy full rossz ez a sor.
				works.forEach(work => { console.log(work); }); //ez a foreach hibát dob a consol-ra, mintha üres lenne a works.
				
				//TODO: remove example test works once adding works, deleting works and edit works are implemented.
				const movieExample: Work = { id: 3, third_party_id: 'tp789', title: 'Movie', status: 'Pending', type: 'movie', current_progress: 10, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(movieExample);
				const TVShowExample: Work = { id: 3, third_party_id: 'tp789', title: 'TVShow', status: 'Pending', type: 'TVshow', current_progress: 20, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(TVShowExample);
				const YouTubeVideoExample: Work = { id: 3, third_party_id: 'tp789', title: 'YouTubeVideo', status: 'Pending', type: 'YouTubeVideo', current_progress: 80, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(YouTubeVideoExample);
				const TVShowExampleSecond: Work = { id: 3, third_party_id: 'tp789', title: 'TVShowSecond', status: 'Pending', type: 'TVshow', current_progress: 15, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(TVShowExampleSecond);
				const bookExample: Work = { id: 3, third_party_id: 'tp789', title: 'Book', status: 'Pending', type: 'book', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(bookExample);
				const YouTubeVideoExampleSecond: Work = { id: 3, third_party_id: 'tp789', title: 'Cooking', status: 'Pending', type: 'YouTubeVideo', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(YouTubeVideoExampleSecond);

				originalWorks=works;
			}
		});
	});

	let sortedByTitle = false;
	let sortedByType = false;
	let sortedByProgress = false;

	function sortByTitle() {
		works = originalWorks;

		sortedByType=false;
		sortedByProgress=false;

		if (sortedByTitle) {
			sortedByTitle = false;
			console.log("not sorted by title");
		} else {
			sortedByTitle = true;

			// Sort and create a new reference for works
			works = [...works].sort((a, b) => {
				if (a.title < b.title) return -1; // a comes before b
				if (a.title > b.title) return 1;  // a comes after b
				return 0; // a and b are equal
			});
			
			console.log("sort by title");
		}
	}

	function sortByType() {
		works = originalWorks;

		sortedByTitle=false;
		sortedByProgress=false;

		if (sortedByType) {
			sortedByType = false;
			console.log("not sorted by type");
		} else {
			sortedByType = true;

			// Sort and create a new reference for works
			works = [...works].sort((a, b) => {
				if (a.type < b.type) return -1; // a comes before b
				if (a.type > b.type) return 1;  // a comes after b
				return 0; // a and b are equal
			});
			
			console.log("sort by type");
		}
	}

	function sortByProgress() {
		works = originalWorks;

		sortedByTitle=false;
		sortedByType=false;

		if (sortedByProgress) {
			sortedByProgress = false;
			console.log("not sorted by progress");
		} else {
			sortedByProgress = true;

			// Sort and create a new reference for works
			works = [...works].sort((a, b) => {
				if (a.target_progress/a.current_progress < b.target_progress/b.current_progress) return -1; // a comes before b
				if (a.target_progress/a.current_progress > b.target_progress/b.current_progress) return 1;  // a comes after b
				return 0; // a and b are equal
			});
			
			console.log("sort by progress");
		}
	}

	function shareMedia() {
		//TODO: write share media function
	}

	function editMedia() {
		//TODO: write edit media function
	}

	function deleteMedia() {
		//TODO: write delete media function
	}
</script>

<div class="relative z-0 flex flex-grow items-center justify-center px-40 py-3">
	<div class="h-full w-full rounded-lg bg-white flex justify-center items-start">
		<div class="grid  grid-cols-4 w-full">
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={sortByTitle}
				>
					<span>Title</span>
					<img src="/chevron-down.png" alt="Chevron-down" class="h-4 w-4" />
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={sortByType}
				>
					<span>Type</span>
					<img src="/chevrons-up-down.png" alt="Chevron-down" class="h-4 w-4" />
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={sortByProgress}
				>
					<span>Progress</span>
					<img src="/chevrons-up-down.png" alt="Chevron-down" class="h-4 w-4" />
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2"></div>
			{#each works as work}
				<div class="border-0 p-2 text-black text-lg Ubuntu-font flex justify-center items-center">{work.title}</div>
				<div class="border-0 p-2 text-black text-lg Ubuntu-font flex justify-center items-center">{work.type}</div>
				<div class="border-0 p-2 flex justify-center items-center">
					{#if work.target_progress==work.current_progress}
					<img src="/done.png" alt="Done" class="w-4 h-4" />
					{/if}
					{#if work.target_progress!=work.current_progress && work.type=="book"}
					<div class=" text-black text-lg Ubuntu-font">
						{work.current_progress}/{work.target_progress} pages
					</div>
					{/if}
					{#if work.target_progress!=work.current_progress && work.type=="TVshow"}
					<div class=" text-black text-lg Ubuntu-font">
						{work.current_progress}/{work.target_progress} episodes
					</div>
					{/if}
					{#if work.target_progress!=work.current_progress && work.type=="movie"}
					<div class=" text-black text-lg Ubuntu-font">
						{work.current_progress}/{work.target_progress} mins
					</div>
					{/if}
					{#if work.target_progress!=work.current_progress && work.type=="YouTubeVideo"}
					<div class=" text-black text-lg Ubuntu-font">
						{work.current_progress} mins
					</div>
					{/if}
				</div>
				<div class="flex justify-center items-center space-x-5 p-2">
					<button class="share-button" on:click={shareMedia}>
						<img src="/share.png" alt="Share" class="w-5 h-5" />
					</button>
					<button class="edit-button" on:click={editMedia}>
						<img src="/edit.png" alt="Edit" class="w-5 h-5" />
					</button>
					<button class="delete-button" on:click={deleteMedia}>
						<img src="/trash.png" alt="Delete" class="w-5 h-5" />
					</button>
				</div>
			{/each}
		</div>
	</div>
</div>
