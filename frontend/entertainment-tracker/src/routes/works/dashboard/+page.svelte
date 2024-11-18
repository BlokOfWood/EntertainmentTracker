<script lang=ts>
	import type { Work, UpdateWorkRequest } from '$lib/api.model';
	import { deleteWork, getWorks, updateWork } from '$lib/works.api';
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
				const movieExample: Work = { id: 23132, third_party_id: 'tp789', title: 'Movie', status: 'Pending', type: 'movie', current_progress: 10, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(movieExample);
				const bookExample4: Work = { id: 1343456, third_party_id: 'ISBN12:32749632', title: 'The Shadow over Innsmouth', status: 'Pending', type: 'book', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(bookExample4);
				const TVShowExample: Work = { id: 235478, third_party_id: 'tp789', title: 'TVShow', status: 'Pending', type: 'TVshow', current_progress: 20, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(TVShowExample);
				const bookExample3: Work = { id: 1343456, third_party_id: 'ISBN12:32749632', title: 'Strange Case of Dr Jekyll and Mr Hyde', status: 'Pending', type: 'book', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(bookExample3);
				const YouTubeVideoExample: Work = { id: 4365879, third_party_id: 'https://youtu.be/dQw4w9WgXcQ?feature=shared', title: 'YouTubeVideo', status: 'Pending', type: 'YouTubeVideo', current_progress: 80, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(YouTubeVideoExample);
				const TVShowExampleSecond: Work = { id: 345768, third_party_id: 'tp789', title: 'TVShowSecond', status: 'Pending', type: 'TVshow', current_progress: 15, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(TVShowExampleSecond);
				const bookExample: Work = { id: 1343456, third_party_id: 'ISBN12:32749632', title: 'Book', status: 'Pending', type: 'book', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(bookExample);
				const YouTubeVideoExampleSecond: Work = { id: 354678987, third_party_id: 'https://youtu.be/Y1ujpoDlgRU?feature=shared', title: 'YouTubeVideo2', status: 'Pending', type: 'YouTubeVideo', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(YouTubeVideoExampleSecond);
				const bookExample2: Work = { id: 1343456, third_party_id: 'ISBN12:32749632', title: 'Dracula', status: 'Pending', type: 'book', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(bookExample2);

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

	let currentWork!: Work;
	let mediaArtSource="/placeholderForEditMediaCoverArt.png";
	let description="";
	let categories!: String[];
	let author="";
	let YTURL="https://www.youtube.com/embed/dQw4w9WgXcQ?si=dourAMMy3-5pBbJr";

	let newProgress=-1;
	let newVideoTitle="";

	function editMedia(work: Work) {
		currentWork=work;
		checkingDashboard=false;
		editingMedia=true;
		//TODO: get the proper description for the media
		description="Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.";
		//TODO: get the proper categories for the media
		let newCategories=[
			"crimi",
			"fantasy",
			"horror"
		]
		categories=newCategories;

		if(work.type=="book"){
			//TODO: get the author properly
			author="XY";
		}

		if(work.type!="YouTubeVideo"){
			//TODO: get the poster/coverart properly
			mediaArtSource="/placeholderForEditMediaCoverArt.png";
		}
		else {
			YTURL=work.third_party_id;

			//Converts the shared URL to the embed URL~
			let videoId = YTURL.split("https://youtu.be/")[1].split("?")[0];

			YTURL="https://www.youtube.com/embed/" + videoId + "?si=dourAMMy3-5pBbJr";
		}

	}

	function setNewProgress(event: Event){
		const target = event.target as HTMLInputElement;
		const value = target.value;
		newProgress=Number(value);
	}

	function setNewTitle(event: Event){
		const target = event.target as HTMLInputElement;
		const value = target.value;
		newVideoTitle=value;
	}

	function mediaEdited(){
		let newDetails: UpdateWorkRequest = {
			title: currentWork.title,
			type: currentWork.type,
			status: currentWork.status,
			current_progress: currentWork.current_progress,
			target_progress: currentWork.target_progress
		};

		if(newProgress!=-1){
			newDetails.current_progress=newProgress;
			console.log("New progress: " + newDetails.current_progress);
		}

		if(currentWork.type=="YouTubeVideo" && newVideoTitle!=""){
			newDetails.title=newVideoTitle;
			console.log("New title for video: " + newDetails.title);
		}

		//TODO: uncomment this
		updateWork(currentWork.id, newDetails);

		//reset these values so it can be checked wether the user filled the fields or not
		newProgress=-1;
		newVideoTitle="";

		returnToDashboard();
	}

	let idOfDeletedWork!: number;

	function deleteMedia(id: number) {
		console.log("deleting media");
		idOfDeletedWork=id;

		const popup = document.getElementById('popup');
		if (popup){
			popup.classList.remove('hidden');
		}
	}

	function mediaDeleted(){
		const popup = document.getElementById('popup');
		if (popup){
			popup.classList.add('hidden');
		}
		
		//TODO: uncomment this
		//deleteWork(idOfDeletedWork);
		console.log("Media deleted");
		returnToDashboard();
	}

	function cancelDeletingMedia(){
		const popup = document.getElementById('popup');
		if (popup){
			popup.classList.add('hidden');
		}
	}

	function returnToDashboard(){
		checkingDashboard=true;
		sharingMedia=false;
		editingMedia=false;

		getWorks().then((response) => {
			if (response.ok) {
				console.log('Works fetched successfully');
				console.log(response.body);
				works=response.body.mediaEntries;

				//TODO: remove example test works once adding works, deleting works and edit works are implemented.
				const movieExample: Work = { id: 23132, third_party_id: 'tp789', title: 'Movie', status: 'Pending', type: 'movie', current_progress: 10, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(movieExample);
				const bookExample4: Work = { id: 1343456, third_party_id: 'ISBN12:32749632', title: 'The Shadow over Innsmouth', status: 'Pending', type: 'book', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(bookExample4);
				const TVShowExample: Work = { id: 235478, third_party_id: 'tp789', title: 'TVShow', status: 'Pending', type: 'TVshow', current_progress: 20, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(TVShowExample);
				const bookExample3: Work = { id: 1343456, third_party_id: 'ISBN12:32749632', title: 'Strange Case of Dr Jekyll and Mr Hyde', status: 'Pending', type: 'book', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(bookExample3);
				const YouTubeVideoExample: Work = { id: 4365879, third_party_id: 'https://youtu.be/dQw4w9WgXcQ?feature=shared', title: 'YouTubeVideo', status: 'Pending', type: 'YouTubeVideo', current_progress: 80, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(YouTubeVideoExample);
				const TVShowExampleSecond: Work = { id: 345768, third_party_id: 'tp789', title: 'TVShowSecond', status: 'Pending', type: 'TVshow', current_progress: 15, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(TVShowExampleSecond);
				const bookExample: Work = { id: 1343456, third_party_id: 'ISBN12:32749632', title: 'Book', status: 'Pending', type: 'book', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(bookExample);
				const YouTubeVideoExampleSecond: Work = { id: 354678987, third_party_id: 'https://youtu.be/Y1ujpoDlgRU?feature=shared', title: 'YouTubeVideo2', status: 'Pending', type: 'YouTubeVideo', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(YouTubeVideoExampleSecond);
				const bookExample2: Work = { id: 1343456, third_party_id: 'ISBN12:32749632', title: 'Dracula', status: 'Pending', type: 'book', current_progress: 40, target_progress: 100, version: 1, created_at: Date.now(), updated_at: new Date() };
				works.push(bookExample2);

				originalWorks=works;
			}
		})
	}
</script>

{#if checkingDashboard}
<div class="relative z-0 flex flex-grow items-center justify-center px-40 py-3 h-full">
	<div class="h-full w-full rounded-lg bg-white flex flex-col overflow-auto">
		<div class="grid grid-cols-4 w-full sticky top-0 bg-white z-10">
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
		<div class="grid grid-cols-4 w-full">
			{#each works as work}
				<div class="border-0 p-2 text-black text-lg Ubuntu-font flex justify-center items-center text-center">{work.title}</div>
				<div class="border-0 p-2 text-black text-lg Ubuntu-font flex justify-center items-center text-center">{work.type}</div>
				<div class="border-0 p-2 flex justify-center items-center">
					{#if work.target_progress==work.current_progress}
					<img src="/done.png" alt="Done" class="w-4 h-4" />
					{/if}
					{#if work.target_progress!=work.current_progress && work.type=="book"}
					<div class="text-black text-lg Ubuntu-font">
						{work.current_progress}/{work.target_progress} pages
					</div>
					{/if}
					{#if work.target_progress!=work.current_progress && work.type=="TVshow"}
					<div class="text-black text-lg Ubuntu-font">
						{work.current_progress}/{work.target_progress} episodes
					</div>
					{/if}
					{#if work.target_progress!=work.current_progress && work.type=="movie"}
					<div class="text-black text-lg Ubuntu-font">
						{work.current_progress}/{work.target_progress} mins
					</div>
					{/if}
					{#if work.target_progress!=work.current_progress && work.type=="YouTubeVideo"}
					<div class="text-black text-lg Ubuntu-font">
						{work.current_progress} mins
					</div>
					{/if}
				</div>
				<div class="flex justify-center items-center space-x-5 p-2">
					<button class="share-button" on:click={shareMedia}>
						<img src="/share.png" alt="Share" class="w-5 h-5" />
					</button>
					<button class="edit-button" on:click={() => editMedia(work)}>
						<img src="/edit.png" alt="Edit" class="w-5 h-5" />
					</button>
					<button class="delete-button" on:click={() => deleteMedia(work.id)}>
						<img src="/trash.png" alt="Delete" class="w-5 h-5" />
					</button>
				</div>
			{/each}
		</div>
	</div>
</div>
<div id="popup" class="fixed inset-0 bg-background bg-opacity-75 hidden flex items-center justify-center">
    <div class="bg-white rounded-lg p-6 w-1/3">
        <div class="Ubuntu-font text-lg font-bold mb-4 text-center">Delete Media Entry</div>
        <div class="Ubuntu-font mb-4">Are you sure you'd like to delete this media entry?</div>
		<div class="flex justify-center">
			<button class="bg-red-500 text-white Ubuntu-font px-4 py-2 rounded mr-4" on:click={mediaDeleted}>
				Delete
			</button>
			<button class="text-white bg-slate-500 Ubuntu-font px-4 py-2 rounded" on:click={cancelDeletingMedia}>
				Cancel
			</button>
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
<div class="relative z-0 flex flex-grow items-center justify-center py-3">
	<div class="h-full w-full rounded-lg bg-white max-w-screen-sm">
		<div class="flex flex-col">
			<div class="flex items-start justify-between p-2">
				<button class="flex items-center justify-center border-0 pl-6 pt-4" on:click={returnToDashboard}>
					<img src="/back-button.png" alt="Return to dashboard" class="w-5 h-5"/>
				</button>
			</div>
			<div></div>
		</div>
		{#if currentWork.type!=="YouTubeVideo"}
		<div class="flex p-4 items-start ">
			<div class="flex flex-col ml-10 mr-4">
				<img src={mediaArtSource} alt="Return to dashboard" class="w-full rounded-md"/>
				{#if currentWork.type==="book"}
					<div class="text-center text-xxs Ubuntu-font pt-1">
						{currentWork.third_party_id}
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
					{currentWork.title}
				</div>
				{#if currentWork.type==="book"}
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
					{#if currentWork.type==="TVshow"}
						<input type="progress" placeholder="Episode number / {currentWork.target_progress}" on:input={setNewProgress} class="border rounded-md p-1.5 mr-1 text-sm" /> 
					{/if}
					{#if currentWork.type==="book"}
						<input type="progress" placeholder="Page number / {currentWork.target_progress}" on:input={setNewProgress} class="border rounded-md p-1.5 mr-1 text-sm" /> 
					{/if}
					{#if currentWork.type==="movie"}
						<input type="progress" placeholder="Minutes / {currentWork.target_progress}" on:input={setNewProgress} class="border rounded-md p-1.5 mr-1 text-sm" /> 
					{/if}
					<button class="bg-background text-white rounded-md py-1.5 px-6 text-sm" on:click={mediaEdited}> 
						Save
					</button>
				</div>
			</div>
		</div>
		{/if}
		{#if currentWork.type==="YouTubeVideo"}
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
					<input type="displayTitle" placeholder="Name" on:input={setNewTitle} class="border rounded-md p-1.5 mr-1 text-sm w-64"/> 
				</div>
				<div class="Ubuntu-font text-xs pt-2 pb-1">
					Progress
				</div>
				<div class="flex">
					<input type="progress" placeholder="Minutes" on:input={setNewProgress} class="border rounded-md p-1.5 mr-1 text-sm" /> 
					<button class="bg-background text-white rounded-md py-1.5 px-6 text-sm" on:click={mediaEdited}> 
						Save
					</button>
				</div>
			</div>
		{/if}
	</div>
</div>
{/if}

