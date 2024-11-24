<script lang="ts">
    import { onMount } from 'svelte';
    import { updateWork } from '$lib/works.api';
    import type { Book, Movie, TvShow, Work, UpdateWorkRequest } from '$lib/api.model';
    import { getBookByISBN, getTVShowByIMDb, getTVShowByIMDbId, getMovieByIMDb, getMovie } from '$lib/addmedia.api';
    import { writable } from 'svelte/store';
    import { goto } from '$app/navigation';

    let currentWork!: Work;

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


    onMount(() => {
        currentWork = history.state['sveltekit:states'].work;
        console.log(currentWork);

        if (currentWork.type == 'book') {
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

		if (currentWork.type == 'show') {
			if(currentWork.third_party_id!=""){
				getTVShowByIMDbId(Number(currentWork.third_party_id)).then(response => {
					const currentTVShow = response.body.tvshow; // Get the book object directly
					tvShow.set(response.body.tvshow); // Use .set to update the store
					if (currentTVShow) { // Check if currentBook is not null
					description = currentTVShow.overview; // Access description
					mediaArtSource = currentTVShow.thumbnail; // Access thumbnail
					categories = currentTVShow.genres; // Access categories

				} else {
					// Handle the case where currentBook is null
					console.error("TwShow data is not available");

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

		if (currentWork.type == 'movie') {
			if (currentWork.third_party_id!=""){
				getMovie(Number(currentWork.third_party_id)).then(response => {
					const currentMovie = response.body.movie; // Get the book object directly
					movie.set(response.body.movie); // Use .set to update the store
					if (currentMovie) { // Check if currentBook is not null
					description = currentMovie.overview; // Access description
					mediaArtSource = currentMovie.thumbnail; // Access thumbnail
					categories = currentMovie.genres; // Access categories

				} else {
					// Handle the case where currentBook is null
					console.error("Book data is not available");

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

		if (currentWork.type == 'youtube') {
			YTURL = `https://www.youtube.com/embed/${currentWork.third_party_id}?si=dourAMMy3-5pBbJr`;
		}
    });

    function setNewProgress(event: Event) {
		const target = event.target as HTMLInputElement;
		const value = target.value;
		newProgress = Number(value);
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

		updateWork(currentWork.id, newDetails);

		//reset these values so it can be checked wether the user filled the fields or not
		newProgress = -1;

		returnToDashboard();
	}	

    function returnToDashboard(){
        goto('/works/dashboard');
    }
</script>

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
                        {#if currentWork.type === 'show'}
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
        {#if currentWork && currentWork.type === 'youtube'}
            <div class="ml-24 mr-24 mt-6 flex flex-col">
                <div class="flex items-center justify-center">
                    <iframe title="video" class="aspect-[18/10] w-full" src="{YTURL}"> </iframe>
                </div>
                <div class="Ubuntu-font pb-2 pt-6 text-sm font-bold">Progress</div>
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