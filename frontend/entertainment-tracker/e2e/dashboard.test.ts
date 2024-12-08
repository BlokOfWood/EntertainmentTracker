import { test, expect } from '@playwright/test';

test.describe('Dashboard Page', () => {
    test.beforeEach(async ({ page }) => {
        // Mock API responses
        await page.route('**/mediaentries', route => {
            route.fulfill({
                status: 200,
                body: JSON.stringify({
                    mediaEntries: [
                        {
                            id: 1,
                            third_party_id: '123',
                            title: 'Test Book',
                            type: 'book',
                            status: 'watching',
                            current_progress: 50,
                            target_progress: 100,
                            version: 1,
                            created_at: '2023-01-01T00:00:00Z',
                            updated_at: '2023-01-01T00:00:00Z'
                        },
                        {
                            id: 2,
                            third_party_id: '456',
                            title: 'Test Movie',
                            type: 'movie',
                            status: 'completed',
                            current_progress: 120,
                            target_progress: 120,
                            version: 1,
                            created_at: '2023-01-01T00:00:00Z',
                            updated_at: '2023-01-01T00:00:00Z'
                        }
                    ]
                })
            });
        });

        await page.route('**/shared', route => {
            route.fulfill({
                status: 200,
                body: JSON.stringify({
                    sharedEntries: []
                })
            });
        });

        await page.route('localhost:5000/user/login', route => {
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

        // Set up authenticated state
        await page.goto('/user/login');
        await page.fill('input[type="email"]', 'test@example.com');
        await page.fill('input[type="password"]', 'password');
        await page.click('button:has-text("Login")');
        await page.waitForURL('/works/dashboard');
    });

    test('should display media entries', async ({ page }) => {
        await page.waitForSelector('.grid-cols-6 button.Ubuntu-font:has-text("Test Book")');
        const titles = await page.locator('.grid-cols-6 button.Ubuntu-font').allTextContents();
        expect(titles).toContain('Test Book');
        expect(titles).toContain('Test Movie');
    });

    test('should open edit modal', async ({ page }) => {
        await page.waitForSelector('.grid-cols-6 button.Ubuntu-font:has-text("Test Book")');
        await page.click('.edit-button');
        await expect(page.locator('#edit-popup')).toBeVisible();
    });

    test('should open share modal', async ({ page }) => {
        await page.waitForSelector('.grid-cols-6 button.Ubuntu-font:has-text("Test Book")');
        await page.click('.share-button');
        await expect(page.locator('#share-popup')).toBeVisible();
    });

    test('should open delete modal', async ({ page }) => {
        await page.waitForSelector('.grid-cols-6 button.Ubuntu-font:has-text("Test Book")');
        await page.click('.delete-button');
        await expect(page.locator('#delete-popup')).toBeVisible();
    });
});
