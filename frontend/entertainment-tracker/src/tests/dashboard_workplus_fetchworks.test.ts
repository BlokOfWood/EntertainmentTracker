import { describe, it, expect, vi, afterEach } from 'vitest';
import type { WorkPlus } from '../routes/works/dashboard/dashboard';
import { fetchWorks, getMediaArtSource } from '../routes/works/dashboard/dashboard';
import type { Work, SharedWork } from '$lib/api.model';
import { getWorks, getSharedWorks } from '$lib/works.api';
import { getBookByGoogleId, getBookByISBN, getMovie, getTVShowByIMDbId } from '$lib/addmedia.api'

describe('WorkPlus Type', () => {
  it('should create a valid WorkPlus object', () => {
      const validWork: Work = {
          id: 1,
          third_party_id: 'abc123',
          title: 'Sample Work',
          status: 'in_progress',
          type: 'book', // Updated to use one of the specified WorkType values
          current_progress: 50,
          target_progress: 100,
          version: 1,
          created_at: Date.now(),
          updated_at: new Date(),
      };

      const workPlus: WorkPlus = {
          work: validWork,
          shared: true,
          sharedBy: 'user123',
          thumbnail: 'http://example.com/thumbnail.jpg',
      };

      expect(workPlus).toHaveProperty('work');
      expect(workPlus.work).toEqual(validWork);
      expect(workPlus).toHaveProperty('shared', true);
      expect(workPlus).toHaveProperty('sharedBy', 'user123');
      expect(workPlus).toHaveProperty('thumbnail', 'http://example.com/thumbnail.jpg');
  });
});

vi.mock('$lib/works.api', () => ({
    getWorks: vi.fn(),
    getSharedWorks: vi.fn(),
}));

vi.mock('$lib/addmedia.api', () => ({
    getBookByISBN: vi.fn().mockResolvedValue({
        body: {
            book: {
                thumbnail: 'isbn_thumbnail_url',
            },
        },
    }),
    getTVShowByIMDbId: vi.fn().mockResolvedValue({
        body: {
            tvshow: {
                thumbnail: 'tvshow_thumbnail_url',
            },
        },
    }),
    getMovie: vi.fn().mockResolvedValue({
        body: {
            movie: {
                thumbnail: 'movie_thumbnail_url',
            },
        },
    }),
    getBookByGoogleId: vi.fn().mockResolvedValue({
        body: {
            book: {
                thumbnail: 'google_thumbnail_url',
            },
        },
    }),
}));

// Mock the media art source function
vi.mock('../routes/works/dashboard/dashboard', async (importOriginal) => {
    const actual = await importOriginal() as { fetchWorks: typeof fetchWorks; getMediaArtSource: typeof getMediaArtSource };
    return {
        ...actual,
        getMediaArtSource: vi.fn().mockImplementation(async (work) => {
            if (work.type === 'book') {
                return 'google_thumbnail_url'; // or 'isbn_thumbnail_url' based on your logic
            } else if (work.type === 'movie') {
                return 'movie_thumbnail_url';
            } else if (work.type === 'show') {
                return 'tvshow_thumbnail_url';
            } else if (work.type === 'youtube') {
                return 'youtube_thumbnail_url';
            }
            return '';
        }),
    };
});

describe('fetchWorks', () => {
    it('should fetch and return works correctly', async () => {
        // Mock data for not shared works
        const notSharedWorks = [
            {
                id: 1,
                third_party_id: '123',
                title: 'Not Shared Book',
                status: 'reading',
                type: 'book',
                current_progress: 50,
                target_progress: 100,
                version: 1,
                created_at: Date.now(),
                updated_at: new Date(),
            },
            {
                id: 2,
                third_party_id: '123',
                title: 'Not Shared TVShow',
                status: 'completed',
                type: 'show',
                current_progress: 50,
                target_progress: 100,
                version: 1,
                created_at: Date.now(),
                updated_at: new Date(),
            },
            {
                id: 2,
                third_party_id: '123',
                title: 'Not Shared YouTubeVideo',
                status: 'completed',
                type: 'youtube',
                current_progress: 50,
                target_progress: 100,
                version: 1,
                created_at: Date.now(),
                updated_at: new Date(),
            },
        ];

        // Mock data for shared works
        const sharedWorks = [
            {
                media_entry: {
                    id: 2,
                    third_party_id: '456',
                    title: 'Shared Movie',
                    status: 'completed',
                    type: 'movie',
                    current_progress: 1,
                    target_progress: 1,
                    version: 1,
                    created_at: Date.now(),
                    updated_at: new Date(),
                },
            },
        ];

        // Mock the API responses
        (getWorks as vi.Mock).mockResolvedValue({ ok: true, body: { mediaEntries: notSharedWorks } });
        (getSharedWorks as vi.Mock).mockResolvedValue({ ok: true, body: { sharedEntries: sharedWorks } });

        // Mock the media art source function
        (getMediaArtSource as vi.Mock).mockResolvedValue('thumbnail_url');

        // Call the fetchWorks function
        const works = await fetchWorks();

        // Assertions
        expect(works).toHaveLength(4);
        expect(works[0]).toEqual({
            work: notSharedWorks[0],
            shared: false,
            sharedBy: '',
            thumbnail: 'google_thumbnail_url',
        });
        expect(works[1]).toEqual({
            work: notSharedWorks[1],
            shared: false,
            sharedBy: '',
            thumbnail: 'tvshow_thumbnail_url',
        });
        expect(works[2]).toEqual({
            work: notSharedWorks[2],
            shared: false,
            sharedBy: '',
            thumbnail: 'https://www.youtube.com/embed/123?si=dourAMMy3-5pBbJr',
        });
        expect(works[3]).toEqual({
            work: sharedWorks[0].media_entry,
            shared: true,
            sharedBy: '',
            thumbnail: 'movie_thumbnail_url',
        });
    });
});