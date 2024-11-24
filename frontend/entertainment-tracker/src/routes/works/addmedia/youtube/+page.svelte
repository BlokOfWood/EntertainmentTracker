<script lang="ts">
	import { goto } from '$app/navigation';
	import { getYoutubeVideo } from '$lib/addmedia.api';
	import { createWork } from '$lib/works.api';
	import { onMount } from 'svelte';
	
	let query = '';

	let disabled = false;
	let error: string | null = null;

	onMount(async () => {
		disabled = false;
	});

	async function addVideo() {
		disabled = true;

		try {
			const url = new URL(query);
			if (url.hostname !== 'www.youtube.com' && url.hostname !== 'youtube.com' && url.hostname !== 'm.youtube.com' && url.hostname !== 'youtu.be' && url.hostname !== 'www.youtu.be') {
				console.log(url.hostname);
				error = 'Not a YouTube link.';
				disabled = false;
				return;
			}

			let videoId;
		    if(url.searchParams.has('v')) {
				videoId = url.searchParams.get('v');
			} else {
				videoId = url.pathname.substring(1);
			}

			if (videoId === null) {
				error = 'Invalid YouTube link.';
				disabled = false;
				return;
			}

			const video = (await getYoutubeVideo(videoId)).body.video;
			if (video === null) {
				error = 'Video not found';
				disabled = false;
				return;
			} else {
				console.log(video);
			}
			const addResult = await createWork({
				title: video.title,
				type: 'youtube',
				current_progress: 0,
				target_progress: video.duration,
				third_party_id: video.video_id,
				status: 'watching'
			});

			if (addResult.statusCode < 400) {
				await goto('/works/dashboard');
			}
		} catch (e) {
			console.error(e);

			error = 'Invalid URL';
			disabled = false;
			return;
		}
	}
</script>

<div>
	<form on:submit={addVideo} class="flex flex-col items-center gap-6">
		<div class="w-max text-left">
			<label class="Inter-font block text-sm" for="query-field">YouTube URL</label>
			<input
				bind:value={query}
				id="query-field"
				class="mt-1 w-96"
				type="text"
				placeholder="YouTube URL"
			/>
			{#if error !== null}
				<div class="text-delete Inter-font mt-0.5 text-sm">{error}</div>
			{/if}
		</div>
		<button class="fancy-button mb-6">Add Video</button>
	</form>
</div>
