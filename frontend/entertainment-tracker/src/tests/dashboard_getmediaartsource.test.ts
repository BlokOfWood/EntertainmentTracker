import { describe, it, expect, vi, afterEach } from 'vitest';
import type { WorkPlus } from '../routes/works/dashboard/dashboard';
import { fetchWorks, getMediaArtSource } from '../routes/works/dashboard/dashboard';
import type { Work, SharedWork } from '$lib/api.model';
import { getWorks, getSharedWorks } from '$lib/works.api';
import { getBookByGoogleId, getBookByISBN, getMovie, getTVShowByIMDbId } from '$lib/addmedia.api'

/*-------------TESTING GET MEDIA ART SOURCE----------------*/

// Mock the API calls
vi.mock('$lib/addmedia.api', () => ({
    getBookByGoogleId: vi.fn(),
    getBookByISBN: vi.fn(),
    getMovie: vi.fn(),
    getTVShowByIMDbId: vi.fn(),
}));

describe('getMediaArtSource', () => {
    afterEach(() => {
        vi.clearAllMocks();
    });

    const book : Work = {
        id: 1,
        third_party_id: '12345',
        title: 'Sample Work',
        status: 'in-progress',
        type: 'book', // Replace with actual WorkType if necessary
        current_progress: 50,
        target_progress: 100,
        version: 1,
        created_at: Date.now(),
        updated_at: new Date(),
     };

    it('should return the thumbnail for a book from Google Books API', async () => {
        
        const mockResponse = { body: { book: { thumbnail: 'google-thumbnail-url' } } };
        
        (getBookByGoogleId as vi.Mock).mockResolvedValue(mockResponse);

        const result = await getMediaArtSource(book);
        expect(result).toBe('google-thumbnail-url');
    });

    it('should return an empty string if no book data is available', async () => {
        const googleResponse = { body: { book: null } };
        const isbnResponse = { body: { book: null } };
        
        (getBookByGoogleId as vi.Mock).mockResolvedValue(googleResponse);
        (getBookByISBN as vi.Mock).mockResolvedValue(isbnResponse);

        const result = await getMediaArtSource(book);
        expect(result).toBe('');
    });

    const movie : Work = {
        id: 1,
        third_party_id: '12345',
        title: 'Sample Work',
        status: 'in-progress',
        type: 'movie', // Replace with actual WorkType if necessary
        current_progress: 50,
        target_progress: 100,
        version: 1,
        created_at: Date.now(),
        updated_at: new Date(),
     };

    it('should return the thumbnail for a movie', async () => {
        const work = { type: 'movie', third_party_id: 'movieId' };
        const mockResponse = { body: { movie: { thumbnail: 'movie-thumbnail-url' } } };
        
        (getMovie as vi.Mock).mockResolvedValue(mockResponse);

        const result = await getMediaArtSource(movie);
        expect(result).toBe('movie-thumbnail-url');
    });

    it('should return an empty string if no movie data is available', async () => {
        const work = { type: 'movie', third_party_id: 'noMovieId' };
        const mockResponse = { body: { movie: null } };
        
        (getMovie as vi.Mock).mockResolvedValue(mockResponse);

        const result = await getMediaArtSource(movie);
        expect(result).toBe('');
    });

    const show : Work = {
        id: 1,
        third_party_id: '12345',
        title: 'Sample Work',
        status: 'in-progress',
        type: 'show', // Replace with actual WorkType if necessary
        current_progress: 50,
        target_progress: 100,
        version: 1,
        created_at: Date.now(),
        updated_at: new Date(),
     };

    it('should return the thumbnail for a TV show', async () => {
        const work = { type: 'show', third_party_id: 'showId' };
        const mockResponse = { body: { tvshow: { thumbnail: 'show-thumbnail-url' } } };
        
        (getTVShowByIMDbId as vi.Mock).mockResolvedValue(mockResponse);

        const result = await getMediaArtSource(show);
        expect(result).toBe('show-thumbnail-url');
    });

    it('should return an empty string if no show data is available', async () => {
        const work = { type: 'show', third_party_id: 'noShowId' };
        const mockResponse = { body: { tvshow: null } };
        
        (getTVShowByIMDbId as vi.Mock).mockResolvedValue(mockResponse);

        const result = await getMediaArtSource(show);
        expect(result).toBe('');
    });

    const youtube : Work = {
        id: 1,
        third_party_id: 'youtubeId',
        title: 'Sample Work',
        status: 'in-progress',
        type: 'youtube', // Replace with actual WorkType if necessary
        current_progress: 50,
        target_progress: 100,
        version: 1,
        created_at: Date.now(),
        updated_at: new Date(),
     };

    it('should return the YouTube embed URL for a YouTube video', async () => {
        const result = await getMediaArtSource(youtube);
        expect(result).toBe('https://www.youtube.com/embed/youtubeId?si=dourAMMy3-5pBbJr');
    });
});
