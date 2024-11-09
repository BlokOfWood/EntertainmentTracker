export type ApiResponse<T> = {
    statusCode: number;
    ok: boolean;
    body: T;
}

export type RegisterRequest = {
    email: string;
    username: string;
    password: string;
}

export type LoginRequest = {
    email: string;
    password: string;
}