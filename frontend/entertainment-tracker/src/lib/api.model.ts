export type ApiOptions = Partial<{
	skipAuth: boolean;
	responseFormat: 'json' | 'none';
}>;

export type ApiResponse<T> = {
	statusCode: number;
	ok: boolean;
	body: T;
};

export type RegisterRequest = {
	email: string;
	name: string;
	password: string;
};

export type LoginRequest = {
	email: string;
	password: string;
};

export type LoginResponse = {
	authentication_token: {
		token: string;
		expiry: Date;
	};
};

export type UpdateWorkRequest = {
	title: string;
	type: string;
	status: string;
	current_progress: number;
	target_progress: number;
};

export type CreateWorkRequest = {
	third_party_id: string;
	title: string;
	type: string;
	status: string;
	current_progress: number;
	target_progress: number;
};

export type Work = {
	id: number;
	third_party_id: string;
	title: string;
	status: string;
	type: string;
	current_progress: number;
	target_progress: number;
	version: number;
	created_at: number;
	updated_at: Date;
};

export type AuthToken = LoginResponse['authentication_token'];
