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

export type Work = {
	
};

export type AuthToken = LoginResponse['authentication_token'];
