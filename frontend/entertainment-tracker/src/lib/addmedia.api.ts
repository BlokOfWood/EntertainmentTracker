import api from './api';
import type {
	ApiResponse,
	Book,
	BookSearchResponse,
	Movie,
	MovieSearchResponse,
	TvSearchResponse,
	TvShow,
	YoutubeVideoReponse
} from './api.model';

export async function searchMoviesByTitle(
	title: string
): ApiResponse<{ movies: MovieSearchResponse[] }> {
	return api.get(`/search/movies?q=${title}`);
}

export async function getMovieByIMDb(imdbID: string): ApiResponse<{ movie: Movie }> {
	return api.get(`/find/movie?id=${imdbID}`);
}

export async function getMovie(id: number): ApiResponse<{ movie: Movie }> {
	return api.get(`/movies/${id}`);
}

export async function searchTVShowsByTitle(
	title: string
): ApiResponse<{ tvshows: TvSearchResponse[] }> {
	return api.get(`/search/tvshows?q=${title}`);
}

export async function getTVShowByIMDb(imdbID: string): ApiResponse<{ tvshow: TvShow }> {
	return api.get(`/find/tvshow?id=${imdbID}`);
}

export async function getTVShowByIMDbId(id: number): ApiResponse<{ tvshow: TvShow }> {
	return api.get(`/tvshows/${id}`);
}

export async function searchBooksByTitle(
	title: string
): ApiResponse<{ books: BookSearchResponse[] }> {
	return api.get(`/search/books?q=${title}`);
}

export async function getBookByISBN(isbn: string): ApiResponse<{ book: Book }> {
	return api.get(`/find/book?id=${isbn}`);
}

export async function getBookByGoogleId(id: string): ApiResponse<{ book: Book }> {
	return api.get(`/books/${id}`);
}

export async function getYoutubeVideo(id: string): ApiResponse<{video: YoutubeVideoReponse}>{
	return api.get(`/youtube?id=${id}`);
}