import api from './api';
import type { ApiResponse, LoginRequest, RegisterRequest } from './api.model';

export async function register(request: RegisterRequest): Promise<ApiResponse<void>> {
	return api.post('/users/register', request);
}

export async function login(request: LoginRequest): Promise<ApiResponse<void>> {
	return api.post('/users/login', request);
}

export async function logout(): Promise<ApiResponse<void>> {
	return api.get('/users/logout');
}
