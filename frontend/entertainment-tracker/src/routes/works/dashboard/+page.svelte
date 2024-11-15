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

	let checkingDashboard=true;
	let sharingMedia=false;
	let editingMedia=false;

	function shareMedia() {
		checkingDashboard=false;
		sharingMedia=true;
	}

	let friendEmail='';

	function shareWithFriend(){
		console.log('Sharing with:', friendEmail);

		//TODO: implement the proper function, so it's not just placeholder
	}

	let title="";
	let type="";
	let mediaArtSource="/placeholderForEditMediaCoverArt.png";
	let description="";
	let categories!: String[];
	let aimProgress!: number;
	let author="";
	let ISBNnumber=""
	let YTURL="https://www.youtube.com/embed/dQw4w9WgXcQ?si=dourAMMy3-5pBbJr";

	let newProgress!: number;
	let newVideoTitle="hehe";

	function editMedia(currentTitle: string, currentType: string, goal: number) {
		checkingDashboard=false;
		editingMedia=true;
		title=currentTitle;
		type=currentType;
		aimProgress=goal;
		//TODO: get the proper description for the media
		description="Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.";
		//TODO: get the proper categories for the media
		let newCategories=[
			"crimi",
			"fantasy",
			"horror"
		]
		categories=newCategories;

		if(type=="book"){
			//TODO: get the author properly
			author="XY";
			//TODO: get the author properly
			ISBNnumber="ISBN1234567890";
		}

		if(type!="YouTubeVideo"){
			//TODO: get the poster/coverart properly
			mediaArtSource="/placeholderForEditMediaCoverArt.png";
		}
		else {
			//TODO: get the proper url that's added by the user
			YTURL="https://youtu.be/dQw4w9WgXcQ?feature=shared"

			//TODO: REMOVE. These are for tests to see if the embedding works for diff vids
			//YTURL="https://youtu.be/XGNeFrnhlvw?feature=shared"
			//YTURL="https://youtu.be/Y1ujpoDlgRU?feature=shared"

			//Converts the shared URL to the embed URL~
			let videoId = YTURL.split("https://youtu.be/")[1].split("?")[0];

			YTURL="https://www.youtube.com/embed/" + videoId + "?si=dourAMMy3-5pBbJr";
		}

	}

	function mediaEdited(){
		if(type=="YouTubeVideo"){
			console.log("New title for video: " + newVideoTitle);
		}
		else{
			console.log("New progress: " + newProgress);
		}
	}

	function deleteMedia() {
		//TODO: write delete media function
	}

	function returnToDashboard(){
		checkingDashboard=true;
		sharingMedia=false;
		editingMedia=false;
	}
</script>

{#if checkingDashboard}
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
					<button class="edit-button" on:click={() => editMedia(work.title, work.type, work.target_progress)}>
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
{/if}
{#if sharingMedia}
<div class="relative z-0 flex flex-grow items-center justify-center px-80 py-3">
	<div class="h-full w-full rounded-lg bg-white">
		<div class="flex flex-col">
			<div class="flex items-start justify-between p-2">
				<button class="flex items-center justify-center border-0 pl-6 pt-4" on:click={returnToDashboard}>
					<img src="/back-button.png" alt="Return to dashboard" class="w-5 h-5"/>
				</button>
				<div class="flex-grow text-center text-lg font-bold Ubuntu-font pt-3"> 
					Share your progress with a friend
				</div>
			</div>
		</div>
		<div class="w-full p-2 flex flex-col justify-center items-center">
			<div class="inline-flex flex-col ml-20">
				<div class="Ubuntu-font text-sm p-1">
					Your friend's email:
				</div>
				<div>
					<input type="email" placeholder="Email" bind:value={friendEmail} class="border rounded-md p-1.5 mr-1" /> 
					<button class="bg-background text-white rounded-md py-1.5 px-6" on:click={shareWithFriend}> 
						Share
					</button>
				</div>
			</div>
		</div>
	</div>
</div>
{/if}
{#if editingMedia}
<div class="relative z-0 flex flex-grow items-center justify-center px-80 py-3">
	<div class="h-full w-full rounded-lg bg-white">
		<div class="flex flex-col">
			<div class="flex items-start justify-between p-2">
				<button class="flex items-center justify-center border-0 pl-6 pt-4" on:click={returnToDashboard}>
					<img src="/back-button.png" alt="Return to dashboard" class="w-5 h-5"/>
				</button>
			</div>
			<div></div>
		</div>
		{#if type!=="YouTubeVideo"}
		<div class="flex p-4 items-start ">
			<div class="flex flex-col ml-10 mr-4">
				<img src={mediaArtSource} alt="Return to dashboard" class="w-full rounded-md"/>
				{#if type==="book"}
					<div class="text-center text-xxs Ubuntu-font pt-1">
						{ISBNnumber}
					</div>
				{/if}
				<div class="text-center text-xxs Ubuntu-font p-1">
					Categories:
				</div>
				{#each categories as category}
					<div class="text-center text-xxs Ubuntu-font pb-1">
						{category}
					</div>
				{/each}
			</div>
			<div class="mr-8 ml-4">
				<div class="text-center text-sm font-bold Ubuntu-font">
					{title}
				</div>
				{#if type==="book"}
					<div class="text-center text-xs font-bold Ubuntu-font">
						{author}
					</div>
				{/if}
				<div class="text-start text-xxs Ubuntu-font pt-3" style="line-height: 2.5;">
					{description}
				</div>
				<div class="Ubuntu-font text-sm font-bold pt-6 pb-2">
					Progress
				</div>
				<div>
					{#if type==="TVshow"}
						<input type="progress" placeholder="Episode number / {aimProgress}" bind:value={newProgress} class="border rounded-md p-1.5 mr-1 text-sm" /> 
					{/if}
					{#if type==="book"}
						<input type="progress" placeholder="Page number / {aimProgress}" bind:value={newProgress} class="border rounded-md p-1.5 mr-1 text-sm" /> 
					{/if}
					{#if type==="movie"}
						<input type="progress" placeholder="Minutes / {aimProgress}" bind:value={newProgress} class="border rounded-md p-1.5 mr-1 text-sm" /> 
					{/if}
					<button class="bg-background text-white rounded-md py-1.5 px-6 text-sm" on:click={mediaEdited}> 
						Save
					</button>
				</div>
			</div>
		</div>
		{/if}
		{#if type==="YouTubeVideo"}
			<div class="flex flex-col mt-6 ml-24 mr-24">
				<div class="flex justify-center items-center">
					<iframe title="video" class="w-full aspect-[18/10]"
						src={YTURL}>
					</iframe>
				</div>
				<div class="Ubuntu-font text-xs pt-6">
					Set Display name
				</div>
				<div class="flex pt-1">
					<input type="displayTitle" placeholder="Name" bind:value={newVideoTitle} class="border rounded-md p-1.5 mr-1 text-sm w-64"/> 
					<button class="bg-background text-white rounded-md py-1.5 px-6 text-sm" on:click={mediaEdited}> 
						Save
					</button>
				</div>
			</div>
		{/if}
	</div>
</div>
{/if}

