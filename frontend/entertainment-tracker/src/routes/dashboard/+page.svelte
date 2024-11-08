<script>
    import Header from '../header.svelte';
    import { writable } from 'svelte/store';
    
    //TODO: get the media from backend to loop through it

    //TODO: these are all placeholders, take them out

    let fullLenght=100;
    let process=20;
    let mediaType="YouTubeVideo";
</script>

<div class="relative h-screen flex flex-col bg-background">
    <Header/>
    <div class="relative flex-grow flex items-center justify-center py-3 px-40 z-0">
        <div class="w-full h-full bg-white rounded-lg grid grid-cols-4">
            <div class="border-0 p-2 flex justify-center items-start">
                <div class="text-black text-lg font-bold Ubuntu-font flex items-center space-x-2">
                    <span>Title</span>
                    <img src="/chevron-down.png" alt="Chevron-down" class="w-4 h-4" />
                </div>
            </div>
            <div class="border-0 p-2 flex justify-center items-start">
                <div class="text-black text-lg font-bold Ubuntu-font flex items-center space-x-2">
                    <span>Type</span>
                    <img src="/chevrons-up-down.png" alt="Chevron-down" class="w-4 h-4" />
                </div>
            </div>
            <div class="border-0 p-2 flex justify-center items-start">
                <div class="text-black text-lg font-bold Ubuntu-font flex items-center space-x-2">
                    <span>Progress</span>
                    <img src="/chevrons-up-down.png" alt="Chevron-down" class="w-4 h-4" />
                </div>
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
                        {process}/{fullLenght} pages
                    </div>
                    {/if}
                    {#if medium.fullLength!=medium.process && medium.type=="TVshow"}
                    <div class=" text-black text-lg Ubuntu-font">
                        {process}/{fullLenght} episodes
                    </div>
                    {/if}
                    {#if medium.fullLength!=medium.process && medium.type=="movie"}
                    <div class=" text-black text-lg Ubuntu-font">
                        {process}/{fullLenght} mins
                    </div>
                    {/if}
                    {#if medium.fullLength!=medium.process && medium.type=="YouTubeVideo"}
                    <div class=" text-black text-lg Ubuntu-font">
                        {process} mins
                    </div>
                    {/if}
                </div>
                <div class="flex items-center space-x-5">
                    TODO: give the images their funcionality /either as button or href, we'll see, also put this in comment once the foreach is no longer commented out/
                    <img src="/share.png" alt="Share" class="w-5 h-5" />
                    <img src="/edit.png" alt="Edit" class="w-5 h-5" />
                    <img src="/trash.png" alt="Delete" class="w-5 h-5" />
                </div>
            </div>
            </div>
            {/each}-->
        </div>
    </div>
</div>
