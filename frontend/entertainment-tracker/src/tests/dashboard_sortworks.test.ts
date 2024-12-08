import { describe, it, expect } from 'vitest';
import { sortByTitle, sortByType, sortByProgress, sortByShared } from '../routes/works/dashboard/dashboard';
import type { WorkPlus } from '../routes/works/dashboard/dashboard';
import type { Work } from '$lib/api.model'

const originalWorks: WorkPlus[] = [
  {
    work: {
      id: 1,
      third_party_id: '123',
      title: 'Work B',
      status: 'in_progress',
      type: 'book',
      current_progress: 50,
      target_progress: 100,
      version: 1,
      created_at: 1672531200,
      updated_at: new Date('2023-01-01'),
    },
    shared: true,
    sharedBy: '',
    thumbnail: '',
  },
  {
    work: {
      id: 2,
      third_party_id: '456',
      title: 'Work A',
      status: 'completed',
      type: 'movie',
      current_progress: 100,
      target_progress: 100,
      version: 2,
      created_at: 1672531200,
      updated_at: new Date('2023-01-03'),
    },
    shared: false,
    sharedBy: '',
    thumbnail: '',
  },
  {
    work: {
      id: 1,
      third_party_id: '123',
      title: 'Work C',
      status: 'in_progress',
      type: 'book',
      current_progress: 20,
      target_progress: 100,
      version: 1,
      created_at: 1672531200,
      updated_at: new Date('2023-01-02'),
    },
    shared: true,
    sharedBy: '',
    thumbnail: '',
  },
];

describe('sortByTitle function', () => {
  it('should reset sortedWorks and log message when sortedByTitle is 2 (not sorted)', () => {
    const sortedByTitle = 2;

    const { sortedWorks, sortedByTitle: returnedSortedByTitle } = sortByTitle(sortedByTitle, originalWorks);

    expect(sortedWorks).toEqual(originalWorks);
    expect(returnedSortedByTitle).toBe(0);
    // You might need to mock console.log for Vitest, depending on your setup.
    // For example, you could use a library like `vitest-mock-console`
    // console.log.mock.calls[0][0].should.equal('not sorted by title');
  });

  it('should sort works by title in ascending order when sortedByTitle is 0', () => {
    const sortedByTitle = 0;

    const { sortedWorks } = sortByTitle(sortedByTitle, originalWorks);

    expect(sortedWorks).toEqual([
      {
        work: {
          id: 2,
          third_party_id: '456',
          title: 'Work A',
          status: 'completed',
          type: 'movie',
          current_progress: 100,
          target_progress: 100,
          version: 2,
          created_at: 1672531200,
          updated_at: new Date('2023-01-03'),
        },
        shared: false,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 1,
          third_party_id: '123',
          title: 'Work B',
          status: 'in_progress',
          type: 'book',
          current_progress: 50,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-01'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 1,
          third_party_id: '123',
          title: 'Work C',
          status: 'in_progress',
          type: 'book',
          current_progress: 20,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-02'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
    ]);
    // console.log.mock.calls[0][0].should.equal('sort by title (asc)');
  });

  it('should sort works by title in descending order when sortedByTitle is 1', () => {
    const sortedByTitle = 1;

    const { sortedWorks } = sortByTitle(sortedByTitle, originalWorks);

    expect(sortedWorks).toEqual([
      {
        work: {
          id: 1,
          third_party_id: '123',
          title: 'Work C',
          status: 'in_progress',
          type: 'book',
          current_progress: 20,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-02'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 1,
          third_party_id: '123',
          title: 'Work B',
          status: 'in_progress',
          type: 'book',
          current_progress: 50,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-01'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 2,
          third_party_id: '456',
          title: 'Work A',
          status: 'completed',
          type: 'movie',
          current_progress: 100,
          target_progress: 100,
          version: 2,
          created_at: 1672531200,
          updated_at: new Date('2023-01-03'),
        },
        shared: false,
        sharedBy: '',
        thumbnail: '',
      },
    ]);
    // console.log.mock.calls[0][0].should.equal('sort by title (desc)');
  });
});

describe('sortByType function', () => {
    it('should reset sortedWorks and log message when sortedByType is 2 (not sorted)', () => {
      const sortedByType = 2;
  
      const { sortedWorks, sortedByType: returnedSortedByType } = sortByType(sortedByType, originalWorks);
  
      expect(sortedWorks).toEqual(originalWorks);
      expect(returnedSortedByType).toBe(0);
      // You might need to mock console.log for Vitest, depending on your setup.
      // For example, you could use a library like `vitest-mock-console`
      // console.log.mock.calls[0][0].should.equal('not sorted by type');
    });
  
    it('should sort works by type in ascending order when sortedByType is 0', () => {
      const sortedByType = 0;
  
      const { sortedWorks } = sortByType(sortedByType, originalWorks);
  
      expect(sortedWorks).toEqual([
        {
          work: {
            id: 1,
            third_party_id: '123',
            title: 'Work B',
            status: 'in_progress',
            type: 'book',
            current_progress: 50,
            target_progress: 100,
            version: 1,
            created_at: 1672531200,
            updated_at: new Date('2023-01-01'),
          },
          shared: true,
          sharedBy: '',
          thumbnail: '',
        },
        {
          work: {
            id: 1,
            third_party_id: '123',
            title: 'Work C',
            status: 'in_progress',
            type: 'book',
            current_progress: 20,
            target_progress: 100,
            version: 1,
            created_at: 1672531200,
            updated_at: new Date('2023-01-02'),
          },
          shared: true,
          sharedBy: '',
          thumbnail: '',
        },
        {
          work: {
            id: 2,
            third_party_id: '456',
            title: 'Work A',
            status: 'completed',
            type: 'movie',
            current_progress: 100,
            target_progress: 100,
            version: 2,
            created_at: 1672531200,
            updated_at: new Date('2023-01-03'),
          },
          shared: false,
          sharedBy: '',
          thumbnail: '',
        },
      ]);
      // console.log.mock.calls[0][0].should.equal('sort by type (asc)');
    });
  
    it('should sort works by type in descending order when sortedByType is 1', () => {
      const sortedByType = 1;
  
      const { sortedWorks } = sortByType(sortedByType, originalWorks);
  
      expect(sortedWorks).toEqual([
        {
          work: {
            id: 2,
            third_party_id: '456',
            title: 'Work A',
            status: 'completed',
            type: 'movie',
            current_progress: 100,
            target_progress: 100,
            version: 2,
            created_at: 1672531200,
            updated_at: new Date('2023-01-03'),
          },
          shared: false,
          sharedBy: '',
          thumbnail: '',
        },
        {
          work: {
            id: 1,
            third_party_id: '123',
            title: 'Work B',
            status: 'in_progress',
            type: 'book',
            current_progress: 50,
            target_progress: 100,
            version: 1,
            created_at: 1672531200,
            updated_at: new Date('2023-01-01'),
          },
          shared: true,
          sharedBy: '',
          thumbnail: '',
        },
        {
          work: {
            id: 1,
            third_party_id: '123',
            title: 'Work C',
            status: 'in_progress',
            type: 'book',
            current_progress: 20,
            target_progress: 100,
            version: 1,
            created_at: 1672531200,
            updated_at: new Date('2023-01-02'),
          },
          shared: true,
          sharedBy: '',
          thumbnail: '',
        },
      ]);
      // console.log.mock.calls[0][0].should.equal('sort by type (desc)');
    });
  });


describe('sortByProgress function', () => {

  it('should reset sortedWorks and log message when sortedByProgress is 2 (not sorted)', () => {
    const sortedByProgress = 2;

    const { sortedWorks, sortedByProgress: returnedSortedByProgress } = sortByProgress(sortedByProgress, originalWorks);

    expect(sortedWorks).toEqual(originalWorks);
    expect(returnedSortedByProgress).toBe(0);
    // You might need to mock console.log for Vitest, depending on your setup.
    // console.log.mock.calls[0][0].should.equal('not sorted by progress');
  });

  it('should sort works by progress in ascending order when sortedByProgress is 1', () => {
    const sortedByProgress = 1;

    const { sortedWorks } = sortByProgress(sortedByProgress, originalWorks);

    expect(sortedWorks).toEqual([
      {
        work: {
          id: 1,
          third_party_id: '123',
          title: 'Work C',
          status: 'in_progress',
          type: 'book',
          current_progress: 20,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-02'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 1,
          third_party_id: '123',
          title: 'Work B',
          status: 'in_progress',
          type: 'book',
          current_progress: 50,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-01'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 2,
          third_party_id: '456',
          title: 'Work A',
          status: 'completed',
          type: 'movie',
          current_progress: 100,
          target_progress: 100,
          version: 2,
          created_at: 1672531200,
          updated_at: new Date('2023-01-03'),
        },
        shared: false,
        sharedBy: '',
        thumbnail: '',
      },
    ]);
    // console.log.mock.calls[0][0].should.equal('sort by progress (desc)');
  });
  it('should sort works by progress in descending order when sortedByProgress is 0', () => {
    const sortedByProgress = 0;

    const { sortedWorks } = sortByProgress(sortedByProgress, originalWorks);

    expect(sortedWorks).toEqual([
      {
        work: {
          id: 2,
          third_party_id: '456',
          title: 'Work A',
          status: 'completed',
          type: 'movie',
          current_progress: 100,
          target_progress: 100,
          version: 2,
          created_at: 1672531200,
          updated_at: new Date('2023-01-03'),
        },
        shared: false,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 1,
          third_party_id: '123',
          title: 'Work B',
          status: 'in_progress',
          type: 'book',
          current_progress: 50,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-01'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 1,
          third_party_id: '123',
          title: 'Work C',
          status: 'in_progress',
          type: 'book',
          current_progress: 20,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-02'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
    ]);
    // console.log.mock.calls[0][0].should.equal('sort by progress (asc)');
  });

});

describe('sortByShared function', () => {
  const originalWorks: WorkPlus[] = [
    {
      work: {
        id: 1,
        third_party_id: '123',
        title: 'Work B',
        status: 'in_progress',
        type: 'book',
        current_progress: 50,
        target_progress: 100,
        version: 1,
        created_at: 1672531200,
        updated_at: new Date('2023-01-01'),
      },
      shared: true,
      sharedBy: '',
      thumbnail: '',
    },
    {
      work: {
        id: 2,
        third_party_id: '456',
        title: 'Work A',
        status: 'completed',
        type: 'movie',
        current_progress: 100,
        target_progress: 100,
        version: 2,
        created_at: 1672531200,
        updated_at: new Date('2023-01-02'),
      },
      shared: false,
      sharedBy: '',
      thumbnail: '',
    },
    {
      work: {
        id: 3,
        third_party_id: '789',
        title: 'Work C',
        status: 'in_progress',
        type: 'book',
        current_progress: 20,
        target_progress: 100,
        version: 1,
        created_at: 1672531200,
        updated_at: new Date('2023-01-01'),
      },
      shared: true,
      sharedBy: '',
      thumbnail: '',
    },
  ];

  it('should reset sortedWorks and log message when sortedByShared is 2 (not sorted)', () => {
    const sortedByShared = 2;

    const { sortedWorks, sortedByShared: returnedSortedByShared } = sortByShared(sortedByShared, originalWorks);

    expect(sortedWorks).toEqual(originalWorks);
    expect(returnedSortedByShared).toBe(0);
    // You might need to mock console.log for Vitest, depending on your setup.
    // console.log.mock.calls[0][0].should.equal('not sorted by shared');
  });

  it('should sort works by shared status in ascending order when sortedByShared is 0', () => {
    const sortedByShared = 0;

    const { sortedWorks } = sortByShared(sortedByShared, originalWorks);

    expect(sortedWorks).toEqual([
      {
        work: {
          id: 1,
          third_party_id: '123',
          title: 'Work B',
          status: 'in_progress',
          type: 'book',
          current_progress: 50,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-01'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 3,
          third_party_id: '789',
          title: 'Work C',
          status: 'in_progress',
          type: 'book',
          current_progress: 20,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-01'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 2,
          third_party_id: '456',
          title: 'Work A',
          status: 'completed',
          type: 'movie',
          current_progress: 100,
          target_progress: 100,
          version: 2,
          created_at: 1672531200,
          updated_at: new Date('2023-01-02'),
        },
        shared: false,
        sharedBy: '',
        thumbnail: '',
      },
    ]);
    // console.log.mock.calls[0][0].should.equal('sort by title (asc)');
  });

  it('should sort works by shared status in descending order when sortedByShared is 1', () => {
    const sortedByShared = 1;

    const { sortedWorks } = sortByShared(sortedByShared, originalWorks);

    expect(sortedWorks).toEqual([
      {
        work: {
          id: 2,
          third_party_id: '456',
          title: 'Work A',
          status: 'completed',
          type: 'movie',
          current_progress: 100,
          target_progress: 100,
          version: 2,
          created_at: 1672531200,
          updated_at: new Date('2023-01-02'),
        },
        shared: false,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 1,
          third_party_id: '123',
          title: 'Work B',
          status: 'in_progress',
          type: 'book',
          current_progress: 50,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-01'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
      {
        work: {
          id: 3,
          third_party_id: '789',
          title: 'Work C',
          status: 'in_progress',
          type: 'book',
          current_progress: 20,
          target_progress: 100,
          version: 1,
          created_at: 1672531200,
          updated_at: new Date('2023-01-01'),
        },
        shared: true,
        sharedBy: '',
        thumbnail: '',
      },
    ]);
    // console.log.mock.calls[0][0].should.equal('sort by title (desc)');
  });
});
