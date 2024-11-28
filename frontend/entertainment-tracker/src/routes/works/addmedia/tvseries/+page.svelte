<script lang="ts">
	import { goto } from '$app/navigation';
	import { getTVShowByIMDbId, searchTVShowsByTitle } from '$lib/addmedia.api';
	import type { TvSearchResponse } from '$lib/api.model';
	import { createWork } from '$lib/works.api';
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	let query = '';
	let displayedSeries = writable(undefined as TvSearchResponse[] | null | undefined);

	let disabled = false;

	async function searchTvSeries() {
		let response: TvSearchResponse[] = [];

		response = (await searchTVShowsByTitle(query)).body.tvshows;

		$displayedSeries = response;
	}

	onMount(async () => {
		disabled = false;
	});

	async function addSeries(idx: number) {
		const tvSeries = $displayedSeries;
		if (tvSeries === null) return;

		const chosenSeries = (await getTVShowByIMDbId(tvSeries[idx].id)).body.tvshow;

		disabled = true;
		const addResult = await createWork({
			title: tvSeries[idx].title,
			type: 'show',
			current_progress: 0,
			target_progress: chosenSeries.number_of_episodes,
			third_party_id: chosenSeries.id.toString(),
			status: 'watching'
		});

		if (addResult.statusCode < 400) {
			await goto('/works/dashboard');
		}
	}
</script>

<div>
	<form on:submit={searchTvSeries} class="flex flex-col items-center gap-6">
		<div class="w-fit text-left">
			<label class="Inter-font block text-sm" for="query-field">Title</label>
			<input
				bind:value={query}
				id="query-field"
				class="mt-1"
				type="text"
				placeholder="Series title"
			/>
		</div>
		<button class="fancy-button mb-6">Search</button>
	</form>

	<div
		class="grid grid-cols-[auto_1fr_1fr_1fr_auto] items-center gap-y-4 gap-x-2 overflow-auto text-center text-sm"
	>
		{#if $displayedSeries !== undefined}
			<div>Thumbnail</div>
			<div>Title</div>
			<div>First air date</div>
			<div>Score</div>
			<div></div>
			{#if $displayedSeries === null}
				<div class="col-span-5">No series found</div>
			{:else}
				{#each $displayedSeries as series, idx}
					<div class="text-center">
						<img src={series.thumbnail} alt={series.title} class="h-24 w-16" />
					</div>
					<div>{series.title}</div>
					<div>{series.first_air_date}</div>
					<div>{series.vote_average}</div>
					<div class="text-center">
						<button
							class="h-fit w-fit text-center disabled:opacity-35"
							{disabled}
							on:click={() => addSeries(idx)}
							aria-label="Add"
							><img src="/plus-square.svg" alt="plus sign" />
						</button>
					</div>
				{/each}
			{/if}
		{/if}
	</div>
</div>
