import type { Work, SharedWork} from '$lib/api.model';
import { getWorks, getSharedWorks } from '$lib/works.api';
import { getBookByISBN, getTVShowByIMDbId, getMovie, getBookByGoogleId } from '$lib/addmedia.api';


export type WorkPlus = {
    work: Work;
    shared: boolean;
    sharedBy: string;
    thumbnail: string;
};

export async function fetchWorks() {
    let works: WorkPlus[] = []; // Now just an array of WorkPlus

    let currentNotSharedWorks: Work[] = [];

    const responseNotShared = await getWorks();
    if (responseNotShared.ok) {
        console.log('not shared works fetched successfully');
        currentNotSharedWorks = responseNotShared.body.mediaEntries;

        // Use a for...of loop to handle async await
        for (const currentNotSharedWork of currentNotSharedWorks) {
            let source = await getMediaArtSource(currentNotSharedWork);
            
            const workPlus: WorkPlus = {
                work: currentNotSharedWork,
                shared: false,
                sharedBy: '',
                thumbnail: source
            };
            works.push(workPlus); // Push directly to the works array
        }
        
    }

    let currentSharedWorks: SharedWork[] = [];

    const responseShared = await getSharedWorks();
    if (responseShared.ok) {
        console.log('shared works fetched successfully');
        currentSharedWorks = responseShared.body.sharedEntries;
        if (currentSharedWorks != null) {
            // Use a for...of loop here as well
            for (const currentSharedWork of currentSharedWorks) {
                let source = await getMediaArtSource(currentSharedWork.media_entry);
            
                const workPlus: WorkPlus = {
                    work: currentSharedWork.media_entry,
                    shared: true,
                    sharedBy: '',
                    thumbnail: source
                };
                works.push(workPlus); // Push directly to the works array
            }
            
        }
    }

    works.sort((a, b) => {
        if (a.work.created_at > b.work.created_at) return -1; // a comes before b
        if (a.work.created_at < b.work.created_at) return 1; // a comes after b
        return 0; // a and b are equal
    });
    
    return works;
}

export async function getMediaArtSource(work: Work) {
    let source = "";

    if (work.type === 'book') {
        if (work.third_party_id !== "") {
            try {
                const response = await getBookByGoogleId(work.third_party_id);
                const currentBook = response.body.book; 
                if (currentBook) { 
                    source = currentBook.thumbnail;
                    console.log("Thumbnail aquired.");
                } else {
                    console.error("Book data is not available.");
                    const isbnResponse = await getBookByISBN(work.third_party_id);
                    const isbnBook = isbnResponse.body.book; 
                    if (isbnBook) { 
                        source = isbnBook.thumbnail;
                    }
                }
            } catch (error) {
                console.error("Error fetching book data:", error);
            }
        }
    } else if (work.type === 'movie') {
        if (work.third_party_id !== "") {
            try {
                const response = await getMovie(Number(work.third_party_id));
                const currentMovie = response.body.movie; 
                if (currentMovie) { 
                    source = currentMovie.thumbnail;
                    console.log("Thumbnail aquired.");
                } else {
                    console.error("Movie data is not available");
                }    
            } catch (error) {
                console.error("Error fetching movie data:", error);
            }
        }
    } else if (work.type === 'show') {
        if (work.third_party_id !== "") {
            try {
                const response = await getTVShowByIMDbId(Number(work.third_party_id));
                const currentShow = response.body.tvshow; 
                if (currentShow) { 
                    source = currentShow.thumbnail;
                    console.log("Thumbnail aquired.");
                } else {
                    console.error("Show data is not available");
                }    
            } catch (error) {
                console.error("Error fetching show data:", error);
            }
        }
    } else if (work.type === 'youtube') {
        source = `https://www.youtube.com/embed/${work.third_party_id}?si=dourAMMy3-5pBbJr`;
        console.log("Thumbnail aquired.");
    }

    return source;
}

export function sortByTitle(sortedByTitle: number, originalWorks: WorkPlus[]) {
    let sortedWorks: WorkPlus[]; // Declare works variable

    if (sortedByTitle == 2) {
        sortedByTitle = 0;
        sortedWorks = originalWorks; // Reset to original works
        console.log('not sorted by title');
    } else if (sortedByTitle == 0) {
        sortedByTitle = 1;

        // Sort and create a new reference for works in ascending order
        sortedWorks = [...originalWorks].sort((a, b) => {
            if (a.work.title < b.work.title) return -1; // a comes before b
            if (a.work.title > b.work.title) return 1; // a comes after b
            return 0; // a and b are equal
        });

        console.log('sort by title (asc)');
    } else {
        sortedByTitle = 2;

        // Sort and create a new reference for works in descending order
        sortedWorks = [...originalWorks].sort((a, b) => {
            if (a.work.title > b.work.title) return -1; // a comes before b
            if (a.work.title < b.work.title) return 1; // a comes after b
            return 0; // a and b are equal
        });

        console.log('sort by title (desc)');
    }

    return { sortedByTitle, sortedWorks };
}

export function sortByType(sortedByType : number, originalWorks : WorkPlus[]){
    let sortedWorks: WorkPlus[];
    
    if (sortedByType==2) {
        sortedByType = 0;
        sortedWorks = originalWorks;
        console.log('not sorted by type');
    } else if (sortedByType==0){
        sortedByType = 1;

        // Sort and create a new reference for works
        sortedWorks = originalWorks;
        sortedWorks = [...originalWorks].sort((a, b) => {
            if (a.work.type < b.work.type) return -1; // a comes before b
            if (a.work.type > b.work.type) return 1; // a comes after b
            return 0; // a and b are equal
        });

        console.log('sort by type (asc)');
    }
    else{
        sortedByType = 2;

        // Sort and create a new reference for works
        sortedWorks = [...originalWorks].sort((a, b) => {
            if (a.work.type > b.work.type) return -1; // a comes before b
            if (a.work.type < b.work.type) return 1; // a comes after b
            return 0; // a and b are equal
        });

        console.log('sort by type (desc)');
    }

    return { sortedByType, sortedWorks };
}

export function sortByProgress(sortedByProgress : number, originalWorks : WorkPlus[]){
    let sortedWorks: WorkPlus[];
    
    if (sortedByProgress==2) {
        sortedByProgress = 0;
        sortedWorks = originalWorks;
        console.log('not sorted by progress');
    } else if(sortedByProgress==0) {
        sortedByProgress = 1;

        // Sort and create a new reference for works
        sortedWorks = [...originalWorks].sort((a, b) => {
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

        console.log('sort by progress (desc)');
    }
    else{
        sortedByProgress = 2;

        // Sort and create a new reference for works
        sortedWorks = [...originalWorks].sort((a, b) => {
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

        console.log('sort by progress (asc)');
    }

    return { sortedByProgress, sortedWorks };
}

export function sortByShared(sortedByShared : number, originalWorks : WorkPlus[]){
    let sortedWorks: WorkPlus[];
    
    if (sortedByShared==2) {
        sortedByShared = 0;
        sortedWorks=originalWorks;
        console.log('not sorted by shared');
    } else if (sortedByShared==0){
        sortedByShared = 1;

        // Sort and create a new reference for works
        sortedWorks = [...originalWorks].sort((a, b) => {
            if (!a.shared && b.shared) return 1; // a comes before b
            if (a.shared && !b.shared) return -1; // a comes after b
            return 0; // a and b are equal
        });

        console.log('sort by title (asc)');
    }
    else{
        sortedByShared = 2;

        // Sort and create a new reference for works
        sortedWorks = [...originalWorks].sort((a, b) => {
            if (!a.shared && b.shared) return -1; // a comes before b
            if (a.shared && !b.shared) return 1; // a comes after b
            return 0; // a and b are equal
        });

        console.log('sort by title (desc)');
    }

    return { sortedByShared, sortedWorks };
}