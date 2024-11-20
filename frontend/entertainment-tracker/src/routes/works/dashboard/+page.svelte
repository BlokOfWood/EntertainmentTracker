<script context="module">
	export type WorkPlus = {
		work: Work;
		shared: boolean;
		sharedBy: string;
	};
</script>


<script lang="ts">
	import type { Work, SharedWork } from '$lib/api.model';
	import { deleteWork, getWorks, getSharedWorks } from '$lib/works.api';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	let works: WorkPlus[] = [];
	let originalWorks: WorkPlus[] = []; // To store the original order of works

	async function fetchWorks() {
		let currentNotSharedWorks: Work[] = [];

        const responseNotShared = await getWorks();
        if (responseNotShared.ok) {
            console.log('not shared works fetched successfully');
            currentNotSharedWorks = responseNotShared.body.mediaEntries;
            currentNotSharedWorks.forEach((currentNotSharedWork) => {
                console.log(currentNotSharedWork);
				const workPlus : WorkPlus = {
					work: currentNotSharedWork,
					shared: false,
					sharedBy: '',
				}
				works = [...works, workPlus];
            });
            originalWorks = works;

			console.log("not shared works:");
			works.forEach((work)=>{
				console.log(work);
			})
        }

		let currentSharedWorks: SharedWork[] = [];

		const responseShared = await getSharedWorks();
        if (responseShared.ok) {
            console.log('shared works fetched successfully');
			console.log(responseShared.body.sharedEntries);
            currentSharedWorks = responseShared.body.sharedEntries;
            if (currentSharedWorks != null){
				console.log('hii');
				currentSharedWorks.forEach((currentSharedWork) => {
                console.log(currentSharedWork);
				const workPlus : WorkPlus = {
					work: currentSharedWork.media_entry,
					shared: true,
					sharedBy: '',
				}
				works = [...works, workPlus];

				console.log("not works:");
				works.forEach((work)=>{
					console.log(work);
				})
				});
				originalWorks = works;
			}
        }

		works.forEach((work)=>{
			console.log(work.work.title);
		})

		console.log('final works:')
		works.forEach((work)=>{
				console.log(work);
		})
    }

	onMount(async () => {
		fetchWorks();
	});

	let sortedByTitle = false;
	let sortedByType = false;
	let sortedByProgress = false;
	let sortedByShared = false;

	function sortByTitle() {
		works = originalWorks;

		sortedByType = false;
		sortedByProgress = false;
		sortedByShared = false;

		if (sortedByTitle) {
			sortedByTitle = false;
			console.log('not sorted by title');
		} else {
			sortedByTitle = true;

			// Sort and create a new reference for works
			works = [...works].sort((a, b) => {
				if (a.work.title < b.work.title) return -1; // a comes before b
				if (a.work.title > b.work.title) return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by title');
		}
	}

	function sortByType() {
		works = originalWorks;

		sortedByTitle = false;
		sortedByProgress = false;
		sortedByShared = false;

		if (sortedByType) {
			sortedByType = false;
			console.log('not sorted by type');
		} else {
			sortedByType = true;

			// Sort and create a new reference for works
			works = [...works].sort((a, b) => {
				if (a.work.type < b.work.type) return -1; // a comes before b
				if (a.work.type > b.work.type) return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by type');
		}
	}

	function sortByProgress() {
		works = originalWorks;

		sortedByTitle = false;
		sortedByType = false;
		sortedByShared = false;

		if (sortedByProgress) {
			sortedByProgress = false;
			console.log('not sorted by progress');
		} else {
			sortedByProgress = true;

			// Sort and create a new reference for works
			works = [...works].sort((a, b) => {
				if (a.work.target_progress / a.work.current_progress < b.work.target_progress / b.work.current_progress)
					return -1; // a comes before b
				if (a.work.target_progress / a.work.current_progress > b.work.target_progress / b.work.current_progress)
					return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by progress');
		}
	}

	function sortByShared(){
		works = originalWorks;

		sortedByTitle = false;
		sortedByType = false;
		sortedByProgress = false

		if (sortedByShared) {
			sortedByShared = false;
			console.log('not sorted by shared');
		} else {
			sortedByShared = true;

			// Sort and create a new reference for works
			works = [...works].sort((a, b) => {
				if (!a.shared && b.shared) return 1; // a comes before b
				if (a.shared && !b.shared) return -1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by title');
		}
	}

	let currentWork!: Work;

	function shareMedia(work: Work) {
		goto('/works/dashboard/share-media', { state: { work } });
	}

	function editMedia(work: Work) {
		
		currentWork = work;
		goto('/works/dashboard/edit-media', { state: { work } });
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
			fetchWorks();
		});
		
	}

	function cancelDeletingMedia() {
		const popup = document.getElementById('popup');
		if (popup) {
			popup.classList.add('hidden');
		}
	}
</script>


<div class="relative z-0 flex h-full flex-grow items-center justify-center py-3">
	<div class="flex h-full w-full flex-col overflow-auto rounded-lg bg-white max-w-screen-lg">
		<div class="sticky top-0 z-10 grid w-full grid-cols-5  bg-white">
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={sortByTitle}
				>
					<span>Title</span>
					{#if sortedByTitle}
						<img src="/chevron-down-reverse.png" alt="Chevron-down" class="h-4 w-4" />
					{/if}
					{#if !sortedByTitle}
						<img src="/chevron-down.png" alt="Chevron-down" class="h-4 w-4" />
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={sortByType}
				>
					<span>Type</span>
					{#if sortedByType}
						<img src="/chevrons-up-down-inverse.png" alt="Chevron-down" class="h-4 w-4" />
					{/if}
					{#if !sortedByType}
						<img src="/chevrons-up-down.png" alt="Chevron-down" class="h-4 w-4" />
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={sortByProgress}
				>
					<span>Progress</span>
					{#if sortedByProgress}
						<img src="/chevrons-up-down-inverse.png" alt="Chevron-down" class="h-4 w-4" />
					{/if}
					{#if !sortedByProgress}
						<img src="/chevrons-up-down.png" alt="Chevron-down" class="h-4 w-4" />
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={sortByShared}
				>
					<span>Shared</span>
					{#if sortedByShared}
						<img src="/chevrons-up-down-inverse.png" alt="Chevron-down" class="h-4 w-4" />
					{/if}
					{#if !sortedByShared}
						<img src="/chevrons-up-down.png" alt="Chevron-down" class="h-4 w-4" />
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2"></div>
		</div>
		<div class="grid w-full grid-cols-5">
			{#each works as work}
				<div
					class="Ubuntu-font flex items-center justify-center border-0 p-2 text-center text-lg text-black"
				>
					{work.work.title}
				</div>
				<div
					class="Ubuntu-font flex items-center justify-center border-0 p-2 text-center text-lg text-black"
				>
					{work.work.type}
				</div>
				<div class="flex items-center justify-center border-0 p-2">
					{#if work.work.target_progress == work.work.current_progress}
						<img src="/done.png" alt="Done" class="h-4 w-4" />
					{/if}
					{#if work.work.target_progress != work.work.current_progress && work.work.type == 'book'}
						<div class="Ubuntu-font text-lg text-black">
							{work.work.current_progress}/{work.work.target_progress} pages
						</div>
					{/if}
					{#if work.work.target_progress != work.work.current_progress && work.work.type == 'TVshow'}
						<div class="Ubuntu-font text-lg text-black">
							{work.work.current_progress}/{work.work.target_progress} episodes
						</div>
					{/if}
					{#if work.work.target_progress != work.work.current_progress && work.work.type == 'movie'}
						<div class="Ubuntu-font text-lg text-black">
							{work.work.current_progress}/{work.work.target_progress} mins
						</div>
					{/if}
					{#if work.work.target_progress != work.work.current_progress && work.work.type == 'YouTubeVideo'}
						<div class="Ubuntu-font text-lg text-black">
							{work.work.current_progress} mins
						</div>
					{/if}
				</div>
				<div class="flex items-center justify-center border-0 p-2">
					{#if work.shared}
						<img src="/shared.png" alt="Shared" class="h-5 w-5">
					{/if}
				</div>
				<div class="flex items-center justify-center space-x-5 p-2">
					<button class="share-button" on:click={() => shareMedia(work.work)}>
						<img src="/share.png" alt="Share" class="h-5 w-5" />
					</button>
					<button class="edit-button" on:click={() => editMedia(work.work)}>
						<img src="/edit.png" alt="Edit" class="h-5 w-5" />
					</button>
					<button class="delete-button" on:click={() => deleteMedia(work.work.id)}>
						<img src="/trash.png" alt="Delete" class="h-5 w-5" />
					</button>
				</div>
			{/each}

		</div>
	</div>
</div>
<div
	id="popup"
	class="bg-background fixed inset-0 flex hidden items-center justify-center bg-opacity-75">
	<div class="w-1/3 rounded-lg bg-white p-6">
		<div class="Ubuntu-font mb-4 text-center text-lg font-bold">Delete Media Entry</div>
		<div class="Ubuntu-font mb-4">Are you sure you'd like to delete this media entry?</div>
		<div class="flex justify-center">
			<button
				class="Ubuntu-font text-sm mr-4 rounded bg-background px-4 py-2 text-white"
				on:click={cancelDeletingMedia}
			>
				Cancel
			</button>
			<button
				class="Ubuntu-font text-sm rounded bg-delete px-4 py-2 text-white"
				on:click={mediaDeleted}
			>
				Delete
			</button>
		</div>
	</div>
</div>