import type { ApiOptions, ApiResponse, AuthToken } from './api.model';

class Api {
	private apiBaseAddress = 'http://localhost:5000/v1';
	private token: AuthToken | null;

	public async get<T>(endpoint: string, options?: ApiOptions): Promise<ApiResponse<T>> {
		const headers = new Headers();
		if (this.token !== null && !options?.skipAuth) {
			headers.append('Authorization', `Bearer ${this.token.token}`);
		}

		const response = await fetch(this.apiBaseAddress + endpoint, {
			headers
		});

		return this.processResponse(response, options);
	}

	public async post<T>(
		endpoint: string,
		data: object,
		options?: ApiOptions
	): Promise<ApiResponse<T>> {
		const headers = new Headers();
		headers.append('Content-Type', 'application/json');
		if (this.token !== null && !options?.skipAuth) {
			headers.append('Authorization', `Bearer ${this.token.token}`);
		}

		const response = await fetch(this.apiBaseAddress + endpoint, {
			method: 'POST',
			headers,
			body: JSON.stringify(data)
		});

		return this.processResponse(response, options);
	}

	public setToken(token: AuthToken | null) {
		this.token = token;

		if (token === null) {
			console.log('removing token');
			localStorage.removeItem('token');
			return;
		}
		localStorage.setItem('token', JSON.stringify(token));
	}

	public get validToken(): boolean {
        console.log(this.token);

		return this.token !== null && this.token.expiry > new Date();
	}

	constructor() {
		this.token = null;
		if (typeof window === 'undefined') return;

		const token = localStorage.getItem('token');
		if (token !== null) {
			try {
				console.log(token);
				this.token = JSON.parse(token);
				this.token!.expiry = new Date(this.token!.expiry);
			} catch (e) {
				console.error('Failed to parse token', e);
				this.token = null;
			}
		}
	}

	private async processResponse<T>(
		response: Response,
		options: ApiOptions | undefined
	): Promise<ApiResponse<T>> {
		let responseBody = null;

		switch (options?.responseFormat) {
			case 'none':
				break;
			default:
				responseBody = await response.json();
				break;
		}

		return {
			statusCode: response.status,
			ok: response.ok,
			body: responseBody as T
		};
	}
}

export const api = new Api();
export default api;
