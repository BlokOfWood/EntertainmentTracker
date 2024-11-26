<script context="module">
	export type WorkPlus = {
		work: Work;
		shared: boolean;
		sharedBy: string;
	};
</script>

<script lang="ts">
	import { BookMarked, Users, Share, Trash2, ChevronDown, ChevronUp, ChevronsUpDown } from 'lucide-svelte';
	import type { Work, SharedWork } from '$lib/api.model';
	import { deleteWork, getWorks, getSharedWorks } from '$lib/works.api';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { writable, type Writable } from 'svelte/store';

	let works: Writable<WorkPlus[]> = writable([]);
	let originalWorks: WorkPlus[] = []; // To store the original order of works

	async function fetchWorks() {
		works.update(() => []);

		let currentNotSharedWorks: Work[] = [];

		const responseNotShared = await getWorks();
		if (responseNotShared.ok) {
			console.log('not shared works fetched successfully');
			currentNotSharedWorks = responseNotShared.body.mediaEntries;
			currentNotSharedWorks.forEach((currentNotSharedWork) => {
				const workPlus: WorkPlus = {
					work: currentNotSharedWork,
					shared: false,
					sharedBy: ''
				};
				works.update((existing) => [...existing, workPlus]);
			});
			originalWorks = $works;
		}

		let currentSharedWorks: SharedWork[] = [];

		const responseShared = await getSharedWorks();
		if (responseShared.ok) {
			console.log('shared works fetched successfully');
			currentSharedWorks = responseShared.body.sharedEntries;
			if (currentSharedWorks != null) {
				currentSharedWorks.forEach((currentSharedWork) => {
					const workPlus: WorkPlus = {
						work: currentSharedWork.media_entry,
						shared: true,
						sharedBy: ''
					};
					works.update((existing) => [...existing, workPlus]);
				});
				originalWorks = $works;
			}
		}

		$works = [...$works].sort((a, b) => {
				if (a.work.created_at > b.work.created_at) return -1; // a comes before b
				if (a.work.created_at < b.work.created_at) return 1; // a comes after b
				return 0; // a and b are equal
			});
		
		originalWorks=$works;
	}

	onMount(async () => {
		fetchWorks();
	});

	let sortedByTitle = 0;
	let sortedByType = 0;
	let sortedByProgress = 0;
	let sortedByShared = 0;

	function sortByTitle() {
		sortedByType = 0;
		sortedByProgress = 0;
		sortedByShared = 0;

		if (sortedByTitle==2) {
			sortedByTitle = 0;
			$works = originalWorks;
			console.log('not sorted by title');
		} else if(sortedByTitle==0){
			sortedByTitle = sortedByTitle+1;

			// Sort and create a new reference for works
			$works = [...originalWorks].sort((a, b) => {
				if (a.work.title < b.work.title) return -1; // a comes before b
				if (a.work.title > b.work.title) return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by title (asc)');
		}
		else{
			sortedByTitle = sortedByTitle+1;

			// Sort and create a new reference for works
			$works = [...originalWorks].sort((a, b) => {
				if (a.work.title > b.work.title) return -1; // a comes before b
				if (a.work.title < b.work.title) return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by title (desc)');
		}
	}

	function sortByType() {
		$works = originalWorks;

		sortedByTitle = 0;
		sortedByProgress = 0;
		sortedByShared = 0;

		if (sortedByType==2) {
			sortedByType = 0;
			console.log('not sorted by type');
		} else if (sortedByType==0){
			sortedByType = 1;

			// Sort and create a new reference for works
			$works = [...$works].sort((a, b) => {
				if (a.work.type < b.work.type) return -1; // a comes before b
				if (a.work.type > b.work.type) return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by type (asc)');
		}
		else{
			sortedByType = 2;

			// Sort and create a new reference for works
			$works = [...$works].sort((a, b) => {
				if (a.work.type > b.work.type) return -1; // a comes before b
				if (a.work.type < b.work.type) return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by type (desc)');
		}
	}

	function sortByProgress() {
		$works = originalWorks;

		sortedByTitle = 0;
		sortedByType = 0;
		sortedByShared = 0;

		if (sortedByProgress==2) {
			sortedByProgress = 0;
			console.log('not sorted by progress');
		} else if(sortedByProgress==0) {
			sortedByProgress = 1;

			// Sort and create a new reference for works
			$works = [...$works].sort((a, b) => {
				if (
					a.work.target_progress / a.work.current_progress < b.work.target_progress / b.work.current_progress
				)
					return -1; // a comes before b
				if (
					a.work.target_progress / a.work.current_progress > b.work.target_progress / b.work.current_progress
				)
					return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by progress (asc)');
		}
		else{
			sortedByProgress = 2;

			// Sort and create a new reference for works
			$works = [...$works].sort((a, b) => {
				if (
					a.work.target_progress / a.work.current_progress > b.work.target_progress / b.work.current_progress
				)
					return -1; // a comes before b
				if (
					a.work.target_progress / a.work.current_progress < b.work.target_progress / b.work.current_progress
				)
					return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by progress');
		}
	}

	function sortByShared() {
		$works = originalWorks;

		sortedByTitle = 0;
		sortedByType = 0;
		sortedByProgress = 0;

		if (sortedByShared==2) {
			sortedByShared = 0;
			console.log('not sorted by shared');
		} else if (sortedByShared==0){
			sortedByShared = 1;

			// Sort and create a new reference for works
			$works = [...$works].sort((a, b) => {
				if (!a.shared && b.shared) return 1; // a comes before b
				if (a.shared && !b.shared) return -1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by title (asc)');
		}
		else{
			sortedByShared = 2;

			// Sort and create a new reference for works
			$works = [...$works].sort((a, b) => {
				if (!a.shared && b.shared) return -1; // a comes before b
				if (a.shared && !b.shared) return 1; // a comes after b
				return 0; // a and b are equal
			});

			console.log('sort by title (desc)');
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

	function openDeleteModal(id: number) {
		console.log('deleting media');
		idOfDeletedWork = id;

		const popup = document.getElementById('popup');
		if (popup) {
			popup.classList.remove('hidden');
		}
	}

	async function deleteMedia() {
		const popup = document.getElementById('popup');
		if (popup) {
			popup.classList.add('hidden');
		}

		await deleteWork(idOfDeletedWork);
		await fetchWorks();
	}

	function cancelDeletingMedia() {
		const popup = document.getElementById('popup');
		if (popup) {
			popup.classList.add('hidden');
		}
	}
</script>

<div class="relative z-0 flex h-full flex-grow items-center justify-center py-3">
	<div class="flex h-full w-full max-w-screen-lg flex-col overflow-auto rounded-lg bg-white">
		<div class="sticky top-0 z-10 grid w-full grid-cols-5 bg-white">
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={sortByTitle}
				>
					<span>Title</span>
					{#if sortedByTitle==0}
						<ChevronsUpDown />
					{/if}
					{#if sortedByTitle==1}
						<ChevronUp />
					{/if}
					{#if sortedByTitle==2}
						<ChevronDown />
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={sortByType}
				>
					<span>Type</span>
					{#if sortedByType==0}
						<ChevronsUpDown />
					{/if}
					{#if sortedByType==1}
						<ChevronUp />
					{/if}
					{#if sortedByType==2}
						<ChevronDown />
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={sortByProgress}
				>
					<span>Progress</span>
					{#if sortedByProgress==0}
						<ChevronsUpDown />
					{/if}
					{#if sortedByProgress==1}
						<ChevronUp />
					{/if}
					{#if sortedByProgress==2}
						<ChevronDown />
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={sortByShared}
				>
					<span>Shared</span>
					{#if sortedByShared==0}
						<ChevronsUpDown />
					{/if}
					{#if sortedByShared==1}
						<ChevronUp />
					{/if}
					{#if sortedByShared==2}
						<ChevronDown />
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2"></div>
		</div>
		<div class="grid w-full grid-cols-5">
			{#each $works as work}
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
							{work.work.current_progress}/{work.work.target_progress} pages ({Math.round(work.work.current_progress/work.work.target_progress*100)}%)
						</div>
					{/if}
					{#if work.work.target_progress != work.work.current_progress && work.work.type == 'show'}
						<div class="Ubuntu-font text-lg text-black">
							{work.work.current_progress}/{work.work.target_progress} episodes ({Math.round(work.work.current_progress/work.work.target_progress*100)}%)
						</div>
					{/if}
					{#if work.work.target_progress != work.work.current_progress && work.work.type == 'movie'}
						<div class="Ubuntu-font text-lg text-black">
							{work.work.current_progress}/{work.work.target_progress} mins ({Math.round(work.work.current_progress/work.work.target_progress*100)}%)
						</div>
					{/if}
					{#if work.work.target_progress != work.work.current_progress && work.work.type == 'youtube'}
						<div class="Ubuntu-font text-lg text-black">
							{work.work.current_progress} mins ({Math.round(work.work.current_progress/work.work.target_progress*100)}%)
						</div>
					{/if}
				</div>
				<div class="flex items-center justify-center border-0 p-2">
					{#if work.shared}
						<Users />
					{/if}
				</div>
				<div class="flex items-center justify-center space-x-5 p-2">
					{#if work.shared}
						<div class="share-button opacity-20">
							<Share />
						</div>
						<div class="edit-button opacity-20">
							<BookMarked />
						</div>
						<div class="delete-button opacity-20">
							<Trash2 color=red />
						</div>
					{/if}
					{#if !work.shared}
						<button class="share-button" on:click={() => shareMedia(work.work)}>
							<Share />
						</button>
						<button class="edit-button" on:click={() => editMedia(work.work)}>
							<BookMarked />
						</button>
						<button class="delete-button" on:click={() => openDeleteModal(work.work.id)}>
							<Trash2 color=red />
						</button>
					{/if}
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
				class="Ubuntu-font bg-background mr-4 rounded px-4 py-2 text-sm text-white"
				on:click={cancelDeletingMedia}
			>
				Cancel
			</button>
			<button
				class="Ubuntu-font bg-delete rounded px-4 py-2 text-sm text-white"
				on:click={deleteMedia}
			>
				Delete
			</button>
		</div>
	</div>
</div>
