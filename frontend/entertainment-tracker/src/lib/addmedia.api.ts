/**
 * 
	router.HandlerFunc(http.MethodGet, "/v1/search/movies", app.requireAuthenticatedUser(app.searchMoviesByTitleHandler))
	router.HandlerFunc(http.MethodGet, "/v1/find/movie", app.requireAuthenticatedUser(app.getMoviesByIMDbHandler))
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.requireAuthenticatedUser(app.getMovieHandler))

	router.HandlerFunc(http.MethodGet, "/v1/search/tvshows", app.requireAuthenticatedUser(app.searchTVShowsByTitleHandler))
	router.HandlerFunc(http.MethodGet, "/v1/find/tvshow", app.requireAuthenticatedUser(app.getTVShowsByIMDbHandler))
	router.HandlerFunc(http.MethodGet, "/v1/tvshows/:id", app.requireAuthenticatedUser(app.getTVShowHandler))

	router.HandlerFunc(http.MethodGet, "/v1/search/books", app.requireAuthenticatedUser(app.searchBooksByTitleHandler))
	router.HandlerFunc(http.MethodGet, "/v1/find/book", app.requireAuthenticatedUser(app.searchBooksByISBNHandler))
	router.HandlerFunc(http.MethodGet, "/v1/books/:id", app.requireAuthenticatedUser(app.getBookHandler))
 */

import api from './api';
import type {
	ApiResponse,
	Book,
	BookSearchResponse,
	Movie,
	MovieSearchResponse,
	TvSearchResponse,
	TvShow
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
