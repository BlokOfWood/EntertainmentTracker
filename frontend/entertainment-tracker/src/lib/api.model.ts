export type WorkType = 'movie' | 'book' | 'show' | 'youtube';

export type ApiOptions = Partial<{
	skipAuth: boolean;
	responseFormat: 'json' | 'none';
}>;

export type ApiResponse<T> = Promise<{
	statusCode: number;
	ok: boolean;
	body: T;
}>;

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

export type User = {
	id: number;
	email: string;
	name: string;
	created_at: Date;
};

export type UpdateWorkRequest = {
	title: string;
	type: WorkType;
	status: string;
	current_progress: number;
	target_progress: number;
};

export type CreateWorkRequest = {
	third_party_id: string;
	title: string;
	type: WorkType;
	status: string;
	current_progress: number;
	target_progress: number;
};

export type ShareWorkRequest = {
	/**
	 * The ID of the work to share.
	 */
	media_entry: number;
	/**
	 * The email of the user to share the work with.	
	 */
	share_with: string;
};

export type SharedWork = {
	id: number;
	shared_by: number;
	shared_with: number;
	media_entry: Work;
	created_at: Date;
};

export type Work = {
	id: number;
	third_party_id: string;
	title: string;
	status: string;
	type: WorkType;
	current_progress: number;
	target_progress: number;
	version: number;
	created_at: number;
	updated_at: Date;
};

export type MovieSearchResponse = {
	id: number;
	title: string;
	release_date: string;
	popularity: number;
	vote_average: number;
	thumbnail: string;
};

export type Movie = {
	id: number;
	title: string;
	release_date: string;
	overview: string;
	popularity: number;
	thumbnail: string;
	genres: string[];
	runtime: number;
};

export type TvSearchResponse = {
	id: number;
	title: string;
	first_air_date: string;
	popularity: number;
	vote_average: number;
	thumbnail: string;
};

export type TvShow = {
	id: number;
	title: string;
	first_air_date: string;
	overview: string;
	popularity: number;
	thumbnail: string;
	genres: string[];
	number_of_seasons: number;
	number_of_episodes: number;
};

export type BookSearchResponse = {
	id: string;
	isbn: string;
	title: string;
	author: string;
	page_count: number;
	thumbnail: string;
};

export type Book = {
	id: string;
	isbn: string;
	title: string;
	author: string;
	description: string;
	page_count: number;
	thumbnail: string;
	categories: string[];
	published_date: string;
	publisher: string;
	language: string;
}

export type YoutubeVideoReponse = {
	video_id: string;
	video_url: string;
	title: string;
	channel: string;
	thumbnail: string;
	duration: number;
	published: string;
};

export type AuthToken = LoginResponse['authentication_token'];
