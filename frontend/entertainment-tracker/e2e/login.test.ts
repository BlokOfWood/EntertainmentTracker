import { expect, test } from '@playwright/test';

test.describe('Login Page', () => {
	test.beforeEach(async ({ page }) => {
		await page.route('**/users/login', route => {
			const request = route.request();
			const postData = JSON.parse(request.postData() || '{}');

			if (postData.email === 'invalid@example.com') {
				route.fulfill({
					status: 401,
					body: JSON.stringify({ message: 'Invalid login!' })
				});
			} else if (postData.email === 'test1@test.com' && postData.password === 'testtest') {
				const response = {
					authentication_token: {
						token: 'mocked-jwt-token',
						expiry: new Date(Date.now() + 3600 * 1000) // 1 hour expiry
					}
				};
				route.fulfill({
					status: 200,
					body: JSON.stringify(response)
				});
			} else {
				route.fulfill({
					status: 401,
					body: JSON.stringify({ message: 'Invalid login!' })
				});
			}
		});

		await page.goto('/user/login');
	});

	test('should display the login page', async ({ page }) => {
		await expect(page).toHaveURL('/user/login');
		await expect(page.locator('h1')).toHaveText('Login');
	});

	test('should show error message with invalid credentials', async ({ page }) => {
		await page.fill('input#email', 'invalid@example.com');
		await page.fill('input#password', 'invalidpassword');
		await page.click('button:has-text("Login")');
		await expect(page.locator('text=Invalid login!')).toBeVisible({ timeout: 10000 });
	});

	test('should navigate to the registration page', async ({ page }) => {
		await page.click('a:has-text("Register")');
		await expect(page).toHaveURL('/user/register');
	});

	test('should login successfully with test credentials', async ({ page }) => {
		await page.fill('input#email', 'test1@test.com');
		await page.fill('input#password', 'testtest');
		await page.click('button:has-text("Login")');
		await expect(page).toHaveURL('/works/dashboard');
	});
});
