<script lang="ts">
	import { goto } from '$app/navigation';
	import { register as sendRegisterRequest, login } from '$lib/user.api';

	let email = '';
	let username = '';
	let password = '';

	let currentError: string | null = null;

	function register() {
		const requestBody = {
			email,
			name: username,
			password
		};

		sendRegisterRequest(requestBody).then((response) => {
			if (response.ok) {
				login({ email, password });
				goto('/works/dashboard');
			} else {
				if (response.statusCode === 409) {
					currentError = 'User with given email address or username already exists!';
				} else {
					currentError = 'Invalid registration!';
				}
			}
		});
	}
</script>

<div class="absolute left-4 top-4 z-10 flex items-center gap-1 pl-2">
	<div class="w-fit text-right text-white">
		<img src="/mediamindlogo.png" alt="MediaMind Logo" class="h-6 w-6" />
	</div>
	<div class="Ubuntu-font w-fit pr-8 text-left text-3xl text-white">MediaMind</div>
</div>

<div
	class="bg-background Ubuntu-font absolute inset-0 p-2 flex items-center justify-center font-[Inter] overflow-auto"
>
	<form
		on:submit={register}
		class="Inter-font w-full max-w-[400px] z-20 rounded-2xl bg-white p-6 text-center overflow-auto min-h-0 max-h-full"
	>
		<h1 class="mb-8 text-left text-3xl">Register</h1>

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

		<div class="my-4">
			<label class="block text-left" for="username">Username</label>
			<input
				bind:value={username}
				class="mt-2 w-full p-1"
				id="username"
				type="text"
				placeholder="Username"
			/>
		</div>

		<div class="my-4">
			<label class="block text-left" for="password">Password</label>
			<input
				bind:value={password}
				class="mt-2 w-full p-1"
				id="password"
				type="password"
				placeholder="Password"
			/>
		</div>

		{#if currentError !== null}
			<div class="text-delete mb-4 text-left">{currentError}</div>
		{/if}

		<button class="bg-background rounded-md px-4 py-2 text-sm text-white">Register</button>

		<div class="mt-8 flex justify-between text-lg">
			<div>Already have an account</div>
			<a class="font-bold underline" href="/user/login">Log In</a>
		</div>
	</form>
</div>
