import type { ApiOptions, ApiResponse, AuthToken } from './api.model';

class Api {
	private apiBaseAddress = 'http://localhost:5000/v1';
	private token: AuthToken | null;

	public async get<T>(endpoint: string, options?: ApiOptions): ApiResponse<T> {
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
	): ApiResponse<T> {
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

	public async patch<T>(
		endpoint: string,
		data: object,
		options?: ApiOptions
	): ApiResponse<T> {
		const headers = new Headers();
		headers.append('Content-Type', 'application/json');
		if (this.token !== null && !options?.skipAuth) {
			headers.append('Authorization', `Bearer ${this.token.token}`);
		}

		const response = await fetch(this.apiBaseAddress + endpoint, {
			method: 'PATCH',
			headers,
			body: JSON.stringify(data)
		});

		return this.processResponse(response, options);
	}

	public async delete<T>(endpoint: string, options?: ApiOptions):ApiResponse<T> {
		const headers = new Headers();
		if (this.token !== null && !options?.skipAuth) {
			headers.append('Authorization', `Bearer ${this.token.token}`);
		}

		const response = await fetch(this.apiBaseAddress + endpoint, {
			method: 'DELETE',
			headers
		});

		return this.processResponse(response, options);
	}

	public setToken(token: AuthToken | null) {
		if (token === null) {
			this.token = null;
			localStorage.removeItem('token');
			return;
		}

		this.token = {
			token: token!.token,
			expiry: new Date(token.expiry)
		};

		localStorage.setItem('token', JSON.stringify(this.token));
	}

	public get validToken(): boolean {
		return this.token !== null && this.token.expiry > new Date();
	}

	constructor() {
		this.token = null;
		if (typeof window === 'undefined') return;

		const token = localStorage.getItem('token');
		if (token !== null) {
			try {
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
	): ApiResponse<T> {
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
