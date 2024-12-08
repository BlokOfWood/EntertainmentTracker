import { test, expect } from '@playwright/test';

test.describe('Register', () => {
	test.beforeEach(async ({ page }) => {
		await page.route('**/users/register', route => {
			const request = route.request();
			const postData = JSON.parse(request.postData() || '{}');

			if (postData.email === 'existing@example.com') {
				route.fulfill({
					status: 409,
					body: JSON.stringify({ message: 'User with given email address or username already exists!' })
				});
			} else if (postData.email === 'invalid@example.com') {
				route.fulfill({
					status: 400,
					body: JSON.stringify({ message: 'Invalid registration!' })
				});
			} else {
				route.fulfill({
					status: 200,
					body: JSON.stringify({ message: 'Registration successful!' })
				});
			}
		});

		await page.route('**/users/login', route => {
			route.fulfill({
				status: 200,
				body: JSON.stringify({ message: 'Login successful!' })
			});
		});
	});

	test('should render the registration form', async ({ page }) => {
		await page.goto('/user/register');
		await expect(page.locator('input[placeholder="Email"]')).toBeVisible();
		await expect(page.locator('input[placeholder="Username"]')).toBeVisible();
		await expect(page.locator('input[placeholder="Password"]')).toBeVisible();
		await expect(page.locator('button:has-text("Register")')).toBeVisible();
	});

	test('should show error message if user already exists', async ({ page }) => {
		await page.goto('/user/register');
		await page.fill('input[placeholder="Email"]', 'existing@example.com');
		await page.fill('input[placeholder="Username"]', 'testuser');
		await page.fill('input[placeholder="Password"]', 'password');
		await page.click('button:has-text("Register")');
		await expect(page.locator('text=User with given email address or username already exists!')).toBeVisible();
	});

	test('should show error message for invalid registration', async ({ page }) => {
		await page.goto('/user/register');
		await page.fill('input[placeholder="Email"]', 'invalid@example.com');
		await page.fill('input[placeholder="Username"]', 'testuser');
		await page.fill('input[placeholder="Password"]', 'password');
		await page.click('button:has-text("Register")');
		await expect(page.locator('text=Invalid registration!')).toBeVisible();
	});

	test('should redirect to dashboard on successful registration', async ({ page }) => {
		await page.goto('/user/register');
		await page.fill('input[placeholder="Email"]', 'test@example.com');
		await page.fill('input[placeholder="Username"]', 'testuser');
		await page.fill('input[placeholder="Password"]', 'password');
		await page.click('button:has-text("Register")');
		await page.waitForURL('/works/dashboard');
	});
});
