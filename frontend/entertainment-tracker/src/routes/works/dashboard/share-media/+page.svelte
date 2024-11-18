<script lang="ts">
    import { onMount } from 'svelte';
    import { shareWork } from '$lib/works.api';
    import type { Work, ShareWorkRequest } from '$lib/api.model';
    import { goto } from '$app/navigation';

    let currentWork!: Work;

    onMount(() => {
        currentWork = history.state['sveltekit:states'].work;
        console.log(currentWork);
    });

    let friendEmail = '';

	function shareWithFriend() {
		console.log('Sharing with:', friendEmail);

		let sharedWork: ShareWorkRequest = {
			media_entry: currentWork.id,
			share_with: friendEmail
		};

		shareWork(sharedWork)
	}

    function returnToDashboard(){
        goto('/works/dashboard');
    }
</script>

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