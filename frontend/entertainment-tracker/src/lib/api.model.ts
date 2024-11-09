export type ApiResponse<T> = {
    statusCode: number;
    body: T;
}

export type RegisterRequest = {
    email: string;
    password: string;
    name: string;
}

export type LoginRequest = {
    email: string;
    password: string;
}