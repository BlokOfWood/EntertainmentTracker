import { expect, test } from '@playwright/test';

test.describe('Login Page', () => {
	test.beforeEach(async ({ page }) => {
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
