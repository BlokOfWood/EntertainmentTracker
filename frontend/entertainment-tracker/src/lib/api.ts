import type { ApiResponse } from "./api.model";

class Api {
	private apiBaseAddress = 'http://localhost:3000/v1';

    public async get<T>(endpoint: string): Promise<ApiResponse<T>> {
        const response = await fetch(this.apiBaseAddress + endpoint);
        const responseBody = await response.json();
        return {
            statusCode: response.status,
            ok: response.ok,
            body: responseBody as T,
        };
    }

    public async post<T>(endpoint: string, data: object): Promise<ApiResponse<T>> {
        const response = await fetch(this.apiBaseAddress + endpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });
        const responseBody = await response.json();
        return {
            statusCode: response.status,
            ok: response.ok,
            body: responseBody as T,
        };
    }
}

const api = new Api();

export default api;