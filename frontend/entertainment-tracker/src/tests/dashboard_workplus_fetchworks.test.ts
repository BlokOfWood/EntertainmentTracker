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