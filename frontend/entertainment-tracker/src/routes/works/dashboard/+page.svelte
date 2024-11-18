<script lang="ts">
	import { getBookByISBN, getTVShowByIMDb, getMovieByIMDb } from '$lib/addmedia.api';
	import type { Work, UpdateWorkRequest, Book, Movie, TvShow, ShareWorkRequest } from '$lib/api.model';
	import { deleteWork, getWorks, updateWork, shareWork } from '$lib/works.api';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	let works: Work[] = [];
	let originalWorks: Work[] = []; // To store the original order of works

	onMount(async () => {
		getWorks().then((response) => {
			if (response.ok) {
				console.log('Works fetched successfully');
				works = response.body.mediaEntries;
				works.forEach((work) => {
					console.log(work);
				});

				originalWorks = works;
			}
		});
	});

	let sortedByTitle = false;
	let sortedByType = false;
	let sortedByProgress = false;

	function sortByTitle() {
		works = originalWorks;

		sortedByType = false;
		sortedByProgress = false;

		if (sortedByTitle) {
			sortedByTitle = false;
			console.log('not sorted by title');
		} else {
			sortedByTitle = true;

			// Sort and create a new reference for works
			works = [...works].sort((a, b) => {
				if (a.title < b.title) return -1; // a comes before b
				if (a.title > b.title) return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by title');
		}
	}

	function sortByType() {
		works = originalWorks;

		sortedByTitle = false;
		sortedByProgress = false;

		if (sortedByType) {
			sortedByType = false;
			console.log('not sorted by type');
		} else {
			sortedByType = true;

			// Sort and create a new reference for works
			works = [...works].sort((a, b) => {
				if (a.type < b.type) return -1; // a comes before b
				if (a.type > b.type) return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by type');
		}
	}

	function sortByProgress() {
		works = originalWorks;

		sortedByTitle = false;
		sortedByType = false;

		if (sortedByProgress) {
			sortedByProgress = false;
			console.log('not sorted by progress');
		} else {
			sortedByProgress = true;

			// Sort and create a new reference for works
			works = [...works].sort((a, b) => {
				if (a.target_progress / a.current_progress < b.target_progress / b.current_progress)
					return -1; // a comes before b
				if (a.target_progress / a.current_progress > b.target_progress / b.current_progress)
					return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by progress');
		}
	}

	
	let currentWork!: Work;

	let checkingDashboard = true;
	let sharingMedia = false;
	let editingMedia = false;

	function shareMedia(work: Work) {
		checkingDashboard = false;
		sharingMedia = true;
		currentWork=work
	}

	let friendEmail = '';

	function shareWithFriend() {
		console.log('Sharing with:', friendEmail);

		let sharedWork: ShareWorkRequest = {
			media_entry: currentWork.id,
			share_with: friendEmail
		};

		shareWork(sharedWork)
	}

	let mediaArtSource = '/placeholderForEditMediaCoverArt.png';
	let description = '';
	let categories!: String[];
	let author = '';
	let YTURL = 'https://www.youtube.com/embed/dQw4w9WgXcQ?si=dourAMMy3-5pBbJr';

	let newProgress = -1;
	let newVideoTitle = '';

	let book = writable(undefined as Book| null | undefined);
	let movie = writable(undefined as Movie| null | undefined);
	let tvShow = writable(undefined as TvShow| null | undefined);

	function editMedia(work: Work) {
		currentWork = work;
		checkingDashboard = false;
		editingMedia = true;

		if (work.type == 'book') {
			if(currentWork.third_party_id!=""){
				getBookByISBN(currentWork.third_party_id).then(response => {
					const currentBook = response.body.book; // Get the book object directly
					book.set(response.body.book); // Use .set to update the store
					if (currentBook) { // Check if currentBook is not null
					author = currentBook.author; // Access author directly
					description = currentBook.description; // Access description
					mediaArtSource = currentBook.thumbnail; // Access thumbnail
					categories = currentBook.categories; // Access categories

				} else {
					// Handle the case where currentBook is null
					console.error("Book data is not available");
				}    
				});   
			}
			else{
				const mockbook : Book = {
						id: "",
						isbn: "",
						title: "",
						author: "",
						description: "",
						page_count: 0,
						thumbnail: "",
						categories: [],
						published_date: "",
						publisher: "",
						language: ""
					}
					book.set(mockbook);
			}
		}

		if (work.type == 'tvShow') {
			if(currentWork.third_party_id!=""){
				getTVShowByIMDb(currentWork.third_party_id).then(response => {
					const currentTVShow = response.body.tvshow; // Get the book object directly
					tvShow.set(response.body.tvshow); // Use .set to update the store
					if (currentTVShow) { // Check if currentBook is not null
					description = currentTVShow.overview; // Access description
					mediaArtSource = currentTVShow.thumbnail; // Access thumbnail
					categories = currentTVShow.genres; // Access categories

				} else {
					// Handle the case where currentBook is null
					console.error("TwShow data is not available");
				}    
				});   
			}
			else{
				const mockTVShow : TvShow = {
						id: 0,
						title: "",
						first_air_date: "",
						overview: "",
						popularity: 0,
						thumbnail: "",
						genres: [],
						number_of_seasons: 0,
						number_of_episodes: 0
					}
					tvShow.set(mockTVShow);
			}
		}

		if (work.type == 'movie') {
			if (currentWork.third_party_id!=""){
				getMovieByIMDb(currentWork.third_party_id).then(response => {
					const currentMovie = response.body.movie; // Get the book object directly
					movie.set(response.body.movie); // Use .set to update the store
					if (currentMovie) { // Check if currentBook is not null
					description = currentMovie.overview; // Access description
					mediaArtSource = currentMovie.thumbnail; // Access thumbnail
					categories = currentMovie.genres; // Access categories

				} else {
					// Handle the case where currentBook is null
					console.error("Book data is not available");
				}    
				});   
			}
			else{
				const mockMovie : Movie = {
						id: 0,
						title: "",
						release_date: "",
						overview: "",
						popularity: 0,
						thumbnail: "",
						genres: [],
						runtime: 0
					}
					movie.set(mockMovie);
			}
		}

		if (work.type == 'YouTubeVideo') {
			YTURL = work.third_party_id;

			//Converts the shared URL to the embed URL~
			let videoId = YTURL.split('https://youtu.be/')[1].split('?')[0];

			YTURL = 'https://www.youtube.com/embed/' + videoId + '?si=dourAMMy3-5pBbJr';
		}
	}

	function setNewProgress(event: Event) {
		const target = event.target as HTMLInputElement;
		const value = target.value;
		newProgress = Number(value);
	}

	function setNewTitle(event: Event) {
		const target = event.target as HTMLInputElement;
		const value = target.value;
		newVideoTitle = value;
	}

	function mediaEdited() {
		let newDetails: UpdateWorkRequest = {
			title: currentWork.title,
			type: currentWork.type,
			status: currentWork.status,
			current_progress: currentWork.current_progress,
			target_progress: currentWork.target_progress
		};

		if (newProgress != -1) {
			newDetails.current_progress = newProgress;
			console.log('New progress: ' + newDetails.current_progress);
		}

		if (currentWork.type == 'YouTubeVideo' && newVideoTitle != '') {
			newDetails.title = newVideoTitle;
			console.log('New title for video: ' + newDetails.title);
		}

		updateWork(currentWork.id, newDetails);

		//reset these values so it can be checked wether the user filled the fields or not
		newProgress = -1;
		newVideoTitle = '';

		returnToDashboard();
	}

	let idOfDeletedWork!: number;

	function deleteMedia(id: number) {
		console.log('deleting media');
		idOfDeletedWork = id;

		const popup = document.getElementById('popup');
		if (popup) {
			popup.classList.remove('hidden');
		}
	}

	function mediaDeleted() {
		const popup = document.getElementById('popup');
		if (popup) {
			popup.classList.add('hidden');
		}

		deleteWork(idOfDeletedWork).then(() => {
			getWorks().then((response) => {
				if (response.ok) {
					console.log('Works fetched successfully');
					works = response.body.mediaEntries;

					originalWorks = works;
				}
			});
		});
		returnToDashboard();
	}

	function cancelDeletingMedia() {
		const popup = document.getElementById('popup');
		if (popup) {
			popup.classList.add('hidden');
		}
	}

	function returnToDashboard() {
		checkingDashboard = true;
		sharingMedia = false;
		editingMedia = false;

		getWorks().then((response) => {
			if (response.ok) {
				console.log('Works fetched successfully');
				works = response.body.mediaEntries;

				originalWorks = works;
			}
		});

		description="";
		author="";
	}
</script>

{#if checkingDashboard}
	<div class="relative z-0 flex h-full flex-grow items-center justify-center py-3">
		<div class="flex h-full w-full flex-col overflow-auto rounded-lg bg-white max-w-screen-md">
			<div class="sticky top-0 z-10 grid w-full grid-cols-4">
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
			</div>
			<div class="grid w-full grid-cols-4">
				{#each works as work}
					<div
						class="Ubuntu-font flex items-center justify-center border-0 p-2 text-center text-lg text-black"
					>
						{work.title}
					</div>
					<div
						class="Ubuntu-font flex items-center justify-center border-0 p-2 text-center text-lg text-black"
					>
						{work.type}
					</div>
					<div class="flex items-center justify-center border-0 p-2">
						{#if work.target_progress == work.current_progress}
							<img src="/done.png" alt="Done" class="h-4 w-4" />
						{/if}
						{#if work.target_progress != work.current_progress && work.type == 'book'}
							<div class="Ubuntu-font text-lg text-black">
								{work.current_progress}/{work.target_progress} pages
							</div>
						{/if}
						{#if work.target_progress != work.current_progress && work.type == 'TVshow'}
							<div class="Ubuntu-font text-lg text-black">
								{work.current_progress}/{work.target_progress} episodes
							</div>
						{/if}
						{#if work.target_progress != work.current_progress && work.type == 'movie'}
							<div class="Ubuntu-font text-lg text-black">
								{work.current_progress}/{work.target_progress} mins
							</div>
						{/if}
						{#if work.target_progress != work.current_progress && work.type == 'YouTubeVideo'}
							<div class="Ubuntu-font text-lg text-black">
								{work.current_progress} mins
							</div>
						{/if}
					</div>
					<div class="flex items-center justify-center space-x-5 p-2">
						<button class="share-button" on:click={() => shareMedia(work)}>
							<img src="/share.png" alt="Share" class="h-5 w-5" />
						</button>
						<button class="edit-button" on:click={() => editMedia(work)}>
							<img src="/edit.png" alt="Edit" class="h-5 w-5" />
						</button>
						<button class="delete-button" on:click={() => deleteMedia(work.id)}>
							<img src="/trash.png" alt="Delete" class="h-5 w-5" />
						</button>
					</div>
				{/each}
			</div>
		</div>
	</div>
	<div
		id="popup"
		class="bg-background fixed inset-0 flex hidden items-center justify-center bg-opacity-75"
	>
		<div class="w-1/3 rounded-lg bg-white p-6">
			<div class="Ubuntu-font mb-4 text-center text-lg font-bold">Delete Media Entry</div>
			<div class="Ubuntu-font mb-4">Are you sure you'd like to delete this media entry?</div>
			<div class="flex justify-center">
				<button
					class="Ubuntu-font mr-4 rounded bg-red-500 px-4 py-2 text-white"
					on:click={mediaDeleted}
				>
					Delete
				</button>
				<button
					class="Ubuntu-font rounded bg-slate-500 px-4 py-2 text-white"
					on:click={cancelDeletingMedia}
				>
					Cancel
				</button>
			</div>
		</div>
	</div>
{/if}
{#if sharingMedia}
	<div class="relative z-0 flex h-full flex-grow items-center justify-center py-3">
		<div class="h-full w-full overflow-auto rounded-lg bg-white max-w-screen-sm">
			<div class="flex flex-col">
				<div class="flex items-start justify-between p-2">
					<button
						class="flex items-center justify-center border-0 pl-6 pt-4"
						on:click={returnToDashboard}
					>
						<img src="/back-button.png" alt="Return to dashboard" class="h-5 w-5" />
					</button>
					<div class="Ubuntu-font flex-grow pt-3 text-center text-lg font-bold">
						Share your progress with a friend
					</div>
				</div>
			</div>
			<div class="flex w-full flex-col items-center justify-center p-2">
				<div class="ml-20 inline-flex flex-col">
					<div class="Ubuntu-font p-1 text-sm">Your friend's email:</div>
					<div>
						<input
							type="email"
							placeholder="Email"
							bind:value={friendEmail}
							class="mr-1 rounded-md border p-1.5"
						/>
						<button
							class="bg-background rounded-md px-6 py-1.5 text-white"
							on:click={shareWithFriend}
						>
							Share
						</button>
					</div>
				</div>
			</div>
		</div>
	</div>
{/if}
{#if editingMedia}
	<div class="relative z-0 flex flex-grow h-full items-center justify-center py-3">
		<div class="h-full w-full max-w-screen-sm rounded-lg bg-white overflow-auto">
			<div class="flex flex-col">
				<div class="flex items-start justify-between p-2">
					<button
						class="flex items-center justify-center border-0 pl-6 pt-4"
						on:click={returnToDashboard}
					>
						<img src="/back-button.png" alt="Return to dashboard" class="h-5 w-5" />
					</button>
				</div>
				<div></div>
			</div>
			{#if $book !== undefined || $movie !== undefined || $tvShow !== undefined}
				<div class="flex items-start p-4">
					<div class="ml-10 mr-4 flex flex-col min-w-32">
						{#if currentWork.third_party_id!==""}
							<img src={mediaArtSource} alt="Cover art" class="rounded-md h-auto" />
							{#if currentWork.type === 'book'}
								<div class="text-xxs Ubuntu-font pt-1 text-center">
									{currentWork.third_party_id}
								</div>
							{/if}
							<div class="text-xxs Ubuntu-font p-1 text-center">Categories:</div>
							{#each categories as category}
								<div class="text-xxs Ubuntu-font pb-1 text-center">
									{category}
								</div>
							{/each}
						{/if}
					</div>
					<div class="ml-4 mr-8">
						<div class="Ubuntu-font text-center text-sm font-bold">
							{currentWork.title}
						</div>
						{#if currentWork.type === 'book'}
							<div class="Ubuntu-font text-center text-xs font-bold">
								{author}
							</div>
						{/if}
						<div class="text-xxs Ubuntu-font pt-3 text-start" style="line-height: 2.5;">
							{description}
						</div>
						<div class="Ubuntu-font pb-2 pt-6 text-sm font-bold">Progress</div>
						<div>
							{#if currentWork.type === 'TVshow'}
								<input
									type="progress"
									placeholder="Episode number / {currentWork.target_progress}"
									on:input={setNewProgress}
									class="mr-1 rounded-md border p-1.5 text-sm"
								/>
							{/if}
							{#if currentWork.type === 'book'}
								<input
									type="progress"
									placeholder="Page number / {currentWork.target_progress}"
									on:input={setNewProgress}
									class="mr-1 rounded-md border p-1.5 text-sm"
								/>
							{/if}
							{#if currentWork.type === 'movie'}
								<input
									type="progress"
									placeholder="Minutes / {currentWork.target_progress}"
									on:input={setNewProgress}
									class="mr-1 rounded-md border p-1.5 text-sm"
								/>
							{/if}
							<button
								class="bg-background rounded-md px-6 py-1.5 text-sm text-white"
								on:click={mediaEdited}
							>
								Save
							</button>
						</div>
					</div>
				</div>
			{/if}
			{#if currentWork.type === 'YouTubeVideo'}
				<div class="ml-24 mr-24 mt-6 flex flex-col">
					<div class="flex items-center justify-center">
						<iframe title="video" class="aspect-[18/10] w-full" src={YTURL}> </iframe>
					</div>
					<div class="Ubuntu-font pt-6 text-xs">Set Display name</div>
					<div class="flex pt-1">
						<input
							type="displayTitle"
							placeholder="Name"
							on:input={setNewTitle}
							class="mr-1 w-64 rounded-md border p-1.5 text-sm"
						/>
					</div>
					<div class="Ubuntu-font pb-1 pt-2 text-xs">Progress</div>
					<div class="flex">
						<input
							type="progress"
							placeholder="Minutes"
							on:input={setNewProgress}
							class="mr-1 rounded-md border p-1.5 text-sm"
						/>
						<button
							class="bg-background rounded-md px-6 py-1.5 text-sm text-white"
							on:click={mediaEdited}
						>
							Save
						</button>
					</div>
				</div>
			{/if}
		</div>
	</div>
{/if}
