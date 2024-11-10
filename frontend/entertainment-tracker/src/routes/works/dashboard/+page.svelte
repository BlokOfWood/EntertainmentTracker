<script>
	import Header from '$lib/header.svelte';
	import { getWorks } from '$lib/works.api';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	//TODO: get the media from backend to loop through it

	onMount(async () => {
		getWorks().then((response) => {
			if (response.ok) {
				console.log('Works fetched successfully');
				console.log(response.body);
			}
		});
	});

	function sortByTitle() {
		//TODO: write sort by title function
	}

	function sortByType() {
		//TODO: write sort by type function
	}

	function sortByProgress() {
		//TODO: write sort by progress
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

<div class="bg-background relative flex h-screen flex-col">
	<Header />
	<div class="relative z-0 flex flex-grow items-center justify-center px-40 py-3">
		<div class="grid h-full w-full grid-cols-4 rounded-lg bg-white">
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

			<!--TODO: check and correct it once media can be fetched from backend and replace the variables with the proper ones
            I just did it with parasztlogika, couldn't test it-->
			<!--{#each media as medium}
                <div class="border-0 p-2 text-black text-lg Ubuntu-font">{medium.title}</div>
                <div class="border-0 p-2 text-black text-lg Ubuntu-font">{medium.type}</div>
                <div class="border-0 p-2 text-black text-lg Ubuntu-font">{medium.progress}</div>
                <div class="border-0 p-2 flex justify-center items-start">
                    {#if medium.fullLength==medium.process}
                    <img src="/done.png" alt="Done" class="w-4 h-4" />
                    {/if}
                    {#if medium.fullLength!=medium.process && medium.type=="book"}
                    <div class=" text-black text-lg Ubuntu-font">
                        {medium.process}/{medium.fullLenght} pages
                    </div>
                    {/if}
                    {#if medium.fullLength!=medium.process && medium.type=="TVshow"}
                    <div class=" text-black text-lg Ubuntu-font">
                        {medium.process}/{medium.fullLenght} episodes
                    </div>
                    {/if}
                    {#if medium.fullLength!=medium.process && medium.type=="movie"}
                    <div class=" text-black text-lg Ubuntu-font">
                        {medium.process}/{medium.fullLenght} mins
                    </div>
                    {/if}
                    {#if medium.fullLength!=medium.process && medium.type=="YouTubeVideo"}
                    <div class=" text-black text-lg Ubuntu-font">
                        {medium.process} mins
                    </div>
                    {/if}
                </div>
            <div class="flex justify-center items-start space-x-5 p-2">
                TODO: give the images their funcionality /either as button or href, we'll see, also put this in comment once the foreach is no longer commented out/
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
            {/each}-->
		</div>
	</div>
</div>
