import { test, expect } from '@playwright/test';

test.describe('Add Media', () => {
	test.beforeEach(async ({ page }) => {
		await page.route('localhost:5000/**/login', route => {
			route.fulfill({
				status: 200,
				body: JSON.stringify({
					authentication_token: {
						token: 'mock-token',
						expiry: new Date(Date.now() + 3600 * 1000).toISOString()
					}
				})
			});
		});
		await page.route('**/youtube/*', route => {
			route.fulfill({
				status: 200,
				body: JSON.stringify({
					video: {
						video_id: 'dQw4w9WgXcQ',
						title: 'Rick Astley - Never Gonna Give You Up',
						duration: 212
					}
				})
			});
		});
		await page.route('http://localhost:5000/v1/search/tvshows/', route => {
			route.fulfill({
				status: 200,
				body: JSON.stringify({
					tvshows: [
						{
							id: 'tt0903747',
							title: 'Breaking Bad',
							first_air_date: '2008-01-20',
							vote_average: 9.5,
							thumbnail: 'https://example.com/breakingbad.jpg'
						}
					]
				})
			});
		});
		await page.route('http://localhost:5000/v1/search/movies', route => {
			route.fulfill({
				status: 200,
				body: JSON.stringify({
					movies: [
						{
							id: 'tt1375666',
							title: 'Inception',
							release_date: '2010-07-16',
							vote_average: 8.8,
							thumbnail: 'https://example.com/inception.jpg'
						}
					]
				})
			});
		});
		await page.route('http://localhost:5000/v1/search/books/', (route) => {
			route.fulfill({
				status: 200,
				body: JSON.stringify({
					books: [
						{
							id: '1',
							isbn: '9780141182636',
							title: 'The Great Gatsby',
							author: 'F. Scott Fitzgerald',
							page_count: 180,
							thumbnail: 'https://example.com/gatsby.jpg'
						}
					]
				})
			});
		});
		await page.goto('/user/login');
		await page.fill('#email', 'test@example.com');
		await page.fill('#password', 'password');
		await page.click('button:has-text("Login")');
		await expect(page).toHaveURL('/works/dashboard');
		await page.goto('/works/addmedia');
	});

	test('should add a YouTube video', async ({ page }) => {
		await page.goto('/works/addmedia/youtube');
		await page.fill('#query-field', 'https://www.youtube.com/watch?v=dQw4w9WgXcQ');
		await page.click('button:has-text("Add Video")');
		await expect(page).toHaveURL('/works/dashboard');
	});

	test('should search and add a TV series', async ({ page }) => {
		await page.goto('/works/addmedia/tvseries');
		await page.fill('#query-field', 'Breaking Bad');
		await page.click('button:has-text("Search")');
		await page.click('button[aria-label="Add"]:first-of-type');
		await expect(page).toHaveURL('/works/dashboard');
	});

	test('should search and add a movie', async ({ page }) => {
		await page.goto('/works/addmedia/movies');
		await page.fill('#query-field', 'Inception');
		await page.click('button:has-text("Search")');
		await page.click('button[aria-label="Add"]:first-of-type');
		await expect(page).toHaveURL('/works/dashboard');
	});

	test('should search and add a book by title', async ({ page }) => {
		await page.goto('/works/addmedia/books');
		await page.fill('#query-field', 'The Great Gatsby');
		await page.click('button:has-text("Search")');
		await page.click('button[aria-label="Add"]:first-of-type');
		await expect(page).toHaveURL('/works/dashboard');
	});

	test('should search and add a book by ISBN', async ({ page }) => {
		await page.goto('/works/addmedia/books');
		await page.fill('#query-field', '9780141182636');
		await page.click('button:has-text("Search")');
		await page.click('button[aria-label="Add"]:first-of-type');
		await expect(page).toHaveURL('/works/dashboard');
	});
});
