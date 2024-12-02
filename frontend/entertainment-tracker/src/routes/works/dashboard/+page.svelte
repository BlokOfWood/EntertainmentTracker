<script lang="ts">
	import { BookMarked, Users, Share, Trash2, ChevronDown, ChevronUp, ChevronsUpDown, CircleCheck } from 'lucide-svelte';
	import type { Work, SharedWork, UpdateWorkRequest, ShareWorkRequest} from '$lib/api.model';
	import { deleteWork, deleteSharedWork, getWorks, getSharedWorks, updateWork, shareWork } from '$lib/works.api';
	import { getBookByISBN, getTVShowByIMDbId, getMovie, getBookByGoogleId } from '$lib/addmedia.api';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { writable, type Writable } from 'svelte/store';
	import type { WorkPlus } from './dashboard'
	import {fetchWorks, sortByTitle, sortByType, sortByProgress, sortByShared} from './dashboard'

	//-------------LOADING IN WORKS AND DISPLAYING THEM-------------------

	let works: Writable<WorkPlus[]> = writable([]);
	let originalWorks: WorkPlus[] = []; // To store the original order of works

	onMount(async () => {
    	await handleFetchWorks();
	});

	async function handleFetchWorks(){
		// Await the fetchWorks to get the data
		const fetchedWorks = await fetchWorks();
		works.set(fetchedWorks); // Set the fetched works to the writable store
		originalWorks = [...fetchedWorks]; // Store a copy of the original works
	}

	//-------------------SORTING WORKS--------------------------

	let sortedByTitle = 0;
	let sortedByType = 0;
	let sortedByProgress = 0;
	let sortedByShared = 0;

	function handleSortByTitle(){
		const { sortedByTitle: newSortedByTitle, sortedWorks } = sortByTitle(sortedByTitle, originalWorks);

        // Update the writable stores
        sortedByTitle = newSortedByTitle; // Update the local variable
        works.set(sortedWorks);

		sortedByType = 0;
    	sortedByProgress = 0;
    	sortedByShared = 0;
	}

	function handleSortByType(){
		const { sortedByType: newSortedByType, sortedWorks } = sortByType(sortedByType, originalWorks);

		sortedByType = newSortedByType; // Update the local variable
        works.set(sortedWorks);

		sortedByTitle = 0;
		sortedByProgress = 0;
		sortedByShared = 0;
	}

	function handleSortByProgress(){
		const { sortedByProgress: newSortedByProgress, sortedWorks } = sortByProgress(sortedByProgress, originalWorks);

		sortedByProgress = newSortedByProgress; // Update the local variable
        works.set(sortedWorks);

		sortedByTitle = 0;
		sortedByType = 0;
		sortedByShared = 0;
	}

	function handleSortByShared(){
		const { sortedByShared: newSortedByShared, sortedWorks } = sortByShared(sortedByShared, originalWorks);

		sortedByShared = newSortedByShared; // Update the local variable
        works.set(sortedWorks);

		sortedByTitle = 0;
		sortedByType = 0;
		sortedByProgress = 0;
	}

	//-------CHECKING, SHARING, EDITING, DELETING WORKS -------------------

	let currentWork!: Work;

	//-----------CHECKING WORKS-------------------------------

	function checkMedia(work: Work) {
		currentWork = work;
		goto('/works/dashboard/edit-media', { state: { work } });
	}

	//---------SHARING WORKS-------------------------------

	function openShareMediaModal(work: Work){
		//goto('/works/dashboard/share-media', { state: { work } });

		currentWork = work;

		const popup = document.getElementById('share-popup');
		if (popup) {
			popup.classList.remove('hidden');
		}
	}

	let friendEmail="";
	let shareMessage: Writable<string>=writable("");

	async function shareMedia() {

		let sharedWork: ShareWorkRequest = {
			media_entry: currentWork.id,
			share_with: friendEmail
		};

		try {
			const response = await shareWork(sharedWork); // Assuming this returns an object with statusCode and ok

			// Check if the response is OK
			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.statusCode}`); // Use statusCode instead of status
			}

			// If the function succeeds, you can handle success here
			const popup = document.getElementById('share-popup');
			if (popup) {
				popup.classList.add('hidden');
			}
		} catch (error) {
			console.error('Error occurred while sharing work:', error);

			// Check the error message for specific status codes
			if (error instanceof Error) {
			// Check the error message for specific status codes
			if (error.message.includes('404')) {
				$shareMessage = "User  doesn't exist."; // Set message for 404 error
			} else if (error.message.includes('409')) {
				$shareMessage = "Already shared this media with this user."; // Set message for 409 error
			} else {
				$shareMessage = "Wrong email address."; // Fallback message for other errors
			}
			} else {
				// Handle the case where error is not an instance of Error
				$shareMessage = "Wrong email address.";
			}
		}
	}

	function cancelShareMedia(){
		$shareMessage="";
		const popup = document.getElementById('share-popup');
		if (popup) {
			popup.classList.add('hidden');
		}
	}

	//-------------EDITING WORKS---------------------------------

	let work_progressValue=0;
	let work_target_progress=0;
	let work_type="";
	let work_title="";
	let work_progressValueYT="";

	function editMediaModal(work: Work) {
		work_progressValue=work.current_progress;
		work_target_progress=work.target_progress;
		work_type=work.type;
		work_title=work.title;

		if(work.type=="youtube"){
			convertProgressToYTProgress(work.current_progress);
		}

		currentWork=work;

		const popup = document.getElementById('edit-popup');
		if (popup) {
			popup.classList.remove('hidden');
		}
	}

	function convertProgressToYTProgress(progress : number){
		const hours = Math.floor(progress / 3600);
			const minutes = Math.floor((progress % 3600) / 60);
			const secs = progress % 60;

			if(progress>3600){
				work_progressValueYT=`${String(hours).padStart(2, '0')}:${String(minutes).padStart(2, '0')}:${String(secs).padStart(2, '0')}`;
			}
			else{
				work_progressValueYT=`${String(minutes).padStart(2, '0')}:${String(secs).padStart(2, '0')}`;
			}
	}

	let editMessage: Writable<string>=writable("");

	async function editMedia(){
		if(currentWork.type=="youtube"){
			work_progressValue=convertYTProgressToSeconds(work_progressValueYT);
		}

		if (Number(work_progressValue)!=currentWork.current_progress){
			let newDetails: UpdateWorkRequest = {
			title: currentWork.title,
			type: currentWork.type,
			status: currentWork.status,
			current_progress: Number(work_progressValue),
			target_progress: currentWork.target_progress
			};

			console.log('New progress: ' + newDetails.current_progress);
			
			try {
				const response = await updateWork(currentWork.id, newDetails); // Assuming this returns an object with statusCode and ok

				// Check if the response is OK
				if (!response.ok) {
					throw new Error(`HTTP error! status: ${response.statusCode}`); // Use statusCode instead of status
				}

				// If the function succeeds, you can handle success here
				const popup = document.getElementById('edit-popup');
				if (popup) {
					popup.classList.add('hidden');
				}
				await handleFetchWorks();
			} catch (error) {
				console.error('Error occurred while sharing work:', error);

				$editMessage = "Invalid progress.";
				work_progressValue=currentWork.current_progress;
			}
		}
	}

	function convertYTProgressToSeconds(YTprogress : string){
		const parts = YTprogress.split(':');
		let seconds = 0;

		if (parts.length === 3) {
			seconds += parseInt(parts[0]) * 3600; // hours to seconds
			seconds += parseInt(parts[1]) * 60;   // minutes to seconds
			seconds += parseInt(parts[2]);        // seconds
		}
		if (parts.length == 2) {
			console.log("has 2 parts")
			seconds += parseInt(parts[0]) * 60;   // minutes to seconds
			seconds += parseInt(parts[1]);        // seconds
		}

		console.log(seconds);
		return seconds;
	}

	function cancelEditMedia(){
		const popup = document.getElementById('edit-popup');
		if (popup) {
			popup.classList.add('hidden');
		}

		$editMessage="";
	}

	//----------------------DELETE MEDIA-------------------------------------

	let idOfDeletedWork!: number;
	let isWorkAboutToBeDeletedShared!: boolean;

	function openDeleteModal(id: number, isShared: boolean) {
		console.log('deleting media');
		idOfDeletedWork = id;
		isWorkAboutToBeDeletedShared=isShared;

		const popup = document.getElementById('delete-popup');
		if (popup) {
			popup.classList.remove('hidden');
		}
	}

	async function deleteMedia() {
		const popup = document.getElementById('delete-popup');
		if (popup) {
			popup.classList.add('hidden');
		}

		if(isWorkAboutToBeDeletedShared){
			await deleteSharedWork(idOfDeletedWork);
			console.log("Media unshared with me.")
		}
		else{
			await deleteWork(idOfDeletedWork);
			console.log("Deleted media.")
		}
		
		await handleFetchWorks();
	}

	function cancelDeletingMedia() {
		const popup = document.getElementById('delete-popup');
		if (popup) {
			popup.classList.add('hidden');
		}
	}
</script>

<!--DASHBOARD-->

<div class="relative z-0 flex h-full flex-grow items-center justify-center py-3">
	<div class="flex h-full w-full max-w-screen-lg flex-col overflow-auto rounded-lg bg-white">
		<div class="sticky top-0 z-10 grid w-full grid-cols-6 bg-white">
			<div></div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={handleSortByTitle}
				>
					<span>Title</span>
					{#if sortedByTitle==0}
						<ChevronsUpDown size={20}/>
					{/if}
					{#if sortedByTitle==1}
						<ChevronUp size={20}/>
					{/if}
					{#if sortedByTitle==2}
						<ChevronDown size={20}/>
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={handleSortByType}
				>
					<span>Type</span>
					{#if sortedByType==0}
						<ChevronsUpDown size={20}/>
					{/if}
					{#if sortedByType==1}
						<ChevronUp size={20}/>
					{/if}
					{#if sortedByType==2}
						<ChevronDown size={20}/>
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={handleSortByProgress}
				>
					<span>Progress</span>
					{#if sortedByProgress==0}
						<ChevronsUpDown size={20}/>
					{/if}
					{#if sortedByProgress==1}
						<ChevronUp size={20}/>
					{/if}
					{#if sortedByProgress==2}
						<ChevronDown size={20}/>
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2">
				<button
					class="Ubuntu-font flex items-center space-x-2 text-lg font-bold text-black"
					on:click={handleSortByShared}
				>
					<span>Shared</span>
					{#if sortedByShared==0}
						<ChevronsUpDown size={20}/>
					{/if}
					{#if sortedByShared==1}
						<ChevronUp size={20}/>
					{/if}
					{#if sortedByShared==2}
						<ChevronDown size={20}/>
					{/if}
				</button>
			</div>
			<div class="flex items-start justify-center border-0 p-2"></div>
		</div>
		<div class="flex-col">
			{#each $works as work}
				<div class="grid w-full grid-cols-6 gap-y-5 p-2">
					<div class="flex items-center justify-center">
						{#if work.work.type!=='youtube'}
							<img src={work.thumbnail} alt="Cover art" class="rounded-md h-auto w-full max-w-[40%]" />
						{/if}
						{#if work.work.type==='youtube'}
							<iframe title="video" class="aspect-[18/10] max-w-[60%]" src="{work.thumbnail}"> </iframe>
						{/if}
					</div>
					<button
						class="Ubuntu-font flex items-center justify-center border-0 p-2 text-center text-lg text-black hover:underline"
						on:click={() => checkMedia(work.work)}>
						{work.work.title}
					</button>
					<div
						class="Ubuntu-font flex items-center justify-center border-0 p-2 text-center text-lg text-black"
					>
						{work.work.type}
					</div>
					<div class="flex items-center justify-center border-0 p-2">
						{#if work.work.target_progress == work.work.current_progress}
							<CircleCheck size={20} />
						{/if}
						{#if work.work.target_progress != work.work.current_progress && work.work.type == 'book'}
							<div class="Ubuntu-font text-lg text-black text-center">
								{work.work.current_progress}/{work.work.target_progress} pages
								({Math.round(work.work.current_progress/work.work.target_progress*100)}%)
							</div>
						{/if}
						{#if work.work.target_progress != work.work.current_progress && work.work.type == 'show'}
							<div class="Ubuntu-font text-lg text-black text-center">
								{work.work.current_progress}/{work.work.target_progress} episodes
								({Math.round(work.work.current_progress/work.work.target_progress*100)}%)
							</div>
						{/if}
						{#if work.work.target_progress != work.work.current_progress && work.work.type == 'movie'}
							<div class="Ubuntu-font text-lg text-black text-center">
								{work.work.current_progress}/{work.work.target_progress} mins
								({Math.round(work.work.current_progress/work.work.target_progress*100)}%)
							</div>
						{/if}
						{#if work.work.target_progress != work.work.current_progress && work.work.type == 'youtube'}
							<div class="Ubuntu-font text-lg text-black text-center">
								{Math.floor(work.work.current_progress/60)} mins
								({Math.round(work.work.current_progress/work.work.target_progress*100)}%)
							</div>
						{/if}
					</div>
					<div class="flex items-center justify-center border-0 p-2">
						{#if work.shared}
							<Users size={20}/>
						{/if}
					</div>
					<div class="flex items-center justify-center space-x-5 p-2">
						{#if work.shared}
							<button class="delete-button" on:click={() => openDeleteModal(work.work.id, work.shared)}>
								<Trash2 color=red size={20}/>
							</button>
						{/if}
						{#if !work.shared}
							<button class="share-button" on:click={() => openShareMediaModal(work.work)}>
								<Share size={20}/>
							</button>
							<button class="edit-button" on:click={() => editMediaModal(work.work)}>
								<BookMarked size={20}/>
							</button>
							<button class="delete-button" on:click={() => openDeleteModal(work.work.id, work.shared)}>
								<Trash2 color=red size={20}/>
							</button>
						{/if}
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>

<!--DELETE MODAL-->

<div
	id="delete-popup"
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

<!--EDIT MODAL-->

<div
	id="edit-popup"
	class="bg-background fixed inset-0 flex hidden items-center justify-center bg-opacity-75"
>
	<div class="w-2/5 rounded-lg bg-white p-6">
		<div class="Ubuntu-font mb-4 text-center text-lg font-bold">Edit {work_title} status</div>
		<div class="flex flex-col items-center">
			<!-- Progress Bar -->
			{#if work_type!=="youtube"}
			<div class="flex flex-col items-center w-[70%]">
				<div class="relative w-full">
					<div class="h-2 bg-slider rounded-lg">
						<div class="h-full bg-background rounded-lg"
							style="width: {(work_progressValue/work_target_progress)*100}%; position: relative;">
						<!-- Orb at the end of the green part -->
						<div
							class="absolute top-1/2 transform -translate-y-1/2 w-4 h-4 bg-white border-2 border-black rounded-full"
							style="left: calc(100% - 0.5rem);"
						></div>
						</div>
					</div>
					<input
						id="progress-slider"
						type="range"
						min="0"
						max={work_target_progress}
						bind:value={work_progressValue}
						class="absolute top-0 left-0 w-full h-2 bg-transparent cursor-pointer"
						style="pointer-events: all; opacity: 0; z-index: 1;"
					/>
				</div>
			</div>
			{/if}
			{#if work_type==="youtube"}
			<div class="flex flex-col items-center w-[70%]">
				<div class="relative w-full">
					<div class="h-2 bg-slider rounded-lg">
						<div class="h-full bg-background rounded-lg"
							style="width: {(work_progressValue/work_target_progress)*100}%; position: relative;">
						<!-- Orb at the end of the green part -->
						<div
							class="absolute top-1/2 transform -translate-y-1/2 w-4 h-4 bg-white border-2 border-black rounded-full"
							style="left: calc(100% - 0.5rem);"
						></div>
						</div>
					</div>
					<input
						id="progress-slider"
						type="range"
						min="0"
						max={work_target_progress}
						bind:value={work_progressValue}
						class="absolute top-0 left-0 w-full h-2 bg-transparent cursor-pointer"
						style="pointer-events: all; opacity: 0; z-index: 1;"
						on:input={() => convertProgressToYTProgress(work_progressValue)}
					/>
				</div>
			</div>
			{/if}
			<div class="w-[80%] flex p-2  justify-center">
				<!-- Input Field -->
				{#if work_type!=="youtube"}
					<div class="flex items-center justify-start w-full">
						<input
						type="number"
						min=0
						max={work_target_progress}
						bind:value={work_progressValue}
						class="mt-2 p-2 border border-gray-300 rounded text-center"
						/>
						{#if work_type==="book"}
							<div class="mt-2 p-3 Ubuntu-font text-sm">
								pages
							</div>
						{/if}
						{#if work_type==="show"}
							<div class="mt-2 p-3 Ubuntu-font text-sm">
								episodes
							</div>
						{/if}
						{#if work_type==="movie"}
							<div class="mt-2 p-3 Ubuntu-font text-sm">
								mins
							</div>
						{/if}
					</div>
					<div class="flex items-center justify-end Ubuntu-font text-sm mt-2 w-full">
						<span class="px-1">{work_target_progress}</span>
						{#if work_type==="book"}
							<span>pages</span>
						{/if}
						{#if work_type==="show"}
							<span>episodes</span>
						{/if}
						{#if work_type==="movie"}
							<span>mins</span>
						{/if}
						
					</div>
				{/if}
				{#if work_type==="youtube"}
					<div class="flex items-center justify-start w-full">
						{#if work_target_progress>3600}
							<input
							type="text"
							bind:value={work_progressValueYT}
							class="mt-2 p-2 border border-gray-300 rounded text-center w-1/3"
							placeholder="hh:mm:ss"
							/>
						{/if}
						{#if work_target_progress<3600}
							<input
							type="text"
							bind:value={work_progressValueYT}
							class="mt-2 p-2 border border-gray-300 rounded text-center w-1/3"
							placeholder="mm:ss"
							/>
						{/if}
						<div class="flex items-center justify-end Ubuntu-font text-sm mt-2 w-full">
							{#if work_target_progress>3600}
								<span class="px-1">{Math.floor(work_target_progress / 60 / 60)}:{Math.floor(work_target_progress / 60 % 60)}:{work_target_progress % 60 % 60}</span>
							{/if}
							{#if work_target_progress<3600}
								<span class="px-1">{Math.floor(work_target_progress / 60)}:{work_target_progress % 60 % 60}</span>
							{/if}
						</div>
					</div>
				{/if}
			</div>
		  </div>
		  <div id='edit-message' class="Ubuntu-font pt-1 pb-2 text-sm text-delete w-full text-center">{$editMessage}</div>
		  <div class="flex justify-center">
			<button
				class="Ubuntu-font bg-cancel mr-4 rounded px-4 py-2 text-sm text-white"
				on:click={cancelEditMedia}
			>
				Cancel
			</button>
			<button
				class="Ubuntu-font bg-background rounded px-4 py-2 text-sm text-white"
				on:click={editMedia}
			>
				Save
			</button>
		</div>
	</div>
</div>

<!--SHARE MODAL-->

<div
	id="share-popup"
	class="bg-background fixed inset-0 flex hidden items-center justify-center bg-opacity-75"
>
	<div class="w-2/5 rounded-lg bg-white p-6 flex-col flex items-center justify-center">
		<div class="Ubuntu-font mb-4 text-center text-xl font-bold">Share your progress with a Friend</div>
		<div class="inline-flex flex-col items-start w-[60%]">
			<div class="Ubuntu-font p-1 text-sm">Your friend's email:</div>
			<input
				type="email"
				placeholder="Email"
				bind:value={friendEmail}
				class="mr-1 rounded-md border p-1.5 w-full"
			/>
		</div>
		<div id='share-message' class="Ubuntu-font p-2 text-sm text-delete">{$shareMessage}</div>
		<div class="flex justify-center">
			<button
				class="Ubuntu-font bg-cancel mr-4 rounded px-4 py-2 text-sm text-white"
				on:click={cancelShareMedia}
			>
				Cancel
			</button>
			<button
				class="Ubuntu-font bg-background rounded px-4 py-2 text-sm text-white"
				on:click={shareMedia}
			>
				Share
			</button>
		</div>
	</div>
</div>
