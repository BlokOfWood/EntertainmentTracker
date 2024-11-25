import api from './api';
import type { ApiResponse, LoginRequest, LoginResponse, RegisterRequest } from './api.model';

export async function register(request: RegisterRequest): ApiResponse<void> {
	return api.post('/users/register', request, { skipAuth: true });
}

export async function login(request: LoginRequest): ApiResponse<LoginResponse> {
	return api.post('/users/login', request, { skipAuth: true });
}

export async function logout(): ApiResponse<void> {
	return api.get('/users/logout', { responseFormat: 'none' });
}
