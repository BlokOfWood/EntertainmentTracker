<script lang="ts">
	import { goto } from '$app/navigation';
	import api from '$lib/api';
	import { login as sendLoginRequest } from '$lib/user.api';

	let email = '';
	let password = '';

	async function login() {
		const requestBody = {
			email,
			password
		};

		const response = await sendLoginRequest(requestBody);
		if (response.ok) {
			console.log('Login successful');
			api.setToken(response.body.authentication_token);
			await goto('/works/dashboard');
		} else {
			console.log('Login failed');
			alert('Invalid email or password');
		}
	}
</script>

<div class="absolute left-4 top-4 z-10 flex items-center gap-1 pl-2">
	<div class="w-fit text-right text-white">
		<img src="/mediamindlogo.png" alt="MediaMind Logo" class="h-6 w-6" />
	</div>
	<div class="Ubuntu-font w-fit pr-8 text-left text-3xl text-white">MediaMind</div>
</div>

<div
	class="bg-background Ubuntu-font absolute inset-0 flex items-center justify-center font-[Inter]"
>
	<form
		on:submit={login}
		class="Inter-font w-full max-w-[400px] rounded-2xl bg-white p-6 text-center"
	>
		<h1 class="mb-8 text-left text-3xl">Login</h1>

		<div>
			<label class="block text-left" for="email">Email</label>
			<input
				bind:value={email}
				class="mt-2 w-full p-1"
				id="email"
				type="email"
				placeholder="Email"
			/>
		</div>

		<div class="mb-8 mt-4">
			<label class="block text-left" for="password">Password</label>
			<input
				bind:value={password}
				class="mt-2 w-full p-1"
				id="password"
				type="password"
				placeholder="Password"
			/>
		</div>

		<button class="bg-background rounded-md px-4 py-2 text-sm text-white">Login</button>
	</form>
</div>
